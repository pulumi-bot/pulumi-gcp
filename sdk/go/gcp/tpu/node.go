// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package tpu

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Cloud TPU instance.
//
// To get more information about Node, see:
//
// * [API documentation](https://cloud.google.com/tpu/docs/reference/rest/)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/tpu/docs/)
type Node struct {
	pulumi.CustomResourceState

	// The type of hardware accelerators associated with this node.
	AcceleratorType pulumi.StringOutput `pulumi:"acceleratorType"`
	// The CIDR block that the TPU node will use when selecting an IP
	// address. This CIDR block must be a /29 block; the Compute Engine
	// networks API forbids a smaller block, and using a larger block would
	// be wasteful (a node can only consume one IP address).
	// Errors will occur if the CIDR block has already been used for a
	// currently existing TPU node, the CIDR block conflicts with any
	// subnetworks in the user's provided network, or the provided network
	// is peered with another network that is using that CIDR block.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Resource labels to represent user provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The immutable name of the TPU.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of a network to peer the TPU node to. It must be a
	// preexisting Compute Engine network inside of the project on which
	// this API has been activated. If none is provided, "default" will be
	// used.
	Network pulumi.StringOutput `pulumi:"network"`
	// The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the
	// node first reach out to the first (index 0) entry.
	NetworkEndpoints NodeNetworkEndpointArrayOutput `pulumi:"networkEndpoints"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Sets the scheduling options for this TPU instance.  Structure is documented below.
	SchedulingConfig NodeSchedulingConfigPtrOutput `pulumi:"schedulingConfig"`
	// The service account used to run the tensor flow services within the node. To share resources, including Google Cloud
	// Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// The version of Tensorflow running in the Node.
	TensorflowVersion pulumi.StringOutput `pulumi:"tensorflowVersion"`
	// The GCP location for the TPU.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewNode registers a new resource with the given unique name, arguments, and options.
func NewNode(ctx *pulumi.Context,
	name string, args *NodeArgs, opts ...pulumi.ResourceOption) (*Node, error) {
	if args == nil || args.AcceleratorType == nil {
		return nil, errors.New("missing required argument 'AcceleratorType'")
	}
	if args == nil || args.CidrBlock == nil {
		return nil, errors.New("missing required argument 'CidrBlock'")
	}
	if args == nil || args.TensorflowVersion == nil {
		return nil, errors.New("missing required argument 'TensorflowVersion'")
	}
	if args == nil || args.Zone == nil {
		return nil, errors.New("missing required argument 'Zone'")
	}
	if args == nil {
		args = &NodeArgs{}
	}
	var resource Node
	err := ctx.RegisterResource("gcp:tpu/node:Node", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNode gets an existing Node resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeState, opts ...pulumi.ResourceOption) (*Node, error) {
	var resource Node
	err := ctx.ReadResource("gcp:tpu/node:Node", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Node resources.
type nodeState struct {
	// The type of hardware accelerators associated with this node.
	AcceleratorType *string `pulumi:"acceleratorType"`
	// The CIDR block that the TPU node will use when selecting an IP
	// address. This CIDR block must be a /29 block; the Compute Engine
	// networks API forbids a smaller block, and using a larger block would
	// be wasteful (a node can only consume one IP address).
	// Errors will occur if the CIDR block has already been used for a
	// currently existing TPU node, the CIDR block conflicts with any
	// subnetworks in the user's provided network, or the provided network
	// is peered with another network that is using that CIDR block.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description *string `pulumi:"description"`
	// Resource labels to represent user provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// The immutable name of the TPU.
	Name *string `pulumi:"name"`
	// The name of a network to peer the TPU node to. It must be a
	// preexisting Compute Engine network inside of the project on which
	// this API has been activated. If none is provided, "default" will be
	// used.
	Network *string `pulumi:"network"`
	// The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the
	// node first reach out to the first (index 0) entry.
	NetworkEndpoints []NodeNetworkEndpoint `pulumi:"networkEndpoints"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Sets the scheduling options for this TPU instance.  Structure is documented below.
	SchedulingConfig *NodeSchedulingConfig `pulumi:"schedulingConfig"`
	// The service account used to run the tensor flow services within the node. To share resources, including Google Cloud
	// Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// The version of Tensorflow running in the Node.
	TensorflowVersion *string `pulumi:"tensorflowVersion"`
	// The GCP location for the TPU.
	Zone *string `pulumi:"zone"`
}

