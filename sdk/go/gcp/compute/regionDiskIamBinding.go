// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Compute Engine Disk. Each of these resources serves a different use case:
//
// * `compute.DiskIamPolicy`: Authoritative. Sets the IAM policy for the disk and replaces any existing policy already attached.
// * `compute.DiskIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the disk are preserved.
// * `compute.DiskIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the disk are preserved.
//
// > **Note:** `compute.DiskIamPolicy` **cannot** be used in conjunction with `compute.DiskIamBinding` and `compute.DiskIamMember` or they will fight over what your policy should be.
//
// > **Note:** `compute.DiskIamBinding` resources **can be** used in conjunction with `compute.DiskIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/zones/{{zone}}/disks/{{name}} * {{project}}/{{zone}}/{{name}} * {{zone}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine disk IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/regionDiskIamBinding:RegionDiskIamBinding editor "projects/{{project}}/zones/{{zone}}/disks/{{disk}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/regionDiskIamBinding:RegionDiskIamBinding editor "projects/{{project}}/zones/{{zone}}/disks/{{disk}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/regionDiskIamBinding:RegionDiskIamBinding editor projects/{{project}}/zones/{{zone}}/disks/{{disk}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type RegionDiskIamBinding struct {
	pulumi.CustomResourceState

	Condition RegionDiskIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	Region  pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `compute.DiskIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewRegionDiskIamBinding registers a new resource with the given unique name, arguments, and options.
func NewRegionDiskIamBinding(ctx *pulumi.Context,
	name string, args *RegionDiskIamBindingArgs, opts ...pulumi.ResourceOption) (*RegionDiskIamBinding, error) {
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &RegionDiskIamBindingArgs{}
	}
	var resource RegionDiskIamBinding
	err := ctx.RegisterResource("gcp:compute/regionDiskIamBinding:RegionDiskIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionDiskIamBinding gets an existing RegionDiskIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionDiskIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionDiskIamBindingState, opts ...pulumi.ResourceOption) (*RegionDiskIamBinding, error) {
	var resource RegionDiskIamBinding
	err := ctx.ReadResource("gcp:compute/regionDiskIamBinding:RegionDiskIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionDiskIamBinding resources.
type regionDiskIamBindingState struct {
	Condition *RegionDiskIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `compute.DiskIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type RegionDiskIamBindingState struct {
	Condition RegionDiskIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.DiskIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (RegionDiskIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionDiskIamBindingState)(nil)).Elem()
}

type regionDiskIamBindingArgs struct {
	Condition *RegionDiskIamBindingCondition `pulumi:"condition"`
	Members   []string                       `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `compute.DiskIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a RegionDiskIamBinding resource.
type RegionDiskIamBindingArgs struct {
	Condition RegionDiskIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.DiskIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (RegionDiskIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionDiskIamBindingArgs)(nil)).Elem()
}

type RegionDiskIamBindingInput interface {
	pulumi.Input

	ToRegionDiskIamBindingOutput() RegionDiskIamBindingOutput
	ToRegionDiskIamBindingOutputWithContext(ctx context.Context) RegionDiskIamBindingOutput
}

func (RegionDiskIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionDiskIamBinding)(nil)).Elem()
}

func (i RegionDiskIamBinding) ToRegionDiskIamBindingOutput() RegionDiskIamBindingOutput {
	return i.ToRegionDiskIamBindingOutputWithContext(context.Background())
}

func (i RegionDiskIamBinding) ToRegionDiskIamBindingOutputWithContext(ctx context.Context) RegionDiskIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionDiskIamBindingOutput)
}

type RegionDiskIamBindingOutput struct {
	*pulumi.OutputState
}

func (RegionDiskIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionDiskIamBindingOutput)(nil)).Elem()
}

func (o RegionDiskIamBindingOutput) ToRegionDiskIamBindingOutput() RegionDiskIamBindingOutput {
	return o
}

func (o RegionDiskIamBindingOutput) ToRegionDiskIamBindingOutputWithContext(ctx context.Context) RegionDiskIamBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegionDiskIamBindingOutput{})
}
