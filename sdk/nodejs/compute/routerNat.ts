// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A NAT service created in a router.
 *
 * To get more information about RouterNat, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/routers)
 * * How-to Guides
 *     * [Google Cloud Router](https://cloud.google.com/router/docs/)
 *
 * ## Example Usage
 * ### Router Nat Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const net = new gcp.compute.Network("net", {});
 * const subnet = new gcp.compute.Subnetwork("subnet", {
 *     network: net.id,
 *     ipCidrRange: "10.0.0.0/16",
 *     region: "us-central1",
 * });
 * const router = new gcp.compute.Router("router", {
 *     region: subnet.region,
 *     network: net.id,
 *     bgp: {
 *         asn: 64514,
 *     },
 * });
 * const nat = new gcp.compute.RouterNat("nat", {
 *     router: router.name,
 *     region: router.region,
 *     natIpAllocateOption: "AUTO_ONLY",
 *     sourceSubnetworkIpRangesToNat: "ALL_SUBNETWORKS_ALL_IP_RANGES",
 *     logConfig: {
 *         enable: true,
 *         filter: "ERRORS_ONLY",
 *     },
 * });
 * ```
 * ### Router Nat Manual Ips
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const net = new gcp.compute.Network("net", {});
 * const subnet = new gcp.compute.Subnetwork("subnet", {
 *     network: net.id,
 *     ipCidrRange: "10.0.0.0/16",
 *     region: "us-central1",
 * });
 * const router = new gcp.compute.Router("router", {
 *     region: subnet.region,
 *     network: net.id,
 * });
 * const address: gcp.compute.Address[];
 * for (const range = {value: 0}; range.value < 2; range.value++) {
 *     address.push(new gcp.compute.Address(`address-${range.value}`, {region: subnet.region}));
 * }
 * const natManual = new gcp.compute.RouterNat("natManual", {
 *     router: router.name,
 *     region: router.region,
 *     natIpAllocateOption: "MANUAL_ONLY",
 *     natIps: address.map(__item => __item.selfLink),
 *     sourceSubnetworkIpRangesToNat: "LIST_OF_SUBNETWORKS",
 *     subnetworks: [{
 *         name: subnet.id,
 *         sourceIpRangesToNats: ["ALL_IP_RANGES"],
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * RouterNat can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/routerNat:RouterNat default projects/{{project}}/regions/{{region}}/routers/{{router}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/routerNat:RouterNat default {{project}}/{{region}}/{{router}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/routerNat:RouterNat default {{region}}/{{router}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/routerNat:RouterNat default {{router}}/{{name}}
 * ```
 */
export class RouterNat extends pulumi.CustomResource {
    /**
     * Get an existing RouterNat resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterNatState, opts?: pulumi.CustomResourceOptions): RouterNat {
        return new RouterNat(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/routerNat:RouterNat';

    /**
     * Returns true if the given object is an instance of RouterNat.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterNat {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterNat.__pulumiType;
    }

    /**
     * A list of URLs of the IP resources to be drained. These IPs must be
     * valid static external IPs that have been assigned to the NAT.
     */
    public readonly drainNatIps!: pulumi.Output<string[] | undefined>;
    /**
     * Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
     */
    public readonly icmpIdleTimeoutSec!: pulumi.Output<number | undefined>;
    /**
     * Configuration for logging on NAT
     * Structure is documented below.
     */
    public readonly logConfig!: pulumi.Output<outputs.compute.RouterNatLogConfig | undefined>;
    /**
     * Minimum number of ports allocated to a VM from this NAT.
     */
    public readonly minPortsPerVm!: pulumi.Output<number | undefined>;
    /**
     * Self-link of subnetwork to NAT
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * How external IPs should be allocated for this NAT. Valid values are
     * `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
     * Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
     * Possible values are `MANUAL_ONLY` and `AUTO_ONLY`.
     */
    public readonly natIpAllocateOption!: pulumi.Output<string>;
    /**
     * Self-links of NAT IPs. Only valid if natIpAllocateOption
     * is set to MANUAL_ONLY.
     */
    public readonly natIps!: pulumi.Output<string[] | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Region where the router and NAT reside.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The name of the Cloud Router in which this NAT will be configured.
     */
    public readonly router!: pulumi.Output<string>;
    /**
     * How NAT should be configured per Subnetwork.
     * If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
     * IP ranges in every Subnetwork are allowed to Nat.
     * If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
     * ranges in every Subnetwork are allowed to Nat.
     * `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
     * (specified in the field subnetwork below). Note that if this field
     * contains ALL_SUBNETWORKS_ALL_IP_RANGES or
     * ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
     * other RouterNat section in any Router for this network in this region.
     * Possible values are `ALL_SUBNETWORKS_ALL_IP_RANGES`, `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, and `LIST_OF_SUBNETWORKS`.
     */
    public readonly sourceSubnetworkIpRangesToNat!: pulumi.Output<string>;
    /**
     * One or more subnetwork NAT configurations. Only used if
     * `sourceSubnetworkIpRangesToNat` is set to `LIST_OF_SUBNETWORKS`
     * Structure is documented below.
     */
    public readonly subnetworks!: pulumi.Output<outputs.compute.RouterNatSubnetwork[] | undefined>;
    /**
     * Timeout (in seconds) for TCP established connections.
     * Defaults to 1200s if not set.
     */
    public readonly tcpEstablishedIdleTimeoutSec!: pulumi.Output<number | undefined>;
    /**
     * Timeout (in seconds) for TCP transitory connections.
     * Defaults to 30s if not set.
     */
    public readonly tcpTransitoryIdleTimeoutSec!: pulumi.Output<number | undefined>;
    /**
     * Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
     */
    public readonly udpIdleTimeoutSec!: pulumi.Output<number | undefined>;

    /**
     * Create a RouterNat resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouterNatArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterNatArgs | RouterNatState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RouterNatState | undefined;
            inputs["drainNatIps"] = state ? state.drainNatIps : undefined;
            inputs["icmpIdleTimeoutSec"] = state ? state.icmpIdleTimeoutSec : undefined;
            inputs["logConfig"] = state ? state.logConfig : undefined;
            inputs["minPortsPerVm"] = state ? state.minPortsPerVm : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["natIpAllocateOption"] = state ? state.natIpAllocateOption : undefined;
            inputs["natIps"] = state ? state.natIps : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["router"] = state ? state.router : undefined;
            inputs["sourceSubnetworkIpRangesToNat"] = state ? state.sourceSubnetworkIpRangesToNat : undefined;
            inputs["subnetworks"] = state ? state.subnetworks : undefined;
            inputs["tcpEstablishedIdleTimeoutSec"] = state ? state.tcpEstablishedIdleTimeoutSec : undefined;
            inputs["tcpTransitoryIdleTimeoutSec"] = state ? state.tcpTransitoryIdleTimeoutSec : undefined;
            inputs["udpIdleTimeoutSec"] = state ? state.udpIdleTimeoutSec : undefined;
        } else {
            const args = argsOrState as RouterNatArgs | undefined;
            if ((!args || args.natIpAllocateOption === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'natIpAllocateOption'");
            }
            if ((!args || args.router === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'router'");
            }
            if ((!args || args.sourceSubnetworkIpRangesToNat === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'sourceSubnetworkIpRangesToNat'");
            }
            inputs["drainNatIps"] = args ? args.drainNatIps : undefined;
            inputs["icmpIdleTimeoutSec"] = args ? args.icmpIdleTimeoutSec : undefined;
            inputs["logConfig"] = args ? args.logConfig : undefined;
            inputs["minPortsPerVm"] = args ? args.minPortsPerVm : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["natIpAllocateOption"] = args ? args.natIpAllocateOption : undefined;
            inputs["natIps"] = args ? args.natIps : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["router"] = args ? args.router : undefined;
            inputs["sourceSubnetworkIpRangesToNat"] = args ? args.sourceSubnetworkIpRangesToNat : undefined;
            inputs["subnetworks"] = args ? args.subnetworks : undefined;
            inputs["tcpEstablishedIdleTimeoutSec"] = args ? args.tcpEstablishedIdleTimeoutSec : undefined;
            inputs["tcpTransitoryIdleTimeoutSec"] = args ? args.tcpTransitoryIdleTimeoutSec : undefined;
            inputs["udpIdleTimeoutSec"] = args ? args.udpIdleTimeoutSec : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RouterNat.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterNat resources.
 */
export interface RouterNatState {
    /**
     * A list of URLs of the IP resources to be drained. These IPs must be
     * valid static external IPs that have been assigned to the NAT.
     */
    readonly drainNatIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
     */
    readonly icmpIdleTimeoutSec?: pulumi.Input<number>;
    /**
     * Configuration for logging on NAT
     * Structure is documented below.
     */
    readonly logConfig?: pulumi.Input<inputs.compute.RouterNatLogConfig>;
    /**
     * Minimum number of ports allocated to a VM from this NAT.
     */
    readonly minPortsPerVm?: pulumi.Input<number>;
    /**
     * Self-link of subnetwork to NAT
     */
    readonly name?: pulumi.Input<string>;
    /**
     * How external IPs should be allocated for this NAT. Valid values are
     * `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
     * Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
     * Possible values are `MANUAL_ONLY` and `AUTO_ONLY`.
     */
    readonly natIpAllocateOption?: pulumi.Input<string>;
    /**
     * Self-links of NAT IPs. Only valid if natIpAllocateOption
     * is set to MANUAL_ONLY.
     */
    readonly natIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Region where the router and NAT reside.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The name of the Cloud Router in which this NAT will be configured.
     */
    readonly router?: pulumi.Input<string>;
    /**
     * How NAT should be configured per Subnetwork.
     * If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
     * IP ranges in every Subnetwork are allowed to Nat.
     * If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
     * ranges in every Subnetwork are allowed to Nat.
     * `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
     * (specified in the field subnetwork below). Note that if this field
     * contains ALL_SUBNETWORKS_ALL_IP_RANGES or
     * ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
     * other RouterNat section in any Router for this network in this region.
     * Possible values are `ALL_SUBNETWORKS_ALL_IP_RANGES`, `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, and `LIST_OF_SUBNETWORKS`.
     */
    readonly sourceSubnetworkIpRangesToNat?: pulumi.Input<string>;
    /**
     * One or more subnetwork NAT configurations. Only used if
     * `sourceSubnetworkIpRangesToNat` is set to `LIST_OF_SUBNETWORKS`
     * Structure is documented below.
     */
    readonly subnetworks?: pulumi.Input<pulumi.Input<inputs.compute.RouterNatSubnetwork>[]>;
    /**
     * Timeout (in seconds) for TCP established connections.
     * Defaults to 1200s if not set.
     */
    readonly tcpEstablishedIdleTimeoutSec?: pulumi.Input<number>;
    /**
     * Timeout (in seconds) for TCP transitory connections.
     * Defaults to 30s if not set.
     */
    readonly tcpTransitoryIdleTimeoutSec?: pulumi.Input<number>;
    /**
     * Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
     */
    readonly udpIdleTimeoutSec?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RouterNat resource.
 */
export interface RouterNatArgs {
    /**
     * A list of URLs of the IP resources to be drained. These IPs must be
     * valid static external IPs that have been assigned to the NAT.
     */
    readonly drainNatIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
     */
    readonly icmpIdleTimeoutSec?: pulumi.Input<number>;
    /**
     * Configuration for logging on NAT
     * Structure is documented below.
     */
    readonly logConfig?: pulumi.Input<inputs.compute.RouterNatLogConfig>;
    /**
     * Minimum number of ports allocated to a VM from this NAT.
     */
    readonly minPortsPerVm?: pulumi.Input<number>;
    /**
     * Self-link of subnetwork to NAT
     */
    readonly name?: pulumi.Input<string>;
    /**
     * How external IPs should be allocated for this NAT. Valid values are
     * `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
     * Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
     * Possible values are `MANUAL_ONLY` and `AUTO_ONLY`.
     */
    readonly natIpAllocateOption: pulumi.Input<string>;
    /**
     * Self-links of NAT IPs. Only valid if natIpAllocateOption
     * is set to MANUAL_ONLY.
     */
    readonly natIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Region where the router and NAT reside.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The name of the Cloud Router in which this NAT will be configured.
     */
    readonly router: pulumi.Input<string>;
    /**
     * How NAT should be configured per Subnetwork.
     * If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
     * IP ranges in every Subnetwork are allowed to Nat.
     * If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
     * ranges in every Subnetwork are allowed to Nat.
     * `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
     * (specified in the field subnetwork below). Note that if this field
     * contains ALL_SUBNETWORKS_ALL_IP_RANGES or
     * ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
     * other RouterNat section in any Router for this network in this region.
     * Possible values are `ALL_SUBNETWORKS_ALL_IP_RANGES`, `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, and `LIST_OF_SUBNETWORKS`.
     */
    readonly sourceSubnetworkIpRangesToNat: pulumi.Input<string>;
    /**
     * One or more subnetwork NAT configurations. Only used if
     * `sourceSubnetworkIpRangesToNat` is set to `LIST_OF_SUBNETWORKS`
     * Structure is documented below.
     */
    readonly subnetworks?: pulumi.Input<pulumi.Input<inputs.compute.RouterNatSubnetwork>[]>;
    /**
     * Timeout (in seconds) for TCP established connections.
     * Defaults to 1200s if not set.
     */
    readonly tcpEstablishedIdleTimeoutSec?: pulumi.Input<number>;
    /**
     * Timeout (in seconds) for TCP transitory connections.
     * Defaults to 30s if not set.
     */
    readonly tcpTransitoryIdleTimeoutSec?: pulumi.Input<number>;
    /**
     * Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
     */
    readonly udpIdleTimeoutSec?: pulumi.Input<number>;
}
