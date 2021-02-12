// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A named resource to which messages are sent by publishers.
 *
 * To get more information about Topic, see:
 *
 * * [API documentation](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.topics)
 * * How-to Guides
 *     * [Managing Topics](https://cloud.google.com/pubsub/docs/admin#managing_topics)
 *
 * > **Note:** You can retrieve the email of the Google Managed Pub/Sub Service Account used for forwarding
 * by using the `gcp.projects.ServiceIdentity` resource.
 *
 * ## Example Usage
 * ### Pubsub Topic Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const example = new gcp.pubsub.Topic("example", {
 *     labels: {
 *         foo: "bar",
 *     },
 * });
 * ```
 * ### Pubsub Topic Cmek
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const keyRing = new gcp.kms.KeyRing("keyRing", {location: "global"});
 * const cryptoKey = new gcp.kms.CryptoKey("cryptoKey", {keyRing: keyRing.id});
 * const example = new gcp.pubsub.Topic("example", {kmsKeyName: cryptoKey.id});
 * ```
 * ### Pubsub Topic Geo Restricted
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const example = new gcp.pubsub.Topic("example", {
 *     messageStoragePolicy: {
 *         allowedPersistenceRegions: ["europe-west3"],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Topic can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:pubsub/topic:Topic default projects/{{project}}/topics/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:pubsub/topic:Topic default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:pubsub/topic:Topic default {{name}}
 * ```
 */
export class Topic extends pulumi.CustomResource {
    /**
     * Get an existing Topic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TopicState, opts?: pulumi.CustomResourceOptions): Topic {
        return new Topic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:pubsub/topic:Topic';

    /**
     * Returns true if the given object is an instance of Topic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Topic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Topic.__pulumiType;
    }

    /**
     * The resource name of the Cloud KMS CryptoKey to be used to protect access
     * to messages published on this topic. Your project's PubSub service account
     * (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
     * `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
     * The expected format is `projects/*&#47;locations/*&#47;keyRings/*&#47;cryptoKeys/*`
     */
    public readonly kmsKeyName!: pulumi.Output<string | undefined>;
    /**
     * A set of key/value label pairs to assign to this Topic.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Policy constraining the set of Google Cloud Platform regions where
     * messages published to the topic may be stored. If not present, then no
     * constraints are in effect.
     * Structure is documented below.
     */
    public readonly messageStoragePolicy!: pulumi.Output<outputs.pubsub.TopicMessageStoragePolicy>;
    /**
     * Name of the topic.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a Topic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TopicArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TopicArgs | TopicState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TopicState | undefined;
            inputs["kmsKeyName"] = state ? state.kmsKeyName : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["messageStoragePolicy"] = state ? state.messageStoragePolicy : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as TopicArgs | undefined;
            inputs["kmsKeyName"] = args ? args.kmsKeyName : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["messageStoragePolicy"] = args ? args.messageStoragePolicy : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Topic.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Topic resources.
 */
export interface TopicState {
    /**
     * The resource name of the Cloud KMS CryptoKey to be used to protect access
     * to messages published on this topic. Your project's PubSub service account
     * (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
     * `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
     * The expected format is `projects/*&#47;locations/*&#47;keyRings/*&#47;cryptoKeys/*`
     */
    readonly kmsKeyName?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to this Topic.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Policy constraining the set of Google Cloud Platform regions where
     * messages published to the topic may be stored. If not present, then no
     * constraints are in effect.
     * Structure is documented below.
     */
    readonly messageStoragePolicy?: pulumi.Input<inputs.pubsub.TopicMessageStoragePolicy>;
    /**
     * Name of the topic.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Topic resource.
 */
export interface TopicArgs {
    /**
     * The resource name of the Cloud KMS CryptoKey to be used to protect access
     * to messages published on this topic. Your project's PubSub service account
     * (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
     * `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
     * The expected format is `projects/*&#47;locations/*&#47;keyRings/*&#47;cryptoKeys/*`
     */
    readonly kmsKeyName?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to this Topic.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Policy constraining the set of Google Cloud Platform regions where
     * messages published to the topic may be stored. If not present, then no
     * constraints are in effect.
     * Structure is documented below.
     */
    readonly messageStoragePolicy?: pulumi.Input<inputs.pubsub.TopicMessageStoragePolicy>;
    /**
     * Name of the topic.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
