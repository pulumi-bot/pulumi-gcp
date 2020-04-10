// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Sql
{
    /// <summary>
    /// Creates a new Google SQL SSL Cert on a Google SQL Instance. For more information, see the [official documentation](https://cloud.google.com/sql/), or the [JSON API](https://cloud.google.com/sql/docs/mysql/admin-api/v1beta4/sslCerts).
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/sql_ssl_cert.html.markdown.
    /// </summary>
    public partial class SslCert : Pulumi.CustomResource
    {
        /// <summary>
        /// The actual certificate data for this client certificate.
        /// </summary>
        [Output("cert")]
        public Output<string> Cert { get; private set; } = null!;

        /// <summary>
        /// The serial number extracted from the certificate data.
        /// </summary>
        [Output("certSerialNumber")]
        public Output<string> CertSerialNumber { get; private set; } = null!;

        /// <summary>
        /// The common name to be used in the certificate to identify the
        /// client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
        /// </summary>
        [Output("commonName")]
        public Output<string> CommonName { get; private set; } = null!;

        /// <summary>
        /// The time when the certificate was created in RFC 3339 format,
        /// for example 2012-11-15T16:19:00.094Z.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The time when the certificate expires in RFC 3339 format,
        /// for example 2012-11-15T16:19:00.094Z.
        /// </summary>
        [Output("expirationTime")]
        public Output<string> ExpirationTime { get; private set; } = null!;

        /// <summary>
        /// The name of the Cloud SQL instance. Changing this
        /// forces a new resource to be created.
        /// </summary>
        [Output("instance")]
        public Output<string> Instance { get; private set; } = null!;

        /// <summary>
        /// The private key associated with the client certificate.
        /// </summary>
        [Output("privateKey")]
        public Output<string> PrivateKey { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The CA cert of the server this client cert was generated from.
        /// </summary>
        [Output("serverCaCert")]
        public Output<string> ServerCaCert { get; private set; } = null!;

        /// <summary>
        /// The SHA1 Fingerprint of the certificate.
        /// </summary>
        [Output("sha1Fingerprint")]
        public Output<string> Sha1Fingerprint { get; private set; } = null!;


        /// <summary>
        /// Create a SslCert resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SslCert(string name, SslCertArgs args, CustomResourceOptions? options = null)
            : base("gcp:sql/sslCert:SslCert", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SslCert(string name, Input<string> id, SslCertState? state = null, CustomResourceOptions? options = null)
            : base("gcp:sql/sslCert:SslCert", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SslCert resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SslCert Get(string name, Input<string> id, SslCertState? state = null, CustomResourceOptions? options = null)
        {
            return new SslCert(name, id, state, options);
        }
    }

    public sealed class SslCertArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The common name to be used in the certificate to identify the
        /// client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
        /// </summary>
        [Input("commonName", required: true)]
        public Input<string> CommonName { get; set; } = null!;

        /// <summary>
        /// The name of the Cloud SQL instance. Changing this
        /// forces a new resource to be created.
        /// </summary>
        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public SslCertArgs()
        {
        }
    }

    public sealed class SslCertState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The actual certificate data for this client certificate.
        /// </summary>
        [Input("cert")]
        public Input<string>? Cert { get; set; }

        /// <summary>
        /// The serial number extracted from the certificate data.
        /// </summary>
        [Input("certSerialNumber")]
        public Input<string>? CertSerialNumber { get; set; }

        /// <summary>
        /// The common name to be used in the certificate to identify the
        /// client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
        /// </summary>
        [Input("commonName")]
        public Input<string>? CommonName { get; set; }

        /// <summary>
        /// The time when the certificate was created in RFC 3339 format,
        /// for example 2012-11-15T16:19:00.094Z.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The time when the certificate expires in RFC 3339 format,
        /// for example 2012-11-15T16:19:00.094Z.
        /// </summary>
        [Input("expirationTime")]
        public Input<string>? ExpirationTime { get; set; }

        /// <summary>
        /// The name of the Cloud SQL instance. Changing this
        /// forces a new resource to be created.
        /// </summary>
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        /// <summary>
        /// The private key associated with the client certificate.
        /// </summary>
        [Input("privateKey")]
        public Input<string>? PrivateKey { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The CA cert of the server this client cert was generated from.
        /// </summary>
        [Input("serverCaCert")]
        public Input<string>? ServerCaCert { get; set; }

        /// <summary>
        /// The SHA1 Fingerprint of the certificate.
        /// </summary>
        [Input("sha1Fingerprint")]
        public Input<string>? Sha1Fingerprint { get; set; }

        public SslCertState()
        {
        }
    }
}
