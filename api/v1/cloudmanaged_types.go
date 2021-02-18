package v1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CloudManagedSpec struct {
	// Type of the cluster
	// +kubebuilder:validation:Enum=postgresql;mysql;redis
	Type string `json:"type"`
	// Amount of replicas
	// +kubebuilder:validation:Maximum=5
	Replicas int32 `json:"replicas"`
	// Resources (requests/limits)
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
	// Volume size
	VolumeSize string `json:"volumeSize,omitempty"`
	// 2 or 3 digits: 5.7 or 5.7.31
	// +kubebuilder:validation:Pattern=^\d+\.\d+[\.\d+]*$
	Version           string            `json:"version,omitempty"`
	AdvancedConf      map[string]string `json:"advancedConf,omitempty"`
	MaintenanceWindow `json:"maintenanceWindow,omitempty"`
}

type MaintenanceWindow struct {
	// start hour, UTC zone is assumed
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=23
	// +kubebuilder:validation:Type=integer
	StartHour int `json:"start,omitempty"`
	// day of the week
	// +kubebuilder:validation:Enum=Monday;Tuesday;Wednesday;Thursday;Friday;Saturday;Sunday
	Weekday string `json:"weekday,omitempty"`
	// window duration in hours
	// +kubebuilder:validation:Type=integer
	// +kubebuilder:validation:Default=4
	DurationHours int `json:"duration,omitempty"`
}

// CloudManagedStatus defines the observed state of CloudManaged
type CloudManagedStatus struct {
	Status string `json:"status"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.status",description="The cluster status"
// +kubebuilder:printcolumn:name="Type",type=string,JSONPath=`.spec.type`,description="The cluster type"
// +kubebuilder:printcolumn:name="Replicas",type=integer,JSONPath=`.spec.replicas`,description="The number of desired replicas"
// +kubebuilder:printcolumn:name="Volume",type=string,JSONPath=`.spec.volumeSize`,description="Volume size for the cluster"
// +kubebuilder:printcolumn:name="CPU Request",type=string,JSONPath=`.spec.resources.requests.cpu`,description="CPU request"
// +kubebuilder:printcolumn:name="Memory Request",type=string,JSONPath=`.spec.resources.requests.memory`,description="Memory request"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:shortName=cl
type CloudManaged struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudManagedSpec   `json:"spec,omitempty"`
	Status CloudManagedStatus `json:"status,omitempty"`
}

// CloudManagedList contains a list of CloudManaged
// +kubebuilder:object:root=true
type CloudManagedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudManaged `json:"items"`
}

func (cm *CloudManaged) IsEqual(newStatus string) bool {
	return cm.Status.Status == newStatus
}

func (cm *CloudManaged) SetStatus(newStatus string) {
	cm.Status.Status = newStatus
}

func (cm *CloudManaged) GetStatus() string {
	return cm.Status.Status
}

// TODO: Figure out workaround in https://github.com/kubernetes-sigs/kubebuilder/issues/1501, not it's a blocker
// for implementation default values based on webhook (https://book.kubebuilder.io/cronjob-tutorial/webhook-implementation.html)
func (cm *CloudManaged) InitDefaults(defaults Defaults) bool {
	dirty := false
	if cm.Spec.Resources.Requests == nil {
		cm.Spec.Resources.Requests = defaults.Resources.Requests
		dirty = true
	}
	if cm.Spec.Resources.Limits == nil {
		cm.Spec.Resources.Limits = defaults.Resources.Limits
		dirty = true
	}
	if cm.Spec.VolumeSize == "" {
		cm.Spec.VolumeSize = defaults.VolumeSize
		dirty = true
	}
	if cm.Spec.Version == "" {
		cm.Spec.Version = defaults.Version
		dirty = true
	}

	return dirty
}

func init() {
	SchemeBuilder.Register(&CloudManaged{}, &CloudManagedList{})
}
