// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudTasks
{
    /// <summary>
    /// A named resource to which messages are sent by publishers.
    /// 
    /// ## Example Usage
    /// </summary>
    public partial class Queue : Pulumi.CustomResource
    {
        /// <summary>
        /// Overrides for task-level appEngineRouting. These settings apply only
        /// to App Engine tasks in this queue
        /// Structure is documented below.
        /// </summary>
        [Output("appEngineRoutingOverride")]
        public Output<Outputs.QueueAppEngineRoutingOverride?> AppEngineRoutingOverride { get; private set; } = null!;

        /// <summary>
        /// The location of the queue
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The queue name.
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
        /// Rate limits for task dispatches.
        /// The queue's actual dispatch rate is the result of:
        /// * Number of tasks in the queue
        /// * User-specified throttling: rateLimits, retryConfig, and the queue's state.
        /// * System throttling due to 429 (Too Many Requests) or 503 (Service
        /// Unavailable) responses from the worker, high error rates, or to
        /// smooth sudden large traffic spikes.
        /// Structure is documented below.
        /// </summary>
        [Output("rateLimits")]
        public Output<Outputs.QueueRateLimits> RateLimits { get; private set; } = null!;

        /// <summary>
        /// Settings that determine the retry behavior.
        /// Structure is documented below.
        /// </summary>
        [Output("retryConfig")]
        public Output<Outputs.QueueRetryConfig> RetryConfig { get; private set; } = null!;

        /// <summary>
        /// Configuration options for writing logs to Stackdriver Logging.
        /// Structure is documented below.
        /// </summary>
        [Output("stackdriverLoggingConfig")]
        public Output<Outputs.QueueStackdriverLoggingConfig?> StackdriverLoggingConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Queue resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Queue(string name, QueueArgs args, CustomResourceOptions? options = null)
            : base("gcp:cloudtasks/queue:Queue", name, args ?? new QueueArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Queue(string name, Input<string> id, QueueState? state = null, CustomResourceOptions? options = null)
            : base("gcp:cloudtasks/queue:Queue", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Queue resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Queue Get(string name, Input<string> id, QueueState? state = null, CustomResourceOptions? options = null)
        {
            return new Queue(name, id, state, options);
        }
    }

    public sealed class QueueArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Overrides for task-level appEngineRouting. These settings apply only
        /// to App Engine tasks in this queue
        /// Structure is documented below.
        /// </summary>
        [Input("appEngineRoutingOverride")]
        public Input<Inputs.QueueAppEngineRoutingOverrideArgs>? AppEngineRoutingOverride { get; set; }

        /// <summary>
        /// The location of the queue
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The queue name.
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
        /// Rate limits for task dispatches.
        /// The queue's actual dispatch rate is the result of:
        /// * Number of tasks in the queue
        /// * User-specified throttling: rateLimits, retryConfig, and the queue's state.
        /// * System throttling due to 429 (Too Many Requests) or 503 (Service
        /// Unavailable) responses from the worker, high error rates, or to
        /// smooth sudden large traffic spikes.
        /// Structure is documented below.
        /// </summary>
        [Input("rateLimits")]
        public Input<Inputs.QueueRateLimitsArgs>? RateLimits { get; set; }

        /// <summary>
        /// Settings that determine the retry behavior.
        /// Structure is documented below.
        /// </summary>
        [Input("retryConfig")]
        public Input<Inputs.QueueRetryConfigArgs>? RetryConfig { get; set; }

        /// <summary>
        /// Configuration options for writing logs to Stackdriver Logging.
        /// Structure is documented below.
        /// </summary>
        [Input("stackdriverLoggingConfig")]
        public Input<Inputs.QueueStackdriverLoggingConfigArgs>? StackdriverLoggingConfig { get; set; }

        public QueueArgs()
        {
        }
    }

    public sealed class QueueState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Overrides for task-level appEngineRouting. These settings apply only
        /// to App Engine tasks in this queue
        /// Structure is documented below.
        /// </summary>
        [Input("appEngineRoutingOverride")]
        public Input<Inputs.QueueAppEngineRoutingOverrideGetArgs>? AppEngineRoutingOverride { get; set; }

        /// <summary>
        /// The location of the queue
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The queue name.
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
        /// Rate limits for task dispatches.
        /// The queue's actual dispatch rate is the result of:
        /// * Number of tasks in the queue
        /// * User-specified throttling: rateLimits, retryConfig, and the queue's state.
        /// * System throttling due to 429 (Too Many Requests) or 503 (Service
        /// Unavailable) responses from the worker, high error rates, or to
        /// smooth sudden large traffic spikes.
        /// Structure is documented below.
        /// </summary>
        [Input("rateLimits")]
        public Input<Inputs.QueueRateLimitsGetArgs>? RateLimits { get; set; }

        /// <summary>
        /// Settings that determine the retry behavior.
        /// Structure is documented below.
        /// </summary>
        [Input("retryConfig")]
        public Input<Inputs.QueueRetryConfigGetArgs>? RetryConfig { get; set; }

        /// <summary>
        /// Configuration options for writing logs to Stackdriver Logging.
        /// Structure is documented below.
        /// </summary>
        [Input("stackdriverLoggingConfig")]
        public Input<Inputs.QueueStackdriverLoggingConfigGetArgs>? StackdriverLoggingConfig { get; set; }

        public QueueState()
        {
        }
    }
}
