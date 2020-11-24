# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['CryptoKey']


class CryptoKey(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_ring: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 purpose: Optional[pulumi.Input[str]] = None,
                 rotation_period: Optional[pulumi.Input[str]] = None,
                 skip_initial_version_creation: Optional[pulumi.Input[bool]] = None,
                 version_template: Optional[pulumi.Input[pulumi.InputType['CryptoKeyVersionTemplateArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A `CryptoKey` represents a logical key that can be used for cryptographic operations.

        > **Note:** CryptoKeys cannot be deleted from Google Cloud Platform.
        Destroying a provider-managed CryptoKey will remove it from state
        and delete all CryptoKeyVersions, rendering the key unusable, but *will
        not delete the resource from the project.* When the provider destroys these keys,
        any data previously encrypted with these keys will be irrecoverable.
        For this reason, it is strongly recommended that you add lifecycle hooks
        to the resource to prevent accidental destruction.

        To get more information about CryptoKey, see:

        * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys)
        * How-to Guides
            * [Creating a key](https://cloud.google.com/kms/docs/creating-keys#create_a_key)

        ## Example Usage
        ### Kms Crypto Key Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        keyring = gcp.kms.KeyRing("keyring", location="global")
        example_key = gcp.kms.CryptoKey("example-key",
            key_ring=keyring.id,
            rotation_period="100000s")
        ```
        ### Kms Crypto Key Asymmetric Sign

        ```python
        import pulumi
        import pulumi_gcp as gcp

        keyring = gcp.kms.KeyRing("keyring", location="global")
        example_asymmetric_sign_key = gcp.kms.CryptoKey("example-asymmetric-sign-key",
            key_ring=keyring.id,
            purpose="ASYMMETRIC_SIGN",
            version_template={
                "algorithm": "EC_SIGN_P384_SHA384",
            })
        ```

        ## Import

        CryptoKey can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:kms/cryptoKey:CryptoKey default {{key_ring}}/cryptoKeys/{{name}}
        ```

        ```sh
         $ pulumi import gcp:kms/cryptoKey:CryptoKey default {{key_ring}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_ring: The KeyRing that this key belongs to.
               Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels with user-defined metadata to apply to this resource.
        :param pulumi.Input[str] name: The resource name for the CryptoKey.
        :param pulumi.Input[str] purpose: The immutable purpose of this CryptoKey. See the
               [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
               for possible inputs.
               Default value is `ENCRYPT_DECRYPT`.
               Possible values are `ENCRYPT_DECRYPT`, `ASYMMETRIC_SIGN`, and `ASYMMETRIC_DECRYPT`.
        :param pulumi.Input[str] rotation_period: Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
               The first rotation will take place after the specified period. The rotation period has
               the format of a decimal number with up to 9 fractional digits, followed by the
               letter `s` (seconds). It must be greater than a day (ie, 86400).
        :param pulumi.Input[bool] skip_initial_version_creation: If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
               You must use the `kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
        :param pulumi.Input[pulumi.InputType['CryptoKeyVersionTemplateArgs']] version_template: A template describing settings for new crypto key versions.
               Structure is documented below.
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

            if key_ring is None:
                raise TypeError("Missing required property 'key_ring'")
            __props__['key_ring'] = key_ring
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['purpose'] = purpose
            __props__['rotation_period'] = rotation_period
            __props__['skip_initial_version_creation'] = skip_initial_version_creation
            __props__['version_template'] = version_template
            __props__['self_link'] = None
        super(CryptoKey, __self__).__init__(
            'gcp:kms/cryptoKey:CryptoKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            key_ring: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            purpose: Optional[pulumi.Input[str]] = None,
            rotation_period: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            skip_initial_version_creation: Optional[pulumi.Input[bool]] = None,
            version_template: Optional[pulumi.Input[pulumi.InputType['CryptoKeyVersionTemplateArgs']]] = None) -> 'CryptoKey':
        """
        Get an existing CryptoKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_ring: The KeyRing that this key belongs to.
               Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels with user-defined metadata to apply to this resource.
        :param pulumi.Input[str] name: The resource name for the CryptoKey.
        :param pulumi.Input[str] purpose: The immutable purpose of this CryptoKey. See the
               [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
               for possible inputs.
               Default value is `ENCRYPT_DECRYPT`.
               Possible values are `ENCRYPT_DECRYPT`, `ASYMMETRIC_SIGN`, and `ASYMMETRIC_DECRYPT`.
        :param pulumi.Input[str] rotation_period: Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
               The first rotation will take place after the specified period. The rotation period has
               the format of a decimal number with up to 9 fractional digits, followed by the
               letter `s` (seconds). It must be greater than a day (ie, 86400).
        :param pulumi.Input[bool] skip_initial_version_creation: If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
               You must use the `kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
        :param pulumi.Input[pulumi.InputType['CryptoKeyVersionTemplateArgs']] version_template: A template describing settings for new crypto key versions.
               Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["key_ring"] = key_ring
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["purpose"] = purpose
        __props__["rotation_period"] = rotation_period
        __props__["self_link"] = self_link
        __props__["skip_initial_version_creation"] = skip_initial_version_creation
        __props__["version_template"] = version_template
        return CryptoKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="keyRing")
    def key_ring(self) -> pulumi.Output[str]:
        """
        The KeyRing that this key belongs to.
        Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        """
        return pulumi.get(self, "key_ring")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Labels with user-defined metadata to apply to this resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name for the CryptoKey.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def purpose(self) -> pulumi.Output[Optional[str]]:
        """
        The immutable purpose of this CryptoKey. See the
        [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
        for possible inputs.
        Default value is `ENCRYPT_DECRYPT`.
        Possible values are `ENCRYPT_DECRYPT`, `ASYMMETRIC_SIGN`, and `ASYMMETRIC_DECRYPT`.
        """
        return pulumi.get(self, "purpose")

    @property
    @pulumi.getter(name="rotationPeriod")
    def rotation_period(self) -> pulumi.Output[Optional[str]]:
        """
        Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
        The first rotation will take place after the specified period. The rotation period has
        the format of a decimal number with up to 9 fractional digits, followed by the
        letter `s` (seconds). It must be greater than a day (ie, 86400).
        """
        return pulumi.get(self, "rotation_period")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="skipInitialVersionCreation")
    def skip_initial_version_creation(self) -> pulumi.Output[Optional[bool]]:
        """
        If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
        You must use the `kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
        """
        return pulumi.get(self, "skip_initial_version_creation")

    @property
    @pulumi.getter(name="versionTemplate")
    def version_template(self) -> pulumi.Output['outputs.CryptoKeyVersionTemplate']:
        """
        A template describing settings for new crypto key versions.
        Structure is documented below.
        """
        return pulumi.get(self, "version_template")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

