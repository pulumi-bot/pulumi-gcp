// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about a Google Cloud Organization.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/organization.html.markdown.
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
    readonly domain?: string;
    /**
     * The name of the Organization in the form `{organization_id}` or `organizations/{organization_id}`.
     */
    readonly organization?: string;
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
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
