// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get the email address of a project's unique Google Cloud Storage service account.
 *
 * Each Google Cloud project has a unique service account for use with Google Cloud Storage. Only this
 * special service account can be used to set up `gcp.storage.Notification` resources.
 *
 * For more information see
 * [the API reference](https://cloud.google.com/storage/docs/json_api/v1/projects/serviceAccount).
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const gcsAccount = gcp.storage.getProjectServiceAccount({});
 * const binding = new gcp.pubsub.TopicIAMBinding("binding", {
 *     topic: google_pubsub_topic.topic.name,
 *     role: "roles/pubsub.publisher",
 *     members: [gcsAccount.then(gcsAccount => `serviceAccount:${gcsAccount.emailAddress}`)],
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/google_storage_project_service_account.html.markdown.
 */
export function getProjectServiceAccount(args?: GetProjectServiceAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectServiceAccountResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:storage/getProjectServiceAccount:getProjectServiceAccount", {
        "project": args.project,
        "userProject": args.userProject,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectServiceAccount.
 */
export interface GetProjectServiceAccountArgs {
    /**
     * The project the unique service account was created for. If it is not provided, the provider project is used.
     */
    readonly project?: string;
    /**
     * The project the lookup originates from. This field is used if you are making the request
     * from a different account than the one you are finding the service account for.
     */
    readonly userProject?: string;
}

/**
 * A collection of values returned by getProjectServiceAccount.
 */
export interface GetProjectServiceAccountResult {
    /**
     * The email address of the service account. This value is often used to refer to the service account
     * in order to grant IAM permissions.
     */
    readonly emailAddress: string;
    readonly project: string;
    readonly userProject?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
