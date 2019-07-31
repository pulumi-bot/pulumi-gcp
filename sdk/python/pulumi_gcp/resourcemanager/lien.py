# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Lien(pulumi.CustomResource):
    create_time: pulumi.Output[str]
    name: pulumi.Output[str]
    origin: pulumi.Output[str]
    parent: pulumi.Output[str]
    reason: pulumi.Output[str]
    restrictions: pulumi.Output[list]
    def __init__(__self__, resource_name, opts=None, origin=None, parent=None, reason=None, restrictions=None, __name__=None, __opts__=None):
        """
        A Lien represents an encumbrance on the actions that can be performed on a resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/resource_manager_lien.html.markdown.
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

        if origin is None:
            raise TypeError("Missing required property 'origin'")
        __props__['origin'] = origin

        if parent is None:
            raise TypeError("Missing required property 'parent'")
        __props__['parent'] = parent

        if reason is None:
            raise TypeError("Missing required property 'reason'")
        __props__['reason'] = reason

        if restrictions is None:
            raise TypeError("Missing required property 'restrictions'")
        __props__['restrictions'] = restrictions

        __props__['create_time'] = None
        __props__['name'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Lien, __self__).__init__(
            'gcp:resourcemanager/lien:Lien',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

