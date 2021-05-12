// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get list of the Cloud Identity Groups under a customer or namespace.
 *
 * https://cloud.google.com/identity/docs/concepts/overview#groups
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const groups = pulumi.output(gcp.cloudidentity.getGroups({
 *     parent: "customers/A01b123xz",
 * }, { async: true }));
 * ```
 */
export function getGroups(args: GetGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:cloudidentity/getGroups:getGroups", {
        "parent": args.parent,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroups.
 */
export interface GetGroupsArgs {
    /**
     * The parent resource under which to list all Groups. Must be of the form identitysources/{identity_source_id} for external- identity-mapped groups or customers/{customer_id} for Google Groups.
     */
    parent: string;
}

/**
 * A collection of values returned by getGroups.
 */
export interface GetGroupsResult {
    /**
     * The list of groups under the provided customer or namespace. Structure is documented below.
     */
    readonly groups: outputs.cloudidentity.GetGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly parent: string;
}
