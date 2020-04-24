// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get information about a Google Billing Account.
func GetBillingAccount(ctx *pulumi.Context, args *GetBillingAccountArgs, opts ...pulumi.InvokeOption) (*GetBillingAccountResult, error) {
	var rv GetBillingAccountResult
	err := ctx.Invoke("gcp:organizations/getBillingAccount:getBillingAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBillingAccount.
type GetBillingAccountArgs struct {
	// The name of the billing account in the form `{billing_account_id}` or `billingAccounts/{billing_account_id}`.
	BillingAccount *string `pulumi:"billingAccount"`
	// The display name of the billing account.
	DisplayName *string `pulumi:"displayName"`
	// `true` if the billing account is open, `false` if the billing account is closed.
	Open *bool `pulumi:"open"`
}

// A collection of values returned by getBillingAccount.
type GetBillingAccountResult struct {
	BillingAccount *string `pulumi:"billingAccount"`
	DisplayName    string  `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The resource name of the billing account in the form `billingAccounts/{billing_account_id}`.
	Name string `pulumi:"name"`
	Open bool   `pulumi:"open"`
	// The IDs of any projects associated with the billing account.
	ProjectIds []string `pulumi:"projectIds"`
}
