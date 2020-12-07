// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a organization-level logging bucket config. For more information see
 * [the official logging documentation](https://cloud.google.com/logging/docs/) and
 * [Storing Logs](https://cloud.google.com/logging/docs/storage).
 *
 * > **Note:** Logging buckets are automatically created for a given folder, project, organization, billingAccount and cannot be deleted. Creating a resource of this type will acquire and update the resource that already exists at the desired location. These buckets cannot be removed so deleting this resource will remove the bucket config from your state but will leave the logging bucket unchanged. The buckets that are currently automatically created are "_Default" and "_Required".
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const default = gcp.organizations.getOrganization({
 *     organization: "123456789",
 * });
 * const basic = new gcp.logging.OrganizationBucketConfig("basic", {
 *     organization: _default.then(_default => _default.organization),
 *     location: "global",
 *     retentionDays: 30,
 *     bucketId: "_Default",
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported using the following format
 *
 * ```sh
 *  $ pulumi import gcp:logging/organizationBucketConfig:OrganizationBucketConfig default organizations/{{organization}}/locations/{{location}}/buckets/{{bucket_id}}
 * ```
 */
export class OrganizationBucketConfig extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationBucketConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationBucketConfigState, opts?: pulumi.CustomResourceOptions): OrganizationBucketConfig {
        return new OrganizationBucketConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:logging/organizationBucketConfig:OrganizationBucketConfig';

    /**
     * Returns true if the given object is an instance of OrganizationBucketConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationBucketConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationBucketConfig.__pulumiType;
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
     * The resource name of the bucket. For example: "organizations/my-organization-id/locations/my-location/buckets/my-bucket-id"
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The parent resource that contains the logging bucket.
     */
    public readonly organization!: pulumi.Output<string>;
    /**
     * Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
     */
    public readonly retentionDays!: pulumi.Output<number | undefined>;

    /**
     * Create a OrganizationBucketConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationBucketConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationBucketConfigArgs | OrganizationBucketConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as OrganizationBucketConfigState | undefined;
            inputs["bucketId"] = state ? state.bucketId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["lifecycleState"] = state ? state.lifecycleState : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["organization"] = state ? state.organization : undefined;
            inputs["retentionDays"] = state ? state.retentionDays : undefined;
        } else {
            const args = argsOrState as OrganizationBucketConfigArgs | undefined;
            if ((!args || args.bucketId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'bucketId'");
            }
            if ((!args || args.location === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.organization === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'organization'");
            }
            inputs["bucketId"] = args ? args.bucketId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["organization"] = args ? args.organization : undefined;
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
        super(OrganizationBucketConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationBucketConfig resources.
 */
export interface OrganizationBucketConfigState {
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
     * The resource name of the bucket. For example: "organizations/my-organization-id/locations/my-location/buckets/my-bucket-id"
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The parent resource that contains the logging bucket.
     */
    readonly organization?: pulumi.Input<string>;
    /**
     * Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
     */
    readonly retentionDays?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a OrganizationBucketConfig resource.
 */
export interface OrganizationBucketConfigArgs {
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
    readonly organization: pulumi.Input<string>;
    /**
     * Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
     */
    readonly retentionDays?: pulumi.Input<number>;
}
