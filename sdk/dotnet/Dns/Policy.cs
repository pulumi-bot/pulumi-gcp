// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dns
{
    /// <summary>
    /// A policy is a collection of DNS rules applied to one or more Virtual
    /// Private Cloud resources.
    /// 
    /// To get more information about Policy, see:
    /// 
    /// * [API documentation](https://cloud.google.com/dns/docs/reference/v1beta2/policies)
    /// * How-to Guides
    ///     * [Using DNS server policies](https://cloud.google.com/dns/zones/#using-dns-server-policies)
    /// 
    /// ## Example Usage
    /// ### Dns Policy Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var network_1 = new Gcp.Compute.Network("network-1", new Gcp.Compute.NetworkArgs
    ///         {
    ///             AutoCreateSubnetworks = false,
    ///         });
    ///         var network_2 = new Gcp.Compute.Network("network-2", new Gcp.Compute.NetworkArgs
    ///         {
    ///             AutoCreateSubnetworks = false,
    ///         });
    ///         var example_policy = new Gcp.Dns.Policy("example-policy", new Gcp.Dns.PolicyArgs
    ///         {
    ///             EnableInboundForwarding = true,
    ///             EnableLogging = true,
    ///             AlternativeNameServerConfig = new Gcp.Dns.Inputs.PolicyAlternativeNameServerConfigArgs
    ///             {
    ///                 TargetNameServers = 
    ///                 {
    ///                     new Gcp.Dns.Inputs.PolicyAlternativeNameServerConfigTargetNameServerArgs
    ///                     {
    ///                         Ipv4Address = "172.16.1.10",
    ///                         ForwardingPath = "private",
    ///                     },
    ///                     new Gcp.Dns.Inputs.PolicyAlternativeNameServerConfigTargetNameServerArgs
    ///                     {
    ///                         Ipv4Address = "172.16.1.20",
    ///                     },
    ///                 },
    ///             },
    ///             Networks = 
    ///             {
    ///                 new Gcp.Dns.Inputs.PolicyNetworkArgs
    ///                 {
    ///                     NetworkUrl = network_1.Id,
    ///                 },
    ///                 new Gcp.Dns.Inputs.PolicyNetworkArgs
    ///                 {
    ///                     NetworkUrl = network_2.Id,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Policy can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dns/policy:Policy default projects/{{project}}/policies/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dns/policy:Policy default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dns/policy:Policy default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:dns/policy:Policy")]
    public partial class Policy : Pulumi.CustomResource
    {
        /// <summary>
        /// Sets an alternative name server for the associated networks.
        /// When specified, all DNS queries are forwarded to a name server that you choose.
        /// Names such as .internal are not available when an alternative name server is specified.
        /// Structure is documented below.
        /// </summary>
        [Output("alternativeNameServerConfig")]
        public Output<Outputs.PolicyAlternativeNameServerConfig?> AlternativeNameServerConfig { get; private set; } = null!;

        /// <summary>
        /// A textual description field. Defaults to 'Managed by Pulumi'.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Allows networks bound to this policy to receive DNS queries sent
        /// by VMs or applications over VPN connections. When enabled, a
        /// virtual IP address will be allocated from each of the sub-networks
        /// that are bound to this policy.
        /// </summary>
        [Output("enableInboundForwarding")]
        public Output<bool?> EnableInboundForwarding { get; private set; } = null!;

        /// <summary>
        /// Controls whether logging is enabled for the networks bound to this policy.
        /// Defaults to no logging if not set.
        /// </summary>
        [Output("enableLogging")]
        public Output<bool?> EnableLogging { get; private set; } = null!;

        /// <summary>
        /// User assigned name for this policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of network names specifying networks to which this policy is applied.
        /// Structure is documented below.
        /// </summary>
        [Output("networks")]
        public Output<ImmutableArray<Outputs.PolicyNetwork>> Networks { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a Policy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Policy(string name, PolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:dns/policy:Policy", name, args ?? new PolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Policy(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:dns/policy:Policy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Policy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Policy Get(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Policy(name, id, state, options);
        }
    }

    public sealed class PolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sets an alternative name server for the associated networks.
        /// When specified, all DNS queries are forwarded to a name server that you choose.
        /// Names such as .internal are not available when an alternative name server is specified.
        /// Structure is documented below.
        /// </summary>
        [Input("alternativeNameServerConfig")]
        public Input<Inputs.PolicyAlternativeNameServerConfigArgs>? AlternativeNameServerConfig { get; set; }

        /// <summary>
        /// A textual description field. Defaults to 'Managed by Pulumi'.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Allows networks bound to this policy to receive DNS queries sent
        /// by VMs or applications over VPN connections. When enabled, a
        /// virtual IP address will be allocated from each of the sub-networks
        /// that are bound to this policy.
        /// </summary>
        [Input("enableInboundForwarding")]
        public Input<bool>? EnableInboundForwarding { get; set; }

        /// <summary>
        /// Controls whether logging is enabled for the networks bound to this policy.
        /// Defaults to no logging if not set.
        /// </summary>
        [Input("enableLogging")]
        public Input<bool>? EnableLogging { get; set; }

        /// <summary>
        /// User assigned name for this policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networks")]
        private InputList<Inputs.PolicyNetworkArgs>? _networks;

        /// <summary>
        /// List of network names specifying networks to which this policy is applied.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.PolicyNetworkArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.PolicyNetworkArgs>());
            set => _networks = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public PolicyArgs()
        {
        }
    }

    public sealed class PolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sets an alternative name server for the associated networks.
        /// When specified, all DNS queries are forwarded to a name server that you choose.
        /// Names such as .internal are not available when an alternative name server is specified.
        /// Structure is documented below.
        /// </summary>
        [Input("alternativeNameServerConfig")]
        public Input<Inputs.PolicyAlternativeNameServerConfigGetArgs>? AlternativeNameServerConfig { get; set; }

        /// <summary>
        /// A textual description field. Defaults to 'Managed by Pulumi'.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Allows networks bound to this policy to receive DNS queries sent
        /// by VMs or applications over VPN connections. When enabled, a
        /// virtual IP address will be allocated from each of the sub-networks
        /// that are bound to this policy.
        /// </summary>
        [Input("enableInboundForwarding")]
        public Input<bool>? EnableInboundForwarding { get; set; }

        /// <summary>
        /// Controls whether logging is enabled for the networks bound to this policy.
        /// Defaults to no logging if not set.
        /// </summary>
        [Input("enableLogging")]
        public Input<bool>? EnableLogging { get; set; }

        /// <summary>
        /// User assigned name for this policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networks")]
        private InputList<Inputs.PolicyNetworkGetArgs>? _networks;

        /// <summary>
        /// List of network names specifying networks to which this policy is applied.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.PolicyNetworkGetArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.PolicyNetworkGetArgs>());
            set => _networks = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public PolicyState()
        {
        }
    }
}
