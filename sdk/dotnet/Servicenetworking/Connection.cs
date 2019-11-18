// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Servicenetworking
{
    /// <summary>
    /// Manages a private VPC connection with a GCP service provider. For more information see
    /// [the official documentation](https://cloud.google.com/vpc/docs/configure-private-services-access#creating-connection)
    /// and
    /// [API](https://cloud.google.com/service-infrastructure/docs/service-networking/reference/rest/v1/services.connections).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/service_networking_connection.html.markdown.
    /// </summary>
    public partial class Connection : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of VPC network connected with service producers using VPC peering.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        [Output("peering")]
        public Output<string> Peering { get; private set; } = null!;

        /// <summary>
        /// Named IP address range(s) of PEERING type reserved for
        /// this service provider. Note that invoking this method with a different range when connection
        /// is already established will not reallocate already provisioned service producer subnetworks.
        /// </summary>
        [Output("reservedPeeringRanges")]
        public Output<ImmutableArray<string>> ReservedPeeringRanges { get; private set; } = null!;

        /// <summary>
        /// Provider peering service that is managing peering connectivity for a
        /// service provider organization. For Google services that support this functionality it is
        /// 'servicenetworking.googleapis.com'.
        /// </summary>
        [Output("service")]
        public Output<string> Service { get; private set; } = null!;


        /// <summary>
        /// Create a Connection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Connection(string name, ConnectionArgs args, CustomResourceOptions? options = null)
            : base("gcp:servicenetworking/connection:Connection", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Connection(string name, Input<string> id, ConnectionState? state = null, CustomResourceOptions? options = null)
            : base("gcp:servicenetworking/connection:Connection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Connection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Connection Get(string name, Input<string> id, ConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new Connection(name, id, state, options);
        }
    }

    public sealed class ConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of VPC network connected with service producers using VPC peering.
        /// </summary>
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        [Input("reservedPeeringRanges", required: true)]
        private InputList<string>? _reservedPeeringRanges;

        /// <summary>
        /// Named IP address range(s) of PEERING type reserved for
        /// this service provider. Note that invoking this method with a different range when connection
        /// is already established will not reallocate already provisioned service producer subnetworks.
        /// </summary>
        public InputList<string> ReservedPeeringRanges
        {
            get => _reservedPeeringRanges ?? (_reservedPeeringRanges = new InputList<string>());
            set => _reservedPeeringRanges = value;
        }

        /// <summary>
        /// Provider peering service that is managing peering connectivity for a
        /// service provider organization. For Google services that support this functionality it is
        /// 'servicenetworking.googleapis.com'.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public ConnectionArgs()
        {
        }
    }

    public sealed class ConnectionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of VPC network connected with service producers using VPC peering.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        [Input("peering")]
        public Input<string>? Peering { get; set; }

        [Input("reservedPeeringRanges")]
        private InputList<string>? _reservedPeeringRanges;

        /// <summary>
        /// Named IP address range(s) of PEERING type reserved for
        /// this service provider. Note that invoking this method with a different range when connection
        /// is already established will not reallocate already provisioned service producer subnetworks.
        /// </summary>
        public InputList<string> ReservedPeeringRanges
        {
            get => _reservedPeeringRanges ?? (_reservedPeeringRanges = new InputList<string>());
            set => _reservedPeeringRanges = value;
        }

        /// <summary>
        /// Provider peering service that is managing peering connectivity for a
        /// service provider organization. For Google services that support this functionality it is
        /// 'servicenetworking.googleapis.com'.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        public ConnectionState()
        {
        }
    }
}
