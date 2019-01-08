// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getProjectServiceAccount(args?: GetProjectServiceAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectServiceAccountResult> {
    args = args || {};
    return pulumi.runtime.invoke("gcp:storage/getProjectServiceAccount:getProjectServiceAccount", {
        "project": args.project,
        "userProject": args.userProject,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectServiceAccount.
 */
export interface GetProjectServiceAccountArgs {
    readonly project?: string;
    readonly userProject?: string;
}

/**
 * A collection of values returned by getProjectServiceAccount.
 */
export interface GetProjectServiceAccountResult {
    readonly emailAddress: string;
    readonly project: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
