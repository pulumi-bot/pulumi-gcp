// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Cloud Endpoints Service. Each of these resources serves a different use case:
 * 
 * * `gcp.endpoints.ServiceIamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
 * * `gcp.endpoints.ServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
 * * `gcp.endpoints.ServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.
 * 
 * > **Note:** `gcp.endpoints.ServiceIamPolicy` **cannot** be used in conjunction with `gcp.endpoints.ServiceIamBinding` and `gcp.endpoints.ServiceIamMember` or they will fight over what your policy should be.
 * 
 * > **Note:** `gcp.endpoints.ServiceIamBinding` resources **can be** used in conjunction with `gcp.endpoints.ServiceIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * 
 * 
 * ## google\_endpoints\_service\_iam\_policy
 * 
 * {{ % example typescript % }}
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
 * const policy = new gcp.endpoints.ServiceIamPolicy("policy", {
 *     serviceName: google_endpoints_service.endpoints_service.service_name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * {{ % /example % }}
 * 
 * ## google\_endpoints\_service\_iam\_binding
 * 
 * {{ % example typescript % }}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const binding = new gcp.endpoints.ServiceIamBinding("binding", {
 *     serviceName: google_endpoints_service.endpoints_service.service_name,
 *     role: "roles/viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 * {{ % /example % }}
 * 
 * ## google\_endpoints\_service\_iam\_member
 * 
 * {{ % example typescript % }}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const member = new gcp.endpoints.ServiceIamMember("member", {
 *     serviceName: google_endpoints_service.endpoints_service.service_name,
 *     role: "roles/viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 * {{ % /example % }}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/endpoints_service_iam.html.markdown.
 */
export class ServiceIamBinding extends pulumi.CustomResource {
    /**
     * Get an existing ServiceIamBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceIamBindingState, opts?: pulumi.CustomResourceOptions): ServiceIamBinding {
        return new ServiceIamBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:endpoints/serviceIamBinding:ServiceIamBinding';

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

    public readonly condition!: pulumi.Output<outputs.endpoints.ServiceIamBindingCondition | undefined>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The role that should be applied. Only one
     * `gcp.endpoints.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The name of the service. Used to find the parent resource to bind the IAM policy to
     */
    public readonly serviceName!: pulumi.Output<string>;

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
        if (opts && opts.id) {
            const state = argsOrState as ServiceIamBindingState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["serviceName"] = state ? state.serviceName : undefined;
        } else {
            const args = argsOrState as ServiceIamBindingArgs | undefined;
            if (!args || args.members === undefined) {
                throw new Error("Missing required property 'members'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ServiceIamBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceIamBinding resources.
 */
export interface ServiceIamBindingState {
    readonly condition?: pulumi.Input<inputs.endpoints.ServiceIamBindingCondition>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role that should be applied. Only one
     * `gcp.endpoints.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * The name of the service. Used to find the parent resource to bind the IAM policy to
     */
    readonly serviceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceIamBinding resource.
 */
export interface ServiceIamBindingArgs {
    readonly condition?: pulumi.Input<inputs.endpoints.ServiceIamBindingCondition>;
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role that should be applied. Only one
     * `gcp.endpoints.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
    /**
     * The name of the service. Used to find the parent resource to bind the IAM policy to
     */
    readonly serviceName: pulumi.Input<string>;
}
