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

__all__ = ['Deployment']


class Deployment(pulumi.CustomResource):
    create_policy: pulumi.Output[Optional[str]] = pulumi.property("createPolicy")
    """
    Set the policy to use for creating new resources. Only used on
    create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
    `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
    the deployment will fail. Note that updating this field does not
    actually affect the deployment, just how it is updated.
    """

    delete_policy: pulumi.Output[Optional[str]] = pulumi.property("deletePolicy")
    """
    Set the policy to use for deleting new resources on update/delete.
    Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
    resource is deleted after removal from Deployment Manager. If
    `ABANDON`, the resource is only removed from Deployment Manager
    and is not actually deleted. Note that updating this field does not
    actually change the deployment, just how it is updated.
    """

    deployment_id: pulumi.Output[str] = pulumi.property("deploymentId")
    """
    Unique identifier for deployment. Output only.
    """

    description: pulumi.Output[Optional[str]] = pulumi.property("description")
    """
    Optional user-provided description of deployment.
    """

    labels: pulumi.Output[Optional[List['outputs.DeploymentLabel']]] = pulumi.property("labels")
    """
    Key-value pairs to apply to this labels.  Structure is documented below.
    """

    manifest: pulumi.Output[str] = pulumi.property("manifest")
    """
    Output only. URL of the manifest representing the last manifest that was successfully deployed.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    The name of the template to import, as declared in the YAML
    configuration.
    """

    preview: pulumi.Output[Optional[bool]] = pulumi.property("preview")
    """
    If set to true, a deployment is created with "shell" resources
    that are not actually instantiated. This allows you to preview a
    deployment. It can be updated to false to actually deploy
    with real resources.
    ~>**NOTE**: Deployment Manager does not allow update
    of a deployment in preview (unless updating to preview=false). Thus,
    the provider will force-recreate deployments if either preview is updated
    to true or if other fields are updated while preview is true.
    """

    project: pulumi.Output[str] = pulumi.property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """

    self_link: pulumi.Output[str] = pulumi.property("selfLink")
    """
    Output only. Server defined URL for the resource.
    """

    target: pulumi.Output['outputs.DeploymentTarget'] = pulumi.property("target")
    """
    Parameters that define your deployment, including the deployment
    configuration and relevant templates.  Structure is documented below.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_policy: Optional[pulumi.Input[str]] = None,
                 delete_policy: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['DeploymentLabelArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 preview: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[pulumi.InputType['DeploymentTargetArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A collection of resources that are deployed and managed together using
        a configuration file

        > **Warning:** This resource is intended only to manage a Deployment resource,
        and attempts to manage the Deployment's resources in the provider as well
        will likely result in errors or unexpected behavior as the two tools
        fight over ownership. We strongly discourage doing so unless you are an
        experienced user of both tools.

        In addition, due to limitations of the API, the provider will treat
        deployments in preview as recreate-only for any update operation other
        than actually deploying an in-preview deployment (i.e. `preview=true` to
        `preview=false`).

        ## Example Usage
        ### Deployment Manager Deployment Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        deployment = gcp.deploymentmanager.Deployment("deployment",
            target={
                "config": {
                    "content": (lambda path: open(path).read())("path/to/config.yml"),
                },
            },
            labels=[{
                "key": "foo",
                "value": "bar",
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_policy: Set the policy to use for creating new resources. Only used on
               create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
               `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
               the deployment will fail. Note that updating this field does not
               actually affect the deployment, just how it is updated.
        :param pulumi.Input[str] delete_policy: Set the policy to use for deleting new resources on update/delete.
               Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
               resource is deleted after removal from Deployment Manager. If
               `ABANDON`, the resource is only removed from Deployment Manager
               and is not actually deleted. Note that updating this field does not
               actually change the deployment, just how it is updated.
        :param pulumi.Input[str] description: Optional user-provided description of deployment.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['DeploymentLabelArgs']]]] labels: Key-value pairs to apply to this labels.  Structure is documented below.
        :param pulumi.Input[str] name: The name of the template to import, as declared in the YAML
               configuration.
        :param pulumi.Input[bool] preview: If set to true, a deployment is created with "shell" resources
               that are not actually instantiated. This allows you to preview a
               deployment. It can be updated to false to actually deploy
               with real resources.
               ~>**NOTE**: Deployment Manager does not allow update
               of a deployment in preview (unless updating to preview=false). Thus,
               the provider will force-recreate deployments if either preview is updated
               to true or if other fields are updated while preview is true.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['DeploymentTargetArgs']] target: Parameters that define your deployment, including the deployment
               configuration and relevant templates.  Structure is documented below.
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

            __props__['create_policy'] = create_policy
            __props__['delete_policy'] = delete_policy
            __props__['description'] = description
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['preview'] = preview
            __props__['project'] = project
            if target is None:
                raise TypeError("Missing required property 'target'")
            __props__['target'] = target
            __props__['deployment_id'] = None
            __props__['manifest'] = None
            __props__['self_link'] = None
        super(Deployment, __self__).__init__(
            'gcp:deploymentmanager/deployment:Deployment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            create_policy: Optional[pulumi.Input[str]] = None,
            delete_policy: Optional[pulumi.Input[str]] = None,
            deployment_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['DeploymentLabelArgs']]]]] = None,
            manifest: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            preview: Optional[pulumi.Input[bool]] = None,
            project: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            target: Optional[pulumi.Input[pulumi.InputType['DeploymentTargetArgs']]] = None) -> 'Deployment':
        """
        Get an existing Deployment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_policy: Set the policy to use for creating new resources. Only used on
               create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
               `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
               the deployment will fail. Note that updating this field does not
               actually affect the deployment, just how it is updated.
        :param pulumi.Input[str] delete_policy: Set the policy to use for deleting new resources on update/delete.
               Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
               resource is deleted after removal from Deployment Manager. If
               `ABANDON`, the resource is only removed from Deployment Manager
               and is not actually deleted. Note that updating this field does not
               actually change the deployment, just how it is updated.
        :param pulumi.Input[str] deployment_id: Unique identifier for deployment. Output only.
        :param pulumi.Input[str] description: Optional user-provided description of deployment.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['DeploymentLabelArgs']]]] labels: Key-value pairs to apply to this labels.  Structure is documented below.
        :param pulumi.Input[str] manifest: Output only. URL of the manifest representing the last manifest that was successfully deployed.
        :param pulumi.Input[str] name: The name of the template to import, as declared in the YAML
               configuration.
        :param pulumi.Input[bool] preview: If set to true, a deployment is created with "shell" resources
               that are not actually instantiated. This allows you to preview a
               deployment. It can be updated to false to actually deploy
               with real resources.
               ~>**NOTE**: Deployment Manager does not allow update
               of a deployment in preview (unless updating to preview=false). Thus,
               the provider will force-recreate deployments if either preview is updated
               to true or if other fields are updated while preview is true.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: Output only. Server defined URL for the resource.
        :param pulumi.Input[pulumi.InputType['DeploymentTargetArgs']] target: Parameters that define your deployment, including the deployment
               configuration and relevant templates.  Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["create_policy"] = create_policy
        __props__["delete_policy"] = delete_policy
        __props__["deployment_id"] = deployment_id
        __props__["description"] = description
        __props__["labels"] = labels
        __props__["manifest"] = manifest
        __props__["name"] = name
        __props__["preview"] = preview
        __props__["project"] = project
        __props__["self_link"] = self_link
        __props__["target"] = target
        return Deployment(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

