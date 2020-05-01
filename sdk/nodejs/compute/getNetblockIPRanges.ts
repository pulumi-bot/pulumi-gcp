// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the IP addresses from different special IP ranges on Google Cloud Platform.
 * 
 * ## Example Usage - Cloud Ranges
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const netblock = gcp.compute.getNetblockIPRanges({});
 * export const cidrBlocks = netblock.then(netblock => netblock.cidrBlocks);
 * export const cidrBlocksIpv4 = netblock.then(netblock => netblock.cidrBlocksIpv4s);
 * export const cidrBlocksIpv6 = netblock.then(netblock => netblock.cidrBlocksIpv6s);
 * ```
 * 
 * ## Example Usage - Allow Health Checks
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const legacy-hcs = gcp.compute.getNetblockIPRanges({
 *     rangeType: "legacy-health-checkers",
 * });
 * const default = new gcp.compute.Network("default", {});
 * const allow-hcs = new gcp.compute.Firewall("allow-hcs", {
 *     network: default.name,
 *     allow: [{
 *         protocol: "tcp",
 *         ports: ["80"],
 *     }],
 *     sourceRanges: legacy-hcs.then(legacy_hcs => legacy_hcs.cidrBlocksIpv4s),
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_google_netblock_ip_ranges.html.markdown.
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
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
