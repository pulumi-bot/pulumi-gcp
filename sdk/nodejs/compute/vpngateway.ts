// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Represents a VPN gateway running in GCP. This virtual device is managed
 * by Google, but used only by you.
 *
 * To get more information about VpnGateway, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/targetVpnGateways)
 *
 * > **Warning:** Classic VPN is deprecating certain functionality on October 31, 2021. For more information,
 * see the [Classic VPN partial deprecation page](https://cloud.google.com/network-connectivity/docs/vpn/deprecations/classic-vpn-deprecation).
 *
 * ## Example Usage
 * ### Target Vpn Gateway Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const network1 = new gcp.compute.Network("network1", {});
 * const targetGateway = new gcp.compute.VPNGateway("targetGateway", {network: network1.id});
 * const vpnStaticIp = new gcp.compute.Address("vpnStaticIp", {});
 * const frEsp = new gcp.compute.ForwardingRule("frEsp", {
 *     ipProtocol: "ESP",
 *     ipAddress: vpnStaticIp.address,
 *     target: targetGateway.id,
 * });
 * const frUdp500 = new gcp.compute.ForwardingRule("frUdp500", {
 *     ipProtocol: "UDP",
 *     portRange: "500",
 *     ipAddress: vpnStaticIp.address,
 *     target: targetGateway.id,
 * });
 * const frUdp4500 = new gcp.compute.ForwardingRule("frUdp4500", {
 *     ipProtocol: "UDP",
 *     portRange: "4500",
 *     ipAddress: vpnStaticIp.address,
 *     target: targetGateway.id,
 * });
 * const tunnel1 = new gcp.compute.VPNTunnel("tunnel1", {
 *     peerIp: "15.0.0.120",
 *     sharedSecret: "a secret message",
 *     targetVpnGateway: targetGateway.id,
 * }, {
 *     dependsOn: [
 *         frEsp,
 *         frUdp500,
 *         frUdp4500,
 *     ],
 * });
 * const route1 = new gcp.compute.Route("route1", {
 *     network: network1.name,
 *     destRange: "15.0.0.0/24",
 *     priority: 1000,
 *     nextHopVpnTunnel: tunnel1.id,
 * });
 * ```
 *
 * ## Import
 *
 * VpnGateway can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/vPNGateway:VPNGateway default projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/vPNGateway:VPNGateway default {{project}}/{{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/vPNGateway:VPNGateway default {{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/vPNGateway:VPNGateway default {{name}}
 * ```
 */
export class VPNGateway extends pulumi.CustomResource {
    /**
     * Get an existing VPNGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VPNGatewayState, opts?: pulumi.CustomResourceOptions): VPNGateway {
        return new VPNGateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/vPNGateway:VPNGateway';

    /**
     * Returns true if the given object is an instance of VPNGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VPNGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VPNGateway.__pulumiType;
    }

    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The unique identifier for the resource.
     */
    public /*out*/ readonly gatewayId!: pulumi.Output<number>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The network this VPN gateway is accepting traffic for.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The region this gateway should sit in.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a VPNGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VPNGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VPNGatewayArgs | VPNGatewayState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VPNGatewayState | undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["gatewayId"] = state ? state.gatewayId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as VPNGatewayArgs | undefined;
            if ((!args || args.network === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'network'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["gatewayId"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VPNGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VPNGateway resources.
 */
export interface VPNGatewayState {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The unique identifier for the resource.
     */
    readonly gatewayId?: pulumi.Input<number>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The network this VPN gateway is accepting traffic for.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region this gateway should sit in.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VPNGateway resource.
 */
export interface VPNGatewayArgs {
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The network this VPN gateway is accepting traffic for.
     */
    readonly network: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region this gateway should sit in.
     */
    readonly region?: pulumi.Input<string>;
}
