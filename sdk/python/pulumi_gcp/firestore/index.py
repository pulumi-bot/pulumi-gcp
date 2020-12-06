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

__all__ = ['Index']


class Index(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 collection: Optional[pulumi.Input[str]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IndexFieldArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 query_scope: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Cloud Firestore indexes enable simple and complex queries against documents in a database.
         This resource manages composite indexes and not single
        field indexes.

        To get more information about Index, see:

        * [API documentation](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.collectionGroups.indexes)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/firestore/docs/query-data/indexing)

        > **Warning:** This resource creates a Firestore Index on a project that already has
        Firestore enabled. If you haven't already enabled it, you can create a
        `appengine.Application` resource with `database_type` set to
        `"CLOUD_FIRESTORE"` to do so. Your Firestore location will be the same as
        the App Engine location specified.

        ## Example Usage
        ### Firestore Index Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_index = gcp.firestore.Index("my-index",
            collection="chatrooms",
            fields=[
                gcp.firestore.IndexFieldArgs(
                    field_path="name",
                    order="ASCENDING",
                ),
                gcp.firestore.IndexFieldArgs(
                    field_path="description",
                    order="DESCENDING",
                ),
                gcp.firestore.IndexFieldArgs(
                    field_path="__name__",
                    order="DESCENDING",
                ),
            ],
            project="my-project-name")
        ```

        ## Import

        Index can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:firestore/index:Index default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] collection: The collection being indexed.
        :param pulumi.Input[str] database: The Firestore database id. Defaults to `"(default)"`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IndexFieldArgs']]]] fields: The fields supported by this index. The last field entry is always for
               the field path `__name__`. If, on creation, `__name__` was not
               specified as the last field, it will be added automatically with the
               same direction as that of the last field defined. If the final field
               in a composite index is not directional, the `__name__` will be
               ordered `"ASCENDING"` (unless explicitly specified otherwise).
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] query_scope: The scope at which a query is run.
               Default value is `COLLECTION`.
               Possible values are `COLLECTION` and `COLLECTION_GROUP`.
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

            if collection is None:
                raise TypeError("Missing required property 'collection'")
            __props__['collection'] = collection
            __props__['database'] = database
            if fields is None:
                raise TypeError("Missing required property 'fields'")
            __props__['fields'] = fields
            __props__['project'] = project
            __props__['query_scope'] = query_scope
            __props__['name'] = None
        super(Index, __self__).__init__(
            'gcp:firestore/index:Index',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            collection: Optional[pulumi.Input[str]] = None,
            database: Optional[pulumi.Input[str]] = None,
            fields: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IndexFieldArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            query_scope: Optional[pulumi.Input[str]] = None) -> 'Index':
        """
        Get an existing Index resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] collection: The collection being indexed.
        :param pulumi.Input[str] database: The Firestore database id. Defaults to `"(default)"`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IndexFieldArgs']]]] fields: The fields supported by this index. The last field entry is always for
               the field path `__name__`. If, on creation, `__name__` was not
               specified as the last field, it will be added automatically with the
               same direction as that of the last field defined. If the final field
               in a composite index is not directional, the `__name__` will be
               ordered `"ASCENDING"` (unless explicitly specified otherwise).
               Structure is documented below.
        :param pulumi.Input[str] name: A server defined name for this index. Format:
               'projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}'
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] query_scope: The scope at which a query is run.
               Default value is `COLLECTION`.
               Possible values are `COLLECTION` and `COLLECTION_GROUP`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["collection"] = collection
        __props__["database"] = database
        __props__["fields"] = fields
        __props__["name"] = name
        __props__["project"] = project
        __props__["query_scope"] = query_scope
        return Index(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def collection(self) -> pulumi.Output[str]:
        """
        The collection being indexed.
        """
        return pulumi.get(self, "collection")

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[Optional[str]]:
        """
        The Firestore database id. Defaults to `"(default)"`.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Output[Sequence['outputs.IndexField']]:
        """
        The fields supported by this index. The last field entry is always for
        the field path `__name__`. If, on creation, `__name__` was not
        specified as the last field, it will be added automatically with the
        same direction as that of the last field defined. If the final field
        in a composite index is not directional, the `__name__` will be
        ordered `"ASCENDING"` (unless explicitly specified otherwise).
        Structure is documented below.
        """
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A server defined name for this index. Format:
        'projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}'
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="queryScope")
    def query_scope(self) -> pulumi.Output[Optional[str]]:
        """
        The scope at which a query is run.
        Default value is `COLLECTION`.
        Possible values are `COLLECTION` and `COLLECTION_GROUP`.
        """
        return pulumi.get(self, "query_scope")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

