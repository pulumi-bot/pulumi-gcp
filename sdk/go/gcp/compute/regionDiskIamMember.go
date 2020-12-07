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
// ## google\_compute\_disk\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/viewer",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewDiskIamPolicy(ctx, "policy", &compute.DiskIamPolicyArgs{
// 			Project:    pulumi.Any(google_compute_disk.Default.Project),
// 			Zone:       pulumi.Any(google_compute_disk.Default.Zone),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_compute\_disk\_iam\_binding
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
// 		_, err := compute.NewDiskIamBinding(ctx, "binding", &compute.DiskIamBindingArgs{
// 			Project: pulumi.Any(google_compute_disk.Default.Project),
// 			Zone:    pulumi.Any(google_compute_disk.Default.Zone),
// 			Role:    pulumi.String("roles/viewer"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
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
// ## google\_compute\_disk\_iam\_member
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
// 		_, err := compute.NewDiskIamMember(ctx, "member", &compute.DiskIamMemberArgs{
// 			Project: pulumi.Any(google_compute_disk.Default.Project),
// 			Zone:    pulumi.Any(google_compute_disk.Default.Zone),
// 			Role:    pulumi.String("roles/viewer"),
// 			Member:  pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/zones/{{zone}}/disks/{{name}} * {{project}}/{{zone}}/{{name}} * {{zone}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine disk IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/regionDiskIamMember:RegionDiskIamMember editor "projects/{{project}}/zones/{{zone}}/disks/{{disk}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/regionDiskIamMember:RegionDiskIamMember editor "projects/{{project}}/zones/{{zone}}/disks/{{disk}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/regionDiskIamMember:RegionDiskIamMember editor projects/{{project}}/zones/{{zone}}/disks/{{disk}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type RegionDiskIamMember struct {
	pulumi.CustomResourceState

	Condition RegionDiskIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
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

// NewRegionDiskIamMember registers a new resource with the given unique name, arguments, and options.
func NewRegionDiskIamMember(ctx *pulumi.Context,
	name string, args *RegionDiskIamMemberArgs, opts ...pulumi.ResourceOption) (*RegionDiskIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource RegionDiskIamMember
	err := ctx.RegisterResource("gcp:compute/regionDiskIamMember:RegionDiskIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionDiskIamMember gets an existing RegionDiskIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionDiskIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionDiskIamMemberState, opts ...pulumi.ResourceOption) (*RegionDiskIamMember, error) {
	var resource RegionDiskIamMember
	err := ctx.ReadResource("gcp:compute/regionDiskIamMember:RegionDiskIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionDiskIamMember resources.
type regionDiskIamMemberState struct {
	Condition *RegionDiskIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
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

type RegionDiskIamMemberState struct {
	Condition RegionDiskIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
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

func (RegionDiskIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionDiskIamMemberState)(nil)).Elem()
}

type regionDiskIamMemberArgs struct {
	Condition *RegionDiskIamMemberCondition `pulumi:"condition"`
	Member    string                        `pulumi:"member"`
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

// The set of arguments for constructing a RegionDiskIamMember resource.
type RegionDiskIamMemberArgs struct {
	Condition RegionDiskIamMemberConditionPtrInput
	Member    pulumi.StringInput
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

func (RegionDiskIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionDiskIamMemberArgs)(nil)).Elem()
}

type RegionDiskIamMemberInput interface {
	pulumi.Input

	ToRegionDiskIamMemberOutput() RegionDiskIamMemberOutput
	ToRegionDiskIamMemberOutputWithContext(ctx context.Context) RegionDiskIamMemberOutput
}

func (RegionDiskIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionDiskIamMember)(nil)).Elem()
}

func (i RegionDiskIamMember) ToRegionDiskIamMemberOutput() RegionDiskIamMemberOutput {
	return i.ToRegionDiskIamMemberOutputWithContext(context.Background())
}

func (i RegionDiskIamMember) ToRegionDiskIamMemberOutputWithContext(ctx context.Context) RegionDiskIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionDiskIamMemberOutput)
}

type RegionDiskIamMemberOutput struct {
	*pulumi.OutputState
}

func (RegionDiskIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionDiskIamMemberOutput)(nil)).Elem()
}

func (o RegionDiskIamMemberOutput) ToRegionDiskIamMemberOutput() RegionDiskIamMemberOutput {
	return o
}

func (o RegionDiskIamMemberOutput) ToRegionDiskIamMemberOutputWithContext(ctx context.Context) RegionDiskIamMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegionDiskIamMemberOutput{})
}
