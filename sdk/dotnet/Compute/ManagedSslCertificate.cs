// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// An SslCertificate resource, used for HTTPS load balancing.  This resource
    /// represents a certificate for which the certificate secrets are created and
    /// managed by Google.
    /// 
    /// For a resource where you provide the key, see the
    /// SSL Certificate resource.
    /// 
    /// To get more information about ManagedSslCertificate, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/sslCertificates)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/load-balancing/docs/ssl-certificates)
    /// 
    /// &gt; **Warning:** This resource should be used with extreme caution!  Provisioning an SSL
    /// certificate is complex.  Ensure that you understand the lifecycle of a
    /// certificate before attempting complex tasks like cert rotation automatically.
    /// This resource will "return" as soon as the certificate object is created,
    /// but post-creation the certificate object will go through a "provisioning"
    /// process.  The provisioning process can complete only when the domain name
    /// for which the certificate is created points to a target pool which, itself,
    /// points at the certificate.  Depending on your DNS provider, this may take
    /// some time, and migrating from self-managed certificates to Google-managed
    /// certificates may entail some downtime while the certificate provisions.
    /// 
    /// In conclusion: Be extremely cautious.
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// ManagedSslCertificate can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/managedSslCertificate:ManagedSslCertificate default projects/{{project}}/global/sslCertificates/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/managedSslCertificate:ManagedSslCertificate default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/managedSslCertificate:ManagedSslCertificate default {{name}}
    /// ```
    /// </summary>
    public partial class ManagedSslCertificate : Pulumi.CustomResource
    {
        /// <summary>
        /// The unique identifier for the resource.
        /// </summary>
        [Output("certificateId")]
        public Output<int> CertificateId { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Expire time of the certificate.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// Properties relevant to a managed certificate.  These will be used if the
        /// certificate is managed (as indicated by a value of `MANAGED` in `type`).
        /// Structure is documented below.
        /// </summary>
        [Output("managed")]
        public Output<Outputs.ManagedSslCertificateManaged?> Managed { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
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
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Domains associated with the certificate via Subject Alternative Name.
        /// </summary>
        [Output("subjectAlternativeNames")]
        public Output<ImmutableArray<string>> SubjectAlternativeNames { get; private set; } = null!;

        /// <summary>
        /// Enum field whose value is always `MANAGED` - used to signal to the API
        /// which type this is.
        /// Default value is `MANAGED`.
        /// Possible values are `MANAGED`.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedSslCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedSslCertificate(string name, ManagedSslCertificateArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:compute/managedSslCertificate:ManagedSslCertificate", name, args ?? new ManagedSslCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedSslCertificate(string name, Input<string> id, ManagedSslCertificateState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/managedSslCertificate:ManagedSslCertificate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "gcp:compute/mangedSslCertificate:MangedSslCertificate"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagedSslCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedSslCertificate Get(string name, Input<string> id, ManagedSslCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new ManagedSslCertificate(name, id, state, options);
        }
    }

    public sealed class ManagedSslCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique identifier for the resource.
        /// </summary>
        [Input("certificateId")]
        public Input<int>? CertificateId { get; set; }

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Properties relevant to a managed certificate.  These will be used if the
        /// certificate is managed (as indicated by a value of `MANAGED` in `type`).
        /// Structure is documented below.
        /// </summary>
        [Input("managed")]
        public Input<Inputs.ManagedSslCertificateManagedArgs>? Managed { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
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
        /// Enum field whose value is always `MANAGED` - used to signal to the API
        /// which type this is.
        /// Default value is `MANAGED`.
        /// Possible values are `MANAGED`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ManagedSslCertificateArgs()
        {
        }
    }

    public sealed class ManagedSslCertificateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique identifier for the resource.
        /// </summary>
        [Input("certificateId")]
        public Input<int>? CertificateId { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Expire time of the certificate.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        /// <summary>
        /// Properties relevant to a managed certificate.  These will be used if the
        /// certificate is managed (as indicated by a value of `MANAGED` in `type`).
        /// Structure is documented below.
        /// </summary>
        [Input("managed")]
        public Input<Inputs.ManagedSslCertificateManagedGetArgs>? Managed { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
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
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        [Input("subjectAlternativeNames")]
        private InputList<string>? _subjectAlternativeNames;

        /// <summary>
        /// Domains associated with the certificate via Subject Alternative Name.
        /// </summary>
        public InputList<string> SubjectAlternativeNames
        {
            get => _subjectAlternativeNames ?? (_subjectAlternativeNames = new InputList<string>());
            set => _subjectAlternativeNames = value;
        }

        /// <summary>
        /// Enum field whose value is always `MANAGED` - used to signal to the API
        /// which type this is.
        /// Default value is `MANAGED`.
        /// Possible values are `MANAGED`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ManagedSslCertificateState()
        {
        }
    }
}
