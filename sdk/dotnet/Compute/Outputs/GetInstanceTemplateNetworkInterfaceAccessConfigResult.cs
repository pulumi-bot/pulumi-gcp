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
    public sealed class GetInstanceTemplateNetworkInterfaceAccessConfigResult
    {
        /// <summary>
        /// The IP address that will be 1:1 mapped to the instance's
        /// network ip. If not given, one will be generated.
        /// </summary>
        public readonly string NatIp;
        /// <summary>
        /// The [networking tier][network-tier] used for configuring
        /// this instance template. This field can take the following values: PREMIUM or
        /// STANDARD. If this field is not specified, it is assumed to be PREMIUM.
        /// </summary>
        public readonly string NetworkTier;
        public readonly string PublicPtrDomainName;

        [OutputConstructor]
        private GetInstanceTemplateNetworkInterfaceAccessConfigResult(
            string natIp,

            string networkTier,

            string publicPtrDomainName)
        {
            NatIp = natIp;
            NetworkTier = networkTier;
            PublicPtrDomainName = publicPtrDomainName;
        }
    }
}
