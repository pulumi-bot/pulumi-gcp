// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get a Compute Region Instance Group within GCE.
 * For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/latest/regionInstanceGroups).
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const group = gcp.compute.getRegionInstanceGroup({
 *     name: "instance-group-name",
 * });
 * ```
 * 
 * The most common use of this datasource will be to fetch information about the instances inside regional managed instance groups, for instance:
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/compute_region_instance_group.html.markdown.
 */
export function getRegionInstanceGroup(args?: GetRegionInstanceGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionInstanceGroupResult> & GetRegionInstanceGroupResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetRegionInstanceGroupResult> = pulumi.runtime.invoke("gcp:compute/getRegionInstanceGroup:getRegionInstanceGroup", {
        "name": args.name,
        "project": args.project,
        "region": args.region,
        "selfLink": args.selfLink,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getRegionInstanceGroup.
 */
export interface GetRegionInstanceGroupArgs {
    /**
     * The name of the instance group.  One of `name` or `selfLink` must be provided.
     */
    readonly name?: string;
    /**
     * The ID of the project in which the resource belongs.
     * If `selfLink` is provided, this value is ignored.  If neither `selfLink`
     * nor `project` are provided, the provider project is used.
     */
    readonly project?: string;
    /**
     * The region in which the resource belongs.  If `selfLink`
     * is provided, this value is ignored.  If neither `selfLink` nor `region` are
     * provided, the provider region is used.
     */
    readonly region?: string;
    /**
     * The link to the instance group.  One of `name` or `selfLink` must be provided.
     */
    readonly selfLink?: string;
}

/**
 * A collection of values returned by getRegionInstanceGroup.
 */
export interface GetRegionInstanceGroupResult {
    /**
     * List of instances in the group, as a list of resources, each containing:
     */
    readonly instances: outputs.compute.GetRegionInstanceGroupInstance[];
    /**
     * String port name
     */
    readonly name: string;
    readonly project: string;
    readonly region: string;
    readonly selfLink: string;
    /**
     * The number of instances in the group.
     */
    readonly size: number;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
