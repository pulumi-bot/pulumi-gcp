// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Allows management of Organization policies for a Google Project. For more information see
 * [the official
 * documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview)
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const policy = gcp.projects.getOrganizationPolicy({
 *     project: "project-id",
 *     constraint: "constraints/serviceuser.services",
 * });
 * export const version = policy.then(policy => policy.version);
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_google_project_organization_policy.html.markdown.
 */
export function getOrganizationPolicy(args: GetOrganizationPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetOrganizationPolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:projects/getOrganizationPolicy:getOrganizationPolicy", {
        "constraint": args.constraint,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrganizationPolicy.
 */
export interface GetOrganizationPolicyArgs {
    /**
     * (Required) The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
     */
    readonly constraint: string;
    /**
     * The project ID.
     */
    readonly project: string;
}

/**
 * A collection of values returned by getOrganizationPolicy.
 */
export interface GetOrganizationPolicyResult {
    readonly booleanPolicies: outputs.projects.GetOrganizationPolicyBooleanPolicy[];
    readonly constraint: string;
    readonly etag: string;
    readonly listPolicies: outputs.projects.GetOrganizationPolicyListPolicy[];
    readonly project: string;
    readonly restorePolicies: outputs.projects.GetOrganizationPolicyRestorePolicy[];
    readonly updateTime: string;
    readonly version: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
