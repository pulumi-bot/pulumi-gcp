// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies the value of the fixed delay interval.  Structure is documented below.
        /// </summary>
        [Input("fixedDelay", required: true)]
        public Input<Inputs.RegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayGetArgs> FixedDelay { get; set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The percentage of traffic (connections/operations/requests) which will be
        /// aborted as part of fault injection. The value must be between 0.0 and 100.0
        /// inclusive.
        /// </summary>
        [Input("percentage", required: true)]
        public Input<double> Percentage { get; set; } = null!;

        public RegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayGetArgs()
        {
        }
    }
}
