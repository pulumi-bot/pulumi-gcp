# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class InboundSamlConfig(pulumi.CustomResource):
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
    SAML IdP configuration when the project acts as the relying party  Structure is documented below.

      * `idpCertificates` (`list`) - The IdP's certificate data to verify the signature in the SAMLResponse issued by the IDP.  Structure is documented below.
        * `x509Certificate` (`str`) - -
          The x509 certificate

      * `idpEntityId` (`str`) - Unique identifier for all SAML entities
      * `signRequest` (`bool`) - Indicates if outbounding SAMLRequest should be signed.
      * `ssoUrl` (`str`) - URL to send Authentication request to.
    """
    name: pulumi.Output[str]
    """
    The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
    hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
    alphanumeric character, and have at least 2 characters.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    sp_config: pulumi.Output[dict]
    """
    SAML SP (Service Provider) configuration when the project acts as the relying party to receive
    and accept an authentication assertion issued by a SAML identity provider.  Structure is documented below.

      * `callbackUri` (`str`) - Callback URI where responses from IDP are handled. Must start with `https://`.
      * `spCertificates` (`list`) - -
        The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP.  Structure is documented below.
        * `x509Certificate` (`str`) - -
          The x509 certificate

      * `spEntityId` (`str`) - Unique identifier for all SAML entities.
    """
    def __init__(__self__, resource_name, opts=None, display_name=None, enabled=None, idp_config=None, name=None, project=None, sp_config=None, __props__=None, __name__=None, __opts__=None):
        """
        Inbound SAML configuration for a Identity Toolkit project.

        You must enable the
        [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
        the marketplace prior to using this resource.



        ## Example Usage - Identity Platform Inbound Saml Config Basic


        {{ % example python % }}
        ```python
        import pulumi
        import pulumi_gcp as gcp

        saml_config = gcp.identityplatform.InboundSamlConfig("samlConfig",
            display_name="Display Name",
            idp_config={
                "idpEntityId": "tf-idp",
                "signRequest": True,
                "ssoUrl": "https://example.com",
                "idp_certificates": [{
                    "x509Certificate": (lambda path: open(path).read())("test-fixtures/rsa_cert.pem"),
                }],
            },
            sp_config={
                "spEntityId": "tf-sp",
                "callbackUri": "https://example.com",
            })
        ```
        {{ % /example % }}

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Human friendly display name.
        :param pulumi.Input[bool] enabled: If this config allows users to sign in with the provider.
        :param pulumi.Input[dict] idp_config: SAML IdP configuration when the project acts as the relying party  Structure is documented below.
        :param pulumi.Input[str] name: The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
               hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
               alphanumeric character, and have at least 2 characters.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] sp_config: SAML SP (Service Provider) configuration when the project acts as the relying party to receive
               and accept an authentication assertion issued by a SAML identity provider.  Structure is documented below.

        The **idp_config** object supports the following:

          * `idpCertificates` (`pulumi.Input[list]`) - The IdP's certificate data to verify the signature in the SAMLResponse issued by the IDP.  Structure is documented below.
            * `x509Certificate` (`pulumi.Input[str]`) - -
              The x509 certificate

          * `idpEntityId` (`pulumi.Input[str]`) - Unique identifier for all SAML entities
          * `signRequest` (`pulumi.Input[bool]`) - Indicates if outbounding SAMLRequest should be signed.
          * `ssoUrl` (`pulumi.Input[str]`) - URL to send Authentication request to.

        The **sp_config** object supports the following:

          * `callbackUri` (`pulumi.Input[str]`) - Callback URI where responses from IDP are handled. Must start with `https://`.
          * `spCertificates` (`pulumi.Input[list]`) - -
            The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP.  Structure is documented below.
            * `x509Certificate` (`pulumi.Input[str]`) - -
              The x509 certificate

          * `spEntityId` (`pulumi.Input[str]`) - Unique identifier for all SAML entities.
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
        super(InboundSamlConfig, __self__).__init__(
            'gcp:identityplatform/inboundSamlConfig:InboundSamlConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, display_name=None, enabled=None, idp_config=None, name=None, project=None, sp_config=None):
        """
        Get an existing InboundSamlConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Human friendly display name.
        :param pulumi.Input[bool] enabled: If this config allows users to sign in with the provider.
        :param pulumi.Input[dict] idp_config: SAML IdP configuration when the project acts as the relying party  Structure is documented below.
        :param pulumi.Input[str] name: The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
               hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
               alphanumeric character, and have at least 2 characters.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] sp_config: SAML SP (Service Provider) configuration when the project acts as the relying party to receive
               and accept an authentication assertion issued by a SAML identity provider.  Structure is documented below.

        The **idp_config** object supports the following:

          * `idpCertificates` (`pulumi.Input[list]`) - The IdP's certificate data to verify the signature in the SAMLResponse issued by the IDP.  Structure is documented below.
            * `x509Certificate` (`pulumi.Input[str]`) - -
              The x509 certificate

          * `idpEntityId` (`pulumi.Input[str]`) - Unique identifier for all SAML entities
          * `signRequest` (`pulumi.Input[bool]`) - Indicates if outbounding SAMLRequest should be signed.
          * `ssoUrl` (`pulumi.Input[str]`) - URL to send Authentication request to.

        The **sp_config** object supports the following:

          * `callbackUri` (`pulumi.Input[str]`) - Callback URI where responses from IDP are handled. Must start with `https://`.
          * `spCertificates` (`pulumi.Input[list]`) - -
            The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP.  Structure is documented below.
            * `x509Certificate` (`pulumi.Input[str]`) - -
              The x509 certificate

          * `spEntityId` (`pulumi.Input[str]`) - Unique identifier for all SAML entities.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["display_name"] = display_name
        __props__["enabled"] = enabled
        __props__["idp_config"] = idp_config
        __props__["name"] = name
        __props__["project"] = project
        __props__["sp_config"] = sp_config
        return InboundSamlConfig(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

