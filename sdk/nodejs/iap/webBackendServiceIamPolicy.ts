// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
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
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/iap_web_backend_service_iam.html.markdown.
 */
export class WebBackendServiceIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing WebBackendServiceIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
        if (opts && opts.id) {
            const state = argsOrState as WebBackendServiceIamPolicyState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["webBackendService"] = state ? state.webBackendService : undefined;
        } else {
            const args = argsOrState as WebBackendServiceIamPolicyArgs | undefined;
            if (!args || args.policyData === undefined) {
                throw new Error("Missing required property 'policyData'");
            }
            if (!args || args.webBackendService === undefined) {
                throw new Error("Missing required property 'webBackendService'");
            }
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["webBackendService"] = args ? args.webBackendService : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
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
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly webBackendService?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebBackendServiceIamPolicy resource.
 */
export interface WebBackendServiceIamPolicyArgs {
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
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly webBackendService: pulumi.Input<string>;
}
