// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Healthcare HL7v2 store. Each of these resources serves a different use case:
//
// * `healthcare.Hl7StoreIamPolicy`: Authoritative. Sets the IAM policy for the HL7v2 store and replaces any existing policy already attached.
// * `healthcare.Hl7StoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the HL7v2 store are preserved.
// * `healthcare.Hl7StoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the HL7v2 store are preserved.
//
// > **Note:** `healthcare.Hl7StoreIamPolicy` **cannot** be used in conjunction with `healthcare.Hl7StoreIamBinding` and `healthcare.Hl7StoreIamMember` or they will fight over what your policy should be.
//
// > **Note:** `healthcare.Hl7StoreIamBinding` resources **can be** used in conjunction with `healthcare.Hl7StoreIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_healthcare\_hl7\_v2\_store\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/organizations"
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
// 		_, err = healthcare.NewHl7StoreIamPolicy(ctx, "hl7V2Store", &healthcare.Hl7StoreIamPolicyArgs{
// 			Hl7V2StoreId: pulumi.String("your-hl7-v2-store-id"),
// 			PolicyData:   pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_healthcare\_hl7\_v2\_store\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := healthcare.NewHl7StoreIamBinding(ctx, "hl7V2Store", &healthcare.Hl7StoreIamBindingArgs{
// 			Hl7V2StoreId: pulumi.String("your-hl7-v2-store-id"),
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
// ## google\_healthcare\_hl7\_v2\_store\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := healthcare.NewHl7StoreIamMember(ctx, "hl7V2Store", &healthcare.Hl7StoreIamMemberArgs{
// 			Hl7V2StoreId: pulumi.String("your-hl7-v2-store-id"),
// 			Member:       pulumi.String("user:jane@example.com"),
// 			Role:         pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Hl7StoreIamBinding struct {
	pulumi.CustomResourceState

	Condition Hl7StoreIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the HL7v2 store's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The HL7v2 store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
	// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	Hl7V2StoreId pulumi.StringOutput      `pulumi:"hl7V2StoreId"`
	Members      pulumi.StringArrayOutput `pulumi:"members"`
	// The role that should be applied. Only one
	// `healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewHl7StoreIamBinding registers a new resource with the given unique name, arguments, and options.
func NewHl7StoreIamBinding(ctx *pulumi.Context,
	name string, args *Hl7StoreIamBindingArgs, opts ...pulumi.ResourceOption) (*Hl7StoreIamBinding, error) {
	if args == nil || args.Hl7V2StoreId == nil {
		return nil, errors.New("missing required argument 'Hl7V2StoreId'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &Hl7StoreIamBindingArgs{}
	}
	var resource Hl7StoreIamBinding
	err := ctx.RegisterResource("gcp:healthcare/hl7StoreIamBinding:Hl7StoreIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHl7StoreIamBinding gets an existing Hl7StoreIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHl7StoreIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Hl7StoreIamBindingState, opts ...pulumi.ResourceOption) (*Hl7StoreIamBinding, error) {
	var resource Hl7StoreIamBinding
	err := ctx.ReadResource("gcp:healthcare/hl7StoreIamBinding:Hl7StoreIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Hl7StoreIamBinding resources.
type hl7StoreIamBindingState struct {
	Condition *Hl7StoreIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the HL7v2 store's IAM policy.
	Etag *string `pulumi:"etag"`
	// The HL7v2 store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
	// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	Hl7V2StoreId *string  `pulumi:"hl7V2StoreId"`
	Members      []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type Hl7StoreIamBindingState struct {
	Condition Hl7StoreIamBindingConditionPtrInput
	// (Computed) The etag of the HL7v2 store's IAM policy.
	Etag pulumi.StringPtrInput
	// The HL7v2 store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
	// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	Hl7V2StoreId pulumi.StringPtrInput
	Members      pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (Hl7StoreIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*hl7StoreIamBindingState)(nil)).Elem()
}

type hl7StoreIamBindingArgs struct {
	Condition *Hl7StoreIamBindingCondition `pulumi:"condition"`
	// The HL7v2 store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
	// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	Hl7V2StoreId string   `pulumi:"hl7V2StoreId"`
	Members      []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a Hl7StoreIamBinding resource.
type Hl7StoreIamBindingArgs struct {
	Condition Hl7StoreIamBindingConditionPtrInput
	// The HL7v2 store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
	// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	Hl7V2StoreId pulumi.StringInput
	Members      pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (Hl7StoreIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hl7StoreIamBindingArgs)(nil)).Elem()
}
