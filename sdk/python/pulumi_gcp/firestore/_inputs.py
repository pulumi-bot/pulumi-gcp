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

@pulumi.input_type
class IndexFieldArgs:
    array_config: Optional[pulumi.Input[str]] = pulumi.input_property("arrayConfig")
    """
    Indicates that this field supports operations on arrayValues. Only one of `order` and `arrayConfig` can
    be specified.
    """
    field_path: Optional[pulumi.Input[str]] = pulumi.input_property("fieldPath")
    """
    Name of the field.
    """
    order: Optional[pulumi.Input[str]] = pulumi.input_property("order")
    """
    Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
    Only one of `order` and `arrayConfig` can be specified.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, array_config: Optional[pulumi.Input[str]] = None, field_path: Optional[pulumi.Input[str]] = None, order: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] array_config: Indicates that this field supports operations on arrayValues. Only one of `order` and `arrayConfig` can
               be specified.
        :param pulumi.Input[str] field_path: Name of the field.
        :param pulumi.Input[str] order: Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
               Only one of `order` and `arrayConfig` can be specified.
        """
        __self__.array_config = array_config
        __self__.field_path = field_path
        __self__.order = order

