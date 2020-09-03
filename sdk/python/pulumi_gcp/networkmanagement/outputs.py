# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'ConnectivityTestDestination',
    'ConnectivityTestSource',
]

@pulumi.output_type
class ConnectivityTestDestination(dict):
    def __init__(__self__, *,
                 instance: Optional[str] = None,
                 ip_address: Optional[str] = None,
                 network: Optional[str] = None,
                 port: Optional[float] = None,
                 project_id: Optional[str] = None):
        """
        :param str instance: A Compute Engine instance URI.
        :param str ip_address: The IP address of the endpoint, which can be an external or
               internal IP. An IPv6 address is only allowed when the test's
               destination is a global load balancer VIP.
        :param str network: A Compute Engine network URI.
        :param float port: The IP protocol port of the endpoint. Only applicable when
               protocol is TCP or UDP.
        :param str project_id: Project ID where the endpoint is located. The Project ID can be
               derived from the URI if you provide a VM instance or network URI.
               The following are two cases where you must provide the project ID:
               1. Only the IP address is specified, and the IP address is within
               a GCP project. 2. When you are using Shared VPC and the IP address
               that you provide is from the service project. In this case, the
               network that the IP address resides in is defined in the host
               project.
        """
        if instance is not None:
            pulumi.set(__self__, "instance", instance)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter
    def instance(self) -> Optional[str]:
        """
        A Compute Engine instance URI.
        """
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[str]:
        """
        The IP address of the endpoint, which can be an external or
        internal IP. An IPv6 address is only allowed when the test's
        destination is a global load balancer VIP.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter
    def network(self) -> Optional[str]:
        """
        A Compute Engine network URI.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter
    def port(self) -> Optional[float]:
        """
        The IP protocol port of the endpoint. Only applicable when
        protocol is TCP or UDP.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
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
        return pulumi.get(self, "project_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ConnectivityTestSource(dict):
    def __init__(__self__, *,
                 instance: Optional[str] = None,
                 ip_address: Optional[str] = None,
                 network: Optional[str] = None,
                 network_type: Optional[str] = None,
                 port: Optional[float] = None,
                 project_id: Optional[str] = None):
        """
        :param str instance: A Compute Engine instance URI.
        :param str ip_address: The IP address of the endpoint, which can be an external or
               internal IP. An IPv6 address is only allowed when the test's
               destination is a global load balancer VIP.
        :param str network: A Compute Engine network URI.
        :param str network_type: Type of the network where the endpoint is located.
               Possible values are `GCP_NETWORK` and `NON_GCP_NETWORK`.
        :param float port: The IP protocol port of the endpoint. Only applicable when
               protocol is TCP or UDP.
        :param str project_id: Project ID where the endpoint is located. The Project ID can be
               derived from the URI if you provide a VM instance or network URI.
               The following are two cases where you must provide the project ID:
               1. Only the IP address is specified, and the IP address is within
               a GCP project. 2. When you are using Shared VPC and the IP address
               that you provide is from the service project. In this case, the
               network that the IP address resides in is defined in the host
               project.
        """
        if instance is not None:
            pulumi.set(__self__, "instance", instance)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if network_type is not None:
            pulumi.set(__self__, "network_type", network_type)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter
    def instance(self) -> Optional[str]:
        """
        A Compute Engine instance URI.
        """
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[str]:
        """
        The IP address of the endpoint, which can be an external or
        internal IP. An IPv6 address is only allowed when the test's
        destination is a global load balancer VIP.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter
    def network(self) -> Optional[str]:
        """
        A Compute Engine network URI.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[str]:
        """
        Type of the network where the endpoint is located.
        Possible values are `GCP_NETWORK` and `NON_GCP_NETWORK`.
        """
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter
    def port(self) -> Optional[float]:
        """
        The IP protocol port of the endpoint. Only applicable when
        protocol is TCP or UDP.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
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
        return pulumi.get(self, "project_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


