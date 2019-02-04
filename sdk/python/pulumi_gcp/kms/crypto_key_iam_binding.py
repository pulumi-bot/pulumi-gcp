# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class CryptoKeyIAMBinding(pulumi.CustomResource):
    crypto_key_id: pulumi.Output[str]
    """
    The crypto key ID, in the form
    `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
    `{location_name}/{key_ring_name}/{crypto_key_name}`.
    In the second form, the provider's project setting will be used as a fallback.
    """
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the crypto key's IAM policy.
    """
    members: pulumi.Output[list]
    """
    A list of users that the role should apply to.
    """
    role: pulumi.Output[str]
    """
    The role that should be applied. Only one
    `google_kms_crypto_key_iam_binding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    def __init__(__self__, resource_name, opts=None, crypto_key_id=None, members=None, role=None, __name__=None, __opts__=None):
        """
        Allows creation and management of a single binding within IAM policy for
        an existing Google Cloud KMS crypto key.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] crypto_key_id: The crypto key ID, in the form
               `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
               `{location_name}/{key_ring_name}/{crypto_key_name}`.
               In the second form, the provider's project setting will be used as a fallback.
        :param pulumi.Input[list] members: A list of users that the role should apply to.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `google_kms_crypto_key_iam_binding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not crypto_key_id:
            raise TypeError('Missing required property crypto_key_id')
        __props__['crypto_key_id'] = crypto_key_id

        if not members:
            raise TypeError('Missing required property members')
        __props__['members'] = members

        if not role:
            raise TypeError('Missing required property role')
        __props__['role'] = role

        __props__['etag'] = None

        super(CryptoKeyIAMBinding, __self__).__init__(
            'gcp:kms/cryptoKeyIAMBinding:CryptoKeyIAMBinding',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

