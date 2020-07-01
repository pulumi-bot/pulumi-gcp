# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class IAMCustomRole(pulumi.CustomResource):
    deleted: pulumi.Output[bool]
    """
    (Optional) The current deleted state of the role.
    """
    description: pulumi.Output[str]
    """
    A human-readable description for the role.
    """
    name: pulumi.Output[str]
    """
    The name of the role in the format `projects/{{project}}/roles/{{role_id}}`. Like `id`, this field can be used as a reference in other resources such as IAM role bindings.
    """
    permissions: pulumi.Output[list]
    """
    The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
    """
    project: pulumi.Output[str]
    """
    The project that the service account will be created in.
    Defaults to the provider project configuration.
    """
    role_id: pulumi.Output[str]
    """
    The camel case role id to use for this role. Cannot contain `-` characters.
    """
    stage: pulumi.Output[str]
    """
    The current launch stage of the role.
    Defaults to `GA`.
    List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
    """
    title: pulumi.Output[str]
    """
    A human-readable title for the role.
    """
    def __init__(__self__, resource_name, opts=None, description=None, permissions=None, project=None, role_id=None, stage=None, title=None, __props__=None, __name__=None, __opts__=None):
        """
        Allows management of a customized Cloud IAM project role. For more information see
        [the official documentation](https://cloud.google.com/iam/docs/understanding-custom-roles)
        and
        [API](https://cloud.google.com/iam/reference/rest/v1/projects.roles).

        > **Warning:** Note that custom roles in GCP have the concept of a soft-delete. There are two issues that may arise
         from this and how roles are propagated. 1) creating a role may involve undeleting and then updating a role with the
         same name, possibly causing confusing behavior between undelete and update. 2) A deleted role is permanently deleted
         after 7 days, but it can take up to 30 more days (i.e. between 7 and 37 days after deletion) before the role name is
         made available again. This means a deleted role that has been deleted for more than 7 days cannot be changed at all
         by the provider, and new roles cannot share that name.

        ## Example Usage

        This snippet creates a customized IAM role.

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_custom_role = gcp.projects.IAMCustomRole("my-custom-role",
            description="A description",
            permissions=[
                "iam.roles.list",
                "iam.roles.create",
                "iam.roles.delete",
            ],
            role_id="myCustomRole",
            title="My Custom Role")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A human-readable description for the role.
        :param pulumi.Input[list] permissions: The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
        :param pulumi.Input[str] project: The project that the service account will be created in.
               Defaults to the provider project configuration.
        :param pulumi.Input[str] role_id: The camel case role id to use for this role. Cannot contain `-` characters.
        :param pulumi.Input[str] stage: The current launch stage of the role.
               Defaults to `GA`.
               List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
        :param pulumi.Input[str] title: A human-readable title for the role.
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

            __props__['description'] = description
            if permissions is None:
                raise TypeError("Missing required property 'permissions'")
            __props__['permissions'] = permissions
            __props__['project'] = project
            if role_id is None:
                raise TypeError("Missing required property 'role_id'")
            __props__['role_id'] = role_id
            __props__['stage'] = stage
            if title is None:
                raise TypeError("Missing required property 'title'")
            __props__['title'] = title
            __props__['deleted'] = None
            __props__['name'] = None
        super(IAMCustomRole, __self__).__init__(
            'gcp:projects/iAMCustomRole:IAMCustomRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, deleted=None, description=None, name=None, permissions=None, project=None, role_id=None, stage=None, title=None):
        """
        Get an existing IAMCustomRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] deleted: (Optional) The current deleted state of the role.
        :param pulumi.Input[str] description: A human-readable description for the role.
        :param pulumi.Input[str] name: The name of the role in the format `projects/{{project}}/roles/{{role_id}}`. Like `id`, this field can be used as a reference in other resources such as IAM role bindings.
        :param pulumi.Input[list] permissions: The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
        :param pulumi.Input[str] project: The project that the service account will be created in.
               Defaults to the provider project configuration.
        :param pulumi.Input[str] role_id: The camel case role id to use for this role. Cannot contain `-` characters.
        :param pulumi.Input[str] stage: The current launch stage of the role.
               Defaults to `GA`.
               List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
        :param pulumi.Input[str] title: A human-readable title for the role.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["deleted"] = deleted
        __props__["description"] = description
        __props__["name"] = name
        __props__["permissions"] = permissions
        __props__["project"] = project
        __props__["role_id"] = role_id
        __props__["stage"] = stage
        __props__["title"] = title
        return IAMCustomRole(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
