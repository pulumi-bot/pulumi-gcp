// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package billing

import (
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
//     Use `import` and inspect the preview output to ensure
//     your existing members are preserved.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/billing_account_iam_binding.html.markdown.
type AccountIamBinding struct {
	s *pulumi.ResourceState
}

// NewAccountIamBinding registers a new resource with the given unique name, arguments, and options.
func NewAccountIamBinding(ctx *pulumi.Context,
	name string, args *AccountIamBindingArgs, opts ...pulumi.ResourceOpt) (*AccountIamBinding, error) {
	if args == nil || args.BillingAccountId == nil {
		return nil, errors.New("missing required argument 'BillingAccountId'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["billingAccountId"] = nil
		inputs["condition"] = nil
		inputs["members"] = nil
		inputs["role"] = nil
	} else {
		inputs["billingAccountId"] = args.BillingAccountId
		inputs["condition"] = args.Condition
		inputs["members"] = args.Members
		inputs["role"] = args.Role
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:billing/accountIamBinding:AccountIamBinding", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AccountIamBinding{s: s}, nil
}

// GetAccountIamBinding gets an existing AccountIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountIamBinding(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AccountIamBindingState, opts ...pulumi.ResourceOpt) (*AccountIamBinding, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["billingAccountId"] = state.BillingAccountId
		inputs["condition"] = state.Condition
		inputs["etag"] = state.Etag
		inputs["members"] = state.Members
		inputs["role"] = state.Role
	}
	s, err := ctx.ReadResource("gcp:billing/accountIamBinding:AccountIamBinding", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AccountIamBinding{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AccountIamBinding) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AccountIamBinding) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The billing account id.
func (r *AccountIamBinding) BillingAccountId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["billingAccountId"])
}

func (r *AccountIamBinding) Condition() *pulumi.Output {
	return r.s.State["condition"]
}

// (Computed) The etag of the billing account's IAM policy.
func (r *AccountIamBinding) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
func (r *AccountIamBinding) Members() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["members"])
}

// The role that should be applied.
func (r *AccountIamBinding) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// Input properties used for looking up and filtering AccountIamBinding resources.
type AccountIamBindingState struct {
	// The billing account id.
	BillingAccountId interface{}
	Condition interface{}
	// (Computed) The etag of the billing account's IAM policy.
	Etag interface{}
	// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members interface{}
	// The role that should be applied.
	Role interface{}
}

// The set of arguments for constructing a AccountIamBinding resource.
type AccountIamBindingArgs struct {
	// The billing account id.
	BillingAccountId interface{}
	Condition interface{}
	// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members interface{}
	// The role that should be applied.
	Role interface{}
}
