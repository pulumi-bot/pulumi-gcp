// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class URLMapHostRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of this test case.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("hosts", required: true)]
        private InputList<string>? _hosts;

        /// <summary>
        /// The list of host patterns to match. They must be valid hostnames, except * will
        /// match any string of ([a-z0-9-.]*). In that case, * must be the first character
        /// and must be followed in the pattern by either - or ..
        /// </summary>
        public InputList<string> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<string>());
            set => _hosts = value;
        }

        /// <summary>
        /// The name of the PathMatcher to use to match the path portion of the URL if the
        /// hostRule matches the URL's host portion.
        /// </summary>
        [Input("pathMatcher", required: true)]
        public Input<string> PathMatcher { get; set; } = null!;

        public URLMapHostRuleArgs()
        {
        }
    }
}
