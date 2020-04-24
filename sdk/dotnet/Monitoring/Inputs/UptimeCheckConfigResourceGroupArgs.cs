// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring.Inputs
{

    public sealed class UptimeCheckConfigResourceGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The group of resources being monitored. Should be the `name` of a group
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The resource type of the group members.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        public UptimeCheckConfigResourceGroupArgs()
        {
        }
    }
}
