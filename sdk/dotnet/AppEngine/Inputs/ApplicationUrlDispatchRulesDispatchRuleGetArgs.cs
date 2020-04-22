// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Inputs
{

    public sealed class ApplicationUrlDispatchRulesDispatchRuleGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Domain name to match against. The wildcard "*" is supported if specified before a period: "*.".
        /// Defaults to matching all domains: "*".
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
        /// The sum of the lengths of the domain and path may not exceed 100 characters.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
        /// The sum of the lengths of the domain and path may not exceed 100 characters.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public ApplicationUrlDispatchRulesDispatchRuleGetArgs()
        {
        }
    }
}
