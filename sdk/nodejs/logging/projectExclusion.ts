// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class ProjectExclusion extends pulumi.CustomResource {
    /**
     * Get an existing ProjectExclusion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectExclusionState, opts?: pulumi.CustomResourceOptions): ProjectExclusion {
        return new ProjectExclusion(name, <any>state, { ...opts, id: id });
    }

    public readonly description: pulumi.Output<string | undefined>;
    public readonly disabled: pulumi.Output<boolean | undefined>;
    public readonly filter: pulumi.Output<string>;
    public readonly name: pulumi.Output<string>;
    public readonly project: pulumi.Output<string>;

    /**
     * Create a ProjectExclusion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectExclusionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectExclusionArgs | ProjectExclusionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ProjectExclusionState = argsOrState as ProjectExclusionState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["disabled"] = state ? state.disabled : undefined;
            inputs["filter"] = state ? state.filter : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as ProjectExclusionArgs | undefined;
            if (!args || args.filter === undefined) {
                throw new Error("Missing required property 'filter'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["disabled"] = args ? args.disabled : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        super("gcp:logging/projectExclusion:ProjectExclusion", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectExclusion resources.
 */
export interface ProjectExclusionState {
    readonly description?: pulumi.Input<string>;
    readonly disabled?: pulumi.Input<boolean>;
    readonly filter?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectExclusion resource.
 */
export interface ProjectExclusionArgs {
    readonly description?: pulumi.Input<string>;
    readonly disabled?: pulumi.Input<boolean>;
    readonly filter: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
}
