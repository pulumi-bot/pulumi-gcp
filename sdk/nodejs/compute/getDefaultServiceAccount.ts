// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve default service account for this project
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/google_compute_default_service_account.html.markdown.
 */
export function getDefaultServiceAccount(args?: GetDefaultServiceAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetDefaultServiceAccountResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:compute/getDefaultServiceAccount:getDefaultServiceAccount", {
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getDefaultServiceAccount.
 */
export interface GetDefaultServiceAccountArgs {
    /**
     * The project ID. If it is not provided, the provider project is used.
     */
    readonly project?: string;
}

/**
 * A collection of values returned by getDefaultServiceAccount.
 */
export interface GetDefaultServiceAccountResult {
    /**
     * The display name for the service account.
     */
    readonly displayName: string;
    /**
     * Email address of the default service account used by VMs running in this project
     */
    readonly email: string;
    /**
     * The fully-qualified name of the service account.
     */
    readonly name: string;
    readonly project: string;
    /**
     * The unique id of the service account.
     */
    readonly uniqueId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
