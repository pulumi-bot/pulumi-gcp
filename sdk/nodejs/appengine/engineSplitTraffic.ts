// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Traffic routing configuration for versions within a single service. Traffic splits define how traffic directed to the service is assigned to versions.
 *
 *
 * To get more information about ServiceSplitTraffic, see:
 *
 * * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services)
 *
 * ## Example Usage - App Engine Service Split Traffic
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const bucket = new gcp.storage.Bucket("bucket", {});
 * const object = new gcp.storage.BucketObject("object", {
 *     bucket: bucket.name,
 *     source: new pulumi.asset.FileAsset("./test-fixtures/appengine/hello-world.zip"),
 * });
 * const liveappV1 = new gcp.appengine.StandardAppVersion("liveappV1", {
 *     versionId: "v1",
 *     service: "liveapp",
 *     deleteServiceOnDestroy: true,
 *     runtime: "nodejs10",
 *     entrypoint: {
 *         shell: "node ./app.js",
 *     },
 *     deployment: {
 *         zip: {
 *             sourceUrl: pulumi.interpolate`https://storage.googleapis.com/${bucket.name}/${object.name}`,
 *         },
 *     },
 *     envVariables: {
 *         port: "8080",
 *     },
 * });
 * const liveappV2 = new gcp.appengine.StandardAppVersion("liveappV2", {
 *     versionId: "v2",
 *     service: "liveapp",
 *     noopOnDestroy: true,
 *     runtime: "nodejs10",
 *     entrypoint: {
 *         shell: "node ./app.js",
 *     },
 *     deployment: {
 *         zip: {
 *             sourceUrl: pulumi.interpolate`https://storage.googleapis.com/${bucket.name}/${object.name}`,
 *         },
 *     },
 *     envVariables: {
 *         port: "8080",
 *     },
 * });
 * const liveapp = new gcp.appengine.EngineSplitTraffic("liveapp", {
 *     service: liveappV2.service,
 *     migrateTraffic: false,
 *     split: {
 *         shardBy: "IP",
 *         allocations: pulumi.all([liveappV1.versionId, liveappV2.versionId]).apply(([liveappV1VersionId, liveappV2VersionId]) => {
 *             [liveappV1VersionId]: 0.75,
 *             [liveappV2VersionId]: 0.25,
 *         }),
 *     },
 * });
 * ```
 */
export class EngineSplitTraffic extends pulumi.CustomResource {
    /**
     * Get an existing EngineSplitTraffic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EngineSplitTrafficState, opts?: pulumi.CustomResourceOptions): EngineSplitTraffic {
        return new EngineSplitTraffic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:appengine/engineSplitTraffic:EngineSplitTraffic';

    /**
     * Returns true if the given object is an instance of EngineSplitTraffic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EngineSplitTraffic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EngineSplitTraffic.__pulumiType;
    }

    /**
     * If set to true traffic will be migrated to this version.
     */
    public readonly migrateTraffic!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The name of the service these settings apply to.
     */
    public readonly service!: pulumi.Output<string>;
    /**
     * Mapping that defines fractional HTTP traffic diversion to different versions within the service.  Structure is documented below.
     */
    public readonly split!: pulumi.Output<outputs.appengine.EngineSplitTrafficSplit>;

    /**
     * Create a EngineSplitTraffic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EngineSplitTrafficArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EngineSplitTrafficArgs | EngineSplitTrafficState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EngineSplitTrafficState | undefined;
            inputs["migrateTraffic"] = state ? state.migrateTraffic : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["service"] = state ? state.service : undefined;
            inputs["split"] = state ? state.split : undefined;
        } else {
            const args = argsOrState as EngineSplitTrafficArgs | undefined;
            if (!args || args.service === undefined) {
                throw new Error("Missing required property 'service'");
            }
            if (!args || args.split === undefined) {
                throw new Error("Missing required property 'split'");
            }
            inputs["migrateTraffic"] = args ? args.migrateTraffic : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["service"] = args ? args.service : undefined;
            inputs["split"] = args ? args.split : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EngineSplitTraffic.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EngineSplitTraffic resources.
 */
export interface EngineSplitTrafficState {
    /**
     * If set to true traffic will be migrated to this version.
     */
    readonly migrateTraffic?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The name of the service these settings apply to.
     */
    readonly service?: pulumi.Input<string>;
    /**
     * Mapping that defines fractional HTTP traffic diversion to different versions within the service.  Structure is documented below.
     */
    readonly split?: pulumi.Input<inputs.appengine.EngineSplitTrafficSplit>;
}

/**
 * The set of arguments for constructing a EngineSplitTraffic resource.
 */
export interface EngineSplitTrafficArgs {
    /**
     * If set to true traffic will be migrated to this version.
     */
    readonly migrateTraffic?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The name of the service these settings apply to.
     */
    readonly service: pulumi.Input<string>;
    /**
     * Mapping that defines fractional HTTP traffic diversion to different versions within the service.  Structure is documented below.
     */
    readonly split: pulumi.Input<inputs.appengine.EngineSplitTrafficSplit>;
}
