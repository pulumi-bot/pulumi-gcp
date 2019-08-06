// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * When managing IAM roles, you can treat a service account either as a resource or as an identity. This resource is to add iam policy bindings to a service account resource **to configure permissions for who can edit the service account**. To configure permissions for a service account to act as an identity that can manage other GCP resources, use the googleProjectIam set of resources.
 * 
 * Three different resources help you manage your IAM policy for a service account. Each of these resources serves a different use case:
 * 
 * * `gcp.serviceAccount.IAMPolicy`: Authoritative. Sets the IAM policy for the service account and replaces any existing policy already attached.
 * * `gcp.serviceAccount.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service account are preserved.
 * * `gcp.serviceAccount.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service account are preserved.
 * 
 * > **Note:** `gcp.serviceAccount.IAMPolicy` **cannot** be used in conjunction with `gcp.serviceAccount.IAMBinding` and `gcp.serviceAccount.IAMMember` or they will fight over what your policy should be.
 * 
 * > **Note:** `gcp.serviceAccount.IAMBinding` resources **can be** used in conjunction with `gcp.serviceAccount.IAMMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_service\_account\_iam\_policy
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const admin = pulumi.output(gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         members: ["user:jane@example.com"],
 *         role: "roles/iam.serviceAccountUser",
 *     }],
 * }));
 * const sa = new gcp.serviceAccount.Account("sa", {
 *     accountId: "my-service-account",
 *     displayName: "A service account that only Jane can interact with",
 * });
 * const adminAccountIam = new gcp.serviceAccount.IAMPolicy("admin-account-iam", {
 *     policyData: admin.policyData,
 *     serviceAccountId: sa.name,
 * });
 * ```
 * 
 * ## google\_service\_account\_iam\_binding
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const sa = new gcp.serviceAccount.Account("sa", {
 *     accountId: "my-service-account",
 *     displayName: "A service account that only Jane can use",
 * });
 * const adminAccountIam = new gcp.serviceAccount.IAMBinding("admin-account-iam", {
 *     members: ["user:jane@example.com"],
 *     role: "roles/iam.serviceAccountUser",
 *     serviceAccountId: sa.name,
 * });
 * ```
 * 
 * ## google\_service\_account\_iam\_member
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const defaultDefaultServiceAccount = pulumi.output(gcp.compute.getDefaultServiceAccount({}));
 * const sa = new gcp.serviceAccount.Account("sa", {
 *     accountId: "my-service-account",
 *     displayName: "A service account that Jane can use",
 * });
 * const adminAccountIam = new gcp.serviceAccount.IAMMember("admin-account-iam", {
 *     member: "user:jane@example.com",
 *     role: "roles/iam.serviceAccountUser",
 *     serviceAccountId: sa.name,
 * });
 * // Allow SA service account use the default GCE account
 * const gceDefaultAccountIam = new gcp.serviceAccount.IAMMember("gce-default-account-iam", {
 *     member: pulumi.interpolate`serviceAccount:${sa.email}`,
 *     role: "roles/iam.serviceAccountUser",
 *     serviceAccountId: defaultDefaultServiceAccount.name,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/service_account_iam_member.html.markdown.
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

    /** @internal */
    public static readonly __pulumiType = 'gcp:serviceAccount/iAMMember:IAMMember';

    /**
     * Returns true if the given object is an instance of IAMMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IAMMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IAMMember.__pulumiType;
    }

    /**
     * (Computed) The etag of the service account IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly member!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The fully-qualified name of the service account to apply policy to.
     */
    public readonly serviceAccountId!: pulumi.Output<string>;

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
            const state = argsOrState as IAMMemberState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["serviceAccountId"] = state ? state.serviceAccountId : undefined;
        } else {
            const args = argsOrState as IAMMemberArgs | undefined;
            if (!args || args.member === undefined) {
                throw new Error("Missing required property 'member'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            if (!args || args.serviceAccountId === undefined) {
                throw new Error("Missing required property 'serviceAccountId'");
            }
            inputs["member"] = args ? args.member : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["serviceAccountId"] = args ? args.serviceAccountId : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IAMMember.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IAMMember resources.
 */
export interface IAMMemberState {
    /**
     * (Computed) The etag of the service account IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly member?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * The fully-qualified name of the service account to apply policy to.
     */
    readonly serviceAccountId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IAMMember resource.
 */
export interface IAMMemberArgs {
    readonly member: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
    /**
     * The fully-qualified name of the service account to apply policy to.
     */
    readonly serviceAccountId: pulumi.Input<string>;
}
