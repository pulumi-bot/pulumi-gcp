// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Health Checks determine whether instances are responsive and able to do work.
 * They are an important part of a comprehensive load balancing configuration,
 * as they enable monitoring instances behind load balancers.
 *
 * Health Checks poll instances at a specified interval. Instances that
 * do not respond successfully to some number of probes in a row are marked
 * as unhealthy. No new connections are sent to unhealthy instances,
 * though existing connections will continue. The health check will
 * continue to poll unhealthy instances. If an instance later responds
 * successfully to some number of consecutive probes, it is marked
 * healthy again and can receive new connections.
 *
 *
 * To get more information about HealthCheck, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/healthChecks)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/load-balancing/docs/health-checks)
 *
 * ## Example Usage - Health Check Tcp
 * {{% example %}}
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const tcpHealthCheck = new gcp.compute.HealthCheck("tcp-health-check", {
 *     checkIntervalSec: 1,
 *     tcpHealthCheck: {
 *         port: 80,
 *     },
 *     timeoutSec: 1,
 * });
 * ```
 * {{% /example %}}
 * ## Example Usage - Health Check Tcp Full
 * {{% example %}}
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const tcpHealthCheck = new gcp.compute.HealthCheck("tcp-health-check", {
 *     checkIntervalSec: 1,
 *     description: "Health check via tcp",
 *     healthyThreshold: 4,
 *     tcpHealthCheck: {
 *         portName: "health-check-port",
 *         portSpecification: "USE_NAMED_PORT",
 *         proxyHeader: "NONE",
 *         request: "ARE YOU HEALTHY?",
 *         response: "I AM HEALTHY",
 *     },
 *     timeoutSec: 1,
 *     unhealthyThreshold: 5,
 * });
 * ```
 * {{% /example %}}
 * ## Example Usage - Health Check Ssl
 * {{% example %}}
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const sslHealthCheck = new gcp.compute.HealthCheck("ssl-health-check", {
 *     checkIntervalSec: 1,
 *     sslHealthCheck: {
 *         port: 443,
 *     },
 *     timeoutSec: 1,
 * });
 * ```
 * {{% /example %}}
 * ## Example Usage - Health Check Ssl Full
 * {{% example %}}
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const sslHealthCheck = new gcp.compute.HealthCheck("ssl-health-check", {
 *     checkIntervalSec: 1,
 *     description: "Health check via ssl",
 *     healthyThreshold: 4,
 *     sslHealthCheck: {
 *         portName: "health-check-port",
 *         portSpecification: "USE_NAMED_PORT",
 *         proxyHeader: "NONE",
 *         request: "ARE YOU HEALTHY?",
 *         response: "I AM HEALTHY",
 *     },
 *     timeoutSec: 1,
 *     unhealthyThreshold: 5,
 * });
 * ```
 * {{% /example %}}
 * ## Example Usage - Health Check Http
 * {{% example %}}
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const httpHealthCheck = new gcp.compute.HealthCheck("http-health-check", {
 *     checkIntervalSec: 1,
 *     httpHealthCheck: {
 *         port: 80,
 *     },
 *     timeoutSec: 1,
 * });
 * ```
 * {{% /example %}}
 * ## Example Usage - Health Check Http Full
 * {{% example %}}
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const httpHealthCheck = new gcp.compute.HealthCheck("http-health-check", {
 *     checkIntervalSec: 1,
 *     description: "Health check via http",
 *     healthyThreshold: 4,
 *     httpHealthCheck: {
 *         host: "1.2.3.4",
 *         portName: "health-check-port",
 *         portSpecification: "USE_NAMED_PORT",
 *         proxyHeader: "NONE",
 *         requestPath: "/mypath",
 *         response: "I AM HEALTHY",
 *     },
 *     timeoutSec: 1,
 *     unhealthyThreshold: 5,
 * });
 * ```
 * {{% /example %}}
 * ## Example Usage - Health Check Https
 * {{% example %}}
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const httpsHealthCheck = new gcp.compute.HealthCheck("https-health-check", {
 *     checkIntervalSec: 1,
 *     httpsHealthCheck: {
 *         port: 443,
 *     },
 *     timeoutSec: 1,
 * });
 * ```
 * {{% /example %}}
 * ## Example Usage - Health Check Https Full
 * {{% example %}}
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const httpsHealthCheck = new gcp.compute.HealthCheck("https-health-check", {
 *     checkIntervalSec: 1,
 *     description: "Health check via https",
 *     healthyThreshold: 4,
 *     httpsHealthCheck: {
 *         host: "1.2.3.4",
 *         portName: "health-check-port",
 *         portSpecification: "USE_NAMED_PORT",
 *         proxyHeader: "NONE",
 *         requestPath: "/mypath",
 *         response: "I AM HEALTHY",
 *     },
 *     timeoutSec: 1,
 *     unhealthyThreshold: 5,
 * });
 * ```
 * {{% /example %}}
 * ## Example Usage - Health Check Http2
 * {{% example %}}
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const http2HealthCheck = new gcp.compute.HealthCheck("http2-health-check", {
 *     checkIntervalSec: 1,
 *     http2HealthCheck: {
 *         port: 443,
 *     },
 *     timeoutSec: 1,
 * });
 * ```
 * {{% /example %}}
 * ## Example Usage - Health Check Http2 Full
 * {{% example %}}
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const http2HealthCheck = new gcp.compute.HealthCheck("http2-health-check", {
 *     checkIntervalSec: 1,
 *     description: "Health check via http2",
 *     healthyThreshold: 4,
 *     http2HealthCheck: {
 *         host: "1.2.3.4",
 *         portName: "health-check-port",
 *         portSpecification: "USE_NAMED_PORT",
 *         proxyHeader: "NONE",
 *         requestPath: "/mypath",
 *         response: "I AM HEALTHY",
 *     },
 *     timeoutSec: 1,
 *     unhealthyThreshold: 5,
 * });
 * ```
 * {{% /example %}}
 * ## Example Usage - Health Check With Logging
 * {{% example %}}
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const health-check-with-logging = new gcp.compute.HealthCheck("health-check-with-logging", {
 *     timeoutSec: 1,
 *     checkIntervalSec: 1,
 *     tcp_health_check: {
 *         port: "22",
 *     },
 *     log_config: {
 *         enable: true,
 *     },
 * });
 * ```
 *
 * {{% /example %}}
 */
export class HealthCheck extends pulumi.CustomResource {
    /**
     * Get an existing HealthCheck resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HealthCheckState, opts?: pulumi.CustomResourceOptions): HealthCheck {
        return new HealthCheck(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/healthCheck:HealthCheck';

    /**
     * Returns true if the given object is an instance of HealthCheck.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HealthCheck {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HealthCheck.__pulumiType;
    }

    /**
     * How often (in seconds) to send a health check. The default value is 5
     * seconds.
     */
    public readonly checkIntervalSec!: pulumi.Output<number | undefined>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A so-far unhealthy instance will be marked healthy after this many
     * consecutive successes. The default value is 2.
     */
    public readonly healthyThreshold!: pulumi.Output<number | undefined>;
    /**
     * A nested object resource  Structure is documented below.
     */
    public readonly http2HealthCheck!: pulumi.Output<outputs.compute.HealthCheckHttp2HealthCheck | undefined>;
    /**
     * A nested object resource  Structure is documented below.
     */
    public readonly httpHealthCheck!: pulumi.Output<outputs.compute.HealthCheckHttpHealthCheck | undefined>;
    /**
     * A nested object resource  Structure is documented below.
     */
    public readonly httpsHealthCheck!: pulumi.Output<outputs.compute.HealthCheckHttpsHealthCheck | undefined>;
    /**
     * Configure logging on this health check.  Structure is documented below.
     */
    public readonly logConfig!: pulumi.Output<outputs.compute.HealthCheckLogConfig | undefined>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
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
     * A nested object resource  Structure is documented below.
     */
    public readonly sslHealthCheck!: pulumi.Output<outputs.compute.HealthCheckSslHealthCheck | undefined>;
    /**
     * A nested object resource  Structure is documented below.
     */
    public readonly tcpHealthCheck!: pulumi.Output<outputs.compute.HealthCheckTcpHealthCheck | undefined>;
    /**
     * How long (in seconds) to wait before claiming failure.
     * The default value is 5 seconds.  It is invalid for timeoutSec to have
     * greater value than checkIntervalSec.
     */
    public readonly timeoutSec!: pulumi.Output<number | undefined>;
    /**
     * The type of the health check. One of HTTP, HTTPS, TCP, or SSL.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * A so-far healthy instance will be marked unhealthy after this many
     * consecutive failures. The default value is 2.
     */
    public readonly unhealthyThreshold!: pulumi.Output<number | undefined>;

    /**
     * Create a HealthCheck resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: HealthCheckArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HealthCheckArgs | HealthCheckState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as HealthCheckState | undefined;
            inputs["checkIntervalSec"] = state ? state.checkIntervalSec : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["healthyThreshold"] = state ? state.healthyThreshold : undefined;
            inputs["http2HealthCheck"] = state ? state.http2HealthCheck : undefined;
            inputs["httpHealthCheck"] = state ? state.httpHealthCheck : undefined;
            inputs["httpsHealthCheck"] = state ? state.httpsHealthCheck : undefined;
            inputs["logConfig"] = state ? state.logConfig : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["sslHealthCheck"] = state ? state.sslHealthCheck : undefined;
            inputs["tcpHealthCheck"] = state ? state.tcpHealthCheck : undefined;
            inputs["timeoutSec"] = state ? state.timeoutSec : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["unhealthyThreshold"] = state ? state.unhealthyThreshold : undefined;
        } else {
            const args = argsOrState as HealthCheckArgs | undefined;
            inputs["checkIntervalSec"] = args ? args.checkIntervalSec : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["healthyThreshold"] = args ? args.healthyThreshold : undefined;
            inputs["http2HealthCheck"] = args ? args.http2HealthCheck : undefined;
            inputs["httpHealthCheck"] = args ? args.httpHealthCheck : undefined;
            inputs["httpsHealthCheck"] = args ? args.httpsHealthCheck : undefined;
            inputs["logConfig"] = args ? args.logConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["sslHealthCheck"] = args ? args.sslHealthCheck : undefined;
            inputs["tcpHealthCheck"] = args ? args.tcpHealthCheck : undefined;
            inputs["timeoutSec"] = args ? args.timeoutSec : undefined;
            inputs["unhealthyThreshold"] = args ? args.unhealthyThreshold : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(HealthCheck.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HealthCheck resources.
 */
export interface HealthCheckState {
    /**
     * How often (in seconds) to send a health check. The default value is 5
     * seconds.
     */
    readonly checkIntervalSec?: pulumi.Input<number>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A so-far unhealthy instance will be marked healthy after this many
     * consecutive successes. The default value is 2.
     */
    readonly healthyThreshold?: pulumi.Input<number>;
    /**
     * A nested object resource  Structure is documented below.
     */
    readonly http2HealthCheck?: pulumi.Input<inputs.compute.HealthCheckHttp2HealthCheck>;
    /**
     * A nested object resource  Structure is documented below.
     */
    readonly httpHealthCheck?: pulumi.Input<inputs.compute.HealthCheckHttpHealthCheck>;
    /**
     * A nested object resource  Structure is documented below.
     */
    readonly httpsHealthCheck?: pulumi.Input<inputs.compute.HealthCheckHttpsHealthCheck>;
    /**
     * Configure logging on this health check.  Structure is documented below.
     */
    readonly logConfig?: pulumi.Input<inputs.compute.HealthCheckLogConfig>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * A nested object resource  Structure is documented below.
     */
    readonly sslHealthCheck?: pulumi.Input<inputs.compute.HealthCheckSslHealthCheck>;
    /**
     * A nested object resource  Structure is documented below.
     */
    readonly tcpHealthCheck?: pulumi.Input<inputs.compute.HealthCheckTcpHealthCheck>;
    /**
     * How long (in seconds) to wait before claiming failure.
     * The default value is 5 seconds.  It is invalid for timeoutSec to have
     * greater value than checkIntervalSec.
     */
    readonly timeoutSec?: pulumi.Input<number>;
    /**
     * The type of the health check. One of HTTP, HTTPS, TCP, or SSL.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * A so-far healthy instance will be marked unhealthy after this many
     * consecutive failures. The default value is 2.
     */
    readonly unhealthyThreshold?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a HealthCheck resource.
 */
export interface HealthCheckArgs {
    /**
     * How often (in seconds) to send a health check. The default value is 5
     * seconds.
     */
    readonly checkIntervalSec?: pulumi.Input<number>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A so-far unhealthy instance will be marked healthy after this many
     * consecutive successes. The default value is 2.
     */
    readonly healthyThreshold?: pulumi.Input<number>;
    /**
     * A nested object resource  Structure is documented below.
     */
    readonly http2HealthCheck?: pulumi.Input<inputs.compute.HealthCheckHttp2HealthCheck>;
    /**
     * A nested object resource  Structure is documented below.
     */
    readonly httpHealthCheck?: pulumi.Input<inputs.compute.HealthCheckHttpHealthCheck>;
    /**
     * A nested object resource  Structure is documented below.
     */
    readonly httpsHealthCheck?: pulumi.Input<inputs.compute.HealthCheckHttpsHealthCheck>;
    /**
     * Configure logging on this health check.  Structure is documented below.
     */
    readonly logConfig?: pulumi.Input<inputs.compute.HealthCheckLogConfig>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A nested object resource  Structure is documented below.
     */
    readonly sslHealthCheck?: pulumi.Input<inputs.compute.HealthCheckSslHealthCheck>;
    /**
     * A nested object resource  Structure is documented below.
     */
    readonly tcpHealthCheck?: pulumi.Input<inputs.compute.HealthCheckTcpHealthCheck>;
    /**
     * How long (in seconds) to wait before claiming failure.
     * The default value is 5 seconds.  It is invalid for timeoutSec to have
     * greater value than checkIntervalSec.
     */
    readonly timeoutSec?: pulumi.Input<number>;
    /**
     * A so-far healthy instance will be marked unhealthy after this many
     * consecutive failures. The default value is 2.
     */
    readonly unhealthyThreshold?: pulumi.Input<number>;
}
