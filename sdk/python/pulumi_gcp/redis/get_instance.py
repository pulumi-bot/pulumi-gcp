# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class GetInstanceResult:
    """
    A collection of values returned by getInstance.
    """
    def __init__(__self__, alternative_location_id=None, authorized_network=None, connect_mode=None, create_time=None, current_location_id=None, display_name=None, host=None, id=None, labels=None, location_id=None, memory_size_gb=None, name=None, port=None, project=None, redis_configs=None, redis_version=None, region=None, reserved_ip_range=None, tier=None):
        if alternative_location_id and not isinstance(alternative_location_id, str):
            raise TypeError("Expected argument 'alternative_location_id' to be a str")
        __self__.alternative_location_id = alternative_location_id
        if authorized_network and not isinstance(authorized_network, str):
            raise TypeError("Expected argument 'authorized_network' to be a str")
        __self__.authorized_network = authorized_network
        if connect_mode and not isinstance(connect_mode, str):
            raise TypeError("Expected argument 'connect_mode' to be a str")
        __self__.connect_mode = connect_mode
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        __self__.create_time = create_time
        if current_location_id and not isinstance(current_location_id, str):
            raise TypeError("Expected argument 'current_location_id' to be a str")
        __self__.current_location_id = current_location_id
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        __self__.display_name = display_name
        if host and not isinstance(host, str):
            raise TypeError("Expected argument 'host' to be a str")
        __self__.host = host
        """
        Hostname or IP address of the exposed Redis endpoint used by clients
        to connect to the service.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        __self__.labels = labels
        if location_id and not isinstance(location_id, str):
            raise TypeError("Expected argument 'location_id' to be a str")
        __self__.location_id = location_id
        if memory_size_gb and not isinstance(memory_size_gb, float):
            raise TypeError("Expected argument 'memory_size_gb' to be a float")
        __self__.memory_size_gb = memory_size_gb
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if port and not isinstance(port, float):
            raise TypeError("Expected argument 'port' to be a float")
        __self__.port = port
        """
        The port number of the exposed Redis endpoint.
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if redis_configs and not isinstance(redis_configs, dict):
            raise TypeError("Expected argument 'redis_configs' to be a dict")
        __self__.redis_configs = redis_configs
        if redis_version and not isinstance(redis_version, str):
            raise TypeError("Expected argument 'redis_version' to be a str")
        __self__.redis_version = redis_version
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        if reserved_ip_range and not isinstance(reserved_ip_range, str):
            raise TypeError("Expected argument 'reserved_ip_range' to be a str")
        __self__.reserved_ip_range = reserved_ip_range
        if tier and not isinstance(tier, str):
            raise TypeError("Expected argument 'tier' to be a str")
        __self__.tier = tier


class AwaitableGetInstanceResult(GetInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceResult(
            alternative_location_id=self.alternative_location_id,
            authorized_network=self.authorized_network,
            connect_mode=self.connect_mode,
            create_time=self.create_time,
            current_location_id=self.current_location_id,
            display_name=self.display_name,
            host=self.host,
            id=self.id,
            labels=self.labels,
            location_id=self.location_id,
            memory_size_gb=self.memory_size_gb,
            name=self.name,
            port=self.port,
            project=self.project,
            redis_configs=self.redis_configs,
            redis_version=self.redis_version,
            region=self.region,
            reserved_ip_range=self.reserved_ip_range,
            tier=self.tier)


def get_instance(name=None, project=None, region=None, opts=None):
    """
    Get information about a Google Cloud Redis instance. For more information see
    the [official documentation](https://cloud.google.com/memorystore/docs/redis)
    and [API](https://cloud.google.com/memorystore/docs/redis/apis).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    default = gcp.redis.get_instance(name="my-redis-instance")
    ```


    :param str name: The name of a Redis instance.
    :param str project: The project in which the resource belongs. If it
           is not provided, the provider project is used.
    :param str region: The region in which the resource belongs. If it
           is not provided, the provider region is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:redis/getInstance:getInstance', __args__, opts=opts).value

    return AwaitableGetInstanceResult(
        alternative_location_id=__ret__.get('alternativeLocationId'),
        authorized_network=__ret__.get('authorizedNetwork'),
        connect_mode=__ret__.get('connectMode'),
        create_time=__ret__.get('createTime'),
        current_location_id=__ret__.get('currentLocationId'),
        display_name=__ret__.get('displayName'),
        host=__ret__.get('host'),
        id=__ret__.get('id'),
        labels=__ret__.get('labels'),
        location_id=__ret__.get('locationId'),
        memory_size_gb=__ret__.get('memorySizeGb'),
        name=__ret__.get('name'),
        port=__ret__.get('port'),
        project=__ret__.get('project'),
        redis_configs=__ret__.get('redisConfigs'),
        redis_version=__ret__.get('redisVersion'),
        region=__ret__.get('region'),
        reserved_ip_range=__ret__.get('reservedIpRange'),
        tier=__ret__.get('tier'))
