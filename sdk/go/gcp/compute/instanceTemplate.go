// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a VM instance template resource within GCE. For more information see
// [the official documentation](https://cloud.google.com/compute/docs/instance-templates)
// and
// [API](https://cloud.google.com/compute/docs/reference/latest/instanceTemplates).
type InstanceTemplate struct {
	pulumi.CustomResourceState

	// Whether to allow sending and receiving of
	// packets with non-matching source or destination IPs. This defaults to false.
	CanIpForward pulumi.BoolPtrOutput `pulumi:"canIpForward"`
	// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
	// to create.
	ConfidentialInstanceConfig InstanceTemplateConfidentialInstanceConfigOutput `pulumi:"confidentialInstanceConfig"`
	// A brief description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Disks to attach to instances created from this template.
	// This can be specified multiple times for multiple disks. Structure is
	// documented below.
	Disks InstanceTemplateDiskArrayOutput `pulumi:"disks"`
	// Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
	// **Note**: `allowStoppingForUpdate` must be set to true in order to update this field.
	EnableDisplay pulumi.BoolPtrOutput `pulumi:"enableDisplay"`
	// List of the type and count of accelerator cards attached to the instance. Structure documented below.
	GuestAccelerators InstanceTemplateGuestAcceleratorArrayOutput `pulumi:"guestAccelerators"`
	// A brief description to use for instances
	// created from this template.
	InstanceDescription pulumi.StringPtrOutput `pulumi:"instanceDescription"`
	// A set of key/value label pairs to assign to instances
	// created from this template,
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The machine type to create.
	MachineType pulumi.StringOutput `pulumi:"machineType"`
	// Metadata key/value pairs to make available from
	// within instances created from this template.
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// The unique fingerprint of the metadata.
	MetadataFingerprint pulumi.StringOutput `pulumi:"metadataFingerprint"`
	// An alternative to using the
	// startup-script metadata key, mostly to match the computeInstance resource.
	// This replaces the startup-script metadata key on the created instance and
	// thus the two mechanisms are not allowed to be used simultaneously.
	MetadataStartupScript pulumi.StringPtrOutput `pulumi:"metadataStartupScript"`
	// Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
	// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
	MinCpuPlatform pulumi.StringPtrOutput `pulumi:"minCpuPlatform"`
	// The name of the instance template. If you leave
	// this blank, the provider will auto-generate a unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// Networks to attach to instances created from
	// this template. This can be specified multiple times for multiple networks.
	// Structure is documented below.
	NetworkInterfaces InstanceTemplateNetworkInterfaceArrayOutput `pulumi:"networkInterfaces"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// An instance template is a global resource that is not
	// bound to a zone or a region. However, you can still specify some regional
	// resources in an instance template, which restricts the template to the
	// region where that resource resides. For example, a custom `subnetwork`
	// resource is tied to a specific region. Defaults to the region of the
	// Provider if no value is given.
	Region pulumi.StringOutput `pulumi:"region"`
	// The scheduling strategy to use. More details about
	// this configuration option are detailed below.
	Scheduling InstanceTemplateSchedulingOutput `pulumi:"scheduling"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Service account to attach to the instance. Structure is documented below.
	ServiceAccount InstanceTemplateServiceAccountPtrOutput `pulumi:"serviceAccount"`
	// Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
	// **Note**: `shieldedInstanceConfig` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
	ShieldedInstanceConfig InstanceTemplateShieldedInstanceConfigOutput `pulumi:"shieldedInstanceConfig"`
	// Tags to attach to the instance.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The unique fingerprint of the tags.
	TagsFingerprint pulumi.StringOutput `pulumi:"tagsFingerprint"`
}

