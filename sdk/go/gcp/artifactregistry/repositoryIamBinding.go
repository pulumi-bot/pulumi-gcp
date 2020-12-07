// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactregistry

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Artifact Registry Repository. Each of these resources serves a different use case:
//
// * `artifactregistry.RepositoryIamPolicy`: Authoritative. Sets the IAM policy for the repository and replaces any existing policy already attached.
// * `artifactregistry.RepositoryIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the repository are preserved.
// * `artifactregistry.RepositoryIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the repository are preserved.
//
// > **Note:** `artifactregistry.RepositoryIamPolicy` **cannot** be used in conjunction with `artifactregistry.RepositoryIamBinding` and `artifactregistry.RepositoryIamMember` or they will fight over what your policy should be.
//
// > **Note:** `artifactregistry.RepositoryIamBinding` resources **can be** used in conjunction with `artifactregistry.RepositoryIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_artifact\_registry\_repository\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/artifactregistry"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 		_, err = artifactregistry.NewRepositoryIamPolicy(ctx, "policy", &artifactregistry.RepositoryIamPolicyArgs{
// 			Project:    pulumi.Any(google_artifact_registry_repository.My - repo.Project),
// 			Location:   pulumi.Any(google_artifact_registry_repository.My - repo.Location),
// 			Repository: pulumi.Any(google_artifact_registry_repository.My - repo.Name),
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
// ## google\_artifact\_registry\_repository\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/artifactregistry"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactregistry.NewRepositoryIamBinding(ctx, "binding", &artifactregistry.RepositoryIamBindingArgs{
// 			Project:    pulumi.Any(google_artifact_registry_repository.My - repo.Project),
// 			Location:   pulumi.Any(google_artifact_registry_repository.My - repo.Location),
// 			Repository: pulumi.Any(google_artifact_registry_repository.My - repo.Name),
// 			Role:       pulumi.String("roles/viewer"),
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
// ## google\_artifact\_registry\_repository\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/artifactregistry"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactregistry.NewRepositoryIamMember(ctx, "member", &artifactregistry.RepositoryIamMemberArgs{
// 			Project:    pulumi.Any(google_artifact_registry_repository.My - repo.Project),
// 			Location:   pulumi.Any(google_artifact_registry_repository.My - repo.Location),
// 			Repository: pulumi.Any(google_artifact_registry_repository.My - repo.Name),
// 			Role:       pulumi.String("roles/viewer"),
// 			Member:     pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/repositories/{{repository}} * {{project}}/{{location}}/{{repository}} * {{location}}/{{repository}} * {{repository}} Any variables not passed in the import command will be taken from the provider configuration. Artifact Registry repository IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:artifactregistry/repositoryIamBinding:RepositoryIamBinding editor "projects/{{project}}/locations/{{location}}/repositories/{{repository}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:artifactregistry/repositoryIamBinding:RepositoryIamBinding editor "projects/{{project}}/locations/{{location}}/repositories/{{repository}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:artifactregistry/repositoryIamBinding:RepositoryIamBinding editor projects/{{project}}/locations/{{location}}/repositories/{{repository}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type RepositoryIamBinding struct {
	pulumi.CustomResourceState

	Condition RepositoryIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the location this repository is located in.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput      `pulumi:"location"`
	Members  pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	Repository pulumi.StringOutput `pulumi:"repository"`
	// The role that should be applied. Only one
	// `artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewRepositoryIamBinding registers a new resource with the given unique name, arguments, and options.
func NewRepositoryIamBinding(ctx *pulumi.Context,
	name string, args *RepositoryIamBindingArgs, opts ...pulumi.ResourceOption) (*RepositoryIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource RepositoryIamBinding
	err := ctx.RegisterResource("gcp:artifactregistry/repositoryIamBinding:RepositoryIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryIamBinding gets an existing RepositoryIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryIamBindingState, opts ...pulumi.ResourceOption) (*RepositoryIamBinding, error) {
	var resource RepositoryIamBinding
	err := ctx.ReadResource("gcp:artifactregistry/repositoryIamBinding:RepositoryIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryIamBinding resources.
type repositoryIamBindingState struct {
	Condition *RepositoryIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The name of the location this repository is located in.
	// Used to find the parent resource to bind the IAM policy to
	Location *string  `pulumi:"location"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	Repository *string `pulumi:"repository"`
	// The role that should be applied. Only one
	// `artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type RepositoryIamBindingState struct {
	Condition RepositoryIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The name of the location this repository is located in.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Repository pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (RepositoryIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryIamBindingState)(nil)).Elem()
}

type repositoryIamBindingArgs struct {
	Condition *RepositoryIamBindingCondition `pulumi:"condition"`
	// The name of the location this repository is located in.
	// Used to find the parent resource to bind the IAM policy to
	Location *string  `pulumi:"location"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	Repository string `pulumi:"repository"`
	// The role that should be applied. Only one
	// `artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a RepositoryIamBinding resource.
type RepositoryIamBindingArgs struct {
	Condition RepositoryIamBindingConditionPtrInput
	// The name of the location this repository is located in.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Repository pulumi.StringInput
	// The role that should be applied. Only one
	// `artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (RepositoryIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryIamBindingArgs)(nil)).Elem()
}

type RepositoryIamBindingInput interface {
	pulumi.Input

	ToRepositoryIamBindingOutput() RepositoryIamBindingOutput
	ToRepositoryIamBindingOutputWithContext(ctx context.Context) RepositoryIamBindingOutput
}

func (RepositoryIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryIamBinding)(nil)).Elem()
}

func (i RepositoryIamBinding) ToRepositoryIamBindingOutput() RepositoryIamBindingOutput {
	return i.ToRepositoryIamBindingOutputWithContext(context.Background())
}

func (i RepositoryIamBinding) ToRepositoryIamBindingOutputWithContext(ctx context.Context) RepositoryIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamBindingOutput)
}

type RepositoryIamBindingOutput struct {
	*pulumi.OutputState
}

func (RepositoryIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryIamBindingOutput)(nil)).Elem()
}

func (o RepositoryIamBindingOutput) ToRepositoryIamBindingOutput() RepositoryIamBindingOutput {
	return o
}

func (o RepositoryIamBindingOutput) ToRepositoryIamBindingOutputWithContext(ctx context.Context) RepositoryIamBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RepositoryIamBindingOutput{})
}
