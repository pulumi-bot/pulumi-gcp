// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A `CryptoKey` represents a logical key that can be used for cryptographic operations.
 *
 * > **Note:** CryptoKeys cannot be deleted from Google Cloud Platform.
 * Destroying a provider-managed CryptoKey will remove it from state
 * and delete all CryptoKeyVersions, rendering the key unusable, but *will
 * not delete the resource from the project.* When the provider destroys these keys,
 * any data previously encrypted with these keys will be irrecoverable.
 * For this reason, it is strongly recommended that you add lifecycle hooks
 * to the resource to prevent accidental destruction.
 *
 * To get more information about CryptoKey, see:
 *
 * * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys)
 * * How-to Guides
 *     * [Creating a key](https://cloud.google.com/kms/docs/creating-keys#create_a_key)
 *
 * ## Example Usage
 * ### Kms Crypto Key Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const keyring = new gcp.kms.KeyRing("keyring", {location: "global"});
 * const example_key = new gcp.kms.CryptoKey("example-key", {
 *     keyRing: keyring.id,
 *     rotationPeriod: "100000s",
 * });
 * ```
 * ### Kms Crypto Key Asymmetric Sign
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const keyring = new gcp.kms.KeyRing("keyring", {location: "global"});
 * const example_asymmetric_sign_key = new gcp.kms.CryptoKey("example-asymmetric-sign-key", {
 *     keyRing: keyring.id,
 *     purpose: "ASYMMETRIC_SIGN",
 *     versionTemplate: {
 *         algorithm: "EC_SIGN_P384_SHA384",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * CryptoKey can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:kms/cryptoKey:CryptoKey default {{key_ring}}/cryptoKeys/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:kms/cryptoKey:CryptoKey default {{key_ring}}/{{name}}
 * ```
 */
export class CryptoKey extends pulumi.CustomResource {
    /**
     * Get an existing CryptoKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CryptoKeyState, opts?: pulumi.CustomResourceOptions): CryptoKey {
        return new CryptoKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:kms/cryptoKey:CryptoKey';

    /**
     * Returns true if the given object is an instance of CryptoKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CryptoKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CryptoKey.__pulumiType;
    }

    /**
     * The KeyRing that this key belongs to.
     * Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
     */
    public readonly keyRing!: pulumi.Output<string>;
    /**
     * Labels with user-defined metadata to apply to this resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The resource name for the CryptoKey.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The immutable purpose of this CryptoKey. See the
     * [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
     * for possible inputs.
     * Default value is `ENCRYPT_DECRYPT`.
     * Possible values are `ENCRYPT_DECRYPT`, `ASYMMETRIC_SIGN`, and `ASYMMETRIC_DECRYPT`.
     */
    public readonly purpose!: pulumi.Output<string | undefined>;
    /**
     * Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
     * The first rotation will take place after the specified period. The rotation period has
     * the format of a decimal number with up to 9 fractional digits, followed by the
     * letter `s` (seconds). It must be greater than a day (ie, 86400).
     */
    public readonly rotationPeriod!: pulumi.Output<string | undefined>;
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
     * You must use the `gcp.kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
     */
    public readonly skipInitialVersionCreation!: pulumi.Output<boolean | undefined>;
    /**
     * A template describing settings for new crypto key versions.
     * Structure is documented below.
     */
    public readonly versionTemplate!: pulumi.Output<outputs.kms.CryptoKeyVersionTemplate>;

    /**
     * Create a CryptoKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CryptoKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CryptoKeyArgs | CryptoKeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CryptoKeyState | undefined;
            inputs["keyRing"] = state ? state.keyRing : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["purpose"] = state ? state.purpose : undefined;
            inputs["rotationPeriod"] = state ? state.rotationPeriod : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["skipInitialVersionCreation"] = state ? state.skipInitialVersionCreation : undefined;
            inputs["versionTemplate"] = state ? state.versionTemplate : undefined;
        } else {
            const args = argsOrState as CryptoKeyArgs | undefined;
            if ((!args || args.keyRing === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'keyRing'");
            }
            inputs["keyRing"] = args ? args.keyRing : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["purpose"] = args ? args.purpose : undefined;
            inputs["rotationPeriod"] = args ? args.rotationPeriod : undefined;
            inputs["skipInitialVersionCreation"] = args ? args.skipInitialVersionCreation : undefined;
            inputs["versionTemplate"] = args ? args.versionTemplate : undefined;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CryptoKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CryptoKey resources.
 */
export interface CryptoKeyState {
    /**
     * The KeyRing that this key belongs to.
     * Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
     */
    readonly keyRing?: pulumi.Input<string>;
    /**
     * Labels with user-defined metadata to apply to this resource.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name for the CryptoKey.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The immutable purpose of this CryptoKey. See the
     * [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
     * for possible inputs.
     * Default value is `ENCRYPT_DECRYPT`.
     * Possible values are `ENCRYPT_DECRYPT`, `ASYMMETRIC_SIGN`, and `ASYMMETRIC_DECRYPT`.
     */
    readonly purpose?: pulumi.Input<string>;
    /**
     * Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
     * The first rotation will take place after the specified period. The rotation period has
     * the format of a decimal number with up to 9 fractional digits, followed by the
     * letter `s` (seconds). It must be greater than a day (ie, 86400).
     */
    readonly rotationPeriod?: pulumi.Input<string>;
    readonly selfLink?: pulumi.Input<string>;
    /**
     * If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
     * You must use the `gcp.kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
     */
    readonly skipInitialVersionCreation?: pulumi.Input<boolean>;
    /**
     * A template describing settings for new crypto key versions.
     * Structure is documented below.
     */
    readonly versionTemplate?: pulumi.Input<inputs.kms.CryptoKeyVersionTemplate>;
}

/**
 * The set of arguments for constructing a CryptoKey resource.
 */
export interface CryptoKeyArgs {
    /**
     * The KeyRing that this key belongs to.
     * Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
     */
    readonly keyRing: pulumi.Input<string>;
    /**
     * Labels with user-defined metadata to apply to this resource.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name for the CryptoKey.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The immutable purpose of this CryptoKey. See the
     * [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
     * for possible inputs.
     * Default value is `ENCRYPT_DECRYPT`.
     * Possible values are `ENCRYPT_DECRYPT`, `ASYMMETRIC_SIGN`, and `ASYMMETRIC_DECRYPT`.
     */
    readonly purpose?: pulumi.Input<string>;
    /**
     * Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
     * The first rotation will take place after the specified period. The rotation period has
     * the format of a decimal number with up to 9 fractional digits, followed by the
     * letter `s` (seconds). It must be greater than a day (ie, 86400).
     */
    readonly rotationPeriod?: pulumi.Input<string>;
    /**
     * If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
     * You must use the `gcp.kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
     */
    readonly skipInitialVersionCreation?: pulumi.Input<boolean>;
    /**
     * A template describing settings for new crypto key versions.
     * Structure is documented below.
     */
    readonly versionTemplate?: pulumi.Input<inputs.kms.CryptoKeyVersionTemplate>;
}
