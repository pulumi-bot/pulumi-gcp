// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Healthcare dataset. Each of these resources serves a different use case:
//
// * `healthcare.DatasetIamPolicy`: Authoritative. Sets the IAM policy for the dataset and replaces any existing policy already attached.
// * `healthcare.DatasetIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the dataset are preserved.
// * `healthcare.DatasetIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the dataset are preserved.
//
// > **Note:** `healthcare.DatasetIamPolicy` **cannot** be used in conjunction with `healthcare.DatasetIamBinding` and `healthcare.DatasetIamMember` or they will fight over what your policy should be.
//
// > **Note:** `healthcare.DatasetIamBinding` resources **can be** used in conjunction with `healthcare.DatasetIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_healthcare\_dataset\_iam\_policy
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
// 		_, err = healthcare.NewDatasetIamPolicy(ctx, "dataset", &healthcare.DatasetIamPolicyArgs{
// 			DatasetId:  pulumi.String("your-dataset-id"),
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
// ## google\_healthcare\_dataset\_iam\_binding
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
// 		_, err = healthcare.NewDatasetIamBinding(ctx, "dataset", &healthcare.DatasetIamBindingArgs{
// 			DatasetId: pulumi.String("your-dataset-id"),
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
// ## google\_healthcare\_dataset\_iam\_member
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
// 		_, err = healthcare.NewDatasetIamMember(ctx, "dataset", &healthcare.DatasetIamMemberArgs{
// 			DatasetId: pulumi.String("your-dataset-id"),
// 			Member:    pulumi.String("user:jane@example.com"),
// 			Role:      pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DatasetIamBinding struct {
	pulumi.CustomResourceState

	Condition DatasetIamBindingConditionPtrOutput `pulumi:"condition"`
	// The dataset ID, in the form
	// `{project_id}/{location_name}/{dataset_name}` or
	// `{location_name}/{dataset_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DatasetId pulumi.StringOutput `pulumi:"datasetId"`
	// (Computed) The etag of the dataset's IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The role that should be applied. Only one
	// `healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewDatasetIamBinding registers a new resource with the given unique name, arguments, and options.
func NewDatasetIamBinding(ctx *pulumi.Context,
	name string, args *DatasetIamBindingArgs, opts ...pulumi.ResourceOption) (*DatasetIamBinding, error) {
	if args == nil || args.DatasetId == nil {
		return nil, errors.New("missing required argument 'DatasetId'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &DatasetIamBindingArgs{}
	}
	var resource DatasetIamBinding
	err := ctx.RegisterResource("gcp:healthcare/datasetIamBinding:DatasetIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetIamBinding gets an existing DatasetIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetIamBindingState, opts ...pulumi.ResourceOption) (*DatasetIamBinding, error) {
	var resource DatasetIamBinding
	err := ctx.ReadResource("gcp:healthcare/datasetIamBinding:DatasetIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetIamBinding resources.
type datasetIamBindingState struct {
	Condition *DatasetIamBindingCondition `pulumi:"condition"`
	// The dataset ID, in the form
	// `{project_id}/{location_name}/{dataset_name}` or
	// `{location_name}/{dataset_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DatasetId *string `pulumi:"datasetId"`
	// (Computed) The etag of the dataset's IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type DatasetIamBindingState struct {
	Condition DatasetIamBindingConditionPtrInput
	// The dataset ID, in the form
	// `{project_id}/{location_name}/{dataset_name}` or
	// `{location_name}/{dataset_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DatasetId pulumi.StringPtrInput
	// (Computed) The etag of the dataset's IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (DatasetIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetIamBindingState)(nil)).Elem()
}

type datasetIamBindingArgs struct {
	Condition *DatasetIamBindingCondition `pulumi:"condition"`
	// The dataset ID, in the form
	// `{project_id}/{location_name}/{dataset_name}` or
	// `{location_name}/{dataset_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DatasetId string   `pulumi:"datasetId"`
	Members   []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a DatasetIamBinding resource.
type DatasetIamBindingArgs struct {
	Condition DatasetIamBindingConditionPtrInput
	// The dataset ID, in the form
	// `{project_id}/{location_name}/{dataset_name}` or
	// `{location_name}/{dataset_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DatasetId pulumi.StringInput
	Members   pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (DatasetIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetIamBindingArgs)(nil)).Elem()
}
