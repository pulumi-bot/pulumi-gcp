// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Get Cloud Run locations available for a project.
 *
 * To get more information about Cloud Run, see:
 *
 * * [API documentation](https://cloud.google.com/run/docs/reference/rest/v1/projects.locations)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/run/docs/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const available = pulumi.output(gcp.cloudrun.getLocations({ async: true }));
 * ```
 */
export function getLocations(args?: GetLocationsArgs, opts?: pulumi.InvokeOptions): Promise<GetLocationsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:cloudrun/getLocations:getLocations", {
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getLocations.
 */
export interface GetLocationsArgs {
    /**
     * The project to list versions for. If it
     * is not provided, the provider project is used.
     */
    project?: string;
}

/**
 * A collection of values returned by getLocations.
 */
export interface GetLocationsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of Cloud Run locations available for the given project.
     */
    readonly locations: string[];
    readonly project: string;
}
