// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
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
 * const all = pulumi.output(gcp.composer.getImageVersions({}));
 * const test = new gcp.composer.Environment("test", {
 *     config: {
 *         softwareConfig: {
 *             imageVersion: all.imageVersions[0].imageVersionId,
 *         },
 *     },
 *     region: "us-central1",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/composer_image_versions.html.markdown.
 */
export function getImageVersions(args?: GetImageVersionsArgs, opts?: pulumi.InvokeOptions): Promise<GetImageVersionsResult> & GetImageVersionsResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetImageVersionsResult> = pulumi.runtime.invoke("gcp:composer/getImageVersions:getImageVersions", {
        "project": args.project,
        "region": args.region,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getImageVersions.
 */
export interface GetImageVersionsArgs {
    /**
     * The ID of the project to list versions in.
     * If it is not provided, the provider project is used.
     */
    readonly project?: string;
    /**
     * The location to list versions in.
     * If it is not provider, the provider region is used.
     */
    readonly region?: string;
}

/**
 * A collection of values returned by getImageVersions.
 */
export interface GetImageVersionsResult {
    /**
     * A list of composer image versions available in the given project and location. Each `imageVersion` contains:
     */
    readonly imageVersions: { imageVersionId: string, supportedPythonVersions: string[] }[];
    readonly project: string;
    readonly region: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
