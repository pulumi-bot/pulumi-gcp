// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dns.Outputs
{

    [OutputType]
    public sealed class PolicyAlternativeNameServerConfig
    {
        /// <summary>
        /// -
        /// (Required)
        /// Sets an alternative name server for the associated networks. When specified,
        /// all DNS queries are forwarded to a name server that you choose. Names such as .internal
        /// are not available when an alternative name server is specified.  Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.PolicyAlternativeNameServerConfigTargetNameServer> TargetNameServers;

        [OutputConstructor]
        private PolicyAlternativeNameServerConfig(ImmutableArray<Outputs.PolicyAlternativeNameServerConfigTargetNameServer> targetNameServers)
        {
            TargetNameServers = targetNameServers;
        }
    }
}
