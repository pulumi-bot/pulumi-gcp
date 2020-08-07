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
    'ServiceApi',
    'ServiceApiMethod',
    'ServiceEndpoint',
    'ServiceIamBindingCondition',
    'ServiceIamMemberCondition',
]

@pulumi.output_type
class ServiceApi(dict):
    @property
    @pulumi.getter
    def methods(self) -> Optional[List['outputs.ServiceApiMethod']]:
        ...

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def syntax(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceApiMethod(dict):
    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter(name="requestType")
    def request_type(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter(name="responseType")
    def response_type(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def syntax(self) -> Optional[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceEndpoint(dict):
    @property
    @pulumi.getter
    def address(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceIamBindingCondition(dict):
    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def expression(self) -> str:
        ...

    @property
    @pulumi.getter
    def title(self) -> str:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceIamMemberCondition(dict):
    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def expression(self) -> str:
        ...

    @property
    @pulumi.getter
    def title(self) -> str:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


