// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A description of the conditions under which some aspect of your system is
 * considered to be "unhealthy" and the ways to notify people or services
 * about this state.
 * 
 * 
 * To get more information about AlertPolicy, see:
 * 
 * * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.alertPolicies)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/monitoring/alerts/)
 * 
 * ## Example Usage - Monitoring Alert Policy Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const alertPolicy = new gcp.monitoring.AlertPolicy("alert_policy", {
 *     combiner: "OR",
 *     conditions: [{
 *         conditionThreshold: {
 *             aggregations: [{
 *                 alignmentPeriod: "60s",
 *                 perSeriesAligner: "ALIGN_RATE",
 *             }],
 *             comparison: "COMPARISON_GT",
 *             duration: "60s",
 *             filter: "metric.type=\"compute.googleapis.com/instance/disk/write_bytes_count\" AND resource.type=\"gce_instance\"",
 *         },
 *         displayName: "test condition",
 *     }],
 *     displayName: "My Alert Policy",
 * });
 * ```
 */
export class AlertPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AlertPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlertPolicyState, opts?: pulumi.CustomResourceOptions): AlertPolicy {
        return new AlertPolicy(name, <any>state, { ...opts, id: id });
    }

    public readonly combiner!: pulumi.Output<string>;
    public readonly conditions!: pulumi.Output<{ conditionAbsent?: { aggregations?: { alignmentPeriod?: string, crossSeriesReducer?: string, groupByFields?: string[], perSeriesAligner?: string }[], duration: string, filter?: string, trigger?: { count?: number, percent?: number } }, conditionThreshold?: { aggregations?: { alignmentPeriod?: string, crossSeriesReducer?: string, groupByFields?: string[], perSeriesAligner?: string }[], comparison: string, denominatorAggregations?: { alignmentPeriod?: string, crossSeriesReducer?: string, groupByFields?: string[], perSeriesAligner?: string }[], denominatorFilter?: string, duration: string, filter?: string, thresholdValue?: number, trigger?: { count?: number, percent?: number } }, displayName: string, name: string }[]>;
    public /*out*/ readonly creationRecord!: pulumi.Output<{ mutateTime: string, mutatedBy: string }>;
    public readonly displayName!: pulumi.Output<string>;
    public readonly documentation!: pulumi.Output<{ content?: string, mimeType?: string } | undefined>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly labels!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly notificationChannels!: pulumi.Output<string[] | undefined>;
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a AlertPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlertPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlertPolicyArgs | AlertPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AlertPolicyState | undefined;
            inputs["combiner"] = state ? state.combiner : undefined;
            inputs["conditions"] = state ? state.conditions : undefined;
            inputs["creationRecord"] = state ? state.creationRecord : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["documentation"] = state ? state.documentation : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notificationChannels"] = state ? state.notificationChannels : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as AlertPolicyArgs | undefined;
            if (!args || args.combiner === undefined) {
                throw new Error("Missing required property 'combiner'");
            }
            if (!args || args.conditions === undefined) {
                throw new Error("Missing required property 'conditions'");
            }
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            inputs["combiner"] = args ? args.combiner : undefined;
            inputs["conditions"] = args ? args.conditions : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["documentation"] = args ? args.documentation : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["notificationChannels"] = args ? args.notificationChannels : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["creationRecord"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        super("gcp:monitoring/alertPolicy:AlertPolicy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AlertPolicy resources.
 */
export interface AlertPolicyState {
    readonly combiner?: pulumi.Input<string>;
    readonly conditions?: pulumi.Input<pulumi.Input<{ conditionAbsent?: pulumi.Input<{ aggregations?: pulumi.Input<pulumi.Input<{ alignmentPeriod?: pulumi.Input<string>, crossSeriesReducer?: pulumi.Input<string>, groupByFields?: pulumi.Input<pulumi.Input<string>[]>, perSeriesAligner?: pulumi.Input<string> }>[]>, duration: pulumi.Input<string>, filter?: pulumi.Input<string>, trigger?: pulumi.Input<{ count?: pulumi.Input<number>, percent?: pulumi.Input<number> }> }>, conditionThreshold?: pulumi.Input<{ aggregations?: pulumi.Input<pulumi.Input<{ alignmentPeriod?: pulumi.Input<string>, crossSeriesReducer?: pulumi.Input<string>, groupByFields?: pulumi.Input<pulumi.Input<string>[]>, perSeriesAligner?: pulumi.Input<string> }>[]>, comparison: pulumi.Input<string>, denominatorAggregations?: pulumi.Input<pulumi.Input<{ alignmentPeriod?: pulumi.Input<string>, crossSeriesReducer?: pulumi.Input<string>, groupByFields?: pulumi.Input<pulumi.Input<string>[]>, perSeriesAligner?: pulumi.Input<string> }>[]>, denominatorFilter?: pulumi.Input<string>, duration: pulumi.Input<string>, filter?: pulumi.Input<string>, thresholdValue?: pulumi.Input<number>, trigger?: pulumi.Input<{ count?: pulumi.Input<number>, percent?: pulumi.Input<number> }> }>, displayName: pulumi.Input<string>, name?: pulumi.Input<string> }>[]>;
    readonly creationRecord?: pulumi.Input<{ mutateTime?: pulumi.Input<string>, mutatedBy?: pulumi.Input<string> }>;
    readonly displayName?: pulumi.Input<string>;
    readonly documentation?: pulumi.Input<{ content?: pulumi.Input<string>, mimeType?: pulumi.Input<string> }>;
    readonly enabled?: pulumi.Input<boolean>;
    readonly labels?: pulumi.Input<pulumi.Input<string>[]>;
    readonly name?: pulumi.Input<string>;
    readonly notificationChannels?: pulumi.Input<pulumi.Input<string>[]>;
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AlertPolicy resource.
 */
export interface AlertPolicyArgs {
    readonly combiner: pulumi.Input<string>;
    readonly conditions: pulumi.Input<pulumi.Input<{ conditionAbsent?: pulumi.Input<{ aggregations?: pulumi.Input<pulumi.Input<{ alignmentPeriod?: pulumi.Input<string>, crossSeriesReducer?: pulumi.Input<string>, groupByFields?: pulumi.Input<pulumi.Input<string>[]>, perSeriesAligner?: pulumi.Input<string> }>[]>, duration: pulumi.Input<string>, filter?: pulumi.Input<string>, trigger?: pulumi.Input<{ count?: pulumi.Input<number>, percent?: pulumi.Input<number> }> }>, conditionThreshold?: pulumi.Input<{ aggregations?: pulumi.Input<pulumi.Input<{ alignmentPeriod?: pulumi.Input<string>, crossSeriesReducer?: pulumi.Input<string>, groupByFields?: pulumi.Input<pulumi.Input<string>[]>, perSeriesAligner?: pulumi.Input<string> }>[]>, comparison: pulumi.Input<string>, denominatorAggregations?: pulumi.Input<pulumi.Input<{ alignmentPeriod?: pulumi.Input<string>, crossSeriesReducer?: pulumi.Input<string>, groupByFields?: pulumi.Input<pulumi.Input<string>[]>, perSeriesAligner?: pulumi.Input<string> }>[]>, denominatorFilter?: pulumi.Input<string>, duration: pulumi.Input<string>, filter?: pulumi.Input<string>, thresholdValue?: pulumi.Input<number>, trigger?: pulumi.Input<{ count?: pulumi.Input<number>, percent?: pulumi.Input<number> }> }>, displayName: pulumi.Input<string>, name?: pulumi.Input<string> }>[]>;
    readonly displayName: pulumi.Input<string>;
    readonly documentation?: pulumi.Input<{ content?: pulumi.Input<string>, mimeType?: pulumi.Input<string> }>;
    readonly enabled?: pulumi.Input<boolean>;
    readonly labels?: pulumi.Input<pulumi.Input<string>[]>;
    readonly notificationChannels?: pulumi.Input<pulumi.Input<string>[]>;
    readonly project?: pulumi.Input<string>;
}
