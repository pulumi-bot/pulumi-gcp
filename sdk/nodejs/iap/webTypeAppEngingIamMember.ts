// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/appengine-{{appId}} * {{project}}/{{appId}} * {{appId}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy webtypeappengine IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:iap/webTypeAppEngingIamMember:WebTypeAppEngingIamMember editor "projects/{{project}}/iap_web/appengine-{{appId}} roles/iap.httpsResourceAccessor user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:iap/webTypeAppEngingIamMember:WebTypeAppEngingIamMember editor "projects/{{project}}/iap_web/appengine-{{appId}} roles/iap.httpsResourceAccessor"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:iap/webTypeAppEngingIamMember:WebTypeAppEngingIamMember editor projects/{{project}}/iap_web/appengine-{{appId}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class WebTypeAppEngingIamMember extends pulumi.CustomResource {
    /**
     * Get an existing WebTypeAppEngingIamMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebTypeAppEngingIamMemberState, opts?: pulumi.CustomResourceOptions): WebTypeAppEngingIamMember {
        return new WebTypeAppEngingIamMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:iap/webTypeAppEngingIamMember:WebTypeAppEngingIamMember';

    /**
     * Returns true if the given object is an instance of WebTypeAppEngingIamMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebTypeAppEngingIamMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebTypeAppEngingIamMember.__pulumiType;
    }

    /**
     * Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    public readonly condition!: pulumi.Output<outputs.iap.WebTypeAppEngingIamMemberCondition | undefined>;
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
     * `gcp.iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a WebTypeAppEngingIamMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebTypeAppEngingIamMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebTypeAppEngingIamMemberArgs | WebTypeAppEngingIamMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebTypeAppEngingIamMemberState | undefined;
            inputs["appId"] = state ? state.appId : undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as WebTypeAppEngingIamMemberArgs | undefined;
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            if ((!args || args.member === undefined) && !opts.urn) {
                throw new Error("Missing required property 'member'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            inputs["appId"] = args ? args.appId : undefined;
            inputs["condition"] = args ? args.condition : undefined;
            inputs["member"] = args ? args.member : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WebTypeAppEngingIamMember.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebTypeAppEngingIamMember resources.
 */
export interface WebTypeAppEngingIamMemberState {
    /**
     * Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
     */
    readonly appId?: pulumi.Input<string>;
    /**
     * ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    readonly condition?: pulumi.Input<inputs.iap.WebTypeAppEngingIamMemberCondition>;
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
     * `gcp.iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebTypeAppEngingIamMember resource.
 */
export interface WebTypeAppEngingIamMemberArgs {
    /**
     * Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
     */
    readonly appId: pulumi.Input<string>;
    /**
     * ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    readonly condition?: pulumi.Input<inputs.iap.WebTypeAppEngingIamMemberCondition>;
    readonly member: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
