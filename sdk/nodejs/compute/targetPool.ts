// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Target Pool within GCE. This is a collection of instances used as
 * target of a network load balancer (Forwarding Rule). For more information see
 * [the official
 * documentation](https://cloud.google.com/compute/docs/load-balancing/network/target-pools)
 * and [API](https://cloud.google.com/compute/docs/reference/latest/targetPools).
 * 
 * 
 * {{% examples %}}
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_target_pool.html.markdown.
 */
export class TargetPool extends pulumi.CustomResource {
    /**
     * Get an existing TargetPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TargetPoolState, opts?: pulumi.CustomResourceOptions): TargetPool {
        return new TargetPool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/targetPool:TargetPool';

    /**
     * Returns true if the given object is an instance of TargetPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TargetPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TargetPool.__pulumiType;
    }

    /**
     * URL to the backup target pool. Must also set
     * failover\_ratio.
     */
    public readonly backupPool!: pulumi.Output<string | undefined>;
    /**
     * Textual description field.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Ratio (0 to 1) of failed nodes before using the
     * backup pool (which must also be set).
     */
    public readonly failoverRatio!: pulumi.Output<number | undefined>;
    /**
     * List of zero or one health check name or self_link. Only
     * legacy `gcp.compute.HttpHealthCheck` is supported.
     */
    public readonly healthChecks!: pulumi.Output<string | undefined>;
    /**
     * List of instances in the pool. They can be given as
     * URLs, or in the form of "zone/name". Note that the instances need not exist
     * at the time of target pool creation, so there is no need to use
     * interpolators to create a dependency on the instances from the
     * target pool.
     */
    public readonly instances!: pulumi.Output<string[]>;
    /**
     * A unique name for the resource, required by GCE. Changing
     * this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Where the target pool resides. Defaults to project
     * region.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * How to distribute load. Options are "NONE" (no
     * affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and
     * "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE").
     */
    public readonly sessionAffinity!: pulumi.Output<string | undefined>;

    /**
     * Create a TargetPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TargetPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TargetPoolArgs | TargetPoolState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TargetPoolState | undefined;
            inputs["backupPool"] = state ? state.backupPool : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["failoverRatio"] = state ? state.failoverRatio : undefined;
            inputs["healthChecks"] = state ? state.healthChecks : undefined;
            inputs["instances"] = state ? state.instances : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["sessionAffinity"] = state ? state.sessionAffinity : undefined;
        } else {
            const args = argsOrState as TargetPoolArgs | undefined;
            inputs["backupPool"] = args ? args.backupPool : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["failoverRatio"] = args ? args.failoverRatio : undefined;
            inputs["healthChecks"] = args ? args.healthChecks : undefined;
            inputs["instances"] = args ? args.instances : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["sessionAffinity"] = args ? args.sessionAffinity : undefined;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TargetPool.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TargetPool resources.
 */
export interface TargetPoolState {
    /**
     * URL to the backup target pool. Must also set
     * failover\_ratio.
     */
    readonly backupPool?: pulumi.Input<string>;
    /**
     * Textual description field.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Ratio (0 to 1) of failed nodes before using the
     * backup pool (which must also be set).
     */
    readonly failoverRatio?: pulumi.Input<number>;
    /**
     * List of zero or one health check name or self_link. Only
     * legacy `gcp.compute.HttpHealthCheck` is supported.
     */
    readonly healthChecks?: pulumi.Input<string>;
    /**
     * List of instances in the pool. They can be given as
     * URLs, or in the form of "zone/name". Note that the instances need not exist
     * at the time of target pool creation, so there is no need to use
     * interpolators to create a dependency on the instances from the
     * target pool.
     */
    readonly instances?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A unique name for the resource, required by GCE. Changing
     * this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Where the target pool resides. Defaults to project
     * region.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * How to distribute load. Options are "NONE" (no
     * affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and
     * "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE").
     */
    readonly sessionAffinity?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TargetPool resource.
 */
export interface TargetPoolArgs {
    /**
     * URL to the backup target pool. Must also set
     * failover\_ratio.
     */
    readonly backupPool?: pulumi.Input<string>;
    /**
     * Textual description field.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Ratio (0 to 1) of failed nodes before using the
     * backup pool (which must also be set).
     */
    readonly failoverRatio?: pulumi.Input<number>;
    /**
     * List of zero or one health check name or self_link. Only
     * legacy `gcp.compute.HttpHealthCheck` is supported.
     */
    readonly healthChecks?: pulumi.Input<string>;
    /**
     * List of instances in the pool. They can be given as
     * URLs, or in the form of "zone/name". Note that the instances need not exist
     * at the time of target pool creation, so there is no need to use
     * interpolators to create a dependency on the instances from the
     * target pool.
     */
    readonly instances?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A unique name for the resource, required by GCE. Changing
     * this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Where the target pool resides. Defaults to project
     * region.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * How to distribute load. Options are "NONE" (no
     * affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and
     * "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE").
     */
    readonly sessionAffinity?: pulumi.Input<string>;
}
