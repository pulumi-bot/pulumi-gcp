// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for API Gateway Gateway. Each of these resources serves a different use case:
 *
 * * `gcp.apigateway.GatewayIamPolicy`: Authoritative. Sets the IAM policy for the gateway and replaces any existing policy already attached.
 * * `gcp.apigateway.GatewayIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the gateway are preserved.
 * * `gcp.apigateway.GatewayIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the gateway are preserved.
 *
 * > **Note:** `gcp.apigateway.GatewayIamPolicy` **cannot** be used in conjunction with `gcp.apigateway.GatewayIamBinding` and `gcp.apigateway.GatewayIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.apigateway.GatewayIamBinding` resources **can be** used in conjunction with `gcp.apigateway.GatewayIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_api\_gateway\_gateway\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/apigateway.viewer",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.apigateway.GatewayIamPolicy("policy", {
 *     project: google_api_gateway_gateway.api_gw.project,
 *     region: google_api_gateway_gateway.api_gw.region,
 *     gateway: google_api_gateway_gateway.api_gw.gateway_id,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## google\_api\_gateway\_gateway\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.apigateway.GatewayIamBinding("binding", {
 *     project: google_api_gateway_gateway.api_gw.project,
 *     region: google_api_gateway_gateway.api_gw.region,
 *     gateway: google_api_gateway_gateway.api_gw.gateway_id,
 *     role: "roles/apigateway.viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * ## google\_api\_gateway\_gateway\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.apigateway.GatewayIamMember("member", {
 *     project: google_api_gateway_gateway.api_gw.project,
 *     region: google_api_gateway_gateway.api_gw.region,
 *     gateway: google_api_gateway_gateway.api_gw.gateway_id,
 *     role: "roles/apigateway.viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{region}}/gateways/{{name}} * {{project}}/{{region}}/{{name}} * {{region}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. API Gateway gateway IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy editor "projects/{{project}}/locations/{{region}}/gateways/{{gateway}} roles/apigateway.viewer user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy editor "projects/{{project}}/locations/{{region}}/gateways/{{gateway}} roles/apigateway.viewer"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy editor projects/{{project}}/locations/{{region}}/gateways/{{gateway}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class GatewayIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing GatewayIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GatewayIamPolicyState, opts?: pulumi.CustomResourceOptions): GatewayIamPolicy {
        return new GatewayIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy';

    /**
     * Returns true if the given object is an instance of GatewayIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GatewayIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GatewayIamPolicy.__pulumiType;
    }

    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly gateway!: pulumi.Output<string>;
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
     * The region of the gateway for the API.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
     * region is specified, it is taken from the provider configuration.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a GatewayIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GatewayIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GatewayIamPolicyArgs | GatewayIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GatewayIamPolicyState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["gateway"] = state ? state.gateway : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as GatewayIamPolicyArgs | undefined;
            if ((!args || args.gateway === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gateway'");
            }
            if ((!args || args.policyData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyData'");
            }
            inputs["gateway"] = args ? args.gateway : undefined;
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(GatewayIamPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GatewayIamPolicy resources.
 */
export interface GatewayIamPolicyState {
    /**
     * (Computed) The etag of the IAM policy.
     */
    etag?: pulumi.Input<string>;
    gateway?: pulumi.Input<string>;
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
     * The region of the gateway for the API.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
     * region is specified, it is taken from the provider configuration.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GatewayIamPolicy resource.
 */
export interface GatewayIamPolicyArgs {
    gateway: pulumi.Input<string>;
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
     * The region of the gateway for the API.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
     * region is specified, it is taken from the provider configuration.
     */
    region?: pulumi.Input<string>;
}
