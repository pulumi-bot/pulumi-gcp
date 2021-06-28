// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class InstanceFromMachineImageBootDiskInitializeParamsArgs : Pulumi.ResourceArgs
    {
        [Input("image")]
        public Input<string>? Image { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        [Input("size")]
        public Input<int>? Size { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public InstanceFromMachineImageBootDiskInitializeParamsArgs()
        {
        }
    }
}
