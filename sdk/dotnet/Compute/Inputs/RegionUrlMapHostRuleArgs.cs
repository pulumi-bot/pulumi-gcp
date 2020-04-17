// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionUrlMapHostRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("hosts", required: true)]
        private InputList<string>? _hosts;

        /// <summary>
        /// -
        /// (Required)
        /// The list of host patterns to match. They must be valid
        /// hostnames, except * will match any string of ([a-z0-9-.]*). In
        /// that case, * must be the first character and must be followed in
        /// the pattern by either - or ..
        /// </summary>
        public InputList<string> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<string>());
            set => _hosts = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// The list of named PathMatchers to use against the URL.  Structure is documented below.
        /// </summary>
        [Input("pathMatcher", required: true)]
        public Input<string> PathMatcher { get; set; } = null!;

        public RegionUrlMapHostRuleArgs()
        {
        }
    }
}
