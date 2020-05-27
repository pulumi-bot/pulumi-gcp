// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Binary Authorization Attestor. Each of these resources serves a different use case:
 *
 * * `gcp.binaryauthorization.AttestorIamPolicy`: Authoritative. Sets the IAM policy for the attestor and replaces any existing policy already attached.
 * * `gcp.binaryauthorization.AttestorIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the attestor are preserved.
 * * `gcp.binaryauthorization.AttestorIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the attestor are preserved.
 *
 * > **Note:** `gcp.binaryauthorization.AttestorIamPolicy` **cannot** be used in conjunction with `gcp.binaryauthorization.AttestorIamBinding` and `gcp.binaryauthorization.AttestorIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.binaryauthorization.AttestorIamBinding` resources **can be** used in conjunction with `gcp.binaryauthorization.AttestorIamMember` resources **only if** they do not grant privilege to the same role.
 *
 *
 *
 * ## google\_binary\_authorization\_attestor\_iam\_policy
 * {{% example %}}
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     binding: [{
 *         role: "roles/viewer",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.binaryauthorization.AttestorIamPolicy("policy", {
 *     project: google_binary_authorization_attestor.attestor.project,
 *     attestor: google_binary_authorization_attestor.attestor.name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * {{% /example %}}
 * ## google\_binary\_authorization\_attestor\_iam\_binding
 * {{% example %}}
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.binaryauthorization.AttestorIamBinding("binding", {
 *     project: google_binary_authorization_attestor.attestor.project,
 *     attestor: google_binary_authorization_attestor.attestor.name,
 *     role: "roles/viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * {{% /example %}}
 * ## google\_binary\_authorization\_attestor\_iam\_member
 * {{% example %}}
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.binaryauthorization.AttestorIamMember("member", {
 *     project: google_binary_authorization_attestor.attestor.project,
 *     attestor: google_binary_authorization_attestor.attestor.name,
 *     role: "roles/viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * {{% /example %}}
 */
export class AttestorIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AttestorIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AttestorIamPolicyState, opts?: pulumi.CustomResourceOptions): AttestorIamPolicy {
        return new AttestorIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:binaryauthorization/attestorIamPolicy:AttestorIamPolicy';

    /**
     * Returns true if the given object is an instance of AttestorIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AttestorIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AttestorIamPolicy.__pulumiType;
    }

    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly attestor!: pulumi.Output<string>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    public readonly policyData!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a AttestorIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AttestorIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AttestorIamPolicyArgs | AttestorIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AttestorIamPolicyState | undefined;
            inputs["attestor"] = state ? state.attestor : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as AttestorIamPolicyArgs | undefined;
            if (!args || args.attestor === undefined) {
                throw new Error("Missing required property 'attestor'");
            }
            if (!args || args.policyData === undefined) {
                throw new Error("Missing required property 'policyData'");
            }
            inputs["attestor"] = args ? args.attestor : undefined;
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
        super(AttestorIamPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AttestorIamPolicy resources.
 */
export interface AttestorIamPolicyState {
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly attestor?: pulumi.Input<string>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AttestorIamPolicy resource.
 */
export interface AttestorIamPolicyArgs {
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly attestor: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
