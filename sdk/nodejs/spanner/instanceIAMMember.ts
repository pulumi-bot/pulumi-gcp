// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for a Spanner instance. Each of these resources serves a different use case:
 * 
 * * `gcp.spanner.InstanceIAMPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
 * 
 * > **Warning:** It's entirely possibly to lock yourself out of your instance using `gcp.spanner.InstanceIAMPolicy`. Any permissions granted by default will be removed unless you include them in your config.
 * 
 * * `gcp.spanner.InstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
 * * `gcp.spanner.InstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
 * 
 * > **Note:** `gcp.spanner.InstanceIAMPolicy` **cannot** be used in conjunction with `gcp.spanner.InstanceIAMBinding` and `gcp.spanner.InstanceIAMMember` or they will fight over what your policy should be.
 * 
 * > **Note:** `gcp.spanner.InstanceIAMBinding` resources **can be** used in conjunction with `gcp.spanner.InstanceIAMMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_spanner\_instance\_iam\_binding
 * 
 * {{% examples %}}
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const instance = new gcp.spanner.InstanceIAMBinding("instance", {
 *     instance: "your-instance-name",
 *     members: ["user:jane@example.com"],
 *     role: "roles/compute.networkUser",
 * });
 * ```
 * 
 * {{% /examples %}}
 * ## google\_spanner\_instance\_iam\_member
 * 
 * {{% examples %}}
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const instance = new gcp.spanner.InstanceIAMMember("instance", {
 *     instance: "your-instance-name",
 *     member: "user:jane@example.com",
 *     role: "roles/compute.networkUser",
 * });
 * ```
 * 
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_instance_iam.html.markdown.
 */
export class InstanceIAMMember extends pulumi.CustomResource {
    /**
     * Get an existing InstanceIAMMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceIAMMemberState, opts?: pulumi.CustomResourceOptions): InstanceIAMMember {
        return new InstanceIAMMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:spanner/instanceIAMMember:InstanceIAMMember';

    /**
     * Returns true if the given object is an instance of InstanceIAMMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceIAMMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceIAMMember.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.spanner.InstanceIAMMemberCondition | undefined>;
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
     * `gcp.spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a InstanceIAMMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceIAMMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceIAMMemberArgs | InstanceIAMMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceIAMMemberState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["instance"] = state ? state.instance : undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as InstanceIAMMemberArgs | undefined;
            if (!args || args.instance === undefined) {
                throw new Error("Missing required property 'instance'");
            }
            if (!args || args.member === undefined) {
                throw new Error("Missing required property 'member'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["instance"] = args ? args.instance : undefined;
            inputs["member"] = args ? args.member : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(InstanceIAMMember.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceIAMMember resources.
 */
export interface InstanceIAMMemberState {
    readonly condition?: pulumi.Input<inputs.spanner.InstanceIAMMemberCondition>;
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
     * `gcp.spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceIAMMember resource.
 */
export interface InstanceIAMMemberArgs {
    readonly condition?: pulumi.Input<inputs.spanner.InstanceIAMMemberCondition>;
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
     * `gcp.spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
