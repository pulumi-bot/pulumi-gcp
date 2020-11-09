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
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/externalVpnGateways)
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// ExternalVpnGateway can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default projects/{{project}}/global/externalVpnGateways/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default {{name}}
    /// ```
    /// </summary>
    public partial class ExternalVpnGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A list of interfaces on this external VPN gateway.
        /// Structure is documented below.
        /// </summary>
        [Output("interfaces")]
        public Output<ImmutableArray<Outputs.ExternalVpnGatewayInterface>> Interfaces { get; private set; } = null!;

        /// <summary>
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
        /// Indicates the redundancy type of this external VPN gateway
        /// Possible values are `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, and `TWO_IPS_REDUNDANCY`.
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
            : base("gcp:compute/externalVpnGateway:ExternalVpnGateway", name, args ?? new ExternalVpnGatewayArgs(), MakeResourceOptions(options, ""))
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
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("interfaces")]
        private InputList<Inputs.ExternalVpnGatewayInterfaceArgs>? _interfaces;

        /// <summary>
        /// A list of interfaces on this external VPN gateway.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ExternalVpnGatewayInterfaceArgs> Interfaces
        {
            get => _interfaces ?? (_interfaces = new InputList<Inputs.ExternalVpnGatewayInterfaceArgs>());
            set => _interfaces = value;
        }

        /// <summary>
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
        /// Indicates the redundancy type of this external VPN gateway
        /// Possible values are `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, and `TWO_IPS_REDUNDANCY`.
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
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("interfaces")]
        private InputList<Inputs.ExternalVpnGatewayInterfaceGetArgs>? _interfaces;

        /// <summary>
        /// A list of interfaces on this external VPN gateway.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ExternalVpnGatewayInterfaceGetArgs> Interfaces
        {
            get => _interfaces ?? (_interfaces = new InputList<Inputs.ExternalVpnGatewayInterfaceGetArgs>());
            set => _interfaces = value;
        }

        /// <summary>
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
        /// Indicates the redundancy type of this external VPN gateway
        /// Possible values are `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, and `TWO_IPS_REDUNDANCY`.
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
}
