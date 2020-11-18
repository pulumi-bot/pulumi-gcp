// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type MachineImageIamBinding struct {
	pulumi.CustomResourceState

	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition MachineImageIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	MachineImage pulumi.StringOutput      `pulumi:"machineImage"`
	Members      pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewMachineImageIamBinding registers a new resource with the given unique name, arguments, and options.
func NewMachineImageIamBinding(ctx *pulumi.Context,
	name string, args *MachineImageIamBindingArgs, opts ...pulumi.ResourceOption) (*MachineImageIamBinding, error) {
	if args == nil || args.MachineImage == nil {
		return nil, errors.New("missing required argument 'MachineImage'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &MachineImageIamBindingArgs{}
	}
	var resource MachineImageIamBinding
	err := ctx.RegisterResource("gcp:compute/machineImageIamBinding:MachineImageIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineImageIamBinding gets an existing MachineImageIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineImageIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineImageIamBindingState, opts ...pulumi.ResourceOption) (*MachineImageIamBinding, error) {
	var resource MachineImageIamBinding
	err := ctx.ReadResource("gcp:compute/machineImageIamBinding:MachineImageIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineImageIamBinding resources.
type machineImageIamBindingState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *MachineImageIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	MachineImage *string  `pulumi:"machineImage"`
	Members      []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type MachineImageIamBindingState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition MachineImageIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	MachineImage pulumi.StringPtrInput
	Members      pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (MachineImageIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineImageIamBindingState)(nil)).Elem()
}

type machineImageIamBindingArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *MachineImageIamBindingCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	MachineImage string   `pulumi:"machineImage"`
	Members      []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a MachineImageIamBinding resource.
type MachineImageIamBindingArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition MachineImageIamBindingConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	MachineImage pulumi.StringInput
	Members      pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (MachineImageIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineImageIamBindingArgs)(nil)).Elem()
}

type MachineImageIamBindingInput interface {
	pulumi.Input

	ToMachineImageIamBindingOutput() MachineImageIamBindingOutput
	ToMachineImageIamBindingOutputWithContext(ctx context.Context) MachineImageIamBindingOutput
}

func (MachineImageIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineImageIamBinding)(nil)).Elem()
}

func (i MachineImageIamBinding) ToMachineImageIamBindingOutput() MachineImageIamBindingOutput {
	return i.ToMachineImageIamBindingOutputWithContext(context.Background())
}

func (i MachineImageIamBinding) ToMachineImageIamBindingOutputWithContext(ctx context.Context) MachineImageIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImageIamBindingOutput)
}

type MachineImageIamBindingOutput struct {
	*pulumi.OutputState
}

func (MachineImageIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineImageIamBindingOutput)(nil)).Elem()
}

func (o MachineImageIamBindingOutput) ToMachineImageIamBindingOutput() MachineImageIamBindingOutput {
	return o
}

func (o MachineImageIamBindingOutput) ToMachineImageIamBindingOutputWithContext(ctx context.Context) MachineImageIamBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MachineImageIamBindingOutput{})
}
