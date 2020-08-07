# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'RepositoryIamBindingConditionArgs',
    'RepositoryIamMemberConditionArgs',
    'RepositoryPubsubConfigArgs',
]

@pulumi.input_type
class RepositoryIamBindingConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        ...

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        ...

    @title.setter
    def title(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        ...

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class RepositoryIamMemberConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        ...

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        ...

    @title.setter
    def title(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        ...

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class RepositoryPubsubConfigArgs:
    def __init__(__self__, *,
                 message_format: pulumi.Input[str],
                 topic: pulumi.Input[str],
                 service_account_email: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] message_format: The format of the Cloud Pub/Sub messages.
               - PROTOBUF: The message payload is a serialized protocol buffer of SourceRepoEvent.
               - JSON: The message payload is a JSON string of SourceRepoEvent.
        :param pulumi.Input[str] topic: The identifier for this object. Format specified above.
        :param pulumi.Input[str] service_account_email: Email address of the service account used for publishing Cloud Pub/Sub messages.
               This service account needs to be in the same project as the PubsubConfig. When added,
               the caller needs to have iam.serviceAccounts.actAs permission on this service account.
               If unspecified, it defaults to the compute engine default service account.
        """
        pulumi.set(__self__, "messageFormat", message_format)
        pulumi.set(__self__, "topic", topic)
        pulumi.set(__self__, "serviceAccountEmail", service_account_email)

    @property
    @pulumi.getter(name="messageFormat")
    def message_format(self) -> pulumi.Input[str]:
        """
        The format of the Cloud Pub/Sub messages.
        - PROTOBUF: The message payload is a serialized protocol buffer of SourceRepoEvent.
        - JSON: The message payload is a JSON string of SourceRepoEvent.
        """
        ...

    @message_format.setter
    def message_format(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def topic(self) -> pulumi.Input[str]:
        """
        The identifier for this object. Format specified above.
        """
        ...

    @topic.setter
    def topic(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> Optional[pulumi.Input[str]]:
        """
        Email address of the service account used for publishing Cloud Pub/Sub messages.
        This service account needs to be in the same project as the PubsubConfig. When added,
        the caller needs to have iam.serviceAccounts.actAs permission on this service account.
        If unspecified, it defaults to the compute engine default service account.
        """
        ...

    @service_account_email.setter
    def service_account_email(self, value: Optional[pulumi.Input[str]]):
        ...


