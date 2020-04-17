// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class URLMapHostRule
    {
        /// <summary>
        /// -
        /// (Optional)
        /// An optional description of this resource. Provide this property when you create
        /// the resource.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// -
        /// (Required)
        /// The list of host patterns to match. They must be valid hostnames, except * will
        /// match any string of ([a-z0-9-.]*). In that case, * must be the first character
        /// and must be followed in the pattern by either - or ..
        /// </summary>
        public readonly ImmutableArray<string> Hosts;
        /// <summary>
        /// -
        /// (Optional)
        /// The list of named PathMatchers to use against the URL.  Structure is documented below.
        /// </summary>
        public readonly string PathMatcher;

        [OutputConstructor]
        private URLMapHostRule(
            string? description,

            ImmutableArray<string> hosts,

            string pathMatcher)
        {
            Description = description;
            Hosts = hosts;
            PathMatcher = pathMatcher;
        }
    }
}
