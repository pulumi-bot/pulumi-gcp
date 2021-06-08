# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ServiceArgs', 'Service']

@pulumi.input_type
class ServiceArgs:
    def __init__(__self__, *,
                 namespace: pulumi.Input[str],
                 service_id: pulumi.Input[str],
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Service resource.
        :param pulumi.Input[str] namespace: The resource name of the namespace this service will belong to.
        :param pulumi.Input[str] service_id: The Resource ID must be 1-63 characters long, including digits,
               lowercase letters or the hyphen character.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Metadata for the service. This data can be consumed
               by service clients. The entire metadata dictionary may contain
               up to 2000 characters, spread across all key-value pairs.
               Metadata that goes beyond any these limits will be rejected.
        """
        pulumi.set(__self__, "namespace", namespace)
        pulumi.set(__self__, "service_id", service_id)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Input[str]:
        """
        The resource name of the namespace this service will belong to.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Input[str]:
        """
        The Resource ID must be 1-63 characters long, including digits,
        lowercase letters or the hyphen character.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_id", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Metadata for the service. This data can be consumed
        by service clients. The entire metadata dictionary may contain
        up to 2000 characters, spread across all key-value pairs.
        Metadata that goes beyond any these limits will be rejected.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "metadata", value)


@pulumi.input_type
class _ServiceState:
    def __init__(__self__, *,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Service resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Metadata for the service. This data can be consumed
               by service clients. The entire metadata dictionary may contain
               up to 2000 characters, spread across all key-value pairs.
               Metadata that goes beyond any these limits will be rejected.
        :param pulumi.Input[str] name: The resource name for the service in the format 'projects/*/locations/*/namespaces/*/services/*'.
        :param pulumi.Input[str] namespace: The resource name of the namespace this service will belong to.
        :param pulumi.Input[str] service_id: The Resource ID must be 1-63 characters long, including digits,
               lowercase letters or the hyphen character.
        """
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Metadata for the service. This data can be consumed
        by service clients. The entire metadata dictionary may contain
        up to 2000 characters, spread across all key-value pairs.
        Metadata that goes beyond any these limits will be rejected.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name for the service in the format 'projects/*/locations/*/namespaces/*/services/*'.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the namespace this service will belong to.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Resource ID must be 1-63 characters long, including digits,
        lowercase letters or the hyphen character.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_id", value)


class Service(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        An individual service. A service contains a name and optional metadata.

        To get more information about Service, see:

        * [API documentation](https://cloud.google.com/service-directory/docs/reference/rest/v1beta1/projects.locations.namespaces.services)
        * How-to Guides
            * [Configuring a service](https://cloud.google.com/service-directory/docs/configuring-service-directory#configuring_a_service)

        ## Example Usage
        ### Service Directory Service Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example_namespace = gcp.servicedirectory.Namespace("exampleNamespace",
            namespace_id="example-namespace",
            location="us-central1",
            opts=pulumi.ResourceOptions(provider=google_beta))
        example_service = gcp.servicedirectory.Service("exampleService",
            service_id="example-service",
            namespace=example_namespace.id,
            metadata={
                "stage": "prod",
                "region": "us-central1",
            },
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        Service can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:servicedirectory/service:Service default projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}
        ```

        ```sh
         $ pulumi import gcp:servicedirectory/service:Service default {{project}}/{{location}}/{{namespace_id}}/{{service_id}}
        ```

        ```sh
         $ pulumi import gcp:servicedirectory/service:Service default {{location}}/{{namespace_id}}/{{service_id}}
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Metadata for the service. This data can be consumed
               by service clients. The entire metadata dictionary may contain
               up to 2000 characters, spread across all key-value pairs.
               Metadata that goes beyond any these limits will be rejected.
        :param pulumi.Input[str] namespace: The resource name of the namespace this service will belong to.
        :param pulumi.Input[str] service_id: The Resource ID must be 1-63 characters long, including digits,
               lowercase letters or the hyphen character.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: ServiceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An individual service. A service contains a name and optional metadata.

        To get more information about Service, see:

        * [API documentation](https://cloud.google.com/service-directory/docs/reference/rest/v1beta1/projects.locations.namespaces.services)
        * How-to Guides
            * [Configuring a service](https://cloud.google.com/service-directory/docs/configuring-service-directory#configuring_a_service)

        ## Example Usage
        ### Service Directory Service Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example_namespace = gcp.servicedirectory.Namespace("exampleNamespace",
            namespace_id="example-namespace",
            location="us-central1",
            opts=pulumi.ResourceOptions(provider=google_beta))
        example_service = gcp.servicedirectory.Service("exampleService",
            service_id="example-service",
            namespace=example_namespace.id,
            metadata={
                "stage": "prod",
                "region": "us-central1",
            },
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        Service can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:servicedirectory/service:Service default projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}
        ```

        ```sh
         $ pulumi import gcp:servicedirectory/service:Service default {{project}}/{{location}}/{{namespace_id}}/{{service_id}}
        ```

        ```sh
         $ pulumi import gcp:servicedirectory/service:Service default {{location}}/{{namespace_id}}/{{service_id}}
        ```

        :param str resource_name_: The name of the resource.
        :param ServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = ServiceArgs.__new__(ServiceArgs)

            __props__.__dict__["metadata"] = metadata
            if namespace is None and not opts.urn:
                raise TypeError("Missing required property 'namespace'")
            __props__.__dict__["namespace"] = namespace
            if service_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_id'")
            __props__.__dict__["service_id"] = service_id
            __props__.__dict__["name"] = None
        super(Service, __self__).__init__(
            'gcp:servicedirectory/service:Service',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            service_id: Optional[pulumi.Input[str]] = None) -> 'Service':
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Metadata for the service. This data can be consumed
               by service clients. The entire metadata dictionary may contain
               up to 2000 characters, spread across all key-value pairs.
               Metadata that goes beyond any these limits will be rejected.
        :param pulumi.Input[str] name: The resource name for the service in the format 'projects/*/locations/*/namespaces/*/services/*'.
        :param pulumi.Input[str] namespace: The resource name of the namespace this service will belong to.
        :param pulumi.Input[str] service_id: The Resource ID must be 1-63 characters long, including digits,
               lowercase letters or the hyphen character.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceState.__new__(_ServiceState)

        __props__.__dict__["metadata"] = metadata
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["service_id"] = service_id
        return Service(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Metadata for the service. This data can be consumed
        by service clients. The entire metadata dictionary may contain
        up to 2000 characters, spread across all key-value pairs.
        Metadata that goes beyond any these limits will be rejected.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name for the service in the format 'projects/*/locations/*/namespaces/*/services/*'.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[str]:
        """
        The resource name of the namespace this service will belong to.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Output[str]:
        """
        The Resource ID must be 1-63 characters long, including digits,
        lowercase letters or the hyphen character.
        """
        return pulumi.get(self, "service_id")

