// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a VM instance resource within GCE. For more information see
// [the official documentation](https://cloud.google.com/compute/docs/instances)
// and
// [API](https://cloud.google.com/compute/docs/reference/latest/instances).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewInstance(ctx, "_default", &compute.InstanceArgs{
// 			BootDisk: &compute.InstanceBootDiskArgs{
// 				InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
// 					Image: pulumi.String("debian-cloud/debian-9"),
// 				},
// 			},
// 			MachineType: pulumi.String("e2-medium"),
// 			Metadata: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 			MetadataStartupScript: pulumi.String("echo hi > /test.txt"),
// 			NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
// 				&compute.InstanceNetworkInterfaceArgs{
// 					AccessConfigs: compute.InstanceNetworkInterfaceAccessConfigArray{
// 						nil,
// 					},
// 					Network: pulumi.String("default"),
// 				},
// 			},
// 			ScratchDisks: compute.InstanceScratchDiskArray{
// 				&compute.InstanceScratchDiskArgs{
// 					Interface: pulumi.String("SCSI"),
// 				},
// 			},
// 			ServiceAccount: &compute.InstanceServiceAccountArgs{
// 				Scopes: pulumi.StringArray{
// 					pulumi.String("userinfo-email"),
// 					pulumi.String("compute-ro"),
// 					pulumi.String("storage-ro"),
// 				},
// 			},
// 			Tags: pulumi.StringArray{
// 				pulumi.String("foo"),
// 				pulumi.String("bar"),
// 			},
// 			Zone: pulumi.String("us-central1-a"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Instances can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/instance:Instance default projects/{{project}}/zones/{{zone}}/instances/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/instance:Instance default {{project}}/{{zone}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/instance:Instance default {{name}}
// ```
//
//  [custom-vm-types]https://cloud.google.com/dataproc/docs/concepts/compute/custom-machine-types [network-tier]https://cloud.google.com/network-tiers/docs/overview [extended-custom-vm-type]https://cloud.google.com/compute/docs/instances/creating-instance-with-custom-machine-type#extendedmemory
type Instance struct {
	pulumi.CustomResourceState

	// If true, allows this prvider to stop the instance to update its properties.
	// If you try to update a property that requires stopping the instance without setting this field, the update will fail.
	AllowStoppingForUpdate pulumi.BoolPtrOutput `pulumi:"allowStoppingForUpdate"`
	// Additional disks to attach to the instance. Can be repeated multiple times for multiple disks. Structure is documented below.
	AttachedDisks InstanceAttachedDiskArrayOutput `pulumi:"attachedDisks"`
	// The boot disk for the instance.
	// Structure is documented below.
	BootDisk InstanceBootDiskOutput `pulumi:"bootDisk"`
	// Whether to allow sending and receiving of
	// packets with non-matching source or destination IPs.
	// This defaults to false.
	CanIpForward pulumi.BoolPtrOutput `pulumi:"canIpForward"`
	// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
	// to create.
	ConfidentialInstanceConfig InstanceConfidentialInstanceConfigOutput `pulumi:"confidentialInstanceConfig"`
	// The CPU platform used by this instance.
	CpuPlatform pulumi.StringOutput `pulumi:"cpuPlatform"`
	// Current status of the instance.
	CurrentStatus pulumi.StringOutput `pulumi:"currentStatus"`
	// Enable deletion protection on this instance. Defaults to false.
	// **Note:** you must disable deletion protection before removing the resource (e.g., via `pulumi destroy`), or the instance cannot be deleted and the provider run will not complete successfully.
	DeletionProtection pulumi.BoolPtrOutput `pulumi:"deletionProtection"`
	// A brief description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Desired status of the instance. Either
	// `"RUNNING"` or `"TERMINATED"`.
	DesiredStatus pulumi.StringPtrOutput `pulumi:"desiredStatus"`
	// Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	EnableDisplay pulumi.BoolPtrOutput `pulumi:"enableDisplay"`
	// List of the type and count of accelerator cards attached to the instance. Structure documented below.
	// **Note:** GPU accelerators can only be used with `onHostMaintenance` option set to TERMINATE.
	GuestAccelerators InstanceGuestAcceleratorArrayOutput `pulumi:"guestAccelerators"`
	// A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid.
	// Valid format is a series of labels 1-63 characters long matching the regular expression `a-z`, concatenated with periods.
	// The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	Hostname pulumi.StringPtrOutput `pulumi:"hostname"`
	// The server-assigned unique identifier of this instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The unique fingerprint of the labels.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// A map of key/value label pairs to assign to the instance.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The machine type to create.
	MachineType pulumi.StringOutput `pulumi:"machineType"`
	// Metadata key/value pairs to make available from
	// within the instance. Ssh keys attached in the Cloud Console will be removed.
	// Add them to your config in order to keep them attached to your instance.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The unique fingerprint of the metadata.
	MetadataFingerprint pulumi.StringOutput `pulumi:"metadataFingerprint"`
	// An alternative to using the
	// startup-script metadata key, except this one forces the instance to be
	// recreated (thus re-running the script) if it is changed. This replaces the
	// startup-script metadata key on the created instance and thus the two
	// mechanisms are not allowed to be used simultaneously.  Users are free to use
	// either mechanism - the only distinction is that this separate attribute
	// willl cause a recreate on modification.  On import, `metadataStartupScript`
	// will be set, but `metadata.startup-script` will not - if you choose to use the
	// other mechanism, you will see a diff immediately after import, which will cause a
	// destroy/recreate operation.  You may want to modify your state file manually
	// using `pulumi stack` commands, depending on your use case.
	MetadataStartupScript pulumi.StringPtrOutput `pulumi:"metadataStartupScript"`
	// Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as
	// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	MinCpuPlatform pulumi.StringOutput `pulumi:"minCpuPlatform"`
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Networks to attach to the instance. This can
	// be specified multiple times. Structure is documented below.
	NetworkInterfaces InstanceNetworkInterfaceArrayOutput `pulumi:"networkInterfaces"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// -- A list of short names or selfLinks of resource policies to attach to the instance. Modifying this list will cause the instance to recreate. Currently a max of 1 resource policy is supported.
	ResourcePolicies pulumi.StringPtrOutput `pulumi:"resourcePolicies"`
	// The scheduling strategy to use. More details about
	// this configuration option are detailed below.
	Scheduling InstanceSchedulingOutput `pulumi:"scheduling"`
	// Scratch disks to attach to the instance. This can be
	// specified multiple times for multiple scratch disks. Structure is documented below.
	ScratchDisks InstanceScratchDiskArrayOutput `pulumi:"scratchDisks"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Service account to attach to the instance.
	// Structure is documented below.
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	ServiceAccount InstanceServiceAccountPtrOutput `pulumi:"serviceAccount"`
	// Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
	// **Note**: `shieldedInstanceConfig` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	ShieldedInstanceConfig InstanceShieldedInstanceConfigOutput `pulumi:"shieldedInstanceConfig"`
	// A list of network tags to attach to the instance.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The unique fingerprint of the tags.
	TagsFingerprint pulumi.StringOutput `pulumi:"tagsFingerprint"`
	// The zone that the machine should be created in.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BootDisk == nil {
		return nil, errors.New("invalid value for required argument 'BootDisk'")
	}
	if args.MachineType == nil {
		return nil, errors.New("invalid value for required argument 'MachineType'")
	}
	if args.NetworkInterfaces == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInterfaces'")
	}
	var resource Instance
	err := ctx.RegisterResource("gcp:compute/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("gcp:compute/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// If true, allows this prvider to stop the instance to update its properties.
	// If you try to update a property that requires stopping the instance without setting this field, the update will fail.
	AllowStoppingForUpdate *bool `pulumi:"allowStoppingForUpdate"`
	// Additional disks to attach to the instance. Can be repeated multiple times for multiple disks. Structure is documented below.
	AttachedDisks []InstanceAttachedDisk `pulumi:"attachedDisks"`
	// The boot disk for the instance.
	// Structure is documented below.
	BootDisk *InstanceBootDisk `pulumi:"bootDisk"`
	// Whether to allow sending and receiving of
	// packets with non-matching source or destination IPs.
	// This defaults to false.
	CanIpForward *bool `pulumi:"canIpForward"`
	// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
	// to create.
	ConfidentialInstanceConfig *InstanceConfidentialInstanceConfig `pulumi:"confidentialInstanceConfig"`
	// The CPU platform used by this instance.
	CpuPlatform *string `pulumi:"cpuPlatform"`
	// Current status of the instance.
	CurrentStatus *string `pulumi:"currentStatus"`
	// Enable deletion protection on this instance. Defaults to false.
	// **Note:** you must disable deletion protection before removing the resource (e.g., via `pulumi destroy`), or the instance cannot be deleted and the provider run will not complete successfully.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// A brief description of this resource.
	Description *string `pulumi:"description"`
	// Desired status of the instance. Either
	// `"RUNNING"` or `"TERMINATED"`.
	DesiredStatus *string `pulumi:"desiredStatus"`
	// Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	EnableDisplay *bool `pulumi:"enableDisplay"`
	// List of the type and count of accelerator cards attached to the instance. Structure documented below.
	// **Note:** GPU accelerators can only be used with `onHostMaintenance` option set to TERMINATE.
	GuestAccelerators []InstanceGuestAccelerator `pulumi:"guestAccelerators"`
	// A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid.
	// Valid format is a series of labels 1-63 characters long matching the regular expression `a-z`, concatenated with periods.
	// The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	Hostname *string `pulumi:"hostname"`
	// The server-assigned unique identifier of this instance.
	InstanceId *string `pulumi:"instanceId"`
	// The unique fingerprint of the labels.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// A map of key/value label pairs to assign to the instance.
	Labels map[string]string `pulumi:"labels"`
	// The machine type to create.
	MachineType *string `pulumi:"machineType"`
	// Metadata key/value pairs to make available from
	// within the instance. Ssh keys attached in the Cloud Console will be removed.
	// Add them to your config in order to keep them attached to your instance.
	Metadata map[string]string `pulumi:"metadata"`
	// The unique fingerprint of the metadata.
	MetadataFingerprint *string `pulumi:"metadataFingerprint"`
	// An alternative to using the
	// startup-script metadata key, except this one forces the instance to be
	// recreated (thus re-running the script) if it is changed. This replaces the
	// startup-script metadata key on the created instance and thus the two
	// mechanisms are not allowed to be used simultaneously.  Users are free to use
	// either mechanism - the only distinction is that this separate attribute
	// willl cause a recreate on modification.  On import, `metadataStartupScript`
	// will be set, but `metadata.startup-script` will not - if you choose to use the
	// other mechanism, you will see a diff immediately after import, which will cause a
	// destroy/recreate operation.  You may want to modify your state file manually
	// using `pulumi stack` commands, depending on your use case.
	MetadataStartupScript *string `pulumi:"metadataStartupScript"`
	// Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as
	// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	MinCpuPlatform *string `pulumi:"minCpuPlatform"`
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Networks to attach to the instance. This can
	// be specified multiple times. Structure is documented below.
	NetworkInterfaces []InstanceNetworkInterface `pulumi:"networkInterfaces"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// -- A list of short names or selfLinks of resource policies to attach to the instance. Modifying this list will cause the instance to recreate. Currently a max of 1 resource policy is supported.
	ResourcePolicies *string `pulumi:"resourcePolicies"`
	// The scheduling strategy to use. More details about
	// this configuration option are detailed below.
	Scheduling *InstanceScheduling `pulumi:"scheduling"`
	// Scratch disks to attach to the instance. This can be
	// specified multiple times for multiple scratch disks. Structure is documented below.
	ScratchDisks []InstanceScratchDisk `pulumi:"scratchDisks"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Service account to attach to the instance.
	// Structure is documented below.
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	ServiceAccount *InstanceServiceAccount `pulumi:"serviceAccount"`
	// Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
	// **Note**: `shieldedInstanceConfig` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	ShieldedInstanceConfig *InstanceShieldedInstanceConfig `pulumi:"shieldedInstanceConfig"`
	// A list of network tags to attach to the instance.
	Tags []string `pulumi:"tags"`
	// The unique fingerprint of the tags.
	TagsFingerprint *string `pulumi:"tagsFingerprint"`
	// The zone that the machine should be created in.
	Zone *string `pulumi:"zone"`
}

