// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package container

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a node pool in a Google Kubernetes Engine (GKE) cluster separately from
// the cluster control plane. For more information see [the official documentation](https://cloud.google.com/container-engine/docs/node-pools)
// and [the API reference](https://cloud.google.com/container-engine/reference/rest/v1/projects.zones.clusters.nodePools).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/container_node_pool.html.markdown.
type NodePool struct {
	s *pulumi.ResourceState
}

// NewNodePool registers a new resource with the given unique name, arguments, and options.
func NewNodePool(ctx *pulumi.Context,
	name string, args *NodePoolArgs, opts ...pulumi.ResourceOpt) (*NodePool, error) {
	if args == nil || args.Cluster == nil {
		return nil, errors.New("missing required argument 'Cluster'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["autoscaling"] = nil
		inputs["cluster"] = nil
		inputs["initialNodeCount"] = nil
		inputs["location"] = nil
		inputs["management"] = nil
		inputs["maxPodsPerNode"] = nil
		inputs["name"] = nil
		inputs["namePrefix"] = nil
		inputs["nodeConfig"] = nil
		inputs["nodeCount"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["version"] = nil
		inputs["zone"] = nil
	} else {
		inputs["autoscaling"] = args.Autoscaling
		inputs["cluster"] = args.Cluster
		inputs["initialNodeCount"] = args.InitialNodeCount
		inputs["location"] = args.Location
		inputs["management"] = args.Management
		inputs["maxPodsPerNode"] = args.MaxPodsPerNode
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["nodeConfig"] = args.NodeConfig
		inputs["nodeCount"] = args.NodeCount
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["version"] = args.Version
		inputs["zone"] = args.Zone
	}
	inputs["instanceGroupUrls"] = nil
	s, err := ctx.RegisterResource("gcp:container/nodePool:NodePool", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NodePool{s: s}, nil
}

// GetNodePool gets an existing NodePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodePool(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NodePoolState, opts ...pulumi.ResourceOpt) (*NodePool, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["autoscaling"] = state.Autoscaling
		inputs["cluster"] = state.Cluster
		inputs["initialNodeCount"] = state.InitialNodeCount
		inputs["instanceGroupUrls"] = state.InstanceGroupUrls
		inputs["location"] = state.Location
		inputs["management"] = state.Management
		inputs["maxPodsPerNode"] = state.MaxPodsPerNode
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
		inputs["nodeConfig"] = state.NodeConfig
		inputs["nodeCount"] = state.NodeCount
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["version"] = state.Version
		inputs["zone"] = state.Zone
	}
	s, err := ctx.ReadResource("gcp:container/nodePool:NodePool", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NodePool{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *NodePool) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *NodePool) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Configuration required by cluster autoscaler to adjust
// the size of the node pool to the current cluster usage. Structure is documented below.
func (r *NodePool) Autoscaling() *pulumi.Output {
	return r.s.State["autoscaling"]
}

// The cluster to create the node pool for.  Cluster must be present in `zone` provided for zonal clusters.
func (r *NodePool) Cluster() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["cluster"])
}

// The initial node count for the pool. Changing this will force
// recreation of the resource.
func (r *NodePool) InitialNodeCount() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["initialNodeCount"])
}

func (r *NodePool) InstanceGroupUrls() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["instanceGroupUrls"])
}

// The location (region or zone) in which the cluster
// resides.
func (r *NodePool) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// Node management configuration, wherein auto-repair and
// auto-upgrade is configured. Structure is documented below.
func (r *NodePool) Management() *pulumi.Output {
	return r.s.State["management"]
}

// ) The maximum number of pods per node in this node pool.
// Note that this does not work on node pools which are "route-based" - that is, node
// pools belonging to clusters that do not have IP Aliasing enabled.
// See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
// for more information.
func (r *NodePool) MaxPodsPerNode() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["maxPodsPerNode"])
}

// The name of the node pool. If left blank, this provider will
// auto-generate a unique name.
func (r *NodePool) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *NodePool) NamePrefix() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["namePrefix"])
}

// The node configuration of the pool. See
// container.Cluster for schema.
func (r *NodePool) NodeConfig() *pulumi.Output {
	return r.s.State["nodeConfig"]
}

// The number of nodes per instance group. This field can be used to
// update the number of nodes per instance group but should not be used alongside `autoscaling`.
func (r *NodePool) NodeCount() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["nodeCount"])
}

