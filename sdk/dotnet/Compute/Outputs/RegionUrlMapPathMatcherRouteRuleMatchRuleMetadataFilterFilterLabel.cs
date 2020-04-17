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
    public sealed class RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel
    {
        /// <summary>
        /// -
        /// (Required)
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// -
        /// (Required)
        /// The value of the label must match the specified value. value can have a maximum
        /// length of 1024 characters.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private RegionUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
