// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Get all of the trusted Certificate Authorities (CAs) for the specified SQL database instance. For more information see the
 * [official documentation](https://cloud.google.com/sql/)
 * and
 * [API](https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1beta4/instances/listServerCas).
 */
export function getCaCerts(args: GetCaCertsArgs, opts?: pulumi.InvokeOptions): Promise<GetCaCertsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:sql/getCaCerts:getCaCerts", {
        "instance": args.instance,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getCaCerts.
 */
export interface GetCaCertsArgs {
    /**
     * The name or self link of the instance.
     */
    instance: string;
    /**
     * The ID of the project in which the resource belongs. If `project` is not provided, the provider project is used.
     */
    project?: string;
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
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instance: string;
    readonly project: string;
}

export function getCaCertsOutput(args: GetCaCertsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCaCertsResult> {
    return pulumi.output(args).apply(a => getCaCerts(a, opts))
}

/**
 * A collection of arguments for invoking getCaCerts.
 */
export interface GetCaCertsOutputArgs {
    /**
     * The name or self link of the instance.
     */
    instance: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If `project` is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
