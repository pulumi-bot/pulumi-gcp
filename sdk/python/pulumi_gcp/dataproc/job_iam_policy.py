# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class JobIAMPolicy(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the jobs's IAM policy.
    """
    job_id: pulumi.Output[str]
    policy_data: pulumi.Output[str]
    """
    The policy data generated by a `organizations.getIAMPolicy` data source.
    """
    project: pulumi.Output[str]
    """
    The project in which the job belongs. If it
    is not provided, the provider will use a default.
    """
    region: pulumi.Output[str]
    """
    The region in which the job belongs. If it
    is not provided, the provider will use a default.
    """
    def __init__(__self__, resource_name, opts=None, job_id=None, policy_data=None, project=None, region=None, __props__=None, __name__=None, __opts__=None):
        """
        Three different resources help you manage IAM policies on dataproc jobs. Each of these resources serves a different use case:

        * `dataproc.JobIAMPolicy`: Authoritative. Sets the IAM policy for the job and replaces any existing policy already attached.
        * `dataproc.JobIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the job are preserved.
        * `dataproc.JobIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the job are preserved.

        > **Note:** `dataproc.JobIAMPolicy` **cannot** be used in conjunction with `dataproc.JobIAMBinding` and `dataproc.JobIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the job as `dataproc.JobIAMPolicy` replaces the entire policy.

        > **Note:** `dataproc.JobIAMBinding` resources **can be** used in conjunction with `dataproc.JobIAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\_pubsub\_subscription\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(binding=[{
            "role": "roles/editor",
            "members": ["user:jane@example.com"],
        }])
        editor = gcp.dataproc.JobIAMPolicy("editor",
            project="your-project",
            region="your-region",
            job_id="your-dataproc-job",
            policy_data=admin.policy_data)
        ```

        ## google\_pubsub\_subscription\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.dataproc.JobIAMBinding("editor",
            job_id="your-dataproc-job",
            members=["user:jane@example.com"],
            role="roles/editor")
        ```

        ## google\_pubsub\_subscription\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.dataproc.JobIAMMember("editor",
            job_id="your-dataproc-job",
            member="user:jane@example.com",
            role="roles/editor")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_data: The policy data generated by a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] project: The project in which the job belongs. If it
               is not provided, the provider will use a default.
        :param pulumi.Input[str] region: The region in which the job belongs. If it
               is not provided, the provider will use a default.
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

            if job_id is None:
                raise TypeError("Missing required property 'job_id'")
            __props__['job_id'] = job_id
            if policy_data is None:
                raise TypeError("Missing required property 'policy_data'")
            __props__['policy_data'] = policy_data
            __props__['project'] = project
            __props__['region'] = region
            __props__['etag'] = None
        super(JobIAMPolicy, __self__).__init__(
            'gcp:dataproc/jobIAMPolicy:JobIAMPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, etag=None, job_id=None, policy_data=None, project=None, region=None):
        """
        Get an existing JobIAMPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the jobs's IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] project: The project in which the job belongs. If it
               is not provided, the provider will use a default.
        :param pulumi.Input[str] region: The region in which the job belongs. If it
               is not provided, the provider will use a default.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["etag"] = etag
        __props__["job_id"] = job_id
        __props__["policy_data"] = policy_data
        __props__["project"] = project
        __props__["region"] = region
        return JobIAMPolicy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
