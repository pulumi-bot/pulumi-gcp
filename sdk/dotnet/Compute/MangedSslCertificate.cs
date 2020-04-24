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
    /// 
    /// Deprecated: gcp.MangedSslCertificate has been deprecated in favour of gcp.ManagedSslCertificate
    /// </summary>
    [Obsolete(@"gcp.MangedSslCertificate has been deprecated in favour of gcp.ManagedSslCertificate")]
    public partial class MangedSslCertificate : Pulumi.CustomResource
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
        /// Properties relevant to a managed certificate. These will be used if the certificate is managed (as indicated by a value
        /// of 'MANAGED' in 'type').
        /// </summary>
        [Output("managed")]
        public Output<Outputs.MangedSslCertificateManaged?> Managed { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
        /// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
        /// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
        /// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. These are in the same
        /// namespace as the managed SSL certificates.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Domains associated with the certificate via Subject Alternative Name.
        /// </summary>
        [Output("subjectAlternativeNames")]
        public Output<ImmutableArray<string>> SubjectAlternativeNames { get; private set; } = null!;

        /// <summary>
        /// Enum field whose value is always 'MANAGED' - used to signal to the API which type this is.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a MangedSslCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MangedSslCertificate(string name, MangedSslCertificateArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:compute/mangedSslCertificate:MangedSslCertificate", name, args ?? new MangedSslCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MangedSslCertificate(string name, Input<string> id, MangedSslCertificateState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/mangedSslCertificate:MangedSslCertificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MangedSslCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MangedSslCertificate Get(string name, Input<string> id, MangedSslCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new MangedSslCertificate(name, id, state, options);
        }
    }

    public sealed class MangedSslCertificateArgs : Pulumi.ResourceArgs
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
        /// Properties relevant to a managed certificate. These will be used if the certificate is managed (as indicated by a value
        /// of 'MANAGED' in 'type').
        /// </summary>
        [Input("managed")]
        public Input<Inputs.MangedSslCertificateManagedArgs>? Managed { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
        /// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
        /// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
        /// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. These are in the same
        /// namespace as the managed SSL certificates.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Enum field whose value is always 'MANAGED' - used to signal to the API which type this is.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public MangedSslCertificateArgs()
        {
        }
    }

    public sealed class MangedSslCertificateState : Pulumi.ResourceArgs
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
        /// Properties relevant to a managed certificate. These will be used if the certificate is managed (as indicated by a value
        /// of 'MANAGED' in 'type').
        /// </summary>
        [Input("managed")]
        public Input<Inputs.MangedSslCertificateManagedGetArgs>? Managed { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
        /// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
        /// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
        /// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. These are in the same
        /// namespace as the managed SSL certificates.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

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
        /// Enum field whose value is always 'MANAGED' - used to signal to the API which type this is.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public MangedSslCertificateState()
        {
        }
    }
}
