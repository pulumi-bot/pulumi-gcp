# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['RegionAutoscaler']


class RegionAutoscaler(pulumi.CustomResource):
    autoscaling_policy: pulumi.Output['outputs.RegionAutoscalerAutoscalingPolicy'] = pulumi.property("autoscalingPolicy")
    """
    The configuration parameters for the autoscaling algorithm. You can
    define one or more of the policies for an autoscaler: cpuUtilization,
    customMetricUtilizations, and loadBalancingUtilization.
    If none of these are specified, the default will be to autoscale based
    on cpuUtilization to 0.6 or 60%.  Structure is documented below.
    """

    creation_timestamp: pulumi.Output[str] = pulumi.property("creationTimestamp")
    """
    Creation timestamp in RFC3339 text format.
    """

    description: pulumi.Output[Optional[str]] = pulumi.property("description")
    """
    An optional description of this resource.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    The identifier (type) of the Stackdriver Monitoring metric.
    The metric cannot have negative values.
    The metric must have a value type of INT64 or DOUBLE.
    """

    project: pulumi.Output[str] = pulumi.property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """

    region: pulumi.Output[str] = pulumi.property("region")
    """
    URL of the region where the instance group resides.
    """

    self_link: pulumi.Output[str] = pulumi.property("selfLink")
    """
    The URI of the created resource.
    """

    target: pulumi.Output[str] = pulumi.property("target")
    """
    Fraction of backend capacity utilization (set in HTTP(s) load
    balancing configuration) that autoscaler should maintain. Must
    be a positive float value. If not defined, the default is 0.8.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscaling_policy: Optional[pulumi.Input[pulumi.InputType['RegionAutoscalerAutoscalingPolicyArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Represents an Autoscaler resource.

        Autoscalers allow you to automatically scale virtual machine instances in
        managed instance groups according to an autoscaling policy that you
        define.

        To get more information about RegionAutoscaler, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/regionAutoscalers)
        * How-to Guides
            * [Autoscaling Groups of Instances](https://cloud.google.com/compute/docs/autoscaler/)

        ## Example Usage
        ### Region Autoscaler Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        debian9 = gcp.compute.get_image(family="debian-9",
            project="debian-cloud")
        foobar_instance_template = gcp.compute.InstanceTemplate("foobarInstanceTemplate",
            machine_type="n1-standard-1",
            can_ip_forward=False,
            tags=[
                "foo",
                "bar",
            ],
            disks=[{
                "sourceImage": debian9.id,
            }],
            network_interfaces=[{
                "network": "default",
            }],
            metadata={
                "foo": "bar",
            },
            service_account={
                "scopes": [
                    "userinfo-email",
                    "compute-ro",
                    "storage-ro",
                ],
            })
        foobar_target_pool = gcp.compute.TargetPool("foobarTargetPool")
        foobar_region_instance_group_manager = gcp.compute.RegionInstanceGroupManager("foobarRegionInstanceGroupManager",
            region="us-central1",
            versions=[{
                "instanceTemplate": foobar_instance_template.id,
                "name": "primary",
            }],
            target_pools=[foobar_target_pool.id],
            base_instance_name="foobar")
        foobar_region_autoscaler = gcp.compute.RegionAutoscaler("foobarRegionAutoscaler",
            region="us-central1",
            target=foobar_region_instance_group_manager.id,
            autoscaling_policy={
                "maxReplicas": 5,
                "minReplicas": 1,
                "cooldownPeriod": 60,
                "cpuUtilization": {
                    "target": 0.5,
                },
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['RegionAutoscalerAutoscalingPolicyArgs']] autoscaling_policy: The configuration parameters for the autoscaling algorithm. You can
               define one or more of the policies for an autoscaler: cpuUtilization,
               customMetricUtilizations, and loadBalancingUtilization.
               If none of these are specified, the default will be to autoscale based
               on cpuUtilization to 0.6 or 60%.  Structure is documented below.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] name: The identifier (type) of the Stackdriver Monitoring metric.
               The metric cannot have negative values.
               The metric must have a value type of INT64 or DOUBLE.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: URL of the region where the instance group resides.
        :param pulumi.Input[str] target: Fraction of backend capacity utilization (set in HTTP(s) load
               balancing configuration) that autoscaler should maintain. Must
               be a positive float value. If not defined, the default is 0.8.
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

            if autoscaling_policy is None:
                raise TypeError("Missing required property 'autoscaling_policy'")
            __props__['autoscaling_policy'] = autoscaling_policy
            __props__['description'] = description
            __props__['name'] = name
            __props__['project'] = project
            __props__['region'] = region
            if target is None:
                raise TypeError("Missing required property 'target'")
            __props__['target'] = target
            __props__['creation_timestamp'] = None
            __props__['self_link'] = None
        super(RegionAutoscaler, __self__).__init__(
            'gcp:compute/regionAutoscaler:RegionAutoscaler',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            autoscaling_policy: Optional[pulumi.Input[pulumi.InputType['RegionAutoscalerAutoscalingPolicyArgs']]] = None,
            creation_timestamp: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            target: Optional[pulumi.Input[str]] = None) -> 'RegionAutoscaler':
        """
        Get an existing RegionAutoscaler resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['RegionAutoscalerAutoscalingPolicyArgs']] autoscaling_policy: The configuration parameters for the autoscaling algorithm. You can
               define one or more of the policies for an autoscaler: cpuUtilization,
               customMetricUtilizations, and loadBalancingUtilization.
               If none of these are specified, the default will be to autoscale based
               on cpuUtilization to 0.6 or 60%.  Structure is documented below.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] name: The identifier (type) of the Stackdriver Monitoring metric.
               The metric cannot have negative values.
               The metric must have a value type of INT64 or DOUBLE.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: URL of the region where the instance group resides.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] target: Fraction of backend capacity utilization (set in HTTP(s) load
               balancing configuration) that autoscaler should maintain. Must
               be a positive float value. If not defined, the default is 0.8.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["autoscaling_policy"] = autoscaling_policy
        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["name"] = name
        __props__["project"] = project
        __props__["region"] = region
        __props__["self_link"] = self_link
        __props__["target"] = target
        return RegionAutoscaler(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

