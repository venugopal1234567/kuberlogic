/*
 * CloudLinux Software Inc 2019-2021 All Rights Reserved
 */

package controllers

import (
	"context"
	"k8s.io/client-go/rest"
	"k8s.io/utils/pointer"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	kuberlogiccomv1alpha1 "github.com/kuberlogic/kuberlogic/modules/dynamic-operator/api/v1alpha1"
	cfg2 "github.com/kuberlogic/kuberlogic/modules/dynamic-operator/cfg"
	"github.com/kuberlogic/kuberlogic/modules/dynamic-operator/plugin/commons"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	velero "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	k8sClient     client.Client
	testEnv       *envtest.Environment
	ctx           context.Context
	cancel        context.CancelFunc
	pluginClients []*plugin.Client

	kuberlogicNamespace = os.Getenv("NAMESPACE")
)

var pluginMap = map[string]plugin.Plugin{
	"docker-compose": &commons.Plugin{},
}

type fakeExecutor struct {
	err error
}

func (f *fakeExecutor) Stream(o remotecommand.StreamOptions) error {
	return f.err
}

func TestControllerAPIs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	ctx, cancel = context.WithCancel(context.TODO())

	By("bootstrapping test environment")

	expectNoError(kuberlogiccomv1alpha1.AddToScheme(scheme.Scheme))
	expectNoError(velero.AddToScheme(scheme.Scheme))

	if useExistingCluster() {
		expectNoError(os.Unsetenv("KUBEBUILDER_ASSETS"))
		testEnv = &envtest.Environment{
			UseExistingCluster: pointer.Bool(useExistingCluster()),
		}
		cfg, err := testEnv.Start()
		expectNoError(err)
		Expect(cfg).NotTo(BeNil())

		k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
		expectNoError(err)
		Expect(k8sClient).NotTo(BeNil())
	} else {
		testEnv = &envtest.Environment{
			CRDDirectoryPaths: []string{
				filepath.Join("..", "config", "crd", "bases"),
				filepath.Join("..", "config", "velero", "crd"),
			},
		}

		cfg, err := testEnv.Start()
		expectNoError(err)
		Expect(cfg).NotTo(BeNil())

		k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
		expectNoError(err)
		Expect(k8sClient).NotTo(BeNil())

		createNamespace(k8sClient, kuberlogicNamespace)

		config, err := cfg2.NewConfig()
		expectNoError(err)

		k8sManager, err := ctrl.NewManager(cfg, ctrl.Options{
			Scheme: scheme.Scheme,
		})
		expectNoError(err)

		logger := createLogger()
		pluginInstances := initializePlugins(config, logger)

		setupReconciler(k8sManager, cfg, config, pluginInstances)
		if config.Backups.Enabled {
			setupBackupControllers(k8sClient, k8sManager, config)
		}

		startManager(k8sManager)
	}
}, 60)

var _ = AfterSuite(func() {
	cancel()
	By("tearing down the test environment")
	for _, cl := range pluginClients {
		cl.Kill()
	}
	expectNoError(testEnv.Stop())
})

func useExistingCluster() bool {
	return os.Getenv("USE_EXISTING_CLUSTER") == "true"
}

func expectNoError(err error) {
	Expect(err).NotTo(HaveOccurred())
}

func createNamespace(client client.Client, name string) {
	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: name}}
	expectNoError(client.Create(ctx, ns))
}

func createLogger() hclog.Logger {
	return hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})
}

func initializePlugins(config *cfg2.Config, logger hclog.Logger) map[string]commons.PluginService {
	pluginInstances := make(map[string]commons.PluginService)
	for _, item := range config.Plugins {
		pluginClient := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: commons.HandshakeConfig,
			Plugins:         pluginMap,
			Cmd:             exec.Command(item.Path),
			Logger:          logger,
		})
		pluginClients = append(pluginClients, pluginClient)

		rpcClient, err := pluginClient.Client()
		expectNoError(err)
		raw, err := rpcClient.Dispense(item.Name)
		expectNoError(err)
		pluginInstances[item.Name] = raw.(commons.PluginService)
	}
	return pluginInstances
}

func setupReconciler(mgr ctrl.Manager, cfg *rest.Config, config *cfg2.Config, pluginInstances map[string]commons.PluginService) {
	dependantObjects := collectDependantObjects(pluginInstances)
	expectNoError((&KuberLogicServiceReconciler{
		Client:     mgr.GetClient(),
		Scheme:     mgr.GetScheme(),
		RESTConfig: cfg, // No change, as cfg is already *rest.Config
		Plugins:    pluginInstances,
		Cfg:        config,
	}).SetupWithManager(mgr, dependantObjects...))
}

func collectDependantObjects(pluginInstances map[string]commons.PluginService) []client.Object {
	var dependantObjects []client.Object
	for _, instance := range pluginInstances {
		for _, obj := range instance.Types().Objects {
			dependantObjects = append(dependantObjects, obj) // Convert to client.Object
		}
	}
	return dependantObjects
}

func setupBackupControllers(client client.Client, mgr ctrl.Manager, config *cfg2.Config) {
	veleroNs := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "velero"}}
	expectNoError(client.Create(ctx, veleroNs))

	backupStorage := &velero.BackupStorageLocation{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "default",
			Namespace: veleroNs.Name,
		},
		Spec: velero.BackupStorageLocationSpec{
			Default:    true,
			Provider:   "aws",
			AccessMode: "ReadWrite",
			StorageType: velero.StorageType{
				ObjectStorage: &velero.ObjectStorageLocation{
					Bucket: "test",
				},
			},
		},
		Status: velero.BackupStorageLocationStatus{
			Phase:              velero.BackupStorageLocationPhaseAvailable,
			LastValidationTime: &metav1.Time{Time: time.Now()},
		},
	}
	expectNoError(client.Create(ctx, backupStorage))

	expectNoError((&KuberlogicServiceBackupReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
		Cfg:    config,
	}).SetupWithManager(mgr))

	expectNoError((&KuberlogicServiceRestoreReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
		Cfg:    config,
	}).SetupWithManager(mgr))

	expectNoError((&KuberlogicServiceBackupScheduleReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
		Cfg:    config,
	}).SetupWithManager(mgr))
}

func startManager(mgr ctrl.Manager) {
	go func() {
		defer GinkgoRecover()
		expectNoError(mgr.Start(ctx))
	}()
}
