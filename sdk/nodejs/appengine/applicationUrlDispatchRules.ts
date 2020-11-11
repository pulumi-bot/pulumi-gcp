// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Rules to match an HTTP request and dispatch that request to a service.
 *
 * To get more information about ApplicationUrlDispatchRules, see:
 *
 * * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps#UrlDispatchRule)
 *
 * ## Example Usage
 */
export class ApplicationUrlDispatchRules extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationUrlDispatchRules resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationUrlDispatchRulesState, opts?: pulumi.CustomResourceOptions): ApplicationUrlDispatchRules {
        return new ApplicationUrlDispatchRules(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules';

    /**
     * Returns true if the given object is an instance of ApplicationUrlDispatchRules.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationUrlDispatchRules {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationUrlDispatchRules.__pulumiType;
    }

    /**
     * Rules to match an HTTP request and dispatch that request to a service.
     * Structure is documented below.
     */
    public readonly dispatchRules!: pulumi.Output<outputs.appengine.ApplicationUrlDispatchRulesDispatchRule[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a ApplicationUrlDispatchRules resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationUrlDispatchRulesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationUrlDispatchRulesArgs | ApplicationUrlDispatchRulesState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ApplicationUrlDispatchRulesState | undefined;
            inputs["dispatchRules"] = state ? state.dispatchRules : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as ApplicationUrlDispatchRulesArgs | undefined;
            if (!args || args.dispatchRules === undefined) {
                throw new Error("Missing required property 'dispatchRules'");
            }
            inputs["dispatchRules"] = args ? args.dispatchRules : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ApplicationUrlDispatchRules.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationUrlDispatchRules resources.
 */
export interface ApplicationUrlDispatchRulesState {
    /**
     * Rules to match an HTTP request and dispatch that request to a service.
     * Structure is documented below.
     */
    readonly dispatchRules?: pulumi.Input<pulumi.Input<inputs.appengine.ApplicationUrlDispatchRulesDispatchRule>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApplicationUrlDispatchRules resource.
 */
export interface ApplicationUrlDispatchRulesArgs {
    /**
     * Rules to match an HTTP request and dispatch that request to a service.
     * Structure is documented below.
     */
    readonly dispatchRules: pulumi.Input<pulumi.Input<inputs.appengine.ApplicationUrlDispatchRulesDispatchRule>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
