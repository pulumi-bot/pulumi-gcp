// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Entry Metadata. A Data Catalog Entry resource represents another resource in Google Cloud Platform
// (such as a BigQuery dataset or a Pub/Sub topic) or outside of Google Cloud Platform. Clients can use
// the linkedResource field in the Entry resource to refer to the original resource ID of the source system.
//
// An Entry resource contains resource details, such as its schema. An Entry can also be used to attach
// flexible metadata, such as a Tag.
//
// To get more information about Entry, see:
//
// * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/data-catalog/docs)
//
// {{% examples %}}
// ## Example Usage
// {{% example %}}
// ### Data Catalog Entry Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/datacatalog"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		entryGroup, err := datacatalog.NewEntryGroup(ctx, "entryGroup", &datacatalog.EntryGroupArgs{
// 			EntryGroupId: pulumi.String("my_group"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datacatalog.NewEntry(ctx, "basicEntry", &datacatalog.EntryArgs{
// 			EntryGroup:          entryGroup.ID(),
// 			EntryId:             pulumi.String("my_entry"),
// 			UserSpecifiedType:   pulumi.String("my_custom_type"),
// 			UserSpecifiedSystem: pulumi.String("SomethingExternal"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// {{% /example %}}
// {{% /examples %}}
type Entry struct {
	pulumi.CustomResourceState

	// Specification for a group of BigQuery tables with name pattern [prefix]YYYYMMDD. Context:
	// https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
	BigqueryDateShardedSpec EntryBigqueryDateShardedSpecOutput `pulumi:"bigqueryDateShardedSpec"`
	// Specification that applies to a BigQuery table. This is only valid on entries of type TABLE.
	BigqueryTableSpec EntryBigqueryTableSpecOutput `pulumi:"bigqueryTableSpec"`
	// Entry description, which can consist of several sentences or paragraphs that describe entry contents.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Display information such as title and description. A short name to identify the entry,
	// for example, "Analytics Data - Jan 2011".
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The name of the entry group this entry is in.
	EntryGroup pulumi.StringOutput `pulumi:"entryGroup"`
	// The id of the entry to create.
	EntryId pulumi.StringOutput `pulumi:"entryId"`
	// Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.  Structure is documented below.
	GcsFilesetSpec EntryGcsFilesetSpecPtrOutput `pulumi:"gcsFilesetSpec"`
	// This field indicates the entry's source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.
	IntegratedSystem pulumi.StringOutput `pulumi:"integratedSystem"`
	// The resource this metadata entry refers to.
	// For Google Cloud Platform resources, linkedResource is the full name of the resource.
	// For example, the linkedResource for a table resource from BigQuery is:
	// //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
	// Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
	// this field is optional and defaults to an empty string.
	LinkedResource pulumi.StringOutput `pulumi:"linkedResource"`
	// The Data Catalog resource name of the entry in URL format. Example:
	// projects/{project_id}/locations/{location}/entryGroups/{entryGroupId}/entries/{entryId}. Note that this Entry and its
	// child resources may not actually be stored in the location in this name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
	// attached to it. See
	// https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
	// for what fields this schema can contain.
	Schema pulumi.StringPtrOutput `pulumi:"schema"`
	// The type of the entry. Only used for Entries with types in the EntryType enum.
	// Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// This field indicates the entry's source system that Data Catalog does not integrate with.
	// userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
	// and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	UserSpecifiedSystem pulumi.StringPtrOutput `pulumi:"userSpecifiedSystem"`
	// Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
	// When creating an entry, users should check the enum values first, if nothing matches the entry
	// to be created, then provide a custom value, for example "mySpecialType".
	// userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
	// numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	UserSpecifiedType pulumi.StringPtrOutput `pulumi:"userSpecifiedType"`
}

