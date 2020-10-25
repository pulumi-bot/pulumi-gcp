// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get info about a Google Cloud Redis instance.
 */
export function getInstance(args: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:redis/getInstance:getInstance", {
        "name": args.name,
        "project": args.project,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceArgs {
    /**
     * The name of a Redis instance.
     */
    readonly name: string;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: string;
    /**
     * The region in which the resource belongs. If it
     * is not provided, the provider region is used.
     */
    readonly region?: string;
}

/**
 * A collection of values returned by getInstance.
 */
export interface GetInstanceResult {
    readonly alternativeLocationId: string;
    readonly authEnabled: boolean;
    readonly authorizedNetwork: string;
    readonly connectMode: string;
    readonly createTime: string;
    readonly currentLocationId: string;
    readonly displayName: string;
    readonly host: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly labels: {[key: string]: string};
    readonly locationId: string;
    readonly memorySizeGb: number;
    readonly name: string;
    readonly persistenceIamIdentity: string;
    readonly port: number;
    readonly project?: string;
    readonly redisConfigs: {[key: string]: string};
    readonly redisVersion: string;
    readonly region?: string;
    readonly reservedIpRange: string;
    readonly tier: string;
}
