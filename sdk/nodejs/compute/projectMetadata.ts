// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Manages metadata common to all instances for a project in GCE. For more information see
 * [the official documentation](https://cloud.google.com/compute/docs/storing-retrieving-metadata)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/latest/projects/setCommonInstanceMetadata).
 * 
 * ~> **Note:**  If you want to manage only single key/value pairs within the project metadata
 * rather than the entire set, then use
 * [google_compute_project_metadata_item](compute_project_metadata_item.html).
 */
export class ProjectMetadata extends pulumi.CustomResource {
    /**
     * Get an existing ProjectMetadata resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectMetadataState): ProjectMetadata {
        return new ProjectMetadata(name, <any>state, { id });
    }

    /**
     * A series of key value pairs. Changing this resource
     * updates the GCE state.
     */
    public readonly metadata: pulumi.Output<{[key: string]: string}>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;

    /**
     * Create a ProjectMetadata resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectMetadataArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: ProjectMetadataArgs | ProjectMetadataState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ProjectMetadataState = argsOrState as ProjectMetadataState | undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as ProjectMetadataArgs | undefined;
            if (!args || args.metadata === undefined) {
                throw new Error("Missing required property 'metadata'");
            }
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        super("gcp:compute/projectMetadata:ProjectMetadata", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectMetadata resources.
 */
export interface ProjectMetadataState {
    /**
     * A series of key value pairs. Changing this resource
     * updates the GCE state.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectMetadata resource.
 */
export interface ProjectMetadataArgs {
    /**
     * A series of key value pairs. Changing this resource
     * updates the GCE state.
     */
    readonly metadata: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
