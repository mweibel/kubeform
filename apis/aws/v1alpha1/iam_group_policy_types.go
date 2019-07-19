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

type IamGroupPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamGroupPolicySpec   `json:"spec,omitempty"`
	Status            IamGroupPolicyStatus `json:"status,omitempty"`
}

type IamGroupPolicySpec struct {
	Group string `json:"group" tf:"group"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix  string                    `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	Policy      string                    `json:"policy" tf:"policy"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type IamGroupPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamGroupPolicyList is a list of IamGroupPolicys
type IamGroupPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamGroupPolicy CRD objects
	Items []IamGroupPolicy `json:"items,omitempty"`
}
