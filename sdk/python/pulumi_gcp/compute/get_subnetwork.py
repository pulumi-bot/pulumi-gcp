# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetSubnetworkResult:
    """
    A collection of values returned by getSubnetwork.
    """
    def __init__(__self__, description=None, gateway_address=None, ip_cidr_range=None, name=None, network=None, private_ip_google_access=None, project=None, region=None, secondary_ip_ranges=None, self_link=None, id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        Description of this subnetwork.
        """
        if gateway_address and not isinstance(gateway_address, str):
            raise TypeError("Expected argument 'gateway_address' to be a str")
        __self__.gateway_address = gateway_address
        """
        The IP address of the gateway.
        """
        if ip_cidr_range and not isinstance(ip_cidr_range, str):
            raise TypeError("Expected argument 'ip_cidr_range' to be a str")
        __self__.ip_cidr_range = ip_cidr_range
        """
        The range of IP addresses belonging to this subnetwork
        secondary range.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if network and not isinstance(network, str):
            raise TypeError("Expected argument 'network' to be a str")
        __self__.network = network
        """
        The network name or resource link to the parent
        network of this subnetwork.
        """
        if private_ip_google_access and not isinstance(private_ip_google_access, bool):
            raise TypeError("Expected argument 'private_ip_google_access' to be a bool")
        __self__.private_ip_google_access = private_ip_google_access
        """
        Whether the VMs in this subnet
        can access Google services without assigned external IP
        addresses.
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        if secondary_ip_ranges and not isinstance(secondary_ip_ranges, list):
            raise TypeError("Expected argument 'secondary_ip_ranges' to be a list")
        __self__.secondary_ip_ranges = secondary_ip_ranges
        """
        An array of configurations for secondary IP ranges for
        VM instances contained in this subnetwork. Structure is documented below.
        """
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        __self__.self_link = self_link
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        delattr(self, "__await__")
        delattr(self, "__iter__")
        return self

    __iter__ = __await__

def get_subnetwork(name=None,project=None,region=None,self_link=None,opts=None):
    """
    Get a subnetwork within GCE from its name and region.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/compute_subnetwork.html.markdown.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['project'] = project
    __args__['region'] = region
    __args__['selfLink'] = self_link
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getSubnetwork:getSubnetwork', __args__, opts=opts).value

    return GetSubnetworkResult(
        description=__ret__.get('description'),
        gateway_address=__ret__.get('gatewayAddress'),
        ip_cidr_range=__ret__.get('ipCidrRange'),
        name=__ret__.get('name'),
        network=__ret__.get('network'),
        private_ip_google_access=__ret__.get('privateIpGoogleAccess'),
        project=__ret__.get('project'),
        region=__ret__.get('region'),
        secondary_ip_ranges=__ret__.get('secondaryIpRanges'),
        self_link=__ret__.get('selfLink'),
        id=__ret__.get('id'))
