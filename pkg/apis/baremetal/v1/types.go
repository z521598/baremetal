package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BareMetalJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BareMetalJobSpec `json:"spec"`
}

type BareMetalJobSpec struct {
	ResourceType string `json:"resourceType"`
	ResourceId   string `json:"resourceId"`
	TaskType     string `json:"taskType"`
	Commands     string `json:"commands"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BareMetalJobList is a list of BareMetalJob resources
type BareMetalJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []BareMetalJob `json:"items"`
}
