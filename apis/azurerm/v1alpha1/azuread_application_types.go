package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AzureadApplication struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzureadApplicationSpec   `json:"spec,omitempty"`
	Status            AzureadApplicationStatus `json:"status,omitempty"`
}

type AzureadApplicationSpec struct {
	// +optional
	AvailableToOtherTenants bool   `json:"available_to_other_tenants,omitempty"`
	Name                    string `json:"name"`
	// +optional
	Oauth2AllowImplicitFlow bool `json:"oauth2_allow_implicit_flow,omitempty"`
}

type AzureadApplicationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzureadApplicationList is a list of AzureadApplications
type AzureadApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzureadApplication CRD objects
	Items []AzureadApplication `json:"items,omitempty"`
}