// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a RuntimeConfig resource in Google Cloud. For more information, see the
 * [official documentation](https://cloud.google.com/deployment-manager/runtime-configurator/),
 * or the
 * [JSON API](https://cloud.google.com/deployment-manager/runtime-configurator/reference/rest/).
 *
 * ## Import
 *
 * Runtime Configs can be imported using the `name` or full config name, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:runtimeconfig/config:Config myconfig myconfig
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:runtimeconfig/config:Config myconfig projects/my-gcp-project/configs/myconfig
 * ```
 *
 *  When importing using only the name, the provider project must be set.
 */
export class Config extends pulumi.CustomResource {
    /**
     * Get an existing Config resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigState, opts?: pulumi.CustomResourceOptions): Config {
        return new Config(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:runtimeconfig/config:Config';

    /**
     * Returns true if the given object is an instance of Config.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Config {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Config.__pulumiType;
    }

    /**
     * The description to associate with the runtime
     * config.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the runtime config.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a Config resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigArgs | ConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ConfigState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as ConfigArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Config.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Config resources.
 */
export interface ConfigState {
    /**
     * The description to associate with the runtime
     * config.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the runtime config.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Config resource.
 */
export interface ConfigArgs {
    /**
     * The description to associate with the runtime
     * config.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the runtime config.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
