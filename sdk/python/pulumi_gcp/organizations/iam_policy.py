# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class IAMPolicy(pulumi.CustomResource):
    etag: pulumi.Output[str]
    org_id: pulumi.Output[str]
    """
    The numeric ID of the organization in which you want to create a custom role.
    """
    policy_data: pulumi.Output[str]
    """
    The `organizations.getIAMPolicy` data source that represents
    the IAM policy that will be applied to the organization. This policy overrides any existing
    policy applied to the organization.
    """
    def __init__(__self__, resource_name, opts=None, org_id=None, policy_data=None, __props__=None, __name__=None, __opts__=None):
        """
        Allows management of the entire IAM policy for an existing Google Cloud Platform Organization.

        !> **Warning:** New organizations have several default policies which will,
           without extreme caution, be **overwritten** by use of this resource.
           The safest alternative is to use multiple `organizations.IAMBinding`
           resources.  It is easy to use this resource to remove your own access to
           an organization, which will require a call to Google Support to have
           fixed, and can take multiple days to resolve.  If you do use this resource,
           the best way to be sure that you are not making dangerous changes is to start
           by importing your existing policy, and examining the diff very closely.

        > **Note:** This resource __must not__ be used in conjunction with
           `organizations.IAMMember` or `organizations.IAMBinding`
           or they will fight over what your policy should be.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(binding=[{
            "role": "roles/editor",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.organizations.IAMPolicy("policy",
            org_id="123456789",
            policy_data=admin.policy_data)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org_id: The numeric ID of the organization in which you want to create a custom role.
        :param pulumi.Input[str] policy_data: The `organizations.getIAMPolicy` data source that represents
               the IAM policy that will be applied to the organization. This policy overrides any existing
               policy applied to the organization.
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

            if org_id is None:
                raise TypeError("Missing required property 'org_id'")
            __props__['org_id'] = org_id
            if policy_data is None:
                raise TypeError("Missing required property 'policy_data'")
            __props__['policy_data'] = policy_data
            __props__['etag'] = None
        super(IAMPolicy, __self__).__init__(
            'gcp:organizations/iAMPolicy:IAMPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, etag=None, org_id=None, policy_data=None):
        """
        Get an existing IAMPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org_id: The numeric ID of the organization in which you want to create a custom role.
        :param pulumi.Input[str] policy_data: The `organizations.getIAMPolicy` data source that represents
               the IAM policy that will be applied to the organization. This policy overrides any existing
               policy applied to the organization.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["etag"] = etag
        __props__["org_id"] = org_id
        __props__["policy_data"] = policy_data
        return IAMPolicy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
