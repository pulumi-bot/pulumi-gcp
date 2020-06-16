// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Logging
{
    /// <summary>
    /// Manages a billing account logging exclusion. For more information see
    /// [the official documentation](https://cloud.google.com/logging/docs/) and
    /// [Excluding Logs](https://cloud.google.com/logging/docs/exclusions).
    /// 
    /// Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
    /// granted to the credentials used with the provider.
    /// 
    /// {{% examples %}}
    /// ## Example Usage
    /// {{% example %}}
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var my_exclusion = new Gcp.Logging.BillingAccountExclusion("my-exclusion", new Gcp.Logging.BillingAccountExclusionArgs
    ///         {
    ///             BillingAccount = "ABCDEF-012345-GHIJKL",
    ///             Description = "Exclude GCE instance debug logs",
    ///             Filter = "resource.type = gce_instance AND severity &lt;= DEBUG",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// {{% /example %}}
    /// {{% /examples %}}
    /// </summary>
    public partial class BillingAccountExclusion : Pulumi.CustomResource
    {
        /// <summary>
        /// The billing account to create the exclusion for.
        /// </summary>
        [Output("billingAccount")]
        public Output<string> BillingAccount { get; private set; } = null!;

        /// <summary>
        /// A human-readable description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether this exclusion rule should be disabled or not. This defaults to
        /// false.
        /// </summary>
        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
        /// write a filter.
        /// </summary>
        [Output("filter")]
        public Output<string> Filter { get; private set; } = null!;

        /// <summary>
        /// The name of the logging exclusion.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a BillingAccountExclusion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BillingAccountExclusion(string name, BillingAccountExclusionArgs args, CustomResourceOptions? options = null)
            : base("gcp:logging/billingAccountExclusion:BillingAccountExclusion", name, args ?? new BillingAccountExclusionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BillingAccountExclusion(string name, Input<string> id, BillingAccountExclusionState? state = null, CustomResourceOptions? options = null)
            : base("gcp:logging/billingAccountExclusion:BillingAccountExclusion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BillingAccountExclusion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BillingAccountExclusion Get(string name, Input<string> id, BillingAccountExclusionState? state = null, CustomResourceOptions? options = null)
        {
            return new BillingAccountExclusion(name, id, state, options);
        }
    }

    public sealed class BillingAccountExclusionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The billing account to create the exclusion for.
        /// </summary>
        [Input("billingAccount", required: true)]
        public Input<string> BillingAccount { get; set; } = null!;

        /// <summary>
        /// A human-readable description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether this exclusion rule should be disabled or not. This defaults to
        /// false.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
        /// write a filter.
        /// </summary>
        [Input("filter", required: true)]
        public Input<string> Filter { get; set; } = null!;

        /// <summary>
        /// The name of the logging exclusion.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public BillingAccountExclusionArgs()
        {
        }
    }

    public sealed class BillingAccountExclusionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The billing account to create the exclusion for.
        /// </summary>
        [Input("billingAccount")]
        public Input<string>? BillingAccount { get; set; }

        /// <summary>
        /// A human-readable description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether this exclusion rule should be disabled or not. This defaults to
        /// false.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
        /// write a filter.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The name of the logging exclusion.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public BillingAccountExclusionState()
        {
        }
    }
}
