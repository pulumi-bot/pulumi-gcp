// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * #     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
 * #
 * # ----------------------------------------------------------------------------
 * #
 * #     This file is automatically generated by Magic Modules and manual
 * #     changes will be clobbered when the file is regenerated.
 * #
 * #     Please read more about how to change this file in
 * #     .github/CONTRIBUTING.md.
 * #
 * # ----------------------------------------------------------------------------
 * layout: "google"
 * page_title: "Google: google_compute_region_autoscaler"
 * sidebar_current: "docs-google-compute-region-autoscaler"
 * description: |-
 *   Represents an Autoscaler resource.
 * ---
 * 
 * # google\_compute\_region\_autoscaler
 * 
 * Represents an Autoscaler resource.
 * 
 * Autoscalers allow you to automatically scale virtual machine instances in
 * managed instance groups according to an autoscaling policy that you
 * define.
 * 
 * 
 * To get more information about RegionAutoscaler, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/regionAutoscalers)
 * * How-to Guides
 *     * [Autoscaling Groups of Instances](https://cloud.google.com/compute/docs/autoscaler/)
 */
export class RegionAutoscaler extends pulumi.CustomResource {
    /**
     * Get an existing RegionAutoscaler resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegionAutoscalerState, opts?: pulumi.CustomResourceOptions): RegionAutoscaler {
        return new RegionAutoscaler(name, <any>state, { ...opts, id: id });
    }

    public readonly autoscalingPolicy: pulumi.Output<{ cooldownPeriod?: number, cpuUtilization: { target: number }, loadBalancingUtilization?: { target: number }, maxReplicas: number, metrics?: { name: string, target: number, type: string }[], minReplicas: number }>;
    public /*out*/ readonly creationTimestamp: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly project: pulumi.Output<string>;
    public readonly region: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    public readonly target: pulumi.Output<string>;

    /**
     * Create a RegionAutoscaler resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionAutoscalerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegionAutoscalerArgs | RegionAutoscalerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: RegionAutoscalerState = argsOrState as RegionAutoscalerState | undefined;
            inputs["autoscalingPolicy"] = state ? state.autoscalingPolicy : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["target"] = state ? state.target : undefined;
        } else {
            const args = argsOrState as RegionAutoscalerArgs | undefined;
            if (!args || args.autoscalingPolicy === undefined) {
                throw new Error("Missing required property 'autoscalingPolicy'");
            }
            if (!args || args.target === undefined) {
                throw new Error("Missing required property 'target'");
            }
            inputs["autoscalingPolicy"] = args ? args.autoscalingPolicy : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["target"] = args ? args.target : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        super("gcp:compute/regionAutoscaler:RegionAutoscaler", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegionAutoscaler resources.
 */
export interface RegionAutoscalerState {
    readonly autoscalingPolicy?: pulumi.Input<{ cooldownPeriod?: pulumi.Input<number>, cpuUtilization?: pulumi.Input<{ target: pulumi.Input<number> }>, loadBalancingUtilization?: pulumi.Input<{ target: pulumi.Input<number> }>, maxReplicas: pulumi.Input<number>, metrics?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, target: pulumi.Input<number>, type: pulumi.Input<string> }>[]>, minReplicas: pulumi.Input<number> }>;
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly target?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RegionAutoscaler resource.
 */
export interface RegionAutoscalerArgs {
    readonly autoscalingPolicy: pulumi.Input<{ cooldownPeriod?: pulumi.Input<number>, cpuUtilization?: pulumi.Input<{ target: pulumi.Input<number> }>, loadBalancingUtilization?: pulumi.Input<{ target: pulumi.Input<number> }>, maxReplicas: pulumi.Input<number>, metrics?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, target: pulumi.Input<number>, type: pulumi.Input<string> }>[]>, minReplicas: pulumi.Input<number> }>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly target: pulumi.Input<string>;
}
