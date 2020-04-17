// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring.Outputs
{

    [OutputType]
    public sealed class UptimeCheckConfigTcpCheck
    {
        /// <summary>
        /// -
        /// (Required)
        /// The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) to construct the full URL.
        /// </summary>
        public readonly int Port;

        [OutputConstructor]
        private UptimeCheckConfigTcpCheck(int port)
        {
            Port = port;
        }
    }
}
