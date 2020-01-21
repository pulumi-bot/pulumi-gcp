// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getFolder

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get information about a Google Cloud Folder.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/folder.html.markdown.
func GetFolder(ctx *pulumi.Context, args *GetFolderArgs, opts ...pulumi.InvokeOption) (*GetFolderResult, error) {
	var rv GetFolderResult
	err := ctx.Invoke("gcp:organizations/getFolder:getFolder", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFolder.
type GetFolderArgs struct {
	// The name of the Folder in the form `{folder_id}` or `folders/{folder_id}`.
	Folder string `pulumi:"folder"`
	// `true` to find the organization that the folder belongs, `false` to avoid the lookup. It searches up the tree. (defaults to `false`)
	LookupOrganization *bool `pulumi:"lookupOrganization"`
}


// A collection of values returned by getFolder.
type GetFolderResult struct {
	// Timestamp when the Organization was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime string `pulumi:"createTime"`
	// The folder's display name.
	DisplayName string `pulumi:"displayName"`
	Folder string `pulumi:"folder"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Folder's current lifecycle state.
	LifecycleState string `pulumi:"lifecycleState"`
	LookupOrganization *bool `pulumi:"lookupOrganization"`
	// The resource name of the Folder in the form `folders/{folder_id}`.
	Name string `pulumi:"name"`
	// If `lookupOrganization` is enable, the resource name of the Organization that the folder belongs.
	Organization string `pulumi:"organization"`
	// The resource name of the parent Folder or Organization.
	Parent string `pulumi:"parent"`
}

