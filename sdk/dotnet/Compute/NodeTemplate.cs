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
    /// Represents a NodeTemplate resource. Node templates specify properties
    /// for creating sole-tenant nodes, such as node type, vCPU and memory
    /// requirements, node affinity labels, and region.
    /// 
    /// To get more information about NodeTemplate, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/nodeTemplates)
    /// * How-to Guides
    ///     * [Sole-Tenant Nodes](https://cloud.google.com/compute/docs/nodes/)
    /// 
    /// {{% examples %}}
    /// ## Example Usage
    /// {{% example %}}
    /// ### Node Template Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var central1a = Output.Create(Gcp.Compute.GetNodeTypes.InvokeAsync(new Gcp.Compute.GetNodeTypesArgs
    ///         {
    ///             Zone = "us-central1-a",
    ///         }));
    ///         var template = new Gcp.Compute.NodeTemplate("template", new Gcp.Compute.NodeTemplateArgs
    ///         {
    ///             Region = "us-central1",
    ///             NodeType = central1a.Apply(central1a =&gt; central1a.Names[0]),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// {{% /example %}}
    /// {{% /examples %}}
    /// </summary>
    public partial class NodeTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional textual description of the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Labels to use for node affinity, which will be used in
        /// instance scheduling.
        /// </summary>
        [Output("nodeAffinityLabels")]
        public Output<ImmutableDictionary<string, string>?> NodeAffinityLabels { get; private set; } = null!;

        /// <summary>
        /// Node type to use for nodes group that are created from this template.
        /// Only one of nodeTypeFlexibility and nodeType can be specified.
        /// </summary>
        [Output("nodeType")]
        public Output<string?> NodeType { get; private set; } = null!;

        /// <summary>
        /// Flexible properties for the desired node type. Node groups that
        /// use this node template will create nodes of a type that matches
        /// these properties. Only one of nodeTypeFlexibility and nodeType can
        /// be specified.  Structure is documented below.
        /// </summary>
        [Output("nodeTypeFlexibility")]
        public Output<Outputs.NodeTemplateNodeTypeFlexibility?> NodeTypeFlexibility { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Region where nodes using the node template will be created.
        /// If it is not provided, the provider region is used.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// The server binding policy for nodes using this template. Determines
        /// where the nodes should restart following a maintenance event.  Structure is documented below.
        /// </summary>
        [Output("serverBinding")]
        public Output<Outputs.NodeTemplateServerBinding> ServerBinding { get; private set; } = null!;


        /// <summary>
        /// Create a NodeTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NodeTemplate(string name, NodeTemplateArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:compute/nodeTemplate:NodeTemplate", name, args ?? new NodeTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NodeTemplate(string name, Input<string> id, NodeTemplateState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/nodeTemplate:NodeTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NodeTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NodeTemplate Get(string name, Input<string> id, NodeTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new NodeTemplate(name, id, state, options);
        }
    }

    public sealed class NodeTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional textual description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nodeAffinityLabels")]
        private InputMap<string>? _nodeAffinityLabels;

        /// <summary>
        /// Labels to use for node affinity, which will be used in
        /// instance scheduling.
        /// </summary>
        public InputMap<string> NodeAffinityLabels
        {
            get => _nodeAffinityLabels ?? (_nodeAffinityLabels = new InputMap<string>());
            set => _nodeAffinityLabels = value;
        }

        /// <summary>
        /// Node type to use for nodes group that are created from this template.
        /// Only one of nodeTypeFlexibility and nodeType can be specified.
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        /// <summary>
        /// Flexible properties for the desired node type. Node groups that
        /// use this node template will create nodes of a type that matches
        /// these properties. Only one of nodeTypeFlexibility and nodeType can
        /// be specified.  Structure is documented below.
        /// </summary>
        [Input("nodeTypeFlexibility")]
        public Input<Inputs.NodeTemplateNodeTypeFlexibilityArgs>? NodeTypeFlexibility { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Region where nodes using the node template will be created.
        /// If it is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The server binding policy for nodes using this template. Determines
        /// where the nodes should restart following a maintenance event.  Structure is documented below.
        /// </summary>
        [Input("serverBinding")]
        public Input<Inputs.NodeTemplateServerBindingArgs>? ServerBinding { get; set; }

        public NodeTemplateArgs()
        {
        }
    }

    public sealed class NodeTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional textual description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nodeAffinityLabels")]
        private InputMap<string>? _nodeAffinityLabels;

        /// <summary>
        /// Labels to use for node affinity, which will be used in
        /// instance scheduling.
        /// </summary>
        public InputMap<string> NodeAffinityLabels
        {
            get => _nodeAffinityLabels ?? (_nodeAffinityLabels = new InputMap<string>());
            set => _nodeAffinityLabels = value;
        }

        /// <summary>
        /// Node type to use for nodes group that are created from this template.
        /// Only one of nodeTypeFlexibility and nodeType can be specified.
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        /// <summary>
        /// Flexible properties for the desired node type. Node groups that
        /// use this node template will create nodes of a type that matches
        /// these properties. Only one of nodeTypeFlexibility and nodeType can
        /// be specified.  Structure is documented below.
        /// </summary>
        [Input("nodeTypeFlexibility")]
        public Input<Inputs.NodeTemplateNodeTypeFlexibilityGetArgs>? NodeTypeFlexibility { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Region where nodes using the node template will be created.
        /// If it is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// The server binding policy for nodes using this template. Determines
        /// where the nodes should restart following a maintenance event.  Structure is documented below.
        /// </summary>
        [Input("serverBinding")]
        public Input<Inputs.NodeTemplateServerBindingGetArgs>? ServerBinding { get; set; }

        public NodeTemplateState()
        {
        }
    }
}
