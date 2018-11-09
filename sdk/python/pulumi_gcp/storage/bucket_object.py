# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class BucketObject(pulumi.CustomResource):
    """
    Creates a new object inside an existing bucket in Google cloud storage service (GCS). 
    [ACLs](https://cloud.google.com/storage/docs/access-control/lists) can be applied using the `google_storage_object_acl` resource.
     For more information see 
    [the official documentation](https://cloud.google.com/storage/docs/key-terms#objects) 
    and 
    [API](https://cloud.google.com/storage/docs/json_api/v1/objects).
    
    """
    def __init__(__self__, __name__, __opts__=None, bucket=None, cache_control=None, content=None, content_disposition=None, content_encoding=None, content_language=None, content_type=None, detect_md5hash=None, name=None, source=None, storage_class=None):
        """Create a BucketObject resource with the given unique name, props, and options."""
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

        __props__['cache_control'] = cache_control

        __props__['content'] = content

        __props__['content_disposition'] = content_disposition

        __props__['content_encoding'] = content_encoding

        __props__['content_language'] = content_language

        __props__['content_type'] = content_type

        __props__['detect_md5hash'] = detect_md5hash

        __props__['name'] = name

        __props__['source'] = source

        __props__['storage_class'] = storage_class

        __props__['crc32c'] = None
        __props__['md5hash'] = None

        super(BucketObject, __self__).__init__(
            'gcp:storage/bucketObject:BucketObject',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

