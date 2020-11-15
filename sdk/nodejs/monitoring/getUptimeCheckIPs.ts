// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Returns the list of IP addresses that checkers run from. For more information see
 * the [official documentation](https://cloud.google.com/monitoring/uptime-checks#get-ips).
 */
export function getUptimeCheckIPs(opts?: pulumi.InvokeOptions): Promise<GetUptimeCheckIPsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:monitoring/getUptimeCheckIPs:getUptimeCheckIPs", {
    }, opts);
}

/**
 * A collection of values returned by getUptimeCheckIPs.
 */
export interface GetUptimeCheckIPsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of uptime check IPs used by Stackdriver Monitoring. Each `uptimeCheckIp` contains:
     */
    readonly uptimeCheckIps: outputs.monitoring.GetUptimeCheckIPsUptimeCheckIp[];
}
