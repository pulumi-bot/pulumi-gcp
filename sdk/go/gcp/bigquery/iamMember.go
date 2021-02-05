// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

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

func (*IamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*IamMember)(nil))
}

func (i *IamMember) ToIamMemberOutput() IamMemberOutput {
	return i.ToIamMemberOutputWithContext(context.Background())
}

func (i *IamMember) ToIamMemberOutputWithContext(ctx context.Context) IamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamMemberOutput)
}

func (i *IamMember) ToIamMemberPtrOutput() IamMemberPtrOutput {
	return i.ToIamMemberPtrOutputWithContext(context.Background())
}

func (i *IamMember) ToIamMemberPtrOutputWithContext(ctx context.Context) IamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamMemberPtrOutput)
}

type IamMemberPtrInput interface {
	pulumi.Input

	ToIamMemberPtrOutput() IamMemberPtrOutput
	ToIamMemberPtrOutputWithContext(ctx context.Context) IamMemberPtrOutput
}

type iamMemberPtrType IamMemberArgs

func (*iamMemberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IamMember)(nil))
}

func (i *iamMemberPtrType) ToIamMemberPtrOutput() IamMemberPtrOutput {
	return i.ToIamMemberPtrOutputWithContext(context.Background())
}

func (i *iamMemberPtrType) ToIamMemberPtrOutputWithContext(ctx context.Context) IamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamMemberPtrOutput)
}

// IamMemberArrayInput is an input type that accepts IamMemberArray and IamMemberArrayOutput values.
// You can construct a concrete instance of `IamMemberArrayInput` via:
//
//          IamMemberArray{ IamMemberArgs{...} }
type IamMemberArrayInput interface {
	pulumi.Input

	ToIamMemberArrayOutput() IamMemberArrayOutput
	ToIamMemberArrayOutputWithContext(context.Context) IamMemberArrayOutput
}

type IamMemberArray []IamMemberInput

func (IamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*IamMember)(nil))
}

func (i IamMemberArray) ToIamMemberArrayOutput() IamMemberArrayOutput {
	return i.ToIamMemberArrayOutputWithContext(context.Background())
}

func (i IamMemberArray) ToIamMemberArrayOutputWithContext(ctx context.Context) IamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamMemberArrayOutput)
}

// IamMemberMapInput is an input type that accepts IamMemberMap and IamMemberMapOutput values.
// You can construct a concrete instance of `IamMemberMapInput` via:
//
//          IamMemberMap{ "key": IamMemberArgs{...} }
type IamMemberMapInput interface {
	pulumi.Input

	ToIamMemberMapOutput() IamMemberMapOutput
	ToIamMemberMapOutputWithContext(context.Context) IamMemberMapOutput
}

type IamMemberMap map[string]IamMemberInput

func (IamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*IamMember)(nil))
}

func (i IamMemberMap) ToIamMemberMapOutput() IamMemberMapOutput {
	return i.ToIamMemberMapOutputWithContext(context.Background())
}

func (i IamMemberMap) ToIamMemberMapOutputWithContext(ctx context.Context) IamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamMemberMapOutput)
}

type IamMemberOutput struct {
	*pulumi.OutputState
}

func (IamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IamMember)(nil))
}

func (o IamMemberOutput) ToIamMemberOutput() IamMemberOutput {
	return o
}

func (o IamMemberOutput) ToIamMemberOutputWithContext(ctx context.Context) IamMemberOutput {
	return o
}

func (o IamMemberOutput) ToIamMemberPtrOutput() IamMemberPtrOutput {
	return o.ToIamMemberPtrOutputWithContext(context.Background())
}

func (o IamMemberOutput) ToIamMemberPtrOutputWithContext(ctx context.Context) IamMemberPtrOutput {
	return o.ApplyT(func(v IamMember) *IamMember {
		return &v
	}).(IamMemberPtrOutput)
}

type IamMemberPtrOutput struct {
	*pulumi.OutputState
}

func (IamMemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamMember)(nil))
}

func (o IamMemberPtrOutput) ToIamMemberPtrOutput() IamMemberPtrOutput {
	return o
}

func (o IamMemberPtrOutput) ToIamMemberPtrOutputWithContext(ctx context.Context) IamMemberPtrOutput {
	return o
}

type IamMemberArrayOutput struct{ *pulumi.OutputState }

func (IamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IamMember)(nil))
}

func (o IamMemberArrayOutput) ToIamMemberArrayOutput() IamMemberArrayOutput {
	return o
}

func (o IamMemberArrayOutput) ToIamMemberArrayOutputWithContext(ctx context.Context) IamMemberArrayOutput {
	return o
}

func (o IamMemberArrayOutput) Index(i pulumi.IntInput) IamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IamMember {
		return vs[0].([]IamMember)[vs[1].(int)]
	}).(IamMemberOutput)
}

type IamMemberMapOutput struct{ *pulumi.OutputState }

func (IamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IamMember)(nil))
}

func (o IamMemberMapOutput) ToIamMemberMapOutput() IamMemberMapOutput {
	return o
}

func (o IamMemberMapOutput) ToIamMemberMapOutputWithContext(ctx context.Context) IamMemberMapOutput {
	return o
}

func (o IamMemberMapOutput) MapIndex(k pulumi.StringInput) IamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IamMember {
		return vs[0].(map[string]IamMember)[vs[1].(string)]
	}).(IamMemberOutput)
}

func init() {
	pulumi.RegisterOutputType(IamMemberOutput{})
	pulumi.RegisterOutputType(IamMemberPtrOutput{})
	pulumi.RegisterOutputType(IamMemberArrayOutput{})
	pulumi.RegisterOutputType(IamMemberMapOutput{})
}
