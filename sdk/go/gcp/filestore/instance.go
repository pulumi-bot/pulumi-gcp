// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package filestore

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Google Cloud Filestore instance.
//
// To get more information about Instance, see:
//
// * [API documentation](https://cloud.google.com/filestore/docs/reference/rest/v1beta1/projects.locations.instances/create)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/filestore/docs/creating-instances)
//     * [Use with Kubernetes](https://cloud.google.com/filestore/docs/accessing-fileshares)
//     * [Copying Data In/Out](https://cloud.google.com/filestore/docs/copying-data)
//
// ## Example Usage
type Instance struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A description of the instance.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// File system shares on the instance. For this version, only a
	// single file share is supported.
	// Structure is documented below.
	FileShares InstanceFileSharesOutput `pulumi:"fileShares"`
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of the fileshare (16 characters or less)
	Name pulumi.StringOutput `pulumi:"name"`
	// VPC networks to which the instance is connected. For this version,
	// only a single network is supported.
	// Structure is documented below.
	Networks InstanceNetworkArrayOutput `pulumi:"networks"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The service tier of the instance.
	// Possible values are `TIER_UNSPECIFIED`, `STANDARD`, `PREMIUM`, `BASIC_HDD`, `BASIC_SSD`, and `HIGH_SCALE_SSD`.
	Tier pulumi.StringOutput `pulumi:"tier"`
	// The name of the Filestore zone of the instance.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileShares == nil {
		return nil, errors.New("invalid value for required argument 'FileShares'")
	}
	if args.Networks == nil {
		return nil, errors.New("invalid value for required argument 'Networks'")
	}
	if args.Tier == nil {
		return nil, errors.New("invalid value for required argument 'Tier'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	var resource Instance
	err := ctx.RegisterResource("gcp:filestore/instance:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:filestore/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Creation timestamp in RFC3339 text format.
	CreateTime *string `pulumi:"createTime"`
	// A description of the instance.
	Description *string `pulumi:"description"`
	// Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
	Etag *string `pulumi:"etag"`
	// File system shares on the instance. For this version, only a
	// single file share is supported.
	// Structure is documented below.
	FileShares *InstanceFileShares `pulumi:"fileShares"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// The name of the fileshare (16 characters or less)
	Name *string `pulumi:"name"`
	// VPC networks to which the instance is connected. For this version,
	// only a single network is supported.
	// Structure is documented below.
	Networks []InstanceNetwork `pulumi:"networks"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The service tier of the instance.
	// Possible values are `TIER_UNSPECIFIED`, `STANDARD`, `PREMIUM`, `BASIC_HDD`, `BASIC_SSD`, and `HIGH_SCALE_SSD`.
	Tier *string `pulumi:"tier"`
	// The name of the Filestore zone of the instance.
	Zone *string `pulumi:"zone"`
}

type InstanceState struct {
	// Creation timestamp in RFC3339 text format.
	CreateTime pulumi.StringPtrInput
	// A description of the instance.
	Description pulumi.StringPtrInput
	// Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
	Etag pulumi.StringPtrInput
	// File system shares on the instance. For this version, only a
	// single file share is supported.
	// Structure is documented below.
	FileShares InstanceFileSharesPtrInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// The name of the fileshare (16 characters or less)
	Name pulumi.StringPtrInput
	// VPC networks to which the instance is connected. For this version,
	// only a single network is supported.
	// Structure is documented below.
	Networks InstanceNetworkArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The service tier of the instance.
	// Possible values are `TIER_UNSPECIFIED`, `STANDARD`, `PREMIUM`, `BASIC_HDD`, `BASIC_SSD`, and `HIGH_SCALE_SSD`.
	Tier pulumi.StringPtrInput
	// The name of the Filestore zone of the instance.
	Zone pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// A description of the instance.
	Description *string `pulumi:"description"`
	// File system shares on the instance. For this version, only a
	// single file share is supported.
	// Structure is documented below.
	FileShares InstanceFileShares `pulumi:"fileShares"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// The name of the fileshare (16 characters or less)
	Name *string `pulumi:"name"`
	// VPC networks to which the instance is connected. For this version,
	// only a single network is supported.
	// Structure is documented below.
	Networks []InstanceNetwork `pulumi:"networks"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The service tier of the instance.
	// Possible values are `TIER_UNSPECIFIED`, `STANDARD`, `PREMIUM`, `BASIC_HDD`, `BASIC_SSD`, and `HIGH_SCALE_SSD`.
	Tier string `pulumi:"tier"`
	// The name of the Filestore zone of the instance.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// A description of the instance.
	Description pulumi.StringPtrInput
	// File system shares on the instance. For this version, only a
	// single file share is supported.
	// Structure is documented below.
	FileShares InstanceFileSharesInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// The name of the fileshare (16 characters or less)
	Name pulumi.StringPtrInput
	// VPC networks to which the instance is connected. For this version,
	// only a single network is supported.
	// Structure is documented below.
	Networks InstanceNetworkArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The service tier of the instance.
	// Possible values are `TIER_UNSPECIFIED`, `STANDARD`, `PREMIUM`, `BASIC_HDD`, `BASIC_SSD`, and `HIGH_SCALE_SSD`.
	Tier pulumi.StringInput
	// The name of the Filestore zone of the instance.
	Zone pulumi.StringInput
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
