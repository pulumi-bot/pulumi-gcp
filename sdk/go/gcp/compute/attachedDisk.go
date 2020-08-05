// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Persistent disks can be attached to a compute instance using the `attachedDisk`
// section within the compute instance configuration.
// However there may be situations where managing the attached disks via the compute
// instance config isn't preferable or possible, such as attaching dynamic
// numbers of disks using the `count` variable.
//
// To get more information about attaching disks, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instances/attachDisk)
// * How-to Guides
//     * [Adding a persistent disk](https://cloud.google.com/compute/docs/disks/add-persistent-disk)
//
// **Note:** When using `compute.AttachedDisk` you **must** use `lifecycle.ignore_changes = ["attachedDisk"]` on the `compute.Instance` resource that has the disks attached. Otherwise the two resources will fight for control of the attached disk block.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultInstance, err := compute.NewInstance(ctx, "defaultInstance", &compute.InstanceArgs{
// 			MachineType: pulumi.String("n1-standard-1"),
// 			Zone:        pulumi.String("us-west1-a"),
// 			BootDisk: &compute.InstanceBootDiskArgs{
// 				InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
// 					Image: pulumi.String("debian-cloud/debian-9"),
// 				},
// 			},
// 			NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
// 				&compute.InstanceNetworkInterfaceArgs{
// 					Network: pulumi.String("default"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewAttachedDisk(ctx, "defaultAttachedDisk", &compute.AttachedDiskArgs{
// 			Disk:     pulumi.Any(google_compute_disk.Default.Id),
// 			Instance: defaultInstance.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AttachedDisk struct {
	pulumi.CustomResourceState

	// Specifies a unique device name of your choice that is
	// reflected into the /dev/disk/by-id/google-* tree of a Linux operating
	// system running within the instance. This name can be used to
	// reference the device for mounting, resizing, and so on, from within
	// the instance.
	DeviceName pulumi.StringOutput `pulumi:"deviceName"`
	// `name` or `selfLink` of the disk that will be attached.
	Disk pulumi.StringOutput `pulumi:"disk"`
	// `name` or `selfLink` of the compute instance that the disk will be attached to.
	// If the `selfLink` is provided then `zone` and `project` are extracted from the
	// self link. If only the name is used then `zone` and `project` must be defined
	// as properties on the resource or provider.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The mode in which to attach this disk, either READ_WRITE or
	// READ_ONLY. If not specified, the default is to attach the disk in
	// READ_WRITE mode.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The project that the referenced compute instance is a part of. If `instance` is referenced by its
	// `selfLink` the project defined in the link will take precedence.
	Project pulumi.StringOutput `pulumi:"project"`
	// The zone that the referenced compute instance is located within. If `instance` is referenced by its
	// `selfLink` the zone defined in the link will take precedence.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewAttachedDisk registers a new resource with the given unique name, arguments, and options.
