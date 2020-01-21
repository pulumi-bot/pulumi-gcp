// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getIAMPolicy

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/organizations/getIAMPolicyAuditConfig"
	"https:/github.com/pulumi/pulumi-gcp/organizations/getIAMPolicyAuditConfigAuditLogConfig"
	"https:/github.com/pulumi/pulumi-gcp/organizations/getIAMPolicyBinding"
	"https:/github.com/pulumi/pulumi-gcp/organizations/getIAMPolicyBindingCondition"
)

// Generates an IAM policy document that may be referenced by and applied to
// other Google Cloud Platform resources, such as the `organizations.Project` resource.
// 
// 
// This data source is used to define IAM policies to apply to other resources.
// Currently, defining a policy through a datasource and referencing that policy
// from another resource is the only way to apply an IAM policy to a resource.
// 
// **Note:** Several restrictions apply when setting IAM policies through this API.
// See the [setIamPolicy docs](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy)
// for a list of these restrictions.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/iam_policy.html.markdown.
func GetIAMPolicy(ctx *pulumi.Context, args *GetIAMPolicyArgs, opts ...pulumi.InvokeOption) (*GetIAMPolicyResult, error) {
	var rv GetIAMPolicyResult
	err := ctx.Invoke("gcp:organizations/getIAMPolicy:getIAMPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIAMPolicy.
type GetIAMPolicyArgs struct {
	// A nested configuration block that defines logging additional configuration for your project.
	AuditConfigs []organizationsgetIAMPolicyAuditConfig.GetIAMPolicyAuditConfig `pulumi:"auditConfigs"`
	// A nested configuration block (described below)
	// defining a binding to be included in the policy document. Multiple
	// `binding` arguments are supported.
	Bindings []organizationsgetIAMPolicyBinding.GetIAMPolicyBinding `pulumi:"bindings"`
}


// A collection of values returned by getIAMPolicy.
type GetIAMPolicyResult struct {
	AuditConfigs []organizationsgetIAMPolicyAuditConfig.GetIAMPolicyAuditConfig `pulumi:"auditConfigs"`
	Bindings []organizationsgetIAMPolicyBinding.GetIAMPolicyBinding `pulumi:"bindings"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The above bindings serialized in a format suitable for
	// referencing from a resource that supports IAM.
	PolicyData string `pulumi:"policyData"`
}

