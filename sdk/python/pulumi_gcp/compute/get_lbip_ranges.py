# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetLBIPRangesResult:
    """
    A collection of values returned by getLBIPRanges.
    """
    def __init__(__self__, http_ssl_tcp_internals=None, id=None, networks=None):
        if http_ssl_tcp_internals and not isinstance(http_ssl_tcp_internals, list):
            raise TypeError("Expected argument 'http_ssl_tcp_internals' to be a list")
        __self__.http_ssl_tcp_internals = http_ssl_tcp_internals
        """
        The IP ranges used for health checks when **HTTP(S), SSL proxy, TCP proxy, and Internal load balancing** is used
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if networks and not isinstance(networks, list):
            raise TypeError("Expected argument 'networks' to be a list")
        __self__.networks = networks
        """
        The IP ranges used for health checks when **Network load balancing** is used
        """
class AwaitableGetLBIPRangesResult(GetLBIPRangesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLBIPRangesResult(
            http_ssl_tcp_internals=self.http_ssl_tcp_internals,
            id=self.id,
            networks=self.networks)

def get_lbip_ranges(opts=None):
    """
    Use this data source to access IP ranges in your firewall rules.

    https://cloud.google.com/compute/docs/load-balancing/health-checks#health_check_source_ips_and_firewall_rules

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    ranges = gcp.compute.get_lbip_ranges()
    lb = gcp.compute.Firewall("lb",
        network=google_compute_network["main"]["name"],
        allow=[{
            "protocol": "tcp",
            "ports": ["80"],
        }],
        source_ranges=ranges.networks,
        target_tags=["InstanceBehindLoadBalancer"])
    ```
    """
    __args__ = dict()


    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getLBIPRanges:getLBIPRanges', __args__, opts=opts).value

    return AwaitableGetLBIPRangesResult(
        http_ssl_tcp_internals=__ret__.get('httpSslTcpInternals'),
        id=__ret__.get('id'),
        networks=__ret__.get('networks'))
