// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows management of a Google Cloud Platform folder. For more information see
// [the official documentation](https://cloud.google.com/resource-manager/docs/creating-managing-folders)
// and
// [API](https://cloud.google.com/resource-manager/reference/rest/v2/folders).
//
// A folder can contain projects, other folders, or a combination of both. You can use folders to group projects under an organization in a hierarchy. For example, your organization might contain multiple departments, each with its own set of Cloud Platform resources. Folders allows you to group these resources on a per-department basis. Folders are used to group resources that share common IAM policies.
//
// Folders created live inside an Organization. See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.
//
// The service account used to run the provider when creating a `organizations.Folder`
// resource must have `roles/resourcemanager.folderCreator`. See the
// [Access Control for Folders Using IAM](https://cloud.google.com/resource-manager/docs/access-control-folders)
// doc for more information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		department1, err := organizations.NewFolder(ctx, "department1", &organizations.FolderArgs{
// 			DisplayName: pulumi.String("Department 1"),
// 			Parent:      pulumi.String("organizations/1234567"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = organizations.NewFolder(ctx, "team-abc", &organizations.FolderArgs{
// 			DisplayName: pulumi.String("Team ABC"),
// 			Parent:      department1.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Folders can be imported using the folder's id, e.g. # Both syntaxes are valid
//
// ```sh
//  $ pulumi import gcp:organizations/folder:Folder department1 1234567
// ```
//
// ```sh
//  $ pulumi import gcp:organizations/folder:Folder department1 folders/1234567
// ```
type Folder struct {
	pulumi.CustomResourceState

	// Timestamp when the Folder was created. Assigned by the server.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The folder’s display name.
	// A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The folder id from the name "folders/{folder_id}"
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// The lifecycle state of the folder such as `ACTIVE` or `DELETE_REQUESTED`.
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	// The resource name of the Folder. Its format is folders/{folder_id}.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource name of the parent Folder or Organization.
	// Must be of the form `folders/{folder_id}` or `organizations/{org_id}`.
	Parent pulumi.StringOutput `pulumi:"parent"`
}

// NewFolder registers a new resource with the given unique name, arguments, and options.
func NewFolder(ctx *pulumi.Context,
	name string, args *FolderArgs, opts ...pulumi.ResourceOption) (*Folder, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	var resource Folder
	err := ctx.RegisterResource("gcp:organizations/folder:Folder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolder gets an existing Folder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderState, opts ...pulumi.ResourceOption) (*Folder, error) {
	var resource Folder
	err := ctx.ReadResource("gcp:organizations/folder:Folder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Folder resources.
type folderState struct {
	// Timestamp when the Folder was created. Assigned by the server.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime *string `pulumi:"createTime"`
	// The folder’s display name.
	// A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
	DisplayName *string `pulumi:"displayName"`
	// The folder id from the name "folders/{folder_id}"
	FolderId *string `pulumi:"folderId"`
	// The lifecycle state of the folder such as `ACTIVE` or `DELETE_REQUESTED`.
	LifecycleState *string `pulumi:"lifecycleState"`
	// The resource name of the Folder. Its format is folders/{folder_id}.
	Name *string `pulumi:"name"`
	// The resource name of the parent Folder or Organization.
	// Must be of the form `folders/{folder_id}` or `organizations/{org_id}`.
	Parent *string `pulumi:"parent"`
}

type FolderState struct {
	// Timestamp when the Folder was created. Assigned by the server.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringPtrInput
	// The folder’s display name.
	// A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
	DisplayName pulumi.StringPtrInput
	// The folder id from the name "folders/{folder_id}"
	FolderId pulumi.StringPtrInput
	// The lifecycle state of the folder such as `ACTIVE` or `DELETE_REQUESTED`.
	LifecycleState pulumi.StringPtrInput
	// The resource name of the Folder. Its format is folders/{folder_id}.
	Name pulumi.StringPtrInput
	// The resource name of the parent Folder or Organization.
	// Must be of the form `folders/{folder_id}` or `organizations/{org_id}`.
	Parent pulumi.StringPtrInput
}

func (FolderState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderState)(nil)).Elem()
}

type folderArgs struct {
	// The folder’s display name.
	// A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
	DisplayName string `pulumi:"displayName"`
	// The resource name of the parent Folder or Organization.
	// Must be of the form `folders/{folder_id}` or `organizations/{org_id}`.
	Parent string `pulumi:"parent"`
}

// The set of arguments for constructing a Folder resource.
type FolderArgs struct {
	// The folder’s display name.
	// A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
	DisplayName pulumi.StringInput
	// The resource name of the parent Folder or Organization.
	// Must be of the form `folders/{folder_id}` or `organizations/{org_id}`.
	Parent pulumi.StringInput
}

func (FolderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderArgs)(nil)).Elem()
}

type FolderInput interface {
	pulumi.Input

	ToFolderOutput() FolderOutput
	ToFolderOutputWithContext(ctx context.Context) FolderOutput
}

func (*Folder) ElementType() reflect.Type {
	return reflect.TypeOf((*Folder)(nil))
}

func (i *Folder) ToFolderOutput() FolderOutput {
	return i.ToFolderOutputWithContext(context.Background())
}

func (i *Folder) ToFolderOutputWithContext(ctx context.Context) FolderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderOutput)
}

type FolderOutput struct {
	*pulumi.OutputState
}

func (FolderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Folder)(nil))
}

func (o FolderOutput) ToFolderOutput() FolderOutput {
	return o
}

func (o FolderOutput) ToFolderOutputWithContext(ctx context.Context) FolderOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FolderOutput{})
}
