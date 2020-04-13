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
    /// Represents a VPN gateway managed outside of GCP.
    /// 
    /// To get more information about ExternalVpnGateway, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/externalVpnGateways)
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_external_vpn_gateway.html.markdown.
    /// </summary>
    public partial class ExternalVpnGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// -
        /// (Optional)
        /// An optional description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// A list of interfaces on this external VPN gateway.  Structure is documented below.
        /// </summary>
        [Output("interfaces")]
        public Output<ImmutableArray<Outputs.ExternalVpnGatewayInterfaces>> Interfaces { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035.  Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
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
        /// -
        /// (Optional)
        /// Indicates the redundancy type of this external VPN gateway
        /// </summary>
        [Output("redundancyType")]
        public Output<string?> RedundancyType { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a ExternalVpnGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExternalVpnGateway(string name, ExternalVpnGatewayArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:compute/externalVpnGateway:ExternalVpnGateway", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ExternalVpnGateway(string name, Input<string> id, ExternalVpnGatewayState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/externalVpnGateway:ExternalVpnGateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ExternalVpnGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExternalVpnGateway Get(string name, Input<string> id, ExternalVpnGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new ExternalVpnGateway(name, id, state, options);
        }
    }

    public sealed class ExternalVpnGatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("interfaces")]
        private InputList<Inputs.ExternalVpnGatewayInterfacesArgs>? _interfaces;

        /// <summary>
        /// -
        /// (Optional)
        /// A list of interfaces on this external VPN gateway.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.ExternalVpnGatewayInterfacesArgs> Interfaces
        {
            get => _interfaces ?? (_interfaces = new InputList<Inputs.ExternalVpnGatewayInterfacesArgs>());
            set => _interfaces = value;
        }

        /// <summary>
        /// -
        /// (Required)
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035.  Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
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
        /// -
        /// (Optional)
        /// Indicates the redundancy type of this external VPN gateway
        /// </summary>
        [Input("redundancyType")]
        public Input<string>? RedundancyType { get; set; }

        public ExternalVpnGatewayArgs()
        {
        }
    }

    public sealed class ExternalVpnGatewayState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("interfaces")]
        private InputList<Inputs.ExternalVpnGatewayInterfacesGetArgs>? _interfaces;

        /// <summary>
        /// -
        /// (Optional)
        /// A list of interfaces on this external VPN gateway.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.ExternalVpnGatewayInterfacesGetArgs> Interfaces
        {
            get => _interfaces ?? (_interfaces = new InputList<Inputs.ExternalVpnGatewayInterfacesGetArgs>());
            set => _interfaces = value;
        }

        /// <summary>
        /// -
        /// (Required)
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035.  Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
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
        /// -
        /// (Optional)
        /// Indicates the redundancy type of this external VPN gateway
        /// </summary>
        [Input("redundancyType")]
        public Input<string>? RedundancyType { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public ExternalVpnGatewayState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ExternalVpnGatewayInterfacesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The numberic ID for this interface. Allowed values are based on the redundancy type
        /// of this external VPN gateway
        /// * `0 - SINGLE_IP_INTERNALLY_REDUNDANT`
        /// * `0, 1 - TWO_IPS_REDUNDANCY`
        /// * `0, 1, 2, 3 - FOUR_IPS_REDUNDANCY`
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// IP address of the interface in the external VPN gateway.
        /// Only IPv4 is supported. This IP address can be either from
        /// your on-premise gateway or another Cloud provider’s VPN gateway,
        /// it cannot be an IP address from Google Compute Engine.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        public ExternalVpnGatewayInterfacesArgs()
        {
        }
    }

    public sealed class ExternalVpnGatewayInterfacesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The numberic ID for this interface. Allowed values are based on the redundancy type
        /// of this external VPN gateway
        /// * `0 - SINGLE_IP_INTERNALLY_REDUNDANT`
        /// * `0, 1 - TWO_IPS_REDUNDANCY`
        /// * `0, 1, 2, 3 - FOUR_IPS_REDUNDANCY`
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// IP address of the interface in the external VPN gateway.
        /// Only IPv4 is supported. This IP address can be either from
        /// your on-premise gateway or another Cloud provider’s VPN gateway,
        /// it cannot be an IP address from Google Compute Engine.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        public ExternalVpnGatewayInterfacesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ExternalVpnGatewayInterfaces
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The numberic ID for this interface. Allowed values are based on the redundancy type
        /// of this external VPN gateway
        /// * `0 - SINGLE_IP_INTERNALLY_REDUNDANT`
        /// * `0, 1 - TWO_IPS_REDUNDANCY`
        /// * `0, 1, 2, 3 - FOUR_IPS_REDUNDANCY`
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// -
        /// (Optional)
        /// IP address of the interface in the external VPN gateway.
        /// Only IPv4 is supported. This IP address can be either from
        /// your on-premise gateway or another Cloud provider’s VPN gateway,
        /// it cannot be an IP address from Google Compute Engine.
        /// </summary>
        public readonly string? IpAddress;

        [OutputConstructor]
        private ExternalVpnGatewayInterfaces(
            int? id,
            string? ipAddress)
        {
            Id = id;
            IpAddress = ipAddress;
        }
    }
    }
}
