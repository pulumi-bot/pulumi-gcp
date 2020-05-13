// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a project-level logging bucket config. For more information see
 * [the official logging documentation](https://cloud.google.com/logging/docs/) and
 * [Storing Logs](https://cloud.google.com/logging/docs/storage).
 * 
 * > **Note:** Logging buckets are automatically created for a given folder, project, organization, billingAccount and cannot be deleted. Creating a resource of this type will acquire and update the resource that already exists at the desired location. These buckets cannot be removed so deleting this resource will remove the bucket config from your state but will leave the logging bucket unchanged. The buckets that are currently automatically created are "_Default" and "_Required".
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const default = new gcp.organizations.Project("default", {
 *     projectId: "your-project-id",
 *     orgId: "123456789",
 * });
 * const basic = new gcp.logging.ProjectBucketConfig("basic", {
 *     project: default.name,
 *     location: "global",
 *     retentionDays: 30,
 *     bucketId: "_Default",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/logging_project_bucket_config.html.markdown.
 */
export class ProjectBucketConfig extends pulumi.CustomResource {
    /**
     * Get an existing ProjectBucketConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectBucketConfigState, opts?: pulumi.CustomResourceOptions): ProjectBucketConfig {
        return new ProjectBucketConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:logging/projectBucketConfig:ProjectBucketConfig';

    /**
     * Returns true if the given object is an instance of ProjectBucketConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectBucketConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectBucketConfig.__pulumiType;
    }

    /**
     * The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
     */
    public readonly bucketId!: pulumi.Output<string>;
    /**
     * Describes this bucket.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
     */
    public /*out*/ readonly lifecycleState!: pulumi.Output<string>;
    /**
     * The location of the bucket. The supported locations are: "global" "us-central1"
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The parent resource that contains the logging bucket.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
     */
    public readonly retentionDays!: pulumi.Output<number | undefined>;

    /**
     * Create a ProjectBucketConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectBucketConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectBucketConfigArgs | ProjectBucketConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProjectBucketConfigState | undefined;
            inputs["bucketId"] = state ? state.bucketId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["lifecycleState"] = state ? state.lifecycleState : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["retentionDays"] = state ? state.retentionDays : undefined;
        } else {
            const args = argsOrState as ProjectBucketConfigArgs | undefined;
            if (!args || args.bucketId === undefined) {
                throw new Error("Missing required property 'bucketId'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.project === undefined) {
                throw new Error("Missing required property 'project'");
            }
            inputs["bucketId"] = args ? args.bucketId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["retentionDays"] = args ? args.retentionDays : undefined;
            inputs["lifecycleState"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ProjectBucketConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectBucketConfig resources.
 */
export interface ProjectBucketConfigState {
    /**
     * The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
     */
    readonly bucketId?: pulumi.Input<string>;
    /**
     * Describes this bucket.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
     */
    readonly lifecycleState?: pulumi.Input<string>;
    /**
     * The location of the bucket. The supported locations are: "global" "us-central1"
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The parent resource that contains the logging bucket.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
     */
    readonly retentionDays?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ProjectBucketConfig resource.
 */
export interface ProjectBucketConfigArgs {
    /**
     * The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
     */
    readonly bucketId: pulumi.Input<string>;
    /**
     * Describes this bucket.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The location of the bucket. The supported locations are: "global" "us-central1"
     */
    readonly location: pulumi.Input<string>;
    /**
     * The parent resource that contains the logging bucket.
     */
    readonly project: pulumi.Input<string>;
    /**
     * Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
     */
    readonly retentionDays?: pulumi.Input<number>;
}
