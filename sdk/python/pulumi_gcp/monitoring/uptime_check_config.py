# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class UptimeCheckConfig(pulumi.CustomResource):
    content_matchers: pulumi.Output[list]
    """
    The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and
    other entries will be ignored. The server will look for an exact match of the string in the page response's content.
    This field is optional and should only be specified if a content match is required.

      * `content` (`str`)
    """
    display_name: pulumi.Output[str]
    """
    A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver
    Workspace in order to make it easier to identify; however, uniqueness is not enforced.
    """
    http_check: pulumi.Output[dict]
    """
    Contains information needed to make an HTTP or HTTPS check.

      * `authInfo` (`dict`)
        * `password` (`str`)
        * `username` (`str`)

      * `headers` (`dict`)
      * `maskHeaders` (`bool`)
      * `path` (`str`)
      * `port` (`float`)
      * `useSsl` (`bool`)
      * `validateSsl` (`bool`)
    """
    monitored_resource: pulumi.Output[dict]
    """
    The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The
    following monitored resource types are supported for uptime checks: uptime_url gce_instance gae_app aws_ec2_instance
    aws_elb_load_balancer

      * `labels` (`dict`)
      * `type` (`str`)
    """
    name: pulumi.Output[str]
    """
    A unique resource name for this UptimeCheckConfig. The format is
    projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].
    """
    period: pulumi.Output[str]
    """
    How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5
    minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    resource_group: pulumi.Output[dict]
    """
    The group resource associated with the configuration.

      * `groupId` (`str`)
      * `resourceType` (`str`)
    """
    selected_regions: pulumi.Output[list]
    """
    The list of regions from which the check will be run. Some regions contain one location, and others contain more than
    one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error
    message is returned. Not specifying this field will result in uptime checks running from all regions.
    """
    tcp_check: pulumi.Output[dict]
    """
    Contains information needed to make a TCP check.

      * `port` (`float`)
    """
    timeout: pulumi.Output[str]
    """
    The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats
    https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
    """
    uptime_check_id: pulumi.Output[str]
    """
    The id of the uptime check
    """
    def __init__(__self__, resource_name, opts=None, content_matchers=None, display_name=None, http_check=None, monitored_resource=None, period=None, project=None, resource_group=None, selected_regions=None, tcp_check=None, timeout=None, __props__=None, __name__=None, __opts__=None):
        """
        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/monitoring_uptime_check_config.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] content_matchers: The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and
               other entries will be ignored. The server will look for an exact match of the string in the page response's content.
               This field is optional and should only be specified if a content match is required.
        :param pulumi.Input[str] display_name: A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver
               Workspace in order to make it easier to identify; however, uniqueness is not enforced.
        :param pulumi.Input[dict] http_check: Contains information needed to make an HTTP or HTTPS check.
        :param pulumi.Input[dict] monitored_resource: The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The
               following monitored resource types are supported for uptime checks: uptime_url gce_instance gae_app aws_ec2_instance
               aws_elb_load_balancer
        :param pulumi.Input[str] period: How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5
               minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] resource_group: The group resource associated with the configuration.
        :param pulumi.Input[list] selected_regions: The list of regions from which the check will be run. Some regions contain one location, and others contain more than
               one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error
               message is returned. Not specifying this field will result in uptime checks running from all regions.
        :param pulumi.Input[dict] tcp_check: Contains information needed to make a TCP check.
        :param pulumi.Input[str] timeout: The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats
               https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration

        The **content_matchers** object supports the following:

          * `content` (`pulumi.Input[str]`)

        The **http_check** object supports the following:

          * `authInfo` (`pulumi.Input[dict]`)
            * `password` (`pulumi.Input[str]`)
            * `username` (`pulumi.Input[str]`)

          * `headers` (`pulumi.Input[dict]`)
          * `maskHeaders` (`pulumi.Input[bool]`)
          * `path` (`pulumi.Input[str]`)
          * `port` (`pulumi.Input[float]`)
          * `useSsl` (`pulumi.Input[bool]`)
          * `validateSsl` (`pulumi.Input[bool]`)

        The **monitored_resource** object supports the following:

          * `labels` (`pulumi.Input[dict]`)
          * `type` (`pulumi.Input[str]`)

        The **resource_group** object supports the following:

          * `groupId` (`pulumi.Input[str]`)
          * `resourceType` (`pulumi.Input[str]`)

        The **tcp_check** object supports the following:

          * `port` (`pulumi.Input[float]`)
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

            __props__['content_matchers'] = content_matchers
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['http_check'] = http_check
            __props__['monitored_resource'] = monitored_resource
            __props__['period'] = period
            __props__['project'] = project
            __props__['resource_group'] = resource_group
            __props__['selected_regions'] = selected_regions
            __props__['tcp_check'] = tcp_check
            if timeout is None:
                raise TypeError("Missing required property 'timeout'")
            __props__['timeout'] = timeout
            __props__['name'] = None
            __props__['uptime_check_id'] = None
        super(UptimeCheckConfig, __self__).__init__(
            'gcp:monitoring/uptimeCheckConfig:UptimeCheckConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, content_matchers=None, display_name=None, http_check=None, monitored_resource=None, name=None, period=None, project=None, resource_group=None, selected_regions=None, tcp_check=None, timeout=None, uptime_check_id=None):
        """
        Get an existing UptimeCheckConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] content_matchers: The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and
               other entries will be ignored. The server will look for an exact match of the string in the page response's content.
               This field is optional and should only be specified if a content match is required.
        :param pulumi.Input[str] display_name: A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver
               Workspace in order to make it easier to identify; however, uniqueness is not enforced.
        :param pulumi.Input[dict] http_check: Contains information needed to make an HTTP or HTTPS check.
        :param pulumi.Input[dict] monitored_resource: The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The
               following monitored resource types are supported for uptime checks: uptime_url gce_instance gae_app aws_ec2_instance
               aws_elb_load_balancer
        :param pulumi.Input[str] name: A unique resource name for this UptimeCheckConfig. The format is
               projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].
        :param pulumi.Input[str] period: How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5
               minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] resource_group: The group resource associated with the configuration.
        :param pulumi.Input[list] selected_regions: The list of regions from which the check will be run. Some regions contain one location, and others contain more than
               one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error
               message is returned. Not specifying this field will result in uptime checks running from all regions.
        :param pulumi.Input[dict] tcp_check: Contains information needed to make a TCP check.
        :param pulumi.Input[str] timeout: The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats
               https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
        :param pulumi.Input[str] uptime_check_id: The id of the uptime check

        The **content_matchers** object supports the following:

          * `content` (`pulumi.Input[str]`)

        The **http_check** object supports the following:

          * `authInfo` (`pulumi.Input[dict]`)
            * `password` (`pulumi.Input[str]`)
            * `username` (`pulumi.Input[str]`)

          * `headers` (`pulumi.Input[dict]`)
          * `maskHeaders` (`pulumi.Input[bool]`)
          * `path` (`pulumi.Input[str]`)
          * `port` (`pulumi.Input[float]`)
          * `useSsl` (`pulumi.Input[bool]`)
          * `validateSsl` (`pulumi.Input[bool]`)

        The **monitored_resource** object supports the following:

          * `labels` (`pulumi.Input[dict]`)
          * `type` (`pulumi.Input[str]`)

        The **resource_group** object supports the following:

          * `groupId` (`pulumi.Input[str]`)
          * `resourceType` (`pulumi.Input[str]`)

        The **tcp_check** object supports the following:

          * `port` (`pulumi.Input[float]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["content_matchers"] = content_matchers
        __props__["display_name"] = display_name
        __props__["http_check"] = http_check
        __props__["monitored_resource"] = monitored_resource
        __props__["name"] = name
        __props__["period"] = period
        __props__["project"] = project
        __props__["resource_group"] = resource_group
        __props__["selected_regions"] = selected_regions
        __props__["tcp_check"] = tcp_check
        __props__["timeout"] = timeout
        __props__["uptime_check_id"] = uptime_check_id
        return UptimeCheckConfig(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

