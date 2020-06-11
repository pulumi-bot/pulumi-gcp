# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Intent(pulumi.CustomResource):
    action: pulumi.Output[str]
    """
    The name of the action associated with the intent.
    Note: The action name must not contain whitespaces.
    """
    default_response_platforms: pulumi.Output[list]
    """
    The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED
    (i.e. default platform).
    """
    display_name: pulumi.Output[str]
    """
    The name of this intent to be displayed on the console.
    """
    events: pulumi.Output[list]
    """
    The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of
    the contexts must be present in the active user session for an event to trigger this intent. See the
    [events reference](https://cloud.google.com/dialogflow/docs/events-overview) for more details.
    """
    followup_intent_infos: pulumi.Output[list]
    """
    Information about all followup intents that have this intent as a direct or indirect parent. We populate this field only
    in the output.

      * `followupIntentName` (`str`)
      * `parent_followup_intent_name` (`str`) - The unique identifier of the parent intent in the chain of followup intents.
        Format: projects/<Project ID>/agent/intents/<Intent ID>.
    """
    input_context_names: pulumi.Output[list]
    """
    The list of context names required for this intent to be triggered.
    Format: projects/<Project ID>/agent/sessions/-/contexts/<Context ID>.
    """
    is_fallback: pulumi.Output[bool]
    """
    Indicates whether this is a fallback intent.
    """
    ml_disabled: pulumi.Output[bool]
    """
    Indicates whether Machine Learning is disabled for the intent.
    Note: If mlDisabled setting is set to true, then this intent is not taken into account during inference in ML
    ONLY match mode. Also, auto-markup in the UI is turned off.
    """
    name: pulumi.Output[str]
    """
    The unique identifier of this intent. Format: projects/<Project ID>/agent/intents/<Intent ID>.
    """
    parent_followup_intent_name: pulumi.Output[str]
    """
    The unique identifier of the parent intent in the chain of followup intents.
    Format: projects/<Project ID>/agent/intents/<Intent ID>.
    """
    priority: pulumi.Output[float]
    """
    The priority of this intent. Higher numbers represent higher priorities.
    - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds
    to the Normal priority in the console.
    - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    reset_contexts: pulumi.Output[bool]
    """
    Indicates whether to delete all contexts in the current session when this intent is matched.
    """
    root_followup_intent_name: pulumi.Output[str]
    """
    The unique identifier of the root intent in the chain of followup intents. It identifies the correct followup intents
    chain for this intent. Format: projects/<Project ID>/agent/intents/<Intent ID>.
    """
    webhook_state: pulumi.Output[str]
    """
    Indicates whether webhooks are enabled for the intent.
    * WEBHOOK_STATE_ENABLED: Webhook is enabled in the agent and in the intent.
    * WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING: Webhook is enabled in the agent and in the intent. Also, each slot
    filling prompt is forwarded to the webhook.
    """
    def __init__(__self__, resource_name, opts=None, action=None, default_response_platforms=None, display_name=None, events=None, input_context_names=None, is_fallback=None, ml_disabled=None, parent_followup_intent_name=None, priority=None, project=None, reset_contexts=None, webhook_state=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents a Dialogflow intent. Intents convert a number of user expressions or patterns into an action. An action
        is an extraction of a user command or sentence semantics.


        To get more information about Intent, see:

        * [API documentation](https://cloud.google.com/dialogflow/docs/reference/rest/v2/projects.agent.intents)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/dialogflow/docs/)

        ## Example Usage - Dialogflow Intent Basic


        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic_agent = gcp.diagflow.Agent("basicAgent",
            display_name="example_agent",
            default_language_code="en",
            time_zone="America/New_York")
        basic_intent = gcp.diagflow.Intent("basicIntent", display_name="basic-intent")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The name of the action associated with the intent.
               Note: The action name must not contain whitespaces.
        :param pulumi.Input[list] default_response_platforms: The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED
               (i.e. default platform).
        :param pulumi.Input[str] display_name: The name of this intent to be displayed on the console.
        :param pulumi.Input[list] events: The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of
               the contexts must be present in the active user session for an event to trigger this intent. See the
               [events reference](https://cloud.google.com/dialogflow/docs/events-overview) for more details.
        :param pulumi.Input[list] input_context_names: The list of context names required for this intent to be triggered.
               Format: projects/<Project ID>/agent/sessions/-/contexts/<Context ID>.
        :param pulumi.Input[bool] is_fallback: Indicates whether this is a fallback intent.
        :param pulumi.Input[bool] ml_disabled: Indicates whether Machine Learning is disabled for the intent.
               Note: If mlDisabled setting is set to true, then this intent is not taken into account during inference in ML
               ONLY match mode. Also, auto-markup in the UI is turned off.
        :param pulumi.Input[str] parent_followup_intent_name: The unique identifier of the parent intent in the chain of followup intents.
               Format: projects/<Project ID>/agent/intents/<Intent ID>.
        :param pulumi.Input[float] priority: The priority of this intent. Higher numbers represent higher priorities.
               - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds
               to the Normal priority in the console.
               - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[bool] reset_contexts: Indicates whether to delete all contexts in the current session when this intent is matched.
        :param pulumi.Input[str] webhook_state: Indicates whether webhooks are enabled for the intent.
               * WEBHOOK_STATE_ENABLED: Webhook is enabled in the agent and in the intent.
               * WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING: Webhook is enabled in the agent and in the intent. Also, each slot
               filling prompt is forwarded to the webhook.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['action'] = action
            __props__['default_response_platforms'] = default_response_platforms
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['events'] = events
            __props__['input_context_names'] = input_context_names
            __props__['is_fallback'] = is_fallback
            __props__['ml_disabled'] = ml_disabled
            __props__['parent_followup_intent_name'] = parent_followup_intent_name
            __props__['priority'] = priority
            __props__['project'] = project
            __props__['reset_contexts'] = reset_contexts
            __props__['webhook_state'] = webhook_state
            __props__['followup_intent_infos'] = None
            __props__['name'] = None
            __props__['root_followup_intent_name'] = None
        super(Intent, __self__).__init__(
            'gcp:diagflow/intent:Intent',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, action=None, default_response_platforms=None, display_name=None, events=None, followup_intent_infos=None, input_context_names=None, is_fallback=None, ml_disabled=None, name=None, parent_followup_intent_name=None, priority=None, project=None, reset_contexts=None, root_followup_intent_name=None, webhook_state=None):
        """
        Get an existing Intent resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The name of the action associated with the intent.
               Note: The action name must not contain whitespaces.
        :param pulumi.Input[list] default_response_platforms: The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED
               (i.e. default platform).
        :param pulumi.Input[str] display_name: The name of this intent to be displayed on the console.
        :param pulumi.Input[list] events: The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of
               the contexts must be present in the active user session for an event to trigger this intent. See the
               [events reference](https://cloud.google.com/dialogflow/docs/events-overview) for more details.
        :param pulumi.Input[list] followup_intent_infos: Information about all followup intents that have this intent as a direct or indirect parent. We populate this field only
               in the output.
        :param pulumi.Input[list] input_context_names: The list of context names required for this intent to be triggered.
               Format: projects/<Project ID>/agent/sessions/-/contexts/<Context ID>.
        :param pulumi.Input[bool] is_fallback: Indicates whether this is a fallback intent.
        :param pulumi.Input[bool] ml_disabled: Indicates whether Machine Learning is disabled for the intent.
               Note: If mlDisabled setting is set to true, then this intent is not taken into account during inference in ML
               ONLY match mode. Also, auto-markup in the UI is turned off.
        :param pulumi.Input[str] name: The unique identifier of this intent. Format: projects/<Project ID>/agent/intents/<Intent ID>.
        :param pulumi.Input[str] parent_followup_intent_name: The unique identifier of the parent intent in the chain of followup intents.
               Format: projects/<Project ID>/agent/intents/<Intent ID>.
        :param pulumi.Input[float] priority: The priority of this intent. Higher numbers represent higher priorities.
               - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds
               to the Normal priority in the console.
               - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[bool] reset_contexts: Indicates whether to delete all contexts in the current session when this intent is matched.
        :param pulumi.Input[str] root_followup_intent_name: The unique identifier of the root intent in the chain of followup intents. It identifies the correct followup intents
               chain for this intent. Format: projects/<Project ID>/agent/intents/<Intent ID>.
        :param pulumi.Input[str] webhook_state: Indicates whether webhooks are enabled for the intent.
               * WEBHOOK_STATE_ENABLED: Webhook is enabled in the agent and in the intent.
               * WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING: Webhook is enabled in the agent and in the intent. Also, each slot
               filling prompt is forwarded to the webhook.

        The **followup_intent_infos** object supports the following:

          * `followupIntentName` (`pulumi.Input[str]`)
          * `parent_followup_intent_name` (`pulumi.Input[str]`) - The unique identifier of the parent intent in the chain of followup intents.
            Format: projects/<Project ID>/agent/intents/<Intent ID>.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["action"] = action
        __props__["default_response_platforms"] = default_response_platforms
        __props__["display_name"] = display_name
        __props__["events"] = events
        __props__["followup_intent_infos"] = followup_intent_infos
        __props__["input_context_names"] = input_context_names
        __props__["is_fallback"] = is_fallback
        __props__["ml_disabled"] = ml_disabled
        __props__["name"] = name
        __props__["parent_followup_intent_name"] = parent_followup_intent_name
        __props__["priority"] = priority
        __props__["project"] = project
        __props__["reset_contexts"] = reset_contexts
        __props__["root_followup_intent_name"] = root_followup_intent_name
        __props__["webhook_state"] = webhook_state
        return Intent(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
