// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source fetches the project name, and provides the appropriate URLs to use for container registry for this project.
 * 
 * The URLs are computed entirely offline - as long as the project exists, they will be valid, but this data source does not contact Google Container Registry (GCR) at any point.
 * 
 * {{% examples %}}
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/google_container_registry_image.html.markdown.
 */
export function getRegistryImage(args: GetRegistryImageArgs, opts?: pulumi.InvokeOptions): Promise<GetRegistryImageResult> & GetRegistryImageResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetRegistryImageResult> = pulumi.runtime.invoke("gcp:container/getRegistryImage:getRegistryImage", {
        "digest": args.digest,
        "name": args.name,
        "project": args.project,
        "region": args.region,
        "tag": args.tag,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getRegistryImage.
 */
export interface GetRegistryImageArgs {
    readonly digest?: string;
    readonly name: string;
    readonly project?: string;
    readonly region?: string;
    readonly tag?: string;
}

/**
 * A collection of values returned by getRegistryImage.
 */
export interface GetRegistryImageResult {
    readonly digest?: string;
    readonly imageUrl: string;
    readonly name: string;
    readonly project: string;
    readonly region?: string;
    readonly tag?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
