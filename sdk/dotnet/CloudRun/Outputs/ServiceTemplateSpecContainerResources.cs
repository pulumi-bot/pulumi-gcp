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
    public sealed class ServiceTemplateSpecContainerResources
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Limits describes the maximum amount of compute resources allowed.
        /// The values of the map is string form of the 'quantity' k8s type:
        /// https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Limits;
        /// <summary>
        /// -
        /// (Optional)
        /// Requests describes the minimum amount of compute resources required.
        /// If Requests is omitted for a container, it defaults to Limits if that is
        /// explicitly specified, otherwise to an implementation-defined value.
        /// The values of the map is string form of the 'quantity' k8s type:
        /// https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Requests;

        [OutputConstructor]
        private ServiceTemplateSpecContainerResources(
            ImmutableDictionary<string, string>? limits,

            ImmutableDictionary<string, string>? requests)
        {
            Limits = limits;
            Requests = requests;
        }
    }
}
