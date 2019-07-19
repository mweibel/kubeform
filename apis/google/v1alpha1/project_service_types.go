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

type ProjectService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectServiceSpec   `json:"spec,omitempty"`
	Status            ProjectServiceStatus `json:"status,omitempty"`
}

type ProjectServiceSpec struct {
	// +optional
	DisableOnDestroy bool `json:"disableOnDestroy,omitempty" tf:"disable_on_destroy,omitempty"`
	// +optional
	Project     string                    `json:"project,omitempty" tf:"project,omitempty"`
	Service     string                    `json:"service" tf:"service"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ProjectServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ProjectServiceList is a list of ProjectServices
type ProjectServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ProjectService CRD objects
	Items []ProjectService `json:"items,omitempty"`
}
