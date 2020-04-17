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
    public sealed class DatasetAccessView
    {
        /// <summary>
        /// -
        /// (Required)
        /// A unique ID for this dataset, without the project name. The ID
        /// must contain only letters (a-z, A-Z), numbers (0-9), or
        /// underscores (_). The maximum length is 1,024 characters.
        /// </summary>
        public readonly string DatasetId;
        /// <summary>
        /// -
        /// (Required)
        /// The ID of the project containing this table.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// -
        /// (Required)
        /// The ID of the table. The ID must contain only letters (a-z,
        /// A-Z), numbers (0-9), or underscores (_). The maximum length
        /// is 1,024 characters.
        /// </summary>
        public readonly string TableId;

        [OutputConstructor]
        private DatasetAccessView(
            string datasetId,

            string projectId,

            string tableId)
        {
            DatasetId = datasetId;
            ProjectId = projectId;
            TableId = tableId;
        }
    }
}
