// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type InstanceFromMachineImage struct {
	pulumi.CustomResourceState

	// If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
	// stopping the instance without setting this field, the update will fail.
	AllowStoppingForUpdate pulumi.BoolOutput `pulumi:"allowStoppingForUpdate"`
	// List of disks attached to the instance
	AttachedDisks InstanceFromMachineImageAttachedDiskArrayOutput `pulumi:"attachedDisks"`
	// The boot disk for the instance.
	BootDisks InstanceFromMachineImageBootDiskArrayOutput `pulumi:"bootDisks"`
	// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
	CanIpForward pulumi.BoolOutput `pulumi:"canIpForward"`
	// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
	// to create.
	ConfidentialInstanceConfig InstanceFromMachineImageConfidentialInstanceConfigOutput `pulumi:"confidentialInstanceConfig"`
	// The CPU platform used by this instance.
	CpuPlatform pulumi.StringOutput `pulumi:"cpuPlatform"`
	// Current status of the instance.
	CurrentStatus pulumi.StringOutput `pulumi:"currentStatus"`
	// Whether deletion protection is enabled on this instance.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`
	// A brief description of the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Desired status of the instance. Either "RUNNING" or "TERMINATED".
	DesiredStatus pulumi.StringOutput `pulumi:"desiredStatus"`
	// Whether the instance has virtual displays enabled.
	EnableDisplay pulumi.BoolOutput `pulumi:"enableDisplay"`
	// List of the type and count of accelerator cards attached to the instance.
	GuestAccelerators InstanceFromMachineImageGuestAcceleratorArrayOutput `pulumi:"guestAccelerators"`
	// A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
	// labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
	// entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The server-assigned unique identifier of this instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The unique fingerprint of the labels.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// A set of key/value label pairs assigned to the instance.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The machine type to create.
	MachineType pulumi.StringOutput `pulumi:"machineType"`
	// Metadata key/value pairs made available within the instance.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The unique fingerprint of the metadata.
	MetadataFingerprint pulumi.StringOutput `pulumi:"metadataFingerprint"`
	// Metadata startup scripts made available within the instance.
	MetadataStartupScript pulumi.StringOutput `pulumi:"metadataStartupScript"`
	// The minimum CPU platform specified for the VM instance.
	MinCpuPlatform pulumi.StringOutput `pulumi:"minCpuPlatform"`
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The networks attached to the instance.
	NetworkInterfaces InstanceFromMachineImageNetworkInterfaceArrayOutput `pulumi:"networkInterfaces"`
	// The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
	// self_link nor project are provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// A list of short names or self_links of resource policies to attach to the instance. Modifying this list will cause the
	// instance to recreate. Currently a max of 1 resource policy is supported.
	ResourcePolicies pulumi.StringOutput `pulumi:"resourcePolicies"`
	// The scheduling strategy being used by the instance.
	Scheduling InstanceFromMachineImageSchedulingOutput `pulumi:"scheduling"`
	// The scratch disks attached to the instance.
	ScratchDisks InstanceFromMachineImageScratchDiskArrayOutput `pulumi:"scratchDisks"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The service account to attach to the instance.
	ServiceAccount InstanceFromMachineImageServiceAccountOutput `pulumi:"serviceAccount"`
	// The shielded vm config being used by the instance.
	ShieldedInstanceConfig InstanceFromMachineImageShieldedInstanceConfigOutput `pulumi:"shieldedInstanceConfig"`
	// Name or self link of a machine
	// image to create the instance based on.
	SourceMachineImage pulumi.StringOutput `pulumi:"sourceMachineImage"`
	// The list of tags attached to the instance.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The unique fingerprint of the tags.
	TagsFingerprint pulumi.StringOutput `pulumi:"tagsFingerprint"`
	// The zone that the machine should be created in. If not
	// set, the provider zone is used.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceFromMachineImage registers a new resource with the given unique name, arguments, and options.
func NewInstanceFromMachineImage(ctx *pulumi.Context,
	name string, args *InstanceFromMachineImageArgs, opts ...pulumi.ResourceOption) (*InstanceFromMachineImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.SourceMachineImage == nil {
		return nil, errors.New("invalid value for required argument 'SourceMachineImage'")
	}
	var resource InstanceFromMachineImage
	err := ctx.RegisterResource("gcp:compute/instanceFromMachineImage:InstanceFromMachineImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceFromMachineImage gets an existing InstanceFromMachineImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceFromMachineImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceFromMachineImageState, opts ...pulumi.ResourceOption) (*InstanceFromMachineImage, error) {
	var resource InstanceFromMachineImage
	err := ctx.ReadResource("gcp:compute/instanceFromMachineImage:InstanceFromMachineImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceFromMachineImage resources.
type instanceFromMachineImageState struct {
	// If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
	// stopping the instance without setting this field, the update will fail.
	AllowStoppingForUpdate *bool `pulumi:"allowStoppingForUpdate"`
	// List of disks attached to the instance
	AttachedDisks []InstanceFromMachineImageAttachedDisk `pulumi:"attachedDisks"`
	// The boot disk for the instance.
	BootDisks []InstanceFromMachineImageBootDisk `pulumi:"bootDisks"`
	// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
	CanIpForward *bool `pulumi:"canIpForward"`
	// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
	// to create.
	ConfidentialInstanceConfig *InstanceFromMachineImageConfidentialInstanceConfig `pulumi:"confidentialInstanceConfig"`
	// The CPU platform used by this instance.
	CpuPlatform *string `pulumi:"cpuPlatform"`
	// Current status of the instance.
	CurrentStatus *string `pulumi:"currentStatus"`
	// Whether deletion protection is enabled on this instance.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// A brief description of the resource.
	Description *string `pulumi:"description"`
	// Desired status of the instance. Either "RUNNING" or "TERMINATED".
	DesiredStatus *string `pulumi:"desiredStatus"`
	// Whether the instance has virtual displays enabled.
	EnableDisplay *bool `pulumi:"enableDisplay"`
	// List of the type and count of accelerator cards attached to the instance.
	GuestAccelerators []InstanceFromMachineImageGuestAccelerator `pulumi:"guestAccelerators"`
	// A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
	// labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
	// entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	Hostname *string `pulumi:"hostname"`
	// The server-assigned unique identifier of this instance.
	InstanceId *string `pulumi:"instanceId"`
	// The unique fingerprint of the labels.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// A set of key/value label pairs assigned to the instance.
	Labels map[string]string `pulumi:"labels"`
	// The machine type to create.
	MachineType *string `pulumi:"machineType"`
	// Metadata key/value pairs made available within the instance.
	Metadata map[string]string `pulumi:"metadata"`
	// The unique fingerprint of the metadata.
	MetadataFingerprint *string `pulumi:"metadataFingerprint"`
	// Metadata startup scripts made available within the instance.
	MetadataStartupScript *string `pulumi:"metadataStartupScript"`
	// The minimum CPU platform specified for the VM instance.
	MinCpuPlatform *string `pulumi:"minCpuPlatform"`
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The networks attached to the instance.
	NetworkInterfaces []InstanceFromMachineImageNetworkInterface `pulumi:"networkInterfaces"`
	// The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
	// self_link nor project are provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A list of short names or self_links of resource policies to attach to the instance. Modifying this list will cause the
	// instance to recreate. Currently a max of 1 resource policy is supported.
	ResourcePolicies *string `pulumi:"resourcePolicies"`
	// The scheduling strategy being used by the instance.
	Scheduling *InstanceFromMachineImageScheduling `pulumi:"scheduling"`
	// The scratch disks attached to the instance.
	ScratchDisks []InstanceFromMachineImageScratchDisk `pulumi:"scratchDisks"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The service account to attach to the instance.
	ServiceAccount *InstanceFromMachineImageServiceAccount `pulumi:"serviceAccount"`
	// The shielded vm config being used by the instance.
	ShieldedInstanceConfig *InstanceFromMachineImageShieldedInstanceConfig `pulumi:"shieldedInstanceConfig"`
	// Name or self link of a machine
	// image to create the instance based on.
	SourceMachineImage *string `pulumi:"sourceMachineImage"`
	// The list of tags attached to the instance.
	Tags []string `pulumi:"tags"`
	// The unique fingerprint of the tags.
	TagsFingerprint *string `pulumi:"tagsFingerprint"`
	// The zone that the machine should be created in. If not
	// set, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

