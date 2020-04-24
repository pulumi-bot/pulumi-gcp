// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.SourceRepo.Inputs
{

    public sealed class RepositoryPubsubConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The format of the Cloud Pub/Sub messages.
        /// - PROTOBUF: The message payload is a serialized protocol buffer of SourceRepoEvent.
        /// - JSON: The message payload is a JSON string of SourceRepoEvent.
        /// </summary>
        [Input("messageFormat", required: true)]
        public Input<string> MessageFormat { get; set; } = null!;

        /// <summary>
        /// Email address of the service account used for publishing Cloud Pub/Sub messages.
        /// This service account needs to be in the same project as the PubsubConfig. When added,
        /// the caller needs to have iam.serviceAccounts.actAs permission on this service account.
        /// If unspecified, it defaults to the compute engine default service account.
        /// </summary>
        [Input("serviceAccountEmail")]
        public Input<string>? ServiceAccountEmail { get; set; }

        /// <summary>
        /// The identifier for this object. Format specified above.
        /// </summary>
        [Input("topic", required: true)]
        public Input<string> Topic { get; set; } = null!;

        public RepositoryPubsubConfigArgs()
        {
        }
    }
}
