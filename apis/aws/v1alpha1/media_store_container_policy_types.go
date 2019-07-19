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

type MediaStoreContainerPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MediaStoreContainerPolicySpec   `json:"spec,omitempty"`
	Status            MediaStoreContainerPolicyStatus `json:"status,omitempty"`
}

type MediaStoreContainerPolicySpec struct {
	ContainerName string                    `json:"containerName" tf:"container_name"`
	Policy        string                    `json:"policy" tf:"policy"`
	ProviderRef   core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type MediaStoreContainerPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MediaStoreContainerPolicyList is a list of MediaStoreContainerPolicys
type MediaStoreContainerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MediaStoreContainerPolicy CRD objects
	Items []MediaStoreContainerPolicy `json:"items,omitempty"`
}
