// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A global network endpoint group contains endpoints that reside outside of Google Cloud.
 * Currently a global network endpoint group can only support a single endpoint.
 * 
 * 
 * To get more information about GlobalNetworkEndpointGroup, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/networkEndpointGroups)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/load-balancing/docs/negs/internet-neg-concepts)
 * 
 * ## Example Usage - Global Network Endpoint Group
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const neg = new gcp.compute.GlobalNetworkEndpointGroup("neg", {
 *     defaultPort: 90,
 *     name: "my-lb-neg",
 *     networkEndpointType: "INTERNET_FQDN_PORT",
 * });
 * ```
 * ## Example Usage - Global Network Endpoint Group Ip Address
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const neg = new gcp.compute.GlobalNetworkEndpointGroup("neg", {
 *     defaultPort: 90,
 *     name: "my-lb-neg",
 *     networkEndpointType: "INTERNET_IP_PORT",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_global_network_endpoint_group.html.markdown.
 */
export class GlobalNetworkEndpointGroup extends pulumi.CustomResource {
    /**
     * Get an existing GlobalNetworkEndpointGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GlobalNetworkEndpointGroupState, opts?: pulumi.CustomResourceOptions): GlobalNetworkEndpointGroup {
        return new GlobalNetworkEndpointGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/globalNetworkEndpointGroup:GlobalNetworkEndpointGroup';

    /**
     * Returns true if the given object is an instance of GlobalNetworkEndpointGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GlobalNetworkEndpointGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GlobalNetworkEndpointGroup.__pulumiType;
    }

    /**
     * The default port used if the port number is not specified in the
     * network endpoint.
     */
    public readonly defaultPort!: pulumi.Output<number | undefined>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Type of network endpoints in this network endpoint group. Supported values are:
     * * INTERNET_IP_PORT
     * * INTERNET_FQDN_PORT
     */
    public readonly networkEndpointType!: pulumi.Output<string>;
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
     * Create a GlobalNetworkEndpointGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GlobalNetworkEndpointGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GlobalNetworkEndpointGroupArgs | GlobalNetworkEndpointGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GlobalNetworkEndpointGroupState | undefined;
            inputs["defaultPort"] = state ? state.defaultPort : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkEndpointType"] = state ? state.networkEndpointType : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as GlobalNetworkEndpointGroupArgs | undefined;
            if (!args || args.networkEndpointType === undefined) {
                throw new Error("Missing required property 'networkEndpointType'");
            }
            inputs["defaultPort"] = args ? args.defaultPort : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkEndpointType"] = args ? args.networkEndpointType : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GlobalNetworkEndpointGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GlobalNetworkEndpointGroup resources.
 */
export interface GlobalNetworkEndpointGroupState {
    /**
     * The default port used if the port number is not specified in the
     * network endpoint.
     */
    readonly defaultPort?: pulumi.Input<number>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Type of network endpoints in this network endpoint group. Supported values are:
     * * INTERNET_IP_PORT
     * * INTERNET_FQDN_PORT
     */
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
}

/**
 * The set of arguments for constructing a GlobalNetworkEndpointGroup resource.
 */
export interface GlobalNetworkEndpointGroupArgs {
    /**
     * The default port used if the port number is not specified in the
     * network endpoint.
     */
    readonly defaultPort?: pulumi.Input<number>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Type of network endpoints in this network endpoint group. Supported values are:
     * * INTERNET_IP_PORT
     * * INTERNET_FQDN_PORT
     */
    readonly networkEndpointType: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
