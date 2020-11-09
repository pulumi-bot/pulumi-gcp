// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Get the DNSKEY and DS records of DNSSEC-signed managed zones. For more information see the
 * [official documentation](https://cloud.google.com/dns/docs/dnskeys/)
 * and [API](https://cloud.google.com/dns/docs/reference/v1/dnsKeys).
 */
export function getKeys(args: GetKeysArgs, opts?: pulumi.InvokeOptions): Promise<GetKeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:dns/getKeys:getKeys", {
        "managedZone": args.managedZone,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getKeys.
 */
export interface GetKeysArgs {
    /**
     * The name or id of the Cloud DNS managed zone.
     */
    readonly managedZone: string;
    /**
     * The ID of the project in which the resource belongs. If `project` is not provided, the provider project is used.
     */
    readonly project?: string;
}

/**
 * A collection of values returned by getKeys.
 */
export interface GetKeysResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of Key-signing key (KSK) records. Structure is documented below. Additionally, the DS record is provided:
     */
    readonly keySigningKeys: outputs.dns.GetKeysKeySigningKey[];
    readonly managedZone: string;
    readonly project: string;
    /**
     * A list of Zone-signing key (ZSK) records. Structure is documented below.
     */
    readonly zoneSigningKeys: outputs.dns.GetKeysZoneSigningKey[];
}
