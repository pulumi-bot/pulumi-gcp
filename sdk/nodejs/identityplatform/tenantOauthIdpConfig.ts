// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * OIDC IdP configuration for a Identity Toolkit project within a tenant.
 *
 * You must enable the
 * [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
 * the marketplace prior to using this resource.
 *
 *
 *
 * ## Example Usage - Identity Platform Tenant Oauth Idp Config Basic
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const tenant = new gcp.identityplatform.Tenant("tenant", {displayName: "tenant"});
 * const tenantOauthIdpConfig = new gcp.identityplatform.TenantOauthIdpConfig("tenantOauthIdpConfig", {
 *     tenant: tenant.name,
 *     displayName: "Display Name",
 *     clientId: "client-id",
 *     issuer: "issuer",
 *     enabled: true,
 *     clientSecret: "secret",
 * });
 * ```
 */
export class TenantOauthIdpConfig extends pulumi.CustomResource {
    /**
     * Get an existing TenantOauthIdpConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TenantOauthIdpConfigState, opts?: pulumi.CustomResourceOptions): TenantOauthIdpConfig {
        return new TenantOauthIdpConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:identityplatform/tenantOauthIdpConfig:TenantOauthIdpConfig';

    /**
     * Returns true if the given object is an instance of TenantOauthIdpConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TenantOauthIdpConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TenantOauthIdpConfig.__pulumiType;
    }

    /**
     * The client id of an OAuth client.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * The client secret of the OAuth client, to enable OIDC code flow.
     */
    public readonly clientSecret!: pulumi.Output<string | undefined>;
    /**
     * Human friendly display name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * If this config allows users to sign in with the provider.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * For OIDC Idps, the issuer identifier.
     */
    public readonly issuer!: pulumi.Output<string>;
    /**
     * The name of the OauthIdpConfig. Must start with `oidc.`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The name of the tenant where this OIDC IDP configuration resource exists
     */
    public readonly tenant!: pulumi.Output<string>;

    /**
     * Create a TenantOauthIdpConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TenantOauthIdpConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TenantOauthIdpConfigArgs | TenantOauthIdpConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TenantOauthIdpConfigState | undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientSecret"] = state ? state.clientSecret : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["issuer"] = state ? state.issuer : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["tenant"] = state ? state.tenant : undefined;
        } else {
            const args = argsOrState as TenantOauthIdpConfigArgs | undefined;
            if (!args || args.clientId === undefined) {
                throw new Error("Missing required property 'clientId'");
            }
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.issuer === undefined) {
                throw new Error("Missing required property 'issuer'");
            }
            if (!args || args.tenant === undefined) {
                throw new Error("Missing required property 'tenant'");
            }
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientSecret"] = args ? args.clientSecret : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["issuer"] = args ? args.issuer : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["tenant"] = args ? args.tenant : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TenantOauthIdpConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TenantOauthIdpConfig resources.
 */
export interface TenantOauthIdpConfigState {
    /**
     * The client id of an OAuth client.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The client secret of the OAuth client, to enable OIDC code flow.
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * Human friendly display name.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * If this config allows users to sign in with the provider.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * For OIDC Idps, the issuer identifier.
     */
    readonly issuer?: pulumi.Input<string>;
    /**
     * The name of the OauthIdpConfig. Must start with `oidc.`.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The name of the tenant where this OIDC IDP configuration resource exists
     */
    readonly tenant?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TenantOauthIdpConfig resource.
 */
export interface TenantOauthIdpConfigArgs {
    /**
     * The client id of an OAuth client.
     */
    readonly clientId: pulumi.Input<string>;
    /**
     * The client secret of the OAuth client, to enable OIDC code flow.
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * Human friendly display name.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * If this config allows users to sign in with the provider.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * For OIDC Idps, the issuer identifier.
     */
    readonly issuer: pulumi.Input<string>;
    /**
     * The name of the OauthIdpConfig. Must start with `oidc.`.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The name of the tenant where this OIDC IDP configuration resource exists
     */
    readonly tenant: pulumi.Input<string>;
}
