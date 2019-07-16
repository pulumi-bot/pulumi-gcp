// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_project_metadata_item.html.markdown.
 */
export class ProjectMetadataItem extends pulumi.CustomResource {
    /**
     * Get an existing ProjectMetadataItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectMetadataItemState, opts?: pulumi.CustomResourceOptions): ProjectMetadataItem {
        return new ProjectMetadataItem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/projectMetadataItem:ProjectMetadataItem';

    /**
     * Returns true if the given object is an instance of ProjectMetadataItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectMetadataItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectMetadataItem.__pulumiType;
    }

    /**
     * The metadata key to set.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The value to set for the given metadata key.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a ProjectMetadataItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectMetadataItemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectMetadataItemArgs | ProjectMetadataItemState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProjectMetadataItemState | undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as ProjectMetadataItemArgs | undefined;
            if (!args || args.key === undefined) {
                throw new Error("Missing required property 'key'");
            }
            if (!args || args.value === undefined) {
                throw new Error("Missing required property 'value'");
            }
            inputs["key"] = args ? args.key : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["value"] = args ? args.value : undefined;
        }
        super(ProjectMetadataItem.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectMetadataItem resources.
 */
export interface ProjectMetadataItemState {
    /**
     * The metadata key to set.
     */
    readonly key?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The value to set for the given metadata key.
     */
    readonly value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectMetadataItem resource.
 */
export interface ProjectMetadataItemArgs {
    /**
     * The metadata key to set.
     */
    readonly key: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The value to set for the given metadata key.
     */
    readonly value: pulumi.Input<string>;
}
