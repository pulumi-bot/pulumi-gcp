// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* {{policy_tag}} Any variables not passed in the import command will be taken from the provider configuration. Data catalog policytag IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:datacatalog/policyTagIamMember:PolicyTagIamMember editor "{{policy_tag}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:datacatalog/policyTagIamMember:PolicyTagIamMember editor "{{policy_tag}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:datacatalog/policyTagIamMember:PolicyTagIamMember editor {{policy_tag}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type PolicyTagIamMember struct {
	pulumi.CustomResourceState

	Condition PolicyTagIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	PolicyTag pulumi.StringOutput `pulumi:"policyTag"`
	// The role that should be applied. Only one
	// `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewPolicyTagIamMember registers a new resource with the given unique name, arguments, and options.
func NewPolicyTagIamMember(ctx *pulumi.Context,
	name string, args *PolicyTagIamMemberArgs, opts ...pulumi.ResourceOption) (*PolicyTagIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.PolicyTag == nil {
		return nil, errors.New("invalid value for required argument 'PolicyTag'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource PolicyTagIamMember
	err := ctx.RegisterResource("gcp:datacatalog/policyTagIamMember:PolicyTagIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyTagIamMember gets an existing PolicyTagIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyTagIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyTagIamMemberState, opts ...pulumi.ResourceOption) (*PolicyTagIamMember, error) {
	var resource PolicyTagIamMember
	err := ctx.ReadResource("gcp:datacatalog/policyTagIamMember:PolicyTagIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyTagIamMember resources.
type policyTagIamMemberState struct {
	Condition *PolicyTagIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	PolicyTag *string `pulumi:"policyTag"`
	// The role that should be applied. Only one
	// `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type PolicyTagIamMemberState struct {
	Condition PolicyTagIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	PolicyTag pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (PolicyTagIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyTagIamMemberState)(nil)).Elem()
}

type policyTagIamMemberArgs struct {
	Condition *PolicyTagIamMemberCondition `pulumi:"condition"`
	Member    string                       `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	PolicyTag string `pulumi:"policyTag"`
	// The role that should be applied. Only one
	// `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a PolicyTagIamMember resource.
type PolicyTagIamMemberArgs struct {
	Condition PolicyTagIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	PolicyTag pulumi.StringInput
	// The role that should be applied. Only one
	// `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (PolicyTagIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyTagIamMemberArgs)(nil)).Elem()
}

type PolicyTagIamMemberInput interface {
	pulumi.Input

	ToPolicyTagIamMemberOutput() PolicyTagIamMemberOutput
	ToPolicyTagIamMemberOutputWithContext(ctx context.Context) PolicyTagIamMemberOutput
}

func (*PolicyTagIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyTagIamMember)(nil))
}

func (i *PolicyTagIamMember) ToPolicyTagIamMemberOutput() PolicyTagIamMemberOutput {
	return i.ToPolicyTagIamMemberOutputWithContext(context.Background())
}

func (i *PolicyTagIamMember) ToPolicyTagIamMemberOutputWithContext(ctx context.Context) PolicyTagIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTagIamMemberOutput)
}

func (i *PolicyTagIamMember) ToPolicyTagIamMemberPtrOutput() PolicyTagIamMemberPtrOutput {
	return i.ToPolicyTagIamMemberPtrOutputWithContext(context.Background())
}

func (i *PolicyTagIamMember) ToPolicyTagIamMemberPtrOutputWithContext(ctx context.Context) PolicyTagIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTagIamMemberPtrOutput)
}

type PolicyTagIamMemberPtrInput interface {
	pulumi.Input

	ToPolicyTagIamMemberPtrOutput() PolicyTagIamMemberPtrOutput
	ToPolicyTagIamMemberPtrOutputWithContext(ctx context.Context) PolicyTagIamMemberPtrOutput
}

type policyTagIamMemberPtrType PolicyTagIamMemberArgs

func (*policyTagIamMemberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyTagIamMember)(nil))
}

func (i *policyTagIamMemberPtrType) ToPolicyTagIamMemberPtrOutput() PolicyTagIamMemberPtrOutput {
	return i.ToPolicyTagIamMemberPtrOutputWithContext(context.Background())
}

func (i *policyTagIamMemberPtrType) ToPolicyTagIamMemberPtrOutputWithContext(ctx context.Context) PolicyTagIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTagIamMemberPtrOutput)
}

// PolicyTagIamMemberArrayInput is an input type that accepts PolicyTagIamMemberArray and PolicyTagIamMemberArrayOutput values.
// You can construct a concrete instance of `PolicyTagIamMemberArrayInput` via:
//
//          PolicyTagIamMemberArray{ PolicyTagIamMemberArgs{...} }
type PolicyTagIamMemberArrayInput interface {
	pulumi.Input

	ToPolicyTagIamMemberArrayOutput() PolicyTagIamMemberArrayOutput
	ToPolicyTagIamMemberArrayOutputWithContext(context.Context) PolicyTagIamMemberArrayOutput
}

type PolicyTagIamMemberArray []PolicyTagIamMemberInput

