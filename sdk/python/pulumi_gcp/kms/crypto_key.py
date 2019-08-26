# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class CryptoKey(pulumi.CustomResource):
    key_ring: pulumi.Output[str]
    labels: pulumi.Output[dict]
    name: pulumi.Output[str]
    purpose: pulumi.Output[str]
    rotation_period: pulumi.Output[str]
    self_link: pulumi.Output[str]
    version_template: pulumi.Output[dict]
    def __init__(__self__, resource_name, opts=None, key_ring=None, labels=None, name=None, purpose=None, rotation_period=None, version_template=None, __props__=None, __name__=None, __opts__=None):
        """
        A `CryptoKey` represents a logical key that can be used for cryptographic operations.
        
        
        > **Note:** CryptoKeys cannot be deleted from Google Cloud Platform.
        Destroying a CryptoKey will remove it from state
        and delete all CryptoKeyVersions, rendering the key unusable, but *will
        not delete the resource on the server.* When this provider destroys these keys,
        any data previously encrypted with these keys will be irrecoverable.
        For this reason, it is strongly recommended that you add lifecycle hooks
        to the resource to prevent accidental destruction.
        
        
        To get more information about CryptoKey, see:
        
        * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys)
        * How-to Guides
            * [Creating a key](https://cloud.google.com/kms/docs/creating-keys#create_a_key)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        
        The **version_template** object supports the following:
        
          * `algorithm` (`pulumi.Input[str]`)
          * `protection_level` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/kms_crypto_key.html.markdown.
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

            if key_ring is None:
                raise TypeError("Missing required property 'key_ring'")
            __props__['key_ring'] = key_ring
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['purpose'] = purpose
            __props__['rotation_period'] = rotation_period
            __props__['version_template'] = version_template
            __props__['self_link'] = None
        super(CryptoKey, __self__).__init__(
            'gcp:kms/cryptoKey:CryptoKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, key_ring=None, labels=None, name=None, purpose=None, rotation_period=None, self_link=None, version_template=None):
        """
        Get an existing CryptoKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        
        The **version_template** object supports the following:
        
          * `algorithm` (`pulumi.Input[str]`)
          * `protection_level` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/kms_crypto_key.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["key_ring"] = key_ring
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["purpose"] = purpose
        __props__["rotation_period"] = rotation_period
        __props__["self_link"] = self_link
        __props__["version_template"] = version_template
        return CryptoKey(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

