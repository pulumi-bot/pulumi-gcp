// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get information about a Google Compute Image. Check that your service account has the `compute.imageUser` role if you want to share [custom images](https://cloud.google.com/compute/docs/images/sharing-images-across-projects) from another project. If you want to use [public images][pubimg], do not forget to specify the dedicated project. For more information see
 * [the official documentation](https://cloud.google.com/compute/docs/images) and its [API](https://cloud.google.com/compute/docs/reference/latest/images).
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const myImage = gcp.compute.getImage({
 *     family: "debian-9",
 *     project: "debian-cloud",
 * });
 * // ...
 * const _default = new gcp.compute.Instance("default", {boot_disk: {
 *     initialize_params: {
 *         image: myImage.then(myImage => myImage.selfLink),
 *     },
 * }});
 * ```
 */
export function getImage(args?: GetImageArgs, opts?: pulumi.InvokeOptions): Promise<GetImageResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:compute/getImage:getImage", {
        "family": args.family,
        "name": args.name,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getImage.
 */
export interface GetImageArgs {
    /**
     * The family name of the image.
     */
    readonly family?: string;
    /**
     * or `family` - (Required) The name of a specific image or a family.
     * Exactly one of `name` of `family` must be specified. If `name` is specified, it will fetch
     * the corresponding image. If `family` is specified, it will returns the latest image
     * that is part of an image family and is not deprecated.
     */
    readonly name?: string;
    /**
     * The project in which the resource belongs. If it is not
     * provided, the provider project is used. If you are using a
     * [public base image][pubimg], be sure to specify the correct Image Project.
     */
    readonly project?: string;
}

/**
 * A collection of values returned by getImage.
 */
export interface GetImageResult {
    /**
     * The size of the image tar.gz archive stored in Google Cloud Storage in bytes.
     */
    readonly archiveSizeBytes: number;
    /**
     * The creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this image.
     */
    readonly description: string;
    /**
     * The size of the image when restored onto a persistent disk in gigabytes.
     */
    readonly diskSizeGb: number;
    /**
     * The family name of the image.
     */
    readonly family: string;
    /**
     * The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4)
     * encoded SHA-256 hash of the [customer-supplied encryption key](https://cloud.google.com/compute/docs/disks/customer-supplied-encryption)
     * that protects this image.
     */
    readonly imageEncryptionKeySha256: string;
    /**
     * The unique identifier for the image.
     */
    readonly imageId: string;
    /**
     * A fingerprint for the labels being applied to this image.
     */
    readonly labelFingerprint: string;
    /**
     * A map of labels applied to this image.
     */
    readonly labels: {[key: string]: string};
    /**
     * A list of applicable license URI.
     */
    readonly licenses: string[];
    /**
     * The name of the image.
     */
    readonly name: string;
    readonly project: string;
    /**
     * The URI of the image.
     */
    readonly selfLink: string;
    /**
     * The URL of the source disk used to create this image.
     */
    readonly sourceDisk: string;
    /**
     * The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4)
     * encoded SHA-256 hash of the [customer-supplied encryption key](https://cloud.google.com/compute/docs/disks/customer-supplied-encryption)
     * that protects this image.
     */
    readonly sourceDiskEncryptionKeySha256: string;
    /**
     * The ID value of the disk used to create this image.
     */
    readonly sourceDiskId: string;
    /**
     * The ID value of the image used to create this image.
     */
    readonly sourceImageId: string;
    /**
     * The status of the image. Possible values are **FAILED**, **PENDING**, or **READY**.
     */
    readonly status: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
