// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides available node types for Compute Engine sole-tenant nodes in a zone
 * for a given project. For more information, see [the official documentation](https://cloud.google.com/compute/docs/nodes/#types) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/nodeTypes).
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const central1b = gcp.compute.getNodeTypes({
 *     zone: "us-central1-b",
 * });
 * const tmpl = new gcp.compute.NodeTemplate("tmpl", {
 *     region: "us-central1",
 *     nodeType: data.google_compute_node_types.types.names[0],
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/google_compute_node_types.html.markdown.
 */
export function getNodeTypes(args?: GetNodeTypesArgs, opts?: pulumi.InvokeOptions): Promise<GetNodeTypesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:compute/getNodeTypes:getNodeTypes", {
        "project": args.project,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getNodeTypes.
 */
export interface GetNodeTypesArgs {
    /**
     * ID of the project to list available node types for.
     * Should match the project the nodes of this type will be deployed to.
     * Defaults to the project that the provider is authenticated with.
     */
    readonly project?: string;
    /**
     * The zone to list node types for. Should be in zone of intended node groups and region of referencing node template. If `zone` is not specified, the provider-level zone must be set and is used
     * instead.
     */
    readonly zone?: string;
}

/**
 * A collection of values returned by getNodeTypes.
 */
export interface GetNodeTypesResult {
    /**
     * A list of node types available in the given zone and project.
     */
    readonly names: string[];
    readonly project: string;
    readonly zone: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
