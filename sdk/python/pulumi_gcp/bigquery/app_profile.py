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

__all__ = ['AppProfile']


class AppProfile(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_profile_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ignore_warnings: Optional[pulumi.Input[bool]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 multi_cluster_routing_use_any: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 single_cluster_routing: Optional[pulumi.Input[pulumi.InputType['AppProfileSingleClusterRoutingArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        App profile is a configuration object describing how Cloud Bigtable should treat traffic from a particular end user application.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_profile_id: The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
        :param pulumi.Input[str] description: Long form description of the use case for this app profile.
        :param pulumi.Input[bool] ignore_warnings: If true, ignore safety checks when deleting/updating the app profile.
        :param pulumi.Input[str] instance: The name of the instance to create the app profile within.
        :param pulumi.Input[bool] multi_cluster_routing_use_any: If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
               in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
               consistency to improve availability.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['AppProfileSingleClusterRoutingArgs']] single_cluster_routing: Use a single-cluster routing policy.
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

            if app_profile_id is None:
                raise TypeError("Missing required property 'app_profile_id'")
            __props__['app_profile_id'] = app_profile_id
            __props__['description'] = description
            __props__['ignore_warnings'] = ignore_warnings
            __props__['instance'] = instance
            __props__['multi_cluster_routing_use_any'] = multi_cluster_routing_use_any
            __props__['project'] = project
            __props__['single_cluster_routing'] = single_cluster_routing
            __props__['name'] = None
        super(AppProfile, __self__).__init__(
            'gcp:bigquery/appProfile:AppProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_profile_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            ignore_warnings: Optional[pulumi.Input[bool]] = None,
            instance: Optional[pulumi.Input[str]] = None,
            multi_cluster_routing_use_any: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            single_cluster_routing: Optional[pulumi.Input[pulumi.InputType['AppProfileSingleClusterRoutingArgs']]] = None) -> 'AppProfile':
        """
        Get an existing AppProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_profile_id: The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
        :param pulumi.Input[str] description: Long form description of the use case for this app profile.
        :param pulumi.Input[bool] ignore_warnings: If true, ignore safety checks when deleting/updating the app profile.
        :param pulumi.Input[str] instance: The name of the instance to create the app profile within.
        :param pulumi.Input[bool] multi_cluster_routing_use_any: If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
               in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
               consistency to improve availability.
        :param pulumi.Input[str] name: The unique name of the requested app profile. Values are of the form
               'projects/<project>/instances/<instance>/appProfiles/<appProfileId>'.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['AppProfileSingleClusterRoutingArgs']] single_cluster_routing: Use a single-cluster routing policy.
               Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_profile_id"] = app_profile_id
        __props__["description"] = description
        __props__["ignore_warnings"] = ignore_warnings
        __props__["instance"] = instance
        __props__["multi_cluster_routing_use_any"] = multi_cluster_routing_use_any
        __props__["name"] = name
        __props__["project"] = project
        __props__["single_cluster_routing"] = single_cluster_routing
        return AppProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appProfileId")
    def app_profile_id(self) -> str:
        """
        The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
        """
        return pulumi.get(self, "app_profile_id")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Long form description of the use case for this app profile.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ignoreWarnings")
    def ignore_warnings(self) -> Optional[bool]:
        """
        If true, ignore safety checks when deleting/updating the app profile.
        """
        return pulumi.get(self, "ignore_warnings")

    @property
    @pulumi.getter
    def instance(self) -> Optional[str]:
        """
        The name of the instance to create the app profile within.
        """
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter(name="multiClusterRoutingUseAny")
    def multi_cluster_routing_use_any(self) -> Optional[bool]:
        """
        If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
        in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
        consistency to improve availability.
        """
        return pulumi.get(self, "multi_cluster_routing_use_any")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The unique name of the requested app profile. Values are of the form
        'projects/<project>/instances/<instance>/appProfiles/<appProfileId>'.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="singleClusterRouting")
    def single_cluster_routing(self) -> Optional['outputs.AppProfileSingleClusterRouting']:
        """
        Use a single-cluster routing policy.
        Structure is documented below.
        """
        return pulumi.get(self, "single_cluster_routing")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

