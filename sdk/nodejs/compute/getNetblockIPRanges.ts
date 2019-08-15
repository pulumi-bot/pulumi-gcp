// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../types/input";
import * as outputApi from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the IP ranges from the sender policy framework (SPF) record of \_cloud-netblocks.googleusercontent
 * 
 * https://cloud.google.com/compute/docs/faq#where_can_i_find_product_name_short_ip_ranges
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const netblock = gcp.compute.getNetblockIPRanges({});
 * 
 * export const cidrBlocks = netblock.cidrBlocks;
 * export const cidrBlocksIpv4 = netblock.cidrBlocksIpv4s;
 * export const cidrBlocksIpv6 = netblock.cidrBlocksIpv6s;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/netblock_ip_ranges.html.markdown.
 */
export function getNetblockIPRanges(opts?: pulumi.InvokeOptions): Promise<GetNetblockIPRangesResult> & GetNetblockIPRangesResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetNetblockIPRangesResult> = pulumi.runtime.invoke("gcp:compute/getNetblockIPRanges:getNetblockIPRanges", {
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of values returned by getNetblockIPRanges.
 */
export interface GetNetblockIPRangesResult {
    /**
     * Retrieve list of all CIDR blocks.
     */
    readonly cidrBlocks: string[];
    /**
     * Retrieve list of the IP4 CIDR blocks
     */
    readonly cidrBlocksIpv4s: string[];
    /**
     * Retrieve list of the IP6 CIDR blocks.
     */
    readonly cidrBlocksIpv6s: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
