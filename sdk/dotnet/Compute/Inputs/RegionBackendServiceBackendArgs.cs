// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionBackendServiceBackendArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the balancing mode for this backend. Defaults to CONNECTION.
        /// </summary>
        [Input("balancingMode")]
        public Input<string>? BalancingMode { get; set; }

        /// <summary>
        /// A multiplier applied to the group's maximum servicing capacity
        /// (based on UTILIZATION, RATE or CONNECTION).
        /// ~&gt;**NOTE**: This field cannot be set for
        /// INTERNAL region backend services (default loadBalancingScheme),
        /// but is required for non-INTERNAL backend service. The total
        /// capacity_scaler for all backends must be non-zero.
        /// A setting of 0 means the group is completely drained, offering
        /// 0% of its available Capacity. Valid range is [0.0,1.0].
        /// </summary>
        [Input("capacityScaler")]
        public Input<double>? CapacityScaler { get; set; }

        /// <summary>
        /// An optional description of this resource.
        /// Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// This field designates whether this is a failover backend. More
        /// than one failover backend can be configured for a given RegionBackendService.
        /// </summary>
        [Input("failover")]
        public Input<bool>? Failover { get; set; }

        /// <summary>
        /// The fully-qualified URL of an Instance Group or Network Endpoint
        /// Group resource. In case of instance group this defines the list
        /// of instances that serve traffic. Member virtual machine
        /// instances from each instance group must live in the same zone as
        /// the instance group itself. No two backends in a backend service
        /// are allowed to use same Instance Group resource.
        /// For Network Endpoint Groups this defines list of endpoints. All
        /// endpoints of Network Endpoint Group must be hosted on instances
        /// located in the same zone as the Network Endpoint Group.
        /// Backend services cannot mix Instance Group and
        /// Network Endpoint Group backends.
        /// When the `load_balancing_scheme` is INTERNAL, only instance groups
        /// are supported.
        /// Note that you must specify an Instance Group or Network Endpoint
        /// Group resource using the fully-qualified URL, rather than a
        /// partial URL.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        /// <summary>
        /// The maximum number of connections to the backend cluster.
        /// Defaults to 1024.
        /// </summary>
        [Input("maxConnections")]
        public Input<int>? MaxConnections { get; set; }

        /// <summary>
        /// The max number of simultaneous connections that a single backend
        /// network endpoint can handle. Cannot be set
        /// for INTERNAL backend services.
        /// This is used to calculate the capacity of the group. Can be
        /// used in either CONNECTION or UTILIZATION balancing modes. For
        /// CONNECTION mode, either maxConnections or
        /// maxConnectionsPerEndpoint must be set.
        /// </summary>
        [Input("maxConnectionsPerEndpoint")]
        public Input<int>? MaxConnectionsPerEndpoint { get; set; }

        /// <summary>
        /// The max number of simultaneous connections that a single
        /// backend instance can handle. Cannot be set for INTERNAL backend
        /// services.
        /// This is used to calculate the capacity of the group.
        /// Can be used in either CONNECTION or UTILIZATION balancing modes.
        /// For CONNECTION mode, either maxConnections or
        /// maxConnectionsPerInstance must be set.
        /// </summary>
        [Input("maxConnectionsPerInstance")]
        public Input<int>? MaxConnectionsPerInstance { get; set; }

        /// <summary>
        /// The max requests per second (RPS) of the group. Cannot be set
        /// for INTERNAL backend services.
        /// Can be used with either RATE or UTILIZATION balancing modes,
        /// but required if RATE mode. Either maxRate or one
        /// of maxRatePerInstance or maxRatePerEndpoint, as appropriate for
        /// group type, must be set.
        /// </summary>
        [Input("maxRate")]
        public Input<int>? MaxRate { get; set; }

        /// <summary>
        /// The max requests per second (RPS) that a single backend network
        /// endpoint can handle. This is used to calculate the capacity of
        /// the group. Can be used in either balancing mode. For RATE mode,
        /// either maxRate or maxRatePerEndpoint must be set. Cannot be set
        /// for INTERNAL backend services.
        /// </summary>
        [Input("maxRatePerEndpoint")]
        public Input<double>? MaxRatePerEndpoint { get; set; }

        /// <summary>
        /// The max requests per second (RPS) that a single backend
        /// instance can handle. This is used to calculate the capacity of
        /// the group. Can be used in either balancing mode. For RATE mode,
        /// either maxRate or maxRatePerInstance must be set. Cannot be set
        /// for INTERNAL backend services.
        /// </summary>
        [Input("maxRatePerInstance")]
        public Input<double>? MaxRatePerInstance { get; set; }

        /// <summary>
        /// Used when balancingMode is UTILIZATION. This ratio defines the
        /// CPU utilization target for the group. Valid range is [0.0, 1.0].
        /// Cannot be set for INTERNAL backend services.
        /// </summary>
        [Input("maxUtilization")]
        public Input<double>? MaxUtilization { get; set; }

        public RegionBackendServiceBackendArgs()
        {
        }
    }
}
