# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetObjectSignedUrlResult:
    """
    A collection of values returned by getObjectSignedUrl.
    """
    def __init__(__self__, bucket=None, content_md5=None, content_type=None, credentials=None, duration=None, extension_headers=None, http_method=None, id=None, path=None, signed_url=None):
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
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        __self__.path = path
        if signed_url and not isinstance(signed_url, str):
            raise TypeError("Expected argument 'signed_url' to be a str")
        __self__.signed_url = signed_url
        """
        The signed URL that can be used to access the storage object without authentication.
        """
class AwaitableGetObjectSignedUrlResult(GetObjectSignedUrlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetObjectSignedUrlResult(
            bucket=self.bucket,
            content_md5=self.content_md5,
            content_type=self.content_type,
            credentials=self.credentials,
            duration=self.duration,
            extension_headers=self.extension_headers,
            http_method=self.http_method,
            id=self.id,
            path=self.path,
            signed_url=self.signed_url)

def get_object_signed_url(bucket=None,content_md5=None,content_type=None,credentials=None,duration=None,extension_headers=None,http_method=None,path=None,opts=None):
    """
    The Google Cloud storage signed URL data source generates a signed URL for a given storage object. Signed URLs provide a way to give time-limited read or write access to anyone in possession of the URL, regardless of whether they have a Google account.

    For more info about signed URL's is available [here](https://cloud.google.com/storage/docs/access-control/signed-urls).

    {{% examples %}}
    ## Example Usage
    {{% example %}}

    ```python
    import pulumi
    import pulumi_gcp as gcp

    artifact = gcp.storage.get_object_signed_url(bucket="install_binaries",
        path="path/to/install_file.bin")
    vm = gcp.compute.Instance("vm")
    ```
    {{% /example %}}
    {{% /examples %}}
    ## Full Example

    ```python
    import pulumi
    import pulumi_gcp as gcp

    get_url = gcp.storage.get_object_signed_url(bucket="fried_chicken",
        path="path/to/file",
        content_md5="pRviqwS4c4OTJRTe03FD1w==",
        content_type="text/plain",
        duration="2d",
        credentials=(lambda path: open(path).read())("path/to/credentials.json"),
        extension_headers={
            "x-goog-if-generation-match": 1,
        })
    ```


    :param str bucket: The name of the bucket to read the object from
    :param str content_md5: The [MD5 digest](https://cloud.google.com/storage/docs/hashes-etags#_MD5) value in Base64.
           Typically retrieved from `google_storage_bucket_object.object.md5hash` attribute.
           If you provide this in the datasource, the client (e.g. browser, curl) must provide the `Content-MD5` HTTP header with this same value in its request.
    :param str content_type: If you specify this in the datasource, the client must provide the `Content-Type` HTTP header with the same value in its request.
    :param str credentials: What Google service account credentials json should be used to sign the URL.
           This data source checks the following locations for credentials, in order of preference: data source `credentials` attribute, provider `credentials` attribute and finally the GOOGLE_APPLICATION_CREDENTIALS environment variable.
    :param str duration: For how long shall the signed URL be valid (defaults to 1 hour - i.e. `1h`).
           See [here](https://golang.org/pkg/time/#ParseDuration) for info on valid duration formats.
    :param dict extension_headers: As needed. The server checks to make sure that the client provides matching values in requests using the signed URL.
           Any header starting with `x-goog-` is accepted but see the [Google Docs](https://cloud.google.com/storage/docs/xml-api/reference-headers) for list of headers that are supported by Google.
    :param str http_method: What HTTP Method will the signed URL allow (defaults to `GET`)
    :param str path: The full path to the object inside the bucket
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
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:storage/getObjectSignedUrl:getObjectSignedUrl', __args__, opts=opts).value

    return AwaitableGetObjectSignedUrlResult(
        bucket=__ret__.get('bucket'),
        content_md5=__ret__.get('contentMd5'),
        content_type=__ret__.get('contentType'),
        credentials=__ret__.get('credentials'),
        duration=__ret__.get('duration'),
        extension_headers=__ret__.get('extensionHeaders'),
        http_method=__ret__.get('httpMethod'),
        id=__ret__.get('id'),
        path=__ret__.get('path'),
        signed_url=__ret__.get('signedUrl'))
