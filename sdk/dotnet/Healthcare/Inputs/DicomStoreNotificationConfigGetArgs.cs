// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Healthcare.Inputs
{

    public sealed class DicomStoreNotificationConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.
        /// PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
        /// It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
        /// was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
        /// project. cloud-healthcare@system.gserviceaccount.com must have publisher permissions on the given
        /// Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.
        /// </summary>
        [Input("pubsubTopic", required: true)]
        public Input<string> PubsubTopic { get; set; } = null!;

        public DicomStoreNotificationConfigGetArgs()
        {
        }
    }
}
