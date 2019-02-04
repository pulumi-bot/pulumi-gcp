# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Subnetwork(pulumi.CustomResource):
    creation_timestamp: pulumi.Output[str]
    description: pulumi.Output[str]
    enable_flow_logs: pulumi.Output[bool]
    fingerprint: pulumi.Output[str]
    gateway_address: pulumi.Output[str]
    ip_cidr_range: pulumi.Output[str]
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
    def __init__(__self__, __name__, __opts__=None, description=None, enable_flow_logs=None, ip_cidr_range=None, name=None, network=None, private_ip_google_access=None, project=None, region=None, secondary_ip_ranges=None):
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
        
        <div class = "oics-button" style="float: right; margin: 0 0 -15px">
          <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=subnetwork_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
            <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
          </a>
        </div>
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] description
        :param pulumi.Input[bool] enable_flow_logs
        :param pulumi.Input[str] ip_cidr_range
        :param pulumi.Input[str] name
        :param pulumi.Input[str] network
        :param pulumi.Input[bool] private_ip_google_access
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region
        :param pulumi.Input[list] secondary_ip_ranges
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        __props__['enable_flow_logs'] = enable_flow_logs

        if ip_cidr_range is None:
            raise TypeError('Missing required property ip_cidr_range')
        __props__['ip_cidr_range'] = ip_cidr_range

        __props__['name'] = name

        if network is None:
            raise TypeError('Missing required property network')
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
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

