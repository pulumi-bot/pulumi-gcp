// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package billing

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows management of the entire IAM policy for an existing Google Cloud Platform Billing Account.
//
// > **Warning:** Billing accounts have a default user that can be **overwritten**
// by use of this resource. The safest alternative is to use multiple `billing.AccountIamBinding`
//    resources. If you do use this resource, the best way to be sure that you are
//    not making dangerous changes is to start by importing your existing policy,
//    and examining the diff very closely.
//
// > **Note:** This resource __must not__ be used in conjunction with
//    `billing.AccountIamMember` or `billing.AccountIamBinding`
//    or they will fight over what your policy should be.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/billing"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Binding: []map[string]interface{}{
// 				map[string]interface{}{
// 					"role": "roles/billing.viewer",
// 					"members": []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = billing.NewAccountIamPolicy(ctx, "policy", &billing.AccountIamPolicyArgs{
// 			BillingAccountId: pulumi.String("00AA00-000AAA-00AA0A"),
// 			PolicyData:       pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AccountIamPolicy struct {
	pulumi.CustomResourceState

	// The billing account id.
	BillingAccountId pulumi.StringOutput `pulumi:"billingAccountId"`
	Etag             pulumi.StringOutput `pulumi:"etag"`
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the billing account. This policy overrides any existing
	// policy applied to the billing account.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
}

// NewAccountIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewAccountIamPolicy(ctx *pulumi.Context,
	name string, args *AccountIamPolicyArgs, opts ...pulumi.ResourceOption) (*AccountIamPolicy, error) {
	if args == nil || args.BillingAccountId == nil {
		return nil, errors.New("missing required argument 'BillingAccountId'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil {
		args = &AccountIamPolicyArgs{}
	}
	var resource AccountIamPolicy
	err := ctx.RegisterResource("gcp:billing/accountIamPolicy:AccountIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountIamPolicy gets an existing AccountIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountIamPolicyState, opts ...pulumi.ResourceOption) (*AccountIamPolicy, error) {
	var resource AccountIamPolicy
	err := ctx.ReadResource("gcp:billing/accountIamPolicy:AccountIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountIamPolicy resources.
type accountIamPolicyState struct {
	// The billing account id.
	BillingAccountId *string `pulumi:"billingAccountId"`
	Etag             *string `pulumi:"etag"`
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the billing account. This policy overrides any existing
	// policy applied to the billing account.
	PolicyData *string `pulumi:"policyData"`
}

type AccountIamPolicyState struct {
	// The billing account id.
	BillingAccountId pulumi.StringPtrInput
	Etag             pulumi.StringPtrInput
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the billing account. This policy overrides any existing
	// policy applied to the billing account.
	PolicyData pulumi.StringPtrInput
}

func (AccountIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountIamPolicyState)(nil)).Elem()
}

type accountIamPolicyArgs struct {
	// The billing account id.
	BillingAccountId string `pulumi:"billingAccountId"`
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the billing account. This policy overrides any existing
	// policy applied to the billing account.
	PolicyData string `pulumi:"policyData"`
}

// The set of arguments for constructing a AccountIamPolicy resource.
type AccountIamPolicyArgs struct {
	// The billing account id.
	BillingAccountId pulumi.StringInput
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the billing account. This policy overrides any existing
	// policy applied to the billing account.
	PolicyData pulumi.StringInput
}

func (AccountIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountIamPolicyArgs)(nil)).Elem()
}
