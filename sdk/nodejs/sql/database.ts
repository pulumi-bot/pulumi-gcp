// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Represents a SQL database inside the Cloud SQL instance, hosted in
 * Google's cloud.
 *
 * ## Example Usage
 * ### Sql Database Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const instance = new gcp.sql.DatabaseInstance("instance", {
 *     region: "us-central1",
 *     settings: {
 *         tier: "db-f1-micro",
 *     },
 * });
 * const database = new gcp.sql.Database("database", {instance: instance.name});
 * ```
 */
export class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseState, opts?: pulumi.CustomResourceOptions): Database {
        return new Database(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:sql/database:Database';

    /**
     * Returns true if the given object is an instance of Database.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Database {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Database.__pulumiType;
    }

    /**
     * The charset value. See MySQL's
     * [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)
     * and Postgres' [Character Set Support](https://www.postgresql.org/docs/9.6/static/multibyte.html)
     * for more details and supported values. Postgres databases only support
     * a value of `UTF8` at creation time.
     */
    public readonly charset!: pulumi.Output<string>;
    /**
     * The collation value. See MySQL's
     * [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)
     * and Postgres' [Collation Support](https://www.postgresql.org/docs/9.6/static/collation.html)
     * for more details and supported values. Postgres databases only support
     * a value of `en_US.UTF8` at creation time.
     */
    public readonly collation!: pulumi.Output<string>;
    /**
     * The name of the Cloud SQL instance. This does not include the project
     * ID.
     */
    public readonly instance!: pulumi.Output<string>;
    /**
     * The name of the database in the Cloud SQL instance.
     * This does not include the project ID or instance name.
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
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseArgs | DatabaseState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DatabaseState | undefined;
            inputs["charset"] = state ? state.charset : undefined;
            inputs["collation"] = state ? state.collation : undefined;
            inputs["instance"] = state ? state.instance : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as DatabaseArgs | undefined;
            if (!args || args.instance === undefined) {
                throw new Error("Missing required property 'instance'");
            }
            inputs["charset"] = args ? args.charset : undefined;
            inputs["collation"] = args ? args.collation : undefined;
            inputs["instance"] = args ? args.instance : undefined;
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
        super(Database.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Database resources.
 */
export interface DatabaseState {
    /**
     * The charset value. See MySQL's
     * [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)
     * and Postgres' [Character Set Support](https://www.postgresql.org/docs/9.6/static/multibyte.html)
     * for more details and supported values. Postgres databases only support
     * a value of `UTF8` at creation time.
     */
    readonly charset?: pulumi.Input<string>;
    /**
     * The collation value. See MySQL's
     * [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)
     * and Postgres' [Collation Support](https://www.postgresql.org/docs/9.6/static/collation.html)
     * for more details and supported values. Postgres databases only support
     * a value of `en_US.UTF8` at creation time.
     */
    readonly collation?: pulumi.Input<string>;
    /**
     * The name of the Cloud SQL instance. This does not include the project
     * ID.
     */
    readonly instance?: pulumi.Input<string>;
    /**
     * The name of the database in the Cloud SQL instance.
     * This does not include the project ID or instance name.
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
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * The charset value. See MySQL's
     * [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)
     * and Postgres' [Character Set Support](https://www.postgresql.org/docs/9.6/static/multibyte.html)
     * for more details and supported values. Postgres databases only support
     * a value of `UTF8` at creation time.
     */
    readonly charset?: pulumi.Input<string>;
    /**
     * The collation value. See MySQL's
     * [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)
     * and Postgres' [Collation Support](https://www.postgresql.org/docs/9.6/static/collation.html)
     * for more details and supported values. Postgres databases only support
     * a value of `en_US.UTF8` at creation time.
     */
    readonly collation?: pulumi.Input<string>;
    /**
     * The name of the Cloud SQL instance. This does not include the project
     * ID.
     */
    readonly instance: pulumi.Input<string>;
    /**
     * The name of the database in the Cloud SQL instance.
     * This does not include the project ID or instance name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
