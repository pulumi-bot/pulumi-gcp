// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The Google Compute Engine Instance Group Manager API creates and manages pools
 * of homogeneous Compute Engine virtual machine instances from a common instance
 * template. For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/manager)
 * and [API](https://cloud.google.com/compute/docs/reference/latest/instanceGroupManagers)
 * 
 * > **Note:** Use [google_compute_region_instance_group_manager](https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager.html) to create a regional (multi-zone) instance group manager.
 * 
 * ## Example Usage with top level instance template
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const google_compute_health_check_autohealing = new gcp.compute.HealthCheck("autohealing", {
 *     checkIntervalSec: 5,
 *     healthyThreshold: 2,
 *     httpHealthCheck: {
 *         port: Number.parseFloat("8080"),
 *         requestPath: "/healthz",
 *     },
 *     name: "autohealing-health-check",
 *     timeoutSec: 5,
 *     unhealthyThreshold: 10,
 * });
 * const google_compute_instance_group_manager_appserver = new gcp.compute.InstanceGroupManager("appserver", {
 *     autoHealingPolicies: {
 *         healthCheck: google_compute_health_check_autohealing.selfLink,
 *         initialDelaySec: 300,
 *     },
 *     baseInstanceName: "app",
 *     instanceTemplate: google_compute_instance_template_appserver.selfLink,
 *     name: "appserver-igm",
 *     namedPorts: [{
 *         name: "customHTTP",
 *         port: 8888,
 *     }],
 *     targetPools: [google_compute_target_pool_appserver.selfLink],
 *     targetSize: 2,
 *     updateStrategy: "NONE",
 *     zone: "us-central1-a",
 * });
 * ```
 * 
 * ## Example Usage with multiple Versions
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const google_compute_instance_group_manager_appserver = new gcp.compute.InstanceGroupManager("appserver", {
 *     baseInstanceName: "app",
 *     name: "appserver-igm",
 *     targetSize: 5,
 *     updateStrategy: "NONE",
 *     versions: [
 *         {
 *             instanceTemplate: google_compute_instance_template_appserver.selfLink,
 *         },
 *         {
 *             instanceTemplate: google_compute_instance_template_appserver_canary.selfLink,
 *             targetSize: {
 *                 fixed: 1,
 *             },
 *         },
 *     ],
 *     zone: "us-central1-a",
 * });
 * ```
 */
export class InstanceGroupManager extends pulumi.CustomResource {
    /**
     * Get an existing InstanceGroupManager resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceGroupManagerState, opts?: pulumi.CustomResourceOptions): InstanceGroupManager {
        return new InstanceGroupManager(name, <any>state, { ...opts, id: id });
    }

    /**
     * The autohealing policies for this managed instance
     * group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
     * This property is in beta, and should be used with the terraform-provider-google-beta provider.
     * See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
     */
    public readonly autoHealingPolicies: pulumi.Output<{ healthCheck: string, initialDelaySec: number } | undefined>;
    /**
     * The base instance name to use for
     * instances in this group. The value must be a valid
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
     * are lowercase letters, numbers, and hyphens (-). Instances are named by
     * appending a hyphen and a random four-character string to the base instance
     * name.
     */
    public readonly baseInstanceName: pulumi.Output<string>;
    /**
     * An optional textual description of the instance
     * group manager.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * The fingerprint of the instance group manager.
     */
    public /*out*/ readonly fingerprint: pulumi.Output<string>;
    /**
     * The full URL of the instance group created by the manager.
     */
    public /*out*/ readonly instanceGroup: pulumi.Output<string>;
    /**
     * - The full URL to an instance template from which all new instances of this version will be created.
     */
    public readonly instanceTemplate: pulumi.Output<string | undefined>;
    /**
     * - Version name.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The named port configuration. See the section below
     * for details on configuration.
     */
    public readonly namedPorts: pulumi.Output<{ name: string, port: number }[] | undefined>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/instanceGroupManagers/patch)
     * This property is in beta, and should be used with the terraform-provider-google-beta provider.
     * See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
     * - - -
     */
    public readonly rollingUpdatePolicy: pulumi.Output<{ maxSurgeFixed?: number, maxSurgePercent?: number, maxUnavailableFixed?: number, maxUnavailablePercent?: number, minReadySec?: number, minimalAction: string, type: string } | undefined>;
    /**
     * The URL of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    /**
     * The full URL of all target pools to which new
     * instances in the group are added. Updating the target pools attribute does
     * not affect existing instances.
     */
    public readonly targetPools: pulumi.Output<string[] | undefined>;
    /**
     * - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
     */
    public readonly targetSize: pulumi.Output<number>;
    /**
     * If the `instance_template`
     * resource is modified, a value of `"NONE"` will prevent any of the managed
     * instances from being restarted by Terraform. A value of `"REPLACE"` will
     * restart all of the instances at once. `"ROLLING_UPDATE"` is supported as a beta feature.
     * A value of `"ROLLING_UPDATE"` requires `rolling_update_policy` block to be set
     */
    public readonly updateStrategy: pulumi.Output<string | undefined>;
    /**
     * Application versions managed by this instance group. Each
     * version deals with a specific instance template, allowing canary release scenarios.
     * Conflicts with `instance_template`. Structure is documented below. Beware that
     * exactly one version must not specify a target size. It means that versions with
     * a target size will respect the setting, and the one without target size will
     * be applied to all remaining Instances (top level target_size - each version target_size).
     * This property is in beta, and should be used with the terraform-provider-google-beta provider.
     * See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
     */
    public readonly versions: pulumi.Output<{ instanceTemplate: string, name: string, targetSize?: { fixed?: number, percent?: number } }[]>;
    /**
     * Whether to wait for all instances to be created/updated before
     * returning. Note that if this is set to true and the operation does not succeed, Terraform will
     * continue trying until it times out.
     */
    public readonly waitForInstances: pulumi.Output<boolean | undefined>;
    /**
     * The zone that instances in this group should be created
     * in.
     */
    public readonly zone: pulumi.Output<string>;

    /**
     * Create a InstanceGroupManager resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceGroupManagerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceGroupManagerArgs | InstanceGroupManagerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: InstanceGroupManagerState = argsOrState as InstanceGroupManagerState | undefined;
            inputs["autoHealingPolicies"] = state ? state.autoHealingPolicies : undefined;
            inputs["baseInstanceName"] = state ? state.baseInstanceName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["fingerprint"] = state ? state.fingerprint : undefined;
            inputs["instanceGroup"] = state ? state.instanceGroup : undefined;
            inputs["instanceTemplate"] = state ? state.instanceTemplate : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namedPorts"] = state ? state.namedPorts : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["rollingUpdatePolicy"] = state ? state.rollingUpdatePolicy : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["targetPools"] = state ? state.targetPools : undefined;
            inputs["targetSize"] = state ? state.targetSize : undefined;
            inputs["updateStrategy"] = state ? state.updateStrategy : undefined;
            inputs["versions"] = state ? state.versions : undefined;
            inputs["waitForInstances"] = state ? state.waitForInstances : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceGroupManagerArgs | undefined;
            if (!args || args.baseInstanceName === undefined) {
                throw new Error("Missing required property 'baseInstanceName'");
            }
            inputs["autoHealingPolicies"] = args ? args.autoHealingPolicies : undefined;
            inputs["baseInstanceName"] = args ? args.baseInstanceName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["instanceTemplate"] = args ? args.instanceTemplate : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namedPorts"] = args ? args.namedPorts : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["rollingUpdatePolicy"] = args ? args.rollingUpdatePolicy : undefined;
            inputs["targetPools"] = args ? args.targetPools : undefined;
            inputs["targetSize"] = args ? args.targetSize : undefined;
            inputs["updateStrategy"] = args ? args.updateStrategy : undefined;
            inputs["versions"] = args ? args.versions : undefined;
            inputs["waitForInstances"] = args ? args.waitForInstances : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["instanceGroup"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        super("gcp:compute/instanceGroupManager:InstanceGroupManager", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceGroupManager resources.
 */
export interface InstanceGroupManagerState {
    /**
     * The autohealing policies for this managed instance
     * group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
     * This property is in beta, and should be used with the terraform-provider-google-beta provider.
     * See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
     */
    readonly autoHealingPolicies?: pulumi.Input<{ healthCheck: pulumi.Input<string>, initialDelaySec: pulumi.Input<number> }>;
    /**
     * The base instance name to use for
     * instances in this group. The value must be a valid
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
     * are lowercase letters, numbers, and hyphens (-). Instances are named by
     * appending a hyphen and a random four-character string to the base instance
     * name.
     */
    readonly baseInstanceName?: pulumi.Input<string>;
    /**
     * An optional textual description of the instance
     * group manager.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The fingerprint of the instance group manager.
     */
    readonly fingerprint?: pulumi.Input<string>;
    /**
     * The full URL of the instance group created by the manager.
     */
    readonly instanceGroup?: pulumi.Input<string>;
    /**
     * - The full URL to an instance template from which all new instances of this version will be created.
     */
    readonly instanceTemplate?: pulumi.Input<string>;
    /**
     * - Version name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The named port configuration. See the section below
     * for details on configuration.
     */
    readonly namedPorts?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, port: pulumi.Input<number> }>[]>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/instanceGroupManagers/patch)
     * This property is in beta, and should be used with the terraform-provider-google-beta provider.
     * See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
     * - - -
     */
    readonly rollingUpdatePolicy?: pulumi.Input<{ maxSurgeFixed?: pulumi.Input<number>, maxSurgePercent?: pulumi.Input<number>, maxUnavailableFixed?: pulumi.Input<number>, maxUnavailablePercent?: pulumi.Input<number>, minReadySec?: pulumi.Input<number>, minimalAction: pulumi.Input<string>, type: pulumi.Input<string> }>;
    /**
     * The URL of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * The full URL of all target pools to which new
     * instances in the group are added. Updating the target pools attribute does
     * not affect existing instances.
     */
    readonly targetPools?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
     */
    readonly targetSize?: pulumi.Input<number>;
    /**
     * If the `instance_template`
     * resource is modified, a value of `"NONE"` will prevent any of the managed
     * instances from being restarted by Terraform. A value of `"REPLACE"` will
     * restart all of the instances at once. `"ROLLING_UPDATE"` is supported as a beta feature.
     * A value of `"ROLLING_UPDATE"` requires `rolling_update_policy` block to be set
     */
    readonly updateStrategy?: pulumi.Input<string>;
    /**
     * Application versions managed by this instance group. Each
     * version deals with a specific instance template, allowing canary release scenarios.
     * Conflicts with `instance_template`. Structure is documented below. Beware that
     * exactly one version must not specify a target size. It means that versions with
     * a target size will respect the setting, and the one without target size will
     * be applied to all remaining Instances (top level target_size - each version target_size).
     * This property is in beta, and should be used with the terraform-provider-google-beta provider.
     * See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
     */
    readonly versions?: pulumi.Input<pulumi.Input<{ instanceTemplate: pulumi.Input<string>, name: pulumi.Input<string>, targetSize?: pulumi.Input<{ fixed?: pulumi.Input<number>, percent?: pulumi.Input<number> }> }>[]>;
    /**
     * Whether to wait for all instances to be created/updated before
     * returning. Note that if this is set to true and the operation does not succeed, Terraform will
     * continue trying until it times out.
     */
    readonly waitForInstances?: pulumi.Input<boolean>;
    /**
     * The zone that instances in this group should be created
     * in.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceGroupManager resource.
 */
