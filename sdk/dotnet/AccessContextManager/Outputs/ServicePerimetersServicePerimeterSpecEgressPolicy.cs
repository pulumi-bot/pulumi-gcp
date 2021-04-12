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
    public sealed class ServicePerimetersServicePerimeterSpecEgressPolicy
    {
        /// <summary>
        /// / Defines conditions on the source of a request causing this `EgressPolicy` to apply.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.ServicePerimetersServicePerimeterSpecEgressPolicyEgressFrom? EgressFrom;
        /// <summary>
        /// / Defines the conditions on the `ApiOperation` and destination resources that cause this `EgressPolicy` to apply.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.ServicePerimetersServicePerimeterSpecEgressPolicyEgressTo? EgressTo;

        [OutputConstructor]
        private ServicePerimetersServicePerimeterSpecEgressPolicy(
            Outputs.ServicePerimetersServicePerimeterSpecEgressPolicyEgressFrom? egressFrom,

            Outputs.ServicePerimetersServicePerimeterSpecEgressPolicyEgressTo? egressTo)
        {
            EgressFrom = egressFrom;
            EgressTo = egressTo;
        }
    }
}