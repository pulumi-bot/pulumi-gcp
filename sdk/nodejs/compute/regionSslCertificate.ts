// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A RegionSslCertificate resource, used for HTTPS load balancing. This resource
 * provides a mechanism to upload an SSL key and certificate to
 * the load balancer to serve secure connections from the user.
 *
 * To get more information about RegionSslCertificate, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/regionSslCertificates)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/load-balancing/docs/ssl-certificates)
 *
 * > **Warning:** All arguments including `certificate` and `privateKey` will be stored in the raw
 * state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).
 *
 * {{% examples %}}
 * ## Example Usage
 * {{% example %}}
 * ### Region Ssl Certificate Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * import * from "fs";
 *
 * const _default = new gcp.compute.RegionSslCertificate("default", {
 *     region: "us-central1",
 *     namePrefix: "my-certificate-",
 *     description: "a description",
 *     privateKey: fs.readFileSync("path/to/private.key"),
 *     certificate: fs.readFileSync("path/to/certificate.crt"),
 * });
 * ```
 * {{% /example %}}
 * {{% example %}}
 * ### Region Ssl Certificate Target Https Proxies
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * import * from "fs";
 *
 * // Using with Region Target HTTPS Proxies
 * //
 * // SSL certificates cannot be updated after creation. In order to apply
 * // the specified configuration, the provider will destroy the existing
 * // resource and create a replacement. To effectively use an SSL
 * // certificate resource with a Target HTTPS Proxy resource, it's
 * // recommended to specify create_before_destroy in a lifecycle block.
 * // Either omit the Instance Template name attribute, specify a partial
 * // name with name_prefix, or use random_id resource. Example:
 * const defaultRegionSslCertificate = new gcp.compute.RegionSslCertificate("defaultRegionSslCertificate", {
 *     region: "us-central1",
 *     namePrefix: "my-certificate-",
 *     privateKey: fs.readFileSync("path/to/private.key"),
 *     certificate: fs.readFileSync("path/to/certificate.crt"),
 * });
 * const defaultRegionHealthCheck = new gcp.compute.RegionHealthCheck("defaultRegionHealthCheck", {
 *     region: "us-central1",
 *     http_health_check: {
 *         port: 80,
 *     },
 * });
 * const defaultRegionBackendService = new gcp.compute.RegionBackendService("defaultRegionBackendService", {
 *     region: "us-central1",
 *     protocol: "HTTP",
 *     timeoutSec: 10,
 *     healthChecks: [defaultRegionHealthCheck.id],
 * });
 * const defaultRegionUrlMap = new gcp.compute.RegionUrlMap("defaultRegionUrlMap", {
 *     region: "us-central1",
 *     description: "a description",
 *     defaultService: defaultRegionBackendService.id,
 *     host_rule: [{
 *         hosts: ["mysite.com"],
 *         pathMatcher: "allpaths",
 *     }],
 *     path_matcher: [{
 *         name: "allpaths",
 *         defaultService: defaultRegionBackendService.id,
 *         path_rule: [{
 *             paths: ["/*"],
 *             service: defaultRegionBackendService.id,
 *         }],
 *     }],
 * });
 * const defaultRegionTargetHttpsProxy = new gcp.compute.RegionTargetHttpsProxy("defaultRegionTargetHttpsProxy", {
 *     region: "us-central1",
 *     urlMap: defaultRegionUrlMap.id,
 *     sslCertificates: [defaultRegionSslCertificate.id],
 * });
 * ```
 * {{% /example %}}
 * {{% /examples %}}
 */
export class RegionSslCertificate extends pulumi.CustomResource {
    /**
     * Get an existing RegionSslCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegionSslCertificateState, opts?: pulumi.CustomResourceOptions): RegionSslCertificate {
        return new RegionSslCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/regionSslCertificate:RegionSslCertificate';

    /**
     * Returns true if the given object is an instance of RegionSslCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionSslCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionSslCertificate.__pulumiType;
    }

    /**
     * The certificate in PEM format.
     * The certificate chain must be no greater than 5 certs long.
     * The chain must include at least one intermediate cert.  **Note**: This property is sensitive and will not be displayed in the plan.
     */
    public readonly certificate!: pulumi.Output<string>;
    /**
     * The unique identifier for the resource.
     */
    public /*out*/ readonly certificateId!: pulumi.Output<number>;
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
     * Creates a unique name beginning with the
     * specified prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * The write-only private key in PEM format.  **Note**: This property is sensitive and will not be displayed in the plan.
     */
    public readonly privateKey!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The Region in which the created regional ssl certificate should reside.
     * If it is not provided, the provider region is used.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a RegionSslCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionSslCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegionSslCertificateArgs | RegionSslCertificateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RegionSslCertificateState | undefined;
            inputs["certificate"] = state ? state.certificate : undefined;
            inputs["certificateId"] = state ? state.certificateId : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["privateKey"] = state ? state.privateKey : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as RegionSslCertificateArgs | undefined;
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
            inputs["region"] = args ? args.region : undefined;
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
        super(RegionSslCertificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegionSslCertificate resources.
 */
export interface RegionSslCertificateState {
    /**
     * The certificate in PEM format.
     * The certificate chain must be no greater than 5 certs long.
     * The chain must include at least one intermediate cert.  **Note**: This property is sensitive and will not be displayed in the plan.
     */
    readonly certificate?: pulumi.Input<string>;
    /**
     * The unique identifier for the resource.
     */
    readonly certificateId?: pulumi.Input<number>;
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
     * Creates a unique name beginning with the
     * specified prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * The write-only private key in PEM format.  **Note**: This property is sensitive and will not be displayed in the plan.
     */
    readonly privateKey?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The Region in which the created regional ssl certificate should reside.
     * If it is not provided, the provider region is used.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RegionSslCertificate resource.
 */
export interface RegionSslCertificateArgs {
    /**
     * The certificate in PEM format.
     * The certificate chain must be no greater than 5 certs long.
     * The chain must include at least one intermediate cert.  **Note**: This property is sensitive and will not be displayed in the plan.
     */
    readonly certificate: pulumi.Input<string>;
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
     * Creates a unique name beginning with the
     * specified prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * The write-only private key in PEM format.  **Note**: This property is sensitive and will not be displayed in the plan.
     */
    readonly privateKey: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The Region in which the created regional ssl certificate should reside.
     * If it is not provided, the provider region is used.
     */
    readonly region?: pulumi.Input<string>;
}
