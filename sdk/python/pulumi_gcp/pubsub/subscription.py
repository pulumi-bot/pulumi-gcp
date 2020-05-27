# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Subscription(pulumi.CustomResource):
    ack_deadline_seconds: pulumi.Output[float]
    """
    This value is the maximum time after a subscriber receives a message
    before the subscriber should acknowledge the message. After message
    delivery but before the ack deadline expires and before the message is
    acknowledged, it is an outstanding message and will not be delivered
    again during that time (on a best-effort basis).
    For pull subscriptions, this value is used as the initial value for
    the ack deadline. To override this value for a given message, call
    subscriptions.modifyAckDeadline with the corresponding ackId if using
    pull. The minimum custom deadline you can specify is 10 seconds. The
    maximum custom deadline you can specify is 600 seconds (10 minutes).
    If this parameter is 0, a default value of 10 seconds is used.
    For push delivery, this value is also used to set the request timeout
    for the call to the push endpoint.
    If the subscriber never acknowledges the message, the Pub/Sub system
    will eventually redeliver the message.
    """
    dead_letter_policy: pulumi.Output[dict]
    """
    A policy that specifies the conditions for dead lettering messages in
    this subscription. If dead_letter_policy is not set, dead lettering
    is disabled.
    The Cloud Pub/Sub service account associated with this subscriptions's
    parent project (i.e.,
    service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
    permission to Acknowledge() messages on this subscription.  Structure is documented below.

      * `deadLetterTopic` (`str`) - The name of the topic to which dead letter messages should be published.
        Format is `projects/{project}/topics/{topic}`.
        The Cloud Pub/Sub service\naccount associated with the enclosing subscription's
        parent project (i.e.,
        service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
        permission to Publish() to this topic.
        The operation will fail if the topic does not exist.
        Users should ensure that there is a subscription attached to this topic
        since messages published to a topic with no subscriptions are lost.
      * `maxDeliveryAttempts` (`float`) - The maximum number of delivery attempts for any message. The value must be
        between 5 and 100.
        The number of delivery attempts is defined as 1 + (the sum of number of
        NACKs and number of times the acknowledgement deadline has been exceeded for the message).
        A NACK is any call to ModifyAckDeadline with a 0 deadline. Note that
        client libraries may automatically extend ack_deadlines.
        This field will be honored on a best effort basis.
        If this parameter is 0, a default value of 5 is used.
    """
    expiration_policy: pulumi.Output[dict]
    """
    A policy that specifies the conditions for this subscription's expiration.
    A subscription is considered active as long as any connected subscriber
    is successfully consuming messages from the subscription or is issuing
    operations on the subscription. If expirationPolicy is not set, a default
    policy with ttl of 31 days will be used.  If it is set but ttl is "", the
    resource never expires.  The minimum allowed value for expirationPolicy.ttl
    is 1 day.  Structure is documented below.

      * `ttl` (`str`) - Specifies the "time-to-live" duration for an associated resource. The
        resource expires if it is not active for a period of ttl.
        If ttl is not set, the associated resource never expires.
        A duration in seconds with up to nine fractional digits, terminated by 's'.
        Example - "3.5s".
    """
    labels: pulumi.Output[dict]
    """
    A set of key/value label pairs to assign to this Subscription.
    """
    message_retention_duration: pulumi.Output[str]
    """
    How long to retain unacknowledged messages in the subscription's
    backlog, from the moment a message is published. If
    retainAckedMessages is true, then this also configures the retention
    of acknowledged messages, and thus configures how far back in time a
    subscriptions.seek can be done. Defaults to 7 days. Cannot be more
    than 7 days (`"604800s"`) or less than 10 minutes (`"600s"`).
    A duration in seconds with up to nine fractional digits, terminated
    by 's'. Example: `"600.5s"`.
    """
    name: pulumi.Output[str]
    """
    Name of the subscription.
    """
    path: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    push_config: pulumi.Output[dict]
    """
    If push delivery is used with this subscription, this field is used to
    configure it. An empty pushConfig signifies that the subscriber will
    pull and ack messages using API methods.  Structure is documented below.

      * `attributes` (`dict`) - Endpoint configuration attributes.
        Every endpoint has a set of API supported attributes that can
        be used to control different aspects of the message delivery.
        The currently supported attribute is x-goog-version, which you
        can use to change the format of the pushed message. This
        attribute indicates the version of the data expected by
        the endpoint. This controls the shape of the pushed message
        (i.e., its fields and metadata). The endpoint version is
        based on the version of the Pub/Sub API.
        If not present during the subscriptions.create call,
        it will default to the version of the API used to make
        such call. If not present during a subscriptions.modifyPushConfig
        call, its value will not be changed. subscriptions.get
        calls will always return a valid version, even if the
        subscription was created without this attribute.
        The possible values for this attribute are:
        - v1beta1: uses the push format defined in the v1beta1 Pub/Sub API.
        - v1 or v1beta2: uses the push format defined in the v1 Pub/Sub API.
      * `oidcToken` (`dict`) - If specified, Pub/Sub will generate and attach an OIDC JWT token as
        an Authorization header in the HTTP request for every pushed message.  Structure is documented below.
        * `audience` (`str`) - Audience to be used when generating OIDC token. The audience claim
          identifies the recipients that the JWT is intended for. The audience
          value is a single case-sensitive string. Having multiple values (array)
          for the audience field is not supported. More info about the OIDC JWT
          token audience here: https://tools.ietf.org/html/rfc7519#section-4.1.3
          Note: if not specified, the Push endpoint URL will be used.
        * `service_account_email` (`str`) - Service account email to be used for generating the OIDC token.
          The caller (for subscriptions.create, subscriptions.patch, and
          subscriptions.modifyPushConfig RPCs) must have the
          iam.serviceAccounts.actAs permission for the service account.

      * `pushEndpoint` (`str`) - A URL locating the endpoint to which messages should be pushed.
        For example, a Webhook endpoint might use
        "https://example.com/push".
    """
    retain_acked_messages: pulumi.Output[bool]
    """
    Indicates whether to retain acknowledged messages. If `true`, then
    messages are not expunged from the subscription's backlog, even if
    they are acknowledged, until they fall out of the
    messageRetentionDuration window.
    """
    topic: pulumi.Output[str]
    """
    A reference to a Topic resource.
    """
    def __init__(__self__, resource_name, opts=None, ack_deadline_seconds=None, dead_letter_policy=None, expiration_policy=None, labels=None, message_retention_duration=None, name=None, project=None, push_config=None, retain_acked_messages=None, topic=None, __props__=None, __name__=None, __opts__=None):
        """
        A named resource representing the stream of messages from a single,
        specific topic, to be delivered to the subscribing application.


        To get more information about Subscription, see:

        * [API documentation](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.subscriptions)
        * How-to Guides
            * [Managing Subscriptions](https://cloud.google.com/pubsub/docs/admin#managing_subscriptions)

        ## Example Usage - Pubsub Subscription Push
        {{% example %}}


        ```python
        import pulumi
        import pulumi_gcp as gcp

        example_topic = gcp.pubsub.Topic("exampleTopic")
        example_subscription = gcp.pubsub.Subscription("exampleSubscription",
            topic=example_topic.name,
            ack_deadline_seconds=20,
            labels={
                "foo": "bar",
            },
            push_config={
                "pushEndpoint": "https://example.com/push",
                "attributes": {
                    "x-goog-version": "v1",
                },
            })
        ```
        {{% /example %}}
        ## Example Usage - Pubsub Subscription Pull
        {{% example %}}


        ```python
        import pulumi
        import pulumi_gcp as gcp

        example_topic = gcp.pubsub.Topic("exampleTopic")
        example_subscription = gcp.pubsub.Subscription("exampleSubscription",
            topic=example_topic.name,
            labels={
                "foo": "bar",
            },
            message_retention_duration="1200s",
            retain_acked_messages=True,
            ack_deadline_seconds=20,
            expiration_policy={
                "ttl": "300000.5s",
            })
        ```
        {{% /example %}}
        ## Example Usage - Pubsub Subscription Different Project
        {{% example %}}


        ```python
        import pulumi
        import pulumi_gcp as gcp

        example_topic = gcp.pubsub.Topic("exampleTopic", project="topic-project")
        example_subscription = gcp.pubsub.Subscription("exampleSubscription",
            project="subscription-project",
            topic=example_topic.name)
        ```
        {{% /example %}}
        ## Example Usage - Pubsub Subscription Dead Letter
        {{% example %}}


        ```python
        import pulumi
        import pulumi_gcp as gcp

        example_topic = gcp.pubsub.Topic("exampleTopic")
        example_dead_letter = gcp.pubsub.Topic("exampleDeadLetter")
        example_subscription = gcp.pubsub.Subscription("exampleSubscription",
            topic=example_topic.name,
            dead_letter_policy={
                "deadLetterTopic": example_dead_letter.id,
                "maxDeliveryAttempts": 10,
            })
        ```

        {{% /example %}}

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] ack_deadline_seconds: This value is the maximum time after a subscriber receives a message
               before the subscriber should acknowledge the message. After message
               delivery but before the ack deadline expires and before the message is
               acknowledged, it is an outstanding message and will not be delivered
               again during that time (on a best-effort basis).
               For pull subscriptions, this value is used as the initial value for
               the ack deadline. To override this value for a given message, call
               subscriptions.modifyAckDeadline with the corresponding ackId if using
               pull. The minimum custom deadline you can specify is 10 seconds. The
               maximum custom deadline you can specify is 600 seconds (10 minutes).
               If this parameter is 0, a default value of 10 seconds is used.
               For push delivery, this value is also used to set the request timeout
               for the call to the push endpoint.
               If the subscriber never acknowledges the message, the Pub/Sub system
               will eventually redeliver the message.
        :param pulumi.Input[dict] dead_letter_policy: A policy that specifies the conditions for dead lettering messages in
               this subscription. If dead_letter_policy is not set, dead lettering
               is disabled.
               The Cloud Pub/Sub service account associated with this subscriptions's
               parent project (i.e.,
               service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
               permission to Acknowledge() messages on this subscription.  Structure is documented below.
        :param pulumi.Input[dict] expiration_policy: A policy that specifies the conditions for this subscription's expiration.
               A subscription is considered active as long as any connected subscriber
               is successfully consuming messages from the subscription or is issuing
               operations on the subscription. If expirationPolicy is not set, a default
               policy with ttl of 31 days will be used.  If it is set but ttl is "", the
               resource never expires.  The minimum allowed value for expirationPolicy.ttl
               is 1 day.  Structure is documented below.
        :param pulumi.Input[dict] labels: A set of key/value label pairs to assign to this Subscription.
        :param pulumi.Input[str] message_retention_duration: How long to retain unacknowledged messages in the subscription's
               backlog, from the moment a message is published. If
               retainAckedMessages is true, then this also configures the retention
               of acknowledged messages, and thus configures how far back in time a
               subscriptions.seek can be done. Defaults to 7 days. Cannot be more
               than 7 days (`"604800s"`) or less than 10 minutes (`"600s"`).
               A duration in seconds with up to nine fractional digits, terminated
               by 's'. Example: `"600.5s"`.
        :param pulumi.Input[str] name: Name of the subscription.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] push_config: If push delivery is used with this subscription, this field is used to
               configure it. An empty pushConfig signifies that the subscriber will
               pull and ack messages using API methods.  Structure is documented below.
        :param pulumi.Input[bool] retain_acked_messages: Indicates whether to retain acknowledged messages. If `true`, then
               messages are not expunged from the subscription's backlog, even if
               they are acknowledged, until they fall out of the
               messageRetentionDuration window.
        :param pulumi.Input[str] topic: A reference to a Topic resource.

        The **dead_letter_policy** object supports the following:

          * `deadLetterTopic` (`pulumi.Input[str]`) - The name of the topic to which dead letter messages should be published.
            Format is `projects/{project}/topics/{topic}`.
            The Cloud Pub/Sub service\naccount associated with the enclosing subscription's
            parent project (i.e.,
            service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
            permission to Publish() to this topic.
            The operation will fail if the topic does not exist.
            Users should ensure that there is a subscription attached to this topic
            since messages published to a topic with no subscriptions are lost.
          * `maxDeliveryAttempts` (`pulumi.Input[float]`) - The maximum number of delivery attempts for any message. The value must be
            between 5 and 100.
            The number of delivery attempts is defined as 1 + (the sum of number of
            NACKs and number of times the acknowledgement deadline has been exceeded for the message).
            A NACK is any call to ModifyAckDeadline with a 0 deadline. Note that
            client libraries may automatically extend ack_deadlines.
            This field will be honored on a best effort basis.
            If this parameter is 0, a default value of 5 is used.

        The **expiration_policy** object supports the following:

          * `ttl` (`pulumi.Input[str]`) - Specifies the "time-to-live" duration for an associated resource. The
            resource expires if it is not active for a period of ttl.
            If ttl is not set, the associated resource never expires.
            A duration in seconds with up to nine fractional digits, terminated by 's'.
            Example - "3.5s".

        The **push_config** object supports the following:

          * `attributes` (`pulumi.Input[dict]`) - Endpoint configuration attributes.
            Every endpoint has a set of API supported attributes that can
            be used to control different aspects of the message delivery.
            The currently supported attribute is x-goog-version, which you
            can use to change the format of the pushed message. This
            attribute indicates the version of the data expected by
            the endpoint. This controls the shape of the pushed message
            (i.e., its fields and metadata). The endpoint version is
            based on the version of the Pub/Sub API.
            If not present during the subscriptions.create call,
            it will default to the version of the API used to make
            such call. If not present during a subscriptions.modifyPushConfig
            call, its value will not be changed. subscriptions.get
            calls will always return a valid version, even if the
            subscription was created without this attribute.
            The possible values for this attribute are:
            - v1beta1: uses the push format defined in the v1beta1 Pub/Sub API.
            - v1 or v1beta2: uses the push format defined in the v1 Pub/Sub API.
          * `oidcToken` (`pulumi.Input[dict]`) - If specified, Pub/Sub will generate and attach an OIDC JWT token as
            an Authorization header in the HTTP request for every pushed message.  Structure is documented below.
            * `audience` (`pulumi.Input[str]`) - Audience to be used when generating OIDC token. The audience claim
              identifies the recipients that the JWT is intended for. The audience
              value is a single case-sensitive string. Having multiple values (array)
              for the audience field is not supported. More info about the OIDC JWT
              token audience here: https://tools.ietf.org/html/rfc7519#section-4.1.3
              Note: if not specified, the Push endpoint URL will be used.
            * `service_account_email` (`pulumi.Input[str]`) - Service account email to be used for generating the OIDC token.
              The caller (for subscriptions.create, subscriptions.patch, and
              subscriptions.modifyPushConfig RPCs) must have the
              iam.serviceAccounts.actAs permission for the service account.

          * `pushEndpoint` (`pulumi.Input[str]`) - A URL locating the endpoint to which messages should be pushed.
            For example, a Webhook endpoint might use
            "https://example.com/push".
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

            __props__['ack_deadline_seconds'] = ack_deadline_seconds
            __props__['dead_letter_policy'] = dead_letter_policy
            __props__['expiration_policy'] = expiration_policy
            __props__['labels'] = labels
            __props__['message_retention_duration'] = message_retention_duration
            __props__['name'] = name
            __props__['project'] = project
            __props__['push_config'] = push_config
            __props__['retain_acked_messages'] = retain_acked_messages
            if topic is None:
                raise TypeError("Missing required property 'topic'")
            __props__['topic'] = topic
            __props__['path'] = None
        super(Subscription, __self__).__init__(
            'gcp:pubsub/subscription:Subscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, ack_deadline_seconds=None, dead_letter_policy=None, expiration_policy=None, labels=None, message_retention_duration=None, name=None, path=None, project=None, push_config=None, retain_acked_messages=None, topic=None):
        """
        Get an existing Subscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] ack_deadline_seconds: This value is the maximum time after a subscriber receives a message
               before the subscriber should acknowledge the message. After message
               delivery but before the ack deadline expires and before the message is
               acknowledged, it is an outstanding message and will not be delivered
               again during that time (on a best-effort basis).
               For pull subscriptions, this value is used as the initial value for
               the ack deadline. To override this value for a given message, call
               subscriptions.modifyAckDeadline with the corresponding ackId if using
               pull. The minimum custom deadline you can specify is 10 seconds. The
               maximum custom deadline you can specify is 600 seconds (10 minutes).
               If this parameter is 0, a default value of 10 seconds is used.
               For push delivery, this value is also used to set the request timeout
               for the call to the push endpoint.
               If the subscriber never acknowledges the message, the Pub/Sub system
               will eventually redeliver the message.
        :param pulumi.Input[dict] dead_letter_policy: A policy that specifies the conditions for dead lettering messages in
               this subscription. If dead_letter_policy is not set, dead lettering
               is disabled.
               The Cloud Pub/Sub service account associated with this subscriptions's
               parent project (i.e.,
               service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
               permission to Acknowledge() messages on this subscription.  Structure is documented below.
        :param pulumi.Input[dict] expiration_policy: A policy that specifies the conditions for this subscription's expiration.
               A subscription is considered active as long as any connected subscriber
               is successfully consuming messages from the subscription or is issuing
               operations on the subscription. If expirationPolicy is not set, a default
               policy with ttl of 31 days will be used.  If it is set but ttl is "", the
               resource never expires.  The minimum allowed value for expirationPolicy.ttl
               is 1 day.  Structure is documented below.
        :param pulumi.Input[dict] labels: A set of key/value label pairs to assign to this Subscription.
        :param pulumi.Input[str] message_retention_duration: How long to retain unacknowledged messages in the subscription's
               backlog, from the moment a message is published. If
               retainAckedMessages is true, then this also configures the retention
               of acknowledged messages, and thus configures how far back in time a
               subscriptions.seek can be done. Defaults to 7 days. Cannot be more
               than 7 days (`"604800s"`) or less than 10 minutes (`"600s"`).
               A duration in seconds with up to nine fractional digits, terminated
               by 's'. Example: `"600.5s"`.
        :param pulumi.Input[str] name: Name of the subscription.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] push_config: If push delivery is used with this subscription, this field is used to
               configure it. An empty pushConfig signifies that the subscriber will
               pull and ack messages using API methods.  Structure is documented below.
        :param pulumi.Input[bool] retain_acked_messages: Indicates whether to retain acknowledged messages. If `true`, then
               messages are not expunged from the subscription's backlog, even if
               they are acknowledged, until they fall out of the
               messageRetentionDuration window.
        :param pulumi.Input[str] topic: A reference to a Topic resource.

        The **dead_letter_policy** object supports the following:

          * `deadLetterTopic` (`pulumi.Input[str]`) - The name of the topic to which dead letter messages should be published.
            Format is `projects/{project}/topics/{topic}`.
            The Cloud Pub/Sub service\naccount associated with the enclosing subscription's
            parent project (i.e.,
            service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
            permission to Publish() to this topic.
            The operation will fail if the topic does not exist.
            Users should ensure that there is a subscription attached to this topic
            since messages published to a topic with no subscriptions are lost.
          * `maxDeliveryAttempts` (`pulumi.Input[float]`) - The maximum number of delivery attempts for any message. The value must be
            between 5 and 100.
            The number of delivery attempts is defined as 1 + (the sum of number of
            NACKs and number of times the acknowledgement deadline has been exceeded for the message).
            A NACK is any call to ModifyAckDeadline with a 0 deadline. Note that
            client libraries may automatically extend ack_deadlines.
            This field will be honored on a best effort basis.
            If this parameter is 0, a default value of 5 is used.

        The **expiration_policy** object supports the following:

          * `ttl` (`pulumi.Input[str]`) - Specifies the "time-to-live" duration for an associated resource. The
            resource expires if it is not active for a period of ttl.
            If ttl is not set, the associated resource never expires.
            A duration in seconds with up to nine fractional digits, terminated by 's'.
            Example - "3.5s".

        The **push_config** object supports the following:

          * `attributes` (`pulumi.Input[dict]`) - Endpoint configuration attributes.
            Every endpoint has a set of API supported attributes that can
            be used to control different aspects of the message delivery.
            The currently supported attribute is x-goog-version, which you
            can use to change the format of the pushed message. This
            attribute indicates the version of the data expected by
            the endpoint. This controls the shape of the pushed message
            (i.e., its fields and metadata). The endpoint version is
            based on the version of the Pub/Sub API.
            If not present during the subscriptions.create call,
            it will default to the version of the API used to make
            such call. If not present during a subscriptions.modifyPushConfig
            call, its value will not be changed. subscriptions.get
            calls will always return a valid version, even if the
            subscription was created without this attribute.
            The possible values for this attribute are:
            - v1beta1: uses the push format defined in the v1beta1 Pub/Sub API.
            - v1 or v1beta2: uses the push format defined in the v1 Pub/Sub API.
          * `oidcToken` (`pulumi.Input[dict]`) - If specified, Pub/Sub will generate and attach an OIDC JWT token as
            an Authorization header in the HTTP request for every pushed message.  Structure is documented below.
            * `audience` (`pulumi.Input[str]`) - Audience to be used when generating OIDC token. The audience claim
              identifies the recipients that the JWT is intended for. The audience
              value is a single case-sensitive string. Having multiple values (array)
              for the audience field is not supported. More info about the OIDC JWT
              token audience here: https://tools.ietf.org/html/rfc7519#section-4.1.3
              Note: if not specified, the Push endpoint URL will be used.
            * `service_account_email` (`pulumi.Input[str]`) - Service account email to be used for generating the OIDC token.
              The caller (for subscriptions.create, subscriptions.patch, and
              subscriptions.modifyPushConfig RPCs) must have the
              iam.serviceAccounts.actAs permission for the service account.

          * `pushEndpoint` (`pulumi.Input[str]`) - A URL locating the endpoint to which messages should be pushed.
            For example, a Webhook endpoint might use
            "https://example.com/push".
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["ack_deadline_seconds"] = ack_deadline_seconds
        __props__["dead_letter_policy"] = dead_letter_policy
        __props__["expiration_policy"] = expiration_policy
        __props__["labels"] = labels
        __props__["message_retention_duration"] = message_retention_duration
        __props__["name"] = name
        __props__["path"] = path
        __props__["project"] = project
        __props__["push_config"] = push_config
        __props__["retain_acked_messages"] = retain_acked_messages
        __props__["topic"] = topic
        return Subscription(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

