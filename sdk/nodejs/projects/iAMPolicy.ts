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
 * ~> **Note:** `google_project_iam_policy` **cannot** be used in conjunction with `google_project_iam_binding` and `google_project_iam_member` or they will fight over what your policy should be.
 * 
 * ~> **Note:** `google_project_iam_binding` resources **can be** used in conjunction with `google_project_iam_member` resources **only if** they do not grant privilege to the same role.
 */
export class IAMPolicy extends pulumi.CustomResource {
    /**
     * Get an existing IAMPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IAMPolicyState, opts?: pulumi.CustomResourceOptions): IAMPolicy {
        return new IAMPolicy(name, <any>state, { ...opts, id: id });
    }

    /**
     * (Optional, only for `google_project_iam_policy`)
     * A boolean value indicating if this policy
     * should overwrite any existing IAM policy on the project. When set to true,
     * **any policies not in your config file will be removed**. This can **lock
     * you out** of your project until an Organization Administrator grants you
     * access again, so please exercise caution. If this argument is `true` and you
     * want to delete the resource, you must set the `disable_project` argument to
     * `true`, acknowledging that the project will be inaccessible to anyone but the
     * Organization Admins, as it will no longer have an IAM policy. Rather than using
     * this, you should use `google_project_iam_binding` and
     * `google_project_iam_member`.
     */
    public readonly authoritative: pulumi.Output<boolean | undefined>;
    /**
     * (Optional, only for `google_project_iam_policy`)
     * A boolean value that must be set to `true`
     * if you want to delete a `google_project_iam_policy` that is authoritative.
     */
    public readonly disableProject: pulumi.Output<boolean | undefined>;
    /**
     * (Computed) The etag of the project's IAM policy.
     */
    public /*out*/ readonly etag: pulumi.Output<string>;
    /**
     * The `google_iam_policy` data source that represents
     * the IAM policy that will be applied to the project. The policy will be
     * merged with any existing policy applied to the project.
     */
    public readonly policyData: pulumi.Output<string>;
    /**
     * The project ID. If not specified, uses the
     * ID of the project configured with the provider.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * (DEPRECATED) (Computed, only for `google_project_iam_policy`)
     * The IAM policy that will be restored when a
     * non-authoritative policy resource is deleted.
     */
    public /*out*/ readonly restorePolicy: pulumi.Output<string>;

    /**
     * Create a IAMPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IAMPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IAMPolicyArgs | IAMPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: IAMPolicyState = argsOrState as IAMPolicyState | undefined;
            inputs["authoritative"] = state ? state.authoritative : undefined;
            inputs["disableProject"] = state ? state.disableProject : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["restorePolicy"] = state ? state.restorePolicy : undefined;
        } else {
            const args = argsOrState as IAMPolicyArgs | undefined;
            if (!args || args.policyData === undefined) {
                throw new Error("Missing required property 'policyData'");
            }
            inputs["authoritative"] = args ? args.authoritative : undefined;
            inputs["disableProject"] = args ? args.disableProject : undefined;
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["restorePolicy"] = undefined /*out*/;
        }
        super("gcp:projects/iAMPolicy:IAMPolicy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IAMPolicy resources.
 */
export interface IAMPolicyState {
    /**
     * (Optional, only for `google_project_iam_policy`)
     * A boolean value indicating if this policy
     * should overwrite any existing IAM policy on the project. When set to true,
     * **any policies not in your config file will be removed**. This can **lock
     * you out** of your project until an Organization Administrator grants you
     * access again, so please exercise caution. If this argument is `true` and you
     * want to delete the resource, you must set the `disable_project` argument to
     * `true`, acknowledging that the project will be inaccessible to anyone but the
     * Organization Admins, as it will no longer have an IAM policy. Rather than using
     * this, you should use `google_project_iam_binding` and
     * `google_project_iam_member`.
     */
    readonly authoritative?: pulumi.Input<boolean>;
    /**
     * (Optional, only for `google_project_iam_policy`)
     * A boolean value that must be set to `true`
     * if you want to delete a `google_project_iam_policy` that is authoritative.
     */
    readonly disableProject?: pulumi.Input<boolean>;
    /**
     * (Computed) The etag of the project's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The `google_iam_policy` data source that represents
     * the IAM policy that will be applied to the project. The policy will be
     * merged with any existing policy applied to the project.
     */
    readonly policyData?: pulumi.Input<string>;
    /**
     * The project ID. If not specified, uses the
     * ID of the project configured with the provider.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * (DEPRECATED) (Computed, only for `google_project_iam_policy`)
     * The IAM policy that will be restored when a
     * non-authoritative policy resource is deleted.
     */
    readonly restorePolicy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IAMPolicy resource.
 */
export interface IAMPolicyArgs {
    /**
     * (Optional, only for `google_project_iam_policy`)
     * A boolean value indicating if this policy
     * should overwrite any existing IAM policy on the project. When set to true,
     * **any policies not in your config file will be removed**. This can **lock
     * you out** of your project until an Organization Administrator grants you
     * access again, so please exercise caution. If this argument is `true` and you
     * want to delete the resource, you must set the `disable_project` argument to
     * `true`, acknowledging that the project will be inaccessible to anyone but the
     * Organization Admins, as it will no longer have an IAM policy. Rather than using
     * this, you should use `google_project_iam_binding` and
     * `google_project_iam_member`.
     */
    readonly authoritative?: pulumi.Input<boolean>;
    /**
     * (Optional, only for `google_project_iam_policy`)
     * A boolean value that must be set to `true`
     * if you want to delete a `google_project_iam_policy` that is authoritative.
     */
    readonly disableProject?: pulumi.Input<boolean>;
    /**
     * The `google_iam_policy` data source that represents
     * the IAM policy that will be applied to the project. The policy will be
     * merged with any existing policy applied to the project.
     */
    readonly policyData: pulumi.Input<string>;
    /**
     * The project ID. If not specified, uses the
     * ID of the project configured with the provider.
     */
    readonly project?: pulumi.Input<string>;
}