// The ID of the project in which to create the node pool. If blank,
// the provider-configured project will be used.
func (r *NodePool) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The region in which the cluster resides (for
// regional clusters). `zone` has been deprecated in favor of `location`.
func (r *NodePool) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The Kubernetes version for the nodes in this pool. Note that if this field
// and `autoUpgrade` are both specified, they will fight each other for what the node version should
// be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
// recommended that you specify explicit versions as this provider will see spurious diffs
// when fuzzy versions are used. See the `googleContainerEngineVersions` data source's
// `versionPrefix` field to approximate fuzzy versions.
func (r *NodePool) Version() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["version"])
}

// The zone in which the cluster resides. `zone`
// has been deprecated in favor of `location`.
func (r *NodePool) Zone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zone"])
}

// Input properties used for looking up and filtering NodePool resources.
type NodePoolState struct {
	// Configuration required by cluster autoscaler to adjust
	// the size of the node pool to the current cluster usage. Structure is documented below.
	Autoscaling interface{}
	// The cluster to create the node pool for.  Cluster must be present in `zone` provided for zonal clusters.
	Cluster interface{}
	// The initial node count for the pool. Changing this will force
	// recreation of the resource.
	InitialNodeCount interface{}
	InstanceGroupUrls interface{}
	// The location (region or zone) in which the cluster
	// resides.
	Location interface{}
	// Node management configuration, wherein auto-repair and
	// auto-upgrade is configured. Structure is documented below.
	Management interface{}
	// ) The maximum number of pods per node in this node pool.
	// Note that this does not work on node pools which are "route-based" - that is, node
	// pools belonging to clusters that do not have IP Aliasing enabled.
	// See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
	// for more information.
	MaxPodsPerNode interface{}
	// The name of the node pool. If left blank, this provider will
	// auto-generate a unique name.
	Name interface{}
	NamePrefix interface{}
	// The node configuration of the pool. See
	// container.Cluster for schema.
	NodeConfig interface{}
	// The number of nodes per instance group. This field can be used to
	// update the number of nodes per instance group but should not be used alongside `autoscaling`.
	NodeCount interface{}
	// The ID of the project in which to create the node pool. If blank,
	// the provider-configured project will be used.
	Project interface{}
	// The region in which the cluster resides (for
	// regional clusters). `zone` has been deprecated in favor of `location`.
	Region interface{}
	// The Kubernetes version for the nodes in this pool. Note that if this field
	// and `autoUpgrade` are both specified, they will fight each other for what the node version should
	// be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
	// recommended that you specify explicit versions as this provider will see spurious diffs
	// when fuzzy versions are used. See the `googleContainerEngineVersions` data source's
	// `versionPrefix` field to approximate fuzzy versions.
	Version interface{}
	// The zone in which the cluster resides. `zone`
	// has been deprecated in favor of `location`.
	Zone interface{}
}

// The set of arguments for constructing a NodePool resource.
type NodePoolArgs struct {
	// Configuration required by cluster autoscaler to adjust
	// the size of the node pool to the current cluster usage. Structure is documented below.
	Autoscaling interface{}
	// The cluster to create the node pool for.  Cluster must be present in `zone` provided for zonal clusters.
	Cluster interface{}
	// The initial node count for the pool. Changing this will force
	// recreation of the resource.
	InitialNodeCount interface{}
	// The location (region or zone) in which the cluster
	// resides.
	Location interface{}
	// Node management configuration, wherein auto-repair and
	// auto-upgrade is configured. Structure is documented below.
	Management interface{}
	// ) The maximum number of pods per node in this node pool.
	// Note that this does not work on node pools which are "route-based" - that is, node
	// pools belonging to clusters that do not have IP Aliasing enabled.
	// See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
	// for more information.
	MaxPodsPerNode interface{}
	// The name of the node pool. If left blank, this provider will
	// auto-generate a unique name.
	Name interface{}
	NamePrefix interface{}
	// The node configuration of the pool. See
	// container.Cluster for schema.
	NodeConfig interface{}
	// The number of nodes per instance group. This field can be used to
	// update the number of nodes per instance group but should not be used alongside `autoscaling`.
	NodeCount interface{}
	// The ID of the project in which to create the node pool. If blank,
	// the provider-configured project will be used.
	Project interface{}
	// The region in which the cluster resides (for
	// regional clusters). `zone` has been deprecated in favor of `location`.
	Region interface{}
	// The Kubernetes version for the nodes in this pool. Note that if this field
	// and `autoUpgrade` are both specified, they will fight each other for what the node version should
	// be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
	// recommended that you specify explicit versions as this provider will see spurious diffs
	// when fuzzy versions are used. See the `googleContainerEngineVersions` data source's
	// `versionPrefix` field to approximate fuzzy versions.
	Version interface{}
	// The zone in which the cluster resides. `zone`
	// has been deprecated in favor of `location`.
	Zone interface{}
}
