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
    public sealed class FirewallDeny
    {
        /// <summary>
        /// -
        /// (Optional)
        /// An optional list of ports to which this rule applies. This field
        /// is only applicable for UDP or TCP protocol. Each entry must be
        /// either an integer or a range. If not specified, this rule
        /// applies to connections through any port.
        /// Example inputs include: ["22"], ["80","443"], and
        /// ["12345-12349"].
        /// </summary>
        public readonly ImmutableArray<string> Ports;
        /// <summary>
        /// -
        /// (Required)
        /// The IP protocol to which this rule applies. The protocol type is
        /// required when creating a firewall rule. This value can either be
        /// one of the following well known protocol strings (tcp, udp,
        /// icmp, esp, ah, sctp, ipip), or the IP protocol number.
        /// </summary>
        public readonly string Protocol;

        [OutputConstructor]
        private FirewallDeny(
            ImmutableArray<string> ports,

            string protocol)
        {
            Ports = ports;
            Protocol = protocol;
        }
    }
}
