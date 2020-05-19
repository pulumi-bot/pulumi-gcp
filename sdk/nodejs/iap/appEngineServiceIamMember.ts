// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Identity-Aware Proxy AppEngineService. Each of these resources serves a different use case:
 *
 * * `gcp.iap.AppEngineServiceIamPolicy`: Authoritative. Sets the IAM policy for the appengineservice and replaces any existing policy already attached.
 * * `gcp.iap.AppEngineServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the appengineservice are preserved.
 * * `gcp.iap.AppEngineServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the appengineservice are preserved.
 *
 * > **Note:** `gcp.iap.AppEngineServiceIamPolicy` **cannot** be used in conjunction with `gcp.iap.AppEngineServiceIamBinding` and `gcp.iap.AppEngineServiceIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.iap.AppEngineServiceIamBinding` resources **can be** used in conjunction with `gcp.iap.AppEngineServiceIamMember` resources **only if** they do not grant privilege to the same role.
 *
 *
 *
 * ## google\_iap\_app\_engine\_service\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     binding: [{
 *         role: "roles/iap.httpsResourceAccessor",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.iap.AppEngineServiceIamPolicy("policy", {
 *     project: google_app_engine_standard_app_version.version.project,
 *     appId: google_app_engine_standard_app_version.version.project,
 *     service: google_app_engine_standard_app_version.version.service,
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
 *         role: "roles/iap.httpsResourceAccessor",
 *         members: ["user:jane@example.com"],
 *         condition: {
 *             title: "expiresAfter20191231",
 *             description: "Expiring at midnight of 2019-12-31",
 *             expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         },
 *     }],
 * });
 * const policy = new gcp.iap.AppEngineServiceIamPolicy("policy", {
 *     project: google_app_engine_standard_app_version.version.project,
 *     appId: google_app_engine_standard_app_version.version.project,
 *     service: google_app_engine_standard_app_version.version.service,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * ## google\_iap\_app\_engine\_service\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.iap.AppEngineServiceIamBinding("binding", {
 *     appId: google_app_engine_standard_app_version_version.project,
 *     members: ["user:jane@example.com"],
 *     project: google_app_engine_standard_app_version_version.project,
 *     role: "roles/iap.httpsResourceAccessor",
 *     service: google_app_engine_standard_app_version_version.service,
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.iap.AppEngineServiceIamBinding("binding", {
 *     appId: google_app_engine_standard_app_version_version.project,
 *     condition: {
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         title: "expiresAfter20191231",
 *     },
 *     members: ["user:jane@example.com"],
 *     project: google_app_engine_standard_app_version_version.project,
 *     role: "roles/iap.httpsResourceAccessor",
 *     service: google_app_engine_standard_app_version_version.service,
 * });
 * ```
 * ## google\_iap\_app\_engine\_service\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.iap.AppEngineServiceIamMember("member", {
 *     appId: google_app_engine_standard_app_version_version.project,
 *     member: "user:jane@example.com",
 *     project: google_app_engine_standard_app_version_version.project,
 *     role: "roles/iap.httpsResourceAccessor",
 *     service: google_app_engine_standard_app_version_version.service,
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.iap.AppEngineServiceIamMember("member", {
 *     appId: google_app_engine_standard_app_version_version.project,
 *     condition: {
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         title: "expiresAfter20191231",
 *     },
 *     member: "user:jane@example.com",
 *     project: google_app_engine_standard_app_version_version.project,
 *     role: "roles/iap.httpsResourceAccessor",
 *     service: google_app_engine_standard_app_version_version.service,
 * });
 * ```
 */
export class AppEngineServiceIamMember extends pulumi.CustomResource {
    /**
     * Get an existing AppEngineServiceIamMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppEngineServiceIamMemberState, opts?: pulumi.CustomResourceOptions): AppEngineServiceIamMember {
        return new AppEngineServiceIamMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:iap/appEngineServiceIamMember:AppEngineServiceIamMember';

    /**
     * Returns true if the given object is an instance of AppEngineServiceIamMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppEngineServiceIamMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppEngineServiceIamMember.__pulumiType;
    }

    /**
     * Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    public readonly condition!: pulumi.Output<outputs.iap.AppEngineServiceIamMemberCondition | undefined>;
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
     * The role that should be applied. Only one
     * `gcp.iap.AppEngineServiceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
     */
    public readonly service!: pulumi.Output<string>;

    /**
     * Create a AppEngineServiceIamMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppEngineServiceIamMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppEngineServiceIamMemberArgs | AppEngineServiceIamMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AppEngineServiceIamMemberState | undefined;
            inputs["appId"] = state ? state.appId : undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["service"] = state ? state.service : undefined;
        } else {
            const args = argsOrState as AppEngineServiceIamMemberArgs | undefined;
            if (!args || args.appId === undefined) {
                throw new Error("Missing required property 'appId'");
            }
            if (!args || args.member === undefined) {
                throw new Error("Missing required property 'member'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            if (!args || args.service === undefined) {
                throw new Error("Missing required property 'service'");
            }
            inputs["appId"] = args ? args.appId : undefined;
            inputs["condition"] = args ? args.condition : undefined;
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
        super(AppEngineServiceIamMember.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppEngineServiceIamMember resources.
 */
export interface AppEngineServiceIamMemberState {
    /**
     * Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
     */
    readonly appId?: pulumi.Input<string>;
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    readonly condition?: pulumi.Input<inputs.iap.AppEngineServiceIamMemberCondition>;
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
     * The role that should be applied. Only one
     * `gcp.iap.AppEngineServiceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
     */
    readonly service?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppEngineServiceIamMember resource.
 */
export interface AppEngineServiceIamMemberArgs {
    /**
     * Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
     */
    readonly appId: pulumi.Input<string>;
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    readonly condition?: pulumi.Input<inputs.iap.AppEngineServiceIamMemberCondition>;
    readonly member: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.iap.AppEngineServiceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
    /**
     * Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
     */
    readonly service: pulumi.Input<string>;
}
