// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get information about a BackendBucket.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const myBackendBucket = gcp.compute.getBackendBucket({
 *     name: "my-backend",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_google_compute_backend_bucket.html.markdown.
 */
export function getBackendBucket(args: GetBackendBucketArgs, opts?: pulumi.InvokeOptions): Promise<GetBackendBucketResult> & GetBackendBucketResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetBackendBucketResult> = pulumi.runtime.invoke("gcp:compute/getBackendBucket:getBackendBucket", {
        "name": args.name,
        "project": args.project,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getBackendBucket.
 */
export interface GetBackendBucketArgs {
    /**
     * Name of the resource.
     */
    readonly name: string;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: string;
}

/**
 * A collection of values returned by getBackendBucket.
 */
export interface GetBackendBucketResult {
    /**
     * Cloud Storage bucket name.
     */
    readonly bucketName: string;
    /**
     * Cloud CDN configuration for this Backend Bucket. Structure is documented below.
     */
    readonly cdnPolicies: outputs.compute.GetBackendBucketCdnPolicy[];
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional textual description of the resource; provided by the client when the resource is created.
     */
    readonly description: string;
    /**
     * Whether Cloud CDN is enabled for this BackendBucket.
     */
    readonly enableCdn: boolean;
    readonly name: string;
    readonly project?: string;
    /**
     * The URI of the created resource.
     */
    readonly selfLink: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