// NewInstanceTemplate registers a new resource with the given unique name, arguments, and options.
func NewInstanceTemplate(ctx *pulumi.Context,
	name string, args *InstanceTemplateArgs, opts ...pulumi.ResourceOption) (*InstanceTemplate, error) {
	if args == nil || args.Disks == nil {
		return nil, errors.New("missing required argument 'Disks'")
	}
	if args == nil || args.MachineType == nil {
		return nil, errors.New("missing required argument 'MachineType'")
	}
	if args == nil {
		args = &InstanceTemplateArgs{}
	}
	var resource InstanceTemplate
	err := ctx.RegisterResource("gcp:compute/instanceTemplate:InstanceTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceTemplate gets an existing InstanceTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceTemplateState, opts ...pulumi.ResourceOption) (*InstanceTemplate, error) {
	var resource InstanceTemplate
	err := ctx.ReadResource("gcp:compute/instanceTemplate:InstanceTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceTemplate resources.
type instanceTemplateState struct {
	// Whether to allow sending and receiving of
	// packets with non-matching source or destination IPs. This defaults to false.
	CanIpForward *bool `pulumi:"canIpForward"`
	// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
	// to create.
	ConfidentialInstanceConfig *InstanceTemplateConfidentialInstanceConfig `pulumi:"confidentialInstanceConfig"`
	// A brief description of this resource.
	Description *string `pulumi:"description"`
	// Disks to attach to instances created from this template.
	// This can be specified multiple times for multiple disks. Structure is
	// documented below.
	Disks []InstanceTemplateDisk `pulumi:"disks"`
	// Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
	// **Note**: `allowStoppingForUpdate` must be set to true in order to update this field.
	EnableDisplay *bool `pulumi:"enableDisplay"`
	// List of the type and count of accelerator cards attached to the instance. Structure documented below.
	GuestAccelerators []InstanceTemplateGuestAccelerator `pulumi:"guestAccelerators"`
	// A brief description to use for instances
	// created from this template.
	InstanceDescription *string `pulumi:"instanceDescription"`
	// A set of key/value label pairs to assign to instances
	// created from this template,
	Labels map[string]string `pulumi:"labels"`
	// The machine type to create.
	MachineType *string `pulumi:"machineType"`
	// Metadata key/value pairs to make available from
	// within instances created from this template.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The unique fingerprint of the metadata.
	MetadataFingerprint *string `pulumi:"metadataFingerprint"`
	// An alternative to using the
	// startup-script metadata key, mostly to match the computeInstance resource.
	// This replaces the startup-script metadata key on the created instance and
	// thus the two mechanisms are not allowed to be used simultaneously.
	MetadataStartupScript *string `pulumi:"metadataStartupScript"`
	// Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
	// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
	MinCpuPlatform *string `pulumi:"minCpuPlatform"`
	// The name of the instance template. If you leave
	// this blank, the provider will auto-generate a unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Networks to attach to instances created from
	// this template. This can be specified multiple times for multiple networks.
	// Structure is documented below.
	NetworkInterfaces []InstanceTemplateNetworkInterface `pulumi:"networkInterfaces"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// An instance template is a global resource that is not
	// bound to a zone or a region. However, you can still specify some regional
	// resources in an instance template, which restricts the template to the
	// region where that resource resides. For example, a custom `subnetwork`
	// resource is tied to a specific region. Defaults to the region of the
	// Provider if no value is given.
	Region *string `pulumi:"region"`
	// The scheduling strategy to use. More details about
	// this configuration option are detailed below.
	Scheduling *InstanceTemplateScheduling `pulumi:"scheduling"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Service account to attach to the instance. Structure is documented below.
	ServiceAccount *InstanceTemplateServiceAccount `pulumi:"serviceAccount"`
	// Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
	// **Note**: `shieldedInstanceConfig` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
	ShieldedInstanceConfig *InstanceTemplateShieldedInstanceConfig `pulumi:"shieldedInstanceConfig"`
	// Tags to attach to the instance.
	Tags []string `pulumi:"tags"`
	// The unique fingerprint of the tags.
	TagsFingerprint *string `pulumi:"tagsFingerprint"`
}

type InstanceTemplateState struct {
	// Whether to allow sending and receiving of
	// packets with non-matching source or destination IPs. This defaults to false.
	CanIpForward pulumi.BoolPtrInput
	// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
	// to create.
	ConfidentialInstanceConfig InstanceTemplateConfidentialInstanceConfigPtrInput
	// A brief description of this resource.
	Description pulumi.StringPtrInput
	// Disks to attach to instances created from this template.
	// This can be specified multiple times for multiple disks. Structure is
	// documented below.
	Disks InstanceTemplateDiskArrayInput
	// Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
	// **Note**: `allowStoppingForUpdate` must be set to true in order to update this field.
	EnableDisplay pulumi.BoolPtrInput
	// List of the type and count of accelerator cards attached to the instance. Structure documented below.
	GuestAccelerators InstanceTemplateGuestAcceleratorArrayInput
	// A brief description to use for instances
	// created from this template.
	InstanceDescription pulumi.StringPtrInput
	// A set of key/value label pairs to assign to instances
	// created from this template,
	Labels pulumi.StringMapInput
	// The machine type to create.
	MachineType pulumi.StringPtrInput
	// Metadata key/value pairs to make available from
	// within instances created from this template.
	Metadata pulumi.MapInput
	// The unique fingerprint of the metadata.
	MetadataFingerprint pulumi.StringPtrInput
	// An alternative to using the
	// startup-script metadata key, mostly to match the computeInstance resource.
	// This replaces the startup-script metadata key on the created instance and
	// thus the two mechanisms are not allowed to be used simultaneously.
	MetadataStartupScript pulumi.StringPtrInput
	// Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
	// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
	MinCpuPlatform pulumi.StringPtrInput
	// The name of the instance template. If you leave
	// this blank, the provider will auto-generate a unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Networks to attach to instances created from
	// this template. This can be specified multiple times for multiple networks.
	// Structure is documented below.
	NetworkInterfaces InstanceTemplateNetworkInterfaceArrayInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// An instance template is a global resource that is not
	// bound to a zone or a region. However, you can still specify some regional
	// resources in an instance template, which restricts the template to the
	// region where that resource resides. For example, a custom `subnetwork`
	// resource is tied to a specific region. Defaults to the region of the
	// Provider if no value is given.
	Region pulumi.StringPtrInput
	// The scheduling strategy to use. More details about
	// this configuration option are detailed below.
	Scheduling InstanceTemplateSchedulingPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Service account to attach to the instance. Structure is documented below.
	ServiceAccount InstanceTemplateServiceAccountPtrInput
	// Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
	// **Note**: `shieldedInstanceConfig` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
	ShieldedInstanceConfig InstanceTemplateShieldedInstanceConfigPtrInput
	// Tags to attach to the instance.
	Tags pulumi.StringArrayInput
	// The unique fingerprint of the tags.
	TagsFingerprint pulumi.StringPtrInput
}

func (InstanceTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceTemplateState)(nil)).Elem()
}

type instanceTemplateArgs struct {
	// Whether to allow sending and receiving of
	// packets with non-matching source or destination IPs. This defaults to false.
	CanIpForward *bool `pulumi:"canIpForward"`
	// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
	// to create.
	ConfidentialInstanceConfig *InstanceTemplateConfidentialInstanceConfig `pulumi:"confidentialInstanceConfig"`
	// A brief description of this resource.
	Description *string `pulumi:"description"`
	// Disks to attach to instances created from this template.
	// This can be specified multiple times for multiple disks. Structure is
	// documented below.
	Disks []InstanceTemplateDisk `pulumi:"disks"`
	// Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
	// **Note**: `allowStoppingForUpdate` must be set to true in order to update this field.
	EnableDisplay *bool `pulumi:"enableDisplay"`
	// List of the type and count of accelerator cards attached to the instance. Structure documented below.
	GuestAccelerators []InstanceTemplateGuestAccelerator `pulumi:"guestAccelerators"`
	// A brief description to use for instances
	// created from this template.
	InstanceDescription *string `pulumi:"instanceDescription"`
	// A set of key/value label pairs to assign to instances
	// created from this template,
	Labels map[string]string `pulumi:"labels"`
	// The machine type to create.
	MachineType string `pulumi:"machineType"`
	// Metadata key/value pairs to make available from
	// within instances created from this template.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// An alternative to using the
	// startup-script metadata key, mostly to match the computeInstance resource.
	// This replaces the startup-script metadata key on the created instance and
	// thus the two mechanisms are not allowed to be used simultaneously.
	MetadataStartupScript *string `pulumi:"metadataStartupScript"`
	// Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
	// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
	MinCpuPlatform *string `pulumi:"minCpuPlatform"`
	// The name of the instance template. If you leave
	// this blank, the provider will auto-generate a unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Networks to attach to instances created from
	// this template. This can be specified multiple times for multiple networks.
	// Structure is documented below.
	NetworkInterfaces []InstanceTemplateNetworkInterface `pulumi:"networkInterfaces"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// An instance template is a global resource that is not
	// bound to a zone or a region. However, you can still specify some regional
	// resources in an instance template, which restricts the template to the
	// region where that resource resides. For example, a custom `subnetwork`
	// resource is tied to a specific region. Defaults to the region of the
	// Provider if no value is given.
	Region *string `pulumi:"region"`
	// The scheduling strategy to use. More details about
	// this configuration option are detailed below.
	Scheduling *InstanceTemplateScheduling `pulumi:"scheduling"`
	// Service account to attach to the instance. Structure is documented below.
	ServiceAccount *InstanceTemplateServiceAccount `pulumi:"serviceAccount"`
	// Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
	// **Note**: `shieldedInstanceConfig` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
	ShieldedInstanceConfig *InstanceTemplateShieldedInstanceConfig `pulumi:"shieldedInstanceConfig"`
	// Tags to attach to the instance.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a InstanceTemplate resource.
type InstanceTemplateArgs struct {
	// Whether to allow sending and receiving of
	// packets with non-matching source or destination IPs. This defaults to false.
	CanIpForward pulumi.BoolPtrInput
	// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
	// to create.
	ConfidentialInstanceConfig InstanceTemplateConfidentialInstanceConfigPtrInput
	// A brief description of this resource.
	Description pulumi.StringPtrInput
	// Disks to attach to instances created from this template.
	// This can be specified multiple times for multiple disks. Structure is
	// documented below.
	Disks InstanceTemplateDiskArrayInput
	// Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
	// **Note**: `allowStoppingForUpdate` must be set to true in order to update this field.
	EnableDisplay pulumi.BoolPtrInput
	// List of the type and count of accelerator cards attached to the instance. Structure documented below.
	GuestAccelerators InstanceTemplateGuestAcceleratorArrayInput
	// A brief description to use for instances
	// created from this template.
	InstanceDescription pulumi.StringPtrInput
	// A set of key/value label pairs to assign to instances
	// created from this template,
	Labels pulumi.StringMapInput
	// The machine type to create.
	MachineType pulumi.StringInput
	// Metadata key/value pairs to make available from
	// within instances created from this template.
	Metadata pulumi.MapInput
	// An alternative to using the
	// startup-script metadata key, mostly to match the computeInstance resource.
	// This replaces the startup-script metadata key on the created instance and
	// thus the two mechanisms are not allowed to be used simultaneously.
	MetadataStartupScript pulumi.StringPtrInput
	// Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
	// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
	MinCpuPlatform pulumi.StringPtrInput
	// The name of the instance template. If you leave
	// this blank, the provider will auto-generate a unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Networks to attach to instances created from
	// this template. This can be specified multiple times for multiple networks.
	// Structure is documented below.
	NetworkInterfaces InstanceTemplateNetworkInterfaceArrayInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// An instance template is a global resource that is not
	// bound to a zone or a region. However, you can still specify some regional
	// resources in an instance template, which restricts the template to the
	// region where that resource resides. For example, a custom `subnetwork`
	// resource is tied to a specific region. Defaults to the region of the
	// Provider if no value is given.
	Region pulumi.StringPtrInput
	// The scheduling strategy to use. More details about
	// this configuration option are detailed below.
	Scheduling InstanceTemplateSchedulingPtrInput
	// Service account to attach to the instance. Structure is documented below.
	ServiceAccount InstanceTemplateServiceAccountPtrInput
	// Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
	// **Note**: `shieldedInstanceConfig` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
	ShieldedInstanceConfig InstanceTemplateShieldedInstanceConfigPtrInput
	// Tags to attach to the instance.
	Tags pulumi.StringArrayInput
}

func (InstanceTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceTemplateArgs)(nil)).Elem()
}

type InstanceTemplateInput interface {
	pulumi.Input

	ToInstanceTemplateOutput() InstanceTemplateOutput
	ToInstanceTemplateOutputWithContext(ctx context.Context) InstanceTemplateOutput
}

func (InstanceTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTemplate)(nil)).Elem()
}

func (i InstanceTemplate) ToInstanceTemplateOutput() InstanceTemplateOutput {
	return i.ToInstanceTemplateOutputWithContext(context.Background())
}

func (i InstanceTemplate) ToInstanceTemplateOutputWithContext(ctx context.Context) InstanceTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTemplateOutput)
}

type InstanceTemplateOutput struct {
	*pulumi.OutputState
}

func (InstanceTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTemplateOutput)(nil)).Elem()
}

func (o InstanceTemplateOutput) ToInstanceTemplateOutput() InstanceTemplateOutput {
	return o
}

func (o InstanceTemplateOutput) ToInstanceTemplateOutputWithContext(ctx context.Context) InstanceTemplateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InstanceTemplateOutput{})
}
