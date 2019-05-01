# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetForwardingRuleResult:
    """
    A collection of values returned by getForwardingRule.
    """
    def __init__(__self__, backend_service=None, description=None, ip_address=None, ip_protocol=None, load_balancing_scheme=None, name=None, network=None, port_range=None, ports=None, project=None, region=None, self_link=None, subnetwork=None, target=None, id=None):
        if backend_service and not isinstance(backend_service, str):
            raise TypeError("Expected argument 'backend_service' to be a str")
        __self__.backend_service = backend_service
        """
        Backend service, if this forwarding rule has one.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        Description of this forwarding rule.
        """
        if ip_address and not isinstance(ip_address, str):
            raise TypeError("Expected argument 'ip_address' to be a str")
        __self__.ip_address = ip_address
        """
        IP address of this forwarding rule.
        """
        if ip_protocol and not isinstance(ip_protocol, str):
            raise TypeError("Expected argument 'ip_protocol' to be a str")
        __self__.ip_protocol = ip_protocol
        """
        IP protocol of this forwarding rule.
        """
        if load_balancing_scheme and not isinstance(load_balancing_scheme, str):
            raise TypeError("Expected argument 'load_balancing_scheme' to be a str")
        __self__.load_balancing_scheme = load_balancing_scheme
        """
        Type of load balancing of this forwarding rule.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if network and not isinstance(network, str):
            raise TypeError("Expected argument 'network' to be a str")
        __self__.network = network
        """
        Network of this forwarding rule.
        """
        if port_range and not isinstance(port_range, str):
            raise TypeError("Expected argument 'port_range' to be a str")
        __self__.port_range = port_range
        """
        Port range, if this forwarding rule has one.
        """
        if ports and not isinstance(ports, list):
            raise TypeError("Expected argument 'ports' to be a list")
        __self__.ports = ports
        """
        List of ports to use for internal load balancing, if this forwarding rule has any.
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        """
        Region of this forwarding rule.
        """
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        __self__.self_link = self_link
        """
        The URI of the resource.
        """
        if subnetwork and not isinstance(subnetwork, str):
            raise TypeError("Expected argument 'subnetwork' to be a str")
        __self__.subnetwork = subnetwork
        """
        Subnetwork of this forwarding rule.
        """
        if target and not isinstance(target, str):
            raise TypeError("Expected argument 'target' to be a str")
        __self__.target = target
        """
        URL of the target pool, if this forwarding rule has one.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_forwarding_rule(name=None,project=None,region=None,opts=None):
    """
    Get a forwarding rule within GCE from its name.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['project'] = project
    __args__['region'] = region
 .   if opts is None:
         opts = pulumi.ResourceOptions()
     if opts.version is None:
         opts.version = utilities.get_version()
    __ret__ = await pulumi.runtime.invoke('gcp:compute/getForwardingRule:getForwardingRule', __args__, opts=opts)

    return GetForwardingRuleResult(
        backend_service=__ret__.get('backendService'),
        description=__ret__.get('description'),
        ip_address=__ret__.get('ipAddress'),
        ip_protocol=__ret__.get('ipProtocol'),
        load_balancing_scheme=__ret__.get('loadBalancingScheme'),
        name=__ret__.get('name'),
        network=__ret__.get('network'),
        port_range=__ret__.get('portRange'),
        ports=__ret__.get('ports'),
        project=__ret__.get('project'),
        region=__ret__.get('region'),
        self_link=__ret__.get('selfLink'),
        subnetwork=__ret__.get('subnetwork'),
        target=__ret__.get('target'),
        id=__ret__.get('id'))
