// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a VM instance resource within GCE. For more information see
// [the official documentation](https://cloud.google.com/compute/docs/instances)
// and
// [API](https://cloud.google.com/compute/docs/reference/latest/instances).
//
// This resource is specifically to create a compute instance from a given
// `sourceInstanceTemplate`. To create an instance without a template, use the
// `compute.Instance` resource.
type InstanceFromTemplate struct {
	pulumi.CustomResourceState

	AllowStoppingForUpdate pulumi.BoolOutput                               `pulumi:"allowStoppingForUpdate"`
	AttachedDisks          InstanceFromTemplateAttachedDiskArrayOutput     `pulumi:"attachedDisks"`
	BootDisk               InstanceFromTemplateBootDiskOutput              `pulumi:"bootDisk"`
	CanIpForward           pulumi.BoolOutput                               `pulumi:"canIpForward"`
	CpuPlatform            pulumi.StringOutput                             `pulumi:"cpuPlatform"`
	CurrentStatus          pulumi.StringOutput                             `pulumi:"currentStatus"`
	DeletionProtection     pulumi.BoolOutput                               `pulumi:"deletionProtection"`
	Description            pulumi.StringOutput                             `pulumi:"description"`
	DesiredStatus          pulumi.StringOutput                             `pulumi:"desiredStatus"`
	EnableDisplay          pulumi.BoolOutput                               `pulumi:"enableDisplay"`
	GuestAccelerators      InstanceFromTemplateGuestAcceleratorArrayOutput `pulumi:"guestAccelerators"`
	Hostname               pulumi.StringOutput                             `pulumi:"hostname"`
	InstanceId             pulumi.StringOutput                             `pulumi:"instanceId"`
	LabelFingerprint       pulumi.StringOutput                             `pulumi:"labelFingerprint"`
	Labels                 pulumi.StringMapOutput                          `pulumi:"labels"`
	MachineType            pulumi.StringOutput                             `pulumi:"machineType"`
	Metadata               pulumi.StringMapOutput                          `pulumi:"metadata"`
	MetadataFingerprint    pulumi.StringOutput                             `pulumi:"metadataFingerprint"`
	MetadataStartupScript  pulumi.StringOutput                             `pulumi:"metadataStartupScript"`
	MinCpuPlatform         pulumi.StringOutput                             `pulumi:"minCpuPlatform"`
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name                   pulumi.StringOutput                              `pulumi:"name"`
	NetworkInterfaces      InstanceFromTemplateNetworkInterfaceArrayOutput  `pulumi:"networkInterfaces"`
	Project                pulumi.StringOutput                              `pulumi:"project"`
	Scheduling             InstanceFromTemplateSchedulingOutput             `pulumi:"scheduling"`
	ScratchDisks           InstanceFromTemplateScratchDiskArrayOutput       `pulumi:"scratchDisks"`
	SelfLink               pulumi.StringOutput                              `pulumi:"selfLink"`
	ServiceAccount         InstanceFromTemplateServiceAccountOutput         `pulumi:"serviceAccount"`
	ShieldedInstanceConfig InstanceFromTemplateShieldedInstanceConfigOutput `pulumi:"shieldedInstanceConfig"`
	// Name or self link of an instance
	// template to create the instance based on.
	SourceInstanceTemplate pulumi.StringOutput      `pulumi:"sourceInstanceTemplate"`
	Tags                   pulumi.StringArrayOutput `pulumi:"tags"`
	TagsFingerprint        pulumi.StringOutput      `pulumi:"tagsFingerprint"`
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
	AllowStoppingForUpdate *bool                                  `pulumi:"allowStoppingForUpdate"`
	AttachedDisks          []InstanceFromTemplateAttachedDisk     `pulumi:"attachedDisks"`
	BootDisk               *InstanceFromTemplateBootDisk          `pulumi:"bootDisk"`
	CanIpForward           *bool                                  `pulumi:"canIpForward"`
	CpuPlatform            *string                                `pulumi:"cpuPlatform"`
	CurrentStatus          *string                                `pulumi:"currentStatus"`
	DeletionProtection     *bool                                  `pulumi:"deletionProtection"`
	Description            *string                                `pulumi:"description"`
	DesiredStatus          *string                                `pulumi:"desiredStatus"`
	EnableDisplay          *bool                                  `pulumi:"enableDisplay"`
	GuestAccelerators      []InstanceFromTemplateGuestAccelerator `pulumi:"guestAccelerators"`
	Hostname               *string                                `pulumi:"hostname"`
	InstanceId             *string                                `pulumi:"instanceId"`
	LabelFingerprint       *string                                `pulumi:"labelFingerprint"`
	Labels                 map[string]string                      `pulumi:"labels"`
	MachineType            *string                                `pulumi:"machineType"`
	Metadata               map[string]string                      `pulumi:"metadata"`
	MetadataFingerprint    *string                                `pulumi:"metadataFingerprint"`
	MetadataStartupScript  *string                                `pulumi:"metadataStartupScript"`
	MinCpuPlatform         *string                                `pulumi:"minCpuPlatform"`
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name                   *string                                     `pulumi:"name"`
	NetworkInterfaces      []InstanceFromTemplateNetworkInterface      `pulumi:"networkInterfaces"`
	Project                *string                                     `pulumi:"project"`
	Scheduling             *InstanceFromTemplateScheduling             `pulumi:"scheduling"`
	ScratchDisks           []InstanceFromTemplateScratchDisk           `pulumi:"scratchDisks"`
	SelfLink               *string                                     `pulumi:"selfLink"`
	ServiceAccount         *InstanceFromTemplateServiceAccount         `pulumi:"serviceAccount"`
	ShieldedInstanceConfig *InstanceFromTemplateShieldedInstanceConfig `pulumi:"shieldedInstanceConfig"`
	// Name or self link of an instance
	// template to create the instance based on.
	SourceInstanceTemplate *string  `pulumi:"sourceInstanceTemplate"`
	Tags                   []string `pulumi:"tags"`
	TagsFingerprint        *string  `pulumi:"tagsFingerprint"`
	// The zone that the machine should be created in. If not
	// set, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

type InstanceFromTemplateState struct {
	AllowStoppingForUpdate pulumi.BoolPtrInput
	AttachedDisks          InstanceFromTemplateAttachedDiskArrayInput
	BootDisk               InstanceFromTemplateBootDiskPtrInput
	CanIpForward           pulumi.BoolPtrInput
	CpuPlatform            pulumi.StringPtrInput
	CurrentStatus          pulumi.StringPtrInput
	DeletionProtection     pulumi.BoolPtrInput
	Description            pulumi.StringPtrInput
	DesiredStatus          pulumi.StringPtrInput
	EnableDisplay          pulumi.BoolPtrInput
	GuestAccelerators      InstanceFromTemplateGuestAcceleratorArrayInput
	Hostname               pulumi.StringPtrInput
	InstanceId             pulumi.StringPtrInput
	LabelFingerprint       pulumi.StringPtrInput
	Labels                 pulumi.StringMapInput
	MachineType            pulumi.StringPtrInput
	Metadata               pulumi.StringMapInput
	MetadataFingerprint    pulumi.StringPtrInput
	MetadataStartupScript  pulumi.StringPtrInput
	MinCpuPlatform         pulumi.StringPtrInput
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name                   pulumi.StringPtrInput
	NetworkInterfaces      InstanceFromTemplateNetworkInterfaceArrayInput
	Project                pulumi.StringPtrInput
	Scheduling             InstanceFromTemplateSchedulingPtrInput
	ScratchDisks           InstanceFromTemplateScratchDiskArrayInput
	SelfLink               pulumi.StringPtrInput
	ServiceAccount         InstanceFromTemplateServiceAccountPtrInput
	ShieldedInstanceConfig InstanceFromTemplateShieldedInstanceConfigPtrInput
	// Name or self link of an instance
	// template to create the instance based on.
	SourceInstanceTemplate pulumi.StringPtrInput
	Tags                   pulumi.StringArrayInput
	TagsFingerprint        pulumi.StringPtrInput
	// The zone that the machine should be created in. If not
	// set, the provider zone is used.
	Zone pulumi.StringPtrInput
}

func (InstanceFromTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceFromTemplateState)(nil)).Elem()
}

