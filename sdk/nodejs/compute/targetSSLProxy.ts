// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Represents a TargetSslProxy resource, which is used by one or more
 * global forwarding rule to route incoming SSL requests to a backend
 * service.
 *
 * To get more information about TargetSslProxy, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetSslProxies)
 * * How-to Guides
 *     * [Setting Up SSL proxy for Google Cloud Load Balancing](https://cloud.google.com/compute/docs/load-balancing/tcp-ssl/)
 *
 * ## Example Usage
 * ### Target Ssl Proxy Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * import * from "fs";
 *
 * const defaultSSLCertificate = new gcp.compute.SSLCertificate("defaultSSLCertificate", {
 *     privateKey: fs.readFileSync("path/to/private.key"),
 *     certificate: fs.readFileSync("path/to/certificate.crt"),
 * });
 * const defaultHealthCheck = new gcp.compute.HealthCheck("defaultHealthCheck", {
 *     checkIntervalSec: 1,
 *     timeoutSec: 1,
 *     tcpHealthCheck: {
 *         port: "443",
 *     },
 * });
 * const defaultBackendService = new gcp.compute.BackendService("defaultBackendService", {
 *     protocol: "SSL",
 *     healthChecks: [defaultHealthCheck.id],
 * });
 * const defaultTargetSSLProxy = new gcp.compute.TargetSSLProxy("defaultTargetSSLProxy", {
 *     backendService: defaultBackendService.id,
 *     sslCertificates: [defaultSSLCertificate.id],
 * });
 * ```
 *
 * ## Import
 *
 * TargetSslProxy can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/targetSSLProxy:TargetSSLProxy default projects/{{project}}/global/targetSslProxies/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/targetSSLProxy:TargetSSLProxy default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/targetSSLProxy:TargetSSLProxy default {{name}}
 * ```
 */
export class TargetSSLProxy extends pulumi.CustomResource {
    /**
     * Get an existing TargetSSLProxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TargetSSLProxyState, opts?: pulumi.CustomResourceOptions): TargetSSLProxy {
        return new TargetSSLProxy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/targetSSLProxy:TargetSSLProxy';

    /**
     * Returns true if the given object is an instance of TargetSSLProxy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TargetSSLProxy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TargetSSLProxy.__pulumiType;
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
     * A list of SslCertificate resources that are used to authenticate
     * connections between users and the load balancer. At least one
     * SSL certificate must be specified.
     */
    public readonly sslCertificates!: pulumi.Output<string[]>;
    /**
     * A reference to the SslPolicy resource that will be associated with
     * the TargetSslProxy resource. If not set, the TargetSslProxy
     * resource will not have any SSL policy configured.
     */
    public readonly sslPolicy!: pulumi.Output<string | undefined>;

    /**
     * Create a TargetSSLProxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TargetSSLProxyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TargetSSLProxyArgs | TargetSSLProxyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TargetSSLProxyState | undefined;
            inputs["backendService"] = state ? state.backendService : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["proxyHeader"] = state ? state.proxyHeader : undefined;
            inputs["proxyId"] = state ? state.proxyId : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["sslCertificates"] = state ? state.sslCertificates : undefined;
            inputs["sslPolicy"] = state ? state.sslPolicy : undefined;
        } else {
            const args = argsOrState as TargetSSLProxyArgs | undefined;
            if ((!args || args.backendService === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backendService'");
            }
            if ((!args || args.sslCertificates === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sslCertificates'");
            }
            inputs["backendService"] = args ? args.backendService : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["proxyHeader"] = args ? args.proxyHeader : undefined;
            inputs["sslCertificates"] = args ? args.sslCertificates : undefined;
            inputs["sslPolicy"] = args ? args.sslPolicy : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["proxyId"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(TargetSSLProxy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TargetSSLProxy resources.
 */
export interface TargetSSLProxyState {
    /**
     * A reference to the BackendService resource.
     */
    backendService?: pulumi.Input<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    description?: pulumi.Input<string>;
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
     * Specifies the type of proxy header to append before sending data to
     * the backend.
     * Default value is `NONE`.
     * Possible values are `NONE` and `PROXY_V1`.
     */
    proxyHeader?: pulumi.Input<string>;
    /**
     * The unique identifier for the resource.
     */
    proxyId?: pulumi.Input<number>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * A list of SslCertificate resources that are used to authenticate
     * connections between users and the load balancer. At least one
     * SSL certificate must be specified.
     */
    sslCertificates?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A reference to the SslPolicy resource that will be associated with
     * the TargetSslProxy resource. If not set, the TargetSslProxy
     * resource will not have any SSL policy configured.
     */
    sslPolicy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TargetSSLProxy resource.
 */
export interface TargetSSLProxyArgs {
    /**
     * A reference to the BackendService resource.
     */
    backendService: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    description?: pulumi.Input<string>;
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
     * Specifies the type of proxy header to append before sending data to
     * the backend.
     * Default value is `NONE`.
     * Possible values are `NONE` and `PROXY_V1`.
     */
    proxyHeader?: pulumi.Input<string>;
    /**
     * A list of SslCertificate resources that are used to authenticate
     * connections between users and the load balancer. At least one
     * SSL certificate must be specified.
     */
    sslCertificates: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A reference to the SslPolicy resource that will be associated with
     * the TargetSslProxy resource. If not set, the TargetSslProxy
     * resource will not have any SSL policy configured.
     */
    sslPolicy?: pulumi.Input<string>;
}
