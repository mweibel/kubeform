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

type IamServiceLinkedRole struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamServiceLinkedRoleSpec   `json:"spec,omitempty"`
	Status            IamServiceLinkedRoleStatus `json:"status,omitempty"`
}

type IamServiceLinkedRoleSpec struct {
	AwsServiceName string `json:"aws_service_name"`
	// +optional
	CustomSuffix string `json:"custom_suffix,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
}

type IamServiceLinkedRoleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamServiceLinkedRoleList is a list of IamServiceLinkedRoles
type IamServiceLinkedRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamServiceLinkedRole CRD objects
	Items []IamServiceLinkedRole `json:"items,omitempty"`
}