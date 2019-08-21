# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AccessPolicy(pulumi.CustomResource):
    create_time: pulumi.Output[str]
    name: pulumi.Output[str]
    parent: pulumi.Output[str]
    title: pulumi.Output[str]
    update_time: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, parent=None, title=None, __props__=None, __name__=None, __opts__=None):
        """
        AccessPolicy is a container for AccessLevels (which define the necessary
        attributes to use GCP services) and ServicePerimeters (which define
        regions of services able to freely pass data within a perimeter). An
        access policy is globally visible within an organization, and the
        restrictions it specifies apply to all projects within an organization.
        
        
        To get more information about AccessPolicy, see:
        
        * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies)
        * How-to Guides
            * [Access Policy Quickstart](https://cloud.google.com/access-context-manager/docs/quickstart)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/access_context_manager_access_policy.html.markdown.
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

            if parent is None:
                raise TypeError("Missing required property 'parent'")
            __props__['parent'] = parent
            if title is None:
                raise TypeError("Missing required property 'title'")
            __props__['title'] = title
            __props__['create_time'] = None
            __props__['name'] = None
            __props__['update_time'] = None
        super(AccessPolicy, __self__).__init__(
            'gcp:accesscontextmanager/accessPolicy:AccessPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, create_time=None, name=None, parent=None, title=None, update_time=None):
        """
        Get an existing AccessPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/access_context_manager_access_policy.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["create_time"] = create_time
        __props__["name"] = name
        __props__["parent"] = parent
        __props__["title"] = title
        __props__["update_time"] = update_time
        return AccessPolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

