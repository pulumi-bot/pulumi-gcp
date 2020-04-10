// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get all of the trusted Certificate Authorities (CAs) for the specified SQL database instance. For more information see the
 * [official documentation](https://cloud.google.com/sql/)
 * and
 * [API](https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1beta4/instances/listServerCas).
 * 
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_google_sql_ca_certs.html.markdown.
 */
export function getCaCerts(args: GetCaCertsArgs, opts?: pulumi.InvokeOptions): Promise<GetCaCertsResult> & GetCaCertsResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetCaCertsResult> = pulumi.runtime.invoke("gcp:sql/getCaCerts:getCaCerts", {
        "instance": args.instance,
        "project": args.project,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getCaCerts.
 */
export interface GetCaCertsArgs {
    /**
     * The name or self link of the instance.
     */
    readonly instance: string;
    /**
     * The ID of the project in which the resource belongs. If `project` is not provided, the provider project is used.
     */
    readonly project?: string;
}

/**
 * A collection of values returned by getCaCerts.
 */
export interface GetCaCertsResult {
    /**
     * SHA1 fingerprint of the currently active CA certificate.
     */
    readonly activeVersion: string;
    /**
     * A list of server CA certificates for the instance. Each contains:
     */
    readonly certs: outputs.sql.GetCaCertsCert[];
    readonly instance: string;
    readonly project: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
