# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetObjectSignedUrlResult:
    """
    A collection of values returned by getObjectSignedUrl.
    """
    def __init__(__self__, bucket=None, content_md5=None, content_type=None, credentials=None, duration=None, extension_headers=None, http_method=None, path=None, signed_url=None, id=None):
        if bucket and not isinstance(bucket, str):
            raise TypeError("Expected argument 'bucket' to be a str")
        __self__.bucket = bucket
        if content_md5 and not isinstance(content_md5, str):
            raise TypeError("Expected argument 'content_md5' to be a str")
        __self__.content_md5 = content_md5
        if content_type and not isinstance(content_type, str):
            raise TypeError("Expected argument 'content_type' to be a str")
        __self__.content_type = content_type
        if credentials and not isinstance(credentials, str):
            raise TypeError("Expected argument 'credentials' to be a str")
        __self__.credentials = credentials
        if duration and not isinstance(duration, str):
            raise TypeError("Expected argument 'duration' to be a str")
        __self__.duration = duration
        if extension_headers and not isinstance(extension_headers, dict):
            raise TypeError("Expected argument 'extension_headers' to be a dict")
        __self__.extension_headers = extension_headers
        if http_method and not isinstance(http_method, str):
            raise TypeError("Expected argument 'http_method' to be a str")
        __self__.http_method = http_method
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        __self__.path = path
        if signed_url and not isinstance(signed_url, str):
            raise TypeError("Expected argument 'signed_url' to be a str")
        __self__.signed_url = signed_url
        """
        The signed URL that can be used to access the storage object without authentication.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        delattr(self, "__await__")
        delattr(self, "__iter__")
        return self

    __iter__ = __await__

def get_object_signed_url(bucket=None,content_md5=None,content_type=None,credentials=None,duration=None,extension_headers=None,http_method=None,path=None,opts=None):
    """
    The Google Cloud storage signed URL data source generates a signed URL for a given storage object. Signed URLs provide a way to give time-limited read or write access to anyone in possession of the URL, regardless of whether they have a Google account.
    
    For more info about signed URL's is available [here](https://cloud.google.com/storage/docs/access-control/signed-urls).

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/storage_object_signed_url.html.markdown.
    """
    __args__ = dict()

    __args__['bucket'] = bucket
    __args__['contentMd5'] = content_md5
    __args__['contentType'] = content_type
    __args__['credentials'] = credentials
    __args__['duration'] = duration
    __args__['extensionHeaders'] = extension_headers
    __args__['httpMethod'] = http_method
    __args__['path'] = path
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:storage/getObjectSignedUrl:getObjectSignedUrl', __args__, opts=opts).value

    return GetObjectSignedUrlResult(
        bucket=__ret__.get('bucket'),
        content_md5=__ret__.get('contentMd5'),
        content_type=__ret__.get('contentType'),
        credentials=__ret__.get('credentials'),
        duration=__ret__.get('duration'),
        extension_headers=__ret__.get('extensionHeaders'),
        http_method=__ret__.get('httpMethod'),
        path=__ret__.get('path'),
        signed_url=__ret__.get('signedUrl'),
        id=__ret__.get('id'))
