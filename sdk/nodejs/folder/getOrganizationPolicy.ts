// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Allows management of Organization policies for a Google Folder. For more information see
 * [the official
 * documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.folder.getOrganizationPolicy({
 *     folder: "folders/folderid",
 *     constraint: "constraints/compute.trustedImageProjects",
 * });
 * export const version = policy.then(policy => policy.version);
 * ```
 */
export function getOrganizationPolicy(args: GetOrganizationPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetOrganizationPolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:folder/getOrganizationPolicy:getOrganizationPolicy", {
        "constraint": args.constraint,
        "folder": args.folder,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrganizationPolicy.
 */
export interface GetOrganizationPolicyArgs {
    /**
     * (Required) The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
     */
    constraint: string;
    /**
     * The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
     */
    folder: string;
}

/**
 * A collection of values returned by getOrganizationPolicy.
 */
export interface GetOrganizationPolicyResult {
    readonly booleanPolicies: outputs.folder.GetOrganizationPolicyBooleanPolicy[];
    readonly constraint: string;
    readonly etag: string;
    readonly folder: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly listPolicies: outputs.folder.GetOrganizationPolicyListPolicy[];
    readonly restorePolicies: outputs.folder.GetOrganizationPolicyRestorePolicy[];
    readonly updateTime: string;
    readonly version: number;
}

export function getOrganizationPolicyApply(args: GetOrganizationPolicyApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrganizationPolicyResult> {
    return pulumi.output(args).apply(a => getOrganizationPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getOrganizationPolicy.
 */
export interface GetOrganizationPolicyApplyArgs {
    /**
     * (Required) The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
     */
    constraint: pulumi.Input<string>;
    /**
     * The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
     */
    folder: pulumi.Input<string>;
}
