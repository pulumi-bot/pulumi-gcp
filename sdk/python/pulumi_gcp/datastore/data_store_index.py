# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['DataStoreIndexArgs', 'DataStoreIndex']

@pulumi.input_type
class DataStoreIndexArgs:
    def __init__(__self__, *,
                 kind: pulumi.Input[str],
                 ancestor: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 properties: Optional[pulumi.Input[Sequence[pulumi.Input['DataStoreIndexPropertyArgs']]]] = None):
        """
        The set of arguments for constructing a DataStoreIndex resource.
        :param pulumi.Input[str] kind: The entity kind which the index applies to.
        :param pulumi.Input[str] ancestor: Policy for including ancestors in the index.
               Default value is `NONE`.
               Possible values are `NONE` and `ALL_ANCESTORS`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Sequence[pulumi.Input['DataStoreIndexPropertyArgs']]] properties: An ordered list of properties to index on.
               Structure is documented below.
        """
        pulumi.set(__self__, "kind", kind)
        if ancestor is not None:
            pulumi.set(__self__, "ancestor", ancestor)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[str]:
        """
        The entity kind which the index applies to.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[str]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def ancestor(self) -> Optional[pulumi.Input[str]]:
        """
        Policy for including ancestors in the index.
        Default value is `NONE`.
        Possible values are `NONE` and `ALL_ANCESTORS`.
        """
        return pulumi.get(self, "ancestor")

    @ancestor.setter
    def ancestor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ancestor", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataStoreIndexPropertyArgs']]]]:
        """
        An ordered list of properties to index on.
        Structure is documented below.
        """
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataStoreIndexPropertyArgs']]]]):
        pulumi.set(self, "properties", value)


class DataStoreIndex(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ancestor: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 properties: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataStoreIndexPropertyArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Describes a composite index for Cloud Datastore.

        To get more information about Index, see:

        * [API documentation](https://cloud.google.com/datastore/docs/reference/admin/rest/v1/projects.indexes)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/datastore/docs/concepts/indexes)

        > **Warning:** This resource creates a Datastore Index on a project that has already
        enabled a Datastore-compatible database. If you haven't already enabled
        one, you can create a `appengine.Application` resource with
        `database_type` set to `"CLOUD_DATASTORE_COMPATIBILITY"` to do so. Your
        Datastore location will be the same as the App Engine location specified.

        ## Example Usage
        ### Datastore Index

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.datastore.DataStoreIndex("default",
            kind="foo",
            properties=[
                gcp.datastore.DataStoreIndexPropertyArgs(
                    direction="ASCENDING",
                    name="property_a",
                ),
                gcp.datastore.DataStoreIndexPropertyArgs(
                    direction="ASCENDING",
                    name="property_b",
                ),
            ])
        ```

        ## Import

        Index can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:datastore/dataStoreIndex:DataStoreIndex default projects/{{project}}/indexes/{{index_id}}
        ```

        ```sh
         $ pulumi import gcp:datastore/dataStoreIndex:DataStoreIndex default {{project}}/{{index_id}}
        ```

        ```sh
         $ pulumi import gcp:datastore/dataStoreIndex:DataStoreIndex default {{index_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ancestor: Policy for including ancestors in the index.
               Default value is `NONE`.
               Possible values are `NONE` and `ALL_ANCESTORS`.
        :param pulumi.Input[str] kind: The entity kind which the index applies to.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataStoreIndexPropertyArgs']]]] properties: An ordered list of properties to index on.
               Structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataStoreIndexArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Describes a composite index for Cloud Datastore.

        To get more information about Index, see:

        * [API documentation](https://cloud.google.com/datastore/docs/reference/admin/rest/v1/projects.indexes)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/datastore/docs/concepts/indexes)

        > **Warning:** This resource creates a Datastore Index on a project that has already
        enabled a Datastore-compatible database. If you haven't already enabled
        one, you can create a `appengine.Application` resource with
        `database_type` set to `"CLOUD_DATASTORE_COMPATIBILITY"` to do so. Your
        Datastore location will be the same as the App Engine location specified.

        ## Example Usage
        ### Datastore Index

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.datastore.DataStoreIndex("default",
            kind="foo",
            properties=[
                gcp.datastore.DataStoreIndexPropertyArgs(
                    direction="ASCENDING",
                    name="property_a",
                ),
                gcp.datastore.DataStoreIndexPropertyArgs(
                    direction="ASCENDING",
                    name="property_b",
                ),
            ])
        ```

        ## Import

        Index can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:datastore/dataStoreIndex:DataStoreIndex default projects/{{project}}/indexes/{{index_id}}
        ```

        ```sh
         $ pulumi import gcp:datastore/dataStoreIndex:DataStoreIndex default {{project}}/{{index_id}}
        ```

        ```sh
         $ pulumi import gcp:datastore/dataStoreIndex:DataStoreIndex default {{index_id}}
        ```

        :param str resource_name: The name of the resource.
        :param DataStoreIndexArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataStoreIndexArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ancestor: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 properties: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataStoreIndexPropertyArgs']]]]] = None,
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

            __props__['ancestor'] = ancestor
            if kind is None and not opts.urn:
                raise TypeError("Missing required property 'kind'")
            __props__['kind'] = kind
            __props__['project'] = project
            __props__['properties'] = properties
            __props__['index_id'] = None
        super(DataStoreIndex, __self__).__init__(
            'gcp:datastore/dataStoreIndex:DataStoreIndex',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ancestor: Optional[pulumi.Input[str]] = None,
            index_id: Optional[pulumi.Input[str]] = None,
            kind: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            properties: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataStoreIndexPropertyArgs']]]]] = None) -> 'DataStoreIndex':
        """
        Get an existing DataStoreIndex resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ancestor: Policy for including ancestors in the index.
               Default value is `NONE`.
               Possible values are `NONE` and `ALL_ANCESTORS`.
        :param pulumi.Input[str] index_id: The index id.
        :param pulumi.Input[str] kind: The entity kind which the index applies to.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataStoreIndexPropertyArgs']]]] properties: An ordered list of properties to index on.
               Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["ancestor"] = ancestor
        __props__["index_id"] = index_id
        __props__["kind"] = kind
        __props__["project"] = project
        __props__["properties"] = properties
        return DataStoreIndex(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def ancestor(self) -> pulumi.Output[Optional[str]]:
        """
        Policy for including ancestors in the index.
        Default value is `NONE`.
        Possible values are `NONE` and `ALL_ANCESTORS`.
        """
        return pulumi.get(self, "ancestor")

    @property
    @pulumi.getter(name="indexId")
    def index_id(self) -> pulumi.Output[str]:
        """
        The index id.
        """
        return pulumi.get(self, "index_id")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        The entity kind which the index applies to.
        """
        return pulumi.get(self, "kind")

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
    def properties(self) -> pulumi.Output[Optional[Sequence['outputs.DataStoreIndexProperty']]]:
        """
        An ordered list of properties to index on.
        Structure is documented below.
        """
        return pulumi.get(self, "properties")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

