# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class EntryGroup(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
    """
    display_name: pulumi.Output[str]
    """
    A short name to identify the entry group, for example, "analytics data - jan 2011".
    """
    entry_group_id: pulumi.Output[str]
    """
    The id of the entry group to create. The id must begin with a letter or underscore,
    contain only English letters, numbers and underscores, and be at most 64 characters.
    """
    name: pulumi.Output[str]
    """
    The resource name of the entry group in URL format. Example:
    projects/{project}/locations/{location}/entryGroups/{entryGroupId}
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    EntryGroup location region.
    """
    def __init__(__self__, resource_name, opts=None, description=None, display_name=None, entry_group_id=None, project=None, region=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a EntryGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
        :param pulumi.Input[str] display_name: A short name to identify the entry group, for example, "analytics data - jan 2011".
        :param pulumi.Input[str] entry_group_id: The id of the entry group to create. The id must begin with a letter or underscore,
               contain only English letters, numbers and underscores, and be at most 64 characters.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: EntryGroup location region.
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
            __props__['display_name'] = display_name
            if entry_group_id is None:
                raise TypeError("Missing required property 'entry_group_id'")
            __props__['entry_group_id'] = entry_group_id
            __props__['project'] = project
            __props__['region'] = region
            __props__['name'] = None
        super(EntryGroup, __self__).__init__(
            'gcp:datacatalog/entryGroup:EntryGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, display_name=None, entry_group_id=None, name=None, project=None, region=None):
        """
        Get an existing EntryGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
        :param pulumi.Input[str] display_name: A short name to identify the entry group, for example, "analytics data - jan 2011".
        :param pulumi.Input[str] entry_group_id: The id of the entry group to create. The id must begin with a letter or underscore,
               contain only English letters, numbers and underscores, and be at most 64 characters.
        :param pulumi.Input[str] name: The resource name of the entry group in URL format. Example:
               projects/{project}/locations/{location}/entryGroups/{entryGroupId}
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: EntryGroup location region.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["entry_group_id"] = entry_group_id
        __props__["name"] = name
        __props__["project"] = project
        __props__["region"] = region
        return EntryGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
