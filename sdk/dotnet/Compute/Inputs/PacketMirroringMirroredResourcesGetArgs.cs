// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class PacketMirroringMirroredResourcesGetArgs : Pulumi.ResourceArgs
    {
        [Input("instances")]
        private InputList<Inputs.PacketMirroringMirroredResourcesInstanceGetArgs>? _instances;

        /// <summary>
        /// -
        /// (Optional)
        /// All the listed instances will be mirrored.  Specify at most 50.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.PacketMirroringMirroredResourcesInstanceGetArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<Inputs.PacketMirroringMirroredResourcesInstanceGetArgs>());
            set => _instances = value;
        }

        [Input("subnetworks")]
        private InputList<Inputs.PacketMirroringMirroredResourcesSubnetworkGetArgs>? _subnetworks;

        /// <summary>
        /// -
        /// (Optional)
        /// All instances in one of these subnetworks will be mirrored.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.PacketMirroringMirroredResourcesSubnetworkGetArgs> Subnetworks
        {
            get => _subnetworks ?? (_subnetworks = new InputList<Inputs.PacketMirroringMirroredResourcesSubnetworkGetArgs>());
            set => _subnetworks = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// -
        /// (Optional)
        /// All instances with these tags will be mirrored.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public PacketMirroringMirroredResourcesGetArgs()
        {
        }
    }
}
