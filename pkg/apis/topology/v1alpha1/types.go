package v1alpha1

import (
	_ "github.com/gogo/protobuf/gogoproto"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster,shortName=node-res-topo
// +kubebuilder:metadata:annotations="api-approved.kubernetes.io=https://github.com/kubernetes/enhancements/pull/1870"
// +kubebuilder:storageversion

// NodeResourceTopology describes node resources and their topology.
type NodeResourceTopology struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Zones      ZoneList      `json:"zones"`
	Attributes AttributeList `json:"attributes,omitempty"`
}

// Zone represents a resource topology zone, e.g. socket, node, die or core.
// +protobuf=true
type Zone struct {
	Name       string           `json:"name" protobuf:"bytes,1,opt,name=name"`
	Type       string           `json:"type" protobuf:"bytes,2,opt,name=type"`
	Parent     string           `json:"parent,omitempty" protobuf:"bytes,3,opt,name=parent"`
	Costs      CostList         `json:"costs,omitempty" protobuf:"bytes,4,rep,name=costs"`
	Attributes AttributeList    `json:"attributes,omitempty" protobuf:"bytes,5,rep,name=attributes"`
	Resources  ResourceInfoList `json:"resources,omitempty" protobuf:"bytes,6,rep,name=resources"`
}

// ZoneList contains an array of Zone objects.
// +protobuf=true
type ZoneList []Zone

// ResourceInfo contains information about one resource type.
// +protobuf=true
type ResourceInfo struct {
	// Name of the resource.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// Capacity of the resource, corresponding to capacity in node status, i.e.
	// total amount of this resource that the node has.
	Capacity resource.Quantity `json:"capacity" protobuf:"bytes,2,opt,name=capacity"`
	// Allocatable quantity of the resource, corresponding to allocatable in
	// node status, i.e. total amount of this resource available to be used by
	// pods.
	Allocatable resource.Quantity `json:"allocatable" protobuf:"bytes,3,opt,name=allocatable"`
	// Available is the amount of this resource currently available for new (to
	// be scheduled) pods, i.e. Allocatable minus the resources reserved by
	// currently running pods.
	Available resource.Quantity `json:"available" protobuf:"bytes,4,opt,name=available"`
}

// ResourceInfoList contains an array of ResourceInfo objects.
// +protobuf=true
type ResourceInfoList []ResourceInfo

// CostInfo describes the cost (or distance) between two Zones.
// +protobuf=true
type CostInfo struct {
	Name  string `json:"name" protobuf:"bytes,1,opt,name=name"`
	Value int64  `json:"value" protobuf:"varint,2,opt,name=value"`
}

// CostList contains an array of CostInfo objects.
// +protobuf=true
type CostList []CostInfo

// AttributeInfo contains one attribute of a Zone.
// +protobuf=true
type AttributeInfo struct {
	Name  string `json:"name" protobuf:"bytes,1,opt,name=name"`
	Value string `json:"value" protobuf:"bytes,2,opt,name=value"`
}

// AttributeList contains an array of AttributeInfo objects.
// +protobuf=true
type AttributeList []AttributeInfo

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeResourceTopologyList is a list of NodeResourceTopology resources
type NodeResourceTopologyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NodeResourceTopology `json:"items"`
}
