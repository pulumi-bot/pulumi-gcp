// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a VM instance template resource within GCE. For more information see
// [the official documentation](https://cloud.google.com/compute/docs/instance-templates)
// and
// [API](https://cloud.google.com/compute/docs/reference/latest/instanceTemplates).
type InstanceTemplate struct {
	s *pulumi.ResourceState
}

// NewInstanceTemplate registers a new resource with the given unique name, arguments, and options.
func NewInstanceTemplate(ctx *pulumi.Context,
	name string, args *InstanceTemplateArgs, opts ...pulumi.ResourceOpt) (*InstanceTemplate, error) {
	if args == nil || args.Disks == nil {
		return nil, errors.New("missing required argument 'Disks'")
	}
	if args == nil || args.MachineType == nil {
		return nil, errors.New("missing required argument 'MachineType'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["canIpForward"] = nil
		inputs["description"] = nil
		inputs["disks"] = nil
		inputs["guestAccelerators"] = nil
		inputs["instanceDescription"] = nil
		inputs["labels"] = nil
		inputs["machineType"] = nil
		inputs["metadata"] = nil
		inputs["metadataStartupScript"] = nil
		inputs["minCpuPlatform"] = nil
		inputs["name"] = nil
		inputs["namePrefix"] = nil
		inputs["networkInterfaces"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["schedulings"] = nil
		inputs["serviceAccount"] = nil
		inputs["tags"] = nil
	} else {
		inputs["canIpForward"] = args.CanIpForward
		inputs["description"] = args.Description
		inputs["disks"] = args.Disks
		inputs["guestAccelerators"] = args.GuestAccelerators
		inputs["instanceDescription"] = args.InstanceDescription
		inputs["labels"] = args.Labels
		inputs["machineType"] = args.MachineType
		inputs["metadata"] = args.Metadata
		inputs["metadataStartupScript"] = args.MetadataStartupScript
		inputs["minCpuPlatform"] = args.MinCpuPlatform
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["networkInterfaces"] = args.NetworkInterfaces
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["schedulings"] = args.Schedulings
		inputs["serviceAccount"] = args.ServiceAccount
		inputs["tags"] = args.Tags
	}
	inputs["metadataFingerprint"] = nil
	inputs["selfLink"] = nil
	inputs["tagsFingerprint"] = nil
	s, err := ctx.RegisterResource("gcp:compute/instanceTemplate:InstanceTemplate", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &InstanceTemplate{s: s}, nil
}

// GetInstanceTemplate gets an existing InstanceTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceTemplate(ctx *pulumi.Context,
	name string, id pulumi.ID, state *InstanceTemplateState, opts ...pulumi.ResourceOpt) (*InstanceTemplate, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["canIpForward"] = state.CanIpForward
		inputs["description"] = state.Description
		inputs["disks"] = state.Disks
		inputs["guestAccelerators"] = state.GuestAccelerators
		inputs["instanceDescription"] = state.InstanceDescription
		inputs["labels"] = state.Labels
		inputs["machineType"] = state.MachineType
		inputs["metadata"] = state.Metadata
		inputs["metadataFingerprint"] = state.MetadataFingerprint
		inputs["metadataStartupScript"] = state.MetadataStartupScript
		inputs["minCpuPlatform"] = state.MinCpuPlatform
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
		inputs["networkInterfaces"] = state.NetworkInterfaces
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["schedulings"] = state.Schedulings
		inputs["selfLink"] = state.SelfLink
		inputs["serviceAccount"] = state.ServiceAccount
		inputs["tags"] = state.Tags
		inputs["tagsFingerprint"] = state.TagsFingerprint
	}
	s, err := ctx.ReadResource("gcp:compute/instanceTemplate:InstanceTemplate", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &InstanceTemplate{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *InstanceTemplate) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *InstanceTemplate) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Whether to allow sending and receiving of
// packets with non-matching source or destination IPs. This defaults to false.
func (r *InstanceTemplate) CanIpForward() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["canIpForward"])
}

// A brief description of this resource.
func (r *InstanceTemplate) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// Disks to attach to instances created from this template.
// This can be specified multiple times for multiple disks. Structure is
// documented below.
func (r *InstanceTemplate) Disks() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["disks"])
}

// List of the type and count of accelerator cards attached to the instance. Structure documented below.
func (r *InstanceTemplate) GuestAccelerators() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["guestAccelerators"])
}

// A brief description to use for instances
// created from this template.
func (r *InstanceTemplate) InstanceDescription() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instanceDescription"])
}

// A set of key/value label pairs to assign to instances
// created from this template,
func (r *InstanceTemplate) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

// The machine type to create.
func (r *InstanceTemplate) MachineType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["machineType"])
}

// Metadata key/value pairs to make available from
// within instances created from this template.
func (r *InstanceTemplate) Metadata() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["metadata"])
}

// The unique fingerprint of the metadata.
func (r *InstanceTemplate) MetadataFingerprint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["metadataFingerprint"])
}

// An alternative to using the
// startup-script metadata key, mostly to match the compute_instance resource.
// This replaces the startup-script metadata key on the created instance and
// thus the two mechanisms are not allowed to be used simultaneously.
func (r *InstanceTemplate) MetadataStartupScript() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["metadataStartupScript"])
}

// Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
func (r *InstanceTemplate) MinCpuPlatform() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["minCpuPlatform"])
}

// The name of the instance template. If you leave
// this blank, Terraform will auto-generate a unique name.
func (r *InstanceTemplate) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Creates a unique name beginning with the specified
// prefix. Conflicts with `name`.
func (r *InstanceTemplate) NamePrefix() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["namePrefix"])
}

