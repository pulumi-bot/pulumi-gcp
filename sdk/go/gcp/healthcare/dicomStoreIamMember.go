// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Healthcare DICOM store. Each of these resources serves a different use case:
//
// * `healthcare.DicomStoreIamPolicy`: Authoritative. Sets the IAM policy for the DICOM store and replaces any existing policy already attached.
// * `healthcare.DicomStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the DICOM store are preserved.
// * `healthcare.DicomStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the DICOM store are preserved.
//
// > **Note:** `healthcare.DicomStoreIamPolicy` **cannot** be used in conjunction with `healthcare.DicomStoreIamBinding` and `healthcare.DicomStoreIamMember` or they will fight over what your policy should be.
//
// > **Note:** `healthcare.DicomStoreIamBinding` resources **can be** used in conjunction with `healthcare.DicomStoreIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_healthcare\_dicom\_store\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
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
// 		_, err = healthcare.NewDicomStoreIamPolicy(ctx, "dicomStore", &healthcare.DicomStoreIamPolicyArgs{
// 			DicomStoreId: pulumi.String("your-dicom-store-id"),
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
// ## google\_healthcare\_dicom\_store\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := healthcare.NewDicomStoreIamBinding(ctx, "dicomStore", &healthcare.DicomStoreIamBindingArgs{
// 			DicomStoreId: pulumi.String("your-dicom-store-id"),
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
// ## google\_healthcare\_dicom\_store\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := healthcare.NewDicomStoreIamMember(ctx, "dicomStore", &healthcare.DicomStoreIamMemberArgs{
// 			DicomStoreId: pulumi.String("your-dicom-store-id"),
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
//
// ## Import
//
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
//
// This member resource can be imported using the `dicom_store_id`, role, and account e.g.
//
// ```sh
//  $ pulumi import gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember dicom_store_iam "your-project-id/location-name/dataset-name/dicom-store-name roles/viewer user:foo@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiers; the resource in question and the role.
//
// This binding resource can be imported using the `dicom_store_id` and role, e.g.
//
// ```sh
//  $ pulumi import gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember dicom_store_iam "your-project-id/location-name/dataset-name/dicom-store-name roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `dicom_store_id`, role, and account e.g.
//
// ```sh
//  $ pulumi import gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember dicom_store_iam your-project-id/location-name/dataset-name/dicom-store-name
// ```
type DicomStoreIamMember struct {
	pulumi.CustomResourceState

	Condition DicomStoreIamMemberConditionPtrOutput `pulumi:"condition"`
	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId pulumi.StringOutput `pulumi:"dicomStoreId"`
	// (Computed) The etag of the DICOM store's IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The role that should be applied. Only one
	// `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewDicomStoreIamMember registers a new resource with the given unique name, arguments, and options.
func NewDicomStoreIamMember(ctx *pulumi.Context,
	name string, args *DicomStoreIamMemberArgs, opts ...pulumi.ResourceOption) (*DicomStoreIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DicomStoreId == nil {
		return nil, errors.New("invalid value for required argument 'DicomStoreId'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource DicomStoreIamMember
	err := ctx.RegisterResource("gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDicomStoreIamMember gets an existing DicomStoreIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDicomStoreIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DicomStoreIamMemberState, opts ...pulumi.ResourceOption) (*DicomStoreIamMember, error) {
	var resource DicomStoreIamMember
	err := ctx.ReadResource("gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DicomStoreIamMember resources.
type dicomStoreIamMemberState struct {
	Condition *DicomStoreIamMemberCondition `pulumi:"condition"`
	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId *string `pulumi:"dicomStoreId"`
	// (Computed) The etag of the DICOM store's IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The role that should be applied. Only one
	// `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type DicomStoreIamMemberState struct {
	Condition DicomStoreIamMemberConditionPtrInput
	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId pulumi.StringPtrInput
	// (Computed) The etag of the DICOM store's IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (DicomStoreIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*dicomStoreIamMemberState)(nil)).Elem()
}

type dicomStoreIamMemberArgs struct {
	Condition *DicomStoreIamMemberCondition `pulumi:"condition"`
	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId string `pulumi:"dicomStoreId"`
	Member       string `pulumi:"member"`
	// The role that should be applied. Only one
	// `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a DicomStoreIamMember resource.
type DicomStoreIamMemberArgs struct {
	Condition DicomStoreIamMemberConditionPtrInput
	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId pulumi.StringInput
	Member       pulumi.StringInput
	// The role that should be applied. Only one
	// `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (DicomStoreIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dicomStoreIamMemberArgs)(nil)).Elem()
}

type DicomStoreIamMemberInput interface {
	pulumi.Input

	ToDicomStoreIamMemberOutput() DicomStoreIamMemberOutput
	ToDicomStoreIamMemberOutputWithContext(ctx context.Context) DicomStoreIamMemberOutput
}

func (DicomStoreIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*DicomStoreIamMember)(nil)).Elem()
}

func (i DicomStoreIamMember) ToDicomStoreIamMemberOutput() DicomStoreIamMemberOutput {
	return i.ToDicomStoreIamMemberOutputWithContext(context.Background())
}

func (i DicomStoreIamMember) ToDicomStoreIamMemberOutputWithContext(ctx context.Context) DicomStoreIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DicomStoreIamMemberOutput)
}

type DicomStoreIamMemberOutput struct {
	*pulumi.OutputState
}

func (DicomStoreIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DicomStoreIamMemberOutput)(nil)).Elem()
}

func (o DicomStoreIamMemberOutput) ToDicomStoreIamMemberOutput() DicomStoreIamMemberOutput {
	return o
}

func (o DicomStoreIamMemberOutput) ToDicomStoreIamMemberOutputWithContext(ctx context.Context) DicomStoreIamMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DicomStoreIamMemberOutput{})
}
