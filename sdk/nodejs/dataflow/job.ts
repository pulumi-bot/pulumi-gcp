// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: JobState, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, <any>state, { ...opts, id: id });
    }

    public readonly maxWorkers: pulumi.Output<number | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly onDelete: pulumi.Output<string | undefined>;
    public readonly parameters: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly project: pulumi.Output<string | undefined>;
    public readonly region: pulumi.Output<string | undefined>;
    public /*out*/ readonly state: pulumi.Output<string>;
    public readonly tempGcsLocation: pulumi.Output<string>;
    public readonly templateGcsPath: pulumi.Output<string>;
    public readonly zone: pulumi.Output<string | undefined>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JobArgs | JobState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: JobState = argsOrState as JobState | undefined;
            inputs["maxWorkers"] = state ? state.maxWorkers : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["onDelete"] = state ? state.onDelete : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["tempGcsLocation"] = state ? state.tempGcsLocation : undefined;
            inputs["templateGcsPath"] = state ? state.templateGcsPath : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as JobArgs | undefined;
            if (!args || args.tempGcsLocation === undefined) {
                throw new Error("Missing required property 'tempGcsLocation'");
            }
            if (!args || args.templateGcsPath === undefined) {
                throw new Error("Missing required property 'templateGcsPath'");
            }
            inputs["maxWorkers"] = args ? args.maxWorkers : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["onDelete"] = args ? args.onDelete : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["tempGcsLocation"] = args ? args.tempGcsLocation : undefined;
            inputs["templateGcsPath"] = args ? args.templateGcsPath : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["state"] = undefined /*out*/;
        }
        super("gcp:dataflow/job:Job", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Job resources.
 */
export interface JobState {
    readonly maxWorkers?: pulumi.Input<number>;
    readonly name?: pulumi.Input<string>;
    readonly onDelete?: pulumi.Input<string>;
    readonly parameters?: pulumi.Input<{[key: string]: any}>;
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly state?: pulumi.Input<string>;
    readonly tempGcsLocation?: pulumi.Input<string>;
    readonly templateGcsPath?: pulumi.Input<string>;
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    readonly maxWorkers?: pulumi.Input<number>;
    readonly name?: pulumi.Input<string>;
    readonly onDelete?: pulumi.Input<string>;
    readonly parameters?: pulumi.Input<{[key: string]: any}>;
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly tempGcsLocation: pulumi.Input<string>;
    readonly templateGcsPath: pulumi.Input<string>;
    readonly zone?: pulumi.Input<string>;
}
