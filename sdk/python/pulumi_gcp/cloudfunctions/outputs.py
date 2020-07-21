# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from ._inputs import *
from . import outputs

@pulumi.output_type
class FunctionEventTrigger(dict):
    event_type: str = pulumi.output_property("eventType")
    """
    The type of event to observe. For example: `"google.storage.object.finalize"`.
    See the documentation on [calling Cloud Functions](https://cloud.google.com/functions/docs/calling/) for a
    full reference of accepted triggers.
    """
    failure_policy: Optional['outputs.FunctionEventTriggerFailurePolicy'] = pulumi.output_property("failurePolicy")
    """
    Specifies policy for failed executions. Structure is documented below.
    """
    resource: str = pulumi.output_property("resource")
    """
    Required. The name or partial URI of the resource from
    which to observe events. For example, `"myBucket"` or `"projects/my-project/topics/my-topic"`
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class FunctionEventTriggerFailurePolicy(dict):
    retry: bool = pulumi.output_property("retry")
    """
    Whether the function should be retried on failure. Defaults to `false`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class FunctionIamBindingCondition(dict):
    description: Optional[str] = pulumi.output_property("description")
    expression: str = pulumi.output_property("expression")
    title: str = pulumi.output_property("title")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class FunctionIamMemberCondition(dict):
    description: Optional[str] = pulumi.output_property("description")
    expression: str = pulumi.output_property("expression")
    title: str = pulumi.output_property("title")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class FunctionSourceRepository(dict):
    deployed_url: Optional[str] = pulumi.output_property("deployedUrl")
    url: str = pulumi.output_property("url")
    """
    The URL pointing to the hosted repository where the function is defined. There are supported Cloud Source Repository URLs in the following formats:
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetFunctionEventTrigger(dict):
    event_type: str = pulumi.output_property("eventType")
    """
    The type of event to observe. For example: `"google.storage.object.finalize"`.
    See the documentation on [calling Cloud Functions](https://cloud.google.com/functions/docs/calling/)
    for a full reference of accepted triggers.
    """
    failure_policies: List['outputs.GetFunctionEventTriggerFailurePolicy'] = pulumi.output_property("failurePolicies")
    """
    Policy for failed executions. Structure is documented below.
    """
    resource: str = pulumi.output_property("resource")
    """
    The name of the resource whose events are being observed, for example, `"myBucket"`
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetFunctionEventTriggerFailurePolicy(dict):
    retry: bool = pulumi.output_property("retry")
    """
    Whether the function should be retried on failure.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetFunctionSourceRepository(dict):
    deployed_url: str = pulumi.output_property("deployedUrl")
    url: str = pulumi.output_property("url")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


