// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package billing

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AccountIamMember struct {
	pulumi.CustomResourceState

	BillingAccountId pulumi.StringOutput                `pulumi:"billingAccountId"`
	Condition        AccountIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag             pulumi.StringOutput                `pulumi:"etag"`
	Member           pulumi.StringOutput                `pulumi:"member"`
	Role             pulumi.StringOutput                `pulumi:"role"`
}

// NewAccountIamMember registers a new resource with the given unique name, arguments, and options.
func NewAccountIamMember(ctx *pulumi.Context,
	name string, args *AccountIamMemberArgs, opts ...pulumi.ResourceOption) (*AccountIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountId == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountId'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource AccountIamMember
	err := ctx.RegisterResource("gcp:billing/accountIamMember:AccountIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountIamMember gets an existing AccountIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountIamMemberState, opts ...pulumi.ResourceOption) (*AccountIamMember, error) {
	var resource AccountIamMember
	err := ctx.ReadResource("gcp:billing/accountIamMember:AccountIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountIamMember resources.
type accountIamMemberState struct {
	BillingAccountId *string                    `pulumi:"billingAccountId"`
	Condition        *AccountIamMemberCondition `pulumi:"condition"`
	Etag             *string                    `pulumi:"etag"`
	Member           *string                    `pulumi:"member"`
	Role             *string                    `pulumi:"role"`
}

type AccountIamMemberState struct {
	BillingAccountId pulumi.StringPtrInput
	Condition        AccountIamMemberConditionPtrInput
	Etag             pulumi.StringPtrInput
	Member           pulumi.StringPtrInput
	Role             pulumi.StringPtrInput
}

func (AccountIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountIamMemberState)(nil)).Elem()
}

type accountIamMemberArgs struct {
	BillingAccountId string                     `pulumi:"billingAccountId"`
	Condition        *AccountIamMemberCondition `pulumi:"condition"`
	Member           string                     `pulumi:"member"`
	Role             string                     `pulumi:"role"`
}

// The set of arguments for constructing a AccountIamMember resource.
type AccountIamMemberArgs struct {
	BillingAccountId pulumi.StringInput
	Condition        AccountIamMemberConditionPtrInput
	Member           pulumi.StringInput
	Role             pulumi.StringInput
}

func (AccountIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountIamMemberArgs)(nil)).Elem()
}

type AccountIamMemberInput interface {
	pulumi.Input

	ToAccountIamMemberOutput() AccountIamMemberOutput
	ToAccountIamMemberOutputWithContext(ctx context.Context) AccountIamMemberOutput
}

func (*AccountIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountIamMember)(nil))
}

func (i *AccountIamMember) ToAccountIamMemberOutput() AccountIamMemberOutput {
	return i.ToAccountIamMemberOutputWithContext(context.Background())
}

func (i *AccountIamMember) ToAccountIamMemberOutputWithContext(ctx context.Context) AccountIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIamMemberOutput)
}

type AccountIamMemberOutput struct {
	*pulumi.OutputState
}

func (AccountIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountIamMember)(nil))
}

func (o AccountIamMemberOutput) ToAccountIamMemberOutput() AccountIamMemberOutput {
	return o
}

func (o AccountIamMemberOutput) ToAccountIamMemberOutputWithContext(ctx context.Context) AccountIamMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AccountIamMemberOutput{})
}
