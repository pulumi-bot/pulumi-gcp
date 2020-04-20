// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery.Outputs
{

    [OutputType]
    public sealed class AppProfileSingleClusterRouting
    {
        /// <summary>
        /// -
        /// (Optional)
        /// If true, CheckAndMutateRow and ReadModifyWriteRow requests are allowed by this app profile.
        /// It is unsafe to send these requests to the same table/row/column in multiple clusters.
        /// </summary>
        public readonly bool? AllowTransactionalWrites;
        /// <summary>
        /// -
        /// (Required)
        /// The cluster to which read/write requests should be routed.
        /// </summary>
        public readonly string ClusterId;

        [OutputConstructor]
        private AppProfileSingleClusterRouting(
            bool? allowTransactionalWrites,

            string clusterId)
        {
            AllowTransactionalWrites = allowTransactionalWrites;
            ClusterId = clusterId;
        }
    }
}
