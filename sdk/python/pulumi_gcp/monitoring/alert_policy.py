# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class AlertPolicy(pulumi.CustomResource):
    combiner: pulumi.Output[str]
    conditions: pulumi.Output[list]
    creation_record: pulumi.Output[dict]
    display_name: pulumi.Output[str]
    enabled: pulumi.Output[bool]
    labels: pulumi.Output[list]
    name: pulumi.Output[str]
    notification_channels: pulumi.Output[list]
    project: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, combiner=None, conditions=None, display_name=None, enabled=None, labels=None, notification_channels=None, project=None, __name__=None, __opts__=None):
        """
        A description of the conditions under which some aspect of your system is
        considered to be "unhealthy" and the ways to notify people or services
        about this state.
        
        
        To get more information about AlertPolicy, see:
        
        * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.alertPolicies)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/monitoring/alerts/)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if combiner is None:
            raise TypeError('Missing required property combiner')
        __props__['combiner'] = combiner

        if conditions is None:
            raise TypeError('Missing required property conditions')
        __props__['conditions'] = conditions

        if display_name is None:
            raise TypeError('Missing required property display_name')
        __props__['display_name'] = display_name

        if enabled is None:
            raise TypeError('Missing required property enabled')
        __props__['enabled'] = enabled

        __props__['labels'] = labels

        __props__['notification_channels'] = notification_channels

        __props__['project'] = project

        __props__['creation_record'] = None
        __props__['name'] = None

        super(AlertPolicy, __self__).__init__(
            'gcp:monitoring/alertPolicy:AlertPolicy',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

