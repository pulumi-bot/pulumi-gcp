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
class ConnectivityTestDestinationArgs:
    instance: Optional[pulumi.Input[str]] = pulumi.input_property("instance")
    """
    A Compute Engine instance URI.
    """
    ip_address: Optional[pulumi.Input[str]] = pulumi.input_property("ipAddress")
    """
    The IP address of the endpoint, which can be an external or
    internal IP. An IPv6 address is only allowed when the test's
    destination is a global load balancer VIP.
    """
    network: Optional[pulumi.Input[str]] = pulumi.input_property("network")
    """
    A Compute Engine network URI.
    """
    port: Optional[pulumi.Input[float]] = pulumi.input_property("port")
    """
    The IP protocol port of the endpoint. Only applicable when
    protocol is TCP or UDP.
    """
    project_id: Optional[pulumi.Input[str]] = pulumi.input_property("projectId")
    """
    Project ID where the endpoint is located. The Project ID can be
    derived from the URI if you provide a VM instance or network URI.
    The following are two cases where you must provide the project ID:
    1. Only the IP address is specified, and the IP address is within
    a GCP project. 2. When you are using Shared VPC and the IP address
    that you provide is from the service project. In this case, the
    network that the IP address resides in is defined in the host
    project.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, instance: Optional[pulumi.Input[str]] = None, ip_address: Optional[pulumi.Input[str]] = None, network: Optional[pulumi.Input[str]] = None, port: Optional[pulumi.Input[float]] = None, project_id: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] instance: A Compute Engine instance URI.
        :param pulumi.Input[str] ip_address: The IP address of the endpoint, which can be an external or
               internal IP. An IPv6 address is only allowed when the test's
               destination is a global load balancer VIP.
        :param pulumi.Input[str] network: A Compute Engine network URI.
        :param pulumi.Input[float] port: The IP protocol port of the endpoint. Only applicable when
               protocol is TCP or UDP.
        :param pulumi.Input[str] project_id: Project ID where the endpoint is located. The Project ID can be
               derived from the URI if you provide a VM instance or network URI.
               The following are two cases where you must provide the project ID:
               1. Only the IP address is specified, and the IP address is within
               a GCP project. 2. When you are using Shared VPC and the IP address
               that you provide is from the service project. In this case, the
               network that the IP address resides in is defined in the host
               project.
        """
        __self__.instance = instance
        __self__.ip_address = ip_address
        __self__.network = network
        __self__.port = port
        __self__.project_id = project_id

@pulumi.input_type
class ConnectivityTestSourceArgs:
    instance: Optional[pulumi.Input[str]] = pulumi.input_property("instance")
    """
    A Compute Engine instance URI.
    """
    ip_address: Optional[pulumi.Input[str]] = pulumi.input_property("ipAddress")
    """
    The IP address of the endpoint, which can be an external or
    internal IP. An IPv6 address is only allowed when the test's
    destination is a global load balancer VIP.
    """
    network: Optional[pulumi.Input[str]] = pulumi.input_property("network")
    """
    A Compute Engine network URI.
    """
    network_type: Optional[pulumi.Input[str]] = pulumi.input_property("networkType")
    """
    Type of the network where the endpoint is located.
    """
    port: Optional[pulumi.Input[float]] = pulumi.input_property("port")
    """
    The IP protocol port of the endpoint. Only applicable when
    protocol is TCP or UDP.
    """
    project_id: Optional[pulumi.Input[str]] = pulumi.input_property("projectId")
    """
    Project ID where the endpoint is located. The Project ID can be
    derived from the URI if you provide a VM instance or network URI.
    The following are two cases where you must provide the project ID:
    1. Only the IP address is specified, and the IP address is within
    a GCP project. 2. When you are using Shared VPC and the IP address
    that you provide is from the service project. In this case, the
    network that the IP address resides in is defined in the host
    project.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, instance: Optional[pulumi.Input[str]] = None, ip_address: Optional[pulumi.Input[str]] = None, network: Optional[pulumi.Input[str]] = None, network_type: Optional[pulumi.Input[str]] = None, port: Optional[pulumi.Input[float]] = None, project_id: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] instance: A Compute Engine instance URI.
        :param pulumi.Input[str] ip_address: The IP address of the endpoint, which can be an external or
               internal IP. An IPv6 address is only allowed when the test's
               destination is a global load balancer VIP.
        :param pulumi.Input[str] network: A Compute Engine network URI.
        :param pulumi.Input[str] network_type: Type of the network where the endpoint is located.
        :param pulumi.Input[float] port: The IP protocol port of the endpoint. Only applicable when
               protocol is TCP or UDP.
        :param pulumi.Input[str] project_id: Project ID where the endpoint is located. The Project ID can be
               derived from the URI if you provide a VM instance or network URI.
               The following are two cases where you must provide the project ID:
               1. Only the IP address is specified, and the IP address is within
               a GCP project. 2. When you are using Shared VPC and the IP address
               that you provide is from the service project. In this case, the
               network that the IP address resides in is defined in the host
               project.
        """
        __self__.instance = instance
        __self__.ip_address = ip_address
        __self__.network = network
        __self__.network_type = network_type
        __self__.port = port
        __self__.project_id = project_id

