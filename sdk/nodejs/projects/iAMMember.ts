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
 */
export class IAMMember extends pulumi.CustomResource {
    /**
     * Get an existing IAMMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IAMMemberState, opts?: pulumi.CustomResourceOptions): IAMMember {
        return new IAMMember(name, <any>state, { ...opts, id: id });
    }

    /**
     * (Computed) The etag of the project's IAM policy.
     */
    public /*out*/ readonly etag: pulumi.Output<string>;
    public readonly member: pulumi.Output<string>;
    /**
     * The project ID. If not specified, uses the
     * ID of the project configured with the provider.
     */
    public readonly project: pulumi.Output<string | undefined>;
    /**
     * The role that should be applied. Only one
     * `google_project_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role: pulumi.Output<string>;

    /**
     * Create a IAMMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IAMMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IAMMemberArgs | IAMMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: IAMMemberState = argsOrState as IAMMemberState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as IAMMemberArgs | undefined;
            if (!args || args.member === undefined) {
                throw new Error("Missing required property 'member'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["member"] = args ? args.member : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        super("gcp:projects/iAMMember:IAMMember", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IAMMember resources.
 */
export interface IAMMemberState {
    /**
     * (Computed) The etag of the project's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly member?: pulumi.Input<string>;
    /**
     * The project ID. If not specified, uses the
     * ID of the project configured with the provider.
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
 * The set of arguments for constructing a IAMMember resource.
 */
export interface IAMMemberArgs {
    readonly member: pulumi.Input<string>;
    /**
     * The project ID. If not specified, uses the
     * ID of the project configured with the provider.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `google_project_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
