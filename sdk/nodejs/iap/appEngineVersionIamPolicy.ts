// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Identity-Aware Proxy AppEngineVersion. Each of these resources serves a different use case:
 *
 * * `gcp.iap.AppEngineVersionIamPolicy`: Authoritative. Sets the IAM policy for the appengineversion and replaces any existing policy already attached.
 * * `gcp.iap.AppEngineVersionIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the appengineversion are preserved.
 * * `gcp.iap.AppEngineVersionIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the appengineversion are preserved.
 *
 * > **Note:** `gcp.iap.AppEngineVersionIamPolicy` **cannot** be used in conjunction with `gcp.iap.AppEngineVersionIamBinding` and `gcp.iap.AppEngineVersionIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.iap.AppEngineVersionIamBinding` resources **can be** used in conjunction with `gcp.iap.AppEngineVersionIamMember` resources **only if** they do not grant privilege to the same role.
 *
 *
 *
 * ## google\_iap\_app\_engine\_version\_iam\_policy
 * {{% example %}}
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
 * const policy = new gcp.iap.AppEngineVersionIamPolicy("policy", {
 *     project: google_app_engine_standard_app_version.version.project,
 *     appId: google_app_engine_standard_app_version.version.project,
 *     service: google_app_engine_standard_app_version.version.service,
 *     versionId: google_app_engine_standard_app_version.version.version_id,
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
 * const policy = new gcp.iap.AppEngineVersionIamPolicy("policy", {
 *     project: google_app_engine_standard_app_version.version.project,
 *     appId: google_app_engine_standard_app_version.version.project,
 *     service: google_app_engine_standard_app_version.version.service,
 *     versionId: google_app_engine_standard_app_version.version.version_id,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * {{% /example %}}
 * ## google\_iap\_app\_engine\_version\_iam\_binding
 * {{% example %}}
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.iap.AppEngineVersionIamBinding("binding", {
 *     appId: google_app_engine_standard_app_version_version.project,
 *     members: ["user:jane@example.com"],
 *     project: google_app_engine_standard_app_version_version.project,
 *     role: "roles/iap.httpsResourceAccessor",
 *     service: google_app_engine_standard_app_version_version.service,
 *     versionId: google_app_engine_standard_app_version_version.versionId,
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.iap.AppEngineVersionIamBinding("binding", {
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
 *     versionId: google_app_engine_standard_app_version_version.versionId,
 * });
 * ```
 * {{% /example %}}
 * ## google\_iap\_app\_engine\_version\_iam\_member
 * {{% example %}}
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.iap.AppEngineVersionIamMember("member", {
 *     appId: google_app_engine_standard_app_version_version.project,
 *     member: "user:jane@example.com",
 *     project: google_app_engine_standard_app_version_version.project,
 *     role: "roles/iap.httpsResourceAccessor",
 *     service: google_app_engine_standard_app_version_version.service,
 *     versionId: google_app_engine_standard_app_version_version.versionId,
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.iap.AppEngineVersionIamMember("member", {
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
 *     versionId: google_app_engine_standard_app_version_version.versionId,
 * });
 * ```
 * {{% /example %}}
 */
export class AppEngineVersionIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AppEngineVersionIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppEngineVersionIamPolicyState, opts?: pulumi.CustomResourceOptions): AppEngineVersionIamPolicy {
        return new AppEngineVersionIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:iap/appEngineVersionIamPolicy:AppEngineVersionIamPolicy';

    /**
     * Returns true if the given object is an instance of AppEngineVersionIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppEngineVersionIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppEngineVersionIamPolicy.__pulumiType;
    }

    /**
     * Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
     */
    public readonly appId!: pulumi.Output<string>;
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
     * Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
     */
    public readonly service!: pulumi.Output<string>;
    /**
     * Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
     */
    public readonly versionId!: pulumi.Output<string>;

    /**
     * Create a AppEngineVersionIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppEngineVersionIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppEngineVersionIamPolicyArgs | AppEngineVersionIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AppEngineVersionIamPolicyState | undefined;
            inputs["appId"] = state ? state.appId : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["service"] = state ? state.service : undefined;
            inputs["versionId"] = state ? state.versionId : undefined;
        } else {
            const args = argsOrState as AppEngineVersionIamPolicyArgs | undefined;
            if (!args || args.appId === undefined) {
                throw new Error("Missing required property 'appId'");
            }
            if (!args || args.policyData === undefined) {
                throw new Error("Missing required property 'policyData'");
            }
            if (!args || args.service === undefined) {
                throw new Error("Missing required property 'service'");
            }
            if (!args || args.versionId === undefined) {
                throw new Error("Missing required property 'versionId'");
            }
            inputs["appId"] = args ? args.appId : undefined;
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["service"] = args ? args.service : undefined;
            inputs["versionId"] = args ? args.versionId : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AppEngineVersionIamPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppEngineVersionIamPolicy resources.
 */
export interface AppEngineVersionIamPolicyState {
    /**
     * Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
     */
    readonly appId?: pulumi.Input<string>;
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
    /**
     * Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
     */
    readonly service?: pulumi.Input<string>;
    /**
     * Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
     */
    readonly versionId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppEngineVersionIamPolicy resource.
 */
export interface AppEngineVersionIamPolicyArgs {
    /**
     * Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
     */
    readonly appId: pulumi.Input<string>;
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
    /**
     * Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
     */
    readonly service: pulumi.Input<string>;
    /**
     * Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
     */
    readonly versionId: pulumi.Input<string>;
}
