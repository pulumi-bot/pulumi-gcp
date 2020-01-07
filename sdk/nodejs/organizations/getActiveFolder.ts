// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get an active folder within GCP by `displayName` and `parent`.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const department1 = gcp.organizations.getActiveFolder({
 *     displayName: "Department 1",
 *     parent: "organizations/1234567",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/active_folder.html.markdown.
 */
export function getActiveFolder(args: GetActiveFolderArgs, opts?: pulumi.InvokeOptions): Promise<GetActiveFolderResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:organizations/getActiveFolder:getActiveFolder", {
        "displayName": args.displayName,
        "parent": args.parent,
    }, opts);
}

/**
 * A collection of arguments for invoking getActiveFolder.
 */
export interface GetActiveFolderArgs {
    /**
     * The folder's display name.
     */
    readonly displayName: string;
    /**
     * The resource name of the parent Folder or Organization.
     */
    readonly parent: string;
}

/**
 * A collection of values returned by getActiveFolder.
 */
export interface GetActiveFolderResult {
    readonly displayName: string;
    /**
     * The resource name of the Folder. This uniquely identifies the folder.
     */
    readonly name: string;
    readonly parent: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
