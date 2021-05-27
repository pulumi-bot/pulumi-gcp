# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetEngineVersionsResult',
    'AwaitableGetEngineVersionsResult',
    'get_engine_versions',
]

@pulumi.output_type
class GetEngineVersionsResult:
    """
    A collection of values returned by getEngineVersions.
    """
    def __init__(__self__, default_cluster_version=None, id=None, latest_master_version=None, latest_node_version=None, location=None, project=None, release_channel_default_version=None, valid_master_versions=None, valid_node_versions=None, version_prefix=None):
        if default_cluster_version and not isinstance(default_cluster_version, str):
            raise TypeError("Expected argument 'default_cluster_version' to be a str")
        pulumi.set(__self__, "default_cluster_version", default_cluster_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if latest_master_version and not isinstance(latest_master_version, str):
            raise TypeError("Expected argument 'latest_master_version' to be a str")
        pulumi.set(__self__, "latest_master_version", latest_master_version)
        if latest_node_version and not isinstance(latest_node_version, str):
            raise TypeError("Expected argument 'latest_node_version' to be a str")
        pulumi.set(__self__, "latest_node_version", latest_node_version)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if release_channel_default_version and not isinstance(release_channel_default_version, dict):
            raise TypeError("Expected argument 'release_channel_default_version' to be a dict")
        pulumi.set(__self__, "release_channel_default_version", release_channel_default_version)
        if valid_master_versions and not isinstance(valid_master_versions, list):
            raise TypeError("Expected argument 'valid_master_versions' to be a list")
        pulumi.set(__self__, "valid_master_versions", valid_master_versions)
        if valid_node_versions and not isinstance(valid_node_versions, list):
            raise TypeError("Expected argument 'valid_node_versions' to be a list")
        pulumi.set(__self__, "valid_node_versions", valid_node_versions)
        if version_prefix and not isinstance(version_prefix, str):
            raise TypeError("Expected argument 'version_prefix' to be a str")
        pulumi.set(__self__, "version_prefix", version_prefix)

    @property
    @pulumi.getter(name="defaultClusterVersion")
    def default_cluster_version(self) -> str:
        """
        Version of Kubernetes the service deploys by default.
        """
        return pulumi.get(self, "default_cluster_version")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="latestMasterVersion")
    def latest_master_version(self) -> str:
        """
        The latest version available in the given zone for use with master instances.
        """
        return pulumi.get(self, "latest_master_version")

    @property
    @pulumi.getter(name="latestNodeVersion")
    def latest_node_version(self) -> str:
        """
        The latest version available in the given zone for use with node instances.
        """
        return pulumi.get(self, "latest_node_version")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="releaseChannelDefaultVersion")
    def release_channel_default_version(self) -> Mapping[str, str]:
        """
        A map from a release channel name to the channel's default version.
        """
        return pulumi.get(self, "release_channel_default_version")

    @property
    @pulumi.getter(name="validMasterVersions")
    def valid_master_versions(self) -> Sequence[str]:
        """
        A list of versions available in the given zone for use with master instances.
        """
        return pulumi.get(self, "valid_master_versions")

    @property
    @pulumi.getter(name="validNodeVersions")
    def valid_node_versions(self) -> Sequence[str]:
        """
        A list of versions available in the given zone for use with node instances.
        """
        return pulumi.get(self, "valid_node_versions")

    @property
    @pulumi.getter(name="versionPrefix")
    def version_prefix(self) -> Optional[str]:
        return pulumi.get(self, "version_prefix")


class AwaitableGetEngineVersionsResult(GetEngineVersionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEngineVersionsResult(
            default_cluster_version=self.default_cluster_version,
            id=self.id,
            latest_master_version=self.latest_master_version,
            latest_node_version=self.latest_node_version,
            location=self.location,
            project=self.project,
            release_channel_default_version=self.release_channel_default_version,
            valid_master_versions=self.valid_master_versions,
            valid_node_versions=self.valid_node_versions,
            version_prefix=self.version_prefix)


def get_engine_versions(location: Optional[str] = None,
                        project: Optional[str] = None,
                        version_prefix: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEngineVersionsResult:
    """
    Provides access to available Google Kubernetes Engine versions in a zone or region for a given project.

    > If you are using the `container.getEngineVersions` datasource with a
    regional cluster, ensure that you have provided a region as the `location` to
    the datasource. A region can have a different set of supported versions than
    its component zones, and not all zones in a region are guaranteed to
    support the same version.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    central1b = gcp.container.get_engine_versions(location="us-central1-b",
        version_prefix="1.12.")
    foo = gcp.container.Cluster("foo",
        location="us-central1-b",
        node_version=central1b.latest_node_version,
        initial_node_count=1,
        master_auth=gcp.container.ClusterMasterAuthArgs(
            username="mr.yoda",
            password="adoy.rm",
        ))
    pulumi.export("stableChannelVersion", central1b.release_channel_default_version["STABLE"])
    ```


    :param str location: The location (region or zone) to list versions for.
           Must exactly match the location the cluster will be deployed in, or listed
           versions may not be available. If `location`, `region`, and `zone` are not
           specified, the provider-level zone must be set and is used instead.
    :param str project: ID of the project to list available cluster versions for. Should match the project the cluster will be deployed to.
           Defaults to the project that the provider is authenticated with.
    :param str version_prefix: If provided, the provider will only return versions
           that match the string prefix. For example, `1.11.` will match all `1.11` series
           releases. Since this is just a string match, it's recommended that you append a
           `.` after minor versions to ensure that prefixes such as `1.1` don't match
           versions like `1.12.5-gke.10` accidentally. See [the docs on versioning schema](https://cloud.google.com/kubernetes-engine/versioning-and-upgrades#versioning_scheme)
           for full details on how version strings are formatted.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['versionPrefix'] = version_prefix
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:container/getEngineVersions:getEngineVersions', __args__, opts=opts, typ=GetEngineVersionsResult).value

    return AwaitableGetEngineVersionsResult(
        default_cluster_version=__ret__.default_cluster_version,
        id=__ret__.id,
        latest_master_version=__ret__.latest_master_version,
        latest_node_version=__ret__.latest_node_version,
        location=__ret__.location,
        project=__ret__.project,
        release_channel_default_version=__ret__.release_channel_default_version,
        valid_master_versions=__ret__.valid_master_versions,
        valid_node_versions=__ret__.valid_node_versions,
        version_prefix=__ret__.version_prefix)


@_utilities.lift_output_func(get_engine_versions)
def get_engine_versions_apply(location: Optional[pulumi.Input[str]] = None,
                              project: Optional[pulumi.Input[str]] = None,
                              version_prefix: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEngineVersionsResult]:
    ...
