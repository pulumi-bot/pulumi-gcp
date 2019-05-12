// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Get service account public key. For more information, see [the official documentation](https://cloud.google.com/iam/docs/creating-managing-service-account-keys) and [API](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys/get).
 * 
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const myaccount = new gcp.serviceAccount.Account("myaccount", {
 *     accountId: "dev-foo-account",
 * });
 * const mykeyKey = new gcp.serviceAccount.Key("mykey", {
 *     serviceAccountId: myaccount.name,
 * });
 * const mykeyAccountKey = mykeyKey.name.apply(name => gcp.serviceAccount.getAccountKey({
 *     name: name,
 *     publicKeyType: "TYPE_X509_PEM_FILE",
 * }));
 * ```
 */
export function getAccountKey(args: GetAccountKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountKeyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:serviceAccount/getAccountKey:getAccountKey", {
        "name": args.name,
        "project": args.project,
        "publicKeyType": args.publicKeyType,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccountKey.
 */
export interface GetAccountKeyArgs {
    /**
     * The name of the service account key. This must have format
     * `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}/keys/{KEYID}`, where `{ACCOUNT}`
     * is the email address or unique id of the service account.
     */
    readonly name: string;
    /**
     * The ID of the project that the service account will be created in.
     * Defaults to the provider project configuration.
     */
    readonly project?: string;
    /**
     * The output format of the public key requested. X509_PEM is the default output format.
     */
    readonly publicKeyType?: string;
}

/**
 * A collection of values returned by getAccountKey.
 */
export interface GetAccountKeyResult {
    readonly keyAlgorithm: string;
    readonly name: string;
    readonly project?: string;
    /**
     * The public key, base64 encoded
     */
    readonly publicKey: string;
    readonly publicKeyType?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
