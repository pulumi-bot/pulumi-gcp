// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../types/input";
import * as outputApi from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for pubsub subscription. Each of these resources serves a different use case:
 * 
 * * `gcp.pubsub.SubscriptionIAMPolicy`: Authoritative. Sets the IAM policy for the subscription and replaces any existing policy already attached.
 * * `gcp.pubsub.SubscriptionIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subscription are preserved.
 * * `gcp.pubsub.SubscriptionIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subscription are preserved.
 * 
 * > **Note:** `gcp.pubsub.SubscriptionIAMPolicy` **cannot** be used in conjunction with `gcp.pubsub.SubscriptionIAMBinding` and `gcp.pubsub.SubscriptionIAMMember` or they will fight over what your policy should be.
 * 
 * > **Note:** `gcp.pubsub.SubscriptionIAMBinding` resources **can be** used in conjunction with `gcp.pubsub.SubscriptionIAMMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_pubsub\_subscription\_iam\_policy
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const admin = pulumi.output(gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         members: ["user:jane@example.com"],
 *         role: "roles/editor",
 *     }],
 * }));
 * const editor = new gcp.pubsub.SubscriptionIAMPolicy("editor", {
 *     policyData: admin.policyData,
 *     subscription: "your-subscription-name",
 * });
 * ```
 * 
 * ## google\_pubsub\_subscription\_iam\_binding
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const editor = new gcp.pubsub.SubscriptionIAMBinding("editor", {
 *     members: ["user:jane@example.com"],
 *     role: "roles/editor",
 *     subscription: "your-subscription-name",
 * });
 * ```
 * 
 * ## google\_pubsub\_subscription\_iam\_member
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const editor = new gcp.pubsub.SubscriptionIAMMember("editor", {
 *     member: "user:jane@example.com",
 *     role: "roles/editor",
 *     subscription: "your-subscription-name",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/pubsub_subscription_iam_binding.html.markdown.
 */
export class SubscriptionIAMBinding extends pulumi.CustomResource {
    /**
     * Get an existing SubscriptionIAMBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubscriptionIAMBindingState, opts?: pulumi.CustomResourceOptions): SubscriptionIAMBinding {
        return new SubscriptionIAMBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:pubsub/subscriptionIAMBinding:SubscriptionIAMBinding';

    /**
     * Returns true if the given object is an instance of SubscriptionIAMBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SubscriptionIAMBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SubscriptionIAMBinding.__pulumiType;
    }

    /**
     * (Computed) The etag of the subscription's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The subscription name or id to bind to attach IAM policy to.
     */
    public readonly subscription!: pulumi.Output<string>;

    /**
     * Create a SubscriptionIAMBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubscriptionIAMBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubscriptionIAMBindingArgs | SubscriptionIAMBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SubscriptionIAMBindingState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["subscription"] = state ? state.subscription : undefined;
        } else {
            const args = argsOrState as SubscriptionIAMBindingArgs | undefined;
            if (!args || args.members === undefined) {
                throw new Error("Missing required property 'members'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            if (!args || args.subscription === undefined) {
                throw new Error("Missing required property 'subscription'");
            }
            inputs["members"] = args ? args.members : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["subscription"] = args ? args.subscription : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SubscriptionIAMBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SubscriptionIAMBinding resources.
 */
export interface SubscriptionIAMBindingState {
    /**
     * (Computed) The etag of the subscription's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * The subscription name or id to bind to attach IAM policy to.
     */
    readonly subscription?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SubscriptionIAMBinding resource.
 */
export interface SubscriptionIAMBindingArgs {
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
    /**
     * The subscription name or id to bind to attach IAM policy to.
     */
    readonly subscription: pulumi.Input<string>;
}
