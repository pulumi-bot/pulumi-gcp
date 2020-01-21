// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package instanceFromTemplate

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/compute/InstanceFromTemplateAttachedDisk"
	"https:/github.com/pulumi/pulumi-gcp/compute/InstanceFromTemplateBootDisk"
	"https:/github.com/pulumi/pulumi-gcp/compute/InstanceFromTemplateBootDiskInitializeParams"
	"https:/github.com/pulumi/pulumi-gcp/compute/InstanceFromTemplateGuestAccelerator"
	"https:/github.com/pulumi/pulumi-gcp/compute/InstanceFromTemplateNetworkInterface"
	"https:/github.com/pulumi/pulumi-gcp/compute/InstanceFromTemplateNetworkInterfaceAccessConfig"
	"https:/github.com/pulumi/pulumi-gcp/compute/InstanceFromTemplateNetworkInterfaceAliasIpRange"
	"https:/github.com/pulumi/pulumi-gcp/compute/InstanceFromTemplateScheduling"
	"https:/github.com/pulumi/pulumi-gcp/compute/InstanceFromTemplateSchedulingNodeAffinity"
	"https:/github.com/pulumi/pulumi-gcp/compute/InstanceFromTemplateScratchDisk"
	"https:/github.com/pulumi/pulumi-gcp/compute/InstanceFromTemplateServiceAccount"
	"https:/github.com/pulumi/pulumi-gcp/compute/InstanceFromTemplateShieldedInstanceConfig"
)

