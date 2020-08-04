// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.NetworkManagement
{
    /// <summary>
    /// A connectivity test are a static analysis of your resource configurations
    /// that enables you to evaluate connectivity to and from Google Cloud
    /// resources in your Virtual Private Cloud (VPC) network.
    /// 
    /// To get more information about ConnectivityTest, see:
    /// 
    /// * [API documentation](https://cloud.google.com/network-intelligence-center/docs/connectivity-tests/reference/networkmanagement/rest/v1/projects.locations.global.connectivityTests)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/network-intelligence-center/docs)
    /// 
    /// ## Example Usage
    /// ### Network Management Connectivity Test Instances
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var vpc = new Gcp.Compute.Network("vpc", new Gcp.Compute.NetworkArgs
    ///         {
    ///         });
    ///         var debian9 = Output.Create(Gcp.Compute.GetImage.InvokeAsync(new Gcp.Compute.GetImageArgs
    ///         {
    ///             Family = "debian-9",
    ///             Project = "debian-cloud",
    ///         }));
    ///         var source = new Gcp.Compute.Instance("source", new Gcp.Compute.InstanceArgs
    ///         {
    ///             MachineType = "n1-standard-1",
    ///             BootDisk = new Gcp.Compute.Inputs.InstanceBootDiskArgs
    ///             {
    ///                 InitializeParams = new Gcp.Compute.Inputs.InstanceBootDiskInitializeParamsArgs
    ///                 {
    ///                     Image = debian9.Apply(debian9 =&gt; debian9.Id),
    ///                 },
    ///             },
    ///             NetworkInterfaces = 
    ///             {
    ///                 new Gcp.Compute.Inputs.InstanceNetworkInterfaceArgs
    ///                 {
    ///                     Network = vpc.Id,
    ///                     AccessConfigs = 
    ///                     {
    ///                         ,
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///         var destination = new Gcp.Compute.Instance("destination", new Gcp.Compute.InstanceArgs
    ///         {
    ///             MachineType = "n1-standard-1",
    ///             BootDisk = new Gcp.Compute.Inputs.InstanceBootDiskArgs
    ///             {
    ///                 InitializeParams = new Gcp.Compute.Inputs.InstanceBootDiskInitializeParamsArgs
    ///                 {
    ///                     Image = debian9.Apply(debian9 =&gt; debian9.Id),
    ///                 },
    ///             },
    ///             NetworkInterfaces = 
    ///             {
    ///                 new Gcp.Compute.Inputs.InstanceNetworkInterfaceArgs
    ///                 {
    ///                     Network = vpc.Id,
    ///                     AccessConfigs = 
    ///                     {
    ///                         ,
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///         var instance_test = new Gcp.NetworkManagement.ConnectivityTest("instance-test", new Gcp.NetworkManagement.ConnectivityTestArgs
    ///         {
    ///             Source = new Gcp.NetworkManagement.Inputs.ConnectivityTestSourceArgs
    ///             {
    ///                 Instance = source.Id,
    ///             },
    ///             Destination = new Gcp.NetworkManagement.Inputs.ConnectivityTestDestinationArgs
    ///             {
    ///                 Instance = destination.Id,
    ///             },
    ///             Protocol = "TCP",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Network Management Connectivity Test Addresses
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var vpc = new Gcp.Compute.Network("vpc", new Gcp.Compute.NetworkArgs
    ///         {
    ///         });
    ///         var subnet = new Gcp.Compute.Subnetwork("subnet", new Gcp.Compute.SubnetworkArgs
    ///         {
    ///             IpCidrRange = "10.0.0.0/16",
    ///             Region = "us-central1",
    ///             Network = vpc.Id,
    ///         });
    ///         var source_addr = new Gcp.Compute.Address("source-addr", new Gcp.Compute.AddressArgs
    ///         {
    ///             Subnetwork = subnet.Id,
    ///             AddressType = "INTERNAL",
    ///             Address = "10.0.42.42",
    ///             Region = "us-central1",
    ///         });
    ///         var dest_addr = new Gcp.Compute.Address("dest-addr", new Gcp.Compute.AddressArgs
    ///         {
    ///             Subnetwork = subnet.Id,
    ///             AddressType = "INTERNAL",
    ///             Address = "10.0.43.43",
    ///             Region = "us-central1",
    ///         });
    ///         var address_test = new Gcp.NetworkManagement.ConnectivityTest("address-test", new Gcp.NetworkManagement.ConnectivityTestArgs
    ///         {
    ///             Source = new Gcp.NetworkManagement.Inputs.ConnectivityTestSourceArgs
    ///             {
    ///                 IpAddress = source_addr.IPAddress,
    ///                 ProjectId = source_addr.Project,
    ///                 Network = vpc.Id,
    ///                 NetworkType = "GCP_NETWORK",
    ///             },
    ///             Destination = new Gcp.NetworkManagement.Inputs.ConnectivityTestDestinationArgs
    ///             {
    ///                 IpAddress = dest_addr.IPAddress,
    ///                 ProjectId = dest_addr.Project,
    ///                 Network = vpc.Id,
    ///             },
    ///             Protocol = "UDP",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ConnectivityTest : Pulumi.CustomResource
    {
        /// <summary>
        /// The user-supplied description of the Connectivity Test.
        /// Maximum of 512 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Required. Destination specification of the Connectivity Test.
        /// You can use a combination of destination IP address, Compute
        /// Engine VM instance, or VPC network to uniquely identify the
        /// destination location.
        /// Even if the destination IP address is not unique, the source IP
        /// location is unique. Usually, the analysis can infer the destination
        /// endpoint from route information.
        /// If the destination you specify is a VM instance and the instance has
        /// multiple network interfaces, then you must also specify either a
        /// destination IP address or VPC network to identify the destination
        /// interface.
        /// A reachability analysis proceeds even if the destination location
        /// is ambiguous. However, the result can include endpoints that you
        /// don't intend to test.  Structure is documented below.
        /// </summary>
        [Output("destination")]
        public Output<Outputs.ConnectivityTestDestination> Destination { get; private set; } = null!;

        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Unique name for the connectivity test.
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
        /// IP Protocol of the test. When not provided, "TCP" is assumed.
        /// </summary>
        [Output("protocol")]
        public Output<string?> Protocol { get; private set; } = null!;

        /// <summary>
        /// Other projects that may be relevant for reachability analysis.
        /// This is applicable to scenarios where a test can cross project
        /// boundaries.
        /// </summary>
        [Output("relatedProjects")]
        public Output<ImmutableArray<string>> RelatedProjects { get; private set; } = null!;

        /// <summary>
        /// Required. Source specification of the Connectivity Test.
        /// You can use a combination of source IP address, virtual machine
        /// (VM) instance, or Compute Engine network to uniquely identify the
        /// source location.
        /// Examples: If the source IP address is an internal IP address within
        /// a Google Cloud Virtual Private Cloud (VPC) network, then you must
        /// also specify the VPC network. Otherwise, specify the VM instance,
        /// which already contains its internal IP address and VPC network
        /// information.
        /// If the source of the test is within an on-premises network, then
        /// you must provide the destination VPC network.
        /// If the source endpoint is a Compute Engine VM instance with multiple
        /// network interfaces, the instance itself is not sufficient to
        /// identify the endpoint. So, you must also specify the source IP
        /// address or VPC network.
        /// A reachability analysis proceeds even if the source location is
        /// ambiguous. However, the test result may include endpoints that
        /// you don't intend to test.  Structure is documented below.
        /// </summary>
        [Output("source")]
        public Output<Outputs.ConnectivityTestSource> Source { get; private set; } = null!;


        /// <summary>
        /// Create a ConnectivityTest resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConnectivityTest(string name, ConnectivityTestArgs args, CustomResourceOptions? options = null)
            : base("gcp:networkmanagement/connectivityTest:ConnectivityTest", name, args ?? new ConnectivityTestArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConnectivityTest(string name, Input<string> id, ConnectivityTestState? state = null, CustomResourceOptions? options = null)
            : base("gcp:networkmanagement/connectivityTest:ConnectivityTest", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConnectivityTest resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConnectivityTest Get(string name, Input<string> id, ConnectivityTestState? state = null, CustomResourceOptions? options = null)
        {
            return new ConnectivityTest(name, id, state, options);
        }
    }

    public sealed class ConnectivityTestArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The user-supplied description of the Connectivity Test.
        /// Maximum of 512 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Required. Destination specification of the Connectivity Test.
        /// You can use a combination of destination IP address, Compute
        /// Engine VM instance, or VPC network to uniquely identify the
        /// destination location.
        /// Even if the destination IP address is not unique, the source IP
        /// location is unique. Usually, the analysis can infer the destination
        /// endpoint from route information.
        /// If the destination you specify is a VM instance and the instance has
        /// multiple network interfaces, then you must also specify either a
        /// destination IP address or VPC network to identify the destination
        /// interface.
        /// A reachability analysis proceeds even if the destination location
        /// is ambiguous. However, the result can include endpoints that you
        /// don't intend to test.  Structure is documented below.
        /// </summary>
        [Input("destination", required: true)]
        public Input<Inputs.ConnectivityTestDestinationArgs> Destination { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Unique name for the connectivity test.
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
        /// IP Protocol of the test. When not provided, "TCP" is assumed.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("relatedProjects")]
        private InputList<string>? _relatedProjects;

        /// <summary>
        /// Other projects that may be relevant for reachability analysis.
        /// This is applicable to scenarios where a test can cross project
        /// boundaries.
        /// </summary>
        public InputList<string> RelatedProjects
        {
            get => _relatedProjects ?? (_relatedProjects = new InputList<string>());
            set => _relatedProjects = value;
        }

        /// <summary>
        /// Required. Source specification of the Connectivity Test.
        /// You can use a combination of source IP address, virtual machine
        /// (VM) instance, or Compute Engine network to uniquely identify the
        /// source location.
        /// Examples: If the source IP address is an internal IP address within
        /// a Google Cloud Virtual Private Cloud (VPC) network, then you must
        /// also specify the VPC network. Otherwise, specify the VM instance,
        /// which already contains its internal IP address and VPC network
        /// information.
        /// If the source of the test is within an on-premises network, then
        /// you must provide the destination VPC network.
        /// If the source endpoint is a Compute Engine VM instance with multiple
        /// network interfaces, the instance itself is not sufficient to
        /// identify the endpoint. So, you must also specify the source IP
        /// address or VPC network.
        /// A reachability analysis proceeds even if the source location is
        /// ambiguous. However, the test result may include endpoints that
        /// you don't intend to test.  Structure is documented below.
        /// </summary>
        [Input("source", required: true)]
        public Input<Inputs.ConnectivityTestSourceArgs> Source { get; set; } = null!;

        public ConnectivityTestArgs()
        {
        }
    }

    public sealed class ConnectivityTestState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The user-supplied description of the Connectivity Test.
        /// Maximum of 512 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Required. Destination specification of the Connectivity Test.
        /// You can use a combination of destination IP address, Compute
        /// Engine VM instance, or VPC network to uniquely identify the
        /// destination location.
        /// Even if the destination IP address is not unique, the source IP
        /// location is unique. Usually, the analysis can infer the destination
        /// endpoint from route information.
        /// If the destination you specify is a VM instance and the instance has
        /// multiple network interfaces, then you must also specify either a
        /// destination IP address or VPC network to identify the destination
        /// interface.
        /// A reachability analysis proceeds even if the destination location
        /// is ambiguous. However, the result can include endpoints that you
        /// don't intend to test.  Structure is documented below.
        /// </summary>
        [Input("destination")]
        public Input<Inputs.ConnectivityTestDestinationGetArgs>? Destination { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Unique name for the connectivity test.
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
        /// IP Protocol of the test. When not provided, "TCP" is assumed.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("relatedProjects")]
        private InputList<string>? _relatedProjects;

        /// <summary>
        /// Other projects that may be relevant for reachability analysis.
        /// This is applicable to scenarios where a test can cross project
        /// boundaries.
        /// </summary>
        public InputList<string> RelatedProjects
        {
            get => _relatedProjects ?? (_relatedProjects = new InputList<string>());
            set => _relatedProjects = value;
        }

        /// <summary>
        /// Required. Source specification of the Connectivity Test.
        /// You can use a combination of source IP address, virtual machine
        /// (VM) instance, or Compute Engine network to uniquely identify the
        /// source location.
        /// Examples: If the source IP address is an internal IP address within
        /// a Google Cloud Virtual Private Cloud (VPC) network, then you must
        /// also specify the VPC network. Otherwise, specify the VM instance,
        /// which already contains its internal IP address and VPC network
        /// information.
        /// If the source of the test is within an on-premises network, then
        /// you must provide the destination VPC network.
        /// If the source endpoint is a Compute Engine VM instance with multiple
        /// network interfaces, the instance itself is not sufficient to
        /// identify the endpoint. So, you must also specify the source IP
        /// address or VPC network.
        /// A reachability analysis proceeds even if the source location is
        /// ambiguous. However, the test result may include endpoints that
        /// you don't intend to test.  Structure is documented below.
        /// </summary>
        [Input("source")]
        public Input<Inputs.ConnectivityTestSourceGetArgs>? Source { get; set; }

        public ConnectivityTestState()
        {
        }
    }
}
