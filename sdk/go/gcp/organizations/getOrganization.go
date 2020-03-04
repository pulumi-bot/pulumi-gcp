// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package organizations

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get information about a Google Cloud Organization.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/google_organization.html.markdown.
func GetOrganization(ctx *pulumi.Context, args *GetOrganizationArgs, opts ...pulumi.InvokeOption) (*GetOrganizationResult, error) {
	var rv GetOrganizationResult
	err := ctx.Invoke("gcp:organizations/getOrganization:getOrganization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganization.
type GetOrganizationArgs struct {
	// The domain name of the Organization.
	Domain *string `pulumi:"domain"`
	// The name of the Organization in the form `{organization_id}` or `organizations/{organization_id}`.
	Organization *string `pulumi:"organization"`
}


// A collection of values returned by getOrganization.
type GetOrganizationResult struct {
	// Timestamp when the Organization was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime string `pulumi:"createTime"`
	// The Google for Work customer ID of the Organization.
	DirectoryCustomerId string `pulumi:"directoryCustomerId"`
	Domain string `pulumi:"domain"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Organization's current lifecycle state.
	LifecycleState string `pulumi:"lifecycleState"`
	// The resource name of the Organization in the form `organizations/{organization_id}`.
	Name string `pulumi:"name"`
	// The Organization ID.
	OrgId string `pulumi:"orgId"`
	Organization *string `pulumi:"organization"`
}

