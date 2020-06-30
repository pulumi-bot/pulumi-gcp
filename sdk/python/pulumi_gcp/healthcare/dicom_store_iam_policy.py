# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class DicomStoreIamPolicy(pulumi.CustomResource):
    dicom_store_id: pulumi.Output[str]
    """
    The DICOM store ID, in the form
    `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
    `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
    project setting will be used as a fallback.
    """
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the DICOM store's IAM policy.
    """
    policy_data: pulumi.Output[str]
    """
    The policy data generated by
    a `organizations.getIAMPolicy` data source.
    """
    def __init__(__self__, resource_name, opts=None, dicom_store_id=None, policy_data=None, __props__=None, __name__=None, __opts__=None):
        """
        Three different resources help you manage your IAM policy for Healthcare DICOM store. Each of these resources serves a different use case:

        * `healthcare.DicomStoreIamPolicy`: Authoritative. Sets the IAM policy for the DICOM store and replaces any existing policy already attached.
        * `healthcare.DicomStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the DICOM store are preserved.
        * `healthcare.DicomStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the DICOM store are preserved.

        > **Note:** `healthcare.DicomStoreIamPolicy` **cannot** be used in conjunction with `healthcare.DicomStoreIamBinding` and `healthcare.DicomStoreIamMember` or they will fight over what your policy should be.

        > **Note:** `healthcare.DicomStoreIamBinding` resources **can be** used in conjunction with `healthcare.DicomStoreIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\_healthcare\_dicom\_store\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(binding=[{
            "role": "roles/editor",
            "members": ["user:jane@example.com"],
        }])
        dicom_store = gcp.healthcare.DicomStoreIamPolicy("dicomStore",
            dicom_store_id="your-dicom-store-id",
            policy_data=admin.policy_data)
        ```

        ## google\_healthcare\_dicom\_store\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dicom_store = gcp.healthcare.DicomStoreIamBinding("dicomStore",
            dicom_store_id="your-dicom-store-id",
            members=["user:jane@example.com"],
            role="roles/editor")
        ```

        ## google\_healthcare\_dicom\_store\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dicom_store = gcp.healthcare.DicomStoreIamMember("dicomStore",
            dicom_store_id="your-dicom-store-id",
            member="user:jane@example.com",
            role="roles/editor")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dicom_store_id: The DICOM store ID, in the form
               `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
               `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
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

            if dicom_store_id is None:
                raise TypeError("Missing required property 'dicom_store_id'")
            __props__['dicom_store_id'] = dicom_store_id
            if policy_data is None:
                raise TypeError("Missing required property 'policy_data'")
            __props__['policy_data'] = policy_data
            __props__['etag'] = None
        super(DicomStoreIamPolicy, __self__).__init__(
            'gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, dicom_store_id=None, etag=None, policy_data=None):
        """
        Get an existing DicomStoreIamPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dicom_store_id: The DICOM store ID, in the form
               `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
               `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] etag: (Computed) The etag of the DICOM store's IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dicom_store_id"] = dicom_store_id
        __props__["etag"] = etag
        __props__["policy_data"] = policy_data
        return DicomStoreIamPolicy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
