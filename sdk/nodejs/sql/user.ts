// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a new Google SQL User on a Google SQL User Instance. For more information, see the [official documentation](https://cloud.google.com/sql/), or the [JSON API](https://cloud.google.com/sql/docs/admin-api/v1beta4/users).
 *
 * ## Example Usage
 *
 * Example creating a SQL User.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * import * as random from "@pulumi/random";
 *
 * const dbNameSuffix = new random.RandomId("dbNameSuffix", {byteLength: 4});
 * const master = new gcp.sql.DatabaseInstance("master", {settings: {
 *     tier: "db-f1-micro",
 * }});
 * const users = new gcp.sql.User("users", {
 *     instance: master.name,
 *     host: "me.com",
 *     password: "changeme",
 * });
 * ```
 *
 * Example creating a Cloud IAM User.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * import * as random from "@pulumi/random";
 *
 * const dbNameSuffix = new random.RandomId("dbNameSuffix", {byteLength: 4});
 * const master = new gcp.sql.DatabaseInstance("master", {
 *     databaseVersion: "POSTGRES_9_6",
 *     settings: {
 *         tier: "db-f1-micro",
 *         datagbaseFlags: [{
 *             name: "cloudsql.iam_authentication",
 *             value: "on",
 *         }],
 *     },
 * });
 * const users = new gcp.sql.User("users", {
 *     instance: master.name,
 *     type: "CLOUD_IAM_USER",
 * });
 * ```
 *
 * ## Import
 *
 * SQL users for MySQL databases can be imported using the `project`, `instance`, `host` and `name`, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:sql/user:User users my-project/master-instance/my-domain.com/me
 * ```
 *
 *  SQL users for PostgreSQL databases can be imported using the `project`, `instance` and `name`, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:sql/user:User users my-project/master-instance/me
 * ```
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:sql/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * The deletion policy for the user.
     * Setting `ABANDON` allows the resource to be abandoned rather than deleted. This is useful
     * for Postgres, where users cannot be deleted from the API if they have been granted SQL roles.
     */
    public readonly deletionPolicy!: pulumi.Output<string | undefined>;
    /**
     * The host the user can connect from. This is only supported
     * for MySQL instances. Don't set this field for PostgreSQL instances.
     * Can be an IP address. Changing this forces a new resource to be created.
     */
    public readonly host!: pulumi.Output<string | undefined>;
    /**
     * The name of the Cloud SQL instance. Changing this
     * forces a new resource to be created.
     */
    public readonly instance!: pulumi.Output<string>;
    /**
     * The name of the user. Changing this forces a new resource
     * to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The password for the user. Can be updated. For Postgres
     * instances this is a Required field.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The user type. It determines the method to authenticate the
     * user during login. The default is the database's built-in user type. Flags
     * include "BUILT_IN", "CLOUD_IAM_USER", or "CLOUD_IAM_SERVICE_ACCOUNT".
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            inputs["deletionPolicy"] = state ? state.deletionPolicy : undefined;
            inputs["host"] = state ? state.host : undefined;
            inputs["instance"] = state ? state.instance : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.instance === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instance'");
            }
            inputs["deletionPolicy"] = args ? args.deletionPolicy : undefined;
            inputs["host"] = args ? args.host : undefined;
            inputs["instance"] = args ? args.instance : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(User.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * The deletion policy for the user.
     * Setting `ABANDON` allows the resource to be abandoned rather than deleted. This is useful
     * for Postgres, where users cannot be deleted from the API if they have been granted SQL roles.
     */
    readonly deletionPolicy?: pulumi.Input<string>;
    /**
     * The host the user can connect from. This is only supported
     * for MySQL instances. Don't set this field for PostgreSQL instances.
     * Can be an IP address. Changing this forces a new resource to be created.
     */
    readonly host?: pulumi.Input<string>;
    /**
     * The name of the Cloud SQL instance. Changing this
     * forces a new resource to be created.
     */
    readonly instance?: pulumi.Input<string>;
    /**
     * The name of the user. Changing this forces a new resource
     * to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The password for the user. Can be updated. For Postgres
     * instances this is a Required field.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The user type. It determines the method to authenticate the
     * user during login. The default is the database's built-in user type. Flags
     * include "BUILT_IN", "CLOUD_IAM_USER", or "CLOUD_IAM_SERVICE_ACCOUNT".
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * The deletion policy for the user.
     * Setting `ABANDON` allows the resource to be abandoned rather than deleted. This is useful
     * for Postgres, where users cannot be deleted from the API if they have been granted SQL roles.
     */
    readonly deletionPolicy?: pulumi.Input<string>;
    /**
     * The host the user can connect from. This is only supported
     * for MySQL instances. Don't set this field for PostgreSQL instances.
     * Can be an IP address. Changing this forces a new resource to be created.
     */
    readonly host?: pulumi.Input<string>;
    /**
     * The name of the Cloud SQL instance. Changing this
     * forces a new resource to be created.
     */
    readonly instance: pulumi.Input<string>;
    /**
     * The name of the user. Changing this forces a new resource
     * to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The password for the user. Can be updated. For Postgres
     * instances this is a Required field.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The user type. It determines the method to authenticate the
     * user during login. The default is the database's built-in user type. Flags
     * include "BUILT_IN", "CLOUD_IAM_USER", or "CLOUD_IAM_SERVICE_ACCOUNT".
     */
    readonly type?: pulumi.Input<string>;
}
