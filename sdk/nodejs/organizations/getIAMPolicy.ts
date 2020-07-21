// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Generates an IAM policy document that may be referenced by and applied to
 * other Google Cloud Platform resources, such as the `gcp.organizations.Project` resource.
 *
 * **Note:** Several restrictions apply when setting IAM policies through this API.
 * See the [setIamPolicy docs](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy)
 * for a list of these restrictions.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = pulumi.output(gcp.organizations.getIAMPolicy({
 *     auditConfigs: [{
 *         auditLogConfigs: [
 *             {
 *                 exemptedMembers: ["user:you@domain.com"],
 *                 logType: "DATA_READ",
 *             },
 *             {
 *                 logType: "DATA_WRITE",
 *             },
 *             {
 *                 logType: "ADMIN_READ",
 *             },
 *         ],
 *         service: "cloudkms.googleapis.com",
 *     }],
 *     bindings: [
 *         {
 *             members: ["serviceAccount:your-custom-sa@your-project.iam.gserviceaccount.com"],
 *             role: "roles/compute.instanceAdmin",
 *         },
 *         {
 *             members: ["user:alice@gmail.com"],
 *             role: "roles/storage.objectViewer",
 *         },
 *     ],
 * }, { async: true }));
 * ```
 *
 * This data source is used to define IAM policies to apply to other resources.
 * Currently, defining a policy through a datasource and referencing that policy
 * from another resource is the only way to apply an IAM policy to a resource.
 */
export function getIAMPolicy(args?: GetIAMPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetIAMPolicyResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:organizations/getIAMPolicy:getIAMPolicy", {
        "auditConfigs": args.auditConfigs,
        "bindings": args.bindings,
    }, opts);
}

/**
 * A collection of arguments for invoking getIAMPolicy.
 */
export interface GetIAMPolicyArgs {
    /**
     * A nested configuration block that defines logging additional configuration for your project.
     */
    readonly auditConfigs?: inputs.organizations.GetIAMPolicyAuditConfig[];
    /**
     * A nested configuration block (described below)
     * defining a binding to be included in the policy document. Multiple
     * `binding` arguments are supported.
     */
    readonly bindings?: inputs.organizations.GetIAMPolicyBinding[];
}

/**
 * A collection of values returned by getIAMPolicy.
 */
export interface GetIAMPolicyResult {
    readonly auditConfigs?: outputs.organizations.GetIAMPolicyAuditConfig[];
    readonly bindings?: outputs.organizations.GetIAMPolicyBinding[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The above bindings serialized in a format suitable for
     * referencing from a resource that supports IAM.
     */
    readonly policyData: string;
}
