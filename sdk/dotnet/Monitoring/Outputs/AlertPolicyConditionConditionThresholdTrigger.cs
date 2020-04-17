// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring.Outputs
{

    [OutputType]
    public sealed class AlertPolicyConditionConditionThresholdTrigger
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The absolute number of time series
        /// that must fail the predicate for the
        /// condition to be triggered.
        /// </summary>
        public readonly int? Count;
        /// <summary>
        /// -
        /// (Optional)
        /// The percentage of time series that
        /// must fail the predicate for the
        /// condition to be triggered.
        /// </summary>
        public readonly double? Percent;

        [OutputConstructor]
        private AlertPolicyConditionConditionThresholdTrigger(
            int? count,

            double? percent)
        {
            Count = count;
            Percent = percent;
        }
    }
}
