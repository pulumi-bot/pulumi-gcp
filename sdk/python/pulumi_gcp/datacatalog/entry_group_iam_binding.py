# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class EntryGroupIamBinding(pulumi.CustomResource):
    condition: pulumi.Output[dict]
    entry_group: pulumi.Output[str]
    """
    Used to find the parent resource to bind the IAM policy to
    """
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the IAM policy.
    """
    members: pulumi.Output[list]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
    """
    region: pulumi.Output[str]
    role: pulumi.Output[str]
    """
    The role that should be applied. Only one
    `datacatalog.EntryGroupIamBinding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    def __init__(__self__, resource_name, opts=None, condition=None, entry_group=None, members=None, project=None, region=None, role=None, __props__=None, __name__=None, __opts__=None):
        """
        Three different resources help you manage your IAM policy for Data catalog EntryGroup. Each of these resources serves a different use case:

        * `datacatalog.EntryGroupIamPolicy`: Authoritative. Sets the IAM policy for the entrygroup and replaces any existing policy already attached.
        * `datacatalog.EntryGroupIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the entrygroup are preserved.
        * `datacatalog.EntryGroupIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the entrygroup are preserved.

        > **Note:** `datacatalog.EntryGroupIamPolicy` **cannot** be used in conjunction with `datacatalog.EntryGroupIamBinding` and `datacatalog.EntryGroupIamMember` or they will fight over what your policy should be.

        > **Note:** `datacatalog.EntryGroupIamBinding` resources **can be** used in conjunction with `datacatalog.EntryGroupIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\_data\_catalog\_entry\_group\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(binding=[{
            "role": "roles/viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.datacatalog.EntryGroupIamPolicy("policy",
            entry_group=google_data_catalog_entry_group["basic_entry_group"]["name"],
            policy_data=admin.policy_data)
        ```

        ## google\_data\_catalog\_entry\_group\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.datacatalog.EntryGroupIamBinding("binding",
            entry_group=google_data_catalog_entry_group["basic_entry_group"]["name"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## google\_data\_catalog\_entry\_group\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.datacatalog.EntryGroupIamMember("member",
            entry_group=google_data_catalog_entry_group["basic_entry_group"]["name"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] entry_group: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `datacatalog.EntryGroupIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.

        The **condition** object supports the following:

          * `description` (`pulumi.Input[str]`)
          * `expression` (`pulumi.Input[str]`)
          * `title` (`pulumi.Input[str]`)
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

            __props__['condition'] = condition
            if entry_group is None:
                raise TypeError("Missing required property 'entry_group'")
            __props__['entry_group'] = entry_group
            if members is None:
                raise TypeError("Missing required property 'members'")
            __props__['members'] = members
            __props__['project'] = project
            __props__['region'] = region
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['etag'] = None
        super(EntryGroupIamBinding, __self__).__init__(
            'gcp:datacatalog/entryGroupIamBinding:EntryGroupIamBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, condition=None, entry_group=None, etag=None, members=None, project=None, region=None, role=None):
        """
        Get an existing EntryGroupIamBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] entry_group: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `datacatalog.EntryGroupIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.

        The **condition** object supports the following:

          * `description` (`pulumi.Input[str]`)
          * `expression` (`pulumi.Input[str]`)
          * `title` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["condition"] = condition
        __props__["entry_group"] = entry_group
        __props__["etag"] = etag
        __props__["members"] = members
        __props__["project"] = project
        __props__["region"] = region
        __props__["role"] = role
        return EntryGroupIamBinding(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
