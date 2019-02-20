# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class BucketObject(pulumi.CustomResource):
    bucket: pulumi.Output[str]
    """
    The name of the containing bucket.
    """
    cache_control: pulumi.Output[str]
    """
    [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
    directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
    """
    content: pulumi.Output[str]
    """
    Data as `string` to be uploaded. Must be defined if
    `source` is not.
    """
    content_disposition: pulumi.Output[str]
    """
    [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
    """
    content_encoding: pulumi.Output[str]
    """
    [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
    """
    content_language: pulumi.Output[str]
    """
    [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
    """
    content_type: pulumi.Output[str]
    """
    [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
    """
    crc32c: pulumi.Output[str]
    """
    (Computed) Base 64 CRC32 hash of the uploaded data.
    """
    detect_md5hash: pulumi.Output[str]
    md5hash: pulumi.Output[str]
    """
    (Computed) Base 64 MD5 hash of the uploaded data.
    """
    name: pulumi.Output[str]
    """
    The name of the object.
    """
    source: pulumi.Output[pulumi.Archive]
    """
    A path to the data you want to upload. Must be defined
    if `content` is not.
    """
    storage_class: pulumi.Output[str]
    """
    The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
    Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`. If not provided, this defaults to the bucket's default
    storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
    """
    def __init__(__self__, resource_name, opts=None, bucket=None, cache_control=None, content=None, content_disposition=None, content_encoding=None, content_language=None, content_type=None, detect_md5hash=None, name=None, source=None, storage_class=None, __name__=None, __opts__=None):
        """
        Creates a new object inside an existing bucket in Google cloud storage service (GCS). 
        [ACLs](https://cloud.google.com/storage/docs/access-control/lists) can be applied using the `google_storage_object_acl` resource.
         For more information see 
        [the official documentation](https://cloud.google.com/storage/docs/key-terms#objects) 
        and 
        [API](https://cloud.google.com/storage/docs/json_api/v1/objects).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the containing bucket.
        :param pulumi.Input[str] cache_control: [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
               directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
        :param pulumi.Input[str] content: Data as `string` to be uploaded. Must be defined if
               `source` is not.
        :param pulumi.Input[str] content_disposition: [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
        :param pulumi.Input[str] content_encoding: [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
        :param pulumi.Input[str] content_language: [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
        :param pulumi.Input[str] content_type: [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
        :param pulumi.Input[str] name: The name of the object.
        :param pulumi.Input[pulumi.Archive] source: A path to the data you want to upload. Must be defined
               if `content` is not.
        :param pulumi.Input[str] storage_class: The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
               Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`. If not provided, this defaults to the bucket's default
               storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
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

        if bucket is None:
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
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

