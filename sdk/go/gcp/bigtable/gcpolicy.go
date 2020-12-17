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
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigtable"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		instance, err := bigtable.NewInstance(ctx, "instance", &bigtable.InstanceArgs{
// 			Clusters: bigtable.InstanceClusterArray{
// 				&bigtable.InstanceClusterArgs{
// 					ClusterId:   pulumi.String("tf-instance-cluster"),
// 					Zone:        pulumi.String("us-central1-b"),
// 					NumNodes:    pulumi.Int(3),
// 					StorageType: pulumi.String("HDD"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		table, err := bigtable.NewTable(ctx, "table", &bigtable.TableArgs{
// 			InstanceName: instance.Name,
// 			ColumnFamilies: bigtable.TableColumnFamilyArray{
// 				&bigtable.TableColumnFamilyArgs{
// 					Family: pulumi.String("name"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigtable.NewGCPolicy(ctx, "policy", &bigtable.GCPolicyArgs{
// 			InstanceName: instance.Name,
// 			Table:        table.Name,
// 			ColumnFamily: pulumi.String("name"),
// 			MaxAges: bigtable.GCPolicyMaxAgeArray{
// 				&bigtable.GCPolicyMaxAgeArgs{
// 					Days: pulumi.Int(7),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Multiple conditions is also supported. `UNION` when any of its sub-policies apply (OR). `INTERSECTION` when all its sub-policies apply (AND)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigtable"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := bigtable.NewGCPolicy(ctx, "policy", &bigtable.GCPolicyArgs{
// 			InstanceName: pulumi.Any(google_bigtable_instance.Instance.Name),
// 			Table:        pulumi.Any(google_bigtable_table.Table.Name),
// 			ColumnFamily: pulumi.String("name"),
// 			Mode:         pulumi.String("UNION"),
// 			MaxAges: bigtable.GCPolicyMaxAgeArray{
// 				&bigtable.GCPolicyMaxAgeArgs{
// 					Days: pulumi.Int(7),
// 				},
// 			},
// 			MaxVersions: bigtable.GCPolicyMaxVersionArray{
// 				&bigtable.GCPolicyMaxVersionArgs{
// 					Number: pulumi.Int(10),
// 				},
// 			},
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
// This resource does not support import.
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ColumnFamily == nil {
		return nil, errors.New("invalid value for required argument 'ColumnFamily'")
	}
	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.Table == nil {
		return nil, errors.New("invalid value for required argument 'Table'")
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

func (*GCPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*GCPolicy)(nil))
}

func (i *GCPolicy) ToGCPolicyOutput() GCPolicyOutput {
	return i.ToGCPolicyOutputWithContext(context.Background())
}

func (i *GCPolicy) ToGCPolicyOutputWithContext(ctx context.Context) GCPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GCPolicyOutput)
}

func (i *GCPolicy) ToGCPolicyPtrOutput() GCPolicyPtrOutput {
	return i.ToGCPolicyPtrOutputWithContext(context.Background())
}

func (i *GCPolicy) ToGCPolicyPtrOutputWithContext(ctx context.Context) GCPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GCPolicyPtrOutput)
}

type GCPolicyPtrInput interface {
	pulumi.Input

	ToGCPolicyPtrOutput() GCPolicyPtrOutput
	ToGCPolicyPtrOutputWithContext(ctx context.Context) GCPolicyPtrOutput
}

type gcpolicyPtrType GCPolicyArgs

func (*gcpolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GCPolicy)(nil))
}

func (i *gcpolicyPtrType) ToGCPolicyPtrOutput() GCPolicyPtrOutput {
	return i.ToGCPolicyPtrOutputWithContext(context.Background())
}

func (i *gcpolicyPtrType) ToGCPolicyPtrOutputWithContext(ctx context.Context) GCPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GCPolicyOutput).ToGCPolicyPtrOutput()
}

// GCPolicyArrayInput is an input type that accepts GCPolicyArray and GCPolicyArrayOutput values.
// You can construct a concrete instance of `GCPolicyArrayInput` via:
//
//          GCPolicyArray{ GCPolicyArgs{...} }
type GCPolicyArrayInput interface {
	pulumi.Input

	ToGCPolicyArrayOutput() GCPolicyArrayOutput
	ToGCPolicyArrayOutputWithContext(context.Context) GCPolicyArrayOutput
}

