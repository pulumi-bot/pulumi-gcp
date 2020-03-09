# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AutoscalingPolicy(pulumi.CustomResource):
    basic_algorithm: pulumi.Output[dict]
    """
    Basic algorithm for autoscaling.

      * `cooldownPeriod` (`str`)
      * `yarnConfig` (`dict`)
        * `gracefulDecommissionTimeout` (`str`)
        * `scaleDownFactor` (`float`)
        * `scaleDownMinWorkerFraction` (`float`)
        * `scaleUpFactor` (`float`)
        * `scaleUpMinWorkerFraction` (`float`)
    """
    location: pulumi.Output[str]
    """
    The location where the autoscaling poicy should reside. The default value is 'global'.
    """
    name: pulumi.Output[str]
    """
    The "resource name" of the autoscaling policy.
    """
    policy_id: pulumi.Output[str]
    """
    The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot
    begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    secondary_worker_config: pulumi.Output[dict]
    """
    Describes how the autoscaler will operate for secondary workers.

      * `maxInstances` (`float`)
      * `minInstances` (`float`)
      * `weight` (`float`)
    """
    worker_config: pulumi.Output[dict]
    """
    Describes how the autoscaler will operate for primary workers.

      * `maxInstances` (`float`)
      * `minInstances` (`float`)
      * `weight` (`float`)
    """
    def __init__(__self__, resource_name, opts=None, basic_algorithm=None, location=None, policy_id=None, project=None, secondary_worker_config=None, worker_config=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a AutoscalingPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] basic_algorithm: Basic algorithm for autoscaling.
        :param pulumi.Input[str] location: The location where the autoscaling poicy should reside. The default value is 'global'.
        :param pulumi.Input[str] policy_id: The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot
               begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] secondary_worker_config: Describes how the autoscaler will operate for secondary workers.
        :param pulumi.Input[dict] worker_config: Describes how the autoscaler will operate for primary workers.

        The **basic_algorithm** object supports the following:

          * `cooldownPeriod` (`pulumi.Input[str]`)
          * `yarnConfig` (`pulumi.Input[dict]`)
            * `gracefulDecommissionTimeout` (`pulumi.Input[str]`)
            * `scaleDownFactor` (`pulumi.Input[float]`)
            * `scaleDownMinWorkerFraction` (`pulumi.Input[float]`)
            * `scaleUpFactor` (`pulumi.Input[float]`)
            * `scaleUpMinWorkerFraction` (`pulumi.Input[float]`)

        The **secondary_worker_config** object supports the following:

          * `maxInstances` (`pulumi.Input[float]`)
          * `minInstances` (`pulumi.Input[float]`)
          * `weight` (`pulumi.Input[float]`)

        The **worker_config** object supports the following:

          * `maxInstances` (`pulumi.Input[float]`)
          * `minInstances` (`pulumi.Input[float]`)
          * `weight` (`pulumi.Input[float]`)
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

            __props__['basic_algorithm'] = basic_algorithm
            __props__['location'] = location
            if policy_id is None:
                raise TypeError("Missing required property 'policy_id'")
            __props__['policy_id'] = policy_id
            __props__['project'] = project
            __props__['secondary_worker_config'] = secondary_worker_config
            __props__['worker_config'] = worker_config
            __props__['name'] = None
        super(AutoscalingPolicy, __self__).__init__(
            'gcp:dataproc/autoscalingPolicy:AutoscalingPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, basic_algorithm=None, location=None, name=None, policy_id=None, project=None, secondary_worker_config=None, worker_config=None):
        """
        Get an existing AutoscalingPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] basic_algorithm: Basic algorithm for autoscaling.
        :param pulumi.Input[str] location: The location where the autoscaling poicy should reside. The default value is 'global'.
        :param pulumi.Input[str] name: The "resource name" of the autoscaling policy.
        :param pulumi.Input[str] policy_id: The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot
               begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] secondary_worker_config: Describes how the autoscaler will operate for secondary workers.
        :param pulumi.Input[dict] worker_config: Describes how the autoscaler will operate for primary workers.

        The **basic_algorithm** object supports the following:

          * `cooldownPeriod` (`pulumi.Input[str]`)
          * `yarnConfig` (`pulumi.Input[dict]`)
            * `gracefulDecommissionTimeout` (`pulumi.Input[str]`)
            * `scaleDownFactor` (`pulumi.Input[float]`)
            * `scaleDownMinWorkerFraction` (`pulumi.Input[float]`)
            * `scaleUpFactor` (`pulumi.Input[float]`)
            * `scaleUpMinWorkerFraction` (`pulumi.Input[float]`)

        The **secondary_worker_config** object supports the following:

          * `maxInstances` (`pulumi.Input[float]`)
          * `minInstances` (`pulumi.Input[float]`)
          * `weight` (`pulumi.Input[float]`)

        The **worker_config** object supports the following:

          * `maxInstances` (`pulumi.Input[float]`)
          * `minInstances` (`pulumi.Input[float]`)
          * `weight` (`pulumi.Input[float]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["basic_algorithm"] = basic_algorithm
        __props__["location"] = location
        __props__["name"] = name
        __props__["policy_id"] = policy_id
        __props__["project"] = project
        __props__["secondary_worker_config"] = secondary_worker_config
        __props__["worker_config"] = worker_config
        return AutoscalingPolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

