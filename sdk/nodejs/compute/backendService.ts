// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A Backend Service defines a group of virtual machines that will serve
 * traffic for load balancing. This resource is a global backend service,
 * appropriate for external load balancing. For internal load balancing, use
 * a regional backend service instead.
 * 
 * 
 * To get more information about BackendService, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/v1/backendServices)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/backend-service)
 * 
 * ## Example Usage - Backend Service Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const defaultHttpHealthCheck = new gcp.compute.HttpHealthCheck("default", {
 *     checkIntervalSec: 1,
 *     requestPath: "/",
 *     timeoutSec: 1,
 * });
 * const defaultBackendService = new gcp.compute.BackendService("default", {
 *     healthChecks: defaultHttpHealthCheck.selfLink,
 * });
 * ```
 */
export class BackendService extends pulumi.CustomResource {
    /**
     * Get an existing BackendService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackendServiceState, opts?: pulumi.CustomResourceOptions): BackendService {
        return new BackendService(name, <any>state, { ...opts, id: id });
    }

    public readonly affinityCookieTtlSec!: pulumi.Output<number | undefined>;
    public readonly backends!: pulumi.Output<{ balancingMode?: string, capacityScaler?: number, description?: string, group?: string, maxConnections?: number, maxConnectionsPerInstance?: number, maxRate?: number, maxRatePerInstance?: number, maxUtilization?: number }[] | undefined>;
    public readonly cdnPolicy!: pulumi.Output<{ cacheKeyPolicy?: { includeHost?: boolean, includeProtocol?: boolean, includeQueryString?: boolean, queryStringBlacklists?: string[], queryStringWhitelists?: string[] }, signedUrlCacheMaxAgeSec?: number }>;
    public readonly connectionDrainingTimeoutSec!: pulumi.Output<number | undefined>;
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    public readonly customRequestHeaders!: pulumi.Output<string[] | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly enableCdn!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    public readonly healthChecks!: pulumi.Output<string>;
    public readonly iap!: pulumi.Output<{ oauth2ClientId: string, oauth2ClientSecret: string, oauth2ClientSecretSha256: string } | undefined>;
    public readonly loadBalancingScheme!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly portName!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public readonly protocol!: pulumi.Output<string>;
    public readonly securityPolicy!: pulumi.Output<string | undefined>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    public readonly sessionAffinity!: pulumi.Output<string>;
    public readonly timeoutSec!: pulumi.Output<number>;

    /**
     * Create a BackendService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackendServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackendServiceArgs | BackendServiceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BackendServiceState | undefined;
            inputs["affinityCookieTtlSec"] = state ? state.affinityCookieTtlSec : undefined;
            inputs["backends"] = state ? state.backends : undefined;
            inputs["cdnPolicy"] = state ? state.cdnPolicy : undefined;
            inputs["connectionDrainingTimeoutSec"] = state ? state.connectionDrainingTimeoutSec : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["customRequestHeaders"] = state ? state.customRequestHeaders : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enableCdn"] = state ? state.enableCdn : undefined;
            inputs["fingerprint"] = state ? state.fingerprint : undefined;
            inputs["healthChecks"] = state ? state.healthChecks : undefined;
            inputs["iap"] = state ? state.iap : undefined;
            inputs["loadBalancingScheme"] = state ? state.loadBalancingScheme : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["portName"] = state ? state.portName : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["securityPolicy"] = state ? state.securityPolicy : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["sessionAffinity"] = state ? state.sessionAffinity : undefined;
            inputs["timeoutSec"] = state ? state.timeoutSec : undefined;
        } else {
            const args = argsOrState as BackendServiceArgs | undefined;
            if (!args || args.healthChecks === undefined) {
                throw new Error("Missing required property 'healthChecks'");
            }
            inputs["affinityCookieTtlSec"] = args ? args.affinityCookieTtlSec : undefined;
            inputs["backends"] = args ? args.backends : undefined;
            inputs["cdnPolicy"] = args ? args.cdnPolicy : undefined;
            inputs["connectionDrainingTimeoutSec"] = args ? args.connectionDrainingTimeoutSec : undefined;
            inputs["customRequestHeaders"] = args ? args.customRequestHeaders : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enableCdn"] = args ? args.enableCdn : undefined;
            inputs["healthChecks"] = args ? args.healthChecks : undefined;
            inputs["iap"] = args ? args.iap : undefined;
            inputs["loadBalancingScheme"] = args ? args.loadBalancingScheme : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["portName"] = args ? args.portName : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["securityPolicy"] = args ? args.securityPolicy : undefined;
            inputs["sessionAffinity"] = args ? args.sessionAffinity : undefined;
            inputs["timeoutSec"] = args ? args.timeoutSec : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        super("gcp:compute/backendService:BackendService", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackendService resources.
 */
export interface BackendServiceState {
    readonly affinityCookieTtlSec?: pulumi.Input<number>;
    readonly backends?: pulumi.Input<pulumi.Input<{ balancingMode?: pulumi.Input<string>, capacityScaler?: pulumi.Input<number>, description?: pulumi.Input<string>, group?: pulumi.Input<string>, maxConnections?: pulumi.Input<number>, maxConnectionsPerInstance?: pulumi.Input<number>, maxRate?: pulumi.Input<number>, maxRatePerInstance?: pulumi.Input<number>, maxUtilization?: pulumi.Input<number> }>[]>;
    readonly cdnPolicy?: pulumi.Input<{ cacheKeyPolicy?: pulumi.Input<{ includeHost?: pulumi.Input<boolean>, includeProtocol?: pulumi.Input<boolean>, includeQueryString?: pulumi.Input<boolean>, queryStringBlacklists?: pulumi.Input<pulumi.Input<string>[]>, queryStringWhitelists?: pulumi.Input<pulumi.Input<string>[]> }>, signedUrlCacheMaxAgeSec?: pulumi.Input<number> }>;
    readonly connectionDrainingTimeoutSec?: pulumi.Input<number>;
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly customRequestHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    readonly description?: pulumi.Input<string>;
    readonly enableCdn?: pulumi.Input<boolean>;
    readonly fingerprint?: pulumi.Input<string>;
    readonly healthChecks?: pulumi.Input<string>;
    readonly iap?: pulumi.Input<{ oauth2ClientId: pulumi.Input<string>, oauth2ClientSecret: pulumi.Input<string>, oauth2ClientSecretSha256?: pulumi.Input<string> }>;
    readonly loadBalancingScheme?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly portName?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly protocol?: pulumi.Input<string>;
    readonly securityPolicy?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly sessionAffinity?: pulumi.Input<string>;
    readonly timeoutSec?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a BackendService resource.
 */
export interface BackendServiceArgs {
    readonly affinityCookieTtlSec?: pulumi.Input<number>;
    readonly backends?: pulumi.Input<pulumi.Input<{ balancingMode?: pulumi.Input<string>, capacityScaler?: pulumi.Input<number>, description?: pulumi.Input<string>, group?: pulumi.Input<string>, maxConnections?: pulumi.Input<number>, maxConnectionsPerInstance?: pulumi.Input<number>, maxRate?: pulumi.Input<number>, maxRatePerInstance?: pulumi.Input<number>, maxUtilization?: pulumi.Input<number> }>[]>;
    readonly cdnPolicy?: pulumi.Input<{ cacheKeyPolicy?: pulumi.Input<{ includeHost?: pulumi.Input<boolean>, includeProtocol?: pulumi.Input<boolean>, includeQueryString?: pulumi.Input<boolean>, queryStringBlacklists?: pulumi.Input<pulumi.Input<string>[]>, queryStringWhitelists?: pulumi.Input<pulumi.Input<string>[]> }>, signedUrlCacheMaxAgeSec?: pulumi.Input<number> }>;
    readonly connectionDrainingTimeoutSec?: pulumi.Input<number>;
    readonly customRequestHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    readonly description?: pulumi.Input<string>;
    readonly enableCdn?: pulumi.Input<boolean>;
    readonly healthChecks: pulumi.Input<string>;
    readonly iap?: pulumi.Input<{ oauth2ClientId: pulumi.Input<string>, oauth2ClientSecret: pulumi.Input<string>, oauth2ClientSecretSha256?: pulumi.Input<string> }>;
    readonly loadBalancingScheme?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly portName?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly protocol?: pulumi.Input<string>;
    readonly securityPolicy?: pulumi.Input<string>;
    readonly sessionAffinity?: pulumi.Input<string>;
    readonly timeoutSec?: pulumi.Input<number>;
}
