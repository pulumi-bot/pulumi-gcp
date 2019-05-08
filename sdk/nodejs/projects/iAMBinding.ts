// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for a project. Each of these resources serves a different use case:
 * 
 * * `google_project_iam_policy`: Authoritative. Sets the IAM policy for the project and replaces any existing policy already attached.
 * * `google_project_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the project are preserved.
 * * `google_project_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the project are preserved.
 * 
 * > **Note:** `google_project_iam_policy` **cannot** be used in conjunction with `google_project_iam_binding` and `google_project_iam_member` or they will fight over what your policy should be.
 * 
 * > **Note:** `google_project_iam_binding` resources **can be** used in conjunction with `google_project_iam_member` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_project\_iam\_policy
 * 
 * > **Be careful!** You can accidentally lock yourself out of your project
 *    using this resource. Deleting a `google_project_iam_policy` removes access
 *    from anyone without organization-level access to the project. Proceed with caution.
 *    It's not recommended to use `google_project_iam_policy` with your provider project
 *    to avoid locking yourself out, and it should generally only be used with projects
 *    fully managed by Terraform.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const admin = pulumi.output(gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         members: ["user:jane@example.com"],
 *         role: "roles/editor",
 *     }],
 * }));
 * const project = new gcp.projects.IAMPolicy("project", {
 *     policyData: admin.apply(admin => admin.policyData),
 *     project: "your-project-id",
 * });
 * ```
 * 
 * ## google\_project\_iam\_binding
 * 
 * > **Note:** If `role` is set to `roles/owner` and you don't specify a user or service account you have access to in `members`, you can lock yourself out of your project.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const project = new gcp.projects.IAMBinding("project", {
 *     members: ["user:jane@example.com"],
 *     project: "your-project-id",
 *     role: "roles/editor",
 * });
 * ```
 * 
 * ## google\_project\_iam\_member
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const project = new gcp.projects.IAMMember("project", {
 *     member: "user:jane@example.com",
 *     project: "your-project-id",
 *     role: "roles/editor",
 * });
 * ```
 */
export class IAMBinding extends pulumi.CustomResource {
    /**
     * Get an existing IAMBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IAMBindingState, opts?: pulumi.CustomResourceOptions): IAMBinding {
        return new IAMBinding(name, <any>state, { ...opts, id: id });
    }

    /**
     * (Computed) The etag of the project's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The project ID. If not specified for `google_project_iam_binding`
     * or `google_project_iam_member`, uses the ID of the project configured with the provider.
     * Required for `google_project_iam_policy` - you must explicitly set the project, and it
     * will not be inferred from the provider.
     */
    public readonly project!: pulumi.Output<string | undefined>;
    /**
     * The role that should be applied. Only one
     * `google_project_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a IAMBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IAMBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IAMBindingArgs | IAMBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IAMBindingState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as IAMBindingArgs | undefined;
            if (!args || args.members === undefined) {
                throw new Error("Missing required property 'members'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
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
        super("gcp:projects/iAMBinding:IAMBinding", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IAMBinding resources.
 */
export interface IAMBindingState {
    /**
     * (Computed) The etag of the project's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The project ID. If not specified for `google_project_iam_binding`
     * or `google_project_iam_member`, uses the ID of the project configured with the provider.
     * Required for `google_project_iam_policy` - you must explicitly set the project, and it
     * will not be inferred from the provider.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `google_project_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IAMBinding resource.
 */
export interface IAMBindingArgs {
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The project ID. If not specified for `google_project_iam_binding`
     * or `google_project_iam_member`, uses the ID of the project configured with the provider.
     * Required for `google_project_iam_policy` - you must explicitly set the project, and it
     * will not be inferred from the provider.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `google_project_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
