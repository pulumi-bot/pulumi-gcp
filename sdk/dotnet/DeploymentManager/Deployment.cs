// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DeploymentManager
{
    /// <summary>
    /// A collection of resources that are deployed and managed together using
    /// a configuration file
    /// 
    /// 
    /// 
    /// &gt; **Warning:** This resource is intended only to manage a Deployment resource,
    /// and attempts to manage the Deployment's resources in the provider as well
    /// will likely result in errors or unexpected behavior as the two tools
    /// fight over ownership. We strongly discourage doing so unless you are an
    /// experienced user of both tools.
    /// 
    /// In addition, due to limitations of the API, the provider will treat
    /// deployments in preview as recreate-only for any update operation other
    /// than actually deploying an in-preview deployment (i.e. `preview=true` to
    /// `preview=false`).
    /// 
    /// ## Example Usage - Deployment Manager Deployment Basic
    /// {{% example %}}
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var deployment = new Gcp.DeploymentManager.Deployment("deployment", new Gcp.DeploymentManager.DeploymentArgs
    ///         {
    ///             Target = new Gcp.DeploymentManager.Inputs.DeploymentTargetArgs
    ///             {
    ///                 Config = new Gcp.DeploymentManager.Inputs.DeploymentTargetConfigArgs
    ///                 {
    ///                     Content = "TODO: ReadFile",
    ///                 },
    ///             },
    ///             Labels = 
    ///             {
    ///                 new Gcp.DeploymentManager.Inputs.DeploymentLabelArgs
    ///                 {
    ///                     Key = "foo",
    ///                     Value = "bar",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// {{% /example %}}
    /// </summary>
    public partial class Deployment : Pulumi.CustomResource
    {
        /// <summary>
        /// Set the policy to use for creating new resources. Only used on
        /// create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
        /// `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
        /// the deployment will fail. Note that updating this field does not
        /// actually affect the deployment, just how it is updated.
        /// </summary>
        [Output("createPolicy")]
        public Output<string?> CreatePolicy { get; private set; } = null!;

        /// <summary>
        /// Set the policy to use for deleting new resources on update/delete.
        /// Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
        /// resource is deleted after removal from Deployment Manager. If
        /// `ABANDON`, the resource is only removed from Deployment Manager
        /// and is not actually deleted. Note that updating this field does not
        /// actually change the deployment, just how it is updated.
        /// </summary>
        [Output("deletePolicy")]
        public Output<string?> DeletePolicy { get; private set; } = null!;

        /// <summary>
        /// Unique identifier for deployment. Output only.
        /// </summary>
        [Output("deploymentId")]
        public Output<string> DeploymentId { get; private set; } = null!;

        /// <summary>
        /// Optional user-provided description of deployment.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Key-value pairs to apply to this labels.  Structure is documented below.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<Outputs.DeploymentLabel>> Labels { get; private set; } = null!;

        /// <summary>
        /// Output only. URL of the manifest representing the last manifest that was successfully deployed.
        /// </summary>
        [Output("manifest")]
        public Output<string> Manifest { get; private set; } = null!;

        /// <summary>
        /// The name of the template to import, as declared in the YAML
        /// configuration.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// If set to true, a deployment is created with "shell" resources
        /// that are not actually instantiated. This allows you to preview a
        /// deployment. It can be updated to false to actually deploy
        /// with real resources.
        /// ~&gt;**NOTE**: Deployment Manager does not allow update
        /// of a deployment in preview (unless updating to preview=false). Thus,
        /// the provider will force-recreate deployments if either preview is updated
        /// to true or if other fields are updated while preview is true.
        /// </summary>
        [Output("preview")]
        public Output<bool?> Preview { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Output only. Server defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Parameters that define your deployment, including the deployment
        /// configuration and relevant templates.  Structure is documented below.
        /// </summary>
        [Output("target")]
        public Output<Outputs.DeploymentTarget> Target { get; private set; } = null!;


        /// <summary>
        /// Create a Deployment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Deployment(string name, DeploymentArgs args, CustomResourceOptions? options = null)
            : base("gcp:deploymentmanager/deployment:Deployment", name, args ?? new DeploymentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Deployment(string name, Input<string> id, DeploymentState? state = null, CustomResourceOptions? options = null)
            : base("gcp:deploymentmanager/deployment:Deployment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Deployment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Deployment Get(string name, Input<string> id, DeploymentState? state = null, CustomResourceOptions? options = null)
        {
            return new Deployment(name, id, state, options);
        }
    }

    public sealed class DeploymentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set the policy to use for creating new resources. Only used on
        /// create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
        /// `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
        /// the deployment will fail. Note that updating this field does not
        /// actually affect the deployment, just how it is updated.
        /// </summary>
        [Input("createPolicy")]
        public Input<string>? CreatePolicy { get; set; }

        /// <summary>
        /// Set the policy to use for deleting new resources on update/delete.
        /// Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
        /// resource is deleted after removal from Deployment Manager. If
        /// `ABANDON`, the resource is only removed from Deployment Manager
        /// and is not actually deleted. Note that updating this field does not
        /// actually change the deployment, just how it is updated.
        /// </summary>
        [Input("deletePolicy")]
        public Input<string>? DeletePolicy { get; set; }

        /// <summary>
        /// Optional user-provided description of deployment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputList<Inputs.DeploymentLabelArgs>? _labels;

        /// <summary>
        /// Key-value pairs to apply to this labels.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.DeploymentLabelArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.DeploymentLabelArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// The name of the template to import, as declared in the YAML
        /// configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// If set to true, a deployment is created with "shell" resources
        /// that are not actually instantiated. This allows you to preview a
        /// deployment. It can be updated to false to actually deploy
        /// with real resources.
        /// ~&gt;**NOTE**: Deployment Manager does not allow update
        /// of a deployment in preview (unless updating to preview=false). Thus,
        /// the provider will force-recreate deployments if either preview is updated
        /// to true or if other fields are updated while preview is true.
        /// </summary>
        [Input("preview")]
        public Input<bool>? Preview { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Parameters that define your deployment, including the deployment
        /// configuration and relevant templates.  Structure is documented below.
        /// </summary>
        [Input("target", required: true)]
        public Input<Inputs.DeploymentTargetArgs> Target { get; set; } = null!;

        public DeploymentArgs()
        {
        }
    }

    public sealed class DeploymentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set the policy to use for creating new resources. Only used on
        /// create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
        /// `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
        /// the deployment will fail. Note that updating this field does not
        /// actually affect the deployment, just how it is updated.
        /// </summary>
        [Input("createPolicy")]
        public Input<string>? CreatePolicy { get; set; }

        /// <summary>
        /// Set the policy to use for deleting new resources on update/delete.
        /// Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
        /// resource is deleted after removal from Deployment Manager. If
        /// `ABANDON`, the resource is only removed from Deployment Manager
        /// and is not actually deleted. Note that updating this field does not
        /// actually change the deployment, just how it is updated.
        /// </summary>
        [Input("deletePolicy")]
        public Input<string>? DeletePolicy { get; set; }

        /// <summary>
        /// Unique identifier for deployment. Output only.
        /// </summary>
        [Input("deploymentId")]
        public Input<string>? DeploymentId { get; set; }

        /// <summary>
        /// Optional user-provided description of deployment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputList<Inputs.DeploymentLabelGetArgs>? _labels;

        /// <summary>
        /// Key-value pairs to apply to this labels.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.DeploymentLabelGetArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.DeploymentLabelGetArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// Output only. URL of the manifest representing the last manifest that was successfully deployed.
        /// </summary>
        [Input("manifest")]
        public Input<string>? Manifest { get; set; }

        /// <summary>
        /// The name of the template to import, as declared in the YAML
        /// configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// If set to true, a deployment is created with "shell" resources
        /// that are not actually instantiated. This allows you to preview a
        /// deployment. It can be updated to false to actually deploy
        /// with real resources.
        /// ~&gt;**NOTE**: Deployment Manager does not allow update
        /// of a deployment in preview (unless updating to preview=false). Thus,
        /// the provider will force-recreate deployments if either preview is updated
        /// to true or if other fields are updated while preview is true.
        /// </summary>
        [Input("preview")]
        public Input<bool>? Preview { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Output only. Server defined URL for the resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// Parameters that define your deployment, including the deployment
        /// configuration and relevant templates.  Structure is documented below.
        /// </summary>
        [Input("target")]
        public Input<Inputs.DeploymentTargetGetArgs>? Target { get; set; }

        public DeploymentState()
        {
        }
    }
}
