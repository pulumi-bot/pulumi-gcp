// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

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

    public readonly autoHealingPolicies: pulumi.Output<{ healthCheck: string, initialDelaySec: number } | undefined>;
    public readonly baseInstanceName: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public /*out*/ readonly fingerprint: pulumi.Output<string>;
    public /*out*/ readonly instanceGroup: pulumi.Output<string>;
    public readonly instanceTemplate: pulumi.Output<string | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly namedPorts: pulumi.Output<{ name: string, port: number }[] | undefined>;
    public readonly project: pulumi.Output<string>;
    public readonly rollingUpdatePolicy: pulumi.Output<{ maxSurgeFixed?: number, maxSurgePercent?: number, maxUnavailableFixed?: number, maxUnavailablePercent?: number, minReadySec?: number, minimalAction: string, type: string } | undefined>;
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    public readonly targetPools: pulumi.Output<string[] | undefined>;
    public readonly targetSize: pulumi.Output<number>;
    public readonly updateStrategy: pulumi.Output<string | undefined>;
    public readonly versions: pulumi.Output<{ instanceTemplate: string, name: string, targetSize?: { fixed?: number, percent?: number } }[]>;
    public readonly waitForInstances: pulumi.Output<boolean | undefined>;
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
    readonly autoHealingPolicies?: pulumi.Input<{ healthCheck: pulumi.Input<string>, initialDelaySec: pulumi.Input<number> }>;
    readonly baseInstanceName?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly fingerprint?: pulumi.Input<string>;
    readonly instanceGroup?: pulumi.Input<string>;
    readonly instanceTemplate?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly namedPorts?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, port: pulumi.Input<number> }>[]>;
    readonly project?: pulumi.Input<string>;
    readonly rollingUpdatePolicy?: pulumi.Input<{ maxSurgeFixed?: pulumi.Input<number>, maxSurgePercent?: pulumi.Input<number>, maxUnavailableFixed?: pulumi.Input<number>, maxUnavailablePercent?: pulumi.Input<number>, minReadySec?: pulumi.Input<number>, minimalAction: pulumi.Input<string>, type: pulumi.Input<string> }>;
    readonly selfLink?: pulumi.Input<string>;
    readonly targetPools?: pulumi.Input<pulumi.Input<string>[]>;
    readonly targetSize?: pulumi.Input<number>;
    readonly updateStrategy?: pulumi.Input<string>;
    readonly versions?: pulumi.Input<pulumi.Input<{ instanceTemplate: pulumi.Input<string>, name: pulumi.Input<string>, targetSize?: pulumi.Input<{ fixed?: pulumi.Input<number>, percent?: pulumi.Input<number> }> }>[]>;
    readonly waitForInstances?: pulumi.Input<boolean>;
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceGroupManager resource.
 */
export interface InstanceGroupManagerArgs {
    readonly autoHealingPolicies?: pulumi.Input<{ healthCheck: pulumi.Input<string>, initialDelaySec: pulumi.Input<number> }>;
    readonly baseInstanceName: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly instanceTemplate?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly namedPorts?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, port: pulumi.Input<number> }>[]>;
    readonly project?: pulumi.Input<string>;
    readonly rollingUpdatePolicy?: pulumi.Input<{ maxSurgeFixed?: pulumi.Input<number>, maxSurgePercent?: pulumi.Input<number>, maxUnavailableFixed?: pulumi.Input<number>, maxUnavailablePercent?: pulumi.Input<number>, minReadySec?: pulumi.Input<number>, minimalAction: pulumi.Input<string>, type: pulumi.Input<string> }>;
    readonly targetPools?: pulumi.Input<pulumi.Input<string>[]>;
    readonly targetSize?: pulumi.Input<number>;
    readonly updateStrategy?: pulumi.Input<string>;
    readonly versions?: pulumi.Input<pulumi.Input<{ instanceTemplate: pulumi.Input<string>, name: pulumi.Input<string>, targetSize?: pulumi.Input<{ fixed?: pulumi.Input<number>, percent?: pulumi.Input<number> }> }>[]>;
    readonly waitForInstances?: pulumi.Input<boolean>;
    readonly zone?: pulumi.Input<string>;
}
