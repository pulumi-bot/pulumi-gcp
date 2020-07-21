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
class InstanceNetworkConfigArgs:
    ip_allocation: pulumi.Input[str] = pulumi.input_property("ipAllocation")
    """
    The IP range in CIDR notation to use for the managed Data Fusion instance
    nodes. This range must not overlap with any other ranges used in the Data Fusion instance network.
    """
    network: pulumi.Input[str] = pulumi.input_property("network")
    """
    Name of the network in the project with which the tenant project
    will be peered for executing pipelines. In case of shared VPC where the network resides in another host
    project the network should specified in the form of projects/{host-project-id}/global/networks/{network}
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, ip_allocation: pulumi.Input[str], network: pulumi.Input[str]) -> None:
        """
        :param pulumi.Input[str] ip_allocation: The IP range in CIDR notation to use for the managed Data Fusion instance
               nodes. This range must not overlap with any other ranges used in the Data Fusion instance network.
        :param pulumi.Input[str] network: Name of the network in the project with which the tenant project
               will be peered for executing pipelines. In case of shared VPC where the network resides in another host
               project the network should specified in the form of projects/{host-project-id}/global/networks/{network}
        """
        __self__.ip_allocation = ip_allocation
        __self__.network = network

