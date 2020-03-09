# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class TenantInboundSamlConfig(pulumi.CustomResource):
    display_name: pulumi.Output[str]
    """
    Human friendly display name.
    """
    enabled: pulumi.Output[bool]
    """
    If this config allows users to sign in with the provider.
    """
    idp_config: pulumi.Output[dict]
    """
    SAML IdP configuration when the project acts as the relying party

      * `idpCertificates` (`list`)
        * `x509Certificate` (`str`)

      * `idpEntityId` (`str`)
      * `signRequest` (`bool`)
      * `ssoUrl` (`str`)
    """
    name: pulumi.Output[str]
    """
    The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters, hyphens,
    underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an alphanumeric
    character, and have at least 2 characters.
    """
    project: pulumi.Output[str]
    sp_config: pulumi.Output[dict]
    """
    SAML SP (Service Provider) configuration when the project acts as the relying party to receive and accept an
    authentication assertion issued by a SAML identity provider.

      * `callbackUri` (`str`)
      * `spCertificates` (`list`)
        * `x509Certificate` (`str`)

      * `spEntityId` (`str`)
    """
    tenant: pulumi.Output[str]
    """
    The name of the tenant where this inbound SAML config resource exists
    """
    def __init__(__self__, resource_name, opts=None, display_name=None, enabled=None, idp_config=None, name=None, project=None, sp_config=None, tenant=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a TenantInboundSamlConfig resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Human friendly display name.
        :param pulumi.Input[bool] enabled: If this config allows users to sign in with the provider.
        :param pulumi.Input[dict] idp_config: SAML IdP configuration when the project acts as the relying party
        :param pulumi.Input[str] name: The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters, hyphens,
               underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an alphanumeric
               character, and have at least 2 characters.
        :param pulumi.Input[dict] sp_config: SAML SP (Service Provider) configuration when the project acts as the relying party to receive and accept an
               authentication assertion issued by a SAML identity provider.
        :param pulumi.Input[str] tenant: The name of the tenant where this inbound SAML config resource exists

        The **idp_config** object supports the following:

          * `idpCertificates` (`pulumi.Input[list]`)
            * `x509Certificate` (`pulumi.Input[str]`)

          * `idpEntityId` (`pulumi.Input[str]`)
          * `signRequest` (`pulumi.Input[bool]`)
          * `ssoUrl` (`pulumi.Input[str]`)

        The **sp_config** object supports the following:

          * `callbackUri` (`pulumi.Input[str]`)
          * `spCertificates` (`pulumi.Input[list]`)
            * `x509Certificate` (`pulumi.Input[str]`)

          * `spEntityId` (`pulumi.Input[str]`)
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
    def get(resource_name, id, opts=None, display_name=None, enabled=None, idp_config=None, name=None, project=None, sp_config=None, tenant=None):
        """
        Get an existing TenantInboundSamlConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Human friendly display name.
        :param pulumi.Input[bool] enabled: If this config allows users to sign in with the provider.
        :param pulumi.Input[dict] idp_config: SAML IdP configuration when the project acts as the relying party
        :param pulumi.Input[str] name: The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters, hyphens,
               underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an alphanumeric
               character, and have at least 2 characters.
        :param pulumi.Input[dict] sp_config: SAML SP (Service Provider) configuration when the project acts as the relying party to receive and accept an
               authentication assertion issued by a SAML identity provider.
        :param pulumi.Input[str] tenant: The name of the tenant where this inbound SAML config resource exists

        The **idp_config** object supports the following:

          * `idpCertificates` (`pulumi.Input[list]`)
            * `x509Certificate` (`pulumi.Input[str]`)

          * `idpEntityId` (`pulumi.Input[str]`)
          * `signRequest` (`pulumi.Input[bool]`)
          * `ssoUrl` (`pulumi.Input[str]`)

        The **sp_config** object supports the following:

          * `callbackUri` (`pulumi.Input[str]`)
          * `spCertificates` (`pulumi.Input[list]`)
            * `x509Certificate` (`pulumi.Input[str]`)

          * `spEntityId` (`pulumi.Input[str]`)
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
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

