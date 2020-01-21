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
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_ha_vpn_gateway.html.markdown.
    /// </summary>
    public partial class HaVpnGateway : Pulumi.CustomResource
    {
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
        /// The network this VPN gateway is accepting traffic for.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The region this gateway should sit in.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// A list of interfaces on this VPN gateway.
        /// </summary>
        [Output("vpnInterfaces")]
        public Output<ImmutableArray<Outputs.HaVpnGatewayVpnInterfaces>> VpnInterfaces { get; private set; } = null!;


        /// <summary>
        /// Create a HaVpnGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HaVpnGateway(string name, HaVpnGatewayArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/haVpnGateway:HaVpnGateway", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private HaVpnGateway(string name, Input<string> id, HaVpnGatewayState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/haVpnGateway:HaVpnGateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HaVpnGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HaVpnGateway Get(string name, Input<string> id, HaVpnGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new HaVpnGateway(name, id, state, options);
        }
    }

    public sealed class HaVpnGatewayArgs : Pulumi.ResourceArgs
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
        /// The network this VPN gateway is accepting traffic for.
        /// </summary>
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region this gateway should sit in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public HaVpnGatewayArgs()
        {
        }
    }

    public sealed class HaVpnGatewayState : Pulumi.ResourceArgs
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
        /// The network this VPN gateway is accepting traffic for.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region this gateway should sit in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        [Input("vpnInterfaces")]
        private InputList<Inputs.HaVpnGatewayVpnInterfacesGetArgs>? _vpnInterfaces;

        /// <summary>
        /// A list of interfaces on this VPN gateway.
        /// </summary>
        public InputList<Inputs.HaVpnGatewayVpnInterfacesGetArgs> VpnInterfaces
        {
            get => _vpnInterfaces ?? (_vpnInterfaces = new InputList<Inputs.HaVpnGatewayVpnInterfacesGetArgs>());
            set => _vpnInterfaces = value;
        }

        public HaVpnGatewayState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class HaVpnGatewayVpnInterfacesGetArgs : Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        public HaVpnGatewayVpnInterfacesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class HaVpnGatewayVpnInterfaces
    {
        public readonly int? Id;
        public readonly string? IpAddress;

        [OutputConstructor]
        private HaVpnGatewayVpnInterfaces(
            int? id,
            string? ipAddress)
        {
            Id = id;
            IpAddress = ipAddress;
        }
    }
    }
}
