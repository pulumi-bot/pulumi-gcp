// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * VPN tunnel resource.
 *
 *
 * To get more information about VpnTunnel, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/vpnTunnels)
 * * How-to Guides
 *     * [Cloud VPN Overview](https://cloud.google.com/vpn/docs/concepts/overview)
 *     * [Networks and Tunnel Routing](https://cloud.google.com/vpn/docs/concepts/choosing-networks-routing)
 *
 * > **Warning:** All arguments including `sharedSecret` will be stored in the raw
 * state as plain-text.
 *
 * ## Example Usage - Vpn Tunnel Basic
 *
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
 * });
 * const route1 = new gcp.compute.Route("route1", {
 *     network: network1.name,
 *     destRange: "15.0.0.0/24",
 *     priority: 1000,
 *     nextHopVpnTunnel: tunnel1.id,
 * });
 * ```
 * ## Example Usage - Vpn Tunnel Beta
 *
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
 *     labels: {
 *         foo: "bar",
 *     },
 * });
 * const route1 = new gcp.compute.Route("route1", {
 *     network: network1.name,
 *     destRange: "15.0.0.0/24",
 *     priority: 1000,
 *     nextHopVpnTunnel: tunnel1.id,
 * });
 * ```
 */
export class VPNTunnel extends pulumi.CustomResource {
    /**
     * Get an existing VPNTunnel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VPNTunnelState, opts?: pulumi.CustomResourceOptions): VPNTunnel {
        return new VPNTunnel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/vPNTunnel:VPNTunnel';

    /**
     * Returns true if the given object is an instance of VPNTunnel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VPNTunnel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VPNTunnel.__pulumiType;
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
     * Detailed status message for the VPN tunnel.
     */
    public /*out*/ readonly detailedStatus!: pulumi.Output<string>;
    /**
     * IKE protocol version to use when establishing the VPN tunnel with
     * peer VPN gateway.
     * Acceptable IKE versions are 1 or 2. Default version is 2.
     */
    public readonly ikeVersion!: pulumi.Output<number | undefined>;
    /**
     * The fingerprint used for optimistic locking of this resource. Used internally during updates.
     */
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * Labels to apply to this VpnTunnel.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Local traffic selector to use when establishing the VPN tunnel with
     * peer VPN gateway. The value should be a CIDR formatted string,
     * for example `192.168.0.0/16`. The ranges should be disjoint.
     * Only IPv4 is supported.
     */
    public readonly localTrafficSelectors!: pulumi.Output<string[]>;
    /**
     * Name of the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63
     * characters long and match the regular expression
     * `a-z?` which means the first character
     * must be a lowercase letter, and all following characters must
     * be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * URL of the peer side external VPN gateway to which this VPN tunnel is connected.
     */
    public readonly peerExternalGateway!: pulumi.Output<string | undefined>;
    /**
     * The interface ID of the external VPN gateway to which this VPN tunnel is connected.
     */
    public readonly peerExternalGatewayInterface!: pulumi.Output<number | undefined>;
    /**
     * URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
     * If provided, the VPN tunnel will automatically use the same vpnGatewayInterface
     * ID in the peer GCP VPN gateway.
     * This field must reference a `gcp.compute.HaVpnGateway` resource.
     */
    public readonly peerGcpGateway!: pulumi.Output<string | undefined>;
    /**
     * IP address of the peer VPN gateway. Only IPv4 is supported.
     */
    public readonly peerIp!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The region where the tunnel is located. If unset, is set to the region of `targetVpnGateway`.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Remote traffic selector to use when establishing the VPN tunnel with
     * peer VPN gateway. The value should be a CIDR formatted string,
     * for example `192.168.0.0/16`. The ranges should be disjoint.
     * Only IPv4 is supported.
     */
    public readonly remoteTrafficSelectors!: pulumi.Output<string[]>;
    /**
     * URL of router resource to be used for dynamic routing.
     */
    public readonly router!: pulumi.Output<string | undefined>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Shared secret used to set the secure session between the Cloud VPN
     * gateway and the peer VPN gateway.  **Note**: This property is sensitive and will not be displayed in the plan.
     */
    public readonly sharedSecret!: pulumi.Output<string>;
    /**
     * Hash of the shared secret.
     */
    public /*out*/ readonly sharedSecretHash!: pulumi.Output<string>;
    /**
     * URL of the Target VPN gateway with which this VPN tunnel is
     * associated.
     */
    public readonly targetVpnGateway!: pulumi.Output<string | undefined>;
    /**
     * The unique identifier for the resource. This identifier is defined by the server.
     */
    public /*out*/ readonly tunnelId!: pulumi.Output<string>;
    /**
     * URL of the VPN gateway with which this VPN tunnel is associated.
     * This must be used if a High Availability VPN gateway resource is created.
     * This field must reference a `gcp.compute.HaVpnGateway` resource.
     */
    public readonly vpnGateway!: pulumi.Output<string | undefined>;
    /**
     * The interface ID of the VPN gateway with which this VPN tunnel is associated.
     */
    public readonly vpnGatewayInterface!: pulumi.Output<number | undefined>;

