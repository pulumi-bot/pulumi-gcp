// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine
{
    /// <summary>
    /// Rules to match an HTTP request and dispatch that request to a service.
    /// 
    /// To get more information about ApplicationUrlDispatchRules, see:
    /// 
    /// * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps#UrlDispatchRule)
    /// 
    /// ## Example Usage
    /// ### App Engine Application Url Dispatch Rules Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bucket = new Gcp.Storage.Bucket("bucket", new Gcp.Storage.BucketArgs
    ///         {
    ///         });
    ///         var @object = new Gcp.Storage.BucketObject("object", new Gcp.Storage.BucketObjectArgs
    ///         {
    ///             Bucket = bucket.Name,
    ///             Source = new FileAsset("./test-fixtures/appengine/hello-world.zip"),
    ///         });
    ///         var adminV3 = new Gcp.AppEngine.StandardAppVersion("adminV3", new Gcp.AppEngine.StandardAppVersionArgs
    ///         {
    ///             VersionId = "v3",
    ///             Service = "admin",
    ///             Runtime = "nodejs10",
    ///             Entrypoint = new Gcp.AppEngine.Inputs.StandardAppVersionEntrypointArgs
    ///             {
    ///                 Shell = "node ./app.js",
    ///             },
    ///             Deployment = new Gcp.AppEngine.Inputs.StandardAppVersionDeploymentArgs
    ///             {
    ///                 Zip = new Gcp.AppEngine.Inputs.StandardAppVersionDeploymentZipArgs
    ///                 {
    ///                     SourceUrl = Output.Tuple(bucket.Name, @object.Name).Apply(values =&gt;
    ///                     {
    ///                         var bucketName = values.Item1;
    ///                         var objectName = values.Item2;
    ///                         return $"https://storage.googleapis.com/{bucketName}/{objectName}";
    ///                     }),
    ///                 },
    ///             },
    ///             EnvVariables = 
    ///             {
    ///                 { "port", "8080" },
    ///             },
    ///             NoopOnDestroy = true,
    ///         });
    ///         var webService = new Gcp.AppEngine.ApplicationUrlDispatchRules("webService", new Gcp.AppEngine.ApplicationUrlDispatchRulesArgs
    ///         {
    ///             DispatchRules = 
    ///             {
    ///                 new Gcp.AppEngine.Inputs.ApplicationUrlDispatchRulesDispatchRuleArgs
    ///                 {
    ///                     Domain = "*",
    ///                     Path = "/*",
    ///                     Service = "default",
    ///                 },
    ///                 new Gcp.AppEngine.Inputs.ApplicationUrlDispatchRulesDispatchRuleArgs
    ///                 {
    ///                     Domain = "*",
    ///                     Path = "/admin/*",
    ///                     Service = adminV3.Service,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ApplicationUrlDispatchRules : Pulumi.CustomResource
    {
        /// <summary>
        /// Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.
        /// </summary>
        [Output("dispatchRules")]
        public Output<ImmutableArray<Outputs.ApplicationUrlDispatchRulesDispatchRule>> DispatchRules { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationUrlDispatchRules resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationUrlDispatchRules(string name, ApplicationUrlDispatchRulesArgs args, CustomResourceOptions? options = null)
            : base("gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules", name, args ?? new ApplicationUrlDispatchRulesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationUrlDispatchRules(string name, Input<string> id, ApplicationUrlDispatchRulesState? state = null, CustomResourceOptions? options = null)
            : base("gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationUrlDispatchRules resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationUrlDispatchRules Get(string name, Input<string> id, ApplicationUrlDispatchRulesState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationUrlDispatchRules(name, id, state, options);
        }
    }

    public sealed class ApplicationUrlDispatchRulesArgs : Pulumi.ResourceArgs
    {
        [Input("dispatchRules", required: true)]
        private InputList<Inputs.ApplicationUrlDispatchRulesDispatchRuleArgs>? _dispatchRules;

        /// <summary>
        /// Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.ApplicationUrlDispatchRulesDispatchRuleArgs> DispatchRules
        {
            get => _dispatchRules ?? (_dispatchRules = new InputList<Inputs.ApplicationUrlDispatchRulesDispatchRuleArgs>());
            set => _dispatchRules = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ApplicationUrlDispatchRulesArgs()
        {
        }
    }

    public sealed class ApplicationUrlDispatchRulesState : Pulumi.ResourceArgs
    {
        [Input("dispatchRules")]
        private InputList<Inputs.ApplicationUrlDispatchRulesDispatchRuleGetArgs>? _dispatchRules;

        /// <summary>
        /// Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.ApplicationUrlDispatchRulesDispatchRuleGetArgs> DispatchRules
        {
            get => _dispatchRules ?? (_dispatchRules = new InputList<Inputs.ApplicationUrlDispatchRulesDispatchRuleGetArgs>());
            set => _dispatchRules = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ApplicationUrlDispatchRulesState()
        {
        }
    }
}
