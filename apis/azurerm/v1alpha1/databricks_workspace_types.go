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

type DatabricksWorkspace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabricksWorkspaceSpec   `json:"spec,omitempty"`
	Status            DatabricksWorkspaceStatus `json:"status,omitempty"`
}

type DatabricksWorkspaceSpec struct {
	Location string `json:"location" tf:"location"`
	// +optional
	ManagedResourceGroupName string `json:"managedResourceGroupName,omitempty" tf:"managed_resource_group_name,omitempty"`
	Name                     string `json:"name" tf:"name"`
	ResourceGroupName        string `json:"resourceGroupName" tf:"resource_group_name"`
	Sku                      string `json:"sku" tf:"sku"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DatabricksWorkspaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DatabricksWorkspaceList is a list of DatabricksWorkspaces
type DatabricksWorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DatabricksWorkspace CRD objects
	Items []DatabricksWorkspace `json:"items,omitempty"`
}
