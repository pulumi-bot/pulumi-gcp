# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['AlertPolicy']


class AlertPolicy(pulumi.CustomResource):
    combiner: pulumi.Output[str] = pulumi.output_property("combiner")
    """
    How to combine the results of multiple conditions to
    determine if an incident should be opened.
    """
    conditions: pulumi.Output[List['outputs.AlertPolicyCondition']] = pulumi.output_property("conditions")
    """
    A list of conditions for the policy. The conditions are combined by
    AND or OR according to the combiner field. If the combined conditions
    evaluate to true, then an incident is created. A policy can have from
    one to six conditions.  Structure is documented below.
    """
    creation_record: pulumi.Output['outputs.AlertPolicyCreationRecord'] = pulumi.output_property("creationRecord")
    """
    A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be
    ignored.
    """
    display_name: pulumi.Output[str] = pulumi.output_property("displayName")
    """
    A short name or phrase used to identify the
    condition in dashboards, notifications, and
    incidents. To avoid confusion, don't use the same
    display name for multiple conditions in the same
    policy.
    """
    documentation: pulumi.Output[Optional['outputs.AlertPolicyDocumentation']] = pulumi.output_property("documentation")
    """
    A short name or phrase used to identify the policy in dashboards,
    notifications, and incidents. To avoid confusion, don't use the same
    display name for multiple policies in the same project. The name is
    limited to 512 Unicode characters.  Structure is documented below.
    """
    enabled: pulumi.Output[Optional[bool]] = pulumi.output_property("enabled")
    """
    Whether or not the policy is enabled. The default is true.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    -
    The unique resource name for this condition.
    Its syntax is:
    projects/[PROJECT_ID]/alertPolicies/[POLICY_ID]/conditions/[CONDITION_ID]
    [CONDITION_ID] is assigned by Stackdriver Monitoring when
    the condition is created as part of a new or updated alerting
    policy.
    """
    notification_channels: pulumi.Output[Optional[List[str]]] = pulumi.output_property("notificationChannels")
    """
    Identifies the notification channels to which notifications should be
    sent when incidents are opened or closed or when new violations occur
    on an already opened incident. Each element of this array corresponds
    to the name field in each of the NotificationChannel objects that are
    returned from the notificationChannels.list method. The syntax of the
    entries in this field is
    `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]`
    """
    project: pulumi.Output[str] = pulumi.output_property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    user_labels: pulumi.Output[Optional[Dict[str, str]]] = pulumi.output_property("userLabels")
    """
    This field is intended to be used for organizing and identifying the AlertPolicy
    objects.The field can contain up to 64 entries. Each key and value is limited
    to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
    can contain only lowercase letters, numerals, underscores, and dashes. Keys
    must begin with a letter.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, combiner: Optional[pulumi.Input[str]] = None, conditions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['AlertPolicyConditionArgs']]]]] = None, display_name: Optional[pulumi.Input[str]] = None, documentation: Optional[pulumi.Input[pulumi.InputType['AlertPolicyDocumentationArgs']]] = None, enabled: Optional[pulumi.Input[bool]] = None, notification_channels: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, project: Optional[pulumi.Input[str]] = None, user_labels: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        A description of the conditions under which some aspect of your system is
        considered to be "unhealthy" and the ways to notify people or services
        about this state.

        To get more information about AlertPolicy, see:

        * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.alertPolicies)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/monitoring/alerts/)

        ## Example Usage
        ### Monitoring Alert Policy Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        alert_policy = gcp.monitoring.AlertPolicy("alertPolicy",
            combiner="OR",
            conditions=[{
                "conditionThreshold": {
                    "aggregations": [{
                        "alignmentPeriod": "60s",
                        "perSeriesAligner": "ALIGN_RATE",
                    }],
                    "comparison": "COMPARISON_GT",
                    "duration": "60s",
                    "filter": "metric.type=\"compute.googleapis.com/instance/disk/write_bytes_count\" AND resource.type=\"gce_instance\"",
                },
                "display_name": "test condition",
            }],
            display_name="My Alert Policy",
            user_labels={
                "foo": "bar",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] combiner: How to combine the results of multiple conditions to
               determine if an incident should be opened.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['AlertPolicyConditionArgs']]]] conditions: A list of conditions for the policy. The conditions are combined by
               AND or OR according to the combiner field. If the combined conditions
               evaluate to true, then an incident is created. A policy can have from
               one to six conditions.  Structure is documented below.
        :param pulumi.Input[str] display_name: A short name or phrase used to identify the
               condition in dashboards, notifications, and
               incidents. To avoid confusion, don't use the same
               display name for multiple conditions in the same
               policy.
        :param pulumi.Input[pulumi.InputType['AlertPolicyDocumentationArgs']] documentation: A short name or phrase used to identify the policy in dashboards,
               notifications, and incidents. To avoid confusion, don't use the same
               display name for multiple policies in the same project. The name is
               limited to 512 Unicode characters.  Structure is documented below.
        :param pulumi.Input[bool] enabled: Whether or not the policy is enabled. The default is true.
        :param pulumi.Input[List[pulumi.Input[str]]] notification_channels: Identifies the notification channels to which notifications should be
               sent when incidents are opened or closed or when new violations occur
               on an already opened incident. Each element of this array corresponds
               to the name field in each of the NotificationChannel objects that are
               returned from the notificationChannels.list method. The syntax of the
               entries in this field is
               `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]`
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] user_labels: This field is intended to be used for organizing and identifying the AlertPolicy
               objects.The field can contain up to 64 entries. Each key and value is limited
               to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
               can contain only lowercase letters, numerals, underscores, and dashes. Keys
               must begin with a letter.
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
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, combiner: Optional[pulumi.Input[str]] = None, conditions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['AlertPolicyConditionArgs']]]]] = None, creation_record: Optional[pulumi.Input[pulumi.InputType['AlertPolicyCreationRecordArgs']]] = None, display_name: Optional[pulumi.Input[str]] = None, documentation: Optional[pulumi.Input[pulumi.InputType['AlertPolicyDocumentationArgs']]] = None, enabled: Optional[pulumi.Input[bool]] = None, name: Optional[pulumi.Input[str]] = None, notification_channels: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, project: Optional[pulumi.Input[str]] = None, user_labels: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None) -> 'AlertPolicy':
        """
        Get an existing AlertPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] combiner: How to combine the results of multiple conditions to
               determine if an incident should be opened.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['AlertPolicyConditionArgs']]]] conditions: A list of conditions for the policy. The conditions are combined by
               AND or OR according to the combiner field. If the combined conditions
               evaluate to true, then an incident is created. A policy can have from
               one to six conditions.  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['AlertPolicyCreationRecordArgs']] creation_record: A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be
               ignored.
        :param pulumi.Input[str] display_name: A short name or phrase used to identify the
               condition in dashboards, notifications, and
               incidents. To avoid confusion, don't use the same
               display name for multiple conditions in the same
               policy.
        :param pulumi.Input[pulumi.InputType['AlertPolicyDocumentationArgs']] documentation: A short name or phrase used to identify the policy in dashboards,
               notifications, and incidents. To avoid confusion, don't use the same
               display name for multiple policies in the same project. The name is
               limited to 512 Unicode characters.  Structure is documented below.
        :param pulumi.Input[bool] enabled: Whether or not the policy is enabled. The default is true.
        :param pulumi.Input[str] name: -
               The unique resource name for this condition.
               Its syntax is:
               projects/[PROJECT_ID]/alertPolicies/[POLICY_ID]/conditions/[CONDITION_ID]
               [CONDITION_ID] is assigned by Stackdriver Monitoring when
               the condition is created as part of a new or updated alerting
               policy.
        :param pulumi.Input[List[pulumi.Input[str]]] notification_channels: Identifies the notification channels to which notifications should be
               sent when incidents are opened or closed or when new violations occur
               on an already opened incident. Each element of this array corresponds
               to the name field in each of the NotificationChannel objects that are
               returned from the notificationChannels.list method. The syntax of the
               entries in this field is
               `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]`
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] user_labels: This field is intended to be used for organizing and identifying the AlertPolicy
               objects.The field can contain up to 64 entries. Each key and value is limited
               to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
               can contain only lowercase letters, numerals, underscores, and dashes. Keys
               must begin with a letter.
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
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

