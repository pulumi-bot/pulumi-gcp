// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Get a subnetwork within GCE from its name and region.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const my_subnetwork = pulumi.output(gcp.compute.getSubnetwork({
 *     name: "default-us-east1",
 *     region: "us-east1",
 * }));
 * ```
 */
export function getSubnetwork(args?: GetSubnetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetSubnetworkResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:compute/getSubnetwork:getSubnetwork", {
        "name": args.name,
        "project": args.project,
        "region": args.region,
        "selfLink": args.selfLink,
    }, opts);
}

/**
 * A collection of arguments for invoking getSubnetwork.
 */
export interface GetSubnetworkArgs {
    /**
     * The name of the subnetwork. One of `name` or `self_link`
     * must be specified.
     */
    readonly name?: string;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: string;
    /**
     * The region this subnetwork has been created in. If
     * unspecified, this defaults to the region configured in the provider.
     */
    readonly region?: string;
    /**
     * The self link of the subnetwork. If `self_link` is
     * specified, `name`, `project`, and `region` are ignored.
     */
    readonly selfLink?: string;
}

/**
 * A collection of values returned by getSubnetwork.
 */
export interface GetSubnetworkResult {
    /**
     * Description of this subnetwork.
     */
    readonly description: string;
    /**
     * The IP address of the gateway.
     */
    readonly gatewayAddress: string;
    /**
     * The range of IP addresses belonging to this subnetwork
     * secondary range.
     */
    readonly ipCidrRange: string;
    readonly name?: string;
    /**
     * The network name or resource link to the parent
     * network of this subnetwork.
     */
    readonly network: string;
    /**
     * Whether the VMs in this subnet
     * can access Google services without assigned external IP
     * addresses.
     */
    readonly privateIpGoogleAccess: boolean;
    readonly project: string;
    readonly region: string;
    /**
     * An array of configurations for secondary IP ranges for
     * VM instances contained in this subnetwork. Structure is documented below.
     */
    readonly secondaryIpRanges: { ipCidrRange: string, rangeName: string }[];
    readonly selfLink: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
