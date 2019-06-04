// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for pubsub subscription. Each of these resources serves a different use case:
 * 
 * * `google_pubsub_subscription_iam_policy`: Authoritative. Sets the IAM policy for the subscription and replaces any existing policy already attached.
 * * `google_pubsub_subscription_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subscription are preserved.
 * * `google_pubsub_subscription_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subscription are preserved.
 * 
 * > **Note:** `google_pubsub_subscription_iam_policy` **cannot** be used in conjunction with `google_pubsub_subscription_iam_binding` and `google_pubsub_subscription_iam_member` or they will fight over what your policy should be.
 * 
 * > **Note:** `google_pubsub_subscription_iam_binding` resources **can be** used in conjunction with `google_pubsub_subscription_iam_member` resources **only if** they do not grant privilege to the same role.
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
 */
export class SubscriptionIAMPolicy extends pulumi.CustomResource {
    /**
     * Get an existing SubscriptionIAMPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubscriptionIAMPolicyState, opts?: pulumi.CustomResourceOptions): SubscriptionIAMPolicy {
        return new SubscriptionIAMPolicy(name, <any>state, { ...opts, id: id });
    }

    private static readonly __pulumiType = 'gcp:pubsub/subscriptionIAMPolicy:SubscriptionIAMPolicy';

    /**
     * Returns true if the given object is an instance of SubscriptionIAMPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SubscriptionIAMPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }

        const t = obj['__pulumiType'];
        if (typeof t !== 'string') {
            return false;
        }

        return t === SubscriptionIAMPolicy.__pulumiType;
    }

    /**
     * (Computed) The etag of the subscription's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The policy data generated by
     * a `google_iam_policy` data source.
     */
    public readonly policyData!: pulumi.Output<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The subscription name or id to bind to attach IAM policy to.
     */
    public readonly subscription!: pulumi.Output<string>;

    /**
     * Create a SubscriptionIAMPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubscriptionIAMPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubscriptionIAMPolicyArgs | SubscriptionIAMPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SubscriptionIAMPolicyState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["subscription"] = state ? state.subscription : undefined;
        } else {
            const args = argsOrState as SubscriptionIAMPolicyArgs | undefined;
            if (!args || args.policyData === undefined) {
                throw new Error("Missing required property 'policyData'");
            }
            if (!args || args.subscription === undefined) {
                throw new Error("Missing required property 'subscription'");
            }
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["subscription"] = args ? args.subscription : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        super(SubscriptionIAMPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SubscriptionIAMPolicy resources.
 */
export interface SubscriptionIAMPolicyState {
    /**
     * (Computed) The etag of the subscription's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `google_iam_policy` data source.
     */
    readonly policyData?: pulumi.Input<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The subscription name or id to bind to attach IAM policy to.
     */
    readonly subscription?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SubscriptionIAMPolicy resource.
 */
export interface SubscriptionIAMPolicyArgs {
    /**
     * The policy data generated by
     * a `google_iam_policy` data source.
     */
    readonly policyData: pulumi.Input<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The subscription name or id to bind to attach IAM policy to.
     */
    readonly subscription: pulumi.Input<string>;
}