// NewEntry registers a new resource with the given unique name, arguments, and options.
func NewEntry(ctx *pulumi.Context,
	name string, args *EntryArgs, opts ...pulumi.ResourceOption) (*Entry, error) {
	if args == nil || args.EntryGroup == nil {
		return nil, errors.New("missing required argument 'EntryGroup'")
	}
	if args == nil || args.EntryId == nil {
		return nil, errors.New("missing required argument 'EntryId'")
	}
	if args == nil {
		args = &EntryArgs{}
	}
	var resource Entry
	err := ctx.RegisterResource("gcp:datacatalog/entry:Entry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntry gets an existing Entry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntryState, opts ...pulumi.ResourceOption) (*Entry, error) {
	var resource Entry
	err := ctx.ReadResource("gcp:datacatalog/entry:Entry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Entry resources.
type entryState struct {
	// Specification for a group of BigQuery tables with name pattern [prefix]YYYYMMDD. Context:
	// https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
	BigqueryDateShardedSpec *EntryBigqueryDateShardedSpec `pulumi:"bigqueryDateShardedSpec"`
	// Specification that applies to a BigQuery table. This is only valid on entries of type TABLE.
	BigqueryTableSpec *EntryBigqueryTableSpec `pulumi:"bigqueryTableSpec"`
	// Entry description, which can consist of several sentences or paragraphs that describe entry contents.
	Description *string `pulumi:"description"`
	// Display information such as title and description. A short name to identify the entry,
	// for example, "Analytics Data - Jan 2011".
	DisplayName *string `pulumi:"displayName"`
	// The name of the entry group this entry is in.
	EntryGroup *string `pulumi:"entryGroup"`
	// The id of the entry to create.
	EntryId *string `pulumi:"entryId"`
	// Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.  Structure is documented below.
	GcsFilesetSpec *EntryGcsFilesetSpec `pulumi:"gcsFilesetSpec"`
	// This field indicates the entry's source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.
	IntegratedSystem *string `pulumi:"integratedSystem"`
	// The resource this metadata entry refers to.
	// For Google Cloud Platform resources, linkedResource is the full name of the resource.
	// For example, the linkedResource for a table resource from BigQuery is:
	// //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
	// Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
	// this field is optional and defaults to an empty string.
	LinkedResource *string `pulumi:"linkedResource"`
	// The Data Catalog resource name of the entry in URL format. Example:
	// projects/{project_id}/locations/{location}/entryGroups/{entryGroupId}/entries/{entryId}. Note that this Entry and its
	// child resources may not actually be stored in the location in this name.
	Name *string `pulumi:"name"`
	// Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
	// attached to it. See
	// https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
	// for what fields this schema can contain.
	Schema *string `pulumi:"schema"`
	// The type of the entry. Only used for Entries with types in the EntryType enum.
	// Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
	Type *string `pulumi:"type"`
	// This field indicates the entry's source system that Data Catalog does not integrate with.
	// userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
	// and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	UserSpecifiedSystem *string `pulumi:"userSpecifiedSystem"`
	// Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
	// When creating an entry, users should check the enum values first, if nothing matches the entry
	// to be created, then provide a custom value, for example "mySpecialType".
	// userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
	// numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	UserSpecifiedType *string `pulumi:"userSpecifiedType"`
}

type EntryState struct {
	// Specification for a group of BigQuery tables with name pattern [prefix]YYYYMMDD. Context:
	// https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
	BigqueryDateShardedSpec EntryBigqueryDateShardedSpecPtrInput
	// Specification that applies to a BigQuery table. This is only valid on entries of type TABLE.
	BigqueryTableSpec EntryBigqueryTableSpecPtrInput
	// Entry description, which can consist of several sentences or paragraphs that describe entry contents.
	Description pulumi.StringPtrInput
	// Display information such as title and description. A short name to identify the entry,
	// for example, "Analytics Data - Jan 2011".
	DisplayName pulumi.StringPtrInput
	// The name of the entry group this entry is in.
	EntryGroup pulumi.StringPtrInput
	// The id of the entry to create.
	EntryId pulumi.StringPtrInput
	// Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.  Structure is documented below.
	GcsFilesetSpec EntryGcsFilesetSpecPtrInput
	// This field indicates the entry's source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.
	IntegratedSystem pulumi.StringPtrInput
	// The resource this metadata entry refers to.
	// For Google Cloud Platform resources, linkedResource is the full name of the resource.
	// For example, the linkedResource for a table resource from BigQuery is:
	// //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
	// Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
	// this field is optional and defaults to an empty string.
	LinkedResource pulumi.StringPtrInput
	// The Data Catalog resource name of the entry in URL format. Example:
	// projects/{project_id}/locations/{location}/entryGroups/{entryGroupId}/entries/{entryId}. Note that this Entry and its
	// child resources may not actually be stored in the location in this name.
	Name pulumi.StringPtrInput
	// Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
	// attached to it. See
	// https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
	// for what fields this schema can contain.
	Schema pulumi.StringPtrInput
	// The type of the entry. Only used for Entries with types in the EntryType enum.
	// Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
	Type pulumi.StringPtrInput
	// This field indicates the entry's source system that Data Catalog does not integrate with.
	// userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
	// and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	UserSpecifiedSystem pulumi.StringPtrInput
	// Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
	// When creating an entry, users should check the enum values first, if nothing matches the entry
	// to be created, then provide a custom value, for example "mySpecialType".
	// userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
	// numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	UserSpecifiedType pulumi.StringPtrInput
}

func (EntryState) ElementType() reflect.Type {
	return reflect.TypeOf((*entryState)(nil)).Elem()
}

type entryArgs struct {
	// Entry description, which can consist of several sentences or paragraphs that describe entry contents.
	Description *string `pulumi:"description"`
	// Display information such as title and description. A short name to identify the entry,
	// for example, "Analytics Data - Jan 2011".
	DisplayName *string `pulumi:"displayName"`
	// The name of the entry group this entry is in.
	EntryGroup string `pulumi:"entryGroup"`
	// The id of the entry to create.
	EntryId string `pulumi:"entryId"`
	// Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.  Structure is documented below.
	GcsFilesetSpec *EntryGcsFilesetSpec `pulumi:"gcsFilesetSpec"`
	// The resource this metadata entry refers to.
	// For Google Cloud Platform resources, linkedResource is the full name of the resource.
	// For example, the linkedResource for a table resource from BigQuery is:
	// //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
	// Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
	// this field is optional and defaults to an empty string.
	LinkedResource *string `pulumi:"linkedResource"`
	// Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
	// attached to it. See
	// https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
	// for what fields this schema can contain.
	Schema *string `pulumi:"schema"`
	// The type of the entry. Only used for Entries with types in the EntryType enum.
	// Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
	Type *string `pulumi:"type"`
	// This field indicates the entry's source system that Data Catalog does not integrate with.
	// userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
	// and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	UserSpecifiedSystem *string `pulumi:"userSpecifiedSystem"`
	// Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
	// When creating an entry, users should check the enum values first, if nothing matches the entry
	// to be created, then provide a custom value, for example "mySpecialType".
	// userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
	// numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	UserSpecifiedType *string `pulumi:"userSpecifiedType"`
}

// The set of arguments for constructing a Entry resource.
type EntryArgs struct {
	// Entry description, which can consist of several sentences or paragraphs that describe entry contents.
	Description pulumi.StringPtrInput
	// Display information such as title and description. A short name to identify the entry,
	// for example, "Analytics Data - Jan 2011".
	DisplayName pulumi.StringPtrInput
	// The name of the entry group this entry is in.
	EntryGroup pulumi.StringInput
	// The id of the entry to create.
	EntryId pulumi.StringInput
	// Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.  Structure is documented below.
	GcsFilesetSpec EntryGcsFilesetSpecPtrInput
	// The resource this metadata entry refers to.
	// For Google Cloud Platform resources, linkedResource is the full name of the resource.
	// For example, the linkedResource for a table resource from BigQuery is:
	// //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
	// Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
	// this field is optional and defaults to an empty string.
	LinkedResource pulumi.StringPtrInput
	// Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
	// attached to it. See
	// https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
	// for what fields this schema can contain.
	Schema pulumi.StringPtrInput
	// The type of the entry. Only used for Entries with types in the EntryType enum.
	// Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
	Type pulumi.StringPtrInput
	// This field indicates the entry's source system that Data Catalog does not integrate with.
	// userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
	// and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	UserSpecifiedSystem pulumi.StringPtrInput
	// Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
	// When creating an entry, users should check the enum values first, if nothing matches the entry
	// to be created, then provide a custom value, for example "mySpecialType".
	// userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
	// numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	UserSpecifiedType pulumi.StringPtrInput
}

func (EntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entryArgs)(nil)).Elem()
}
