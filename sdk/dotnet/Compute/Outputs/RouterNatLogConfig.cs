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
    public sealed class RouterNatLogConfig
    {
        /// <summary>
        /// Indicates whether or not to export logs.
        /// </summary>
        public readonly bool Enable;
        /// <summary>
        /// Specifies the desired filtering of logs on this NAT. Valid
        /// values are: `"ERRORS_ONLY"`, `"TRANSLATIONS_ONLY"`, `"ALL"`
        /// </summary>
        public readonly string Filter;

        [OutputConstructor]
        private RouterNatLogConfig(
            bool enable,

            string filter)
        {
            Enable = enable;
            Filter = filter;
        }
    }
}
