# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'GetRegistryImageResult',
    'AwaitableGetRegistryImageResult',
    'get_registry_image',
]

@pulumi.output_type
class GetRegistryImageResult:
    """
    A collection of values returned by getRegistryImage.
    """
    def __init__(__self__, digest=None, id=None, image_url=None, name=None, project=None, region=None, tag=None):
        if digest and not isinstance(digest, str):
            raise TypeError("Expected argument 'digest' to be a str")
        pulumi.set(__self__, "digest", digest)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_url and not isinstance(image_url, str):
            raise TypeError("Expected argument 'image_url' to be a str")
        pulumi.set(__self__, "image_url", image_url)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        pulumi.set(__self__, "tag", tag)

    @property
    @pulumi.getter
    def digest(self) -> Optional[str]:
        return pulumi.get(self, "digest")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageUrl")
    def image_url(self) -> str:
        return pulumi.get(self, "image_url")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> str:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def tag(self) -> Optional[str]:
        return pulumi.get(self, "tag")


class AwaitableGetRegistryImageResult(GetRegistryImageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegistryImageResult(
            digest=self.digest,
            id=self.id,
            image_url=self.image_url,
            name=self.name,
            project=self.project,
            region=self.region,
            tag=self.tag)


def get_registry_image(digest: Optional[str] = None,
                       name: Optional[str] = None,
                       project: Optional[str] = None,
                       region: Optional[str] = None,
                       tag: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegistryImageResult:
    """
    This data source fetches the project name, and provides the appropriate URLs to use for container registry for this project.

    The URLs are computed entirely offline - as long as the project exists, they will be valid, but this data source does not contact Google Container Registry (GCR) at any point.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    debian = gcp.container.get_registry_image(name="debian")
    pulumi.export("gcrLocation", debian.image_url)
    ```
    """
    __args__ = dict()
    __args__['digest'] = digest
    __args__['name'] = name
    __args__['project'] = project
    __args__['region'] = region
    __args__['tag'] = tag
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:container/getRegistryImage:getRegistryImage', __args__, opts=opts, typ=GetRegistryImageResult).value

    return AwaitableGetRegistryImageResult(
        digest=__ret__.digest,
        id=__ret__.id,
        image_url=__ret__.image_url,
        name=__ret__.name,
        project=__ret__.project,
        region=__ret__.region,
        tag=__ret__.tag)
