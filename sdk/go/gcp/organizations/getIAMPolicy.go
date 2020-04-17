// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Generates an IAM policy document that may be referenced by and applied to
// other Google Cloud Platform resources, such as the `organizations.Project` resource.
//
// **Note:** Several restrictions apply when setting IAM policies through this API.
// See the [setIamPolicy docs](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy)
// for a list of these restrictions.
//
//
// This data source is used to define IAM policies to apply to other resources.
// Currently, defining a policy through a datasource and referencing that policy
// from another resource is the only way to apply an IAM policy to a resource.
func LookupIAMPolicy(ctx *pulumi.Context, args *LookupIAMPolicyArgs, opts ...pulumi.InvokeOption) (*LookupIAMPolicyResult, error) {
	var rv LookupIAMPolicyResult
	err := ctx.Invoke("gcp:organizations/getIAMPolicy:getIAMPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIAMPolicy.
type LookupIAMPolicyArgs struct {
	// A nested configuration block that defines logging additional configuration for your project.
	// * `service` (Required) Defines a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
	// * `auditLogConfigs` (Required) A nested block that defines the operations you'd like to log.
	// * `logType` (Required) Defines the logging level. `DATA_READ`, `DATA_WRITE` and `ADMIN_READ` capture different types of events. See [the audit configuration documentation](https://cloud.google.com/resource-manager/reference/rest/Shared.Types/AuditConfig) for more details.
	// * `exemptedMembers` (Optional) Specifies the identities that are exempt from these types of logging operations. Follows the same format of the `members` array for `binding`.
	AuditConfigs []GetIAMPolicyAuditConfig `pulumi:"auditConfigs"`
	// A nested configuration block (described below)
	// defining a binding to be included in the policy document. Multiple
	// `binding` arguments are supported.
	Bindings []GetIAMPolicyBinding `pulumi:"bindings"`
}

// A collection of values returned by getIAMPolicy.
type LookupIAMPolicyResult struct {
	AuditConfigs []GetIAMPolicyAuditConfig `pulumi:"auditConfigs"`
	Bindings     []GetIAMPolicyBinding     `pulumi:"bindings"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The above bindings serialized in a format suitable for
	// referencing from a resource that supports IAM.
	PolicyData string `pulumi:"policyData"`
}
