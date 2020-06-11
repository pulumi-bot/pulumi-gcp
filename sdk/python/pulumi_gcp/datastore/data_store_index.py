# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class DataStoreIndex(pulumi.CustomResource):
    ancestor: pulumi.Output[str]
    """
    Policy for including ancestors in the index.
    """
    index_id: pulumi.Output[str]
    """
    The index id.
    """
    kind: pulumi.Output[str]
    """
    The entity kind which the index applies to.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    properties: pulumi.Output[list]
    """
    An ordered list of properties to index on.  Structure is documented below.

      * `direction` (`str`) - The direction the index should optimize for sorting.
      * `name` (`str`) - The property name to index.
    """
    def __init__(__self__, resource_name, opts=None, ancestor=None, kind=None, project=None, properties=None, __props__=None, __name__=None, __opts__=None):
        """
        Describes a composite index for Cloud Datastore.


        To get more information about Index, see:

        * [API documentation](https://cloud.google.com/datastore/docs/reference/admin/rest/v1/projects.indexes)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/datastore/docs/concepts/indexes)

        ## Example Usage - Datastore Index


        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.datastore.DataStoreIndex("default",
            kind="foo",
            properties=[
                {
                    "direction": "ASCENDING",
                    "name": "property_a",
                },
                {
                    "direction": "ASCENDING",
                    "name": "property_b",
                },
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ancestor: Policy for including ancestors in the index.
        :param pulumi.Input[str] kind: The entity kind which the index applies to.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[list] properties: An ordered list of properties to index on.  Structure is documented below.

        The **properties** object supports the following:

          * `direction` (`pulumi.Input[str]`) - The direction the index should optimize for sorting.
          * `name` (`pulumi.Input[str]`) - The property name to index.
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

            __props__['ancestor'] = ancestor
            if kind is None:
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
    def get(resource_name, id, opts=None, ancestor=None, index_id=None, kind=None, project=None, properties=None):
        """
        Get an existing DataStoreIndex resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ancestor: Policy for including ancestors in the index.
        :param pulumi.Input[str] index_id: The index id.
        :param pulumi.Input[str] kind: The entity kind which the index applies to.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[list] properties: An ordered list of properties to index on.  Structure is documented below.

        The **properties** object supports the following:

          * `direction` (`pulumi.Input[str]`) - The direction the index should optimize for sorting.
          * `name` (`pulumi.Input[str]`) - The property name to index.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["ancestor"] = ancestor
        __props__["index_id"] = index_id
        __props__["kind"] = kind
        __props__["project"] = project
        __props__["properties"] = properties
        return DataStoreIndex(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
