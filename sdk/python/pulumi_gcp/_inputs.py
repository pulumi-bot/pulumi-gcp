# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from . import _utilities, _tables
from ._inputs import *
from . import outputs

@pulumi.input_type
class ProviderBatchingArgs:
    enable_batching: Optional[pulumi.Input[bool]] = pulumi.input_property("enableBatching")
    send_after: Optional[pulumi.Input[str]] = pulumi.input_property("sendAfter")

    # pylint: disable=no-self-argument
    def __init__(__self__, *, enable_batching: Optional[pulumi.Input[bool]] = None, send_after: Optional[pulumi.Input[str]] = None) -> None:
        __self__.enable_batching = enable_batching
        __self__.send_after = send_after

