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

__all__ = ['StandardAppVersion']


class StandardAppVersion(pulumi.CustomResource):
    automatic_scaling: pulumi.Output[Optional['outputs.StandardAppVersionAutomaticScaling']] = pulumi.output_property("automaticScaling")
    """
    Automatic scaling is based on request rate, response latencies, and other application metrics.  Structure is documented below.
    """
    basic_scaling: pulumi.Output[Optional['outputs.StandardAppVersionBasicScaling']] = pulumi.output_property("basicScaling")
    """
    Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.  Structure is documented below.
    """
    delete_service_on_destroy: pulumi.Output[Optional[bool]] = pulumi.output_property("deleteServiceOnDestroy")
    """
    If set to `true`, the service will be deleted if it is the last version.
    """
    deployment: pulumi.Output['outputs.StandardAppVersionDeployment'] = pulumi.output_property("deployment")
    """
    Code and application artifacts that make up this version.  Structure is documented below.
    """
    entrypoint: pulumi.Output[Optional['outputs.StandardAppVersionEntrypoint']] = pulumi.output_property("entrypoint")
    """
    The entrypoint for the application.  Structure is documented below.
    """
    env_variables: pulumi.Output[Optional[Dict[str, str]]] = pulumi.output_property("envVariables")
    """
    Environment variables available to the application.
    """
    handlers: pulumi.Output[List['outputs.StandardAppVersionHandler']] = pulumi.output_property("handlers")
    """
    An ordered list of URL-matching patterns that should be applied to incoming requests.
    The first matching URL handles the request and other request handlers are not attempted.  Structure is documented below.
    """
    inbound_services: pulumi.Output[Optional[List[str]]] = pulumi.output_property("inboundServices")
    """
    Before an application can receive email or XMPP messages, the application must be configured to enable the service.
    """
    instance_class: pulumi.Output[str] = pulumi.output_property("instanceClass")
    """
    Instance class that is used to run this version. Valid values are
    AutomaticScaling: F1, F2, F4, F4_1G
    BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
    Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
    """
    libraries: pulumi.Output[Optional[List['outputs.StandardAppVersionLibrary']]] = pulumi.output_property("libraries")
    """
    Configuration for third-party Python runtime libraries that are required by the application.  Structure is documented below.
    """
    manual_scaling: pulumi.Output[Optional['outputs.StandardAppVersionManualScaling']] = pulumi.output_property("manualScaling")
    """
    A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.  Structure is documented below.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    Name of the library. Example "django".
    """
    noop_on_destroy: pulumi.Output[Optional[bool]] = pulumi.output_property("noopOnDestroy")
    """
    If set to `true`, the application version will not be deleted.
    """
    project: pulumi.Output[str] = pulumi.output_property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    runtime: pulumi.Output[str] = pulumi.output_property("runtime")
    """
    Desired runtime. Example python27.
    """
    runtime_api_version: pulumi.Output[Optional[str]] = pulumi.output_property("runtimeApiVersion")
    """
    The version of the API in the given runtime environment.
    Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
    """
    service: pulumi.Output[str] = pulumi.output_property("service")
    """
    AppEngine service resource
    """
    threadsafe: pulumi.Output[Optional[bool]] = pulumi.output_property("threadsafe")
    """
    Whether multiple requests can be dispatched to this version at once.
    """
    version_id: pulumi.Output[Optional[str]] = pulumi.output_property("versionId")
    """
    Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, automatic_scaling: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionAutomaticScalingArgs']]] = None, basic_scaling: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionBasicScalingArgs']]] = None, delete_service_on_destroy: Optional[pulumi.Input[bool]] = None, deployment: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionDeploymentArgs']]] = None, entrypoint: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionEntrypointArgs']]] = None, env_variables: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, handlers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['StandardAppVersionHandlerArgs']]]]] = None, inbound_services: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, instance_class: Optional[pulumi.Input[str]] = None, libraries: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['StandardAppVersionLibraryArgs']]]]] = None, manual_scaling: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionManualScalingArgs']]] = None, noop_on_destroy: Optional[pulumi.Input[bool]] = None, project: Optional[pulumi.Input[str]] = None, runtime: Optional[pulumi.Input[str]] = None, runtime_api_version: Optional[pulumi.Input[str]] = None, service: Optional[pulumi.Input[str]] = None, threadsafe: Optional[pulumi.Input[bool]] = None, version_id: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
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
        ### App Engine Standard App Version

        ```python
        import pulumi
        import pulumi_gcp as gcp

        bucket = gcp.storage.Bucket("bucket")
        object = gcp.storage.BucketObject("object",
            bucket=bucket.name,
            source=pulumi.FileAsset("./test-fixtures/appengine/hello-world.zip"))
        myapp_v1 = gcp.appengine.StandardAppVersion("myappV1",
            version_id="v1",
            service="myapp",
            runtime="nodejs10",
            entrypoint={
                "shell": "node ./app.js",
            },
            deployment={
                "zip": {
                    "sourceUrl": pulumi.Output.all(bucket.name, object.name).apply(lambda bucketName, objectName: f"https://storage.googleapis.com/{bucket_name}/{object_name}"),
                },
            },
            env_variables={
                "port": "8080",
            },
            automatic_scaling={
                "maxConcurrentRequests": 10,
                "minIdleInstances": 1,
                "maxIdleInstances": 3,
                "minPendingLatency": "1s",
                "maxPendingLatency": "5s",
                "standardSchedulerSettings": {
                    "targetCpuUtilization": 0.5,
                    "targetThroughputUtilization": 0.75,
                    "minInstances": 2,
                    "max_instances": 10,
                },
            },
            delete_service_on_destroy=True)
        myapp_v2 = gcp.appengine.StandardAppVersion("myappV2",
            version_id="v2",
            service="myapp",
            runtime="nodejs10",
            entrypoint={
                "shell": "node ./app.js",
            },
            deployment={
                "zip": {
                    "sourceUrl": pulumi.Output.all(bucket.name, object.name).apply(lambda bucketName, objectName: f"https://storage.googleapis.com/{bucket_name}/{object_name}"),
                },
            },
            env_variables={
                "port": "8080",
            },
            basic_scaling={
                "max_instances": 5,
            },
            noop_on_destroy=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionAutomaticScalingArgs']] automatic_scaling: Automatic scaling is based on request rate, response latencies, and other application metrics.  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionBasicScalingArgs']] basic_scaling: Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.  Structure is documented below.
        :param pulumi.Input[bool] delete_service_on_destroy: If set to `true`, the service will be deleted if it is the last version.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionDeploymentArgs']] deployment: Code and application artifacts that make up this version.  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionEntrypointArgs']] entrypoint: The entrypoint for the application.  Structure is documented below.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] env_variables: Environment variables available to the application.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['StandardAppVersionHandlerArgs']]]] handlers: An ordered list of URL-matching patterns that should be applied to incoming requests.
               The first matching URL handles the request and other request handlers are not attempted.  Structure is documented below.
        :param pulumi.Input[List[pulumi.Input[str]]] inbound_services: Before an application can receive email or XMPP messages, the application must be configured to enable the service.
        :param pulumi.Input[str] instance_class: Instance class that is used to run this version. Valid values are
               AutomaticScaling: F1, F2, F4, F4_1G
               BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
               Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['StandardAppVersionLibraryArgs']]]] libraries: Configuration for third-party Python runtime libraries that are required by the application.  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionManualScalingArgs']] manual_scaling: A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.  Structure is documented below.
        :param pulumi.Input[bool] noop_on_destroy: If set to `true`, the application version will not be deleted.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] runtime: Desired runtime. Example python27.
        :param pulumi.Input[str] runtime_api_version: The version of the API in the given runtime environment.
               Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
        :param pulumi.Input[str] service: AppEngine service resource
        :param pulumi.Input[bool] threadsafe: Whether multiple requests can be dispatched to this version at once.
        :param pulumi.Input[str] version_id: Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
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
            __props__['name'] = None
        super(StandardAppVersion, __self__).__init__(
            'gcp:appengine/standardAppVersion:StandardAppVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, automatic_scaling: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionAutomaticScalingArgs']]] = None, basic_scaling: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionBasicScalingArgs']]] = None, delete_service_on_destroy: Optional[pulumi.Input[bool]] = None, deployment: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionDeploymentArgs']]] = None, entrypoint: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionEntrypointArgs']]] = None, env_variables: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, handlers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['StandardAppVersionHandlerArgs']]]]] = None, inbound_services: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, instance_class: Optional[pulumi.Input[str]] = None, libraries: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['StandardAppVersionLibraryArgs']]]]] = None, manual_scaling: Optional[pulumi.Input[pulumi.InputType['StandardAppVersionManualScalingArgs']]] = None, name: Optional[pulumi.Input[str]] = None, noop_on_destroy: Optional[pulumi.Input[bool]] = None, project: Optional[pulumi.Input[str]] = None, runtime: Optional[pulumi.Input[str]] = None, runtime_api_version: Optional[pulumi.Input[str]] = None, service: Optional[pulumi.Input[str]] = None, threadsafe: Optional[pulumi.Input[bool]] = None, version_id: Optional[pulumi.Input[str]] = None) -> 'StandardAppVersion':
        """
        Get an existing StandardAppVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionAutomaticScalingArgs']] automatic_scaling: Automatic scaling is based on request rate, response latencies, and other application metrics.  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionBasicScalingArgs']] basic_scaling: Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.  Structure is documented below.
        :param pulumi.Input[bool] delete_service_on_destroy: If set to `true`, the service will be deleted if it is the last version.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionDeploymentArgs']] deployment: Code and application artifacts that make up this version.  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionEntrypointArgs']] entrypoint: The entrypoint for the application.  Structure is documented below.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] env_variables: Environment variables available to the application.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['StandardAppVersionHandlerArgs']]]] handlers: An ordered list of URL-matching patterns that should be applied to incoming requests.
               The first matching URL handles the request and other request handlers are not attempted.  Structure is documented below.
        :param pulumi.Input[List[pulumi.Input[str]]] inbound_services: Before an application can receive email or XMPP messages, the application must be configured to enable the service.
        :param pulumi.Input[str] instance_class: Instance class that is used to run this version. Valid values are
               AutomaticScaling: F1, F2, F4, F4_1G
               BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
               Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['StandardAppVersionLibraryArgs']]]] libraries: Configuration for third-party Python runtime libraries that are required by the application.  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['StandardAppVersionManualScalingArgs']] manual_scaling: A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.  Structure is documented below.
        :param pulumi.Input[str] name: Name of the library. Example "django".
        :param pulumi.Input[bool] noop_on_destroy: If set to `true`, the application version will not be deleted.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] runtime: Desired runtime. Example python27.
        :param pulumi.Input[str] runtime_api_version: The version of the API in the given runtime environment.
               Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
        :param pulumi.Input[str] service: AppEngine service resource
        :param pulumi.Input[bool] threadsafe: Whether multiple requests can be dispatched to this version at once.
        :param pulumi.Input[str] version_id: Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
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
        return StandardAppVersion(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

