// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > **Warning:** These resources are in beta, and should be used with the terraform-provider-google-beta provider.
 * See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta resources.
 * 
 * Three different resources help you manage your IAM policy for IAP Tunnel Instance. Each of these resources serves a different use case:
 * 
 * * `google_iap_tunnel_instance_iam_policy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
 * * `google_iap_tunnel_instance_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
 * * `google_iap_tunnel_instance_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
 * 
 * > **Note:** `google_iap_tunnel_instance_iam_policy` **cannot** be used in conjunction with `google_iap_tunnel_instance_iam_binding` and `google_iap_tunnel_instance_iam_member` or they will fight over what your policy should be.
 * 
 * > **Note:** `google_iap_tunnel_instance_iam_binding` resources **can be** used in conjunction with `google_iap_tunnel_instance_iam_member` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_iap\_tunnel\_instance\_iam\_policy
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
 * const instance = new gcp.iap.TunnelInstanceIAMPolicy("instance", {
 *     instance: "your-instance-name",
 *     policyData: admin.policyData,
 * });
 * ```
 * 
 * ## google\_iap\_tunnel\_instance\_iam\_binding
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const instance = new gcp.iap.TunnelInstanceIAMBinding("instance", {
 *     instance: "your-instance-name",
 *     members: ["user:jane@example.com"],
 *     role: "roles/compute.networkUser",
 * });
 * ```
 * 
 * ## google\_iap\_tunnel\_instance\_iam\_member
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const instance = new gcp.iap.TunnelInstanceIAMMember("instance", {
 *     instance: "your-instance-name",
 *     member: "user:jane@example.com",
 *     role: "roles/compute.networkUser",
 * });
 * ```
 */
export class TunnelInstanceIAMMember extends pulumi.CustomResource {
    /**
     * Get an existing TunnelInstanceIAMMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TunnelInstanceIAMMemberState, opts?: pulumi.CustomResourceOptions): TunnelInstanceIAMMember {
        return new TunnelInstanceIAMMember(name, <any>state, { ...opts, id: id });
    }

    private static readonly __pulumiType = 'gcp:iap/tunnelInstanceIAMMember:TunnelInstanceIAMMember';

    /**
     * Returns true if the given object is an instance of TunnelInstanceIAMMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TunnelInstanceIAMMember {
        if (obj === undefined || obj === null) {
            return false;
        }

        const t = obj['__pulumiType'];
        if (typeof t !== 'string') {
            return false;
        }

        return t === TunnelInstanceIAMMember.__pulumiType;
    }

    /**
     * (Computed) The etag of the instance's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the instance.
     */
    public readonly instance!: pulumi.Output<string>;
    public readonly member!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `google_iap_tunnel_instance_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The zone of the instance. If
     * unspecified, this defaults to the zone configured in the provider.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a TunnelInstanceIAMMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TunnelInstanceIAMMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TunnelInstanceIAMMemberArgs | TunnelInstanceIAMMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TunnelInstanceIAMMemberState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["instance"] = state ? state.instance : undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as TunnelInstanceIAMMemberArgs | undefined;
            if (!args || args.instance === undefined) {
                throw new Error("Missing required property 'instance'");
            }
            if (!args || args.member === undefined) {
                throw new Error("Missing required property 'member'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["instance"] = args ? args.instance : undefined;
            inputs["member"] = args ? args.member : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        super(TunnelInstanceIAMMember.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TunnelInstanceIAMMember resources.
 */
export interface TunnelInstanceIAMMemberState {
    /**
     * (Computed) The etag of the instance's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The name of the instance.
     */
    readonly instance?: pulumi.Input<string>;
    readonly member?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `google_iap_tunnel_instance_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * The zone of the instance. If
     * unspecified, this defaults to the zone configured in the provider.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TunnelInstanceIAMMember resource.
 */
export interface TunnelInstanceIAMMemberArgs {
    /**
     * The name of the instance.
     */
    readonly instance: pulumi.Input<string>;
    readonly member: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `google_iap_tunnel_instance_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
    /**
     * The zone of the instance. If
     * unspecified, this defaults to the zone configured in the provider.
     */
    readonly zone?: pulumi.Input<string>;
}
