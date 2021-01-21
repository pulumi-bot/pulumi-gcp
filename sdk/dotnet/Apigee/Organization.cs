// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Apigee
{
    /// <summary>
    /// An `Organization` is the top-level container in Apigee.
    /// 
    /// To get more information about Organization, see:
    /// 
    /// * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations)
    /// * How-to Guides
    ///     * [Creating an API organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org)
    /// 
    /// ## Example Usage
    /// ### Apigee Organization Cloud Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var current = Output.Create(Gcp.Organizations.GetClientConfig.InvokeAsync());
    ///         var apigeeNetwork = new Gcp.Compute.Network("apigeeNetwork", new Gcp.Compute.NetworkArgs
    ///         {
    ///         });
    ///         var apigeeRange = new Gcp.Compute.GlobalAddress("apigeeRange", new Gcp.Compute.GlobalAddressArgs
    ///         {
    ///             Purpose = "VPC_PEERING",
    ///             AddressType = "INTERNAL",
    ///             PrefixLength = 16,
    ///             Network = apigeeNetwork.Id,
    ///         });
    ///         var apigeeVpcConnection = new Gcp.ServiceNetworking.Connection("apigeeVpcConnection", new Gcp.ServiceNetworking.ConnectionArgs
    ///         {
    ///             Network = apigeeNetwork.Id,
    ///             Service = "servicenetworking.googleapis.com",
    ///             ReservedPeeringRanges = 
    ///             {
    ///                 apigeeRange.Name,
    ///             },
    ///         });
    ///         var org = new Gcp.Apigee.Organization("org", new Gcp.Apigee.OrganizationArgs
    ///         {
    ///             AnalyticsRegion = "us-central1",
    ///             ProjectId = current.Apply(current =&gt; current.Project),
    ///             AuthorizedNetwork = apigeeNetwork.Id,
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 apigeeVpcConnection,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Apigee Organization Cloud Full
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var current = Output.Create(Gcp.Organizations.GetClientConfig.InvokeAsync());
    ///         var apigeeNetwork = new Gcp.Compute.Network("apigeeNetwork", new Gcp.Compute.NetworkArgs
    ///         {
    ///         });
    ///         var apigeeRange = new Gcp.Compute.GlobalAddress("apigeeRange", new Gcp.Compute.GlobalAddressArgs
    ///         {
    ///             Purpose = "VPC_PEERING",
    ///             AddressType = "INTERNAL",
    ///             PrefixLength = 16,
    ///             Network = apigeeNetwork.Id,
    ///         });
    ///         var apigeeVpcConnection = new Gcp.ServiceNetworking.Connection("apigeeVpcConnection", new Gcp.ServiceNetworking.ConnectionArgs
    ///         {
    ///             Network = apigeeNetwork.Id,
    ///             Service = "servicenetworking.googleapis.com",
    ///             ReservedPeeringRanges = 
    ///             {
    ///                 apigeeRange.Name,
    ///             },
    ///         });
    ///         var apigeeKeyring = new Gcp.Kms.KeyRing("apigeeKeyring", new Gcp.Kms.KeyRingArgs
    ///         {
    ///             Location = "us-central1",
    ///         });
    ///         var apigeeKey = new Gcp.Kms.CryptoKey("apigeeKey", new Gcp.Kms.CryptoKeyArgs
    ///         {
    ///             KeyRing = apigeeKeyring.Id,
    ///         });
    ///         var apigeeSa = new Gcp.Projects.ServiceIdentity("apigeeSa", new Gcp.Projects.ServiceIdentityArgs
    ///         {
    ///             Project = google_project.Project.Project_id,
    ///             Service = google_project_service.Apigee.Service,
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///         var apigeeSaKeyuser = new Gcp.Kms.CryptoKeyIAMBinding("apigeeSaKeyuser", new Gcp.Kms.CryptoKeyIAMBindingArgs
    ///         {
    ///             CryptoKeyId = apigeeKey.Id,
    ///             Role = "roles/cloudkms.cryptoKeyEncrypterDecrypter",
    ///             Members = 
    ///             {
    ///                 apigeeSa.Email.Apply(email =&gt; $"serviceAccount:{email}"),
    ///             },
    ///         });
    ///         var org = new Gcp.Apigee.Organization("org", new Gcp.Apigee.OrganizationArgs
    ///         {
    ///             AnalyticsRegion = "us-central1",
    ///             DisplayName = "apigee-org",
    ///             Description = "Terraform-provisioned Apigee Org.",
    ///             ProjectId = current.Apply(current =&gt; current.Project),
    ///             AuthorizedNetwork = apigeeNetwork.Id,
    ///             RuntimeDatabaseEncryptionKeyName = google_kms_key.Apigee_key.Id,
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 apigeeVpcConnection,
    ///                 apigeeSaKeyuser,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Organization can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:apigee/organization:Organization default organizations/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:apigee/organization:Organization default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:apigee/organization:Organization")]
    public partial class Organization : Pulumi.CustomResource
    {
        /// <summary>
        /// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
        /// </summary>
        [Output("analyticsRegion")]
        public Output<string?> AnalyticsRegion { get; private set; } = null!;

        /// <summary>
        /// Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
        /// See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
        /// Valid only when `RuntimeType` is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".
        /// </summary>
        [Output("authorizedNetwork")]
        public Output<string?> AuthorizedNetwork { get; private set; } = null!;

        /// <summary>
        /// Output only. Base64-encoded public certificate for the root CA of the Apigee organization. Valid only when 'RuntimeType'
        /// is CLOUD. A base64-encoded string.
        /// </summary>
        [Output("caCertificate")]
        public Output<string> CaCertificate { get; private set; } = null!;

        /// <summary>
        /// Description of the Apigee organization.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the Apigee organization.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Output only. Name of the Apigee organization.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project ID associated with the Apigee organization.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
        /// Update is not allowed after the organization is created.
        /// If not specified, a Google-Managed encryption key will be used.
        /// Valid only when `RuntimeType` is CLOUD. For example: `projects/foo/locations/us/keyRings/bar/cryptoKeys/baz`.
        /// </summary>
        [Output("runtimeDatabaseEncryptionKeyName")]
        public Output<string?> RuntimeDatabaseEncryptionKeyName { get; private set; } = null!;

        /// <summary>
        /// Runtime type of the Apigee organization based on the Apigee subscription purchased.
        /// Default value is `CLOUD`.
        /// Possible values are `CLOUD` and `HYBRID`.
        /// </summary>
        [Output("runtimeType")]
        public Output<string?> RuntimeType { get; private set; } = null!;

        /// <summary>
        /// Output only. Subscription type of the Apigee organization. Valid values include trial (free, limited, and for evaluation
        /// purposes only) or paid (full subscription has been purchased).
        /// </summary>
        [Output("subscriptionType")]
        public Output<string> SubscriptionType { get; private set; } = null!;


        /// <summary>
        /// Create a Organization resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Organization(string name, OrganizationArgs args, CustomResourceOptions? options = null)
            : base("gcp:apigee/organization:Organization", name, args ?? new OrganizationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Organization(string name, Input<string> id, OrganizationState? state = null, CustomResourceOptions? options = null)
            : base("gcp:apigee/organization:Organization", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Organization resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Organization Get(string name, Input<string> id, OrganizationState? state = null, CustomResourceOptions? options = null)
        {
            return new Organization(name, id, state, options);
        }
    }

    public sealed class OrganizationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
        /// </summary>
        [Input("analyticsRegion")]
        public Input<string>? AnalyticsRegion { get; set; }

        /// <summary>
        /// Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
        /// See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
        /// Valid only when `RuntimeType` is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".
        /// </summary>
        [Input("authorizedNetwork")]
        public Input<string>? AuthorizedNetwork { get; set; }

        /// <summary>
        /// Description of the Apigee organization.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the Apigee organization.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The project ID associated with the Apigee organization.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
        /// Update is not allowed after the organization is created.
        /// If not specified, a Google-Managed encryption key will be used.
        /// Valid only when `RuntimeType` is CLOUD. For example: `projects/foo/locations/us/keyRings/bar/cryptoKeys/baz`.
        /// </summary>
        [Input("runtimeDatabaseEncryptionKeyName")]
        public Input<string>? RuntimeDatabaseEncryptionKeyName { get; set; }

        /// <summary>
        /// Runtime type of the Apigee organization based on the Apigee subscription purchased.
        /// Default value is `CLOUD`.
        /// Possible values are `CLOUD` and `HYBRID`.
        /// </summary>
        [Input("runtimeType")]
        public Input<string>? RuntimeType { get; set; }

        public OrganizationArgs()
        {
        }
    }

    public sealed class OrganizationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
        /// </summary>
        [Input("analyticsRegion")]
        public Input<string>? AnalyticsRegion { get; set; }

        /// <summary>
        /// Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
        /// See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
        /// Valid only when `RuntimeType` is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".
        /// </summary>
        [Input("authorizedNetwork")]
        public Input<string>? AuthorizedNetwork { get; set; }

        /// <summary>
        /// Output only. Base64-encoded public certificate for the root CA of the Apigee organization. Valid only when 'RuntimeType'
        /// is CLOUD. A base64-encoded string.
        /// </summary>
        [Input("caCertificate")]
        public Input<string>? CaCertificate { get; set; }

        /// <summary>
        /// Description of the Apigee organization.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the Apigee organization.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Output only. Name of the Apigee organization.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project ID associated with the Apigee organization.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
        /// Update is not allowed after the organization is created.
        /// If not specified, a Google-Managed encryption key will be used.
        /// Valid only when `RuntimeType` is CLOUD. For example: `projects/foo/locations/us/keyRings/bar/cryptoKeys/baz`.
        /// </summary>
        [Input("runtimeDatabaseEncryptionKeyName")]
        public Input<string>? RuntimeDatabaseEncryptionKeyName { get; set; }

        /// <summary>
        /// Runtime type of the Apigee organization based on the Apigee subscription purchased.
        /// Default value is `CLOUD`.
        /// Possible values are `CLOUD` and `HYBRID`.
        /// </summary>
        [Input("runtimeType")]
        public Input<string>? RuntimeType { get; set; }

        /// <summary>
        /// Output only. Subscription type of the Apigee organization. Valid values include trial (free, limited, and for evaluation
        /// purposes only) or paid (full subscription has been purchased).
        /// </summary>
        [Input("subscriptionType")]
        public Input<string>? SubscriptionType { get; set; }

        public OrganizationState()
        {
        }
    }
}