type InstanceFromMachineImageState struct {
	// If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
	// stopping the instance without setting this field, the update will fail.
	AllowStoppingForUpdate pulumi.BoolPtrInput
	// List of disks attached to the instance
	AttachedDisks InstanceFromMachineImageAttachedDiskArrayInput
	// The boot disk for the instance.
	BootDisks InstanceFromMachineImageBootDiskArrayInput
	// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
	CanIpForward pulumi.BoolPtrInput
	// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
	// to create.
	ConfidentialInstanceConfig InstanceFromMachineImageConfidentialInstanceConfigPtrInput
	// The CPU platform used by this instance.
	CpuPlatform pulumi.StringPtrInput
	// Current status of the instance.
	CurrentStatus pulumi.StringPtrInput
	// Whether deletion protection is enabled on this instance.
	DeletionProtection pulumi.BoolPtrInput
	// A brief description of the resource.
	Description pulumi.StringPtrInput
	// Desired status of the instance. Either "RUNNING" or "TERMINATED".
	DesiredStatus pulumi.StringPtrInput
	// Whether the instance has virtual displays enabled.
	EnableDisplay pulumi.BoolPtrInput
	// List of the type and count of accelerator cards attached to the instance.
	GuestAccelerators InstanceFromMachineImageGuestAcceleratorArrayInput
	// A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
	// labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
	// entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	Hostname pulumi.StringPtrInput
	// The server-assigned unique identifier of this instance.
	InstanceId pulumi.StringPtrInput
	// The unique fingerprint of the labels.
	LabelFingerprint pulumi.StringPtrInput
	// A set of key/value label pairs assigned to the instance.
	Labels pulumi.StringMapInput
	// The machine type to create.
	MachineType pulumi.StringPtrInput
	// Metadata key/value pairs made available within the instance.
	Metadata pulumi.StringMapInput
	// The unique fingerprint of the metadata.
	MetadataFingerprint pulumi.StringPtrInput
	// Metadata startup scripts made available within the instance.
	MetadataStartupScript pulumi.StringPtrInput
	// The minimum CPU platform specified for the VM instance.
	MinCpuPlatform pulumi.StringPtrInput
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The networks attached to the instance.
	NetworkInterfaces InstanceFromMachineImageNetworkInterfaceArrayInput
	// The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
	// self_link nor project are provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A list of short names or self_links of resource policies to attach to the instance. Modifying this list will cause the
	// instance to recreate. Currently a max of 1 resource policy is supported.
	ResourcePolicies pulumi.StringPtrInput
	// The scheduling strategy being used by the instance.
	Scheduling InstanceFromMachineImageSchedulingPtrInput
	// The scratch disks attached to the instance.
	ScratchDisks InstanceFromMachineImageScratchDiskArrayInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The service account to attach to the instance.
	ServiceAccount InstanceFromMachineImageServiceAccountPtrInput
	// The shielded vm config being used by the instance.
	ShieldedInstanceConfig InstanceFromMachineImageShieldedInstanceConfigPtrInput
	// Name or self link of a machine
	// image to create the instance based on.
	SourceMachineImage pulumi.StringPtrInput
	// The list of tags attached to the instance.
	Tags pulumi.StringArrayInput
	// The unique fingerprint of the tags.
	TagsFingerprint pulumi.StringPtrInput
	// The zone that the machine should be created in. If not
	// set, the provider zone is used.
	Zone pulumi.StringPtrInput
}

