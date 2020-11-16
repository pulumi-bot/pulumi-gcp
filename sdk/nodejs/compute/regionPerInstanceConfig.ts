// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A config defined for a single managed instance that belongs to an instance group manager. It preserves the instance name
 * across instance group manager operations and can define stateful disks or metadata that are unique to the instance.
 * This resource works with regional instance group managers.
 *
 * To get more information about RegionPerInstanceConfig, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/compute/docs/instance-groups/stateful-migs#per-instance_configs)
 *
 * ## Example Usage
 */
export class RegionPerInstanceConfig extends pulumi.CustomResource {
    /**
     * Get an existing RegionPerInstanceConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegionPerInstanceConfigState, opts?: pulumi.CustomResourceOptions): RegionPerInstanceConfig {
        return new RegionPerInstanceConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/regionPerInstanceConfig:RegionPerInstanceConfig';

    /**
     * Returns true if the given object is an instance of RegionPerInstanceConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionPerInstanceConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionPerInstanceConfig.__pulumiType;
    }

    /**
     * The minimal action to perform on the instance during an update.
     * Default is `NONE`. Possible values are:
     * * REPLACE
     * * RESTART
     * * REFRESH
     * * NONE
     */
    public readonly minimalAction!: pulumi.Output<string | undefined>;
    /**
     * The most disruptive action to perform on the instance during an update.
     * Default is `REPLACE`. Possible values are:
     * * REPLACE
     * * RESTART
     * * REFRESH
     * * NONE
     */
    public readonly mostDisruptiveAllowedAction!: pulumi.Output<string | undefined>;
    /**
     * The name for this per-instance config and its corresponding instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The preserved state for this instance.
     * Structure is documented below.
     */
    public readonly preservedState!: pulumi.Output<outputs.compute.RegionPerInstanceConfigPreservedState | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Region where the containing instance group manager is located
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The region instance group manager this instance config is part of.
     */
    public readonly regionInstanceGroupManager!: pulumi.Output<string>;
    /**
     * When true, deleting this config will immediately remove any specified state from the underlying instance.
     * When false, deleting this config will *not* immediately remove any state from the underlying instance.
     * State will be removed on the next instance recreation or update.
     */
    public readonly removeInstanceStateOnDestroy!: pulumi.Output<boolean | undefined>;

    /**
     * Create a RegionPerInstanceConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionPerInstanceConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegionPerInstanceConfigArgs | RegionPerInstanceConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RegionPerInstanceConfigState | undefined;
            inputs["minimalAction"] = state ? state.minimalAction : undefined;
            inputs["mostDisruptiveAllowedAction"] = state ? state.mostDisruptiveAllowedAction : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["preservedState"] = state ? state.preservedState : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["regionInstanceGroupManager"] = state ? state.regionInstanceGroupManager : undefined;
            inputs["removeInstanceStateOnDestroy"] = state ? state.removeInstanceStateOnDestroy : undefined;
        } else {
            const args = argsOrState as RegionPerInstanceConfigArgs | undefined;
            if (!args || args.region === undefined) {
                throw new Error("Missing required property 'region'");
            }
            if (!args || args.regionInstanceGroupManager === undefined) {
                throw new Error("Missing required property 'regionInstanceGroupManager'");
            }
            inputs["minimalAction"] = args ? args.minimalAction : undefined;
            inputs["mostDisruptiveAllowedAction"] = args ? args.mostDisruptiveAllowedAction : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["preservedState"] = args ? args.preservedState : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["regionInstanceGroupManager"] = args ? args.regionInstanceGroupManager : undefined;
            inputs["removeInstanceStateOnDestroy"] = args ? args.removeInstanceStateOnDestroy : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RegionPerInstanceConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegionPerInstanceConfig resources.
 */
export interface RegionPerInstanceConfigState {
    /**
     * The minimal action to perform on the instance during an update.
     * Default is `NONE`. Possible values are:
     * * REPLACE
     * * RESTART
     * * REFRESH
     * * NONE
     */
    readonly minimalAction?: pulumi.Input<string>;
    /**
     * The most disruptive action to perform on the instance during an update.
     * Default is `REPLACE`. Possible values are:
     * * REPLACE
     * * RESTART
     * * REFRESH
     * * NONE
     */
    readonly mostDisruptiveAllowedAction?: pulumi.Input<string>;
    /**
     * The name for this per-instance config and its corresponding instance.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The preserved state for this instance.
     * Structure is documented below.
     */
    readonly preservedState?: pulumi.Input<inputs.compute.RegionPerInstanceConfigPreservedState>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Region where the containing instance group manager is located
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The region instance group manager this instance config is part of.
     */
    readonly regionInstanceGroupManager?: pulumi.Input<string>;
    /**
     * When true, deleting this config will immediately remove any specified state from the underlying instance.
     * When false, deleting this config will *not* immediately remove any state from the underlying instance.
     * State will be removed on the next instance recreation or update.
     */
    readonly removeInstanceStateOnDestroy?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a RegionPerInstanceConfig resource.
 */
export interface RegionPerInstanceConfigArgs {
    /**
     * The minimal action to perform on the instance during an update.
     * Default is `NONE`. Possible values are:
     * * REPLACE
     * * RESTART
     * * REFRESH
     * * NONE
     */
    readonly minimalAction?: pulumi.Input<string>;
    /**
     * The most disruptive action to perform on the instance during an update.
     * Default is `REPLACE`. Possible values are:
     * * REPLACE
     * * RESTART
     * * REFRESH
     * * NONE
     */
    readonly mostDisruptiveAllowedAction?: pulumi.Input<string>;
    /**
     * The name for this per-instance config and its corresponding instance.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The preserved state for this instance.
     * Structure is documented below.
     */
    readonly preservedState?: pulumi.Input<inputs.compute.RegionPerInstanceConfigPreservedState>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Region where the containing instance group manager is located
     */
    readonly region: pulumi.Input<string>;
    /**
     * The region instance group manager this instance config is part of.
     */
    readonly regionInstanceGroupManager: pulumi.Input<string>;
    /**
     * When true, deleting this config will immediately remove any specified state from the underlying instance.
     * When false, deleting this config will *not* immediately remove any state from the underlying instance.
     * State will be removed on the next instance recreation or update.
     */
    readonly removeInstanceStateOnDestroy?: pulumi.Input<boolean>;
}
