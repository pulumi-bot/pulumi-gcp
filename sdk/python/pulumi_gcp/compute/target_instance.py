# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class TargetInstance(pulumi.CustomResource):
    creation_timestamp: pulumi.Output[str]
    """
    Creation timestamp in RFC3339 text format.
    """
    description: pulumi.Output[str]
    """
    An optional description of this resource.
    """
    instance: pulumi.Output[str]
    """
    The Compute instance VM handling traffic for this target instance.
    Accepts the instance self-link, relative path
    (e.g. `projects/project/zones/zone/instances/instance`) or name. If
    name is given, the zone will default to the given zone or
    the provider-default zone and the project will default to the
    provider-level project.
    """
    name: pulumi.Output[str]
    """
    Name of the resource. Provided by the client when the resource is
    created. The name must be 1-63 characters long, and comply with
    RFC1035. Specifically, the name must be 1-63 characters long and match
    the regular expression `a-z?` which means the
    first character must be a lowercase letter, and all following
    characters must be a dash, lowercase letter, or digit, except the last
    character, which cannot be a dash.
    """
    nat_policy: pulumi.Output[str]
    """
    NAT option controlling how IPs are NAT'ed to the instance.
    Currently only NO_NAT (default value) is supported.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    zone: pulumi.Output[str]
    """
    URL of the zone where the target instance resides.
    """
    def __init__(__self__, resource_name, opts=None, description=None, instance=None, name=None, nat_policy=None, project=None, zone=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents a TargetInstance resource which defines an endpoint instance
        that terminates traffic of certain protocols. In particular, they are used
        in Protocol Forwarding, where forwarding rules can send packets to a
        non-NAT'ed target instance. Each target instance contains a single
        virtual machine instance that receives and handles traffic from the
        corresponding forwarding rules.


        To get more information about TargetInstance, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetInstances)
        * How-to Guides
            * [Using Protocol Forwarding](https://cloud.google.com/compute/docs/protocol-forwarding)

        ## Example Usage

        ### Target Instance Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        vmimage = gcp.compute.get_image(family="debian-9",
            project="debian-cloud")
        target_vm = gcp.compute.Instance("target-vm",
            machine_type="n1-standard-1",
            zone="us-central1-a",
            boot_disk={
                "initialize_params": {
                    "image": vmimage.self_link,
                },
            },
            network_interface=[{
                "network": "default",
            }])
        default = gcp.compute.TargetInstance("default", instance=target_vm.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] instance: The Compute instance VM handling traffic for this target instance.
               Accepts the instance self-link, relative path
               (e.g. `projects/project/zones/zone/instances/instance`) or name. If
               name is given, the zone will default to the given zone or
               the provider-default zone and the project will default to the
               provider-level project.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] nat_policy: NAT option controlling how IPs are NAT'ed to the instance.
               Currently only NO_NAT (default value) is supported.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] zone: URL of the zone where the target instance resides.
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
            if instance is None:
                raise TypeError("Missing required property 'instance'")
            __props__['instance'] = instance
            __props__['name'] = name
            __props__['nat_policy'] = nat_policy
            __props__['project'] = project
            __props__['zone'] = zone
            __props__['creation_timestamp'] = None
            __props__['self_link'] = None
        super(TargetInstance, __self__).__init__(
            'gcp:compute/targetInstance:TargetInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, creation_timestamp=None, description=None, instance=None, name=None, nat_policy=None, project=None, self_link=None, zone=None):
        """
        Get an existing TargetInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] instance: The Compute instance VM handling traffic for this target instance.
               Accepts the instance self-link, relative path
               (e.g. `projects/project/zones/zone/instances/instance`) or name. If
               name is given, the zone will default to the given zone or
               the provider-default zone and the project will default to the
               provider-level project.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] nat_policy: NAT option controlling how IPs are NAT'ed to the instance.
               Currently only NO_NAT (default value) is supported.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] zone: URL of the zone where the target instance resides.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["instance"] = instance
        __props__["name"] = name
        __props__["nat_policy"] = nat_policy
        __props__["project"] = project
        __props__["self_link"] = self_link
        __props__["zone"] = zone
        return TargetInstance(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

