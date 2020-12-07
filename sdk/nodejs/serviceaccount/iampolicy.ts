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
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/iam.serviceAccountUser",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const sa = new gcp.serviceAccount.Account("sa", {
 *     accountId: "my-service-account",
 *     displayName: "A service account that only Jane can interact with",
 * });
 * const admin_account_iam = new gcp.serviceAccount.IAMPolicy("admin-account-iam", {
 *     serviceAccountId: sa.name,
 *     policyData: admin.then(admin => admin.policyData),
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
 * const admin_account_iam = new gcp.serviceAccount.IAMBinding("admin-account-iam", {
 *     serviceAccountId: sa.name,
 *     role: "roles/iam.serviceAccountUser",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const sa = new gcp.serviceAccount.Account("sa", {
 *     accountId: "my-service-account",
 *     displayName: "A service account that only Jane can use",
 * });
 * const admin_account_iam = new gcp.serviceAccount.IAMBinding("admin-account-iam", {
 *     condition: {
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         title: "expires_after_2019_12_31",
 *     },
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
 * const default = gcp.compute.getDefaultServiceAccount({});
 * const sa = new gcp.serviceAccount.Account("sa", {
 *     accountId: "my-service-account",
 *     displayName: "A service account that Jane can use",
 * });
 * const admin_account_iam = new gcp.serviceAccount.IAMMember("admin-account-iam", {
 *     serviceAccountId: sa.name,
 *     role: "roles/iam.serviceAccountUser",
 *     member: "user:jane@example.com",
 * });
 * // Allow SA service account use the default GCE account
 * const gce_default_account_iam = new gcp.serviceAccount.IAMMember("gce-default-account-iam", {
 *     serviceAccountId: _default.then(_default => _default.name),
 *     role: "roles/iam.serviceAccountUser",
 *     member: pulumi.interpolate`serviceAccount:${sa.email}`,
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const sa = new gcp.serviceAccount.Account("sa", {
 *     accountId: "my-service-account",
 *     displayName: "A service account that Jane can use",
 * });
 * const admin_account_iam = new gcp.serviceAccount.IAMMember("admin-account-iam", {
 *     condition: {
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         title: "expires_after_2019_12_31",
 *     },
 *     member: "user:jane@example.com",
 *     role: "roles/iam.serviceAccountUser",
 *     serviceAccountId: sa.name,
 * });
 * ```
 *
 * ## Import
 *
 * Service account IAM resources can be imported using the project, service account email, role, member identity, and condition (beta).
 *
 * ```sh
 *  $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam projects/{your-project-id}/serviceAccounts/{your-service-account-email}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} iam.serviceAccountUser"
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/editor user:foo@example.com"
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`. With conditions
 *
 * ```sh
 *  $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} iam.serviceAccountUser expires_after_2019_12_31"
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} iam.serviceAccountUser user:foo@example.com expires_after_2019_12_31"
 * ```
 */
export class IAMPolicy extends pulumi.CustomResource {
    /**
     * Get an existing IAMPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IAMPolicyState, opts?: pulumi.CustomResourceOptions): IAMPolicy {
        return new IAMPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:serviceAccount/iAMPolicy:IAMPolicy';

    /**
     * Returns true if the given object is an instance of IAMPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IAMPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IAMPolicy.__pulumiType;
    }

    /**
     * (Computed) The etag of the service account IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    public readonly policyData!: pulumi.Output<string>;
    /**
     * The fully-qualified name of the service account to apply policy to.
     */
    public readonly serviceAccountId!: pulumi.Output<string>;

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
            const state = argsOrState as IAMPolicyState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
            inputs["serviceAccountId"] = state ? state.serviceAccountId : undefined;
        } else {
            const args = argsOrState as IAMPolicyArgs | undefined;
            if ((!args || args.policyData === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'policyData'");
            }
            if ((!args || args.serviceAccountId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'serviceAccountId'");
            }
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["serviceAccountId"] = args ? args.serviceAccountId : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IAMPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IAMPolicy resources.
 */
export interface IAMPolicyState {
    /**
     * (Computed) The etag of the service account IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData?: pulumi.Input<string>;
    /**
     * The fully-qualified name of the service account to apply policy to.
     */
    readonly serviceAccountId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IAMPolicy resource.
 */
export interface IAMPolicyArgs {
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: pulumi.Input<string>;
    /**
     * The fully-qualified name of the service account to apply policy to.
     */
    readonly serviceAccountId: pulumi.Input<string>;
}
