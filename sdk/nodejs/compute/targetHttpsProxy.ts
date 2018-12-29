// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Represents a TargetHttpsProxy resource, which is used by one or more
 * global forwarding rule to route incoming HTTPS requests to a URL map.
 * 
 * 
 * To get more information about TargetHttpsProxy, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/latest/targetHttpsProxies)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * import * as fs from "fs";
 * 
 * const google_compute_http_health_check_default = new gcp.compute.HttpHealthCheck("default", {
 *     checkIntervalSec: 1,
 *     name: "http-health-check",
 *     requestPath: "/",
 *     timeoutSec: 1,
 * });
 * const google_compute_ssl_certificate_default = new gcp.compute.SSLCertificate("default", {
 *     certificate: fs.readFileSync("path/to/certificate.crt", "utf-8"),
 *     name: "my-certificate",
 *     privateKey: fs.readFileSync("path/to/private.key", "utf-8"),
 * });
 * const google_compute_backend_service_default = new gcp.compute.BackendService("default", {
 *     healthChecks: google_compute_http_health_check_default.selfLink,
 *     name: "backend-service",
 *     portName: "http",
 *     protocol: "HTTP",
 *     timeoutSec: 10,
 * });
 * const google_compute_url_map_default = new gcp.compute.URLMap("default", {
 *     defaultService: google_compute_backend_service_default.selfLink,
 *     description: "a description",
 *     hostRules: [{
 *         hosts: ["mysite.com"],
 *         pathMatcher: "allpaths",
 *     }],
 *     name: "url-map",
 *     pathMatchers: [{
 *         defaultService: google_compute_backend_service_default.selfLink,
 *         name: "allpaths",
 *         pathRules: [{
 *             paths: ["/*"],
 *             service: google_compute_backend_service_default.selfLink,
 *         }],
 *     }],
 * });
 * const google_compute_target_https_proxy_default = new gcp.compute.TargetHttpsProxy("default", {
 *     name: "test-proxy",
 *     sslCertificates: [google_compute_ssl_certificate_default.selfLink],
 *     urlMap: google_compute_url_map_default.selfLink,
 * });
 * ```
 */
export class TargetHttpsProxy extends pulumi.CustomResource {
    /**
     * Get an existing TargetHttpsProxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TargetHttpsProxyState, opts?: pulumi.CustomResourceOptions): TargetHttpsProxy {
        return new TargetHttpsProxy(name, <any>state, { ...opts, id: id });
    }

    public /*out*/ readonly creationTimestamp: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly name: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    public /*out*/ readonly proxyId: pulumi.Output<number>;
    public readonly quicOverride: pulumi.Output<string | undefined>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    public readonly sslCertificates: pulumi.Output<string[]>;
    public readonly sslPolicy: pulumi.Output<string | undefined>;
    public readonly urlMap: pulumi.Output<string>;

    /**
     * Create a TargetHttpsProxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TargetHttpsProxyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TargetHttpsProxyArgs | TargetHttpsProxyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: TargetHttpsProxyState = argsOrState as TargetHttpsProxyState | undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["proxyId"] = state ? state.proxyId : undefined;
            inputs["quicOverride"] = state ? state.quicOverride : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["sslCertificates"] = state ? state.sslCertificates : undefined;
            inputs["sslPolicy"] = state ? state.sslPolicy : undefined;
            inputs["urlMap"] = state ? state.urlMap : undefined;
        } else {
            const args = argsOrState as TargetHttpsProxyArgs | undefined;
            if (!args || args.sslCertificates === undefined) {
                throw new Error("Missing required property 'sslCertificates'");
            }
            if (!args || args.urlMap === undefined) {
                throw new Error("Missing required property 'urlMap'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["quicOverride"] = args ? args.quicOverride : undefined;
            inputs["sslCertificates"] = args ? args.sslCertificates : undefined;
            inputs["sslPolicy"] = args ? args.sslPolicy : undefined;
            inputs["urlMap"] = args ? args.urlMap : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["proxyId"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        super("gcp:compute/targetHttpsProxy:TargetHttpsProxy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TargetHttpsProxy resources.
 */
export interface TargetHttpsProxyState {
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly proxyId?: pulumi.Input<number>;
    readonly quicOverride?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly sslCertificates?: pulumi.Input<pulumi.Input<string>[]>;
    readonly sslPolicy?: pulumi.Input<string>;
    readonly urlMap?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TargetHttpsProxy resource.
 */
export interface TargetHttpsProxyArgs {
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly quicOverride?: pulumi.Input<string>;
    readonly sslCertificates: pulumi.Input<pulumi.Input<string>[]>;
    readonly sslPolicy?: pulumi.Input<string>;
    readonly urlMap: pulumi.Input<string>;
}
