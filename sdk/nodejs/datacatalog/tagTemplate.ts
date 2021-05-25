// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A tag template defines a tag, which can have one or more typed fields.
 * The template is used to create and attach the tag to GCP resources.
 *
 * To get more information about TagTemplate, see:
 *
 * * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.tagTemplates)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/data-catalog/docs)
 *
 * ## Example Usage
 * ### Data Catalog Tag Template Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basicTagTemplate = new gcp.datacatalog.TagTemplate("basic_tag_template", {
 *     displayName: "Demo Tag Template",
 *     fields: [
 *         {
 *             displayName: "Source of data asset",
 *             fieldId: "source",
 *             isRequired: true,
 *             type: {
 *                 primitiveType: "STRING",
 *             },
 *         },
 *         {
 *             displayName: "Number of rows in the data asset",
 *             fieldId: "num_rows",
 *             type: {
 *                 primitiveType: "DOUBLE",
 *             },
 *         },
 *         {
 *             displayName: "PII type",
 *             fieldId: "pii_type",
 *             type: {
 *                 enumType: {
 *                     allowedValues: [
 *                         {
 *                             displayName: "EMAIL",
 *                         },
 *                         {
 *                             displayName: "SOCIAL SECURITY NUMBER",
 *                         },
 *                         {
 *                             displayName: "NONE",
 *                         },
 *                     ],
 *                 },
 *             },
 *         },
 *     ],
 *     forceDelete: false,
 *     region: "us-central1",
 *     tagTemplateId: "my_template",
 * });
 * ```
 *
 * ## Import
 *
 * TagTemplate can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:datacatalog/tagTemplate:TagTemplate default {{name}}
 * ```
 */
export class TagTemplate extends pulumi.CustomResource {
    /**
     * Get an existing TagTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TagTemplateState, opts?: pulumi.CustomResourceOptions): TagTemplate {
        return new TagTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:datacatalog/tagTemplate:TagTemplate';

    /**
     * Returns true if the given object is an instance of TagTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TagTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TagTemplate.__pulumiType;
    }

    /**
     * The display name for this template.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.
     * Structure is documented below.
     */
    public readonly fields!: pulumi.Output<outputs.datacatalog.TagTemplateField[]>;
    /**
     * This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
     */
    public readonly forceDelete!: pulumi.Output<boolean | undefined>;
    /**
     * -
     * The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Template location region.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The id of the tag template to create.
     */
    public readonly tagTemplateId!: pulumi.Output<string>;

    /**
     * Create a TagTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TagTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TagTemplateArgs | TagTemplateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TagTemplateState | undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["fields"] = state ? state.fields : undefined;
            inputs["forceDelete"] = state ? state.forceDelete : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["tagTemplateId"] = state ? state.tagTemplateId : undefined;
        } else {
            const args = argsOrState as TagTemplateArgs | undefined;
            if ((!args || args.fields === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fields'");
            }
            if ((!args || args.tagTemplateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tagTemplateId'");
            }
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["fields"] = args ? args.fields : undefined;
            inputs["forceDelete"] = args ? args.forceDelete : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["tagTemplateId"] = args ? args.tagTemplateId : undefined;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(TagTemplate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TagTemplate resources.
 */
export interface TagTemplateState {
    /**
     * The display name for this template.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.
     * Structure is documented below.
     */
    fields?: pulumi.Input<pulumi.Input<inputs.datacatalog.TagTemplateField>[]>;
    /**
     * This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * -
     * The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Template location region.
     */
    region?: pulumi.Input<string>;
    /**
     * The id of the tag template to create.
     */
    tagTemplateId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TagTemplate resource.
 */
export interface TagTemplateArgs {
    /**
     * The display name for this template.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.
     * Structure is documented below.
     */
    fields: pulumi.Input<pulumi.Input<inputs.datacatalog.TagTemplateField>[]>;
    /**
     * This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Template location region.
     */
    region?: pulumi.Input<string>;
    /**
     * The id of the tag template to create.
     */
    tagTemplateId: pulumi.Input<string>;
}
