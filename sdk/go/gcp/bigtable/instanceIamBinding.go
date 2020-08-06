// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigtable

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage IAM policies on bigtable instances. Each of these resources serves a different use case:
//
// * `bigtable.InstanceIamPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
// * `bigtable.InstanceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
// * `bigtable.InstanceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
//
// > **Note:** `bigtable.InstanceIamPolicy` **cannot** be used in conjunction with `bigtable.InstanceIamBinding` and `bigtable.InstanceIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the instance as `bigtable.InstanceIamPolicy` replaces the entire policy.
//
// > **Note:** `bigtable.InstanceIamBinding` resources **can be** used in conjunction with `bigtable.InstanceIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_bigtable\_instance\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/bigtable"
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/editor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigtable.NewInstanceIamPolicy(ctx, "editor", &bigtable.InstanceIamPolicyArgs{
// 			Project:    pulumi.String("your-project"),
// 			Instance:   pulumi.String("your-bigtable-instance"),
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
// ## google\_bigtable\_instance\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/bigtable"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := bigtable.NewInstanceIamBinding(ctx, "editor", &bigtable.InstanceIamBindingArgs{
// 			Instance: pulumi.String("your-bigtable-instance"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role: pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_bigtable\_instance\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/bigtable"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := bigtable.NewInstanceIamMember(ctx, "editor", &bigtable.InstanceIamMemberArgs{
// 			Instance: pulumi.String("your-bigtable-instance"),
// 			Member:   pulumi.String("user:jane@example.com"),
// 			Role:     pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type InstanceIamBinding struct {
	pulumi.CustomResourceState

	Condition InstanceIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the instances's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance pulumi.StringOutput      `pulumi:"instance"`
	Members  pulumi.StringArrayOutput `pulumi:"members"`
	// The project in which the instance belongs. If it
	// is not provided, a default will be supplied.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewInstanceIamBinding registers a new resource with the given unique name, arguments, and options.
func NewInstanceIamBinding(ctx *pulumi.Context,
	name string, args *InstanceIamBindingArgs, opts ...pulumi.ResourceOption) (*InstanceIamBinding, error) {
	if args == nil || args.Instance == nil {
		return nil, errors.New("missing required argument 'Instance'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &InstanceIamBindingArgs{}
	}
	var resource InstanceIamBinding
	err := ctx.RegisterResource("gcp:bigtable/instanceIamBinding:InstanceIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceIamBinding gets an existing InstanceIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceIamBindingState, opts ...pulumi.ResourceOption) (*InstanceIamBinding, error) {
	var resource InstanceIamBinding
	err := ctx.ReadResource("gcp:bigtable/instanceIamBinding:InstanceIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceIamBinding resources.
type instanceIamBindingState struct {
	Condition *InstanceIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the instances's IAM policy.
	Etag *string `pulumi:"etag"`
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance *string  `pulumi:"instance"`
	Members  []string `pulumi:"members"`
	// The project in which the instance belongs. If it
	// is not provided, a default will be supplied.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type InstanceIamBindingState struct {
	Condition InstanceIamBindingConditionPtrInput
	// (Computed) The etag of the instances's IAM policy.
	Etag pulumi.StringPtrInput
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The project in which the instance belongs. If it
	// is not provided, a default will be supplied.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (InstanceIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIamBindingState)(nil)).Elem()
}

type instanceIamBindingArgs struct {
	Condition *InstanceIamBindingCondition `pulumi:"condition"`
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance string   `pulumi:"instance"`
	Members  []string `pulumi:"members"`
	// The project in which the instance belongs. If it
	// is not provided, a default will be supplied.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a InstanceIamBinding resource.
type InstanceIamBindingArgs struct {
	Condition InstanceIamBindingConditionPtrInput
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance pulumi.StringInput
	Members  pulumi.StringArrayInput
	// The project in which the instance belongs. If it
	// is not provided, a default will be supplied.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (InstanceIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIamBindingArgs)(nil)).Elem()
}
