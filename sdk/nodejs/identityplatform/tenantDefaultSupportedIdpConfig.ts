// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configurations options for the tenant for authenticating with a the standard set of Identity Toolkit-trusted IDPs.
 *
 * You must enable the
 * [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
 * the marketplace prior to using this resource.
 *
 * ## Example Usage
 * ### Identity Platform Tenant Default Supported Idp Config Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const tenant = new gcp.identityplatform.Tenant("tenant", {displayName: "tenant"});
 * const idpConfig = new gcp.identityplatform.TenantDefaultSupportedIdpConfig("idpConfig", {
 *     enabled: true,
 *     tenant: tenant.name,
 *     idpId: "playgames.google.com",
 *     clientId: "my-client-id",
 *     clientSecret: "secret",
 * });
 * ```
 *
 * ## Import
 *
 * TenantDefaultSupportedIdpConfig can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:identityplatform/tenantDefaultSupportedIdpConfig:TenantDefaultSupportedIdpConfig default projects/{{project}}/tenants/{{tenant}}/defaultSupportedIdpConfigs/{{idp_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:identityplatform/tenantDefaultSupportedIdpConfig:TenantDefaultSupportedIdpConfig default {{project}}/{{tenant}}/{{idp_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:identityplatform/tenantDefaultSupportedIdpConfig:TenantDefaultSupportedIdpConfig default {{tenant}}/{{idp_id}}
 * ```
 */
export class TenantDefaultSupportedIdpConfig extends pulumi.CustomResource {
    /**
     * Get an existing TenantDefaultSupportedIdpConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TenantDefaultSupportedIdpConfigState, opts?: pulumi.CustomResourceOptions): TenantDefaultSupportedIdpConfig {
        return new TenantDefaultSupportedIdpConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:identityplatform/tenantDefaultSupportedIdpConfig:TenantDefaultSupportedIdpConfig';

    /**
     * Returns true if the given object is an instance of TenantDefaultSupportedIdpConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TenantDefaultSupportedIdpConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TenantDefaultSupportedIdpConfig.__pulumiType;
    }

    /**
     * OAuth client ID
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * OAuth client secret
     */
    public readonly clientSecret!: pulumi.Output<string>;
    /**
     * If this IDP allows the user to sign in
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * ID of the IDP. Possible values include:
     * * `apple.com`
     * * `facebook.com`
     * * `gc.apple.com`
     * * `github.com`
     * * `google.com`
     * * `linkedin.com`
     * * `microsoft.com`
     * * `playgames.google.com`
     * * `twitter.com`
     * * `yahoo.com`
     */
    public readonly idpId!: pulumi.Output<string>;
    /**
     * The name of the default supported IDP config resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The name of the tenant where this DefaultSupportedIdpConfig resource exists
     */
    public readonly tenant!: pulumi.Output<string>;

    /**
     * Create a TenantDefaultSupportedIdpConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TenantDefaultSupportedIdpConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TenantDefaultSupportedIdpConfigArgs | TenantDefaultSupportedIdpConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TenantDefaultSupportedIdpConfigState | undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientSecret"] = state ? state.clientSecret : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["idpId"] = state ? state.idpId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["tenant"] = state ? state.tenant : undefined;
        } else {
            const args = argsOrState as TenantDefaultSupportedIdpConfigArgs | undefined;
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.clientSecret === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientSecret'");
            }
            if ((!args || args.idpId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'idpId'");
            }
            if ((!args || args.tenant === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenant'");
            }
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientSecret"] = args ? args.clientSecret : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["idpId"] = args ? args.idpId : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["tenant"] = args ? args.tenant : undefined;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(TenantDefaultSupportedIdpConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TenantDefaultSupportedIdpConfig resources.
 */
export interface TenantDefaultSupportedIdpConfigState {
    /**
     * OAuth client ID
     */
    clientId?: pulumi.Input<string>;
    /**
     * OAuth client secret
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * If this IDP allows the user to sign in
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * ID of the IDP. Possible values include:
     * * `apple.com`
     * * `facebook.com`
     * * `gc.apple.com`
     * * `github.com`
     * * `google.com`
     * * `linkedin.com`
     * * `microsoft.com`
     * * `playgames.google.com`
     * * `twitter.com`
     * * `yahoo.com`
     */
    idpId?: pulumi.Input<string>;
    /**
     * The name of the default supported IDP config resource
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The name of the tenant where this DefaultSupportedIdpConfig resource exists
     */
    tenant?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TenantDefaultSupportedIdpConfig resource.
 */
export interface TenantDefaultSupportedIdpConfigArgs {
    /**
     * OAuth client ID
     */
    clientId: pulumi.Input<string>;
    /**
     * OAuth client secret
     */
    clientSecret: pulumi.Input<string>;
    /**
     * If this IDP allows the user to sign in
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * ID of the IDP. Possible values include:
     * * `apple.com`
     * * `facebook.com`
     * * `gc.apple.com`
     * * `github.com`
     * * `google.com`
     * * `linkedin.com`
     * * `microsoft.com`
     * * `playgames.google.com`
     * * `twitter.com`
     * * `yahoo.com`
     */
    idpId: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The name of the tenant where this DefaultSupportedIdpConfig resource exists
     */
    tenant: pulumi.Input<string>;
}
