// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage IAM policies on dataproc clusters. Each of these resources serves a different use case:
 * 
 * * `google_dataproc_cluster_iam_policy`: Authoritative. Sets the IAM policy for the cluster and replaces any existing policy already attached.
 * * `google_dataproc_cluster_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the cluster are preserved.
 * * `google_dataproc_cluster_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the cluster are preserved.
 * 
 * > **Note:** `google_dataproc_cluster_iam_policy` **cannot** be used in conjunction with `google_dataproc_cluster_iam_binding` and `google_dataproc_cluster_iam_member` or they will fight over what your policy should be. In addition, be careful not to accidentaly unset ownership of the cluster as `google_dataproc_cluster_iam_policy` replaces the entire policy.
 * 
 * > **Note:** `google_dataproc_cluster_iam_binding` resources **can be** used in conjunction with `google_dataproc_cluster_iam_member` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_pubsub\_subscription\_iam\_policy
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const admin = pulumi.output(gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         members: ["user:jane@example.com"],
 *         role: "roles/editor",
 *     }],
 * }));
 * const editor = new gcp.dataproc.ClusterIAMPolicy("editor", {
 *     cluster: "your-dataproc-cluster",
 *     policyData: admin.policyData,
 *     project: "your-project",
 *     region: "your-region",
 * });
 * ```
 * 
 * ## google\_pubsub\_subscription\_iam\_binding
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const editor = new gcp.dataproc.ClusterIAMBinding("editor", {
 *     cluster: "your-dataproc-cluster",
 *     members: ["user:jane@example.com"],
 *     role: "roles/editor",
 * });
 * ```
 * 
 * ## google\_pubsub\_subscription\_iam\_member
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const editor = new gcp.dataproc.ClusterIAMMember("editor", {
 *     cluster: "your-dataproc-cluster",
 *     member: "user:jane@example.com",
 *     role: "roles/editor",
 * });
 * ```
 */
export class ClusterIAMPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ClusterIAMPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterIAMPolicyState, opts?: pulumi.CustomResourceOptions): ClusterIAMPolicy {
        return new ClusterIAMPolicy(name, <any>state, { ...opts, id: id });
    }

    private static readonly __pulumiType = 'gcp:dataproc/clusterIAMPolicy:ClusterIAMPolicy';

    /**
     * Returns true if the given object is an instance of ClusterIAMPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterIAMPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }

        const t = obj['__pulumiType'];
        if (typeof t !== 'string') {
            return false;
        }

        return t === ClusterIAMPolicy.__pulumiType;
    }

    /**
     * The name or relative resource id of the cluster to manage IAM policies for.
     */
    public readonly cluster!: pulumi.Output<string>;
    /**
     * (Computed) The etag of the clusters's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The policy data generated by a `google_iam_policy` data source.
     */
    public readonly policyData!: pulumi.Output<string>;
    /**
     * The project in which the cluster belongs. If it
     * is not provided, Terraform will use the provider default.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The region in which the cluster belongs. If it
     * is not provided, Terraform will use the provider default.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a ClusterIAMPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterIAMPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterIAMPolicyArgs | ClusterIAMPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClusterIAMPolicyState | undefined;
            inputs["cluster"] = state ? state.cluster : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as ClusterIAMPolicyArgs | undefined;
            if (!args || args.cluster === undefined) {
                throw new Error("Missing required property 'cluster'");
            }
            if (!args || args.policyData === undefined) {
                throw new Error("Missing required property 'policyData'");
            }
            inputs["cluster"] = args ? args.cluster : undefined;
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        super(ClusterIAMPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClusterIAMPolicy resources.
 */
export interface ClusterIAMPolicyState {
    /**
     * The name or relative resource id of the cluster to manage IAM policies for.
     */
    readonly cluster?: pulumi.Input<string>;
    /**
     * (Computed) The etag of the clusters's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The policy data generated by a `google_iam_policy` data source.
     */
    readonly policyData?: pulumi.Input<string>;
    /**
     * The project in which the cluster belongs. If it
     * is not provided, Terraform will use the provider default.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region in which the cluster belongs. If it
     * is not provided, Terraform will use the provider default.
     */
    readonly region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClusterIAMPolicy resource.
 */
export interface ClusterIAMPolicyArgs {
    /**
     * The name or relative resource id of the cluster to manage IAM policies for.
     */
    readonly cluster: pulumi.Input<string>;
    /**
     * The policy data generated by a `google_iam_policy` data source.
     */
    readonly policyData: pulumi.Input<string>;
    /**
     * The project in which the cluster belongs. If it
     * is not provided, Terraform will use the provider default.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region in which the cluster belongs. If it
     * is not provided, Terraform will use the provider default.
     */
    readonly region?: pulumi.Input<string>;
}
