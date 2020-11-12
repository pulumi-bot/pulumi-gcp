// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Represents a TargetTcpProxy resource, which is used by one or more
 * global forwarding rule to route incoming TCP requests to a Backend
 * service.
 *
 * To get more information about TargetTcpProxy, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetTcpProxies)
 * * How-to Guides
 *     * [Setting Up TCP proxy for Google Cloud Load Balancing](https://cloud.google.com/compute/docs/load-balancing/tcp-ssl/tcp-proxy)
 *
 * ## Example Usage
 *
 * ## Import
 *
 * TargetTcpProxy can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/targetTCPProxy:TargetTCPProxy default projects/{{project}}/global/targetTcpProxies/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/targetTCPProxy:TargetTCPProxy default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/targetTCPProxy:TargetTCPProxy default {{name}}
 * ```
 */
export class TargetTCPProxy extends pulumi.CustomResource {
    /**
     * Get an existing TargetTCPProxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TargetTCPProxyState, opts?: pulumi.CustomResourceOptions): TargetTCPProxy {
        return new TargetTCPProxy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/targetTCPProxy:TargetTCPProxy';

    /**
     * Returns true if the given object is an instance of TargetTCPProxy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TargetTCPProxy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TargetTCPProxy.__pulumiType;
    }

    /**
     * A reference to the BackendService resource.
     */
    public readonly backendService!: pulumi.Output<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
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
     * Specifies the type of proxy header to append before sending data to
     * the backend.
     * Default value is `NONE`.
     * Possible values are `NONE` and `PROXY_V1`.
     */
    public readonly proxyHeader!: pulumi.Output<string | undefined>;
    /**
     * The unique identifier for the resource.
     */
    public /*out*/ readonly proxyId!: pulumi.Output<number>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a TargetTCPProxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TargetTCPProxyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TargetTCPProxyArgs | TargetTCPProxyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TargetTCPProxyState | undefined;
            inputs["backendService"] = state ? state.backendService : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["proxyHeader"] = state ? state.proxyHeader : undefined;
            inputs["proxyId"] = state ? state.proxyId : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as TargetTCPProxyArgs | undefined;
            if (!args || args.backendService === undefined) {
                throw new Error("Missing required property 'backendService'");
            }
            inputs["backendService"] = args ? args.backendService : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["proxyHeader"] = args ? args.proxyHeader : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["proxyId"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TargetTCPProxy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TargetTCPProxy resources.
 */
export interface TargetTCPProxyState {
    /**
     * A reference to the BackendService resource.
     */
    readonly backendService?: pulumi.Input<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Specifies the type of proxy header to append before sending data to
     * the backend.
     * Default value is `NONE`.
     * Possible values are `NONE` and `PROXY_V1`.
     */
    readonly proxyHeader?: pulumi.Input<string>;
    /**
     * The unique identifier for the resource.
     */
    readonly proxyId?: pulumi.Input<number>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TargetTCPProxy resource.
 */
export interface TargetTCPProxyArgs {
    /**
     * A reference to the BackendService resource.
     */
    readonly backendService: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Specifies the type of proxy header to append before sending data to
     * the backend.
     * Default value is `NONE`.
     * Possible values are `NONE` and `PROXY_V1`.
     */
    readonly proxyHeader?: pulumi.Input<string>;
}