func (InstanceFromMachineImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceFromMachineImageState)(nil)).Elem()
}

type instanceFromMachineImageArgs struct {
	// If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
	// stopping the instance without setting this field, the update will fail.
	AllowStoppingForUpdate *bool `pulumi:"allowStoppingForUpdate"`
	// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
	CanIpForward *bool `pulumi:"canIpForward"`
	// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
	// to create.
	ConfidentialInstanceConfig *InstanceFromMachineImageConfidentialInstanceConfig `pulumi:"confidentialInstanceConfig"`
	// Whether deletion protection is enabled on this instance.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// A brief description of the resource.
	Description *string `pulumi:"description"`
	// Desired status of the instance. Either "RUNNING" or "TERMINATED".
	DesiredStatus *string `pulumi:"desiredStatus"`
	// Whether the instance has virtual displays enabled.
	EnableDisplay *bool `pulumi:"enableDisplay"`
	// List of the type and count of accelerator cards attached to the instance.
	GuestAccelerators []InstanceFromMachineImageGuestAccelerator `pulumi:"guestAccelerators"`
	// A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
	// labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
	// entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	Hostname *string `pulumi:"hostname"`
	// A set of key/value label pairs assigned to the instance.
	Labels map[string]string `pulumi:"labels"`
	// The machine type to create.
	MachineType *string `pulumi:"machineType"`
	// Metadata key/value pairs made available within the instance.
	Metadata map[string]string `pulumi:"metadata"`
	// Metadata startup scripts made available within the instance.
	MetadataStartupScript *string `pulumi:"metadataStartupScript"`
	// The minimum CPU platform specified for the VM instance.
	MinCpuPlatform *string `pulumi:"minCpuPlatform"`
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The networks attached to the instance.
	NetworkInterfaces []InstanceFromMachineImageNetworkInterface `pulumi:"networkInterfaces"`
	// The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
	// self_link nor project are provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A list of short names or self_links of resource policies to attach to the instance. Modifying this list will cause the
	// instance to recreate. Currently a max of 1 resource policy is supported.
	ResourcePolicies *string `pulumi:"resourcePolicies"`
	// The scheduling strategy being used by the instance.
	Scheduling *InstanceFromMachineImageScheduling `pulumi:"scheduling"`
	// The service account to attach to the instance.
	ServiceAccount *InstanceFromMachineImageServiceAccount `pulumi:"serviceAccount"`
	// The shielded vm config being used by the instance.
	ShieldedInstanceConfig *InstanceFromMachineImageShieldedInstanceConfig `pulumi:"shieldedInstanceConfig"`
	// Name or self link of a machine
	// image to create the instance based on.
	SourceMachineImage string `pulumi:"sourceMachineImage"`
	// The list of tags attached to the instance.
	Tags []string `pulumi:"tags"`
	// The zone that the machine should be created in. If not
	// set, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceFromMachineImage resource.
type InstanceFromMachineImageArgs struct {
	// If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
	// stopping the instance without setting this field, the update will fail.
	AllowStoppingForUpdate pulumi.BoolPtrInput
	// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
	CanIpForward pulumi.BoolPtrInput
	// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
	// to create.
	ConfidentialInstanceConfig InstanceFromMachineImageConfidentialInstanceConfigPtrInput
	// Whether deletion protection is enabled on this instance.
	DeletionProtection pulumi.BoolPtrInput
	// A brief description of the resource.
	Description pulumi.StringPtrInput
	// Desired status of the instance. Either "RUNNING" or "TERMINATED".
	DesiredStatus pulumi.StringPtrInput
	// Whether the instance has virtual displays enabled.
	EnableDisplay pulumi.BoolPtrInput
	// List of the type and count of accelerator cards attached to the instance.
	GuestAccelerators InstanceFromMachineImageGuestAcceleratorArrayInput
	// A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
	// labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
	// entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	Hostname pulumi.StringPtrInput
	// A set of key/value label pairs assigned to the instance.
	Labels pulumi.StringMapInput
	// The machine type to create.
	MachineType pulumi.StringPtrInput
	// Metadata key/value pairs made available within the instance.
	Metadata pulumi.StringMapInput
	// Metadata startup scripts made available within the instance.
	MetadataStartupScript pulumi.StringPtrInput
	// The minimum CPU platform specified for the VM instance.
	MinCpuPlatform pulumi.StringPtrInput
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The networks attached to the instance.
	NetworkInterfaces InstanceFromMachineImageNetworkInterfaceArrayInput
	// The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
	// self_link nor project are provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A list of short names or self_links of resource policies to attach to the instance. Modifying this list will cause the
	// instance to recreate. Currently a max of 1 resource policy is supported.
	ResourcePolicies pulumi.StringPtrInput
	// The scheduling strategy being used by the instance.
	Scheduling InstanceFromMachineImageSchedulingPtrInput
	// The service account to attach to the instance.
	ServiceAccount InstanceFromMachineImageServiceAccountPtrInput
	// The shielded vm config being used by the instance.
	ShieldedInstanceConfig InstanceFromMachineImageShieldedInstanceConfigPtrInput
	// Name or self link of a machine
	// image to create the instance based on.
	SourceMachineImage pulumi.StringInput
	// The list of tags attached to the instance.
	Tags pulumi.StringArrayInput
	// The zone that the machine should be created in. If not
	// set, the provider zone is used.
	Zone pulumi.StringPtrInput
}

func (InstanceFromMachineImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceFromMachineImageArgs)(nil)).Elem()
}
