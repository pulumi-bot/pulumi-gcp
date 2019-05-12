// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieve information about a set of projects based on a filter. See the
 * [REST API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list)
 * for more details.
 * 
 * ## Example Usage - searching for projects about to be deleted in an org
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const my_org_projects = pulumi.output(gcp.projects.getProject({
 *     filter: "parent.id:012345678910 lifecycleState:DELETE_REQUESTED",
 * }));
 * const deletion_candidate = my_org_projects.apply(my_org_projects => gcp.organizations.getProject({
 *     projectId: my_org_projects.projects[0].projectId,
 * }));
 * ```
 */
export function getProject(args: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:projects/getProject:getProject", {
        "filter": args.filter,
    }, opts);
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectArgs {
    /**
     * A string filter as defined in the [REST API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list#query-parameters).
     */
    readonly filter: string;
}

/**
 * A collection of values returned by getProject.
 */
export interface GetProjectResult {
    readonly filter: string;
    /**
     * A list of projects matching the provided filter. Structure is defined below.
     */
    readonly projects: { projectId: string }[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
