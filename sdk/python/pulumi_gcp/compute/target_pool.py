# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['TargetPool']


class TargetPool(pulumi.CustomResource):
    backup_pool: pulumi.Output[Optional[str]] = pulumi.property("backupPool")
    """
    URL to the backup target pool. Must also set
    failover\_ratio.
    """

    description: pulumi.Output[Optional[str]] = pulumi.property("description")
    """
    Textual description field.
    """

    failover_ratio: pulumi.Output[Optional[float]] = pulumi.property("failoverRatio")
    """
    Ratio (0 to 1) of failed nodes before using the
    backup pool (which must also be set).
    """

    health_checks: pulumi.Output[Optional[str]] = pulumi.property("healthChecks")
    """
    List of zero or one health check name or self_link. Only
    legacy `compute.HttpHealthCheck` is supported.
    """

    instances: pulumi.Output[List[str]] = pulumi.property("instances")
    """
    List of instances in the pool. They can be given as
    URLs, or in the form of "zone/name". Note that the instances need not exist
    at the time of target pool creation, so there is no need to use the
    interpolation to create a dependency on the instances from the
    target pool.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    A unique name for the resource, required by GCE. Changing
    this forces a new resource to be created.
    """

    project: pulumi.Output[str] = pulumi.property("project")
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """

    region: pulumi.Output[str] = pulumi.property("region")
    """
    Where the target pool resides. Defaults to project
    region.
    """

    self_link: pulumi.Output[str] = pulumi.property("selfLink")
    """
    The URI of the created resource.
    """

    session_affinity: pulumi.Output[Optional[str]] = pulumi.property("sessionAffinity")
    """
    How to distribute load. Options are "NONE" (no
    affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and
    "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE").
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_pool: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 failover_ratio: Optional[pulumi.Input[float]] = None,
                 health_checks: Optional[pulumi.Input[str]] = None,
                 instances: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 session_affinity: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Target Pool within GCE. This is a collection of instances used as
        target of a network load balancer (Forwarding Rule). For more information see
        [the official
        documentation](https://cloud.google.com/compute/docs/load-balancing/network/target-pools)
        and [API](https://cloud.google.com/compute/docs/reference/latest/targetPools).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_http_health_check = gcp.compute.HttpHealthCheck("defaultHttpHealthCheck",
            request_path="/",
            check_interval_sec=1,
            timeout_sec=1)
        default_target_pool = gcp.compute.TargetPool("defaultTargetPool",
            instances=[
                "us-central1-a/myinstance1",
                "us-central1-b/myinstance2",
            ],
            health_checks=[default_http_health_check.name])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_pool: URL to the backup target pool. Must also set
               failover\_ratio.
        :param pulumi.Input[str] description: Textual description field.
        :param pulumi.Input[float] failover_ratio: Ratio (0 to 1) of failed nodes before using the
               backup pool (which must also be set).
        :param pulumi.Input[str] health_checks: List of zero or one health check name or self_link. Only
               legacy `compute.HttpHealthCheck` is supported.
        :param pulumi.Input[List[pulumi.Input[str]]] instances: List of instances in the pool. They can be given as
               URLs, or in the form of "zone/name". Note that the instances need not exist
               at the time of target pool creation, so there is no need to use the
               interpolation to create a dependency on the instances from the
               target pool.
        :param pulumi.Input[str] name: A unique name for the resource, required by GCE. Changing
               this forces a new resource to be created.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] region: Where the target pool resides. Defaults to project
               region.
        :param pulumi.Input[str] session_affinity: How to distribute load. Options are "NONE" (no
               affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and
               "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE").
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['backup_pool'] = backup_pool
            __props__['description'] = description
            __props__['failover_ratio'] = failover_ratio
            __props__['health_checks'] = health_checks
            __props__['instances'] = instances
            __props__['name'] = name
            __props__['project'] = project
            __props__['region'] = region
            __props__['session_affinity'] = session_affinity
            __props__['self_link'] = None
        super(TargetPool, __self__).__init__(
            'gcp:compute/targetPool:TargetPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_pool: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            failover_ratio: Optional[pulumi.Input[float]] = None,
            health_checks: Optional[pulumi.Input[str]] = None,
            instances: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            session_affinity: Optional[pulumi.Input[str]] = None) -> 'TargetPool':
        """
        Get an existing TargetPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_pool: URL to the backup target pool. Must also set
               failover\_ratio.
        :param pulumi.Input[str] description: Textual description field.
        :param pulumi.Input[float] failover_ratio: Ratio (0 to 1) of failed nodes before using the
               backup pool (which must also be set).
        :param pulumi.Input[str] health_checks: List of zero or one health check name or self_link. Only
               legacy `compute.HttpHealthCheck` is supported.
        :param pulumi.Input[List[pulumi.Input[str]]] instances: List of instances in the pool. They can be given as
               URLs, or in the form of "zone/name". Note that the instances need not exist
               at the time of target pool creation, so there is no need to use the
               interpolation to create a dependency on the instances from the
               target pool.
        :param pulumi.Input[str] name: A unique name for the resource, required by GCE. Changing
               this forces a new resource to be created.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] region: Where the target pool resides. Defaults to project
               region.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] session_affinity: How to distribute load. Options are "NONE" (no
               affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and
               "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE").
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backup_pool"] = backup_pool
        __props__["description"] = description
        __props__["failover_ratio"] = failover_ratio
        __props__["health_checks"] = health_checks
        __props__["instances"] = instances
        __props__["name"] = name
        __props__["project"] = project
        __props__["region"] = region
        __props__["self_link"] = self_link
        __props__["session_affinity"] = session_affinity
        return TargetPool(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

