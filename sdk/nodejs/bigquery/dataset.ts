// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a dataset resource for Google BigQuery. For more information see
 * [the official documentation](https://cloud.google.com/bigquery/docs/) and
 * [API](https://cloud.google.com/bigquery/docs/reference/rest/v2/datasets).
 * 
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const defaultDataset = new gcp.bigquery.Dataset("default", {
 *     accesses: [
 *         {
 *             domain: "example.com",
 *             role: "READER",
 *         },
 *         {
 *             groupByEmail: "writers@example.com",
 *             role: "WRITER",
 *         },
 *     ],
 *     datasetId: "foo",
 *     defaultTableExpirationMs: 3600000,
 *     description: "This is a test description",
 *     friendlyName: "test",
 *     labels: {
 *         env: "default",
 *     },
 *     location: "EU",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/bigquery_dataset.html.markdown.
 */
export class Dataset extends pulumi.CustomResource {
    /**
     * Get an existing Dataset resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatasetState, opts?: pulumi.CustomResourceOptions): Dataset {
        return new Dataset(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:bigquery/dataset:Dataset';

    /**
     * Returns true if the given object is an instance of Dataset.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Dataset {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Dataset.__pulumiType;
    }

    /**
     * An array of objects that define dataset access for
     * one or more entities. Structure is documented below.
     */
    public readonly accesses!: pulumi.Output<{ domain?: string, groupByEmail?: string, role?: string, specialGroup?: string, userByEmail?: string, view?: { datasetId: string, projectId: string, tableId: string } }[]>;
    /**
     * The time when this dataset was created, in milliseconds since the epoch.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<number>;
    /**
     * The ID of the dataset containing this table.
     */
    public readonly datasetId!: pulumi.Output<string>;
    /**
     * The default partition expiration
     * for all partitioned tables in the dataset, in milliseconds.
     */
    public readonly defaultPartitionExpirationMs!: pulumi.Output<number | undefined>;
    /**
     * The default lifetime of all
     * tables in the dataset, in milliseconds. The minimum value is 3600000
     * milliseconds (one hour).
     */
    public readonly defaultTableExpirationMs!: pulumi.Output<number | undefined>;
    /**
     * If set to `true`, delete all the
     * tables in the dataset when destroying the resource; otherwise, destroying
     * the resource will fail if tables are present.
     */
    public readonly deleteContentsOnDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * A user-friendly description of the dataset.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A hash of the resource.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * A descriptive name for the dataset.
     */
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    /**
     * A mapping of labels to assign to the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The date when this dataset or any of its tables was last modified,
     * in milliseconds since the epoch.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<number>;
    /**
     * The geographic location where the dataset should reside.
     * See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a Dataset resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatasetArgs | DatasetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DatasetState | undefined;
            inputs["accesses"] = state ? state.accesses : undefined;
            inputs["creationTime"] = state ? state.creationTime : undefined;
            inputs["datasetId"] = state ? state.datasetId : undefined;
            inputs["defaultPartitionExpirationMs"] = state ? state.defaultPartitionExpirationMs : undefined;
            inputs["defaultTableExpirationMs"] = state ? state.defaultTableExpirationMs : undefined;
            inputs["deleteContentsOnDestroy"] = state ? state.deleteContentsOnDestroy : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["friendlyName"] = state ? state.friendlyName : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["lastModifiedTime"] = state ? state.lastModifiedTime : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as DatasetArgs | undefined;
            if (!args || args.datasetId === undefined) {
                throw new Error("Missing required property 'datasetId'");
            }
            inputs["accesses"] = args ? args.accesses : undefined;
            inputs["datasetId"] = args ? args.datasetId : undefined;
            inputs["defaultPartitionExpirationMs"] = args ? args.defaultPartitionExpirationMs : undefined;
            inputs["defaultTableExpirationMs"] = args ? args.defaultTableExpirationMs : undefined;
            inputs["deleteContentsOnDestroy"] = args ? args.deleteContentsOnDestroy : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["friendlyName"] = args ? args.friendlyName : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Dataset.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Dataset resources.
 */
export interface DatasetState {
    /**
     * An array of objects that define dataset access for
     * one or more entities. Structure is documented below.
     */
    readonly accesses?: pulumi.Input<pulumi.Input<{ domain?: pulumi.Input<string>, groupByEmail?: pulumi.Input<string>, role?: pulumi.Input<string>, specialGroup?: pulumi.Input<string>, userByEmail?: pulumi.Input<string>, view?: pulumi.Input<{ datasetId: pulumi.Input<string>, projectId: pulumi.Input<string>, tableId: pulumi.Input<string> }> }>[]>;
    /**
     * The time when this dataset was created, in milliseconds since the epoch.
     */
    readonly creationTime?: pulumi.Input<number>;
    /**
     * The ID of the dataset containing this table.
     */
    readonly datasetId?: pulumi.Input<string>;
    /**
     * The default partition expiration
     * for all partitioned tables in the dataset, in milliseconds.
     */
    readonly defaultPartitionExpirationMs?: pulumi.Input<number>;
    /**
     * The default lifetime of all
     * tables in the dataset, in milliseconds. The minimum value is 3600000
     * milliseconds (one hour).
     */
    readonly defaultTableExpirationMs?: pulumi.Input<number>;
    /**
     * If set to `true`, delete all the
     * tables in the dataset when destroying the resource; otherwise, destroying
     * the resource will fail if tables are present.
     */
    readonly deleteContentsOnDestroy?: pulumi.Input<boolean>;
    /**
     * A user-friendly description of the dataset.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A hash of the resource.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * A descriptive name for the dataset.
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * A mapping of labels to assign to the resource.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The date when this dataset or any of its tables was last modified,
     * in milliseconds since the epoch.
     */
    readonly lastModifiedTime?: pulumi.Input<number>;
    /**
     * The geographic location where the dataset should reside.
     * See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Dataset resource.
 */
export interface DatasetArgs {
    /**
     * An array of objects that define dataset access for
     * one or more entities. Structure is documented below.
     */
    readonly accesses?: pulumi.Input<pulumi.Input<{ domain?: pulumi.Input<string>, groupByEmail?: pulumi.Input<string>, role?: pulumi.Input<string>, specialGroup?: pulumi.Input<string>, userByEmail?: pulumi.Input<string>, view?: pulumi.Input<{ datasetId: pulumi.Input<string>, projectId: pulumi.Input<string>, tableId: pulumi.Input<string> }> }>[]>;
    /**
     * The ID of the dataset containing this table.
     */
    readonly datasetId: pulumi.Input<string>;
    /**
     * The default partition expiration
     * for all partitioned tables in the dataset, in milliseconds.
     */
    readonly defaultPartitionExpirationMs?: pulumi.Input<number>;
    /**
     * The default lifetime of all
     * tables in the dataset, in milliseconds. The minimum value is 3600000
     * milliseconds (one hour).
     */
    readonly defaultTableExpirationMs?: pulumi.Input<number>;
    /**
     * If set to `true`, delete all the
     * tables in the dataset when destroying the resource; otherwise, destroying
     * the resource will fail if tables are present.
     */
    readonly deleteContentsOnDestroy?: pulumi.Input<boolean>;
    /**
     * A user-friendly description of the dataset.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A descriptive name for the dataset.
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * A mapping of labels to assign to the resource.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The geographic location where the dataset should reside.
     * See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
