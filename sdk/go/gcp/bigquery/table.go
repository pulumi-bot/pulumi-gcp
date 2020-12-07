// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a table resource in a dataset for Google BigQuery. For more information see
// [the official documentation](https://cloud.google.com/bigquery/docs/) and
// [API](https://cloud.google.com/bigquery/docs/reference/rest/v2/tables).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultDataset, err := bigquery.NewDataset(ctx, "defaultDataset", &bigquery.DatasetArgs{
// 			DatasetId:                pulumi.String("foo"),
// 			FriendlyName:             pulumi.String("test"),
// 			Description:              pulumi.String("This is a test description"),
// 			Location:                 pulumi.String("EU"),
// 			DefaultTableExpirationMs: pulumi.Int(3600000),
// 			Labels: pulumi.StringMap{
// 				"env": pulumi.String("default"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewTable(ctx, "defaultTable", &bigquery.TableArgs{
// 			DatasetId: defaultDataset.DatasetId,
// 			TableId:   pulumi.String("bar"),
// 			TimePartitioning: &bigquery.TableTimePartitioningArgs{
// 				Type: pulumi.String("DAY"),
// 			},
// 			Labels: pulumi.StringMap{
// 				"env": pulumi.String("default"),
// 			},
// 			Schema: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "[\n", "  {\n", "    \"name\": \"permalink\",\n", "    \"type\": \"STRING\",\n", "    \"mode\": \"NULLABLE\",\n", "    \"description\": \"The Permalink\"\n", "  },\n", "  {\n", "    \"name\": \"state\",\n", "    \"type\": \"STRING\",\n", "    \"mode\": \"NULLABLE\",\n", "    \"description\": \"State where the head office is located\"\n", "  }\n", "]\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewTable(ctx, "sheet", &bigquery.TableArgs{
// 			DatasetId: defaultDataset.DatasetId,
// 			TableId:   pulumi.String("sheet"),
// 			ExternalDataConfiguration: &bigquery.TableExternalDataConfigurationArgs{
// 				Autodetect:   pulumi.Bool(true),
// 				SourceFormat: pulumi.String("GOOGLE_SHEETS"),
// 				GoogleSheetsOptions: &bigquery.TableExternalDataConfigurationGoogleSheetsOptionsArgs{
// 					SkipLeadingRows: pulumi.Int(1),
// 				},
// 				SourceUris: pulumi.StringArray{
// 					pulumi.String("https://docs.google.com/spreadsheets/d/123456789012345"),
// 				},
// 			},
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
// BigQuery tables can be imported using the `project`, `dataset_id`, and `table_id`, e.g.
//
// ```sh
//  $ pulumi import gcp:bigquery/table:Table default gcp-project/foo/bar
// ```
type Table struct {
	pulumi.CustomResourceState

	// Specifies column names to use for data clustering.
	// Up to four top-level columns are allowed, and should be specified in
	// descending priority order.
	Clusterings pulumi.StringArrayOutput `pulumi:"clusterings"`
	// The time when this table was created, in milliseconds since the epoch.
	CreationTime pulumi.IntOutput `pulumi:"creationTime"`
	// The dataset ID to create the table in.
	// Changing this forces a new resource to be created.
	DatasetId pulumi.StringOutput `pulumi:"datasetId"`
	// The field description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies how the table should be encrypted.
	// If left blank, the table will be encrypted with a Google-managed key; that process
	// is transparent to the user.  Structure is documented below.
	EncryptionConfiguration TableEncryptionConfigurationPtrOutput `pulumi:"encryptionConfiguration"`
	// A hash of the resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The time when this table expires, in
	// milliseconds since the epoch. If not present, the table will persist
	// indefinitely. Expired tables will be deleted and their storage
	// reclaimed.
	ExpirationTime pulumi.IntOutput `pulumi:"expirationTime"`
	// Describes the data format,
	// location, and other properties of a table stored outside of BigQuery.
	// By defining these properties, the data source can then be queried as
	// if it were a standard BigQuery table. Structure is documented below.
	ExternalDataConfiguration TableExternalDataConfigurationPtrOutput `pulumi:"externalDataConfiguration"`
	// A descriptive name for the table.
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// A mapping of labels to assign to the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The time when this table was last modified, in milliseconds since the epoch.
	LastModifiedTime pulumi.IntOutput `pulumi:"lastModifiedTime"`
	// The geographic location where the table resides. This value is inherited from the dataset.
	Location pulumi.StringOutput `pulumi:"location"`
	// If specified, configures this table as a materialized view.
	// Structure is documented below.
	MaterializedView TableMaterializedViewPtrOutput `pulumi:"materializedView"`
	// The size of this table in bytes, excluding any data in the streaming buffer.
	NumBytes pulumi.IntOutput `pulumi:"numBytes"`
	// The number of bytes in the table that are considered "long-term storage".
	NumLongTermBytes pulumi.IntOutput `pulumi:"numLongTermBytes"`
	// The number of rows of data in this table, excluding any data in the streaming buffer.
	NumRows pulumi.IntOutput `pulumi:"numRows"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// If specified, configures range-based
	// partitioning for this table. Structure is documented below.
	RangePartitioning TableRangePartitioningPtrOutput `pulumi:"rangePartitioning"`
	// A JSON schema for the external table. Schema is required
	// for CSV and JSON formats if autodetect is not on. Schema is disallowed
	// for Google Cloud Bigtable, Cloud Datastore backups, Avro, ORC and Parquet formats.
	// ~>**NOTE:** Because this field expects a JSON string, any changes to the
	// string will create a diff, even if the JSON itself hasn't changed.
	// Furthermore drift for this field cannot not be detected because BigQuery
	// only uses this schema to compute the effective schema for the table, therefore
	// any changes on the configured value will force the table to be recreated.
	// This schema is effectively only applied when creating a table from an external
	// datasource, after creation the computed schema will be stored in
	// `google_bigquery_table.schema`
	Schema pulumi.StringOutput `pulumi:"schema"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// A unique ID for the resource.
	// Changing this forces a new resource to be created.
	TableId pulumi.StringOutput `pulumi:"tableId"`
	// If specified, configures time-based
	// partitioning for this table. Structure is documented below.
	TimePartitioning TableTimePartitioningPtrOutput `pulumi:"timePartitioning"`
	// The supported types are DAY, HOUR, MONTH, and YEAR,
	// which will generate one partition per day, hour, month, and year, respectively.
	Type pulumi.StringOutput `pulumi:"type"`
	// If specified, configures this table as a view.
	// Structure is documented below.
	View TableViewPtrOutput `pulumi:"view"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	if args.TableId == nil {
		return nil, errors.New("invalid value for required argument 'TableId'")
	}
	var resource Table
	err := ctx.RegisterResource("gcp:bigquery/table:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTable gets an existing Table resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("gcp:bigquery/table:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type tableState struct {
	// Specifies column names to use for data clustering.
	// Up to four top-level columns are allowed, and should be specified in
	// descending priority order.
	Clusterings []string `pulumi:"clusterings"`
	// The time when this table was created, in milliseconds since the epoch.
	CreationTime *int `pulumi:"creationTime"`
	// The dataset ID to create the table in.
	// Changing this forces a new resource to be created.
	DatasetId *string `pulumi:"datasetId"`
	// The field description.
	Description *string `pulumi:"description"`
	// Specifies how the table should be encrypted.
	// If left blank, the table will be encrypted with a Google-managed key; that process
	// is transparent to the user.  Structure is documented below.
	EncryptionConfiguration *TableEncryptionConfiguration `pulumi:"encryptionConfiguration"`
	// A hash of the resource.
	Etag *string `pulumi:"etag"`
	// The time when this table expires, in
	// milliseconds since the epoch. If not present, the table will persist
	// indefinitely. Expired tables will be deleted and their storage
	// reclaimed.
	ExpirationTime *int `pulumi:"expirationTime"`
	// Describes the data format,
	// location, and other properties of a table stored outside of BigQuery.
	// By defining these properties, the data source can then be queried as
	// if it were a standard BigQuery table. Structure is documented below.
	ExternalDataConfiguration *TableExternalDataConfiguration `pulumi:"externalDataConfiguration"`
	// A descriptive name for the table.
	FriendlyName *string `pulumi:"friendlyName"`
	// A mapping of labels to assign to the resource.
	Labels map[string]string `pulumi:"labels"`
	// The time when this table was last modified, in milliseconds since the epoch.
	LastModifiedTime *int `pulumi:"lastModifiedTime"`
	// The geographic location where the table resides. This value is inherited from the dataset.
	Location *string `pulumi:"location"`
	// If specified, configures this table as a materialized view.
	// Structure is documented below.
	MaterializedView *TableMaterializedView `pulumi:"materializedView"`
	// The size of this table in bytes, excluding any data in the streaming buffer.
	NumBytes *int `pulumi:"numBytes"`
	// The number of bytes in the table that are considered "long-term storage".
	NumLongTermBytes *int `pulumi:"numLongTermBytes"`
	// The number of rows of data in this table, excluding any data in the streaming buffer.
	NumRows *int `pulumi:"numRows"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// If specified, configures range-based
	// partitioning for this table. Structure is documented below.
	RangePartitioning *TableRangePartitioning `pulumi:"rangePartitioning"`
	// A JSON schema for the external table. Schema is required
	// for CSV and JSON formats if autodetect is not on. Schema is disallowed
	// for Google Cloud Bigtable, Cloud Datastore backups, Avro, ORC and Parquet formats.
	// ~>**NOTE:** Because this field expects a JSON string, any changes to the
	// string will create a diff, even if the JSON itself hasn't changed.
	// Furthermore drift for this field cannot not be detected because BigQuery
	// only uses this schema to compute the effective schema for the table, therefore
	// any changes on the configured value will force the table to be recreated.
	// This schema is effectively only applied when creating a table from an external
	// datasource, after creation the computed schema will be stored in
	// `google_bigquery_table.schema`
	Schema *string `pulumi:"schema"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// A unique ID for the resource.
	// Changing this forces a new resource to be created.
	TableId *string `pulumi:"tableId"`
	// If specified, configures time-based
	// partitioning for this table. Structure is documented below.
	TimePartitioning *TableTimePartitioning `pulumi:"timePartitioning"`
	// The supported types are DAY, HOUR, MONTH, and YEAR,
	// which will generate one partition per day, hour, month, and year, respectively.
	Type *string `pulumi:"type"`
	// If specified, configures this table as a view.
	// Structure is documented below.
	View *TableView `pulumi:"view"`
}

type TableState struct {
	// Specifies column names to use for data clustering.
	// Up to four top-level columns are allowed, and should be specified in
	// descending priority order.
	Clusterings pulumi.StringArrayInput
	// The time when this table was created, in milliseconds since the epoch.
	CreationTime pulumi.IntPtrInput
	// The dataset ID to create the table in.
	// Changing this forces a new resource to be created.
	DatasetId pulumi.StringPtrInput
	// The field description.
	Description pulumi.StringPtrInput
	// Specifies how the table should be encrypted.
	// If left blank, the table will be encrypted with a Google-managed key; that process
	// is transparent to the user.  Structure is documented below.
	EncryptionConfiguration TableEncryptionConfigurationPtrInput
	// A hash of the resource.
	Etag pulumi.StringPtrInput
	// The time when this table expires, in
	// milliseconds since the epoch. If not present, the table will persist
	// indefinitely. Expired tables will be deleted and their storage
	// reclaimed.
	ExpirationTime pulumi.IntPtrInput
	// Describes the data format,
	// location, and other properties of a table stored outside of BigQuery.
	// By defining these properties, the data source can then be queried as
	// if it were a standard BigQuery table. Structure is documented below.
	ExternalDataConfiguration TableExternalDataConfigurationPtrInput
	// A descriptive name for the table.
	FriendlyName pulumi.StringPtrInput
	// A mapping of labels to assign to the resource.
	Labels pulumi.StringMapInput
	// The time when this table was last modified, in milliseconds since the epoch.
	LastModifiedTime pulumi.IntPtrInput
	// The geographic location where the table resides. This value is inherited from the dataset.
	Location pulumi.StringPtrInput
	// If specified, configures this table as a materialized view.
	// Structure is documented below.
	MaterializedView TableMaterializedViewPtrInput
	// The size of this table in bytes, excluding any data in the streaming buffer.
	NumBytes pulumi.IntPtrInput
	// The number of bytes in the table that are considered "long-term storage".
	NumLongTermBytes pulumi.IntPtrInput
	// The number of rows of data in this table, excluding any data in the streaming buffer.
	NumRows pulumi.IntPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// If specified, configures range-based
	// partitioning for this table. Structure is documented below.
	RangePartitioning TableRangePartitioningPtrInput
	// A JSON schema for the external table. Schema is required
	// for CSV and JSON formats if autodetect is not on. Schema is disallowed
	// for Google Cloud Bigtable, Cloud Datastore backups, Avro, ORC and Parquet formats.
	// ~>**NOTE:** Because this field expects a JSON string, any changes to the
	// string will create a diff, even if the JSON itself hasn't changed.
	// Furthermore drift for this field cannot not be detected because BigQuery
	// only uses this schema to compute the effective schema for the table, therefore
	// any changes on the configured value will force the table to be recreated.
	// This schema is effectively only applied when creating a table from an external
	// datasource, after creation the computed schema will be stored in
	// `google_bigquery_table.schema`
	Schema pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// A unique ID for the resource.
	// Changing this forces a new resource to be created.
	TableId pulumi.StringPtrInput
	// If specified, configures time-based
	// partitioning for this table. Structure is documented below.
	TimePartitioning TableTimePartitioningPtrInput
	// The supported types are DAY, HOUR, MONTH, and YEAR,
	// which will generate one partition per day, hour, month, and year, respectively.
	Type pulumi.StringPtrInput
	// If specified, configures this table as a view.
	// Structure is documented below.
	View TableViewPtrInput
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	// Specifies column names to use for data clustering.
	// Up to four top-level columns are allowed, and should be specified in
	// descending priority order.
	Clusterings []string `pulumi:"clusterings"`
	// The dataset ID to create the table in.
	// Changing this forces a new resource to be created.
	DatasetId string `pulumi:"datasetId"`
	// The field description.
	Description *string `pulumi:"description"`
	// Specifies how the table should be encrypted.
	// If left blank, the table will be encrypted with a Google-managed key; that process
	// is transparent to the user.  Structure is documented below.
	EncryptionConfiguration *TableEncryptionConfiguration `pulumi:"encryptionConfiguration"`
	// The time when this table expires, in
	// milliseconds since the epoch. If not present, the table will persist
	// indefinitely. Expired tables will be deleted and their storage
	// reclaimed.
	ExpirationTime *int `pulumi:"expirationTime"`
	// Describes the data format,
	// location, and other properties of a table stored outside of BigQuery.
	// By defining these properties, the data source can then be queried as
	// if it were a standard BigQuery table. Structure is documented below.
	ExternalDataConfiguration *TableExternalDataConfiguration `pulumi:"externalDataConfiguration"`
	// A descriptive name for the table.
	FriendlyName *string `pulumi:"friendlyName"`
	// A mapping of labels to assign to the resource.
	Labels map[string]string `pulumi:"labels"`
	// If specified, configures this table as a materialized view.
	// Structure is documented below.
	MaterializedView *TableMaterializedView `pulumi:"materializedView"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// If specified, configures range-based
	// partitioning for this table. Structure is documented below.
	RangePartitioning *TableRangePartitioning `pulumi:"rangePartitioning"`
	// A JSON schema for the external table. Schema is required
	// for CSV and JSON formats if autodetect is not on. Schema is disallowed
	// for Google Cloud Bigtable, Cloud Datastore backups, Avro, ORC and Parquet formats.
	// ~>**NOTE:** Because this field expects a JSON string, any changes to the
	// string will create a diff, even if the JSON itself hasn't changed.
	// Furthermore drift for this field cannot not be detected because BigQuery
	// only uses this schema to compute the effective schema for the table, therefore
	// any changes on the configured value will force the table to be recreated.
	// This schema is effectively only applied when creating a table from an external
	// datasource, after creation the computed schema will be stored in
	// `google_bigquery_table.schema`
	Schema *string `pulumi:"schema"`
	// A unique ID for the resource.
	// Changing this forces a new resource to be created.
	TableId string `pulumi:"tableId"`
	// If specified, configures time-based
	// partitioning for this table. Structure is documented below.
	TimePartitioning *TableTimePartitioning `pulumi:"timePartitioning"`
	// If specified, configures this table as a view.
	// Structure is documented below.
	View *TableView `pulumi:"view"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	// Specifies column names to use for data clustering.
	// Up to four top-level columns are allowed, and should be specified in
	// descending priority order.
	Clusterings pulumi.StringArrayInput
	// The dataset ID to create the table in.
	// Changing this forces a new resource to be created.
	DatasetId pulumi.StringInput
	// The field description.
	Description pulumi.StringPtrInput
	// Specifies how the table should be encrypted.
	// If left blank, the table will be encrypted with a Google-managed key; that process
	// is transparent to the user.  Structure is documented below.
	EncryptionConfiguration TableEncryptionConfigurationPtrInput
	// The time when this table expires, in
	// milliseconds since the epoch. If not present, the table will persist
	// indefinitely. Expired tables will be deleted and their storage
	// reclaimed.
	ExpirationTime pulumi.IntPtrInput
	// Describes the data format,
	// location, and other properties of a table stored outside of BigQuery.
	// By defining these properties, the data source can then be queried as
	// if it were a standard BigQuery table. Structure is documented below.
	ExternalDataConfiguration TableExternalDataConfigurationPtrInput
	// A descriptive name for the table.
	FriendlyName pulumi.StringPtrInput
	// A mapping of labels to assign to the resource.
	Labels pulumi.StringMapInput
	// If specified, configures this table as a materialized view.
	// Structure is documented below.
	MaterializedView TableMaterializedViewPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// If specified, configures range-based
	// partitioning for this table. Structure is documented below.
	RangePartitioning TableRangePartitioningPtrInput
	// A JSON schema for the external table. Schema is required
	// for CSV and JSON formats if autodetect is not on. Schema is disallowed
	// for Google Cloud Bigtable, Cloud Datastore backups, Avro, ORC and Parquet formats.
	// ~>**NOTE:** Because this field expects a JSON string, any changes to the
	// string will create a diff, even if the JSON itself hasn't changed.
	// Furthermore drift for this field cannot not be detected because BigQuery
	// only uses this schema to compute the effective schema for the table, therefore
	// any changes on the configured value will force the table to be recreated.
	// This schema is effectively only applied when creating a table from an external
	// datasource, after creation the computed schema will be stored in
	// `google_bigquery_table.schema`
	Schema pulumi.StringPtrInput
	// A unique ID for the resource.
	// Changing this forces a new resource to be created.
	TableId pulumi.StringInput
	// If specified, configures time-based
	// partitioning for this table. Structure is documented below.
	TimePartitioning TableTimePartitioningPtrInput
	// If specified, configures this table as a view.
	// Structure is documented below.
	View TableViewPtrInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}

type TableInput interface {
	pulumi.Input

	ToTableOutput() TableOutput
	ToTableOutputWithContext(ctx context.Context) TableOutput
}

func (Table) ElementType() reflect.Type {
	return reflect.TypeOf((*Table)(nil)).Elem()
}

func (i Table) ToTableOutput() TableOutput {
	return i.ToTableOutputWithContext(context.Background())
}

func (i Table) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableOutput)
}

type TableOutput struct {
	*pulumi.OutputState
}

func (TableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableOutput)(nil)).Elem()
}

func (o TableOutput) ToTableOutput() TableOutput {
	return o
}

func (o TableOutput) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TableOutput{})
}
