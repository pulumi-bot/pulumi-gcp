# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['StandardAppVersion']


class StandardAppVersion(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 automatic_scaling: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionAutomaticScalingArgs']]] = None,
                 basic_scaling: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionBasicScalingArgs']]] = None,
                 delete_service_on_destroy: Optional[pulumi.Input[bool]] = None,
                 deployment: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionDeploymentArgs']]] = None,
                 entrypoint: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionEntrypointArgs']]] = None,
                 env_variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 handlers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StandardAppVersionHandlerArgs']]]]] = None,
                 inbound_services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_class: Optional[pulumi.Input[str]] = None,
                 libraries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StandardAppVersionLibraryArgs']]]]] = None,
                 manual_scaling: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionManualScalingArgs']]] = None,
                 noop_on_destroy: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 runtime: Optional[pulumi.Input[str]] = None,
                 runtime_api_version: Optional[pulumi.Input[str]] = None,
                 service: Optional[pulumi.Input[str]] = None,
                 threadsafe: Optional[pulumi.Input[bool]] = None,
                 version_id: Optional[pulumi.Input[str]] = None,
                 vpc_access_connector: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionVpcAccessConnectorArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Standard App Version resource to create a new version of standard GAE Application.
        Learn about the differences between the standard environment and the flexible environment
        at https://cloud.google.com/appengine/docs/the-appengine-environments.
        Currently supporting Zip and File Containers.

        To get more information about StandardAppVersion, see:

        * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/appengine/docs/standard)

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionAutomaticScalingArgs']] automatic_scaling: Automatic scaling is based on request rate, response latencies, and other application metrics.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionBasicScalingArgs']] basic_scaling: Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
               Structure is documented below.
        :param pulumi.Input[bool] delete_service_on_destroy: If set to `true`, the service will be deleted if it is the last version.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionDeploymentArgs']] deployment: Code and application artifacts that make up this version.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionEntrypointArgs']] entrypoint: The entrypoint for the application.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] env_variables: Environment variables available to the application.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StandardAppVersionHandlerArgs']]]] handlers: An ordered list of URL-matching patterns that should be applied to incoming requests.
               The first matching URL handles the request and other request handlers are not attempted.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] inbound_services: A list of the types of messages that this application is able to receive.
               Each value may be one of `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, and `INBOUND_SERVICE_WARMUP`.
        :param pulumi.Input[str] instance_class: Instance class that is used to run this version. Valid values are
               AutomaticScaling: F1, F2, F4, F4_1G
               BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
               Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StandardAppVersionLibraryArgs']]]] libraries: Configuration for third-party Python runtime libraries that are required by the application.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionManualScalingArgs']] manual_scaling: A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
               Structure is documented below.
        :param pulumi.Input[bool] noop_on_destroy: If set to `true`, the application version will not be deleted.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] runtime: Desired runtime. Example python27.
        :param pulumi.Input[str] runtime_api_version: The version of the API in the given runtime environment.
               Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
        :param pulumi.Input[str] service: AppEngine service resource
        :param pulumi.Input[bool] threadsafe: Whether multiple requests can be dispatched to this version at once.
        :param pulumi.Input[str] version_id: Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
        :param pulumi.Input[pulumi.InputType['StandardAppVersionVpcAccessConnectorArgs']] vpc_access_connector: Enables VPC connectivity for standard apps.
               Structure is documented below.
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

            __props__['automatic_scaling'] = automatic_scaling
            __props__['basic_scaling'] = basic_scaling
            __props__['delete_service_on_destroy'] = delete_service_on_destroy
            if deployment is None:
                raise TypeError("Missing required property 'deployment'")
            __props__['deployment'] = deployment
            __props__['entrypoint'] = entrypoint
            __props__['env_variables'] = env_variables
            __props__['handlers'] = handlers
            __props__['inbound_services'] = inbound_services
            __props__['instance_class'] = instance_class
            __props__['libraries'] = libraries
            __props__['manual_scaling'] = manual_scaling
            __props__['noop_on_destroy'] = noop_on_destroy
            __props__['project'] = project
            if runtime is None:
                raise TypeError("Missing required property 'runtime'")
            __props__['runtime'] = runtime
            __props__['runtime_api_version'] = runtime_api_version
            if service is None:
                raise TypeError("Missing required property 'service'")
            __props__['service'] = service
            __props__['threadsafe'] = threadsafe
            __props__['version_id'] = version_id
            __props__['vpc_access_connector'] = vpc_access_connector
            __props__['name'] = None
        super(StandardAppVersion, __self__).__init__(
            'gcp:appengine/standardAppVersion:StandardAppVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            automatic_scaling: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionAutomaticScalingArgs']]] = None,
            basic_scaling: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionBasicScalingArgs']]] = None,
            delete_service_on_destroy: Optional[pulumi.Input[bool]] = None,
            deployment: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionDeploymentArgs']]] = None,
            entrypoint: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionEntrypointArgs']]] = None,
            env_variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            handlers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StandardAppVersionHandlerArgs']]]]] = None,
            inbound_services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            instance_class: Optional[pulumi.Input[str]] = None,
            libraries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StandardAppVersionLibraryArgs']]]]] = None,
            manual_scaling: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionManualScalingArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            noop_on_destroy: Optional[pulumi.Input[bool]] = None,
            project: Optional[pulumi.Input[str]] = None,
            runtime: Optional[pulumi.Input[str]] = None,
            runtime_api_version: Optional[pulumi.Input[str]] = None,
            service: Optional[pulumi.Input[str]] = None,
            threadsafe: Optional[pulumi.Input[bool]] = None,
            version_id: Optional[pulumi.Input[str]] = None,
            vpc_access_connector: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionVpcAccessConnectorArgs']]] = None) -> 'StandardAppVersion':
        """
        Get an existing StandardAppVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionAutomaticScalingArgs']] automatic_scaling: Automatic scaling is based on request rate, response latencies, and other application metrics.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionBasicScalingArgs']] basic_scaling: Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
               Structure is documented below.
        :param pulumi.Input[bool] delete_service_on_destroy: If set to `true`, the service will be deleted if it is the last version.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionDeploymentArgs']] deployment: Code and application artifacts that make up this version.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionEntrypointArgs']] entrypoint: The entrypoint for the application.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] env_variables: Environment variables available to the application.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StandardAppVersionHandlerArgs']]]] handlers: An ordered list of URL-matching patterns that should be applied to incoming requests.
               The first matching URL handles the request and other request handlers are not attempted.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] inbound_services: A list of the types of messages that this application is able to receive.
               Each value may be one of `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, and `INBOUND_SERVICE_WARMUP`.
        :param pulumi.Input[str] instance_class: Instance class that is used to run this version. Valid values are
               AutomaticScaling: F1, F2, F4, F4_1G
               BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
               Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StandardAppVersionLibraryArgs']]]] libraries: Configuration for third-party Python runtime libraries that are required by the application.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionManualScalingArgs']] manual_scaling: A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
               Structure is documented below.
        :param pulumi.Input[str] name: Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
        :param pulumi.Input[bool] noop_on_destroy: If set to `true`, the application version will not be deleted.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] runtime: Desired runtime. Example python27.
        :param pulumi.Input[str] runtime_api_version: The version of the API in the given runtime environment.
               Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
        :param pulumi.Input[str] service: AppEngine service resource
        :param pulumi.Input[bool] threadsafe: Whether multiple requests can be dispatched to this version at once.
        :param pulumi.Input[str] version_id: Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
        :param pulumi.Input[pulumi.InputType['StandardAppVersionVpcAccessConnectorArgs']] vpc_access_connector: Enables VPC connectivity for standard apps.
               Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["automatic_scaling"] = automatic_scaling
        __props__["basic_scaling"] = basic_scaling
        __props__["delete_service_on_destroy"] = delete_service_on_destroy
        __props__["deployment"] = deployment
        __props__["entrypoint"] = entrypoint
        __props__["env_variables"] = env_variables
        __props__["handlers"] = handlers
        __props__["inbound_services"] = inbound_services
        __props__["instance_class"] = instance_class
        __props__["libraries"] = libraries
        __props__["manual_scaling"] = manual_scaling
        __props__["name"] = name
        __props__["noop_on_destroy"] = noop_on_destroy
        __props__["project"] = project
        __props__["runtime"] = runtime
        __props__["runtime_api_version"] = runtime_api_version
        __props__["service"] = service
        __props__["threadsafe"] = threadsafe
        __props__["version_id"] = version_id
        __props__["vpc_access_connector"] = vpc_access_connector
        return StandardAppVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="automaticScaling")
    def automatic_scaling(self) -> pulumi.Output[Optional['outputs.StandardAppVersionAutomaticScaling']]:
        """
        Automatic scaling is based on request rate, response latencies, and other application metrics.
        Structure is documented below.
        """
        return pulumi.get(self, "automatic_scaling")

    @property
    @pulumi.getter(name="basicScaling")
    def basic_scaling(self) -> pulumi.Output[Optional['outputs.StandardAppVersionBasicScaling']]:
        """
        Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
        Structure is documented below.
        """
        return pulumi.get(self, "basic_scaling")

    @property
    @pulumi.getter(name="deleteServiceOnDestroy")
    def delete_service_on_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        If set to `true`, the service will be deleted if it is the last version.
        """
        return pulumi.get(self, "delete_service_on_destroy")

    @property
    @pulumi.getter
    def deployment(self) -> pulumi.Output['outputs.StandardAppVersionDeployment']:
        """
        Code and application artifacts that make up this version.
        Structure is documented below.
        """
        return pulumi.get(self, "deployment")

    @property
    @pulumi.getter
    def entrypoint(self) -> pulumi.Output[Optional['outputs.StandardAppVersionEntrypoint']]:
        """
        The entrypoint for the application.
        Structure is documented below.
        """
        return pulumi.get(self, "entrypoint")

    @property
    @pulumi.getter(name="envVariables")
    def env_variables(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Environment variables available to the application.
        """
        return pulumi.get(self, "env_variables")

    @property
    @pulumi.getter
    def handlers(self) -> pulumi.Output[Sequence['outputs.StandardAppVersionHandler']]:
        """
        An ordered list of URL-matching patterns that should be applied to incoming requests.
        The first matching URL handles the request and other request handlers are not attempted.
        Structure is documented below.
        """
        return pulumi.get(self, "handlers")

    @property
    @pulumi.getter(name="inboundServices")
    def inbound_services(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of the types of messages that this application is able to receive.
        Each value may be one of `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, and `INBOUND_SERVICE_WARMUP`.
        """
        return pulumi.get(self, "inbound_services")

    @property
    @pulumi.getter(name="instanceClass")
    def instance_class(self) -> pulumi.Output[str]:
        """
        Instance class that is used to run this version. Valid values are
        AutomaticScaling: F1, F2, F4, F4_1G
        BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
        Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
        """
        return pulumi.get(self, "instance_class")

    @property
    @pulumi.getter
    def libraries(self) -> pulumi.Output[Optional[Sequence['outputs.StandardAppVersionLibrary']]]:
        """
        Configuration for third-party Python runtime libraries that are required by the application.
        Structure is documented below.
        """
        return pulumi.get(self, "libraries")

    @property
    @pulumi.getter(name="manualScaling")
    def manual_scaling(self) -> pulumi.Output[Optional['outputs.StandardAppVersionManualScaling']]:
        """
        A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
        Structure is documented below.
        """
        return pulumi.get(self, "manual_scaling")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="noopOnDestroy")
    def noop_on_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        If set to `true`, the application version will not be deleted.
        """
        return pulumi.get(self, "noop_on_destroy")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def runtime(self) -> pulumi.Output[str]:
        """
        Desired runtime. Example python27.
        """
        return pulumi.get(self, "runtime")

    @property
    @pulumi.getter(name="runtimeApiVersion")
    def runtime_api_version(self) -> pulumi.Output[Optional[str]]:
        """
        The version of the API in the given runtime environment.
        Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
        """
        return pulumi.get(self, "runtime_api_version")

    @property
    @pulumi.getter
    def service(self) -> pulumi.Output[str]:
        """
        AppEngine service resource
        """
        return pulumi.get(self, "service")

    @property
    @pulumi.getter
    def threadsafe(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether multiple requests can be dispatched to this version at once.
        """
        return pulumi.get(self, "threadsafe")

    @property
    @pulumi.getter(name="versionId")
    def version_id(self) -> pulumi.Output[Optional[str]]:
        """
        Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
        """
        return pulumi.get(self, "version_id")

    @property
    @pulumi.getter(name="vpcAccessConnector")
    def vpc_access_connector(self) -> pulumi.Output[Optional['outputs.StandardAppVersionVpcAccessConnector']]:
        """
        Enables VPC connectivity for standard apps.
        Structure is documented below.
        """
        return pulumi.get(self, "vpc_access_connector")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

