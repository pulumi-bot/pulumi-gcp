// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Inputs
{

    public sealed class FlexibleAppVersionResourcesVolumeGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Volume size in gigabytes.
        /// </summary>
        [Input("sizeGb", required: true)]
        public Input<int> SizeGb { get; set; } = null!;

        /// <summary>
        /// Underlying volume type, e.g. 'tmpfs'.
        /// </summary>
        [Input("volumeType", required: true)]
        public Input<string> VolumeType { get; set; } = null!;

        public FlexibleAppVersionResourcesVolumeGetArgs()
        {
        }
    }
}
