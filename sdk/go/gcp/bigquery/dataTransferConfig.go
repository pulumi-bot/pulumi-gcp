// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package bigquery

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/bigquery_data_transfer_config.html.markdown.
type DataTransferConfig struct {
	pulumi.CustomResourceState

	// The number of days to look back to automatically refresh the data. For example, if dataRefreshWindowDays = 10, then
	// every day BigQuery reingests data for [today-10, today-1], rather than ingesting data for just [today-1]. Only valid if
	// the data source supports the feature. Set the value to 0 to use the default value.
	DataRefreshWindowDays pulumi.IntPtrOutput `pulumi:"dataRefreshWindowDays"`
	// The data source id. Cannot be changed once the transfer config is created.
	DataSourceId pulumi.StringOutput `pulumi:"dataSourceId"`
	// The BigQuery target dataset id.
	DestinationDatasetId pulumi.StringOutput `pulumi:"destinationDatasetId"`
	// When set to true, no runs are scheduled for a given transfer.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// The user specified display name for the transfer config.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is
	// US.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource name of the transfer config. Transfer config names have the form
	// projects/{projectId}/locations/{location}/transferConfigs/{configId}. Where configId is usually a uuid, but this is not
	// required. The name is ignored when creating a transfer config.
	Name pulumi.StringOutput `pulumi:"name"`
	// These parameters are specific to each data source.
	Params pulumi.StringMapOutput `pulumi:"params"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Data transfer schedule. If the data source does not support a custom schedule, this should be empty. If it is empty, the
	// default value for the data source will be used. The specified times are in UTC. Examples of valid format: 1st,3rd monday
	// of month 15:30, every wed,fri of jan, jun 13:15, and first sunday of quarter 00:00. See more explanation about the
	// format here: https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
	// NOTE: the granularity should be at least 8 hours, or less frequent.
	Schedule pulumi.StringPtrOutput `pulumi:"schedule"`
}

// NewDataTransferConfig registers a new resource with the given unique name, arguments, and options.
func NewDataTransferConfig(ctx *pulumi.Context,
	name string, args *DataTransferConfigArgs, opts ...pulumi.ResourceOption) (*DataTransferConfig, error) {
	if args == nil || args.DataSourceId == nil {
		return nil, errors.New("missing required argument 'DataSourceId'")
	}
	if args == nil || args.DestinationDatasetId == nil {
		return nil, errors.New("missing required argument 'DestinationDatasetId'")
	}
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.Params == nil {
		return nil, errors.New("missing required argument 'Params'")
	}
	if args == nil {
		args = &DataTransferConfigArgs{}
	}
	var resource DataTransferConfig
	err := ctx.RegisterResource("gcp:bigquery/dataTransferConfig:DataTransferConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataTransferConfig gets an existing DataTransferConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataTransferConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataTransferConfigState, opts ...pulumi.ResourceOption) (*DataTransferConfig, error) {
	var resource DataTransferConfig
	err := ctx.ReadResource("gcp:bigquery/dataTransferConfig:DataTransferConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataTransferConfig resources.
type dataTransferConfigState struct {
	// The number of days to look back to automatically refresh the data. For example, if dataRefreshWindowDays = 10, then
	// every day BigQuery reingests data for [today-10, today-1], rather than ingesting data for just [today-1]. Only valid if
	// the data source supports the feature. Set the value to 0 to use the default value.
	DataRefreshWindowDays *int `pulumi:"dataRefreshWindowDays"`
	// The data source id. Cannot be changed once the transfer config is created.
	DataSourceId *string `pulumi:"dataSourceId"`
	// The BigQuery target dataset id.
	DestinationDatasetId *string `pulumi:"destinationDatasetId"`
	// When set to true, no runs are scheduled for a given transfer.
	Disabled *bool `pulumi:"disabled"`
	// The user specified display name for the transfer config.
	DisplayName *string `pulumi:"displayName"`
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is
	// US.
	Location *string `pulumi:"location"`
	// The resource name of the transfer config. Transfer config names have the form
	// projects/{projectId}/locations/{location}/transferConfigs/{configId}. Where configId is usually a uuid, but this is not
	// required. The name is ignored when creating a transfer config.
	Name *string `pulumi:"name"`
	// These parameters are specific to each data source.
	Params map[string]string `pulumi:"params"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Data transfer schedule. If the data source does not support a custom schedule, this should be empty. If it is empty, the
	// default value for the data source will be used. The specified times are in UTC. Examples of valid format: 1st,3rd monday
	// of month 15:30, every wed,fri of jan, jun 13:15, and first sunday of quarter 00:00. See more explanation about the
	// format here: https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
	// NOTE: the granularity should be at least 8 hours, or less frequent.
	Schedule *string `pulumi:"schedule"`
}

