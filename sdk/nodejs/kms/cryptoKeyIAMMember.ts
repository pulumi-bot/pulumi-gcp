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
 * {{ % example typescript % }}
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
 * {{ % /example % }}
 * 
 * With IAM Conditions:
 * 
 * {{ % example typescript % }}
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
 * {{ % /example % }}
 * 
 * {{ % example typescript % }}
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
 * {{ % /example % }}
 * 
 * With IAM Conditions:
 * 
 * {{ % example typescript % }}
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
 * {{ % /example % }}
 * 
 * {{ % example typescript % }}
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
 * {{ % /example % }}
 * 
 * With IAM Conditions:
 * 
 * {{ % example typescript % }}
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
 * {{ % /example % }}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/google_kms_crypto_key_iam.html.markdown.
 */
export class CryptoKeyIAMMember extends pulumi.CustomResource {
    /**
     * Get an existing CryptoKeyIAMMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CryptoKeyIAMMemberState, opts?: pulumi.CustomResourceOptions): CryptoKeyIAMMember {
        return new CryptoKeyIAMMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:kms/cryptoKeyIAMMember:CryptoKeyIAMMember';

    /**
     * Returns true if the given object is an instance of CryptoKeyIAMMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CryptoKeyIAMMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CryptoKeyIAMMember.__pulumiType;
    }

    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    public readonly condition!: pulumi.Output<outputs.kms.CryptoKeyIAMMemberCondition | undefined>;
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
    public readonly member!: pulumi.Output<string>;
    /**
     * The role that should be applied. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a CryptoKeyIAMMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CryptoKeyIAMMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CryptoKeyIAMMemberArgs | CryptoKeyIAMMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CryptoKeyIAMMemberState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["cryptoKeyId"] = state ? state.cryptoKeyId : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as CryptoKeyIAMMemberArgs | undefined;
            if (!args || args.cryptoKeyId === undefined) {
                throw new Error("Missing required property 'cryptoKeyId'");
            }
            if (!args || args.member === undefined) {
                throw new Error("Missing required property 'member'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["cryptoKeyId"] = args ? args.cryptoKeyId : undefined;
            inputs["member"] = args ? args.member : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CryptoKeyIAMMember.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CryptoKeyIAMMember resources.
 */
export interface CryptoKeyIAMMemberState {
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    readonly condition?: pulumi.Input<inputs.kms.CryptoKeyIAMMemberCondition>;
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
    readonly member?: pulumi.Input<string>;
    /**
     * The role that should be applied. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CryptoKeyIAMMember resource.
 */
export interface CryptoKeyIAMMemberArgs {
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    readonly condition?: pulumi.Input<inputs.kms.CryptoKeyIAMMemberCondition>;
    /**
     * The crypto key ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
     * `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
     * the provider's project setting will be used as a fallback.
     */
    readonly cryptoKeyId: pulumi.Input<string>;
    readonly member: pulumi.Input<string>;
    /**
     * The role that should be applied. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
