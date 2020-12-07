// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Healthcare FHIR store. Each of these resources serves a different use case:
//
// * `healthcare.FhirStoreIamPolicy`: Authoritative. Sets the IAM policy for the FHIR store and replaces any existing policy already attached.
// * `healthcare.FhirStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the FHIR store are preserved.
// * `healthcare.FhirStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the FHIR store are preserved.
//
// > **Note:** `healthcare.FhirStoreIamPolicy` **cannot** be used in conjunction with `healthcare.FhirStoreIamBinding` and `healthcare.FhirStoreIamMember` or they will fight over what your policy should be.
//
// > **Note:** `healthcare.FhirStoreIamBinding` resources **can be** used in conjunction with `healthcare.FhirStoreIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_healthcare\_fhir\_store\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 		_, err = healthcare.NewFhirStoreIamPolicy(ctx, "fhirStore", &healthcare.FhirStoreIamPolicyArgs{
// 			FhirStoreId: pulumi.String("your-fhir-store-id"),
// 			PolicyData:  pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_healthcare\_fhir\_store\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := healthcare.NewFhirStoreIamBinding(ctx, "fhirStore", &healthcare.FhirStoreIamBindingArgs{
// 			FhirStoreId: pulumi.String("your-fhir-store-id"),
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
// ## google\_healthcare\_fhir\_store\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := healthcare.NewFhirStoreIamMember(ctx, "fhirStore", &healthcare.FhirStoreIamMemberArgs{
// 			FhirStoreId: pulumi.String("your-fhir-store-id"),
// 			Member:      pulumi.String("user:jane@example.com"),
// 			Role:        pulumi.String("roles/editor"),
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
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
//
// This member resource can be imported using the `fhir_store_id`, role, and account e.g.
//
// ```sh
//  $ pulumi import gcp:healthcare/fhirStoreIamBinding:FhirStoreIamBinding fhir_store_iam "your-project-id/location-name/dataset-name/fhir-store-name roles/viewer user:foo@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiers; the resource in question and the role.
//
// This binding resource can be imported using the `fhir_store_id` and role, e.g.
//
// ```sh
//  $ pulumi import gcp:healthcare/fhirStoreIamBinding:FhirStoreIamBinding fhir_store_iam "your-project-id/location-name/dataset-name/fhir-store-name roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `fhir_store_id`, role, and account e.g.
//
// ```sh
//  $ pulumi import gcp:healthcare/fhirStoreIamBinding:FhirStoreIamBinding fhir_store_iam your-project-id/location-name/dataset-name/fhir-store-name
// ```
type FhirStoreIamBinding struct {
	pulumi.CustomResourceState

	Condition FhirStoreIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the FHIR store's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId pulumi.StringOutput      `pulumi:"fhirStoreId"`
	Members     pulumi.StringArrayOutput `pulumi:"members"`
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewFhirStoreIamBinding registers a new resource with the given unique name, arguments, and options.
func NewFhirStoreIamBinding(ctx *pulumi.Context,
	name string, args *FhirStoreIamBindingArgs, opts ...pulumi.ResourceOption) (*FhirStoreIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FhirStoreId == nil {
		return nil, errors.New("invalid value for required argument 'FhirStoreId'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource FhirStoreIamBinding
	err := ctx.RegisterResource("gcp:healthcare/fhirStoreIamBinding:FhirStoreIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFhirStoreIamBinding gets an existing FhirStoreIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFhirStoreIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FhirStoreIamBindingState, opts ...pulumi.ResourceOption) (*FhirStoreIamBinding, error) {
	var resource FhirStoreIamBinding
	err := ctx.ReadResource("gcp:healthcare/fhirStoreIamBinding:FhirStoreIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FhirStoreIamBinding resources.
type fhirStoreIamBindingState struct {
	Condition *FhirStoreIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the FHIR store's IAM policy.
	Etag *string `pulumi:"etag"`
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId *string  `pulumi:"fhirStoreId"`
	Members     []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type FhirStoreIamBindingState struct {
	Condition FhirStoreIamBindingConditionPtrInput
	// (Computed) The etag of the FHIR store's IAM policy.
	Etag pulumi.StringPtrInput
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId pulumi.StringPtrInput
	Members     pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (FhirStoreIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirStoreIamBindingState)(nil)).Elem()
}

type fhirStoreIamBindingArgs struct {
	Condition *FhirStoreIamBindingCondition `pulumi:"condition"`
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId string   `pulumi:"fhirStoreId"`
	Members     []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a FhirStoreIamBinding resource.
type FhirStoreIamBindingArgs struct {
	Condition FhirStoreIamBindingConditionPtrInput
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId pulumi.StringInput
	Members     pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (FhirStoreIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirStoreIamBindingArgs)(nil)).Elem()
}

type FhirStoreIamBindingInput interface {
	pulumi.Input

	ToFhirStoreIamBindingOutput() FhirStoreIamBindingOutput
	ToFhirStoreIamBindingOutputWithContext(ctx context.Context) FhirStoreIamBindingOutput
}

func (FhirStoreIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirStoreIamBinding)(nil)).Elem()
}

func (i FhirStoreIamBinding) ToFhirStoreIamBindingOutput() FhirStoreIamBindingOutput {
	return i.ToFhirStoreIamBindingOutputWithContext(context.Background())
}

func (i FhirStoreIamBinding) ToFhirStoreIamBindingOutputWithContext(ctx context.Context) FhirStoreIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirStoreIamBindingOutput)
}

type FhirStoreIamBindingOutput struct {
	*pulumi.OutputState
}

func (FhirStoreIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirStoreIamBindingOutput)(nil)).Elem()
}

func (o FhirStoreIamBindingOutput) ToFhirStoreIamBindingOutput() FhirStoreIamBindingOutput {
	return o
}

func (o FhirStoreIamBindingOutput) ToFhirStoreIamBindingOutputWithContext(ctx context.Context) FhirStoreIamBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FhirStoreIamBindingOutput{})
}
