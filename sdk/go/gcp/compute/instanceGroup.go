// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_instance_group.html.markdown.
type InstanceGroup struct {
	s *pulumi.ResourceState
}

// NewInstanceGroup registers a new resource with the given unique name, arguments, and options.
func NewInstanceGroup(ctx *pulumi.Context,
	name string, args *InstanceGroupArgs, opts ...pulumi.ResourceOpt) (*InstanceGroup, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["instances"] = nil
		inputs["name"] = nil
		inputs["namedPorts"] = nil
		inputs["network"] = nil
		inputs["project"] = nil
		inputs["zone"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["instances"] = args.Instances
		inputs["name"] = args.Name
		inputs["namedPorts"] = args.NamedPorts
		inputs["network"] = args.Network
		inputs["project"] = args.Project
		inputs["zone"] = args.Zone
	}
	inputs["selfLink"] = nil
	inputs["size"] = nil
	s, err := ctx.RegisterResource("gcp:compute/instanceGroup:InstanceGroup", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &InstanceGroup{s: s}, nil
}

// GetInstanceGroup gets an existing InstanceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *InstanceGroupState, opts ...pulumi.ResourceOpt) (*InstanceGroup, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["instances"] = state.Instances
		inputs["name"] = state.Name
		inputs["namedPorts"] = state.NamedPorts
		inputs["network"] = state.Network
		inputs["project"] = state.Project
		inputs["selfLink"] = state.SelfLink
		inputs["size"] = state.Size
		inputs["zone"] = state.Zone
	}
	s, err := ctx.ReadResource("gcp:compute/instanceGroup:InstanceGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &InstanceGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *InstanceGroup) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *InstanceGroup) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// An optional textual description of the instance
// group.
func (r *InstanceGroup) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// List of instances in the group. They should be given
// as self_link URLs. When adding instances they must all be in the same
// network and zone as the instance group.
func (r *InstanceGroup) Instances() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["instances"])
}

// The name of the instance group. Must be 1-63
// characters long and comply with
// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
// include lowercase letters, numbers, and hyphens.
func (r *InstanceGroup) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The named port configuration. See the section below
// for details on configuration.
func (r *InstanceGroup) NamedPorts() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["namedPorts"])
}

// The URL of the network the instance group is in. If
// this is different from the network where the instances are in, the creation
// fails. Defaults to the network where the instances are in (if neither
// `network` nor `instances` is specified, this field will be blank).
func (r *InstanceGroup) Network() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["network"])
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *InstanceGroup) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The URI of the created resource.
func (r *InstanceGroup) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

// The number of instances in the group.
func (r *InstanceGroup) Size() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["size"])
}

// The zone that this instance group should be created in.
func (r *InstanceGroup) Zone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zone"])
}

// Input properties used for looking up and filtering InstanceGroup resources.
type InstanceGroupState struct {
	// An optional textual description of the instance
	// group.
	Description interface{}
	// List of instances in the group. They should be given
	// as self_link URLs. When adding instances they must all be in the same
	// network and zone as the instance group.
	Instances interface{}
	// The name of the instance group. Must be 1-63
	// characters long and comply with
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
	// include lowercase letters, numbers, and hyphens.
	Name interface{}
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts interface{}
	// The URL of the network the instance group is in. If
	// this is different from the network where the instances are in, the creation
	// fails. Defaults to the network where the instances are in (if neither
	// `network` nor `instances` is specified, this field will be blank).
	Network interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The URI of the created resource.
	SelfLink interface{}
	// The number of instances in the group.
	Size interface{}
	// The zone that this instance group should be created in.
	Zone interface{}
}

// The set of arguments for constructing a InstanceGroup resource.
type InstanceGroupArgs struct {
	// An optional textual description of the instance
	// group.
	Description interface{}
	// List of instances in the group. They should be given
	// as self_link URLs. When adding instances they must all be in the same
	// network and zone as the instance group.
	Instances interface{}
	// The name of the instance group. Must be 1-63
	// characters long and comply with
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
	// include lowercase letters, numbers, and hyphens.
	Name interface{}
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts interface{}
	// The URL of the network the instance group is in. If
	// this is different from the network where the instances are in, the creation
	// fails. Defaults to the network where the instances are in (if neither
	// `network` nor `instances` is specified, this field will be blank).
	Network interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The zone that this instance group should be created in.
	Zone interface{}
}
