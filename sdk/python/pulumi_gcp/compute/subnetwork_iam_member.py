# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class SubnetworkIAMMember(pulumi.CustomResource):
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
    member: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    The GCP region for this subnetwork.
    Used to find the parent resource to bind the IAM policy to. If not specified,
    the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
    region is specified, it is taken from the provider configuration.
    """
    role: pulumi.Output[str]
    """
    The role that should be applied. Only one
    `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    subnetwork: pulumi.Output[str]
    """
    Used to find the parent resource to bind the IAM policy to
    """
    def __init__(__self__, resource_name, opts=None, condition=None, member=None, project=None, region=None, role=None, subnetwork=None, __props__=None, __name__=None, __opts__=None):
        """
        Three different resources help you manage your IAM policy for Compute Engine Subnetwork. Each of these resources serves a different use case:

        * `compute.SubnetworkIAMPolicy`: Authoritative. Sets the IAM policy for the subnetwork and replaces any existing policy already attached.
        * `compute.SubnetworkIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subnetwork are preserved.
        * `compute.SubnetworkIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subnetwork are preserved.

        > **Note:** `compute.SubnetworkIAMPolicy` **cannot** be used in conjunction with `compute.SubnetworkIAMBinding` and `compute.SubnetworkIAMMember` or they will fight over what your policy should be.

        > **Note:** `compute.SubnetworkIAMBinding` resources **can be** used in conjunction with `compute.SubnetworkIAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\_compute\_subnetwork\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(gcp.organizations.GetIAMPolicyArgsArgs(
            bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
                role="roles/compute.networkUser",
                members=["user:jane@example.com"],
            )],
        ))
        policy = gcp.compute.SubnetworkIAMPolicy("policy",
            project=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["project"],
            region=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["region"],
            subnetwork=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["name"],
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(gcp.organizations.GetIAMPolicyArgsArgs(
            bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
                role="roles/compute.networkUser",
                members=["user:jane@example.com"],
                condition=gcp.organizations.GetIAMPolicyBindingConditionArgs(
                    title="expires_after_2019_12_31",
                    description="Expiring at midnight of 2019-12-31",
                    expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                ),
            )],
        ))
        policy = gcp.compute.SubnetworkIAMPolicy("policy",
            project=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["project"],
            region=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["region"],
            subnetwork=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["name"],
            policy_data=admin.policy_data)
        ```
        ## google\_compute\_subnetwork\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.compute.SubnetworkIAMBinding("binding",
            project=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["project"],
            region=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["region"],
            subnetwork=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["name"],
            role="roles/compute.networkUser",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.compute.SubnetworkIAMBinding("binding",
            project=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["project"],
            region=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["region"],
            subnetwork=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["name"],
            role="roles/compute.networkUser",
            members=["user:jane@example.com"],
            condition=gcp.compute.SubnetworkIAMBindingConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            ))
        ```
        ## google\_compute\_subnetwork\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.compute.SubnetworkIAMMember("member",
            project=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["project"],
            region=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["region"],
            subnetwork=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["name"],
            role="roles/compute.networkUser",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.compute.SubnetworkIAMMember("member",
            project=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["project"],
            region=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["region"],
            subnetwork=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["name"],
            role="roles/compute.networkUser",
            member="user:jane@example.com",
            condition=gcp.compute.SubnetworkIAMMemberConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] condition: ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] region: The GCP region for this subnetwork.
               Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
               region is specified, it is taken from the provider configuration.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] subnetwork: Used to find the parent resource to bind the IAM policy to

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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['condition'] = condition
            if member is None:
                raise TypeError("Missing required property 'member'")
            __props__['member'] = member
            __props__['project'] = project
            __props__['region'] = region
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            if subnetwork is None:
                raise TypeError("Missing required property 'subnetwork'")
            __props__['subnetwork'] = subnetwork
            __props__['etag'] = None
        super(SubnetworkIAMMember, __self__).__init__(
            'gcp:compute/subnetworkIAMMember:SubnetworkIAMMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, condition=None, etag=None, member=None, project=None, region=None, role=None, subnetwork=None):
        """
        Get an existing SubnetworkIAMMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] condition: ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] region: The GCP region for this subnetwork.
               Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
               region is specified, it is taken from the provider configuration.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] subnetwork: Used to find the parent resource to bind the IAM policy to

        The **condition** object supports the following:

          * `description` (`pulumi.Input[str]`) - An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
          * `expression` (`pulumi.Input[str]`) - Textual representation of an expression in Common Expression Language syntax.
          * `title` (`pulumi.Input[str]`) - A title for the expression, i.e. a short string describing its purpose.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["condition"] = condition
        __props__["etag"] = etag
        __props__["member"] = member
        __props__["project"] = project
        __props__["region"] = region
        __props__["role"] = role
        __props__["subnetwork"] = subnetwork
        return SubnetworkIAMMember(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
