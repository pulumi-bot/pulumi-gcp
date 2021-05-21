// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Identity-Aware Proxy WebBackendService. Each of these resources serves a different use case:
 *
 * * `gcp.iap.WebBackendServiceIamPolicy`: Authoritative. Sets the IAM policy for the webbackendservice and replaces any existing policy already attached.
 * * `gcp.iap.WebBackendServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webbackendservice are preserved.
 * * `gcp.iap.WebBackendServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webbackendservice are preserved.
 *
 * > **Note:** `gcp.iap.WebBackendServiceIamPolicy` **cannot** be used in conjunction with `gcp.iap.WebBackendServiceIamBinding` and `gcp.iap.WebBackendServiceIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.iap.WebBackendServiceIamBinding` resources **can be** used in conjunction with `gcp.iap.WebBackendServiceIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_iap\_web\_backend\_service\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/iap.httpsResourceAccessor",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.iap.WebBackendServiceIamPolicy("policy", {
 *     project: google_compute_backend_service["default"].project,
 *     webBackendService: google_compute_backend_service["default"].name,
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
 *         role: "roles/iap.httpsResourceAccessor",
 *         members: ["user:jane@example.com"],
 *         condition: {
 *             title: "expires_after_2019_12_31",
 *             description: "Expiring at midnight of 2019-12-31",
 *             expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         },
 *     }],
 * });
 * const policy = new gcp.iap.WebBackendServiceIamPolicy("policy", {
 *     project: google_compute_backend_service["default"].project,
 *     webBackendService: google_compute_backend_service["default"].name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * ## google\_iap\_web\_backend\_service\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.iap.WebBackendServiceIamBinding("binding", {
 *     project: google_compute_backend_service["default"].project,
 *     webBackendService: google_compute_backend_service["default"].name,
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
 * const binding = new gcp.iap.WebBackendServiceIamBinding("binding", {
 *     project: google_compute_backend_service["default"].project,
 *     webBackendService: google_compute_backend_service["default"].name,
 *     role: "roles/iap.httpsResourceAccessor",
 *     members: ["user:jane@example.com"],
 *     condition: {
 *         title: "expires_after_2019_12_31",
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *     },
 * });
 * ```
 * ## google\_iap\_web\_backend\_service\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.iap.WebBackendServiceIamMember("member", {
 *     project: google_compute_backend_service["default"].project,
 *     webBackendService: google_compute_backend_service["default"].name,
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
 * const member = new gcp.iap.WebBackendServiceIamMember("member", {
 *     project: google_compute_backend_service["default"].project,
 *     webBackendService: google_compute_backend_service["default"].name,
 *     role: "roles/iap.httpsResourceAccessor",
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
 * For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/compute/services/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy webbackendservice IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy editor "projects/{{project}}/iap_web/compute/services/{{web_backend_service}} roles/iap.httpsResourceAccessor user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy editor "projects/{{project}}/iap_web/compute/services/{{web_backend_service}} roles/iap.httpsResourceAccessor"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy editor projects/{{project}}/iap_web/compute/services/{{web_backend_service}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class WebBackendServiceIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing WebBackendServiceIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebBackendServiceIamPolicyState, opts?: pulumi.CustomResourceOptions): WebBackendServiceIamPolicy {
        return new WebBackendServiceIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy';

    /**
     * Returns true if the given object is an instance of WebBackendServiceIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebBackendServiceIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebBackendServiceIamPolicy.__pulumiType;
    }

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
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly webBackendService!: pulumi.Output<string>;

    /**
     * Create a WebBackendServiceIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebBackendServiceIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebBackendServiceIamPolicyArgs | WebBackendServiceIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebBackendServiceIamPolicyState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["webBackendService"] = state ? state.webBackendService : undefined;
        } else {
            const args = argsOrState as WebBackendServiceIamPolicyArgs | undefined;
            if ((!args || args.policyData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyData'");
            }
            if ((!args || args.webBackendService === undefined) && !opts.urn) {
                throw new Error("Missing required property 'webBackendService'");
            }
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["webBackendService"] = args ? args.webBackendService : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WebBackendServiceIamPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebBackendServiceIamPolicy resources.
 */
export interface WebBackendServiceIamPolicyState {
    /**
     * (Computed) The etag of the IAM policy.
     */
    etag?: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    policyData?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    webBackendService?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebBackendServiceIamPolicy resource.
 */
export interface WebBackendServiceIamPolicyArgs {
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    policyData: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    webBackendService: pulumi.Input<string>;
}
