// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a NodeGroup resource to manage a group of sole-tenant nodes.
//
// To get more information about NodeGroup, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/nodeGroups)
// * How-to Guides
//     * [Sole-Tenant Nodes](https://cloud.google.com/compute/docs/nodes/)
//
// > **Warning:** Due to limitations of the API, this provider cannot update the
// number of nodes in a node group and changes to node group size either
// through provider config or through external changes will cause
// the provider to delete and recreate the node group.
//
// ## Example Usage
type NodeGroup struct {
	pulumi.CustomResourceState

	// -
	// If you use sole-tenant nodes for your workloads, you can use the node
	// group autoscaler to automatically manage the sizes of your node groups.  Structure is documented below.
	AutoscalingPolicy NodeGroupAutoscalingPolicyOutput `pulumi:"autoscalingPolicy"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional textual description of the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL of the node template to which this node group belongs.
	NodeTemplate pulumi.StringOutput `pulumi:"nodeTemplate"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The total number of nodes in the node group.
	Size pulumi.IntOutput `pulumi:"size"`
	// Zone where this node group is located
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewNodeGroup registers a new resource with the given unique name, arguments, and options.
func NewNodeGroup(ctx *pulumi.Context,
	name string, args *NodeGroupArgs, opts ...pulumi.ResourceOption) (*NodeGroup, error) {
	if args == nil || args.NodeTemplate == nil {
		return nil, errors.New("missing required argument 'NodeTemplate'")
	}
	if args == nil || args.Size == nil {
		return nil, errors.New("missing required argument 'Size'")
	}
	if args == nil {
		args = &NodeGroupArgs{}
	}
	var resource NodeGroup
	err := ctx.RegisterResource("gcp:compute/nodeGroup:NodeGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodeGroup gets an existing NodeGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodeGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeGroupState, opts ...pulumi.ResourceOption) (*NodeGroup, error) {
	var resource NodeGroup
	err := ctx.ReadResource("gcp:compute/nodeGroup:NodeGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NodeGroup resources.
type nodeGroupState struct {
	// -
	// If you use sole-tenant nodes for your workloads, you can use the node
	// group autoscaler to automatically manage the sizes of your node groups.  Structure is documented below.
	AutoscalingPolicy *NodeGroupAutoscalingPolicy `pulumi:"autoscalingPolicy"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional textual description of the resource.
	Description *string `pulumi:"description"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// The URL of the node template to which this node group belongs.
	NodeTemplate *string `pulumi:"nodeTemplate"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The total number of nodes in the node group.
	Size *int `pulumi:"size"`
	// Zone where this node group is located
	Zone *string `pulumi:"zone"`
}

type NodeGroupState struct {
	// -
	// If you use sole-tenant nodes for your workloads, you can use the node
	// group autoscaler to automatically manage the sizes of your node groups.  Structure is documented below.
	AutoscalingPolicy NodeGroupAutoscalingPolicyPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional textual description of the resource.
	Description pulumi.StringPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// The URL of the node template to which this node group belongs.
	NodeTemplate pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The total number of nodes in the node group.
	Size pulumi.IntPtrInput
	// Zone where this node group is located
	Zone pulumi.StringPtrInput
}

func (NodeGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeGroupState)(nil)).Elem()
}

type nodeGroupArgs struct {
	// -
	// If you use sole-tenant nodes for your workloads, you can use the node
	// group autoscaler to automatically manage the sizes of your node groups.  Structure is documented below.
	AutoscalingPolicy *NodeGroupAutoscalingPolicy `pulumi:"autoscalingPolicy"`
	// An optional textual description of the resource.
	Description *string `pulumi:"description"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// The URL of the node template to which this node group belongs.
	NodeTemplate string `pulumi:"nodeTemplate"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The total number of nodes in the node group.
	Size int `pulumi:"size"`
	// Zone where this node group is located
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a NodeGroup resource.
type NodeGroupArgs struct {
	// -
	// If you use sole-tenant nodes for your workloads, you can use the node
	// group autoscaler to automatically manage the sizes of your node groups.  Structure is documented below.
	AutoscalingPolicy NodeGroupAutoscalingPolicyPtrInput
	// An optional textual description of the resource.
	Description pulumi.StringPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// The URL of the node template to which this node group belongs.
	NodeTemplate pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The total number of nodes in the node group.
	Size pulumi.IntInput
	// Zone where this node group is located
	Zone pulumi.StringPtrInput
}

func (NodeGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeGroupArgs)(nil)).Elem()
}