type instanceFromTemplateArgs struct {
	AllowStoppingForUpdate *bool                                      `pulumi:"allowStoppingForUpdate"`
	AttachedDisks          []InstanceFromTemplateAttachedDiskArgs     `pulumi:"attachedDisks"`
	BootDisk               *InstanceFromTemplateBootDiskArgs          `pulumi:"bootDisk"`
	CanIpForward           *bool                                      `pulumi:"canIpForward"`
	DeletionProtection     *bool                                      `pulumi:"deletionProtection"`
	Description            *string                                    `pulumi:"description"`
	DesiredStatus          *string                                    `pulumi:"desiredStatus"`
	EnableDisplay          *bool                                      `pulumi:"enableDisplay"`
	GuestAccelerators      []InstanceFromTemplateGuestAcceleratorArgs `pulumi:"guestAccelerators"`
	Hostname               *string                                    `pulumi:"hostname"`
	Labels                 map[string]string                          `pulumi:"labels"`
	MachineType            *string                                    `pulumi:"machineType"`
	Metadata               map[string]string                          `pulumi:"metadata"`
	MetadataStartupScript  *string                                    `pulumi:"metadataStartupScript"`
	MinCpuPlatform         *string                                    `pulumi:"minCpuPlatform"`
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name                   *string                                         `pulumi:"name"`
	NetworkInterfaces      []InstanceFromTemplateNetworkInterfaceArgs      `pulumi:"networkInterfaces"`
	Project                *string                                         `pulumi:"project"`
	Scheduling             *InstanceFromTemplateSchedulingArgs             `pulumi:"scheduling"`
	ScratchDisks           []InstanceFromTemplateScratchDiskArgs           `pulumi:"scratchDisks"`
	ServiceAccount         *InstanceFromTemplateServiceAccountArgs         `pulumi:"serviceAccount"`
	ShieldedInstanceConfig *InstanceFromTemplateShieldedInstanceConfigArgs `pulumi:"shieldedInstanceConfig"`
	// Name or self link of an instance
	// template to create the instance based on.
	SourceInstanceTemplate string   `pulumi:"sourceInstanceTemplate"`
	Tags                   []string `pulumi:"tags"`
	// The zone that the machine should be created in. If not
	// set, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceFromTemplate resource.
type InstanceFromTemplateArgs struct {
	AllowStoppingForUpdate pulumi.BoolPtrInput
	AttachedDisks          InstanceFromTemplateAttachedDiskArgsArrayInput
	BootDisk               InstanceFromTemplateBootDiskArgsPtrInput
	CanIpForward           pulumi.BoolPtrInput
	DeletionProtection     pulumi.BoolPtrInput
	Description            pulumi.StringPtrInput
	DesiredStatus          pulumi.StringPtrInput
	EnableDisplay          pulumi.BoolPtrInput
	GuestAccelerators      InstanceFromTemplateGuestAcceleratorArgsArrayInput
	Hostname               pulumi.StringPtrInput
	Labels                 pulumi.StringMapInput
	MachineType            pulumi.StringPtrInput
	Metadata               pulumi.StringMapInput
	MetadataStartupScript  pulumi.StringPtrInput
	MinCpuPlatform         pulumi.StringPtrInput
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name                   pulumi.StringPtrInput
	NetworkInterfaces      InstanceFromTemplateNetworkInterfaceArgsArrayInput
	Project                pulumi.StringPtrInput
	Scheduling             InstanceFromTemplateSchedulingArgsPtrInput
	ScratchDisks           InstanceFromTemplateScratchDiskArgsArrayInput
	ServiceAccount         InstanceFromTemplateServiceAccountArgsPtrInput
	ShieldedInstanceConfig InstanceFromTemplateShieldedInstanceConfigArgsPtrInput
	// Name or self link of an instance
	// template to create the instance based on.
	SourceInstanceTemplate pulumi.StringInput
	Tags                   pulumi.StringArrayInput
	// The zone that the machine should be created in. If not
	// set, the provider zone is used.
	Zone pulumi.StringPtrInput
}

func (InstanceFromTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceFromTemplateArgs)(nil)).Elem()
}
