# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Registry(pulumi.CustomResource):
    bucket_self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    location: pulumi.Output[str]
    """
    The location of the registry. One of `ASIA`, `EU`, `US` or not specified. See [the official documentation](https://cloud.google.com/container-registry/docs/pushing-and-pulling#pushing_an_image_to_a_registry) for more information on registry locations.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
    """
    def __init__(__self__, resource_name, opts=None, location=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        Ensures that the Google Cloud Storage bucket that backs Google Container Registry exists. Creating this resource will create the backing bucket if it does not exist, or do nothing if the bucket already exists. Destroying this resource does *NOT* destroy the backing bucket. For more information see [the official documentation](https://cloud.google.com/container-registry/docs/overview)

        This resource can be used to ensure that the GCS bucket exists prior to assigning permissions. For more information see the [access control page](https://cloud.google.com/container-registry/docs/access-control) for GCR.


        ## Example Usage



        ```python
        import pulumi
        import pulumi_gcp as gcp

        registry = gcp.container.Registry("registry",
            location="EU",
            project="my-project")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location of the registry. One of `ASIA`, `EU`, `US` or not specified. See [the official documentation](https://cloud.google.com/container-registry/docs/pushing-and-pulling#pushing_an_image_to_a_registry) for more information on registry locations.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
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

            __props__['location'] = location
            __props__['project'] = project
            __props__['bucket_self_link'] = None
        super(Registry, __self__).__init__(
            'gcp:container/registry:Registry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, bucket_self_link=None, location=None, project=None):
        """
        Get an existing Registry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket_self_link: The URI of the created resource.
        :param pulumi.Input[str] location: The location of the registry. One of `ASIA`, `EU`, `US` or not specified. See [the official documentation](https://cloud.google.com/container-registry/docs/pushing-and-pulling#pushing_an_image_to_a_registry) for more information on registry locations.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bucket_self_link"] = bucket_self_link
        __props__["location"] = location
        __props__["project"] = project
        return Registry(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
