// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Describes a composite index for Cloud Datastore.
 *
 * To get more information about Index, see:
 *
 * * [API documentation](https://cloud.google.com/datastore/docs/reference/admin/rest/v1/projects.indexes)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/datastore/docs/concepts/indexes)
 *
 * > **Warning:** This resource creates a Datastore Index on a project that has already
 * enabled a Datastore-compatible database. If you haven't already enabled
 * one, you can create a `gcp.appengine.Application` resource with
 * `databaseType` set to `"CLOUD_DATASTORE_COMPATIBILITY"` to do so. Your
 * Datastore location will be the same as the App Engine location specified.
 *
 * ## Example Usage
 * ### Datastore Index
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultDataStoreIndex = new gcp.datastore.DataStoreIndex("default", {
 *     kind: "foo",
 *     properties: [
 *         {
 *             direction: "ASCENDING",
 *             name: "property_a",
 *         },
 *         {
 *             direction: "ASCENDING",
 *             name: "property_b",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Index can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:datastore/dataStoreIndex:DataStoreIndex default projects/{{project}}/indexes/{{index_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:datastore/dataStoreIndex:DataStoreIndex default {{project}}/{{index_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:datastore/dataStoreIndex:DataStoreIndex default {{index_id}}
 * ```
 */
export class DataStoreIndex extends pulumi.CustomResource {
    /**
     * Get an existing DataStoreIndex resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataStoreIndexState, opts?: pulumi.CustomResourceOptions): DataStoreIndex {
        return new DataStoreIndex(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:datastore/dataStoreIndex:DataStoreIndex';

    /**
     * Returns true if the given object is an instance of DataStoreIndex.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataStoreIndex {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataStoreIndex.__pulumiType;
    }

    /**
     * Policy for including ancestors in the index.
     * Default value is `NONE`.
     * Possible values are `NONE` and `ALL_ANCESTORS`.
     */
    public readonly ancestor!: pulumi.Output<string | undefined>;
    /**
     * The index id.
     */
    public /*out*/ readonly indexId!: pulumi.Output<string>;
    /**
     * The entity kind which the index applies to.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * An ordered list of properties to index on.
     * Structure is documented below.
     */
    public readonly properties!: pulumi.Output<outputs.datastore.DataStoreIndexProperty[] | undefined>;

    /**
     * Create a DataStoreIndex resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataStoreIndexArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataStoreIndexArgs | DataStoreIndexState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DataStoreIndexState | undefined;
            inputs["ancestor"] = state ? state.ancestor : undefined;
            inputs["indexId"] = state ? state.indexId : undefined;
            inputs["kind"] = state ? state.kind : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["properties"] = state ? state.properties : undefined;
        } else {
            const args = argsOrState as DataStoreIndexArgs | undefined;
            if (!args || args.kind === undefined) {
                throw new Error("Missing required property 'kind'");
            }
            inputs["ancestor"] = args ? args.ancestor : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["indexId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DataStoreIndex.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataStoreIndex resources.
 */
export interface DataStoreIndexState {
    /**
     * Policy for including ancestors in the index.
     * Default value is `NONE`.
     * Possible values are `NONE` and `ALL_ANCESTORS`.
     */
    readonly ancestor?: pulumi.Input<string>;
    /**
     * The index id.
     */
    readonly indexId?: pulumi.Input<string>;
    /**
     * The entity kind which the index applies to.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * An ordered list of properties to index on.
     * Structure is documented below.
     */
    readonly properties?: pulumi.Input<pulumi.Input<inputs.datastore.DataStoreIndexProperty>[]>;
}

/**
 * The set of arguments for constructing a DataStoreIndex resource.
 */
export interface DataStoreIndexArgs {
    /**
     * Policy for including ancestors in the index.
     * Default value is `NONE`.
     * Possible values are `NONE` and `ALL_ANCESTORS`.
     */
    readonly ancestor?: pulumi.Input<string>;
    /**
     * The entity kind which the index applies to.
     */
    readonly kind: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * An ordered list of properties to index on.
     * Structure is documented below.
     */
    readonly properties?: pulumi.Input<pulumi.Input<inputs.datastore.DataStoreIndexProperty>[]>;
}
