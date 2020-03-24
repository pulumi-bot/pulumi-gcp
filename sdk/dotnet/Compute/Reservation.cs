// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// Represents a reservation resource. A reservation ensures that capacity is
    /// held in a specific zone even if the reserved VMs are not running.
    /// 
    /// Reservations apply only to Compute Engine, Cloud Dataproc, and Google
    /// Kubernetes Engine VM usage.Reservations do not apply to `f1-micro` or
    /// `g1-small` machine types, preemptible VMs, sole tenant nodes, or other
    /// services not listed above
    /// like Cloud SQL and Dataflow.
    /// 
    /// 
    /// To get more information about Reservation, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/reservations)
    /// * How-to Guides
    ///     * [Reserving zonal resources](https://cloud.google.com/compute/docs/instances/reserving-zonal-resources)
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_reservation.html.markdown.
    /// </summary>
    public partial class Reservation : Pulumi.CustomResource
    {
        /// <summary>
        /// Full or partial URL to a parent commitment. This field displays for reservations that are tied to a
        /// commitment.
        /// </summary>
        [Output("commitment")]
        public Output<string> Commitment { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters
        /// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
        /// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be
        /// a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Reservation for instances with specific machine shapes.
        /// </summary>
        [Output("specificReservation")]
        public Output<Outputs.ReservationSpecificReservation> SpecificReservation { get; private set; } = null!;

        /// <summary>
        /// When set to true, only VMs that target this reservation by name can consume this reservation. Otherwise, it
        /// can be consumed by VMs with affinity for any reservation. Defaults to false.
        /// </summary>
        [Output("specificReservationRequired")]
        public Output<bool?> SpecificReservationRequired { get; private set; } = null!;

        /// <summary>
        /// The status of the reservation.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The zone where the reservation is made.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Reservation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Reservation(string name, ReservationArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/reservation:Reservation", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Reservation(string name, Input<string> id, ReservationState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/reservation:Reservation", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Reservation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Reservation Get(string name, Input<string> id, ReservationState? state = null, CustomResourceOptions? options = null)
        {
            return new Reservation(name, id, state, options);
        }
    }

    public sealed class ReservationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters
        /// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
        /// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be
        /// a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Reservation for instances with specific machine shapes.
        /// </summary>
        [Input("specificReservation", required: true)]
        public Input<Inputs.ReservationSpecificReservationArgs> SpecificReservation { get; set; } = null!;

        /// <summary>
        /// When set to true, only VMs that target this reservation by name can consume this reservation. Otherwise, it
        /// can be consumed by VMs with affinity for any reservation. Defaults to false.
        /// </summary>
        [Input("specificReservationRequired")]
        public Input<bool>? SpecificReservationRequired { get; set; }

        /// <summary>
        /// The zone where the reservation is made.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public ReservationArgs()
        {
        }
    }

    public sealed class ReservationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Full or partial URL to a parent commitment. This field displays for reservations that are tied to a
        /// commitment.
        /// </summary>
        [Input("commitment")]
        public Input<string>? Commitment { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters
        /// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
        /// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be
        /// a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// Reservation for instances with specific machine shapes.
        /// </summary>
        [Input("specificReservation")]
        public Input<Inputs.ReservationSpecificReservationGetArgs>? SpecificReservation { get; set; }

        /// <summary>
        /// When set to true, only VMs that target this reservation by name can consume this reservation. Otherwise, it
        /// can be consumed by VMs with affinity for any reservation. Defaults to false.
        /// </summary>
        [Input("specificReservationRequired")]
        public Input<bool>? SpecificReservationRequired { get; set; }

        /// <summary>
        /// The status of the reservation.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The zone where the reservation is made.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ReservationState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ReservationSpecificReservationArgs : Pulumi.ResourceArgs
    {
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        [Input("inUseCount")]
        public Input<int>? InUseCount { get; set; }

        [Input("instanceProperties", required: true)]
        public Input<ReservationSpecificReservationInstancePropertiesArgs> InstanceProperties { get; set; } = null!;

        public ReservationSpecificReservationArgs()
        {
        }
    }

    public sealed class ReservationSpecificReservationGetArgs : Pulumi.ResourceArgs
    {
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        [Input("inUseCount")]
        public Input<int>? InUseCount { get; set; }

        [Input("instanceProperties", required: true)]
        public Input<ReservationSpecificReservationInstancePropertiesGetArgs> InstanceProperties { get; set; } = null!;

        public ReservationSpecificReservationGetArgs()
        {
        }
    }

    public sealed class ReservationSpecificReservationInstancePropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("guestAccelerators")]
        private InputList<ReservationSpecificReservationInstancePropertiesGuestAcceleratorsArgs>? _guestAccelerators;
        public InputList<ReservationSpecificReservationInstancePropertiesGuestAcceleratorsArgs> GuestAccelerators
        {
            get => _guestAccelerators ?? (_guestAccelerators = new InputList<ReservationSpecificReservationInstancePropertiesGuestAcceleratorsArgs>());
            set => _guestAccelerators = value;
        }

        [Input("localSsds")]
        private InputList<ReservationSpecificReservationInstancePropertiesLocalSsdsArgs>? _localSsds;
        public InputList<ReservationSpecificReservationInstancePropertiesLocalSsdsArgs> LocalSsds
        {
            get => _localSsds ?? (_localSsds = new InputList<ReservationSpecificReservationInstancePropertiesLocalSsdsArgs>());
            set => _localSsds = value;
        }

        [Input("machineType", required: true)]
        public Input<string> MachineType { get; set; } = null!;

        [Input("minCpuPlatform")]
        public Input<string>? MinCpuPlatform { get; set; }

        public ReservationSpecificReservationInstancePropertiesArgs()
        {
        }
    }

    public sealed class ReservationSpecificReservationInstancePropertiesGetArgs : Pulumi.ResourceArgs
    {
        [Input("guestAccelerators")]
        private InputList<ReservationSpecificReservationInstancePropertiesGuestAcceleratorsGetArgs>? _guestAccelerators;
        public InputList<ReservationSpecificReservationInstancePropertiesGuestAcceleratorsGetArgs> GuestAccelerators
        {
            get => _guestAccelerators ?? (_guestAccelerators = new InputList<ReservationSpecificReservationInstancePropertiesGuestAcceleratorsGetArgs>());
            set => _guestAccelerators = value;
        }

        [Input("localSsds")]
        private InputList<ReservationSpecificReservationInstancePropertiesLocalSsdsGetArgs>? _localSsds;
        public InputList<ReservationSpecificReservationInstancePropertiesLocalSsdsGetArgs> LocalSsds
        {
            get => _localSsds ?? (_localSsds = new InputList<ReservationSpecificReservationInstancePropertiesLocalSsdsGetArgs>());
            set => _localSsds = value;
        }

        [Input("machineType", required: true)]
        public Input<string> MachineType { get; set; } = null!;

        [Input("minCpuPlatform")]
        public Input<string>? MinCpuPlatform { get; set; }

        public ReservationSpecificReservationInstancePropertiesGetArgs()
        {
        }
    }

    public sealed class ReservationSpecificReservationInstancePropertiesGuestAcceleratorsArgs : Pulumi.ResourceArgs
    {
        [Input("acceleratorCount", required: true)]
        public Input<int> AcceleratorCount { get; set; } = null!;

        [Input("acceleratorType", required: true)]
        public Input<string> AcceleratorType { get; set; } = null!;

        public ReservationSpecificReservationInstancePropertiesGuestAcceleratorsArgs()
        {
        }
    }

    public sealed class ReservationSpecificReservationInstancePropertiesGuestAcceleratorsGetArgs : Pulumi.ResourceArgs
    {
        [Input("acceleratorCount", required: true)]
        public Input<int> AcceleratorCount { get; set; } = null!;

        [Input("acceleratorType", required: true)]
        public Input<string> AcceleratorType { get; set; } = null!;

        public ReservationSpecificReservationInstancePropertiesGuestAcceleratorsGetArgs()
        {
        }
    }

    public sealed class ReservationSpecificReservationInstancePropertiesLocalSsdsArgs : Pulumi.ResourceArgs
    {
        [Input("diskSizeGb", required: true)]
        public Input<int> DiskSizeGb { get; set; } = null!;

        [Input("interface")]
        public Input<string>? Interface { get; set; }

        public ReservationSpecificReservationInstancePropertiesLocalSsdsArgs()
        {
        }
    }

    public sealed class ReservationSpecificReservationInstancePropertiesLocalSsdsGetArgs : Pulumi.ResourceArgs
    {
        [Input("diskSizeGb", required: true)]
        public Input<int> DiskSizeGb { get; set; } = null!;

        [Input("interface")]
        public Input<string>? Interface { get; set; }

        public ReservationSpecificReservationInstancePropertiesLocalSsdsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ReservationSpecificReservation
    {
        public readonly int Count;
        public readonly int InUseCount;
        public readonly ReservationSpecificReservationInstanceProperties InstanceProperties;

        [OutputConstructor]
        private ReservationSpecificReservation(
            int count,
            int inUseCount,
            ReservationSpecificReservationInstanceProperties instanceProperties)
        {
            Count = count;
            InUseCount = inUseCount;
            InstanceProperties = instanceProperties;
        }
    }

    [OutputType]
    public sealed class ReservationSpecificReservationInstanceProperties
    {
        public readonly ImmutableArray<ReservationSpecificReservationInstancePropertiesGuestAccelerators> GuestAccelerators;
        public readonly ImmutableArray<ReservationSpecificReservationInstancePropertiesLocalSsds> LocalSsds;
        public readonly string MachineType;
        public readonly string MinCpuPlatform;

        [OutputConstructor]
        private ReservationSpecificReservationInstanceProperties(
            ImmutableArray<ReservationSpecificReservationInstancePropertiesGuestAccelerators> guestAccelerators,
            ImmutableArray<ReservationSpecificReservationInstancePropertiesLocalSsds> localSsds,
            string machineType,
            string minCpuPlatform)
        {
            GuestAccelerators = guestAccelerators;
            LocalSsds = localSsds;
            MachineType = machineType;
            MinCpuPlatform = minCpuPlatform;
        }
    }

    [OutputType]
    public sealed class ReservationSpecificReservationInstancePropertiesGuestAccelerators
    {
        public readonly int AcceleratorCount;
        public readonly string AcceleratorType;

        [OutputConstructor]
        private ReservationSpecificReservationInstancePropertiesGuestAccelerators(
            int acceleratorCount,
            string acceleratorType)
        {
            AcceleratorCount = acceleratorCount;
            AcceleratorType = acceleratorType;
        }
    }

    [OutputType]
    public sealed class ReservationSpecificReservationInstancePropertiesLocalSsds
    {
        public readonly int DiskSizeGb;
        public readonly string? Interface;

        [OutputConstructor]
        private ReservationSpecificReservationInstancePropertiesLocalSsds(
            int diskSizeGb,
            string? @interface)
        {
            DiskSizeGb = diskSizeGb;
            Interface = @interface;
        }
    }
    }
}
