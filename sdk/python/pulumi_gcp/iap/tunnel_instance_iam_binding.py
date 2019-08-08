# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class TunnelInstanceIAMBinding(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the instance's IAM policy.
    """
    instance: pulumi.Output[str]
    """
    The name of the instance.
    """
    members: pulumi.Output[list]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    role: pulumi.Output[str]
    """
    The role that should be applied. Only one
    `iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    zone: pulumi.Output[str]
    """
    The zone of the instance. If
    unspecified, this defaults to the zone configured in the provider.
    """
    def __init__(__self__, resource_name, opts=None, instance=None, members=None, project=None, role=None, zone=None, __name__=None, __opts__=None):
        """
        Create a TunnelInstanceIAMBinding resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance: The name of the instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] zone: The zone of the instance. If
               unspecified, this defaults to the zone configured in the provider.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/iap_tunnel_instance_iam_binding.html.markdown.
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

        if instance is None:
            raise TypeError("Missing required property 'instance'")
        __props__['instance'] = instance
        if members is None:
            raise TypeError("Missing required property 'members'")
        __props__['members'] = members
        __props__['project'] = project
        if role is None:
            raise TypeError("Missing required property 'role'")
        __props__['role'] = role
        __props__['zone'] = zone
        __props__['etag'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(TunnelInstanceIAMBinding, __self__).__init__(
            'gcp:iap/tunnelInstanceIAMBinding:TunnelInstanceIAMBinding',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

