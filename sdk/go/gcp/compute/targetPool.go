// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Target Pool within GCE. This is a collection of instances used as
// target of a network load balancer (Forwarding Rule). For more information see
// [the official
// documentation](https://cloud.google.com/compute/docs/load-balancing/network/target-pools)
// and [API](https://cloud.google.com/compute/docs/reference/latest/targetPools).
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
// 		defaultHttpHealthCheck, err := compute.NewHttpHealthCheck(ctx, "defaultHttpHealthCheck", &compute.HttpHealthCheckArgs{
// 			RequestPath:      pulumi.String("/"),
// 			CheckIntervalSec: pulumi.Int(1),
// 			TimeoutSec:       pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewTargetPool(ctx, "defaultTargetPool", &compute.TargetPoolArgs{
// 			Instances: pulumi.StringArray{
// 				pulumi.String("us-central1-a/myinstance1"),
// 				pulumi.String("us-central1-b/myinstance2"),
// 			},
// 			HealthChecks: pulumi.String(pulumi.String{
// 				defaultHttpHealthCheck.Name,
// 			}),
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
// Target pools can be imported using any of the following formats
//
// ```sh
//  $ pulumi import gcp:compute/targetPool:TargetPool default projects/{{project}}/regions/{{region}}/targetPools/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/targetPool:TargetPool default {{project}}/{{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/targetPool:TargetPool default {{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/targetPool:TargetPool default {{name}}
// ```
type TargetPool struct {
	pulumi.CustomResourceState

	// URL to the backup target pool. Must also set
	// failover\_ratio.
	BackupPool pulumi.StringPtrOutput `pulumi:"backupPool"`
	// Textual description field.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Ratio (0 to 1) of failed nodes before using the
	// backup pool (which must also be set).
	FailoverRatio pulumi.Float64PtrOutput `pulumi:"failoverRatio"`
	// List of zero or one health check name or self_link. Only
	// legacy `compute.HttpHealthCheck` is supported.
	HealthChecks pulumi.StringPtrOutput `pulumi:"healthChecks"`
	// List of instances in the pool. They can be given as
	// URLs, or in the form of "zone/name". Note that the instances need not exist
	// at the time of target pool creation, so there is no need to use the
	// interpolation to create a dependency on the instances from the
	// target pool.
	Instances pulumi.StringArrayOutput `pulumi:"instances"`
	// A unique name for the resource, required by GCE. Changing
	// this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Where the target pool resides. Defaults to project
	// region.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// How to distribute load. Options are "NONE" (no
	// affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and
	// "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE").
	SessionAffinity pulumi.StringPtrOutput `pulumi:"sessionAffinity"`
}

