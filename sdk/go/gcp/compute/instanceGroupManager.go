// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Google Compute Engine Instance Group Manager API creates and manages pools
// of homogeneous Compute Engine virtual machine instances from a common instance
// template. For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/manager)
// and [API](https://cloud.google.com/compute/docs/reference/latest/instanceGroupManagers)
//
// > **Note:** Use [compute.RegionInstanceGroupManager](https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager.html) to create a regional (multi-zone) instance group manager.
//
// ## Example Usage
// ### With Multiple Versions (`Google-Beta` Provider)
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
// 		_, err := compute.NewInstanceGroupManager(ctx, "appserver", &compute.InstanceGroupManagerArgs{
// 			BaseInstanceName: pulumi.String("app"),
// 			Zone:             pulumi.String("us-central1-a"),
// 			TargetSize:       pulumi.Int(5),
// 			Versions: compute.InstanceGroupManagerVersionArray{
// 				&compute.InstanceGroupManagerVersionArgs{
// 					Name:             pulumi.String("appserver"),
// 					InstanceTemplate: pulumi.Any(google_compute_instance_template.Appserver.Id),
// 				},
// 				&compute.InstanceGroupManagerVersionArgs{
// 					Name:             pulumi.String("appserver-canary"),
// 					InstanceTemplate: pulumi.Any(google_compute_instance_template.Appserver - canary.Id),
// 					TargetSize: &compute.InstanceGroupManagerVersionTargetSizeArgs{
// 						Fixed: pulumi.Int(1),
// 					},
// 				},
// 			},
// 		}, pulumi.Provider(google_beta))
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
// Instance group managers can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager appserver projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager appserver {{project}}/{{zone}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager appserver {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager appserver {{name}}
// ```
type InstanceGroupManager struct {
	pulumi.CustomResourceState

	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies InstanceGroupManagerAutoHealingPoliciesPtrOutput `pulumi:"autoHealingPolicies"`
	// The base instance name to use for
	// instances in this group. The value must be a valid
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
	// are lowercase letters, numbers, and hyphens (-). Instances are named by
	// appending a hyphen and a random four-character string to the base instance
	// name.
	BaseInstanceName pulumi.StringOutput `pulumi:"baseInstanceName"`
	// An optional textual description of the instance
	// group manager.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The fingerprint of the instance group manager.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The full URL of the instance group created by the manager.
	InstanceGroup pulumi.StringOutput `pulumi:"instanceGroup"`
	// - Version name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts InstanceGroupManagerNamedPortArrayOutput `pulumi:"namedPorts"`
	Operation  pulumi.StringOutput                      `pulumi:"operation"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URL of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
	StatefulDisks InstanceGroupManagerStatefulDiskArrayOutput `pulumi:"statefulDisks"`
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools pulumi.StringArrayOutput `pulumi:"targetPools"`
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize pulumi.IntOutput `pulumi:"targetSize"`
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
	UpdatePolicy InstanceGroupManagerUpdatePolicyOutput `pulumi:"updatePolicy"`
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions InstanceGroupManagerVersionArrayOutput `pulumi:"versions"`
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, this provider will
	// continue trying until it times out.
	WaitForInstances pulumi.BoolPtrOutput `pulumi:"waitForInstances"`
	// The zone that instances in this group should be created
	// in.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceGroupManager registers a new resource with the given unique name, arguments, and options.
func NewInstanceGroupManager(ctx *pulumi.Context,
	name string, args *InstanceGroupManagerArgs, opts ...pulumi.ResourceOption) (*InstanceGroupManager, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BaseInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'BaseInstanceName'")
	}
	if args.Versions == nil {
		return nil, errors.New("invalid value for required argument 'Versions'")
	}
	var resource InstanceGroupManager
	err := ctx.RegisterResource("gcp:compute/instanceGroupManager:InstanceGroupManager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceGroupManager gets an existing InstanceGroupManager resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceGroupManager(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceGroupManagerState, opts ...pulumi.ResourceOption) (*InstanceGroupManager, error) {
	var resource InstanceGroupManager
	err := ctx.ReadResource("gcp:compute/instanceGroupManager:InstanceGroupManager", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceGroupManager resources.
type instanceGroupManagerState struct {
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies *InstanceGroupManagerAutoHealingPolicies `pulumi:"autoHealingPolicies"`
	// The base instance name to use for
	// instances in this group. The value must be a valid
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
	// are lowercase letters, numbers, and hyphens (-). Instances are named by
	// appending a hyphen and a random four-character string to the base instance
	// name.
	BaseInstanceName *string `pulumi:"baseInstanceName"`
	// An optional textual description of the instance
	// group manager.
	Description *string `pulumi:"description"`
	// The fingerprint of the instance group manager.
	Fingerprint *string `pulumi:"fingerprint"`
	// The full URL of the instance group created by the manager.
	InstanceGroup *string `pulumi:"instanceGroup"`
	// - Version name.
	Name *string `pulumi:"name"`
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts []InstanceGroupManagerNamedPort `pulumi:"namedPorts"`
	Operation  *string                         `pulumi:"operation"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URL of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
	StatefulDisks []InstanceGroupManagerStatefulDisk `pulumi:"statefulDisks"`
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools []string `pulumi:"targetPools"`
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize *int `pulumi:"targetSize"`
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
	UpdatePolicy *InstanceGroupManagerUpdatePolicy `pulumi:"updatePolicy"`
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions []InstanceGroupManagerVersion `pulumi:"versions"`
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, this provider will
	// continue trying until it times out.
	WaitForInstances *bool `pulumi:"waitForInstances"`
	// The zone that instances in this group should be created
	// in.
	Zone *string `pulumi:"zone"`
}

