// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Datasets allow you to organize and control access to your tables.
//
// To get more information about Dataset, see:
//
// * [API documentation](https://cloud.google.com/bigquery/docs/reference/rest/v2/datasets)
// * How-to Guides
//     * [Datasets Intro](https://cloud.google.com/bigquery/docs/datasets-intro)
//
// ## Example Usage
type Dataset struct {
	pulumi.CustomResourceState

	// An array of objects that define dataset access for one or more entities.  Structure is documented below.
	Accesses DatasetAccessTypeArrayOutput `pulumi:"accesses"`
	// The time when this dataset was created, in milliseconds since the epoch.
	CreationTime pulumi.IntOutput `pulumi:"creationTime"`
	// The ID of the dataset containing this table.
	DatasetId pulumi.StringOutput `pulumi:"datasetId"`
	// The default encryption key for all tables in the dataset. Once this property is set,
	// all newly-created partitioned tables in the dataset will have encryption key set to
	// this value, unless table creation request (or query) overrides the key.  Structure is documented below.
	DefaultEncryptionConfiguration DatasetDefaultEncryptionConfigurationPtrOutput `pulumi:"defaultEncryptionConfiguration"`
	// The default partition expiration for all partitioned tables in
	// the dataset, in milliseconds.
	DefaultPartitionExpirationMs pulumi.IntPtrOutput `pulumi:"defaultPartitionExpirationMs"`
	// The default lifetime of all tables in the dataset, in milliseconds.
	// The minimum value is 3600000 milliseconds (one hour).
	DefaultTableExpirationMs pulumi.IntPtrOutput `pulumi:"defaultTableExpirationMs"`
	// If set to `true`, delete all the tables in the
	// dataset when destroying the resource; otherwise,
	// destroying the resource will fail if tables are present.
	DeleteContentsOnDestroy pulumi.BoolPtrOutput `pulumi:"deleteContentsOnDestroy"`
	// A user-friendly description of the dataset
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A hash of the resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A descriptive name for the dataset
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// The labels associated with this dataset. You can use these to
	// organize and group your datasets
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The date when this dataset or any of its tables was last modified, in milliseconds since the epoch.
	LastModifiedTime pulumi.IntOutput `pulumi:"lastModifiedTime"`
	// The geographic location where the dataset should reside.
	// See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewDataset registers a new resource with the given unique name, arguments, and options.
func NewDataset(ctx *pulumi.Context,
	name string, args *DatasetArgs, opts ...pulumi.ResourceOption) (*Dataset, error) {
	if args == nil || args.DatasetId == nil {
		return nil, errors.New("missing required argument 'DatasetId'")
	}
	if args == nil {
		args = &DatasetArgs{}
	}
	var resource Dataset
	err := ctx.RegisterResource("gcp:bigquery/dataset:Dataset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataset gets an existing Dataset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetState, opts ...pulumi.ResourceOption) (*Dataset, error) {
	var resource Dataset
	err := ctx.ReadResource("gcp:bigquery/dataset:Dataset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dataset resources.
type datasetState struct {
	// An array of objects that define dataset access for one or more entities.  Structure is documented below.
	Accesses []DatasetAccessType `pulumi:"accesses"`
	// The time when this dataset was created, in milliseconds since the epoch.
	CreationTime *int `pulumi:"creationTime"`
	// The ID of the dataset containing this table.
	DatasetId *string `pulumi:"datasetId"`
	// The default encryption key for all tables in the dataset. Once this property is set,
	// all newly-created partitioned tables in the dataset will have encryption key set to
	// this value, unless table creation request (or query) overrides the key.  Structure is documented below.
	DefaultEncryptionConfiguration *DatasetDefaultEncryptionConfiguration `pulumi:"defaultEncryptionConfiguration"`
	// The default partition expiration for all partitioned tables in
	// the dataset, in milliseconds.
	DefaultPartitionExpirationMs *int `pulumi:"defaultPartitionExpirationMs"`
	// The default lifetime of all tables in the dataset, in milliseconds.
	// The minimum value is 3600000 milliseconds (one hour).
	DefaultTableExpirationMs *int `pulumi:"defaultTableExpirationMs"`
	// If set to `true`, delete all the tables in the
	// dataset when destroying the resource; otherwise,
	// destroying the resource will fail if tables are present.
	DeleteContentsOnDestroy *bool `pulumi:"deleteContentsOnDestroy"`
	// A user-friendly description of the dataset
	Description *string `pulumi:"description"`
	// A hash of the resource.
	Etag *string `pulumi:"etag"`
	// A descriptive name for the dataset
	FriendlyName *string `pulumi:"friendlyName"`
	// The labels associated with this dataset. You can use these to
	// organize and group your datasets
	Labels map[string]string `pulumi:"labels"`
	// The date when this dataset or any of its tables was last modified, in milliseconds since the epoch.
	LastModifiedTime *int `pulumi:"lastModifiedTime"`
	// The geographic location where the dataset should reside.
	// See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type DatasetState struct {
	// An array of objects that define dataset access for one or more entities.  Structure is documented below.
	Accesses DatasetAccessTypeArrayInput
	// The time when this dataset was created, in milliseconds since the epoch.
	CreationTime pulumi.IntPtrInput
	// The ID of the dataset containing this table.
	DatasetId pulumi.StringPtrInput
	// The default encryption key for all tables in the dataset. Once this property is set,
	// all newly-created partitioned tables in the dataset will have encryption key set to
	// this value, unless table creation request (or query) overrides the key.  Structure is documented below.
	DefaultEncryptionConfiguration DatasetDefaultEncryptionConfigurationPtrInput
	// The default partition expiration for all partitioned tables in
	// the dataset, in milliseconds.
	DefaultPartitionExpirationMs pulumi.IntPtrInput
	// The default lifetime of all tables in the dataset, in milliseconds.
	// The minimum value is 3600000 milliseconds (one hour).
	DefaultTableExpirationMs pulumi.IntPtrInput
	// If set to `true`, delete all the tables in the
	// dataset when destroying the resource; otherwise,
	// destroying the resource will fail if tables are present.
	DeleteContentsOnDestroy pulumi.BoolPtrInput
	// A user-friendly description of the dataset
	Description pulumi.StringPtrInput
	// A hash of the resource.
	Etag pulumi.StringPtrInput
	// A descriptive name for the dataset
	FriendlyName pulumi.StringPtrInput
	// The labels associated with this dataset. You can use these to
	// organize and group your datasets
	Labels pulumi.StringMapInput
	// The date when this dataset or any of its tables was last modified, in milliseconds since the epoch.
	LastModifiedTime pulumi.IntPtrInput
	// The geographic location where the dataset should reside.
	// See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (DatasetState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetState)(nil)).Elem()
}

type datasetArgs struct {
	// An array of objects that define dataset access for one or more entities.  Structure is documented below.
	Accesses []DatasetAccessType `pulumi:"accesses"`
	// The ID of the dataset containing this table.
	DatasetId string `pulumi:"datasetId"`
	// The default encryption key for all tables in the dataset. Once this property is set,
	// all newly-created partitioned tables in the dataset will have encryption key set to
	// this value, unless table creation request (or query) overrides the key.  Structure is documented below.
	DefaultEncryptionConfiguration *DatasetDefaultEncryptionConfiguration `pulumi:"defaultEncryptionConfiguration"`
	// The default partition expiration for all partitioned tables in
	// the dataset, in milliseconds.
	DefaultPartitionExpirationMs *int `pulumi:"defaultPartitionExpirationMs"`
	// The default lifetime of all tables in the dataset, in milliseconds.
	// The minimum value is 3600000 milliseconds (one hour).
	DefaultTableExpirationMs *int `pulumi:"defaultTableExpirationMs"`
	// If set to `true`, delete all the tables in the
	// dataset when destroying the resource; otherwise,
	// destroying the resource will fail if tables are present.
	DeleteContentsOnDestroy *bool `pulumi:"deleteContentsOnDestroy"`
	// A user-friendly description of the dataset
	Description *string `pulumi:"description"`
	// A descriptive name for the dataset
	FriendlyName *string `pulumi:"friendlyName"`
	// The labels associated with this dataset. You can use these to
	// organize and group your datasets
	Labels map[string]string `pulumi:"labels"`
	// The geographic location where the dataset should reside.
	// See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Dataset resource.
type DatasetArgs struct {
	// An array of objects that define dataset access for one or more entities.  Structure is documented below.
	Accesses DatasetAccessTypeArrayInput
	// The ID of the dataset containing this table.
	DatasetId pulumi.StringInput
	// The default encryption key for all tables in the dataset. Once this property is set,
	// all newly-created partitioned tables in the dataset will have encryption key set to
	// this value, unless table creation request (or query) overrides the key.  Structure is documented below.
	DefaultEncryptionConfiguration DatasetDefaultEncryptionConfigurationPtrInput
	// The default partition expiration for all partitioned tables in
	// the dataset, in milliseconds.
	DefaultPartitionExpirationMs pulumi.IntPtrInput
	// The default lifetime of all tables in the dataset, in milliseconds.
	// The minimum value is 3600000 milliseconds (one hour).
	DefaultTableExpirationMs pulumi.IntPtrInput
	// If set to `true`, delete all the tables in the
	// dataset when destroying the resource; otherwise,
	// destroying the resource will fail if tables are present.
	DeleteContentsOnDestroy pulumi.BoolPtrInput
	// A user-friendly description of the dataset
	Description pulumi.StringPtrInput
	// A descriptive name for the dataset
	FriendlyName pulumi.StringPtrInput
	// The labels associated with this dataset. You can use these to
	// organize and group your datasets
	Labels pulumi.StringMapInput
	// The geographic location where the dataset should reside.
	// See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetArgs)(nil)).Elem()
}
