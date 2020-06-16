# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class TargetPool(pulumi.CustomResource):
    backup_pool: pulumi.Output[str]
    """
    URL to the backup target pool. Must also set
    failover\_ratio.
    """
    description: pulumi.Output[str]
    """
    Textual description field.
    """
    failover_ratio: pulumi.Output[float]
    """
    Ratio (0 to 1) of failed nodes before using the
    backup pool (which must also be set).
    """
    health_checks: pulumi.Output[str]
    """
    List of zero or one health check name or self_link. Only
    legacy `compute.HttpHealthCheck` is supported.
    """
    instances: pulumi.Output[list]
    """
    List of instances in the pool. They can be given as
    URLs, or in the form of "zone/name". Note that the instances need not exist
    at the time of target pool creation, so there is no need to use the
    interpolation to create a dependency on the instances from the
    target pool.
    """
    name: pulumi.Output[str]
    """
    A unique name for the resource, required by GCE. Changing
    this forces a new resource to be created.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    Where the target pool resides. Defaults to project
    region.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    session_affinity: pulumi.Output[str]
    """
    How to distribute load. Options are "NONE" (no
    affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and
    "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE").
    """
    def __init__(__self__, resource_name, opts=None, backup_pool=None, description=None, failover_ratio=None, health_checks=None, instances=None, name=None, project=None, region=None, session_affinity=None, __props__=None, __name__=None, __opts__=None):
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
        :param pulumi.Input[list] instances: List of instances in the pool. They can be given as
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, backup_pool=None, description=None, failover_ratio=None, health_checks=None, instances=None, name=None, project=None, region=None, self_link=None, session_affinity=None):
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
        :param pulumi.Input[list] instances: List of instances in the pool. They can be given as
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
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
