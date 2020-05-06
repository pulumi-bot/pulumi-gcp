# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SecretVersion(pulumi.CustomResource):
    create_time: pulumi.Output[str]
    """
    The time at which the Secret was created.
    """
    destroy_time: pulumi.Output[str]
    """
    The time at which the Secret was destroyed. Only present if state is DESTROYED.
    """
    enabled: pulumi.Output[bool]
    """
    The current state of the SecretVersion.
    """
    name: pulumi.Output[str]
    """
    The resource name of the SecretVersion. Format: 'projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}'
    """
    secret: pulumi.Output[str]
    """
    Secret Manager secret resource
    """
    secret_data: pulumi.Output[str]
    """
    The secret data. Must be no larger than 64KiB.  **Note**: This property is sensitive and will not be displayed in the plan.
    """
    def __init__(__self__, resource_name, opts=None, enabled=None, secret=None, secret_data=None, __props__=None, __name__=None, __opts__=None):
        """
        A secret version resource.

        > **Warning:** All arguments including `payload.secret_data` will be stored in the raw
        state as plain-text.

        ## Example Usage - Secret Version Basic


        {{ % example python % }}
        ```python
        import pulumi
        import pulumi_gcp as gcp

        secret_basic = gcp.secretmanager.Secret("secret-basic",
            secret_id="secret-version",
            labels={
                "label": "my-label",
            },
            replication={
                "automatic": True,
            })
        secret_version_basic = gcp.secretmanager.SecretVersion("secret-version-basic",
            secret=secret_basic.id,
            secret_data="secret-data")
        ```
        {{ % /example % }}

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: The current state of the SecretVersion.
        :param pulumi.Input[str] secret: Secret Manager secret resource
        :param pulumi.Input[str] secret_data: The secret data. Must be no larger than 64KiB.  **Note**: This property is sensitive and will not be displayed in the plan.
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

            __props__['enabled'] = enabled
            if secret is None:
                raise TypeError("Missing required property 'secret'")
            __props__['secret'] = secret
            __props__['secret_data'] = secret_data
            __props__['create_time'] = None
            __props__['destroy_time'] = None
            __props__['name'] = None
        super(SecretVersion, __self__).__init__(
            'gcp:secretmanager/secretVersion:SecretVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, create_time=None, destroy_time=None, enabled=None, name=None, secret=None, secret_data=None):
        """
        Get an existing SecretVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: The time at which the Secret was created.
        :param pulumi.Input[str] destroy_time: The time at which the Secret was destroyed. Only present if state is DESTROYED.
        :param pulumi.Input[bool] enabled: The current state of the SecretVersion.
        :param pulumi.Input[str] name: The resource name of the SecretVersion. Format: 'projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}'
        :param pulumi.Input[str] secret: Secret Manager secret resource
        :param pulumi.Input[str] secret_data: The secret data. Must be no larger than 64KiB.  **Note**: This property is sensitive and will not be displayed in the plan.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["create_time"] = create_time
        __props__["destroy_time"] = destroy_time
        __props__["enabled"] = enabled
        __props__["name"] = name
        __props__["secret"] = secret
        __props__["secret_data"] = secret_data
        return SecretVersion(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

