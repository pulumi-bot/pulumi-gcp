// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides access to a zone's attributes within Google Cloud DNS.
 * For more information see
 * [the official documentation](https://cloud.google.com/dns/zones/)
 * and
 * [API](https://cloud.google.com/dns/api/v1/managedZones).
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/dns_managed_zone.html.markdown.
 */
export function getManagedZone(args: GetManagedZoneArgs, opts?: pulumi.InvokeOptions): Promise<GetManagedZoneResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:dns/getManagedZone:getManagedZone", {
        "name": args.name,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getManagedZone.
 */
export interface GetManagedZoneArgs {
    /**
     * A unique name for the resource.
     */
    readonly name: string;
    /**
     * The ID of the project for the Google Cloud DNS zone.
     */
    readonly project?: string;
}

/**
 * A collection of values returned by getManagedZone.
 */
export interface GetManagedZoneResult {
    /**
     * A textual description field.
     */
    readonly description: string;
    /**
     * The fully qualified DNS name of this zone, e.g. `example.io.`.
     */
    readonly dnsName: string;
    readonly name: string;
    /**
     * The list of nameservers that will be authoritative for this
     * domain. Use NS records to redirect from your DNS provider to these names,
     * thus making Google Cloud DNS authoritative for this zone.
     */
    readonly nameServers: string[];
    readonly project?: string;
    /**
     * The zone's visibility: public zones are exposed to the Internet,
     * while private zones are visible only to Virtual Private Cloud resources.
     */
    readonly visibility: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