type InstanceState struct {
	// If true, allows this prvider to stop the instance to update its properties.
	// If you try to update a property that requires stopping the instance without setting this field, the update will fail.
	AllowStoppingForUpdate pulumi.BoolPtrInput
	// Additional disks to attach to the instance. Can be repeated multiple times for multiple disks. Structure is documented below.
	AttachedDisks InstanceAttachedDiskArrayInput
	// The boot disk for the instance.
	// Structure is documented below.
	BootDisk InstanceBootDiskPtrInput
	// Whether to allow sending and receiving of
	// packets with non-matching source or destination IPs.
	// This defaults to false.
	CanIpForward pulumi.BoolPtrInput
	// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
	// to create.
	ConfidentialInstanceConfig InstanceConfidentialInstanceConfigPtrInput
	// The CPU platform used by this instance.
	CpuPlatform pulumi.StringPtrInput
	// Current status of the instance.
	CurrentStatus pulumi.StringPtrInput
	// Enable deletion protection on this instance. Defaults to false.
	// **Note:** you must disable deletion protection before removing the resource (e.g., via `pulumi destroy`), or the instance cannot be deleted and the provider run will not complete successfully.
	DeletionProtection pulumi.BoolPtrInput
	// A brief description of this resource.
	Description pulumi.StringPtrInput
	// Desired status of the instance. Either
	// `"RUNNING"` or `"TERMINATED"`.
	DesiredStatus pulumi.StringPtrInput
	// Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	EnableDisplay pulumi.BoolPtrInput
	// List of the type and count of accelerator cards attached to the instance. Structure documented below.
	// **Note:** GPU accelerators can only be used with `onHostMaintenance` option set to TERMINATE.
	GuestAccelerators InstanceGuestAcceleratorArrayInput
	// A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid.
	// Valid format is a series of labels 1-63 characters long matching the regular expression `a-z`, concatenated with periods.
	// The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	Hostname pulumi.StringPtrInput
	// The server-assigned unique identifier of this instance.
	InstanceId pulumi.StringPtrInput
	// The unique fingerprint of the labels.
	LabelFingerprint pulumi.StringPtrInput
	// A map of key/value label pairs to assign to the instance.
	Labels pulumi.StringMapInput
	// The machine type to create.
	MachineType pulumi.StringPtrInput
	// Metadata key/value pairs to make available from
	// within the instance. Ssh keys attached in the Cloud Console will be removed.
	// Add them to your config in order to keep them attached to your instance.
	Metadata pulumi.StringMapInput
	// The unique fingerprint of the metadata.
	MetadataFingerprint pulumi.StringPtrInput
	// An alternative to using the
	// startup-script metadata key, except this one forces the instance to be
	// recreated (thus re-running the script) if it is changed. This replaces the
	// startup-script metadata key on the created instance and thus the two
	// mechanisms are not allowed to be used simultaneously.  Users are free to use
	// either mechanism - the only distinction is that this separate attribute
	// willl cause a recreate on modification.  On import, `metadataStartupScript`
	// will be set, but `metadata.startup-script` will not - if you choose to use the
	// other mechanism, you will see a diff immediately after import, which will cause a
	// destroy/recreate operation.  You may want to modify your state file manually
	// using `pulumi stack` commands, depending on your use case.
	MetadataStartupScript pulumi.StringPtrInput
	// Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as
	// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	MinCpuPlatform pulumi.StringPtrInput
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Networks to attach to the instance. This can
	// be specified multiple times. Structure is documented below.
	NetworkInterfaces InstanceNetworkInterfaceArrayInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// -- A list of short names or selfLinks of resource policies to attach to the instance. Modifying this list will cause the instance to recreate. Currently a max of 1 resource policy is supported.
	ResourcePolicies pulumi.StringPtrInput
	// The scheduling strategy to use. More details about
	// this configuration option are detailed below.
	Scheduling InstanceSchedulingPtrInput
	// Scratch disks to attach to the instance. This can be
	// specified multiple times for multiple scratch disks. Structure is documented below.
	ScratchDisks InstanceScratchDiskArrayInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Service account to attach to the instance.
	// Structure is documented below.
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	ServiceAccount InstanceServiceAccountPtrInput
	// Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
	// **Note**: `shieldedInstanceConfig` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	ShieldedInstanceConfig InstanceShieldedInstanceConfigPtrInput
	// A list of network tags to attach to the instance.
	Tags pulumi.StringArrayInput
	// The unique fingerprint of the tags.
	TagsFingerprint pulumi.StringPtrInput
	// The zone that the machine should be created in.
	Zone pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// If true, allows this prvider to stop the instance to update its properties.
	// If you try to update a property that requires stopping the instance without setting this field, the update will fail.
	AllowStoppingForUpdate *bool `pulumi:"allowStoppingForUpdate"`
	// Additional disks to attach to the instance. Can be repeated multiple times for multiple disks. Structure is documented below.
	AttachedDisks []InstanceAttachedDisk `pulumi:"attachedDisks"`
	// The boot disk for the instance.
	// Structure is documented below.
	BootDisk InstanceBootDisk `pulumi:"bootDisk"`
	// Whether to allow sending and receiving of
	// packets with non-matching source or destination IPs.
	// This defaults to false.
	CanIpForward *bool `pulumi:"canIpForward"`
	// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
	// to create.
	ConfidentialInstanceConfig *InstanceConfidentialInstanceConfig `pulumi:"confidentialInstanceConfig"`
	// Enable deletion protection on this instance. Defaults to false.
	// **Note:** you must disable deletion protection before removing the resource (e.g., via `pulumi destroy`), or the instance cannot be deleted and the provider run will not complete successfully.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// A brief description of this resource.
	Description *string `pulumi:"description"`
	// Desired status of the instance. Either
	// `"RUNNING"` or `"TERMINATED"`.
	DesiredStatus *string `pulumi:"desiredStatus"`
	// Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	EnableDisplay *bool `pulumi:"enableDisplay"`
	// List of the type and count of accelerator cards attached to the instance. Structure documented below.
	// **Note:** GPU accelerators can only be used with `onHostMaintenance` option set to TERMINATE.
	GuestAccelerators []InstanceGuestAccelerator `pulumi:"guestAccelerators"`
	// A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid.
	// Valid format is a series of labels 1-63 characters long matching the regular expression `a-z`, concatenated with periods.
	// The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	Hostname *string `pulumi:"hostname"`
	// A map of key/value label pairs to assign to the instance.
	Labels map[string]string `pulumi:"labels"`
	// The machine type to create.
	MachineType string `pulumi:"machineType"`
	// Metadata key/value pairs to make available from
	// within the instance. Ssh keys attached in the Cloud Console will be removed.
	// Add them to your config in order to keep them attached to your instance.
	Metadata map[string]string `pulumi:"metadata"`
	// An alternative to using the
	// startup-script metadata key, except this one forces the instance to be
	// recreated (thus re-running the script) if it is changed. This replaces the
	// startup-script metadata key on the created instance and thus the two
	// mechanisms are not allowed to be used simultaneously.  Users are free to use
	// either mechanism - the only distinction is that this separate attribute
	// willl cause a recreate on modification.  On import, `metadataStartupScript`
	// will be set, but `metadata.startup-script` will not - if you choose to use the
	// other mechanism, you will see a diff immediately after import, which will cause a
	// destroy/recreate operation.  You may want to modify your state file manually
	// using `pulumi stack` commands, depending on your use case.
	MetadataStartupScript *string `pulumi:"metadataStartupScript"`
	// Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as
	// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	MinCpuPlatform *string `pulumi:"minCpuPlatform"`
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Networks to attach to the instance. This can
	// be specified multiple times. Structure is documented below.
	NetworkInterfaces []InstanceNetworkInterface `pulumi:"networkInterfaces"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// -- A list of short names or selfLinks of resource policies to attach to the instance. Modifying this list will cause the instance to recreate. Currently a max of 1 resource policy is supported.
	ResourcePolicies *string `pulumi:"resourcePolicies"`
	// The scheduling strategy to use. More details about
	// this configuration option are detailed below.
	Scheduling *InstanceScheduling `pulumi:"scheduling"`
	// Scratch disks to attach to the instance. This can be
	// specified multiple times for multiple scratch disks. Structure is documented below.
	ScratchDisks []InstanceScratchDisk `pulumi:"scratchDisks"`
	// Service account to attach to the instance.
	// Structure is documented below.
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	ServiceAccount *InstanceServiceAccount `pulumi:"serviceAccount"`
	// Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
	// **Note**: `shieldedInstanceConfig` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	ShieldedInstanceConfig *InstanceShieldedInstanceConfig `pulumi:"shieldedInstanceConfig"`
	// A list of network tags to attach to the instance.
	Tags []string `pulumi:"tags"`
	// The zone that the machine should be created in.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// If true, allows this prvider to stop the instance to update its properties.
	// If you try to update a property that requires stopping the instance without setting this field, the update will fail.
	AllowStoppingForUpdate pulumi.BoolPtrInput
	// Additional disks to attach to the instance. Can be repeated multiple times for multiple disks. Structure is documented below.
	AttachedDisks InstanceAttachedDiskArrayInput
	// The boot disk for the instance.
	// Structure is documented below.
	BootDisk InstanceBootDiskInput
	// Whether to allow sending and receiving of
	// packets with non-matching source or destination IPs.
	// This defaults to false.
	CanIpForward pulumi.BoolPtrInput
	// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
	// to create.
	ConfidentialInstanceConfig InstanceConfidentialInstanceConfigPtrInput
	// Enable deletion protection on this instance. Defaults to false.
	// **Note:** you must disable deletion protection before removing the resource (e.g., via `pulumi destroy`), or the instance cannot be deleted and the provider run will not complete successfully.
	DeletionProtection pulumi.BoolPtrInput
	// A brief description of this resource.
	Description pulumi.StringPtrInput
	// Desired status of the instance. Either
	// `"RUNNING"` or `"TERMINATED"`.
	DesiredStatus pulumi.StringPtrInput
	// Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	EnableDisplay pulumi.BoolPtrInput
	// List of the type and count of accelerator cards attached to the instance. Structure documented below.
	// **Note:** GPU accelerators can only be used with `onHostMaintenance` option set to TERMINATE.
	GuestAccelerators InstanceGuestAcceleratorArrayInput
	// A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid.
	// Valid format is a series of labels 1-63 characters long matching the regular expression `a-z`, concatenated with periods.
	// The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	Hostname pulumi.StringPtrInput
	// A map of key/value label pairs to assign to the instance.
	Labels pulumi.StringMapInput
	// The machine type to create.
	MachineType pulumi.StringInput
	// Metadata key/value pairs to make available from
	// within the instance. Ssh keys attached in the Cloud Console will be removed.
	// Add them to your config in order to keep them attached to your instance.
	Metadata pulumi.StringMapInput
	// An alternative to using the
	// startup-script metadata key, except this one forces the instance to be
	// recreated (thus re-running the script) if it is changed. This replaces the
	// startup-script metadata key on the created instance and thus the two
	// mechanisms are not allowed to be used simultaneously.  Users are free to use
	// either mechanism - the only distinction is that this separate attribute
	// willl cause a recreate on modification.  On import, `metadataStartupScript`
	// will be set, but `metadata.startup-script` will not - if you choose to use the
	// other mechanism, you will see a diff immediately after import, which will cause a
	// destroy/recreate operation.  You may want to modify your state file manually
	// using `pulumi stack` commands, depending on your use case.
	MetadataStartupScript pulumi.StringPtrInput
	// Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as
	// `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	MinCpuPlatform pulumi.StringPtrInput
	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Networks to attach to the instance. This can
	// be specified multiple times. Structure is documented below.
	NetworkInterfaces InstanceNetworkInterfaceArrayInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// -- A list of short names or selfLinks of resource policies to attach to the instance. Modifying this list will cause the instance to recreate. Currently a max of 1 resource policy is supported.
	ResourcePolicies pulumi.StringPtrInput
	// The scheduling strategy to use. More details about
	// this configuration option are detailed below.
	Scheduling InstanceSchedulingPtrInput
	// Scratch disks to attach to the instance. This can be
	// specified multiple times for multiple scratch disks. Structure is documented below.
	ScratchDisks InstanceScratchDiskArrayInput
	// Service account to attach to the instance.
	// Structure is documented below.
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	ServiceAccount InstanceServiceAccountPtrInput
	// Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
	// **Note**: `shieldedInstanceConfig` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
	// **Note**: `allowStoppingForUpdate` must be set to true or your instance must have a `desiredStatus` of `TERMINATED` in order to update this field.
	ShieldedInstanceConfig InstanceShieldedInstanceConfigPtrInput
	// A list of network tags to attach to the instance.
	Tags pulumi.StringArrayInput
	// The zone that the machine should be created in.
	Zone pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (Instance) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil)).Elem()
}

func (i Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

type InstanceOutput struct {
	*pulumi.OutputState
}

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceOutput)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InstanceOutput{})
}
