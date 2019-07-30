// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * An SslCertificate resource, used for HTTPS load balancing. This resource
 * provides a mechanism to upload an SSL key and certificate to
 * the load balancer to serve secure connections from the user.
 * 
 * 
 * To get more information about SslCertificate, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/sslCertificates)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/load-balancing/docs/ssl-certificates)
 * 
 * ## Example Usage - Ssl Certificate Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const defaultSSLCertificate = new gcp.compute.SSLCertificate("default", {
 *     certificate: fs.readFileSync("path/to/certificate.crt", "utf-8"),
 *     description: "a description",
 *     namePrefix: "my-certificate-",
 *     privateKey: fs.readFileSync("path/to/private.key", "utf-8"),
 * });
 * ```
 * ## Example Usage - Ssl Certificate Random Provider
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as gcp from "@pulumi/gcp";
 * import * as random from "@pulumi/random";
 * 
 * const certificate = new random.RandomId("certificate", {
 *     byteLength: 4,
 *     // For security, do not expose raw certificate values in the output
 *     keepers: {
 *         certificate: (() => {
 *             throw "tf2pulumi error: NYI: call to base64sha256";
 *             return (() => { throw "NYI: call to base64sha256"; })();
 *         })(),
 *         private_key: (() => {
 *             throw "tf2pulumi error: NYI: call to base64sha256";
 *             return (() => { throw "NYI: call to base64sha256"; })();
 *         })(),
 *     },
 *     prefix: "my-certificate-",
 * });
 * // You may also want to control name generation explicitly:
 * const defaultSSLCertificate = new gcp.compute.SSLCertificate("default", {
 *     certificate: fs.readFileSync("path/to/certificate.crt", "utf-8"),
 *     privateKey: fs.readFileSync("path/to/private.key", "utf-8"),
 * });
 * ```
 * ## Example Usage - Ssl Certificate Target Https Proxies
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const defaultHttpHealthCheck = new gcp.compute.HttpHealthCheck("default", {
 *     checkIntervalSec: 1,
 *     requestPath: "/",
 *     timeoutSec: 1,
 * });
 * const defaultSSLCertificate = new gcp.compute.SSLCertificate("default", {
 *     certificate: fs.readFileSync("path/to/certificate.crt", "utf-8"),
 *     namePrefix: "my-certificate-",
 *     privateKey: fs.readFileSync("path/to/private.key", "utf-8"),
 * });
 * const defaultBackendService = new gcp.compute.BackendService("default", {
 *     healthChecks: defaultHttpHealthCheck.selfLink,
 *     portName: "http",
 *     protocol: "HTTP",
 *     timeoutSec: 10,
 * });
 * const defaultURLMap = new gcp.compute.URLMap("default", {
 *     defaultService: defaultBackendService.selfLink,
 *     description: "a description",
 *     hostRules: [{
 *         hosts: ["mysite.com"],
 *         pathMatcher: "allpaths",
 *     }],
 *     pathMatchers: [{
 *         defaultService: defaultBackendService.selfLink,
 *         name: "allpaths",
 *         pathRules: [{
 *             paths: ["/*"],
 *             service: defaultBackendService.selfLink,
 *         }],
 *     }],
 * });
 * const defaultTargetHttpsProxy = new gcp.compute.TargetHttpsProxy("default", {
 *     sslCertificates: [defaultSSLCertificate.selfLink],
 *     urlMap: defaultURLMap.selfLink,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_ssl_certificate.html.markdown.
 */
export class SSLCertificate extends pulumi.CustomResource {
    /**
     * Get an existing SSLCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SSLCertificateState, opts?: pulumi.CustomResourceOptions): SSLCertificate {
        return new SSLCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/sSLCertificate:SSLCertificate';

    /**
     * Returns true if the given object is an instance of SSLCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SSLCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SSLCertificate.__pulumiType;
    }

    public readonly certificate!: pulumi.Output<string>;
    public /*out*/ readonly certificateId!: pulumi.Output<number>;
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the
     * specified prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string>;
    public readonly privateKey!: pulumi.Output<string>;
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
     * Create a SSLCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SSLCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SSLCertificateArgs | SSLCertificateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SSLCertificateState | undefined;
            inputs["certificate"] = state ? state.certificate : undefined;
            inputs["certificateId"] = state ? state.certificateId : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["privateKey"] = state ? state.privateKey : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as SSLCertificateArgs | undefined;
            if (!args || args.certificate === undefined) {
                throw new Error("Missing required property 'certificate'");
            }
            if (!args || args.privateKey === undefined) {
                throw new Error("Missing required property 'privateKey'");
            }
            inputs["certificate"] = args ? args.certificate : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["privateKey"] = args ? args.privateKey : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["certificateId"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SSLCertificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SSLCertificate resources.
 */
export interface SSLCertificateState {
    readonly certificate?: pulumi.Input<string>;
    readonly certificateId?: pulumi.Input<number>;
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the
     * specified prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    readonly privateKey?: pulumi.Input<string>;
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
 * The set of arguments for constructing a SSLCertificate resource.
 */
export interface SSLCertificateArgs {
    readonly certificate: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the
     * specified prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    readonly privateKey: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
