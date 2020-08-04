// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secretmanager

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Secret Manager Secret. Each of these resources serves a different use case:
//
// * `secretmanager.SecretIamPolicy`: Authoritative. Sets the IAM policy for the secret and replaces any existing policy already attached.
// * `secretmanager.SecretIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the secret are preserved.
// * `secretmanager.SecretIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the secret are preserved.
//
// > **Note:** `secretmanager.SecretIamPolicy` **cannot** be used in conjunction with `secretmanager.SecretIamBinding` and `secretmanager.SecretIamMember` or they will fight over what your policy should be.
//
// > **Note:** `secretmanager.SecretIamBinding` resources **can be** used in conjunction with `secretmanager.SecretIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_secret\_manager\_secret\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/secretmanager"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/secretmanager.secretAccessor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = secretmanager.NewSecretIamPolicy(ctx, "policy", &secretmanager.SecretIamPolicyArgs{
// 			Project:    pulumi.Any(google_secret_manager_secret.Secret - basic.Project),
// 			SecretId:   pulumi.Any(google_secret_manager_secret.Secret - basic.Secret_id),
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
// ## google\_secret\_manager\_secret\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/secretmanager"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := secretmanager.NewSecretIamBinding(ctx, "binding", &secretmanager.SecretIamBindingArgs{
// 			Project:  pulumi.Any(google_secret_manager_secret.Secret - basic.Project),
// 			SecretId: pulumi.Any(google_secret_manager_secret.Secret - basic.Secret_id),
// 			Role:     pulumi.String("roles/secretmanager.secretAccessor"),
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
// ## google\_secret\_manager\_secret\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/secretmanager"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := secretmanager.NewSecretIamMember(ctx, "member", &secretmanager.SecretIamMemberArgs{
// 			Project:  pulumi.Any(google_secret_manager_secret.Secret - basic.Project),
// 			SecretId: pulumi.Any(google_secret_manager_secret.Secret - basic.Secret_id),
// 			Role:     pulumi.String("roles/secretmanager.secretAccessor"),
// 			Member:   pulumi.String("user:jane@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SecretIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project  pulumi.StringOutput `pulumi:"project"`
	SecretId pulumi.StringOutput `pulumi:"secretId"`
}

// NewSecretIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewSecretIamPolicy(ctx *pulumi.Context,
	name string, args *SecretIamPolicyArgs, opts ...pulumi.ResourceOption) (*SecretIamPolicy, error) {
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil || args.SecretId == nil {
		return nil, errors.New("missing required argument 'SecretId'")
	}
	if args == nil {
		args = &SecretIamPolicyArgs{}
	}
	var resource SecretIamPolicy
	err := ctx.RegisterResource("gcp:secretmanager/secretIamPolicy:SecretIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretIamPolicy gets an existing SecretIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretIamPolicyState, opts ...pulumi.ResourceOption) (*SecretIamPolicy, error) {
	var resource SecretIamPolicy
	err := ctx.ReadResource("gcp:secretmanager/secretIamPolicy:SecretIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretIamPolicy resources.
type secretIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project  *string `pulumi:"project"`
	SecretId *string `pulumi:"secretId"`
}

type SecretIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project  pulumi.StringPtrInput
	SecretId pulumi.StringPtrInput
}

func (SecretIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretIamPolicyState)(nil)).Elem()
}

type secretIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project  *string `pulumi:"project"`
	SecretId string  `pulumi:"secretId"`
}

// The set of arguments for constructing a SecretIamPolicy resource.
type SecretIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project  pulumi.StringPtrInput
	SecretId pulumi.StringInput
}

func (SecretIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretIamPolicyArgs)(nil)).Elem()
}