type GCPolicyArray []GCPolicyInput

func (GCPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GCPolicy)(nil))
}

func (i GCPolicyArray) ToGCPolicyArrayOutput() GCPolicyArrayOutput {
	return i.ToGCPolicyArrayOutputWithContext(context.Background())
}

func (i GCPolicyArray) ToGCPolicyArrayOutputWithContext(ctx context.Context) GCPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GCPolicyArrayOutput)
}

// GCPolicyMapInput is an input type that accepts GCPolicyMap and GCPolicyMapOutput values.
// You can construct a concrete instance of `GCPolicyMapInput` via:
//
//          GCPolicyMap{ "key": GCPolicyArgs{...} }
type GCPolicyMapInput interface {
	pulumi.Input

	ToGCPolicyMapOutput() GCPolicyMapOutput
	ToGCPolicyMapOutputWithContext(context.Context) GCPolicyMapOutput
}

type GCPolicyMap map[string]GCPolicyInput

func (GCPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GCPolicy)(nil))
}

func (i GCPolicyMap) ToGCPolicyMapOutput() GCPolicyMapOutput {
	return i.ToGCPolicyMapOutputWithContext(context.Background())
}

func (i GCPolicyMap) ToGCPolicyMapOutputWithContext(ctx context.Context) GCPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GCPolicyMapOutput)
}

type GCPolicyOutput struct {
	*pulumi.OutputState
}

func (GCPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GCPolicy)(nil))
}

func (o GCPolicyOutput) ToGCPolicyOutput() GCPolicyOutput {
	return o
}

func (o GCPolicyOutput) ToGCPolicyOutputWithContext(ctx context.Context) GCPolicyOutput {
	return o
}

func (o GCPolicyOutput) ToGCPolicyPtrOutput() GCPolicyPtrOutput {
	return o.ToGCPolicyPtrOutputWithContext(context.Background())
}

func (o GCPolicyOutput) ToGCPolicyPtrOutputWithContext(ctx context.Context) GCPolicyPtrOutput {
	return o.ApplyT(func(v GCPolicy) *GCPolicy {
		return &v
	}).(GCPolicyPtrOutput)
}

type GCPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (GCPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GCPolicy)(nil))
}

func (o GCPolicyPtrOutput) ToGCPolicyPtrOutput() GCPolicyPtrOutput {
	return o
}

func (o GCPolicyPtrOutput) ToGCPolicyPtrOutputWithContext(ctx context.Context) GCPolicyPtrOutput {
	return o
}

type GCPolicyArrayOutput struct{ *pulumi.OutputState }

func (GCPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GCPolicy)(nil))
}

func (o GCPolicyArrayOutput) ToGCPolicyArrayOutput() GCPolicyArrayOutput {
	return o
}

func (o GCPolicyArrayOutput) ToGCPolicyArrayOutputWithContext(ctx context.Context) GCPolicyArrayOutput {
	return o
}

func (o GCPolicyArrayOutput) Index(i pulumi.IntInput) GCPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GCPolicy {
		return vs[0].([]GCPolicy)[vs[1].(int)]
	}).(GCPolicyOutput)
}

type GCPolicyMapOutput struct{ *pulumi.OutputState }

func (GCPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GCPolicy)(nil))
}

func (o GCPolicyMapOutput) ToGCPolicyMapOutput() GCPolicyMapOutput {
	return o
}

func (o GCPolicyMapOutput) ToGCPolicyMapOutputWithContext(ctx context.Context) GCPolicyMapOutput {
	return o
}

func (o GCPolicyMapOutput) MapIndex(k pulumi.StringInput) GCPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GCPolicy {
		return vs[0].(map[string]GCPolicy)[vs[1].(string)]
	}).(GCPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(GCPolicyOutput{})
	pulumi.RegisterOutputType(GCPolicyPtrOutput{})
	pulumi.RegisterOutputType(GCPolicyArrayOutput{})
	pulumi.RegisterOutputType(GCPolicyMapOutput{})
}
