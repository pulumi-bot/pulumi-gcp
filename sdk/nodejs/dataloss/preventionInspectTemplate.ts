// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * An inspect job template.
 *
 * To get more information about InspectTemplate, see:
 *
 * * [API documentation](https://cloud.google.com/dlp/docs/reference/rest/v2/projects.inspectTemplates)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dlp/docs/creating-templates-inspect)
 *
 * ## Example Usage
 */
export class PreventionInspectTemplate extends pulumi.CustomResource {
    /**
     * Get an existing PreventionInspectTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PreventionInspectTemplateState, opts?: pulumi.CustomResourceOptions): PreventionInspectTemplate {
        return new PreventionInspectTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataloss/preventionInspectTemplate:PreventionInspectTemplate';

    /**
     * Returns true if the given object is an instance of PreventionInspectTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PreventionInspectTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PreventionInspectTemplate.__pulumiType;
    }

    /**
     * A description of the inspect template.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * User set display name of the inspect template.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The core content of the template.
     * Structure is documented below.
     */
    public readonly inspectConfig!: pulumi.Output<outputs.dataloss.PreventionInspectTemplateInspectConfig | undefined>;
    /**
     * Resource name of the requested StoredInfoType, for example `organizations/433245324/storedInfoTypes/432452342`
     * or `projects/project-id/storedInfoTypes/432452342`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The parent of the inspect template in any of the following formats:
     * * `projects/{{project}}`
     * * `projects/{{project}}/locations/{{location}}`
     * * `organizations/{{organization_id}}`
     * * `organizations/{{organization_id}}/locations/{{location}}`
     */
    public readonly parent!: pulumi.Output<string>;

    /**
     * Create a PreventionInspectTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PreventionInspectTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PreventionInspectTemplateArgs | PreventionInspectTemplateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PreventionInspectTemplateState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["inspectConfig"] = state ? state.inspectConfig : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parent"] = state ? state.parent : undefined;
        } else {
            const args = argsOrState as PreventionInspectTemplateArgs | undefined;
            if (!args || args.parent === undefined) {
                throw new Error("Missing required property 'parent'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["inspectConfig"] = args ? args.inspectConfig : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PreventionInspectTemplate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PreventionInspectTemplate resources.
 */
export interface PreventionInspectTemplateState {
    /**
     * A description of the inspect template.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * User set display name of the inspect template.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The core content of the template.
     * Structure is documented below.
     */
    readonly inspectConfig?: pulumi.Input<inputs.dataloss.PreventionInspectTemplateInspectConfig>;
    /**
     * Resource name of the requested StoredInfoType, for example `organizations/433245324/storedInfoTypes/432452342`
     * or `projects/project-id/storedInfoTypes/432452342`.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The parent of the inspect template in any of the following formats:
     * * `projects/{{project}}`
     * * `projects/{{project}}/locations/{{location}}`
     * * `organizations/{{organization_id}}`
     * * `organizations/{{organization_id}}/locations/{{location}}`
     */
    readonly parent?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PreventionInspectTemplate resource.
 */
export interface PreventionInspectTemplateArgs {
    /**
     * A description of the inspect template.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * User set display name of the inspect template.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The core content of the template.
     * Structure is documented below.
     */
    readonly inspectConfig?: pulumi.Input<inputs.dataloss.PreventionInspectTemplateInspectConfig>;
    /**
     * The parent of the inspect template in any of the following formats:
     * * `projects/{{project}}`
     * * `projects/{{project}}/locations/{{location}}`
     * * `organizations/{{organization_id}}`
     * * `organizations/{{organization_id}}/locations/{{location}}`
     */
    readonly parent: pulumi.Input<string>;
}
