// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package tpu

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// A Cloud TPU instance.
// 
// 
// To get more information about Node, see:
// 
// * [API documentation](https://cloud.google.com/tpu/docs/reference/rest/)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/tpu/docs/)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/tpu_node.html.markdown.
type Node struct {
	s *pulumi.ResourceState
}

// NewNode registers a new resource with the given unique name, arguments, and options.
func NewNode(ctx *pulumi.Context,
	name string, args *NodeArgs, opts ...pulumi.ResourceOpt) (*Node, error) {
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
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["acceleratorType"] = nil
		inputs["cidrBlock"] = nil
		inputs["description"] = nil
		inputs["labels"] = nil
		inputs["name"] = nil
		inputs["network"] = nil
		inputs["project"] = nil
		inputs["schedulingConfig"] = nil
		inputs["tensorflowVersion"] = nil
		inputs["zone"] = nil
	} else {
		inputs["acceleratorType"] = args.AcceleratorType
		inputs["cidrBlock"] = args.CidrBlock
		inputs["description"] = args.Description
		inputs["labels"] = args.Labels
		inputs["name"] = args.Name
		inputs["network"] = args.Network
		inputs["project"] = args.Project
		inputs["schedulingConfig"] = args.SchedulingConfig
		inputs["tensorflowVersion"] = args.TensorflowVersion
		inputs["zone"] = args.Zone
	}
	inputs["networkEndpoints"] = nil
	inputs["serviceAccount"] = nil
	s, err := ctx.RegisterResource("gcp:tpu/node:Node", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Node{s: s}, nil
}

// GetNode gets an existing Node resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNode(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NodeState, opts ...pulumi.ResourceOpt) (*Node, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["acceleratorType"] = state.AcceleratorType
		inputs["cidrBlock"] = state.CidrBlock
		inputs["description"] = state.Description
		inputs["labels"] = state.Labels
		inputs["name"] = state.Name
		inputs["network"] = state.Network
		inputs["networkEndpoints"] = state.NetworkEndpoints
		inputs["project"] = state.Project
		inputs["schedulingConfig"] = state.SchedulingConfig
		inputs["serviceAccount"] = state.ServiceAccount
		inputs["tensorflowVersion"] = state.TensorflowVersion
		inputs["zone"] = state.Zone
	}
	s, err := ctx.ReadResource("gcp:tpu/node:Node", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Node{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Node) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Node) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *Node) AcceleratorType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["acceleratorType"])
}

func (r *Node) CidrBlock() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["cidrBlock"])
}

func (r *Node) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *Node) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

func (r *Node) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *Node) Network() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["network"])
}

func (r *Node) NetworkEndpoints() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["networkEndpoints"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *Node) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *Node) SchedulingConfig() *pulumi.Output {
	return r.s.State["schedulingConfig"]
}

func (r *Node) ServiceAccount() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceAccount"])
}

func (r *Node) TensorflowVersion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tensorflowVersion"])
}

func (r *Node) Zone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zone"])
}

// Input properties used for looking up and filtering Node resources.
type NodeState struct {
	AcceleratorType interface{}
	CidrBlock interface{}
	Description interface{}
	Labels interface{}
	Name interface{}
	Network interface{}
	NetworkEndpoints interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	SchedulingConfig interface{}
	ServiceAccount interface{}
	TensorflowVersion interface{}
	Zone interface{}
}

// The set of arguments for constructing a Node resource.
type NodeArgs struct {
	AcceleratorType interface{}
	CidrBlock interface{}
	Description interface{}
	Labels interface{}
	Name interface{}
	Network interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	SchedulingConfig interface{}
	TensorflowVersion interface{}
	Zone interface{}
}
