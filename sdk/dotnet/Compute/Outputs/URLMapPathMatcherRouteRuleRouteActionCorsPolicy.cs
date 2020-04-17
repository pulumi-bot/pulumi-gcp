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
    public sealed class URLMapPathMatcherRouteRuleRouteActionCorsPolicy
    {
        /// <summary>
        /// -
        /// (Optional)
        /// In response to a preflight request, setting this to true indicates that the
        /// actual request can include user credentials. This translates to the Access-
        /// Control-Allow-Credentials header. Defaults to false.
        /// </summary>
        public readonly bool? AllowCredentials;
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies the content for the Access-Control-Allow-Headers header.
        /// </summary>
        public readonly ImmutableArray<string> AllowHeaders;
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies the content for the Access-Control-Allow-Methods header.
        /// </summary>
        public readonly ImmutableArray<string> AllowMethods;
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies the regualar expression patterns that match allowed origins. For
        /// regular expression grammar please see en.cppreference.com/w/cpp/regex/ecmascript
        /// An origin is allowed if it matches either allow_origins or allow_origin_regex.
        /// </summary>
        public readonly ImmutableArray<string> AllowOriginRegexes;
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies the list of origins that will be allowed to do CORS requests. An
        /// origin is allowed if it matches either allow_origins or allow_origin_regex.
        /// </summary>
        public readonly ImmutableArray<string> AllowOrigins;
        /// <summary>
        /// -
        /// (Required)
        /// If true, specifies the CORS policy is disabled.
        /// </summary>
        public readonly bool? Disabled;
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies the content for the Access-Control-Expose-Headers header.
        /// </summary>
        public readonly ImmutableArray<string> ExposeHeaders;
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies how long the results of a preflight request can be cached. This
        /// translates to the content for the Access-Control-Max-Age header.
        /// </summary>
        public readonly int? MaxAge;

        [OutputConstructor]
        private URLMapPathMatcherRouteRuleRouteActionCorsPolicy(
            bool? allowCredentials,

            ImmutableArray<string> allowHeaders,

            ImmutableArray<string> allowMethods,

            ImmutableArray<string> allowOriginRegexes,

            ImmutableArray<string> allowOrigins,

            bool? disabled,

            ImmutableArray<string> exposeHeaders,

            int? maxAge)
        {
            AllowCredentials = allowCredentials;
            AllowHeaders = allowHeaders;
            AllowMethods = allowMethods;
            AllowOriginRegexes = allowOriginRegexes;
            AllowOrigins = allowOrigins;
            Disabled = disabled;
            ExposeHeaders = exposeHeaders;
            MaxAge = maxAge;
        }
    }
}
