// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get information about a Google IAM Role.
func LookupRule(ctx *pulumi.Context, args *GetRuleArgs) (*GetRuleResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("gcp:iam/getRule:getRule", inputs)
	if err != nil {
		return nil, err
	}
	return &GetRuleResult{
		IncludedPermissions: outputs["includedPermissions"],
		Stage: outputs["stage"],
		Title: outputs["title"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getRule.
type GetRuleArgs struct {
	// The name of the Role to lookup in the form `roles/{ROLE_NAME}`, `organizations/{ORGANIZATION_ID}/roles/{ROLE_NAME}` or `projects/{PROJECT_ID}/roles/{ROLE_NAME}`
	Name interface{}
}

// A collection of values returned by getRule.
type GetRuleResult struct {
	// specifies the list of one or more permissions to include in the custom role, such as - `iam.roles.get`
	IncludedPermissions interface{}
	// indicates the stage of a role in the launch lifecycle, such as `GA`, `BETA` or `ALPHA`.
	Stage interface{}
	// is a friendly title for the role, such as "Role Viewer"
	Title interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
