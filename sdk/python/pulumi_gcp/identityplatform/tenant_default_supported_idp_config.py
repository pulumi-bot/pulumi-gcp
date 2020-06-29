# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class TenantDefaultSupportedIdpConfig(pulumi.CustomResource):
    client_id: pulumi.Output[str]
    """
    OAuth client ID
    """
    client_secret: pulumi.Output[str]
    """
    OAuth client secret
    """
    enabled: pulumi.Output[bool]
    """
    If this IDP allows the user to sign in
    """
    idp_id: pulumi.Output[str]
    """
    ID of the IDP. Possible values include:
    * `apple.com`
    * `facebook.com`
    * `gc.apple.com`
    * `github.com`
    * `google.com`
    * `linkedin.com`
    * `microsoft.com`
    * `playgames.google.com`
    * `twitter.com`
    * `yahoo.com`
    """
    name: pulumi.Output[str]
    """
    The name of the default supported IDP config resource
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    tenant: pulumi.Output[str]
    """
    The name of the tenant where this DefaultSupportedIdpConfig resource exists
    """
    def __init__(__self__, resource_name, opts=None, client_id=None, client_secret=None, enabled=None, idp_id=None, project=None, tenant=None, __props__=None, __name__=None, __opts__=None):
        """
        Configurations options for the tenant for authenticating with a the standard set of Identity Toolkit-trusted IDPs.

        You must enable the
        [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
        the marketplace prior to using this resource.

        ## Example Usage
        ### Identity Platform Tenant Default Supported Idp Config Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        tenant = gcp.identityplatform.Tenant("tenant", display_name="tenant")
        idp_config = gcp.identityplatform.TenantDefaultSupportedIdpConfig("idpConfig",
            enabled=True,
            tenant=tenant.name,
            idp_id="playgames.google.com",
            client_id="my-client-id",
            client_secret="secret")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: OAuth client ID
        :param pulumi.Input[str] client_secret: OAuth client secret
        :param pulumi.Input[bool] enabled: If this IDP allows the user to sign in
        :param pulumi.Input[str] idp_id: ID of the IDP. Possible values include:
               * `apple.com`
               * `facebook.com`
               * `gc.apple.com`
               * `github.com`
               * `google.com`
               * `linkedin.com`
               * `microsoft.com`
               * `playgames.google.com`
               * `twitter.com`
               * `yahoo.com`
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] tenant: The name of the tenant where this DefaultSupportedIdpConfig resource exists
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
            if client_secret is None:
                raise TypeError("Missing required property 'client_secret'")
            __props__['client_secret'] = client_secret
            __props__['enabled'] = enabled
            if idp_id is None:
                raise TypeError("Missing required property 'idp_id'")
            __props__['idp_id'] = idp_id
            __props__['project'] = project
            if tenant is None:
                raise TypeError("Missing required property 'tenant'")
            __props__['tenant'] = tenant
            __props__['name'] = None
        super(TenantDefaultSupportedIdpConfig, __self__).__init__(
            'gcp:identityplatform/tenantDefaultSupportedIdpConfig:TenantDefaultSupportedIdpConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, client_id=None, client_secret=None, enabled=None, idp_id=None, name=None, project=None, tenant=None):
        """
        Get an existing TenantDefaultSupportedIdpConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: OAuth client ID
        :param pulumi.Input[str] client_secret: OAuth client secret
        :param pulumi.Input[bool] enabled: If this IDP allows the user to sign in
        :param pulumi.Input[str] idp_id: ID of the IDP. Possible values include:
               * `apple.com`
               * `facebook.com`
               * `gc.apple.com`
               * `github.com`
               * `google.com`
               * `linkedin.com`
               * `microsoft.com`
               * `playgames.google.com`
               * `twitter.com`
               * `yahoo.com`
        :param pulumi.Input[str] name: The name of the default supported IDP config resource
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] tenant: The name of the tenant where this DefaultSupportedIdpConfig resource exists
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["client_id"] = client_id
        __props__["client_secret"] = client_secret
        __props__["enabled"] = enabled
        __props__["idp_id"] = idp_id
        __props__["name"] = name
        __props__["project"] = project
        __props__["tenant"] = tenant
        return TenantDefaultSupportedIdpConfig(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
