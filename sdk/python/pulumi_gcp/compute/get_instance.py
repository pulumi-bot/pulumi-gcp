# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from ._inputs import *
from . import outputs


class GetInstanceResult:
    """
    A collection of values returned by getInstance.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, allow_stopping_for_update=None, attached_disks=None, boot_disks=None, can_ip_forward=None, cpu_platform=None, current_status=None, deletion_protection=None, description=None, desired_status=None, enable_display=None, guest_accelerators=None, hostname=None, id=None, instance_id=None, label_fingerprint=None, labels=None, machine_type=None, metadata=None, metadata_fingerprint=None, metadata_startup_script=None, min_cpu_platform=None, name=None, network_interfaces=None, project=None, resource_policies=None, schedulings=None, scratch_disks=None, self_link=None, service_accounts=None, shielded_instance_configs=None, tags=None, tags_fingerprint=None, zone=None) -> None:
        if allow_stopping_for_update and not isinstance(allow_stopping_for_update, bool):
            raise TypeError("Expected argument 'allow_stopping_for_update' to be a bool")
        __self__.allow_stopping_for_update = allow_stopping_for_update
        if attached_disks and not isinstance(attached_disks, list):
            raise TypeError("Expected argument 'attached_disks' to be a list")
        __self__.attached_disks = attached_disks
        """
        List of disks attached to the instance. Structure is documented below.
        """
        if boot_disks and not isinstance(boot_disks, list):
            raise TypeError("Expected argument 'boot_disks' to be a list")
        __self__.boot_disks = boot_disks
        """
        The boot disk for the instance. Structure is documented below.
        """
        if can_ip_forward and not isinstance(can_ip_forward, bool):
            raise TypeError("Expected argument 'can_ip_forward' to be a bool")
        __self__.can_ip_forward = can_ip_forward
        """
        Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
        """
        if cpu_platform and not isinstance(cpu_platform, str):
            raise TypeError("Expected argument 'cpu_platform' to be a str")
        __self__.cpu_platform = cpu_platform
        """
        The CPU platform used by this instance.
        """
        if current_status and not isinstance(current_status, str):
            raise TypeError("Expected argument 'current_status' to be a str")
        __self__.current_status = current_status
        if deletion_protection and not isinstance(deletion_protection, bool):
            raise TypeError("Expected argument 'deletion_protection' to be a bool")
        __self__.deletion_protection = deletion_protection
        """
        Whether deletion protection is enabled on this instance.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        A brief description of the resource.
        """
        if desired_status and not isinstance(desired_status, str):
            raise TypeError("Expected argument 'desired_status' to be a str")
        __self__.desired_status = desired_status
        if enable_display and not isinstance(enable_display, bool):
            raise TypeError("Expected argument 'enable_display' to be a bool")
        __self__.enable_display = enable_display
        if guest_accelerators and not isinstance(guest_accelerators, list):
            raise TypeError("Expected argument 'guest_accelerators' to be a list")
        __self__.guest_accelerators = guest_accelerators
        """
        List of the type and count of accelerator cards attached to the instance. Structure is documented below.
        """
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        __self__.hostname = hostname
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        __self__.instance_id = instance_id
        """
        The server-assigned unique identifier of this instance.
        """
        if label_fingerprint and not isinstance(label_fingerprint, str):
            raise TypeError("Expected argument 'label_fingerprint' to be a str")
        __self__.label_fingerprint = label_fingerprint
        """
        The unique fingerprint of the labels.
        """
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        __self__.labels = labels
        """
        A set of key/value label pairs assigned to the instance.
        """
        if machine_type and not isinstance(machine_type, str):
            raise TypeError("Expected argument 'machine_type' to be a str")
        __self__.machine_type = machine_type
        """
        The machine type to create.
        """
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        __self__.metadata = metadata
        """
        Metadata key/value pairs made available within the instance.
        """
        if metadata_fingerprint and not isinstance(metadata_fingerprint, str):
            raise TypeError("Expected argument 'metadata_fingerprint' to be a str")
        __self__.metadata_fingerprint = metadata_fingerprint
        """
        The unique fingerprint of the metadata.
        """
        if metadata_startup_script and not isinstance(metadata_startup_script, str):
            raise TypeError("Expected argument 'metadata_startup_script' to be a str")
        __self__.metadata_startup_script = metadata_startup_script
        if min_cpu_platform and not isinstance(min_cpu_platform, str):
            raise TypeError("Expected argument 'min_cpu_platform' to be a str")
        __self__.min_cpu_platform = min_cpu_platform
        """
        The minimum CPU platform specified for the VM instance.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if network_interfaces and not isinstance(network_interfaces, list):
            raise TypeError("Expected argument 'network_interfaces' to be a list")
        __self__.network_interfaces = network_interfaces
        """
        The networks attached to the instance. Structure is documented below.
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if resource_policies and not isinstance(resource_policies, list):
            raise TypeError("Expected argument 'resource_policies' to be a list")
        __self__.resource_policies = resource_policies
        if schedulings and not isinstance(schedulings, list):
            raise TypeError("Expected argument 'schedulings' to be a list")
        __self__.schedulings = schedulings
        """
        The scheduling strategy being used by the instance.
        """
        if scratch_disks and not isinstance(scratch_disks, list):
            raise TypeError("Expected argument 'scratch_disks' to be a list")
        __self__.scratch_disks = scratch_disks
        """
        The scratch disks attached to the instance. Structure is documented below.
        """
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        __self__.self_link = self_link
        """
        The URI of the created resource.
        """
        if service_accounts and not isinstance(service_accounts, list):
            raise TypeError("Expected argument 'service_accounts' to be a list")
        __self__.service_accounts = service_accounts
        """
        The service account to attach to the instance. Structure is documented below.
        """
        if shielded_instance_configs and not isinstance(shielded_instance_configs, list):
            raise TypeError("Expected argument 'shielded_instance_configs' to be a list")
        __self__.shielded_instance_configs = shielded_instance_configs
        """
        The shielded vm config being used by the instance. Structure is documented below.
        """
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        __self__.tags = tags
        """
        The list of tags attached to the instance.
        """
        if tags_fingerprint and not isinstance(tags_fingerprint, str):
            raise TypeError("Expected argument 'tags_fingerprint' to be a str")
        __self__.tags_fingerprint = tags_fingerprint
        """
        The unique fingerprint of the tags.
        """
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        __self__.zone = zone


class AwaitableGetInstanceResult(GetInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceResult(
            allow_stopping_for_update=self.allow_stopping_for_update,
            attached_disks=self.attached_disks,
            boot_disks=self.boot_disks,
            can_ip_forward=self.can_ip_forward,
            cpu_platform=self.cpu_platform,
            current_status=self.current_status,
            deletion_protection=self.deletion_protection,
            description=self.description,
            desired_status=self.desired_status,
            enable_display=self.enable_display,
            guest_accelerators=self.guest_accelerators,
            hostname=self.hostname,
            id=self.id,
            instance_id=self.instance_id,
            label_fingerprint=self.label_fingerprint,
            labels=self.labels,
            machine_type=self.machine_type,
            metadata=self.metadata,
            metadata_fingerprint=self.metadata_fingerprint,
            metadata_startup_script=self.metadata_startup_script,
            min_cpu_platform=self.min_cpu_platform,
            name=self.name,
            network_interfaces=self.network_interfaces,
            project=self.project,
            resource_policies=self.resource_policies,
            schedulings=self.schedulings,
            scratch_disks=self.scratch_disks,
            self_link=self.self_link,
            service_accounts=self.service_accounts,
            shielded_instance_configs=self.shielded_instance_configs,
            tags=self.tags,
            tags_fingerprint=self.tags_fingerprint,
            zone=self.zone)


def get_instance(name=None, project=None, self_link=None, zone=None, opts=None):
    """
    Get information about a VM instance resource within GCE. For more information see
    [the official documentation](https://cloud.google.com/compute/docs/instances)
    and
    [API](https://cloud.google.com/compute/docs/reference/latest/instances).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    appserver = gcp.compute.get_instance(name="primary-application-server",
        zone="us-central1-a")
    ```


    :param str name: The name of the instance. One of `name` or `self_link` must be provided.
    :param str project: The ID of the project in which the resource belongs.
           If `self_link` is provided, this value is ignored.  If neither `self_link`
           nor `project` are provided, the provider project is used.
    :param str self_link: The self link of the instance. One of `name` or `self_link` must be provided.
    :param str zone: The zone of the instance. If `self_link` is provided, this
           value is ignored.  If neither `self_link` nor `zone` are provided, the
           provider zone is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    __args__['selfLink'] = self_link
    __args__['zone'] = zone
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getInstance:getInstance', __args__, opts=opts).value

    return AwaitableGetInstanceResult(
        allow_stopping_for_update=__ret__.get('allowStoppingForUpdate'),
        attached_disks=__ret__.get('attachedDisks'),
        boot_disks=__ret__.get('bootDisks'),
        can_ip_forward=__ret__.get('canIpForward'),
        cpu_platform=__ret__.get('cpuPlatform'),
        current_status=__ret__.get('currentStatus'),
        deletion_protection=__ret__.get('deletionProtection'),
        description=__ret__.get('description'),
        desired_status=__ret__.get('desiredStatus'),
        enable_display=__ret__.get('enableDisplay'),
        guest_accelerators=__ret__.get('guestAccelerators'),
        hostname=__ret__.get('hostname'),
        id=__ret__.get('id'),
        instance_id=__ret__.get('instanceId'),
        label_fingerprint=__ret__.get('labelFingerprint'),
        labels=__ret__.get('labels'),
        machine_type=__ret__.get('machineType'),
        metadata=__ret__.get('metadata'),
        metadata_fingerprint=__ret__.get('metadataFingerprint'),
        metadata_startup_script=__ret__.get('metadataStartupScript'),
        min_cpu_platform=__ret__.get('minCpuPlatform'),
        name=__ret__.get('name'),
        network_interfaces=__ret__.get('networkInterfaces'),
        project=__ret__.get('project'),
        resource_policies=__ret__.get('resourcePolicies'),
        schedulings=__ret__.get('schedulings'),
        scratch_disks=__ret__.get('scratchDisks'),
        self_link=__ret__.get('selfLink'),
        service_accounts=__ret__.get('serviceAccounts'),
        shielded_instance_configs=__ret__.get('shieldedInstanceConfigs'),
        tags=__ret__.get('tags'),
        tags_fingerprint=__ret__.get('tagsFingerprint'),
        zone=__ret__.get('zone'))
