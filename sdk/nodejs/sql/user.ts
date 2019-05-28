// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a new Google SQL User on a Google SQL User Instance. For more information, see the [official documentation](https://cloud.google.com/sql/), or the [JSON API](https://cloud.google.com/sql/docs/admin-api/v1beta4/users).
 * 
 * > **Note:** All arguments including the username and password will be stored in the raw state as plain-text.
 * [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html). Passwords will not be retrieved when running
 * "terraform import".
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
 * const dbNameSuffix = new random.RandomId("db_name_suffix", {
 *     byteLength: 4,
 * });
 * const master = new gcp.sql.DatabaseInstance("master", {
 *     settings: {
 *         tier: "D0",
 *     },
 * });
 * const users = new gcp.sql.User("users", {
 *     host: "me.com",
 *     instance: master.name,
 *     password: "changeme",
 * });
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
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

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
     * The password for the user. Can be updated.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

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
        if (opts && opts.id) {
            const state = argsOrState as UserState | undefined;
            inputs["host"] = state ? state.host : undefined;
            inputs["instance"] = state ? state.instance : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if (!args || args.instance === undefined) {
                throw new Error("Missing required property 'instance'");
            }
            inputs["host"] = args ? args.host : undefined;
            inputs["instance"] = args ? args.instance : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        super("gcp:sql/user:User", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
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
     * The password for the user. Can be updated.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
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
     * The password for the user. Can be updated.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