    /**
     * Create a VPNTunnel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VPNTunnelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VPNTunnelArgs | VPNTunnelState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VPNTunnelState | undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["detailedStatus"] = state ? state.detailedStatus : undefined;
            inputs["ikeVersion"] = state ? state.ikeVersion : undefined;
            inputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["localTrafficSelectors"] = state ? state.localTrafficSelectors : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["peerExternalGateway"] = state ? state.peerExternalGateway : undefined;
            inputs["peerExternalGatewayInterface"] = state ? state.peerExternalGatewayInterface : undefined;
            inputs["peerGcpGateway"] = state ? state.peerGcpGateway : undefined;
            inputs["peerIp"] = state ? state.peerIp : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["remoteTrafficSelectors"] = state ? state.remoteTrafficSelectors : undefined;
            inputs["router"] = state ? state.router : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["sharedSecret"] = state ? state.sharedSecret : undefined;
            inputs["sharedSecretHash"] = state ? state.sharedSecretHash : undefined;
            inputs["targetVpnGateway"] = state ? state.targetVpnGateway : undefined;
            inputs["tunnelId"] = state ? state.tunnelId : undefined;
            inputs["vpnGateway"] = state ? state.vpnGateway : undefined;
            inputs["vpnGatewayInterface"] = state ? state.vpnGatewayInterface : undefined;
        } else {
            const args = argsOrState as VPNTunnelArgs | undefined;
            if (!args || args.sharedSecret === undefined) {
                throw new Error("Missing required property 'sharedSecret'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["ikeVersion"] = args ? args.ikeVersion : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["localTrafficSelectors"] = args ? args.localTrafficSelectors : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["peerExternalGateway"] = args ? args.peerExternalGateway : undefined;
            inputs["peerExternalGatewayInterface"] = args ? args.peerExternalGatewayInterface : undefined;
            inputs["peerGcpGateway"] = args ? args.peerGcpGateway : undefined;
            inputs["peerIp"] = args ? args.peerIp : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["remoteTrafficSelectors"] = args ? args.remoteTrafficSelectors : undefined;
            inputs["router"] = args ? args.router : undefined;
            inputs["sharedSecret"] = args ? args.sharedSecret : undefined;
            inputs["targetVpnGateway"] = args ? args.targetVpnGateway : undefined;
            inputs["vpnGateway"] = args ? args.vpnGateway : undefined;
            inputs["vpnGatewayInterface"] = args ? args.vpnGatewayInterface : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["detailedStatus"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["sharedSecretHash"] = undefined /*out*/;
            inputs["tunnelId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VPNTunnel.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VPNTunnel resources.
 */
export interface VPNTunnelState {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Detailed status message for the VPN tunnel.
     */
    readonly detailedStatus?: pulumi.Input<string>;
    /**
     * IKE protocol version to use when establishing the VPN tunnel with
     * peer VPN gateway.
     * Acceptable IKE versions are 1 or 2. Default version is 2.
     */
    readonly ikeVersion?: pulumi.Input<number>;
    /**
     * The fingerprint used for optimistic locking of this resource. Used internally during updates.
     */
    readonly labelFingerprint?: pulumi.Input<string>;
    /**
     * Labels to apply to this VpnTunnel.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Local traffic selector to use when establishing the VPN tunnel with
     * peer VPN gateway. The value should be a CIDR formatted string,
     * for example `192.168.0.0/16`. The ranges should be disjoint.
     * Only IPv4 is supported.
     */
    readonly localTrafficSelectors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63
     * characters long and match the regular expression
     * `a-z?` which means the first character
     * must be a lowercase letter, and all following characters must
     * be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * URL of the peer side external VPN gateway to which this VPN tunnel is connected.
     */
    readonly peerExternalGateway?: pulumi.Input<string>;
    /**
     * The interface ID of the external VPN gateway to which this VPN tunnel is connected.
     */
    readonly peerExternalGatewayInterface?: pulumi.Input<number>;
    /**
     * URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
     * If provided, the VPN tunnel will automatically use the same vpnGatewayInterface
     * ID in the peer GCP VPN gateway.
     * This field must reference a `gcp.compute.HaVpnGateway` resource.
     */
    readonly peerGcpGateway?: pulumi.Input<string>;
    /**
     * IP address of the peer VPN gateway. Only IPv4 is supported.
     */
    readonly peerIp?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region where the tunnel is located. If unset, is set to the region of `targetVpnGateway`.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Remote traffic selector to use when establishing the VPN tunnel with
     * peer VPN gateway. The value should be a CIDR formatted string,
     * for example `192.168.0.0/16`. The ranges should be disjoint.
     * Only IPv4 is supported.
     */
    readonly remoteTrafficSelectors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * URL of router resource to be used for dynamic routing.
     */
    readonly router?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * Shared secret used to set the secure session between the Cloud VPN
     * gateway and the peer VPN gateway.  **Note**: This property is sensitive and will not be displayed in the plan.
     */
    readonly sharedSecret?: pulumi.Input<string>;
    /**
     * Hash of the shared secret.
     */
    readonly sharedSecretHash?: pulumi.Input<string>;
    /**
     * URL of the Target VPN gateway with which this VPN tunnel is
     * associated.
     */
    readonly targetVpnGateway?: pulumi.Input<string>;
    /**
     * The unique identifier for the resource. This identifier is defined by the server.
     */
    readonly tunnelId?: pulumi.Input<string>;
    /**
     * URL of the VPN gateway with which this VPN tunnel is associated.
     * This must be used if a High Availability VPN gateway resource is created.
     * This field must reference a `gcp.compute.HaVpnGateway` resource.
     */
    readonly vpnGateway?: pulumi.Input<string>;
    /**
     * The interface ID of the VPN gateway with which this VPN tunnel is associated.
     */
    readonly vpnGatewayInterface?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a VPNTunnel resource.
 */
export interface VPNTunnelArgs {
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * IKE protocol version to use when establishing the VPN tunnel with
     * peer VPN gateway.
     * Acceptable IKE versions are 1 or 2. Default version is 2.
     */
    readonly ikeVersion?: pulumi.Input<number>;
    /**
     * Labels to apply to this VpnTunnel.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Local traffic selector to use when establishing the VPN tunnel with
     * peer VPN gateway. The value should be a CIDR formatted string,
     * for example `192.168.0.0/16`. The ranges should be disjoint.
     * Only IPv4 is supported.
     */
    readonly localTrafficSelectors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63
     * characters long and match the regular expression
     * `a-z?` which means the first character
     * must be a lowercase letter, and all following characters must
     * be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * URL of the peer side external VPN gateway to which this VPN tunnel is connected.
     */
    readonly peerExternalGateway?: pulumi.Input<string>;
    /**
     * The interface ID of the external VPN gateway to which this VPN tunnel is connected.
     */
    readonly peerExternalGatewayInterface?: pulumi.Input<number>;
    /**
     * URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
     * If provided, the VPN tunnel will automatically use the same vpnGatewayInterface
     * ID in the peer GCP VPN gateway.
     * This field must reference a `gcp.compute.HaVpnGateway` resource.
     */
    readonly peerGcpGateway?: pulumi.Input<string>;
    /**
     * IP address of the peer VPN gateway. Only IPv4 is supported.
     */
    readonly peerIp?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region where the tunnel is located. If unset, is set to the region of `targetVpnGateway`.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Remote traffic selector to use when establishing the VPN tunnel with
     * peer VPN gateway. The value should be a CIDR formatted string,
     * for example `192.168.0.0/16`. The ranges should be disjoint.
     * Only IPv4 is supported.
     */
    readonly remoteTrafficSelectors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * URL of router resource to be used for dynamic routing.
     */
    readonly router?: pulumi.Input<string>;
    /**
     * Shared secret used to set the secure session between the Cloud VPN
     * gateway and the peer VPN gateway.  **Note**: This property is sensitive and will not be displayed in the plan.
     */
    readonly sharedSecret: pulumi.Input<string>;
    /**
     * URL of the Target VPN gateway with which this VPN tunnel is
     * associated.
     */
    readonly targetVpnGateway?: pulumi.Input<string>;
    /**
     * URL of the VPN gateway with which this VPN tunnel is associated.
     * This must be used if a High Availability VPN gateway resource is created.
     * This field must reference a `gcp.compute.HaVpnGateway` resource.
     */
    readonly vpnGateway?: pulumi.Input<string>;
    /**
     * The interface ID of the VPN gateway with which this VPN tunnel is associated.
     */
    readonly vpnGatewayInterface?: pulumi.Input<number>;
}
