# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class AttachedDisk(pulumi.CustomResource):
    device_name: pulumi.Output[str]
    disk: pulumi.Output[str]
    instance: pulumi.Output[str]
    mode: pulumi.Output[str]
    project: pulumi.Output[str]
    zone: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, device_name=None, disk=None, instance=None, mode=None, project=None, zone=None, __name__=None, __opts__=None):
        """
        Persistent disks can be attached to a compute instance using [the `attached_disk`
        section within the compute instance configuration](https://www.terraform.io/docs/providers/google/r/compute_instance.html#attached_disk).
        However there may be situations where managing the attached disks via the compute
        instance config isn't preferable or possible, such as attaching dynamic
        numbers of disks using the `count` variable.
        
        
        To get more information about attaching disks, see:
        
        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instances/attachDisk)
        * [Resource: google_compute_disk](https://www.terraform.io/docs/providers/google/r/compute_disk.html)
        * How-to Guides
            * [Adding a persistent disk](https://cloud.google.com/compute/docs/disks/add-persistent-disk)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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

        __props__['device_name'] = device_name

        if disk is None:
            raise TypeError("Missing required property 'disk'")
        __props__['disk'] = disk

        if instance is None:
            raise TypeError("Missing required property 'instance'")
        __props__['instance'] = instance

        __props__['mode'] = mode

        __props__['project'] = project

        __props__['zone'] = zone

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(AttachedDisk, __self__).__init__(
            'gcp:compute/attachedDisk:AttachedDisk',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

