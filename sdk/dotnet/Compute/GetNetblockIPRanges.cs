// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get the IP addresses from different special IP ranges on Google Cloud Platform.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_google_netblock_ip_ranges.html.markdown.
        /// </summary>
        [Obsolete("Use GetNetblockIPRanges.InvokeAsync() instead")]
        public static Task<GetNetblockIPRangesResult> GetNetblockIPRanges(GetNetblockIPRangesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetblockIPRangesResult>("gcp:compute/getNetblockIPRanges:getNetblockIPRanges", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetNetblockIPRanges
    {
        /// <summary>
        /// Use this data source to get the IP addresses from different special IP ranges on Google Cloud Platform.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_google_netblock_ip_ranges.html.markdown.
        /// </summary>
        public static Task<GetNetblockIPRangesResult> InvokeAsync(GetNetblockIPRangesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetblockIPRangesResult>("gcp:compute/getNetblockIPRanges:getNetblockIPRanges", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetNetblockIPRangesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The type of range for which to provide results.
        /// </summary>
        [Input("rangeType")]
        public string? RangeType { get; set; }

        public GetNetblockIPRangesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetNetblockIPRangesResult
    {
        /// <summary>
        /// Retrieve list of all CIDR blocks.
        /// </summary>
        public readonly ImmutableArray<string> CidrBlocks;
        /// <summary>
        /// Retrieve list of the IPv4 CIDR blocks
        /// </summary>
        public readonly ImmutableArray<string> CidrBlocksIpv4s;
        /// <summary>
        /// Retrieve list of the IPv6 CIDR blocks, if available.
        /// </summary>
        public readonly ImmutableArray<string> CidrBlocksIpv6s;
        public readonly string? RangeType;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetNetblockIPRangesResult(
            ImmutableArray<string> cidrBlocks,
            ImmutableArray<string> cidrBlocksIpv4s,
            ImmutableArray<string> cidrBlocksIpv6s,
            string? rangeType,
            string id)
        {
            CidrBlocks = cidrBlocks;
            CidrBlocksIpv4s = cidrBlocksIpv4s;
            CidrBlocksIpv6s = cidrBlocksIpv6s;
            RangeType = rangeType;
            Id = id;
        }
    }
}
