# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Hl7StoreIamMember(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the HL7v2 store's IAM policy.
    """
    hl7_v2_store_id: pulumi.Output[str]
    """
    The HL7v2 store ID, in the form
    `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
    `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
    project setting will be used as a fallback.
    """
    member: pulumi.Output[str]
    role: pulumi.Output[str]
    """
    The role that should be applied. Only one
    `healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    def __init__(__self__, resource_name, opts=None, hl7_v2_store_id=None, member=None, role=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Hl7StoreIamMember resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hl7_v2_store_id: The HL7v2 store ID, in the form
               `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
               `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store_iam_member.html.markdown.
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

            if hl7_v2_store_id is None:
                raise TypeError("Missing required property 'hl7_v2_store_id'")
            __props__['hl7_v2_store_id'] = hl7_v2_store_id
            if member is None:
                raise TypeError("Missing required property 'member'")
            __props__['member'] = member
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['etag'] = None
        super(Hl7StoreIamMember, __self__).__init__(
            'gcp:healthcare/hl7StoreIamMember:Hl7StoreIamMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, etag=None, hl7_v2_store_id=None, member=None, role=None):
        """
        Get an existing Hl7StoreIamMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the HL7v2 store's IAM policy.
        :param pulumi.Input[str] hl7_v2_store_id: The HL7v2 store ID, in the form
               `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
               `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store_iam_member.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["etag"] = etag
        __props__["hl7_v2_store_id"] = hl7_v2_store_id
        __props__["member"] = member
        __props__["role"] = role
        return Hl7StoreIamMember(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

