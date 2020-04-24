// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Composer.Inputs
{

    public sealed class EnvironmentConfigPrivateEnvironmentConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// If true, access to the public endpoint of the GKE cluster is denied.
        /// </summary>
        [Input("enablePrivateEndpoint")]
        public Input<bool>? EnablePrivateEndpoint { get; set; }

        /// <summary>
        /// The IP range in CIDR notation to use for the hosted master network. This range is used
        /// for assigning internal IP addresses to the cluster master or set of masters and to the
        /// internal load balancer virtual IP. This range must not overlap with any other ranges
        /// in use within the cluster's network.
        /// If left blank, the default value of '172.16.0.0/28' is used.
        /// </summary>
        [Input("masterIpv4CidrBlock")]
        public Input<string>? MasterIpv4CidrBlock { get; set; }

        public EnvironmentConfigPrivateEnvironmentConfigArgs()
        {
        }
    }
}
