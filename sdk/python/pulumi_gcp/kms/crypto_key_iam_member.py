# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class CryptoKeyIAMMember(pulumi.CustomResource):
    """
    Allows creation and management of a single member for a single binding within
    the IAM policy for an existing Google Cloud KMS crypto key.
    
    > **Note:** This resource _must not_ be used in conjunction with
       `google_kms_crypto_key_iam_policy` or they will fight over what your policy
       should be. Similarly, roles controlled by `google_kms_crypto_key_iam_binding`
       should not be assigned to using `google_kms_crypto_key_iam_member`.
    """
    def __init__(__self__, __name__, __opts__=None, crypto_key_id=None, member=None, role=None):
        """Create a CryptoKeyIAMMember resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not crypto_key_id:
            raise TypeError('Missing required property crypto_key_id')
        __props__['crypto_key_id'] = crypto_key_id

        if not member:
            raise TypeError('Missing required property member')
        __props__['member'] = member

        if not role:
            raise TypeError('Missing required property role')
        __props__['role'] = role

        __props__['etag'] = None

        super(CryptoKeyIAMMember, __self__).__init__(
            'gcp:kms/cryptoKeyIAMMember:CryptoKeyIAMMember',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

