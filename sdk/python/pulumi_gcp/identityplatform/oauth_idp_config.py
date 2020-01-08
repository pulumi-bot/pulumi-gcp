# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class OauthIdpConfig(pulumi.CustomResource):
    client_id: pulumi.Output[str]
    client_secret: pulumi.Output[str]
    display_name: pulumi.Output[str]
    enabled: pulumi.Output[bool]
    issuer: pulumi.Output[str]
    name: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    def __init__(__self__, resource_name, opts=None, client_id=None, client_secret=None, display_name=None, enabled=None, issuer=None, name=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a OauthIdpConfig resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/identity_platform_oauth_idp_config.html.markdown.
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

            if client_id is None:
                raise TypeError("Missing required property 'client_id'")
            __props__['client_id'] = client_id
            __props__['client_secret'] = client_secret
            __props__['display_name'] = display_name
            __props__['enabled'] = enabled
            if issuer is None:
                raise TypeError("Missing required property 'issuer'")
            __props__['issuer'] = issuer
            __props__['name'] = name
            __props__['project'] = project
        super(OauthIdpConfig, __self__).__init__(
            'gcp:identityplatform/oauthIdpConfig:OauthIdpConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, client_id=None, client_secret=None, display_name=None, enabled=None, issuer=None, name=None, project=None):
        """
        Get an existing OauthIdpConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/identity_platform_oauth_idp_config.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["client_id"] = client_id
        __props__["client_secret"] = client_secret
        __props__["display_name"] = display_name
        __props__["enabled"] = enabled
        __props__["issuer"] = issuer
        __props__["name"] = name
        __props__["project"] = project
        return OauthIdpConfig(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

