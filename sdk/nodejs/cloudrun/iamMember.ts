// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Cloud Run Service. Each of these resources serves a different use case:
 * 
 * * `gcp.cloudrun.IamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
 * * `gcp.cloudrun.IamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
 * * `gcp.cloudrun.IamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.
 * 
 * > **Note:** `gcp.cloudrun.IamPolicy` **cannot** be used in conjunction with `gcp.cloudrun.IamBinding` and `gcp.cloudrun.IamMember` or they will fight over what your policy should be.
 * 
 * > **Note:** `gcp.cloudrun.IamBinding` resources **can be** used in conjunction with `gcp.cloudrun.IamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * 
 * 
 * ## google\_cloud\_run\_service\_iam\_policy
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
 * const policy = new gcp.cloudrun.IamPolicy("policy", {
 *     location: google_cloud_run_service["default"].location,
 *     project: google_cloud_run_service["default"].project,
 *     service: google_cloud_run_service["default"].name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * {{ % /example % }}
 * 
 * ## google\_cloud\_run\_service\_iam\_binding
 * 
 * {{ % example typescript % }}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const binding = new gcp.cloudrun.IamBinding("binding", {
 *     location: google_cloud_run_service["default"].location,
 *     project: google_cloud_run_service["default"].project,
 *     service: google_cloud_run_service["default"].name,
 *     role: "roles/viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 * {{ % /example % }}
 * 
 * ## google\_cloud\_run\_service\_iam\_member
 * 
 * {{ % example typescript % }}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const member = new gcp.cloudrun.IamMember("member", {
 *     location: google_cloud_run_service["default"].location,
 *     project: google_cloud_run_service["default"].project,
 *     service: google_cloud_run_service["default"].name,
 *     role: "roles/viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 * {{ % /example % }}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloud_run_service_iam.html.markdown.
 */
export class IamMember extends pulumi.CustomResource {
    /**
     * Get an existing IamMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IamMemberState, opts?: pulumi.CustomResourceOptions): IamMember {
        return new IamMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:cloudrun/iamMember:IamMember';

    /**
     * Returns true if the given object is an instance of IamMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IamMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IamMember.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.cloudrun.IamMemberCondition | undefined>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
     */
    public readonly location!: pulumi.Output<string>;
    public readonly member!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly service!: pulumi.Output<string>;

    /**
     * Create a IamMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IamMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IamMemberArgs | IamMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IamMemberState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["service"] = state ? state.service : undefined;
        } else {
            const args = argsOrState as IamMemberArgs | undefined;
            if (!args || args.member === undefined) {
                throw new Error("Missing required property 'member'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            if (!args || args.service === undefined) {
                throw new Error("Missing required property 'service'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["member"] = args ? args.member : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["service"] = args ? args.service : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IamMember.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IamMember resources.
 */
export interface IamMemberState {
    readonly condition?: pulumi.Input<inputs.cloudrun.IamMemberCondition>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
     */
    readonly location?: pulumi.Input<string>;
    readonly member?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly service?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IamMember resource.
 */
export interface IamMemberArgs {
    readonly condition?: pulumi.Input<inputs.cloudrun.IamMemberCondition>;
    /**
     * The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
     */
    readonly location?: pulumi.Input<string>;
    readonly member: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly service: pulumi.Input<string>;
}
