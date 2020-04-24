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
    public sealed class AccessLevelBasicConditionDevicePolicyOsConstraint
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The minimum allowed OS version. If not set, any version
        /// of this OS satisfies the constraint.
        /// Format: "major.minor.patch" such as "10.5.301", "9.2.1".
        /// </summary>
        public readonly string? MinimumVersion;
        /// <summary>
        /// -
        /// (Required)
        /// The operating system type of the device.
        /// </summary>
        public readonly string OsType;

        [OutputConstructor]
        private AccessLevelBasicConditionDevicePolicyOsConstraint(
            string? minimumVersion,

            string osType)
        {
            MinimumVersion = minimumVersion;
            OsType = osType;
        }
    }
}
