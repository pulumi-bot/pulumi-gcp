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
    public sealed class SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig
    {
        /// <summary>
        /// ) If set to true, enables CAAP for L7 DDoS detection.
        /// </summary>
        public readonly bool? Enable;
        /// <summary>
        /// ) Rule visibility can be one of the following: STANDARD - opaque rules. (default) PREMIUM - transparent rules.
        /// </summary>
        public readonly string? RuleVisibility;

        [OutputConstructor]
        private SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig(
            bool? enable,

            string? ruleVisibility)
        {
            Enable = enable;
            RuleVisibility = ruleVisibility;
        }
    }
}