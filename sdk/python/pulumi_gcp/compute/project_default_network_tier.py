# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['ProjectDefaultNetworkTierArgs', 'ProjectDefaultNetworkTier']

@pulumi.input_type
class ProjectDefaultNetworkTierArgs:
    def __init__(__self__, *,
                 network_tier: pulumi.Input[str],
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProjectDefaultNetworkTier resource.
        :param pulumi.Input[str] network_tier: The default network tier to be configured for the project.
               This field can take the following values: `PREMIUM` or `STANDARD`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        """
        pulumi.set(__self__, "network_tier", network_tier)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="networkTier")
    def network_tier(self) -> pulumi.Input[str]:
        """
        The default network tier to be configured for the project.
        This field can take the following values: `PREMIUM` or `STANDARD`.
        """
        return pulumi.get(self, "network_tier")

    @network_tier.setter
    def network_tier(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_tier", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class ProjectDefaultNetworkTier(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 network_tier: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Configures the Google Compute Engine
        [Default Network Tier](https://cloud.google.com/network-tiers/docs/using-network-service-tiers#setting_the_tier_for_all_resources_in_a_project)
        for a project.

        For more information, see,
        [the Project API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/projects/setDefaultNetworkTier).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.compute.ProjectDefaultNetworkTier("default", network_tier="PREMIUM")
        ```

        ## Import

        This resource can be imported using the project ID

        ```sh
         $ pulumi import gcp:compute/projectDefaultNetworkTier:ProjectDefaultNetworkTier default project-id`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] network_tier: The default network tier to be configured for the project.
               This field can take the following values: `PREMIUM` or `STANDARD`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectDefaultNetworkTierArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configures the Google Compute Engine
        [Default Network Tier](https://cloud.google.com/network-tiers/docs/using-network-service-tiers#setting_the_tier_for_all_resources_in_a_project)
        for a project.

        For more information, see,
        [the Project API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/projects/setDefaultNetworkTier).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.compute.ProjectDefaultNetworkTier("default", network_tier="PREMIUM")
        ```

        ## Import

        This resource can be imported using the project ID

        ```sh
         $ pulumi import gcp:compute/projectDefaultNetworkTier:ProjectDefaultNetworkTier default project-id`
        ```

        :param str resource_name: The name of the resource.
        :param ProjectDefaultNetworkTierArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectDefaultNetworkTierArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 network_tier: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
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
            __props__ = dict()

            if network_tier is None and not opts.urn:
                raise TypeError("Missing required property 'network_tier'")
            __props__['network_tier'] = network_tier
            __props__['project'] = project
        super(ProjectDefaultNetworkTier, __self__).__init__(
            'gcp:compute/projectDefaultNetworkTier:ProjectDefaultNetworkTier',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            network_tier: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'ProjectDefaultNetworkTier':
        """
        Get an existing ProjectDefaultNetworkTier resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] network_tier: The default network tier to be configured for the project.
               This field can take the following values: `PREMIUM` or `STANDARD`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["network_tier"] = network_tier
        __props__["project"] = project
        return ProjectDefaultNetworkTier(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="networkTier")
    def network_tier(self) -> pulumi.Output[str]:
        """
        The default network tier to be configured for the project.
        This field can take the following values: `PREMIUM` or `STANDARD`.
        """
        return pulumi.get(self, "network_tier")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