// NewTargetPool registers a new resource with the given unique name, arguments, and options.
func NewTargetPool(ctx *pulumi.Context,
	name string, args *TargetPoolArgs, opts ...pulumi.ResourceOption) (*TargetPool, error) {
	if args == nil {
		args = &TargetPoolArgs{}
	}

	var resource TargetPool
	err := ctx.RegisterResource("gcp:compute/targetPool:TargetPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetPool gets an existing TargetPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetPoolState, opts ...pulumi.ResourceOption) (*TargetPool, error) {
	var resource TargetPool
	err := ctx.ReadResource("gcp:compute/targetPool:TargetPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetPool resources.
type targetPoolState struct {
	// URL to the backup target pool. Must also set
	// failover\_ratio.
	BackupPool *string `pulumi:"backupPool"`
	// Textual description field.
	Description *string `pulumi:"description"`
	// Ratio (0 to 1) of failed nodes before using the
	// backup pool (which must also be set).
	FailoverRatio *float64 `pulumi:"failoverRatio"`
	// List of zero or one health check name or self_link. Only
	// legacy `compute.HttpHealthCheck` is supported.
	HealthChecks *string `pulumi:"healthChecks"`
	// List of instances in the pool. They can be given as
	// URLs, or in the form of "zone/name". Note that the instances need not exist
	// at the time of target pool creation, so there is no need to use the
	// interpolation to create a dependency on the instances from the
	// target pool.
	Instances []string `pulumi:"instances"`
	// A unique name for the resource, required by GCE. Changing
	// this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Where the target pool resides. Defaults to project
	// region.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// How to distribute load. Options are "NONE" (no
	// affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and
	// "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE").
	SessionAffinity *string `pulumi:"sessionAffinity"`
}

type TargetPoolState struct {
	// URL to the backup target pool. Must also set
	// failover\_ratio.
	BackupPool pulumi.StringPtrInput
	// Textual description field.
	Description pulumi.StringPtrInput
	// Ratio (0 to 1) of failed nodes before using the
	// backup pool (which must also be set).
	FailoverRatio pulumi.Float64PtrInput
	// List of zero or one health check name or self_link. Only
	// legacy `compute.HttpHealthCheck` is supported.
	HealthChecks pulumi.StringPtrInput
	// List of instances in the pool. They can be given as
	// URLs, or in the form of "zone/name". Note that the instances need not exist
	// at the time of target pool creation, so there is no need to use the
	// interpolation to create a dependency on the instances from the
	// target pool.
	Instances pulumi.StringArrayInput
	// A unique name for the resource, required by GCE. Changing
	// this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Where the target pool resides. Defaults to project
	// region.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// How to distribute load. Options are "NONE" (no
	// affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and
	// "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE").
	SessionAffinity pulumi.StringPtrInput
}

func (TargetPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetPoolState)(nil)).Elem()
}

type targetPoolArgs struct {
	// URL to the backup target pool. Must also set
	// failover\_ratio.
	BackupPool *string `pulumi:"backupPool"`
	// Textual description field.
	Description *string `pulumi:"description"`
	// Ratio (0 to 1) of failed nodes before using the
	// backup pool (which must also be set).
	FailoverRatio *float64 `pulumi:"failoverRatio"`
	// List of zero or one health check name or self_link. Only
	// legacy `compute.HttpHealthCheck` is supported.
	HealthChecks *string `pulumi:"healthChecks"`
	// List of instances in the pool. They can be given as
	// URLs, or in the form of "zone/name". Note that the instances need not exist
	// at the time of target pool creation, so there is no need to use the
	// interpolation to create a dependency on the instances from the
	// target pool.
	Instances []string `pulumi:"instances"`
	// A unique name for the resource, required by GCE. Changing
	// this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Where the target pool resides. Defaults to project
	// region.
	Region *string `pulumi:"region"`
	// How to distribute load. Options are "NONE" (no
	// affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and
	// "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE").
	SessionAffinity *string `pulumi:"sessionAffinity"`
}

// The set of arguments for constructing a TargetPool resource.
type TargetPoolArgs struct {
	// URL to the backup target pool. Must also set
	// failover\_ratio.
	BackupPool pulumi.StringPtrInput
	// Textual description field.
	Description pulumi.StringPtrInput
	// Ratio (0 to 1) of failed nodes before using the
	// backup pool (which must also be set).
	FailoverRatio pulumi.Float64PtrInput
	// List of zero or one health check name or self_link. Only
	// legacy `compute.HttpHealthCheck` is supported.
	HealthChecks pulumi.StringPtrInput
	// List of instances in the pool. They can be given as
	// URLs, or in the form of "zone/name". Note that the instances need not exist
	// at the time of target pool creation, so there is no need to use the
	// interpolation to create a dependency on the instances from the
	// target pool.
	Instances pulumi.StringArrayInput
	// A unique name for the resource, required by GCE. Changing
	// this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Where the target pool resides. Defaults to project
	// region.
	Region pulumi.StringPtrInput
	// How to distribute load. Options are "NONE" (no
	// affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and
	// "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE").
	SessionAffinity pulumi.StringPtrInput
}

func (TargetPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetPoolArgs)(nil)).Elem()
}

type TargetPoolInput interface {
	pulumi.Input

	ToTargetPoolOutput() TargetPoolOutput
	ToTargetPoolOutputWithContext(ctx context.Context) TargetPoolOutput
}

func (*TargetPool) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetPool)(nil))
}

func (i *TargetPool) ToTargetPoolOutput() TargetPoolOutput {
	return i.ToTargetPoolOutputWithContext(context.Background())
}

func (i *TargetPool) ToTargetPoolOutputWithContext(ctx context.Context) TargetPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetPoolOutput)
}

func (i *TargetPool) ToTargetPoolPtrOutput() TargetPoolPtrOutput {
	return i.ToTargetPoolPtrOutputWithContext(context.Background())
}

func (i *TargetPool) ToTargetPoolPtrOutputWithContext(ctx context.Context) TargetPoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetPoolPtrOutput)
}

type TargetPoolPtrInput interface {
	pulumi.Input

	ToTargetPoolPtrOutput() TargetPoolPtrOutput
	ToTargetPoolPtrOutputWithContext(ctx context.Context) TargetPoolPtrOutput
}

type targetPoolPtrType TargetPoolArgs

func (*targetPoolPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetPool)(nil))
}

func (i *targetPoolPtrType) ToTargetPoolPtrOutput() TargetPoolPtrOutput {
	return i.ToTargetPoolPtrOutputWithContext(context.Background())
}

