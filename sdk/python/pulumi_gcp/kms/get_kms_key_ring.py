# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetKMSKeyRingResult:
    """
    A collection of values returned by getKMSKeyRing.
    """
    def __init__(__self__, id=None, location=None, name=None, project=None, self_link=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
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
        The self link of the created KeyRing. Its format is `projects/{projectId}/locations/{location}/keyRings/{keyRingName}`.
        """
class AwaitableGetKMSKeyRingResult(GetKMSKeyRingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKMSKeyRingResult(
            id=self.id,
            location=self.location,
            name=self.name,
            project=self.project,
            self_link=self.self_link)

def get_kms_key_ring(location=None,name=None,project=None,opts=None):
    """
    Provides access to Google Cloud Platform KMS KeyRing. For more information see
    [the official documentation](https://cloud.google.com/kms/docs/object-hierarchy#key_ring)
    and
    [API](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings).

    A KeyRing is a grouping of CryptoKeys for organizational purposes. A KeyRing belongs to a Google Cloud Platform Project
    and resides in a specific location.

    {{% examples %}}
    ## Example Usage
    {{% example %}}

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_key_ring = gcp.kms.get_kms_key_ring(location="us-central1",
        name="my-key-ring")
    ```
    {{% /example %}}
    {{% /examples %}}


    :param str location: The Google Cloud Platform location for the KeyRing.
           A full list of valid locations can be found by running `gcloud kms locations list`.
    :param str name: The KeyRing's name.
           A KeyRing name must exist within the provided location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
    :param str project: The project in which the resource belongs. If it
           is not provided, the provider project is used.
    """
    __args__ = dict()


    __args__['location'] = location
    __args__['name'] = name
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:kms/getKMSKeyRing:getKMSKeyRing', __args__, opts=opts).value

    return AwaitableGetKMSKeyRingResult(
        id=__ret__.get('id'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        project=__ret__.get('project'),
        self_link=__ret__.get('selfLink'))
