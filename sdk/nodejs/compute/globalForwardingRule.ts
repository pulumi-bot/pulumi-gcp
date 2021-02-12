// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Represents a GlobalForwardingRule resource. Global forwarding rules are
 * used to forward traffic to the correct load balancer for HTTP load
 * balancing. Global forwarding rules can only be used for HTTP load
 * balancing.
 *
 * For more information, see
 * https://cloud.google.com/compute/docs/load-balancing/http/
 *
 * ## Example Usage
 * ### Global Forwarding Rule Http
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultHttpHealthCheck = new gcp.compute.HttpHealthCheck("defaultHttpHealthCheck", {
 *     requestPath: "/",
 *     checkIntervalSec: 1,
 *     timeoutSec: 1,
 * });
 * const defaultBackendService = new gcp.compute.BackendService("defaultBackendService", {
 *     portName: "http",
 *     protocol: "HTTP",
 *     timeoutSec: 10,
 *     healthChecks: [defaultHttpHealthCheck.id],
 * });
 * const defaultURLMap = new gcp.compute.URLMap("defaultURLMap", {
 *     description: "a description",
 *     defaultService: defaultBackendService.id,
 *     hostRules: [{
 *         hosts: ["mysite.com"],
 *         pathMatcher: "allpaths",
 *     }],
 *     pathMatchers: [{
 *         name: "allpaths",
 *         defaultService: defaultBackendService.id,
 *         pathRules: [{
 *             paths: ["/*"],
 *             service: defaultBackendService.id,
 *         }],
 *     }],
 * });
 * const defaultTargetHttpProxy = new gcp.compute.TargetHttpProxy("defaultTargetHttpProxy", {
 *     description: "a description",
 *     urlMap: defaultURLMap.id,
 * });
 * const defaultGlobalForwardingRule = new gcp.compute.GlobalForwardingRule("defaultGlobalForwardingRule", {
 *     target: defaultTargetHttpProxy.id,
 *     portRange: "80",
 * });
 * ```
 * ### Global Forwarding Rule Internal
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const debianImage = gcp.compute.getImage({
 *     family: "debian-9",
 *     project: "debian-cloud",
 * });
 * const instanceTemplate = new gcp.compute.InstanceTemplate("instanceTemplate", {
 *     machineType: "e2-medium",
 *     networkInterfaces: [{
 *         network: "default",
 *     }],
 *     disks: [{
 *         sourceImage: debianImage.then(debianImage => debianImage.selfLink),
 *         autoDelete: true,
 *         boot: true,
 *     }],
 * }, {
 *     provider: google_beta,
 * });
 * const igm = new gcp.compute.InstanceGroupManager("igm", {
 *     versions: [{
 *         instanceTemplate: instanceTemplate.id,
 *         name: "primary",
 *     }],
 *     baseInstanceName: "internal-glb",
 *     zone: "us-central1-f",
 *     targetSize: 1,
 * }, {
 *     provider: google_beta,
 * });
 * const defaultHealthCheck = new gcp.compute.HealthCheck("defaultHealthCheck", {
 *     checkIntervalSec: 1,
 *     timeoutSec: 1,
 *     tcpHealthCheck: {
 *         port: "80",
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * const defaultBackendService = new gcp.compute.BackendService("defaultBackendService", {
 *     portName: "http",
 *     protocol: "HTTP",
 *     timeoutSec: 10,
 *     loadBalancingScheme: "INTERNAL_SELF_MANAGED",
 *     backends: [{
 *         group: igm.instanceGroup,
 *         balancingMode: "RATE",
 *         capacityScaler: 0.4,
 *         maxRatePerInstance: 50,
 *     }],
 *     healthChecks: [defaultHealthCheck.id],
 * }, {
 *     provider: google_beta,
 * });
 * const defaultURLMap = new gcp.compute.URLMap("defaultURLMap", {
 *     description: "a description",
 *     defaultService: defaultBackendService.id,
 *     hostRules: [{
 *         hosts: ["mysite.com"],
 *         pathMatcher: "allpaths",
 *     }],
 *     pathMatchers: [{
 *         name: "allpaths",
 *         defaultService: defaultBackendService.id,
 *         pathRules: [{
 *             paths: ["/*"],
 *             service: defaultBackendService.id,
 *         }],
 *     }],
 * }, {
 *     provider: google_beta,
 * });
 * const defaultTargetHttpProxy = new gcp.compute.TargetHttpProxy("defaultTargetHttpProxy", {
 *     description: "a description",
 *     urlMap: defaultURLMap.id,
 * }, {
 *     provider: google_beta,
 * });
 * const defaultGlobalForwardingRule = new gcp.compute.GlobalForwardingRule("defaultGlobalForwardingRule", {
 *     target: defaultTargetHttpProxy.id,
 *     portRange: "80",
 *     loadBalancingScheme: "INTERNAL_SELF_MANAGED",
 *     ipAddress: "0.0.0.0",
 *     metadataFilters: [{
 *         filterMatchCriteria: "MATCH_ANY",
 *         filterLabels: [{
 *             name: "PLANET",
 *             value: "MARS",
 *         }],
 *     }],
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * GlobalForwardingRule can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/globalForwardingRule:GlobalForwardingRule default projects/{{project}}/global/forwardingRules/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/globalForwardingRule:GlobalForwardingRule default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/globalForwardingRule:GlobalForwardingRule default {{name}}
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
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GlobalForwardingRuleState, opts?: pulumi.CustomResourceOptions): GlobalForwardingRule {
        return new GlobalForwardingRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/globalForwardingRule:GlobalForwardingRule';

    /**
     * Returns true if the given object is an instance of GlobalForwardingRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GlobalForwardingRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GlobalForwardingRule.__pulumiType;
    }

    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The IP address that this forwarding rule is serving on behalf of.
     * Addresses are restricted based on the forwarding rule's load balancing
     * scheme (EXTERNAL or INTERNAL) and scope (global or regional).
     * When the load balancing scheme is EXTERNAL, for global forwarding
     * rules, the address must be a global IP, and for regional forwarding
     * rules, the address must live in the same region as the forwarding
     * rule. If this field is empty, an ephemeral IPv4 address from the same
     * scope (global or regional) will be assigned. A regional forwarding
     * rule supports IPv4 only. A global forwarding rule supports either IPv4
     * or IPv6.
     * When the load balancing scheme is INTERNAL, this can only be an RFC
     * 1918 IP address belonging to the network/subnet configured for the
     * forwarding rule. By default, if this field is empty, an ephemeral
     * internal IP address will be automatically allocated from the IP range
     * of the subnet or network configured for this forwarding rule.
     * An address must be specified by a literal IP address. > **NOTE**: While
     * the API allows you to specify various resource paths for an address resource
     * instead, this provider requires this to specifically be an IP address to
     * avoid needing to fetching the IP address from resource paths on refresh
     * or unnecessary diffs.
     */
    public readonly ipAddress!: pulumi.Output<string>;
    /**
     * The IP protocol to which this rule applies. When the load balancing scheme is
     * INTERNAL_SELF_MANAGED, only TCP is valid.
     * Possible values are `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, and `ICMP`.
     */
    public readonly ipProtocol!: pulumi.Output<string>;
    /**
     * The IP Version that will be used by this global forwarding rule.
     * Possible values are `IPV4` and `IPV6`.
     */
    public readonly ipVersion!: pulumi.Output<string | undefined>;
    /**
     * The fingerprint used for optimistic locking of this resource. Used internally during updates.
     */
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * Labels to apply to this forwarding rule.  A list of key->value pairs.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * This signifies what the GlobalForwardingRule will be used for.
     * The value of INTERNAL_SELF_MANAGED means that this will be used for
     * Internal Global HTTP(S) LB. The value of EXTERNAL means that this
     * will be used for External Global Load Balancing (HTTP(S) LB,
     * External TCP/UDP LB, SSL Proxy)
     * NOTE: Currently global forwarding rules cannot be used for INTERNAL
     * load balancing.
     * Default value is `EXTERNAL`.
     * Possible values are `EXTERNAL` and `INTERNAL_SELF_MANAGED`.
     */
    public readonly loadBalancingScheme!: pulumi.Output<string | undefined>;
    /**
     * Opaque filter criteria used by Loadbalancer to restrict routing
     * configuration to a limited set xDS compliant clients. In their xDS
     * requests to Loadbalancer, xDS clients present node metadata. If a
     * match takes place, the relevant routing configuration is made available
     * to those proxies.
     * For each metadataFilter in this list, if its filterMatchCriteria is set
     * to MATCH_ANY, at least one of the filterLabels must match the
     * corresponding label provided in the metadata. If its filterMatchCriteria
     * is set to MATCH_ALL, then all of its filterLabels must match with
     * corresponding labels in the provided metadata.
     * metadataFilters specified here can be overridden by those specified in
     * the UrlMap that this ForwardingRule references.
     * metadataFilters only applies to Loadbalancers that have their
     * loadBalancingScheme set to INTERNAL_SELF_MANAGED.
     * Structure is documented below.
     */
    public readonly metadataFilters!: pulumi.Output<outputs.compute.GlobalForwardingRuleMetadataFilter[] | undefined>;
    /**
     * Name of the metadata label. The length must be between
     * 1 and 1024 characters, inclusive.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * This field is not used for external load balancing.
     * For INTERNAL_SELF_MANAGED load balancing, this field
     * identifies the network that the load balanced IP should belong to
     * for this global forwarding rule. If this field is not specified,
     * the default network will be used.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * This field is used along with the target field for TargetHttpProxy,
     * TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
     * TargetPool, TargetInstance.
     * Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
     * addressed to ports in the specified range will be forwarded to target.
     * Forwarding rules with the same [IPAddress, IPProtocol] pair must have
     * disjoint port ranges.
     * Some types of forwarding target have constraints on the acceptable
     * ports:
     * * TargetHttpProxy: 80, 8080
     * * TargetHttpsProxy: 443
     * * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
     * 1883, 5222
     * * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
     * 1883, 5222
     * * TargetVpnGateway: 500, 4500
     */
    public readonly portRange!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The URL of the target resource to receive the matched traffic.
     * The forwarded traffic must be of a type appropriate to the target object.
     * For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
     * are valid.
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
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GlobalForwardingRuleState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["ipAddress"] = state ? state.ipAddress : undefined;
            inputs["ipProtocol"] = state ? state.ipProtocol : undefined;
            inputs["ipVersion"] = state ? state.ipVersion : undefined;
            inputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["loadBalancingScheme"] = state ? state.loadBalancingScheme : undefined;
            inputs["metadataFilters"] = state ? state.metadataFilters : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["portRange"] = state ? state.portRange : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["target"] = state ? state.target : undefined;
        } else {
            const args = argsOrState as GlobalForwardingRuleArgs | undefined;
            if ((!args || args.target === undefined) && !opts.urn) {
                throw new Error("Missing required property 'target'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["ipAddress"] = args ? args.ipAddress : undefined;
            inputs["ipProtocol"] = args ? args.ipProtocol : undefined;
            inputs["ipVersion"] = args ? args.ipVersion : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["loadBalancingScheme"] = args ? args.loadBalancingScheme : undefined;
            inputs["metadataFilters"] = args ? args.metadataFilters : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["portRange"] = args ? args.portRange : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["target"] = args ? args.target : undefined;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(GlobalForwardingRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GlobalForwardingRule resources.
 */
export interface GlobalForwardingRuleState {
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The IP address that this forwarding rule is serving on behalf of.
     * Addresses are restricted based on the forwarding rule's load balancing
     * scheme (EXTERNAL or INTERNAL) and scope (global or regional).
     * When the load balancing scheme is EXTERNAL, for global forwarding
     * rules, the address must be a global IP, and for regional forwarding
     * rules, the address must live in the same region as the forwarding
     * rule. If this field is empty, an ephemeral IPv4 address from the same
     * scope (global or regional) will be assigned. A regional forwarding
     * rule supports IPv4 only. A global forwarding rule supports either IPv4
     * or IPv6.
     * When the load balancing scheme is INTERNAL, this can only be an RFC
     * 1918 IP address belonging to the network/subnet configured for the
     * forwarding rule. By default, if this field is empty, an ephemeral
     * internal IP address will be automatically allocated from the IP range
     * of the subnet or network configured for this forwarding rule.
     * An address must be specified by a literal IP address. > **NOTE**: While
     * the API allows you to specify various resource paths for an address resource
     * instead, this provider requires this to specifically be an IP address to
     * avoid needing to fetching the IP address from resource paths on refresh
     * or unnecessary diffs.
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * The IP protocol to which this rule applies. When the load balancing scheme is
     * INTERNAL_SELF_MANAGED, only TCP is valid.
     * Possible values are `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, and `ICMP`.
     */
    readonly ipProtocol?: pulumi.Input<string>;
    /**
     * The IP Version that will be used by this global forwarding rule.
     * Possible values are `IPV4` and `IPV6`.
     */
    readonly ipVersion?: pulumi.Input<string>;
    /**
     * The fingerprint used for optimistic locking of this resource. Used internally during updates.
     */
    readonly labelFingerprint?: pulumi.Input<string>;
    /**
     * Labels to apply to this forwarding rule.  A list of key->value pairs.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * This signifies what the GlobalForwardingRule will be used for.
     * The value of INTERNAL_SELF_MANAGED means that this will be used for
     * Internal Global HTTP(S) LB. The value of EXTERNAL means that this
     * will be used for External Global Load Balancing (HTTP(S) LB,
     * External TCP/UDP LB, SSL Proxy)
     * NOTE: Currently global forwarding rules cannot be used for INTERNAL
     * load balancing.
     * Default value is `EXTERNAL`.
     * Possible values are `EXTERNAL` and `INTERNAL_SELF_MANAGED`.
     */
    readonly loadBalancingScheme?: pulumi.Input<string>;
    /**
     * Opaque filter criteria used by Loadbalancer to restrict routing
     * configuration to a limited set xDS compliant clients. In their xDS
     * requests to Loadbalancer, xDS clients present node metadata. If a
     * match takes place, the relevant routing configuration is made available
     * to those proxies.
     * For each metadataFilter in this list, if its filterMatchCriteria is set
     * to MATCH_ANY, at least one of the filterLabels must match the
     * corresponding label provided in the metadata. If its filterMatchCriteria
     * is set to MATCH_ALL, then all of its filterLabels must match with
     * corresponding labels in the provided metadata.
     * metadataFilters specified here can be overridden by those specified in
     * the UrlMap that this ForwardingRule references.
     * metadataFilters only applies to Loadbalancers that have their
     * loadBalancingScheme set to INTERNAL_SELF_MANAGED.
     * Structure is documented below.
     */
    readonly metadataFilters?: pulumi.Input<pulumi.Input<inputs.compute.GlobalForwardingRuleMetadataFilter>[]>;
    /**
     * Name of the metadata label. The length must be between
     * 1 and 1024 characters, inclusive.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * This field is not used for external load balancing.
     * For INTERNAL_SELF_MANAGED load balancing, this field
     * identifies the network that the load balanced IP should belong to
     * for this global forwarding rule. If this field is not specified,
     * the default network will be used.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * This field is used along with the target field for TargetHttpProxy,
     * TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
     * TargetPool, TargetInstance.
     * Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
     * addressed to ports in the specified range will be forwarded to target.
     * Forwarding rules with the same [IPAddress, IPProtocol] pair must have
     * disjoint port ranges.
     * Some types of forwarding target have constraints on the acceptable
     * ports:
     * * TargetHttpProxy: 80, 8080
     * * TargetHttpsProxy: 443
     * * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
     * 1883, 5222
     * * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
     * 1883, 5222
     * * TargetVpnGateway: 500, 4500
     */
    readonly portRange?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * The URL of the target resource to receive the matched traffic.
     * The forwarded traffic must be of a type appropriate to the target object.
     * For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
     * are valid.
     */
    readonly target?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GlobalForwardingRule resource.
 */
export interface GlobalForwardingRuleArgs {
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The IP address that this forwarding rule is serving on behalf of.
     * Addresses are restricted based on the forwarding rule's load balancing
     * scheme (EXTERNAL or INTERNAL) and scope (global or regional).
     * When the load balancing scheme is EXTERNAL, for global forwarding
     * rules, the address must be a global IP, and for regional forwarding
     * rules, the address must live in the same region as the forwarding
     * rule. If this field is empty, an ephemeral IPv4 address from the same
     * scope (global or regional) will be assigned. A regional forwarding
     * rule supports IPv4 only. A global forwarding rule supports either IPv4
     * or IPv6.
     * When the load balancing scheme is INTERNAL, this can only be an RFC
     * 1918 IP address belonging to the network/subnet configured for the
     * forwarding rule. By default, if this field is empty, an ephemeral
     * internal IP address will be automatically allocated from the IP range
     * of the subnet or network configured for this forwarding rule.
     * An address must be specified by a literal IP address. > **NOTE**: While
     * the API allows you to specify various resource paths for an address resource
     * instead, this provider requires this to specifically be an IP address to
     * avoid needing to fetching the IP address from resource paths on refresh
     * or unnecessary diffs.
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * The IP protocol to which this rule applies. When the load balancing scheme is
     * INTERNAL_SELF_MANAGED, only TCP is valid.
     * Possible values are `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, and `ICMP`.
     */
    readonly ipProtocol?: pulumi.Input<string>;
    /**
     * The IP Version that will be used by this global forwarding rule.
     * Possible values are `IPV4` and `IPV6`.
     */
    readonly ipVersion?: pulumi.Input<string>;
    /**
     * Labels to apply to this forwarding rule.  A list of key->value pairs.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * This signifies what the GlobalForwardingRule will be used for.
     * The value of INTERNAL_SELF_MANAGED means that this will be used for
     * Internal Global HTTP(S) LB. The value of EXTERNAL means that this
     * will be used for External Global Load Balancing (HTTP(S) LB,
     * External TCP/UDP LB, SSL Proxy)
     * NOTE: Currently global forwarding rules cannot be used for INTERNAL
     * load balancing.
     * Default value is `EXTERNAL`.
     * Possible values are `EXTERNAL` and `INTERNAL_SELF_MANAGED`.
     */
    readonly loadBalancingScheme?: pulumi.Input<string>;
    /**
     * Opaque filter criteria used by Loadbalancer to restrict routing
     * configuration to a limited set xDS compliant clients. In their xDS
     * requests to Loadbalancer, xDS clients present node metadata. If a
     * match takes place, the relevant routing configuration is made available
     * to those proxies.
     * For each metadataFilter in this list, if its filterMatchCriteria is set
     * to MATCH_ANY, at least one of the filterLabels must match the
     * corresponding label provided in the metadata. If its filterMatchCriteria
     * is set to MATCH_ALL, then all of its filterLabels must match with
     * corresponding labels in the provided metadata.
     * metadataFilters specified here can be overridden by those specified in
     * the UrlMap that this ForwardingRule references.
     * metadataFilters only applies to Loadbalancers that have their
     * loadBalancingScheme set to INTERNAL_SELF_MANAGED.
     * Structure is documented below.
     */
    readonly metadataFilters?: pulumi.Input<pulumi.Input<inputs.compute.GlobalForwardingRuleMetadataFilter>[]>;
    /**
     * Name of the metadata label. The length must be between
     * 1 and 1024 characters, inclusive.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * This field is not used for external load balancing.
     * For INTERNAL_SELF_MANAGED load balancing, this field
     * identifies the network that the load balanced IP should belong to
     * for this global forwarding rule. If this field is not specified,
     * the default network will be used.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * This field is used along with the target field for TargetHttpProxy,
     * TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
     * TargetPool, TargetInstance.
     * Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
     * addressed to ports in the specified range will be forwarded to target.
     * Forwarding rules with the same [IPAddress, IPProtocol] pair must have
     * disjoint port ranges.
     * Some types of forwarding target have constraints on the acceptable
     * ports:
     * * TargetHttpProxy: 80, 8080
     * * TargetHttpsProxy: 443
     * * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
     * 1883, 5222
     * * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
     * 1883, 5222
     * * TargetVpnGateway: 500, 4500
     */
    readonly portRange?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URL of the target resource to receive the matched traffic.
     * The forwarded traffic must be of a type appropriate to the target object.
     * For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
     * are valid.
     */
    readonly target: pulumi.Input<string>;
}
