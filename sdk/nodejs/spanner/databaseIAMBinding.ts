// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for a Spanner database. Each of these resources serves a different use case:
 * 
 * * `gcp.spanner.DatabaseIAMPolicy`: Authoritative. Sets the IAM policy for the database and replaces any existing policy already attached.
 * 
 * > **Warning:** It's entirely possibly to lock yourself out of your database using `gcp.spanner.DatabaseIAMPolicy`. Any permissions granted by default will be removed unless you include them in your config.
 * 
 * * `gcp.spanner.DatabaseIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the database are preserved.
 * * `gcp.spanner.DatabaseIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the database are preserved.
 * 
 * > **Note:** `gcp.spanner.DatabaseIAMPolicy` **cannot** be used in conjunction with `gcp.spanner.DatabaseIAMBinding` and `gcp.spanner.DatabaseIAMMember` or they will fight over what your policy should be.
 * 
 * > **Note:** `gcp.spanner.DatabaseIAMBinding` resources **can be** used in conjunction with `gcp.spanner.DatabaseIAMMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_spanner\_database\_iam\_policy
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const admin = gcp.organizations.getIAMPolicy({
 *     binding: [{
 *         role: "roles/editor",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const database = new gcp.spanner.DatabaseIAMPolicy("database", {
 *     instance: "your-instance-name",
 *     database: "your-database-name",
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * 
 * ## google\_spanner\_database\_iam\_binding
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const database = new gcp.spanner.DatabaseIAMBinding("database", {
 *     database: "your-database-name",
 *     instance: "your-instance-name",
 *     members: ["user:jane@example.com"],
 *     role: "roles/compute.networkUser",
 * });
 * ```
 * 
 * ## google\_spanner\_database\_iam\_member
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const database = new gcp.spanner.DatabaseIAMMember("database", {
 *     database: "your-database-name",
 *     instance: "your-instance-name",
 *     member: "user:jane@example.com",
 *     role: "roles/compute.networkUser",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database_iam.html.markdown.
 */
export class DatabaseIAMBinding extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseIAMBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseIAMBindingState, opts?: pulumi.CustomResourceOptions): DatabaseIAMBinding {
        return new DatabaseIAMBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:spanner/databaseIAMBinding:DatabaseIAMBinding';

    /**
     * Returns true if the given object is an instance of DatabaseIAMBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseIAMBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseIAMBinding.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.spanner.DatabaseIAMBindingCondition | undefined>;
    /**
     * The name of the Spanner database.
     */
    public readonly database!: pulumi.Output<string>;
    /**
     * (Computed) The etag of the database's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the Spanner instance the database belongs to.
     */
    public readonly instance!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a DatabaseIAMBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseIAMBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseIAMBindingArgs | DatabaseIAMBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DatabaseIAMBindingState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["database"] = state ? state.database : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["instance"] = state ? state.instance : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as DatabaseIAMBindingArgs | undefined;
            if (!args || args.database === undefined) {
                throw new Error("Missing required property 'database'");
            }
            if (!args || args.instance === undefined) {
                throw new Error("Missing required property 'instance'");
            }
            if (!args || args.members === undefined) {
                throw new Error("Missing required property 'members'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["database"] = args ? args.database : undefined;
            inputs["instance"] = args ? args.instance : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DatabaseIAMBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabaseIAMBinding resources.
 */
export interface DatabaseIAMBindingState {
    readonly condition?: pulumi.Input<inputs.spanner.DatabaseIAMBindingCondition>;
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
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatabaseIAMBinding resource.
 */
export interface DatabaseIAMBindingArgs {
    readonly condition?: pulumi.Input<inputs.spanner.DatabaseIAMBindingCondition>;
    /**
     * The name of the Spanner database.
     */
    readonly database: pulumi.Input<string>;
    /**
     * The name of the Spanner instance the database belongs to.
     */
    readonly instance: pulumi.Input<string>;
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
