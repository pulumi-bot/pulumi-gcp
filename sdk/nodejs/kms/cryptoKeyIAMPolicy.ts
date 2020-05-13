// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for KMS crypto key. Each of these resources serves a different use case:
 *
 * * `gcp.kms.CryptoKeyIAMPolicy`: Authoritative. Sets the IAM policy for the crypto key and replaces any existing policy already attached.
 * * `gcp.kms.CryptoKeyIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the crypto key are preserved.
 * * `gcp.kms.CryptoKeyIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the crypto key are preserved.
 *
 * > **Note:** `gcp.kms.CryptoKeyIAMPolicy` **cannot** be used in conjunction with `gcp.kms.CryptoKeyIAMBinding` and `gcp.kms.CryptoKeyIAMMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.kms.CryptoKeyIAMBinding` resources **can be** used in conjunction with `gcp.kms.CryptoKeyIAMMember` resources **only if** they do not grant privilege to the same role.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const keyring = new gcp.kms.KeyRing("keyring", {location: "global"});
 * const key = new gcp.kms.CryptoKey("key", {
 *     keyRing: keyring.id,
 *     rotationPeriod: "100000s",
 * });
 * const admin = gcp.organizations.getIAMPolicy({
 *     binding: [{
 *         role: "roles/cloudkms.cryptoKeyEncrypter",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const cryptoKey = new gcp.kms.CryptoKeyIAMPolicy("cryptoKey", {
 *     cryptoKeyId: key.id,
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
 * const admin = pulumi.output(gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         condition: {
 *             description: "Expiring at midnight of 2019-12-31",
 *             expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *             title: "expiresAfter20191231",
 *         },
 *         members: ["user:jane@example.com"],
 *         role: "roles/cloudkms.cryptoKeyEncrypter",
 *     }],
 * }, { async: true }));
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const cryptoKey = new gcp.kms.CryptoKeyIAMBinding("cryptoKey", {
 *     cryptoKeyId: google_kms_crypto_key.key.id,
 *     role: "roles/cloudkms.cryptoKeyEncrypter",
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
 * const cryptoKey = new gcp.kms.CryptoKeyIAMBinding("cryptoKey", {
 *     cryptoKeyId: google_kms_crypto_key.key.id,
 *     role: "roles/cloudkms.cryptoKeyEncrypter",
 *     members: ["user:jane@example.com"],
 *     condition: {
 *         title: "expiresAfter20191231",
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *     },
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const cryptoKey = new gcp.kms.CryptoKeyIAMMember("cryptoKey", {
 *     cryptoKeyId: google_kms_crypto_key.key.id,
 *     role: "roles/cloudkms.cryptoKeyEncrypter",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const cryptoKey = new gcp.kms.CryptoKeyIAMMember("cryptoKey", {
 *     cryptoKeyId: google_kms_crypto_key.key.id,
 *     role: "roles/cloudkms.cryptoKeyEncrypter",
 *     member: "user:jane@example.com",
 *     condition: {
 *         title: "expiresAfter20191231",
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *     },
 * });
 * ```
 */
export class CryptoKeyIAMPolicy extends pulumi.CustomResource {
    /**
     * Get an existing CryptoKeyIAMPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CryptoKeyIAMPolicyState, opts?: pulumi.CustomResourceOptions): CryptoKeyIAMPolicy {
        return new CryptoKeyIAMPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:kms/cryptoKeyIAMPolicy:CryptoKeyIAMPolicy';

    /**
     * Returns true if the given object is an instance of CryptoKeyIAMPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CryptoKeyIAMPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CryptoKeyIAMPolicy.__pulumiType;
    }

    /**
     * The crypto key ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
     * `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
     * the provider's project setting will be used as a fallback.
     */
    public readonly cryptoKeyId!: pulumi.Output<string>;
    /**
     * (Computed) The etag of the project's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    public readonly policyData!: pulumi.Output<string>;

    /**
     * Create a CryptoKeyIAMPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CryptoKeyIAMPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CryptoKeyIAMPolicyArgs | CryptoKeyIAMPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CryptoKeyIAMPolicyState | undefined;
            inputs["cryptoKeyId"] = state ? state.cryptoKeyId : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
        } else {
            const args = argsOrState as CryptoKeyIAMPolicyArgs | undefined;
            if (!args || args.cryptoKeyId === undefined) {
                throw new Error("Missing required property 'cryptoKeyId'");
            }
            if (!args || args.policyData === undefined) {
                throw new Error("Missing required property 'policyData'");
            }
            inputs["cryptoKeyId"] = args ? args.cryptoKeyId : undefined;
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CryptoKeyIAMPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CryptoKeyIAMPolicy resources.
 */
export interface CryptoKeyIAMPolicyState {
    /**
     * The crypto key ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
     * `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
     * the provider's project setting will be used as a fallback.
     */
    readonly cryptoKeyId?: pulumi.Input<string>;
    /**
     * (Computed) The etag of the project's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CryptoKeyIAMPolicy resource.
 */
export interface CryptoKeyIAMPolicyArgs {
    /**
     * The crypto key ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
     * `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
     * the provider's project setting will be used as a fallback.
     */
    readonly cryptoKeyId: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: pulumi.Input<string>;
}
