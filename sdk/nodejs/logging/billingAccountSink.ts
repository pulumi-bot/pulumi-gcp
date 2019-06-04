// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a billing account logging sink. For more information see
 * [the official documentation](https://cloud.google.com/logging/docs/) and
 * [Exporting Logs in the API](https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
 * 
 * > **Note** You must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
 * [granted on the billing account](https://cloud.google.com/billing/reference/rest/v1/billingAccounts/getIamPolicy) to
 * the credentials used with Terraform. [IAM roles granted on a billing account](https://cloud.google.com/billing/docs/how-to/billing-access) are separate from the
 * typical IAM roles granted on a project.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const log_bucket = new gcp.storage.Bucket("log-bucket", {});
 * const my_sink = new gcp.logging.BillingAccountSink("my-sink", {
 *     billingAccount: "ABCDEF-012345-GHIJKL",
 *     destination: pulumi.interpolate`storage.googleapis.com/${log_bucket.name}`,
 * });
 * const log_writer = new gcp.projects.IAMBinding("log-writer", {
 *     members: [my_sink.writerIdentity],
 *     role: "roles/storage.objectCreator",
 * });
 * ```
 */
export class BillingAccountSink extends pulumi.CustomResource {
    /**
     * Get an existing BillingAccountSink resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BillingAccountSinkState, opts?: pulumi.CustomResourceOptions): BillingAccountSink {
        return new BillingAccountSink(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:logging/billingAccountSink:BillingAccountSink';

    /**
     * Returns true if the given object is an instance of BillingAccountSink.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BillingAccountSink {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BillingAccountSink.__pulumiType;
    }

    /**
     * The billing account exported to the sink.
     */
    public readonly billingAccount!: pulumi.Output<string>;
    /**
     * The destination of the sink (or, in other words, where logs are written to). Can be a
     * Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * ```
     * The writer associated with the sink must have access to write to the above resource.
     */
    public readonly destination!: pulumi.Output<string>;
    /**
     * The filter to apply when exporting logs. Only log entries that match the filter are exported.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    public readonly filter!: pulumi.Output<string | undefined>;
    /**
     * The name of the logging sink.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The identity associated with this sink. This identity must be granted write access to the
     * configured `destination`.
     */
    public /*out*/ readonly writerIdentity!: pulumi.Output<string>;

    /**
     * Create a BillingAccountSink resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BillingAccountSinkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BillingAccountSinkArgs | BillingAccountSinkState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BillingAccountSinkState | undefined;
            inputs["billingAccount"] = state ? state.billingAccount : undefined;
            inputs["destination"] = state ? state.destination : undefined;
            inputs["filter"] = state ? state.filter : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["writerIdentity"] = state ? state.writerIdentity : undefined;
        } else {
            const args = argsOrState as BillingAccountSinkArgs | undefined;
            if (!args || args.billingAccount === undefined) {
                throw new Error("Missing required property 'billingAccount'");
            }
            if (!args || args.destination === undefined) {
                throw new Error("Missing required property 'destination'");
            }
            inputs["billingAccount"] = args ? args.billingAccount : undefined;
            inputs["destination"] = args ? args.destination : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["writerIdentity"] = undefined /*out*/;
        }
        super(BillingAccountSink.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BillingAccountSink resources.
 */
export interface BillingAccountSinkState {
    /**
     * The billing account exported to the sink.
     */
    readonly billingAccount?: pulumi.Input<string>;
    /**
     * The destination of the sink (or, in other words, where logs are written to). Can be a
     * Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * ```
     * The writer associated with the sink must have access to write to the above resource.
     */
    readonly destination?: pulumi.Input<string>;
    /**
     * The filter to apply when exporting logs. Only log entries that match the filter are exported.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    readonly filter?: pulumi.Input<string>;
    /**
     * The name of the logging sink.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The identity associated with this sink. This identity must be granted write access to the
     * configured `destination`.
     */
    readonly writerIdentity?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BillingAccountSink resource.
 */
export interface BillingAccountSinkArgs {
    /**
     * The billing account exported to the sink.
     */
    readonly billingAccount: pulumi.Input<string>;
    /**
     * The destination of the sink (or, in other words, where logs are written to). Can be a
     * Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * ```
     * The writer associated with the sink must have access to write to the above resource.
     */
    readonly destination: pulumi.Input<string>;
    /**
     * The filter to apply when exporting logs. Only log entries that match the filter are exported.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    readonly filter?: pulumi.Input<string>;
    /**
     * The name of the logging sink.
     */
    readonly name?: pulumi.Input<string>;
}
