// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manage a network peering's route settings without managing the peering as
 * a whole. This resource is primarily intended for use with GCP-generated
 * peerings that shouldn't otherwise be managed by other tools. Deleting this
 * resource is a no-op and the peering will not be modified.
 *
 * To get more information about NetworkPeeringRoutesConfig, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/networks/updatePeering)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/vpc/docs/vpc-peering)
 *
 * ## Example Usage
 * ### Network Peering Routes Config Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const networkPrimary = new gcp.compute.Network("networkPrimary", {autoCreateSubnetworks: "false"});
 * const networkSecondary = new gcp.compute.Network("networkSecondary", {autoCreateSubnetworks: "false"});
 * const peeringPrimary = new gcp.compute.NetworkPeering("peeringPrimary", {
 *     network: networkPrimary.id,
 *     peerNetwork: networkSecondary.id,
 *     importCustomRoutes: true,
 *     exportCustomRoutes: true,
 * });
 * const peeringPrimaryRoutes = new gcp.compute.NetworkPeeringRoutesConfig("peeringPrimaryRoutes", {
 *     peering: peeringPrimary.name,
 *     network: networkPrimary.name,
 *     importCustomRoutes: true,
 *     exportCustomRoutes: true,
 * });
 * const peeringSecondary = new gcp.compute.NetworkPeering("peeringSecondary", {
 *     network: networkSecondary.id,
 *     peerNetwork: networkPrimary.id,
 * });
 * ```
 *
 * ## Import
 *
 * NetworkPeeringRoutesConfig can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig default projects/{{project}}/global/networks/{{network}}/networkPeerings/{{peering}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig default {{project}}/{{network}}/{{peering}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig default {{network}}/{{peering}}
 * ```
 */
export class NetworkPeeringRoutesConfig extends pulumi.CustomResource {
    /**
     * Get an existing NetworkPeeringRoutesConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkPeeringRoutesConfigState, opts?: pulumi.CustomResourceOptions): NetworkPeeringRoutesConfig {
        return new NetworkPeeringRoutesConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig';

    /**
     * Returns true if the given object is an instance of NetworkPeeringRoutesConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkPeeringRoutesConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkPeeringRoutesConfig.__pulumiType;
    }

    /**
     * Whether to export the custom routes to the peer network.
     */
    public readonly exportCustomRoutes!: pulumi.Output<boolean>;
    /**
     * Whether to import the custom routes to the peer network.
     */
    public readonly importCustomRoutes!: pulumi.Output<boolean>;
    /**
     * The name of the primary network for the peering.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * Name of the peering.
     */
    public readonly peering!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a NetworkPeeringRoutesConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkPeeringRoutesConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkPeeringRoutesConfigArgs | NetworkPeeringRoutesConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NetworkPeeringRoutesConfigState | undefined;
            inputs["exportCustomRoutes"] = state ? state.exportCustomRoutes : undefined;
            inputs["importCustomRoutes"] = state ? state.importCustomRoutes : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["peering"] = state ? state.peering : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as NetworkPeeringRoutesConfigArgs | undefined;
            if ((!args || args.exportCustomRoutes === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'exportCustomRoutes'");
            }
            if ((!args || args.importCustomRoutes === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'importCustomRoutes'");
            }
            if ((!args || args.network === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'network'");
            }
            if ((!args || args.peering === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'peering'");
            }
            inputs["exportCustomRoutes"] = args ? args.exportCustomRoutes : undefined;
            inputs["importCustomRoutes"] = args ? args.importCustomRoutes : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["peering"] = args ? args.peering : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NetworkPeeringRoutesConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkPeeringRoutesConfig resources.
 */
export interface NetworkPeeringRoutesConfigState {
    /**
     * Whether to export the custom routes to the peer network.
     */
    readonly exportCustomRoutes?: pulumi.Input<boolean>;
    /**
     * Whether to import the custom routes to the peer network.
     */
    readonly importCustomRoutes?: pulumi.Input<boolean>;
    /**
     * The name of the primary network for the peering.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * Name of the peering.
     */
    readonly peering?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkPeeringRoutesConfig resource.
 */
export interface NetworkPeeringRoutesConfigArgs {
    /**
     * Whether to export the custom routes to the peer network.
     */
    readonly exportCustomRoutes: pulumi.Input<boolean>;
    /**
     * Whether to import the custom routes to the peer network.
     */
    readonly importCustomRoutes: pulumi.Input<boolean>;
    /**
     * The name of the primary network for the peering.
     */
    readonly network: pulumi.Input<string>;
    /**
     * Name of the peering.
     */
    readonly peering: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