type NodeState struct {
	// The type of hardware accelerators associated with this node.
	AcceleratorType pulumi.StringPtrInput
	// The CIDR block that the TPU node will use when selecting an IP
	// address. This CIDR block must be a /29 block; the Compute Engine
	// networks API forbids a smaller block, and using a larger block would
	// be wasteful (a node can only consume one IP address).
	// Errors will occur if the CIDR block has already been used for a
	// currently existing TPU node, the CIDR block conflicts with any
	// subnetworks in the user's provided network, or the provided network
	// is peered with another network that is using that CIDR block.
	CidrBlock pulumi.StringPtrInput
	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description pulumi.StringPtrInput
	// Resource labels to represent user provided metadata.
	Labels pulumi.StringMapInput
	// The immutable name of the TPU.
	Name pulumi.StringPtrInput
	// The name of a network to peer the TPU node to. It must be a
	// preexisting Compute Engine network inside of the project on which
	// this API has been activated. If none is provided, "default" will be
	// used.
	Network pulumi.StringPtrInput
	// The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the
	// node first reach out to the first (index 0) entry.
	NetworkEndpoints NodeNetworkEndpointArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Sets the scheduling options for this TPU instance.  Structure is documented below.
	SchedulingConfig NodeSchedulingConfigPtrInput
	// The service account used to run the tensor flow services within the node. To share resources, including Google Cloud
	// Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
	ServiceAccount pulumi.StringPtrInput
	// The version of Tensorflow running in the Node.
	TensorflowVersion pulumi.StringPtrInput
	// The GCP location for the TPU.
	Zone pulumi.StringPtrInput
}

func (NodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeState)(nil)).Elem()
}

type nodeArgs struct {
	// The type of hardware accelerators associated with this node.
	AcceleratorType string `pulumi:"acceleratorType"`
	// The CIDR block that the TPU node will use when selecting an IP
	// address. This CIDR block must be a /29 block; the Compute Engine
	// networks API forbids a smaller block, and using a larger block would
	// be wasteful (a node can only consume one IP address).
	// Errors will occur if the CIDR block has already been used for a
	// currently existing TPU node, the CIDR block conflicts with any
	// subnetworks in the user's provided network, or the provided network
	// is peered with another network that is using that CIDR block.
	CidrBlock string `pulumi:"cidrBlock"`
	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description *string `pulumi:"description"`
	// Resource labels to represent user provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// The immutable name of the TPU.
	Name *string `pulumi:"name"`
	// The name of a network to peer the TPU node to. It must be a
	// preexisting Compute Engine network inside of the project on which
	// this API has been activated. If none is provided, "default" will be
	// used.
	Network *string `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Sets the scheduling options for this TPU instance.  Structure is documented below.
	SchedulingConfig *NodeSchedulingConfig `pulumi:"schedulingConfig"`
	// The version of Tensorflow running in the Node.
	TensorflowVersion string `pulumi:"tensorflowVersion"`
	// The GCP location for the TPU.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a Node resource.
type NodeArgs struct {
	// The type of hardware accelerators associated with this node.
	AcceleratorType pulumi.StringInput
	// The CIDR block that the TPU node will use when selecting an IP
	// address. This CIDR block must be a /29 block; the Compute Engine
	// networks API forbids a smaller block, and using a larger block would
	// be wasteful (a node can only consume one IP address).
	// Errors will occur if the CIDR block has already been used for a
	// currently existing TPU node, the CIDR block conflicts with any
	// subnetworks in the user's provided network, or the provided network
	// is peered with another network that is using that CIDR block.
	CidrBlock pulumi.StringInput
	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description pulumi.StringPtrInput
	// Resource labels to represent user provided metadata.
	Labels pulumi.StringMapInput
	// The immutable name of the TPU.
	Name pulumi.StringPtrInput
	// The name of a network to peer the TPU node to. It must be a
	// preexisting Compute Engine network inside of the project on which
	// this API has been activated. If none is provided, "default" will be
	// used.
	Network pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Sets the scheduling options for this TPU instance.  Structure is documented below.
	SchedulingConfig NodeSchedulingConfigPtrInput
	// The version of Tensorflow running in the Node.
	TensorflowVersion pulumi.StringInput
	// The GCP location for the TPU.
	Zone pulumi.StringInput
}

func (NodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeArgs)(nil)).Elem()
}
