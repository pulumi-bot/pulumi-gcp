// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a new object inside an existing bucket in Google cloud storage service (GCS). 
 * [ACLs](https://cloud.google.com/storage/docs/access-control/lists) can be applied using the `gcp.storage.ObjectACL` resource.
 *  For more information see 
 * [the official documentation](https://cloud.google.com/storage/docs/key-terms#objects) 
 * and 
 * [API](https://cloud.google.com/storage/docs/json_api/v1/objects).
 *
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const picture = new gcp.storage.BucketObject("picture", {
 *     bucket: "image-store",
 *     source: new pulumi.asset.FileAsset("/images/nature/garden-tiger-moth.jpg"),
 * });
 * ```
 */
export class BucketObject extends pulumi.CustomResource {
    /**
     * Get an existing BucketObject resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketObjectState, opts?: pulumi.CustomResourceOptions): BucketObject {
        return new BucketObject(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:storage/bucketObject:BucketObject';

    /**
     * Returns true if the given object is an instance of BucketObject.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketObject {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketObject.__pulumiType;
    }

    /**
     * The name of the containing bucket.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
     * directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
     */
    public readonly cacheControl!: pulumi.Output<string | undefined>;
    /**
     * Data as `string` to be uploaded. Must be defined if `source` is not. **Note**: The `content` field is marked as sensitive.
     */
    public readonly content!: pulumi.Output<string | undefined>;
    /**
     * [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
     */
    public readonly contentDisposition!: pulumi.Output<string | undefined>;
    /**
     * [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
     */
    public readonly contentEncoding!: pulumi.Output<string | undefined>;
    /**
     * [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
     */
    public readonly contentLanguage!: pulumi.Output<string | undefined>;
    /**
     * [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
     */
    public readonly contentType!: pulumi.Output<string>;
    /**
     * (Computed) Base 64 CRC32 hash of the uploaded data.
     */
    public /*out*/ readonly crc32c!: pulumi.Output<string>;
    public readonly detectMd5hash!: pulumi.Output<string | undefined>;
    /**
     * (Computed) Base 64 MD5 hash of the uploaded data.
     */
    public /*out*/ readonly md5hash!: pulumi.Output<string>;
    /**
     * User-provided metadata, in key/value pairs.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the object. If you're interpolating the name of this object, see `outputName` instead.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * (Computed) The name of the object. Use this field in interpolations with `gcp.storage.ObjectACL` to recreate
     * `gcp.storage.ObjectACL` resources when your `gcp.storage.BucketObject` is recreated.
     */
    public /*out*/ readonly outputName!: pulumi.Output<string>;
    /**
     * (Computed) A url reference to this object.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * A path to the data you want to upload. Must be defined
     * if `content` is not.
     */
    public readonly source!: pulumi.Output<pulumi.asset.Asset | pulumi.asset.Archive | undefined>;
    /**
     * The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
     * Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`. If not provided, this defaults to the bucket's default
     * storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
     */
    public readonly storageClass!: pulumi.Output<string>;

    /**
     * Create a BucketObject resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketObjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketObjectArgs | BucketObjectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BucketObjectState | undefined;
            inputs["bucket"] = state ? state.bucket : undefined;
            inputs["cacheControl"] = state ? state.cacheControl : undefined;
            inputs["content"] = state ? state.content : undefined;
            inputs["contentDisposition"] = state ? state.contentDisposition : undefined;
            inputs["contentEncoding"] = state ? state.contentEncoding : undefined;
            inputs["contentLanguage"] = state ? state.contentLanguage : undefined;
            inputs["contentType"] = state ? state.contentType : undefined;
            inputs["crc32c"] = state ? state.crc32c : undefined;
            inputs["detectMd5hash"] = state ? state.detectMd5hash : undefined;
            inputs["md5hash"] = state ? state.md5hash : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["outputName"] = state ? state.outputName : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["source"] = state ? state.source : undefined;
            inputs["storageClass"] = state ? state.storageClass : undefined;
        } else {
            const args = argsOrState as BucketObjectArgs | undefined;
            if (!args || args.bucket === undefined) {
                throw new Error("Missing required property 'bucket'");
            }
            inputs["bucket"] = args ? args.bucket : undefined;
            inputs["cacheControl"] = args ? args.cacheControl : undefined;
            inputs["content"] = args ? args.content : undefined;
            inputs["contentDisposition"] = args ? args.contentDisposition : undefined;
            inputs["contentEncoding"] = args ? args.contentEncoding : undefined;
            inputs["contentLanguage"] = args ? args.contentLanguage : undefined;
            inputs["contentType"] = args ? args.contentType : undefined;
            inputs["detectMd5hash"] = args ? args.detectMd5hash : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["storageClass"] = args ? args.storageClass : undefined;
            inputs["crc32c"] = undefined /*out*/;
            inputs["md5hash"] = undefined /*out*/;
            inputs["outputName"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BucketObject.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BucketObject resources.
 */
export interface BucketObjectState {
    /**
     * The name of the containing bucket.
     */
    readonly bucket?: pulumi.Input<string>;
    /**
     * [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
     * directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
     */
    readonly cacheControl?: pulumi.Input<string>;
    /**
     * Data as `string` to be uploaded. Must be defined if `source` is not. **Note**: The `content` field is marked as sensitive.
     */
    readonly content?: pulumi.Input<string>;
    /**
     * [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
     */
    readonly contentDisposition?: pulumi.Input<string>;
    /**
     * [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
     */
    readonly contentEncoding?: pulumi.Input<string>;
    /**
     * [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
     */
    readonly contentLanguage?: pulumi.Input<string>;
    /**
     * [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
     */
    readonly contentType?: pulumi.Input<string>;
    /**
     * (Computed) Base 64 CRC32 hash of the uploaded data.
     */
    readonly crc32c?: pulumi.Input<string>;
    readonly detectMd5hash?: pulumi.Input<string>;
    /**
     * (Computed) Base 64 MD5 hash of the uploaded data.
     */
    readonly md5hash?: pulumi.Input<string>;
    /**
     * User-provided metadata, in key/value pairs.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the object. If you're interpolating the name of this object, see `outputName` instead.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * (Computed) The name of the object. Use this field in interpolations with `gcp.storage.ObjectACL` to recreate
     * `gcp.storage.ObjectACL` resources when your `gcp.storage.BucketObject` is recreated.
     */
    readonly outputName?: pulumi.Input<string>;
    /**
     * (Computed) A url reference to this object.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * A path to the data you want to upload. Must be defined
     * if `content` is not.
     */
    readonly source?: pulumi.Input<pulumi.asset.Asset | pulumi.asset.Archive>;
    /**
     * The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
     * Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`. If not provided, this defaults to the bucket's default
     * storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
     */
    readonly storageClass?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BucketObject resource.
 */
export interface BucketObjectArgs {
    /**
     * The name of the containing bucket.
     */
    readonly bucket: pulumi.Input<string>;
    /**
     * [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
     * directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
     */
    readonly cacheControl?: pulumi.Input<string>;
    /**
     * Data as `string` to be uploaded. Must be defined if `source` is not. **Note**: The `content` field is marked as sensitive.
     */
    readonly content?: pulumi.Input<string>;
    /**
     * [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
     */
    readonly contentDisposition?: pulumi.Input<string>;
    /**
     * [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
     */
    readonly contentEncoding?: pulumi.Input<string>;
    /**
     * [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
     */
    readonly contentLanguage?: pulumi.Input<string>;
    /**
     * [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
     */
    readonly contentType?: pulumi.Input<string>;
    readonly detectMd5hash?: pulumi.Input<string>;
    /**
     * User-provided metadata, in key/value pairs.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the object. If you're interpolating the name of this object, see `outputName` instead.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A path to the data you want to upload. Must be defined
     * if `content` is not.
     */
    readonly source?: pulumi.Input<pulumi.asset.Asset | pulumi.asset.Archive>;
    /**
     * The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
     * Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`. If not provided, this defaults to the bucket's default
     * storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
     */
    readonly storageClass?: pulumi.Input<string>;
}
