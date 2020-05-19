// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about a Google Cloud Folder.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const myFolder1 = gcp.organizations.getFolder({
 *     folder: "folders/12345",
 *     lookupOrganization: true,
 * });
 * const myFolder2 = gcp.organizations.getFolder({
 *     folder: "folders/23456",
 * });
 * export const myFolder1Organization = myFolder1.then(myFolder1 => myFolder1.organization);
 * export const myFolder2Parent = myFolder2.then(myFolder2 => myFolder2.parent);
 * ```
 */
export function getFolder(args: GetFolderArgs, opts?: pulumi.InvokeOptions): Promise<GetFolderResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:organizations/getFolder:getFolder", {
        "folder": args.folder,
        "lookupOrganization": args.lookupOrganization,
    }, opts);
}

/**
 * A collection of arguments for invoking getFolder.
 */
export interface GetFolderArgs {
    /**
     * The name of the Folder in the form `{folder_id}` or `folders/{folder_id}`.
     */
    readonly folder: string;
    /**
     * `true` to find the organization that the folder belongs, `false` to avoid the lookup. It searches up the tree. (defaults to `false`)
     */
    readonly lookupOrganization?: boolean;
}

/**
 * A collection of values returned by getFolder.
 */
export interface GetFolderResult {
    /**
     * Timestamp when the Organization was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
     */
    readonly createTime: string;
    /**
     * The folder's display name.
     */
    readonly displayName: string;
    readonly folder: string;
    /**
     * The Folder's current lifecycle state.
     */
    readonly lifecycleState: string;
    readonly lookupOrganization?: boolean;
    /**
     * The resource name of the Folder in the form `folders/{folder_id}`.
     */
    readonly name: string;
    /**
     * If `lookupOrganization` is enable, the resource name of the Organization that the folder belongs.
     */
    readonly organization: string;
    /**
     * The resource name of the parent Folder or Organization.
     */
    readonly parent: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
