// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_network_endpoint.html.markdown.
    /// </summary>
    public partial class NetworkEndpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// The name for a specific VM instance that the IP address belongs to. This is required for network endpoints
        /// of type GCE_VM_IP_PORT. The instance must be in the same zone of network endpoint group.
        /// </summary>
        [Output("instance")]
        public Output<string> Instance { get; private set; } = null!;

        /// <summary>
        /// IPv4 address of network endpoint. The IP address must belong to a VM in GCE (either the primary IP or as
        /// part of an aliased IP range).
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// The network endpoint group this endpoint is part of.
        /// </summary>
        [Output("networkEndpointGroup")]
        public Output<string> NetworkEndpointGroup { get; private set; } = null!;

        /// <summary>
        /// Port number of network endpoint.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Zone where the containing network endpoint group is located.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkEndpoint(string name, NetworkEndpointArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/networkEndpoint:NetworkEndpoint", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private NetworkEndpoint(string name, Input<string> id, NetworkEndpointState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/networkEndpoint:NetworkEndpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkEndpoint Get(string name, Input<string> id, NetworkEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkEndpoint(name, id, state, options);
        }
    }

    public sealed class NetworkEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name for a specific VM instance that the IP address belongs to. This is required for network endpoints
        /// of type GCE_VM_IP_PORT. The instance must be in the same zone of network endpoint group.
        /// </summary>
        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        /// <summary>
        /// IPv4 address of network endpoint. The IP address must belong to a VM in GCE (either the primary IP or as
        /// part of an aliased IP range).
        /// </summary>
        [Input("ipAddress", required: true)]
        public Input<string> IpAddress { get; set; } = null!;

        /// <summary>
        /// The network endpoint group this endpoint is part of.
        /// </summary>
        [Input("networkEndpointGroup", required: true)]
        public Input<string> NetworkEndpointGroup { get; set; } = null!;

        /// <summary>
        /// Port number of network endpoint.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Zone where the containing network endpoint group is located.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public NetworkEndpointArgs()
        {
        }
    }

    public sealed class NetworkEndpointState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name for a specific VM instance that the IP address belongs to. This is required for network endpoints
        /// of type GCE_VM_IP_PORT. The instance must be in the same zone of network endpoint group.
        /// </summary>
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        /// <summary>
        /// IPv4 address of network endpoint. The IP address must belong to a VM in GCE (either the primary IP or as
        /// part of an aliased IP range).
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The network endpoint group this endpoint is part of.
        /// </summary>
        [Input("networkEndpointGroup")]
        public Input<string>? NetworkEndpointGroup { get; set; }

        /// <summary>
        /// Port number of network endpoint.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Zone where the containing network endpoint group is located.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public NetworkEndpointState()
        {
        }
    }
}
