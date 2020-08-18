# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetLBIPRangesResult',
    'AwaitableGetLBIPRangesResult',
    'get_lbip_ranges',
]



@pulumi.output_type
class GetLBIPRangesResult:
    """
    A collection of values returned by getLBIPRanges.
    """
    def __init__(__self__, http_ssl_tcp_internals=None, id=None, networks=None):
        if http_ssl_tcp_internals and not isinstance(http_ssl_tcp_internals, list):
            raise TypeError("Expected argument 'http_ssl_tcp_internals' to be a list")
        pulumi.set(__self__, "http_ssl_tcp_internals", http_ssl_tcp_internals)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if networks and not isinstance(networks, list):
            raise TypeError("Expected argument 'networks' to be a list")
        pulumi.set(__self__, "networks", networks)

    @property
    @pulumi.getter(name="httpSslTcpInternals")
    def http_ssl_tcp_internals(self) -> List[str]:
        """
        The IP ranges used for health checks when **HTTP(S), SSL proxy, TCP proxy, and Internal load balancing** is used
        """
        return pulumi.get(self, "http_ssl_tcp_internals")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def networks(self) -> List[str]:
        """
        The IP ranges used for health checks when **Network load balancing** is used
        """
        return pulumi.get(self, "networks")



class AwaitableGetLBIPRangesResult(GetLBIPRangesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLBIPRangesResult(
            http_ssl_tcp_internals=self.http_ssl_tcp_internals,
            id=self.id,
            networks=self.networks)


def get_lbip_ranges(                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLBIPRangesResult:
    """
    Use this data source to access IP ranges in your firewall rules.

    https://cloud.google.com/compute/docs/load-balancing/health-checks#health_check_source_ips_and_firewall_rules
    """
    __args__ = dict()
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getLBIPRanges:getLBIPRanges', __args__, opts=opts, typ=GetLBIPRangesResult).value

    return AwaitableGetLBIPRangesResult(
        http_ssl_tcp_internals=__ret__.http_ssl_tcp_internals,
        id=__ret__.id,
        networks=__ret__.networks)
