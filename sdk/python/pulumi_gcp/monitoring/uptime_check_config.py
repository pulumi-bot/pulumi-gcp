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
    The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.  Structure is documented below.

      * `content` (`str`) - String or regex content to match (max 1024 bytes)
    """
    display_name: pulumi.Output[str]
    """
    A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
    """
    http_check: pulumi.Output[dict]
    """
    Contains information needed to make an HTTP or HTTPS check.  Structure is documented below.

      * `authInfo` (`dict`) - The authentication information. Optional when creating an HTTP check; defaults to empty.  Structure is documented below.
        * `password` (`str`) - The password to authenticate.
        * `username` (`str`) - The username to authenticate.

      * `headers` (`dict`) - The list of headers to send as part of the uptime check request. If two headers have the same key and different values, they should be entered as a single header, with the value being a comma-separated list of all the desired values as described at https://www.w3.org/Protocols/rfc2616/rfc2616.txt (page 31). Entering two separate headers with the same key in a Create call will cause the first to be overwritten by the second. The maximum number of headers allowed is 100.
      * `maskHeaders` (`bool`) - Boolean specifying whether to encrypt the header information. Encryption should be specified for any headers related to authentication that you do not wish to be seen when retrieving the configuration. The server will be responsible for encrypting the headers. On Get/List calls, if mask_headers is set to True then the headers will be obscured with ******.
      * `path` (`str`) - The path to the page to run the check against. Will be combined with the host (specified within the MonitoredResource) and port to construct the full URL. Optional (defaults to "/").
      * `port` (`float`) - The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) to construct the full URL.
      * `useSsl` (`bool`) - If true, use HTTPS instead of HTTP to run the check.
      * `validateSsl` (`bool`) - Boolean specifying whether to include SSL certificate validation as a part of the Uptime check. Only applies to checks where monitoredResource is set to uptime_url. If useSsl is false, setting validateSsl to true has no effect.
    """
    monitored_resource: pulumi.Output[dict]
    """
    The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptime_url  gce_instance  gae_app  aws_ec2_instance  aws_elb_load_balancer  Structure is documented below.

      * `labels` (`dict`) - Values for all of the labels listed in the associated monitored resource descriptor. For example, Compute Engine VM instances use the labels "project_id", "instance_id", and "zone".
      * `type` (`str`) - The monitored resource type. This field must match the type field of a MonitoredResourceDescriptor (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.monitoredResourceDescriptors#MonitoredResourceDescriptor) object. For example, the type of a Compute Engine VM instance is gce_instance. For a list of types, see Monitoring resource types (https://cloud.google.com/monitoring/api/resources) and Logging resource types (https://cloud.google.com/logging/docs/api/v2/resource-list).
    """
    name: pulumi.Output[str]
    """
    A unique resource name for this UptimeCheckConfig. The format is
    projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].
    """
    period: pulumi.Output[str]
    """
    How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    resource_group: pulumi.Output[dict]
    """
    The group resource associated with the configuration.  Structure is documented below.

      * `groupId` (`str`) - The group of resources being monitored. Should be the `name` of a group
      * `resourceType` (`str`) - The resource type of the group members.
    """
    selected_regions: pulumi.Output[list]
    """
    The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
    """
    tcp_check: pulumi.Output[dict]
    """
    Contains information needed to make a TCP check.  Structure is documented below.

      * `port` (`float`) - The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) to construct the full URL.
    """
    timeout: pulumi.Output[str]
    """
    The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
    """
    uptime_check_id: pulumi.Output[str]
    """
    The id of the uptime check
    """
    def __init__(__self__, resource_name, opts=None, content_matchers=None, display_name=None, http_check=None, monitored_resource=None, period=None, project=None, resource_group=None, selected_regions=None, tcp_check=None, timeout=None, __props__=None, __name__=None, __opts__=None):
        """
        This message configures which resources and services to monitor for availability.


        To get more information about UptimeCheckConfig, see:

        * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.uptimeCheckConfigs)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/monitoring/uptime-checks/)

        ## Example Usage - Uptime Check Config Http


        ```python
        import pulumi
        import pulumi_gcp as gcp

        http = gcp.monitoring.UptimeCheckConfig("http",
            content_matchers=[{
                "content": "example",
            }],
            display_name="http-uptime-check",
            http_check={
                "path": "/some-path",
                "port": "8010",
            },
            monitored_resource={
                "labels": {
                    "host": "192.168.1.1",
                    "project_id": "my-project-name",
                },
                "type": "uptime_url",
            },
            timeout="60s")
        ```
        ## Example Usage - Uptime Check Config Https


        ```python
        import pulumi
        import pulumi_gcp as gcp

        https = gcp.monitoring.UptimeCheckConfig("https",
            content_matchers=[{
                "content": "example",
            }],
            display_name="https-uptime-check",
            http_check={
                "path": "/some-path",
                "port": "443",
                "useSsl": True,
                "validateSsl": True,
            },
            monitored_resource={
                "labels": {
                    "host": "192.168.1.1",
                    "project_id": "my-project-name",
                },
                "type": "uptime_url",
            },
            timeout="60s")
        ```
        ## Example Usage - Uptime Check Tcp


        ```python
        import pulumi
        import pulumi_gcp as gcp

        check = gcp.monitoring.Group("check",
            display_name="uptime-check-group",
            filter="resource.metadata.name=has_substring(\"foo\")")
        tcp_group = gcp.monitoring.UptimeCheckConfig("tcpGroup",
            display_name="tcp-uptime-check",
            timeout="60s",
            tcp_check={
                "port": 888,
            },
            resource_group={
                "resourceType": "INSTANCE",
                "groupId": check.name,
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] content_matchers: The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.  Structure is documented below.
        :param pulumi.Input[str] display_name: A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
        :param pulumi.Input[dict] http_check: Contains information needed to make an HTTP or HTTPS check.  Structure is documented below.
        :param pulumi.Input[dict] monitored_resource: The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptime_url  gce_instance  gae_app  aws_ec2_instance  aws_elb_load_balancer  Structure is documented below.
        :param pulumi.Input[str] period: How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] resource_group: The group resource associated with the configuration.  Structure is documented below.
        :param pulumi.Input[list] selected_regions: The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
        :param pulumi.Input[dict] tcp_check: Contains information needed to make a TCP check.  Structure is documented below.
        :param pulumi.Input[str] timeout: The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration

        The **content_matchers** object supports the following:

          * `content` (`pulumi.Input[str]`) - String or regex content to match (max 1024 bytes)

        The **http_check** object supports the following:

          * `authInfo` (`pulumi.Input[dict]`) - The authentication information. Optional when creating an HTTP check; defaults to empty.  Structure is documented below.
            * `password` (`pulumi.Input[str]`) - The password to authenticate.
            * `username` (`pulumi.Input[str]`) - The username to authenticate.

          * `headers` (`pulumi.Input[dict]`) - The list of headers to send as part of the uptime check request. If two headers have the same key and different values, they should be entered as a single header, with the value being a comma-separated list of all the desired values as described at https://www.w3.org/Protocols/rfc2616/rfc2616.txt (page 31). Entering two separate headers with the same key in a Create call will cause the first to be overwritten by the second. The maximum number of headers allowed is 100.
          * `maskHeaders` (`pulumi.Input[bool]`) - Boolean specifying whether to encrypt the header information. Encryption should be specified for any headers related to authentication that you do not wish to be seen when retrieving the configuration. The server will be responsible for encrypting the headers. On Get/List calls, if mask_headers is set to True then the headers will be obscured with ******.
          * `path` (`pulumi.Input[str]`) - The path to the page to run the check against. Will be combined with the host (specified within the MonitoredResource) and port to construct the full URL. Optional (defaults to "/").
          * `port` (`pulumi.Input[float]`) - The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) to construct the full URL.
          * `useSsl` (`pulumi.Input[bool]`) - If true, use HTTPS instead of HTTP to run the check.
          * `validateSsl` (`pulumi.Input[bool]`) - Boolean specifying whether to include SSL certificate validation as a part of the Uptime check. Only applies to checks where monitoredResource is set to uptime_url. If useSsl is false, setting validateSsl to true has no effect.

        The **monitored_resource** object supports the following:

          * `labels` (`pulumi.Input[dict]`) - Values for all of the labels listed in the associated monitored resource descriptor. For example, Compute Engine VM instances use the labels "project_id", "instance_id", and "zone".
          * `type` (`pulumi.Input[str]`) - The monitored resource type. This field must match the type field of a MonitoredResourceDescriptor (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.monitoredResourceDescriptors#MonitoredResourceDescriptor) object. For example, the type of a Compute Engine VM instance is gce_instance. For a list of types, see Monitoring resource types (https://cloud.google.com/monitoring/api/resources) and Logging resource types (https://cloud.google.com/logging/docs/api/v2/resource-list).

        The **resource_group** object supports the following:

          * `groupId` (`pulumi.Input[str]`) - The group of resources being monitored. Should be the `name` of a group
          * `resourceType` (`pulumi.Input[str]`) - The resource type of the group members.

        The **tcp_check** object supports the following:

          * `port` (`pulumi.Input[float]`) - The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) to construct the full URL.
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
        :param pulumi.Input[list] content_matchers: The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.  Structure is documented below.
        :param pulumi.Input[str] display_name: A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
        :param pulumi.Input[dict] http_check: Contains information needed to make an HTTP or HTTPS check.  Structure is documented below.
        :param pulumi.Input[dict] monitored_resource: The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptime_url  gce_instance  gae_app  aws_ec2_instance  aws_elb_load_balancer  Structure is documented below.
        :param pulumi.Input[str] name: A unique resource name for this UptimeCheckConfig. The format is
               projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].
        :param pulumi.Input[str] period: How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] resource_group: The group resource associated with the configuration.  Structure is documented below.
        :param pulumi.Input[list] selected_regions: The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
        :param pulumi.Input[dict] tcp_check: Contains information needed to make a TCP check.  Structure is documented below.
        :param pulumi.Input[str] timeout: The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
        :param pulumi.Input[str] uptime_check_id: The id of the uptime check

        The **content_matchers** object supports the following:

          * `content` (`pulumi.Input[str]`) - String or regex content to match (max 1024 bytes)

        The **http_check** object supports the following:

          * `authInfo` (`pulumi.Input[dict]`) - The authentication information. Optional when creating an HTTP check; defaults to empty.  Structure is documented below.
            * `password` (`pulumi.Input[str]`) - The password to authenticate.
            * `username` (`pulumi.Input[str]`) - The username to authenticate.

          * `headers` (`pulumi.Input[dict]`) - The list of headers to send as part of the uptime check request. If two headers have the same key and different values, they should be entered as a single header, with the value being a comma-separated list of all the desired values as described at https://www.w3.org/Protocols/rfc2616/rfc2616.txt (page 31). Entering two separate headers with the same key in a Create call will cause the first to be overwritten by the second. The maximum number of headers allowed is 100.
          * `maskHeaders` (`pulumi.Input[bool]`) - Boolean specifying whether to encrypt the header information. Encryption should be specified for any headers related to authentication that you do not wish to be seen when retrieving the configuration. The server will be responsible for encrypting the headers. On Get/List calls, if mask_headers is set to True then the headers will be obscured with ******.
          * `path` (`pulumi.Input[str]`) - The path to the page to run the check against. Will be combined with the host (specified within the MonitoredResource) and port to construct the full URL. Optional (defaults to "/").
          * `port` (`pulumi.Input[float]`) - The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) to construct the full URL.
          * `useSsl` (`pulumi.Input[bool]`) - If true, use HTTPS instead of HTTP to run the check.
          * `validateSsl` (`pulumi.Input[bool]`) - Boolean specifying whether to include SSL certificate validation as a part of the Uptime check. Only applies to checks where monitoredResource is set to uptime_url. If useSsl is false, setting validateSsl to true has no effect.

        The **monitored_resource** object supports the following:

          * `labels` (`pulumi.Input[dict]`) - Values for all of the labels listed in the associated monitored resource descriptor. For example, Compute Engine VM instances use the labels "project_id", "instance_id", and "zone".
          * `type` (`pulumi.Input[str]`) - The monitored resource type. This field must match the type field of a MonitoredResourceDescriptor (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.monitoredResourceDescriptors#MonitoredResourceDescriptor) object. For example, the type of a Compute Engine VM instance is gce_instance. For a list of types, see Monitoring resource types (https://cloud.google.com/monitoring/api/resources) and Logging resource types (https://cloud.google.com/logging/docs/api/v2/resource-list).

        The **resource_group** object supports the following:

          * `groupId` (`pulumi.Input[str]`) - The group of resources being monitored. Should be the `name` of a group
          * `resourceType` (`pulumi.Input[str]`) - The resource type of the group members.

        The **tcp_check** object supports the following:

          * `port` (`pulumi.Input[float]`) - The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) to construct the full URL.
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

