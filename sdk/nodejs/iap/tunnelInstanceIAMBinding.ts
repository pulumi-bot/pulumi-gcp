// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Identity-Aware Proxy TunnelInstance. Each of these resources serves a different use case:
 *
 * * `gcp.iap.TunnelInstanceIAMPolicy`: Authoritative. Sets the IAM policy for the tunnelinstance and replaces any existing policy already attached.
 * * `gcp.iap.TunnelInstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the tunnelinstance are preserved.
 * * `gcp.iap.TunnelInstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the tunnelinstance are preserved.
 *
 * > **Note:** `gcp.iap.TunnelInstanceIAMPolicy` **cannot** be used in conjunction with `gcp.iap.TunnelInstanceIAMBinding` and `gcp.iap.TunnelInstanceIAMMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.iap.TunnelInstanceIAMBinding` resources **can be** used in conjunction with `gcp.iap.TunnelInstanceIAMMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_iap\_tunnel\_instance\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/iap.tunnelResourceAccessor",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.iap.TunnelInstanceIAMPolicy("policy", {
 *     project: google_compute_instance.tunnelvm.project,
 *     zone: google_compute_instance.tunnelvm.zone,
 *     instance: google_compute_instance.tunnelvm.name,
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
 *     bindings: [{
 *         role: "roles/iap.tunnelResourceAccessor",
 *         members: ["user:jane@example.com"],
 *         condition: {
 *             title: "expires_after_2019_12_31",
 *             description: "Expiring at midnight of 2019-12-31",
 *             expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         },
 *     }],
 * });
 * const policy = new gcp.iap.TunnelInstanceIAMPolicy("policy", {
 *     project: google_compute_instance.tunnelvm.project,
 *     zone: google_compute_instance.tunnelvm.zone,
 *     instance: google_compute_instance.tunnelvm.name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * ## google\_iap\_tunnel\_instance\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.iap.TunnelInstanceIAMBinding("binding", {
 *     project: google_compute_instance.tunnelvm.project,
 *     zone: google_compute_instance.tunnelvm.zone,
 *     instance: google_compute_instance.tunnelvm.name,
 *     role: "roles/iap.tunnelResourceAccessor",
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
 * const binding = new gcp.iap.TunnelInstanceIAMBinding("binding", {
 *     project: google_compute_instance.tunnelvm.project,
 *     zone: google_compute_instance.tunnelvm.zone,
 *     instance: google_compute_instance.tunnelvm.name,
 *     role: "roles/iap.tunnelResourceAccessor",
 *     members: ["user:jane@example.com"],
 *     condition: {
 *         title: "expires_after_2019_12_31",
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *     },
 * });
 * ```
 * ## google\_iap\_tunnel\_instance\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.iap.TunnelInstanceIAMMember("member", {
 *     project: google_compute_instance.tunnelvm.project,
 *     zone: google_compute_instance.tunnelvm.zone,
 *     instance: google_compute_instance.tunnelvm.name,
 *     role: "roles/iap.tunnelResourceAccessor",
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
 * const member = new gcp.iap.TunnelInstanceIAMMember("member", {
 *     project: google_compute_instance.tunnelvm.project,
 *     zone: google_compute_instance.tunnelvm.zone,
 *     instance: google_compute_instance.tunnelvm.name,
 *     role: "roles/iap.tunnelResourceAccessor",
 *     member: "user:jane@example.com",
 *     condition: {
 *         title: "expires_after_2019_12_31",
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{name}} * projects/{{project}}/zones/{{zone}}/instances/{{name}} * {{project}}/{{zone}}/{{name}} * {{zone}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy tunnelinstance IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:iap/tunnelInstanceIAMBinding:TunnelInstanceIAMBinding editor "projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{tunnel_instance}} roles/iap.tunnelResourceAccessor user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:iap/tunnelInstanceIAMBinding:TunnelInstanceIAMBinding editor "projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{tunnel_instance}} roles/iap.tunnelResourceAccessor"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:iap/tunnelInstanceIAMBinding:TunnelInstanceIAMBinding editor projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{tunnel_instance}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class TunnelInstanceIAMBinding extends pulumi.CustomResource {
    /**
     * Get an existing TunnelInstanceIAMBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TunnelInstanceIAMBindingState, opts?: pulumi.CustomResourceOptions): TunnelInstanceIAMBinding {
        return new TunnelInstanceIAMBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:iap/tunnelInstanceIAMBinding:TunnelInstanceIAMBinding';

    /**
     * Returns true if the given object is an instance of TunnelInstanceIAMBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TunnelInstanceIAMBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TunnelInstanceIAMBinding.__pulumiType;
    }

    /**
     * ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    public readonly condition!: pulumi.Output<outputs.iap.TunnelInstanceIAMBindingCondition | undefined>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly instance!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a TunnelInstanceIAMBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TunnelInstanceIAMBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TunnelInstanceIAMBindingArgs | TunnelInstanceIAMBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TunnelInstanceIAMBindingState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["instance"] = state ? state.instance : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as TunnelInstanceIAMBindingArgs | undefined;
            if ((!args || args.instance === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'instance'");
            }
            if ((!args || args.members === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'members'");
            }
            if ((!args || args.role === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'role'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["instance"] = args ? args.instance : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TunnelInstanceIAMBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TunnelInstanceIAMBinding resources.
 */
export interface TunnelInstanceIAMBindingState {
    /**
     * ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    readonly condition?: pulumi.Input<inputs.iap.TunnelInstanceIAMBindingCondition>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly instance?: pulumi.Input<string>;
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TunnelInstanceIAMBinding resource.
 */
export interface TunnelInstanceIAMBindingArgs {
    /**
     * ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    readonly condition?: pulumi.Input<inputs.iap.TunnelInstanceIAMBindingCondition>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly instance: pulumi.Input<string>;
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
    readonly zone?: pulumi.Input<string>;
}
