# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AlertPolicy(pulumi.CustomResource):
    combiner: pulumi.Output[str]
    """
    How to combine the results of multiple conditions to determine if an incident should be opened.
    """
    conditions: pulumi.Output[list]
    """
    A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner field. If the
    combined conditions evaluate to true, then an incident is created. A policy can have from one to six conditions.

      * `conditionAbsent` (`dict`)
        * `aggregations` (`list`)
          * `alignmentPeriod` (`str`)
          * `crossSeriesReducer` (`str`)
          * `groupByFields` (`list`)
          * `perSeriesAligner` (`str`)

        * `duration` (`str`)
        * `filter` (`str`)
        * `trigger` (`dict`)
          * `count` (`float`)
          * `percent` (`float`)

      * `conditionThreshold` (`dict`)
        * `aggregations` (`list`)
          * `alignmentPeriod` (`str`)
          * `crossSeriesReducer` (`str`)
          * `groupByFields` (`list`)
          * `perSeriesAligner` (`str`)

        * `comparison` (`str`)
        * `denominatorAggregations` (`list`)
          * `alignmentPeriod` (`str`)
          * `crossSeriesReducer` (`str`)
          * `groupByFields` (`list`)
          * `perSeriesAligner` (`str`)

        * `denominatorFilter` (`str`)
        * `duration` (`str`)
        * `filter` (`str`)
        * `thresholdValue` (`float`)
        * `trigger` (`dict`)
          * `count` (`float`)
          * `percent` (`float`)

      * `display_name` (`str`)
      * `name` (`str`)
    """
    creation_record: pulumi.Output[dict]
    """
    A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be
    ignored.

      * `mutateTime` (`str`)
      * `mutatedBy` (`str`)
    """
    display_name: pulumi.Output[str]
    """
    A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion,
    don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode
    characters.
    """
    documentation: pulumi.Output[dict]
    """
    A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion,
    don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode
    characters.

      * `content` (`str`)
      * `mimeType` (`str`)
    """
    enabled: pulumi.Output[bool]
    """
    Whether or not the policy is enabled. The default is true.
    """
    name: pulumi.Output[str]
    """
    The unique resource name for this policy. Its syntax is: projects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]
    """
    notification_channels: pulumi.Output[list]
    """
    Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when
    new violations occur on an already opened incident. Each element of this array corresponds to the name field in each of
    the NotificationChannel objects that are returned from the notificationChannels.list method. The syntax of the entries
    in this field is 'projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]'
    """
    project: pulumi.Output[str]
    user_labels: pulumi.Output[dict]
    """
    This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64
    entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
    can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
    """
    def __init__(__self__, resource_name, opts=None, combiner=None, conditions=None, display_name=None, documentation=None, enabled=None, notification_channels=None, project=None, user_labels=None, __props__=None, __name__=None, __opts__=None):
        """
        A description of the conditions under which some aspect of your system is
        considered to be "unhealthy" and the ways to notify people or services
        about this state.


        To get more information about AlertPolicy, see:

        * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.alertPolicies)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/monitoring/alerts/)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/monitoring_alert_policy.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] combiner: How to combine the results of multiple conditions to determine if an incident should be opened.
        :param pulumi.Input[list] conditions: A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner field. If the
               combined conditions evaluate to true, then an incident is created. A policy can have from one to six conditions.
        :param pulumi.Input[str] display_name: A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion,
               don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode
               characters.
        :param pulumi.Input[dict] documentation: A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion,
               don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode
               characters.
        :param pulumi.Input[bool] enabled: Whether or not the policy is enabled. The default is true.
        :param pulumi.Input[list] notification_channels: Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when
               new violations occur on an already opened incident. Each element of this array corresponds to the name field in each of
               the NotificationChannel objects that are returned from the notificationChannels.list method. The syntax of the entries
               in this field is 'projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]'
        :param pulumi.Input[dict] user_labels: This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64
               entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
               can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.

        The **conditions** object supports the following:

          * `conditionAbsent` (`pulumi.Input[dict]`)
            * `aggregations` (`pulumi.Input[list]`)
              * `alignmentPeriod` (`pulumi.Input[str]`)
              * `crossSeriesReducer` (`pulumi.Input[str]`)
              * `groupByFields` (`pulumi.Input[list]`)
              * `perSeriesAligner` (`pulumi.Input[str]`)

            * `duration` (`pulumi.Input[str]`)
            * `filter` (`pulumi.Input[str]`)
            * `trigger` (`pulumi.Input[dict]`)
              * `count` (`pulumi.Input[float]`)
              * `percent` (`pulumi.Input[float]`)

          * `conditionThreshold` (`pulumi.Input[dict]`)
            * `aggregations` (`pulumi.Input[list]`)
              * `alignmentPeriod` (`pulumi.Input[str]`)
              * `crossSeriesReducer` (`pulumi.Input[str]`)
              * `groupByFields` (`pulumi.Input[list]`)
              * `perSeriesAligner` (`pulumi.Input[str]`)

            * `comparison` (`pulumi.Input[str]`)
            * `denominatorAggregations` (`pulumi.Input[list]`)
              * `alignmentPeriod` (`pulumi.Input[str]`)
              * `crossSeriesReducer` (`pulumi.Input[str]`)
              * `groupByFields` (`pulumi.Input[list]`)
              * `perSeriesAligner` (`pulumi.Input[str]`)

            * `denominatorFilter` (`pulumi.Input[str]`)
            * `duration` (`pulumi.Input[str]`)
            * `filter` (`pulumi.Input[str]`)
            * `thresholdValue` (`pulumi.Input[float]`)
            * `trigger` (`pulumi.Input[dict]`)
              * `count` (`pulumi.Input[float]`)
              * `percent` (`pulumi.Input[float]`)

          * `display_name` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`)

        The **documentation** object supports the following:

          * `content` (`pulumi.Input[str]`)
          * `mimeType` (`pulumi.Input[str]`)
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

            if combiner is None:
                raise TypeError("Missing required property 'combiner'")
            __props__['combiner'] = combiner
            if conditions is None:
                raise TypeError("Missing required property 'conditions'")
            __props__['conditions'] = conditions
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['documentation'] = documentation
            __props__['enabled'] = enabled
            __props__['notification_channels'] = notification_channels
            __props__['project'] = project
            __props__['user_labels'] = user_labels
            __props__['creation_record'] = None
            __props__['name'] = None
        super(AlertPolicy, __self__).__init__(
            'gcp:monitoring/alertPolicy:AlertPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, combiner=None, conditions=None, creation_record=None, display_name=None, documentation=None, enabled=None, name=None, notification_channels=None, project=None, user_labels=None):
        """
        Get an existing AlertPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] combiner: How to combine the results of multiple conditions to determine if an incident should be opened.
        :param pulumi.Input[list] conditions: A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner field. If the
               combined conditions evaluate to true, then an incident is created. A policy can have from one to six conditions.
        :param pulumi.Input[dict] creation_record: A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be
               ignored.
        :param pulumi.Input[str] display_name: A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion,
               don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode
               characters.
        :param pulumi.Input[dict] documentation: A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion,
               don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode
               characters.
        :param pulumi.Input[bool] enabled: Whether or not the policy is enabled. The default is true.
        :param pulumi.Input[str] name: The unique resource name for this policy. Its syntax is: projects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]
        :param pulumi.Input[list] notification_channels: Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when
               new violations occur on an already opened incident. Each element of this array corresponds to the name field in each of
               the NotificationChannel objects that are returned from the notificationChannels.list method. The syntax of the entries
               in this field is 'projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]'
        :param pulumi.Input[dict] user_labels: This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64
               entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
               can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.

        The **conditions** object supports the following:

          * `conditionAbsent` (`pulumi.Input[dict]`)
            * `aggregations` (`pulumi.Input[list]`)
              * `alignmentPeriod` (`pulumi.Input[str]`)
              * `crossSeriesReducer` (`pulumi.Input[str]`)
              * `groupByFields` (`pulumi.Input[list]`)
              * `perSeriesAligner` (`pulumi.Input[str]`)

            * `duration` (`pulumi.Input[str]`)
            * `filter` (`pulumi.Input[str]`)
            * `trigger` (`pulumi.Input[dict]`)
              * `count` (`pulumi.Input[float]`)
              * `percent` (`pulumi.Input[float]`)

          * `conditionThreshold` (`pulumi.Input[dict]`)
            * `aggregations` (`pulumi.Input[list]`)
              * `alignmentPeriod` (`pulumi.Input[str]`)
              * `crossSeriesReducer` (`pulumi.Input[str]`)
              * `groupByFields` (`pulumi.Input[list]`)
              * `perSeriesAligner` (`pulumi.Input[str]`)

            * `comparison` (`pulumi.Input[str]`)
            * `denominatorAggregations` (`pulumi.Input[list]`)
              * `alignmentPeriod` (`pulumi.Input[str]`)
              * `crossSeriesReducer` (`pulumi.Input[str]`)
              * `groupByFields` (`pulumi.Input[list]`)
              * `perSeriesAligner` (`pulumi.Input[str]`)

            * `denominatorFilter` (`pulumi.Input[str]`)
            * `duration` (`pulumi.Input[str]`)
            * `filter` (`pulumi.Input[str]`)
            * `thresholdValue` (`pulumi.Input[float]`)
            * `trigger` (`pulumi.Input[dict]`)
              * `count` (`pulumi.Input[float]`)
              * `percent` (`pulumi.Input[float]`)

          * `display_name` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`)

        The **creation_record** object supports the following:

          * `mutateTime` (`pulumi.Input[str]`)
          * `mutatedBy` (`pulumi.Input[str]`)

        The **documentation** object supports the following:

          * `content` (`pulumi.Input[str]`)
          * `mimeType` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["combiner"] = combiner
        __props__["conditions"] = conditions
        __props__["creation_record"] = creation_record
        __props__["display_name"] = display_name
        __props__["documentation"] = documentation
        __props__["enabled"] = enabled
        __props__["name"] = name
        __props__["notification_channels"] = notification_channels
        __props__["project"] = project
        __props__["user_labels"] = user_labels
        return AlertPolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

