// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public partial class NetworkPeeringRoutesConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to export the custom routes to the peer network.
        /// </summary>
        [Output("exportCustomRoutes")]
        public Output<bool> ExportCustomRoutes { get; private set; } = null!;

        /// <summary>
        /// Whether to import the custom routes to the peer network.
        /// </summary>
        [Output("importCustomRoutes")]
        public Output<bool> ImportCustomRoutes { get; private set; } = null!;

        /// <summary>
        /// The name of the primary network for the peering.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// Name of the peering.
        /// </summary>
        [Output("peering")]
        public Output<string> Peering { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkPeeringRoutesConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkPeeringRoutesConfig(string name, NetworkPeeringRoutesConfigArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private NetworkPeeringRoutesConfig(string name, Input<string> id, NetworkPeeringRoutesConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkPeeringRoutesConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkPeeringRoutesConfig Get(string name, Input<string> id, NetworkPeeringRoutesConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkPeeringRoutesConfig(name, id, state, options);
        }
    }

    public sealed class NetworkPeeringRoutesConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to export the custom routes to the peer network.
        /// </summary>
        [Input("exportCustomRoutes", required: true)]
        public Input<bool> ExportCustomRoutes { get; set; } = null!;

        /// <summary>
        /// Whether to import the custom routes to the peer network.
        /// </summary>
        [Input("importCustomRoutes", required: true)]
        public Input<bool> ImportCustomRoutes { get; set; } = null!;

        /// <summary>
        /// The name of the primary network for the peering.
        /// </summary>
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        /// <summary>
        /// Name of the peering.
        /// </summary>
        [Input("peering", required: true)]
        public Input<string> Peering { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public NetworkPeeringRoutesConfigArgs()
        {
        }
    }

    public sealed class NetworkPeeringRoutesConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to export the custom routes to the peer network.
        /// </summary>
        [Input("exportCustomRoutes")]
        public Input<bool>? ExportCustomRoutes { get; set; }

        /// <summary>
        /// Whether to import the custom routes to the peer network.
        /// </summary>
        [Input("importCustomRoutes")]
        public Input<bool>? ImportCustomRoutes { get; set; }

        /// <summary>
        /// The name of the primary network for the peering.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// Name of the peering.
        /// </summary>
        [Input("peering")]
        public Input<string>? Peering { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public NetworkPeeringRoutesConfigState()
        {
        }
    }
}
