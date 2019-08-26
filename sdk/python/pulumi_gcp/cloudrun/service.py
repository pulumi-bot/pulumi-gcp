# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Service(pulumi.CustomResource):
    location: pulumi.Output[str]
    metadata: pulumi.Output[dict]
    name: pulumi.Output[str]
    project: pulumi.Output[str]
    spec: pulumi.Output[dict]
    status: pulumi.Output[dict]
    def __init__(__self__, resource_name, opts=None, location=None, metadata=None, name=None, project=None, spec=None, __props__=None, __name__=None, __opts__=None):
        """
        **Note:** Cloud Run as a product is in beta, however the REST API is currently still an alpha.
        Please use this with caution as it may change when the API moves to beta.
        
        Service acts as a top-level container that manages a set of Routes and
        Configurations which implement a network service. Service exists to provide a
        singular abstraction which can be access controlled, reasoned about, and
        which encapsulates software lifecycle decisions such as rollout policy and
        team resource ownership. Service acts only as an orchestrator of the
        underlying Routes and Configurations (much as a kubernetes Deployment
        orchestrates ReplicaSets).
        
        The Service's controller will track the statuses of its owned Configuration
        and Route, reflecting their statuses and conditions as its own.
        
        See also:
        https://github.com/knative/serving/blob/master/docs/spec/overview.md#service
        
        To get more information about Service, see:
        
        * [API documentation](https://cloud.google.com/run/docs/reference/rest/v1alpha1/projects.locations.services)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/run/docs/)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        
        The **metadata** object supports the following:
        
          * `annotations` (`pulumi.Input[dict]`)
          * `generation` (`pulumi.Input[float]`)
          * `labels` (`pulumi.Input[dict]`)
          * `namespace` (`pulumi.Input[str]`)
          * `resource_version` (`pulumi.Input[str]`)
          * `self_link` (`pulumi.Input[str]`)
          * `uid` (`pulumi.Input[str]`)
        
        The **spec** object supports the following:
        
          * `container_concurrency` (`pulumi.Input[float]`)
          * `containers` (`pulumi.Input[list]`)
        
            * `args` (`pulumi.Input[list]`)
            * `commands` (`pulumi.Input[list]`)
            * `envs` (`pulumi.Input[list]`)
        
              * `name` (`pulumi.Input[str]`)
              * `value` (`pulumi.Input[str]`)
        
            * `env_froms` (`pulumi.Input[list]`)
        
              * `config_map_ref` (`pulumi.Input[dict]`)
        
                * `local_object_reference` (`pulumi.Input[dict]`)
        
                  * `name` (`pulumi.Input[str]`)
        
                * `optional` (`pulumi.Input[bool]`)
        
              * `prefix` (`pulumi.Input[str]`)
              * `secret_ref` (`pulumi.Input[dict]`)
        
                * `local_object_reference` (`pulumi.Input[dict]`)
        
                  * `name` (`pulumi.Input[str]`)
        
                * `optional` (`pulumi.Input[bool]`)
        
            * `image` (`pulumi.Input[str]`)
            * `resources` (`pulumi.Input[dict]`)
        
              * `limits` (`pulumi.Input[dict]`)
              * `requests` (`pulumi.Input[dict]`)
        
            * `working_dir` (`pulumi.Input[str]`)
        
          * `serving_state` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloud_run_service.html.markdown.
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
        super(Service, __self__).__init__(
            'gcp:cloudrun/service:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, location=None, metadata=None, name=None, project=None, spec=None, status=None):
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        
        The **metadata** object supports the following:
        
          * `annotations` (`pulumi.Input[dict]`)
          * `generation` (`pulumi.Input[float]`)
          * `labels` (`pulumi.Input[dict]`)
          * `namespace` (`pulumi.Input[str]`)
          * `resource_version` (`pulumi.Input[str]`)
          * `self_link` (`pulumi.Input[str]`)
          * `uid` (`pulumi.Input[str]`)
        
        The **spec** object supports the following:
        
          * `container_concurrency` (`pulumi.Input[float]`)
          * `containers` (`pulumi.Input[list]`)
        
            * `args` (`pulumi.Input[list]`)
            * `commands` (`pulumi.Input[list]`)
            * `envs` (`pulumi.Input[list]`)
        
              * `name` (`pulumi.Input[str]`)
              * `value` (`pulumi.Input[str]`)
        
            * `env_froms` (`pulumi.Input[list]`)
        
              * `config_map_ref` (`pulumi.Input[dict]`)
        
                * `local_object_reference` (`pulumi.Input[dict]`)
        
                  * `name` (`pulumi.Input[str]`)
        
                * `optional` (`pulumi.Input[bool]`)
        
              * `prefix` (`pulumi.Input[str]`)
              * `secret_ref` (`pulumi.Input[dict]`)
        
                * `local_object_reference` (`pulumi.Input[dict]`)
        
                  * `name` (`pulumi.Input[str]`)
        
                * `optional` (`pulumi.Input[bool]`)
        
            * `image` (`pulumi.Input[str]`)
            * `resources` (`pulumi.Input[dict]`)
        
              * `limits` (`pulumi.Input[dict]`)
              * `requests` (`pulumi.Input[dict]`)
        
            * `working_dir` (`pulumi.Input[str]`)
        
          * `serving_state` (`pulumi.Input[str]`)
        
        The **status** object supports the following:
        
          * `conditions` (`pulumi.Input[list]`)
        
            * `message` (`pulumi.Input[str]`)
            * `reason` (`pulumi.Input[str]`)
            * `status` (`pulumi.Input[str]`)
            * `type` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloud_run_service.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["location"] = location
        __props__["metadata"] = metadata
        __props__["name"] = name
        __props__["project"] = project
        __props__["spec"] = spec
        __props__["status"] = status
        return Service(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

