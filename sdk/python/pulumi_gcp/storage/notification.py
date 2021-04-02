# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NotificationArgs', 'Notification']

@pulumi.input_type
class NotificationArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 payload_format: pulumi.Input[str],
                 topic: pulumi.Input[str],
                 custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 object_name_prefix: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Notification resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[str] payload_format: The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
        :param pulumi.Input[str] topic: The Cloud PubSub topic to which this subscription publishes. Expects either the
               topic name, assumed to belong to the default GCP provider project, or the project-level name,
               i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`. If the project is not set in the provider,
               you will need to use the project-level name.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_attributes: A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
        :param pulumi.Input[str] object_name_prefix: Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "payload_format", payload_format)
        pulumi.set(__self__, "topic", topic)
        if custom_attributes is not None:
            pulumi.set(__self__, "custom_attributes", custom_attributes)
        if event_types is not None:
            pulumi.set(__self__, "event_types", event_types)
        if object_name_prefix is not None:
            pulumi.set(__self__, "object_name_prefix", object_name_prefix)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter(name="payloadFormat")
    def payload_format(self) -> pulumi.Input[str]:
        """
        The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
        """
        return pulumi.get(self, "payload_format")

    @payload_format.setter
    def payload_format(self, value: pulumi.Input[str]):
        pulumi.set(self, "payload_format", value)

    @property
    @pulumi.getter
    def topic(self) -> pulumi.Input[str]:
        """
        The Cloud PubSub topic to which this subscription publishes. Expects either the
        topic name, assumed to belong to the default GCP provider project, or the project-level name,
        i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`. If the project is not set in the provider,
        you will need to use the project-level name.
        """
        return pulumi.get(self, "topic")

    @topic.setter
    def topic(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic", value)

    @property
    @pulumi.getter(name="customAttributes")
    def custom_attributes(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
        """
        return pulumi.get(self, "custom_attributes")

    @custom_attributes.setter
    def custom_attributes(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "custom_attributes", value)

    @property
    @pulumi.getter(name="eventTypes")
    def event_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
        """
        return pulumi.get(self, "event_types")

    @event_types.setter
    def event_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "event_types", value)

    @property
    @pulumi.getter(name="objectNamePrefix")
    def object_name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
        """
        return pulumi.get(self, "object_name_prefix")

    @object_name_prefix.setter
    def object_name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_name_prefix", value)