type DataTransferConfigState struct {
	// The number of days to look back to automatically refresh the data. For example, if dataRefreshWindowDays = 10, then
	// every day BigQuery reingests data for [today-10, today-1], rather than ingesting data for just [today-1]. Only valid if
	// the data source supports the feature. Set the value to 0 to use the default value.
	DataRefreshWindowDays pulumi.IntPtrInput
	// The data source id. Cannot be changed once the transfer config is created.
	DataSourceId pulumi.StringPtrInput
	// The BigQuery target dataset id.
	DestinationDatasetId pulumi.StringPtrInput
	// When set to true, no runs are scheduled for a given transfer.
	Disabled pulumi.BoolPtrInput
	// The user specified display name for the transfer config.
	DisplayName pulumi.StringPtrInput
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is
	// US.
	Location pulumi.StringPtrInput
	// The resource name of the transfer config. Transfer config names have the form
	// projects/{projectId}/locations/{location}/transferConfigs/{configId}. Where configId is usually a uuid, but this is not
	// required. The name is ignored when creating a transfer config.
	Name pulumi.StringPtrInput
	// These parameters are specific to each data source.
	Params pulumi.StringMapInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Data transfer schedule. If the data source does not support a custom schedule, this should be empty. If it is empty, the
	// default value for the data source will be used. The specified times are in UTC. Examples of valid format: 1st,3rd monday
	// of month 15:30, every wed,fri of jan, jun 13:15, and first sunday of quarter 00:00. See more explanation about the
	// format here: https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
	// NOTE: the granularity should be at least 8 hours, or less frequent.
	Schedule pulumi.StringPtrInput
}

func (DataTransferConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataTransferConfigState)(nil)).Elem()
}

type dataTransferConfigArgs struct {
	// The number of days to look back to automatically refresh the data. For example, if dataRefreshWindowDays = 10, then
	// every day BigQuery reingests data for [today-10, today-1], rather than ingesting data for just [today-1]. Only valid if
	// the data source supports the feature. Set the value to 0 to use the default value.
	DataRefreshWindowDays *int `pulumi:"dataRefreshWindowDays"`
	// The data source id. Cannot be changed once the transfer config is created.
	DataSourceId string `pulumi:"dataSourceId"`
	// The BigQuery target dataset id.
	DestinationDatasetId string `pulumi:"destinationDatasetId"`
	// When set to true, no runs are scheduled for a given transfer.
	Disabled *bool `pulumi:"disabled"`
	// The user specified display name for the transfer config.
	DisplayName string `pulumi:"displayName"`
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is
	// US.
	Location *string `pulumi:"location"`
	// These parameters are specific to each data source.
	Params map[string]string `pulumi:"params"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Data transfer schedule. If the data source does not support a custom schedule, this should be empty. If it is empty, the
	// default value for the data source will be used. The specified times are in UTC. Examples of valid format: 1st,3rd monday
	// of month 15:30, every wed,fri of jan, jun 13:15, and first sunday of quarter 00:00. See more explanation about the
	// format here: https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
	// NOTE: the granularity should be at least 8 hours, or less frequent.
	Schedule *string `pulumi:"schedule"`
}

// The set of arguments for constructing a DataTransferConfig resource.
type DataTransferConfigArgs struct {
	// The number of days to look back to automatically refresh the data. For example, if dataRefreshWindowDays = 10, then
	// every day BigQuery reingests data for [today-10, today-1], rather than ingesting data for just [today-1]. Only valid if
	// the data source supports the feature. Set the value to 0 to use the default value.
	DataRefreshWindowDays pulumi.IntPtrInput
	// The data source id. Cannot be changed once the transfer config is created.
	DataSourceId pulumi.StringInput
	// The BigQuery target dataset id.
	DestinationDatasetId pulumi.StringInput
	// When set to true, no runs are scheduled for a given transfer.
	Disabled pulumi.BoolPtrInput
	// The user specified display name for the transfer config.
	DisplayName pulumi.StringInput
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is
	// US.
	Location pulumi.StringPtrInput
	// These parameters are specific to each data source.
	Params pulumi.StringMapInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Data transfer schedule. If the data source does not support a custom schedule, this should be empty. If it is empty, the
	// default value for the data source will be used. The specified times are in UTC. Examples of valid format: 1st,3rd monday
	// of month 15:30, every wed,fri of jan, jun 13:15, and first sunday of quarter 00:00. See more explanation about the
	// format here: https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
	// NOTE: the granularity should be at least 8 hours, or less frequent.
	Schedule pulumi.StringPtrInput
}

func (DataTransferConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataTransferConfigArgs)(nil)).Elem()
}

