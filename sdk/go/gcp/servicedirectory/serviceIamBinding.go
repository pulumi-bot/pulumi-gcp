// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicedirectory

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Service Directory Service. Each of these resources serves a different use case:
//
// * `servicedirectory.ServiceIamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
// * `servicedirectory.ServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
// * `servicedirectory.ServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.
//
// > **Note:** `servicedirectory.ServiceIamPolicy` **cannot** be used in conjunction with `servicedirectory.ServiceIamBinding` and `servicedirectory.ServiceIamMember` or they will fight over what your policy should be.
//
// > **Note:** `servicedirectory.ServiceIamBinding` resources **can be** used in conjunction with `servicedirectory.ServiceIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_service\_directory\_service\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/servicedirectory"
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
// 		policy, err := servicedirectory.NewServiceIamPolicy(ctx, "policy", &servicedirectory.ServiceIamPolicyArgs{
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
// ## google\_service\_directory\_service\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/servicedirectory"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		binding, err := servicedirectory.NewServiceIamBinding(ctx, "binding", &servicedirectory.ServiceIamBindingArgs{
// 			Role: pulumi.String("roles/viewer"),
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
// ## google\_service\_directory\_service\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/servicedirectory"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		member, err := servicedirectory.NewServiceIamMember(ctx, "member", &servicedirectory.ServiceIamMemberArgs{
// 			Role:   pulumi.String("roles/viewer"),
// 			Member: pulumi.String("user:jane@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ServiceIamBinding struct {
	pulumi.CustomResourceState

	Condition ServiceIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
	// The role that should be applied. Only one
	// `servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewServiceIamBinding registers a new resource with the given unique name, arguments, and options.
func NewServiceIamBinding(ctx *pulumi.Context,
	name string, args *ServiceIamBindingArgs, opts ...pulumi.ResourceOption) (*ServiceIamBinding, error) {
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &ServiceIamBindingArgs{}
	}
	var resource ServiceIamBinding
	err := ctx.RegisterResource("gcp:servicedirectory/serviceIamBinding:ServiceIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceIamBinding gets an existing ServiceIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceIamBindingState, opts ...pulumi.ResourceOption) (*ServiceIamBinding, error) {
	var resource ServiceIamBinding
	err := ctx.ReadResource("gcp:servicedirectory/serviceIamBinding:ServiceIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceIamBinding resources.
type serviceIamBindingState struct {
	Condition *ServiceIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The role that should be applied. Only one
	// `servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ServiceIamBindingState struct {
	Condition ServiceIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ServiceIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceIamBindingState)(nil)).Elem()
}

type serviceIamBindingArgs struct {
	Condition *ServiceIamBindingCondition `pulumi:"condition"`
	Members   []string                    `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The role that should be applied. Only one
	// `servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ServiceIamBinding resource.
type ServiceIamBindingArgs struct {
	Condition ServiceIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (ServiceIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceIamBindingArgs)(nil)).Elem()
}
