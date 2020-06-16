// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a Google Cloud Bigtable GC Policy inside a family. For more information see
 * [the official documentation](https://cloud.google.com/bigtable/) and
 * [API](https://cloud.google.com/bigtable/docs/go/reference).
 *
 * {{% examples %}}
 * ## Example Usage
 * {{% example %}}
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const instance = new gcp.bigtable.Instance("instance", {cluster: [{
 *     clusterId: "tf-instance-cluster",
 *     zone: "us-central1-b",
 *     numNodes: 3,
 *     storageType: "HDD",
 * }]});
 * const table = new gcp.bigtable.Table("table", {
 *     instanceName: instance.name,
 *     column_family: [{
 *         family: "name",
 *     }],
 * });
 * const policy = new gcp.bigtable.GCPolicy("policy", {
 *     instanceName: instance.name,
 *     table: table.name,
 *     columnFamily: "name",
 *     max_age: [{
 *         days: 7,
 *     }],
 * });
 * ```
 *
 * Multiple conditions is also supported. `UNION` when any of its sub-policies apply (OR). `INTERSECTION` when all its sub-policies apply (AND)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = new gcp.bigtable.GCPolicy("policy", {
 *     instanceName: google_bigtable_instance.instance.name,
 *     table: google_bigtable_table.table.name,
 *     columnFamily: "name",
 *     mode: "UNION",
 *     max_age: [{
 *         days: 7,
 *     }],
 *     max_version: [{
 *         number: 10,
 *     }],
 * });
 * ```
 * {{% /example %}}
 * {{% /examples %}}
 */
export class GCPolicy extends pulumi.CustomResource {
    /**
     * Get an existing GCPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GCPolicyState, opts?: pulumi.CustomResourceOptions): GCPolicy {
        return new GCPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:bigtable/gCPolicy:GCPolicy';

    /**
     * Returns true if the given object is an instance of GCPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GCPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GCPolicy.__pulumiType;
    }

    /**
     * The name of the column family.
     */
    public readonly columnFamily!: pulumi.Output<string>;
    /**
     * The name of the Bigtable instance.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * GC policy that applies to all cells older than the given age.
     */
    public readonly maxAges!: pulumi.Output<outputs.bigtable.GCPolicyMaxAge[] | undefined>;
    /**
     * GC policy that applies to all versions of a cell except for the most recent.
     */
    public readonly maxVersions!: pulumi.Output<outputs.bigtable.GCPolicyMaxVersion[] | undefined>;
    /**
     * If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
     */
    public readonly mode!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The name of the table.
     */
    public readonly table!: pulumi.Output<string>;

    /**
     * Create a GCPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GCPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GCPolicyArgs | GCPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GCPolicyState | undefined;
            inputs["columnFamily"] = state ? state.columnFamily : undefined;
            inputs["instanceName"] = state ? state.instanceName : undefined;
            inputs["maxAges"] = state ? state.maxAges : undefined;
            inputs["maxVersions"] = state ? state.maxVersions : undefined;
            inputs["mode"] = state ? state.mode : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["table"] = state ? state.table : undefined;
        } else {
            const args = argsOrState as GCPolicyArgs | undefined;
            if (!args || args.columnFamily === undefined) {
                throw new Error("Missing required property 'columnFamily'");
            }
            if (!args || args.instanceName === undefined) {
                throw new Error("Missing required property 'instanceName'");
            }
            if (!args || args.table === undefined) {
                throw new Error("Missing required property 'table'");
            }
            inputs["columnFamily"] = args ? args.columnFamily : undefined;
            inputs["instanceName"] = args ? args.instanceName : undefined;
            inputs["maxAges"] = args ? args.maxAges : undefined;
            inputs["maxVersions"] = args ? args.maxVersions : undefined;
            inputs["mode"] = args ? args.mode : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["table"] = args ? args.table : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GCPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GCPolicy resources.
 */
export interface GCPolicyState {
    /**
     * The name of the column family.
     */
    readonly columnFamily?: pulumi.Input<string>;
    /**
     * The name of the Bigtable instance.
     */
    readonly instanceName?: pulumi.Input<string>;
    /**
     * GC policy that applies to all cells older than the given age.
     */
    readonly maxAges?: pulumi.Input<pulumi.Input<inputs.bigtable.GCPolicyMaxAge>[]>;
    /**
     * GC policy that applies to all versions of a cell except for the most recent.
     */
    readonly maxVersions?: pulumi.Input<pulumi.Input<inputs.bigtable.GCPolicyMaxVersion>[]>;
    /**
     * If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
     */
    readonly mode?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The name of the table.
     */
    readonly table?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GCPolicy resource.
 */
export interface GCPolicyArgs {
    /**
     * The name of the column family.
     */
    readonly columnFamily: pulumi.Input<string>;
    /**
     * The name of the Bigtable instance.
     */
    readonly instanceName: pulumi.Input<string>;
    /**
     * GC policy that applies to all cells older than the given age.
     */
    readonly maxAges?: pulumi.Input<pulumi.Input<inputs.bigtable.GCPolicyMaxAge>[]>;
    /**
     * GC policy that applies to all versions of a cell except for the most recent.
     */
    readonly maxVersions?: pulumi.Input<pulumi.Input<inputs.bigtable.GCPolicyMaxVersion>[]>;
    /**
     * If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
     */
    readonly mode?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The name of the table.
     */
    readonly table: pulumi.Input<string>;
}
