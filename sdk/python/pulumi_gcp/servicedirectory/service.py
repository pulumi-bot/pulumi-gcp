# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Service']


class Service(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        An individual service. A service contains a name and optional metadata.

        To get more information about Service, see:

        * [API documentation](https://cloud.google.com/service-directory/docs/reference/rest/v1beta1/projects.locations.namespaces.services)
        * How-to Guides
            * [Configuring a service](https://cloud.google.com/service-directory/docs/configuring-service-directory#configuring_a_service)

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Metadata for the service. This data can be consumed
               by service clients. The entire metadata dictionary may contain
               up to 2000 characters, spread across all key-value pairs.
               Metadata that goes beyond any these limits will be rejected.
        :param pulumi.Input[str] namespace: The resource name of the namespace this service will belong to.
        :param pulumi.Input[str] service_id: The Resource ID must be 1-63 characters long, including digits,
               lowercase letters or the hyphen character.
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

            __props__['metadata'] = metadata
            if namespace is None:
                raise TypeError("Missing required property 'namespace'")
            __props__['namespace'] = namespace
            if service_id is None:
                raise TypeError("Missing required property 'service_id'")
            __props__['service_id'] = service_id
            __props__['name'] = None
        super(Service, __self__).__init__(
            'gcp:servicedirectory/service:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            service_id: Optional[pulumi.Input[str]] = None) -> 'Service':
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
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

        __props__ = dict()

        __props__["metadata"] = metadata
        __props__["name"] = name
        __props__["namespace"] = namespace
        __props__["service_id"] = service_id
        return Service(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[Mapping[str, str]]:
        """
        Metadata for the service. This data can be consumed
        by service clients. The entire metadata dictionary may contain
        up to 2000 characters, spread across all key-value pairs.
        Metadata that goes beyond any these limits will be rejected.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name for the service in the format 'projects/*/locations/*/namespaces/*/services/*'.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> str:
        """
        The resource name of the namespace this service will belong to.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> str:
        """
        The Resource ID must be 1-63 characters long, including digits,
        lowercase letters or the hyphen character.
        """
        return pulumi.get(self, "service_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

