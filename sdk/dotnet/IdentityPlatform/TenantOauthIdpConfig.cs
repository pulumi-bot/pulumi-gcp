// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.IdentityPlatform
{
    /// <summary>
    /// OIDC IdP configuration for a Identity Toolkit project within a tenant.
    /// 
    /// You must enable the
    /// [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
    /// the marketplace prior to using this resource.
    /// 
    /// ## Example Usage
    /// ### Identity Platform Tenant Oauth Idp Config Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var tenant = new Gcp.IdentityPlatform.Tenant("tenant", new Gcp.IdentityPlatform.TenantArgs
    ///         {
    ///             DisplayName = "tenant",
    ///         });
    ///         var tenantOauthIdpConfig = new Gcp.IdentityPlatform.TenantOauthIdpConfig("tenantOauthIdpConfig", new Gcp.IdentityPlatform.TenantOauthIdpConfigArgs
    ///         {
    ///             Tenant = tenant.Name,
    ///             DisplayName = "Display Name",
    ///             ClientId = "client-id",
    ///             Issuer = "issuer",
    ///             Enabled = true,
    ///             ClientSecret = "secret",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// TenantOauthIdpConfig can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:identityplatform/tenantOauthIdpConfig:TenantOauthIdpConfig default projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:identityplatform/tenantOauthIdpConfig:TenantOauthIdpConfig default {{project}}/{{tenant}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:identityplatform/tenantOauthIdpConfig:TenantOauthIdpConfig default {{tenant}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:identityplatform/tenantOauthIdpConfig:TenantOauthIdpConfig")]
    public partial class TenantOauthIdpConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// The client id of an OAuth client.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// The client secret of the OAuth client, to enable OIDC code flow.
        /// </summary>
        [Output("clientSecret")]
        public Output<string?> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// Human friendly display name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// If this config allows users to sign in with the provider.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// For OIDC Idps, the issuer identifier.
        /// </summary>
        [Output("issuer")]
        public Output<string> Issuer { get; private set; } = null!;

        /// <summary>
        /// The name of the OauthIdpConfig. Must start with `oidc.`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The name of the tenant where this OIDC IDP configuration resource exists
        /// </summary>
        [Output("tenant")]
        public Output<string> Tenant { get; private set; } = null!;


        /// <summary>
        /// Create a TenantOauthIdpConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TenantOauthIdpConfig(string name, TenantOauthIdpConfigArgs args, CustomResourceOptions? options = null)
            : base("gcp:identityplatform/tenantOauthIdpConfig:TenantOauthIdpConfig", name, args ?? new TenantOauthIdpConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TenantOauthIdpConfig(string name, Input<string> id, TenantOauthIdpConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:identityplatform/tenantOauthIdpConfig:TenantOauthIdpConfig", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TenantOauthIdpConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TenantOauthIdpConfig Get(string name, Input<string> id, TenantOauthIdpConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new TenantOauthIdpConfig(name, id, state, options);
        }
    }

    public sealed class TenantOauthIdpConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client id of an OAuth client.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The client secret of the OAuth client, to enable OIDC code flow.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// Human friendly display name.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// If this config allows users to sign in with the provider.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// For OIDC Idps, the issuer identifier.
        /// </summary>
        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

        /// <summary>
        /// The name of the OauthIdpConfig. Must start with `oidc.`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The name of the tenant where this OIDC IDP configuration resource exists
        /// </summary>
        [Input("tenant", required: true)]
        public Input<string> Tenant { get; set; } = null!;

        public TenantOauthIdpConfigArgs()
        {
        }
    }

    public sealed class TenantOauthIdpConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client id of an OAuth client.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The client secret of the OAuth client, to enable OIDC code flow.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// Human friendly display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// If this config allows users to sign in with the provider.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// For OIDC Idps, the issuer identifier.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// The name of the OauthIdpConfig. Must start with `oidc.`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The name of the tenant where this OIDC IDP configuration resource exists
        /// </summary>
        [Input("tenant")]
        public Input<string>? Tenant { get; set; }

        public TenantOauthIdpConfigState()
        {
        }
    }
}
