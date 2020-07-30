# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class DatasetIamMember(pulumi.CustomResource):
    condition: pulumi.Output[dict]
    dataset_id: pulumi.Output[str]
    """
    The dataset ID, in the form `projects/{project}/datasets/{dataset_id}`
    """
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the dataset's IAM policy.
    """
    member: pulumi.Output[str]
    project: pulumi.Output[str]
    role: pulumi.Output[str]
    """
    The role that should be applied. Only one
    `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    def __init__(__self__, resource_name, opts=None, condition=None, dataset_id=None, member=None, project=None, role=None, __props__=None, __name__=None, __opts__=None):
        """
        Three different resources help you manage your IAM policy for BigQuery dataset. Each of these resources serves a different use case:

        * `bigquery.DatasetIamPolicy`: Authoritative. Sets the IAM policy for the dataset and replaces any existing policy already attached.
        * `bigquery.DatasetIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the dataset are preserved.
        * `bigquery.DatasetIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the dataset are preserved.

        These resources are intended to convert the permissions system for BigQuery datasets to the standard IAM interface. For advanced usages, including [creating authorized views](https://cloud.google.com/bigquery/docs/share-access-views), please use either `bigquery.DatasetAccess` or the `access` field on `bigquery.Dataset`.

        > **Note:** These resources **cannot** be used with `bigquery.DatasetAccess` resources or the `access` field on `bigquery.Dataset` or they will fight over what the policy should be.

        > **Note:** Using any of these resources will remove any authorized view permissions from the dataset. To assign and preserve authorized view permissions use the `bigquery.DatasetAccess` instead.

        > **Note:** Legacy BigQuery roles `OWNER` `WRITER` and `READER` **cannot** be used with any of these IAM resources. Instead use the full role form of: `roles/bigquery.dataOwner` `roles/bigquery.dataEditor` and `roles/bigquery.dataViewer`.

        > **Note:** `bigquery.DatasetIamPolicy` **cannot** be used in conjunction with `bigquery.DatasetIamBinding` and `bigquery.DatasetIamMember` or they will fight over what your policy should be.

        > **Note:** `bigquery.DatasetIamBinding` resources **can be** used in conjunction with `bigquery.DatasetIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\bigquery\_dataset\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        owner = gcp.organizations.get_iam_policy(gcp.organizations.GetIAMPolicyArgsArgs(
            bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
                role="roles/dataOwner",
                members=["user:jane@example.com"],
            )],
        ))
        dataset = gcp.bigquery.DatasetIamPolicy("dataset",
            dataset_id="your-dataset-id",
            policy_data=owner.policy_data)
        ```

        ## google\_bigquery\_dataset\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        reader = gcp.bigquery.DatasetIamBinding("reader",
            dataset_id="your-dataset-id",
            members=["user:jane@example.com"],
            role="roles/bigquery.dataViewer")
        ```

        ## google\_bigquery\_dataset\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.bigquery.DatasetIamMember("editor",
            dataset_id="your-dataset-id",
            member="user:jane@example.com",
            role="roles/bigquery.dataEditor")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset_id: The dataset ID, in the form `projects/{project}/datasets/{dataset_id}`
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['condition'] = condition
            if dataset_id is None:
                raise TypeError("Missing required property 'dataset_id'")
            __props__['dataset_id'] = dataset_id
            if member is None:
                raise TypeError("Missing required property 'member'")
            __props__['member'] = member
            __props__['project'] = project
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['etag'] = None
        super(DatasetIamMember, __self__).__init__(
            'gcp:bigquery/datasetIamMember:DatasetIamMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, condition=None, dataset_id=None, etag=None, member=None, project=None, role=None):
        """
        Get an existing DatasetIamMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset_id: The dataset ID, in the form `projects/{project}/datasets/{dataset_id}`
        :param pulumi.Input[str] etag: (Computed) The etag of the dataset's IAM policy.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `bigquery.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.

        The **condition** object supports the following:

          * `description` (`pulumi.Input[str]`)
          * `expression` (`pulumi.Input[str]`)
          * `title` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["condition"] = condition
        __props__["dataset_id"] = dataset_id
        __props__["etag"] = etag
        __props__["member"] = member
        __props__["project"] = project
        __props__["role"] = role
        return DatasetIamMember(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
