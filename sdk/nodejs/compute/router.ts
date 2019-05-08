// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Represents a Router resource.
 * 
 * 
 * To get more information about Router, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/routers)
 * * How-to Guides
 *     * [Google Cloud Router](https://cloud.google.com/router/docs/)
 * 
 * ## Example Usage - Router Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const foobarNetwork = new gcp.compute.Network("foobar", {
 *     autoCreateSubnetworks: false,
 * });
 * const foobarRouter = new gcp.compute.Router("foobar", {
 *     bgp: {
 *         advertiseMode: "CUSTOM",
 *         advertisedGroups: ["ALL_SUBNETS"],
 *         advertisedIpRanges: [
 *             {
 *                 range: "1.2.3.4",
 *             },
 *             {
 *                 range: "6.7.0.0/16",
 *             },
 *         ],
 *         asn: 64514,
 *     },
 *     network: foobarNetwork.name,
 * });
 * ```
 */
export class Router extends pulumi.CustomResource {
    /**
     * Get an existing Router resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterState, opts?: pulumi.CustomResourceOptions): Router {
        return new Router(name, <any>state, { ...opts, id: id });
    }

    public readonly bgp!: pulumi.Output<{ advertiseMode?: string, advertisedGroups?: string[], advertisedIpRanges?: { description?: string, range?: string }[], asn: number } | undefined>;
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly network!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a Router resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterArgs | RouterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RouterState | undefined;
            inputs["bgp"] = state ? state.bgp : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as RouterArgs | undefined;
            if (!args || args.network === undefined) {
                throw new Error("Missing required property 'network'");
            }
            inputs["bgp"] = args ? args.bgp : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super("gcp:compute/router:Router", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Router resources.
 */
export interface RouterState {
    readonly bgp?: pulumi.Input<{ advertiseMode?: pulumi.Input<string>, advertisedGroups?: pulumi.Input<pulumi.Input<string>[]>, advertisedIpRanges?: pulumi.Input<pulumi.Input<{ description?: pulumi.Input<string>, range?: pulumi.Input<string> }>[]>, asn: pulumi.Input<number> }>;
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly network?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Router resource.
 */
export interface RouterArgs {
    readonly bgp?: pulumi.Input<{ advertiseMode?: pulumi.Input<string>, advertisedGroups?: pulumi.Input<pulumi.Input<string>[]>, advertisedIpRanges?: pulumi.Input<pulumi.Input<{ description?: pulumi.Input<string>, range?: pulumi.Input<string> }>[]>, asn: pulumi.Input<number> }>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly network: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
}
