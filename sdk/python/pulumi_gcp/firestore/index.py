# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Index(pulumi.CustomResource):
    collection: pulumi.Output[str]
    """
    The collection being indexed.
    """
    database: pulumi.Output[str]
    """
    The Firestore database id. Defaults to `"(default)"`.
    """
    fields: pulumi.Output[list]
    """
    The fields supported by this index. The last field entry is always for
    the field path `__name__`. If, on creation, `__name__` was not
    specified as the last field, it will be added automatically with the
    same direction as that of the last field defined. If the final field
    in a composite index is not directional, the `__name__` will be
    ordered `"ASCENDING"` (unless explicitly specified otherwise).  Structure is documented below.

      * `arrayConfig` (`str`) - Indicates that this field supports operations on arrayValues. Only one of `order` and `arrayConfig` can
        be specified.
      * `fieldPath` (`str`) - Name of the field.
      * `order` (`str`) - Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
        Only one of `order` and `arrayConfig` can be specified.
    """
    name: pulumi.Output[str]
    """
    A server defined name for this index. Format:
    'projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}'
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    query_scope: pulumi.Output[str]
    """
    The scope at which a query is run. One of `"COLLECTION"` or
    `"COLLECTION_GROUP"`. Defaults to `"COLLECTION"`.
    """
    def __init__(__self__, resource_name, opts=None, collection=None, database=None, fields=None, project=None, query_scope=None, __props__=None, __name__=None, __opts__=None):
        """
        Cloud Firestore indexes enable simple and complex queries against documents in a database.
         This resource manages composite indexes and not single
        field indexes.


        To get more information about Index, see:

        * [API documentation](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.collectionGroups.indexes)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/firestore/docs/query-data/indexing)

        ## Example Usage - Firestore Index Basic


        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_index = gcp.firestore.Index("my-index",
            collection="chatrooms",
            fields=[
                {
                    "fieldPath": "name",
                    "order": "ASCENDING",
                },
                {
                    "fieldPath": "description",
                    "order": "DESCENDING",
                },
                {
                    "fieldPath": "__name__",
                    "order": "DESCENDING",
                },
            ],
            project="my-project-name")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] collection: The collection being indexed.
        :param pulumi.Input[str] database: The Firestore database id. Defaults to `"(default)"`.
        :param pulumi.Input[list] fields: The fields supported by this index. The last field entry is always for
               the field path `__name__`. If, on creation, `__name__` was not
               specified as the last field, it will be added automatically with the
               same direction as that of the last field defined. If the final field
               in a composite index is not directional, the `__name__` will be
               ordered `"ASCENDING"` (unless explicitly specified otherwise).  Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] query_scope: The scope at which a query is run. One of `"COLLECTION"` or
               `"COLLECTION_GROUP"`. Defaults to `"COLLECTION"`.

        The **fields** object supports the following:

          * `arrayConfig` (`pulumi.Input[str]`) - Indicates that this field supports operations on arrayValues. Only one of `order` and `arrayConfig` can
            be specified.
          * `fieldPath` (`pulumi.Input[str]`) - Name of the field.
          * `order` (`pulumi.Input[str]`) - Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
            Only one of `order` and `arrayConfig` can be specified.
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, collection=None, database=None, fields=None, name=None, project=None, query_scope=None):
        """
        Get an existing Index resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] collection: The collection being indexed.
        :param pulumi.Input[str] database: The Firestore database id. Defaults to `"(default)"`.
        :param pulumi.Input[list] fields: The fields supported by this index. The last field entry is always for
               the field path `__name__`. If, on creation, `__name__` was not
               specified as the last field, it will be added automatically with the
               same direction as that of the last field defined. If the final field
               in a composite index is not directional, the `__name__` will be
               ordered `"ASCENDING"` (unless explicitly specified otherwise).  Structure is documented below.
        :param pulumi.Input[str] name: A server defined name for this index. Format:
               'projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}'
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] query_scope: The scope at which a query is run. One of `"COLLECTION"` or
               `"COLLECTION_GROUP"`. Defaults to `"COLLECTION"`.

        The **fields** object supports the following:

          * `arrayConfig` (`pulumi.Input[str]`) - Indicates that this field supports operations on arrayValues. Only one of `order` and `arrayConfig` can
            be specified.
          * `fieldPath` (`pulumi.Input[str]`) - Name of the field.
          * `order` (`pulumi.Input[str]`) - Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
            Only one of `order` and `arrayConfig` can be specified.
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

