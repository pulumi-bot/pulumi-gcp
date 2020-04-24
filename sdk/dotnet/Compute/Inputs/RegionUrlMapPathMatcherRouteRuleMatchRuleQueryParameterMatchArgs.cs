// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The queryParameterMatch matches if the value of the parameter exactly matches
        /// the contents of exactMatch. Only one of presentMatch, exactMatch and regexMatch
        /// must be set.
        /// </summary>
        [Input("exactMatch")]
        public Input<string>? ExactMatch { get; set; }

        /// <summary>
        /// The name of the query parameter to match. The query parameter must exist in the
        /// request, in the absence of which the request match fails.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies that the queryParameterMatch matches if the request contains the query
        /// parameter, irrespective of whether the parameter has a value or not. Only one of
        /// presentMatch, exactMatch and regexMatch must be set.
        /// </summary>
        [Input("presentMatch")]
        public Input<bool>? PresentMatch { get; set; }

        /// <summary>
        /// The queryParameterMatch matches if the value of the parameter matches the
        /// regular expression specified by regexMatch. For the regular expression grammar,
        /// please see en.cppreference.com/w/cpp/regex/ecmascript  Only one of presentMatch,
        /// exactMatch and regexMatch must be set.
        /// </summary>
        [Input("regexMatch")]
        public Input<string>? RegexMatch { get; set; }

        public RegionUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchArgs()
        {
        }
    }
}
