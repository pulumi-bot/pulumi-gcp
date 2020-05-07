// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get info about a Google Compute SSL Certificate from its name.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const myCert = gcp.compute.getCertificate({
 *     name: "my-cert",
 * });
 * export const certificate = myCert.then(myCert => myCert.certificate);
 * export const certificateId = myCert.then(myCert => myCert.certificateId);
 * export const selfLink = myCert.then(myCert => myCert.selfLink);
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_compute_ssl_certificate.html.markdown.
 */
export function getCertificate(args: GetCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:compute/getCertificate:getCertificate", {
        "name": args.name,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getCertificate.
 */
export interface GetCertificateArgs {
    /**
     * The name of the certificate.
     */
    readonly name: string;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: string;
}

/**
 * A collection of values returned by getCertificate.
 */
export interface GetCertificateResult {
    readonly certificate: string;
    readonly certificateId: number;
    readonly creationTimestamp: string;
    readonly description: string;
    readonly name: string;
    readonly namePrefix: string;
    readonly privateKey: string;
    readonly project?: string;
    readonly selfLink: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
