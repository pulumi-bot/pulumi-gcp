# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['InstanceFromTemplate']


class InstanceFromTemplate(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_stopping_for_update: Optional[pulumi.Input[bool]] = None,
                 attached_disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceFromTemplateAttachedDiskArgs']]]]] = None,
                 boot_disk: Optional[pulumi.Input[pulumi.InputType['InstanceFromTemplateBootDiskArgs']]] = None,
                 can_ip_forward: Optional[pulumi.Input[bool]] = None,
                 confidential_instance_config: Optional[pulumi.Input[pulumi.InputType['InstanceFromTemplateConfidentialInstanceConfigArgs']]] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 desired_status: Optional[pulumi.Input[str]] = None,
                 enable_display: Optional[pulumi.Input[bool]] = None,
                 guest_accelerators: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceFromTemplateGuestAcceleratorArgs']]]]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 machine_type: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 metadata_startup_script: Optional[pulumi.Input[str]] = None,
                 min_cpu_platform: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceFromTemplateNetworkInterfaceArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 resource_policies: Optional[pulumi.Input[str]] = None,
                 scheduling: Optional[pulumi.Input[pulumi.InputType['InstanceFromTemplateSchedulingArgs']]] = None,
                 scratch_disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceFromTemplateScratchDiskArgs']]]]] = None,
                 service_account: Optional[pulumi.Input[pulumi.InputType['InstanceFromTemplateServiceAccountArgs']]] = None,
                 shielded_instance_config: Optional[pulumi.Input[pulumi.InputType['InstanceFromTemplateShieldedInstanceConfigArgs']]] = None,
                 source_instance_template: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a VM instance resource within GCE. For more information see
        [the official documentation](https://cloud.google.com/compute/docs/instances)
        and
        [API](https://cloud.google.com/compute/docs/reference/latest/instances).

        This resource is specifically to create a compute instance from a given
        `source_instance_template`. To create an instance without a template, use the
        `compute.Instance` resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        tpl_instance_template = gcp.compute.InstanceTemplate("tplInstanceTemplate",
            machine_type="e2-medium",
            disks=[{
                "source_image": "debian-cloud/debian-9",
                "autoDelete": True,
                "disk_size_gb": 100,
                "boot": True,
            }],
            network_interfaces=[{
                "network": "default",
            }],
            metadata={
                "foo": "bar",
            },
            can_ip_forward=True)
        tpl_instance_from_template = gcp.compute.InstanceFromTemplate("tplInstanceFromTemplate",
            zone="us-central1-a",
            source_instance_template=tpl_instance_template.id,
            can_ip_forward=False,
            labels={
                "my_key": "my_value",
            })
        ```

        ## Import

        This resource does not support import.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_stopping_for_update: If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
               stopping the instance without setting this field, the update will fail.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceFromTemplateAttachedDiskArgs']]]] attached_disks: List of disks attached to the instance
        :param pulumi.Input[pulumi.InputType['InstanceFromTemplateBootDiskArgs']] boot_disk: The boot disk for the instance.
        :param pulumi.Input[bool] can_ip_forward: Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
        :param pulumi.Input[pulumi.InputType['InstanceFromTemplateConfidentialInstanceConfigArgs']] confidential_instance_config: The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
               to create.
        :param pulumi.Input[bool] deletion_protection: Whether deletion protection is enabled on this instance.
        :param pulumi.Input[str] description: A brief description of the resource.
        :param pulumi.Input[str] desired_status: Desired status of the instance. Either "RUNNING" or "TERMINATED".
        :param pulumi.Input[bool] enable_display: Whether the instance has virtual displays enabled.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceFromTemplateGuestAcceleratorArgs']]]] guest_accelerators: List of the type and count of accelerator cards attached to the instance.
        :param pulumi.Input[str] hostname: A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
               labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
               entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs assigned to the instance.
        :param pulumi.Input[str] machine_type: The machine type to create.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Metadata key/value pairs made available within the instance.
        :param pulumi.Input[str] metadata_startup_script: Metadata startup scripts made available within the instance.
        :param pulumi.Input[str] min_cpu_platform: The minimum CPU platform specified for the VM instance.
        :param pulumi.Input[str] name: A unique name for the resource, required by GCE.
               Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceFromTemplateNetworkInterfaceArgs']]]] network_interfaces: The networks attached to the instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
               self_link nor project are provided, the provider project is used.
        :param pulumi.Input[str] resource_policies: A list of short names or self_links of resource policies to attach to the instance. Modifying this list will cause the
               instance to recreate. Currently a max of 1 resource policy is supported.
        :param pulumi.Input[pulumi.InputType['InstanceFromTemplateSchedulingArgs']] scheduling: The scheduling strategy being used by the instance.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceFromTemplateScratchDiskArgs']]]] scratch_disks: The scratch disks attached to the instance.
        :param pulumi.Input[pulumi.InputType['InstanceFromTemplateServiceAccountArgs']] service_account: The service account to attach to the instance.
        :param pulumi.Input[pulumi.InputType['InstanceFromTemplateShieldedInstanceConfigArgs']] shielded_instance_config: The shielded vm config being used by the instance.
        :param pulumi.Input[str] source_instance_template: Name or self link of an instance
               template to create the instance based on.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The list of tags attached to the instance.
        :param pulumi.Input[str] zone: The zone that the machine should be created in. If not
               set, the provider zone is used.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['allow_stopping_for_update'] = allow_stopping_for_update
            __props__['attached_disks'] = attached_disks
            __props__['boot_disk'] = boot_disk
            __props__['can_ip_forward'] = can_ip_forward
            __props__['confidential_instance_config'] = confidential_instance_config
            __props__['deletion_protection'] = deletion_protection
            __props__['description'] = description
            __props__['desired_status'] = desired_status
            __props__['enable_display'] = enable_display
            __props__['guest_accelerators'] = guest_accelerators
            __props__['hostname'] = hostname
            __props__['labels'] = labels
            __props__['machine_type'] = machine_type
            __props__['metadata'] = metadata
            __props__['metadata_startup_script'] = metadata_startup_script
            __props__['min_cpu_platform'] = min_cpu_platform
            __props__['name'] = name
            __props__['network_interfaces'] = network_interfaces
            __props__['project'] = project
            __props__['resource_policies'] = resource_policies
            __props__['scheduling'] = scheduling
            __props__['scratch_disks'] = scratch_disks
            __props__['service_account'] = service_account
            __props__['shielded_instance_config'] = shielded_instance_config
            if source_instance_template is None:
                raise TypeError("Missing required property 'source_instance_template'")
            __props__['source_instance_template'] = source_instance_template
            __props__['tags'] = tags
            __props__['zone'] = zone
            __props__['cpu_platform'] = None
            __props__['current_status'] = None
            __props__['instance_id'] = None
            __props__['label_fingerprint'] = None
            __props__['metadata_fingerprint'] = None
            __props__['self_link'] = None
            __props__['tags_fingerprint'] = None
        super(InstanceFromTemplate, __self__).__init__(
            'gcp:compute/instanceFromTemplate:InstanceFromTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_stopping_for_update: Optional[pulumi.Input[bool]] = None,
            attached_disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceFromTemplateAttachedDiskArgs']]]]] = None,
            boot_disk: Optional[pulumi.Input[pulumi.InputType['InstanceFromTemplateBootDiskArgs']]] = None,
            can_ip_forward: Optional[pulumi.Input[bool]] = None,
            confidential_instance_config: Optional[pulumi.Input[pulumi.InputType['InstanceFromTemplateConfidentialInstanceConfigArgs']]] = None,
            cpu_platform: Optional[pulumi.Input[str]] = None,
            current_status: Optional[pulumi.Input[str]] = None,
            deletion_protection: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            desired_status: Optional[pulumi.Input[str]] = None,
            enable_display: Optional[pulumi.Input[bool]] = None,
            guest_accelerators: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceFromTemplateGuestAcceleratorArgs']]]]] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            label_fingerprint: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            machine_type: Optional[pulumi.Input[str]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            metadata_fingerprint: Optional[pulumi.Input[str]] = None,
            metadata_startup_script: Optional[pulumi.Input[str]] = None,
            min_cpu_platform: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceFromTemplateNetworkInterfaceArgs']]]]] = None,
            project: Optional[pulumi.Input[str]] = None,
            resource_policies: Optional[pulumi.Input[str]] = None,
            scheduling: Optional[pulumi.Input[pulumi.InputType['InstanceFromTemplateSchedulingArgs']]] = None,
            scratch_disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceFromTemplateScratchDiskArgs']]]]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            service_account: Optional[pulumi.Input[pulumi.InputType['InstanceFromTemplateServiceAccountArgs']]] = None,
            shielded_instance_config: Optional[pulumi.Input[pulumi.InputType['InstanceFromTemplateShieldedInstanceConfigArgs']]] = None,
            source_instance_template: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tags_fingerprint: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'InstanceFromTemplate':
        """
        Get an existing InstanceFromTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_stopping_for_update: If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
               stopping the instance without setting this field, the update will fail.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceFromTemplateAttachedDiskArgs']]]] attached_disks: List of disks attached to the instance
        :param pulumi.Input[pulumi.InputType['InstanceFromTemplateBootDiskArgs']] boot_disk: The boot disk for the instance.
        :param pulumi.Input[bool] can_ip_forward: Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
        :param pulumi.Input[pulumi.InputType['InstanceFromTemplateConfidentialInstanceConfigArgs']] confidential_instance_config: The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
               to create.
        :param pulumi.Input[str] cpu_platform: The CPU platform used by this instance.
        :param pulumi.Input[str] current_status: Current status of the instance.
        :param pulumi.Input[bool] deletion_protection: Whether deletion protection is enabled on this instance.
        :param pulumi.Input[str] description: A brief description of the resource.
        :param pulumi.Input[str] desired_status: Desired status of the instance. Either "RUNNING" or "TERMINATED".
        :param pulumi.Input[bool] enable_display: Whether the instance has virtual displays enabled.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceFromTemplateGuestAcceleratorArgs']]]] guest_accelerators: List of the type and count of accelerator cards attached to the instance.
        :param pulumi.Input[str] hostname: A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
               labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
               entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
        :param pulumi.Input[str] instance_id: The server-assigned unique identifier of this instance.
        :param pulumi.Input[str] label_fingerprint: The unique fingerprint of the labels.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs assigned to the instance.
        :param pulumi.Input[str] machine_type: The machine type to create.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Metadata key/value pairs made available within the instance.
        :param pulumi.Input[str] metadata_fingerprint: The unique fingerprint of the metadata.
        :param pulumi.Input[str] metadata_startup_script: Metadata startup scripts made available within the instance.
        :param pulumi.Input[str] min_cpu_platform: The minimum CPU platform specified for the VM instance.
        :param pulumi.Input[str] name: A unique name for the resource, required by GCE.
               Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceFromTemplateNetworkInterfaceArgs']]]] network_interfaces: The networks attached to the instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
               self_link nor project are provided, the provider project is used.
        :param pulumi.Input[str] resource_policies: A list of short names or self_links of resource policies to attach to the instance. Modifying this list will cause the
               instance to recreate. Currently a max of 1 resource policy is supported.
        :param pulumi.Input[pulumi.InputType['InstanceFromTemplateSchedulingArgs']] scheduling: The scheduling strategy being used by the instance.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceFromTemplateScratchDiskArgs']]]] scratch_disks: The scratch disks attached to the instance.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[pulumi.InputType['InstanceFromTemplateServiceAccountArgs']] service_account: The service account to attach to the instance.
        :param pulumi.Input[pulumi.InputType['InstanceFromTemplateShieldedInstanceConfigArgs']] shielded_instance_config: The shielded vm config being used by the instance.
        :param pulumi.Input[str] source_instance_template: Name or self link of an instance
               template to create the instance based on.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The list of tags attached to the instance.
        :param pulumi.Input[str] tags_fingerprint: The unique fingerprint of the tags.
        :param pulumi.Input[str] zone: The zone that the machine should be created in. If not
               set, the provider zone is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allow_stopping_for_update"] = allow_stopping_for_update
        __props__["attached_disks"] = attached_disks
        __props__["boot_disk"] = boot_disk
        __props__["can_ip_forward"] = can_ip_forward
        __props__["confidential_instance_config"] = confidential_instance_config
        __props__["cpu_platform"] = cpu_platform
        __props__["current_status"] = current_status
        __props__["deletion_protection"] = deletion_protection
        __props__["description"] = description
        __props__["desired_status"] = desired_status
        __props__["enable_display"] = enable_display
        __props__["guest_accelerators"] = guest_accelerators
        __props__["hostname"] = hostname
        __props__["instance_id"] = instance_id
        __props__["label_fingerprint"] = label_fingerprint
        __props__["labels"] = labels
        __props__["machine_type"] = machine_type
        __props__["metadata"] = metadata
        __props__["metadata_fingerprint"] = metadata_fingerprint
        __props__["metadata_startup_script"] = metadata_startup_script
        __props__["min_cpu_platform"] = min_cpu_platform
        __props__["name"] = name
        __props__["network_interfaces"] = network_interfaces
        __props__["project"] = project
        __props__["resource_policies"] = resource_policies
        __props__["scheduling"] = scheduling
        __props__["scratch_disks"] = scratch_disks
        __props__["self_link"] = self_link
        __props__["service_account"] = service_account
        __props__["shielded_instance_config"] = shielded_instance_config
        __props__["source_instance_template"] = source_instance_template
        __props__["tags"] = tags
        __props__["tags_fingerprint"] = tags_fingerprint
        __props__["zone"] = zone
        return InstanceFromTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowStoppingForUpdate")
    def allow_stopping_for_update(self) -> pulumi.Output[bool]:
        """
        If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
        stopping the instance without setting this field, the update will fail.
        """
        return pulumi.get(self, "allow_stopping_for_update")

    @property
    @pulumi.getter(name="attachedDisks")
    def attached_disks(self) -> pulumi.Output[Sequence['outputs.InstanceFromTemplateAttachedDisk']]:
        """
        List of disks attached to the instance
        """
        return pulumi.get(self, "attached_disks")

    @property
    @pulumi.getter(name="bootDisk")
    def boot_disk(self) -> pulumi.Output['outputs.InstanceFromTemplateBootDisk']:
        """
        The boot disk for the instance.
        """
        return pulumi.get(self, "boot_disk")

    @property
    @pulumi.getter(name="canIpForward")
    def can_ip_forward(self) -> pulumi.Output[bool]:
        """
        Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
        """
        return pulumi.get(self, "can_ip_forward")

    @property
    @pulumi.getter(name="confidentialInstanceConfig")
    def confidential_instance_config(self) -> pulumi.Output['outputs.InstanceFromTemplateConfidentialInstanceConfig']:
        """
        The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
        to create.
        """
        return pulumi.get(self, "confidential_instance_config")

    @property
    @pulumi.getter(name="cpuPlatform")
    def cpu_platform(self) -> pulumi.Output[str]:
        """
        The CPU platform used by this instance.
        """
        return pulumi.get(self, "cpu_platform")

    @property
    @pulumi.getter(name="currentStatus")
    def current_status(self) -> pulumi.Output[str]:
        """
        Current status of the instance.
        """
        return pulumi.get(self, "current_status")

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> pulumi.Output[bool]:
        """
        Whether deletion protection is enabled on this instance.
        """
        return pulumi.get(self, "deletion_protection")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        A brief description of the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="desiredStatus")
    def desired_status(self) -> pulumi.Output[str]:
        """
        Desired status of the instance. Either "RUNNING" or "TERMINATED".
        """
        return pulumi.get(self, "desired_status")

    @property
    @pulumi.getter(name="enableDisplay")
    def enable_display(self) -> pulumi.Output[bool]:
        """
        Whether the instance has virtual displays enabled.
        """
        return pulumi.get(self, "enable_display")

    @property
    @pulumi.getter(name="guestAccelerators")
    def guest_accelerators(self) -> pulumi.Output[Sequence['outputs.InstanceFromTemplateGuestAccelerator']]:
        """
        List of the type and count of accelerator cards attached to the instance.
        """
        return pulumi.get(self, "guest_accelerators")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        """
        A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
        labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
        entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The server-assigned unique identifier of this instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="labelFingerprint")
    def label_fingerprint(self) -> pulumi.Output[str]:
        """
        The unique fingerprint of the labels.
        """
        return pulumi.get(self, "label_fingerprint")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A set of key/value label pairs assigned to the instance.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="machineType")
    def machine_type(self) -> pulumi.Output[str]:
        """
        The machine type to create.
        """
        return pulumi.get(self, "machine_type")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Metadata key/value pairs made available within the instance.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="metadataFingerprint")
    def metadata_fingerprint(self) -> pulumi.Output[str]:
        """
        The unique fingerprint of the metadata.
        """
        return pulumi.get(self, "metadata_fingerprint")

    @property
    @pulumi.getter(name="metadataStartupScript")
    def metadata_startup_script(self) -> pulumi.Output[str]:
        """
        Metadata startup scripts made available within the instance.
        """
        return pulumi.get(self, "metadata_startup_script")

    @property
    @pulumi.getter(name="minCpuPlatform")
    def min_cpu_platform(self) -> pulumi.Output[str]:
        """
        The minimum CPU platform specified for the VM instance.
        """
        return pulumi.get(self, "min_cpu_platform")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique name for the resource, required by GCE.
        Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkInterfaces")
    def network_interfaces(self) -> pulumi.Output[Sequence['outputs.InstanceFromTemplateNetworkInterface']]:
        """
        The networks attached to the instance.
        """
        return pulumi.get(self, "network_interfaces")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
        self_link nor project are provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="resourcePolicies")
    def resource_policies(self) -> pulumi.Output[str]:
        """
        A list of short names or self_links of resource policies to attach to the instance. Modifying this list will cause the
        instance to recreate. Currently a max of 1 resource policy is supported.
        """
        return pulumi.get(self, "resource_policies")

    @property
    @pulumi.getter
    def scheduling(self) -> pulumi.Output['outputs.InstanceFromTemplateScheduling']:
        """
        The scheduling strategy being used by the instance.
        """
        return pulumi.get(self, "scheduling")

    @property
    @pulumi.getter(name="scratchDisks")
    def scratch_disks(self) -> pulumi.Output[Sequence['outputs.InstanceFromTemplateScratchDisk']]:
        """
        The scratch disks attached to the instance.
        """
        return pulumi.get(self, "scratch_disks")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> pulumi.Output['outputs.InstanceFromTemplateServiceAccount']:
        """
        The service account to attach to the instance.
        """
        return pulumi.get(self, "service_account")

    @property
    @pulumi.getter(name="shieldedInstanceConfig")
    def shielded_instance_config(self) -> pulumi.Output['outputs.InstanceFromTemplateShieldedInstanceConfig']:
        """
        The shielded vm config being used by the instance.
        """
        return pulumi.get(self, "shielded_instance_config")

    @property
    @pulumi.getter(name="sourceInstanceTemplate")
    def source_instance_template(self) -> pulumi.Output[str]:
        """
        Name or self link of an instance
        template to create the instance based on.
        """
        return pulumi.get(self, "source_instance_template")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of tags attached to the instance.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsFingerprint")
    def tags_fingerprint(self) -> pulumi.Output[str]:
        """
        The unique fingerprint of the tags.
        """
        return pulumi.get(self, "tags_fingerprint")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        The zone that the machine should be created in. If not
        set, the provider zone is used.
        """
        return pulumi.get(self, "zone")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

