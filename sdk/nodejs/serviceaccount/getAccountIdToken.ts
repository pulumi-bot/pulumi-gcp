// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides a Google OpenID Connect (`oidc`) `idToken`.  Tokens issued from this data source are typically used to call external services that accept OIDC tokens for authentication (e.g. [Google Cloud Run](https://cloud.google.com/run/docs/authenticating/service-to-service)).
 *
 * For more information see
 * [OpenID Connect](https://openid.net/specs/openid-connect-core-1_0.html#IDToken).
 *
 * ## Example Usage
 *
 * ### ServiceAccount JSON Credential File.
 *   `gcp.serviceAccount.getAccountIdToken` will use the configured provider credentials
 *
 * ### Service Account Impersonation.
 *   `gcp.serviceAccount.getAccountAccessToken` will use background impersonated credentials provided by `gcp.serviceAccount.getAccountAccessToken`.
 *
 *   Note: to use the following, you must grant `targetServiceAccount` the
 *   `roles/iam.serviceAccountTokenCreator` role on itself.
 */
export function getAccountIdToken(args: GetAccountIdTokenArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountIdTokenResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:serviceAccount/getAccountIdToken:getAccountIdToken", {
        "delegates": args.delegates,
        "includeEmail": args.includeEmail,
        "targetAudience": args.targetAudience,
        "targetServiceAccount": args.targetServiceAccount,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccountIdToken.
 */
export interface GetAccountIdTokenArgs {
    /**
     * Delegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.   Used only when using impersonation mode.
     */
    delegates?: string[];
    /**
     * Include the verified email in the claim. Used only when using impersonation mode.
     */
    includeEmail?: boolean;
    /**
     * The audience claim for the `idToken`.
     */
    targetAudience: string;
    /**
     * The email of the service account being impersonated.  Used only when using impersonation mode.
     */
    targetServiceAccount?: string;
}

/**
 * A collection of values returned by getAccountIdToken.
 */
export interface GetAccountIdTokenResult {
    readonly delegates?: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The `idToken` representing the new generated identity.
     */
    readonly idToken: string;
    readonly includeEmail?: boolean;
    readonly targetAudience: string;
    readonly targetServiceAccount?: string;
}

export function getAccountIdTokenApply(args: GetAccountIdTokenApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccountIdTokenResult> {
    return pulumi.output(args).apply(a => getAccountIdToken(a, opts))
}

/**
 * A collection of arguments for invoking getAccountIdToken.
 */
export interface GetAccountIdTokenApplyArgs {
    /**
     * Delegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.   Used only when using impersonation mode.
     */
    delegates?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Include the verified email in the claim. Used only when using impersonation mode.
     */
    includeEmail?: pulumi.Input<boolean>;
    /**
     * The audience claim for the `idToken`.
     */
    targetAudience: pulumi.Input<string>;
    /**
     * The email of the service account being impersonated.  Used only when using impersonation mode.
     */
    targetServiceAccount?: pulumi.Input<string>;
}
