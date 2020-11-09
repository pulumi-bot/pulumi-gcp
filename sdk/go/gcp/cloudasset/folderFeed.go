// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudasset

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes a Cloud Asset Inventory feed used to to listen to asset updates.
//
// To get more information about FolderFeed, see:
//
// * [API documentation](https://cloud.google.com/asset-inventory/docs/reference/rest/)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/asset-inventory/docs)
//
// ## Example Usage
type FolderFeed struct {
	pulumi.CustomResourceState

	// A list of the full names of the assets to receive updates. You must specify either or both of
	// assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
	// exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
	// See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
	AssetNames pulumi.StringArrayOutput `pulumi:"assetNames"`
	// A list of types of the assets to receive updates. You must specify either or both of assetNames
	// and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
	// the feed. For example: "compute.googleapis.com/Disk"
	// See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
	// supported asset types.
	AssetTypes pulumi.StringArrayOutput `pulumi:"assetTypes"`
	// The project whose identity will be used when sending messages to the
	// destination pubsub topic. It also specifies the project for API
	// enablement check, quota, and billing.
	BillingProject pulumi.StringOutput `pulumi:"billingProject"`
	// A condition which determines whether an asset update should be published. If specified, an asset
	// will be returned only when the expression evaluates to true. When set, expression field
	// must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
	// expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
	// condition are optional.
	// Structure is documented below.
	Condition FolderFeedConditionPtrOutput `pulumi:"condition"`
	// Asset content type. If not specified, no content but the asset name and type will be returned.
	// Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	// This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
	FeedId pulumi.StringOutput `pulumi:"feedId"`
	// Output configuration for asset feed destination.
	// Structure is documented below.
	FeedOutputConfig FolderFeedFeedOutputConfigOutput `pulumi:"feedOutputConfig"`
	// The folder this feed should be created in.
	Folder pulumi.StringOutput `pulumi:"folder"`
	// The ID of the folder where this feed has been created. Both [FOLDER_NUMBER] and folders/[FOLDER_NUMBER] are accepted.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// The format will be folders/{folder_number}/feeds/{client-assigned_feed_identifier}.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewFolderFeed registers a new resource with the given unique name, arguments, and options.
func NewFolderFeed(ctx *pulumi.Context,
	name string, args *FolderFeedArgs, opts ...pulumi.ResourceOption) (*FolderFeed, error) {
	if args == nil || args.BillingProject == nil {
		return nil, errors.New("missing required argument 'BillingProject'")
	}
	if args == nil || args.FeedId == nil {
		return nil, errors.New("missing required argument 'FeedId'")
	}
	if args == nil || args.FeedOutputConfig == nil {
		return nil, errors.New("missing required argument 'FeedOutputConfig'")
	}
	if args == nil || args.Folder == nil {
		return nil, errors.New("missing required argument 'Folder'")
	}
	if args == nil {
		args = &FolderFeedArgs{}
	}
	var resource FolderFeed
	err := ctx.RegisterResource("gcp:cloudasset/folderFeed:FolderFeed", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolderFeed gets an existing FolderFeed resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolderFeed(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderFeedState, opts ...pulumi.ResourceOption) (*FolderFeed, error) {
	var resource FolderFeed
	err := ctx.ReadResource("gcp:cloudasset/folderFeed:FolderFeed", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FolderFeed resources.
type folderFeedState struct {
	// A list of the full names of the assets to receive updates. You must specify either or both of
	// assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
	// exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
	// See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
	AssetNames []string `pulumi:"assetNames"`
	// A list of types of the assets to receive updates. You must specify either or both of assetNames
	// and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
	// the feed. For example: "compute.googleapis.com/Disk"
	// See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
	// supported asset types.
	AssetTypes []string `pulumi:"assetTypes"`
	// The project whose identity will be used when sending messages to the
	// destination pubsub topic. It also specifies the project for API
	// enablement check, quota, and billing.
	BillingProject *string `pulumi:"billingProject"`
	// A condition which determines whether an asset update should be published. If specified, an asset
	// will be returned only when the expression evaluates to true. When set, expression field
	// must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
	// expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
	// condition are optional.
	// Structure is documented below.
	Condition *FolderFeedCondition `pulumi:"condition"`
	// Asset content type. If not specified, no content but the asset name and type will be returned.
	// Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
	ContentType *string `pulumi:"contentType"`
	// This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
	FeedId *string `pulumi:"feedId"`
	// Output configuration for asset feed destination.
	// Structure is documented below.
	FeedOutputConfig *FolderFeedFeedOutputConfig `pulumi:"feedOutputConfig"`
	// The folder this feed should be created in.
	Folder *string `pulumi:"folder"`
	// The ID of the folder where this feed has been created. Both [FOLDER_NUMBER] and folders/[FOLDER_NUMBER] are accepted.
	FolderId *string `pulumi:"folderId"`
	// The format will be folders/{folder_number}/feeds/{client-assigned_feed_identifier}.
	Name *string `pulumi:"name"`
}

type FolderFeedState struct {
	// A list of the full names of the assets to receive updates. You must specify either or both of
	// assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
	// exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
	// See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
	AssetNames pulumi.StringArrayInput
	// A list of types of the assets to receive updates. You must specify either or both of assetNames
	// and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
	// the feed. For example: "compute.googleapis.com/Disk"
	// See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
	// supported asset types.
	AssetTypes pulumi.StringArrayInput
	// The project whose identity will be used when sending messages to the
	// destination pubsub topic. It also specifies the project for API
	// enablement check, quota, and billing.
	BillingProject pulumi.StringPtrInput
	// A condition which determines whether an asset update should be published. If specified, an asset
	// will be returned only when the expression evaluates to true. When set, expression field
	// must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
	// expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
	// condition are optional.
	// Structure is documented below.
	Condition FolderFeedConditionPtrInput
	// Asset content type. If not specified, no content but the asset name and type will be returned.
	// Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
	ContentType pulumi.StringPtrInput
	// This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
	FeedId pulumi.StringPtrInput
	// Output configuration for asset feed destination.
	// Structure is documented below.
	FeedOutputConfig FolderFeedFeedOutputConfigPtrInput
	// The folder this feed should be created in.
	Folder pulumi.StringPtrInput
	// The ID of the folder where this feed has been created. Both [FOLDER_NUMBER] and folders/[FOLDER_NUMBER] are accepted.
	FolderId pulumi.StringPtrInput
	// The format will be folders/{folder_number}/feeds/{client-assigned_feed_identifier}.
	Name pulumi.StringPtrInput
}

func (FolderFeedState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderFeedState)(nil)).Elem()
}

type folderFeedArgs struct {
	// A list of the full names of the assets to receive updates. You must specify either or both of
	// assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
	// exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
	// See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
	AssetNames []string `pulumi:"assetNames"`
	// A list of types of the assets to receive updates. You must specify either or both of assetNames
	// and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
	// the feed. For example: "compute.googleapis.com/Disk"
	// See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
	// supported asset types.
	AssetTypes []string `pulumi:"assetTypes"`
	// The project whose identity will be used when sending messages to the
	// destination pubsub topic. It also specifies the project for API
	// enablement check, quota, and billing.
	BillingProject string `pulumi:"billingProject"`
	// A condition which determines whether an asset update should be published. If specified, an asset
	// will be returned only when the expression evaluates to true. When set, expression field
	// must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
	// expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
	// condition are optional.
	// Structure is documented below.
	Condition *FolderFeedCondition `pulumi:"condition"`
	// Asset content type. If not specified, no content but the asset name and type will be returned.
	// Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
	ContentType *string `pulumi:"contentType"`
	// This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
	FeedId string `pulumi:"feedId"`
	// Output configuration for asset feed destination.
	// Structure is documented below.
	FeedOutputConfig FolderFeedFeedOutputConfig `pulumi:"feedOutputConfig"`
	// The folder this feed should be created in.
	Folder string `pulumi:"folder"`
}

// The set of arguments for constructing a FolderFeed resource.
type FolderFeedArgs struct {
	// A list of the full names of the assets to receive updates. You must specify either or both of
	// assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
	// exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
	// See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
	AssetNames pulumi.StringArrayInput
	// A list of types of the assets to receive updates. You must specify either or both of assetNames
	// and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
	// the feed. For example: "compute.googleapis.com/Disk"
	// See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
	// supported asset types.
	AssetTypes pulumi.StringArrayInput
	// The project whose identity will be used when sending messages to the
	// destination pubsub topic. It also specifies the project for API
	// enablement check, quota, and billing.
	BillingProject pulumi.StringInput
	// A condition which determines whether an asset update should be published. If specified, an asset
	// will be returned only when the expression evaluates to true. When set, expression field
	// must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
	// expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
	// condition are optional.
	// Structure is documented below.
	Condition FolderFeedConditionPtrInput
	// Asset content type. If not specified, no content but the asset name and type will be returned.
	// Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
	ContentType pulumi.StringPtrInput
	// This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
	FeedId pulumi.StringInput
	// Output configuration for asset feed destination.
	// Structure is documented below.
	FeedOutputConfig FolderFeedFeedOutputConfigInput
	// The folder this feed should be created in.
	Folder pulumi.StringInput
}

func (FolderFeedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderFeedArgs)(nil)).Elem()
}

type FolderFeedInput interface {
	pulumi.Input

	ToFolderFeedOutput() FolderFeedOutput
	ToFolderFeedOutputWithContext(ctx context.Context) FolderFeedOutput
}

func (FolderFeed) ElementType() reflect.Type {
	return reflect.TypeOf((*FolderFeed)(nil)).Elem()
}

func (i FolderFeed) ToFolderFeedOutput() FolderFeedOutput {
	return i.ToFolderFeedOutputWithContext(context.Background())
}

func (i FolderFeed) ToFolderFeedOutputWithContext(ctx context.Context) FolderFeedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderFeedOutput)
}

type FolderFeedOutput struct {
	*pulumi.OutputState
}

func (FolderFeedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FolderFeedOutput)(nil)).Elem()
}

func (o FolderFeedOutput) ToFolderFeedOutput() FolderFeedOutput {
	return o
}

func (o FolderFeedOutput) ToFolderFeedOutputWithContext(ctx context.Context) FolderFeedOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FolderFeedOutput{})
}
