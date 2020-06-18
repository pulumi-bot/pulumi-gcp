// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package binaryauthorization

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Binary Authorization Attestor. Each of these resources serves a different use case:
//
// * `binaryauthorization.AttestorIamPolicy`: Authoritative. Sets the IAM policy for the attestor and replaces any existing policy already attached.
// * `binaryauthorization.AttestorIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the attestor are preserved.
// * `binaryauthorization.AttestorIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the attestor are preserved.
//
// > **Note:** `binaryauthorization.AttestorIamPolicy` **cannot** be used in conjunction with `binaryauthorization.AttestorIamBinding` and `binaryauthorization.AttestorIamMember` or they will fight over what your policy should be.
//
// > **Note:** `binaryauthorization.AttestorIamBinding` resources **can be** used in conjunction with `binaryauthorization.AttestorIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_binary\_authorization\_attestor\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/binaryauthorization"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Binding: []map[string]interface{}{
// 				map[string]interface{}{
// 					"role": "roles/viewer",
// 					"members": []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = binaryauthorization.NewAttestorIamPolicy(ctx, "policy", &binaryauthorization.AttestorIamPolicyArgs{
// 			Project:    pulumi.String(google_binary_authorization_attestor.Attestor.Project),
// 			Attestor:   pulumi.String(google_binary_authorization_attestor.Attestor.Name),
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
// ## google\_binary\_authorization\_attestor\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/binaryauthorization"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = binaryauthorization.NewAttestorIamBinding(ctx, "binding", &binaryauthorization.AttestorIamBindingArgs{
// 			Project:  pulumi.String(google_binary_authorization_attestor.Attestor.Project),
// 			Attestor: pulumi.String(google_binary_authorization_attestor.Attestor.Name),
// 			Role:     pulumi.String("roles/viewer"),
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
// ## google\_binary\_authorization\_attestor\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/binaryauthorization"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = binaryauthorization.NewAttestorIamMember(ctx, "member", &binaryauthorization.AttestorIamMemberArgs{
// 			Project:  pulumi.String(google_binary_authorization_attestor.Attestor.Project),
// 			Attestor: pulumi.String(google_binary_authorization_attestor.Attestor.Name),
// 			Role:     pulumi.String("roles/viewer"),
// 			Member:   pulumi.String("user:jane@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AttestorIamMember struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	Attestor  pulumi.StringOutput                 `pulumi:"attestor"`
	Condition AttestorIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewAttestorIamMember registers a new resource with the given unique name, arguments, and options.
func NewAttestorIamMember(ctx *pulumi.Context,
	name string, args *AttestorIamMemberArgs, opts ...pulumi.ResourceOption) (*AttestorIamMember, error) {
	if args == nil || args.Attestor == nil {
		return nil, errors.New("missing required argument 'Attestor'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &AttestorIamMemberArgs{}
	}
	var resource AttestorIamMember
	err := ctx.RegisterResource("gcp:binaryauthorization/attestorIamMember:AttestorIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttestorIamMember gets an existing AttestorIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttestorIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttestorIamMemberState, opts ...pulumi.ResourceOption) (*AttestorIamMember, error) {
	var resource AttestorIamMember
	err := ctx.ReadResource("gcp:binaryauthorization/attestorIamMember:AttestorIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttestorIamMember resources.
type attestorIamMemberState struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor  *string                     `pulumi:"attestor"`
	Condition *AttestorIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type AttestorIamMemberState struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor  pulumi.StringPtrInput
	Condition AttestorIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (AttestorIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*attestorIamMemberState)(nil)).Elem()
}

type attestorIamMemberArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor  string                      `pulumi:"attestor"`
	Condition *AttestorIamMemberCondition `pulumi:"condition"`
	Member    string                      `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a AttestorIamMember resource.
type AttestorIamMemberArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor  pulumi.StringInput
	Condition AttestorIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (AttestorIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attestorIamMemberArgs)(nil)).Elem()
}
