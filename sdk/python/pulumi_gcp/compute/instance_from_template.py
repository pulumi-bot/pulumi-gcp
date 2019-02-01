# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class InstanceFromTemplate(pulumi.CustomResource):
    allow_stopping_for_update: pulumi.Output[bool]
    attached_disks: pulumi.Output[list]
    boot_disk: pulumi.Output[dict]
    can_ip_forward: pulumi.Output[bool]
    cpu_platform: pulumi.Output[str]
    deletion_protection: pulumi.Output[bool]
    description: pulumi.Output[str]
    guest_accelerators: pulumi.Output[list]
    instance_id: pulumi.Output[str]
    label_fingerprint: pulumi.Output[str]
    labels: pulumi.Output[dict]
    machine_type: pulumi.Output[str]
    metadata: pulumi.Output[dict]
    metadata_fingerprint: pulumi.Output[str]
    metadata_startup_script: pulumi.Output[str]
    min_cpu_platform: pulumi.Output[str]
    name: pulumi.Output[str]
    """
    A unique name for the resource, required by GCE.
    Changing this forces a new resource to be created.
    """
    network_interfaces: pulumi.Output[list]
    project: pulumi.Output[str]
    scheduling: pulumi.Output[dict]
    scratch_disks: pulumi.Output[list]
    self_link: pulumi.Output[str]
    service_account: pulumi.Output[dict]
    source_instance_template: pulumi.Output[str]
    """
    Name or self link of an instance
    template to create the instance based on.
    """
    tags: pulumi.Output[list]
    tags_fingerprint: pulumi.Output[str]
    zone: pulumi.Output[str]
    """
    The zone that the machine should be created in. If not
    set, the provider zone is used.
    """
    def __init__(__self__, __name__, __opts__=None, allow_stopping_for_update=None, attached_disks=None, boot_disk=None, can_ip_forward=None, deletion_protection=None, description=None, guest_accelerators=None, labels=None, machine_type=None, metadata=None, metadata_startup_script=None, min_cpu_platform=None, name=None, network_interfaces=None, project=None, scheduling=None, scratch_disks=None, service_account=None, source_instance_template=None, tags=None, zone=None):
        """
        Manages a VM instance resource within GCE. For more information see
        [the official documentation](https://cloud.google.com/compute/docs/instances)
        and
        [API](https://cloud.google.com/compute/docs/reference/latest/instances).
        
        This resource is specifically to create a compute instance from a given
        `source_instance_template`. To create an instance without a template, use the
        `google_compute_instance` resource.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[bool] allow_stopping_for_update
        :param pulumi.Input[list] attached_disks
        :param pulumi.Input[dict] boot_disk
        :param pulumi.Input[bool] can_ip_forward
        :param pulumi.Input[bool] deletion_protection
        :param pulumi.Input[str] description
        :param pulumi.Input[list] guest_accelerators
        :param pulumi.Input[dict] labels
        :param pulumi.Input[str] machine_type
        :param pulumi.Input[dict] metadata
        :param pulumi.Input[str] metadata_startup_script
        :param pulumi.Input[str] min_cpu_platform
        :param pulumi.Input[str] name: A unique name for the resource, required by GCE.
               Changing this forces a new resource to be created.
        :param pulumi.Input[list] network_interfaces
        :param pulumi.Input[str] project
        :param pulumi.Input[dict] scheduling
        :param pulumi.Input[list] scratch_disks
        :param pulumi.Input[dict] service_account
        :param pulumi.Input[str] source_instance_template: Name or self link of an instance
               template to create the instance based on.
        :param pulumi.Input[list] tags
        :param pulumi.Input[str] zone: The zone that the machine should be created in. If not
               set, the provider zone is used.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['allow_stopping_for_update'] = allow_stopping_for_update

        __props__['attached_disks'] = attached_disks

        __props__['boot_disk'] = boot_disk

        __props__['can_ip_forward'] = can_ip_forward

        __props__['deletion_protection'] = deletion_protection

        __props__['description'] = description

        __props__['guest_accelerators'] = guest_accelerators

        __props__['labels'] = labels

        __props__['machine_type'] = machine_type

        __props__['metadata'] = metadata

        __props__['metadata_startup_script'] = metadata_startup_script

        __props__['min_cpu_platform'] = min_cpu_platform

        __props__['name'] = name

        __props__['network_interfaces'] = network_interfaces

        __props__['project'] = project

        __props__['scheduling'] = scheduling

        __props__['scratch_disks'] = scratch_disks

        __props__['service_account'] = service_account

        if not source_instance_template:
            raise TypeError('Missing required property source_instance_template')
        __props__['source_instance_template'] = source_instance_template

        __props__['tags'] = tags

        __props__['zone'] = zone

        __props__['cpu_platform'] = None
        __props__['instance_id'] = None
        __props__['label_fingerprint'] = None
        __props__['metadata_fingerprint'] = None
        __props__['self_link'] = None
        __props__['tags_fingerprint'] = None

        super(InstanceFromTemplate, __self__).__init__(
            'gcp:compute/instanceFromTemplate:InstanceFromTemplate',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

