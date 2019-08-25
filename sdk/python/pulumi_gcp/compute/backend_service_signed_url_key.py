# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class BackendServiceSignedUrlKey(pulumi.CustomResource):
    backend_service: pulumi.Output[str]
    key_value: pulumi.Output[str]
    name: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    def __init__(__self__, resource_name, opts=None, backend_service=None, key_value=None, name=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        A key for signing Cloud CDN signed URLs for Backend Services.
        
        
        To get more information about BackendServiceSignedUrlKey, see:
        
        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/backendServices)
        * How-to Guides
            * [Using Signed URLs](https://cloud.google.com/cdn/docs/using-signed-urls/)
        
        > **Warning:** All arguments including the key's value will be stored in the raw
        state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
        Because the API does not return the sensitive key value,
        we cannot confirm or reverse changes to a key outside of this provider.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_backend_service_signed_url_key.html.markdown.
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

            if backend_service is None:
                raise TypeError("Missing required property 'backend_service'")
            __props__['backend_service'] = backend_service
            if key_value is None:
                raise TypeError("Missing required property 'key_value'")
            __props__['key_value'] = key_value
            __props__['name'] = name
            __props__['project'] = project
        super(BackendServiceSignedUrlKey, __self__).__init__(
            'gcp:compute/backendServiceSignedUrlKey:BackendServiceSignedUrlKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, backend_service=None, key_value=None, name=None, project=None):
        """
        Get an existing BackendServiceSignedUrlKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_backend_service_signed_url_key.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["backend_service"] = backend_service
        __props__["key_value"] = key_value
        __props__["name"] = name
        __props__["project"] = project
        return BackendServiceSignedUrlKey(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

