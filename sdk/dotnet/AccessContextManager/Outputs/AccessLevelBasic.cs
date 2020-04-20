// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AccessContextManager.Outputs
{

    [OutputType]
    public sealed class AccessLevelBasic
    {
        /// <summary>
        /// -
        /// (Optional)
        /// How the conditions list should be combined to determine if a request
        /// is granted this AccessLevel. If AND is used, each Condition in
        /// conditions must be satisfied for the AccessLevel to be applied. If
        /// OR is used, at least one Condition in conditions must be satisfied
        /// for the AccessLevel to be applied. Defaults to AND if unspecified.
        /// </summary>
        public readonly string? CombiningFunction;
        /// <summary>
        /// -
        /// (Required)
        /// A set of requirements for the AccessLevel to be granted.  Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.AccessLevelBasicCondition> Conditions;

        [OutputConstructor]
        private AccessLevelBasic(
            string? combiningFunction,

            ImmutableArray<Outputs.AccessLevelBasicCondition> conditions)
        {
            CombiningFunction = combiningFunction;
            Conditions = conditions;
        }
    }
}
