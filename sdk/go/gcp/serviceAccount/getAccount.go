// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package serviceAccount

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get the service account from a project. For more information see
// the official [API](https://cloud.google.com/compute/docs/access/service-accounts) documentation.
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_google_service_account.html.markdown.
func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("gcp:serviceAccount/getAccount:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccount.
type LookupAccountArgs struct {
	// The Service account id.  (This is the part of the service account's email field that comes before the @ symbol.)
	AccountId string `pulumi:"accountId"`
	// The ID of the project that the service account is present in.
	// Defaults to the provider project configuration.
	Project *string `pulumi:"project"`
}


// A collection of values returned by getAccount.
type LookupAccountResult struct {
	AccountId string `pulumi:"accountId"`
	// The display name for the service account.
	DisplayName string `pulumi:"displayName"`
	// The e-mail address of the service account. This value
	// should be referenced from any `organizations.getIAMPolicy` data sources
	// that would grant the service account privileges.
	Email string `pulumi:"email"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The fully-qualified name of the service account.
	Name string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The unique id of the service account.
	UniqueId string `pulumi:"uniqueId"`
}

