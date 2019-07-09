// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package serviceAccount

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides a google `oauth2` `access_token` for a different service account than the one initially running the script.
// 
// For more information see
// [the official documentation](https://cloud.google.com/iam/docs/creating-short-lived-service-account-credentials) as well as [iamcredentials.generateAccessToken()](https://cloud.google.com/iam/credentials/reference/rest/v1/projects.serviceAccounts/generateAccessToken)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/service_account_access_token.html.markdown.
func LookupAccountAccessToken(ctx *pulumi.Context, args *GetAccountAccessTokenArgs) (*GetAccountAccessTokenResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["delegates"] = args.Delegates
		inputs["lifetime"] = args.Lifetime
		inputs["scopes"] = args.Scopes
		inputs["targetServiceAccount"] = args.TargetServiceAccount
	}
	outputs, err := ctx.Invoke("gcp:serviceAccount/getAccountAccessToken:getAccountAccessToken", inputs)
	if err != nil {
		return nil, err
	}
	return &GetAccountAccessTokenResult{
		AccessToken: outputs["accessToken"],
		Delegates: outputs["delegates"],
		Lifetime: outputs["lifetime"],
		Scopes: outputs["scopes"],
		TargetServiceAccount: outputs["targetServiceAccount"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getAccountAccessToken.
type GetAccountAccessTokenArgs struct {
	// Deegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.  (e.g. `["projects/-/serviceAccounts/delegate-svc-account@project-id.iam.gserviceaccount.com"]`)
	Delegates interface{}
	// Lifetime of the impersonated token (defaults to its max: `3600s`).
	Lifetime interface{}
	// The scopes the new credential should have (e.g. `["storage-ro", "cloud-platform"]`)
	Scopes interface{}
	// The service account _to_ impersonate (e.g. `service_B@your-project-id.iam.gserviceaccount.com`)
	TargetServiceAccount interface{}
}

// A collection of values returned by getAccountAccessToken.
type GetAccountAccessTokenResult struct {
	// The `access_token` representing the new generated identity.
	AccessToken interface{}
	Delegates interface{}
	Lifetime interface{}
	Scopes interface{}
	TargetServiceAccount interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
