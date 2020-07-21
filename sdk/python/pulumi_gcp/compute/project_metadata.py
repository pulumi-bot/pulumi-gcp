# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from ._inputs import *
from . import outputs


class ProjectMetadata(pulumi.CustomResource):
    metadata: pulumi.Output[Dict[str, str]] = pulumi.output_property("metadata")
    """
    A series of key value pairs.
    """
    project: pulumi.Output[str] = pulumi.output_property("project")
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, metadata=None, project=None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Authoritatively manages metadata common to all instances for a project in GCE. For more information see
        [the official documentation](https://cloud.google.com/compute/docs/storing-retrieving-metadata)
        and
        [API](https://cloud.google.com/compute/docs/reference/latest/projects/setCommonInstanceMetadata).

        > **Note:**  This resource manages all project-level metadata including project-level ssh keys.
        Keys unset in config but set on the server will be removed. If you want to manage only single
        key/value pairs within the project metadata rather than the entire set, then use
        google_compute_project_metadata_item.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.compute.ProjectMetadata("default", metadata={
            "13": "42",
            "fizz": "buzz",
            "foo": "bar",
        })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] metadata: A series of key value pairs.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
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

            if metadata is None:
                raise TypeError("Missing required property 'metadata'")
            __props__['metadata'] = metadata
            __props__['project'] = project
        super(ProjectMetadata, __self__).__init__(
            'gcp:compute/projectMetadata:ProjectMetadata',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, metadata=None, project=None):
        """
        Get an existing ProjectMetadata resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] metadata: A series of key value pairs.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["metadata"] = metadata
        __props__["project"] = project
        return ProjectMetadata(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