func (PolicyTagIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*PolicyTagIamMember)(nil))
}

func (i PolicyTagIamMemberArray) ToPolicyTagIamMemberArrayOutput() PolicyTagIamMemberArrayOutput {
	return i.ToPolicyTagIamMemberArrayOutputWithContext(context.Background())
}

func (i PolicyTagIamMemberArray) ToPolicyTagIamMemberArrayOutputWithContext(ctx context.Context) PolicyTagIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTagIamMemberArrayOutput)
}

// PolicyTagIamMemberMapInput is an input type that accepts PolicyTagIamMemberMap and PolicyTagIamMemberMapOutput values.
// You can construct a concrete instance of `PolicyTagIamMemberMapInput` via:
//
//          PolicyTagIamMemberMap{ "key": PolicyTagIamMemberArgs{...} }
type PolicyTagIamMemberMapInput interface {
	pulumi.Input

	ToPolicyTagIamMemberMapOutput() PolicyTagIamMemberMapOutput
	ToPolicyTagIamMemberMapOutputWithContext(context.Context) PolicyTagIamMemberMapOutput
}

type PolicyTagIamMemberMap map[string]PolicyTagIamMemberInput

func (PolicyTagIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*PolicyTagIamMember)(nil))
}

func (i PolicyTagIamMemberMap) ToPolicyTagIamMemberMapOutput() PolicyTagIamMemberMapOutput {
	return i.ToPolicyTagIamMemberMapOutputWithContext(context.Background())
}

func (i PolicyTagIamMemberMap) ToPolicyTagIamMemberMapOutputWithContext(ctx context.Context) PolicyTagIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTagIamMemberMapOutput)
}

type PolicyTagIamMemberOutput struct {
	*pulumi.OutputState
}

func (PolicyTagIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyTagIamMember)(nil))
}

func (o PolicyTagIamMemberOutput) ToPolicyTagIamMemberOutput() PolicyTagIamMemberOutput {
	return o
}

func (o PolicyTagIamMemberOutput) ToPolicyTagIamMemberOutputWithContext(ctx context.Context) PolicyTagIamMemberOutput {
	return o
}

func (o PolicyTagIamMemberOutput) ToPolicyTagIamMemberPtrOutput() PolicyTagIamMemberPtrOutput {
	return o.ToPolicyTagIamMemberPtrOutputWithContext(context.Background())
}

func (o PolicyTagIamMemberOutput) ToPolicyTagIamMemberPtrOutputWithContext(ctx context.Context) PolicyTagIamMemberPtrOutput {
	return o.ApplyT(func(v PolicyTagIamMember) *PolicyTagIamMember {
		return &v
	}).(PolicyTagIamMemberPtrOutput)
}

type PolicyTagIamMemberPtrOutput struct {
	*pulumi.OutputState
}

func (PolicyTagIamMemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyTagIamMember)(nil))
}

func (o PolicyTagIamMemberPtrOutput) ToPolicyTagIamMemberPtrOutput() PolicyTagIamMemberPtrOutput {
	return o
}

func (o PolicyTagIamMemberPtrOutput) ToPolicyTagIamMemberPtrOutputWithContext(ctx context.Context) PolicyTagIamMemberPtrOutput {
	return o
}

type PolicyTagIamMemberArrayOutput struct{ *pulumi.OutputState }

func (PolicyTagIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyTagIamMember)(nil))
}

func (o PolicyTagIamMemberArrayOutput) ToPolicyTagIamMemberArrayOutput() PolicyTagIamMemberArrayOutput {
	return o
}

func (o PolicyTagIamMemberArrayOutput) ToPolicyTagIamMemberArrayOutputWithContext(ctx context.Context) PolicyTagIamMemberArrayOutput {
	return o
}

func (o PolicyTagIamMemberArrayOutput) Index(i pulumi.IntInput) PolicyTagIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyTagIamMember {
		return vs[0].([]PolicyTagIamMember)[vs[1].(int)]
	}).(PolicyTagIamMemberOutput)
}

type PolicyTagIamMemberMapOutput struct{ *pulumi.OutputState }

func (PolicyTagIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PolicyTagIamMember)(nil))
}

func (o PolicyTagIamMemberMapOutput) ToPolicyTagIamMemberMapOutput() PolicyTagIamMemberMapOutput {
	return o
}

func (o PolicyTagIamMemberMapOutput) ToPolicyTagIamMemberMapOutputWithContext(ctx context.Context) PolicyTagIamMemberMapOutput {
	return o
}

func (o PolicyTagIamMemberMapOutput) MapIndex(k pulumi.StringInput) PolicyTagIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PolicyTagIamMember {
		return vs[0].(map[string]PolicyTagIamMember)[vs[1].(string)]
	}).(PolicyTagIamMemberOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyTagIamMemberOutput{})
	pulumi.RegisterOutputType(PolicyTagIamMemberPtrOutput{})
	pulumi.RegisterOutputType(PolicyTagIamMemberArrayOutput{})
	pulumi.RegisterOutputType(PolicyTagIamMemberMapOutput{})
}
