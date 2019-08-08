// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get the email address of a project's unique Google Cloud Storage service account.
// 
// Each Google Cloud project has a unique service account for use with Google Cloud Storage. Only this
// special service account can be used to set up `storage.Notification` resources.
// 
// For more information see
// [the API reference](https://cloud.google.com/storage/docs/json_api/v1/projects/serviceAccount).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/storage_project_service_account.html.markdown.
func LookupProjectServiceAccount(ctx *pulumi.Context, args *GetProjectServiceAccountArgs) (*GetProjectServiceAccountResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["project"] = args.Project
		inputs["userProject"] = args.UserProject
	}
	outputs, err := ctx.Invoke("gcp:storage/getProjectServiceAccount:getProjectServiceAccount", inputs)
	if err != nil {
		return nil, err
	}
	return &GetProjectServiceAccountResult{
		EmailAddress: outputs["emailAddress"],
		Project: outputs["project"],
		UserProject: outputs["userProject"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getProjectServiceAccount.
type GetProjectServiceAccountArgs struct {
	// The project the unique service account was created for. If it is not provided, the provider project is used.
	Project interface{}
	// The project the lookup originates from. This field is used if you are making the request
	// from a different account than the one you are finding the service account for.
	UserProject interface{}
}

// A collection of values returned by getProjectServiceAccount.
type GetProjectServiceAccountResult struct {
	// The email address of the service account. This value is often used to refer to the service account
	// in order to grant IAM permissions.
	EmailAddress interface{}
	Project interface{}
	UserProject interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
