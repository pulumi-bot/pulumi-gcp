// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a group of dissimilar Compute Engine virtual machine instances.
// For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/#unmanaged_instance_groups)
// and [API](https://cloud.google.com/compute/docs/reference/latest/instanceGroups)
//
// ## Example Usage
// ### Empty Instance Group
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewInstanceGroup(ctx, "test", &compute.InstanceGroupArgs{
// 			Description: pulumi.String("Test instance group"),
// 			Zone:        pulumi.String("us-central1-a"),
// 			Network:     pulumi.Any(google_compute_network.Default.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Example Usage - With instances and named ports
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewInstanceGroup(ctx, "webservers", &compute.InstanceGroupArgs{
// 			Description: pulumi.String("Test instance group"),
// 			Instances: pulumi.StringArray{
// 				pulumi.Any(google_compute_instance.Test.Id),
// 				pulumi.Any(google_compute_instance.Test2.Id),
// 			},
// 			NamedPorts: compute.InstanceGroupNamedPortArray{
// 				&compute.InstanceGroupNamedPortArgs{
// 					Name: pulumi.String("http"),
// 					Port: pulumi.Int(8080),
// 				},
// 				&compute.InstanceGroupNamedPortArgs{
// 					Name: pulumi.String("https"),
// 					Port: pulumi.Int(8443),
// 				},
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
// ### Example Usage - Recreating an instance group in use
// Recreating an instance group that's in use by another resource will give a
// `resourceInUseByAnotherResource` error. Use `lifecycle.create_before_destroy`
// as shown in this example to avoid this type of error.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "debian-9"
// 		opt1 := "debian-cloud"
// 		debianImage, err := compute.LookupImage(ctx, &compute.LookupImageArgs{
// 			Family:  &opt0,
// 			Project: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		stagingVm, err := compute.NewInstance(ctx, "stagingVm", &compute.InstanceArgs{
// 			MachineType: pulumi.String("e2-medium"),
// 			Zone:        pulumi.String("us-central1-c"),
// 			BootDisk: &compute.InstanceBootDiskArgs{
// 				InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
// 					Image: pulumi.String(debianImage.SelfLink),
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
// 		stagingGroup, err := compute.NewInstanceGroup(ctx, "stagingGroup", &compute.InstanceGroupArgs{
// 			Zone: pulumi.String("us-central1-c"),
// 			Instances: pulumi.StringArray{
// 				stagingVm.ID(),
// 			},
// 			NamedPorts: compute.InstanceGroupNamedPortArray{
// 				&compute.InstanceGroupNamedPortArgs{
// 					Name: pulumi.String("http"),
// 					Port: pulumi.Int(8080),
// 				},
// 				&compute.InstanceGroupNamedPortArgs{
// 					Name: pulumi.String("https"),
// 					Port: pulumi.Int(8443),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		stagingHealth, err := compute.NewHttpsHealthCheck(ctx, "stagingHealth", &compute.HttpsHealthCheckArgs{
// 			RequestPath: pulumi.String("/health_check"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewBackendService(ctx, "stagingService", &compute.BackendServiceArgs{
// 			PortName: pulumi.String("https"),
// 			Protocol: pulumi.String("HTTPS"),
// 			Backends: compute.BackendServiceBackendArray{
// 				&compute.BackendServiceBackendArgs{
// 					Group: stagingGroup.ID(),
// 				},
// 			},
// 			HealthChecks: pulumi.String{
// 				stagingHealth.ID(),
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
// Instance group can be imported using the `zone` and `name` with an optional `project`, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/instanceGroup:InstanceGroup webservers us-central1-a/terraform-webservers
// ```
//
// ```sh
//  $ pulumi import gcp:compute/instanceGroup:InstanceGroup webservers big-project/us-central1-a/terraform-webservers
// ```
//
// ```sh
//  $ pulumi import gcp:compute/instanceGroup:InstanceGroup webservers projects/big-project/zones/us-central1-a/instanceGroups/terraform-webservers
// ```
type InstanceGroup struct {
	pulumi.CustomResourceState

	// An optional textual description of the instance
	// group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of instances in the group. They should be given
	// as either selfLink or id. When adding instances they must all be in the same
	// network and zone as the instance group.
	Instances pulumi.StringArrayOutput `pulumi:"instances"`
	// The name which the port will be mapped to.
	Name pulumi.StringOutput `pulumi:"name"`
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts InstanceGroupNamedPortTypeArrayOutput `pulumi:"namedPorts"`
	// The URL of the network the instance group is in. If
	// this is different from the network where the instances are in, the creation
	// fails. Defaults to the network where the instances are in (if neither
	// `network` nor `instances` is specified, this field will be blank).
	Network pulumi.StringOutput `pulumi:"network"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The number of instances in the group.
	Size pulumi.IntOutput `pulumi:"size"`
	// The zone that this instance group should be created in.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceGroup registers a new resource with the given unique name, arguments, and options.
func NewInstanceGroup(ctx *pulumi.Context,
	name string, args *InstanceGroupArgs, opts ...pulumi.ResourceOption) (*InstanceGroup, error) {
	if args == nil {
		args = &InstanceGroupArgs{}
	}

	var resource InstanceGroup
	err := ctx.RegisterResource("gcp:compute/instanceGroup:InstanceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceGroup gets an existing InstanceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceGroupState, opts ...pulumi.ResourceOption) (*InstanceGroup, error) {
	var resource InstanceGroup
	err := ctx.ReadResource("gcp:compute/instanceGroup:InstanceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceGroup resources.
type instanceGroupState struct {
	// An optional textual description of the instance
	// group.
	Description *string `pulumi:"description"`
	// List of instances in the group. They should be given
	// as either selfLink or id. When adding instances they must all be in the same
	// network and zone as the instance group.
	Instances []string `pulumi:"instances"`
	// The name which the port will be mapped to.
	Name *string `pulumi:"name"`
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts []InstanceGroupNamedPortType `pulumi:"namedPorts"`
	// The URL of the network the instance group is in. If
	// this is different from the network where the instances are in, the creation
	// fails. Defaults to the network where the instances are in (if neither
	// `network` nor `instances` is specified, this field will be blank).
	Network *string `pulumi:"network"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The number of instances in the group.
	Size *int `pulumi:"size"`
	// The zone that this instance group should be created in.
	Zone *string `pulumi:"zone"`
}

type InstanceGroupState struct {
	// An optional textual description of the instance
	// group.
	Description pulumi.StringPtrInput
	// List of instances in the group. They should be given
	// as either selfLink or id. When adding instances they must all be in the same
	// network and zone as the instance group.
	Instances pulumi.StringArrayInput
	// The name which the port will be mapped to.
	Name pulumi.StringPtrInput
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts InstanceGroupNamedPortTypeArrayInput
	// The URL of the network the instance group is in. If
	// this is different from the network where the instances are in, the creation
	// fails. Defaults to the network where the instances are in (if neither
	// `network` nor `instances` is specified, this field will be blank).
	Network pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The number of instances in the group.
	Size pulumi.IntPtrInput
	// The zone that this instance group should be created in.
	Zone pulumi.StringPtrInput
}

func (InstanceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupState)(nil)).Elem()
}

type instanceGroupArgs struct {
	// An optional textual description of the instance
	// group.
	Description *string `pulumi:"description"`
	// List of instances in the group. They should be given
	// as either selfLink or id. When adding instances they must all be in the same
	// network and zone as the instance group.
	Instances []string `pulumi:"instances"`
	// The name which the port will be mapped to.
	Name *string `pulumi:"name"`
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts []InstanceGroupNamedPortType `pulumi:"namedPorts"`
	// The URL of the network the instance group is in. If
	// this is different from the network where the instances are in, the creation
	// fails. Defaults to the network where the instances are in (if neither
	// `network` nor `instances` is specified, this field will be blank).
	Network *string `pulumi:"network"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The zone that this instance group should be created in.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceGroup resource.
type InstanceGroupArgs struct {
	// An optional textual description of the instance
	// group.
	Description pulumi.StringPtrInput
	// List of instances in the group. They should be given
	// as either selfLink or id. When adding instances they must all be in the same
	// network and zone as the instance group.
	Instances pulumi.StringArrayInput
	// The name which the port will be mapped to.
	Name pulumi.StringPtrInput
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts InstanceGroupNamedPortTypeArrayInput
	// The URL of the network the instance group is in. If
	// this is different from the network where the instances are in, the creation
	// fails. Defaults to the network where the instances are in (if neither
	// `network` nor `instances` is specified, this field will be blank).
	Network pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The zone that this instance group should be created in.
	Zone pulumi.StringPtrInput
}

func (InstanceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupArgs)(nil)).Elem()
}

type InstanceGroupInput interface {
	pulumi.Input

	ToInstanceGroupOutput() InstanceGroupOutput
	ToInstanceGroupOutputWithContext(ctx context.Context) InstanceGroupOutput
}

func (*InstanceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceGroup)(nil))
}

func (i *InstanceGroup) ToInstanceGroupOutput() InstanceGroupOutput {
	return i.ToInstanceGroupOutputWithContext(context.Background())
}

func (i *InstanceGroup) ToInstanceGroupOutputWithContext(ctx context.Context) InstanceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupOutput)
}

func (i *InstanceGroup) ToInstanceGroupPtrOutput() InstanceGroupPtrOutput {
	return i.ToInstanceGroupPtrOutputWithContext(context.Background())
}

func (i *InstanceGroup) ToInstanceGroupPtrOutputWithContext(ctx context.Context) InstanceGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupPtrOutput)
}

type InstanceGroupPtrInput interface {
	pulumi.Input

	ToInstanceGroupPtrOutput() InstanceGroupPtrOutput
	ToInstanceGroupPtrOutputWithContext(ctx context.Context) InstanceGroupPtrOutput
}

type instanceGroupPtrType InstanceGroupArgs

func (*instanceGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceGroup)(nil))
}

func (i *instanceGroupPtrType) ToInstanceGroupPtrOutput() InstanceGroupPtrOutput {
	return i.ToInstanceGroupPtrOutputWithContext(context.Background())
}

func (i *instanceGroupPtrType) ToInstanceGroupPtrOutputWithContext(ctx context.Context) InstanceGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupPtrOutput)
}

// InstanceGroupArrayInput is an input type that accepts InstanceGroupArray and InstanceGroupArrayOutput values.
// You can construct a concrete instance of `InstanceGroupArrayInput` via:
//
//          InstanceGroupArray{ InstanceGroupArgs{...} }
type InstanceGroupArrayInput interface {
	pulumi.Input

	ToInstanceGroupArrayOutput() InstanceGroupArrayOutput
	ToInstanceGroupArrayOutputWithContext(context.Context) InstanceGroupArrayOutput
}

type InstanceGroupArray []InstanceGroupInput

func (InstanceGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*InstanceGroup)(nil))
}

func (i InstanceGroupArray) ToInstanceGroupArrayOutput() InstanceGroupArrayOutput {
	return i.ToInstanceGroupArrayOutputWithContext(context.Background())
}

func (i InstanceGroupArray) ToInstanceGroupArrayOutputWithContext(ctx context.Context) InstanceGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupArrayOutput)
}

// InstanceGroupMapInput is an input type that accepts InstanceGroupMap and InstanceGroupMapOutput values.
// You can construct a concrete instance of `InstanceGroupMapInput` via:
//
//          InstanceGroupMap{ "key": InstanceGroupArgs{...} }
type InstanceGroupMapInput interface {
	pulumi.Input

	ToInstanceGroupMapOutput() InstanceGroupMapOutput
	ToInstanceGroupMapOutputWithContext(context.Context) InstanceGroupMapOutput
}

type InstanceGroupMap map[string]InstanceGroupInput

func (InstanceGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*InstanceGroup)(nil))
}

func (i InstanceGroupMap) ToInstanceGroupMapOutput() InstanceGroupMapOutput {
	return i.ToInstanceGroupMapOutputWithContext(context.Background())
}

func (i InstanceGroupMap) ToInstanceGroupMapOutputWithContext(ctx context.Context) InstanceGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupMapOutput)
}

type InstanceGroupOutput struct {
	*pulumi.OutputState
}

func (InstanceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceGroup)(nil))
}

func (o InstanceGroupOutput) ToInstanceGroupOutput() InstanceGroupOutput {
	return o
}

func (o InstanceGroupOutput) ToInstanceGroupOutputWithContext(ctx context.Context) InstanceGroupOutput {
	return o
}

func (o InstanceGroupOutput) ToInstanceGroupPtrOutput() InstanceGroupPtrOutput {
	return o.ToInstanceGroupPtrOutputWithContext(context.Background())
}

func (o InstanceGroupOutput) ToInstanceGroupPtrOutputWithContext(ctx context.Context) InstanceGroupPtrOutput {
	return o.ApplyT(func(v InstanceGroup) *InstanceGroup {
		return &v
	}).(InstanceGroupPtrOutput)
}

type InstanceGroupPtrOutput struct {
	*pulumi.OutputState
}

func (InstanceGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceGroup)(nil))
}

func (o InstanceGroupPtrOutput) ToInstanceGroupPtrOutput() InstanceGroupPtrOutput {
	return o
}

func (o InstanceGroupPtrOutput) ToInstanceGroupPtrOutputWithContext(ctx context.Context) InstanceGroupPtrOutput {
	return o
}

type InstanceGroupArrayOutput struct{ *pulumi.OutputState }

func (InstanceGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceGroup)(nil))
}

func (o InstanceGroupArrayOutput) ToInstanceGroupArrayOutput() InstanceGroupArrayOutput {
	return o
}

func (o InstanceGroupArrayOutput) ToInstanceGroupArrayOutputWithContext(ctx context.Context) InstanceGroupArrayOutput {
	return o
}

func (o InstanceGroupArrayOutput) Index(i pulumi.IntInput) InstanceGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceGroup {
		return vs[0].([]InstanceGroup)[vs[1].(int)]
	}).(InstanceGroupOutput)
}

type InstanceGroupMapOutput struct{ *pulumi.OutputState }

func (InstanceGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InstanceGroup)(nil))
}

func (o InstanceGroupMapOutput) ToInstanceGroupMapOutput() InstanceGroupMapOutput {
	return o
}

func (o InstanceGroupMapOutput) ToInstanceGroupMapOutputWithContext(ctx context.Context) InstanceGroupMapOutput {
	return o
}

func (o InstanceGroupMapOutput) MapIndex(k pulumi.StringInput) InstanceGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InstanceGroup {
		return vs[0].(map[string]InstanceGroup)[vs[1].(string)]
	}).(InstanceGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceGroupOutput{})
	pulumi.RegisterOutputType(InstanceGroupPtrOutput{})
	pulumi.RegisterOutputType(InstanceGroupArrayOutput{})
	pulumi.RegisterOutputType(InstanceGroupMapOutput{})
}
