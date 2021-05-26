// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Get the IP address from a static address. For more information see
 * the official [API](https://cloud.google.com/compute/docs/reference/latest/addresses/get) documentation.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const myAddress = gcp.compute.getAddress({
 *     name: "foobar",
 * });
 * const prod = new gcp.dns.ManagedZone("prod", {dnsName: "prod.mydomain.com."});
 * const frontend = new gcp.dns.RecordSet("frontend", {
 *     name: pulumi.interpolate`frontend.${prod.dnsName}`,
 *     type: "A",
 *     ttl: 300,
 *     managedZone: prod.name,
 *     rrdatas: [myAddress.then(myAddress => myAddress.address)],
 * });
 * ```
 */
export function getAddress(args: GetAddressArgs, opts?: pulumi.InvokeOptions): Promise<GetAddressResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:compute/getAddress:getAddress", {
        "name": args.name,
        "project": args.project,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getAddress.
 */
export interface GetAddressArgs {
    /**
     * A unique name for the resource, required by GCE.
     */
    name: string;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: string;
    /**
     * The Region in which the created address reside.
     * If it is not provided, the provider region is used.
     */
    region?: string;
}

/**
 * A collection of values returned by getAddress.
 */
export interface GetAddressResult {
    /**
     * The IP of the created resource.
     */
    readonly address: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly project: string;
    readonly region: string;
    /**
     * The URI of the created resource.
     */
    readonly selfLink: string;
    /**
     * Indicates if the address is used. Possible values are: RESERVED or IN_USE.
     */
    readonly status: string;
}
