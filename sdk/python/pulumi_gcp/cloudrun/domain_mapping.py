# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class DomainMapping(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The location of the cloud run instance. eg us-central1
    """
    metadata: pulumi.Output[dict]
    """
    Metadata associated with this DomainMapping.

      * `annotations` (`dict`)
      * `generation` (`float`)
      * `labels` (`dict`)
      * `namespace` (`str`)
      * `resourceVersion` (`str`)
      * `self_link` (`str`)
      * `uid` (`str`)
    """
    name: pulumi.Output[str]
    """
    Name should be a verified domain
    """
    project: pulumi.Output[str]
    spec: pulumi.Output[dict]
    """
    The spec for this DomainMapping.

      * `certificateMode` (`str`)
      * `forceOverride` (`bool`)
      * `routeName` (`str`)
    """
    status: pulumi.Output[dict]
    """
    The current status of the DomainMapping.

      * `conditions` (`list`)
        * `message` (`str`)
        * `reason` (`str`)
        * `status` (`str`)
        * `type` (`str`)

      * `mappedRouteName` (`str`)
      * `observedGeneration` (`float`)
      * `resourceRecords` (`list`)
        * `name` (`str`)
        * `rrdata` (`str`)
        * `type` (`str`)
    """
    def __init__(__self__, resource_name, opts=None, location=None, metadata=None, name=None, project=None, spec=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a DomainMapping resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location of the cloud run instance. eg us-central1
        :param pulumi.Input[dict] metadata: Metadata associated with this DomainMapping.
        :param pulumi.Input[str] name: Name should be a verified domain
        :param pulumi.Input[dict] spec: The spec for this DomainMapping.

        The **metadata** object supports the following:

          * `annotations` (`pulumi.Input[dict]`)
          * `generation` (`pulumi.Input[float]`)
          * `labels` (`pulumi.Input[dict]`)
          * `namespace` (`pulumi.Input[str]`)
          * `resourceVersion` (`pulumi.Input[str]`)
          * `self_link` (`pulumi.Input[str]`)
          * `uid` (`pulumi.Input[str]`)

        The **spec** object supports the following:

          * `certificateMode` (`pulumi.Input[str]`)
          * `forceOverride` (`pulumi.Input[bool]`)
          * `routeName` (`pulumi.Input[str]`)
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

            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if metadata is None:
                raise TypeError("Missing required property 'metadata'")
            __props__['metadata'] = metadata
            __props__['name'] = name
            __props__['project'] = project
            if spec is None:
                raise TypeError("Missing required property 'spec'")
            __props__['spec'] = spec
            __props__['status'] = None
        super(DomainMapping, __self__).__init__(
            'gcp:cloudrun/domainMapping:DomainMapping',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, location=None, metadata=None, name=None, project=None, spec=None, status=None):
        """
        Get an existing DomainMapping resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location of the cloud run instance. eg us-central1
        :param pulumi.Input[dict] metadata: Metadata associated with this DomainMapping.
        :param pulumi.Input[str] name: Name should be a verified domain
        :param pulumi.Input[dict] spec: The spec for this DomainMapping.
        :param pulumi.Input[dict] status: The current status of the DomainMapping.

        The **metadata** object supports the following:

          * `annotations` (`pulumi.Input[dict]`)
          * `generation` (`pulumi.Input[float]`)
          * `labels` (`pulumi.Input[dict]`)
          * `namespace` (`pulumi.Input[str]`)
          * `resourceVersion` (`pulumi.Input[str]`)
          * `self_link` (`pulumi.Input[str]`)
          * `uid` (`pulumi.Input[str]`)

        The **spec** object supports the following:

          * `certificateMode` (`pulumi.Input[str]`)
          * `forceOverride` (`pulumi.Input[bool]`)
          * `routeName` (`pulumi.Input[str]`)

        The **status** object supports the following:

          * `conditions` (`pulumi.Input[list]`)
            * `message` (`pulumi.Input[str]`)
            * `reason` (`pulumi.Input[str]`)
            * `status` (`pulumi.Input[str]`)
            * `type` (`pulumi.Input[str]`)

          * `mappedRouteName` (`pulumi.Input[str]`)
          * `observedGeneration` (`pulumi.Input[float]`)
          * `resourceRecords` (`pulumi.Input[list]`)
            * `name` (`pulumi.Input[str]`)
            * `rrdata` (`pulumi.Input[str]`)
            * `type` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["location"] = location
        __props__["metadata"] = metadata
        __props__["name"] = name
        __props__["project"] = project
        __props__["spec"] = spec
        __props__["status"] = status
        return DomainMapping(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

