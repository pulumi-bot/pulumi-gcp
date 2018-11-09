# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class InstanceTemplate(pulumi.CustomResource):
    """
    Manages a VM instance template resource within GCE. For more information see
    [the official documentation](https://cloud.google.com/compute/docs/instance-templates)
    and
    [API](https://cloud.google.com/compute/docs/reference/latest/instanceTemplates).
    
    """
    def __init__(__self__, __name__, __opts__=None, can_ip_forward=None, description=None, disks=None, guest_accelerators=None, instance_description=None, labels=None, machine_type=None, metadata=None, metadata_startup_script=None, min_cpu_platform=None, name=None, name_prefix=None, network_interfaces=None, project=None, region=None, schedulings=None, service_account=None, tags=None):
        """Create a InstanceTemplate resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['can_ip_forward'] = can_ip_forward

        __props__['description'] = description

        if not disks:
            raise TypeError('Missing required property disks')
        __props__['disks'] = disks

        __props__['guest_accelerators'] = guest_accelerators

        __props__['instance_description'] = instance_description

        __props__['labels'] = labels

        if not machine_type:
            raise TypeError('Missing required property machine_type')
        __props__['machine_type'] = machine_type

        __props__['metadata'] = metadata

        __props__['metadata_startup_script'] = metadata_startup_script

        __props__['min_cpu_platform'] = min_cpu_platform

        __props__['name'] = name

        __props__['name_prefix'] = name_prefix

        __props__['network_interfaces'] = network_interfaces

        __props__['project'] = project

        __props__['region'] = region

        __props__['schedulings'] = schedulings

        __props__['service_account'] = service_account

        __props__['tags'] = tags

        __props__['metadata_fingerprint'] = None
        __props__['self_link'] = None
        __props__['tags_fingerprint'] = None

        super(InstanceTemplate, __self__).__init__(
            'gcp:compute/instanceTemplate:InstanceTemplate',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

