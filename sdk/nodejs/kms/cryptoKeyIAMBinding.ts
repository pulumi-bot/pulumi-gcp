// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows creation and management of a single binding within IAM policy for
 * an existing Google Cloud KMS crypto key.
 * 
 * > **Note:** On create, this resource will overwrite members of any existing roles.
 *     Use `terraform import` and inspect the `terraform plan` output to ensure
 *     your existing members are preserved.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const cryptoKey = new gcp.kms.CryptoKeyIAMBinding("crypto_key", {
 *     cryptoKeyId: "my-gcp-project/us-central1/my-key-ring/my-crypto-key",
 *     members: ["user:alice@gmail.com"],
 *     role: "roles/editor",
 * });
 * ```
 */
export class CryptoKeyIAMBinding extends pulumi.CustomResource {
    /**
     * Get an existing CryptoKeyIAMBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CryptoKeyIAMBindingState, opts?: pulumi.CustomResourceOptions): CryptoKeyIAMBinding {
        return new CryptoKeyIAMBinding(name, <any>state, { ...opts, id: id });
    }

    private static readonly __pulumiType = 'gcp:kms/cryptoKeyIAMBinding:CryptoKeyIAMBinding';

    /**
     * Returns true if the given object is an instance of CryptoKeyIAMBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CryptoKeyIAMBinding {
        if (obj === undefined || obj === null) {
            return false;
        }

        const t = obj['__pulumiType'];
        if (typeof t !== 'string') {
            return false;
        }

        return t === CryptoKeyIAMBinding.__pulumiType;
    }

    /**
     * The crypto key ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
     * `{location_name}/{key_ring_name}/{crypto_key_name}`.
     * In the second form, the provider's project setting will be used as a fallback.
     */
    public readonly cryptoKeyId!: pulumi.Output<string>;
    /**
     * (Computed) The etag of the crypto key's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
     */
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The role that should be applied. Only one
     * `google_kms_crypto_key_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a CryptoKeyIAMBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CryptoKeyIAMBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CryptoKeyIAMBindingArgs | CryptoKeyIAMBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CryptoKeyIAMBindingState | undefined;
            inputs["cryptoKeyId"] = state ? state.cryptoKeyId : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as CryptoKeyIAMBindingArgs | undefined;
            if (!args || args.cryptoKeyId === undefined) {
                throw new Error("Missing required property 'cryptoKeyId'");
            }
            if (!args || args.members === undefined) {
                throw new Error("Missing required property 'members'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["cryptoKeyId"] = args ? args.cryptoKeyId : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        super(CryptoKeyIAMBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CryptoKeyIAMBinding resources.
 */
export interface CryptoKeyIAMBindingState {
    /**
     * The crypto key ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
     * `{location_name}/{key_ring_name}/{crypto_key_name}`.
     * In the second form, the provider's project setting will be used as a fallback.
     */
    readonly cryptoKeyId?: pulumi.Input<string>;
    /**
     * (Computed) The etag of the crypto key's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
     */
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role that should be applied. Only one
     * `google_kms_crypto_key_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CryptoKeyIAMBinding resource.
 */
export interface CryptoKeyIAMBindingArgs {
    /**
     * The crypto key ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
     * `{location_name}/{key_ring_name}/{crypto_key_name}`.
     * In the second form, the provider's project setting will be used as a fallback.
     */
    readonly cryptoKeyId: pulumi.Input<string>;
    /**
     * A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
     */
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role that should be applied. Only one
     * `google_kms_crypto_key_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
