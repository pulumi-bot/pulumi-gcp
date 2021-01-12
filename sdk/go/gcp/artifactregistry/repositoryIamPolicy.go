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
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
//  $ pulumi import gcp:artifactregistry/repositoryIamPolicy:RepositoryIamPolicy editor "projects/{{project}}/locations/{{location}}/repositories/{{repository}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:artifactregistry/repositoryIamPolicy:RepositoryIamPolicy editor "projects/{{project}}/locations/{{location}}/repositories/{{repository}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:artifactregistry/repositoryIamPolicy:RepositoryIamPolicy editor projects/{{project}}/locations/{{location}}/repositories/{{repository}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type RepositoryIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the location this repository is located in.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	Repository pulumi.StringOutput `pulumi:"repository"`
}

// NewRepositoryIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewRepositoryIamPolicy(ctx *pulumi.Context,
	name string, args *RepositoryIamPolicyArgs, opts ...pulumi.ResourceOption) (*RepositoryIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	var resource RepositoryIamPolicy
	err := ctx.RegisterResource("gcp:artifactregistry/repositoryIamPolicy:RepositoryIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryIamPolicy gets an existing RepositoryIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryIamPolicyState, opts ...pulumi.ResourceOption) (*RepositoryIamPolicy, error) {
	var resource RepositoryIamPolicy
	err := ctx.ReadResource("gcp:artifactregistry/repositoryIamPolicy:RepositoryIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryIamPolicy resources.
type repositoryIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The name of the location this repository is located in.
	// Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	Repository *string `pulumi:"repository"`
}

type RepositoryIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The name of the location this repository is located in.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Repository pulumi.StringPtrInput
}

func (RepositoryIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryIamPolicyState)(nil)).Elem()
}

type repositoryIamPolicyArgs struct {
	// The name of the location this repository is located in.
	// Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	Repository string `pulumi:"repository"`
}

// The set of arguments for constructing a RepositoryIamPolicy resource.
type RepositoryIamPolicyArgs struct {
	// The name of the location this repository is located in.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Repository pulumi.StringInput
}

func (RepositoryIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryIamPolicyArgs)(nil)).Elem()
}

type RepositoryIamPolicyInput interface {
	pulumi.Input

	ToRepositoryIamPolicyOutput() RepositoryIamPolicyOutput
	ToRepositoryIamPolicyOutputWithContext(ctx context.Context) RepositoryIamPolicyOutput
}

func (*RepositoryIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryIamPolicy)(nil))
}

func (i *RepositoryIamPolicy) ToRepositoryIamPolicyOutput() RepositoryIamPolicyOutput {
	return i.ToRepositoryIamPolicyOutputWithContext(context.Background())
}

func (i *RepositoryIamPolicy) ToRepositoryIamPolicyOutputWithContext(ctx context.Context) RepositoryIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamPolicyOutput)
}

func (i *RepositoryIamPolicy) ToRepositoryIamPolicyPtrOutput() RepositoryIamPolicyPtrOutput {
	return i.ToRepositoryIamPolicyPtrOutputWithContext(context.Background())
}

func (i *RepositoryIamPolicy) ToRepositoryIamPolicyPtrOutputWithContext(ctx context.Context) RepositoryIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamPolicyPtrOutput)
}

type RepositoryIamPolicyPtrInput interface {
	pulumi.Input

	ToRepositoryIamPolicyPtrOutput() RepositoryIamPolicyPtrOutput
	ToRepositoryIamPolicyPtrOutputWithContext(ctx context.Context) RepositoryIamPolicyPtrOutput
}

type repositoryIamPolicyPtrType RepositoryIamPolicyArgs

func (*repositoryIamPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryIamPolicy)(nil))
}

func (i *repositoryIamPolicyPtrType) ToRepositoryIamPolicyPtrOutput() RepositoryIamPolicyPtrOutput {
	return i.ToRepositoryIamPolicyPtrOutputWithContext(context.Background())
}

func (i *repositoryIamPolicyPtrType) ToRepositoryIamPolicyPtrOutputWithContext(ctx context.Context) RepositoryIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamPolicyOutput).ToRepositoryIamPolicyPtrOutput()
}

// RepositoryIamPolicyArrayInput is an input type that accepts RepositoryIamPolicyArray and RepositoryIamPolicyArrayOutput values.
// You can construct a concrete instance of `RepositoryIamPolicyArrayInput` via:
//
//          RepositoryIamPolicyArray{ RepositoryIamPolicyArgs{...} }
type RepositoryIamPolicyArrayInput interface {
	pulumi.Input

	ToRepositoryIamPolicyArrayOutput() RepositoryIamPolicyArrayOutput
	ToRepositoryIamPolicyArrayOutputWithContext(context.Context) RepositoryIamPolicyArrayOutput
}

type RepositoryIamPolicyArray []RepositoryIamPolicyInput

