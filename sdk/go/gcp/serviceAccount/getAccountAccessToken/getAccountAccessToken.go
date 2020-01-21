// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getAccountAccessToken

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides a google `oauth2` `accessToken` for a different service account than the one initially running the script.
// 
// For more information see
// [the official documentation](https://cloud.google.com/iam/docs/creating-short-lived-service-account-credentials) as well as [iamcredentials.generateAccessToken()](https://cloud.google.com/iam/credentials/reference/rest/v1/projects.serviceAccounts/generateAccessToken)
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/service_account_access_token.html.markdown.
func GetAccountAccessToken(ctx *pulumi.Context, args *GetAccountAccessTokenArgs, opts ...pulumi.InvokeOption) (*GetAccountAccessTokenResult, error) {
	var rv GetAccountAccessTokenResult
	err := ctx.Invoke("gcp:serviceAccount/getAccountAccessToken:getAccountAccessToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccountAccessToken.
type GetAccountAccessTokenArgs struct {
	// Delegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.  (e.g. `["projects/-/serviceAccounts/delegate-svc-account@project-id.iam.gserviceaccount.com"]`)
	Delegates []string `pulumi:"delegates"`
	// Lifetime of the impersonated token (defaults to its max: `3600s`).
	Lifetime *string `pulumi:"lifetime"`
	// The scopes the new credential should have (e.g. `["storage-ro", "cloud-platform"]`)
	Scopes []string `pulumi:"scopes"`
	// The service account _to_ impersonate (e.g. `service_B@your-project-id.iam.gserviceaccount.com`)
	TargetServiceAccount string `pulumi:"targetServiceAccount"`
}


// A collection of values returned by getAccountAccessToken.
type GetAccountAccessTokenResult struct {
	// The `accessToken` representing the new generated identity.
	AccessToken string `pulumi:"accessToken"`
	Delegates []string `pulumi:"delegates"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Lifetime *string `pulumi:"lifetime"`
	Scopes []string `pulumi:"scopes"`
	TargetServiceAccount string `pulumi:"targetServiceAccount"`
}

