// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for a Spanner database. Each of these resources serves a different use case:
 * 
 * * `google_spanner_database_iam_policy`: Authoritative. Sets the IAM policy for the database and replaces any existing policy already attached.
 * 
 * > **Warning:** It's entirely possibly to lock yourself out of your database using `google_spanner_database_iam_policy`. Any permissions granted by default will be removed unless you include them in your config.
 * 
 * * `google_spanner_database_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the database are preserved.
 * * `google_spanner_database_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the database are preserved.
 * 
 * > **Note:** `google_spanner_database_iam_policy` **cannot** be used in conjunction with `google_spanner_database_iam_binding` and `google_spanner_database_iam_member` or they will fight over what your policy should be.
 * 
 * > **Note:** `google_spanner_database_iam_binding` resources **can be** used in conjunction with `google_spanner_database_iam_member` resources **only if** they do not grant privilege to the same role.
 */
export class DatabaseIAMMember extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseIAMMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseIAMMemberState, opts?: pulumi.CustomResourceOptions): DatabaseIAMMember {
        return new DatabaseIAMMember(name, <any>state, { ...opts, id: id });
    }

    /**
     * The name of the Spanner database.
     */
    public readonly database: pulumi.Output<string>;
    /**
     * (Computed) The etag of the database's IAM policy.
     */
    public /*out*/ readonly etag: pulumi.Output<string>;
    /**
     * The name of the Spanner instance the database belongs to.
     */
    public readonly instance: pulumi.Output<string>;
    public readonly member: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `google_spanner_database_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role: pulumi.Output<string>;

    /**
     * Create a DatabaseIAMMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseIAMMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseIAMMemberArgs | DatabaseIAMMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: DatabaseIAMMemberState = argsOrState as DatabaseIAMMemberState | undefined;
            inputs["database"] = state ? state.database : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["instance"] = state ? state.instance : undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as DatabaseIAMMemberArgs | undefined;
            if (!args || args.database === undefined) {
                throw new Error("Missing required property 'database'");
            }
            if (!args || args.instance === undefined) {
                throw new Error("Missing required property 'instance'");
            }
            if (!args || args.member === undefined) {
                throw new Error("Missing required property 'member'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["database"] = args ? args.database : undefined;
            inputs["instance"] = args ? args.instance : undefined;
            inputs["member"] = args ? args.member : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        super("gcp:spanner/databaseIAMMember:DatabaseIAMMember", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabaseIAMMember resources.
 */
export interface DatabaseIAMMemberState {
    /**
     * The name of the Spanner database.
     */
    readonly database?: pulumi.Input<string>;
    /**
     * (Computed) The etag of the database's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The name of the Spanner instance the database belongs to.
     */
    readonly instance?: pulumi.Input<string>;
    readonly member?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `google_spanner_database_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatabaseIAMMember resource.
 */
export interface DatabaseIAMMemberArgs {
    /**
     * The name of the Spanner database.
     */
    readonly database: pulumi.Input<string>;
    /**
     * The name of the Spanner instance the database belongs to.
     */
    readonly instance: pulumi.Input<string>;
    readonly member: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `google_spanner_database_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
