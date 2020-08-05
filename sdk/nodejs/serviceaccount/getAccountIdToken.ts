// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

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
    readonly delegates?: string[];
    /**
     * Include the verified email in the claim. Used only when using impersonation mode.
     */
    readonly includeEmail?: boolean;
    /**
     * The audience claim for the `idToken`.
     */
    readonly targetAudience: string;
    /**
     * The email of the service account being impersonated.  Used only when using impersonation mode.
     */
    readonly targetServiceAccount?: string;
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
