// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Generates an IAM policy document that may be referenced by and applied to
// other Google Cloud Platform IAM resources, such as the `projects.IAMPolicy` resource.
//
// **Note:** Please review the documentation of the resource that you will be using the datasource with. Some resources such as `projects.IAMPolicy` and others have limitations in their API methods which are noted on their respective page.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			AuditConfigs: []organizations.GetIAMPolicyAuditConfig{
// 				organizations.GetIAMPolicyAuditConfig{
// 					AuditLogConfigs: []organizations.GetIAMPolicyAuditConfigAuditLogConfig{
// 						organizations.GetIAMPolicyAuditConfigAuditLogConfig{
// 							ExemptedMembers: []string{
// 								"user:you@domain.com",
// 							},
// 							LogType: "DATA_READ",
// 						},
// 						organizations.GetIAMPolicyAuditConfigAuditLogConfig{
// 							LogType: "DATA_WRITE",
// 						},
// 						organizations.GetIAMPolicyAuditConfigAuditLogConfig{
// 							LogType: "ADMIN_READ",
// 						},
// 					},
// 					Service: "cloudkms.googleapis.com",
// 				},
// 			},
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Members: []string{
// 						"serviceAccount:your-custom-sa@your-project.iam.gserviceaccount.com",
// 					},
// 					Role: "roles/compute.instanceAdmin",
// 				},
// 				organizations.GetIAMPolicyBinding{
// 					Members: []string{
// 						"user:alice@gmail.com",
// 					},
// 					Role: "roles/storage.objectViewer",
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
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
	// A nested configuration block that defines logging additional configuration for your project. This field is only supported on `projects.IAMPolicy`, `folder.IAMPolicy` and `organizations.IAMPolicy`.
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
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The above bindings serialized in a format suitable for
	// referencing from a resource that supports IAM.
	PolicyData string `pulumi:"policyData"`
}