func NewAttachedDisk(ctx *pulumi.Context,
	name string, args *AttachedDiskArgs, opts ...pulumi.ResourceOption) (*AttachedDisk, error) {
	if args == nil || args.Disk == nil {
		return nil, errors.New("missing required argument 'Disk'")
	}
	if args == nil || args.Instance == nil {
		return nil, errors.New("missing required argument 'Instance'")
	}
	if args == nil {
		args = &AttachedDiskArgs{}
	}
	var resource AttachedDisk
	err := ctx.RegisterResource("gcp:compute/attachedDisk:AttachedDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttachedDisk gets an existing AttachedDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttachedDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachedDiskState, opts ...pulumi.ResourceOption) (*AttachedDisk, error) {
	var resource AttachedDisk
	err := ctx.ReadResource("gcp:compute/attachedDisk:AttachedDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttachedDisk resources.
type attachedDiskState struct {
	// Specifies a unique device name of your choice that is
	// reflected into the /dev/disk/by-id/google-* tree of a Linux operating
	// system running within the instance. This name can be used to
	// reference the device for mounting, resizing, and so on, from within
	// the instance.
	DeviceName *string `pulumi:"deviceName"`
	// `name` or `selfLink` of the disk that will be attached.
	Disk *string `pulumi:"disk"`
	// `name` or `selfLink` of the compute instance that the disk will be attached to.
	// If the `selfLink` is provided then `zone` and `project` are extracted from the
	// self link. If only the name is used then `zone` and `project` must be defined
	// as properties on the resource or provider.
	Instance *string `pulumi:"instance"`
	// The mode in which to attach this disk, either READ_WRITE or
	// READ_ONLY. If not specified, the default is to attach the disk in
	// READ_WRITE mode.
	Mode *string `pulumi:"mode"`
	// The project that the referenced compute instance is a part of. If `instance` is referenced by its
	// `selfLink` the project defined in the link will take precedence.
	Project *string `pulumi:"project"`
	// The zone that the referenced compute instance is located within. If `instance` is referenced by its
	// `selfLink` the zone defined in the link will take precedence.
	Zone *string `pulumi:"zone"`
}

type AttachedDiskState struct {
	// Specifies a unique device name of your choice that is
	// reflected into the /dev/disk/by-id/google-* tree of a Linux operating
	// system running within the instance. This name can be used to
	// reference the device for mounting, resizing, and so on, from within
	// the instance.
	DeviceName pulumi.StringPtrInput
	// `name` or `selfLink` of the disk that will be attached.
	Disk pulumi.StringPtrInput
	// `name` or `selfLink` of the compute instance that the disk will be attached to.
	// If the `selfLink` is provided then `zone` and `project` are extracted from the
	// self link. If only the name is used then `zone` and `project` must be defined
	// as properties on the resource or provider.
	Instance pulumi.StringPtrInput
	// The mode in which to attach this disk, either READ_WRITE or
	// READ_ONLY. If not specified, the default is to attach the disk in
	// READ_WRITE mode.
	Mode pulumi.StringPtrInput
	// The project that the referenced compute instance is a part of. If `instance` is referenced by its
	// `selfLink` the project defined in the link will take precedence.
	Project pulumi.StringPtrInput
	// The zone that the referenced compute instance is located within. If `instance` is referenced by its
	// `selfLink` the zone defined in the link will take precedence.
	Zone pulumi.StringPtrInput
}

func (AttachedDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDiskState)(nil)).Elem()
}

type attachedDiskArgs struct {
	// Specifies a unique device name of your choice that is
	// reflected into the /dev/disk/by-id/google-* tree of a Linux operating
	// system running within the instance. This name can be used to
	// reference the device for mounting, resizing, and so on, from within
	// the instance.
	DeviceName *string `pulumi:"deviceName"`
	// `name` or `selfLink` of the disk that will be attached.
	Disk string `pulumi:"disk"`
	// `name` or `selfLink` of the compute instance that the disk will be attached to.
	// If the `selfLink` is provided then `zone` and `project` are extracted from the
	// self link. If only the name is used then `zone` and `project` must be defined
	// as properties on the resource or provider.
	Instance string `pulumi:"instance"`
	// The mode in which to attach this disk, either READ_WRITE or
	// READ_ONLY. If not specified, the default is to attach the disk in
	// READ_WRITE mode.
	Mode *string `pulumi:"mode"`
	// The project that the referenced compute instance is a part of. If `instance` is referenced by its
	// `selfLink` the project defined in the link will take precedence.
	Project *string `pulumi:"project"`
	// The zone that the referenced compute instance is located within. If `instance` is referenced by its
	// `selfLink` the zone defined in the link will take precedence.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a AttachedDisk resource.
type AttachedDiskArgs struct {
	// Specifies a unique device name of your choice that is
	// reflected into the /dev/disk/by-id/google-* tree of a Linux operating
	// system running within the instance. This name can be used to
	// reference the device for mounting, resizing, and so on, from within
	// the instance.
	DeviceName pulumi.StringPtrInput
	// `name` or `selfLink` of the disk that will be attached.
	Disk pulumi.StringInput
	// `name` or `selfLink` of the compute instance that the disk will be attached to.
	// If the `selfLink` is provided then `zone` and `project` are extracted from the
	// self link. If only the name is used then `zone` and `project` must be defined
	// as properties on the resource or provider.
	Instance pulumi.StringInput
	// The mode in which to attach this disk, either READ_WRITE or
	// READ_ONLY. If not specified, the default is to attach the disk in
	// READ_WRITE mode.
	Mode pulumi.StringPtrInput
	// The project that the referenced compute instance is a part of. If `instance` is referenced by its
	// `selfLink` the project defined in the link will take precedence.
	Project pulumi.StringPtrInput
	// The zone that the referenced compute instance is located within. If `instance` is referenced by its
	// `selfLink` the zone defined in the link will take precedence.
	Zone pulumi.StringPtrInput
}

func (AttachedDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDiskArgs)(nil)).Elem()
}
