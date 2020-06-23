// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get information about a Google Cloud Folder.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myFolder1, err := organizations.LookupFolder(ctx, &organizations.LookupFolderArgs{
// 			Folder:             "folders/12345",
// 			LookupOrganization: true,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		myFolder2, err := organizations.LookupFolder(ctx, &organizations.LookupFolderArgs{
// 			Folder: "folders/23456",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("myFolder1Organization", myFolder1.Organization)
// 		ctx.Export("myFolder2Parent", myFolder2.Parent)
// 		return nil
// 	})
// }
// ```
func LookupFolder(ctx *pulumi.Context, args *LookupFolderArgs, opts ...pulumi.InvokeOption) (*LookupFolderResult, error) {
	var rv LookupFolderResult
	err := ctx.Invoke("gcp:organizations/getFolder:getFolder", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFolder.
type LookupFolderArgs struct {
	// The name of the Folder in the form `{folder_id}` or `folders/{folder_id}`.
	Folder string `pulumi:"folder"`
	// `true` to find the organization that the folder belongs, `false` to avoid the lookup. It searches up the tree. (defaults to `false`)
	LookupOrganization *bool `pulumi:"lookupOrganization"`
}

// A collection of values returned by getFolder.
type LookupFolderResult struct {
	// Timestamp when the Organization was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime string `pulumi:"createTime"`
	// The folder's display name.
	DisplayName string `pulumi:"displayName"`
	Folder      string `pulumi:"folder"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Folder's current lifecycle state.
	LifecycleState     string `pulumi:"lifecycleState"`
	LookupOrganization *bool  `pulumi:"lookupOrganization"`
	// The resource name of the Folder in the form `folders/{folder_id}`.
	Name string `pulumi:"name"`
	// If `lookupOrganization` is enable, the resource name of the Organization that the folder belongs.
	Organization string `pulumi:"organization"`
	// The resource name of the parent Folder or Organization.
	Parent string `pulumi:"parent"`
}
