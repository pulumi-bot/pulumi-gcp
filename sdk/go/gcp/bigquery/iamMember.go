// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for BigQuery Table. Each of these resources serves a different use case:
//
// * `bigquery.IamPolicy`: Authoritative. Sets the IAM policy for the table and replaces any existing policy already attached.
// * `bigquery.IamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
// * `bigquery.IamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the table are preserved.
//
// > **Note:** `bigquery.IamPolicy` **cannot** be used in conjunction with `bigquery.IamBinding` and `bigquery.IamMember` or they will fight over what your policy should be.
//
// > **Note:** `bigquery.IamBinding` resources **can be** used in conjunction with `bigquery.IamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_bigquery\_table\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/bigquery.dataOwner",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewIamPolicy(ctx, "policy", &bigquery.IamPolicyArgs{
// 			Project:    pulumi.Any(google_bigquery_table.Test.Project),
// 			DatasetId:  pulumi.Any(google_bigquery_table.Test.Dataset_id),
// 			TableId:    pulumi.Any(google_bigquery_table.Test.Table_id),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/bigquery.dataOwner",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 					Condition: organizations.GetIAMPolicyBindingCondition{
// 						Title:       "expires_after_2019_12_31",
// 						Description: "Expiring at midnight of 2019-12-31",
// 						Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewIamPolicy(ctx, "policy", &bigquery.IamPolicyArgs{
// 			Project:    pulumi.Any(google_bigquery_table.Test.Project),
// 			DatasetId:  pulumi.Any(google_bigquery_table.Test.Dataset_id),
// 			TableId:    pulumi.Any(google_bigquery_table.Test.Table_id),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_bigquery\_table\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := bigquery.NewIamBinding(ctx, "binding", &bigquery.IamBindingArgs{
// 			Project:   pulumi.Any(google_bigquery_table.Test.Project),
// 			DatasetId: pulumi.Any(google_bigquery_table.Test.Dataset_id),
// 			TableId:   pulumi.Any(google_bigquery_table.Test.Table_id),
// 			Role:      pulumi.String("roles/bigquery.dataOwner"),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := bigquery.NewIamBinding(ctx, "binding", &bigquery.IamBindingArgs{
// 			Project:   pulumi.Any(google_bigquery_table.Test.Project),
// 			DatasetId: pulumi.Any(google_bigquery_table.Test.Dataset_id),
// 			TableId:   pulumi.Any(google_bigquery_table.Test.Table_id),
// 			Role:      pulumi.String("roles/bigquery.dataOwner"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Condition: &bigquery.IamBindingConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_bigquery\_table\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := bigquery.NewIamMember(ctx, "member", &bigquery.IamMemberArgs{
// 			Project:   pulumi.Any(google_bigquery_table.Test.Project),
// 			DatasetId: pulumi.Any(google_bigquery_table.Test.Dataset_id),
// 			TableId:   pulumi.Any(google_bigquery_table.Test.Table_id),
// 			Role:      pulumi.String("roles/bigquery.dataOwner"),
// 			Member:    pulumi.String("user:jane@example.com"),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := bigquery.NewIamMember(ctx, "member", &bigquery.IamMemberArgs{
// 			Project:   pulumi.Any(google_bigquery_table.Test.Project),
// 			DatasetId: pulumi.Any(google_bigquery_table.Test.Dataset_id),
// 			TableId:   pulumi.Any(google_bigquery_table.Test.Table_id),
// 			Role:      pulumi.String("roles/bigquery.dataOwner"),
// 			Member:    pulumi.String("user:jane@example.com"),
// 			Condition: &bigquery.IamMemberConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}} * {{project}}/{{dataset_id}}/{{table_id}} * {{dataset_id}}/{{table_id}} * {{table_id}} Any variables not passed in the import command will be taken from the provider configuration. BigQuery table IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:bigquery/iamMember:IamMember editor "projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}} roles/bigquery.dataOwner user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:bigquery/iamMember:IamMember editor "projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}} roles/bigquery.dataOwner"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:bigquery/iamMember:IamMember editor projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type IamMember struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IamMemberConditionPtrOutput `pulumi:"condition"`
	DatasetId pulumi.StringOutput         `pulumi:"datasetId"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role    pulumi.StringOutput `pulumi:"role"`
	TableId pulumi.StringOutput `pulumi:"tableId"`
}

// NewIamMember registers a new resource with the given unique name, arguments, and options.
func NewIamMember(ctx *pulumi.Context,
	name string, args *IamMemberArgs, opts ...pulumi.ResourceOption) (*IamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.TableId == nil {
		return nil, errors.New("invalid value for required argument 'TableId'")
	}
	var resource IamMember
	err := ctx.RegisterResource("gcp:bigquery/iamMember:IamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamMember gets an existing IamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamMemberState, opts ...pulumi.ResourceOption) (*IamMember, error) {
	var resource IamMember
	err := ctx.ReadResource("gcp:bigquery/iamMember:IamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamMember resources.
type iamMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IamMemberCondition `pulumi:"condition"`
	DatasetId *string             `pulumi:"datasetId"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role    *string `pulumi:"role"`
	TableId *string `pulumi:"tableId"`
}

type IamMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IamMemberConditionPtrInput
	DatasetId pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigquery.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role    pulumi.StringPtrInput
	TableId pulumi.StringPtrInput
}

func (IamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamMemberState)(nil)).Elem()
}

type iamMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IamMemberCondition `pulumi:"condition"`
	DatasetId string              `pulumi:"datasetId"`
	Member    string              `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role    string `pulumi:"role"`
	TableId string `pulumi:"tableId"`
}

// The set of arguments for constructing a IamMember resource.
type IamMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IamMemberConditionPtrInput
	DatasetId pulumi.StringInput
	Member    pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigquery.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role    pulumi.StringInput
	TableId pulumi.StringInput
}

func (IamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamMemberArgs)(nil)).Elem()
}

type IamMemberInput interface {
	pulumi.Input

	ToIamMemberOutput() IamMemberOutput
	ToIamMemberOutputWithContext(ctx context.Context) IamMemberOutput
}

func (IamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*IamMember)(nil)).Elem()
}

func (i IamMember) ToIamMemberOutput() IamMemberOutput {
	return i.ToIamMemberOutputWithContext(context.Background())
}

func (i IamMember) ToIamMemberOutputWithContext(ctx context.Context) IamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamMemberOutput)
}

type IamMemberOutput struct {
	*pulumi.OutputState
}

func (IamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IamMemberOutput)(nil)).Elem()
}

func (o IamMemberOutput) ToIamMemberOutput() IamMemberOutput {
	return o
}

func (o IamMemberOutput) ToIamMemberOutputWithContext(ctx context.Context) IamMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IamMemberOutput{})
}
