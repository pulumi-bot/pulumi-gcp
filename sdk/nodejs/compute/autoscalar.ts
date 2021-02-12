// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Represents an Autoscaler resource.
 *
 * Autoscalers allow you to automatically scale virtual machine instances in
 * managed instance groups according to an autoscaling policy that you
 * define.
 *
 * To get more information about Autoscaler, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/autoscalers)
 * * How-to Guides
 *     * [Autoscaling Groups of Instances](https://cloud.google.com/compute/docs/autoscaler/)
 *
 * ## Example Usage
 * ### Autoscaler Single Instance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const debian9 = gcp.compute.getImage({
 *     family: "debian-9",
 *     project: "debian-cloud",
 * });
 * const defaultInstanceTemplate = new gcp.compute.InstanceTemplate("defaultInstanceTemplate", {
 *     machineType: "e2-medium",
 *     canIpForward: false,
 *     tags: [
 *         "foo",
 *         "bar",
 *     ],
 *     disks: [{
 *         sourceImage: debian9.then(debian9 => debian9.id),
 *     }],
 *     networkInterfaces: [{
 *         network: "default",
 *     }],
 *     metadata: {
 *         foo: "bar",
 *     },
 *     serviceAccount: {
 *         scopes: [
 *             "userinfo-email",
 *             "compute-ro",
 *             "storage-ro",
 *         ],
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * const defaultTargetPool = new gcp.compute.TargetPool("defaultTargetPool", {}, {
 *     provider: google_beta,
 * });
 * const defaultInstanceGroupManager = new gcp.compute.InstanceGroupManager("defaultInstanceGroupManager", {
 *     zone: "us-central1-f",
 *     versions: [{
 *         instanceTemplate: defaultInstanceTemplate.id,
 *         name: "primary",
 *     }],
 *     targetPools: [defaultTargetPool.id],
 *     baseInstanceName: "autoscaler-sample",
 * }, {
 *     provider: google_beta,
 * });
 * const defaultAutoscaler = new gcp.compute.Autoscaler("defaultAutoscaler", {
 *     zone: "us-central1-f",
 *     target: defaultInstanceGroupManager.id,
 *     autoscalingPolicy: {
 *         maxReplicas: 5,
 *         minReplicas: 1,
 *         cooldownPeriod: 60,
 *         metrics: [{
 *             name: "pubsub.googleapis.com/subscription/num_undelivered_messages",
 *             filter: "resource.type = pubsub_subscription AND resource.label.subscription_id = our-subscription",
 *             singleInstanceAssignment: 65535,
 *         }],
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Autoscaler Basic
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
 *     machineType: "e2-medium",
 *     canIpForward: false,
 *     tags: [
 *         "foo",
 *         "bar",
 *     ],
 *     disks: [{
 *         sourceImage: debian9.then(debian9 => debian9.id),
 *     }],
 *     networkInterfaces: [{
 *         network: "default",
 *     }],
 *     metadata: {
 *         foo: "bar",
 *     },
 *     serviceAccount: {
 *         scopes: [
 *             "userinfo-email",
 *             "compute-ro",
 *             "storage-ro",
 *         ],
 *     },
 * });
 * const foobarTargetPool = new gcp.compute.TargetPool("foobarTargetPool", {});
 * const foobarInstanceGroupManager = new gcp.compute.InstanceGroupManager("foobarInstanceGroupManager", {
 *     zone: "us-central1-f",
 *     versions: [{
 *         instanceTemplate: foobarInstanceTemplate.id,
 *         name: "primary",
 *     }],
 *     targetPools: [foobarTargetPool.id],
 *     baseInstanceName: "foobar",
 * });
 * const foobarAutoscaler = new gcp.compute.Autoscaler("foobarAutoscaler", {
 *     zone: "us-central1-f",
 *     target: foobarInstanceGroupManager.id,
 *     autoscalingPolicy: {
 *         maxReplicas: 5,
 *         minReplicas: 1,
 *         cooldownPeriod: 60,
 *         cpuUtilization: {
 *             target: 0.5,
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Autoscaler can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/autoscalar:Autoscalar default projects/{{project}}/zones/{{zone}}/autoscalers/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/autoscalar:Autoscalar default {{project}}/{{zone}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/autoscalar:Autoscalar default {{zone}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/autoscalar:Autoscalar default {{name}}
 * ```
 *
 * @deprecated gcp.compute.Autoscalar has been deprecated in favor of gcp.compute.Autoscaler
 */
export class Autoscalar extends pulumi.CustomResource {
    /**
     * Get an existing Autoscalar resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutoscalarState, opts?: pulumi.CustomResourceOptions): Autoscalar {
        pulumi.log.warn("Autoscalar is deprecated: gcp.compute.Autoscalar has been deprecated in favor of gcp.compute.Autoscaler")
        return new Autoscalar(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/autoscalar:Autoscalar';

    /**
     * Returns true if the given object is an instance of Autoscalar.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Autoscalar {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Autoscalar.__pulumiType;
    }

    /**
     * The configuration parameters for the autoscaling algorithm. You can
     * define one or more of the policies for an autoscaler: cpuUtilization,
     * customMetricUtilizations, and loadBalancingUtilization.
     * If none of these are specified, the default will be to autoscale based
     * on cpuUtilization to 0.6 or 60%.
     * Structure is documented below.
     */
    public readonly autoscalingPolicy!: pulumi.Output<outputs.compute.AutoscalarAutoscalingPolicy>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The identifier for this object. Format specified above.
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
     * Fraction of backend capacity utilization (set in HTTP(s) load
     * balancing configuration) that autoscaler should maintain. Must
     * be a positive float value. If not defined, the default is 0.8.
     */
    public readonly target!: pulumi.Output<string>;
    /**
     * URL of the zone where the instance group resides.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Autoscalar resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated gcp.compute.Autoscalar has been deprecated in favor of gcp.compute.Autoscaler */
    constructor(name: string, args: AutoscalarArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated gcp.compute.Autoscalar has been deprecated in favor of gcp.compute.Autoscaler */
    constructor(name: string, argsOrState?: AutoscalarArgs | AutoscalarState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Autoscalar is deprecated: gcp.compute.Autoscalar has been deprecated in favor of gcp.compute.Autoscaler")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AutoscalarState | undefined;
            inputs["autoscalingPolicy"] = state ? state.autoscalingPolicy : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["target"] = state ? state.target : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as AutoscalarArgs | undefined;
            if ((!args || args.autoscalingPolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoscalingPolicy'");
            }
            if ((!args || args.target === undefined) && !opts.urn) {
                throw new Error("Missing required property 'target'");
            }
            inputs["autoscalingPolicy"] = args ? args.autoscalingPolicy : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["target"] = args ? args.target : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Autoscalar.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Autoscalar resources.
 */
export interface AutoscalarState {
    /**
     * The configuration parameters for the autoscaling algorithm. You can
     * define one or more of the policies for an autoscaler: cpuUtilization,
     * customMetricUtilizations, and loadBalancingUtilization.
     * If none of these are specified, the default will be to autoscale based
     * on cpuUtilization to 0.6 or 60%.
     * Structure is documented below.
     */
    readonly autoscalingPolicy?: pulumi.Input<inputs.compute.AutoscalarAutoscalingPolicy>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The identifier for this object. Format specified above.
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
     * Fraction of backend capacity utilization (set in HTTP(s) load
     * balancing configuration) that autoscaler should maintain. Must
     * be a positive float value. If not defined, the default is 0.8.
     */
    readonly target?: pulumi.Input<string>;
    /**
     * URL of the zone where the instance group resides.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Autoscalar resource.
 */
export interface AutoscalarArgs {
    /**
     * The configuration parameters for the autoscaling algorithm. You can
     * define one or more of the policies for an autoscaler: cpuUtilization,
     * customMetricUtilizations, and loadBalancingUtilization.
     * If none of these are specified, the default will be to autoscale based
     * on cpuUtilization to 0.6 or 60%.
     * Structure is documented below.
     */
    readonly autoscalingPolicy: pulumi.Input<inputs.compute.AutoscalarAutoscalingPolicy>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The identifier for this object. Format specified above.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Fraction of backend capacity utilization (set in HTTP(s) load
     * balancing configuration) that autoscaler should maintain. Must
     * be a positive float value. If not defined, the default is 0.8.
     */
    readonly target: pulumi.Input<string>;
    /**
     * URL of the zone where the instance group resides.
     */
    readonly zone?: pulumi.Input<string>;
}
