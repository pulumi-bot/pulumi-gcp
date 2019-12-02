// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Gets an existing object inside an existing bucket in Google Cloud Storage service (GCS).
 * See [the official documentation](https://cloud.google.com/storage/docs/key-terms#objects)
 * and
 * [API](https://cloud.google.com/storage/docs/json_api/v1/objects).
 * 
 * 
 * ## Example Usage
 * 
 * Example picture stored within a folder.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const picture = gcp.storage.getBucketObject({
 *     bucket: "image-store",
 *     name: "folder/butterfly01.jpg",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/storage_bucket_object.html.markdown.
 */
export function getBucketObject(args?: GetBucketObjectArgs, opts?: pulumi.InvokeOptions): Promise<GetBucketObjectResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:storage/getBucketObject:getBucketObject", {
        "bucket": args.bucket,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getBucketObject.
 */
export interface GetBucketObjectArgs {
    /**
     * The name of the containing bucket.
     */
    readonly bucket?: string;
    /**
     * The name of the object.
     */
    readonly name?: string;
}

/**
 * A collection of values returned by getBucketObject.
 */
export interface GetBucketObjectResult {
    readonly bucket?: string;
    /**
     * (Computed) [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
     * directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
     */
    readonly cacheControl: string;
    readonly content: string;
    /**
     * (Computed) [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
     */
    readonly contentDisposition: string;
    /**
     * (Computed) [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
     */
    readonly contentEncoding: string;
    /**
     * (Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
     */
    readonly contentLanguage: string;
    /**
     * (Computed) [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
     */
    readonly contentType: string;
    /**
     * (Computed) Base 64 CRC32 hash of the uploaded data.
     */
    readonly crc32c: string;
    readonly detectMd5hash: string;
    /**
     * (Computed) Base 64 MD5 hash of the uploaded data.
     */
    readonly md5hash: string;
    readonly name?: string;
    readonly outputName: string;
    readonly predefinedAcl: string;
    /**
     * (Computed) A url reference to this object.
     */
    readonly selfLink: string;
    readonly source: string;
    /**
     * (Computed) The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
     * Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`. If not provided, this defaults to the bucket's default
     * storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
     */
    readonly storageClass: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
