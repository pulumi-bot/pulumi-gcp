# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GameServerDeployment(pulumi.CustomResource):
    deployment_id: pulumi.Output[str]
    """
    A unique id for the deployment.
    """
    description: pulumi.Output[str]
    """
    Human readable description of the game server deployment.
    """
    labels: pulumi.Output[dict]
    """
    The labels associated with this game server deployment. Each label is a
    key-value pair.
    """
    location: pulumi.Output[str]
    """
    Location of the Deployment.
    """
    name: pulumi.Output[str]
    """
    The resource id of the game server deployment, eg:
    'projects/{project_id}/locations/{location}/gameServerDeployments/{deployment_id}'. For example,
    'projects/my-project/locations/{location}/gameServerDeployments/my-deployment'.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    def __init__(__self__, resource_name, opts=None, deployment_id=None, description=None, labels=None, location=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        A game server deployment resource.

        To get more information about GameServerDeployment, see:

        * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.gameServerDeployments)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/game-servers/docs)

        ## Example Usage

        ### Game Service Deployment Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.gameservices.GameServerDeployment("default",
            deployment_id="tf-test-deployment",
            description="a deployment description")
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] deployment_id: A unique id for the deployment.
        :param pulumi.Input[str] description: Human readable description of the game server deployment.
        :param pulumi.Input[dict] labels: The labels associated with this game server deployment. Each label is a
               key-value pair.
        :param pulumi.Input[str] location: Location of the Deployment.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if deployment_id is None:
                raise TypeError("Missing required property 'deployment_id'")
            __props__['deployment_id'] = deployment_id
            __props__['description'] = description
            __props__['labels'] = labels
            __props__['location'] = location
            __props__['project'] = project
            __props__['name'] = None
        super(GameServerDeployment, __self__).__init__(
            'gcp:gameservices/gameServerDeployment:GameServerDeployment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, deployment_id=None, description=None, labels=None, location=None, name=None, project=None):
        """
        Get an existing GameServerDeployment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] deployment_id: A unique id for the deployment.
        :param pulumi.Input[str] description: Human readable description of the game server deployment.
        :param pulumi.Input[dict] labels: The labels associated with this game server deployment. Each label is a
               key-value pair.
        :param pulumi.Input[str] location: Location of the Deployment.
        :param pulumi.Input[str] name: The resource id of the game server deployment, eg:
               'projects/{project_id}/locations/{location}/gameServerDeployments/{deployment_id}'. For example,
               'projects/my-project/locations/{location}/gameServerDeployments/my-deployment'.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["deployment_id"] = deployment_id
        __props__["description"] = description
        __props__["labels"] = labels
        __props__["location"] = location
        __props__["name"] = name
        __props__["project"] = project
        return GameServerDeployment(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

