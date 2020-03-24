// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Cloud Pub/Sub Topic. Each of these resources serves a different use case:
 * 
 * * `gcp.pubsub.TopicIAMPolicy`: Authoritative. Sets the IAM policy for the topic and replaces any existing policy already attached.
 * * `gcp.pubsub.TopicIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the topic are preserved.
 * * `gcp.pubsub.TopicIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the topic are preserved.
 * 
 * > **Note:** `gcp.pubsub.TopicIAMPolicy` **cannot** be used in conjunction with `gcp.pubsub.TopicIAMBinding` and `gcp.pubsub.TopicIAMMember` or they will fight over what your policy should be.
 * 
 * > **Note:** `gcp.pubsub.TopicIAMBinding` resources **can be** used in conjunction with `gcp.pubsub.TopicIAMMember` resources **only if** they do not grant privilege to the same role.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/pubsub_topic_iam.html.markdown.
 */
export class TopicIAMMember extends pulumi.CustomResource {
    /**
     * Get an existing TopicIAMMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TopicIAMMemberState, opts?: pulumi.CustomResourceOptions): TopicIAMMember {
        return new TopicIAMMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:pubsub/topicIAMMember:TopicIAMMember';

    /**
     * Returns true if the given object is an instance of TopicIAMMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TopicIAMMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TopicIAMMember.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.pubsub.TopicIAMMemberCondition | undefined>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly member!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly topic!: pulumi.Output<string>;

    /**
     * Create a TopicIAMMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TopicIAMMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TopicIAMMemberArgs | TopicIAMMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TopicIAMMemberState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["topic"] = state ? state.topic : undefined;
        } else {
            const args = argsOrState as TopicIAMMemberArgs | undefined;
            if (!args || args.member === undefined) {
                throw new Error("Missing required property 'member'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            if (!args || args.topic === undefined) {
                throw new Error("Missing required property 'topic'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["member"] = args ? args.member : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["topic"] = args ? args.topic : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TopicIAMMember.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TopicIAMMember resources.
 */
export interface TopicIAMMemberState {
    readonly condition?: pulumi.Input<inputs.pubsub.TopicIAMMemberCondition>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly member?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly topic?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TopicIAMMember resource.
 */
export interface TopicIAMMemberArgs {
    readonly condition?: pulumi.Input<inputs.pubsub.TopicIAMMemberCondition>;
    readonly member: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly topic: pulumi.Input<string>;
}
