// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.PubSub.Inputs
{

    public sealed class TopicMessageStoragePolicyGetArgs : Pulumi.ResourceArgs
    {
        [Input("allowedPersistenceRegions", required: true)]
        private InputList<string>? _allowedPersistenceRegions;

        /// <summary>
        /// -
        /// (Required)
        /// A list of IDs of GCP regions where messages that are published to
        /// the topic may be persisted in storage. Messages published by
        /// publishers running in non-allowed GCP regions (or running outside
        /// of GCP altogether) will be routed for storage in one of the
        /// allowed regions. An empty list means that no regions are allowed,
        /// and is not a valid configuration.
        /// </summary>
        public InputList<string> AllowedPersistenceRegions
        {
            get => _allowedPersistenceRegions ?? (_allowedPersistenceRegions = new InputList<string>());
            set => _allowedPersistenceRegions = value;
        }

        public TopicMessageStoragePolicyGetArgs()
        {
        }
    }
}
