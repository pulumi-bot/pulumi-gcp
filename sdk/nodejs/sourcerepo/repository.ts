// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A repository (or repo) is a Git repository storing versioned source content.
 *
 *
 * To get more information about Repository, see:
 *
 * * [API documentation](https://cloud.google.com/source-repositories/docs/reference/rest/v1/projects.repos)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/source-repositories/)
 *
 * ## Example Usage
 *
 * ### Sourcerepo Repository Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const my_repo = new gcp.sourcerepo.Repository("my-repo", {});
 * ```
 *
 * ### Sourcerepo Repository Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const test_account = new gcp.serviceAccount.Account("test-account", {
 *     accountId: "my-account",
 *     displayName: "Test Service Account",
 * });
 * const topic = new gcp.pubsub.Topic("topic", {});
 * const my_repo = new gcp.sourcerepo.Repository("my-repo", {pubsub_configs: [{
 *     topic: topic.id,
 *     messageFormat: "JSON",
 *     serviceAccountEmail: test_account.email,
 * }]});
 * ```
 */
export class Repository extends pulumi.CustomResource {
    /**
     * Get an existing Repository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryState, opts?: pulumi.CustomResourceOptions): Repository {
        return new Repository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:sourcerepo/repository:Repository';

    /**
     * Returns true if the given object is an instance of Repository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Repository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Repository.__pulumiType;
    }

    /**
     * Resource name of the repository, of the form `{{repo}}`.
     * The repo name may contain slashes. eg, `name/with/slash`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * How this repository publishes a change in the repository through Cloud Pub/Sub.
     * Keyed by the topic names.  Structure is documented below.
     */
    public readonly pubsubConfigs!: pulumi.Output<outputs.sourcerepo.RepositoryPubsubConfig[] | undefined>;
    /**
     * The disk usage of the repo, in bytes.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * URL to clone the repository from Google Cloud Source Repositories.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a Repository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryArgs | RepositoryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RepositoryState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["pubsubConfigs"] = state ? state.pubsubConfigs : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as RepositoryArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["pubsubConfigs"] = args ? args.pubsubConfigs : undefined;
            inputs["size"] = undefined /*out*/;
            inputs["url"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Repository.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Repository resources.
 */
export interface RepositoryState {
    /**
     * Resource name of the repository, of the form `{{repo}}`.
     * The repo name may contain slashes. eg, `name/with/slash`
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * How this repository publishes a change in the repository through Cloud Pub/Sub.
     * Keyed by the topic names.  Structure is documented below.
     */
    readonly pubsubConfigs?: pulumi.Input<pulumi.Input<inputs.sourcerepo.RepositoryPubsubConfig>[]>;
    /**
     * The disk usage of the repo, in bytes.
     */
    readonly size?: pulumi.Input<number>;
    /**
     * URL to clone the repository from Google Cloud Source Repositories.
     */
    readonly url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Repository resource.
 */
export interface RepositoryArgs {
    /**
     * Resource name of the repository, of the form `{{repo}}`.
     * The repo name may contain slashes. eg, `name/with/slash`
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * How this repository publishes a change in the repository through Cloud Pub/Sub.
     * Keyed by the topic names.  Structure is documented below.
     */
    readonly pubsubConfigs?: pulumi.Input<pulumi.Input<inputs.sourcerepo.RepositoryPubsubConfig>[]>;
}
