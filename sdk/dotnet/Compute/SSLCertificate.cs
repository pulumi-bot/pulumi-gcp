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
    /// An SslCertificate resource, used for HTTPS load balancing. This resource
    /// provides a mechanism to upload an SSL key and certificate to
    /// the load balancer to serve secure connections from the user.
    /// 
    /// 
    /// To get more information about SslCertificate, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/sslCertificates)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/load-balancing/docs/ssl-certificates)
    /// 
    /// &gt; **Warning:** All arguments including `certificate` and `private_key` will be stored in the raw
    /// state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).
    /// 
    /// ## Example Usage - Ssl Certificate Basic
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var default = new Gcp.Compute.SSLCertificate("default", new Gcp.Compute.SSLCertificateArgs
    ///         {
    ///             NamePrefix = "my-certificate-",
    ///             Description = "a description",
    ///             PrivateKey = "TODO: ReadFile",
    ///             Certificate = "TODO: ReadFile",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class SSLCertificate : Pulumi.CustomResource
    {
        /// <summary>
        /// The certificate in PEM format.
        /// The certificate chain must be no greater than 5 certs long.
        /// The chain must include at least one intermediate cert.  **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

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
        /// Creates a unique name beginning with the
        /// specified prefix. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// The write-only private key in PEM format.  **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Output("privateKey")]
        public Output<string> PrivateKey { get; private set; } = null!;

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
        /// Create a SSLCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SSLCertificate(string name, SSLCertificateArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/sSLCertificate:SSLCertificate", name, args ?? new SSLCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SSLCertificate(string name, Input<string> id, SSLCertificateState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/sSLCertificate:SSLCertificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SSLCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SSLCertificate Get(string name, Input<string> id, SSLCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new SSLCertificate(name, id, state, options);
        }
    }

    public sealed class SSLCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The certificate in PEM format.
        /// The certificate chain must be no greater than 5 certs long.
        /// The chain must include at least one intermediate cert.  **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Input("certificate", required: true)]
        public Input<string> Certificate { get; set; } = null!;

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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
        /// Creates a unique name beginning with the
        /// specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The write-only private key in PEM format.  **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Input("privateKey", required: true)]
        public Input<string> PrivateKey { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public SSLCertificateArgs()
        {
        }
    }

    public sealed class SSLCertificateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The certificate in PEM format.
        /// The certificate chain must be no greater than 5 certs long.
        /// The chain must include at least one intermediate cert.  **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

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
        /// Creates a unique name beginning with the
        /// specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The write-only private key in PEM format.  **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Input("privateKey")]
        public Input<string>? PrivateKey { get; set; }

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

        public SSLCertificateState()
        {
        }
    }
}
