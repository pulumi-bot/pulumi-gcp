// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static class GetImage
    {
        /// <summary>
        /// Get information about a Google Compute Image. Check that your service account has the `compute.imageUser` role if you want to share [custom images](https://cloud.google.com/compute/docs/images/sharing-images-across-projects) from another project. If you want to use [public images][pubimg], do not forget to specify the dedicated project. For more information see
        /// [the official documentation](https://cloud.google.com/compute/docs/images) and its [API](https://cloud.google.com/compute/docs/reference/latest/images).
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
        ///         var myImage = Output.Create(Gcp.Compute.GetImage.InvokeAsync(new Gcp.Compute.GetImageArgs
        ///         {
        ///             Family = "debian-9",
        ///             Project = "debian-cloud",
        ///         }));
        ///         // ...
        ///         var @default = new Gcp.Compute.Instance("default", new Gcp.Compute.InstanceArgs
        ///         {
        ///             BootDisk = new Gcp.Compute.Inputs.InstanceBootDiskArgs
        ///             {
        ///                 InitializeParams = new Gcp.Compute.Inputs.InstanceBootDiskInitializeParamsArgs
        ///                 {
        ///                     Image = myImage.Apply(myImage =&gt; myImage.SelfLink),
        ///                 },
        ///             },
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetImageResult> InvokeAsync(GetImageArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetImageResult>("gcp:compute/getImage:getImage", args ?? new GetImageArgs(), options.WithVersion());

        public static Output<GetImageResult> Invoke(GetImageOutputArgs? args = null, InvokeOptions? options = null)
        {
            args = args ?? new GetImageOutputArgs();
            return Pulumi.Output.All(
                args.Family.Box(),
                args.Filter.Box(),
                args.Name.Box(),
                args.Project.Box()
            ).Apply(a => {
                    var args = new GetImageArgs();
                    a[0].Set(args, nameof(args.Family));
                    a[1].Set(args, nameof(args.Filter));
                    a[2].Set(args, nameof(args.Name));
                    a[3].Set(args, nameof(args.Project));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetImageArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The family name of the image.
        /// </summary>
        [Input("family")]
        public string? Family { get; set; }

        [Input("filter")]
        public string? Filter { get; set; }

        /// <summary>
        /// The name of the image.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The project in which the resource belongs. If it is not
        /// provided, the provider project is used. If you are using a
        /// [public base image][pubimg], be sure to specify the correct Image Project.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetImageArgs()
        {
        }
    }

    public sealed class GetImageOutputArgs
    {
        /// <summary>
        /// The family name of the image.
        /// </summary>
        [Input("family")]
        public Input<string>? Family { get; set; }

        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The name of the image.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project in which the resource belongs. If it is not
        /// provided, the provider project is used. If you are using a
        /// [public base image][pubimg], be sure to specify the correct Image Project.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetImageOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetImageResult
    {
        /// <summary>
        /// The size of the image tar.gz archive stored in Google Cloud Storage in bytes.
        /// </summary>
        public readonly int ArchiveSizeBytes;
        /// <summary>
        /// The creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this image.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The size of the image when restored onto a persistent disk in gigabytes.
        /// </summary>
        public readonly int DiskSizeGb;
        /// <summary>
        /// The family name of the image.
        /// </summary>
        public readonly string Family;
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4)
        /// encoded SHA-256 hash of the [customer-supplied encryption key](https://cloud.google.com/compute/docs/disks/customer-supplied-encryption)
        /// that protects this image.
        /// </summary>
        public readonly string ImageEncryptionKeySha256;
        /// <summary>
        /// The unique identifier for the image.
        /// </summary>
        public readonly string ImageId;
        /// <summary>
        /// A fingerprint for the labels being applied to this image.
        /// </summary>
        public readonly string LabelFingerprint;
        /// <summary>
        /// A map of labels applied to this image.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// A list of applicable license URI.
        /// </summary>
        public readonly ImmutableArray<string> Licenses;
        /// <summary>
        /// The name of the image.
        /// </summary>
        public readonly string Name;
        public readonly string Project;
        /// <summary>
        /// The URI of the image.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The URL of the source disk used to create this image.
        /// </summary>
        public readonly string SourceDisk;
        /// <summary>
        /// The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4)
        /// encoded SHA-256 hash of the [customer-supplied encryption key](https://cloud.google.com/compute/docs/disks/customer-supplied-encryption)
        /// that protects this image.
        /// </summary>
        public readonly string SourceDiskEncryptionKeySha256;
        /// <summary>
        /// The ID value of the disk used to create this image.
        /// </summary>
        public readonly string SourceDiskId;
        /// <summary>
        /// The ID value of the image used to create this image.
        /// </summary>
        public readonly string SourceImageId;
        /// <summary>
        /// The status of the image. Possible values are **FAILED**, **PENDING**, or **READY**.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetImageResult(
            int archiveSizeBytes,

            string creationTimestamp,

            string description,

            int diskSizeGb,

            string family,

            string? filter,

            string id,

            string imageEncryptionKeySha256,

            string imageId,

            string labelFingerprint,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<string> licenses,

            string name,

            string project,

            string selfLink,

            string sourceDisk,

            string sourceDiskEncryptionKeySha256,

            string sourceDiskId,

            string sourceImageId,

            string status)
        {
            ArchiveSizeBytes = archiveSizeBytes;
            CreationTimestamp = creationTimestamp;
            Description = description;
            DiskSizeGb = diskSizeGb;
            Family = family;
            Filter = filter;
            Id = id;
            ImageEncryptionKeySha256 = imageEncryptionKeySha256;
            ImageId = imageId;
            LabelFingerprint = labelFingerprint;
            Labels = labels;
            Licenses = licenses;
            Name = name;
            Project = project;
            SelfLink = selfLink;
            SourceDisk = sourceDisk;
            SourceDiskEncryptionKeySha256 = sourceDiskEncryptionKeySha256;
            SourceDiskId = sourceDiskId;
            SourceImageId = sourceImageId;
            Status = status;
        }
    }
}
