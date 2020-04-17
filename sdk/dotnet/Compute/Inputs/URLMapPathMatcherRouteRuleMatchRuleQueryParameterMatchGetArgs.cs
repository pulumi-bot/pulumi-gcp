// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class URLMapPathMatcherRouteRuleMatchRuleQueryParameterMatchGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The value should exactly match contents of exactMatch. Only one of exactMatch,
        /// prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch must be set.
        /// </summary>
        [Input("exactMatch")]
        public Input<string>? ExactMatch { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// Name of the resource. Provided by the client when the resource is created. The
        /// name must be 1-63 characters long, and comply with RFC1035. Specifically, the
        /// name must be 1-63 characters long and match the regular expression
        /// `a-z?` which means the first character must be a lowercase
        /// letter, and all following characters must be a dash, lowercase letter, or digit,
        /// except the last character, which cannot be a dash.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// A header with the contents of headerName must exist. The match takes place
        /// whether or not the request's header has a value or not. Only one of exactMatch,
        /// prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch must be set.
        /// </summary>
        [Input("presentMatch")]
        public Input<bool>? PresentMatch { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// For satifying the matchRule condition, the path of the request must satisfy the
        /// regular expression specified in regexMatch after removing any query parameters
        /// and anchor supplied with the original URL. For regular expression grammar please
        /// see en.cppreference.com/w/cpp/regex/ecmascript  Only one of prefixMatch,
        /// fullPathMatch or regexMatch must be specified.
        /// </summary>
        [Input("regexMatch")]
        public Input<string>? RegexMatch { get; set; }

        public URLMapPathMatcherRouteRuleMatchRuleQueryParameterMatchGetArgs()
        {
        }
    }
}
