// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/healthcare"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/healthcare"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/healthcare"
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
type DicomStoreIamBinding struct {
	pulumi.CustomResourceState

	Condition DicomStoreIamBindingConditionPtrOutput `pulumi:"condition"`
	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId pulumi.StringOutput `pulumi:"dicomStoreId"`
	// (Computed) The etag of the DICOM store's IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The role that should be applied. Only one
	// `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewDicomStoreIamBinding registers a new resource with the given unique name, arguments, and options.
func NewDicomStoreIamBinding(ctx *pulumi.Context,
	name string, args *DicomStoreIamBindingArgs, opts ...pulumi.ResourceOption) (*DicomStoreIamBinding, error) {
	if args == nil || args.DicomStoreId == nil {
		return nil, errors.New("missing required argument 'DicomStoreId'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &DicomStoreIamBindingArgs{}
	}
	var resource DicomStoreIamBinding
	err := ctx.RegisterResource("gcp:healthcare/dicomStoreIamBinding:DicomStoreIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDicomStoreIamBinding gets an existing DicomStoreIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDicomStoreIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DicomStoreIamBindingState, opts ...pulumi.ResourceOption) (*DicomStoreIamBinding, error) {
	var resource DicomStoreIamBinding
	err := ctx.ReadResource("gcp:healthcare/dicomStoreIamBinding:DicomStoreIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DicomStoreIamBinding resources.
type dicomStoreIamBindingState struct {
	Condition *DicomStoreIamBindingCondition `pulumi:"condition"`
	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId *string `pulumi:"dicomStoreId"`
	// (Computed) The etag of the DICOM store's IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type DicomStoreIamBindingState struct {
	Condition DicomStoreIamBindingConditionPtrInput
	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId pulumi.StringPtrInput
	// (Computed) The etag of the DICOM store's IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (DicomStoreIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*dicomStoreIamBindingState)(nil)).Elem()
}

type dicomStoreIamBindingArgs struct {
	Condition *DicomStoreIamBindingCondition `pulumi:"condition"`
	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId string   `pulumi:"dicomStoreId"`
	Members      []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a DicomStoreIamBinding resource.
type DicomStoreIamBindingArgs struct {
	Condition DicomStoreIamBindingConditionPtrInput
	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId pulumi.StringInput
	Members      pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (DicomStoreIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dicomStoreIamBindingArgs)(nil)).Elem()
}
