# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetNotificationChannelResult:
    """
    A collection of values returned by getNotificationChannel.
    """
    def __init__(__self__, description=None, display_name=None, enabled=None, id=None, labels=None, name=None, project=None, sensitive_labels=None, type=None, user_labels=None, verification_status=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        __self__.display_name = display_name
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        __self__.enabled = enabled
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        __self__.labels = labels
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if sensitive_labels and not isinstance(sensitive_labels, list):
            raise TypeError("Expected argument 'sensitive_labels' to be a list")
        __self__.sensitive_labels = sensitive_labels
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        if user_labels and not isinstance(user_labels, dict):
            raise TypeError("Expected argument 'user_labels' to be a dict")
        __self__.user_labels = user_labels
        if verification_status and not isinstance(verification_status, str):
            raise TypeError("Expected argument 'verification_status' to be a str")
        __self__.verification_status = verification_status
class AwaitableGetNotificationChannelResult(GetNotificationChannelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNotificationChannelResult(
            description=self.description,
            display_name=self.display_name,
            enabled=self.enabled,
            id=self.id,
            labels=self.labels,
            name=self.name,
            project=self.project,
            sensitive_labels=self.sensitive_labels,
            type=self.type,
            user_labels=self.user_labels,
            verification_status=self.verification_status)

def get_notification_channel(display_name=None,labels=None,project=None,type=None,user_labels=None,opts=None):
    """
    A NotificationChannel is a medium through which an alert is delivered
    when a policy violation is detected. Examples of channels include email, SMS,
    and third-party messaging applications. Fields containing sensitive information
    like authentication tokens or contact info are only partially populated on retrieval.

    To get more information about NotificationChannel, see:

    * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannels)
    * How-to Guides
        * [Notification Options](https://cloud.google.com/monitoring/support/notification-options)
        * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)

    ## Example Usage
    ### Notification Channel Basic

    ```python
    import pulumi
    import pulumi_gcp as gcp

    basic = gcp.monitoring.get_notification_channel(display_name="Test Notification Channel")
    alert_policy = gcp.monitoring.AlertPolicy("alertPolicy",
        display_name="My Alert Policy",
        notification_channels=[basic.name],
        combiner="OR",
        conditions=[{
            "display_name": "test condition",
            "condition_threshold": {
                "filter": "metric.type=\"compute.googleapis.com/instance/disk/write_bytes_count\" AND resource.type=\"gce_instance\"",
                "duration": "60s",
                "comparison": "COMPARISON_GT",
                "aggregations": [{
                    "alignmentPeriod": "60s",
                    "perSeriesAligner": "ALIGN_RATE",
                }],
            },
        }])
    ```


    :param str display_name: The display name for this notification channel.
    :param dict labels: Labels (corresponding to the
           NotificationChannelDescriptor schema) to filter the notification channels by.
    :param str project: The ID of the project in which the resource belongs.
           If it is not provided, the provider project is used.
    :param str type: The type of the notification channel.
    :param dict user_labels: User-provided key-value labels to filter by.
    """
    __args__ = dict()


    __args__['displayName'] = display_name
    __args__['labels'] = labels
    __args__['project'] = project
    __args__['type'] = type
    __args__['userLabels'] = user_labels
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:monitoring/getNotificationChannel:getNotificationChannel', __args__, opts=opts).value

    return AwaitableGetNotificationChannelResult(
        description=__ret__.get('description'),
        display_name=__ret__.get('displayName'),
        enabled=__ret__.get('enabled'),
        id=__ret__.get('id'),
        labels=__ret__.get('labels'),
        name=__ret__.get('name'),
        project=__ret__.get('project'),
        sensitive_labels=__ret__.get('sensitiveLabels'),
        type=__ret__.get('type'),
        user_labels=__ret__.get('userLabels'),
        verification_status=__ret__.get('verificationStatus'))
