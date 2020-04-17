// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudScheduler.Outputs
{

    [OutputType]
    public sealed class JobPubsubTarget
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Attributes for PubsubMessage.
        /// Pubsub message must contain either non-empty data, or at least one attribute.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Attributes;
        /// <summary>
        /// -
        /// (Optional)
        /// The message payload for PubsubMessage.
        /// Pubsub message must contain either non-empty data, or at least one attribute.
        /// </summary>
        public readonly string? Data;
        /// <summary>
        /// -
        /// (Required)
        /// The full resource name for the Cloud Pub/Sub topic to which
        /// messages will be published when a job is delivered. ~&gt;**NOTE**:
        /// The topic name must be in the same format as required by PubSub's
        /// PublishRequest.name, e.g. `projects/my-project/topics/my-topic`.
        /// </summary>
        public readonly string TopicName;

        [OutputConstructor]
        private JobPubsubTarget(
            ImmutableDictionary<string, string>? attributes,

            string? data,

            string topicName)
        {
            Attributes = attributes;
            Data = data;
            TopicName = topicName;
        }
    }
}
