// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Projects can be imported using the `project_id`, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:projects/usageExportBucket:UsageExportBucket my_project your-project-id
 * ```
 */
export class UsageExportBucket extends pulumi.CustomResource {
    /**
     * Get an existing UsageExportBucket resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UsageExportBucketState, opts?: pulumi.CustomResourceOptions): UsageExportBucket {
        return new UsageExportBucket(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:projects/usageExportBucket:UsageExportBucket';

    /**
     * Returns true if the given object is an instance of UsageExportBucket.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UsageExportBucket {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UsageExportBucket.__pulumiType;
    }

    /**
     * The bucket to store reports in.
     */
    public readonly bucketName!: pulumi.Output<string>;
    /**
     * A prefix for the reports, for instance, the project name.
     */
    public readonly prefix!: pulumi.Output<string | undefined>;
    /**
     * The project to set the export bucket on. If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a UsageExportBucket resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UsageExportBucketArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UsageExportBucketArgs | UsageExportBucketState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UsageExportBucketState | undefined;
            inputs["bucketName"] = state ? state.bucketName : undefined;
            inputs["prefix"] = state ? state.prefix : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as UsageExportBucketArgs | undefined;
            if ((!args || args.bucketName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'bucketName'");
            }
            inputs["bucketName"] = args ? args.bucketName : undefined;
            inputs["prefix"] = args ? args.prefix : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(UsageExportBucket.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UsageExportBucket resources.
 */
export interface UsageExportBucketState {
    /**
     * The bucket to store reports in.
     */
    readonly bucketName?: pulumi.Input<string>;
    /**
     * A prefix for the reports, for instance, the project name.
     */
    readonly prefix?: pulumi.Input<string>;
    /**
     * The project to set the export bucket on. If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UsageExportBucket resource.
 */
export interface UsageExportBucketArgs {
    /**
     * The bucket to store reports in.
     */
    readonly bucketName: pulumi.Input<string>;
    /**
     * A prefix for the reports, for instance, the project name.
     */
    readonly prefix?: pulumi.Input<string>;
    /**
     * The project to set the export bucket on. If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
