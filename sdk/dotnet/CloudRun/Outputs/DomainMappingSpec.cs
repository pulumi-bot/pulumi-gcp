// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun.Outputs
{

    [OutputType]
    public sealed class DomainMappingSpec
    {
        /// <summary>
        /// The mode of the certificate.
        /// </summary>
        public readonly string? CertificateMode;
        /// <summary>
        /// If set, the mapping will override any mapping set before this spec was set.
        /// It is recommended that the user leaves this empty to receive an error
        /// warning about a potential conflict and only set it once the respective UI
        /// has given such a warning.
        /// </summary>
        public readonly bool? ForceOverride;
        /// <summary>
        /// The name of the Cloud Run Service that this DomainMapping applies to.
        /// The route must exist.
        /// </summary>
        public readonly string RouteName;

        [OutputConstructor]
        private DomainMappingSpec(
            string? certificateMode,

            bool? forceOverride,

            string routeName)
        {
            CertificateMode = certificateMode;
            ForceOverride = forceOverride;
            RouteName = routeName;
        }
    }
}
