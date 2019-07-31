# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ForwardingRule(pulumi.CustomResource):
    all_ports: pulumi.Output[bool]
    backend_service: pulumi.Output[str]
    creation_timestamp: pulumi.Output[str]
    description: pulumi.Output[str]
    ip_address: pulumi.Output[str]
    ip_protocol: pulumi.Output[str]
    ip_version: pulumi.Output[str]
    label_fingerprint: pulumi.Output[str]
    labels: pulumi.Output[dict]
    load_balancing_scheme: pulumi.Output[str]
    name: pulumi.Output[str]
    network: pulumi.Output[str]
    network_tier: pulumi.Output[str]
    port_range: pulumi.Output[str]
    ports: pulumi.Output[list]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    service_label: pulumi.Output[str]
    service_name: pulumi.Output[str]
    subnetwork: pulumi.Output[str]
    target: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, all_ports=None, backend_service=None, description=None, ip_address=None, ip_protocol=None, ip_version=None, labels=None, load_balancing_scheme=None, name=None, network=None, network_tier=None, port_range=None, ports=None, project=None, region=None, service_label=None, subnetwork=None, target=None, __name__=None, __opts__=None):
        """
        A ForwardingRule resource. A ForwardingRule resource specifies which pool
        of target virtual machines to forward a packet to if it matches the given
        [IPAddress, IPProtocol, portRange] tuple.
        
        
        To get more information about ForwardingRule, see:
        
        * [API documentation](https://cloud.google.com/compute/docs/reference/v1/forwardingRule)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/network/forwarding-rules)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_forwarding_rule.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['all_ports'] = all_ports

        __props__['backend_service'] = backend_service

        __props__['description'] = description

        __props__['ip_address'] = ip_address

        __props__['ip_protocol'] = ip_protocol

        __props__['ip_version'] = ip_version

        __props__['labels'] = labels

        __props__['load_balancing_scheme'] = load_balancing_scheme

        __props__['name'] = name

        __props__['network'] = network

        __props__['network_tier'] = network_tier

        __props__['port_range'] = port_range

        __props__['ports'] = ports

        __props__['project'] = project

        __props__['region'] = region

        __props__['service_label'] = service_label

        __props__['subnetwork'] = subnetwork

        __props__['target'] = target

        __props__['creation_timestamp'] = None
        __props__['label_fingerprint'] = None
        __props__['self_link'] = None
        __props__['service_name'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(ForwardingRule, __self__).__init__(
            'gcp:compute/forwardingRule:ForwardingRule',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

