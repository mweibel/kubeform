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

type RdsClusterInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RdsClusterInstanceSpec   `json:"spec,omitempty"`
	Status            RdsClusterInstanceStatus `json:"status,omitempty"`
}

type RdsClusterInstanceSpec struct {
	// +optional
	ApplyImmediately bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`
	// +optional
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`
	// +optional
	AvailabilityZone  string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
	ClusterIdentifier string `json:"clusterIdentifier" tf:"cluster_identifier"`
	// +optional
	CopyTagsToSnapshot bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot,omitempty"`
	// +optional
	DbParameterGroupName string `json:"dbParameterGroupName,omitempty" tf:"db_parameter_group_name,omitempty"`
	// +optional
	DbSubnetGroupName string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name,omitempty"`
	// +optional
	Engine string `json:"engine,omitempty" tf:"engine,omitempty"`
	// +optional
	EngineVersion string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`
	// +optional
	Identifier string `json:"identifier,omitempty" tf:"identifier,omitempty"`
	// +optional
	IdentifierPrefix string `json:"identifierPrefix,omitempty" tf:"identifier_prefix,omitempty"`
	InstanceClass    string `json:"instanceClass" tf:"instance_class"`
	// +optional
	MonitoringInterval int `json:"monitoringInterval,omitempty" tf:"monitoring_interval,omitempty"`
	// +optional
	MonitoringRoleArn string `json:"monitoringRoleArn,omitempty" tf:"monitoring_role_arn,omitempty"`
	// +optional
	PerformanceInsightsEnabled bool `json:"performanceInsightsEnabled,omitempty" tf:"performance_insights_enabled,omitempty"`
	// +optional
	PerformanceInsightsKmsKeyID string `json:"performanceInsightsKmsKeyID,omitempty" tf:"performance_insights_kms_key_id,omitempty"`
	// +optional
	PreferredBackupWindow string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`
	// +optional
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`
	// +optional
	PromotionTier int `json:"promotionTier,omitempty" tf:"promotion_tier,omitempty"`
	// +optional
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type RdsClusterInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RdsClusterInstanceList is a list of RdsClusterInstances
type RdsClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RdsClusterInstance CRD objects
	Items []RdsClusterInstance `json:"items,omitempty"`
}
