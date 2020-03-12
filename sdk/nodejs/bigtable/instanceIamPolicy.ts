// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage IAM policies on bigtable instances. Each of these resources serves a different use case:
 * 
 * * `gcp.bigtable.InstanceIamPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
 * * `gcp.bigtable.InstanceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
 * * `gcp.bigtable.InstanceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
 * 
 * > **Note:** `gcp.bigtable.InstanceIamPolicy` **cannot** be used in conjunction with `gcp.bigtable.InstanceIamBinding` and `gcp.bigtable.InstanceIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the instance as `gcp.bigtable.InstanceIamPolicy` replaces the entire policy.
 * 
 * > **Note:** `gcp.bigtable.InstanceIamBinding` resources **can be** used in conjunction with `gcp.bigtable.InstanceIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_bigtable\_instance\_iam\_binding
 * 
 * {{% examples %}}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const editor = new gcp.bigtable.InstanceIamBinding("editor", {
 *     instance: "your-bigtable-instance",
 *     members: ["user:jane@example.com"],
 *     role: "roles/editor",
 * });
 * ```
 * {{% /examples %}}
 * 
 * ## google\_bigtable\_instance\_iam\_member
 * 
 * {{% examples %}}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const editor = new gcp.bigtable.InstanceIamMember("editor", {
 *     instance: "your-bigtable-instance",
 *     member: "user:jane@example.com",
 *     role: "roles/editor",
 * });
 * ```
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/bigtable_instance_iam.html.markdown.
 */
export class InstanceIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing InstanceIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceIamPolicyState, opts?: pulumi.CustomResourceOptions): InstanceIamPolicy {
        return new InstanceIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:bigtable/instanceIamPolicy:InstanceIamPolicy';

    /**
     * Returns true if the given object is an instance of InstanceIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceIamPolicy.__pulumiType;
    }

    /**
     * (Computed) The etag of the instances's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name or relative resource id of the instance to manage IAM policies for.
     */
    public readonly instance!: pulumi.Output<string>;
    /**
     * The policy data generated by a `gcp.organizations.getIAMPolicy` data source.
     */
    public readonly policyData!: pulumi.Output<string>;
    /**
     * The project in which the instance belongs. If it
     * is not provided, this provider will use the provider default.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a InstanceIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceIamPolicyArgs | InstanceIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceIamPolicyState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["instance"] = state ? state.instance : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as InstanceIamPolicyArgs | undefined;
            if (!args || args.instance === undefined) {
                throw new Error("Missing required property 'instance'");
            }
            if (!args || args.policyData === undefined) {
                throw new Error("Missing required property 'policyData'");
            }
            inputs["instance"] = args ? args.instance : undefined;
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(InstanceIamPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceIamPolicy resources.
 */
export interface InstanceIamPolicyState {
    /**
     * (Computed) The etag of the instances's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The name or relative resource id of the instance to manage IAM policies for.
     */
    readonly instance?: pulumi.Input<string>;
    /**
     * The policy data generated by a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData?: pulumi.Input<string>;
    /**
     * The project in which the instance belongs. If it
     * is not provided, this provider will use the provider default.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceIamPolicy resource.
 */
export interface InstanceIamPolicyArgs {
    /**
     * The name or relative resource id of the instance to manage IAM policies for.
     */
    readonly instance: pulumi.Input<string>;
    /**
     * The policy data generated by a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: pulumi.Input<string>;
    /**
     * The project in which the instance belongs. If it
     * is not provided, this provider will use the provider default.
     */
    readonly project?: pulumi.Input<string>;
}
