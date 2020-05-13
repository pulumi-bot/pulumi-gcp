// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a network peering within GCE. For more information see
 * [the official documentation](https://cloud.google.com/compute/docs/vpc/vpc-peering)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/latest/networks).
 *
 * > Both network must create a peering with each other for the peering
 * to be functional.
 *
 * > Subnets IP ranges across peered VPC networks cannot overlap.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const default = new gcp.compute.Network("default", {autoCreateSubnetworks: "false"});
 * const other = new gcp.compute.Network("other", {autoCreateSubnetworks: "false"});
 * const peering1 = new gcp.compute.NetworkPeering("peering1", {
 *     network: default.selfLink,
 *     peerNetwork: other.selfLink,
 * });
 * const peering2 = new gcp.compute.NetworkPeering("peering2", {
 *     network: other.selfLink,
 *     peerNetwork: default.selfLink,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_network_peering.html.markdown.
 */
export class NetworkPeering extends pulumi.CustomResource {
    /**
     * Get an existing NetworkPeering resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkPeeringState, opts?: pulumi.CustomResourceOptions): NetworkPeering {
        return new NetworkPeering(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/networkPeering:NetworkPeering';

    /**
     * Returns true if the given object is an instance of NetworkPeering.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkPeering {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkPeering.__pulumiType;
    }

    /**
     *
     * Whether to export the custom routes to the peer network. Defaults to `false`.
     */
    public readonly exportCustomRoutes!: pulumi.Output<boolean | undefined>;
    /**
     *
     * Whether to export the custom routes from the peer network. Defaults to `false`.
     */
    public readonly importCustomRoutes!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the peering.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The primary network of the peering.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * The peer network in the peering. The peer network
     * may belong to a different project.
     */
    public readonly peerNetwork!: pulumi.Output<string>;
    /**
     * State for the peering, either `ACTIVE` or `INACTIVE`. The peering is
     * `ACTIVE` when there's a matching configuration in the peer network.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Details about the current state of the peering.
     */
    public /*out*/ readonly stateDetails!: pulumi.Output<string>;

    /**
     * Create a NetworkPeering resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkPeeringArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkPeeringArgs | NetworkPeeringState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NetworkPeeringState | undefined;
            inputs["exportCustomRoutes"] = state ? state.exportCustomRoutes : undefined;
            inputs["importCustomRoutes"] = state ? state.importCustomRoutes : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["peerNetwork"] = state ? state.peerNetwork : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["stateDetails"] = state ? state.stateDetails : undefined;
        } else {
            const args = argsOrState as NetworkPeeringArgs | undefined;
            if (!args || args.network === undefined) {
                throw new Error("Missing required property 'network'");
            }
            if (!args || args.peerNetwork === undefined) {
                throw new Error("Missing required property 'peerNetwork'");
            }
            inputs["exportCustomRoutes"] = args ? args.exportCustomRoutes : undefined;
            inputs["importCustomRoutes"] = args ? args.importCustomRoutes : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["peerNetwork"] = args ? args.peerNetwork : undefined;
            inputs["state"] = undefined /*out*/;
            inputs["stateDetails"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NetworkPeering.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkPeering resources.
 */
export interface NetworkPeeringState {
    /**
     *
     * Whether to export the custom routes to the peer network. Defaults to `false`.
     */
    readonly exportCustomRoutes?: pulumi.Input<boolean>;
    /**
     *
     * Whether to export the custom routes from the peer network. Defaults to `false`.
     */
    readonly importCustomRoutes?: pulumi.Input<boolean>;
    /**
     * Name of the peering.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The primary network of the peering.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * The peer network in the peering. The peer network
     * may belong to a different project.
     */
    readonly peerNetwork?: pulumi.Input<string>;
    /**
     * State for the peering, either `ACTIVE` or `INACTIVE`. The peering is
     * `ACTIVE` when there's a matching configuration in the peer network.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * Details about the current state of the peering.
     */
    readonly stateDetails?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkPeering resource.
 */
export interface NetworkPeeringArgs {
    /**
     *
     * Whether to export the custom routes to the peer network. Defaults to `false`.
     */
    readonly exportCustomRoutes?: pulumi.Input<boolean>;
    /**
     *
     * Whether to export the custom routes from the peer network. Defaults to `false`.
     */
    readonly importCustomRoutes?: pulumi.Input<boolean>;
    /**
     * Name of the peering.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The primary network of the peering.
     */
    readonly network: pulumi.Input<string>;
    /**
     * The peer network in the peering. The peer network
     * may belong to a different project.
     */
    readonly peerNetwork: pulumi.Input<string>;
}
