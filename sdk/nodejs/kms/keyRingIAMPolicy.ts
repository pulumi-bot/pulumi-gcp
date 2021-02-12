// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for KMS key ring. Each of these resources serves a different use case:
 *
 * * `gcp.kms.KeyRingIAMPolicy`: Authoritative. Sets the IAM policy for the key ring and replaces any existing policy already attached.
 * * `gcp.kms.KeyRingIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the key ring are preserved.
 * * `gcp.kms.KeyRingIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the key ring are preserved.
 *
 * > **Note:** `gcp.kms.KeyRingIAMPolicy` **cannot** be used in conjunction with `gcp.kms.KeyRingIAMBinding` and `gcp.kms.KeyRingIAMMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.kms.KeyRingIAMBinding` resources **can be** used in conjunction with `gcp.kms.KeyRingIAMMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_kms\_key\_ring\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const keyring = new gcp.kms.KeyRing("keyring", {location: "global"});
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/editor",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const keyRing = new gcp.kms.KeyRingIAMPolicy("keyRing", {
 *     keyRingId: keyring.id,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const keyring = new gcp.kms.KeyRing("keyring", {location: "global"});
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/editor",
 *         members: ["user:jane@example.com"],
 *         condition: {
 *             title: "expires_after_2019_12_31",
 *             description: "Expiring at midnight of 2019-12-31",
 *             expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         },
 *     }],
 * });
 * const keyRing = new gcp.kms.KeyRingIAMPolicy("keyRing", {
 *     keyRingId: keyring.id,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## google\_kms\_key\_ring\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const keyRing = new gcp.kms.KeyRingIAMBinding("key_ring", {
 *     keyRingId: "your-key-ring-id",
 *     members: ["user:jane@example.com"],
 *     role: "roles/editor",
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const keyRing = new gcp.kms.KeyRingIAMBinding("key_ring", {
 *     condition: {
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         title: "expires_after_2019_12_31",
 *     },
 *     keyRingId: "your-key-ring-id",
 *     members: ["user:jane@example.com"],
 *     role: "roles/editor",
 * });
 * ```
 *
 * ## google\_kms\_key\_ring\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const keyRing = new gcp.kms.KeyRingIAMMember("key_ring", {
 *     keyRingId: "your-key-ring-id",
 *     member: "user:jane@example.com",
 *     role: "roles/editor",
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const keyRing = new gcp.kms.KeyRingIAMMember("key_ring", {
 *     condition: {
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         title: "expires_after_2019_12_31",
 *     },
 *     keyRingId: "your-key-ring-id",
 *     member: "user:jane@example.com",
 *     role: "roles/editor",
 * });
 * ```
 *
 * ## Import
 *
 * IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
 *
 * This member resource can be imported using the `key_ring_id`, role, and account e.g.
 *
 * ```sh
 *  $ pulumi import gcp:kms/keyRingIAMPolicy:KeyRingIAMPolicy key_ring_iam "your-project-id/location-name/key-ring-name roles/viewer user:foo@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiers; the resource in question and the role.
 *
 * This binding resource can be imported using the `key_ring_id` and role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:kms/keyRingIAMPolicy:KeyRingIAMPolicy key_ring_iam "your-project-id/location-name/key-ring-name roles/viewer"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question.
 *
 * This policy resource can be imported using the `key_ring_id`, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:kms/keyRingIAMPolicy:KeyRingIAMPolicy key_ring_iam your-project-id/location-name/key-ring-name
 * ```
 */
export class KeyRingIAMPolicy extends pulumi.CustomResource {
    /**
     * Get an existing KeyRingIAMPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeyRingIAMPolicyState, opts?: pulumi.CustomResourceOptions): KeyRingIAMPolicy {
        return new KeyRingIAMPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:kms/keyRingIAMPolicy:KeyRingIAMPolicy';

    /**
     * Returns true if the given object is an instance of KeyRingIAMPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KeyRingIAMPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KeyRingIAMPolicy.__pulumiType;
    }

    /**
     * (Computed) The etag of the key ring's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The key ring ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}` or
     * `{location_name}/{key_ring_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    public readonly keyRingId!: pulumi.Output<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    public readonly policyData!: pulumi.Output<string>;

    /**
     * Create a KeyRingIAMPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeyRingIAMPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyRingIAMPolicyArgs | KeyRingIAMPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KeyRingIAMPolicyState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["keyRingId"] = state ? state.keyRingId : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
        } else {
            const args = argsOrState as KeyRingIAMPolicyArgs | undefined;
            if ((!args || args.keyRingId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyRingId'");
            }
            if ((!args || args.policyData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyData'");
            }
            inputs["keyRingId"] = args ? args.keyRingId : undefined;
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(KeyRingIAMPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KeyRingIAMPolicy resources.
 */
export interface KeyRingIAMPolicyState {
    /**
     * (Computed) The etag of the key ring's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The key ring ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}` or
     * `{location_name}/{key_ring_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    readonly keyRingId?: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KeyRingIAMPolicy resource.
 */
export interface KeyRingIAMPolicyArgs {
    /**
     * The key ring ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}` or
     * `{location_name}/{key_ring_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    readonly keyRingId: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: pulumi.Input<string>;
}
