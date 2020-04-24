// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get a Secret Manager secret's version. For more information see the [official documentation](https://cloud.google.com/secret-manager/docs/) and [API](https://cloud.google.com/secret-manager/docs/reference/rest/v1beta1/projects.secrets.versions).
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_google_secret_manager_secret_version.html.markdown.
 */
/** @deprecated gcp.getSecretVersion has been deprecated in favour of gcp.getSecretVersion */
export function getSecretVersion(args: GetSecretVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetSecretVersionResult> {
    pulumi.log.warn("getSecretVersion is deprecated: gcp.getSecretVersion has been deprecated in favour of gcp.getSecretVersion")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:monitoring/getSecretVersion:getSecretVersion", {
        "project": args.project,
        "secret": args.secret,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecretVersion.
 */
export interface GetSecretVersionArgs {
    /**
     * The project to get the secret version for. If it
     * is not provided, the provider project is used.
     */
    readonly project?: string;
    /**
     * The secret to get the secret version for.
     */
    readonly secret: string;
    /**
     * The version of the secret to get. If it
     * is not provided, the latest version is retrieved.
     */
    readonly version?: string;
}

/**
 * A collection of values returned by getSecretVersion.
 */
export interface GetSecretVersionResult {
    /**
     * The time at which the Secret was created.
     */
    readonly createTime: string;
    /**
     * The time at which the Secret was destroyed. Only present if state is DESTROYED.
     */
    readonly destroyTime: string;
    /**
     * True if the current state of the SecretVersion is enabled.
     */
    readonly enabled: boolean;
    /**
     * The resource name of the SecretVersion. Format:
     * `projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}`
     */
    readonly name: string;
    readonly project: string;
    readonly secret: string;
    /**
     * The secret data. No larger than 64KiB.
     */
    readonly secretData: string;
    readonly version: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