func (RepositoryIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RepositoryIamPolicy)(nil))
}

func (i RepositoryIamPolicyArray) ToRepositoryIamPolicyArrayOutput() RepositoryIamPolicyArrayOutput {
	return i.ToRepositoryIamPolicyArrayOutputWithContext(context.Background())
}

func (i RepositoryIamPolicyArray) ToRepositoryIamPolicyArrayOutputWithContext(ctx context.Context) RepositoryIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamPolicyArrayOutput)
}

// RepositoryIamPolicyMapInput is an input type that accepts RepositoryIamPolicyMap and RepositoryIamPolicyMapOutput values.
// You can construct a concrete instance of `RepositoryIamPolicyMapInput` via:
//
//          RepositoryIamPolicyMap{ "key": RepositoryIamPolicyArgs{...} }
type RepositoryIamPolicyMapInput interface {
	pulumi.Input

	ToRepositoryIamPolicyMapOutput() RepositoryIamPolicyMapOutput
	ToRepositoryIamPolicyMapOutputWithContext(context.Context) RepositoryIamPolicyMapOutput
}

type RepositoryIamPolicyMap map[string]RepositoryIamPolicyInput

func (RepositoryIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RepositoryIamPolicy)(nil))
}

func (i RepositoryIamPolicyMap) ToRepositoryIamPolicyMapOutput() RepositoryIamPolicyMapOutput {
	return i.ToRepositoryIamPolicyMapOutputWithContext(context.Background())
}

func (i RepositoryIamPolicyMap) ToRepositoryIamPolicyMapOutputWithContext(ctx context.Context) RepositoryIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamPolicyMapOutput)
}

type RepositoryIamPolicyOutput struct {
	*pulumi.OutputState
}

func (RepositoryIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryIamPolicy)(nil))
}

func (o RepositoryIamPolicyOutput) ToRepositoryIamPolicyOutput() RepositoryIamPolicyOutput {
	return o
}

func (o RepositoryIamPolicyOutput) ToRepositoryIamPolicyOutputWithContext(ctx context.Context) RepositoryIamPolicyOutput {
	return o
}

func (o RepositoryIamPolicyOutput) ToRepositoryIamPolicyPtrOutput() RepositoryIamPolicyPtrOutput {
	return o.ToRepositoryIamPolicyPtrOutputWithContext(context.Background())
}

func (o RepositoryIamPolicyOutput) ToRepositoryIamPolicyPtrOutputWithContext(ctx context.Context) RepositoryIamPolicyPtrOutput {
	return o.ApplyT(func(v RepositoryIamPolicy) *RepositoryIamPolicy {
		return &v
	}).(RepositoryIamPolicyPtrOutput)
}

type RepositoryIamPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (RepositoryIamPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryIamPolicy)(nil))
}

func (o RepositoryIamPolicyPtrOutput) ToRepositoryIamPolicyPtrOutput() RepositoryIamPolicyPtrOutput {
	return o
}

func (o RepositoryIamPolicyPtrOutput) ToRepositoryIamPolicyPtrOutputWithContext(ctx context.Context) RepositoryIamPolicyPtrOutput {
	return o
}

type RepositoryIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (RepositoryIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RepositoryIamPolicy)(nil))
}

func (o RepositoryIamPolicyArrayOutput) ToRepositoryIamPolicyArrayOutput() RepositoryIamPolicyArrayOutput {
	return o
}

func (o RepositoryIamPolicyArrayOutput) ToRepositoryIamPolicyArrayOutputWithContext(ctx context.Context) RepositoryIamPolicyArrayOutput {
	return o
}

func (o RepositoryIamPolicyArrayOutput) Index(i pulumi.IntInput) RepositoryIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RepositoryIamPolicy {
		return vs[0].([]RepositoryIamPolicy)[vs[1].(int)]
	}).(RepositoryIamPolicyOutput)
}

type RepositoryIamPolicyMapOutput struct{ *pulumi.OutputState }

func (RepositoryIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RepositoryIamPolicy)(nil))
}

func (o RepositoryIamPolicyMapOutput) ToRepositoryIamPolicyMapOutput() RepositoryIamPolicyMapOutput {
	return o
}

func (o RepositoryIamPolicyMapOutput) ToRepositoryIamPolicyMapOutputWithContext(ctx context.Context) RepositoryIamPolicyMapOutput {
	return o
}

func (o RepositoryIamPolicyMapOutput) MapIndex(k pulumi.StringInput) RepositoryIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RepositoryIamPolicy {
		return vs[0].(map[string]RepositoryIamPolicy)[vs[1].(string)]
	}).(RepositoryIamPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(RepositoryIamPolicyOutput{})
	pulumi.RegisterOutputType(RepositoryIamPolicyPtrOutput{})
	pulumi.RegisterOutputType(RepositoryIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(RepositoryIamPolicyMapOutput{})
}