type InstanceGroupManagerState struct {
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies InstanceGroupManagerAutoHealingPoliciesPtrInput
	// The base instance name to use for
	// instances in this group. The value must be a valid
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
	// are lowercase letters, numbers, and hyphens (-). Instances are named by
	// appending a hyphen and a random four-character string to the base instance
	// name.
	BaseInstanceName pulumi.StringPtrInput
	// An optional textual description of the instance
	// group manager.
	Description pulumi.StringPtrInput
	// The fingerprint of the instance group manager.
	Fingerprint pulumi.StringPtrInput
	// The full URL of the instance group created by the manager.
	InstanceGroup pulumi.StringPtrInput
	// - Version name.
	Name pulumi.StringPtrInput
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts InstanceGroupManagerNamedPortArrayInput
	Operation  pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URL of the created resource.
	SelfLink pulumi.StringPtrInput
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
	StatefulDisks InstanceGroupManagerStatefulDiskArrayInput
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools pulumi.StringArrayInput
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize pulumi.IntPtrInput
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
	UpdatePolicy InstanceGroupManagerUpdatePolicyPtrInput
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions InstanceGroupManagerVersionArrayInput
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, this provider will
	// continue trying until it times out.
	WaitForInstances pulumi.BoolPtrInput
	// The zone that instances in this group should be created
	// in.
	Zone pulumi.StringPtrInput
}

func (InstanceGroupManagerState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupManagerState)(nil)).Elem()
}

