// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Get a HA VPN Gateway within GCE from its name.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const gateway = pulumi.output(gcp.compute.getHcVpnGateway({
 *     name: "foobar",
 * }, { async: true }));
 * ```
 */
export function getHcVpnGateway(args: GetHcVpnGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetHcVpnGatewayResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:compute/getHcVpnGateway:getHcVpnGateway", {
        "name": args.name,
        "project": args.project,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getHcVpnGateway.
 */
export interface GetHcVpnGatewayArgs {
    /**
     * The name of the forwarding rule.
     */
    name: string;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: string;
    /**
     * The region in which the resource belongs. If it
     * is not provided, the project region is used.
     */
    region?: string;
}

/**
 * A collection of values returned by getHcVpnGateway.
 */
export interface GetHcVpnGatewayResult {
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly network: string;
    readonly project?: string;
    readonly region?: string;
    readonly selfLink: string;
    readonly vpnInterfaces: outputs.compute.GetHcVpnGatewayVpnInterface[];
}
