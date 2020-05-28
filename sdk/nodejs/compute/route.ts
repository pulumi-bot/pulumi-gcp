// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Represents a Route resource.
 *
 * A route is a rule that specifies how certain packets should be handled by
 * the virtual network. Routes are associated with virtual machines by tag,
 * and the set of routes for a particular virtual machine is called its
 * routing table. For each packet leaving a virtual machine, the system
 * searches that virtual machine's routing table for a single best matching
 * route.
 *
 * Routes match packets by destination IP address, preferring smaller or more
 * specific ranges over larger ones. If there is a tie, the system selects
 * the route with the smallest priority value. If there is still a tie, it
 * uses the layer three and four packet headers to select just one of the
 * remaining matching routes. The packet is then forwarded as specified by
 * the nextHop field of the winning route -- either to another virtual
 * machine destination, a virtual machine gateway or a Compute
 * Engine-operated gateway. Packets that do not match any route in the
 * sending virtual machine's routing table will be dropped.
 *
 * A Route resource must have exactly one specification of either
 * nextHopGateway, nextHopInstance, nextHopIp, nextHopVpnTunnel, or
 * nextHopIlb.
 *
 *
 * To get more information about Route, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/routes)
 * * How-to Guides
 *     * [Using Routes](https://cloud.google.com/vpc/docs/using-routes)
 *
 * ## Example Usage - Route Basic
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultNetwork = new gcp.compute.Network("defaultNetwork", {});
 * const defaultRoute = new gcp.compute.Route("defaultRoute", {
 *     destRange: "15.0.0.0/24",
 *     network: defaultNetwork.name,
 *     nextHopIp: "10.132.1.5",
 *     priority: 100,
 * });
 * ```
 * ## Example Usage - Route Ilb
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultNetwork = new gcp.compute.Network("defaultNetwork", {autoCreateSubnetworks: false});
 * const defaultSubnetwork = new gcp.compute.Subnetwork("defaultSubnetwork", {
 *     ipCidrRange: "10.0.1.0/24",
 *     region: "us-central1",
 *     network: defaultNetwork.id,
 * });
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
 * const defaultForwardingRule = new gcp.compute.ForwardingRule("defaultForwardingRule", {
 *     region: "us-central1",
 *     loadBalancingScheme: "INTERNAL",
 *     backendService: backend.id,
 *     allPorts: true,
 *     network: defaultNetwork.name,
 *     subnetwork: defaultSubnetwork.name,
 * });
 * const routeIlb = new gcp.compute.Route("route-ilb", {
 *     destRange: "0.0.0.0/0",
 *     network: defaultNetwork.name,
 *     nextHopIlb: defaultForwardingRule.id,
 *     priority: 2000,
 * });
 * ```
 */
export class Route extends pulumi.CustomResource {
    /**
     * Get an existing Route resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteState, opts?: pulumi.CustomResourceOptions): Route {
        return new Route(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/route:Route';

    /**
     * Returns true if the given object is an instance of Route.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Route {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Route.__pulumiType;
    }

    /**
     * An optional description of this resource. Provide this property
     * when you create the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The destination range of outgoing packets that this route applies to.
     * Only IPv4 is supported.
     */
    public readonly destRange!: pulumi.Output<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The network that this route applies to.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * URL to a gateway that should handle matching packets.
     * Currently, you can only specify the internet gateway, using a full or
     * partial valid URL:
     * * `https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway`
     * * `projects/project/global/gateways/default-internet-gateway`
     * * `global/gateways/default-internet-gateway`
     * * The string `default-internet-gateway`.
     */
    public readonly nextHopGateway!: pulumi.Output<string | undefined>;
    /**
     * The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets.
     * You can only specify the forwarding rule as a partial or full URL. For example, the following are all valid URLs:
     * https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
     * regions/region/forwardingRules/forwardingRule
     * Note that this can only be used when the destinationRange is a public (non-RFC 1918) IP CIDR range.
     */
    public readonly nextHopIlb!: pulumi.Output<string | undefined>;
    /**
     * URL to an instance that should handle matching packets.
     * You can specify this as a full or partial URL. For example:
     * * `https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance`
     * * `projects/project/zones/zone/instances/instance`
     * * `zones/zone/instances/instance`
     * * Just the instance name, with the zone in `nextHopInstanceZone`.
     */
    public readonly nextHopInstance!: pulumi.Output<string | undefined>;
    /**
     * (Optional when `nextHopInstance` is
     * specified)  The zone of the instance specified in
     * `nextHopInstance`.  Omit if `nextHopInstance` is specified as
     * a URL.
     */
    public readonly nextHopInstanceZone!: pulumi.Output<string | undefined>;
    /**
     * Network IP address of an instance that should handle matching packets.
     */
    public readonly nextHopIp!: pulumi.Output<string>;
    /**
     * URL to a Network that should handle matching packets.
     */
    public /*out*/ readonly nextHopNetwork!: pulumi.Output<string>;
    /**
     * URL to a VpnTunnel that should handle matching packets.
     */
    public readonly nextHopVpnTunnel!: pulumi.Output<string | undefined>;
    /**
     * The priority of this route. Priority is used to break ties in cases
     * where there is more than one matching route of equal prefix length.
     * In the case of two routes with equal prefix length, the one with the
     * lowest-numbered priority value wins.
     * Default value is 1000. Valid range is 0 through 65535.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
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
     * A list of instance tags to which this route applies.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Route resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteArgs | RouteState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RouteState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["destRange"] = state ? state.destRange : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["nextHopGateway"] = state ? state.nextHopGateway : undefined;
            inputs["nextHopIlb"] = state ? state.nextHopIlb : undefined;
            inputs["nextHopInstance"] = state ? state.nextHopInstance : undefined;
            inputs["nextHopInstanceZone"] = state ? state.nextHopInstanceZone : undefined;
            inputs["nextHopIp"] = state ? state.nextHopIp : undefined;
            inputs["nextHopNetwork"] = state ? state.nextHopNetwork : undefined;
            inputs["nextHopVpnTunnel"] = state ? state.nextHopVpnTunnel : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as RouteArgs | undefined;
            if (!args || args.destRange === undefined) {
                throw new Error("Missing required property 'destRange'");
            }
            if (!args || args.network === undefined) {
                throw new Error("Missing required property 'network'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["destRange"] = args ? args.destRange : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["nextHopGateway"] = args ? args.nextHopGateway : undefined;
            inputs["nextHopIlb"] = args ? args.nextHopIlb : undefined;
            inputs["nextHopInstance"] = args ? args.nextHopInstance : undefined;
            inputs["nextHopInstanceZone"] = args ? args.nextHopInstanceZone : undefined;
            inputs["nextHopIp"] = args ? args.nextHopIp : undefined;
            inputs["nextHopVpnTunnel"] = args ? args.nextHopVpnTunnel : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["nextHopNetwork"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Route.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Route resources.
 */
export interface RouteState {
    /**
     * An optional description of this resource. Provide this property
     * when you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The destination range of outgoing packets that this route applies to.
     * Only IPv4 is supported.
     */
    readonly destRange?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The network that this route applies to.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * URL to a gateway that should handle matching packets.
     * Currently, you can only specify the internet gateway, using a full or
     * partial valid URL:
     * * `https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway`
     * * `projects/project/global/gateways/default-internet-gateway`
     * * `global/gateways/default-internet-gateway`
     * * The string `default-internet-gateway`.
     */
    readonly nextHopGateway?: pulumi.Input<string>;
    /**
     * The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets.
     * You can only specify the forwarding rule as a partial or full URL. For example, the following are all valid URLs:
     * https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
     * regions/region/forwardingRules/forwardingRule
     * Note that this can only be used when the destinationRange is a public (non-RFC 1918) IP CIDR range.
     */
    readonly nextHopIlb?: pulumi.Input<string>;
    /**
     * URL to an instance that should handle matching packets.
     * You can specify this as a full or partial URL. For example:
     * * `https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance`
     * * `projects/project/zones/zone/instances/instance`
     * * `zones/zone/instances/instance`
     * * Just the instance name, with the zone in `nextHopInstanceZone`.
     */
    readonly nextHopInstance?: pulumi.Input<string>;
    /**
     * (Optional when `nextHopInstance` is
     * specified)  The zone of the instance specified in
     * `nextHopInstance`.  Omit if `nextHopInstance` is specified as
     * a URL.
     */
    readonly nextHopInstanceZone?: pulumi.Input<string>;
    /**
     * Network IP address of an instance that should handle matching packets.
     */
    readonly nextHopIp?: pulumi.Input<string>;
    /**
     * URL to a Network that should handle matching packets.
     */
    readonly nextHopNetwork?: pulumi.Input<string>;
    /**
     * URL to a VpnTunnel that should handle matching packets.
     */
    readonly nextHopVpnTunnel?: pulumi.Input<string>;
    /**
     * The priority of this route. Priority is used to break ties in cases
     * where there is more than one matching route of equal prefix length.
     * In the case of two routes with equal prefix length, the one with the
     * lowest-numbered priority value wins.
     * Default value is 1000. Valid range is 0 through 65535.
     */
    readonly priority?: pulumi.Input<number>;
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
     * A list of instance tags to which this route applies.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Route resource.
 */
export interface RouteArgs {
    /**
     * An optional description of this resource. Provide this property
     * when you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The destination range of outgoing packets that this route applies to.
     * Only IPv4 is supported.
     */
    readonly destRange: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The network that this route applies to.
     */
    readonly network: pulumi.Input<string>;
    /**
     * URL to a gateway that should handle matching packets.
     * Currently, you can only specify the internet gateway, using a full or
     * partial valid URL:
     * * `https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway`
     * * `projects/project/global/gateways/default-internet-gateway`
     * * `global/gateways/default-internet-gateway`
     * * The string `default-internet-gateway`.
     */
    readonly nextHopGateway?: pulumi.Input<string>;
    /**
     * The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets.
     * You can only specify the forwarding rule as a partial or full URL. For example, the following are all valid URLs:
     * https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
     * regions/region/forwardingRules/forwardingRule
     * Note that this can only be used when the destinationRange is a public (non-RFC 1918) IP CIDR range.
     */
    readonly nextHopIlb?: pulumi.Input<string>;
    /**
     * URL to an instance that should handle matching packets.
     * You can specify this as a full or partial URL. For example:
     * * `https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance`
     * * `projects/project/zones/zone/instances/instance`
     * * `zones/zone/instances/instance`
     * * Just the instance name, with the zone in `nextHopInstanceZone`.
     */
    readonly nextHopInstance?: pulumi.Input<string>;
    /**
     * (Optional when `nextHopInstance` is
     * specified)  The zone of the instance specified in
     * `nextHopInstance`.  Omit if `nextHopInstance` is specified as
     * a URL.
     */
    readonly nextHopInstanceZone?: pulumi.Input<string>;
    /**
     * Network IP address of an instance that should handle matching packets.
     */
    readonly nextHopIp?: pulumi.Input<string>;
    /**
     * URL to a VpnTunnel that should handle matching packets.
     */
    readonly nextHopVpnTunnel?: pulumi.Input<string>;
    /**
     * The priority of this route. Priority is used to break ties in cases
     * where there is more than one matching route of equal prefix length.
     * In the case of two routes with equal prefix length, the one with the
     * lowest-numbered priority value wins.
     * Default value is 1000. Valid range is 0 through 65535.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A list of instance tags to which this route applies.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
}
