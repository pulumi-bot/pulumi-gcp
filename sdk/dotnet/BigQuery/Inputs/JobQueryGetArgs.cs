// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery.Inputs
{

    public sealed class JobQueryGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true and query uses legacy SQL dialect, allows the query to produce arbitrarily large result tables at a slight cost in performance.
        /// Requires destinationTable to be set. For standard SQL queries, this flag is ignored and large results are always allowed.
        /// However, you must still set destinationTable when result size exceeds the allowed maximum response size.
        /// </summary>
        [Input("allowLargeResults")]
        public Input<bool>? AllowLargeResults { get; set; }

        /// <summary>
        /// Specifies whether the job is allowed to create new tables. The following values are supported:
        /// CREATE_IF_NEEDED: If the table does not exist, BigQuery creates the table.
        /// CREATE_NEVER: The table must already exist. If it does not, a 'notFound' error is returned in the job result.
        /// Creation, truncation and append actions occur as one atomic update upon job completion
        /// </summary>
        [Input("createDisposition")]
        public Input<string>? CreateDisposition { get; set; }

        /// <summary>
        /// Specifies the default dataset to use for unqualified table names in the query. Note that this does not alter behavior of unqualified dataset names.  Structure is documented below.
        /// </summary>
        [Input("defaultDataset")]
        public Input<Inputs.JobQueryDefaultDatasetGetArgs>? DefaultDataset { get; set; }

        /// <summary>
        /// Custom encryption configuration (e.g., Cloud KMS keys)  Structure is documented below.
        /// </summary>
        [Input("destinationEncryptionConfiguration")]
        public Input<Inputs.JobQueryDestinationEncryptionConfigurationGetArgs>? DestinationEncryptionConfiguration { get; set; }

        /// <summary>
        /// The destination table.  Structure is documented below.
        /// </summary>
        [Input("destinationTable")]
        public Input<Inputs.JobQueryDestinationTableGetArgs>? DestinationTable { get; set; }

        /// <summary>
        /// If true and query uses legacy SQL dialect, flattens all nested and repeated fields in the query results.
        /// allowLargeResults must be true if this is set to false. For standard SQL queries, this flag is ignored and results are never flattened.
        /// </summary>
        [Input("flattenResults")]
        public Input<bool>? FlattenResults { get; set; }

        /// <summary>
        /// Limits the billing tier for this job. Queries that have resource usage beyond this tier will fail (without incurring a charge).
        /// If unspecified, this will be set to your project default.
        /// </summary>
        [Input("maximumBillingTier")]
        public Input<int>? MaximumBillingTier { get; set; }

        /// <summary>
        /// Limits the bytes billed for this job. Queries that will have bytes billed beyond this limit will fail (without incurring a charge).
        /// If unspecified, this will be set to your project default.
        /// </summary>
        [Input("maximumBytesBilled")]
        public Input<string>? MaximumBytesBilled { get; set; }

        /// <summary>
        /// Standard SQL only. Set to POSITIONAL to use positional (?) query parameters or to NAMED to use named (@myparam) query parameters in this query.
        /// </summary>
        [Input("parameterMode")]
        public Input<string>? ParameterMode { get; set; }

        /// <summary>
        /// Specifies a priority for the query.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// Configures a query job.  Structure is documented below.
        /// </summary>
        [Input("query", required: true)]
        public Input<string> Query { get; set; } = null!;

        [Input("schemaUpdateOptions")]
        private InputList<string>? _schemaUpdateOptions;

        /// <summary>
        /// Allows the schema of the destination table to be updated as a side effect of the load job if a schema is autodetected or
        /// supplied in the job configuration. Schema update options are supported in two cases: when writeDisposition is WRITE_APPEND;
        /// when writeDisposition is WRITE_TRUNCATE and the destination table is a partition of a table, specified by partition decorators.
        /// For normal tables, WRITE_TRUNCATE will always overwrite the schema. One or more of the following values are specified:
        /// ALLOW_FIELD_ADDITION: allow adding a nullable field to the schema.
        /// ALLOW_FIELD_RELAXATION: allow relaxing a required field in the original schema to nullable.
        /// </summary>
        public InputList<string> SchemaUpdateOptions
        {
            get => _schemaUpdateOptions ?? (_schemaUpdateOptions = new InputList<string>());
            set => _schemaUpdateOptions = value;
        }

        /// <summary>
        /// Options controlling the execution of scripts.  Structure is documented below.
        /// </summary>
        [Input("scriptOptions")]
        public Input<Inputs.JobQueryScriptOptionsGetArgs>? ScriptOptions { get; set; }

        /// <summary>
        /// Specifies whether to use BigQuery's legacy SQL dialect for this query. The default value is true.
        /// If set to false, the query will use BigQuery's standard SQL.
        /// </summary>
        [Input("useLegacySql")]
        public Input<bool>? UseLegacySql { get; set; }

        /// <summary>
        /// Whether to look for the result in the query cache. The query cache is a best-effort cache that will be flushed whenever
        /// tables in the query are modified. Moreover, the query cache is only available when a query does not have a destination table specified.
        /// The default value is true.
        /// </summary>
        [Input("useQueryCache")]
        public Input<bool>? UseQueryCache { get; set; }

        [Input("userDefinedFunctionResources")]
        private InputList<Inputs.JobQueryUserDefinedFunctionResourceGetArgs>? _userDefinedFunctionResources;

        /// <summary>
        /// Describes user-defined function resources used in the query.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.JobQueryUserDefinedFunctionResourceGetArgs> UserDefinedFunctionResources
        {
            get => _userDefinedFunctionResources ?? (_userDefinedFunctionResources = new InputList<Inputs.JobQueryUserDefinedFunctionResourceGetArgs>());
            set => _userDefinedFunctionResources = value;
        }

        /// <summary>
        /// Specifies the action that occurs if the destination table already exists. The following values are supported:
        /// WRITE_TRUNCATE: If the table already exists, BigQuery overwrites the table data and uses the schema from the query result.
        /// WRITE_APPEND: If the table already exists, BigQuery appends the data to the table.
        /// WRITE_EMPTY: If the table already exists and contains data, a 'duplicate' error is returned in the job result.
        /// Each action is atomic and only occurs if BigQuery is able to complete the job successfully.
        /// Creation, truncation and append actions occur as one atomic update upon job completion.
        /// </summary>
        [Input("writeDisposition")]
        public Input<string>? WriteDisposition { get; set; }

        public JobQueryGetArgs()
        {
        }
    }
}
