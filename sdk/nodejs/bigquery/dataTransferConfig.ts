// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Represents a data transfer configuration. A transfer configuration
 * contains all metadata needed to perform a data transfer.
 *
 *
 * To get more information about Config, see:
 *
 * * [API documentation](https://cloud.google.com/bigquery/docs/reference/datatransfer/rest/v1/projects.locations.transferConfigs/create)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/bigquery/docs/reference/datatransfer/rest/)
 *
 * ## Example Usage - Scheduled Query
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = gcp.organizations.getProject({});
 * const permissions = new gcp.projects.IAMMember("permissions", {
 *     role: "roles/iam.serviceAccountShortTermTokenMinter",
 *     member: project.then(project => `serviceAccount:service-${project.number}@gcp-sa-bigquerydatatransfer.iam.gserviceaccount.com`),
 * });
 * const myDataset = new gcp.bigquery.Dataset("myDataset", {
 *     datasetId: "my_dataset",
 *     friendlyName: "foo",
 *     description: "bar",
 *     location: "asia-northeast1",
 * });
 * const queryConfig = new gcp.bigquery.DataTransferConfig("queryConfig", {
 *     displayName: "my-query",
 *     location: "asia-northeast1",
 *     dataSourceId: "scheduled_query",
 *     schedule: "first sunday of quarter 00:00",
 *     destinationDatasetId: myDataset.datasetId,
 *     params: {
 *         destination_table_name_template: "my-table",
 *         write_disposition: "WRITE_APPEND",
 *         query: "SELECT name FROM tabl WHERE x = 'y'",
 *     },
 * });
 * ```
 */
export class DataTransferConfig extends pulumi.CustomResource {
    /**
     * Get an existing DataTransferConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataTransferConfigState, opts?: pulumi.CustomResourceOptions): DataTransferConfig {
        return new DataTransferConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:bigquery/dataTransferConfig:DataTransferConfig';

    /**
     * Returns true if the given object is an instance of DataTransferConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataTransferConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataTransferConfig.__pulumiType;
    }

    /**
     * The number of days to look back to automatically refresh the data.
     * For example, if dataRefreshWindowDays = 10, then every day BigQuery
     * reingests data for [today-10, today-1], rather than ingesting data for
     * just [today-1]. Only valid if the data source supports the feature.
     * Set the value to 0 to use the default value.
     */
    public readonly dataRefreshWindowDays!: pulumi.Output<number | undefined>;
    /**
     * The data source id. Cannot be changed once the transfer config is created.
     */
    public readonly dataSourceId!: pulumi.Output<string>;
    /**
     * The BigQuery target dataset id.
     */
    public readonly destinationDatasetId!: pulumi.Output<string>;
    /**
     * When set to true, no runs are scheduled for a given transfer.
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * The user specified display name for the transfer config.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The geographic location where the transfer config should reside.
     * Examples: US, EU, asia-northeast1. The default value is US.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The resource name of the transfer config. Transfer config names have the form
     * projects/{projectId}/locations/{location}/transferConfigs/{configId}. Where configId is usually a uuid, but this is
     * not required. The name is ignored when creating a transfer config.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * These parameters are specific to each data source.
     */
    public readonly params!: pulumi.Output<{[key: string]: string}>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Data transfer schedule. If the data source does not support a custom
     * schedule, this should be empty. If it is empty, the default value for
     * the data source will be used. The specified times are in UTC. Examples
     * of valid format: 1st,3rd monday of month 15:30, every wed,fri of jan,
     * jun 13:15, and first sunday of quarter 00:00. See more explanation
     * about the format here:
     * https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
     * NOTE: the granularity should be at least 8 hours, or less frequent.
     */
    public readonly schedule!: pulumi.Output<string | undefined>;
    /**
     * Optional service account name. If this field is set, transfer config will
     * be created with this service account credentials. It requires that
     * requesting user calling this API has permissions to act as this service account.
     */
    public readonly serviceAccountName!: pulumi.Output<string | undefined>;

    /**
     * Create a DataTransferConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataTransferConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataTransferConfigArgs | DataTransferConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DataTransferConfigState | undefined;
            inputs["dataRefreshWindowDays"] = state ? state.dataRefreshWindowDays : undefined;
            inputs["dataSourceId"] = state ? state.dataSourceId : undefined;
            inputs["destinationDatasetId"] = state ? state.destinationDatasetId : undefined;
            inputs["disabled"] = state ? state.disabled : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["params"] = state ? state.params : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["schedule"] = state ? state.schedule : undefined;
            inputs["serviceAccountName"] = state ? state.serviceAccountName : undefined;
        } else {
            const args = argsOrState as DataTransferConfigArgs | undefined;
            if (!args || args.dataSourceId === undefined) {
                throw new Error("Missing required property 'dataSourceId'");
            }
            if (!args || args.destinationDatasetId === undefined) {
                throw new Error("Missing required property 'destinationDatasetId'");
            }
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.params === undefined) {
                throw new Error("Missing required property 'params'");
            }
            inputs["dataRefreshWindowDays"] = args ? args.dataRefreshWindowDays : undefined;
            inputs["dataSourceId"] = args ? args.dataSourceId : undefined;
            inputs["destinationDatasetId"] = args ? args.destinationDatasetId : undefined;
            inputs["disabled"] = args ? args.disabled : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["params"] = args ? args.params : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["serviceAccountName"] = args ? args.serviceAccountName : undefined;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DataTransferConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataTransferConfig resources.
 */
export interface DataTransferConfigState {
    /**
     * The number of days to look back to automatically refresh the data.
     * For example, if dataRefreshWindowDays = 10, then every day BigQuery
     * reingests data for [today-10, today-1], rather than ingesting data for
     * just [today-1]. Only valid if the data source supports the feature.
     * Set the value to 0 to use the default value.
     */
    readonly dataRefreshWindowDays?: pulumi.Input<number>;
    /**
     * The data source id. Cannot be changed once the transfer config is created.
     */
    readonly dataSourceId?: pulumi.Input<string>;
    /**
     * The BigQuery target dataset id.
     */
    readonly destinationDatasetId?: pulumi.Input<string>;
    /**
     * When set to true, no runs are scheduled for a given transfer.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * The user specified display name for the transfer config.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The geographic location where the transfer config should reside.
     * Examples: US, EU, asia-northeast1. The default value is US.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The resource name of the transfer config. Transfer config names have the form
     * projects/{projectId}/locations/{location}/transferConfigs/{configId}. Where configId is usually a uuid, but this is
     * not required. The name is ignored when creating a transfer config.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * These parameters are specific to each data source.
     */
    readonly params?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Data transfer schedule. If the data source does not support a custom
     * schedule, this should be empty. If it is empty, the default value for
     * the data source will be used. The specified times are in UTC. Examples
     * of valid format: 1st,3rd monday of month 15:30, every wed,fri of jan,
     * jun 13:15, and first sunday of quarter 00:00. See more explanation
     * about the format here:
     * https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
     * NOTE: the granularity should be at least 8 hours, or less frequent.
     */
    readonly schedule?: pulumi.Input<string>;
    /**
     * Optional service account name. If this field is set, transfer config will
     * be created with this service account credentials. It requires that
     * requesting user calling this API has permissions to act as this service account.
     */
    readonly serviceAccountName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DataTransferConfig resource.
 */
export interface DataTransferConfigArgs {
    /**
     * The number of days to look back to automatically refresh the data.
     * For example, if dataRefreshWindowDays = 10, then every day BigQuery
     * reingests data for [today-10, today-1], rather than ingesting data for
     * just [today-1]. Only valid if the data source supports the feature.
     * Set the value to 0 to use the default value.
     */
    readonly dataRefreshWindowDays?: pulumi.Input<number>;
    /**
     * The data source id. Cannot be changed once the transfer config is created.
     */
    readonly dataSourceId: pulumi.Input<string>;
    /**
     * The BigQuery target dataset id.
     */
    readonly destinationDatasetId: pulumi.Input<string>;
    /**
     * When set to true, no runs are scheduled for a given transfer.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * The user specified display name for the transfer config.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * The geographic location where the transfer config should reside.
     * Examples: US, EU, asia-northeast1. The default value is US.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * These parameters are specific to each data source.
     */
    readonly params: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Data transfer schedule. If the data source does not support a custom
     * schedule, this should be empty. If it is empty, the default value for
     * the data source will be used. The specified times are in UTC. Examples
     * of valid format: 1st,3rd monday of month 15:30, every wed,fri of jan,
     * jun 13:15, and first sunday of quarter 00:00. See more explanation
     * about the format here:
     * https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
     * NOTE: the granularity should be at least 8 hours, or less frequent.
     */
    readonly schedule?: pulumi.Input<string>;
    /**
     * Optional service account name. If this field is set, transfer config will
     * be created with this service account credentials. It requires that
     * requesting user calling this API has permissions to act as this service account.
     */
    readonly serviceAccountName?: pulumi.Input<string>;
}
