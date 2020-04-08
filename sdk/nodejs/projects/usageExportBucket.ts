// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Allows creation and management of a Google Cloud Platform project.
 * 
 * Projects created with this resource must be associated with an Organization.
 * See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.
 * 
 * The service account used to run this provider when creating a `gcp.organizations.Project`
 * resource must have `roles/resourcemanager.projectCreator`. See the
 * [Access Control for Organizations Using IAM](https://cloud.google.com/resource-manager/docs/access-control-org)
 * doc for more information.
 * 
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/google_project.html.markdown.
 */
export class UsageExportBucket extends pulumi.CustomResource {
    /**
     * Get an existing UsageExportBucket resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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

    public readonly bucketName!: pulumi.Output<string>;
    public readonly prefix!: pulumi.Output<string | undefined>;
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
            if (!args || args.bucketName === undefined) {
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
    readonly bucketName?: pulumi.Input<string>;
    readonly prefix?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UsageExportBucket resource.
 */
export interface UsageExportBucketArgs {
    readonly bucketName: pulumi.Input<string>;
    readonly prefix?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
}
