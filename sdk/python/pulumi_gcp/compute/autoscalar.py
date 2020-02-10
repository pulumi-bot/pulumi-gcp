# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Autoscalar(pulumi.CustomResource):
    autoscaling_policy: pulumi.Output[dict]
    """
    The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an
    autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization. If none of these are specified, the
    default will be to autoscale based on cpuUtilization to 0.6 or 60%.

      * `cooldownPeriod` (`float`)
      * `cpuUtilization` (`dict`)
        * `target` (`float`)

      * `loadBalancingUtilization` (`dict`)
        * `target` (`float`)

      * `maxReplicas` (`float`)
      * `metrics` (`list`)
        * `filter` (`str`)
        * `name` (`str`)
        * `singleInstanceAssignment` (`float`)
        * `target` (`float`)
        * `type` (`str`)

      * `minReplicas` (`float`)
    """
    creation_timestamp: pulumi.Output[str]
    """
    Creation timestamp in RFC3339 text format.
    """
    description: pulumi.Output[str]
    """
    An optional description of this resource.
    """
    name: pulumi.Output[str]
    """
    Name of the resource. The name must be 1-63 characters long and match the regular expression
    '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
    must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
    """
    project: pulumi.Output[str]
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    target: pulumi.Output[str]
    """
    URL of the managed instance group that this autoscaler will scale.
    """
    zone: pulumi.Output[str]
    """
    URL of the zone where the instance group resides.
    """
    def __init__(__self__, resource_name, opts=None, autoscaling_policy=None, description=None, name=None, project=None, target=None, zone=None, __props__=None, __name__=None, __opts__=None):
        """
        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_autoscaler.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] autoscaling_policy: The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an
               autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization. If none of these are specified, the
               default will be to autoscale based on cpuUtilization to 0.6 or 60%.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] name: Name of the resource. The name must be 1-63 characters long and match the regular expression
               '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
               must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] target: URL of the managed instance group that this autoscaler will scale.
        :param pulumi.Input[str] zone: URL of the zone where the instance group resides.

        The **autoscaling_policy** object supports the following:

          * `cooldownPeriod` (`pulumi.Input[float]`)
          * `cpuUtilization` (`pulumi.Input[dict]`)
            * `target` (`pulumi.Input[float]`)

          * `loadBalancingUtilization` (`pulumi.Input[dict]`)
            * `target` (`pulumi.Input[float]`)

          * `maxReplicas` (`pulumi.Input[float]`)
          * `metrics` (`pulumi.Input[list]`)
            * `filter` (`pulumi.Input[str]`)
            * `name` (`pulumi.Input[str]`)
            * `singleInstanceAssignment` (`pulumi.Input[float]`)
            * `target` (`pulumi.Input[float]`)
            * `type` (`pulumi.Input[str]`)

          * `minReplicas` (`pulumi.Input[float]`)
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

            if autoscaling_policy is None:
                raise TypeError("Missing required property 'autoscaling_policy'")
            __props__['autoscaling_policy'] = autoscaling_policy
            __props__['description'] = description
            __props__['name'] = name
            __props__['project'] = project
            if target is None:
                raise TypeError("Missing required property 'target'")
            __props__['target'] = target
            __props__['zone'] = zone
            __props__['creation_timestamp'] = None
            __props__['self_link'] = None
        super(Autoscalar, __self__).__init__(
            'gcp:compute/autoscalar:Autoscalar',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, autoscaling_policy=None, creation_timestamp=None, description=None, name=None, project=None, self_link=None, target=None, zone=None):
        """
        Get an existing Autoscalar resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] autoscaling_policy: The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an
               autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization. If none of these are specified, the
               default will be to autoscale based on cpuUtilization to 0.6 or 60%.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] name: Name of the resource. The name must be 1-63 characters long and match the regular expression
               '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
               must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] target: URL of the managed instance group that this autoscaler will scale.
        :param pulumi.Input[str] zone: URL of the zone where the instance group resides.

        The **autoscaling_policy** object supports the following:

          * `cooldownPeriod` (`pulumi.Input[float]`)
          * `cpuUtilization` (`pulumi.Input[dict]`)
            * `target` (`pulumi.Input[float]`)

          * `loadBalancingUtilization` (`pulumi.Input[dict]`)
            * `target` (`pulumi.Input[float]`)

          * `maxReplicas` (`pulumi.Input[float]`)
          * `metrics` (`pulumi.Input[list]`)
            * `filter` (`pulumi.Input[str]`)
            * `name` (`pulumi.Input[str]`)
            * `singleInstanceAssignment` (`pulumi.Input[float]`)
            * `target` (`pulumi.Input[float]`)
            * `type` (`pulumi.Input[str]`)

          * `minReplicas` (`pulumi.Input[float]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["autoscaling_policy"] = autoscaling_policy
        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["name"] = name
        __props__["project"] = project
        __props__["self_link"] = self_link
        __props__["target"] = target
        __props__["zone"] = zone
        return Autoscalar(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