type instanceGroupManagerArgs struct {
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies *InstanceGroupManagerAutoHealingPolicies `pulumi:"autoHealingPolicies"`
	// The base instance name to use for
	// instances in this group. The value must be a valid
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
	// are lowercase letters, numbers, and hyphens (-). Instances are named by
	// appending a hyphen and a random four-character string to the base instance
	// name.
	BaseInstanceName string `pulumi:"baseInstanceName"`
	// An optional textual description of the instance
	// group manager.
	Description *string `pulumi:"description"`
	// - Version name.
	Name *string `pulumi:"name"`
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts []InstanceGroupManagerNamedPort `pulumi:"namedPorts"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
	StatefulDisks []InstanceGroupManagerStatefulDisk `pulumi:"statefulDisks"`
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools []string `pulumi:"targetPools"`
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize *int `pulumi:"targetSize"`
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
	UpdatePolicy *InstanceGroupManagerUpdatePolicy `pulumi:"updatePolicy"`
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions []InstanceGroupManagerVersion `pulumi:"versions"`
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, this provider will
	// continue trying until it times out.
	WaitForInstances *bool `pulumi:"waitForInstances"`
	// The zone that instances in this group should be created
	// in.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceGroupManager resource.
type InstanceGroupManagerArgs struct {
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies InstanceGroupManagerAutoHealingPoliciesPtrInput
	// The base instance name to use for
	// instances in this group. The value must be a valid
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
	// are lowercase letters, numbers, and hyphens (-). Instances are named by
	// appending a hyphen and a random four-character string to the base instance
	// name.
	BaseInstanceName pulumi.StringInput
	// An optional textual description of the instance
	// group manager.
	Description pulumi.StringPtrInput
	// - Version name.
	Name pulumi.StringPtrInput
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts InstanceGroupManagerNamedPortArrayInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
	StatefulDisks InstanceGroupManagerStatefulDiskArrayInput
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools pulumi.StringArrayInput
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize pulumi.IntPtrInput
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
	UpdatePolicy InstanceGroupManagerUpdatePolicyPtrInput
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions InstanceGroupManagerVersionArrayInput
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, this provider will
	// continue trying until it times out.
	WaitForInstances pulumi.BoolPtrInput
	// The zone that instances in this group should be created
	// in.
	Zone pulumi.StringPtrInput
}

func (InstanceGroupManagerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupManagerArgs)(nil)).Elem()
}

type InstanceGroupManagerInput interface {
	pulumi.Input

	ToInstanceGroupManagerOutput() InstanceGroupManagerOutput
	ToInstanceGroupManagerOutputWithContext(ctx context.Context) InstanceGroupManagerOutput
}

func (*InstanceGroupManager) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceGroupManager)(nil))
}

func (i *InstanceGroupManager) ToInstanceGroupManagerOutput() InstanceGroupManagerOutput {
	return i.ToInstanceGroupManagerOutputWithContext(context.Background())
}

func (i *InstanceGroupManager) ToInstanceGroupManagerOutputWithContext(ctx context.Context) InstanceGroupManagerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupManagerOutput)
}

func (i *InstanceGroupManager) ToInstanceGroupManagerPtrOutput() InstanceGroupManagerPtrOutput {
	return i.ToInstanceGroupManagerPtrOutputWithContext(context.Background())
}

func (i *InstanceGroupManager) ToInstanceGroupManagerPtrOutputWithContext(ctx context.Context) InstanceGroupManagerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupManagerPtrOutput)
}

type InstanceGroupManagerPtrInput interface {
	pulumi.Input

	ToInstanceGroupManagerPtrOutput() InstanceGroupManagerPtrOutput
	ToInstanceGroupManagerPtrOutputWithContext(ctx context.Context) InstanceGroupManagerPtrOutput
}

type instanceGroupManagerPtrType InstanceGroupManagerArgs

func (*instanceGroupManagerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceGroupManager)(nil))
}

func (i *instanceGroupManagerPtrType) ToInstanceGroupManagerPtrOutput() InstanceGroupManagerPtrOutput {
	return i.ToInstanceGroupManagerPtrOutputWithContext(context.Background())
}

func (i *instanceGroupManagerPtrType) ToInstanceGroupManagerPtrOutputWithContext(ctx context.Context) InstanceGroupManagerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupManagerPtrOutput)
}

// InstanceGroupManagerArrayInput is an input type that accepts InstanceGroupManagerArray and InstanceGroupManagerArrayOutput values.
// You can construct a concrete instance of `InstanceGroupManagerArrayInput` via:
//
//          InstanceGroupManagerArray{ InstanceGroupManagerArgs{...} }
type InstanceGroupManagerArrayInput interface {
	pulumi.Input

	ToInstanceGroupManagerArrayOutput() InstanceGroupManagerArrayOutput
	ToInstanceGroupManagerArrayOutputWithContext(context.Context) InstanceGroupManagerArrayOutput
}

type InstanceGroupManagerArray []InstanceGroupManagerInput

func (InstanceGroupManagerArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*InstanceGroupManager)(nil))
}

func (i InstanceGroupManagerArray) ToInstanceGroupManagerArrayOutput() InstanceGroupManagerArrayOutput {
	return i.ToInstanceGroupManagerArrayOutputWithContext(context.Background())
}

func (i InstanceGroupManagerArray) ToInstanceGroupManagerArrayOutputWithContext(ctx context.Context) InstanceGroupManagerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupManagerArrayOutput)
}

// InstanceGroupManagerMapInput is an input type that accepts InstanceGroupManagerMap and InstanceGroupManagerMapOutput values.
// You can construct a concrete instance of `InstanceGroupManagerMapInput` via:
//
//          InstanceGroupManagerMap{ "key": InstanceGroupManagerArgs{...} }
type InstanceGroupManagerMapInput interface {
	pulumi.Input

	ToInstanceGroupManagerMapOutput() InstanceGroupManagerMapOutput
	ToInstanceGroupManagerMapOutputWithContext(context.Context) InstanceGroupManagerMapOutput
}

type InstanceGroupManagerMap map[string]InstanceGroupManagerInput

func (InstanceGroupManagerMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*InstanceGroupManager)(nil))
}

func (i InstanceGroupManagerMap) ToInstanceGroupManagerMapOutput() InstanceGroupManagerMapOutput {
	return i.ToInstanceGroupManagerMapOutputWithContext(context.Background())
}

func (i InstanceGroupManagerMap) ToInstanceGroupManagerMapOutputWithContext(ctx context.Context) InstanceGroupManagerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupManagerMapOutput)
}

type InstanceGroupManagerOutput struct {
	*pulumi.OutputState
}

func (InstanceGroupManagerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceGroupManager)(nil))
}

func (o InstanceGroupManagerOutput) ToInstanceGroupManagerOutput() InstanceGroupManagerOutput {
	return o
}

func (o InstanceGroupManagerOutput) ToInstanceGroupManagerOutputWithContext(ctx context.Context) InstanceGroupManagerOutput {
	return o
}

func (o InstanceGroupManagerOutput) ToInstanceGroupManagerPtrOutput() InstanceGroupManagerPtrOutput {
	return o.ToInstanceGroupManagerPtrOutputWithContext(context.Background())
}

func (o InstanceGroupManagerOutput) ToInstanceGroupManagerPtrOutputWithContext(ctx context.Context) InstanceGroupManagerPtrOutput {
	return o.ApplyT(func(v InstanceGroupManager) *InstanceGroupManager {
		return &v
	}).(InstanceGroupManagerPtrOutput)
}

type InstanceGroupManagerPtrOutput struct {
	*pulumi.OutputState
}

func (InstanceGroupManagerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceGroupManager)(nil))
}

func (o InstanceGroupManagerPtrOutput) ToInstanceGroupManagerPtrOutput() InstanceGroupManagerPtrOutput {
	return o
}

func (o InstanceGroupManagerPtrOutput) ToInstanceGroupManagerPtrOutputWithContext(ctx context.Context) InstanceGroupManagerPtrOutput {
	return o
}

type InstanceGroupManagerArrayOutput struct{ *pulumi.OutputState }

func (InstanceGroupManagerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceGroupManager)(nil))
}

func (o InstanceGroupManagerArrayOutput) ToInstanceGroupManagerArrayOutput() InstanceGroupManagerArrayOutput {
	return o
}

func (o InstanceGroupManagerArrayOutput) ToInstanceGroupManagerArrayOutputWithContext(ctx context.Context) InstanceGroupManagerArrayOutput {
	return o
}

func (o InstanceGroupManagerArrayOutput) Index(i pulumi.IntInput) InstanceGroupManagerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceGroupManager {
		return vs[0].([]InstanceGroupManager)[vs[1].(int)]
	}).(InstanceGroupManagerOutput)
}

type InstanceGroupManagerMapOutput struct{ *pulumi.OutputState }

func (InstanceGroupManagerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InstanceGroupManager)(nil))
}

func (o InstanceGroupManagerMapOutput) ToInstanceGroupManagerMapOutput() InstanceGroupManagerMapOutput {
	return o
}

func (o InstanceGroupManagerMapOutput) ToInstanceGroupManagerMapOutputWithContext(ctx context.Context) InstanceGroupManagerMapOutput {
	return o
}

func (o InstanceGroupManagerMapOutput) MapIndex(k pulumi.StringInput) InstanceGroupManagerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InstanceGroupManager {
		return vs[0].(map[string]InstanceGroupManager)[vs[1].(string)]
	}).(InstanceGroupManagerOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceGroupManagerOutput{})
	pulumi.RegisterOutputType(InstanceGroupManagerPtrOutput{})
	pulumi.RegisterOutputType(InstanceGroupManagerArrayOutput{})
	pulumi.RegisterOutputType(InstanceGroupManagerMapOutput{})
}
