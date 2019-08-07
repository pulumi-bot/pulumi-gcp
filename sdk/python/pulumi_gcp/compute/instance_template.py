# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class InstanceTemplate(pulumi.CustomResource):
    can_ip_forward: pulumi.Output[bool]
    """
    Whether to allow sending and receiving of
    packets with non-matching source or destination IPs. This defaults to false.
    """
    description: pulumi.Output[str]
    """
    A brief description of this resource.
    """
    disks: pulumi.Output[list]
    """
    Disks to attach to instances created from this template.
    This can be specified multiple times for multiple disks. Structure is
    documented below.
    """
    guest_accelerators: pulumi.Output[list]
    """
    List of the type and count of accelerator cards attached to the instance. Structure documented below.
    """
    instance_description: pulumi.Output[str]
    """
    A brief description to use for instances
    created from this template.
    """
    labels: pulumi.Output[dict]
    """
    A set of key/value label pairs to assign to instances
    created from this template,
    """
    machine_type: pulumi.Output[str]
    """
    The machine type to create.
    """
    metadata: pulumi.Output[dict]
    """
    Metadata key/value pairs to make available from
    within instances created from this template.
    """
    metadata_fingerprint: pulumi.Output[str]
    """
    The unique fingerprint of the metadata.
    """
    metadata_startup_script: pulumi.Output[str]
    """
    An alternative to using the
    startup-script metadata key, mostly to match the compute_instance resource.
    This replaces the startup-script metadata key on the created instance and
    thus the two mechanisms are not allowed to be used simultaneously.
    """
    min_cpu_platform: pulumi.Output[str]
    """
    Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
    `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
    """
    name: pulumi.Output[str]
    """
    The name of the instance template. If you leave
    this blank, this provider will auto-generate a unique name.
    """
    name_prefix: pulumi.Output[str]
    """
    Creates a unique name beginning with the specified
    prefix. Conflicts with `name`.
    """
    network_interfaces: pulumi.Output[list]
    """
    Networks to attach to instances created from
    this template. This can be specified multiple times for multiple networks.
    Structure is documented below.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    An instance template is a global resource that is not
    bound to a zone or a region. However, you can still specify some regional
    resources in an instance template, which restricts the template to the
    region where that resource resides. For example, a custom `subnetwork`
    resource is tied to a specific region. Defaults to the region of the
    Provider if no value is given.
    """
    scheduling: pulumi.Output[dict]
    """
    The scheduling strategy to use. More details about
    this configuration option are detailed below.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    service_account: pulumi.Output[dict]
    """
    Service account to attach to the instance. Structure is documented below.
    """
    shielded_instance_config: pulumi.Output[dict]
    """
    Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
    **Note**: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
    """
    tags: pulumi.Output[list]
    """
    Tags to attach to the instance.
    """
    tags_fingerprint: pulumi.Output[str]
    """
    The unique fingerprint of the tags.
    """
    def __init__(__self__, resource_name, opts=None, can_ip_forward=None, description=None, disks=None, guest_accelerators=None, instance_description=None, labels=None, machine_type=None, metadata=None, metadata_startup_script=None, min_cpu_platform=None, name=None, name_prefix=None, network_interfaces=None, project=None, region=None, scheduling=None, service_account=None, shielded_instance_config=None, tags=None, __name__=None, __opts__=None):
        """
        Manages a VM instance template resource within GCE. For more information see
        [the official documentation](https://cloud.google.com/compute/docs/instance-templates)
        and
        [API](https://cloud.google.com/compute/docs/reference/latest/instanceTemplates).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] can_ip_forward: Whether to allow sending and receiving of
               packets with non-matching source or destination IPs. This defaults to false.
        :param pulumi.Input[str] description: A brief description of this resource.
        :param pulumi.Input[list] disks: Disks to attach to instances created from this template.
               This can be specified multiple times for multiple disks. Structure is
               documented below.
        :param pulumi.Input[list] guest_accelerators: List of the type and count of accelerator cards attached to the instance. Structure documented below.
        :param pulumi.Input[str] instance_description: A brief description to use for instances
               created from this template.
        :param pulumi.Input[dict] labels: A set of key/value label pairs to assign to instances
               created from this template,
        :param pulumi.Input[str] machine_type: The machine type to create.
        :param pulumi.Input[dict] metadata: Metadata key/value pairs to make available from
               within instances created from this template.
        :param pulumi.Input[str] metadata_startup_script: An alternative to using the
               startup-script metadata key, mostly to match the compute_instance resource.
               This replaces the startup-script metadata key on the created instance and
               thus the two mechanisms are not allowed to be used simultaneously.
        :param pulumi.Input[str] min_cpu_platform: Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
               `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
        :param pulumi.Input[str] name: The name of the instance template. If you leave
               this blank, this provider will auto-generate a unique name.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified
               prefix. Conflicts with `name`.
        :param pulumi.Input[list] network_interfaces: Networks to attach to instances created from
               this template. This can be specified multiple times for multiple networks.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] region: An instance template is a global resource that is not
               bound to a zone or a region. However, you can still specify some regional
               resources in an instance template, which restricts the template to the
               region where that resource resides. For example, a custom `subnetwork`
               resource is tied to a specific region. Defaults to the region of the
               Provider if no value is given.
        :param pulumi.Input[dict] scheduling: The scheduling strategy to use. More details about
               this configuration option are detailed below.
        :param pulumi.Input[dict] service_account: Service account to attach to the instance. Structure is documented below.
        :param pulumi.Input[dict] shielded_instance_config: Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
               **Note**: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
        :param pulumi.Input[list] tags: Tags to attach to the instance.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_instance_template.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['can_ip_forward'] = can_ip_forward
        __props__['description'] = description
        if disks is None:
            raise TypeError("Missing required property 'disks'")
        __props__['disks'] = disks
        __props__['guest_accelerators'] = guest_accelerators
        __props__['instance_description'] = instance_description
        __props__['labels'] = labels
        if machine_type is None:
            raise TypeError("Missing required property 'machine_type'")
        __props__['machine_type'] = machine_type
        __props__['metadata'] = metadata
        __props__['metadata_startup_script'] = metadata_startup_script
        __props__['min_cpu_platform'] = min_cpu_platform
        __props__['name'] = name
        __props__['name_prefix'] = name_prefix
        __props__['network_interfaces'] = network_interfaces
        __props__['project'] = project
        __props__['region'] = region
        __props__['scheduling'] = scheduling
        __props__['service_account'] = service_account
        __props__['shielded_instance_config'] = shielded_instance_config
        __props__['tags'] = tags
        __props__['metadata_fingerprint'] = None
        __props__['self_link'] = None
        __props__['tags_fingerprint'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(InstanceTemplate, __self__).__init__(
            'gcp:compute/instanceTemplate:InstanceTemplate',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

