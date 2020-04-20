// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// The name of the query parameter to match. The query parameter must exist in the
        /// request, in the absence of which the request match fails.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// The value of the label must match the specified value. value can have a maximum
        /// length of 1024 characters.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelGetArgs()
        {
        }
    }
}
