// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Compute Engine Subnetwork. Each of these resources serves a different use case:
 * 
 * * `gcp.compute.SubnetworkIAMPolicy`: Authoritative. Sets the IAM policy for the subnetwork and replaces any existing policy already attached.
 * * `gcp.compute.SubnetworkIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subnetwork are preserved.
 * * `gcp.compute.SubnetworkIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subnetwork are preserved.
 * 
 * > **Note:** `gcp.compute.SubnetworkIAMPolicy` **cannot** be used in conjunction with `gcp.compute.SubnetworkIAMBinding` and `gcp.compute.SubnetworkIAMMember` or they will fight over what your policy should be.
 * 
 * > **Note:** `gcp.compute.SubnetworkIAMBinding` resources **can be** used in conjunction with `gcp.compute.SubnetworkIAMMember` resources **only if** they do not grant privilege to the same role.
 * 
 * 
 * 
 * ## google\_compute\_subnetwork\_iam\_policy
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const admin = gcp.organizations.getIAMPolicy({
 *     binding: [{
 *         role: "roles/compute.networkUser",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.compute.SubnetworkIAMPolicy("policy", {
 *     project: google_compute_subnetwork["network-with-private-secondary-ip-ranges"].project,
 *     region: google_compute_subnetwork["network-with-private-secondary-ip-ranges"].region,
 *     subnetwork: google_compute_subnetwork["network-with-private-secondary-ip-ranges"].name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * 
 * With IAM Conditions:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const admin = gcp.organizations.getIAMPolicy({
 *     binding: [{
 *         role: "roles/compute.networkUser",
 *         members: ["user:jane@example.com"],
 *         condition: {
 *             title: "expiresAfter20191231",
 *             description: "Expiring at midnight of 2019-12-31",
 *             expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         },
 *     }],
 * });
 * const policy = new gcp.compute.SubnetworkIAMPolicy("policy", {
 *     project: google_compute_subnetwork["network-with-private-secondary-ip-ranges"].project,
 *     region: google_compute_subnetwork["network-with-private-secondary-ip-ranges"].region,
 *     subnetwork: google_compute_subnetwork["network-with-private-secondary-ip-ranges"].name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * ## google\_compute\_subnetwork\_iam\_binding
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const binding = new gcp.compute.SubnetworkIAMBinding("binding", {
 *     project: google_compute_subnetwork["network-with-private-secondary-ip-ranges"].project,
 *     region: google_compute_subnetwork["network-with-private-secondary-ip-ranges"].region,
 *     subnetwork: google_compute_subnetwork["network-with-private-secondary-ip-ranges"].name,
 *     role: "roles/compute.networkUser",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 * 
 * With IAM Conditions:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const binding = new gcp.compute.SubnetworkIAMBinding("binding", {
 *     project: google_compute_subnetwork["network-with-private-secondary-ip-ranges"].project,
 *     region: google_compute_subnetwork["network-with-private-secondary-ip-ranges"].region,
 *     subnetwork: google_compute_subnetwork["network-with-private-secondary-ip-ranges"].name,
 *     role: "roles/compute.networkUser",
 *     members: ["user:jane@example.com"],
 *     condition: {
 *         title: "expiresAfter20191231",
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *     },
 * });
 * ```
 * ## google\_compute\_subnetwork\_iam\_member
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const member = new gcp.compute.SubnetworkIAMMember("member", {
 *     project: google_compute_subnetwork["network-with-private-secondary-ip-ranges"].project,
 *     region: google_compute_subnetwork["network-with-private-secondary-ip-ranges"].region,
 *     subnetwork: google_compute_subnetwork["network-with-private-secondary-ip-ranges"].name,
 *     role: "roles/compute.networkUser",
 *     member: "user:jane@example.com",
 * });
 * ```
 * 
 * With IAM Conditions:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const member = new gcp.compute.SubnetworkIAMMember("member", {
 *     project: google_compute_subnetwork["network-with-private-secondary-ip-ranges"].project,
 *     region: google_compute_subnetwork["network-with-private-secondary-ip-ranges"].region,
 *     subnetwork: google_compute_subnetwork["network-with-private-secondary-ip-ranges"].name,
 *     role: "roles/compute.networkUser",
 *     member: "user:jane@example.com",
 *     condition: {
 *         title: "expiresAfter20191231",
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *     },
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_subnetwork_iam.html.markdown.
 */
export class SubnetworkIAMMember extends pulumi.CustomResource {
    /**
     * Get an existing SubnetworkIAMMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubnetworkIAMMemberState, opts?: pulumi.CustomResourceOptions): SubnetworkIAMMember {
        return new SubnetworkIAMMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/subnetworkIAMMember:SubnetworkIAMMember';

    /**
     * Returns true if the given object is an instance of SubnetworkIAMMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SubnetworkIAMMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SubnetworkIAMMember.__pulumiType;
    }

    /**
     * ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    public readonly condition!: pulumi.Output<outputs.compute.SubnetworkIAMMemberCondition | undefined>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly member!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * URL of the GCP region for this subnetwork.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
     * region is specified, it is taken from the provider configuration.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly subnetwork!: pulumi.Output<string>;

    /**
     * Create a SubnetworkIAMMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetworkIAMMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubnetworkIAMMemberArgs | SubnetworkIAMMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SubnetworkIAMMemberState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["subnetwork"] = state ? state.subnetwork : undefined;
        } else {
            const args = argsOrState as SubnetworkIAMMemberArgs | undefined;
            if (!args || args.member === undefined) {
                throw new Error("Missing required property 'member'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            if (!args || args.subnetwork === undefined) {
                throw new Error("Missing required property 'subnetwork'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["member"] = args ? args.member : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["subnetwork"] = args ? args.subnetwork : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SubnetworkIAMMember.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SubnetworkIAMMember resources.
 */
export interface SubnetworkIAMMemberState {
    /**
     * ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    readonly condition?: pulumi.Input<inputs.compute.SubnetworkIAMMemberCondition>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly member?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * URL of the GCP region for this subnetwork.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
     * region is specified, it is taken from the provider configuration.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly subnetwork?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SubnetworkIAMMember resource.
 */
export interface SubnetworkIAMMemberArgs {
    /**
     * ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    readonly condition?: pulumi.Input<inputs.compute.SubnetworkIAMMemberCondition>;
    readonly member: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * URL of the GCP region for this subnetwork.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
     * region is specified, it is taken from the provider configuration.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly subnetwork: pulumi.Input<string>;
}
