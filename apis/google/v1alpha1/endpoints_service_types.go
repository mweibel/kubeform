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

type EndpointsService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointsServiceSpec   `json:"spec,omitempty"`
	Status            EndpointsServiceStatus `json:"status,omitempty"`
}

type EndpointsServiceSpec struct {
	// +optional
	GrpcConfig string `json:"grpc_config,omitempty"`
	// +optional
	OpenapiConfig string `json:"openapi_config,omitempty"`
	// +optional
	// Deprecated
	ProtocOutput string `json:"protoc_output,omitempty"`
	// +optional
	ProtocOutputBase64 string `json:"protoc_output_base64,omitempty"`
	ServiceName        string `json:"service_name"`
}

type EndpointsServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EndpointsServiceList is a list of EndpointsServices
type EndpointsServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EndpointsService CRD objects
	Items []EndpointsService `json:"items,omitempty"`
}