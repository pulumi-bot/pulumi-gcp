// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class PerInstanceConfigPreservedStateDiskGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A value that prescribes what should happen to the stateful disk when the VM instance is deleted.
        /// The available options are `NEVER` and `ON_PERMANENT_INSTANCE_DELETION`.
        /// `NEVER` detatch the disk when the VM is deleted, but not delete the disk.
        /// `ON_PERMANENT_INSTANCE_DELETION` will delete the stateful disk when the VM is permanently
        /// deleted from the instance group.
        /// </summary>
        [Input("deleteRule")]
        public Input<string>? DeleteRule { get; set; }

        /// <summary>
        /// A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance.
        /// </summary>
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        /// <summary>
        /// The mode of the disk.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The URI of an existing persistent disk to attach under the specified device-name in the format
        /// `projects/project-id/zones/zone/disks/disk-name`.
        /// </summary>
        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        public PerInstanceConfigPreservedStateDiskGetArgs()
        {
        }
    }
}
