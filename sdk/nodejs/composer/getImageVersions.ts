// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides access to available Cloud Composer versions in a region for a given project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const all = gcp.composer.getImageVersions({});
 * const test = new gcp.composer.Environment("test", {
 *     region: "us-central1",
 *     config: {
 *         softwareConfig: {
 *             imageVersion: all.then(all => all.imageVersions[0].imageVersionId),
 *         },
 *     },
 * });
 * ```
 */
export function getImageVersions(args?: GetImageVersionsArgs, opts?: pulumi.InvokeOptions): Promise<GetImageVersionsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:composer/getImageVersions:getImageVersions", {
        "project": args.project,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getImageVersions.
 */
export interface GetImageVersionsArgs {
    /**
     * The ID of the project to list versions in.
     * If it is not provided, the provider project is used.
     */
    project?: string;
    /**
     * The location to list versions in.
     * If it is not provider, the provider region is used.
     */
    region?: string;
}

/**
 * A collection of values returned by getImageVersions.
 */
export interface GetImageVersionsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of composer image versions available in the given project and location. Each `imageVersion` contains:
     */
    readonly imageVersions: outputs.composer.GetImageVersionsImageVersion[];
    readonly project: string;
    readonly region: string;
}
