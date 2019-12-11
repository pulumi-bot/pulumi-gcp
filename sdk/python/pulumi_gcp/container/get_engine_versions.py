# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetEngineVersionsResult:
    """
    A collection of values returned by getEngineVersions.
    """
    def __init__(__self__, default_cluster_version=None, latest_master_version=None, latest_node_version=None, location=None, project=None, valid_master_versions=None, valid_node_versions=None, version_prefix=None, id=None):
        if default_cluster_version and not isinstance(default_cluster_version, str):
            raise TypeError("Expected argument 'default_cluster_version' to be a str")
        __self__.default_cluster_version = default_cluster_version
        """
        Version of Kubernetes the service deploys by default.
        """
        if latest_master_version and not isinstance(latest_master_version, str):
            raise TypeError("Expected argument 'latest_master_version' to be a str")
        __self__.latest_master_version = latest_master_version
        """
        The latest version available in the given zone for use with master instances.
        """
        if latest_node_version and not isinstance(latest_node_version, str):
            raise TypeError("Expected argument 'latest_node_version' to be a str")
        __self__.latest_node_version = latest_node_version
        """
        The latest version available in the given zone for use with node instances.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if valid_master_versions and not isinstance(valid_master_versions, list):
            raise TypeError("Expected argument 'valid_master_versions' to be a list")
        __self__.valid_master_versions = valid_master_versions
        """
        A list of versions available in the given zone for use with master instances.
        """
        if valid_node_versions and not isinstance(valid_node_versions, list):
            raise TypeError("Expected argument 'valid_node_versions' to be a list")
        __self__.valid_node_versions = valid_node_versions
        """
        A list of versions available in the given zone for use with node instances.
        """
        if version_prefix and not isinstance(version_prefix, str):
            raise TypeError("Expected argument 'version_prefix' to be a str")
        __self__.version_prefix = version_prefix
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetEngineVersionsResult(GetEngineVersionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEngineVersionsResult(
            default_cluster_version=self.default_cluster_version,
            latest_master_version=self.latest_master_version,
            latest_node_version=self.latest_node_version,
            location=self.location,
            project=self.project,
            valid_master_versions=self.valid_master_versions,
            valid_node_versions=self.valid_node_versions,
            version_prefix=self.version_prefix,
            id=self.id)

def get_engine_versions(location=None,project=None,version_prefix=None,opts=None):
    """
    ## a---
    
    subcategory: "Kubernetes (Container) Engine"
    layout: "google"
    page_title: "Google: container.getEngineVersions"
    sidebar_current: "docs-google-datasource-container-versions"
    description: |-
      Provides lists of available Google Kubernetes Engine versions for masters and nodes.
    ---
    
    # google\_container\_engine\_versions
    
    Provides access to available Google Kubernetes Engine versions in a zone or region for a given project.
    
    > If you are using the `container.getEngineVersions` datasource with a
    regional cluster, ensure that you have provided a region as the `location` to
    the datasource. A region can have a different set of supported versions than
    its component zones, and not all zones in a region are guaranteed to
    support the same version.
    
    :param str location: The location (region or zone) to list versions for.
           Must exactly match the location the cluster will be deployed in, or listed
           versions may not be available. If `location`, `region`, and `zone` are not
           specified, the provider-level zone must be set and is used instead.
    :param str project: ID of the project to list available cluster versions for. Should match the project the cluster will be deployed to.
           Defaults to the project that the provider is authenticated with.
    :param str version_prefix: If provided, this provider will only return versions
           that match the string prefix. For example, `1.11.` will match all `1.11` series
           releases. Since this is just a string match, it's recommended that you append a
           `.` after minor versions to ensure that prefixes such as `1.1` don't match
           versions like `1.12.5-gke.10` accidentally. See [the docs on versioning schema](https://cloud.google.com/kubernetes-engine/versioning-and-upgrades#versioning_scheme)
           for full details on how version strings are formatted.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/container_engine_versions.html.markdown.
    """
    __args__ = dict()

    __args__['location'] = location
    __args__['project'] = project
    __args__['versionPrefix'] = version_prefix
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:container/getEngineVersions:getEngineVersions', __args__, opts=opts).value

    return AwaitableGetEngineVersionsResult(
        default_cluster_version=__ret__.get('defaultClusterVersion'),
        latest_master_version=__ret__.get('latestMasterVersion'),
        latest_node_version=__ret__.get('latestNodeVersion'),
        location=__ret__.get('location'),
        project=__ret__.get('project'),
        valid_master_versions=__ret__.get('validMasterVersions'),
        valid_node_versions=__ret__.get('validNodeVersions'),
        version_prefix=__ret__.get('versionPrefix'),
        id=__ret__.get('id'))
