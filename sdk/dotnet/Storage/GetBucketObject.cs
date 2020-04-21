// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Storage
{
    public static class GetBucketObject
    {
        /// <summary>
        /// Gets an existing object inside an existing bucket in Google Cloud Storage service (GCS).
        /// See [the official documentation](https://cloud.google.com/storage/docs/key-terms#objects)
        /// and
        /// [API](https://cloud.google.com/storage/docs/json_api/v1/objects).
        /// 
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBucketObjectResult> InvokeAsync(GetBucketObjectArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBucketObjectResult>("gcp:storage/getBucketObject:getBucketObject", args ?? new GetBucketObjectArgs(), options.WithVersion());
    }


    public sealed class GetBucketObjectArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the containing bucket.
        /// </summary>
        [Input("bucket")]
        public string? Bucket { get; set; }

        /// <summary>
        /// The name of the object.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetBucketObjectArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBucketObjectResult
    {
        public readonly string? Bucket;
        /// <summary>
        /// (Computed) [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
        /// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
        /// </summary>
        public readonly string CacheControl;
        public readonly string Content;
        /// <summary>
        /// (Computed) [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
        /// </summary>
        public readonly string ContentDisposition;
        /// <summary>
        /// (Computed) [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
        /// </summary>
        public readonly string ContentEncoding;
        /// <summary>
        /// (Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
        /// </summary>
        public readonly string ContentLanguage;
        /// <summary>
        /// (Computed) [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
        /// </summary>
        public readonly string ContentType;
        /// <summary>
        /// (Computed) Base 64 CRC32 hash of the uploaded data.
        /// </summary>
        public readonly string Crc32c;
        public readonly string DetectMd5hash;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Computed) Base 64 MD5 hash of the uploaded data.
        /// </summary>
        public readonly string Md5hash;
        public readonly ImmutableDictionary<string, string> Metadata;
        public readonly string? Name;
        public readonly string OutputName;
        /// <summary>
        /// (Computed) A url reference to this object.
        /// </summary>
        public readonly string SelfLink;
        public readonly string Source;
        /// <summary>
        /// (Computed) The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
        /// Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`. If not provided, this defaults to the bucket's default
        /// storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
        /// </summary>
        public readonly string StorageClass;

        [OutputConstructor]
        private GetBucketObjectResult(
            string? bucket,

            string cacheControl,

            string content,

            string contentDisposition,

            string contentEncoding,

            string contentLanguage,

            string contentType,

            string crc32c,

            string detectMd5hash,

            string id,

            string md5hash,

            ImmutableDictionary<string, string> metadata,

            string? name,

            string outputName,

            string selfLink,

            string source,

            string storageClass)
        {
            Bucket = bucket;
            CacheControl = cacheControl;
            Content = content;
            ContentDisposition = contentDisposition;
            ContentEncoding = contentEncoding;
            ContentLanguage = contentLanguage;
            ContentType = contentType;
            Crc32c = crc32c;
            DetectMd5hash = detectMd5hash;
            Id = id;
            Md5hash = md5hash;
            Metadata = metadata;
            Name = name;
            OutputName = outputName;
            SelfLink = selfLink;
            Source = source;
            StorageClass = storageClass;
        }
    }
}
