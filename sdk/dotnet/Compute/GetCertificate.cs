// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static class GetCertificate
    {
        /// <summary>
        /// Get info about a Google Compute SSL Certificate from its name.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var myCert = Output.Create(Gcp.Compute.GetCertificate.InvokeAsync(new Gcp.Compute.GetCertificateArgs
        ///         {
        ///             Name = "my-cert",
        ///         }));
        ///         this.Certificate = myCert.Apply(myCert =&gt; myCert.Certificate);
        ///         this.CertificateId = myCert.Apply(myCert =&gt; myCert.CertificateId);
        ///         this.SelfLink = myCert.Apply(myCert =&gt; myCert.SelfLink);
        ///     }
        /// 
        ///     [Output("certificate")]
        ///     public Output&lt;string&gt; Certificate { get; set; }
        ///     [Output("certificateId")]
        ///     public Output&lt;string&gt; CertificateId { get; set; }
        ///     [Output("selfLink")]
        ///     public Output&lt;string&gt; SelfLink { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCertificateResult> InvokeAsync(GetCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCertificateResult>("gcp:compute/getCertificate:getCertificate", args ?? new GetCertificateArgs(), options.WithVersion());

        public static Output<GetCertificateResult> Invoke(GetCertificateOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Name.Box(),
                args.Project.Box()
            ).Apply(a => {
                    var args = new GetCertificateArgs();
                    a[0].Set(args, nameof(args.Name));
                    a[1].Set(args, nameof(args.Project));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetCertificateArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the certificate.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetCertificateArgs()
        {
        }
    }

    public sealed class GetCertificateOutputArgs
    {
        /// <summary>
        /// The name of the certificate.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetCertificateOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCertificateResult
    {
        public readonly string Certificate;
        public readonly int CertificateId;
        public readonly string CreationTimestamp;
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string NamePrefix;
        public readonly string PrivateKey;
        public readonly string? Project;
        public readonly string SelfLink;

        [OutputConstructor]
        private GetCertificateResult(
            string certificate,

            int certificateId,

            string creationTimestamp,

            string description,

            string id,

            string name,

            string namePrefix,

            string privateKey,

            string? project,

            string selfLink)
        {
            Certificate = certificate;
            CertificateId = certificateId;
            CreationTimestamp = creationTimestamp;
            Description = description;
            Id = id;
            Name = name;
            NamePrefix = namePrefix;
            PrivateKey = privateKey;
            Project = project;
            SelfLink = selfLink;
        }
    }
}
