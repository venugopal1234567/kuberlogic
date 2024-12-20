module github.com/kuberlogic/kuberlogic

go 1.16

require (
	github.com/AlecAivazis/survey/v2 v2.3.5
	github.com/chargebee/chargebee-go v2.12.0+incompatible
	github.com/compose-spec/compose-go v1.2.4
	github.com/dustinkirkland/golang-petname v0.0.0-20191129215211-8e5a1ed0cff0
	github.com/getsentry/sentry-go v0.13.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-chi/chi v1.5.4
	github.com/go-chi/cors v1.2.0
	github.com/go-logr/logr v1.2.4
	github.com/go-logr/zapr v1.2.4
	github.com/go-openapi/errors v0.20.2
	github.com/go-openapi/loads v0.21.1
	github.com/go-openapi/runtime v0.23.3
	github.com/go-openapi/spec v0.20.4
	github.com/go-openapi/strfmt v0.21.2
	github.com/go-openapi/swag v0.22.3
	github.com/go-openapi/validate v0.21.0
	github.com/go-test/deep v1.0.8
	github.com/google/uuid v1.3.1
	github.com/hashicorp/go-hclog v1.5.0
	github.com/hashicorp/go-plugin v1.4.3
	github.com/jessevdk/go-flags v1.5.0
	github.com/jetstack/cert-manager v1.6.3
	github.com/kuberlogic/zapsentry v1.6.2
	github.com/novln/docker-parser v1.0.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/ginkgo/v2 v2.9.5
	github.com/onsi/gomega v1.27.7
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron v1.1.0
	github.com/spf13/cobra v1.7.0
	github.com/spf13/viper v1.17.0
	github.com/vmware-tanzu/velero v1.8.1
	github.com/vrischmann/envconfig v1.3.0
	github.com/xeipuuv/gojsonschema v1.2.0
	github.com/zalando/postgres-operator v1.7.1
	go.uber.org/zap v1.24.0
	golang.org/x/net v0.19.0
	k8s.io/api v0.28.6
	k8s.io/apiextensions-apiserver v0.28.6
	k8s.io/apimachinery v0.28.6
	k8s.io/client-go v0.28.6
	k8s.io/utils v0.0.0-20230406110748-d93618cff8a2
	sigs.k8s.io/controller-runtime v0.15.3
)
