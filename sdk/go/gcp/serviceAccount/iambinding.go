// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package serviceAccount

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// When managing IAM roles, you can treat a service account either as a resource or as an identity. This resource is to add iam policy bindings to a service account resource **to configure permissions for who can edit the service account**. To configure permissions for a service account to act as an identity that can manage other GCP resources, use the googleProjectIam set of resources.
//
// Three different resources help you manage your IAM policy for a service account. Each of these resources serves a different use case:
//
// * `serviceAccount.IAMPolicy`: Authoritative. Sets the IAM policy for the service account and replaces any existing policy already attached.
// * `serviceAccount.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service account are preserved.
// * `serviceAccount.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service account are preserved.
//
// > **Note:** `serviceAccount.IAMPolicy` **cannot** be used in conjunction with `serviceAccount.IAMBinding` and `serviceAccount.IAMMember` or they will fight over what your policy should be.
//
// > **Note:** `serviceAccount.IAMBinding` resources **can be** used in conjunction with `serviceAccount.IAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_service\_account\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/serviceAccount"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Binding: []map[string]interface{}{
// 				map[string]interface{}{
// 					"role": "roles/iam.serviceAccountUser",
// 					"members": []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		sa, err := serviceAccount.NewAccount(ctx, "sa", &serviceAccount.AccountArgs{
// 			AccountId:   pulumi.String("my-service-account"),
// 			DisplayName: pulumi.String("A service account that only Jane can interact with"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = serviceAccount.NewIAMPolicy(ctx, "admin-account-iam", &serviceAccount.IAMPolicyArgs{
// 			ServiceAccountId: sa.Name,
// 			PolicyData:       pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_service\_account\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/serviceAccount"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		sa, err := serviceAccount.NewAccount(ctx, "sa", &serviceAccount.AccountArgs{
// 			AccountId:   pulumi.String("my-service-account"),
// 			DisplayName: pulumi.String("A service account that only Jane can use"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = serviceAccount.NewIAMBinding(ctx, "admin-account-iam", &serviceAccount.IAMBindingArgs{
// 			ServiceAccountId: sa.Name,
// 			Role:             pulumi.String("roles/iam.serviceAccountUser"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/serviceAccount"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		sa, err := serviceAccount.NewAccount(ctx, "sa", &serviceAccount.AccountArgs{
// 			AccountId:   pulumi.String("my-service-account"),
// 			DisplayName: pulumi.String("A service account that only Jane can use"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = serviceAccount.NewIAMBinding(ctx, "admin-account-iam", &serviceAccount.IAMBindingArgs{
// 			Condition: &serviceAccount.IAMBindingConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role:             pulumi.String("roles/iam.serviceAccountUser"),
// 			ServiceAccountId: sa.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_service\_account\_iam\_member
//
//
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/serviceAccount"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		sa, err := serviceAccount.NewAccount(ctx, "sa", &serviceAccount.AccountArgs{
// 			AccountId:   pulumi.String("my-service-account"),
// 			DisplayName: pulumi.String("A service account that Jane can use"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = serviceAccount.NewIAMMember(ctx, "admin-account-iam", &serviceAccount.IAMMemberArgs{
// 			Condition: &serviceAccount.IAMMemberConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			Member:           pulumi.String("user:jane@example.com"),
// 			Role:             pulumi.String("roles/iam.serviceAccountUser"),
// 			ServiceAccountId: sa.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type IAMBinding struct {
	pulumi.CustomResourceState

	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the service account IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The role that should be applied. Only one
	// `serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId pulumi.StringOutput `pulumi:"serviceAccountId"`
}

// NewIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewIAMBinding(ctx *pulumi.Context,
	name string, args *IAMBindingArgs, opts ...pulumi.ResourceOption) (*IAMBinding, error) {
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.ServiceAccountId == nil {
		return nil, errors.New("missing required argument 'ServiceAccountId'")
	}
	if args == nil {
		args = &IAMBindingArgs{}
	}
	var resource IAMBinding
	err := ctx.RegisterResource("gcp:serviceAccount/iAMBinding:IAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMBinding gets an existing IAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMBindingState, opts ...pulumi.ResourceOption) (*IAMBinding, error) {
	var resource IAMBinding
	err := ctx.ReadResource("gcp:serviceAccount/iAMBinding:IAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMBinding resources.
type iambindingState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the service account IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId *string `pulumi:"serviceAccountId"`
}

type IAMBindingState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMBindingConditionPtrInput
	// (Computed) The etag of the service account IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId pulumi.StringPtrInput
}

func (IAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*iambindingState)(nil)).Elem()
}

type iambindingArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IAMBindingCondition `pulumi:"condition"`
	Members   []string             `pulumi:"members"`
	// The role that should be applied. Only one
	// `serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId string `pulumi:"serviceAccountId"`
}

// The set of arguments for constructing a IAMBinding resource.
type IAMBindingArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId pulumi.StringInput
}

func (IAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iambindingArgs)(nil)).Elem()
}
