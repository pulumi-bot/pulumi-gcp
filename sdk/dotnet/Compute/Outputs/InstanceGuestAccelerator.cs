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
    public sealed class InstanceGuestAccelerator
    {
        /// <summary>
        /// The number of the guest accelerator cards exposed to this instance.
        /// </summary>
        public readonly int Count;
        /// <summary>
        /// The GCE disk type. May be set to pd-standard or pd-ssd.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private InstanceGuestAccelerator(
            int count,

            string type)
        {
            Count = count;
            Type = type;
        }
    }
}