// Manages a VM instance resource within GCE. For more information see
// [the official documentation](https://cloud.google.com/compute/docs/instances)
// and
// [API](https://cloud.google.com/compute/docs/reference/latest/instances).
// 
// This resource is specifically to create a compute instance from a given
// `sourceInstanceTemplate`. To create an instance without a template, use the
// `compute.Instance` resource.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_instance_from_template.html.markdown.
type InstanceFromTemplate struct {
	pulumi.CustomResourceState

	AllowStoppingForUpdate pulumi.BoolOutput `pulumi:"allowStoppingForUpdate"`
	AttachedDisks computeInstanceFromTemplateAttachedDisk.InstanceFromTemplateAttachedDiskArrayOutput `pulumi:"attachedDisks"`
	BootDisk computeInstanceFromTemplateBootDisk.InstanceFromTemplateBootDiskOutput `pulumi:"bootDisk"`
	CanIpForward pulumi.BoolOutput `pulumi:"canIpForward"`
	CpuPlatform pulumi.StringOutput `pulumi:"cpuPlatform"`
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`
	Description pulumi.StringOutput `pulumi:"description"`
	EnableDisplay pulumi.BoolOutput `pulumi:"enableDisplay"`
	GuestAccelerators computeInstanceFromTemplateGuestAccelerator.InstanceFromTemplateGuestAcceleratorArrayOutput `pulumi:"guestAccelerators"`
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	MachineType pulumi.StringOutput `pulumi:"machineType"`
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	MetadataFingerprint pulumi.StringOutput `pulumi:"metadataFingerprint"`
	MetadataStartupScript pulumi.StringOutput `pulumi:"metadataStartupScript"`
	MinCpuPlatform pulumi.StringOutput `pulumi:"minCpuPlatform"`
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	NetworkInterfaces computeInstanceFromTemplateNetworkInterface.InstanceFromTemplateNetworkInterfaceArrayOutput `pulumi:"networkInterfaces"`
	Project pulumi.StringOutput `pulumi:"project"`
	Scheduling computeInstanceFromTemplateScheduling.InstanceFromTemplateSchedulingOutput `pulumi:"scheduling"`
	ScratchDisks computeInstanceFromTemplateScratchDisk.InstanceFromTemplateScratchDiskArrayOutput `pulumi:"scratchDisks"`
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	ServiceAccount computeInstanceFromTemplateServiceAccount.InstanceFromTemplateServiceAccountOutput `pulumi:"serviceAccount"`
	ShieldedInstanceConfig computeInstanceFromTemplateShieldedInstanceConfig.InstanceFromTemplateShieldedInstanceConfigOutput `pulumi:"shieldedInstanceConfig"`
	// Name or self link of an instance
	// template to create the instance based on.
	SourceInstanceTemplate pulumi.StringOutput `pulumi:"sourceInstanceTemplate"`
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	TagsFingerprint pulumi.StringOutput `pulumi:"tagsFingerprint"`
	// The zone that the machine should be created in. If not
	// set, the provider zone is used.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceFromTemplate registers a new resource with the given unique name, arguments, and options.
func NewInstanceFromTemplate(ctx *pulumi.Context,
	name string, args *InstanceFromTemplateArgs, opts ...pulumi.ResourceOption) (*InstanceFromTemplate, error) {
	if args == nil || args.SourceInstanceTemplate == nil {
		return nil, errors.New("missing required argument 'SourceInstanceTemplate'")
	}
	if args == nil {
		args = &InstanceFromTemplateArgs{}
	}
	var resource InstanceFromTemplate
	err := ctx.RegisterResource("gcp:compute/instanceFromTemplate:InstanceFromTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceFromTemplate gets an existing InstanceFromTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceFromTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceFromTemplateState, opts ...pulumi.ResourceOption) (*InstanceFromTemplate, error) {
	var resource InstanceFromTemplate
	err := ctx.ReadResource("gcp:compute/instanceFromTemplate:InstanceFromTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceFromTemplate resources.
type instanceFromTemplateState struct {
	AllowStoppingForUpdate *bool `pulumi:"allowStoppingForUpdate"`
	AttachedDisks []computeInstanceFromTemplateAttachedDisk.InstanceFromTemplateAttachedDisk `pulumi:"attachedDisks"`
	BootDisk *computeInstanceFromTemplateBootDisk.InstanceFromTemplateBootDisk `pulumi:"bootDisk"`
	CanIpForward *bool `pulumi:"canIpForward"`
	CpuPlatform *string `pulumi:"cpuPlatform"`
	DeletionProtection *bool `pulumi:"deletionProtection"`
	Description *string `pulumi:"description"`
	EnableDisplay *bool `pulumi:"enableDisplay"`
	GuestAccelerators []computeInstanceFromTemplateGuestAccelerator.InstanceFromTemplateGuestAccelerator `pulumi:"guestAccelerators"`
	Hostname *string `pulumi:"hostname"`
	InstanceId *string `pulumi:"instanceId"`
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	Labels map[string]string `pulumi:"labels"`
	MachineType *string `pulumi:"machineType"`
	Metadata map[string]string `pulumi:"metadata"`
	MetadataFingerprint *string `pulumi:"metadataFingerprint"`
	MetadataStartupScript *string `pulumi:"metadataStartupScript"`
	MinCpuPlatform *string `pulumi:"minCpuPlatform"`
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	NetworkInterfaces []computeInstanceFromTemplateNetworkInterface.InstanceFromTemplateNetworkInterface `pulumi:"networkInterfaces"`
	Project *string `pulumi:"project"`
	Scheduling *computeInstanceFromTemplateScheduling.InstanceFromTemplateScheduling `pulumi:"scheduling"`
	ScratchDisks []computeInstanceFromTemplateScratchDisk.InstanceFromTemplateScratchDisk `pulumi:"scratchDisks"`
	SelfLink *string `pulumi:"selfLink"`
	ServiceAccount *computeInstanceFromTemplateServiceAccount.InstanceFromTemplateServiceAccount `pulumi:"serviceAccount"`
	ShieldedInstanceConfig *computeInstanceFromTemplateShieldedInstanceConfig.InstanceFromTemplateShieldedInstanceConfig `pulumi:"shieldedInstanceConfig"`
	// Name or self link of an instance
	// template to create the instance based on.
	SourceInstanceTemplate *string `pulumi:"sourceInstanceTemplate"`
	Tags []string `pulumi:"tags"`
	TagsFingerprint *string `pulumi:"tagsFingerprint"`
	// The zone that the machine should be created in. If not
	// set, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

type InstanceFromTemplateState struct {
	AllowStoppingForUpdate pulumi.BoolPtrInput
	AttachedDisks computeInstanceFromTemplateAttachedDisk.InstanceFromTemplateAttachedDiskArrayInput
	BootDisk computeInstanceFromTemplateBootDisk.InstanceFromTemplateBootDiskPtrInput
	CanIpForward pulumi.BoolPtrInput
	CpuPlatform pulumi.StringPtrInput
	DeletionProtection pulumi.BoolPtrInput
	Description pulumi.StringPtrInput
	EnableDisplay pulumi.BoolPtrInput
	GuestAccelerators computeInstanceFromTemplateGuestAccelerator.InstanceFromTemplateGuestAcceleratorArrayInput
	Hostname pulumi.StringPtrInput
	InstanceId pulumi.StringPtrInput
	LabelFingerprint pulumi.StringPtrInput
	Labels pulumi.StringMapInput
	MachineType pulumi.StringPtrInput
	Metadata pulumi.StringMapInput
	MetadataFingerprint pulumi.StringPtrInput
	MetadataStartupScript pulumi.StringPtrInput
	MinCpuPlatform pulumi.StringPtrInput
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	NetworkInterfaces computeInstanceFromTemplateNetworkInterface.InstanceFromTemplateNetworkInterfaceArrayInput
	Project pulumi.StringPtrInput
	Scheduling computeInstanceFromTemplateScheduling.InstanceFromTemplateSchedulingPtrInput
	ScratchDisks computeInstanceFromTemplateScratchDisk.InstanceFromTemplateScratchDiskArrayInput
	SelfLink pulumi.StringPtrInput
	ServiceAccount computeInstanceFromTemplateServiceAccount.InstanceFromTemplateServiceAccountPtrInput
	ShieldedInstanceConfig computeInstanceFromTemplateShieldedInstanceConfig.InstanceFromTemplateShieldedInstanceConfigPtrInput
	// Name or self link of an instance
	// template to create the instance based on.
	SourceInstanceTemplate pulumi.StringPtrInput
	Tags pulumi.StringArrayInput
	TagsFingerprint pulumi.StringPtrInput
	// The zone that the machine should be created in. If not
	// set, the provider zone is used.
	Zone pulumi.StringPtrInput
}

func (InstanceFromTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceFromTemplateState)(nil)).Elem()
}

type instanceFromTemplateArgs struct {
	AllowStoppingForUpdate *bool `pulumi:"allowStoppingForUpdate"`
	AttachedDisks []computeInstanceFromTemplateAttachedDisk.InstanceFromTemplateAttachedDisk `pulumi:"attachedDisks"`
	BootDisk *computeInstanceFromTemplateBootDisk.InstanceFromTemplateBootDisk `pulumi:"bootDisk"`
	CanIpForward *bool `pulumi:"canIpForward"`
	DeletionProtection *bool `pulumi:"deletionProtection"`
	Description *string `pulumi:"description"`
	EnableDisplay *bool `pulumi:"enableDisplay"`
	GuestAccelerators []computeInstanceFromTemplateGuestAccelerator.InstanceFromTemplateGuestAccelerator `pulumi:"guestAccelerators"`
	Hostname *string `pulumi:"hostname"`
	Labels map[string]string `pulumi:"labels"`
	MachineType *string `pulumi:"machineType"`
	Metadata map[string]string `pulumi:"metadata"`
	MetadataStartupScript *string `pulumi:"metadataStartupScript"`
	MinCpuPlatform *string `pulumi:"minCpuPlatform"`
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	NetworkInterfaces []computeInstanceFromTemplateNetworkInterface.InstanceFromTemplateNetworkInterface `pulumi:"networkInterfaces"`
	Project *string `pulumi:"project"`
	Scheduling *computeInstanceFromTemplateScheduling.InstanceFromTemplateScheduling `pulumi:"scheduling"`
	ScratchDisks []computeInstanceFromTemplateScratchDisk.InstanceFromTemplateScratchDisk `pulumi:"scratchDisks"`
	ServiceAccount *computeInstanceFromTemplateServiceAccount.InstanceFromTemplateServiceAccount `pulumi:"serviceAccount"`
	ShieldedInstanceConfig *computeInstanceFromTemplateShieldedInstanceConfig.InstanceFromTemplateShieldedInstanceConfig `pulumi:"shieldedInstanceConfig"`
	// Name or self link of an instance
	// template to create the instance based on.
	SourceInstanceTemplate string `pulumi:"sourceInstanceTemplate"`
	Tags []string `pulumi:"tags"`
	// The zone that the machine should be created in. If not
	// set, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceFromTemplate resource.
type InstanceFromTemplateArgs struct {
	AllowStoppingForUpdate pulumi.BoolPtrInput
	AttachedDisks computeInstanceFromTemplateAttachedDisk.InstanceFromTemplateAttachedDiskArrayInput
	BootDisk computeInstanceFromTemplateBootDisk.InstanceFromTemplateBootDiskPtrInput
	CanIpForward pulumi.BoolPtrInput
	DeletionProtection pulumi.BoolPtrInput
	Description pulumi.StringPtrInput
	EnableDisplay pulumi.BoolPtrInput
	GuestAccelerators computeInstanceFromTemplateGuestAccelerator.InstanceFromTemplateGuestAcceleratorArrayInput
	Hostname pulumi.StringPtrInput
	Labels pulumi.StringMapInput
	MachineType pulumi.StringPtrInput
	Metadata pulumi.StringMapInput
	MetadataStartupScript pulumi.StringPtrInput
	MinCpuPlatform pulumi.StringPtrInput
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	NetworkInterfaces computeInstanceFromTemplateNetworkInterface.InstanceFromTemplateNetworkInterfaceArrayInput
	Project pulumi.StringPtrInput
	Scheduling computeInstanceFromTemplateScheduling.InstanceFromTemplateSchedulingPtrInput
	ScratchDisks computeInstanceFromTemplateScratchDisk.InstanceFromTemplateScratchDiskArrayInput
	ServiceAccount computeInstanceFromTemplateServiceAccount.InstanceFromTemplateServiceAccountPtrInput
	ShieldedInstanceConfig computeInstanceFromTemplateShieldedInstanceConfig.InstanceFromTemplateShieldedInstanceConfigPtrInput
	// Name or self link of an instance
	// template to create the instance based on.
	SourceInstanceTemplate pulumi.StringInput
	Tags pulumi.StringArrayInput
	// The zone that the machine should be created in. If not
	// set, the provider zone is used.
	Zone pulumi.StringPtrInput
}

func (InstanceFromTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceFromTemplateArgs)(nil)).Elem()
}

