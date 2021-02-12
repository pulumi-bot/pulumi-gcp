// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Service Directory Service. Each of these resources serves a different use case:
 *
 * * `gcp.servicedirectory.ServiceIamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
 * * `gcp.servicedirectory.ServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
 * * `gcp.servicedirectory.ServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.
 *
 * > **Note:** `gcp.servicedirectory.ServiceIamPolicy` **cannot** be used in conjunction with `gcp.servicedirectory.ServiceIamBinding` and `gcp.servicedirectory.ServiceIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.servicedirectory.ServiceIamBinding` resources **can be** used in conjunction with `gcp.servicedirectory.ServiceIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_service\_directory\_service\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/viewer",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.servicedirectory.ServiceIamPolicy("policy", {policyData: admin.then(admin => admin.policyData)});
 * ```
 *
 * ## google\_service\_directory\_service\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.servicedirectory.ServiceIamBinding("binding", {
 *     role: "roles/viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * ## google\_service\_directory\_service\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.servicedirectory.ServiceIamMember("member", {
 *     role: "roles/viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}} * {{project}}/{{location}}/{{namespace_id}}/{{service_id}} * {{location}}/{{namespace_id}}/{{service_id}} Any variables not passed in the import command will be taken from the provider configuration. Service Directory service IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:servicedirectory/serviceIamBinding:ServiceIamBinding editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}} roles/viewer user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:servicedirectory/serviceIamBinding:ServiceIamBinding editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}} roles/viewer"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:servicedirectory/serviceIamBinding:ServiceIamBinding editor projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class ServiceIamBinding extends pulumi.CustomResource {
    /**
     * Get an existing ServiceIamBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceIamBindingState, opts?: pulumi.CustomResourceOptions): ServiceIamBinding {
        return new ServiceIamBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:servicedirectory/serviceIamBinding:ServiceIamBinding';

    /**
     * Returns true if the given object is an instance of ServiceIamBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceIamBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceIamBinding.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.servicedirectory.ServiceIamBindingCondition | undefined>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a ServiceIamBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceIamBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceIamBindingArgs | ServiceIamBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceIamBindingState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as ServiceIamBindingArgs | undefined;
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ServiceIamBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceIamBinding resources.
 */
export interface ServiceIamBindingState {
    readonly condition?: pulumi.Input<inputs.servicedirectory.ServiceIamBindingCondition>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceIamBinding resource.
 */
export interface ServiceIamBindingArgs {
    readonly condition?: pulumi.Input<inputs.servicedirectory.ServiceIamBindingCondition>;
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
