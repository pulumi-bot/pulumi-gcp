// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package billing

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows creation and management of a single binding within IAM policy for
// an existing Google Cloud Platform Billing Account.
//
// > **Note:** This resource __must not__ be used in conjunction with
//    `billing.AccountIamMember` for the __same role__ or they will fight over
//    what your policy should be.
//
// > **Note:** On create, this resource will overwrite members of any existing roles.
//     Use `pulumi import` and inspect the output to ensure
//     your existing members are preserved.
type AccountIamBinding struct {
	pulumi.CustomResourceState

	// The billing account id.
	BillingAccountId pulumi.StringOutput                 `pulumi:"billingAccountId"`
	Condition        AccountIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the billing account's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The role that should be applied.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewAccountIamBinding registers a new resource with the given unique name, arguments, and options.
func NewAccountIamBinding(ctx *pulumi.Context,
	name string, args *AccountIamBindingArgs, opts ...pulumi.ResourceOption) (*AccountIamBinding, error) {
	if args == nil || args.BillingAccountId == nil {
		return nil, errors.New("missing required argument 'BillingAccountId'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &AccountIamBindingArgs{}
	}
	var resource AccountIamBinding
	err := ctx.RegisterResource("gcp:billing/accountIamBinding:AccountIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountIamBinding gets an existing AccountIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountIamBindingState, opts ...pulumi.ResourceOption) (*AccountIamBinding, error) {
	var resource AccountIamBinding
	err := ctx.ReadResource("gcp:billing/accountIamBinding:AccountIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountIamBinding resources.
type accountIamBindingState struct {
	// The billing account id.
	BillingAccountId *string                     `pulumi:"billingAccountId"`
	Condition        *AccountIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the billing account's IAM policy.
	Etag *string `pulumi:"etag"`
	// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members []string `pulumi:"members"`
	// The role that should be applied.
	Role *string `pulumi:"role"`
}

type AccountIamBindingState struct {
	// The billing account id.
	BillingAccountId pulumi.StringPtrInput
	Condition        AccountIamBindingConditionPtrInput
	// (Computed) The etag of the billing account's IAM policy.
	Etag pulumi.StringPtrInput
	// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members pulumi.StringArrayInput
	// The role that should be applied.
	Role pulumi.StringPtrInput
}

func (AccountIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountIamBindingState)(nil)).Elem()
}

type accountIamBindingArgs struct {
	// The billing account id.
	BillingAccountId string                      `pulumi:"billingAccountId"`
	Condition        *AccountIamBindingCondition `pulumi:"condition"`
	// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members []string `pulumi:"members"`
	// The role that should be applied.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a AccountIamBinding resource.
type AccountIamBindingArgs struct {
	// The billing account id.
	BillingAccountId pulumi.StringInput
	Condition        AccountIamBindingConditionPtrInput
	// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members pulumi.StringArrayInput
	// The role that should be applied.
	Role pulumi.StringInput
}

func (AccountIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountIamBindingArgs)(nil)).Elem()
}
