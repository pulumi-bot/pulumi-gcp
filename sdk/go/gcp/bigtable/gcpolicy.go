// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigtable

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a Google Cloud Bigtable GC Policy inside a family. For more information see
// [the official documentation](https://cloud.google.com/bigtable/) and
// [API](https://cloud.google.com/bigtable/docs/go/reference).
type GCPolicy struct {
	pulumi.CustomResourceState

	// The name of the column family.
	ColumnFamily pulumi.StringOutput `pulumi:"columnFamily"`
	// The name of the Bigtable instance.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// GC policy that applies to all cells older than the given age.
	MaxAges GCPolicyMaxAgeArrayOutput `pulumi:"maxAges"`
	// GC policy that applies to all versions of a cell except for the most recent.
	MaxVersions GCPolicyMaxVersionArrayOutput `pulumi:"maxVersions"`
	// If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The name of the table.
	Table pulumi.StringOutput `pulumi:"table"`
}

// NewGCPolicy registers a new resource with the given unique name, arguments, and options.
func NewGCPolicy(ctx *pulumi.Context,
	name string, args *GCPolicyArgs, opts ...pulumi.ResourceOption) (*GCPolicy, error) {
	if args == nil || args.ColumnFamily == nil {
		return nil, errors.New("missing required argument 'ColumnFamily'")
	}
	if args == nil || args.InstanceName == nil {
		return nil, errors.New("missing required argument 'InstanceName'")
	}
	if args == nil || args.Table == nil {
		return nil, errors.New("missing required argument 'Table'")
	}
	if args == nil {
		args = &GCPolicyArgs{}
	}
	var resource GCPolicy
	err := ctx.RegisterResource("gcp:bigtable/gCPolicy:GCPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGCPolicy gets an existing GCPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGCPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GCPolicyState, opts ...pulumi.ResourceOption) (*GCPolicy, error) {
	var resource GCPolicy
	err := ctx.ReadResource("gcp:bigtable/gCPolicy:GCPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GCPolicy resources.
type gcpolicyState struct {
	// The name of the column family.
	ColumnFamily *string `pulumi:"columnFamily"`
	// The name of the Bigtable instance.
	InstanceName *string `pulumi:"instanceName"`
	// GC policy that applies to all cells older than the given age.
	MaxAges []GCPolicyMaxAge `pulumi:"maxAges"`
	// GC policy that applies to all versions of a cell except for the most recent.
	MaxVersions []GCPolicyMaxVersion `pulumi:"maxVersions"`
	// If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
	Mode *string `pulumi:"mode"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The name of the table.
	Table *string `pulumi:"table"`
}

type GCPolicyState struct {
	// The name of the column family.
	ColumnFamily pulumi.StringPtrInput
	// The name of the Bigtable instance.
	InstanceName pulumi.StringPtrInput
	// GC policy that applies to all cells older than the given age.
	MaxAges GCPolicyMaxAgeArrayInput
	// GC policy that applies to all versions of a cell except for the most recent.
	MaxVersions GCPolicyMaxVersionArrayInput
	// If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
	Mode pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The name of the table.
	Table pulumi.StringPtrInput
}

func (GCPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*gcpolicyState)(nil)).Elem()
}

type gcpolicyArgs struct {
	// The name of the column family.
	ColumnFamily string `pulumi:"columnFamily"`
	// The name of the Bigtable instance.
	InstanceName string `pulumi:"instanceName"`
	// GC policy that applies to all cells older than the given age.
	MaxAges []GCPolicyMaxAge `pulumi:"maxAges"`
	// GC policy that applies to all versions of a cell except for the most recent.
	MaxVersions []GCPolicyMaxVersion `pulumi:"maxVersions"`
	// If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
	Mode *string `pulumi:"mode"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The name of the table.
	Table string `pulumi:"table"`
}

// The set of arguments for constructing a GCPolicy resource.
type GCPolicyArgs struct {
	// The name of the column family.
	ColumnFamily pulumi.StringInput
	// The name of the Bigtable instance.
	InstanceName pulumi.StringInput
	// GC policy that applies to all cells older than the given age.
	MaxAges GCPolicyMaxAgeArrayInput
	// GC policy that applies to all versions of a cell except for the most recent.
	MaxVersions GCPolicyMaxVersionArrayInput
	// If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
	Mode pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The name of the table.
	Table pulumi.StringInput
}

func (GCPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gcpolicyArgs)(nil)).Elem()
}

type GCPolicyInput interface {
	pulumi.Input

	ToGCPolicyOutput() GCPolicyOutput
	ToGCPolicyOutputWithContext(ctx context.Context) GCPolicyOutput
}

func (GCPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*GCPolicy)(nil)).Elem()
}

func (i GCPolicy) ToGCPolicyOutput() GCPolicyOutput {
	return i.ToGCPolicyOutputWithContext(context.Background())
}

func (i GCPolicy) ToGCPolicyOutputWithContext(ctx context.Context) GCPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GCPolicyOutput)
}

type GCPolicyOutput struct {
	*pulumi.OutputState
}

func (GCPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GCPolicyOutput)(nil)).Elem()
}

func (o GCPolicyOutput) ToGCPolicyOutput() GCPolicyOutput {
	return o
}

func (o GCPolicyOutput) ToGCPolicyOutputWithContext(ctx context.Context) GCPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GCPolicyOutput{})
}
