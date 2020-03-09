// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Datastore
{
    public partial class DataStoreIndex : Pulumi.CustomResource
    {
        /// <summary>
        /// Policy for including ancestors in the index. Either 'ALL_ANCESTORS' or 'NONE', the default is 'NONE'.
        /// </summary>
        [Output("ancestor")]
        public Output<string?> Ancestor { get; private set; } = null!;

        /// <summary>
        /// The index id.
        /// </summary>
        [Output("indexId")]
        public Output<string> IndexId { get; private set; } = null!;

        /// <summary>
        /// The entity kind which the index applies to.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// An ordered list of properties to index on.
        /// </summary>
        [Output("properties")]
        public Output<ImmutableArray<Outputs.DataStoreIndexProperties>> Properties { get; private set; } = null!;


        /// <summary>
        /// Create a DataStoreIndex resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataStoreIndex(string name, DataStoreIndexArgs args, CustomResourceOptions? options = null)
            : base("gcp:datastore/dataStoreIndex:DataStoreIndex", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private DataStoreIndex(string name, Input<string> id, DataStoreIndexState? state = null, CustomResourceOptions? options = null)
            : base("gcp:datastore/dataStoreIndex:DataStoreIndex", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DataStoreIndex resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataStoreIndex Get(string name, Input<string> id, DataStoreIndexState? state = null, CustomResourceOptions? options = null)
        {
            return new DataStoreIndex(name, id, state, options);
        }
    }

    public sealed class DataStoreIndexArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy for including ancestors in the index. Either 'ALL_ANCESTORS' or 'NONE', the default is 'NONE'.
        /// </summary>
        [Input("ancestor")]
        public Input<string>? Ancestor { get; set; }

        /// <summary>
        /// The entity kind which the index applies to.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("properties")]
        private InputList<Inputs.DataStoreIndexPropertiesArgs>? _properties;

        /// <summary>
        /// An ordered list of properties to index on.
        /// </summary>
        public InputList<Inputs.DataStoreIndexPropertiesArgs> Properties
        {
            get => _properties ?? (_properties = new InputList<Inputs.DataStoreIndexPropertiesArgs>());
            set => _properties = value;
        }

        public DataStoreIndexArgs()
        {
        }
    }

    public sealed class DataStoreIndexState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy for including ancestors in the index. Either 'ALL_ANCESTORS' or 'NONE', the default is 'NONE'.
        /// </summary>
        [Input("ancestor")]
        public Input<string>? Ancestor { get; set; }

        /// <summary>
        /// The index id.
        /// </summary>
        [Input("indexId")]
        public Input<string>? IndexId { get; set; }

        /// <summary>
        /// The entity kind which the index applies to.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("properties")]
        private InputList<Inputs.DataStoreIndexPropertiesGetArgs>? _properties;

        /// <summary>
        /// An ordered list of properties to index on.
        /// </summary>
        public InputList<Inputs.DataStoreIndexPropertiesGetArgs> Properties
        {
            get => _properties ?? (_properties = new InputList<Inputs.DataStoreIndexPropertiesGetArgs>());
            set => _properties = value;
        }

        public DataStoreIndexState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class DataStoreIndexPropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("direction", required: true)]
        public Input<string> Direction { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public DataStoreIndexPropertiesArgs()
        {
        }
    }

    public sealed class DataStoreIndexPropertiesGetArgs : Pulumi.ResourceArgs
    {
        [Input("direction", required: true)]
        public Input<string> Direction { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public DataStoreIndexPropertiesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class DataStoreIndexProperties
    {
        public readonly string Direction;
        public readonly string Name;

        [OutputConstructor]
        private DataStoreIndexProperties(
            string direction,
            string name)
        {
            Direction = direction;
            Name = name;
        }
    }
    }
}
