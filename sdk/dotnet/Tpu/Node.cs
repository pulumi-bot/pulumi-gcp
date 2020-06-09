// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Tpu
{
    /// <summary>
    /// A Cloud TPU instance.
    /// 
    /// 
    /// To get more information about Node, see:
    /// 
    /// * [API documentation](https://cloud.google.com/tpu/docs/reference/rest/)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/tpu/docs/)
    /// 
    /// ## Example Usage
    /// 
    /// ### TPU Node Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var available = Output.Create(Gcp.Tpu.GetTensorflowVersions.InvokeAsync());
    ///         var tpu = new Gcp.Tpu.Node("tpu", new Gcp.Tpu.NodeArgs
    ///         {
    ///             Zone = "us-central1-b",
    ///             AcceleratorType = "v3-8",
    ///             TensorflowVersion = available.Apply(available =&gt; available.Versions[0]),
    ///             CidrBlock = "10.2.0.0/29",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Node : Pulumi.CustomResource
    {
        /// <summary>
        /// The type of hardware accelerators associated with this node.
        /// </summary>
        [Output("acceleratorType")]
        public Output<string> AcceleratorType { get; private set; } = null!;

        /// <summary>
        /// The CIDR block that the TPU node will use when selecting an IP
        /// address. This CIDR block must be a /29 block; the Compute Engine
        /// networks API forbids a smaller block, and using a larger block would
        /// be wasteful (a node can only consume one IP address).
        /// Errors will occur if the CIDR block has already been used for a
        /// currently existing TPU node, the CIDR block conflicts with any
        /// subnetworks in the user's provided network, or the provided network
        /// is peered with another network that is using that CIDR block.
        /// </summary>
        [Output("cidrBlock")]
        public Output<string> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// The user-supplied description of the TPU. Maximum of 512 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Resource labels to represent user provided metadata.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The immutable name of the TPU.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of a network to peer the TPU node to. It must be a
        /// preexisting Compute Engine network inside of the project on which
        /// this API has been activated. If none is provided, "default" will be
        /// used.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the
        /// node first reach out to the first (index 0) entry.
        /// </summary>
        [Output("networkEndpoints")]
        public Output<ImmutableArray<Outputs.NodeNetworkEndpoint>> NetworkEndpoints { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Sets the scheduling options for this TPU instance.  Structure is documented below.
        /// </summary>
        [Output("schedulingConfig")]
        public Output<Outputs.NodeSchedulingConfig?> SchedulingConfig { get; private set; } = null!;

        /// <summary>
        /// The service account used to run the tensor flow services within the node. To share resources, including Google Cloud
        /// Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
        /// </summary>
        [Output("serviceAccount")]
        public Output<string> ServiceAccount { get; private set; } = null!;

        /// <summary>
        /// The version of Tensorflow running in the Node.
        /// </summary>
        [Output("tensorflowVersion")]
        public Output<string> TensorflowVersion { get; private set; } = null!;

        /// <summary>
        /// The GCP location for the TPU.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Node resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Node(string name, NodeArgs args, CustomResourceOptions? options = null)
            : base("gcp:tpu/node:Node", name, args ?? new NodeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Node(string name, Input<string> id, NodeState? state = null, CustomResourceOptions? options = null)
            : base("gcp:tpu/node:Node", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Node resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Node Get(string name, Input<string> id, NodeState? state = null, CustomResourceOptions? options = null)
        {
            return new Node(name, id, state, options);
        }
    }

    public sealed class NodeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of hardware accelerators associated with this node.
        /// </summary>
        [Input("acceleratorType", required: true)]
        public Input<string> AcceleratorType { get; set; } = null!;

        /// <summary>
        /// The CIDR block that the TPU node will use when selecting an IP
        /// address. This CIDR block must be a /29 block; the Compute Engine
        /// networks API forbids a smaller block, and using a larger block would
        /// be wasteful (a node can only consume one IP address).
        /// Errors will occur if the CIDR block has already been used for a
        /// currently existing TPU node, the CIDR block conflicts with any
        /// subnetworks in the user's provided network, or the provided network
        /// is peered with another network that is using that CIDR block.
        /// </summary>
        [Input("cidrBlock", required: true)]
        public Input<string> CidrBlock { get; set; } = null!;

        /// <summary>
        /// The user-supplied description of the TPU. Maximum of 512 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels to represent user provided metadata.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The immutable name of the TPU.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of a network to peer the TPU node to. It must be a
        /// preexisting Compute Engine network inside of the project on which
        /// this API has been activated. If none is provided, "default" will be
        /// used.
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
        /// Sets the scheduling options for this TPU instance.  Structure is documented below.
        /// </summary>
        [Input("schedulingConfig")]
        public Input<Inputs.NodeSchedulingConfigArgs>? SchedulingConfig { get; set; }

        /// <summary>
        /// The version of Tensorflow running in the Node.
        /// </summary>
        [Input("tensorflowVersion", required: true)]
        public Input<string> TensorflowVersion { get; set; } = null!;

        /// <summary>
        /// The GCP location for the TPU.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public NodeArgs()
        {
        }
    }

    public sealed class NodeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of hardware accelerators associated with this node.
        /// </summary>
        [Input("acceleratorType")]
        public Input<string>? AcceleratorType { get; set; }

        /// <summary>
        /// The CIDR block that the TPU node will use when selecting an IP
        /// address. This CIDR block must be a /29 block; the Compute Engine
        /// networks API forbids a smaller block, and using a larger block would
        /// be wasteful (a node can only consume one IP address).
        /// Errors will occur if the CIDR block has already been used for a
        /// currently existing TPU node, the CIDR block conflicts with any
        /// subnetworks in the user's provided network, or the provided network
        /// is peered with another network that is using that CIDR block.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// The user-supplied description of the TPU. Maximum of 512 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels to represent user provided metadata.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The immutable name of the TPU.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of a network to peer the TPU node to. It must be a
        /// preexisting Compute Engine network inside of the project on which
        /// this API has been activated. If none is provided, "default" will be
        /// used.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        [Input("networkEndpoints")]
        private InputList<Inputs.NodeNetworkEndpointGetArgs>? _networkEndpoints;

        /// <summary>
        /// The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the
        /// node first reach out to the first (index 0) entry.
        /// </summary>
        public InputList<Inputs.NodeNetworkEndpointGetArgs> NetworkEndpoints
        {
            get => _networkEndpoints ?? (_networkEndpoints = new InputList<Inputs.NodeNetworkEndpointGetArgs>());
            set => _networkEndpoints = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Sets the scheduling options for this TPU instance.  Structure is documented below.
        /// </summary>
        [Input("schedulingConfig")]
        public Input<Inputs.NodeSchedulingConfigGetArgs>? SchedulingConfig { get; set; }

        /// <summary>
        /// The service account used to run the tensor flow services within the node. To share resources, including Google Cloud
        /// Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
        /// </summary>
        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        /// <summary>
        /// The version of Tensorflow running in the Node.
        /// </summary>
        [Input("tensorflowVersion")]
        public Input<string>? TensorflowVersion { get; set; }

        /// <summary>
        /// The GCP location for the TPU.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public NodeState()
        {
        }
    }
}
