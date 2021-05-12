// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a VPC network or legacy network resource on GCP.
 *
 * To get more information about Network, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/networks)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/vpc/docs/vpc)
 *
 * ## Example Usage
 * ### Network Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const vpcNetwork = new gcp.compute.Network("vpc_network", {});
 * ```
 * ### Network Custom Mtu
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const vpcNetwork = new gcp.compute.Network("vpc_network", {
 *     mtu: 1500,
 * });
 * ```
 *
 * ## Import
 *
 * Network can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/network:Network default projects/{{project}}/global/networks/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/network:Network default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/network:Network default {{name}}
 * ```
 */
export class Network extends pulumi.CustomResource {
    /**
     * Get an existing Network resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkState, opts?: pulumi.CustomResourceOptions): Network {
        return new Network(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/network:Network';

    /**
     * Returns true if the given object is an instance of Network.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Network {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Network.__pulumiType;
    }

    /**
     * When set to `true`, the network is created in "auto subnet mode" and
     * it will create a subnet for each region automatically across the
     * `10.128.0.0/9` address range.
     * When set to `false`, the network is created in "custom subnet mode" so
     * the user can explicitly connect subnetwork resources.
     */
    public readonly autoCreateSubnetworks!: pulumi.Output<boolean | undefined>;
    /**
     * If set to `true`, default routes (`0.0.0.0/0`) will be deleted
     * immediately after network creation. Defaults to `false`.
     */
    public readonly deleteDefaultRoutesOnCreate!: pulumi.Output<boolean | undefined>;
    /**
     * An optional description of this resource. The resource must be
     * recreated to modify this field.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The gateway address for default routing out of the network. This value is selected by GCP.
     */
    public /*out*/ readonly gatewayIpv4!: pulumi.Output<string>;
    /**
     * Maximum Transmission Unit in bytes. The minimum value for this field is 1460
     * and the maximum value is 1500 bytes.
     */
    public readonly mtu!: pulumi.Output<number>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The network-wide routing mode to use. If set to `REGIONAL`, this
     * network's cloud routers will only advertise routes with subnetworks
     * of this network in the same region as the router. If set to `GLOBAL`,
     * this network's cloud routers will advertise routes with all
     * subnetworks of this network, across regions.
     * Possible values are `REGIONAL` and `GLOBAL`.
     */
    public readonly routingMode!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a Network resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkArgs | NetworkState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkState | undefined;
            inputs["autoCreateSubnetworks"] = state ? state.autoCreateSubnetworks : undefined;
            inputs["deleteDefaultRoutesOnCreate"] = state ? state.deleteDefaultRoutesOnCreate : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["gatewayIpv4"] = state ? state.gatewayIpv4 : undefined;
            inputs["mtu"] = state ? state.mtu : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["routingMode"] = state ? state.routingMode : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as NetworkArgs | undefined;
            inputs["autoCreateSubnetworks"] = args ? args.autoCreateSubnetworks : undefined;
            inputs["deleteDefaultRoutesOnCreate"] = args ? args.deleteDefaultRoutesOnCreate : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["mtu"] = args ? args.mtu : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["routingMode"] = args ? args.routingMode : undefined;
            inputs["gatewayIpv4"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Network.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Network resources.
 */
export interface NetworkState {
    /**
     * When set to `true`, the network is created in "auto subnet mode" and
     * it will create a subnet for each region automatically across the
     * `10.128.0.0/9` address range.
     * When set to `false`, the network is created in "custom subnet mode" so
     * the user can explicitly connect subnetwork resources.
     */
    autoCreateSubnetworks?: pulumi.Input<boolean>;
    /**
     * If set to `true`, default routes (`0.0.0.0/0`) will be deleted
     * immediately after network creation. Defaults to `false`.
     */
    deleteDefaultRoutesOnCreate?: pulumi.Input<boolean>;
    /**
     * An optional description of this resource. The resource must be
     * recreated to modify this field.
     */
    description?: pulumi.Input<string>;
    /**
     * The gateway address for default routing out of the network. This value is selected by GCP.
     */
    gatewayIpv4?: pulumi.Input<string>;
    /**
     * Maximum Transmission Unit in bytes. The minimum value for this field is 1460
     * and the maximum value is 1500 bytes.
     */
    mtu?: pulumi.Input<number>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The network-wide routing mode to use. If set to `REGIONAL`, this
     * network's cloud routers will only advertise routes with subnetworks
     * of this network in the same region as the router. If set to `GLOBAL`,
     * this network's cloud routers will advertise routes with all
     * subnetworks of this network, across regions.
     * Possible values are `REGIONAL` and `GLOBAL`.
     */
    routingMode?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Network resource.
 */
export interface NetworkArgs {
    /**
     * When set to `true`, the network is created in "auto subnet mode" and
     * it will create a subnet for each region automatically across the
     * `10.128.0.0/9` address range.
     * When set to `false`, the network is created in "custom subnet mode" so
     * the user can explicitly connect subnetwork resources.
     */
    autoCreateSubnetworks?: pulumi.Input<boolean>;
    /**
     * If set to `true`, default routes (`0.0.0.0/0`) will be deleted
     * immediately after network creation. Defaults to `false`.
     */
    deleteDefaultRoutesOnCreate?: pulumi.Input<boolean>;
    /**
     * An optional description of this resource. The resource must be
     * recreated to modify this field.
     */
    description?: pulumi.Input<string>;
    /**
     * Maximum Transmission Unit in bytes. The minimum value for this field is 1460
     * and the maximum value is 1500 bytes.
     */
    mtu?: pulumi.Input<number>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The network-wide routing mode to use. If set to `REGIONAL`, this
     * network's cloud routers will only advertise routes with subnetworks
     * of this network in the same region as the router. If set to `GLOBAL`,
     * this network's cloud routers will advertise routes with all
     * subnetworks of this network, across regions.
     * Possible values are `REGIONAL` and `GLOBAL`.
     */
    routingMode?: pulumi.Input<string>;
}