@pulumi.input_type
class _NotificationState:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[str]] = None,
                 custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notification_id: Optional[pulumi.Input[str]] = None,
                 object_name_prefix: Optional[pulumi.Input[str]] = None,
                 payload_format: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 topic: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Notification resources.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_attributes: A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
        :param pulumi.Input[str] notification_id: The ID of the created notification.
        :param pulumi.Input[str] object_name_prefix: Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
        :param pulumi.Input[str] payload_format: The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] topic: The Cloud PubSub topic to which this subscription publishes. Expects either the
               topic name, assumed to belong to the default GCP provider project, or the project-level name,
               i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`. If the project is not set in the provider,
               you will need to use the project-level name.
        """
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if custom_attributes is not None:
            pulumi.set(__self__, "custom_attributes", custom_attributes)
        if event_types is not None:
            pulumi.set(__self__, "event_types", event_types)
        if notification_id is not None:
            pulumi.set(__self__, "notification_id", notification_id)
        if object_name_prefix is not None:
            pulumi.set(__self__, "object_name_prefix", object_name_prefix)
        if payload_format is not None:
            pulumi.set(__self__, "payload_format", payload_format)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)
        if topic is not None:
            pulumi.set(__self__, "topic", topic)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter(name="customAttributes")
    def custom_attributes(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
        """
        return pulumi.get(self, "custom_attributes")

    @custom_attributes.setter
    def custom_attributes(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "custom_attributes", value)

    @property
    @pulumi.getter(name="eventTypes")
    def event_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
        """
        return pulumi.get(self, "event_types")

    @event_types.setter
    def event_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "event_types", value)

    @property
    @pulumi.getter(name="notificationId")
    def notification_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the created notification.
        """
        return pulumi.get(self, "notification_id")

    @notification_id.setter
    def notification_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_id", value)

    @property
    @pulumi.getter(name="objectNamePrefix")
    def object_name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
        """
        return pulumi.get(self, "object_name_prefix")

    @object_name_prefix.setter
    def object_name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_name_prefix", value)

    @property
    @pulumi.getter(name="payloadFormat")
    def payload_format(self) -> Optional[pulumi.Input[str]]:
        """
        The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
        """
        return pulumi.get(self, "payload_format")

    @payload_format.setter
    def payload_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payload_format", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)

    @property
    @pulumi.getter
    def topic(self) -> Optional[pulumi.Input[str]]:
        """
        The Cloud PubSub topic to which this subscription publishes. Expects either the
        topic name, assumed to belong to the default GCP provider project, or the project-level name,
        i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`. If the project is not set in the provider,
        you will need to use the project-level name.
        """
        return pulumi.get(self, "topic")

    @topic.setter
    def topic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic", value)


class Notification(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 object_name_prefix: Optional[pulumi.Input[str]] = None,
                 payload_format: Optional[pulumi.Input[str]] = None,
                 topic: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a new notification configuration on a specified bucket, establishing a flow of event notifications from GCS to a Cloud Pub/Sub topic.
         For more information see
        [the official documentation](https://cloud.google.com/storage/docs/pubsub-notifications)
        and
        [API](https://cloud.google.com/storage/docs/json_api/v1/notifications).

        In order to enable notifications, a special Google Cloud Storage service account unique to the project
        must exist and have the IAM permission "projects.topics.publish" for a Cloud Pub/Sub topic in the project.
        This service account is not created automatically when a project is created.
        To ensure the service account exists and obtain its email address for use in granting the correct IAM permission, use the
        [`storage.getProjectServiceAccount`](https://www.terraform.io/docs/providers/google/d/storage_project_service_account.html)
        datasource's `email_address` value, and see below for an example of enabling notifications by granting the correct IAM permission.
        See [the notifications documentation](https://cloud.google.com/storage/docs/gsutil/commands/notification) for more details.

        > **NOTE**: This resource can affect your storage IAM policy. If you are using this in the same config as your storage IAM policy resources, consider
        making this resource dependent on those IAM resources via `depends_on`. This will safeguard against errors due to IAM race conditions.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        gcs_account = gcp.storage.get_project_service_account()
        topic = gcp.pubsub.Topic("topic")
        binding = gcp.pubsub.TopicIAMBinding("binding",
            topic=topic.id,
            role="roles/pubsub.publisher",
            members=[f"serviceAccount:{gcs_account.email_address}"])
        # End enabling notifications
        bucket = gcp.storage.Bucket("bucket")
        notification = gcp.storage.Notification("notification",
            bucket=bucket.name,
            payload_format="JSON_API_V1",
            topic=topic.id,
            event_types=[
                "OBJECT_FINALIZE",
                "OBJECT_METADATA_UPDATE",
            ],
            custom_attributes={
                "new-attribute": "new-attribute-value",
            },
            opts=pulumi.ResourceOptions(depends_on=[binding]))
        # Enable notifications by giving the correct IAM permission to the unique service account.
        ```

        ## Import

        Storage notifications can be imported using the notification `id` in the format `<bucket_name>/notificationConfigs/<id>` e.g.

        ```sh
         $ pulumi import gcp:storage/notification:Notification notification default_bucket/notificationConfigs/102
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_attributes: A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
        :param pulumi.Input[str] object_name_prefix: Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
        :param pulumi.Input[str] payload_format: The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
        :param pulumi.Input[str] topic: The Cloud PubSub topic to which this subscription publishes. Expects either the
               topic name, assumed to belong to the default GCP provider project, or the project-level name,
               i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`. If the project is not set in the provider,
               you will need to use the project-level name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NotificationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new notification configuration on a specified bucket, establishing a flow of event notifications from GCS to a Cloud Pub/Sub topic.
         For more information see
        [the official documentation](https://cloud.google.com/storage/docs/pubsub-notifications)
        and
        [API](https://cloud.google.com/storage/docs/json_api/v1/notifications).

        In order to enable notifications, a special Google Cloud Storage service account unique to the project
        must exist and have the IAM permission "projects.topics.publish" for a Cloud Pub/Sub topic in the project.
        This service account is not created automatically when a project is created.
        To ensure the service account exists and obtain its email address for use in granting the correct IAM permission, use the
        [`storage.getProjectServiceAccount`](https://www.terraform.io/docs/providers/google/d/storage_project_service_account.html)
        datasource's `email_address` value, and see below for an example of enabling notifications by granting the correct IAM permission.
        See [the notifications documentation](https://cloud.google.com/storage/docs/gsutil/commands/notification) for more details.

        > **NOTE**: This resource can affect your storage IAM policy. If you are using this in the same config as your storage IAM policy resources, consider
        making this resource dependent on those IAM resources via `depends_on`. This will safeguard against errors due to IAM race conditions.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        gcs_account = gcp.storage.get_project_service_account()
        topic = gcp.pubsub.Topic("topic")
        binding = gcp.pubsub.TopicIAMBinding("binding",
            topic=topic.id,
            role="roles/pubsub.publisher",
            members=[f"serviceAccount:{gcs_account.email_address}"])
        # End enabling notifications
        bucket = gcp.storage.Bucket("bucket")
        notification = gcp.storage.Notification("notification",
            bucket=bucket.name,
            payload_format="JSON_API_V1",
            topic=topic.id,
            event_types=[
                "OBJECT_FINALIZE",
                "OBJECT_METADATA_UPDATE",
            ],
            custom_attributes={
                "new-attribute": "new-attribute-value",
            },
            opts=pulumi.ResourceOptions(depends_on=[binding]))
        # Enable notifications by giving the correct IAM permission to the unique service account.
        ```

        ## Import

        Storage notifications can be imported using the notification `id` in the format `<bucket_name>/notificationConfigs/<id>` e.g.

        ```sh
         $ pulumi import gcp:storage/notification:Notification notification default_bucket/notificationConfigs/102
        ```

        :param str resource_name: The name of the resource.
        :param NotificationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NotificationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 object_name_prefix: Optional[pulumi.Input[str]] = None,
                 payload_format: Optional[pulumi.Input[str]] = None,
                 topic: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            __props__ = NotificationArgs.__new__(NotificationArgs)

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__['bucket'] = bucket
            __props__.__dict__['custom_attributes'] = custom_attributes
            __props__.__dict__['event_types'] = event_types
            __props__.__dict__['object_name_prefix'] = object_name_prefix
            if payload_format is None and not opts.urn:
                raise TypeError("Missing required property 'payload_format'")
            __props__.__dict__['payload_format'] = payload_format
            if topic is None and not opts.urn:
                raise TypeError("Missing required property 'topic'")
            __props__.__dict__['topic'] = topic
            __props__.__dict__['notification_id'] = None
            __props__.__dict__['self_link'] = None
        super(Notification, __self__).__init__(
            'gcp:storage/notification:Notification',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            notification_id: Optional[pulumi.Input[str]] = None,
            object_name_prefix: Optional[pulumi.Input[str]] = None,
            payload_format: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            topic: Optional[pulumi.Input[str]] = None) -> 'Notification':
        """
        Get an existing Notification resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_attributes: A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
        :param pulumi.Input[str] notification_id: The ID of the created notification.
        :param pulumi.Input[str] object_name_prefix: Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
        :param pulumi.Input[str] payload_format: The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] topic: The Cloud PubSub topic to which this subscription publishes. Expects either the
               topic name, assumed to belong to the default GCP provider project, or the project-level name,
               i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`. If the project is not set in the provider,
               you will need to use the project-level name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NotificationState.__new__(_NotificationState)

        __props__.__dict__['bucket'] = bucket
        __props__.__dict__['custom_attributes'] = custom_attributes
        __props__.__dict__['event_types'] = event_types
        __props__.__dict__['notification_id'] = notification_id
        __props__.__dict__['object_name_prefix'] = object_name_prefix
        __props__.__dict__['payload_format'] = payload_format
        __props__.__dict__['self_link'] = self_link
        __props__.__dict__['topic'] = topic
        return Notification(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="customAttributes")
    def custom_attributes(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
        """
        return pulumi.get(self, "custom_attributes")

    @property
    @pulumi.getter(name="eventTypes")
    def event_types(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
        """
        return pulumi.get(self, "event_types")

    @property
    @pulumi.getter(name="notificationId")
    def notification_id(self) -> pulumi.Output[str]:
        """
        The ID of the created notification.
        """
        return pulumi.get(self, "notification_id")

    @property
    @pulumi.getter(name="objectNamePrefix")
    def object_name_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
        """
        return pulumi.get(self, "object_name_prefix")

    @property
    @pulumi.getter(name="payloadFormat")
    def payload_format(self) -> pulumi.Output[str]:
        """
        The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
        """
        return pulumi.get(self, "payload_format")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def topic(self) -> pulumi.Output[str]:
        """
        The Cloud PubSub topic to which this subscription publishes. Expects either the
        topic name, assumed to belong to the default GCP provider project, or the project-level name,
        i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`. If the project is not set in the provider,
        you will need to use the project-level name.
        """
        return pulumi.get(self, "topic")

