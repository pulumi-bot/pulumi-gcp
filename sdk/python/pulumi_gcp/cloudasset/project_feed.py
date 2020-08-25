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

__all__ = ['ProjectFeed']


class ProjectFeed(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asset_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 asset_types: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 billing_project: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 feed_id: Optional[pulumi.Input[str]] = None,
                 feed_output_config: Optional[pulumi.Input[pulumi.InputType['ProjectFeedFeedOutputConfigArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Describes a Cloud Asset Inventory feed used to to listen to asset updates.

        To get more information about ProjectFeed, see:

        * [API documentation](https://cloud.google.com/asset-inventory/docs/reference/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/asset-inventory/docs)

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] asset_names: A list of the full names of the assets to receive updates. You must specify either or both of
               assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
               exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
               See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
        :param pulumi.Input[List[pulumi.Input[str]]] asset_types: A list of types of the assets to receive updates. You must specify either or both of assetNames
               and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
               the feed. For example: "compute.googleapis.com/Disk"
               See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
               supported asset types.
        :param pulumi.Input[str] billing_project: The project whose identity will be used when sending messages to the
               destination pubsub topic. It also specifies the project for API
               enablement check, quota, and billing. If not specified, the resource's
               project will be used.
        :param pulumi.Input[str] content_type: Asset content type. If not specified, no content but the asset name and type will be returned.
               Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
        :param pulumi.Input[str] feed_id: This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
        :param pulumi.Input[pulumi.InputType['ProjectFeedFeedOutputConfigArgs']] feed_output_config: Output configuration for asset feed destination.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
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

            __props__['asset_names'] = asset_names
            __props__['asset_types'] = asset_types
            __props__['billing_project'] = billing_project
            __props__['content_type'] = content_type
            if feed_id is None:
                raise TypeError("Missing required property 'feed_id'")
            __props__['feed_id'] = feed_id
            if feed_output_config is None:
                raise TypeError("Missing required property 'feed_output_config'")
            __props__['feed_output_config'] = feed_output_config
            __props__['project'] = project
            __props__['name'] = None
        super(ProjectFeed, __self__).__init__(
            'gcp:cloudasset/projectFeed:ProjectFeed',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            asset_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            asset_types: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            billing_project: Optional[pulumi.Input[str]] = None,
            content_type: Optional[pulumi.Input[str]] = None,
            feed_id: Optional[pulumi.Input[str]] = None,
            feed_output_config: Optional[pulumi.Input[pulumi.InputType['ProjectFeedFeedOutputConfigArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'ProjectFeed':
        """
        Get an existing ProjectFeed resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] asset_names: A list of the full names of the assets to receive updates. You must specify either or both of
               assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
               exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
               See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
        :param pulumi.Input[List[pulumi.Input[str]]] asset_types: A list of types of the assets to receive updates. You must specify either or both of assetNames
               and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
               the feed. For example: "compute.googleapis.com/Disk"
               See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
               supported asset types.
        :param pulumi.Input[str] billing_project: The project whose identity will be used when sending messages to the
               destination pubsub topic. It also specifies the project for API
               enablement check, quota, and billing. If not specified, the resource's
               project will be used.
        :param pulumi.Input[str] content_type: Asset content type. If not specified, no content but the asset name and type will be returned.
               Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
        :param pulumi.Input[str] feed_id: This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
        :param pulumi.Input[pulumi.InputType['ProjectFeedFeedOutputConfigArgs']] feed_output_config: Output configuration for asset feed destination.
               Structure is documented below.
        :param pulumi.Input[str] name: The format will be projects/{projectNumber}/feeds/{client-assigned_feed_identifier}.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["asset_names"] = asset_names
        __props__["asset_types"] = asset_types
        __props__["billing_project"] = billing_project
        __props__["content_type"] = content_type
        __props__["feed_id"] = feed_id
        __props__["feed_output_config"] = feed_output_config
        __props__["name"] = name
        __props__["project"] = project
        return ProjectFeed(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="assetNames")
    def asset_names(self) -> Optional[List[str]]:
        """
        A list of the full names of the assets to receive updates. You must specify either or both of
        assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
        exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
        See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
        """
        return pulumi.get(self, "asset_names")

    @property
    @pulumi.getter(name="assetTypes")
    def asset_types(self) -> Optional[List[str]]:
        """
        A list of types of the assets to receive updates. You must specify either or both of assetNames
        and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
        the feed. For example: "compute.googleapis.com/Disk"
        See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
        supported asset types.
        """
        return pulumi.get(self, "asset_types")

    @property
    @pulumi.getter(name="billingProject")
    def billing_project(self) -> Optional[str]:
        """
        The project whose identity will be used when sending messages to the
        destination pubsub topic. It also specifies the project for API
        enablement check, quota, and billing. If not specified, the resource's
        project will be used.
        """
        return pulumi.get(self, "billing_project")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[str]:
        """
        Asset content type. If not specified, no content but the asset name and type will be returned.
        Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter(name="feedId")
    def feed_id(self) -> str:
        """
        This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
        """
        return pulumi.get(self, "feed_id")

    @property
    @pulumi.getter(name="feedOutputConfig")
    def feed_output_config(self) -> 'outputs.ProjectFeedFeedOutputConfig':
        """
        Output configuration for asset feed destination.
        Structure is documented below.
        """
        return pulumi.get(self, "feed_output_config")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The format will be projects/{projectNumber}/feeds/{client-assigned_feed_identifier}.
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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

