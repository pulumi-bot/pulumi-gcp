// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Represents a Location resource.
 *
 * ## Import
 *
 * Location can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:notebooks/location:Location default projects/{{project}}/locations/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:notebooks/location:Location default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:notebooks/location:Location default {{name}}
 * ```
 */
export class Location extends pulumi.CustomResource {
    /**
     * Get an existing Location resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LocationState, opts?: pulumi.CustomResourceOptions): Location {
        return new Location(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:notebooks/location:Location';

    /**
     * Returns true if the given object is an instance of Location.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Location {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Location.__pulumiType;
    }

    /**
     * Name of the Location resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a Location resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LocationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LocationArgs | LocationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LocationState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as LocationArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Location.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Location resources.
 */
export interface LocationState {
    /**
     * Name of the Location resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Location resource.
 */
export interface LocationArgs {
    /**
     * Name of the Location resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
