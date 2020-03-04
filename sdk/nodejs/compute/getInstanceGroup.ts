// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get a Compute Instance Group within GCE.
 * For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/#unmanaged_instance_groups)
 * and [API](https://cloud.google.com/compute/docs/reference/latest/instanceGroups)
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const all = gcp.compute.getInstanceGroup({
 *     name: "instance-group-name",
 *     zone: "us-central1-a",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/google_compute_instance_group.html.markdown.
 */
export function getInstanceGroup(args?: GetInstanceGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceGroupResult> & GetInstanceGroupResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetInstanceGroupResult> = pulumi.runtime.invoke("gcp:compute/getInstanceGroup:getInstanceGroup", {
        "name": args.name,
        "project": args.project,
        "selfLink": args.selfLink,
        "zone": args.zone,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getInstanceGroup.
 */
export interface GetInstanceGroupArgs {
    /**
     * The name of the instance group. Either `name` or `selfLink` must be provided.
     */
    readonly name?: string;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: string;
    /**
     * The self link of the instance group. Either `name` or `selfLink` must be provided.
     */
    readonly selfLink?: string;
    /**
     * The zone of the instance group. If referencing the instance group by name
     * and `zone` is not provided, the provider zone is used.
     */
    readonly zone?: string;
}

/**
 * A collection of values returned by getInstanceGroup.
 */
export interface GetInstanceGroupResult {
    /**
     * Textual description of the instance group.
     */
    readonly description: string;
    /**
     * List of instances in the group.
     */
    readonly instances: string[];
    readonly name?: string;
    /**
     * List of named ports in the group.
     */
    readonly namedPorts: outputs.compute.GetInstanceGroupNamedPort[];
    /**
     * The URL of the network the instance group is in.
     */
    readonly network: string;
    readonly project: string;
    /**
     * The URI of the resource.
     */
    readonly selfLink: string;
    /**
     * The number of instances in the group.
     */
    readonly size: number;
    readonly zone: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
