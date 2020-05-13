// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A ForwardingRule resource. A ForwardingRule resource specifies which pool
 * of target virtual machines to forward a packet to if it matches the given
 * [IPAddress, IPProtocol, portRange] tuple.
 * 
 * 
 * To get more information about ForwardingRule, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/v1/forwardingRules)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/network/forwarding-rules)
 * 
 * ## Example Usage - Forwarding Rule Global Internallb
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const hc = new gcp.compute.HealthCheck("hc", {
 *     checkIntervalSec: 1,
 *     timeoutSec: 1,
 *     tcp_health_check: {
 *         port: "80",
 *     },
 * });
 * const backend = new gcp.compute.RegionBackendService("backend", {
 *     region: "us-central1",
 *     healthChecks: [hc.id],
 * });
 * const defaultNetwork = new gcp.compute.Network("defaultNetwork", {autoCreateSubnetworks: false});
 * const defaultSubnetwork = new gcp.compute.Subnetwork("defaultSubnetwork", {
 *     ipCidrRange: "10.0.0.0/16",
 *     region: "us-central1",
 *     network: defaultNetwork.id,
 * });
 * // Forwarding rule for Internal Load Balancing
 * const defaultForwardingRule = new gcp.compute.ForwardingRule("defaultForwardingRule", {
 *     region: "us-central1",
 *     loadBalancingScheme: "INTERNAL",
 *     backendService: backend.id,
 *     allPorts: true,
 *     allowGlobalAccess: true,
 *     network: defaultNetwork.name,
 *     subnetwork: defaultSubnetwork.name,
 * });
 * ```
 * ## Example Usage - Forwarding Rule Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const defaultTargetPool = new gcp.compute.TargetPool("defaultTargetPool", {});
 * const defaultForwardingRule = new gcp.compute.ForwardingRule("defaultForwardingRule", {
 *     target: defaultTargetPool.id,
 *     portRange: "80",
 * });
 * ```
 * ## Example Usage - Forwarding Rule Internallb
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const hc = new gcp.compute.HealthCheck("hc", {
 *     checkIntervalSec: 1,
 *     timeoutSec: 1,
 *     tcp_health_check: {
 *         port: "80",
 *     },
 * });
 * const backend = new gcp.compute.RegionBackendService("backend", {
 *     region: "us-central1",
 *     healthChecks: [hc.id],
 * });
 * const defaultNetwork = new gcp.compute.Network("defaultNetwork", {autoCreateSubnetworks: false});
 * const defaultSubnetwork = new gcp.compute.Subnetwork("defaultSubnetwork", {
 *     ipCidrRange: "10.0.0.0/16",
 *     region: "us-central1",
 *     network: defaultNetwork.id,
 * });
 * // Forwarding rule for Internal Load Balancing
 * const defaultForwardingRule = new gcp.compute.ForwardingRule("defaultForwardingRule", {
 *     region: "us-central1",
 *     loadBalancingScheme: "INTERNAL",
 *     backendService: backend.id,
 *     allPorts: true,
 *     network: defaultNetwork.name,
 *     subnetwork: defaultSubnetwork.name,
 * });
 * ```
 * ## Example Usage - Forwarding Rule Http Lb
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const debianImage = gcp.compute.getImage({
 *     family: "debian-9",
 *     project: "debian-cloud",
 * });
 * const defaultNetwork = new gcp.compute.Network("defaultNetwork", {
 *     autoCreateSubnetworks: false,
 *     routingMode: "REGIONAL",
 * });
 * const defaultSubnetwork = new gcp.compute.Subnetwork("defaultSubnetwork", {
 *     ipCidrRange: "10.1.2.0/24",
 *     region: "us-central1",
 *     network: defaultNetwork.id,
 * });
 * const instanceTemplate = new gcp.compute.InstanceTemplate("instanceTemplate", {
 *     machineType: "n1-standard-1",
 *     network_interface: [{
 *         network: defaultNetwork.id,
 *         subnetwork: defaultSubnetwork.id,
 *     }],
 *     disk: [{
 *         sourceImage: debianImage.then(debianImage => debianImage.selfLink),
 *         autoDelete: true,
 *         boot: true,
 *     }],
 *     tags: [
 *         "allow-ssh",
 *         "load-balanced-backend",
 *     ],
 * });
 * const rigm = new gcp.compute.RegionInstanceGroupManager("rigm", {
 *     region: "us-central1",
 *     version: [{
 *         instanceTemplate: instanceTemplate.selfLink,
 *         name: "primary",
 *     }],
 *     baseInstanceName: "internal-glb",
 *     targetSize: 1,
 * });
 * const fw1 = new gcp.compute.Firewall("fw1", {
 *     network: defaultNetwork.id,
 *     sourceRanges: ["10.1.2.0/24"],
 *     allow: [
 *         {
 *             protocol: "tcp",
 *         },
 *         {
 *             protocol: "udp",
 *         },
 *         {
 *             protocol: "icmp",
 *         },
 *     ],
 *     direction: "INGRESS",
 * });
 * const fw2 = new gcp.compute.Firewall("fw2", {
 *     network: defaultNetwork.id,
 *     sourceRanges: ["0.0.0.0/0"],
 *     allow: [{
 *         protocol: "tcp",
 *         ports: ["22"],
 *     }],
 *     targetTags: ["allow-ssh"],
 *     direction: "INGRESS",
 * });
 * const fw3 = new gcp.compute.Firewall("fw3", {
 *     network: defaultNetwork.id,
 *     sourceRanges: [
 *         "130.211.0.0/22",
 *         "35.191.0.0/16",
 *     ],
 *     allow: [{
 *         protocol: "tcp",
 *     }],
 *     targetTags: ["load-balanced-backend"],
 *     direction: "INGRESS",
 * });
 * const fw4 = new gcp.compute.Firewall("fw4", {
 *     network: defaultNetwork.id,
 *     sourceRanges: ["10.129.0.0/26"],
 *     targetTags: ["load-balanced-backend"],
 *     allow: [
 *         {
 *             protocol: "tcp",
 *             ports: ["80"],
 *         },
 *         {
 *             protocol: "tcp",
 *             ports: ["443"],
 *         },
 *         {
 *             protocol: "tcp",
 *             ports: ["8000"],
 *         },
 *     ],
 *     direction: "INGRESS",
 * });
 * const defaultRegionHealthCheck = new gcp.compute.RegionHealthCheck("defaultRegionHealthCheck", {
 *     region: "us-central1",
 *     http_health_check: {
 *         portSpecification: "USE_SERVING_PORT",
 *     },
 * });
 * const defaultRegionBackendService = new gcp.compute.RegionBackendService("defaultRegionBackendService", {
 *     loadBalancingScheme: "INTERNAL_MANAGED",
 *     backend: [{
 *         group: rigm.instanceGroup,
 *         balancingMode: "UTILIZATION",
 *         capacityScaler: 1,
 *     }],
 *     region: "us-central1",
 *     protocol: "HTTP",
 *     timeoutSec: 10,
 *     healthChecks: [defaultRegionHealthCheck.id],
 * });
 * const defaultRegionUrlMap = new gcp.compute.RegionUrlMap("defaultRegionUrlMap", {
 *     region: "us-central1",
 *     defaultService: defaultRegionBackendService.id,
 * });
 * const defaultRegionTargetHttpProxy = new gcp.compute.RegionTargetHttpProxy("defaultRegionTargetHttpProxy", {
 *     region: "us-central1",
 *     urlMap: defaultRegionUrlMap.id,
 * });
 * const proxy = new gcp.compute.Subnetwork("proxy", {
 *     ipCidrRange: "10.129.0.0/26",
 *     region: "us-central1",
 *     network: defaultNetwork.id,
 *     purpose: "INTERNAL_HTTPS_LOAD_BALANCER",
 *     role: "ACTIVE",
 * });
 * // Forwarding rule for Internal Load Balancing
 * const defaultForwardingRule = new gcp.compute.ForwardingRule("defaultForwardingRule", {
 *     region: "us-central1",
 *     ipProtocol: "TCP",
 *     loadBalancingScheme: "INTERNAL_MANAGED",
 *     portRange: "80",
 *     target: defaultRegionTargetHttpProxy.id,
 *     network: defaultNetwork.id,
 *     subnetwork: defaultSubnetwork.id,
 *     networkTier: "PREMIUM",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_forwarding_rule.html.markdown.
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

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/forwardingRule:ForwardingRule';

    /**
     * Returns true if the given object is an instance of ForwardingRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ForwardingRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ForwardingRule.__pulumiType;
    }

    /**
     * For internal TCP/UDP load balancing (i.e. load balancing scheme is
     * INTERNAL and protocol is TCP/UDP), set this to true to allow packets
     * addressed to any ports to be forwarded to the backends configured
     * with this forwarding rule. Used with backend service. Cannot be set
     * if port or portRange are set.
     */
    public readonly allPorts!: pulumi.Output<boolean | undefined>;
    /**
     * If true, clients can access ILB from all regions.
     * Otherwise only allows from the local region the ILB is located at.
     */
    public readonly allowGlobalAccess!: pulumi.Output<boolean | undefined>;
    /**
     * A BackendService to receive the matched traffic. This is used only
     * for INTERNAL load balancing.
     */
    public readonly backendService!: pulumi.Output<string | undefined>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
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
     * The IP protocol to which this rule applies.
     * When the load balancing scheme is INTERNAL, only TCP and UDP are
     * valid.
     */
    public readonly ipProtocol!: pulumi.Output<string>;
    /**
     * Indicates whether or not this load balancer can be used
     * as a collector for packet mirroring. To prevent mirroring loops,
     * instances behind this load balancer will not have their traffic
     * mirrored even if a PacketMirroring rule applies to them. This
     * can only be set to true for load balancers that have their
     * loadBalancingScheme set to INTERNAL.
     */
    public readonly isMirroringCollector!: pulumi.Output<boolean | undefined>;
    /**
     * The fingerprint used for optimistic locking of this resource. Used internally during updates.
     */
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * Labels to apply to this forwarding rule.  A list of key->value pairs.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * This signifies what the ForwardingRule will be used for and can be
     * EXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is used for Classic
     * Cloud VPN gateways, protocol forwarding to VMs from an external IP address,
     * and HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP load balancers.
     * INTERNAL is used for protocol forwarding to VMs from an internal IP address,
     * and internal TCP/UDP load balancers.
     * INTERNAL_MANAGED is used for internal HTTP(S) load balancers.
     */
    public readonly loadBalancingScheme!: pulumi.Output<string | undefined>;
    /**
     * Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * For internal load balancing, this field identifies the network that
     * the load balanced IP should belong to for this Forwarding Rule. If
     * this field is not specified, the default network will be used.
     * This field is only used for INTERNAL load balancing.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * The networking tier used for configuring this address. If this field is not
     * specified, it is assumed to be PREMIUM.
     */
    public readonly networkTier!: pulumi.Output<string>;
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
     * This field is used along with the backendService field for internal
     * load balancing.
     * When the load balancing scheme is INTERNAL, a single port or a comma
     * separated list of ports can be configured. Only packets addressed to
     * these ports will be forwarded to the backends configured with this
     * forwarding rule.
     * You may specify a maximum of up to 5 ports.
     */
    public readonly ports!: pulumi.Output<string[] | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * A reference to the region where the regional forwarding rule resides.
     * This field is not applicable to global forwarding rules.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * An optional prefix to the service name for this Forwarding Rule.
     * If specified, will be the first label of the fully qualified service
     * name.
     * The label must be 1-63 characters long, and comply with RFC1035.
     * Specifically, the label must be 1-63 characters long and match the
     * regular expression `a-z?` which means the first
     * character must be a lowercase letter, and all following characters
     * must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     * This field is only used for INTERNAL load balancing.
     */
    public readonly serviceLabel!: pulumi.Output<string | undefined>;
    /**
     * The internal fully qualified service name for this Forwarding Rule. This field is only used for INTERNAL load
     * balancing.
     */
    public /*out*/ readonly serviceName!: pulumi.Output<string>;
    /**
     * The subnetwork that the load balanced IP should belong to for this
     * Forwarding Rule.  This field is only used for INTERNAL load balancing.
     * If the network specified is in auto subnet mode, this field is
     * optional. However, if the network is in custom subnet mode, a
     * subnetwork must be specified.
     */
    public readonly subnetwork!: pulumi.Output<string>;
    /**
     * The URL of the target resource to receive the matched traffic.
     * The target must live in the same region as the forwarding rule.
     * The forwarded traffic must be of a type appropriate to the target
     * object.
     */
    public readonly target!: pulumi.Output<string | undefined>;

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
            const state = argsOrState as ForwardingRuleState | undefined;
            inputs["allPorts"] = state ? state.allPorts : undefined;
            inputs["allowGlobalAccess"] = state ? state.allowGlobalAccess : undefined;
            inputs["backendService"] = state ? state.backendService : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["ipAddress"] = state ? state.ipAddress : undefined;
            inputs["ipProtocol"] = state ? state.ipProtocol : undefined;
            inputs["isMirroringCollector"] = state ? state.isMirroringCollector : undefined;
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
            inputs["allPorts"] = args ? args.allPorts : undefined;
            inputs["allowGlobalAccess"] = args ? args.allowGlobalAccess : undefined;
            inputs["backendService"] = args ? args.backendService : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["ipAddress"] = args ? args.ipAddress : undefined;
            inputs["ipProtocol"] = args ? args.ipProtocol : undefined;
            inputs["isMirroringCollector"] = args ? args.isMirroringCollector : undefined;
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
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ForwardingRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ForwardingRule resources.
 */
export interface ForwardingRuleState {
    /**
     * For internal TCP/UDP load balancing (i.e. load balancing scheme is
     * INTERNAL and protocol is TCP/UDP), set this to true to allow packets
     * addressed to any ports to be forwarded to the backends configured
     * with this forwarding rule. Used with backend service. Cannot be set
     * if port or portRange are set.
     */
    readonly allPorts?: pulumi.Input<boolean>;
    /**
     * If true, clients can access ILB from all regions.
     * Otherwise only allows from the local region the ILB is located at.
     */
    readonly allowGlobalAccess?: pulumi.Input<boolean>;
    /**
     * A BackendService to receive the matched traffic. This is used only
     * for INTERNAL load balancing.
     */
    readonly backendService?: pulumi.Input<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
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
     * The IP protocol to which this rule applies.
     * When the load balancing scheme is INTERNAL, only TCP and UDP are
     * valid.
     */
    readonly ipProtocol?: pulumi.Input<string>;
    /**
     * Indicates whether or not this load balancer can be used
     * as a collector for packet mirroring. To prevent mirroring loops,
     * instances behind this load balancer will not have their traffic
     * mirrored even if a PacketMirroring rule applies to them. This
     * can only be set to true for load balancers that have their
     * loadBalancingScheme set to INTERNAL.
     */
    readonly isMirroringCollector?: pulumi.Input<boolean>;
    /**
     * The fingerprint used for optimistic locking of this resource. Used internally during updates.
     */
    readonly labelFingerprint?: pulumi.Input<string>;
    /**
     * Labels to apply to this forwarding rule.  A list of key->value pairs.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * This signifies what the ForwardingRule will be used for and can be
     * EXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is used for Classic
     * Cloud VPN gateways, protocol forwarding to VMs from an external IP address,
     * and HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP load balancers.
     * INTERNAL is used for protocol forwarding to VMs from an internal IP address,
     * and internal TCP/UDP load balancers.
     * INTERNAL_MANAGED is used for internal HTTP(S) load balancers.
     */
    readonly loadBalancingScheme?: pulumi.Input<string>;
    /**
     * Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * For internal load balancing, this field identifies the network that
     * the load balanced IP should belong to for this Forwarding Rule. If
     * this field is not specified, the default network will be used.
     * This field is only used for INTERNAL load balancing.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * The networking tier used for configuring this address. If this field is not
     * specified, it is assumed to be PREMIUM.
     */
    readonly networkTier?: pulumi.Input<string>;
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
     * This field is used along with the backendService field for internal
     * load balancing.
     * When the load balancing scheme is INTERNAL, a single port or a comma
     * separated list of ports can be configured. Only packets addressed to
     * these ports will be forwarded to the backends configured with this
     * forwarding rule.
     * You may specify a maximum of up to 5 ports.
     */
    readonly ports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A reference to the region where the regional forwarding rule resides.
     * This field is not applicable to global forwarding rules.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * An optional prefix to the service name for this Forwarding Rule.
     * If specified, will be the first label of the fully qualified service
     * name.
     * The label must be 1-63 characters long, and comply with RFC1035.
     * Specifically, the label must be 1-63 characters long and match the
     * regular expression `a-z?` which means the first
     * character must be a lowercase letter, and all following characters
     * must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     * This field is only used for INTERNAL load balancing.
     */
    readonly serviceLabel?: pulumi.Input<string>;
    /**
     * The internal fully qualified service name for this Forwarding Rule. This field is only used for INTERNAL load
     * balancing.
     */
    readonly serviceName?: pulumi.Input<string>;
    /**
     * The subnetwork that the load balanced IP should belong to for this
     * Forwarding Rule.  This field is only used for INTERNAL load balancing.
     * If the network specified is in auto subnet mode, this field is
     * optional. However, if the network is in custom subnet mode, a
     * subnetwork must be specified.
     */
    readonly subnetwork?: pulumi.Input<string>;
    /**
     * The URL of the target resource to receive the matched traffic.
     * The target must live in the same region as the forwarding rule.
     * The forwarded traffic must be of a type appropriate to the target
     * object.
     */
    readonly target?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ForwardingRule resource.
 */
export interface ForwardingRuleArgs {
    /**
     * For internal TCP/UDP load balancing (i.e. load balancing scheme is
     * INTERNAL and protocol is TCP/UDP), set this to true to allow packets
     * addressed to any ports to be forwarded to the backends configured
     * with this forwarding rule. Used with backend service. Cannot be set
     * if port or portRange are set.
     */
    readonly allPorts?: pulumi.Input<boolean>;
    /**
     * If true, clients can access ILB from all regions.
     * Otherwise only allows from the local region the ILB is located at.
     */
    readonly allowGlobalAccess?: pulumi.Input<boolean>;
    /**
     * A BackendService to receive the matched traffic. This is used only
     * for INTERNAL load balancing.
     */
    readonly backendService?: pulumi.Input<string>;
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
     * The IP protocol to which this rule applies.
     * When the load balancing scheme is INTERNAL, only TCP and UDP are
     * valid.
     */
    readonly ipProtocol?: pulumi.Input<string>;
    /**
     * Indicates whether or not this load balancer can be used
     * as a collector for packet mirroring. To prevent mirroring loops,
     * instances behind this load balancer will not have their traffic
     * mirrored even if a PacketMirroring rule applies to them. This
     * can only be set to true for load balancers that have their
     * loadBalancingScheme set to INTERNAL.
     */
    readonly isMirroringCollector?: pulumi.Input<boolean>;
    /**
     * Labels to apply to this forwarding rule.  A list of key->value pairs.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * This signifies what the ForwardingRule will be used for and can be
     * EXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is used for Classic
     * Cloud VPN gateways, protocol forwarding to VMs from an external IP address,
     * and HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP load balancers.
     * INTERNAL is used for protocol forwarding to VMs from an internal IP address,
     * and internal TCP/UDP load balancers.
     * INTERNAL_MANAGED is used for internal HTTP(S) load balancers.
     */
    readonly loadBalancingScheme?: pulumi.Input<string>;
    /**
     * Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * For internal load balancing, this field identifies the network that
     * the load balanced IP should belong to for this Forwarding Rule. If
     * this field is not specified, the default network will be used.
     * This field is only used for INTERNAL load balancing.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * The networking tier used for configuring this address. If this field is not
     * specified, it is assumed to be PREMIUM.
     */
    readonly networkTier?: pulumi.Input<string>;
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
     * This field is used along with the backendService field for internal
     * load balancing.
     * When the load balancing scheme is INTERNAL, a single port or a comma
     * separated list of ports can be configured. Only packets addressed to
     * these ports will be forwarded to the backends configured with this
     * forwarding rule.
     * You may specify a maximum of up to 5 ports.
     */
    readonly ports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A reference to the region where the regional forwarding rule resides.
     * This field is not applicable to global forwarding rules.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * An optional prefix to the service name for this Forwarding Rule.
     * If specified, will be the first label of the fully qualified service
     * name.
     * The label must be 1-63 characters long, and comply with RFC1035.
     * Specifically, the label must be 1-63 characters long and match the
     * regular expression `a-z?` which means the first
     * character must be a lowercase letter, and all following characters
     * must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     * This field is only used for INTERNAL load balancing.
     */
    readonly serviceLabel?: pulumi.Input<string>;
    /**
     * The subnetwork that the load balanced IP should belong to for this
     * Forwarding Rule.  This field is only used for INTERNAL load balancing.
     * If the network specified is in auto subnet mode, this field is
     * optional. However, if the network is in custom subnet mode, a
     * subnetwork must be specified.
     */
    readonly subnetwork?: pulumi.Input<string>;
    /**
     * The URL of the target resource to receive the matched traffic.
     * The target must live in the same region as the forwarding rule.
     * The forwarded traffic must be of a type appropriate to the target
     * object.
     */
    readonly target?: pulumi.Input<string>;
}
