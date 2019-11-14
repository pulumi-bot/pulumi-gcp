# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Account(pulumi.CustomResource):
    account_id: pulumi.Output[str]
    """
    The account id that is used to generate the service
    account email address and a stable unique id. It is unique within a project,
    must be 6-30 characters long, and match the regular expression `a-z`
    to comply with RFC1035. Changing this forces a new service account to be created.
    """
    description: pulumi.Output[str]
    """
    A text description of the service account.
    """
    display_name: pulumi.Output[str]
    """
    The display name for the service account.
    Can be updated without creating a new resource.
    """
    email: pulumi.Output[str]
    """
    The e-mail address of the service account. This value
    should be referenced from any `organizations.getIAMPolicy` data sources
    that would grant the service account privileges.
    """
    name: pulumi.Output[str]
    """
    The fully-qualified name of the service account.
    """
    project: pulumi.Output[str]
    """
    The ID of the project that the service account will be created in.
    Defaults to the provider project configuration.
    """
    unique_id: pulumi.Output[str]
    """
    The unique id of the service account.
    """
    def __init__(__self__, resource_name, opts=None, account_id=None, description=None, display_name=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        Allows management of a [Google Cloud Platform service account](https://cloud.google.com/compute/docs/access/service-accounts)
        
        > Creation of service accounts is eventually consistent, and that can lead to
        errors when you try to apply ACLs to service accounts immediately after
        creation.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The account id that is used to generate the service
               account email address and a stable unique id. It is unique within a project,
               must be 6-30 characters long, and match the regular expression `a-z`
               to comply with RFC1035. Changing this forces a new service account to be created.
        :param pulumi.Input[str] description: A text description of the service account.
        :param pulumi.Input[str] display_name: The display name for the service account.
               Can be updated without creating a new resource.
        :param pulumi.Input[str] project: The ID of the project that the service account will be created in.
               Defaults to the provider project configuration.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/service_account.html.markdown.
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

            if account_id is None:
                raise TypeError("Missing required property 'account_id'")
            __props__['account_id'] = account_id
            __props__['description'] = description
            __props__['display_name'] = display_name
            __props__['project'] = project
            __props__['email'] = None
            __props__['name'] = None
            __props__['unique_id'] = None
        super(Account, __self__).__init__(
            'gcp:serviceAccount/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account_id=None, description=None, display_name=None, email=None, name=None, project=None, unique_id=None):
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The account id that is used to generate the service
               account email address and a stable unique id. It is unique within a project,
               must be 6-30 characters long, and match the regular expression `a-z`
               to comply with RFC1035. Changing this forces a new service account to be created.
        :param pulumi.Input[str] description: A text description of the service account.
        :param pulumi.Input[str] display_name: The display name for the service account.
               Can be updated without creating a new resource.
        :param pulumi.Input[str] email: The e-mail address of the service account. This value
               should be referenced from any `organizations.getIAMPolicy` data sources
               that would grant the service account privileges.
        :param pulumi.Input[str] name: The fully-qualified name of the service account.
        :param pulumi.Input[str] project: The ID of the project that the service account will be created in.
               Defaults to the provider project configuration.
        :param pulumi.Input[str] unique_id: The unique id of the service account.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/service_account.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["account_id"] = account_id
        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["email"] = email
        __props__["name"] = name
        __props__["project"] = project
        __props__["unique_id"] = unique_id
        return Account(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

