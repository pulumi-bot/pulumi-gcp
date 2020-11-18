// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigtable

import (
	"context"
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
type TableIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the tables's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name or relative resource id of the instance that owns the table.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	Project    pulumi.StringOutput `pulumi:"project"`
	// The name or relative resource id of the table to manage IAM policies for.
	Table pulumi.StringOutput `pulumi:"table"`
}

// NewTableIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewTableIamPolicy(ctx *pulumi.Context,
	name string, args *TableIamPolicyArgs, opts ...pulumi.ResourceOption) (*TableIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.Table == nil {
		return nil, errors.New("invalid value for required argument 'Table'")
	}
	var resource TableIamPolicy
	err := ctx.RegisterResource("gcp:bigtable/tableIamPolicy:TableIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTableIamPolicy gets an existing TableIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTableIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableIamPolicyState, opts ...pulumi.ResourceOption) (*TableIamPolicy, error) {
	var resource TableIamPolicy
	err := ctx.ReadResource("gcp:bigtable/tableIamPolicy:TableIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TableIamPolicy resources.
type tableIamPolicyState struct {
	// (Computed) The etag of the tables's IAM policy.
	Etag *string `pulumi:"etag"`
	// The name or relative resource id of the instance that owns the table.
	Instance *string `pulumi:"instance"`
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	Project    *string `pulumi:"project"`
	// The name or relative resource id of the table to manage IAM policies for.
	Table *string `pulumi:"table"`
}

type TableIamPolicyState struct {
	// (Computed) The etag of the tables's IAM policy.
	Etag pulumi.StringPtrInput
	// The name or relative resource id of the instance that owns the table.
	Instance pulumi.StringPtrInput
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	Project    pulumi.StringPtrInput
	// The name or relative resource id of the table to manage IAM policies for.
	Table pulumi.StringPtrInput
}

func (TableIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableIamPolicyState)(nil)).Elem()
}

type tableIamPolicyArgs struct {
	// The name or relative resource id of the instance that owns the table.
	Instance string `pulumi:"instance"`
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData string  `pulumi:"policyData"`
	Project    *string `pulumi:"project"`
	// The name or relative resource id of the table to manage IAM policies for.
	Table string `pulumi:"table"`
}

// The set of arguments for constructing a TableIamPolicy resource.
type TableIamPolicyArgs struct {
	// The name or relative resource id of the instance that owns the table.
	Instance pulumi.StringInput
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	Project    pulumi.StringPtrInput
	// The name or relative resource id of the table to manage IAM policies for.
	Table pulumi.StringInput
}

func (TableIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableIamPolicyArgs)(nil)).Elem()
}

type TableIamPolicyInput interface {
	pulumi.Input

	ToTableIamPolicyOutput() TableIamPolicyOutput
	ToTableIamPolicyOutputWithContext(ctx context.Context) TableIamPolicyOutput
}

func (TableIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*TableIamPolicy)(nil)).Elem()
}

func (i TableIamPolicy) ToTableIamPolicyOutput() TableIamPolicyOutput {
	return i.ToTableIamPolicyOutputWithContext(context.Background())
}

func (i TableIamPolicy) ToTableIamPolicyOutputWithContext(ctx context.Context) TableIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableIamPolicyOutput)
}

type TableIamPolicyOutput struct {
	*pulumi.OutputState
}

func (TableIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableIamPolicyOutput)(nil)).Elem()
}

func (o TableIamPolicyOutput) ToTableIamPolicyOutput() TableIamPolicyOutput {
	return o
}

func (o TableIamPolicyOutput) ToTableIamPolicyOutputWithContext(ctx context.Context) TableIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TableIamPolicyOutput{})
}
