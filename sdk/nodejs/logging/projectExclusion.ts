// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a project-level logging exclusion. For more information see
 * [the official documentation](https://cloud.google.com/logging/docs/) and
 * [Excluding Logs](https://cloud.google.com/logging/docs/exclusions).
 * 
 * Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
 * granted to the credentials used with Terraform.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const my_exclusion = new gcp.logging.ProjectExclusion("my-exclusion", {
 *     description: "Exclude GCE instance debug logs",
 *     // Exclude all DEBUG or lower severity messages relating to instances
 *     filter: "resource.type = gce_instance AND severity <= DEBUG",
 * });
 * ```
 */
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

    /**
     * A human-readable description.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * Whether this exclusion rule should be disabled or not. This defaults to
     * false.
     */
    public readonly disabled: pulumi.Output<boolean | undefined>;
    /**
     * The filter to apply when excluding logs. Only log entries that match the filter are excluded.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
     * write a filter.
     */
    public readonly filter: pulumi.Output<string>;
    /**
     * The name of the logging exclusion.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The project to create the exclusion in. If omitted, the project associated with the provider is
     * used.
     */
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
    /**
     * A human-readable description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Whether this exclusion rule should be disabled or not. This defaults to
     * false.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * The filter to apply when excluding logs. Only log entries that match the filter are excluded.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
     * write a filter.
     */
    readonly filter?: pulumi.Input<string>;
    /**
     * The name of the logging exclusion.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The project to create the exclusion in. If omitted, the project associated with the provider is
     * used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectExclusion resource.
 */
export interface ProjectExclusionArgs {
    /**
     * A human-readable description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Whether this exclusion rule should be disabled or not. This defaults to
     * false.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * The filter to apply when excluding logs. Only log entries that match the filter are excluded.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
     * write a filter.
     */
    readonly filter: pulumi.Input<string>;
    /**
     * The name of the logging exclusion.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The project to create the exclusion in. If omitted, the project associated with the provider is
     * used.
     */
    readonly project?: pulumi.Input<string>;
}
