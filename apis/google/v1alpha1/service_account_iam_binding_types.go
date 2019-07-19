package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ServiceAccountIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAccountIamBindingSpec   `json:"spec,omitempty"`
	Status            ServiceAccountIamBindingStatus `json:"status,omitempty"`
}

type ServiceAccountIamBindingSpec struct {
	// +kubebuilder:validation:UniqueItems=true
	Members          []string                  `json:"members" tf:"members"`
	Role             string                    `json:"role" tf:"role"`
	ServiceAccountID string                    `json:"serviceAccountID" tf:"service_account_id"`
	ProviderRef      core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ServiceAccountIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServiceAccountIamBindingList is a list of ServiceAccountIamBindings
type ServiceAccountIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceAccountIamBinding CRD objects
	Items []ServiceAccountIamBinding `json:"items,omitempty"`
}
