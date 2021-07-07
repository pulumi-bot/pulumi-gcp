// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Iot.Inputs
{

    public sealed class DeviceConfigArgs : Pulumi.ResourceArgs
    {
        [Input("binaryData")]
        public Input<string>? BinaryData { get; set; }

        [Input("cloudUpdateTime")]
        public Input<string>? CloudUpdateTime { get; set; }

        [Input("deviceAckTime")]
        public Input<string>? DeviceAckTime { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        public DeviceConfigArgs()
        {
        }
    }
}
