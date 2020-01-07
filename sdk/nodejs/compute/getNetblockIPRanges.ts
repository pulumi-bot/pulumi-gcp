// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the IP addresses from different special IP ranges on Google Cloud Platform.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/netblock_ip_ranges.html.markdown.
 */
export function getNetblockIPRanges(args?: GetNetblockIPRangesArgs, opts?: pulumi.InvokeOptions): Promise<GetNetblockIPRangesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:compute/getNetblockIPRanges:getNetblockIPRanges", {
        "rangeType": args.rangeType,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetblockIPRanges.
 */
export interface GetNetblockIPRangesArgs {
    /**
     * The type of range for which to provide results.
     */
    readonly rangeType?: string;
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
     * Retrieve list of the IPv4 CIDR blocks
     */
    readonly cidrBlocksIpv4s: string[];
    /**
     * Retrieve list of the IPv6 CIDR blocks, if available.
     */
    readonly cidrBlocksIpv6s: string[];
    readonly rangeType?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
