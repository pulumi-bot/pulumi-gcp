// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Budget configuration for a billing account.
 *
 * To get more information about Budget, see:
 *
 * * [API documentation](https://cloud.google.com/billing/docs/reference/budget/rest/v1beta1/billingAccounts.budgets)
 * * How-to Guides
 *     * [Creating a budget](https://cloud.google.com/billing/docs/how-to/budgets)
 *
 * ## Example Usage
 *
 * ### Billing Budget Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const account = gcp.organizations.getBillingAccount({
 *     billingAccount: "000000-0000000-0000000-000000",
 * });
 * const budget = new gcp.billing.Budget("budget", {
 *     billingAccount: account.then(account => account.id),
 *     displayName: "Example Billing Budget",
 *     amount: {
 *         specified_amount: {
 *             currencyCode: "USD",
 *             units: "100000",
 *         },
 *     },
 *     threshold_rules: [{
 *         thresholdPercent: 0.5,
 *     }],
 * });
 * ```
 *
 * ### Billing Budget Filter
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const account = gcp.organizations.getBillingAccount({
 *     billingAccount: "000000-0000000-0000000-000000",
 * });
 * const budget = new gcp.billing.Budget("budget", {
 *     billingAccount: account.then(account => account.id),
 *     displayName: "Example Billing Budget",
 *     budget_filter: {
 *         projects: ["projects/my-project-name"],
 *         creditTypesTreatment: "EXCLUDE_ALL_CREDITS",
 *         services: ["services/24E6-581D-38E5"],
 *     },
 *     amount: {
 *         specified_amount: {
 *             currencyCode: "USD",
 *             units: "100000",
 *         },
 *     },
 *     threshold_rules: [
 *         {
 *             thresholdPercent: 0.5,
 *         },
 *         {
 *             thresholdPercent: 0.9,
 *             spendBasis: "FORECASTED_SPEND",
 *         },
 *     ],
 * });
 * ```
 */
export class Budget extends pulumi.CustomResource {
    /**
     * Get an existing Budget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BudgetState, opts?: pulumi.CustomResourceOptions): Budget {
        return new Budget(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:billing/budget:Budget';

    /**
     * Returns true if the given object is an instance of Budget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Budget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Budget.__pulumiType;
    }

    /**
     * Defines notifications that are sent on every update to the
     * billing account's spend, regardless of the thresholds defined
     * using threshold rules.  Structure is documented below.
     */
    public readonly allUpdatesRule!: pulumi.Output<outputs.billing.BudgetAllUpdatesRule | undefined>;
    /**
     * The budgeted amount for each usage period.  Structure is documented below.
     */
    public readonly amount!: pulumi.Output<outputs.billing.BudgetAmount>;
    /**
     * ID of the billing account to set a budget on.
     */
    public readonly billingAccount!: pulumi.Output<string>;
    /**
     * Filters that define which resources are used to compute the actual
     * spend against the budget.  Structure is documented below.
     */
    public readonly budgetFilter!: pulumi.Output<outputs.billing.BudgetBudgetFilter | undefined>;
    /**
     * User data for display name in UI. Must be <= 60 chars.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Resource name of the budget. The resource name implies the scope of a budget. Values are of the form
     * billingAccounts/{billingAccountId}/budgets/{budgetId}.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Rules that trigger alerts (notifications of thresholds being
     * crossed) when spend exceeds the specified percentages of the
     * budget.  Structure is documented below.
     */
    public readonly thresholdRules!: pulumi.Output<outputs.billing.BudgetThresholdRule[]>;

    /**
     * Create a Budget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BudgetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BudgetArgs | BudgetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BudgetState | undefined;
            inputs["allUpdatesRule"] = state ? state.allUpdatesRule : undefined;
            inputs["amount"] = state ? state.amount : undefined;
            inputs["billingAccount"] = state ? state.billingAccount : undefined;
            inputs["budgetFilter"] = state ? state.budgetFilter : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["thresholdRules"] = state ? state.thresholdRules : undefined;
        } else {
            const args = argsOrState as BudgetArgs | undefined;
            if (!args || args.amount === undefined) {
                throw new Error("Missing required property 'amount'");
            }
            if (!args || args.billingAccount === undefined) {
                throw new Error("Missing required property 'billingAccount'");
            }
            if (!args || args.thresholdRules === undefined) {
                throw new Error("Missing required property 'thresholdRules'");
            }
            inputs["allUpdatesRule"] = args ? args.allUpdatesRule : undefined;
            inputs["amount"] = args ? args.amount : undefined;
            inputs["billingAccount"] = args ? args.billingAccount : undefined;
            inputs["budgetFilter"] = args ? args.budgetFilter : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["thresholdRules"] = args ? args.thresholdRules : undefined;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Budget.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Budget resources.
 */
export interface BudgetState {
    /**
     * Defines notifications that are sent on every update to the
     * billing account's spend, regardless of the thresholds defined
     * using threshold rules.  Structure is documented below.
     */
    readonly allUpdatesRule?: pulumi.Input<inputs.billing.BudgetAllUpdatesRule>;
    /**
     * The budgeted amount for each usage period.  Structure is documented below.
     */
    readonly amount?: pulumi.Input<inputs.billing.BudgetAmount>;
    /**
     * ID of the billing account to set a budget on.
     */
    readonly billingAccount?: pulumi.Input<string>;
    /**
     * Filters that define which resources are used to compute the actual
     * spend against the budget.  Structure is documented below.
     */
    readonly budgetFilter?: pulumi.Input<inputs.billing.BudgetBudgetFilter>;
    /**
     * User data for display name in UI. Must be <= 60 chars.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Resource name of the budget. The resource name implies the scope of a budget. Values are of the form
     * billingAccounts/{billingAccountId}/budgets/{budgetId}.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Rules that trigger alerts (notifications of thresholds being
     * crossed) when spend exceeds the specified percentages of the
     * budget.  Structure is documented below.
     */
    readonly thresholdRules?: pulumi.Input<pulumi.Input<inputs.billing.BudgetThresholdRule>[]>;
}

/**
 * The set of arguments for constructing a Budget resource.
 */
export interface BudgetArgs {
    /**
     * Defines notifications that are sent on every update to the
     * billing account's spend, regardless of the thresholds defined
     * using threshold rules.  Structure is documented below.
     */
    readonly allUpdatesRule?: pulumi.Input<inputs.billing.BudgetAllUpdatesRule>;
    /**
     * The budgeted amount for each usage period.  Structure is documented below.
     */
    readonly amount: pulumi.Input<inputs.billing.BudgetAmount>;
    /**
     * ID of the billing account to set a budget on.
     */
    readonly billingAccount: pulumi.Input<string>;
    /**
     * Filters that define which resources are used to compute the actual
     * spend against the budget.  Structure is documented below.
     */
    readonly budgetFilter?: pulumi.Input<inputs.billing.BudgetBudgetFilter>;
    /**
     * User data for display name in UI. Must be <= 60 chars.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Rules that trigger alerts (notifications of thresholds being
     * crossed) when spend exceeds the specified percentages of the
     * budget.  Structure is documented below.
     */
    readonly thresholdRules: pulumi.Input<pulumi.Input<inputs.billing.BudgetThresholdRule>[]>;
}
