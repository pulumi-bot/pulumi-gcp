// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates and manages service account key-pairs, which allow the user to establish identity of a service account outside of GCP. For more information, see [the official documentation](https://cloud.google.com/iam/docs/creating-managing-service-account-keys) and [API](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys).
 * 
 * 
 * ## Example Usage, creating a new Key Pair
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const google_service_account_myaccount = new gcp.serviceAccount.Account("myaccount", {
 *     accountId: "myaccount",
 *     displayName: "My Service Account",
 * });
 * const google_service_account_key_mykey = new gcp.serviceAccount.Key("mykey", {
 *     publicKeyType: "TYPE_X509_PEM_FILE",
 *     serviceAccountId: google_service_account_myaccount.name,
 * });
 * ```
 */
export class Key extends pulumi.CustomResource {
    /**
     * Get an existing Key resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeyState, opts?: pulumi.CustomResourceOptions): Key {
        return new Key(name, <any>state, { ...opts, id: id });
    }

    /**
     * The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
     * Valid values are listed at
     * [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
     * (only used on create)
     */
    public readonly keyAlgorithm: pulumi.Output<string | undefined>;
    /**
     * The name used for this key pair
     */
    public /*out*/ readonly name: pulumi.Output<string>;
    /**
     * An optional PGP key to encrypt the resulting private
     * key material. Only used when creating or importing a new key pair. May either be
     * a base64-encoded public key or a `keybase:keybaseusername` string for looking up
     * in Vault.
     */
    public readonly pgpKey: pulumi.Output<string | undefined>;
    /**
     * The private key in JSON format, base64 encoded. This is what you normally get as a file when creating
     * service account keys through the CLI or web console. This is only populated when creating a new key, and when no
     * `pgp_key` is provided.
     */
    public /*out*/ readonly privateKey: pulumi.Output<string>;
    /**
     * The private key material, base 64 encoded and
     * encrypted with the given `pgp_key`. This is only populated when creating a new
     * key and `pgp_key` is supplied
     */
    public /*out*/ readonly privateKeyEncrypted: pulumi.Output<string>;
    /**
     * The MD5 public key fingerprint for the encrypted
     * private key. This is only populated when creating a new key and `pgp_key` is supplied
     */
    public /*out*/ readonly privateKeyFingerprint: pulumi.Output<string>;
    /**
     * The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
     */
    public readonly privateKeyType: pulumi.Output<string | undefined>;
    /**
     * The public key, base64 encoded
     */
    public /*out*/ readonly publicKey: pulumi.Output<string>;
    /**
     * The output format of the public key requested. X509_PEM is the default output format.
     */
    public readonly publicKeyType: pulumi.Output<string | undefined>;
    /**
     * The Service account id of the Key Pair. This can be a string in the format
     * `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`, where `{ACCOUNT}` is the email address or
     * unique id of the service account. If the `{ACCOUNT}` syntax is used, the project will be inferred from the account.
     */
    public readonly serviceAccountId: pulumi.Output<string>;
    /**
     * The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly validAfter: pulumi.Output<string>;
    /**
     * The key can be used before this timestamp.
     * A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly validBefore: pulumi.Output<string>;

    /**
     * Create a Key resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyArgs | KeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: KeyState = argsOrState as KeyState | undefined;
            inputs["keyAlgorithm"] = state ? state.keyAlgorithm : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["pgpKey"] = state ? state.pgpKey : undefined;
            inputs["privateKey"] = state ? state.privateKey : undefined;
            inputs["privateKeyEncrypted"] = state ? state.privateKeyEncrypted : undefined;
            inputs["privateKeyFingerprint"] = state ? state.privateKeyFingerprint : undefined;
            inputs["privateKeyType"] = state ? state.privateKeyType : undefined;
            inputs["publicKey"] = state ? state.publicKey : undefined;
            inputs["publicKeyType"] = state ? state.publicKeyType : undefined;
            inputs["serviceAccountId"] = state ? state.serviceAccountId : undefined;
            inputs["validAfter"] = state ? state.validAfter : undefined;
            inputs["validBefore"] = state ? state.validBefore : undefined;
        } else {
            const args = argsOrState as KeyArgs | undefined;
            if (!args || args.serviceAccountId === undefined) {
                throw new Error("Missing required property 'serviceAccountId'");
            }
            inputs["keyAlgorithm"] = args ? args.keyAlgorithm : undefined;
            inputs["pgpKey"] = args ? args.pgpKey : undefined;
            inputs["privateKeyType"] = args ? args.privateKeyType : undefined;
            inputs["publicKeyType"] = args ? args.publicKeyType : undefined;
            inputs["serviceAccountId"] = args ? args.serviceAccountId : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["privateKey"] = undefined /*out*/;
            inputs["privateKeyEncrypted"] = undefined /*out*/;
            inputs["privateKeyFingerprint"] = undefined /*out*/;
            inputs["publicKey"] = undefined /*out*/;
            inputs["validAfter"] = undefined /*out*/;
            inputs["validBefore"] = undefined /*out*/;
        }
        super("gcp:serviceAccount/key:Key", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Key resources.
 */
export interface KeyState {
    /**
     * The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
     * Valid values are listed at
     * [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
     * (only used on create)
     */
    readonly keyAlgorithm?: pulumi.Input<string>;
    /**
     * The name used for this key pair
     */
    readonly name?: pulumi.Input<string>;
    /**
     * An optional PGP key to encrypt the resulting private
     * key material. Only used when creating or importing a new key pair. May either be
     * a base64-encoded public key or a `keybase:keybaseusername` string for looking up
     * in Vault.
     */
    readonly pgpKey?: pulumi.Input<string>;
    /**
     * The private key in JSON format, base64 encoded. This is what you normally get as a file when creating
     * service account keys through the CLI or web console. This is only populated when creating a new key, and when no
     * `pgp_key` is provided.
     */
    readonly privateKey?: pulumi.Input<string>;
    /**
     * The private key material, base 64 encoded and
     * encrypted with the given `pgp_key`. This is only populated when creating a new
     * key and `pgp_key` is supplied
     */
    readonly privateKeyEncrypted?: pulumi.Input<string>;
    /**
     * The MD5 public key fingerprint for the encrypted
     * private key. This is only populated when creating a new key and `pgp_key` is supplied
     */
    readonly privateKeyFingerprint?: pulumi.Input<string>;
    /**
     * The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
     */
    readonly privateKeyType?: pulumi.Input<string>;
    /**
     * The public key, base64 encoded
     */
    readonly publicKey?: pulumi.Input<string>;
    /**
     * The output format of the public key requested. X509_PEM is the default output format.
     */
    readonly publicKeyType?: pulumi.Input<string>;
    /**
     * The Service account id of the Key Pair. This can be a string in the format
     * `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`, where `{ACCOUNT}` is the email address or
     * unique id of the service account. If the `{ACCOUNT}` syntax is used, the project will be inferred from the account.
     */
    readonly serviceAccountId?: pulumi.Input<string>;
    /**
     * The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
     */
    readonly validAfter?: pulumi.Input<string>;
    /**
     * The key can be used before this timestamp.
     * A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
     */
    readonly validBefore?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Key resource.
 */
export interface KeyArgs {
    /**
     * The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
     * Valid values are listed at
     * [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
     * (only used on create)
     */
    readonly keyAlgorithm?: pulumi.Input<string>;
    /**
     * An optional PGP key to encrypt the resulting private
     * key material. Only used when creating or importing a new key pair. May either be
     * a base64-encoded public key or a `keybase:keybaseusername` string for looking up
     * in Vault.
     */
    readonly pgpKey?: pulumi.Input<string>;
    /**
     * The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
     */
    readonly privateKeyType?: pulumi.Input<string>;
    /**
     * The output format of the public key requested. X509_PEM is the default output format.
     */
    readonly publicKeyType?: pulumi.Input<string>;
    /**
     * The Service account id of the Key Pair. This can be a string in the format
     * `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`, where `{ACCOUNT}` is the email address or
     * unique id of the service account. If the `{ACCOUNT}` syntax is used, the project will be inferred from the account.
     */
    readonly serviceAccountId: pulumi.Input<string>;
}
