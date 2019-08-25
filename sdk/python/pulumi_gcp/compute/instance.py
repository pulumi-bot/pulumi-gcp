# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Instance(pulumi.CustomResource):
    allow_stopping_for_update: pulumi.Output[bool]
    """
    If true, allows this provider to stop the instance to update its properties.
    If you try to update a property that requires stopping the instance without setting this field, the update will fail.
    """
    attached_disks: pulumi.Output[list]
    """
    Additional disks to attach to the instance. Can be repeated multiple times for multiple disks. Structure is documented below.
    """
    boot_disk: pulumi.Output[dict]
    """
    The boot disk for the instance.
    Structure is documented below.
    """
    can_ip_forward: pulumi.Output[bool]
    """
    Whether to allow sending and receiving of
    packets with non-matching source or destination IPs.
    This defaults to false.
    """
    cpu_platform: pulumi.Output[str]
    """
    The CPU platform used by this instance.
    """
    deletion_protection: pulumi.Output[bool]
    """
    Enable deletion protection on this instance. Defaults to false.
    **Note:** you must disable deletion protection before removing the resource, or the instance cannot be deleted and the deployment will not complete successfully.
    """
    description: pulumi.Output[str]
    """
    A brief description of this resource.
    """
    guest_accelerators: pulumi.Output[list]
    """
    List of the type and count of accelerator cards attached to the instance. Structure documented below.
    **Note:** GPU accelerators can only be used with `on_host_maintenance` option set to TERMINATE.
    """
    hostname: pulumi.Output[str]
    """
    A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid.
    Valid format is a series of labels 1-63 characters long matching the regular expression `a-z`, concatenated with periods.
    The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
    """
    instance_id: pulumi.Output[str]
    """
    The server-assigned unique identifier of this instance.
    """
    label_fingerprint: pulumi.Output[str]
    """
    The unique fingerprint of the labels.
    """
    labels: pulumi.Output[dict]
    """
    A set of key/value label pairs to assign to the instance.
    """
    machine_type: pulumi.Output[str]
    """
    The machine type to create.
    """
    metadata: pulumi.Output[dict]
    """
    Metadata key/value pairs to make available from
    within the instance. Ssh keys attached in the Cloud Console will be removed.
    Add them to your config in order to keep them attached to your instance.
    """
    metadata_fingerprint: pulumi.Output[str]
    """
    The unique fingerprint of the metadata.
    """
    metadata_startup_script: pulumi.Output[str]
    """
    An alternative to using the
    startup-script metadata key, except this one forces the instance to be
    recreated (thus re-running the script) if it is changed. This replaces the
    startup-script metadata key on the created instance and thus the two
    mechanisms are not allowed to be used simultaneously.
    """
    min_cpu_platform: pulumi.Output[str]
    """
    Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as
    `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
    **Note**: `allow_stopping_for_update` must be set to true in order to update this field.
    """
    name: pulumi.Output[str]
    """
    A unique name for the resource, required by GCE.
    Changing this forces a new resource to be created.
    """
    network_interfaces: pulumi.Output[list]
    """
    Networks to attach to the instance. This can
    be specified multiple times. Structure is documented below.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    scheduling: pulumi.Output[dict]
    """
    The scheduling strategy to use. More details about
    this configuration option are detailed below.
    """
    scratch_disks: pulumi.Output[list]
    """
    Scratch disks to attach to the instance. This can be
    specified multiple times for multiple scratch disks. Structure is documented below.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    service_account: pulumi.Output[dict]
    """
    Service account to attach to the instance.
    Structure is documented below.
    **Note**: `allow_stopping_for_update` must be set to true in order to update this field.
    """
    shielded_instance_config: pulumi.Output[dict]
    """
    Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
    **Note**: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
    """
    tags: pulumi.Output[list]
    """
    A list of tags to attach to the instance.
    """
    tags_fingerprint: pulumi.Output[str]
    """
    The unique fingerprint of the tags.
    """
    zone: pulumi.Output[str]
    """
    The zone that the machine should be created in.
    """
    def __init__(__self__, resource_name, opts=None, allow_stopping_for_update=None, attached_disks=None, boot_disk=None, can_ip_forward=None, deletion_protection=None, description=None, guest_accelerators=None, hostname=None, labels=None, machine_type=None, metadata=None, metadata_startup_script=None, min_cpu_platform=None, name=None, network_interfaces=None, project=None, scheduling=None, scratch_disks=None, service_account=None, shielded_instance_config=None, tags=None, zone=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a VM instance resource within GCE. For more information see
        [the official documentation](https://cloud.google.com/compute/docs/instances)
        and
        [API](https://cloud.google.com/compute/docs/reference/latest/instances).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_stopping_for_update: If true, allows this provider to stop the instance to update its properties.
               If you try to update a property that requires stopping the instance without setting this field, the update will fail.
        :param pulumi.Input[list] attached_disks: Additional disks to attach to the instance. Can be repeated multiple times for multiple disks. Structure is documented below.
        :param pulumi.Input[dict] boot_disk: The boot disk for the instance.
               Structure is documented below.
        :param pulumi.Input[bool] can_ip_forward: Whether to allow sending and receiving of
               packets with non-matching source or destination IPs.
               This defaults to false.
        :param pulumi.Input[bool] deletion_protection: Enable deletion protection on this instance. Defaults to false.
               **Note:** you must disable deletion protection before removing the resource, or the instance cannot be deleted and the deployment will not complete successfully.
        :param pulumi.Input[str] description: A brief description of this resource.
        :param pulumi.Input[list] guest_accelerators: List of the type and count of accelerator cards attached to the instance. Structure documented below.
               **Note:** GPU accelerators can only be used with `on_host_maintenance` option set to TERMINATE.
        :param pulumi.Input[str] hostname: A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid.
               Valid format is a series of labels 1-63 characters long matching the regular expression `a-z`, concatenated with periods.
               The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] labels: A set of key/value label pairs to assign to the instance.
        :param pulumi.Input[str] machine_type: The machine type to create.
        :param pulumi.Input[dict] metadata: Metadata key/value pairs to make available from
               within the instance. Ssh keys attached in the Cloud Console will be removed.
               Add them to your config in order to keep them attached to your instance.
        :param pulumi.Input[str] metadata_startup_script: An alternative to using the
               startup-script metadata key, except this one forces the instance to be
               recreated (thus re-running the script) if it is changed. This replaces the
               startup-script metadata key on the created instance and thus the two
               mechanisms are not allowed to be used simultaneously.
        :param pulumi.Input[str] min_cpu_platform: Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as
               `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
               **Note**: `allow_stopping_for_update` must be set to true in order to update this field.
        :param pulumi.Input[str] name: A unique name for the resource, required by GCE.
               Changing this forces a new resource to be created.
        :param pulumi.Input[list] network_interfaces: Networks to attach to the instance. This can
               be specified multiple times. Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[dict] scheduling: The scheduling strategy to use. More details about
               this configuration option are detailed below.
        :param pulumi.Input[list] scratch_disks: Scratch disks to attach to the instance. This can be
               specified multiple times for multiple scratch disks. Structure is documented below.
        :param pulumi.Input[dict] service_account: Service account to attach to the instance.
               Structure is documented below.
               **Note**: `allow_stopping_for_update` must be set to true in order to update this field.
        :param pulumi.Input[dict] shielded_instance_config: Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
               **Note**: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
        :param pulumi.Input[list] tags: A list of tags to attach to the instance.
        :param pulumi.Input[str] zone: The zone that the machine should be created in.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_instance.html.markdown.
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

            __props__['allow_stopping_for_update'] = allow_stopping_for_update
            __props__['attached_disks'] = attached_disks
            if boot_disk is None:
                raise TypeError("Missing required property 'boot_disk'")
            __props__['boot_disk'] = boot_disk
            __props__['can_ip_forward'] = can_ip_forward
            __props__['deletion_protection'] = deletion_protection
            __props__['description'] = description
            __props__['guest_accelerators'] = guest_accelerators
            __props__['hostname'] = hostname
            __props__['labels'] = labels
            if machine_type is None:
                raise TypeError("Missing required property 'machine_type'")
            __props__['machine_type'] = machine_type
            __props__['metadata'] = metadata
            __props__['metadata_startup_script'] = metadata_startup_script
            __props__['min_cpu_platform'] = min_cpu_platform
            __props__['name'] = name
            if network_interfaces is None:
                raise TypeError("Missing required property 'network_interfaces'")
            __props__['network_interfaces'] = network_interfaces
            __props__['project'] = project
            __props__['scheduling'] = scheduling
            __props__['scratch_disks'] = scratch_disks
            __props__['service_account'] = service_account
            __props__['shielded_instance_config'] = shielded_instance_config
            __props__['tags'] = tags
            __props__['zone'] = zone
            __props__['cpu_platform'] = None
            __props__['instance_id'] = None
            __props__['label_fingerprint'] = None
            __props__['metadata_fingerprint'] = None
            __props__['self_link'] = None
            __props__['tags_fingerprint'] = None
        super(Instance, __self__).__init__(
            'gcp:compute/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, allow_stopping_for_update=None, attached_disks=None, boot_disk=None, can_ip_forward=None, cpu_platform=None, deletion_protection=None, description=None, guest_accelerators=None, hostname=None, instance_id=None, label_fingerprint=None, labels=None, machine_type=None, metadata=None, metadata_fingerprint=None, metadata_startup_script=None, min_cpu_platform=None, name=None, network_interfaces=None, project=None, scheduling=None, scratch_disks=None, self_link=None, service_account=None, shielded_instance_config=None, tags=None, tags_fingerprint=None, zone=None):
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_stopping_for_update: If true, allows this provider to stop the instance to update its properties.
               If you try to update a property that requires stopping the instance without setting this field, the update will fail.
        :param pulumi.Input[list] attached_disks: Additional disks to attach to the instance. Can be repeated multiple times for multiple disks. Structure is documented below.
        :param pulumi.Input[dict] boot_disk: The boot disk for the instance.
               Structure is documented below.
        :param pulumi.Input[bool] can_ip_forward: Whether to allow sending and receiving of
               packets with non-matching source or destination IPs.
               This defaults to false.
        :param pulumi.Input[str] cpu_platform: The CPU platform used by this instance.
        :param pulumi.Input[bool] deletion_protection: Enable deletion protection on this instance. Defaults to false.
               **Note:** you must disable deletion protection before removing the resource, or the instance cannot be deleted and the deployment will not complete successfully.
        :param pulumi.Input[str] description: A brief description of this resource.
        :param pulumi.Input[list] guest_accelerators: List of the type and count of accelerator cards attached to the instance. Structure documented below.
               **Note:** GPU accelerators can only be used with `on_host_maintenance` option set to TERMINATE.
        :param pulumi.Input[str] hostname: A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid.
               Valid format is a series of labels 1-63 characters long matching the regular expression `a-z`, concatenated with periods.
               The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
        :param pulumi.Input[str] instance_id: The server-assigned unique identifier of this instance.
        :param pulumi.Input[str] label_fingerprint: The unique fingerprint of the labels.
        :param pulumi.Input[dict] labels: A set of key/value label pairs to assign to the instance.
        :param pulumi.Input[str] machine_type: The machine type to create.
        :param pulumi.Input[dict] metadata: Metadata key/value pairs to make available from
               within the instance. Ssh keys attached in the Cloud Console will be removed.
               Add them to your config in order to keep them attached to your instance.
        :param pulumi.Input[str] metadata_fingerprint: The unique fingerprint of the metadata.
        :param pulumi.Input[str] metadata_startup_script: An alternative to using the
               startup-script metadata key, except this one forces the instance to be
               recreated (thus re-running the script) if it is changed. This replaces the
               startup-script metadata key on the created instance and thus the two
               mechanisms are not allowed to be used simultaneously.
        :param pulumi.Input[str] min_cpu_platform: Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as
               `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
               **Note**: `allow_stopping_for_update` must be set to true in order to update this field.
        :param pulumi.Input[str] name: A unique name for the resource, required by GCE.
               Changing this forces a new resource to be created.
        :param pulumi.Input[list] network_interfaces: Networks to attach to the instance. This can
               be specified multiple times. Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[dict] scheduling: The scheduling strategy to use. More details about
               this configuration option are detailed below.
        :param pulumi.Input[list] scratch_disks: Scratch disks to attach to the instance. This can be
               specified multiple times for multiple scratch disks. Structure is documented below.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[dict] service_account: Service account to attach to the instance.
               Structure is documented below.
               **Note**: `allow_stopping_for_update` must be set to true in order to update this field.
        :param pulumi.Input[dict] shielded_instance_config: Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
               **Note**: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
        :param pulumi.Input[list] tags: A list of tags to attach to the instance.
        :param pulumi.Input[str] tags_fingerprint: The unique fingerprint of the tags.
        :param pulumi.Input[str] zone: The zone that the machine should be created in.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_instance.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["allow_stopping_for_update"] = allow_stopping_for_update
        __props__["attached_disks"] = attached_disks
        __props__["boot_disk"] = boot_disk
        __props__["can_ip_forward"] = can_ip_forward
        __props__["cpu_platform"] = cpu_platform
        __props__["deletion_protection"] = deletion_protection
        __props__["description"] = description
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
        __props__["scheduling"] = scheduling
        __props__["scratch_disks"] = scratch_disks
        __props__["self_link"] = self_link
        __props__["service_account"] = service_account
        __props__["shielded_instance_config"] = shielded_instance_config
        __props__["tags"] = tags
        __props__["tags_fingerprint"] = tags_fingerprint
        __props__["zone"] = zone
        return Instance(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

