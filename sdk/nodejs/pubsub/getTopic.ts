// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Get information about a Google Cloud Pub/Sub Topic. For more information see
 * the [official documentation](https://cloud.google.com/pubsub/docs/)
 * and [API](https://cloud.google.com/pubsub/docs/apis).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const my_pubsub_topic = pulumi.output(gcp.pubsub.getTopic({
 *     name: "my-pubsub-topic",
 * }));
 * ```
 */
export function getTopic(args: GetTopicArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:pubsub/getTopic:getTopic", {
        "name": args.name,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getTopic.
 */
export interface GetTopicArgs {
    /**
     * The name of the Cloud Pub/Sub Topic.
     */
    name: string;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: string;
}

/**
 * A collection of values returned by getTopic.
 */
export interface GetTopicResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly kmsKeyName: string;
    readonly labels: {[key: string]: string};
    readonly messageStoragePolicies: outputs.pubsub.GetTopicMessageStoragePolicy[];
    readonly name: string;
    readonly project?: string;
    readonly schemaSettings: outputs.pubsub.GetTopicSchemaSetting[];
}

export function getTopicApply(args: GetTopicApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTopicResult> {
    return pulumi.output(args).apply(a => getTopic(a, opts))
}

/**
 * A collection of arguments for invoking getTopic.
 */
export interface GetTopicApplyArgs {
    /**
     * The name of the Cloud Pub/Sub Topic.
     */
    name: pulumi.Input<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
