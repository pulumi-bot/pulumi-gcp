// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery.Inputs
{

    public sealed class TableRangePartitioningGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The field used to determine how to create a time-based
        /// partition. If time-based partitioning is enabled without this value, the
        /// table is partitioned based on the load time.
        /// </summary>
        [Input("field", required: true)]
        public Input<string> Field { get; set; } = null!;

        /// <summary>
        /// Range of a sheet to query from. Only used when
        /// non-empty. At least one of `range` or `skip_leading_rows` must be set.
        /// Typical format: "sheet_name!top_left_cell_id:bottom_right_cell_id"
        /// For example: "sheet1!A1:B20"
        /// </summary>
        [Input("range", required: true)]
        public Input<Inputs.TableRangePartitioningRangeGetArgs> Range { get; set; } = null!;

        public TableRangePartitioningGetArgs()
        {
        }
    }
}