// Networks to attach to instances created from
// this template. This can be specified multiple times for multiple networks.
// Structure is documented below.
func (r *InstanceTemplate) NetworkInterfaces() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["networkInterfaces"])
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *InstanceTemplate) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// An instance template is a global resource that is not
// bound to a zone or a region. However, you can still specify some regional
// resources in an instance template, which restricts the template to the
// region where that resource resides. For example, a custom `subnetwork`
// resource is tied to a specific region. Defaults to the region of the
// Provider if no value is given.
func (r *InstanceTemplate) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The scheduling strategy to use. More details about
// this configuration option are detailed below.
func (r *InstanceTemplate) Schedulings() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["schedulings"])
}

// The URI of the created resource.
func (r *InstanceTemplate) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

// Service account to attach to the instance. Structure is documented below.
func (r *InstanceTemplate) ServiceAccount() *pulumi.Output {
	return r.s.State["serviceAccount"]
}

// Tags to attach to the instance.
func (r *InstanceTemplate) Tags() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["tags"])
}

// The unique fingerprint of the tags.
func (r *InstanceTemplate) TagsFingerprint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tagsFingerprint"])
}

// Input properties used for looking up and filtering InstanceTemplate resources.
type InstanceTemplateState struct {
	// Whether to allow sending and receiving of
	// packets with non-matching source or destination IPs. This defaults to false.
	CanIpForward interface{}
	// A brief description of this resource.
	Description interface{}
	// Disks to attach to instances created from this template.
	// This can be specified multiple times for multiple disks. Structure is
	// documented below.
	Disks interface{}
	// List of the type and count of accelerator cards attached to the instance. Structure documented below.
	GuestAccelerators interface{}
	// A brief description to use for instances
	// created from this template.
	InstanceDescription interface{}
	// A set of key/value label pairs to assign to instances
	// created from this template,
	Labels interface{}
	// The machine type to create.
	MachineType interface{}
	// Metadata key/value pairs to make available from
	// within instances created from this template.
	Metadata interface{}
	// The unique fingerprint of the metadata.
	MetadataFingerprint interface{}
	// An alternative to using the
	// startup-script metadata key, mostly to match the compute_instance resource.
	// This replaces the startup-script metadata key on the created instance and
	// thus the two mechanisms are not allowed to be used simultaneously.
	MetadataStartupScript interface{}
	// Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
	// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
	MinCpuPlatform interface{}
	// The name of the instance template. If you leave
	// this blank, Terraform will auto-generate a unique name.
	Name interface{}
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix interface{}
	// Networks to attach to instances created from
	// this template. This can be specified multiple times for multiple networks.
	// Structure is documented below.
	NetworkInterfaces interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// An instance template is a global resource that is not
	// bound to a zone or a region. However, you can still specify some regional
	// resources in an instance template, which restricts the template to the
	// region where that resource resides. For example, a custom `subnetwork`
	// resource is tied to a specific region. Defaults to the region of the
	// Provider if no value is given.
	Region interface{}
	// The scheduling strategy to use. More details about
	// this configuration option are detailed below.
	Schedulings interface{}
	// The URI of the created resource.
	SelfLink interface{}
	// Service account to attach to the instance. Structure is documented below.
	ServiceAccount interface{}
	// Tags to attach to the instance.
	Tags interface{}
	// The unique fingerprint of the tags.
	TagsFingerprint interface{}
}

// The set of arguments for constructing a InstanceTemplate resource.
type InstanceTemplateArgs struct {
	// Whether to allow sending and receiving of
	// packets with non-matching source or destination IPs. This defaults to false.
	CanIpForward interface{}
	// A brief description of this resource.
	Description interface{}
	// Disks to attach to instances created from this template.
	// This can be specified multiple times for multiple disks. Structure is
	// documented below.
	Disks interface{}
	// List of the type and count of accelerator cards attached to the instance. Structure documented below.
	GuestAccelerators interface{}
	// A brief description to use for instances
	// created from this template.
	InstanceDescription interface{}
	// A set of key/value label pairs to assign to instances
	// created from this template,
	Labels interface{}
	// The machine type to create.
	MachineType interface{}
	// Metadata key/value pairs to make available from
	// within instances created from this template.
	Metadata interface{}
	// An alternative to using the
	// startup-script metadata key, mostly to match the compute_instance resource.
	// This replaces the startup-script metadata key on the created instance and
	// thus the two mechanisms are not allowed to be used simultaneously.
	MetadataStartupScript interface{}
	// Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
	// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
	MinCpuPlatform interface{}
	// The name of the instance template. If you leave
	// this blank, Terraform will auto-generate a unique name.
	Name interface{}
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix interface{}
	// Networks to attach to instances created from
	// this template. This can be specified multiple times for multiple networks.
	// Structure is documented below.
	NetworkInterfaces interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// An instance template is a global resource that is not
	// bound to a zone or a region. However, you can still specify some regional
	// resources in an instance template, which restricts the template to the
	// region where that resource resides. For example, a custom `subnetwork`
	// resource is tied to a specific region. Defaults to the region of the
	// Provider if no value is given.
	Region interface{}
	// The scheduling strategy to use. More details about
	// this configuration option are detailed below.
	Schedulings interface{}
	// Service account to attach to the instance. Structure is documented below.
	ServiceAccount interface{}
	// Tags to attach to the instance.
	Tags interface{}
}
