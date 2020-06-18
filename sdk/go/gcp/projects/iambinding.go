// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Four different resources help you manage your IAM policy for a project. Each of these resources serves a different use case:
//
// * `projects.IAMPolicy`: Authoritative. Sets the IAM policy for the project and replaces any existing policy already attached.
// * `projects.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the project are preserved.
// * `projects.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the project are preserved.
// * `projects.IAMAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.
//
// > **Note:** `projects.IAMPolicy` **cannot** be used in conjunction with `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig` or they will fight over what your policy should be.
//
// > **Note:** `projects.IAMBinding` resources **can be** used in conjunction with `projects.IAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_project\_iam\_policy
//
// > **Be careful!** You can accidentally lock yourself out of your project
//    using this resource. Deleting a `projects.IAMPolicy` removes access
//    from anyone without organization-level access to the project. Proceed with caution.
//    It's not recommended to use `projects.IAMPolicy` with your provider project
//    to avoid locking yourself out, and it should generally only be used with projects
//    fully managed by this provider. If you do use this resource, it is recommended to **import** the policy before
//    applying the change.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Binding: []map[string]interface{}{
// 				map[string]interface{}{
// 					"role": "roles/editor",
// 					"members": []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = projects.NewIAMPolicy(ctx, "project", &projects.IAMPolicyArgs{
// 			Project:    pulumi.String("your-project-id"),
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
// With IAM Conditions):
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: organizations.getIAMPolicyBindingArray{
// 				&organizations.LookupIAMPolicyBinding{
// 					Condition: &organizations.LookupIAMPolicyBindingCondition{
// 						Description: "Expiring at midnight of 2019-12-31",
// 						Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 						Title:       "expires_after_2019_12_31",
// 					},
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 					Role: "roles/editor",
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = projects.NewIAMPolicy(ctx, "project", &projects.IAMPolicyArgs{
// 			PolicyData: pulumi.String(admin.PolicyData),
// 			Project:    pulumi.String("your-project-id"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_project\_iam\_binding
//
// > **Note:** If `role` is set to `roles/owner` and you don't specify a user or service account you have access to in `members`, you can lock yourself out of your project.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = projects.NewIAMBinding(ctx, "project", &projects.IAMBindingArgs{
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Project: pulumi.String("your-project-id"),
// 			Role:    pulumi.String("roles/editor"),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = projects.NewIAMBinding(ctx, "project", &projects.IAMBindingArgs{
// 			Condition: &projects.IAMBindingConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Project: pulumi.String("your-project-id"),
// 			Role:    pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_project\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = projects.NewIAMMember(ctx, "project", &projects.IAMMemberArgs{
// 			Member:  pulumi.String("user:jane@example.com"),
// 			Project: pulumi.String("your-project-id"),
// 			Role:    pulumi.String("roles/editor"),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = projects.NewIAMMember(ctx, "project", &projects.IAMMemberArgs{
// 			Condition: &projects.IAMMemberConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			Member:  pulumi.String("user:jane@example.com"),
// 			Project: pulumi.String("your-project-id"),
// 			Role:    pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_project\_iam\_audit\_config
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = projects.NewIAMAuditConfig(ctx, "project", &projects.IAMAuditConfigArgs{
// 			AuditLogConfigs: projects.IAMAuditConfigAuditLogConfigArray{
// 				&projects.IAMAuditConfigAuditLogConfigArgs{
// 					LogType: pulumi.String("ADMIN_READ"),
// 				},
// 				&projects.IAMAuditConfigAuditLogConfigArgs{
// 					ExemptedMembers: pulumi.StringArray{
// 						pulumi.String("user:joebloggs@hashicorp.com"),
// 					},
// 					LogType: pulumi.String("DATA_READ"),
// 				},
// 			},
// 			Project: pulumi.String("your-project-id"),
// 			Service: pulumi.String("allServices"),
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

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the project's IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
	// Required for `projects.IAMPolicy` - you must explicitly set the project, and it
	// will not be inferred from the provider.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
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
	if args == nil {
		args = &IAMBindingArgs{}
	}
	var resource IAMBinding
	err := ctx.RegisterResource("gcp:projects/iAMBinding:IAMBinding", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:projects/iAMBinding:IAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMBinding resources.
type iambindingState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the project's IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
	// Required for `projects.IAMPolicy` - you must explicitly set the project, and it
	// will not be inferred from the provider.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type IAMBindingState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMBindingConditionPtrInput
	// (Computed) The etag of the project's IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
	// Required for `projects.IAMPolicy` - you must explicitly set the project, and it
	// will not be inferred from the provider.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (IAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*iambindingState)(nil)).Elem()
}

type iambindingArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IAMBindingCondition `pulumi:"condition"`
	Members   []string             `pulumi:"members"`
	// The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
	// Required for `projects.IAMPolicy` - you must explicitly set the project, and it
	// will not be inferred from the provider.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a IAMBinding resource.
type IAMBindingArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
	// Required for `projects.IAMPolicy` - you must explicitly set the project, and it
	// will not be inferred from the provider.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (IAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iambindingArgs)(nil)).Elem()
}
