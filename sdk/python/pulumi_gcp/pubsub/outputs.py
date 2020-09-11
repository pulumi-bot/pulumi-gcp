# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'SubscriptionDeadLetterPolicy',
    'SubscriptionExpirationPolicy',
    'SubscriptionIAMBindingCondition',
    'SubscriptionIAMMemberCondition',
    'SubscriptionPushConfig',
    'SubscriptionPushConfigOidcToken',
    'SubscriptionRetryPolicy',
    'TopicIAMBindingCondition',
    'TopicIAMMemberCondition',
    'TopicMessageStoragePolicy',
]

@pulumi.output_type
class SubscriptionDeadLetterPolicy(dict):
    def __init__(__self__, *,
                 dead_letter_topic: Optional[str] = None,
                 max_delivery_attempts: Optional[int] = None):
        """
        :param str dead_letter_topic: The name of the topic to which dead letter messages should be published.
               Format is `projects/{project}/topics/{topic}`.
               The Cloud Pub/Sub service account associated with the enclosing subscription's
               parent project (i.e.,
               service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
               permission to Publish() to this topic.
               The operation will fail if the topic does not exist.
               Users should ensure that there is a subscription attached to this topic
               since messages published to a topic with no subscriptions are lost.
        :param int max_delivery_attempts: The maximum number of delivery attempts for any message. The value must be
               between 5 and 100.
               The number of delivery attempts is defined as 1 + (the sum of number of
               NACKs and number of times the acknowledgement deadline has been exceeded for the message).
               A NACK is any call to ModifyAckDeadline with a 0 deadline. Note that
               client libraries may automatically extend ack_deadlines.
               This field will be honored on a best effort basis.
               If this parameter is 0, a default value of 5 is used.
        """
        if dead_letter_topic is not None:
            pulumi.set(__self__, "dead_letter_topic", dead_letter_topic)
        if max_delivery_attempts is not None:
            pulumi.set(__self__, "max_delivery_attempts", max_delivery_attempts)

    @property
    @pulumi.getter(name="deadLetterTopic")
    def dead_letter_topic(self) -> Optional[str]:
        """
        The name of the topic to which dead letter messages should be published.
        Format is `projects/{project}/topics/{topic}`.
        The Cloud Pub/Sub service account associated with the enclosing subscription's
        parent project (i.e.,
        service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
        permission to Publish() to this topic.
        The operation will fail if the topic does not exist.
        Users should ensure that there is a subscription attached to this topic
        since messages published to a topic with no subscriptions are lost.
        """
        return pulumi.get(self, "dead_letter_topic")

    @property
    @pulumi.getter(name="maxDeliveryAttempts")
    def max_delivery_attempts(self) -> Optional[int]:
        """
        The maximum number of delivery attempts for any message. The value must be
        between 5 and 100.
        The number of delivery attempts is defined as 1 + (the sum of number of
        NACKs and number of times the acknowledgement deadline has been exceeded for the message).
        A NACK is any call to ModifyAckDeadline with a 0 deadline. Note that
        client libraries may automatically extend ack_deadlines.
        This field will be honored on a best effort basis.
        If this parameter is 0, a default value of 5 is used.
        """
        return pulumi.get(self, "max_delivery_attempts")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SubscriptionExpirationPolicy(dict):
    def __init__(__self__, *,
                 ttl: str):
        """
        :param str ttl: Specifies the "time-to-live" duration for an associated resource. The
               resource expires if it is not active for a period of ttl.
               If ttl is not set, the associated resource never expires.
               A duration in seconds with up to nine fractional digits, terminated by 's'.
               Example - "3.5s".
        """
        pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter
    def ttl(self) -> str:
        """
        Specifies the "time-to-live" duration for an associated resource. The
        resource expires if it is not active for a period of ttl.
        If ttl is not set, the associated resource never expires.
        A duration in seconds with up to nine fractional digits, terminated by 's'.
        Example - "3.5s".
        """
        return pulumi.get(self, "ttl")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SubscriptionIAMBindingCondition(dict):
    def __init__(__self__, *,
                 expression: str,
                 title: str,
                 description: Optional[str] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> str:
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SubscriptionIAMMemberCondition(dict):
    def __init__(__self__, *,
                 expression: str,
                 title: str,
                 description: Optional[str] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> str:
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SubscriptionPushConfig(dict):
    def __init__(__self__, *,
                 push_endpoint: str,
                 attributes: Optional[Mapping[str, str]] = None,
                 oidc_token: Optional['outputs.SubscriptionPushConfigOidcToken'] = None):
        """
        :param str push_endpoint: A URL locating the endpoint to which messages should be pushed.
               For example, a Webhook endpoint might use
               "https://example.com/push".
        :param Mapping[str, str] attributes: Endpoint configuration attributes.
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
        :param 'SubscriptionPushConfigOidcTokenArgs' oidc_token: If specified, Pub/Sub will generate and attach an OIDC JWT token as
               an Authorization header in the HTTP request for every pushed message.
               Structure is documented below.
        """
        pulumi.set(__self__, "push_endpoint", push_endpoint)
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if oidc_token is not None:
            pulumi.set(__self__, "oidc_token", oidc_token)

    @property
    @pulumi.getter(name="pushEndpoint")
    def push_endpoint(self) -> str:
        """
        A URL locating the endpoint to which messages should be pushed.
        For example, a Webhook endpoint might use
        "https://example.com/push".
        """
        return pulumi.get(self, "push_endpoint")

    @property
    @pulumi.getter
    def attributes(self) -> Optional[Mapping[str, str]]:
        """
        Endpoint configuration attributes.
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
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="oidcToken")
    def oidc_token(self) -> Optional['outputs.SubscriptionPushConfigOidcToken']:
        """
        If specified, Pub/Sub will generate and attach an OIDC JWT token as
        an Authorization header in the HTTP request for every pushed message.
        Structure is documented below.
        """
        return pulumi.get(self, "oidc_token")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SubscriptionPushConfigOidcToken(dict):
    def __init__(__self__, *,
                 service_account_email: str,
                 audience: Optional[str] = None):
        """
        :param str service_account_email: Service account email to be used for generating the OIDC token.
               The caller (for subscriptions.create, subscriptions.patch, and
               subscriptions.modifyPushConfig RPCs) must have the
               iam.serviceAccounts.actAs permission for the service account.
        :param str audience: Audience to be used when generating OIDC token. The audience claim
               identifies the recipients that the JWT is intended for. The audience
               value is a single case-sensitive string. Having multiple values (array)
               for the audience field is not supported. More info about the OIDC JWT
               token audience here: https://tools.ietf.org/html/rfc7519#section-4.1.3
               Note: if not specified, the Push endpoint URL will be used.
        """
        pulumi.set(__self__, "service_account_email", service_account_email)
        if audience is not None:
            pulumi.set(__self__, "audience", audience)

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> str:
        """
        Service account email to be used for generating the OIDC token.
        The caller (for subscriptions.create, subscriptions.patch, and
        subscriptions.modifyPushConfig RPCs) must have the
        iam.serviceAccounts.actAs permission for the service account.
        """
        return pulumi.get(self, "service_account_email")

    @property
    @pulumi.getter
    def audience(self) -> Optional[str]:
        """
        Audience to be used when generating OIDC token. The audience claim
        identifies the recipients that the JWT is intended for. The audience
        value is a single case-sensitive string. Having multiple values (array)
        for the audience field is not supported. More info about the OIDC JWT
        token audience here: https://tools.ietf.org/html/rfc7519#section-4.1.3
        Note: if not specified, the Push endpoint URL will be used.
        """
        return pulumi.get(self, "audience")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SubscriptionRetryPolicy(dict):
    def __init__(__self__, *,
                 maximum_backoff: Optional[str] = None,
                 minimum_backoff: Optional[str] = None):
        """
        :param str maximum_backoff: The maximum delay between consecutive deliveries of a given message. Value should be between 0 and 600 seconds. Defaults to 600 seconds.
               A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        :param str minimum_backoff: The minimum delay between consecutive deliveries of a given message. Value should be between 0 and 600 seconds. Defaults to 10 seconds.
               A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        """
        if maximum_backoff is not None:
            pulumi.set(__self__, "maximum_backoff", maximum_backoff)
        if minimum_backoff is not None:
            pulumi.set(__self__, "minimum_backoff", minimum_backoff)

    @property
    @pulumi.getter(name="maximumBackoff")
    def maximum_backoff(self) -> Optional[str]:
        """
        The maximum delay between consecutive deliveries of a given message. Value should be between 0 and 600 seconds. Defaults to 600 seconds.
        A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        """
        return pulumi.get(self, "maximum_backoff")

    @property
    @pulumi.getter(name="minimumBackoff")
    def minimum_backoff(self) -> Optional[str]:
        """
        The minimum delay between consecutive deliveries of a given message. Value should be between 0 and 600 seconds. Defaults to 10 seconds.
        A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        """
        return pulumi.get(self, "minimum_backoff")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TopicIAMBindingCondition(dict):
    def __init__(__self__, *,
                 expression: str,
                 title: str,
                 description: Optional[str] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> str:
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TopicIAMMemberCondition(dict):
    def __init__(__self__, *,
                 expression: str,
                 title: str,
                 description: Optional[str] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> str:
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TopicMessageStoragePolicy(dict):
    def __init__(__self__, *,
                 allowed_persistence_regions: List[str]):
        """
        :param List[str] allowed_persistence_regions: A list of IDs of GCP regions where messages that are published to
               the topic may be persisted in storage. Messages published by
               publishers running in non-allowed GCP regions (or running outside
               of GCP altogether) will be routed for storage in one of the
               allowed regions. An empty list means that no regions are allowed,
               and is not a valid configuration.
        """
        pulumi.set(__self__, "allowed_persistence_regions", allowed_persistence_regions)

    @property
    @pulumi.getter(name="allowedPersistenceRegions")
    def allowed_persistence_regions(self) -> List[str]:
        """
        A list of IDs of GCP regions where messages that are published to
        the topic may be persisted in storage. Messages published by
        publishers running in non-allowed GCP regions (or running outside
        of GCP altogether) will be routed for storage in one of the
        allowed regions. An empty list means that no regions are allowed,
        and is not a valid configuration.
        """
        return pulumi.get(self, "allowed_persistence_regions")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


