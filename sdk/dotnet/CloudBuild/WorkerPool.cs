// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudBuild
{
    /// <summary>
    /// ## Import
    /// 
    /// WorkerPool can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:cloudbuild/workerPool:WorkerPool default projects/{{project}}/locations/{{location}}/workerPools/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:cloudbuild/workerPool:WorkerPool default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:cloudbuild/workerPool:WorkerPool default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:cloudbuild/workerPool:WorkerPool")]
    public partial class WorkerPool : Pulumi.CustomResource
    {
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        [Output("deleteTime")]
        public Output<string> DeleteTime { get; private set; } = null!;

        /// <summary>
        /// The location for the resource
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// User-defined name of the `WorkerPool`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network configuration for the `WorkerPool`.
        /// </summary>
        [Output("networkConfig")]
        public Output<Outputs.WorkerPoolNetworkConfig?> NetworkConfig { get; private set; } = null!;

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Configuration to be used for a creating workers in the `WorkerPool`.
        /// </summary>
        [Output("workerConfig")]
        public Output<Outputs.WorkerPoolWorkerConfig> WorkerConfig { get; private set; } = null!;


        /// <summary>
        /// Create a WorkerPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkerPool(string name, WorkerPoolArgs args, CustomResourceOptions? options = null)
            : base("gcp:cloudbuild/workerPool:WorkerPool", name, args ?? new WorkerPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkerPool(string name, Input<string> id, WorkerPoolState? state = null, CustomResourceOptions? options = null)
            : base("gcp:cloudbuild/workerPool:WorkerPool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WorkerPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkerPool Get(string name, Input<string> id, WorkerPoolState? state = null, CustomResourceOptions? options = null)
        {
            return new WorkerPool(name, id, state, options);
        }
    }

    public sealed class WorkerPoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The location for the resource
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// User-defined name of the `WorkerPool`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network configuration for the `WorkerPool`.
        /// </summary>
        [Input("networkConfig")]
        public Input<Inputs.WorkerPoolNetworkConfigArgs>? NetworkConfig { get; set; }

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Configuration to be used for a creating workers in the `WorkerPool`.
        /// </summary>
        [Input("workerConfig")]
        public Input<Inputs.WorkerPoolWorkerConfigArgs>? WorkerConfig { get; set; }

        public WorkerPoolArgs()
        {
        }
    }

    public sealed class WorkerPoolState : Pulumi.ResourceArgs
    {
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        [Input("deleteTime")]
        public Input<string>? DeleteTime { get; set; }

        /// <summary>
        /// The location for the resource
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// User-defined name of the `WorkerPool`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network configuration for the `WorkerPool`.
        /// </summary>
        [Input("networkConfig")]
        public Input<Inputs.WorkerPoolNetworkConfigGetArgs>? NetworkConfig { get; set; }

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// Configuration to be used for a creating workers in the `WorkerPool`.
        /// </summary>
        [Input("workerConfig")]
        public Input<Inputs.WorkerPoolWorkerConfigGetArgs>? WorkerConfig { get; set; }

        public WorkerPoolState()
        {
        }
    }
}