// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows management of the entire IAM policy for an existing Google Cloud Platform Organization.
 *
 * !> **Warning:** New organizations have several default policies which will,
 *    without extreme caution, be **overwritten** by use of this resource.
 *    The safest alternative is to use multiple `gcp.organizations.IAMBinding`
 *    resources.  It is easy to use this resource to remove your own access to
 *    an organization, which will require a call to Google Support to have
 *    fixed, and can take multiple days to resolve.  If you do use this resource,
 *    the best way to be sure that you are not making dangerous changes is to start
 *    by importing your existing policy, and examining the diff very closely.
 *
 * > **Note:** This resource __must not__ be used in conjunction with
 *    `gcp.organizations.IAMMember` or `gcp.organizations.IAMBinding`
 *    or they will fight over what your policy should be.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     binding: [{
 *         role: "roles/editor",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.organizations.IAMPolicy("policy", {
 *     orgId: "123456789",
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 */
export class IAMPolicy extends pulumi.CustomResource {
    /**
     * Get an existing IAMPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IAMPolicyState, opts?: pulumi.CustomResourceOptions): IAMPolicy {
        return new IAMPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:organizations/iAMPolicy:IAMPolicy';

    /**
     * Returns true if the given object is an instance of IAMPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IAMPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IAMPolicy.__pulumiType;
    }

    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The numeric ID of the organization in which you want to create a custom role.
     */
    public readonly orgId!: pulumi.Output<string>;
    /**
     * The `gcp.organizations.getIAMPolicy` data source that represents
     * the IAM policy that will be applied to the organization. This policy overrides any existing
     * policy applied to the organization.
     */
    public readonly policyData!: pulumi.Output<string>;

    /**
     * Create a IAMPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IAMPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IAMPolicyArgs | IAMPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IAMPolicyState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["orgId"] = state ? state.orgId : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
        } else {
            const args = argsOrState as IAMPolicyArgs | undefined;
            if (!args || args.orgId === undefined) {
                throw new Error("Missing required property 'orgId'");
            }
            if (!args || args.policyData === undefined) {
                throw new Error("Missing required property 'policyData'");
            }
            inputs["orgId"] = args ? args.orgId : undefined;
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IAMPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IAMPolicy resources.
 */
export interface IAMPolicyState {
    readonly etag?: pulumi.Input<string>;
    /**
     * The numeric ID of the organization in which you want to create a custom role.
     */
    readonly orgId?: pulumi.Input<string>;
    /**
     * The `gcp.organizations.getIAMPolicy` data source that represents
     * the IAM policy that will be applied to the organization. This policy overrides any existing
     * policy applied to the organization.
     */
    readonly policyData?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IAMPolicy resource.
 */
export interface IAMPolicyArgs {
    /**
     * The numeric ID of the organization in which you want to create a custom role.
     */
    readonly orgId: pulumi.Input<string>;
    /**
     * The `gcp.organizations.getIAMPolicy` data source that represents
     * the IAM policy that will be applied to the organization. This policy overrides any existing
     * policy applied to the organization.
     */
    readonly policyData: pulumi.Input<string>;
}
