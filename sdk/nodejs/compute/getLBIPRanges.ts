// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access IP ranges in your firewall rules.
 * 
 * https://cloud.google.com/compute/docs/load-balancing/health-checks#health_check_source_ips_and_firewall_rules
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_compute_lb_ip_ranges.html.markdown.
 */
export function getLBIPRanges(opts?: pulumi.InvokeOptions): Promise<GetLBIPRangesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:compute/getLBIPRanges:getLBIPRanges", {
    }, opts);
}

/**
 * A collection of values returned by getLBIPRanges.
 */
export interface GetLBIPRangesResult {
    /**
     * The IP ranges used for health checks when **HTTP(S), SSL proxy, TCP proxy, and Internal load balancing** is used
     */
    readonly httpSslTcpInternals: string[];
    /**
     * The IP ranges used for health checks when **Network load balancing** is used
     */
    readonly networks: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
