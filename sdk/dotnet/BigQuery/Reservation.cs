// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery
{
    /// <summary>
    /// A reservation is a mechanism used to guarantee BigQuery slots to users.
    /// 
    /// To get more information about Reservation, see:
    /// 
    /// * [API documentation](https://cloud.google.com/bigquery/docs/reference/reservations/rest/v1beta1/projects.locations.reservations/create)
    /// * How-to Guides
    ///     * [Introduction to Reservations](https://cloud.google.com/bigquery/docs/reservations-intro)
    /// 
    /// ## Example Usage
    /// ### Bigquery Reservation Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var reservation = new Gcp.BigQuery.Reservation("reservation", new Gcp.BigQuery.ReservationArgs
    ///         {
    ///             IgnoreIdleSlots = false,
    ///             Location = "asia-northeast1",
    ///             SlotCapacity = 0,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Reservation can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquery/reservation:Reservation default projects/{{project}}/locations/{{location}}/reservations/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquery/reservation:Reservation default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquery/reservation:Reservation default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:bigquery/reservation:Reservation")]
    public partial class Reservation : Pulumi.CustomResource
    {
        /// <summary>
        /// If false, any query using this reservation will use idle slots from other reservations within
        /// the same admin project. If true, a query using this reservation will execute with the slot
        /// capacity specified above at most.
        /// </summary>
        [Output("ignoreIdleSlots")]
        public Output<bool?> IgnoreIdleSlots { get; private set; } = null!;

        /// <summary>
        /// The geographic location where the transfer config should reside.
        /// Examples: US, EU, asia-northeast1. The default value is US.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the reservation. This field must only contain alphanumeric characters or dash.
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
        /// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
        /// unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
        /// </summary>
        [Output("slotCapacity")]
        public Output<int> SlotCapacity { get; private set; } = null!;


        /// <summary>
        /// Create a Reservation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Reservation(string name, ReservationArgs args, CustomResourceOptions? options = null)
            : base("gcp:bigquery/reservation:Reservation", name, args ?? new ReservationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Reservation(string name, Input<string> id, ReservationState? state = null, CustomResourceOptions? options = null)
            : base("gcp:bigquery/reservation:Reservation", name, state, MakeResourceOptions(options, id))
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
        /// If false, any query using this reservation will use idle slots from other reservations within
        /// the same admin project. If true, a query using this reservation will execute with the slot
        /// capacity specified above at most.
        /// </summary>
        [Input("ignoreIdleSlots")]
        public Input<bool>? IgnoreIdleSlots { get; set; }

        /// <summary>
        /// The geographic location where the transfer config should reside.
        /// Examples: US, EU, asia-northeast1. The default value is US.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the reservation. This field must only contain alphanumeric characters or dash.
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
        /// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
        /// unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
        /// </summary>
        [Input("slotCapacity", required: true)]
        public Input<int> SlotCapacity { get; set; } = null!;

        public ReservationArgs()
        {
        }
    }

    public sealed class ReservationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If false, any query using this reservation will use idle slots from other reservations within
        /// the same admin project. If true, a query using this reservation will execute with the slot
        /// capacity specified above at most.
        /// </summary>
        [Input("ignoreIdleSlots")]
        public Input<bool>? IgnoreIdleSlots { get; set; }

        /// <summary>
        /// The geographic location where the transfer config should reside.
        /// Examples: US, EU, asia-northeast1. The default value is US.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the reservation. This field must only contain alphanumeric characters or dash.
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
        /// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
        /// unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
        /// </summary>
        [Input("slotCapacity")]
        public Input<int>? SlotCapacity { get; set; }

        public ReservationState()
        {
        }
    }
}
