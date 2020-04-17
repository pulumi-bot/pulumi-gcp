// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionInstanceGroupManagerNamedPortArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the instance group manager. Must be 1-63
        /// characters long and comply with
        /// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
        /// include lowercase letters, numbers, and hyphens.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The port number.
        /// - - -
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        public RegionInstanceGroupManagerNamedPortArgs()
        {
        }
    }
}
