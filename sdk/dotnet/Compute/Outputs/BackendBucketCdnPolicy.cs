// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class BackendBucketCdnPolicy
    {
        /// <summary>
        /// Maximum number of seconds the response to a signed URL request will
        /// be considered fresh. After this time period,
        /// the response will be revalidated before being served.
        /// When serving responses to signed URL requests,
        /// Cloud CDN will internally behave as though
        /// all responses from this backend had a "Cache-Control: public,
        /// max-age=[TTL]" header, regardless of any existing Cache-Control
        /// header. The actual headers served in responses will not be altered.
        /// </summary>
        public readonly int SignedUrlCacheMaxAgeSec;

        [OutputConstructor]
        private BackendBucketCdnPolicy(int signedUrlCacheMaxAgeSec)
        {
            SignedUrlCacheMaxAgeSec = signedUrlCacheMaxAgeSec;
        }
    }
}