func (i *targetPoolPtrType) ToTargetPoolPtrOutputWithContext(ctx context.Context) TargetPoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetPoolPtrOutput)
}

// TargetPoolArrayInput is an input type that accepts TargetPoolArray and TargetPoolArrayOutput values.
// You can construct a concrete instance of `TargetPoolArrayInput` via:
//
//          TargetPoolArray{ TargetPoolArgs{...} }
type TargetPoolArrayInput interface {
	pulumi.Input

	ToTargetPoolArrayOutput() TargetPoolArrayOutput
	ToTargetPoolArrayOutputWithContext(context.Context) TargetPoolArrayOutput
}

type TargetPoolArray []TargetPoolInput

func (TargetPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*TargetPool)(nil))
}

func (i TargetPoolArray) ToTargetPoolArrayOutput() TargetPoolArrayOutput {
	return i.ToTargetPoolArrayOutputWithContext(context.Background())
}

func (i TargetPoolArray) ToTargetPoolArrayOutputWithContext(ctx context.Context) TargetPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetPoolArrayOutput)
}

// TargetPoolMapInput is an input type that accepts TargetPoolMap and TargetPoolMapOutput values.
// You can construct a concrete instance of `TargetPoolMapInput` via:
//
//          TargetPoolMap{ "key": TargetPoolArgs{...} }
type TargetPoolMapInput interface {
	pulumi.Input

	ToTargetPoolMapOutput() TargetPoolMapOutput
	ToTargetPoolMapOutputWithContext(context.Context) TargetPoolMapOutput
}

type TargetPoolMap map[string]TargetPoolInput

func (TargetPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*TargetPool)(nil))
}

func (i TargetPoolMap) ToTargetPoolMapOutput() TargetPoolMapOutput {
	return i.ToTargetPoolMapOutputWithContext(context.Background())
}

func (i TargetPoolMap) ToTargetPoolMapOutputWithContext(ctx context.Context) TargetPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetPoolMapOutput)
}

type TargetPoolOutput struct {
	*pulumi.OutputState
}

func (TargetPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetPool)(nil))
}

func (o TargetPoolOutput) ToTargetPoolOutput() TargetPoolOutput {
	return o
}

func (o TargetPoolOutput) ToTargetPoolOutputWithContext(ctx context.Context) TargetPoolOutput {
	return o
}

func (o TargetPoolOutput) ToTargetPoolPtrOutput() TargetPoolPtrOutput {
	return o.ToTargetPoolPtrOutputWithContext(context.Background())
}

func (o TargetPoolOutput) ToTargetPoolPtrOutputWithContext(ctx context.Context) TargetPoolPtrOutput {
	return o.ApplyT(func(v TargetPool) *TargetPool {
		return &v
	}).(TargetPoolPtrOutput)
}

type TargetPoolPtrOutput struct {
	*pulumi.OutputState
}

func (TargetPoolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetPool)(nil))
}

func (o TargetPoolPtrOutput) ToTargetPoolPtrOutput() TargetPoolPtrOutput {
	return o
}

func (o TargetPoolPtrOutput) ToTargetPoolPtrOutputWithContext(ctx context.Context) TargetPoolPtrOutput {
	return o
}

type TargetPoolArrayOutput struct{ *pulumi.OutputState }

func (TargetPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetPool)(nil))
}

func (o TargetPoolArrayOutput) ToTargetPoolArrayOutput() TargetPoolArrayOutput {
	return o
}

func (o TargetPoolArrayOutput) ToTargetPoolArrayOutputWithContext(ctx context.Context) TargetPoolArrayOutput {
	return o
}

func (o TargetPoolArrayOutput) Index(i pulumi.IntInput) TargetPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetPool {
		return vs[0].([]TargetPool)[vs[1].(int)]
	}).(TargetPoolOutput)
}

type TargetPoolMapOutput struct{ *pulumi.OutputState }

func (TargetPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TargetPool)(nil))
}

func (o TargetPoolMapOutput) ToTargetPoolMapOutput() TargetPoolMapOutput {
	return o
}

func (o TargetPoolMapOutput) ToTargetPoolMapOutputWithContext(ctx context.Context) TargetPoolMapOutput {
	return o
}

func (o TargetPoolMapOutput) MapIndex(k pulumi.StringInput) TargetPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TargetPool {
		return vs[0].(map[string]TargetPool)[vs[1].(string)]
	}).(TargetPoolOutput)
}

func init() {
	pulumi.RegisterOutputType(TargetPoolOutput{})
	pulumi.RegisterOutputType(TargetPoolPtrOutput{})
	pulumi.RegisterOutputType(TargetPoolArrayOutput{})
	pulumi.RegisterOutputType(TargetPoolMapOutput{})
}
