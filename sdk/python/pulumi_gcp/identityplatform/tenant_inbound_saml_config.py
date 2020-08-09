# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['TenantInboundSamlConfig']


class TenantInboundSamlConfig(pulumi.CustomResource):
    display_name: pulumi.Output[str] = pulumi.property("displayName")
    """
    Human friendly display name.
    """

    enabled: pulumi.Output[Optional[bool]] = pulumi.property("enabled")
    """
    If this config allows users to sign in with the provider.
    """

    idp_config: pulumi.Output['outputs.TenantInboundSamlConfigIdpConfig'] = pulumi.property("idpConfig")
    """
    SAML IdP configuration when the project acts as the relying party  Structure is documented below.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
    hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
    alphanumeric character, and have at least 2 characters.
    """

    project: pulumi.Output[str] = pulumi.property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """

    sp_config: pulumi.Output['outputs.TenantInboundSamlConfigSpConfig'] = pulumi.property("spConfig")
    """
    SAML SP (Service Provider) configuration when the project acts as the relying party to receive
    and accept an authentication assertion issued by a SAML identity provider.  Structure is documented below.
    """

    tenant: pulumi.Output[str] = pulumi.property("tenant")
    """
    The name of the tenant where this inbound SAML config resource exists
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 idp_config: Optional[pulumi.Input[pulumi.InputType['TenantInboundSamlConfigIdpConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 sp_config: Optional[pulumi.Input[pulumi.InputType['TenantInboundSamlConfigSpConfigArgs']]] = None,
                 tenant: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Inbound SAML configuration for a Identity Toolkit tenant.

        You must enable the
        [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
        the marketplace prior to using this resource.

        ## Example Usage
        ### Identity Platform Tenant Inbound Saml Config Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        tenant = gcp.identityplatform.Tenant("tenant", display_name="tenant")
        tenant_saml_config = gcp.identityplatform.TenantInboundSamlConfig("tenantSamlConfig",
            display_name="Display Name",
            tenant=tenant.name,
            idp_config={
                "idpEntityId": "tf-idp",
                "signRequest": True,
                "ssoUrl": "https://example.com",
                "idpCertificates": [{
                    "x509Certificate": (lambda path: open(path).read())("test-fixtures/rsa_cert.pem"),
                }],
            },
            sp_config={
                "spEntityId": "tf-sp",
                "callbackUri": "https://example.com",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Human friendly display name.
        :param pulumi.Input[bool] enabled: If this config allows users to sign in with the provider.
        :param pulumi.Input[pulumi.InputType['TenantInboundSamlConfigIdpConfigArgs']] idp_config: SAML IdP configuration when the project acts as the relying party  Structure is documented below.
        :param pulumi.Input[str] name: The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
               hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
               alphanumeric character, and have at least 2 characters.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['TenantInboundSamlConfigSpConfigArgs']] sp_config: SAML SP (Service Provider) configuration when the project acts as the relying party to receive
               and accept an authentication assertion issued by a SAML identity provider.  Structure is documented below.
        :param pulumi.Input[str] tenant: The name of the tenant where this inbound SAML config resource exists
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['enabled'] = enabled
            if idp_config is None:
                raise TypeError("Missing required property 'idp_config'")
            __props__['idp_config'] = idp_config
            __props__['name'] = name
            __props__['project'] = project
            if sp_config is None:
                raise TypeError("Missing required property 'sp_config'")
            __props__['sp_config'] = sp_config
            if tenant is None:
                raise TypeError("Missing required property 'tenant'")
            __props__['tenant'] = tenant
        super(TenantInboundSamlConfig, __self__).__init__(
            'gcp:identityplatform/tenantInboundSamlConfig:TenantInboundSamlConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            idp_config: Optional[pulumi.Input[pulumi.InputType['TenantInboundSamlConfigIdpConfigArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            sp_config: Optional[pulumi.Input[pulumi.InputType['TenantInboundSamlConfigSpConfigArgs']]] = None,
            tenant: Optional[pulumi.Input[str]] = None) -> 'TenantInboundSamlConfig':
        """
        Get an existing TenantInboundSamlConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Human friendly display name.
        :param pulumi.Input[bool] enabled: If this config allows users to sign in with the provider.
        :param pulumi.Input[pulumi.InputType['TenantInboundSamlConfigIdpConfigArgs']] idp_config: SAML IdP configuration when the project acts as the relying party  Structure is documented below.
        :param pulumi.Input[str] name: The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
               hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
               alphanumeric character, and have at least 2 characters.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['TenantInboundSamlConfigSpConfigArgs']] sp_config: SAML SP (Service Provider) configuration when the project acts as the relying party to receive
               and accept an authentication assertion issued by a SAML identity provider.  Structure is documented below.
        :param pulumi.Input[str] tenant: The name of the tenant where this inbound SAML config resource exists
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["display_name"] = display_name
        __props__["enabled"] = enabled
        __props__["idp_config"] = idp_config
        __props__["name"] = name
        __props__["project"] = project
        __props__["sp_config"] = sp_config
        __props__["tenant"] = tenant
        return TenantInboundSamlConfig(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

