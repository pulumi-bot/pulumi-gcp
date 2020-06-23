# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class InstanceIAMBinding(pulumi.CustomResource):
    condition: pulumi.Output[dict]
    """
    ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
    Structure is documented below.

      * `description` (`str`) - An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
      * `expression` (`str`) - Textual representation of an expression in Common Expression Language syntax.
      * `title` (`str`) - A title for the expression, i.e. a short string describing its purpose.
    """
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the IAM policy.
    """
    instance_name: pulumi.Output[str]
    """
    Used to find the parent resource to bind the IAM policy to
    """
    members: pulumi.Output[list]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
    """
    role: pulumi.Output[str]
    """
    The role that should be applied. Only one
    `compute.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    zone: pulumi.Output[str]
    """
    A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
    the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
    zone is specified, it is taken from the provider configuration.
    """
    def __init__(__self__, resource_name, opts=None, condition=None, instance_name=None, members=None, project=None, role=None, zone=None, __props__=None, __name__=None, __opts__=None):
        """
        Three different resources help you manage your IAM policy for Compute Engine Instance. Each of these resources serves a different use case:

        * `compute.InstanceIAMPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
        * `compute.InstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
        * `compute.InstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.

        > **Note:** `compute.InstanceIAMPolicy` **cannot** be used in conjunction with `compute.InstanceIAMBinding` and `compute.InstanceIAMMember` or they will fight over what your policy should be.

        > **Note:** `compute.InstanceIAMBinding` resources **can be** used in conjunction with `compute.InstanceIAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\_compute\_instance\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(binding=[{
            "role": "roles/compute.osLogin",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.compute.InstanceIAMPolicy("policy",
            project=google_compute_instance["default"]["project"],
            zone=google_compute_instance["default"]["zone"],
            instance_name=google_compute_instance["default"]["name"],
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(binding=[{
            "role": "roles/compute.osLogin",
            "members": ["user:jane@example.com"],
            "condition": {
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            },
        }])
        policy = gcp.compute.InstanceIAMPolicy("policy",
            project=google_compute_instance["default"]["project"],
            zone=google_compute_instance["default"]["zone"],
            instance_name=google_compute_instance["default"]["name"],
            policy_data=admin.policy_data)
        ```
        ## google\_compute\_instance\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.compute.InstanceIAMBinding("binding",
            project=google_compute_instance["default"]["project"],
            zone=google_compute_instance["default"]["zone"],
            instance_name=google_compute_instance["default"]["name"],
            role="roles/compute.osLogin",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.compute.InstanceIAMBinding("binding",
            project=google_compute_instance["default"]["project"],
            zone=google_compute_instance["default"]["zone"],
            instance_name=google_compute_instance["default"]["name"],
            role="roles/compute.osLogin",
            members=["user:jane@example.com"],
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            })
        ```
        ## google\_compute\_instance\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.compute.InstanceIAMMember("member",
            project=google_compute_instance["default"]["project"],
            zone=google_compute_instance["default"]["zone"],
            instance_name=google_compute_instance["default"]["name"],
            role="roles/compute.osLogin",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.compute.InstanceIAMMember("member",
            project=google_compute_instance["default"]["project"],
            zone=google_compute_instance["default"]["zone"],
            instance_name=google_compute_instance["default"]["name"],
            role="roles/compute.osLogin",
            member="user:jane@example.com",
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] condition: ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] instance_name: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `compute.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] zone: A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
               zone is specified, it is taken from the provider configuration.

        The **condition** object supports the following:

          * `description` (`pulumi.Input[str]`) - An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
          * `expression` (`pulumi.Input[str]`) - Textual representation of an expression in Common Expression Language syntax.
          * `title` (`pulumi.Input[str]`) - A title for the expression, i.e. a short string describing its purpose.
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
            if instance_name is None:
                raise TypeError("Missing required property 'instance_name'")
            __props__['instance_name'] = instance_name
            if members is None:
                raise TypeError("Missing required property 'members'")
            __props__['members'] = members
            __props__['project'] = project
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['zone'] = zone
            __props__['etag'] = None
        super(InstanceIAMBinding, __self__).__init__(
            'gcp:compute/instanceIAMBinding:InstanceIAMBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, condition=None, etag=None, instance_name=None, members=None, project=None, role=None, zone=None):
        """
        Get an existing InstanceIAMBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] condition: ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] instance_name: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `compute.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] zone: A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
               zone is specified, it is taken from the provider configuration.

        The **condition** object supports the following:

          * `description` (`pulumi.Input[str]`) - An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
          * `expression` (`pulumi.Input[str]`) - Textual representation of an expression in Common Expression Language syntax.
          * `title` (`pulumi.Input[str]`) - A title for the expression, i.e. a short string describing its purpose.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["condition"] = condition
        __props__["etag"] = etag
        __props__["instance_name"] = instance_name
        __props__["members"] = members
        __props__["project"] = project
        __props__["role"] = role
        __props__["zone"] = zone
        return InstanceIAMBinding(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
