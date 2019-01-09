# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class BucketACL(pulumi.CustomResource):
    """
    Creates a new bucket ACL in Google cloud storage service (GCS). For more information see 
    [the official documentation](https://cloud.google.com/storage/docs/access-control/lists) 
    and 
    [API](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls).
    """
    def __init__(__self__, __name__, __opts__=None, bucket=None, default_acl=None, predefined_acl=None, role_entities=None):
        """Create a BucketACL resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not bucket:
            raise TypeError('Missing required property bucket')
        __props__['bucket'] = bucket

        __props__['default_acl'] = default_acl

        __props__['predefined_acl'] = predefined_acl

        __props__['role_entities'] = role_entities

        super(BucketACL, __self__).__init__(
            'gcp:storage/bucketACL:BucketACL',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

