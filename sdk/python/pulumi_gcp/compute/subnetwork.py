# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Subnetwork(pulumi.CustomResource):
    creation_timestamp: pulumi.Output[str]
    description: pulumi.Output[str]
    enable_flow_logs: pulumi.Output[bool]
    fingerprint: pulumi.Output[str]
    gateway_address: pulumi.Output[str]
    ip_cidr_range: pulumi.Output[str]
    log_config: pulumi.Output[dict]
    name: pulumi.Output[str]
    network: pulumi.Output[str]
    private_ip_google_access: pulumi.Output[bool]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    secondary_ip_ranges: pulumi.Output[list]
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    def __init__(__self__, resource_name, opts=None, description=None, enable_flow_logs=None, ip_cidr_range=None, log_config=None, name=None, network=None, private_ip_google_access=None, project=None, region=None, secondary_ip_ranges=None, __props__=None, __name__=None, __opts__=None):
        """
        A VPC network is a virtual version of the traditional physical networks
        that exist within and between physical data centers. A VPC network
        provides connectivity for your Compute Engine virtual machine (VM)
        instances, Container Engine containers, App Engine Flex services, and
        other network-related resources.
        
        Each GCP project contains one or more VPC networks. Each VPC network is a
        global entity spanning all GCP regions. This global VPC network allows VM
        instances and other resources to communicate with each other via internal,
        private IP addresses.
        
        Each VPC network is subdivided into subnets, and each subnet is contained
        within a single region. You can have more than one subnet in a region for
        a given VPC network. Each subnet has a contiguous private RFC1918 IP
        space. You create instances, containers, and the like in these subnets.
        When you create an instance, you must create it in a subnet, and the
        instance draws its internal IP address from that subnet.
        
        Virtual machine (VM) instances in a VPC network can communicate with
        instances in all other subnets of the same VPC network, regardless of
        region, using their RFC1918 private IP addresses. You can isolate portions
        of the network, even entire subnets, using firewall rules.
        
        
        To get more information about Subnetwork, see:
        
        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/subnetworks)
        * How-to Guides
            * [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
            * [Cloud Networking](https://cloud.google.com/vpc/docs/using-vpc)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        
        The **log_config** object supports the following:
        
          * `aggregation_interval` (`pulumi.Input[str]`)
          * `flow_sampling` (`pulumi.Input[float]`)
          * `metadata` (`pulumi.Input[str]`)
        
        The **secondary_ip_ranges** object supports the following:
        
          * `ip_cidr_range` (`pulumi.Input[str]`)
          * `range_name` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_subnetwork.html.markdown.
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

            __props__['description'] = description
            __props__['enable_flow_logs'] = enable_flow_logs
            if ip_cidr_range is None:
                raise TypeError("Missing required property 'ip_cidr_range'")
            __props__['ip_cidr_range'] = ip_cidr_range
            __props__['log_config'] = log_config
            __props__['name'] = name
            if network is None:
                raise TypeError("Missing required property 'network'")
            __props__['network'] = network
            __props__['private_ip_google_access'] = private_ip_google_access
            __props__['project'] = project
            __props__['region'] = region
            __props__['secondary_ip_ranges'] = secondary_ip_ranges
            __props__['creation_timestamp'] = None
            __props__['fingerprint'] = None
            __props__['gateway_address'] = None
            __props__['self_link'] = None
        super(Subnetwork, __self__).__init__(
            'gcp:compute/subnetwork:Subnetwork',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, creation_timestamp=None, description=None, enable_flow_logs=None, fingerprint=None, gateway_address=None, ip_cidr_range=None, log_config=None, name=None, network=None, private_ip_google_access=None, project=None, region=None, secondary_ip_ranges=None, self_link=None):
        """
        Get an existing Subnetwork resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        
        The **secondary_ip_ranges** object supports the following:
        
          * `ip_cidr_range` (`pulumi.Input[str]`)
          * `range_name` (`pulumi.Input[str]`)
        
        The **log_config** object supports the following:
        
          * `aggregation_interval` (`pulumi.Input[str]`)
          * `flow_sampling` (`pulumi.Input[float]`)
          * `metadata` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_subnetwork.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["enable_flow_logs"] = enable_flow_logs
        __props__["fingerprint"] = fingerprint
        __props__["gateway_address"] = gateway_address
        __props__["ip_cidr_range"] = ip_cidr_range
        __props__["log_config"] = log_config
        __props__["name"] = name
        __props__["network"] = network
        __props__["private_ip_google_access"] = private_ip_google_access
        __props__["project"] = project
        __props__["region"] = region
        __props__["secondary_ip_ranges"] = secondary_ip_ranges
        __props__["self_link"] = self_link
        return Subnetwork(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

