# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['TriggerArgs', 'Trigger']

@pulumi.input_type
class TriggerArgs:
    def __init__(__self__, *,
                 destination: pulumi.Input['TriggerDestinationArgs'],
                 location: pulumi.Input[str],
                 matching_criterias: pulumi.Input[Sequence[pulumi.Input['TriggerMatchingCriteriaArgs']]],
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_account: Optional[pulumi.Input[str]] = None,
                 transports: Optional[pulumi.Input[Sequence[pulumi.Input['TriggerTransportArgs']]]] = None):
        """
        The set of arguments for constructing a Trigger resource.
        :param pulumi.Input['TriggerDestinationArgs'] destination: Required. Destination specifies where the events should be sent to.
        :param pulumi.Input[str] location: The location for the resource
        :param pulumi.Input[Sequence[pulumi.Input['TriggerMatchingCriteriaArgs']]] matching_criterias: Required. The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. User labels attached to the triggers that can be used to group resources.
        :param pulumi.Input[str] name: Required. The resource name of the trigger. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
        :param pulumi.Input[str] project: The project for the resource
        :param pulumi.Input[str] service_account: Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa\\\_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
        :param pulumi.Input[Sequence[pulumi.Input['TriggerTransportArgs']]] transports: Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
        """
        pulumi.set(__self__, "destination", destination)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "matching_criterias", matching_criterias)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if service_account is not None:
            pulumi.set(__self__, "service_account", service_account)
        if transports is not None:
            pulumi.set(__self__, "transports", transports)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Input['TriggerDestinationArgs']:
        """
        Required. Destination specifies where the events should be sent to.
        """
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: pulumi.Input['TriggerDestinationArgs']):
        pulumi.set(self, "destination", value)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        The location for the resource
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="matchingCriterias")
    def matching_criterias(self) -> pulumi.Input[Sequence[pulumi.Input['TriggerMatchingCriteriaArgs']]]:
        """
        Required. The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
        """
        return pulumi.get(self, "matching_criterias")

    @matching_criterias.setter
    def matching_criterias(self, value: pulumi.Input[Sequence[pulumi.Input['TriggerMatchingCriteriaArgs']]]):
        pulumi.set(self, "matching_criterias", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. User labels attached to the triggers that can be used to group resources.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The resource name of the trigger. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project for the resource
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa\\\_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
        """
        return pulumi.get(self, "service_account")

    @service_account.setter
    def service_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account", value)

    @property
    @pulumi.getter
    def transports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TriggerTransportArgs']]]]:
        """
        Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
        """
        return pulumi.get(self, "transports")

    @transports.setter
    def transports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TriggerTransportArgs']]]]):
        pulumi.set(self, "transports", value)


@pulumi.input_type
class _TriggerState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 destination: Optional[pulumi.Input['TriggerDestinationArgs']] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 matching_criterias: Optional[pulumi.Input[Sequence[pulumi.Input['TriggerMatchingCriteriaArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_account: Optional[pulumi.Input[str]] = None,
                 transports: Optional[pulumi.Input[Sequence[pulumi.Input['TriggerTransportArgs']]]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Trigger resources.
        :param pulumi.Input['TriggerDestinationArgs'] destination: Required. Destination specifies where the events should be sent to.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. User labels attached to the triggers that can be used to group resources.
        :param pulumi.Input[str] location: The location for the resource
        :param pulumi.Input[Sequence[pulumi.Input['TriggerMatchingCriteriaArgs']]] matching_criterias: Required. The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
        :param pulumi.Input[str] name: Required. The resource name of the trigger. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
        :param pulumi.Input[str] project: The project for the resource
        :param pulumi.Input[str] service_account: Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa\\\_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
        :param pulumi.Input[Sequence[pulumi.Input['TriggerTransportArgs']]] transports: Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if destination is not None:
            pulumi.set(__self__, "destination", destination)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if matching_criterias is not None:
            pulumi.set(__self__, "matching_criterias", matching_criterias)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if service_account is not None:
            pulumi.set(__self__, "service_account", service_account)
        if transports is not None:
            pulumi.set(__self__, "transports", transports)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def destination(self) -> Optional[pulumi.Input['TriggerDestinationArgs']]:
        """
        Required. Destination specifies where the events should be sent to.
        """
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: Optional[pulumi.Input['TriggerDestinationArgs']]):
        pulumi.set(self, "destination", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. User labels attached to the triggers that can be used to group resources.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location for the resource
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="matchingCriterias")
    def matching_criterias(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TriggerMatchingCriteriaArgs']]]]:
        """
        Required. The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
        """
        return pulumi.get(self, "matching_criterias")

    @matching_criterias.setter
    def matching_criterias(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TriggerMatchingCriteriaArgs']]]]):
        pulumi.set(self, "matching_criterias", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The resource name of the trigger. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project for the resource
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa\\\_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
        """
        return pulumi.get(self, "service_account")

    @service_account.setter
    def service_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account", value)

    @property
    @pulumi.getter
    def transports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TriggerTransportArgs']]]]:
        """
        Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
        """
        return pulumi.get(self, "transports")

    @transports.setter
    def transports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TriggerTransportArgs']]]]):
        pulumi.set(self, "transports", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class Trigger(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination: Optional[pulumi.Input[pulumi.InputType['TriggerDestinationArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 matching_criterias: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TriggerMatchingCriteriaArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_account: Optional[pulumi.Input[str]] = None,
                 transports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TriggerTransportArgs']]]]] = None,
                 __props__=None):
        """
        An event trigger sends messages to the event receiver service deployed on Cloud Run.

        * [API documentation](https://cloud.google.com/eventarc/docs/reference/rest/v1/projects.locations.triggers)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.cloudrun.Service("default",
            location="us-central1",
            metadata=gcp.cloudrun.ServiceMetadataArgs(
                namespace="my-project",
            ),
            template=gcp.cloudrun.ServiceTemplateArgs(
                spec=gcp.cloudrun.ServiceTemplateSpecArgs(
                    containers=[gcp.cloudrun.ServiceTemplateSpecContainerArgs(
                        image="gcr.io/cloudrun/hello",
                        args=["arrgs"],
                    )],
                    container_concurrency=50,
                ),
            ),
            traffics=[gcp.cloudrun.ServiceTrafficArgs(
                percent=100,
                latest_revision=True,
            )])
        trigger = gcp.eventarc.Trigger("trigger",
            location="us-central1",
            matching_criterias=[gcp.eventarc.TriggerMatchingCriteriaArgs(
                attribute="type",
                value="google.cloud.pubsub.topic.v1.messagePublished",
            )],
            destination=gcp.eventarc.TriggerDestinationArgs(
                cloud_run_service=gcp.eventarc.TriggerDestinationCloudRunServiceArgs(
                    service=default.name,
                    region="us-central1",
                ),
            ))
        ```

        ## Import

        Trigger can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:eventarc/trigger:Trigger default projects/{{project}}/locations/{{location}}/triggers/{{name}}
        ```

        ```sh
         $ pulumi import gcp:eventarc/trigger:Trigger default {{project}}/{{location}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:eventarc/trigger:Trigger default {{location}}/{{name}}
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['TriggerDestinationArgs']] destination: Required. Destination specifies where the events should be sent to.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. User labels attached to the triggers that can be used to group resources.
        :param pulumi.Input[str] location: The location for the resource
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TriggerMatchingCriteriaArgs']]]] matching_criterias: Required. The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
        :param pulumi.Input[str] name: Required. The resource name of the trigger. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
        :param pulumi.Input[str] project: The project for the resource
        :param pulumi.Input[str] service_account: Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa\\\_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TriggerTransportArgs']]]] transports: Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: TriggerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An event trigger sends messages to the event receiver service deployed on Cloud Run.

        * [API documentation](https://cloud.google.com/eventarc/docs/reference/rest/v1/projects.locations.triggers)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.cloudrun.Service("default",
            location="us-central1",
            metadata=gcp.cloudrun.ServiceMetadataArgs(
                namespace="my-project",
            ),
            template=gcp.cloudrun.ServiceTemplateArgs(
                spec=gcp.cloudrun.ServiceTemplateSpecArgs(
                    containers=[gcp.cloudrun.ServiceTemplateSpecContainerArgs(
                        image="gcr.io/cloudrun/hello",
                        args=["arrgs"],
                    )],
                    container_concurrency=50,
                ),
            ),
            traffics=[gcp.cloudrun.ServiceTrafficArgs(
                percent=100,
                latest_revision=True,
            )])
        trigger = gcp.eventarc.Trigger("trigger",
            location="us-central1",
            matching_criterias=[gcp.eventarc.TriggerMatchingCriteriaArgs(
                attribute="type",
                value="google.cloud.pubsub.topic.v1.messagePublished",
            )],
            destination=gcp.eventarc.TriggerDestinationArgs(
                cloud_run_service=gcp.eventarc.TriggerDestinationCloudRunServiceArgs(
                    service=default.name,
                    region="us-central1",
                ),
            ))
        ```

        ## Import

        Trigger can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:eventarc/trigger:Trigger default projects/{{project}}/locations/{{location}}/triggers/{{name}}
        ```

        ```sh
         $ pulumi import gcp:eventarc/trigger:Trigger default {{project}}/{{location}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:eventarc/trigger:Trigger default {{location}}/{{name}}
        ```

        :param str resource_name_: The name of the resource.
        :param TriggerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TriggerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination: Optional[pulumi.Input[pulumi.InputType['TriggerDestinationArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 matching_criterias: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TriggerMatchingCriteriaArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_account: Optional[pulumi.Input[str]] = None,
                 transports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TriggerTransportArgs']]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TriggerArgs.__new__(TriggerArgs)

            if destination is None and not opts.urn:
                raise TypeError("Missing required property 'destination'")
            __props__.__dict__["destination"] = destination
            __props__.__dict__["labels"] = labels
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            if matching_criterias is None and not opts.urn:
                raise TypeError("Missing required property 'matching_criterias'")
            __props__.__dict__["matching_criterias"] = matching_criterias
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["service_account"] = service_account
            __props__.__dict__["transports"] = transports
            __props__.__dict__["create_time"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["update_time"] = None
        super(Trigger, __self__).__init__(
            'gcp:eventarc/trigger:Trigger',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            destination: Optional[pulumi.Input[pulumi.InputType['TriggerDestinationArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            matching_criterias: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TriggerMatchingCriteriaArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            service_account: Optional[pulumi.Input[str]] = None,
            transports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TriggerTransportArgs']]]]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'Trigger':
        """
        Get an existing Trigger resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['TriggerDestinationArgs']] destination: Required. Destination specifies where the events should be sent to.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. User labels attached to the triggers that can be used to group resources.
        :param pulumi.Input[str] location: The location for the resource
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TriggerMatchingCriteriaArgs']]]] matching_criterias: Required. The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
        :param pulumi.Input[str] name: Required. The resource name of the trigger. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
        :param pulumi.Input[str] project: The project for the resource
        :param pulumi.Input[str] service_account: Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa\\\_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TriggerTransportArgs']]]] transports: Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TriggerState.__new__(_TriggerState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["destination"] = destination
        __props__.__dict__["etag"] = etag
        __props__.__dict__["labels"] = labels
        __props__.__dict__["location"] = location
        __props__.__dict__["matching_criterias"] = matching_criterias
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["service_account"] = service_account
        __props__.__dict__["transports"] = transports
        __props__.__dict__["update_time"] = update_time
        return Trigger(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Output['outputs.TriggerDestination']:
        """
        Required. Destination specifies where the events should be sent to.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Optional. User labels attached to the triggers that can be used to group resources.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location for the resource
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="matchingCriterias")
    def matching_criterias(self) -> pulumi.Output[Sequence['outputs.TriggerMatchingCriteria']]:
        """
        Required. The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
        """
        return pulumi.get(self, "matching_criterias")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Required. The resource name of the trigger. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project for the resource
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> pulumi.Output[Optional[str]]:
        """
        Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa\\\_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
        """
        return pulumi.get(self, "service_account")

    @property
    @pulumi.getter
    def transports(self) -> pulumi.Output[Sequence['outputs.TriggerTransport']]:
        """
        Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
        """
        return pulumi.get(self, "transports")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "update_time")

