// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigtable

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage IAM policies on bigtable tables. Each of these resources serves a different use case:
//
// * `bigtable.TableIamPolicy`: Authoritative. Sets the IAM policy for the tables and replaces any existing policy already attached.
// * `bigtable.TableIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
// * `bigtable.TableIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the table are preserved.
//
// > **Note:** `bigtable.TableIamPolicy` **cannot** be used in conjunction with `bigtable.TableIamBinding` and `bigtable.TableIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the table as `bigtable.TableIamPolicy` replaces the entire policy.
//
// > **Note:** `bigtable.TableIamBinding` resources **can be** used in conjunction with `bigtable.TableIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## Import
//
// Instance IAM resources can be imported using the project, table name, role and/or member.
//
// ```sh
//  $ pulumi import gcp:bigtable/tableIamBinding:TableIamBinding editor "projects/{project}/tables/{table}"
// ```
//
// ```sh
//  $ pulumi import gcp:bigtable/tableIamBinding:TableIamBinding editor "projects/{project}/tables/{table} roles/editor"
// ```
//
// ```sh
//  $ pulumi import gcp:bigtable/tableIamBinding:TableIamBinding editor "projects/{project}/tables/{table} roles/editor user:jane@example.com"
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TableIamBinding struct {
	pulumi.CustomResourceState

	Condition TableIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the tables's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name or relative resource id of the instance that owns the table.
	Instance pulumi.StringOutput      `pulumi:"instance"`
	Members  pulumi.StringArrayOutput `pulumi:"members"`
	Project  pulumi.StringOutput      `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role pulumi.StringOutput `pulumi:"role"`
	// The name or relative resource id of the table to manage IAM policies for.
	Table pulumi.StringOutput `pulumi:"table"`
}

// NewTableIamBinding registers a new resource with the given unique name, arguments, and options.
func NewTableIamBinding(ctx *pulumi.Context,
	name string, args *TableIamBindingArgs, opts ...pulumi.ResourceOption) (*TableIamBinding, error) {
	if args == nil || args.Instance == nil {
		return nil, errors.New("missing required argument 'Instance'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.Table == nil {
		return nil, errors.New("missing required argument 'Table'")
	}
	if args == nil {
		args = &TableIamBindingArgs{}
	}
	var resource TableIamBinding
	err := ctx.RegisterResource("gcp:bigtable/tableIamBinding:TableIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTableIamBinding gets an existing TableIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTableIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableIamBindingState, opts ...pulumi.ResourceOption) (*TableIamBinding, error) {
	var resource TableIamBinding
	err := ctx.ReadResource("gcp:bigtable/tableIamBinding:TableIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TableIamBinding resources.
type tableIamBindingState struct {
	Condition *TableIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the tables's IAM policy.
	Etag *string `pulumi:"etag"`
	// The name or relative resource id of the instance that owns the table.
	Instance *string  `pulumi:"instance"`
	Members  []string `pulumi:"members"`
	Project  *string  `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role *string `pulumi:"role"`
	// The name or relative resource id of the table to manage IAM policies for.
	Table *string `pulumi:"table"`
}

type TableIamBindingState struct {
	Condition TableIamBindingConditionPtrInput
	// (Computed) The etag of the tables's IAM policy.
	Etag pulumi.StringPtrInput
	// The name or relative resource id of the instance that owns the table.
	Instance pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	Project  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role pulumi.StringPtrInput
	// The name or relative resource id of the table to manage IAM policies for.
	Table pulumi.StringPtrInput
}

func (TableIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableIamBindingState)(nil)).Elem()
}

type tableIamBindingArgs struct {
	Condition *TableIamBindingCondition `pulumi:"condition"`
	// The name or relative resource id of the instance that owns the table.
	Instance string   `pulumi:"instance"`
	Members  []string `pulumi:"members"`
	Project  *string  `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role string `pulumi:"role"`
	// The name or relative resource id of the table to manage IAM policies for.
	Table string `pulumi:"table"`
}

// The set of arguments for constructing a TableIamBinding resource.
type TableIamBindingArgs struct {
	Condition TableIamBindingConditionPtrInput
	// The name or relative resource id of the instance that owns the table.
	Instance pulumi.StringInput
	Members  pulumi.StringArrayInput
	Project  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role pulumi.StringInput
	// The name or relative resource id of the table to manage IAM policies for.
	Table pulumi.StringInput
}

func (TableIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableIamBindingArgs)(nil)).Elem()
}
