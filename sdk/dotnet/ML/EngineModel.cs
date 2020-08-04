// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ML
{
    /// <summary>
    /// Represents a machine learning solution.
    /// 
    /// A model can have multiple versions, each of which is a deployed, trained model
    /// ready to receive prediction requests. The model itself is just a container.
    /// 
    /// ## Example Usage
    /// ### Ml Model Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @default = new Gcp.ML.EngineModel("default", new Gcp.ML.EngineModelArgs
    ///         {
    ///             Description = "My model",
    ///             Regions = "us-central1",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Ml Model Full
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @default = new Gcp.ML.EngineModel("default", new Gcp.ML.EngineModelArgs
    ///         {
    ///             Description = "My model",
    ///             Labels = 
    ///             {
    ///                 { "my_model", "foo" },
    ///             },
    ///             OnlinePredictionConsoleLogging = true,
    ///             OnlinePredictionLogging = true,
    ///             Regions = "us-central1",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class EngineModel : Pulumi.CustomResource
    {
        /// <summary>
        /// The default version of the model. This version will be used to handle
        /// prediction requests that do not specify a version.  Structure is documented below.
        /// </summary>
        [Output("defaultVersion")]
        public Output<Outputs.EngineModelDefaultVersion?> DefaultVersion { get; private set; } = null!;

        /// <summary>
        /// The description specified for the model when it was created.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// One or more labels that you can add, to organize your models.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The name specified for the version when it was created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging
        /// </summary>
        [Output("onlinePredictionConsoleLogging")]
        public Output<bool?> OnlinePredictionConsoleLogging { get; private set; } = null!;

        /// <summary>
        /// If true, online prediction access logs are sent to StackDriver Logging.
        /// </summary>
        [Output("onlinePredictionLogging")]
        public Output<bool?> OnlinePredictionLogging { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The list of regions where the model is going to be deployed.
        /// Currently only one region per model is supported
        /// </summary>
        [Output("regions")]
        public Output<string?> Regions { get; private set; } = null!;


        /// <summary>
        /// Create a EngineModel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EngineModel(string name, EngineModelArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:ml/engineModel:EngineModel", name, args ?? new EngineModelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EngineModel(string name, Input<string> id, EngineModelState? state = null, CustomResourceOptions? options = null)
            : base("gcp:ml/engineModel:EngineModel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EngineModel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EngineModel Get(string name, Input<string> id, EngineModelState? state = null, CustomResourceOptions? options = null)
        {
            return new EngineModel(name, id, state, options);
        }
    }

    public sealed class EngineModelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default version of the model. This version will be used to handle
        /// prediction requests that do not specify a version.  Structure is documented below.
        /// </summary>
        [Input("defaultVersion")]
        public Input<Inputs.EngineModelDefaultVersionArgs>? DefaultVersion { get; set; }

        /// <summary>
        /// The description specified for the model when it was created.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// One or more labels that you can add, to organize your models.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The name specified for the version when it was created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging
        /// </summary>
        [Input("onlinePredictionConsoleLogging")]
        public Input<bool>? OnlinePredictionConsoleLogging { get; set; }

        /// <summary>
        /// If true, online prediction access logs are sent to StackDriver Logging.
        /// </summary>
        [Input("onlinePredictionLogging")]
        public Input<bool>? OnlinePredictionLogging { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The list of regions where the model is going to be deployed.
        /// Currently only one region per model is supported
        /// </summary>
        [Input("regions")]
        public Input<string>? Regions { get; set; }

        public EngineModelArgs()
        {
        }
    }

    public sealed class EngineModelState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default version of the model. This version will be used to handle
        /// prediction requests that do not specify a version.  Structure is documented below.
        /// </summary>
        [Input("defaultVersion")]
        public Input<Inputs.EngineModelDefaultVersionGetArgs>? DefaultVersion { get; set; }

        /// <summary>
        /// The description specified for the model when it was created.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// One or more labels that you can add, to organize your models.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The name specified for the version when it was created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging
        /// </summary>
        [Input("onlinePredictionConsoleLogging")]
        public Input<bool>? OnlinePredictionConsoleLogging { get; set; }

        /// <summary>
        /// If true, online prediction access logs are sent to StackDriver Logging.
        /// </summary>
        [Input("onlinePredictionLogging")]
        public Input<bool>? OnlinePredictionLogging { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The list of regions where the model is going to be deployed.
        /// Currently only one region per model is supported
        /// </summary>
        [Input("regions")]
        public Input<string>? Regions { get; set; }

        public EngineModelState()
        {
        }
    }
}