export interface InstanceGroupManagerArgs {
    /**
     * The autohealing policies for this managed instance
     * group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
     * This property is in beta, and should be used with the terraform-provider-google-beta provider.
     * See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
     */
    readonly autoHealingPolicies?: pulumi.Input<{ healthCheck: pulumi.Input<string>, initialDelaySec: pulumi.Input<number> }>;
    /**
     * The base instance name to use for
     * instances in this group. The value must be a valid
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
     * are lowercase letters, numbers, and hyphens (-). Instances are named by
     * appending a hyphen and a random four-character string to the base instance
     * name.
     */
    readonly baseInstanceName: pulumi.Input<string>;
    /**
     * An optional textual description of the instance
     * group manager.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * - The full URL to an instance template from which all new instances of this version will be created.
     */
    readonly instanceTemplate?: pulumi.Input<string>;
    /**
     * - Version name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The named port configuration. See the section below
     * for details on configuration.
     */
    readonly namedPorts?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, port: pulumi.Input<number> }>[]>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/instanceGroupManagers/patch)
     * This property is in beta, and should be used with the terraform-provider-google-beta provider.
     * See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
     * - - -
     */
    readonly rollingUpdatePolicy?: pulumi.Input<{ maxSurgeFixed?: pulumi.Input<number>, maxSurgePercent?: pulumi.Input<number>, maxUnavailableFixed?: pulumi.Input<number>, maxUnavailablePercent?: pulumi.Input<number>, minReadySec?: pulumi.Input<number>, minimalAction: pulumi.Input<string>, type: pulumi.Input<string> }>;
    /**
     * The full URL of all target pools to which new
     * instances in the group are added. Updating the target pools attribute does
     * not affect existing instances.
     */
    readonly targetPools?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
     */
    readonly targetSize?: pulumi.Input<number>;
    /**
     * If the `instance_template`
     * resource is modified, a value of `"NONE"` will prevent any of the managed
     * instances from being restarted by Terraform. A value of `"REPLACE"` will
     * restart all of the instances at once. `"ROLLING_UPDATE"` is supported as a beta feature.
     * A value of `"ROLLING_UPDATE"` requires `rolling_update_policy` block to be set
     */
    readonly updateStrategy?: pulumi.Input<string>;
    /**
     * Application versions managed by this instance group. Each
     * version deals with a specific instance template, allowing canary release scenarios.
     * Conflicts with `instance_template`. Structure is documented below. Beware that
     * exactly one version must not specify a target size. It means that versions with
     * a target size will respect the setting, and the one without target size will
     * be applied to all remaining Instances (top level target_size - each version target_size).
     * This property is in beta, and should be used with the terraform-provider-google-beta provider.
     * See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
     */
    readonly versions?: pulumi.Input<pulumi.Input<{ instanceTemplate: pulumi.Input<string>, name: pulumi.Input<string>, targetSize?: pulumi.Input<{ fixed?: pulumi.Input<number>, percent?: pulumi.Input<number> }> }>[]>;
    /**
     * Whether to wait for all instances to be created/updated before
     * returning. Note that if this is set to true and the operation does not succeed, Terraform will
     * continue trying until it times out.
     */
    readonly waitForInstances?: pulumi.Input<boolean>;
    /**
     * The zone that instances in this group should be created
     * in.
     */
    readonly zone?: pulumi.Input<string>;
}
