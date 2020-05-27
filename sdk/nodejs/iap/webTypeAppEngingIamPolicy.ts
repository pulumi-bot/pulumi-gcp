// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Identity-Aware Proxy WebTypeAppEngine. Each of these resources serves a different use case:
 *
 * * `gcp.iap.WebTypeAppEngingIamPolicy`: Authoritative. Sets the IAM policy for the webtypeappengine and replaces any existing policy already attached.
 * * `gcp.iap.WebTypeAppEngingIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webtypeappengine are preserved.
 * * `gcp.iap.WebTypeAppEngingIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webtypeappengine are preserved.
 *
 * > **Note:** `gcp.iap.WebTypeAppEngingIamPolicy` **cannot** be used in conjunction with `gcp.iap.WebTypeAppEngingIamBinding` and `gcp.iap.WebTypeAppEngingIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.iap.WebTypeAppEngingIamBinding` resources **can be** used in conjunction with `gcp.iap.WebTypeAppEngingIamMember` resources **only if** they do not grant privilege to the same role.
 *
 *
 *
 * ## google\_iap\_web\_type\_app\_engine\_iam\_policy
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
 * const policy = new gcp.iap.WebTypeAppEngingIamPolicy("policy", {
 *     project: google_app_engine_application.app.project,
 *     appId: google_app_engine_application.app.app_id,
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
 * const policy = new gcp.iap.WebTypeAppEngingIamPolicy("policy", {
 *     project: google_app_engine_application.app.project,
 *     appId: google_app_engine_application.app.app_id,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * {{% /example %}}
 * ## google\_iap\_web\_type\_app\_engine\_iam\_binding
 * {{% example %}}
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.iap.WebTypeAppEngingIamBinding("binding", {
 *     project: google_app_engine_application.app.project,
 *     appId: google_app_engine_application.app.app_id,
 *     role: "roles/iap.httpsResourceAccessor",
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
 * const binding = new gcp.iap.WebTypeAppEngingIamBinding("binding", {
 *     project: google_app_engine_application.app.project,
 *     appId: google_app_engine_application.app.app_id,
 *     role: "roles/iap.httpsResourceAccessor",
 *     members: ["user:jane@example.com"],
 *     condition: {
 *         title: "expiresAfter20191231",
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *     },
 * });
 * ```
 * {{% /example %}}
 * ## google\_iap\_web\_type\_app\_engine\_iam\_member
 * {{% example %}}
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.iap.WebTypeAppEngingIamMember("member", {
 *     project: google_app_engine_application.app.project,
 *     appId: google_app_engine_application.app.app_id,
 *     role: "roles/iap.httpsResourceAccessor",
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
 * const member = new gcp.iap.WebTypeAppEngingIamMember("member", {
 *     project: google_app_engine_application.app.project,
 *     appId: google_app_engine_application.app.app_id,
 *     role: "roles/iap.httpsResourceAccessor",
 *     member: "user:jane@example.com",
 *     condition: {
 *         title: "expiresAfter20191231",
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *     },
 * });
 * ```
 * {{% /example %}}
 */
export class WebTypeAppEngingIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing WebTypeAppEngingIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebTypeAppEngingIamPolicyState, opts?: pulumi.CustomResourceOptions): WebTypeAppEngingIamPolicy {
        return new WebTypeAppEngingIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:iap/webTypeAppEngingIamPolicy:WebTypeAppEngingIamPolicy';

    /**
     * Returns true if the given object is an instance of WebTypeAppEngingIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebTypeAppEngingIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebTypeAppEngingIamPolicy.__pulumiType;
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
     * Create a WebTypeAppEngingIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebTypeAppEngingIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebTypeAppEngingIamPolicyArgs | WebTypeAppEngingIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as WebTypeAppEngingIamPolicyState | undefined;
            inputs["appId"] = state ? state.appId : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as WebTypeAppEngingIamPolicyArgs | undefined;
            if (!args || args.appId === undefined) {
                throw new Error("Missing required property 'appId'");
            }
            if (!args || args.policyData === undefined) {
                throw new Error("Missing required property 'policyData'");
            }
            inputs["appId"] = args ? args.appId : undefined;
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
        super(WebTypeAppEngingIamPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebTypeAppEngingIamPolicy resources.
 */
export interface WebTypeAppEngingIamPolicyState {
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
}

/**
 * The set of arguments for constructing a WebTypeAppEngingIamPolicy resource.
 */
export interface WebTypeAppEngingIamPolicyArgs {
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
}
