// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Sql.Inputs
{

    public sealed class DatabaseInstanceIpAddressArgs : Pulumi.ResourceArgs
    {
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        [Input("timeToRetire")]
        public Input<string>? TimeToRetire { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public DatabaseInstanceIpAddressArgs()
        {
        }
    }
}
