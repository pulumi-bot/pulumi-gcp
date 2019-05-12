// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Global Forwarding Rule within GCE. This binds an ip and port to a target HTTP(s) proxy. For more
 * information see [the official
 * documentation](https://cloud.google.com/compute/docs/load-balancing/http/global-forwarding-rules) and
 * [API](https://cloud.google.com/compute/docs/reference/latest/globalForwardingRules).
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const defaultHttpHealthCheck = new gcp.compute.HttpHealthCheck("default", {
 *     checkIntervalSec: 1,
 *     requestPath: "/",
 *     timeoutSec: 1,
 * });
 * const defaultBackendService = new gcp.compute.BackendService("default", {
 *     healthChecks: defaultHttpHealthCheck.selfLink,
 *     portName: "http",
 *     protocol: "HTTP",
 *     timeoutSec: 10,
 * });
 * const defaultURLMap = new gcp.compute.URLMap("default", {
 *     defaultService: defaultBackendService.selfLink,
 *     description: "a description",
 *     hostRules: [{
 *         hosts: ["mysite.com"],
 *         pathMatcher: "allpaths",
 *     }],
 *     pathMatchers: [{
 *         defaultService: defaultBackendService.selfLink,
 *         name: "allpaths",
 *         pathRules: [{
 *             paths: ["/*"],
 *             service: defaultBackendService.selfLink,
 *         }],
 *     }],
 * });
 * const defaultTargetHttpProxy = new gcp.compute.TargetHttpProxy("default", {
 *     description: "a description",
 *     urlMap: defaultURLMap.selfLink,
 * });
 * const defaultGlobalForwardingRule = new gcp.compute.GlobalForwardingRule("default", {
 *     portRange: "80",
 *     target: defaultTargetHttpProxy.selfLink,
 * });
 * ```
 */
export class GlobalForwardingRule extends pulumi.CustomResource {
    /**
     * Get an existing GlobalForwardingRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GlobalForwardingRuleState, opts?: pulumi.CustomResourceOptions): GlobalForwardingRule {
        return new GlobalForwardingRule(name, <any>state, { ...opts, id: id });
    }

    /**
     * Textual description field.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The static IP. (if not set, an ephemeral IP is
     * used). This should be the literal IP address to be used, not the `self_link`
     * to a `google_compute_global_address` resource. (If using a `google_compute_global_address`
     * resource, use the `address` property instead of the `self_link` property.)
     */
    public readonly ipAddress!: pulumi.Output<string>;
    /**
     * The IP protocol to route, one of "TCP" "UDP" "AH"
     * "ESP" or "SCTP". (default "TCP").
     */
    public readonly ipProtocol!: pulumi.Output<string>;
    /**
     * 
     * The IP Version that will be used by this resource's address. One of `"IPV4"` or `"IPV6"`.
     * You cannot provide this and `ip_address`.
     */
    public readonly ipVersion!: pulumi.Output<string | undefined>;
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * )
     * A set of key/value label pairs to assign to the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A unique name for the resource, required by GCE. Changing
     * this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A range e.g. "1024-2048" or a single port "1024"
     * (defaults to all ports!).
     * Some types of forwarding targets have constraints on the acceptable ports:
     * * Target HTTP proxy: 80, 8080
     * * Target HTTPS proxy: 443
     * * Target TCP proxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222
     * * Target SSL proxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222
     * * Target VPN gateway: 500, 4500
     */
    public readonly portRange!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * URL of target HTTP or HTTPS proxy.
     */
    public readonly target!: pulumi.Output<string>;

    /**
     * Create a GlobalForwardingRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GlobalForwardingRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GlobalForwardingRuleArgs | GlobalForwardingRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GlobalForwardingRuleState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["ipAddress"] = state ? state.ipAddress : undefined;
            inputs["ipProtocol"] = state ? state.ipProtocol : undefined;
            inputs["ipVersion"] = state ? state.ipVersion : undefined;
            inputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["portRange"] = state ? state.portRange : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["target"] = state ? state.target : undefined;
        } else {
            const args = argsOrState as GlobalForwardingRuleArgs | undefined;
            if (!args || args.target === undefined) {
                throw new Error("Missing required property 'target'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["ipAddress"] = args ? args.ipAddress : undefined;
            inputs["ipProtocol"] = args ? args.ipProtocol : undefined;
            inputs["ipVersion"] = args ? args.ipVersion : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["portRange"] = args ? args.portRange : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["target"] = args ? args.target : undefined;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super("gcp:compute/globalForwardingRule:GlobalForwardingRule", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GlobalForwardingRule resources.
 */
export interface GlobalForwardingRuleState {
    /**
     * Textual description field.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The static IP. (if not set, an ephemeral IP is
     * used). This should be the literal IP address to be used, not the `self_link`
     * to a `google_compute_global_address` resource. (If using a `google_compute_global_address`
     * resource, use the `address` property instead of the `self_link` property.)
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * The IP protocol to route, one of "TCP" "UDP" "AH"
     * "ESP" or "SCTP". (default "TCP").
     */
    readonly ipProtocol?: pulumi.Input<string>;
    /**
     * 
     * The IP Version that will be used by this resource's address. One of `"IPV4"` or `"IPV6"`.
     * You cannot provide this and `ip_address`.
     */
    readonly ipVersion?: pulumi.Input<string>;
    readonly labelFingerprint?: pulumi.Input<string>;
    /**
     * )
     * A set of key/value label pairs to assign to the resource.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique name for the resource, required by GCE. Changing
     * this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A range e.g. "1024-2048" or a single port "1024"
     * (defaults to all ports!).
     * Some types of forwarding targets have constraints on the acceptable ports:
     * * Target HTTP proxy: 80, 8080
     * * Target HTTPS proxy: 443
     * * Target TCP proxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222
     * * Target SSL proxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222
     * * Target VPN gateway: 500, 4500
     */
    readonly portRange?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * URL of target HTTP or HTTPS proxy.
     */
    readonly target?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GlobalForwardingRule resource.
 */
export interface GlobalForwardingRuleArgs {
    /**
     * Textual description field.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The static IP. (if not set, an ephemeral IP is
     * used). This should be the literal IP address to be used, not the `self_link`
     * to a `google_compute_global_address` resource. (If using a `google_compute_global_address`
     * resource, use the `address` property instead of the `self_link` property.)
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * The IP protocol to route, one of "TCP" "UDP" "AH"
     * "ESP" or "SCTP". (default "TCP").
     */
    readonly ipProtocol?: pulumi.Input<string>;
    /**
     * 
     * The IP Version that will be used by this resource's address. One of `"IPV4"` or `"IPV6"`.
     * You cannot provide this and `ip_address`.
     */
    readonly ipVersion?: pulumi.Input<string>;
    /**
     * )
     * A set of key/value label pairs to assign to the resource.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique name for the resource, required by GCE. Changing
     * this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A range e.g. "1024-2048" or a single port "1024"
     * (defaults to all ports!).
     * Some types of forwarding targets have constraints on the acceptable ports:
     * * Target HTTP proxy: 80, 8080
     * * Target HTTPS proxy: 443
     * * Target TCP proxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222
     * * Target SSL proxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222
     * * Target VPN gateway: 500, 4500
     */
    readonly portRange?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * URL of target HTTP or HTTPS proxy.
     */
    readonly target: pulumi.Input<string>;
}
