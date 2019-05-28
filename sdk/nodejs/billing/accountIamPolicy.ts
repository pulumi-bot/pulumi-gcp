// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows management of the entire IAM policy for an existing Google Cloud Platform Billing Account.
 * 
 * > **Warning:** Billing accounts have a default user that can be **overwritten**
 * by use of this resource. The safest alternative is to use multiple `google_billing_account_iam_binding`
 *    resources. If you do use this resource, the best way to be sure that you are
 *    not making dangerous changes is to start by importing your existing policy,
 *    and examining the diff very closely.
 * 
 * > **Note:** This resource __must not__ be used in conjunction with
 *    `google_billing_account_iam_member` or `google_billing_account_iam_binding`
 *    or they will fight over what your policy should be.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const admin = pulumi.output(gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         members: ["user:jane@example.com"],
 *         role: "roles/billing.viewer",
 *     }],
 * }));
 * const policy = new gcp.billing.AccountIamPolicy("policy", {
 *     billingAccountId: "00AA00-000AAA-00AA0A",
 *     policyData: admin.policyData,
 * });
 * ```
 */
export class AccountIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AccountIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountIamPolicyState, opts?: pulumi.CustomResourceOptions): AccountIamPolicy {
        return new AccountIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /**
     * The billing account id.
     */
    public readonly billingAccountId!: pulumi.Output<string>;
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The `google_iam_policy` data source that represents
     * the IAM policy that will be applied to the billing account. This policy overrides any existing
     * policy applied to the billing account.
     */
    public readonly policyData!: pulumi.Output<string>;

    /**
     * Create a AccountIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountIamPolicyArgs | AccountIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AccountIamPolicyState | undefined;
            inputs["billingAccountId"] = state ? state.billingAccountId : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
        } else {
            const args = argsOrState as AccountIamPolicyArgs | undefined;
            if (!args || args.billingAccountId === undefined) {
                throw new Error("Missing required property 'billingAccountId'");
            }
            if (!args || args.policyData === undefined) {
                throw new Error("Missing required property 'policyData'");
            }
            inputs["billingAccountId"] = args ? args.billingAccountId : undefined;
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        super("gcp:billing/accountIamPolicy:AccountIamPolicy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccountIamPolicy resources.
 */
export interface AccountIamPolicyState {
    /**
     * The billing account id.
     */
    readonly billingAccountId?: pulumi.Input<string>;
    readonly etag?: pulumi.Input<string>;
    /**
     * The `google_iam_policy` data source that represents
     * the IAM policy that will be applied to the billing account. This policy overrides any existing
     * policy applied to the billing account.
     */
    readonly policyData?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccountIamPolicy resource.
 */
export interface AccountIamPolicyArgs {
    /**
     * The billing account id.
     */
    readonly billingAccountId: pulumi.Input<string>;
    /**
     * The `google_iam_policy` data source that represents
     * the IAM policy that will be applied to the billing account. This policy overrides any existing
     * policy applied to the billing account.
     */
    readonly policyData: pulumi.Input<string>;
}
