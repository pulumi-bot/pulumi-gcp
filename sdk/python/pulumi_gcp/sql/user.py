# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class User(pulumi.CustomResource):
    host: pulumi.Output[str]
    """
    The host the user can connect from. This is only supported
    for MySQL instances. Don't set this field for PostgreSQL instances.
    Can be an IP address. Changing this forces a new resource to be created.
    """
    instance: pulumi.Output[str]
    """
    The name of the Cloud SQL instance. Changing this
    forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The name of the user. Changing this forces a new resource
    to be created.
    """
    password: pulumi.Output[str]
    """
    The password for the user. Can be updated.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    def __init__(__self__, resource_name, opts=None, host=None, instance=None, name=None, password=None, project=None, __name__=None, __opts__=None):
        """
        Creates a new Google SQL User on a Google SQL User Instance. For more information, see the [official documentation](https://cloud.google.com/sql/), or the [JSON API](https://cloud.google.com/sql/docs/admin-api/v1beta4/users).
        
        > **Note:** All arguments including the username and password will be stored in the raw state as plain-text.
        [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html). Passwords will not be retrieved when running import.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host: The host the user can connect from. This is only supported
               for MySQL instances. Don't set this field for PostgreSQL instances.
               Can be an IP address. Changing this forces a new resource to be created.
        :param pulumi.Input[str] instance: The name of the Cloud SQL instance. Changing this
               forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the user. Changing this forces a new resource
               to be created.
        :param pulumi.Input[str] password: The password for the user. Can be updated.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/sql_user.html.markdown.
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

        __props__['host'] = host
        if instance is None:
            raise TypeError("Missing required property 'instance'")
        __props__['instance'] = instance
        __props__['name'] = name
        __props__['password'] = password
        __props__['project'] = project
        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(User, __self__).__init__(
            'gcp:sql/user:User',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

