# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Function']


class Function(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 available_memory_mb: Optional[pulumi.Input[float]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entry_point: Optional[pulumi.Input[str]] = None,
                 environment_variables: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 event_trigger: Optional[pulumi.Input[pulumi.InputType['FunctionEventTriggerArgs']]] = None,
                 https_trigger_url: Optional[pulumi.Input[str]] = None,
                 ingress_settings: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 max_instances: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 runtime: Optional[pulumi.Input[str]] = None,
                 service_account_email: Optional[pulumi.Input[str]] = None,
                 source_archive_bucket: Optional[pulumi.Input[str]] = None,
                 source_archive_object: Optional[pulumi.Input[str]] = None,
                 source_repository: Optional[pulumi.Input[pulumi.InputType['FunctionSourceRepositoryArgs']]] = None,
                 timeout: Optional[pulumi.Input[float]] = None,
                 trigger_http: Optional[pulumi.Input[bool]] = None,
                 vpc_connector: Optional[pulumi.Input[str]] = None,
                 vpc_connector_egress_settings: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a new Cloud Function. For more information see
        [the official documentation](https://cloud.google.com/functions/docs/)
        and
        [API](https://cloud.google.com/functions/docs/apis).

        > **Warning:** As of November 1, 2019, newly created Functions are
        private-by-default and will require [appropriate IAM permissions](https://cloud.google.com/functions/docs/reference/iam/roles)
        to be invoked. See below examples for how to set up the appropriate permissions,
        or view the [Cloud Functions IAM resources](https://www.terraform.io/docs/providers/google/r/cloudfunctions_cloud_function_iam.html)
        for Cloud Functions.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] available_memory_mb: Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
        :param pulumi.Input[str] description: Description of the function.
        :param pulumi.Input[str] entry_point: Name of the function that will be executed when the Google Cloud Function is triggered.
        :param pulumi.Input[Mapping[str, Any]] environment_variables: A set of key/value environment variable pairs to assign to the function.
        :param pulumi.Input[pulumi.InputType['FunctionEventTriggerArgs']] event_trigger: A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `trigger_http`.
        :param pulumi.Input[str] https_trigger_url: URL which triggers function execution. Returned only if `trigger_http` is used.
        :param pulumi.Input[str] ingress_settings: String value that controls what traffic can reach the function. Allowed values are ALLOW_ALL and ALLOW_INTERNAL_ONLY. Changes to this field will recreate the cloud function.
        :param pulumi.Input[Mapping[str, Any]] labels: A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
        :param pulumi.Input[float] max_instances: The limit on the maximum number of function instances that may coexist at a given time.
        :param pulumi.Input[str] name: A user-defined name of the function. Function names must be unique globally.
        :param pulumi.Input[str] project: Project of the function. If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Region of function. Currently can be only "us-central1". If it is not provided, the provider region is used.
        :param pulumi.Input[str] runtime: The runtime in which the function is going to run.
               Eg. `"nodejs8"`, `"nodejs10"`, `"python37"`, `"go111"`, `"go113"`.
        :param pulumi.Input[str] service_account_email: If provided, the self-provided service account to run the function with.
        :param pulumi.Input[str] source_archive_bucket: The GCS bucket containing the zip archive which contains the function.
        :param pulumi.Input[str] source_archive_object: The source archive object (file) in archive bucket.
        :param pulumi.Input[pulumi.InputType['FunctionSourceRepositoryArgs']] source_repository: Represents parameters related to source repository where a function is hosted.
               Cannot be set alongside `source_archive_bucket` or `source_archive_object`. Structure is documented below.
        :param pulumi.Input[float] timeout: Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
        :param pulumi.Input[bool] trigger_http: Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `https_trigger_url`. Cannot be used with `trigger_bucket` and `trigger_topic`.
        :param pulumi.Input[str] vpc_connector: The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is `projects/*/locations/*/connectors/*`.
        :param pulumi.Input[str] vpc_connector_egress_settings: The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are `ALL_TRAFFIC` and `PRIVATE_RANGES_ONLY`. Defaults to `PRIVATE_RANGES_ONLY`. If unset, this field preserves the previously set value.
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

            __props__['available_memory_mb'] = available_memory_mb
            __props__['description'] = description
            __props__['entry_point'] = entry_point
            __props__['environment_variables'] = environment_variables
            __props__['event_trigger'] = event_trigger
            __props__['https_trigger_url'] = https_trigger_url
            __props__['ingress_settings'] = ingress_settings
            __props__['labels'] = labels
            __props__['max_instances'] = max_instances
            __props__['name'] = name
            __props__['project'] = project
            __props__['region'] = region
            if runtime is None:
                raise TypeError("Missing required property 'runtime'")
            __props__['runtime'] = runtime
            __props__['service_account_email'] = service_account_email
            __props__['source_archive_bucket'] = source_archive_bucket
            __props__['source_archive_object'] = source_archive_object
            __props__['source_repository'] = source_repository
            __props__['timeout'] = timeout
            __props__['trigger_http'] = trigger_http
            __props__['vpc_connector'] = vpc_connector
            __props__['vpc_connector_egress_settings'] = vpc_connector_egress_settings
        super(Function, __self__).__init__(
            'gcp:cloudfunctions/function:Function',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            available_memory_mb: Optional[pulumi.Input[float]] = None,
            description: Optional[pulumi.Input[str]] = None,
            entry_point: Optional[pulumi.Input[str]] = None,
            environment_variables: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            event_trigger: Optional[pulumi.Input[pulumi.InputType['FunctionEventTriggerArgs']]] = None,
            https_trigger_url: Optional[pulumi.Input[str]] = None,
            ingress_settings: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            max_instances: Optional[pulumi.Input[float]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            runtime: Optional[pulumi.Input[str]] = None,
            service_account_email: Optional[pulumi.Input[str]] = None,
            source_archive_bucket: Optional[pulumi.Input[str]] = None,
            source_archive_object: Optional[pulumi.Input[str]] = None,
            source_repository: Optional[pulumi.Input[pulumi.InputType['FunctionSourceRepositoryArgs']]] = None,
            timeout: Optional[pulumi.Input[float]] = None,
            trigger_http: Optional[pulumi.Input[bool]] = None,
            vpc_connector: Optional[pulumi.Input[str]] = None,
            vpc_connector_egress_settings: Optional[pulumi.Input[str]] = None) -> 'Function':
        """
        Get an existing Function resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] available_memory_mb: Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
        :param pulumi.Input[str] description: Description of the function.
        :param pulumi.Input[str] entry_point: Name of the function that will be executed when the Google Cloud Function is triggered.
        :param pulumi.Input[Mapping[str, Any]] environment_variables: A set of key/value environment variable pairs to assign to the function.
        :param pulumi.Input[pulumi.InputType['FunctionEventTriggerArgs']] event_trigger: A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `trigger_http`.
        :param pulumi.Input[str] https_trigger_url: URL which triggers function execution. Returned only if `trigger_http` is used.
        :param pulumi.Input[str] ingress_settings: String value that controls what traffic can reach the function. Allowed values are ALLOW_ALL and ALLOW_INTERNAL_ONLY. Changes to this field will recreate the cloud function.
        :param pulumi.Input[Mapping[str, Any]] labels: A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
        :param pulumi.Input[float] max_instances: The limit on the maximum number of function instances that may coexist at a given time.
        :param pulumi.Input[str] name: A user-defined name of the function. Function names must be unique globally.
        :param pulumi.Input[str] project: Project of the function. If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Region of function. Currently can be only "us-central1". If it is not provided, the provider region is used.
        :param pulumi.Input[str] runtime: The runtime in which the function is going to run.
               Eg. `"nodejs8"`, `"nodejs10"`, `"python37"`, `"go111"`, `"go113"`.
        :param pulumi.Input[str] service_account_email: If provided, the self-provided service account to run the function with.
        :param pulumi.Input[str] source_archive_bucket: The GCS bucket containing the zip archive which contains the function.
        :param pulumi.Input[str] source_archive_object: The source archive object (file) in archive bucket.
        :param pulumi.Input[pulumi.InputType['FunctionSourceRepositoryArgs']] source_repository: Represents parameters related to source repository where a function is hosted.
               Cannot be set alongside `source_archive_bucket` or `source_archive_object`. Structure is documented below.
        :param pulumi.Input[float] timeout: Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
        :param pulumi.Input[bool] trigger_http: Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `https_trigger_url`. Cannot be used with `trigger_bucket` and `trigger_topic`.
        :param pulumi.Input[str] vpc_connector: The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is `projects/*/locations/*/connectors/*`.
        :param pulumi.Input[str] vpc_connector_egress_settings: The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are `ALL_TRAFFIC` and `PRIVATE_RANGES_ONLY`. Defaults to `PRIVATE_RANGES_ONLY`. If unset, this field preserves the previously set value.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["available_memory_mb"] = available_memory_mb
        __props__["description"] = description
        __props__["entry_point"] = entry_point
        __props__["environment_variables"] = environment_variables
        __props__["event_trigger"] = event_trigger
        __props__["https_trigger_url"] = https_trigger_url
        __props__["ingress_settings"] = ingress_settings
        __props__["labels"] = labels
        __props__["max_instances"] = max_instances
        __props__["name"] = name
        __props__["project"] = project
        __props__["region"] = region
        __props__["runtime"] = runtime
        __props__["service_account_email"] = service_account_email
        __props__["source_archive_bucket"] = source_archive_bucket
        __props__["source_archive_object"] = source_archive_object
        __props__["source_repository"] = source_repository
        __props__["timeout"] = timeout
        __props__["trigger_http"] = trigger_http
        __props__["vpc_connector"] = vpc_connector
        __props__["vpc_connector_egress_settings"] = vpc_connector_egress_settings
        return Function(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availableMemoryMb")
    def available_memory_mb(self) -> Optional[float]:
        """
        Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
        """
        return pulumi.get(self, "available_memory_mb")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of the function.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="entryPoint")
    def entry_point(self) -> Optional[str]:
        """
        Name of the function that will be executed when the Google Cloud Function is triggered.
        """
        return pulumi.get(self, "entry_point")

    @property
    @pulumi.getter(name="environmentVariables")
    def environment_variables(self) -> Optional[Mapping[str, Any]]:
        """
        A set of key/value environment variable pairs to assign to the function.
        """
        return pulumi.get(self, "environment_variables")

    @property
    @pulumi.getter(name="eventTrigger")
    def event_trigger(self) -> 'outputs.FunctionEventTrigger':
        """
        A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `trigger_http`.
        """
        return pulumi.get(self, "event_trigger")

    @property
    @pulumi.getter(name="httpsTriggerUrl")
    def https_trigger_url(self) -> str:
        """
        URL which triggers function execution. Returned only if `trigger_http` is used.
        """
        return pulumi.get(self, "https_trigger_url")

    @property
    @pulumi.getter(name="ingressSettings")
    def ingress_settings(self) -> Optional[str]:
        """
        String value that controls what traffic can reach the function. Allowed values are ALLOW_ALL and ALLOW_INTERNAL_ONLY. Changes to this field will recreate the cloud function.
        """
        return pulumi.get(self, "ingress_settings")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Mapping[str, Any]]:
        """
        A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="maxInstances")
    def max_instances(self) -> Optional[float]:
        """
        The limit on the maximum number of function instances that may coexist at a given time.
        """
        return pulumi.get(self, "max_instances")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A user-defined name of the function. Function names must be unique globally.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        Project of the function. If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        Region of function. Currently can be only "us-central1". If it is not provided, the provider region is used.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def runtime(self) -> str:
        """
        The runtime in which the function is going to run.
        Eg. `"nodejs8"`, `"nodejs10"`, `"python37"`, `"go111"`, `"go113"`.
        """
        return pulumi.get(self, "runtime")

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> str:
        """
        If provided, the self-provided service account to run the function with.
        """
        return pulumi.get(self, "service_account_email")

    @property
    @pulumi.getter(name="sourceArchiveBucket")
    def source_archive_bucket(self) -> Optional[str]:
        """
        The GCS bucket containing the zip archive which contains the function.
        """
        return pulumi.get(self, "source_archive_bucket")

    @property
    @pulumi.getter(name="sourceArchiveObject")
    def source_archive_object(self) -> Optional[str]:
        """
        The source archive object (file) in archive bucket.
        """
        return pulumi.get(self, "source_archive_object")

    @property
    @pulumi.getter(name="sourceRepository")
    def source_repository(self) -> Optional['outputs.FunctionSourceRepository']:
        """
        Represents parameters related to source repository where a function is hosted.
        Cannot be set alongside `source_archive_bucket` or `source_archive_object`. Structure is documented below.
        """
        return pulumi.get(self, "source_repository")

    @property
    @pulumi.getter
    def timeout(self) -> Optional[float]:
        """
        Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter(name="triggerHttp")
    def trigger_http(self) -> Optional[bool]:
        """
        Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `https_trigger_url`. Cannot be used with `trigger_bucket` and `trigger_topic`.
        """
        return pulumi.get(self, "trigger_http")

    @property
    @pulumi.getter(name="vpcConnector")
    def vpc_connector(self) -> Optional[str]:
        """
        The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is `projects/*/locations/*/connectors/*`.
        """
        return pulumi.get(self, "vpc_connector")

    @property
    @pulumi.getter(name="vpcConnectorEgressSettings")
    def vpc_connector_egress_settings(self) -> str:
        """
        The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are `ALL_TRAFFIC` and `PRIVATE_RANGES_ONLY`. Defaults to `PRIVATE_RANGES_ONLY`. If unset, this field preserves the previously set value.
        """
        return pulumi.get(self, "vpc_connector_egress_settings")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

