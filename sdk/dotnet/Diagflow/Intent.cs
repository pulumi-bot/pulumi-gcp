// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Diagflow
{
    /// <summary>
    /// Represents a Dialogflow intent. Intents convert a number of user expressions or patterns into an action. An action
    /// is an extraction of a user command or sentence semantics.
    /// 
    /// To get more information about Intent, see:
    /// 
    /// * [API documentation](https://cloud.google.com/dialogflow/docs/reference/rest/v2/projects.agent.intents)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/dialogflow/docs/)
    /// 
    /// ## Example Usage
    /// ### Dialogflow Intent Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var basicAgent = new Gcp.Diagflow.Agent("basicAgent", new Gcp.Diagflow.AgentArgs
    ///         {
    ///             DisplayName = "example_agent",
    ///             DefaultLanguageCode = "en",
    ///             TimeZone = "America/New_York",
    ///         });
    ///         var basicIntent = new Gcp.Diagflow.Intent("basicIntent", new Gcp.Diagflow.IntentArgs
    ///         {
    ///             DisplayName = "basic-intent",
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 basicAgent,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Intent : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the action associated with the intent.
        /// Note: The action name must not contain whitespaces.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED
        /// (i.e. default platform).
        /// </summary>
        [Output("defaultResponsePlatforms")]
        public Output<ImmutableArray<string>> DefaultResponsePlatforms { get; private set; } = null!;

        /// <summary>
        /// The name of this intent to be displayed on the console.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of
        /// the contexts must be present in the active user session for an event to trigger this intent. See the
        /// [events reference](https://cloud.google.com/dialogflow/docs/events-overview) for more details.
        /// </summary>
        [Output("events")]
        public Output<ImmutableArray<string>> Events { get; private set; } = null!;

        /// <summary>
        /// Information about all followup intents that have this intent as a direct or indirect parent. We populate this field only
        /// in the output.
        /// </summary>
        [Output("followupIntentInfos")]
        public Output<ImmutableArray<Outputs.IntentFollowupIntentInfo>> FollowupIntentInfos { get; private set; } = null!;

        /// <summary>
        /// The list of context names required for this intent to be triggered.
        /// Format: projects/&lt;Project ID&gt;/agent/sessions/-/contexts/&lt;Context ID&gt;.
        /// </summary>
        [Output("inputContextNames")]
        public Output<ImmutableArray<string>> InputContextNames { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this is a fallback intent.
        /// </summary>
        [Output("isFallback")]
        public Output<bool> IsFallback { get; private set; } = null!;

        /// <summary>
        /// Indicates whether Machine Learning is disabled for the intent.
        /// Note: If mlDisabled setting is set to true, then this intent is not taken into account during inference in ML
        /// ONLY match mode. Also, auto-markup in the UI is turned off.
        /// </summary>
        [Output("mlDisabled")]
        public Output<bool> MlDisabled { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of this intent. Format: projects/&lt;Project ID&gt;/agent/intents/&lt;Intent ID&gt;.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the parent intent in the chain of followup intents.
        /// Format: projects/&lt;Project ID&gt;/agent/intents/&lt;Intent ID&gt;.
        /// </summary>
        [Output("parentFollowupIntentName")]
        public Output<string> ParentFollowupIntentName { get; private set; } = null!;

        /// <summary>
        /// The priority of this intent. Higher numbers represent higher priorities.
        /// - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds
        /// to the Normal priority in the console.
        /// - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to delete all contexts in the current session when this intent is matched.
        /// </summary>
        [Output("resetContexts")]
        public Output<bool> ResetContexts { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the root intent in the chain of followup intents. It identifies the correct followup intents
        /// chain for this intent. Format: projects/&lt;Project ID&gt;/agent/intents/&lt;Intent ID&gt;.
        /// </summary>
        [Output("rootFollowupIntentName")]
        public Output<string> RootFollowupIntentName { get; private set; } = null!;

        /// <summary>
        /// Indicates whether webhooks are enabled for the intent.
        /// * WEBHOOK_STATE_ENABLED: Webhook is enabled in the agent and in the intent.
        /// * WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING: Webhook is enabled in the agent and in the intent. Also, each slot
        /// filling prompt is forwarded to the webhook.
        /// </summary>
        [Output("webhookState")]
        public Output<string> WebhookState { get; private set; } = null!;


        /// <summary>
        /// Create a Intent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Intent(string name, IntentArgs args, CustomResourceOptions? options = null)
            : base("gcp:diagflow/intent:Intent", name, args ?? new IntentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Intent(string name, Input<string> id, IntentState? state = null, CustomResourceOptions? options = null)
            : base("gcp:diagflow/intent:Intent", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Intent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Intent Get(string name, Input<string> id, IntentState? state = null, CustomResourceOptions? options = null)
        {
            return new Intent(name, id, state, options);
        }
    }

    public sealed class IntentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the action associated with the intent.
        /// Note: The action name must not contain whitespaces.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("defaultResponsePlatforms")]
        private InputList<string>? _defaultResponsePlatforms;

        /// <summary>
        /// The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED
        /// (i.e. default platform).
        /// </summary>
        public InputList<string> DefaultResponsePlatforms
        {
            get => _defaultResponsePlatforms ?? (_defaultResponsePlatforms = new InputList<string>());
            set => _defaultResponsePlatforms = value;
        }

        /// <summary>
        /// The name of this intent to be displayed on the console.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        [Input("events")]
        private InputList<string>? _events;

        /// <summary>
        /// The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of
        /// the contexts must be present in the active user session for an event to trigger this intent. See the
        /// [events reference](https://cloud.google.com/dialogflow/docs/events-overview) for more details.
        /// </summary>
        public InputList<string> Events
        {
            get => _events ?? (_events = new InputList<string>());
            set => _events = value;
        }

        [Input("inputContextNames")]
        private InputList<string>? _inputContextNames;

        /// <summary>
        /// The list of context names required for this intent to be triggered.
        /// Format: projects/&lt;Project ID&gt;/agent/sessions/-/contexts/&lt;Context ID&gt;.
        /// </summary>
        public InputList<string> InputContextNames
        {
            get => _inputContextNames ?? (_inputContextNames = new InputList<string>());
            set => _inputContextNames = value;
        }

        /// <summary>
        /// Indicates whether this is a fallback intent.
        /// </summary>
        [Input("isFallback")]
        public Input<bool>? IsFallback { get; set; }

        /// <summary>
        /// Indicates whether Machine Learning is disabled for the intent.
        /// Note: If mlDisabled setting is set to true, then this intent is not taken into account during inference in ML
        /// ONLY match mode. Also, auto-markup in the UI is turned off.
        /// </summary>
        [Input("mlDisabled")]
        public Input<bool>? MlDisabled { get; set; }

        /// <summary>
        /// The unique identifier of the parent intent in the chain of followup intents.
        /// Format: projects/&lt;Project ID&gt;/agent/intents/&lt;Intent ID&gt;.
        /// </summary>
        [Input("parentFollowupIntentName")]
        public Input<string>? ParentFollowupIntentName { get; set; }

        /// <summary>
        /// The priority of this intent. Higher numbers represent higher priorities.
        /// - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds
        /// to the Normal priority in the console.
        /// - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Indicates whether to delete all contexts in the current session when this intent is matched.
        /// </summary>
        [Input("resetContexts")]
        public Input<bool>? ResetContexts { get; set; }

        /// <summary>
        /// Indicates whether webhooks are enabled for the intent.
        /// * WEBHOOK_STATE_ENABLED: Webhook is enabled in the agent and in the intent.
        /// * WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING: Webhook is enabled in the agent and in the intent. Also, each slot
        /// filling prompt is forwarded to the webhook.
        /// </summary>
        [Input("webhookState")]
        public Input<string>? WebhookState { get; set; }

        public IntentArgs()
        {
        }
    }

    public sealed class IntentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the action associated with the intent.
        /// Note: The action name must not contain whitespaces.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("defaultResponsePlatforms")]
        private InputList<string>? _defaultResponsePlatforms;

        /// <summary>
        /// The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED
        /// (i.e. default platform).
        /// </summary>
        public InputList<string> DefaultResponsePlatforms
        {
            get => _defaultResponsePlatforms ?? (_defaultResponsePlatforms = new InputList<string>());
            set => _defaultResponsePlatforms = value;
        }

        /// <summary>
        /// The name of this intent to be displayed on the console.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("events")]
        private InputList<string>? _events;

        /// <summary>
        /// The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of
        /// the contexts must be present in the active user session for an event to trigger this intent. See the
        /// [events reference](https://cloud.google.com/dialogflow/docs/events-overview) for more details.
        /// </summary>
        public InputList<string> Events
        {
            get => _events ?? (_events = new InputList<string>());
            set => _events = value;
        }

        [Input("followupIntentInfos")]
        private InputList<Inputs.IntentFollowupIntentInfoGetArgs>? _followupIntentInfos;

        /// <summary>
        /// Information about all followup intents that have this intent as a direct or indirect parent. We populate this field only
        /// in the output.
        /// </summary>
        public InputList<Inputs.IntentFollowupIntentInfoGetArgs> FollowupIntentInfos
        {
            get => _followupIntentInfos ?? (_followupIntentInfos = new InputList<Inputs.IntentFollowupIntentInfoGetArgs>());
            set => _followupIntentInfos = value;
        }

        [Input("inputContextNames")]
        private InputList<string>? _inputContextNames;

        /// <summary>
        /// The list of context names required for this intent to be triggered.
        /// Format: projects/&lt;Project ID&gt;/agent/sessions/-/contexts/&lt;Context ID&gt;.
        /// </summary>
        public InputList<string> InputContextNames
        {
            get => _inputContextNames ?? (_inputContextNames = new InputList<string>());
            set => _inputContextNames = value;
        }

        /// <summary>
        /// Indicates whether this is a fallback intent.
        /// </summary>
        [Input("isFallback")]
        public Input<bool>? IsFallback { get; set; }

        /// <summary>
        /// Indicates whether Machine Learning is disabled for the intent.
        /// Note: If mlDisabled setting is set to true, then this intent is not taken into account during inference in ML
        /// ONLY match mode. Also, auto-markup in the UI is turned off.
        /// </summary>
        [Input("mlDisabled")]
        public Input<bool>? MlDisabled { get; set; }

        /// <summary>
        /// The unique identifier of this intent. Format: projects/&lt;Project ID&gt;/agent/intents/&lt;Intent ID&gt;.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The unique identifier of the parent intent in the chain of followup intents.
        /// Format: projects/&lt;Project ID&gt;/agent/intents/&lt;Intent ID&gt;.
        /// </summary>
        [Input("parentFollowupIntentName")]
        public Input<string>? ParentFollowupIntentName { get; set; }

        /// <summary>
        /// The priority of this intent. Higher numbers represent higher priorities.
        /// - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds
        /// to the Normal priority in the console.
        /// - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Indicates whether to delete all contexts in the current session when this intent is matched.
        /// </summary>
        [Input("resetContexts")]
        public Input<bool>? ResetContexts { get; set; }

        /// <summary>
        /// The unique identifier of the root intent in the chain of followup intents. It identifies the correct followup intents
        /// chain for this intent. Format: projects/&lt;Project ID&gt;/agent/intents/&lt;Intent ID&gt;.
        /// </summary>
        [Input("rootFollowupIntentName")]
        public Input<string>? RootFollowupIntentName { get; set; }

        /// <summary>
        /// Indicates whether webhooks are enabled for the intent.
        /// * WEBHOOK_STATE_ENABLED: Webhook is enabled in the agent and in the intent.
        /// * WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING: Webhook is enabled in the agent and in the intent. Also, each slot
        /// filling prompt is forwarded to the webhook.
        /// </summary>
        [Input("webhookState")]
        public Input<string>? WebhookState { get; set; }

        public IntentState()
        {
        }
    }
}
