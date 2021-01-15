// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring
{
    public static class GetClusterIstioService
    {
        /// <summary>
        /// A Monitoring Service is the root resource under which operational aspects of a
        /// generic service are accessible. A service is some discrete, autonomous, and
        /// network-accessible unit, designed to solve an individual concern
        /// 
        /// An Cluster Istio monitoring service is automatically created by GCP to monitor
        /// Cluster Istio services.
        /// 
        /// 
        /// To get more information about Service, see:
        /// 
        /// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services)
        /// * How-to Guides
        ///     * [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
        ///     * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### Monitoring Cluster Istio Service
        /// 
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var @default = Output.Create(Gcp.Monitoring.GetClusterIstioService.InvokeAsync(new Gcp.Monitoring.GetClusterIstioServiceArgs
        ///         {
        ///             ClusterName = "west",
        ///             Location = "us-west2-a",
        ///             ServiceName = "istio-policy",
        ///             ServiceNamespace = "istio-system",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetClusterIstioServiceResult> InvokeAsync(GetClusterIstioServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterIstioServiceResult>("gcp:monitoring/getClusterIstioService:getClusterIstioService", args ?? new GetClusterIstioServiceArgs(), options.WithVersion());
    }


    public sealed class GetClusterIstioServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Kubernetes cluster in which this Istio service 
        /// is defined. Corresponds to the clusterName resource label in k8s_cluster resources.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The location of the Kubernetes cluster in which this Istio service 
        /// is defined. Corresponds to the location resource label in k8s_cluster resources.
        /// </summary>
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The name of the Istio service underlying this service.
        /// Corresponds to the destination_service_name metric label in Istio metrics.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        /// <summary>
        /// The namespace of the Istio service underlying this service.
        /// Corresponds to the destination_service_namespace metric label in Istio metrics.
        /// </summary>
        [Input("serviceNamespace", required: true)]
        public string ServiceNamespace { get; set; } = null!;

        public GetClusterIstioServiceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClusterIstioServiceResult
    {
        public readonly string ClusterName;
        public readonly string DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Location;
        public readonly string Name;
        public readonly string? Project;
        public readonly string ServiceId;
        public readonly string ServiceName;
        public readonly string ServiceNamespace;
        public readonly ImmutableArray<Outputs.GetClusterIstioServiceTelemetryResult> Telemetries;

        [OutputConstructor]
        private GetClusterIstioServiceResult(
            string clusterName,

            string displayName,

            string id,

            string location,

            string name,

            string? project,

            string serviceId,

            string serviceName,

            string serviceNamespace,

            ImmutableArray<Outputs.GetClusterIstioServiceTelemetryResult> telemetries)
        {
            ClusterName = clusterName;
            DisplayName = displayName;
            Id = id;
            Location = location;
            Name = name;
            Project = project;
            ServiceId = serviceId;
            ServiceName = serviceName;
            ServiceNamespace = serviceNamespace;
            Telemetries = telemetries;
        }
    }
}