# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GlobalNetworkEndpoint(pulumi.CustomResource):
    fqdn: pulumi.Output[str]
    """
    Fully qualified domain name of network endpoint.
    This can only be specified when network_endpoint_type of the NEG is INTERNET_FQDN_PORT.
    """
    global_network_endpoint_group: pulumi.Output[str]
    """
    The global network endpoint group this endpoint is part of.
    """
    ip_address: pulumi.Output[str]
    """
    IPv4 address external endpoint.
    """
    port: pulumi.Output[float]
    """
    Port number of the external endpoint.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    def __init__(__self__, resource_name, opts=None, fqdn=None, global_network_endpoint_group=None, ip_address=None, port=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        A Global Network endpoint represents a IP address and port combination that exists outside of GCP.
        **NOTE**: Global network endpoints cannot be created outside of a
        global network endpoint group.


        To get more information about GlobalNetworkEndpoint, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/networkEndpointGroups)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/load-balancing/docs/negs/)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fqdn: Fully qualified domain name of network endpoint.
               This can only be specified when network_endpoint_type of the NEG is INTERNET_FQDN_PORT.
        :param pulumi.Input[str] global_network_endpoint_group: The global network endpoint group this endpoint is part of.
        :param pulumi.Input[str] ip_address: IPv4 address external endpoint.
        :param pulumi.Input[float] port: Port number of the external endpoint.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['fqdn'] = fqdn
            if global_network_endpoint_group is None:
                raise TypeError("Missing required property 'global_network_endpoint_group'")
            __props__['global_network_endpoint_group'] = global_network_endpoint_group
            __props__['ip_address'] = ip_address
            if port is None:
                raise TypeError("Missing required property 'port'")
            __props__['port'] = port
            __props__['project'] = project
        super(GlobalNetworkEndpoint, __self__).__init__(
            'gcp:compute/globalNetworkEndpoint:GlobalNetworkEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, fqdn=None, global_network_endpoint_group=None, ip_address=None, port=None, project=None):
        """
        Get an existing GlobalNetworkEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fqdn: Fully qualified domain name of network endpoint.
               This can only be specified when network_endpoint_type of the NEG is INTERNET_FQDN_PORT.
        :param pulumi.Input[str] global_network_endpoint_group: The global network endpoint group this endpoint is part of.
        :param pulumi.Input[str] ip_address: IPv4 address external endpoint.
        :param pulumi.Input[float] port: Port number of the external endpoint.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["fqdn"] = fqdn
        __props__["global_network_endpoint_group"] = global_network_endpoint_group
        __props__["ip_address"] = ip_address
        __props__["port"] = port
        __props__["project"] = project
        return GlobalNetworkEndpoint(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

