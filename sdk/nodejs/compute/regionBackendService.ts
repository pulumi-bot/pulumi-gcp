// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../types/input";
import * as outputApi from "../types/output";
import * as utilities from "../utilities";

/**
 * A Region Backend Service defines a regionally-scoped group of virtual
 * machines that will serve traffic for load balancing.
 * 
 * Region backend services can only be used when using internal load balancing.
 * For external load balancing, use a global backend service instead.
 * 
 * 
 * To get more information about RegionBackendService, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/latest/regionBackendServices)
 * * How-to Guides
 *     * [Internal TCP/UDP Load Balancing](https://cloud.google.com/compute/docs/load-balancing/internal/)
 * 
 * ## Example Usage - Region Backend Service Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const defaultHealthCheck = new gcp.compute.HealthCheck("default", {
 *     checkIntervalSec: 1,
 *     tcpHealthCheck: {
 *         port: 80,
 *     },
 *     timeoutSec: 1,
 * });
 * const defaultRegionBackendService = new gcp.compute.RegionBackendService("default", {
 *     connectionDrainingTimeoutSec: 10,
 *     healthChecks: defaultHealthCheck.selfLink,
 *     region: "us-central1",
 *     sessionAffinity: "CLIENT_IP",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_region_backend_service.html.markdown.
 */
export class RegionBackendService extends pulumi.CustomResource {
    /**
     * Get an existing RegionBackendService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegionBackendServiceState, opts?: pulumi.CustomResourceOptions): RegionBackendService {
        return new RegionBackendService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/regionBackendService:RegionBackendService';

    /**
     * Returns true if the given object is an instance of RegionBackendService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionBackendService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionBackendService.__pulumiType;
    }

    public readonly backends!: pulumi.Output<outputApi.compute.RegionBackendServiceBackend[] | undefined>;
    public readonly connectionDrainingTimeoutSec!: pulumi.Output<number | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly failoverPolicy!: pulumi.Output<outputApi.compute.RegionBackendServiceFailoverPolicy | undefined>;
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    public readonly healthChecks!: pulumi.Output<string>;
    public readonly loadBalancingScheme!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public readonly protocol!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    public readonly sessionAffinity!: pulumi.Output<string>;
    public readonly timeoutSec!: pulumi.Output<number>;

    /**
     * Create a RegionBackendService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionBackendServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegionBackendServiceArgs | RegionBackendServiceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RegionBackendServiceState | undefined;
            inputs["backends"] = state ? state.backends : undefined;
            inputs["connectionDrainingTimeoutSec"] = state ? state.connectionDrainingTimeoutSec : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["failoverPolicy"] = state ? state.failoverPolicy : undefined;
            inputs["fingerprint"] = state ? state.fingerprint : undefined;
            inputs["healthChecks"] = state ? state.healthChecks : undefined;
            inputs["loadBalancingScheme"] = state ? state.loadBalancingScheme : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["sessionAffinity"] = state ? state.sessionAffinity : undefined;
            inputs["timeoutSec"] = state ? state.timeoutSec : undefined;
        } else {
            const args = argsOrState as RegionBackendServiceArgs | undefined;
            if (!args || args.healthChecks === undefined) {
                throw new Error("Missing required property 'healthChecks'");
            }
            inputs["backends"] = args ? args.backends : undefined;
            inputs["connectionDrainingTimeoutSec"] = args ? args.connectionDrainingTimeoutSec : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["failoverPolicy"] = args ? args.failoverPolicy : undefined;
            inputs["healthChecks"] = args ? args.healthChecks : undefined;
            inputs["loadBalancingScheme"] = args ? args.loadBalancingScheme : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["sessionAffinity"] = args ? args.sessionAffinity : undefined;
            inputs["timeoutSec"] = args ? args.timeoutSec : undefined;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RegionBackendService.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegionBackendService resources.
 */
export interface RegionBackendServiceState {
    readonly backends?: pulumi.Input<pulumi.Input<inputApi.compute.RegionBackendServiceBackend>[]>;
    readonly connectionDrainingTimeoutSec?: pulumi.Input<number>;
    readonly description?: pulumi.Input<string>;
    readonly failoverPolicy?: pulumi.Input<inputApi.compute.RegionBackendServiceFailoverPolicy>;
    readonly fingerprint?: pulumi.Input<string>;
    readonly healthChecks?: pulumi.Input<string>;
    readonly loadBalancingScheme?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly protocol?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly sessionAffinity?: pulumi.Input<string>;
    readonly timeoutSec?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RegionBackendService resource.
 */
export interface RegionBackendServiceArgs {
    readonly backends?: pulumi.Input<pulumi.Input<inputApi.compute.RegionBackendServiceBackend>[]>;
    readonly connectionDrainingTimeoutSec?: pulumi.Input<number>;
    readonly description?: pulumi.Input<string>;
    readonly failoverPolicy?: pulumi.Input<inputApi.compute.RegionBackendServiceFailoverPolicy>;
    readonly healthChecks: pulumi.Input<string>;
    readonly loadBalancingScheme?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly protocol?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly sessionAffinity?: pulumi.Input<string>;
    readonly timeoutSec?: pulumi.Input<number>;
}
