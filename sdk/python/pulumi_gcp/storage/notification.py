# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Notification(pulumi.CustomResource):
    """
    Creates a new notification configuration on a specified bucket, establishing a flow of event notifications from GCS to a Cloud Pub/Sub topic.
     For more information see 
    [the official documentation](https://cloud.google.com/storage/docs/pubsub-notifications) 
    and 
    [API](https://cloud.google.com/storage/docs/json_api/v1/notifications).
    
    In order to enable notifications, a special Google Cloud Storage service account unique to the project
    must have the IAM permission "projects.topics.publish" for a Cloud Pub/Sub topic in the project. To get the service
    account's email address, use the `google_storage_project_service_account` datasource's `email_address` value, and see below
    for an example of enabling notifications by granting the correct IAM permission. See
    [the notifications documentation](https://cloud.google.com/storage/docs/gsutil/commands/notification) for more details.
    """
    def __init__(__self__, __name__, __opts__=None, bucket=None, custom_attributes=None, event_types=None, object_name_prefix=None, payload_format=None, topic=None):
        """Create a Notification resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not bucket:
            raise TypeError('Missing required property bucket')
        __props__['bucket'] = bucket

        __props__['custom_attributes'] = custom_attributes

        __props__['event_types'] = event_types

        __props__['object_name_prefix'] = object_name_prefix

        if not payload_format:
            raise TypeError('Missing required property payload_format')
        __props__['payload_format'] = payload_format

        if not topic:
            raise TypeError('Missing required property topic')
        __props__['topic'] = topic

        __props__['self_link'] = None

        super(Notification, __self__).__init__(
            'gcp:storage/notification:Notification',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

