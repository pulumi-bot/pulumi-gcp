# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'DatabaseIAMBindingConditionArgs',
    'DatabaseIAMMemberConditionArgs',
    'InstanceIAMBindingConditionArgs',
    'InstanceIAMMemberConditionArgs',
]

@pulumi.input_type
class DatabaseIAMBindingConditionArgs:
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
class DatabaseIAMMemberConditionArgs:
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
class InstanceIAMBindingConditionArgs:
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
class InstanceIAMMemberConditionArgs:
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


