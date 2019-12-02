// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides access to a Google Cloud Platform KMS CryptoKey. For more information see
 * [the official documentation](https://cloud.google.com/kms/docs/object-hierarchy#key)
 * and
 * [API](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys).
 * 
 * A CryptoKey is an interface to key material which can be used to encrypt and decrypt data. A CryptoKey belongs to a
 * Google Cloud KMS KeyRing.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const myKeyRing = gcp.kms.getKMSKeyRing({
 *     location: "us-central1",
 *     name: "my-key-ring",
 * });
 * const myCryptoKey = gcp.kms.getKMSCryptoKey({
 *     keyRing: myKeyRing.selfLink,
 *     name: "my-crypto-key",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/kms_crypto_key.html.markdown.
 */
export function getKMSCryptoKey(args: GetKMSCryptoKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetKMSCryptoKeyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:kms/getKMSCryptoKey:getKMSCryptoKey", {
        "keyRing": args.keyRing,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getKMSCryptoKey.
 */
export interface GetKMSCryptoKeyArgs {
    /**
     * The `selfLink` of the Google Cloud Platform KeyRing to which the key belongs.
     */
    readonly keyRing: string;
    /**
     * The CryptoKey's name.
     * A CryptoKey’s name belonging to the specified Google Cloud Platform KeyRing and match the regular expression `[a-zA-Z0-9_-]{1,63}`
     */
    readonly name: string;
}

/**
 * A collection of values returned by getKMSCryptoKey.
 */
export interface GetKMSCryptoKeyResult {
    readonly keyRing: string;
    readonly labels: {[key: string]: string};
    readonly name: string;
    /**
     * Defines the cryptographic capabilities of the key.
     */
    readonly purpose: string;
    /**
     * Every time this period passes, generate a new CryptoKeyVersion and set it as
     * the primary. The first rotation will take place after the specified period. The rotation period has the format
     * of a decimal number with up to 9 fractional digits, followed by the letter s (seconds).
     */
    readonly rotationPeriod: string;
    /**
     * The self link of the created CryptoKey. Its format is `projects/{projectId}/locations/{location}/keyRings/{keyRingName}/cryptoKeys/{cryptoKeyName}`.
     */
    readonly selfLink: string;
    readonly versionTemplates: outputs.kms.GetKMSCryptoKeyVersionTemplate[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
