// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring.Inputs
{

    public sealed class AlertPolicyConditionConditionThresholdArgs : Pulumi.ResourceArgs
    {
        [Input("aggregations")]
        private InputList<Inputs.AlertPolicyConditionConditionThresholdAggregationArgs>? _aggregations;

        /// <summary>
        /// -
        /// (Optional)
        /// Specifies the alignment of data points in
        /// individual time series as well as how to
        /// combine the retrieved time series together
        /// (such as when aggregating multiple streams
        /// on each resource to a single stream for each
        /// resource or when aggregating streams across
        /// all members of a group of resources).
        /// Multiple aggregations are applied in the
        /// order specified.This field is similar to the
        /// one in the MetricService.ListTimeSeries
        /// request. It is advisable to use the
        /// ListTimeSeries method when debugging this
        /// field.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.AlertPolicyConditionConditionThresholdAggregationArgs> Aggregations
        {
            get => _aggregations ?? (_aggregations = new InputList<Inputs.AlertPolicyConditionConditionThresholdAggregationArgs>());
            set => _aggregations = value;
        }

        /// <summary>
        /// -
        /// (Required)
        /// The comparison to apply between the time
        /// series (indicated by filter and aggregation)
        /// and the threshold (indicated by
        /// threshold_value). The comparison is applied
        /// on each time series, with the time series on
        /// the left-hand side and the threshold on the
        /// right-hand side. Only COMPARISON_LT and
        /// COMPARISON_GT are supported currently.
        /// </summary>
        [Input("comparison", required: true)]
        public Input<string> Comparison { get; set; } = null!;

        [Input("denominatorAggregations")]
        private InputList<Inputs.AlertPolicyConditionConditionThresholdDenominatorAggregationArgs>? _denominatorAggregations;

        /// <summary>
        /// -
        /// (Optional)
        /// Specifies the alignment of data points in
        /// individual time series selected by
        /// denominatorFilter as well as how to combine
        /// the retrieved time series together (such as
        /// when aggregating multiple streams on each
        /// resource to a single stream for each
        /// resource or when aggregating streams across
        /// all members of a group of resources).When
        /// computing ratios, the aggregations and
        /// denominator_aggregations fields must use the
        /// same alignment period and produce time
        /// series that have the same periodicity and
        /// labels.This field is similar to the one in
        /// the MetricService.ListTimeSeries request. It
        /// is advisable to use the ListTimeSeries
        /// method when debugging this field.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.AlertPolicyConditionConditionThresholdDenominatorAggregationArgs> DenominatorAggregations
        {
            get => _denominatorAggregations ?? (_denominatorAggregations = new InputList<Inputs.AlertPolicyConditionConditionThresholdDenominatorAggregationArgs>());
            set => _denominatorAggregations = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// A filter that identifies a time series that
        /// should be used as the denominator of a ratio
        /// that will be compared with the threshold. If
        /// a denominator_filter is specified, the time
        /// series specified by the filter field will be
        /// used as the numerator.The filter is similar
        /// to the one that is specified in the
        /// MetricService.ListTimeSeries request (that
        /// call is useful to verify the time series
        /// that will be retrieved / processed) and must
        /// specify the metric type and optionally may
        /// contain restrictions on resource type,
        /// resource labels, and metric labels. This
        /// field may not exceed 2048 Unicode characters
        /// in length.
        /// </summary>
        [Input("denominatorFilter")]
        public Input<string>? DenominatorFilter { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// The amount of time that a time series must
        /// violate the threshold to be considered
        /// failing. Currently, only values that are a
        /// multiple of a minute--e.g., 0, 60, 120, or
        /// 300 seconds--are supported. If an invalid
        /// value is given, an error will be returned.
        /// When choosing a duration, it is useful to
        /// keep in mind the frequency of the underlying
        /// time series data (which may also be affected
        /// by any alignments specified in the
        /// aggregations field); a good duration is long
        /// enough so that a single outlier does not
        /// generate spurious alerts, but short enough
        /// that unhealthy states are detected and
        /// alerted on quickly.
        /// </summary>
        [Input("duration", required: true)]
        public Input<string> Duration { get; set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// A filter that identifies which time series
        /// should be compared with the threshold.The
        /// filter is similar to the one that is
        /// specified in the
        /// MetricService.ListTimeSeries request (that
        /// call is useful to verify the time series
        /// that will be retrieved / processed) and must
        /// specify the metric type and optionally may
        /// contain restrictions on resource type,
        /// resource labels, and metric labels. This
        /// field may not exceed 2048 Unicode characters
        /// in length.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// A value against which to compare the time
        /// series.
        /// </summary>
        [Input("thresholdValue")]
        public Input<double>? ThresholdValue { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The number/percent of time series for which
        /// the comparison must hold in order for the
        /// condition to trigger. If unspecified, then
        /// the condition will trigger if the comparison
        /// is true for any of the time series that have
        /// been identified by filter and aggregations,
        /// or by the ratio, if denominator_filter and
        /// denominator_aggregations are specified.  Structure is documented below.
        /// </summary>
        [Input("trigger")]
        public Input<Inputs.AlertPolicyConditionConditionThresholdTriggerArgs>? Trigger { get; set; }

        public AlertPolicyConditionConditionThresholdArgs()
        {
        }
    }
}
