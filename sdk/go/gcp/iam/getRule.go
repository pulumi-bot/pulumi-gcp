// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get information about a Google IAM Role.
func GetRule(ctx *pulumi.Context, args *GetRuleArgs, opts ...pulumi.InvokeOption) (*GetRuleResult, error) {
	var rv GetRuleResult
	err := ctx.Invoke("gcp:iam/getRule:getRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRule.
type GetRuleArgs struct {
	// The name of the Role to lookup in the form `roles/{ROLE_NAME}`, `organizations/{ORGANIZATION_ID}/roles/{ROLE_NAME}` or `projects/{PROJECT_ID}/roles/{ROLE_NAME}`
	Name string `pulumi:"name"`
}

// A collection of values returned by getRule.
type GetRuleResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// specifies the list of one or more permissions to include in the custom role, such as - `iam.roles.get`
	IncludedPermissions []string `pulumi:"includedPermissions"`
	Name                string   `pulumi:"name"`
	// indicates the stage of a role in the launch lifecycle, such as `GA`, `BETA` or `ALPHA`.
	Stage string `pulumi:"stage"`
	// is a friendly title for the role, such as "Role Viewer"
	Title string `pulumi:"title"`
}
