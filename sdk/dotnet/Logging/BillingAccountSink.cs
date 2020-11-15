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
    /// Manages a billing account logging sink. For more information see
    /// [the official documentation](https://cloud.google.com/logging/docs/) and
    /// [Exporting Logs in the API](https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
    /// 
    /// &gt; **Note** You must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
    /// [granted on the billing account](https://cloud.google.com/billing/reference/rest/v1/billingAccounts/getIamPolicy) to
    /// the credentials used with this provider. [IAM roles granted on a billing account](https://cloud.google.com/billing/docs/how-to/billing-access) are separate from the
    /// typical IAM roles granted on a project.
    /// 
    /// ## Import
    /// 
    /// Billing account logging sinks can be imported using this format
    /// 
    /// ```sh
    ///  $ pulumi import gcp:logging/billingAccountSink:BillingAccountSink my_sink billingAccounts/{{billing_account_id}}/sinks/{{sink_id}}
    /// ```
    /// </summary>
    public partial class BillingAccountSink : Pulumi.CustomResource
    {
        /// <summary>
        /// Options that affect sinks exporting data to BigQuery. Structure documented below.
        /// </summary>
        [Output("bigqueryOptions")]
        public Output<Outputs.BillingAccountSinkBigqueryOptions> BigqueryOptions { get; private set; } = null!;

        /// <summary>
        /// The billing account exported to the sink.
        /// </summary>
        [Output("billingAccount")]
        public Output<string> BillingAccount { get; private set; } = null!;

        /// <summary>
        /// The destination of the sink (or, in other words, where logs are written to). Can be a
        /// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// The writer associated with the sink must have access to write to the above resource.
        /// </summary>
        [Output("destination")]
        public Output<string> Destination { get; private set; } = null!;

        /// <summary>
        /// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
        /// one of exclusion_filters it will not be exported.
        /// </summary>
        [Output("exclusions")]
        public Output<ImmutableArray<Outputs.BillingAccountSinkExclusion>> Exclusions { get; private set; } = null!;

        /// <summary>
        /// The filter to apply when exporting logs. Only log entries that match the filter are exported.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        /// write a filter.
        /// </summary>
        [Output("filter")]
        public Output<string?> Filter { get; private set; } = null!;

        /// <summary>
        /// The name of the logging sink.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The identity associated with this sink. This identity must be granted write access to the
        /// configured `destination`.
        /// </summary>
        [Output("writerIdentity")]
        public Output<string> WriterIdentity { get; private set; } = null!;


        /// <summary>
        /// Create a BillingAccountSink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BillingAccountSink(string name, BillingAccountSinkArgs args, CustomResourceOptions? options = null)
            : base("gcp:logging/billingAccountSink:BillingAccountSink", name, args ?? new BillingAccountSinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BillingAccountSink(string name, Input<string> id, BillingAccountSinkState? state = null, CustomResourceOptions? options = null)
            : base("gcp:logging/billingAccountSink:BillingAccountSink", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BillingAccountSink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BillingAccountSink Get(string name, Input<string> id, BillingAccountSinkState? state = null, CustomResourceOptions? options = null)
        {
            return new BillingAccountSink(name, id, state, options);
        }
    }

    public sealed class BillingAccountSinkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Options that affect sinks exporting data to BigQuery. Structure documented below.
        /// </summary>
        [Input("bigqueryOptions")]
        public Input<Inputs.BillingAccountSinkBigqueryOptionsArgs>? BigqueryOptions { get; set; }

        /// <summary>
        /// The billing account exported to the sink.
        /// </summary>
        [Input("billingAccount", required: true)]
        public Input<string> BillingAccount { get; set; } = null!;

        /// <summary>
        /// The destination of the sink (or, in other words, where logs are written to). Can be a
        /// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// The writer associated with the sink must have access to write to the above resource.
        /// </summary>
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

        [Input("exclusions")]
        private InputList<Inputs.BillingAccountSinkExclusionArgs>? _exclusions;

        /// <summary>
        /// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
        /// one of exclusion_filters it will not be exported.
        /// </summary>
        public InputList<Inputs.BillingAccountSinkExclusionArgs> Exclusions
        {
            get => _exclusions ?? (_exclusions = new InputList<Inputs.BillingAccountSinkExclusionArgs>());
            set => _exclusions = value;
        }

        /// <summary>
        /// The filter to apply when exporting logs. Only log entries that match the filter are exported.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        /// write a filter.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The name of the logging sink.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public BillingAccountSinkArgs()
        {
        }
    }

    public sealed class BillingAccountSinkState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Options that affect sinks exporting data to BigQuery. Structure documented below.
        /// </summary>
        [Input("bigqueryOptions")]
        public Input<Inputs.BillingAccountSinkBigqueryOptionsGetArgs>? BigqueryOptions { get; set; }

        /// <summary>
        /// The billing account exported to the sink.
        /// </summary>
        [Input("billingAccount")]
        public Input<string>? BillingAccount { get; set; }

        /// <summary>
        /// The destination of the sink (or, in other words, where logs are written to). Can be a
        /// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// The writer associated with the sink must have access to write to the above resource.
        /// </summary>
        [Input("destination")]
        public Input<string>? Destination { get; set; }

        [Input("exclusions")]
        private InputList<Inputs.BillingAccountSinkExclusionGetArgs>? _exclusions;

        /// <summary>
        /// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
        /// one of exclusion_filters it will not be exported.
        /// </summary>
        public InputList<Inputs.BillingAccountSinkExclusionGetArgs> Exclusions
        {
            get => _exclusions ?? (_exclusions = new InputList<Inputs.BillingAccountSinkExclusionGetArgs>());
            set => _exclusions = value;
        }

        /// <summary>
        /// The filter to apply when exporting logs. Only log entries that match the filter are exported.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        /// write a filter.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The name of the logging sink.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The identity associated with this sink. This identity must be granted write access to the
        /// configured `destination`.
        /// </summary>
        [Input("writerIdentity")]
        public Input<string>? WriterIdentity { get; set; }

        public BillingAccountSinkState()
        {
        }
    }
}
