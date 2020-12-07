// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Cloud Router interface. For more information see
 * [the official documentation](https://cloud.google.com/compute/docs/cloudrouter)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/latest/routers).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const foobar = new gcp.compute.RouterInterface("foobar", {
 *     ipRange: "169.254.1.1/30",
 *     region: "us-central1",
 *     router: "router-1",
 *     vpnTunnel: "tunnel-1",
 * });
 * ```
 *
 * ## Import
 *
 * Router interfaces can be imported using the `region`, `router`, and `name`, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:compute/routerInterface:RouterInterface foobar us-central1/router-1/interface-1
 * ```
 */
export class RouterInterface extends pulumi.CustomResource {
    /**
     * Get an existing RouterInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterInterfaceState, opts?: pulumi.CustomResourceOptions): RouterInterface {
        return new RouterInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/routerInterface:RouterInterface';

    /**
     * Returns true if the given object is an instance of RouterInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterInterface.__pulumiType;
    }

    /**
     * The name or resource link to the
     * VLAN interconnect for this interface. Changing this forces a new interface to
     * be created. Only one of `vpnTunnel` and `interconnectAttachment` can be
     * specified.
     */
    public readonly interconnectAttachment!: pulumi.Output<string | undefined>;
    /**
     * IP address and range of the interface. The IP range must be
     * in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
     */
    public readonly ipRange!: pulumi.Output<string | undefined>;
    /**
     * A unique name for the interface, required by GCE. Changing
     * this forces a new interface to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which this interface's router belongs. If it
     * is not provided, the provider project is used. Changing this forces a new interface to be created.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The region this interface's router sits in. If not specified,
     * the project region will be used. Changing this forces a new interface to be
     * created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The name of the router this interface will be attached to.
     * Changing this forces a new interface to be created.
     */
    public readonly router!: pulumi.Output<string>;
    /**
     * The name or resource link to the VPN tunnel this
     * interface will be linked to. Changing this forces a new interface to be created. Only
     * one of `vpnTunnel` and `interconnectAttachment` can be specified.
     */
    public readonly vpnTunnel!: pulumi.Output<string | undefined>;

    /**
     * Create a RouterInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouterInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterInterfaceArgs | RouterInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RouterInterfaceState | undefined;
            inputs["interconnectAttachment"] = state ? state.interconnectAttachment : undefined;
            inputs["ipRange"] = state ? state.ipRange : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["router"] = state ? state.router : undefined;
            inputs["vpnTunnel"] = state ? state.vpnTunnel : undefined;
        } else {
            const args = argsOrState as RouterInterfaceArgs | undefined;
            if ((!args || args.router === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'router'");
            }
            inputs["interconnectAttachment"] = args ? args.interconnectAttachment : undefined;
            inputs["ipRange"] = args ? args.ipRange : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["router"] = args ? args.router : undefined;
            inputs["vpnTunnel"] = args ? args.vpnTunnel : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RouterInterface.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterInterface resources.
 */
export interface RouterInterfaceState {
    /**
     * The name or resource link to the
     * VLAN interconnect for this interface. Changing this forces a new interface to
     * be created. Only one of `vpnTunnel` and `interconnectAttachment` can be
     * specified.
     */
    readonly interconnectAttachment?: pulumi.Input<string>;
    /**
     * IP address and range of the interface. The IP range must be
     * in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
     */
    readonly ipRange?: pulumi.Input<string>;
    /**
     * A unique name for the interface, required by GCE. Changing
     * this forces a new interface to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which this interface's router belongs. If it
     * is not provided, the provider project is used. Changing this forces a new interface to be created.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region this interface's router sits in. If not specified,
     * the project region will be used. Changing this forces a new interface to be
     * created.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The name of the router this interface will be attached to.
     * Changing this forces a new interface to be created.
     */
    readonly router?: pulumi.Input<string>;
    /**
     * The name or resource link to the VPN tunnel this
     * interface will be linked to. Changing this forces a new interface to be created. Only
     * one of `vpnTunnel` and `interconnectAttachment` can be specified.
     */
    readonly vpnTunnel?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouterInterface resource.
 */
export interface RouterInterfaceArgs {
    /**
     * The name or resource link to the
     * VLAN interconnect for this interface. Changing this forces a new interface to
     * be created. Only one of `vpnTunnel` and `interconnectAttachment` can be
     * specified.
     */
    readonly interconnectAttachment?: pulumi.Input<string>;
    /**
     * IP address and range of the interface. The IP range must be
     * in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
     */
    readonly ipRange?: pulumi.Input<string>;
    /**
     * A unique name for the interface, required by GCE. Changing
     * this forces a new interface to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which this interface's router belongs. If it
     * is not provided, the provider project is used. Changing this forces a new interface to be created.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region this interface's router sits in. If not specified,
     * the project region will be used. Changing this forces a new interface to be
     * created.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The name of the router this interface will be attached to.
     * Changing this forces a new interface to be created.
     */
    readonly router: pulumi.Input<string>;
    /**
     * The name or resource link to the VPN tunnel this
     * interface will be linked to. Changing this forces a new interface to be created. Only
     * one of `vpnTunnel` and `interconnectAttachment` can be specified.
     */
    readonly vpnTunnel?: pulumi.Input<string>;
}
