// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Network endpoint groups (NEGs) are zonal resources that represent
 * collections of IP address and port combinations for GCP resources within a
 * single subnet. Each IP address and port combination is called a network
 * endpoint.
 * 
 * Network endpoint groups can be used as backends in backend services for
 * HTTP(S), TCP proxy, and SSL proxy load balancers. You cannot use NEGs as a
 * backend with internal load balancers. Because NEG backends allow you to
 * specify IP addresses and ports, you can distribute traffic in a granular
 * fashion among applications or containers running within VM instances.
 * 
 * 
 * To get more information about NetworkEndpointGroup, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/networkEndpointGroups)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/load-balancing/docs/negs/)
 * 
 * ## Example Usage - Network Endpoint Group
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const defaultNetwork = new gcp.compute.Network("default", {
 *     autoCreateSubnetworks: false,
 * });
 * const defaultSubnetwork = new gcp.compute.Subnetwork("default", {
 *     ipCidrRange: "10.0.0.0/16",
 *     network: defaultNetwork.selfLink,
 *     region: "us-central1",
 * });
 * const neg = new gcp.compute.NetworkEndpointGroup("neg", {
 *     defaultPort: 90,
 *     network: defaultNetwork.selfLink,
 *     subnetwork: defaultSubnetwork.selfLink,
 *     zone: "us-central1-a",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_network_endpoint_group.html.markdown.
 */
export class NetworkEndpointGroup extends pulumi.CustomResource {
    /**
     * Get an existing NetworkEndpointGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkEndpointGroupState, opts?: pulumi.CustomResourceOptions): NetworkEndpointGroup {
        return new NetworkEndpointGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/networkEndpointGroup:NetworkEndpointGroup';

    /**
     * Returns true if the given object is an instance of NetworkEndpointGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkEndpointGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkEndpointGroup.__pulumiType;
    }

    public readonly defaultPort!: pulumi.Output<number | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly network!: pulumi.Output<string>;
    public readonly networkEndpointType!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    public /*out*/ readonly size!: pulumi.Output<number>;
    public readonly subnetwork!: pulumi.Output<string | undefined>;
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a NetworkEndpointGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkEndpointGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkEndpointGroupArgs | NetworkEndpointGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NetworkEndpointGroupState | undefined;
            inputs["defaultPort"] = state ? state.defaultPort : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["networkEndpointType"] = state ? state.networkEndpointType : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["subnetwork"] = state ? state.subnetwork : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as NetworkEndpointGroupArgs | undefined;
            if (!args || args.network === undefined) {
                throw new Error("Missing required property 'network'");
            }
            inputs["defaultPort"] = args ? args.defaultPort : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["networkEndpointType"] = args ? args.networkEndpointType : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["subnetwork"] = args ? args.subnetwork : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["selfLink"] = undefined /*out*/;
            inputs["size"] = undefined /*out*/;
        }
        super(NetworkEndpointGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkEndpointGroup resources.
 */
export interface NetworkEndpointGroupState {
    readonly defaultPort?: pulumi.Input<number>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly network?: pulumi.Input<string>;
    readonly networkEndpointType?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly size?: pulumi.Input<number>;
    readonly subnetwork?: pulumi.Input<string>;
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkEndpointGroup resource.
 */
export interface NetworkEndpointGroupArgs {
    readonly defaultPort?: pulumi.Input<number>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly network: pulumi.Input<string>;
    readonly networkEndpointType?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly subnetwork?: pulumi.Input<string>;
    readonly zone?: pulumi.Input<string>;
}
