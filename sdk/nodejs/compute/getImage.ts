// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getImage(args?: GetImageArgs, opts?: pulumi.InvokeOptions): Promise<GetImageResult> {
    args = args || {};
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
    readonly family?: string;
    readonly name?: string;
    readonly project?: string;
}

/**
 * A collection of values returned by getImage.
 */
export interface GetImageResult {
    readonly archiveSizeBytes: number;
    readonly creationTimestamp: string;
    readonly description: string;
    readonly diskSizeGb: number;
    readonly family: string;
    readonly imageEncryptionKeySha256: string;
    readonly imageId: string;
    readonly labelFingerprint: string;
    readonly labels: {[key: string]: string};
    readonly licenses: string[];
    readonly name: string;
    readonly project: string;
    readonly selfLink: string;
    readonly sourceDisk: string;
    readonly sourceDiskEncryptionKeySha256: string;
    readonly sourceDiskId: string;
    readonly sourceImageId: string;
    readonly status: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
