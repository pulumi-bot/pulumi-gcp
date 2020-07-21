# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from ._inputs import *
from . import outputs


class IamPolicy(pulumi.CustomResource):
    etag: pulumi.Output[str] = pulumi.output_property("etag")
    """
    (Computed) The etag of the IAM policy.
    """
    location: pulumi.Output[str] = pulumi.output_property("location")
    """
    The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
    """
    policy_data: pulumi.Output[str] = pulumi.output_property("policyData")
    """
    The policy data generated by
    a `organizations.getIAMPolicy` data source.
    """
    project: pulumi.Output[str] = pulumi.output_property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
    """
    service: pulumi.Output[str] = pulumi.output_property("service")
    """
    Used to find the parent resource to bind the IAM policy to
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, location=None, policy_data=None, project=None, service=None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Three different resources help you manage your IAM policy for Cloud Run Service. Each of these resources serves a different use case:

        * `cloudrun.IamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
        * `cloudrun.IamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
        * `cloudrun.IamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.

        > **Note:** `cloudrun.IamPolicy` **cannot** be used in conjunction with `cloudrun.IamBinding` and `cloudrun.IamMember` or they will fight over what your policy should be.

        > **Note:** `cloudrun.IamBinding` resources **can be** used in conjunction with `cloudrun.IamMember` resources **only if** they do not grant privilege to the same role.

        ## google\_cloud\_run\_service\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.cloudrun.IamPolicy("policy",
            location=google_cloud_run_service["default"]["location"],
            project=google_cloud_run_service["default"]["project"],
            service=google_cloud_run_service["default"]["name"],
            policy_data=admin.policy_data)
        ```

        ## google\_cloud\_run\_service\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.cloudrun.IamBinding("binding",
            location=google_cloud_run_service["default"]["location"],
            project=google_cloud_run_service["default"]["project"],
            service=google_cloud_run_service["default"]["name"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## google\_cloud\_run\_service\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.cloudrun.IamMember("member",
            location=google_cloud_run_service["default"]["location"],
            project=google_cloud_run_service["default"]["project"],
            service=google_cloud_run_service["default"]["name"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] service: Used to find the parent resource to bind the IAM policy to
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

            __props__['location'] = location
            if policy_data is None:
                raise TypeError("Missing required property 'policy_data'")
            __props__['policy_data'] = policy_data
            __props__['project'] = project
            if service is None:
                raise TypeError("Missing required property 'service'")
            __props__['service'] = service
            __props__['etag'] = None
        super(IamPolicy, __self__).__init__(
            'gcp:cloudrun/iamPolicy:IamPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, etag=None, location=None, policy_data=None, project=None, service=None):
        """
        Get an existing IamPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] location: The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] service: Used to find the parent resource to bind the IAM policy to
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["etag"] = etag
        __props__["location"] = location
        __props__["policy_data"] = policy_data
        __props__["project"] = project
        __props__["service"] = service
        return IamPolicy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

