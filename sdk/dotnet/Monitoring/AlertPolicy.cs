// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring
{
    /// <summary>
    /// A description of the conditions under which some aspect of your system is
    /// considered to be "unhealthy" and the ways to notify people or services
    /// about this state.
    /// 
    /// 
    /// To get more information about AlertPolicy, see:
    /// 
    /// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.alertPolicies)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/monitoring/alerts/)
    /// 
    /// ## Example Usage - Monitoring Alert Policy Basic
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
    ///         var alertPolicy = new Gcp.Monitoring.AlertPolicy("alertPolicy", new Gcp.Monitoring.AlertPolicyArgs
    ///         {
    ///             Combiner = "OR",
    ///             Conditions = 
    ///             {
    ///                 new Gcp.Monitoring.Inputs.AlertPolicyConditionArgs
    ///                 {
    ///                     ConditionThreshold = new Gcp.Monitoring.Inputs.AlertPolicyConditionConditionThresholdArgs
    ///                     {
    ///                         Aggregations = 
    ///                         {
    ///                             new Gcp.Monitoring.Inputs.AlertPolicyConditionConditionThresholdAggregationArgs
    ///                             {
    ///                                 AlignmentPeriod = "60s",
    ///                                 PerSeriesAligner = "ALIGN_RATE",
    ///                             },
    ///                         },
    ///                         Comparison = "COMPARISON_GT",
    ///                         Duration = "60s",
    ///                         Filter = "metric.type=\"compute.googleapis.com/instance/disk/write_bytes_count\" AND resource.type=\"gce_instance\"",
    ///                     },
    ///                     DisplayName = "test condition",
    ///                 },
    ///             },
    ///             DisplayName = "My Alert Policy",
    ///             UserLabels = 
    ///             {
    ///                 { "foo", "bar" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class AlertPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// How to combine the results of multiple conditions to
        /// determine if an incident should be opened.
        /// </summary>
        [Output("combiner")]
        public Output<string> Combiner { get; private set; } = null!;

        /// <summary>
        /// A list of conditions for the policy. The conditions are combined by
        /// AND or OR according to the combiner field. If the combined conditions
        /// evaluate to true, then an incident is created. A policy can have from
        /// one to six conditions.  Structure is documented below.
        /// </summary>
        [Output("conditions")]
        public Output<ImmutableArray<Outputs.AlertPolicyCondition>> Conditions { get; private set; } = null!;

        /// <summary>
        /// A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be
        /// ignored.
        /// </summary>
        [Output("creationRecord")]
        public Output<Outputs.AlertPolicyCreationRecord> CreationRecord { get; private set; } = null!;

        /// <summary>
        /// A short name or phrase used to identify the
        /// condition in dashboards, notifications, and
        /// incidents. To avoid confusion, don't use the same
        /// display name for multiple conditions in the same
        /// policy.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// A short name or phrase used to identify the policy in dashboards,
        /// notifications, and incidents. To avoid confusion, don't use the same
        /// display name for multiple policies in the same project. The name is
        /// limited to 512 Unicode characters.  Structure is documented below.
        /// </summary>
        [Output("documentation")]
        public Output<Outputs.AlertPolicyDocumentation?> Documentation { get; private set; } = null!;

        /// <summary>
        /// Whether or not the policy is enabled. The default is true.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// -
        /// The unique resource name for this condition.
        /// Its syntax is:
        /// projects/[PROJECT_ID]/alertPolicies/[POLICY_ID]/conditions/[CONDITION_ID]
        /// [CONDITION_ID] is assigned by Stackdriver Monitoring when
        /// the condition is created as part of a new or updated alerting
        /// policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Identifies the notification channels to which notifications should be
        /// sent when incidents are opened or closed or when new violations occur
        /// on an already opened incident. Each element of this array corresponds
        /// to the name field in each of the NotificationChannel objects that are
        /// returned from the notificationChannels.list method. The syntax of the
        /// entries in this field is
        /// `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]`
        /// </summary>
        [Output("notificationChannels")]
        public Output<ImmutableArray<string>> NotificationChannels { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// This field is intended to be used for organizing and identifying the AlertPolicy
        /// objects.The field can contain up to 64 entries. Each key and value is limited
        /// to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
        /// can contain only lowercase letters, numerals, underscores, and dashes. Keys
        /// must begin with a letter.
        /// </summary>
        [Output("userLabels")]
        public Output<ImmutableDictionary<string, string>?> UserLabels { get; private set; } = null!;


        /// <summary>
        /// Create a AlertPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AlertPolicy(string name, AlertPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:monitoring/alertPolicy:AlertPolicy", name, args ?? new AlertPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AlertPolicy(string name, Input<string> id, AlertPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:monitoring/alertPolicy:AlertPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AlertPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AlertPolicy Get(string name, Input<string> id, AlertPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new AlertPolicy(name, id, state, options);
        }
    }

    public sealed class AlertPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How to combine the results of multiple conditions to
        /// determine if an incident should be opened.
        /// </summary>
        [Input("combiner", required: true)]
        public Input<string> Combiner { get; set; } = null!;

        [Input("conditions", required: true)]
        private InputList<Inputs.AlertPolicyConditionArgs>? _conditions;

        /// <summary>
        /// A list of conditions for the policy. The conditions are combined by
        /// AND or OR according to the combiner field. If the combined conditions
        /// evaluate to true, then an incident is created. A policy can have from
        /// one to six conditions.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.AlertPolicyConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.AlertPolicyConditionArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// A short name or phrase used to identify the
        /// condition in dashboards, notifications, and
        /// incidents. To avoid confusion, don't use the same
        /// display name for multiple conditions in the same
        /// policy.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// A short name or phrase used to identify the policy in dashboards,
        /// notifications, and incidents. To avoid confusion, don't use the same
        /// display name for multiple policies in the same project. The name is
        /// limited to 512 Unicode characters.  Structure is documented below.
        /// </summary>
        [Input("documentation")]
        public Input<Inputs.AlertPolicyDocumentationArgs>? Documentation { get; set; }

        /// <summary>
        /// Whether or not the policy is enabled. The default is true.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("notificationChannels")]
        private InputList<string>? _notificationChannels;

        /// <summary>
        /// Identifies the notification channels to which notifications should be
        /// sent when incidents are opened or closed or when new violations occur
        /// on an already opened incident. Each element of this array corresponds
        /// to the name field in each of the NotificationChannel objects that are
        /// returned from the notificationChannels.list method. The syntax of the
        /// entries in this field is
        /// `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]`
        /// </summary>
        public InputList<string> NotificationChannels
        {
            get => _notificationChannels ?? (_notificationChannels = new InputList<string>());
            set => _notificationChannels = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("userLabels")]
        private InputMap<string>? _userLabels;

        /// <summary>
        /// This field is intended to be used for organizing and identifying the AlertPolicy
        /// objects.The field can contain up to 64 entries. Each key and value is limited
        /// to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
        /// can contain only lowercase letters, numerals, underscores, and dashes. Keys
        /// must begin with a letter.
        /// </summary>
        public InputMap<string> UserLabels
        {
            get => _userLabels ?? (_userLabels = new InputMap<string>());
            set => _userLabels = value;
        }

        public AlertPolicyArgs()
        {
        }
    }

    public sealed class AlertPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How to combine the results of multiple conditions to
        /// determine if an incident should be opened.
        /// </summary>
        [Input("combiner")]
        public Input<string>? Combiner { get; set; }

        [Input("conditions")]
        private InputList<Inputs.AlertPolicyConditionGetArgs>? _conditions;

        /// <summary>
        /// A list of conditions for the policy. The conditions are combined by
        /// AND or OR according to the combiner field. If the combined conditions
        /// evaluate to true, then an incident is created. A policy can have from
        /// one to six conditions.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.AlertPolicyConditionGetArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.AlertPolicyConditionGetArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be
        /// ignored.
        /// </summary>
        [Input("creationRecord")]
        public Input<Inputs.AlertPolicyCreationRecordGetArgs>? CreationRecord { get; set; }

        /// <summary>
        /// A short name or phrase used to identify the
        /// condition in dashboards, notifications, and
        /// incidents. To avoid confusion, don't use the same
        /// display name for multiple conditions in the same
        /// policy.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// A short name or phrase used to identify the policy in dashboards,
        /// notifications, and incidents. To avoid confusion, don't use the same
        /// display name for multiple policies in the same project. The name is
        /// limited to 512 Unicode characters.  Structure is documented below.
        /// </summary>
        [Input("documentation")]
        public Input<Inputs.AlertPolicyDocumentationGetArgs>? Documentation { get; set; }

        /// <summary>
        /// Whether or not the policy is enabled. The default is true.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// -
        /// The unique resource name for this condition.
        /// Its syntax is:
        /// projects/[PROJECT_ID]/alertPolicies/[POLICY_ID]/conditions/[CONDITION_ID]
        /// [CONDITION_ID] is assigned by Stackdriver Monitoring when
        /// the condition is created as part of a new or updated alerting
        /// policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationChannels")]
        private InputList<string>? _notificationChannels;

        /// <summary>
        /// Identifies the notification channels to which notifications should be
        /// sent when incidents are opened or closed or when new violations occur
        /// on an already opened incident. Each element of this array corresponds
        /// to the name field in each of the NotificationChannel objects that are
        /// returned from the notificationChannels.list method. The syntax of the
        /// entries in this field is
        /// `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]`
        /// </summary>
        public InputList<string> NotificationChannels
        {
            get => _notificationChannels ?? (_notificationChannels = new InputList<string>());
            set => _notificationChannels = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("userLabels")]
        private InputMap<string>? _userLabels;

        /// <summary>
        /// This field is intended to be used for organizing and identifying the AlertPolicy
        /// objects.The field can contain up to 64 entries. Each key and value is limited
        /// to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
        /// can contain only lowercase letters, numerals, underscores, and dashes. Keys
        /// must begin with a letter.
        /// </summary>
        public InputMap<string> UserLabels
        {
            get => _userLabels ?? (_userLabels = new InputMap<string>());
            set => _userLabels = value;
        }

        public AlertPolicyState()
        {
        }
    }
}
