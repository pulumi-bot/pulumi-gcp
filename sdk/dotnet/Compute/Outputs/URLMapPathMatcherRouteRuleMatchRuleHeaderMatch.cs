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
    public sealed class URLMapPathMatcherRouteRuleMatchRuleHeaderMatch
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The queryParameterMatch matches if the value of the parameter exactly matches
        /// the contents of exactMatch. Only one of presentMatch, exactMatch and regexMatch
        /// must be set.
        /// </summary>
        public readonly string? ExactMatch;
        /// <summary>
        /// -
        /// (Required)
        /// The name of the header.
        /// </summary>
        public readonly string HeaderName;
        /// <summary>
        /// -
        /// (Optional)
        /// If set to false, the headerMatch is considered a match if the match criteria
        /// above are met. If set to true, the headerMatch is considered a match if the
        /// match criteria above are NOT met. Defaults to false.
        /// </summary>
        public readonly bool? InvertMatch;
        /// <summary>
        /// -
        /// (Optional)
        /// The value of the header must start with the contents of prefixMatch. Only one of
        /// exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch
        /// must be set.
        /// </summary>
        public readonly string? PrefixMatch;
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies that the queryParameterMatch matches if the request contains the query
        /// parameter, irrespective of whether the parameter has a value or not. Only one of
        /// presentMatch, exactMatch and regexMatch must be set.
        /// </summary>
        public readonly bool? PresentMatch;
        /// <summary>
        /// -
        /// (Optional)
        /// The header value must be an integer and its value must be in the range specified
        /// in rangeMatch. If the header does not contain an integer, number or is empty,
        /// the match fails. For example for a range [-5, 0]   - -3 will match.  - 0 will
        /// not match.  - 0.25 will not match.  - -3someString will not match.   Only one of
        /// exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch
        /// must be set.  Structure is documented below.
        /// </summary>
        public readonly Outputs.URLMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch? RangeMatch;
        /// <summary>
        /// -
        /// (Optional)
        /// The queryParameterMatch matches if the value of the parameter matches the
        /// regular expression specified by regexMatch. For the regular expression grammar,
        /// please see en.cppreference.com/w/cpp/regex/ecmascript  Only one of presentMatch,
        /// exactMatch and regexMatch must be set.
        /// </summary>
        public readonly string? RegexMatch;
        /// <summary>
        /// -
        /// (Optional)
        /// The value of the header must end with the contents of suffixMatch. Only one of
        /// exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch
        /// must be set.
        /// </summary>
        public readonly string? SuffixMatch;

        [OutputConstructor]
        private URLMapPathMatcherRouteRuleMatchRuleHeaderMatch(
            string? exactMatch,

            string headerName,

            bool? invertMatch,

            string? prefixMatch,

            bool? presentMatch,

            Outputs.URLMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch? rangeMatch,

            string? regexMatch,

            string? suffixMatch)
        {
            ExactMatch = exactMatch;
            HeaderName = headerName;
            InvertMatch = invertMatch;
            PrefixMatch = prefixMatch;
            PresentMatch = presentMatch;
            RangeMatch = rangeMatch;
            RegexMatch = regexMatch;
            SuffixMatch = suffixMatch;
        }
    }
}
