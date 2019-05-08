# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetEngineVersionsResult:
    """
    A collection of values returned by getEngineVersions.
    """
    def __init__(__self__, default_cluster_version=None, latest_master_version=None, latest_node_version=None, location=None, project=None, region=None, valid_master_versions=None, valid_node_versions=None, version_prefix=None, zone=None, id=None):
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
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
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
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        __self__.zone = zone
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_engine_versions(location=None,project=None,region=None,version_prefix=None,zone=None,opts=None):
    """
    Provides access to available Google Kubernetes Engine versions in a zone or region for a given project.
    
    > If you are using the `google_container_engine_versions` datasource with a
    regional cluster, ensure that you have provided a region as the `location` to
    the datasource. A region can have a different set of supported versions than
    its component zones, and not all zones in a region are guaranteed to
    support the same version.
    """
    __args__ = dict()

    __args__['location'] = location
    __args__['project'] = project
    __args__['region'] = region
    __args__['versionPrefix'] = version_prefix
    __args__['zone'] = zone
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = await pulumi.runtime.invoke('gcp:container/getEngineVersions:getEngineVersions', __args__, opts=opts)

    return GetEngineVersionsResult(
        default_cluster_version=__ret__.get('defaultClusterVersion'),
        latest_master_version=__ret__.get('latestMasterVersion'),
        latest_node_version=__ret__.get('latestNodeVersion'),
        location=__ret__.get('location'),
        project=__ret__.get('project'),
        region=__ret__.get('region'),
        valid_master_versions=__ret__.get('validMasterVersions'),
        valid_node_versions=__ret__.get('validNodeVersions'),
        version_prefix=__ret__.get('versionPrefix'),
        zone=__ret__.get('zone'),
        id=__ret__.get('id'))
