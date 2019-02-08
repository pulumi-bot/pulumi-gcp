// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Get a VPN gateway within GCE from its name.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const my_vpn_gateway = pulumi.output(gcp.compute.getVPNGateway({
 *     name: "vpn-gateway-us-east1",
 * }));
 * ```
 */
export function getVPNGateway(args: GetVPNGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetVPNGatewayResult> {
    return pulumi.runtime.invoke("gcp:compute/getVPNGateway:getVPNGateway", {
        "name": args.name,
        "project": args.project,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getVPNGateway.
 */
export interface GetVPNGatewayArgs {
    /**
     * The name of the VPN gateway.
     */
    readonly name: string;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: string;
    /**
     * The region in which the resource belongs. If it
     * is not provided, the project region is used.
     */
    readonly region?: string;
}

/**
 * A collection of values returned by getVPNGateway.
 */
export interface GetVPNGatewayResult {
    /**
     * Description of this VPN gateway.
     */
    readonly description: string;
    /**
     * The network of this VPN gateway.
     */
    readonly network: string;
    readonly project: string;
    /**
     * Region of this VPN gateway.
     */
    readonly region: string;
    /**
     * The URI of the resource.
     */
    readonly selfLink: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
