// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A connection allows BigQuery connections to external data sources..
//
// To get more information about Connection, see:
//
// * [API documentation](https://cloud.google.com/bigquery/docs/reference/bigqueryconnection/rest/v1beta1/projects.locations.connections/create)
// * How-to Guides
//     * [Cloud SQL federated queries](https://cloud.google.com/bigquery/docs/cloud-sql-federated-queries)
type Connection struct {
	pulumi.CustomResourceState

	// Cloud SQL properties.  Structure is documented below.
	CloudSql ConnectionCloudSqlOutput `pulumi:"cloudSql"`
	// Optional connection id that should be assigned to the created connection.
	ConnectionId pulumi.StringPtrOutput `pulumi:"connectionId"`
	// A descriptive description for the connection
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A descriptive name for the connection
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// True if the connection has credential assigned.
	HasCredential pulumi.BoolOutput `pulumi:"hasCredential"`
	// The geographic location where the connection should reside.
	// Cloud SQL instance must be in the same location as the connection
	// with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
	// Examples: US, EU, asia-northeast1, us-central1, europe-west1. The default value is US.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource name of the connection in the form of:
	// "projects/{project_id}/locations/{location_id}/connections/{connectionId}"
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil || args.CloudSql == nil {
		return nil, errors.New("missing required argument 'CloudSql'")
	}
	if args == nil {
		args = &ConnectionArgs{}
	}
	var resource Connection
	err := ctx.RegisterResource("gcp:bigquery/connection:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("gcp:bigquery/connection:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
	// Cloud SQL properties.  Structure is documented below.
	CloudSql *ConnectionCloudSql `pulumi:"cloudSql"`
	// Optional connection id that should be assigned to the created connection.
	ConnectionId *string `pulumi:"connectionId"`
	// A descriptive description for the connection
	Description *string `pulumi:"description"`
	// A descriptive name for the connection
	FriendlyName *string `pulumi:"friendlyName"`
	// True if the connection has credential assigned.
	HasCredential *bool `pulumi:"hasCredential"`
	// The geographic location where the connection should reside.
	// Cloud SQL instance must be in the same location as the connection
	// with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
	// Examples: US, EU, asia-northeast1, us-central1, europe-west1. The default value is US.
	Location *string `pulumi:"location"`
	// The resource name of the connection in the form of:
	// "projects/{project_id}/locations/{location_id}/connections/{connectionId}"
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type ConnectionState struct {
	// Cloud SQL properties.  Structure is documented below.
	CloudSql ConnectionCloudSqlPtrInput
	// Optional connection id that should be assigned to the created connection.
	ConnectionId pulumi.StringPtrInput
	// A descriptive description for the connection
	Description pulumi.StringPtrInput
	// A descriptive name for the connection
	FriendlyName pulumi.StringPtrInput
	// True if the connection has credential assigned.
	HasCredential pulumi.BoolPtrInput
	// The geographic location where the connection should reside.
	// Cloud SQL instance must be in the same location as the connection
	// with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
	// Examples: US, EU, asia-northeast1, us-central1, europe-west1. The default value is US.
	Location pulumi.StringPtrInput
	// The resource name of the connection in the form of:
	// "projects/{project_id}/locations/{location_id}/connections/{connectionId}"
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// Cloud SQL properties.  Structure is documented below.
	CloudSql ConnectionCloudSql `pulumi:"cloudSql"`
	// Optional connection id that should be assigned to the created connection.
	ConnectionId *string `pulumi:"connectionId"`
	// A descriptive description for the connection
	Description *string `pulumi:"description"`
	// A descriptive name for the connection
	FriendlyName *string `pulumi:"friendlyName"`
	// The geographic location where the connection should reside.
	// Cloud SQL instance must be in the same location as the connection
	// with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
	// Examples: US, EU, asia-northeast1, us-central1, europe-west1. The default value is US.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// Cloud SQL properties.  Structure is documented below.
	CloudSql ConnectionCloudSqlInput
	// Optional connection id that should be assigned to the created connection.
	ConnectionId pulumi.StringPtrInput
	// A descriptive description for the connection
	Description pulumi.StringPtrInput
	// A descriptive name for the connection
	FriendlyName pulumi.StringPtrInput
	// The geographic location where the connection should reside.
	// Cloud SQL instance must be in the same location as the connection
	// with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
	// Examples: US, EU, asia-northeast1, us-central1, europe-west1. The default value is US.
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}
