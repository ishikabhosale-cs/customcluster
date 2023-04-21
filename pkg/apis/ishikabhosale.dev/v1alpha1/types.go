package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo is a specification for a Foo resource
type Customcluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomclusterSpec   `json:"spec"`
	Status CustomclusterStatus `json:"status,omitempty"`
}

type CustomclusterStatus struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
}

// FooSpec is the spec for a Foo resource
type CustomclusterSpec struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo resources
type CustomclusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Customcluster `json:"items"`
}
