// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Represents an Autoscaler resource.
 *
 * Autoscalers allow you to automatically scale virtual machine instances in
 * managed instance groups according to an autoscaling policy that you
 * define.
 *
 *
 * To get more information about RegionAutoscaler, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/regionAutoscalers)
 * * How-to Guides
 *     * [Autoscaling Groups of Instances](https://cloud.google.com/compute/docs/autoscaler/)
 *
 * ## Example Usage - Region Autoscaler Basic
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const debian9 = gcp.compute.getImage({
 *     family: "debian-9",
 *     project: "debian-cloud",
 * });
 * const foobarInstanceTemplate = new gcp.compute.InstanceTemplate("foobarInstanceTemplate", {
 *     machineType: "n1-standard-1",
 *     canIpForward: false,
 *     tags: [
 *         "foo",
 *         "bar",
 *     ],
 *     disk: [{
 *         sourceImage: debian9.then(debian9 => debian9.selfLink),
 *     }],
 *     network_interface: [{
 *         network: "default",
 *     }],
 *     metadata: {
 *         foo: "bar",
 *     },
 *     service_account: {
 *         scopes: [
 *             "userinfo-email",
 *             "compute-ro",
 *             "storage-ro",
 *         ],
 *     },
 * });
 * const foobarTargetPool = new gcp.compute.TargetPool("foobarTargetPool", {});
 * const foobarRegionInstanceGroupManager = new gcp.compute.RegionInstanceGroupManager("foobarRegionInstanceGroupManager", {
 *     region: "us-central1",
 *     version: [{
 *         instanceTemplate: foobarInstanceTemplate.id,
 *         name: "primary",
 *     }],
 *     targetPools: [foobarTargetPool.id],
 *     baseInstanceName: "foobar",
 * });
 * const foobarRegionAutoscaler = new gcp.compute.RegionAutoscaler("foobarRegionAutoscaler", {
 *     region: "us-central1",
 *     target: foobarRegionInstanceGroupManager.id,
 *     autoscaling_policy: {
 *         maxReplicas: 5,
 *         minReplicas: 1,
 *         cooldownPeriod: 60,
 *         cpu_utilization: {
 *             target: 0.5,
 *         },
 *     },
 * });
 * ```
 */
export class RegionAutoscaler extends pulumi.CustomResource {
    /**
     * Get an existing RegionAutoscaler resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegionAutoscalerState, opts?: pulumi.CustomResourceOptions): RegionAutoscaler {
        return new RegionAutoscaler(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/regionAutoscaler:RegionAutoscaler';

    /**
     * Returns true if the given object is an instance of RegionAutoscaler.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionAutoscaler {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionAutoscaler.__pulumiType;
    }

    /**
     * The configuration parameters for the autoscaling algorithm. You can
     * define one or more of the policies for an autoscaler: cpuUtilization,
     * customMetricUtilizations, and loadBalancingUtilization.
     * If none of these are specified, the default will be to autoscale based
     * on cpuUtilization to 0.6 or 60%.  Structure is documented below.
     */
    public readonly autoscalingPolicy!: pulumi.Output<outputs.compute.RegionAutoscalerAutoscalingPolicy>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The identifier (type) of the Stackdriver Monitoring metric.
     * The metric cannot have negative values.
     * The metric must have a value type of INT64 or DOUBLE.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * URL of the region where the instance group resides.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Fraction of backend capacity utilization (set in HTTP(s) load
     * balancing configuration) that autoscaler should maintain. Must
     * be a positive float value. If not defined, the default is 0.8.
     */
    public readonly target!: pulumi.Output<string>;

    /**
     * Create a RegionAutoscaler resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionAutoscalerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegionAutoscalerArgs | RegionAutoscalerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RegionAutoscalerState | undefined;
            inputs["autoscalingPolicy"] = state ? state.autoscalingPolicy : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["target"] = state ? state.target : undefined;
        } else {
            const args = argsOrState as RegionAutoscalerArgs | undefined;
            if (!args || args.autoscalingPolicy === undefined) {
                throw new Error("Missing required property 'autoscalingPolicy'");
            }
            if (!args || args.target === undefined) {
                throw new Error("Missing required property 'target'");
            }
            inputs["autoscalingPolicy"] = args ? args.autoscalingPolicy : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["target"] = args ? args.target : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RegionAutoscaler.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegionAutoscaler resources.
 */
export interface RegionAutoscalerState {
    /**
     * The configuration parameters for the autoscaling algorithm. You can
     * define one or more of the policies for an autoscaler: cpuUtilization,
     * customMetricUtilizations, and loadBalancingUtilization.
     * If none of these are specified, the default will be to autoscale based
     * on cpuUtilization to 0.6 or 60%.  Structure is documented below.
     */
    readonly autoscalingPolicy?: pulumi.Input<inputs.compute.RegionAutoscalerAutoscalingPolicy>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The identifier (type) of the Stackdriver Monitoring metric.
     * The metric cannot have negative values.
     * The metric must have a value type of INT64 or DOUBLE.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * URL of the region where the instance group resides.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * Fraction of backend capacity utilization (set in HTTP(s) load
     * balancing configuration) that autoscaler should maintain. Must
     * be a positive float value. If not defined, the default is 0.8.
     */
    readonly target?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RegionAutoscaler resource.
 */
export interface RegionAutoscalerArgs {
    /**
     * The configuration parameters for the autoscaling algorithm. You can
     * define one or more of the policies for an autoscaler: cpuUtilization,
     * customMetricUtilizations, and loadBalancingUtilization.
     * If none of these are specified, the default will be to autoscale based
     * on cpuUtilization to 0.6 or 60%.  Structure is documented below.
     */
    readonly autoscalingPolicy: pulumi.Input<inputs.compute.RegionAutoscalerAutoscalingPolicy>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The identifier (type) of the Stackdriver Monitoring metric.
     * The metric cannot have negative values.
     * The metric must have a value type of INT64 or DOUBLE.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * URL of the region where the instance group resides.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Fraction of backend capacity utilization (set in HTTP(s) load
     * balancing configuration) that autoscaler should maintain. Must
     * be a positive float value. If not defined, the default is 0.8.
     */
    readonly target: pulumi.Input<string>;
}
