# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Deployment(pulumi.CustomResource):
    create_policy: pulumi.Output[str]
    """
    Set the policy to use for creating new resources. Only used on
    create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
    `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
    the deployment will fail. Note that updating this field does not
    actually affect the deployment, just how it is updated.
    """
    delete_policy: pulumi.Output[str]
    """
    Set the policy to use for deleting new resources on update/delete.
    Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
    resource is deleted after removal from Deployment Manager. If
    `ABANDON`, the resource is only removed from Deployment Manager
    and is not actually deleted. Note that updating this field does not
    actually change the deployment, just how it is updated.
    """
    deployment_id: pulumi.Output[str]
    """
    Unique identifier for deployment. Output only.
    """
    description: pulumi.Output[str]
    """
    Optional user-provided description of deployment.
    """
    labels: pulumi.Output[list]
    """
    Key-value pairs to apply to this labels.  Structure is documented below.

      * `key` (`str`) - Key for label.
      * `value` (`str`) - Value of label.
    """
    manifest: pulumi.Output[str]
    """
    Output only. URL of the manifest representing the last manifest that was successfully deployed.
    """
    name: pulumi.Output[str]
    """
    The name of the template to import, as declared in the YAML
    configuration.
    """
    preview: pulumi.Output[bool]
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
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    self_link: pulumi.Output[str]
    """
    Output only. Server defined URL for the resource.
    """
    target: pulumi.Output[dict]
    """
    Parameters that define your deployment, including the deployment
    configuration and relevant templates.  Structure is documented below.

      * `config` (`dict`) - The root configuration file to use for this deployment.  Structure is documented below.
        * `content` (`str`) - The full contents of the template that you want to import.

      * `imports` (`list`) - Specifies import files for this configuration. This can be
        used to import templates or other files. For example, you might
        import a text file in order to use the file in a template.  Structure is documented below.
        * `content` (`str`) - The full contents of the template that you want to import.
        * `name` (`str`) - The name of the template to import, as declared in the YAML
          configuration.
    """
    def __init__(__self__, resource_name, opts=None, create_policy=None, delete_policy=None, description=None, labels=None, name=None, preview=None, project=None, target=None, __props__=None, __name__=None, __opts__=None):
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
        :param pulumi.Input[list] labels: Key-value pairs to apply to this labels.  Structure is documented below.
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
        :param pulumi.Input[dict] target: Parameters that define your deployment, including the deployment
               configuration and relevant templates.  Structure is documented below.

        The **labels** object supports the following:

          * `key` (`pulumi.Input[str]`) - Key for label.
          * `value` (`pulumi.Input[str]`) - Value of label.

        The **target** object supports the following:

          * `config` (`pulumi.Input[dict]`) - The root configuration file to use for this deployment.  Structure is documented below.
            * `content` (`pulumi.Input[str]`) - The full contents of the template that you want to import.

          * `imports` (`pulumi.Input[list]`) - Specifies import files for this configuration. This can be
            used to import templates or other files. For example, you might
            import a text file in order to use the file in a template.  Structure is documented below.
            * `content` (`pulumi.Input[str]`) - The full contents of the template that you want to import.
            * `name` (`pulumi.Input[str]`) - The name of the template to import, as declared in the YAML
              configuration.
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
    def get(resource_name, id, opts=None, create_policy=None, delete_policy=None, deployment_id=None, description=None, labels=None, manifest=None, name=None, preview=None, project=None, self_link=None, target=None):
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
        :param pulumi.Input[list] labels: Key-value pairs to apply to this labels.  Structure is documented below.
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
        :param pulumi.Input[dict] target: Parameters that define your deployment, including the deployment
               configuration and relevant templates.  Structure is documented below.

        The **labels** object supports the following:

          * `key` (`pulumi.Input[str]`) - Key for label.
          * `value` (`pulumi.Input[str]`) - Value of label.

        The **target** object supports the following:

          * `config` (`pulumi.Input[dict]`) - The root configuration file to use for this deployment.  Structure is documented below.
            * `content` (`pulumi.Input[str]`) - The full contents of the template that you want to import.

          * `imports` (`pulumi.Input[list]`) - Specifies import files for this configuration. This can be
            used to import templates or other files. For example, you might
            import a text file in order to use the file in a template.  Structure is documented below.
            * `content` (`pulumi.Input[str]`) - The full contents of the template that you want to import.
            * `name` (`pulumi.Input[str]`) - The name of the template to import, as declared in the YAML
              configuration.
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
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

