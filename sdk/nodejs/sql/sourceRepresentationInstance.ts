// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A source representation instance is a Cloud SQL instance that represents
 * the source database server to the Cloud SQL replica. It is visible in the
 * Cloud Console and appears the same as a regular Cloud SQL instance, but it
 * contains no data, requires no configuration or maintenance, and does not
 * affect billing. You cannot update the source representation instance.
 *
 * ## Example Usage
 *
 * ## Import
 *
 * SourceRepresentationInstance can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance default projects/{{project}}/instances/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance default {{name}}
 * ```
 */
export class SourceRepresentationInstance extends pulumi.CustomResource {
    /**
     * Get an existing SourceRepresentationInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SourceRepresentationInstanceState, opts?: pulumi.CustomResourceOptions): SourceRepresentationInstance {
        return new SourceRepresentationInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance';

    /**
     * Returns true if the given object is an instance of SourceRepresentationInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SourceRepresentationInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SourceRepresentationInstance.__pulumiType;
    }

    /**
     * The MySQL version running on your source database server.
     * Possible values are `MYSQL_5_6` and `MYSQL_5_7`.
     */
    public readonly databaseVersion!: pulumi.Output<string>;
    /**
     * The externally accessible IPv4 address for the source database server.
     */
    public readonly host!: pulumi.Output<string>;
    /**
     * The name of the source representation instance. Use any valid Cloud SQL instance name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The externally accessible port for the source database server.
     * Defaults to 3306.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The Region in which the created instance should reside.
     * If it is not provided, the provider region is used.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a SourceRepresentationInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SourceRepresentationInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SourceRepresentationInstanceArgs | SourceRepresentationInstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SourceRepresentationInstanceState | undefined;
            inputs["databaseVersion"] = state ? state.databaseVersion : undefined;
            inputs["host"] = state ? state.host : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as SourceRepresentationInstanceArgs | undefined;
            if (!args || args.databaseVersion === undefined) {
                throw new Error("Missing required property 'databaseVersion'");
            }
            if (!args || args.host === undefined) {
                throw new Error("Missing required property 'host'");
            }
            inputs["databaseVersion"] = args ? args.databaseVersion : undefined;
            inputs["host"] = args ? args.host : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SourceRepresentationInstance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SourceRepresentationInstance resources.
 */
export interface SourceRepresentationInstanceState {
    /**
     * The MySQL version running on your source database server.
     * Possible values are `MYSQL_5_6` and `MYSQL_5_7`.
     */
    readonly databaseVersion?: pulumi.Input<string>;
    /**
     * The externally accessible IPv4 address for the source database server.
     */
    readonly host?: pulumi.Input<string>;
    /**
     * The name of the source representation instance. Use any valid Cloud SQL instance name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The externally accessible port for the source database server.
     * Defaults to 3306.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The Region in which the created instance should reside.
     * If it is not provided, the provider region is used.
     */
    readonly region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SourceRepresentationInstance resource.
 */
export interface SourceRepresentationInstanceArgs {
    /**
     * The MySQL version running on your source database server.
     * Possible values are `MYSQL_5_6` and `MYSQL_5_7`.
     */
    readonly databaseVersion: pulumi.Input<string>;
    /**
     * The externally accessible IPv4 address for the source database server.
     */
    readonly host: pulumi.Input<string>;
    /**
     * The name of the source representation instance. Use any valid Cloud SQL instance name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The externally accessible port for the source database server.
     * Defaults to 3306.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The Region in which the created instance should reside.
     * If it is not provided, the provider region is used.
     */
    readonly region?: pulumi.Input<string>;
}
