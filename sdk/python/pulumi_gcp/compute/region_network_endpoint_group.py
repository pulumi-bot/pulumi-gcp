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

__all__ = ['RegionNetworkEndpointGroup']


class RegionNetworkEndpointGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_engine: Optional[pulumi.Input[pulumi.InputType['RegionNetworkEndpointGroupAppEngineArgs']]] = None,
                 cloud_function: Optional[pulumi.Input[pulumi.InputType['RegionNetworkEndpointGroupCloudFunctionArgs']]] = None,
                 cloud_run: Optional[pulumi.Input[pulumi.InputType['RegionNetworkEndpointGroupCloudRunArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_endpoint_type: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a RegionNetworkEndpointGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['RegionNetworkEndpointGroupAppEngineArgs']] app_engine: Only valid when networkEndpointType is "SERVERLESS".
               Only one of cloud_run, app_engine or cloud_function may be set.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['RegionNetworkEndpointGroupCloudFunctionArgs']] cloud_function: Only valid when networkEndpointType is "SERVERLESS".
               Only one of cloud_run, app_engine or cloud_function may be set.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['RegionNetworkEndpointGroupCloudRunArgs']] cloud_run: Only valid when networkEndpointType is "SERVERLESS".
               Only one of cloud_run, app_engine or cloud_function may be set.
               Structure is documented below.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource.
        :param pulumi.Input[str] name: Name of the resource; provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] network_endpoint_type: Type of network endpoints in this network endpoint group. Defaults to SERVERLESS
               Default value is `SERVERLESS`.
               Possible values are `SERVERLESS`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: A reference to the region where the Serverless NEGs Reside.
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

            __props__['app_engine'] = app_engine
            __props__['cloud_function'] = cloud_function
            __props__['cloud_run'] = cloud_run
            __props__['description'] = description
            __props__['name'] = name
            __props__['network_endpoint_type'] = network_endpoint_type
            __props__['project'] = project
            if region is None:
                raise TypeError("Missing required property 'region'")
            __props__['region'] = region
            __props__['self_link'] = None
        super(RegionNetworkEndpointGroup, __self__).__init__(
            'gcp:compute/regionNetworkEndpointGroup:RegionNetworkEndpointGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_engine: Optional[pulumi.Input[pulumi.InputType['RegionNetworkEndpointGroupAppEngineArgs']]] = None,
            cloud_function: Optional[pulumi.Input[pulumi.InputType['RegionNetworkEndpointGroupCloudFunctionArgs']]] = None,
            cloud_run: Optional[pulumi.Input[pulumi.InputType['RegionNetworkEndpointGroupCloudRunArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_endpoint_type: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None) -> 'RegionNetworkEndpointGroup':
        """
        Get an existing RegionNetworkEndpointGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['RegionNetworkEndpointGroupAppEngineArgs']] app_engine: Only valid when networkEndpointType is "SERVERLESS".
               Only one of cloud_run, app_engine or cloud_function may be set.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['RegionNetworkEndpointGroupCloudFunctionArgs']] cloud_function: Only valid when networkEndpointType is "SERVERLESS".
               Only one of cloud_run, app_engine or cloud_function may be set.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['RegionNetworkEndpointGroupCloudRunArgs']] cloud_run: Only valid when networkEndpointType is "SERVERLESS".
               Only one of cloud_run, app_engine or cloud_function may be set.
               Structure is documented below.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource.
        :param pulumi.Input[str] name: Name of the resource; provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] network_endpoint_type: Type of network endpoints in this network endpoint group. Defaults to SERVERLESS
               Default value is `SERVERLESS`.
               Possible values are `SERVERLESS`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: A reference to the region where the Serverless NEGs Reside.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_engine"] = app_engine
        __props__["cloud_function"] = cloud_function
        __props__["cloud_run"] = cloud_run
        __props__["description"] = description
        __props__["name"] = name
        __props__["network_endpoint_type"] = network_endpoint_type
        __props__["project"] = project
        __props__["region"] = region
        __props__["self_link"] = self_link
        return RegionNetworkEndpointGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appEngine")
    def app_engine(self) -> pulumi.Output[Optional['outputs.RegionNetworkEndpointGroupAppEngine']]:
        """
        Only valid when networkEndpointType is "SERVERLESS".
        Only one of cloud_run, app_engine or cloud_function may be set.
        Structure is documented below.
        """
        return pulumi.get(self, "app_engine")

    @property
    @pulumi.getter(name="cloudFunction")
    def cloud_function(self) -> pulumi.Output[Optional['outputs.RegionNetworkEndpointGroupCloudFunction']]:
        """
        Only valid when networkEndpointType is "SERVERLESS".
        Only one of cloud_run, app_engine or cloud_function may be set.
        Structure is documented below.
        """
        return pulumi.get(self, "cloud_function")

    @property
    @pulumi.getter(name="cloudRun")
    def cloud_run(self) -> pulumi.Output[Optional['outputs.RegionNetworkEndpointGroupCloudRun']]:
        """
        Only valid when networkEndpointType is "SERVERLESS".
        Only one of cloud_run, app_engine or cloud_function may be set.
        Structure is documented below.
        """
        return pulumi.get(self, "cloud_run")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description of this resource. Provide this property when
        you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource; provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035. Specifically, the name must be 1-63 characters long and match
        the regular expression `a-z?` which means the
        first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkEndpointType")
    def network_endpoint_type(self) -> pulumi.Output[Optional[str]]:
        """
        Type of network endpoints in this network endpoint group. Defaults to SERVERLESS
        Default value is `SERVERLESS`.
        Possible values are `SERVERLESS`.
        """
        return pulumi.get(self, "network_endpoint_type")

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
    def region(self) -> pulumi.Output[str]:
        """
        A reference to the region where the Serverless NEGs Reside.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

