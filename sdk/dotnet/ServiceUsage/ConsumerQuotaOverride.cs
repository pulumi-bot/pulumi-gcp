// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ServiceUsage
{
    /// <summary>
    /// A consumer override is applied to the consumer on its own authority to limit its own quota usage.
    /// Consumer overrides cannot be used to grant more quota than would be allowed by admin overrides,
    /// producer overrides, or the default limit of the service.
    /// 
    /// To get more information about ConsumerQuotaOverride, see:
    /// 
    /// * How-to Guides
    ///     * [Getting Started](https://cloud.google.com/service-usage/docs/getting-started)
    ///     * [REST API documentation](https://cloud.google.com/service-usage/docs/reference/rest/v1beta1/services.consumerQuotaMetrics.limits.consumerOverrides)
    /// 
    /// ## Example Usage
    /// ### Consumer Quota Override
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var myProject = new Gcp.Organizations.Project("myProject", new Gcp.Organizations.ProjectArgs
    ///         {
    ///             ProjectId = "quota",
    ///             OrgId = "123456789",
    ///         }, new CustomResourceOptions {
    ///             Provider = google_beta,
    ///         });
    ///         var @override = new Gcp.ServiceUsage.ConsumerQuotaOverride("override", new Gcp.ServiceUsage.ConsumerQuotaOverrideArgs
    ///         {
    ///             Project = myProject.ProjectId,
    ///             Service = "servicemanagement.googleapis.com",
    ///             Metric = "servicemanagement.googleapis.com%2Fdefault_requests",
    ///             Limit = "%2Fmin%2Fproject",
    ///             OverrideValue = "95",
    ///             Force = true,
    ///         }, new CustomResourceOptions {
    ///             Provider = google_beta,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ConsumerQuotaOverride : Pulumi.CustomResource
    {
        /// <summary>
        /// If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
        /// </summary>
        [Output("dimensions")]
        public Output<ImmutableDictionary<string, string>?> Dimensions { get; private set; } = null!;

        /// <summary>
        /// If the new quota would decrease the existing quota by more than 10%, the request is rejected.
        /// If `force` is `true`, that safety check is ignored.
        /// </summary>
        [Output("force")]
        public Output<bool?> Force { get; private set; } = null!;

        /// <summary>
        /// The limit on the metric, e.g. `/project/region`.
        /// </summary>
        [Output("limit")]
        public Output<string> Limit { get; private set; } = null!;

        /// <summary>
        /// The metric that should be limited, e.g. `compute.googleapis.com/cpus`.
        /// </summary>
        [Output("metric")]
        public Output<string> Metric { get; private set; } = null!;

        /// <summary>
        /// The server-generated name of the quota override.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
        /// </summary>
        [Output("overrideValue")]
        public Output<string> OverrideValue { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The service that the metrics belong to, e.g. `compute.googleapis.com`.
        /// </summary>
        [Output("service")]
        public Output<string> Service { get; private set; } = null!;


        /// <summary>
        /// Create a ConsumerQuotaOverride resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConsumerQuotaOverride(string name, ConsumerQuotaOverrideArgs args, CustomResourceOptions? options = null)
            : base("gcp:serviceusage/consumerQuotaOverride:ConsumerQuotaOverride", name, args ?? new ConsumerQuotaOverrideArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConsumerQuotaOverride(string name, Input<string> id, ConsumerQuotaOverrideState? state = null, CustomResourceOptions? options = null)
            : base("gcp:serviceusage/consumerQuotaOverride:ConsumerQuotaOverride", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConsumerQuotaOverride resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConsumerQuotaOverride Get(string name, Input<string> id, ConsumerQuotaOverrideState? state = null, CustomResourceOptions? options = null)
        {
            return new ConsumerQuotaOverride(name, id, state, options);
        }
    }

    public sealed class ConsumerQuotaOverrideArgs : Pulumi.ResourceArgs
    {
        [Input("dimensions")]
        private InputMap<string>? _dimensions;

        /// <summary>
        /// If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
        /// </summary>
        public InputMap<string> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputMap<string>());
            set => _dimensions = value;
        }

        /// <summary>
        /// If the new quota would decrease the existing quota by more than 10%, the request is rejected.
        /// If `force` is `true`, that safety check is ignored.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// The limit on the metric, e.g. `/project/region`.
        /// </summary>
        [Input("limit", required: true)]
        public Input<string> Limit { get; set; } = null!;

        /// <summary>
        /// The metric that should be limited, e.g. `compute.googleapis.com/cpus`.
        /// </summary>
        [Input("metric", required: true)]
        public Input<string> Metric { get; set; } = null!;

        /// <summary>
        /// The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
        /// </summary>
        [Input("overrideValue", required: true)]
        public Input<string> OverrideValue { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The service that the metrics belong to, e.g. `compute.googleapis.com`.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public ConsumerQuotaOverrideArgs()
        {
        }
    }

    public sealed class ConsumerQuotaOverrideState : Pulumi.ResourceArgs
    {
        [Input("dimensions")]
        private InputMap<string>? _dimensions;

        /// <summary>
        /// If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
        /// </summary>
        public InputMap<string> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputMap<string>());
            set => _dimensions = value;
        }

        /// <summary>
        /// If the new quota would decrease the existing quota by more than 10%, the request is rejected.
        /// If `force` is `true`, that safety check is ignored.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// The limit on the metric, e.g. `/project/region`.
        /// </summary>
        [Input("limit")]
        public Input<string>? Limit { get; set; }

        /// <summary>
        /// The metric that should be limited, e.g. `compute.googleapis.com/cpus`.
        /// </summary>
        [Input("metric")]
        public Input<string>? Metric { get; set; }

        /// <summary>
        /// The server-generated name of the quota override.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
        /// </summary>
        [Input("overrideValue")]
        public Input<string>? OverrideValue { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The service that the metrics belong to, e.g. `compute.googleapis.com`.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        public ConsumerQuotaOverrideState()
        {
        }
    }
}
