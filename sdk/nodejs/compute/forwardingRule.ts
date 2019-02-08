// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A ForwardingRule resource. A ForwardingRule resource specifies which pool
 * of target virtual machines to forward a packet to if it matches the given
 * [IPAddress, IPProtocol, portRange] tuple.
 * 
 * 
 * To get more information about ForwardingRule, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/latest/forwardingRule)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/network/forwarding-rules)
 * 
 * <div class = "oics-button" style="float: right; margin: 0 0 -15px">
 *   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=forwarding_rule_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
 *     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
 *   </a>
 * </div>
 * ## Example Usage - Forwarding Rule Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const defaultTargetPool = new gcp.compute.TargetPool("default", {});
 * const defaultForwardingRule = new gcp.compute.ForwardingRule("default", {
 *     portRange: "80",
 *     target: defaultTargetPool.selfLink,
 * });
 * ```
 */
export class ForwardingRule extends pulumi.CustomResource {
    /**
     * Get an existing ForwardingRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ForwardingRuleState, opts?: pulumi.CustomResourceOptions): ForwardingRule {
        return new ForwardingRule(name, <any>state, { ...opts, id: id });
    }

    public readonly backendService: pulumi.Output<string | undefined>;
    public /*out*/ readonly creationTimestamp: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly ipAddress: pulumi.Output<string>;
    public readonly ipProtocol: pulumi.Output<string>;
    public readonly ipVersion: pulumi.Output<string | undefined>;
    public /*out*/ readonly labelFingerprint: pulumi.Output<string>;
    public readonly labels: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly loadBalancingScheme: pulumi.Output<string | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly network: pulumi.Output<string>;
    public readonly networkTier: pulumi.Output<string>;
    public readonly portRange: pulumi.Output<string | undefined>;
    public readonly ports: pulumi.Output<string[] | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    public readonly region: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    public readonly serviceLabel: pulumi.Output<string | undefined>;
    public /*out*/ readonly serviceName: pulumi.Output<string>;
    public readonly subnetwork: pulumi.Output<string>;
    public readonly target: pulumi.Output<string | undefined>;

    /**
     * Create a ForwardingRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ForwardingRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ForwardingRuleArgs | ForwardingRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ForwardingRuleState = argsOrState as ForwardingRuleState | undefined;
            inputs["backendService"] = state ? state.backendService : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["ipAddress"] = state ? state.ipAddress : undefined;
            inputs["ipProtocol"] = state ? state.ipProtocol : undefined;
            inputs["ipVersion"] = state ? state.ipVersion : undefined;
            inputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["loadBalancingScheme"] = state ? state.loadBalancingScheme : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["networkTier"] = state ? state.networkTier : undefined;
            inputs["portRange"] = state ? state.portRange : undefined;
            inputs["ports"] = state ? state.ports : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["serviceLabel"] = state ? state.serviceLabel : undefined;
            inputs["serviceName"] = state ? state.serviceName : undefined;
            inputs["subnetwork"] = state ? state.subnetwork : undefined;
            inputs["target"] = state ? state.target : undefined;
        } else {
            const args = argsOrState as ForwardingRuleArgs | undefined;
            inputs["backendService"] = args ? args.backendService : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["ipAddress"] = args ? args.ipAddress : undefined;
            inputs["ipProtocol"] = args ? args.ipProtocol : undefined;
            inputs["ipVersion"] = args ? args.ipVersion : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["loadBalancingScheme"] = args ? args.loadBalancingScheme : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["networkTier"] = args ? args.networkTier : undefined;
            inputs["portRange"] = args ? args.portRange : undefined;
            inputs["ports"] = args ? args.ports : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["serviceLabel"] = args ? args.serviceLabel : undefined;
            inputs["subnetwork"] = args ? args.subnetwork : undefined;
            inputs["target"] = args ? args.target : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["serviceName"] = undefined /*out*/;
        }
        super("gcp:compute/forwardingRule:ForwardingRule", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ForwardingRule resources.
 */
export interface ForwardingRuleState {
    readonly backendService?: pulumi.Input<string>;
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly ipAddress?: pulumi.Input<string>;
    readonly ipProtocol?: pulumi.Input<string>;
    readonly ipVersion?: pulumi.Input<string>;
    readonly labelFingerprint?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly loadBalancingScheme?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly network?: pulumi.Input<string>;
    readonly networkTier?: pulumi.Input<string>;
    readonly portRange?: pulumi.Input<string>;
    readonly ports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly serviceLabel?: pulumi.Input<string>;
    readonly serviceName?: pulumi.Input<string>;
    readonly subnetwork?: pulumi.Input<string>;
    readonly target?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ForwardingRule resource.
 */
export interface ForwardingRuleArgs {
    readonly backendService?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly ipAddress?: pulumi.Input<string>;
    readonly ipProtocol?: pulumi.Input<string>;
    readonly ipVersion?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly loadBalancingScheme?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly network?: pulumi.Input<string>;
    readonly networkTier?: pulumi.Input<string>;
    readonly portRange?: pulumi.Input<string>;
    readonly ports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly serviceLabel?: pulumi.Input<string>;
    readonly subnetwork?: pulumi.Input<string>;
    readonly target?: pulumi.Input<string>;
}
