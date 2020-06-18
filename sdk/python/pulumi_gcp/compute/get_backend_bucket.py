# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetBackendBucketResult:
    """
    A collection of values returned by getBackendBucket.
    """
    def __init__(__self__, bucket_name=None, cdn_policies=None, creation_timestamp=None, description=None, enable_cdn=None, id=None, name=None, project=None, self_link=None):
        if bucket_name and not isinstance(bucket_name, str):
            raise TypeError("Expected argument 'bucket_name' to be a str")
        __self__.bucket_name = bucket_name
        """
        Cloud Storage bucket name.
        """
        if cdn_policies and not isinstance(cdn_policies, list):
            raise TypeError("Expected argument 'cdn_policies' to be a list")
        __self__.cdn_policies = cdn_policies
        """
        Cloud CDN configuration for this Backend Bucket. Structure is documented below.
        """
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        __self__.creation_timestamp = creation_timestamp
        """
        Creation timestamp in RFC3339 text format.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        An optional textual description of the resource; provided by the client when the resource is created.
        """
        if enable_cdn and not isinstance(enable_cdn, bool):
            raise TypeError("Expected argument 'enable_cdn' to be a bool")
        __self__.enable_cdn = enable_cdn
        """
        Whether Cloud CDN is enabled for this BackendBucket.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        __self__.self_link = self_link
        """
        The URI of the created resource.
        """
class AwaitableGetBackendBucketResult(GetBackendBucketResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackendBucketResult(
            bucket_name=self.bucket_name,
            cdn_policies=self.cdn_policies,
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            enable_cdn=self.enable_cdn,
            id=self.id,
            name=self.name,
            project=self.project,
            self_link=self.self_link)

def get_backend_bucket(name=None,project=None,opts=None):
    """
    Get information about a BackendBucket.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_backend_bucket = gcp.compute.get_backend_bucket(name="my-backend")
    ```


    :param str name: Name of the resource.
    :param str project: The ID of the project in which the resource belongs. If it
           is not provided, the provider project is used.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getBackendBucket:getBackendBucket', __args__, opts=opts).value

    return AwaitableGetBackendBucketResult(
        bucket_name=__ret__.get('bucketName'),
        cdn_policies=__ret__.get('cdnPolicies'),
        creation_timestamp=__ret__.get('creationTimestamp'),
        description=__ret__.get('description'),
        enable_cdn=__ret__.get('enableCdn'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        project=__ret__.get('project'),
        self_link=__ret__.get('selfLink'))
