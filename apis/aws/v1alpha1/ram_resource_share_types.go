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

type RamResourceShare struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RamResourceShareSpec   `json:"spec,omitempty"`
	Status            RamResourceShareStatus `json:"status,omitempty"`
}

type RamResourceShareSpec struct {
	// +optional
	AllowExternalPrincipals bool   `json:"allow_external_principals,omitempty"`
	Name                    string `json:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type RamResourceShareStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RamResourceShareList is a list of RamResourceShares
type RamResourceShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RamResourceShare CRD objects
	Items []RamResourceShare `json:"items,omitempty"`
}