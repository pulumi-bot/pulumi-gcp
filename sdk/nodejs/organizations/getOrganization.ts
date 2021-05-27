// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Get information about a Google Cloud Organization. Note that you must have the `roles/resourcemanager.organizationViewer` role (or equivalent permissions) at the organization level to use this datasource.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const org = gcp.organizations.getOrganization({
 *     domain: "example.com",
 * });
 * const sales = new gcp.organizations.Folder("sales", {
 *     displayName: "Sales",
 *     parent: org.then(org => org.name),
 * });
 * ```
 */
export function getOrganization(args?: GetOrganizationArgs, opts?: pulumi.InvokeOptions): Promise<GetOrganizationResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:organizations/getOrganization:getOrganization", {
        "domain": args.domain,
        "organization": args.organization,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrganization.
 */
export interface GetOrganizationArgs {
    /**
     * The domain name of the Organization.
     */
    domain?: string;
    /**
     * The Organization's numeric ID, including an optional `organizations/` prefix.
     */
    organization?: string;
}

/**
 * A collection of values returned by getOrganization.
 */
export interface GetOrganizationResult {
    /**
     * Timestamp when the Organization was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
     */
    readonly createTime: string;
    /**
     * The Google for Work customer ID of the Organization.
     */
    readonly directoryCustomerId: string;
    readonly domain: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Organization's current lifecycle state.
     */
    readonly lifecycleState: string;
    /**
     * The resource name of the Organization in the form `organizations/{organization_id}`.
     */
    readonly name: string;
    /**
     * The Organization ID.
     */
    readonly orgId: string;
    readonly organization?: string;
}

export function getOrganizationApply(args?: GetOrganizationApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrganizationResult> {
    return pulumi.output(args).apply(a => getOrganization(a, opts))
}

/**
 * A collection of arguments for invoking getOrganization.
 */
export interface GetOrganizationApplyArgs {
    /**
     * The domain name of the Organization.
     */
    domain?: pulumi.Input<string>;
    /**
     * The Organization's numeric ID, including an optional `organizations/` prefix.
     */
    organization?: pulumi.Input<string>;
}
