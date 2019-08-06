# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class SubnetworkIAMBinding(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the subnetwork's IAM policy.
    """
    members: pulumi.Output[list]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    The region of the subnetwork. If
    unspecified, this defaults to the region configured in the provider.
    """
    role: pulumi.Output[str]
    """
    The role that should be applied. Only one
    `google_compute_subnetwork_iam_binding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    subnetwork: pulumi.Output[str]
    """
    The name of the subnetwork.
    """
    def __init__(__self__, resource_name, opts=None, members=None, project=None, region=None, role=None, subnetwork=None, __name__=None, __opts__=None):
        """
        Create a SubnetworkIAMBinding resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the subnetwork. If
               unspecified, this defaults to the region configured in the provider.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `google_compute_subnetwork_iam_binding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] subnetwork: The name of the subnetwork.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_subnetwork_iam_binding.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__

        __props__ = dict()

        if members is None:
            raise TypeError("Missing required property 'members'")
        __props__['members'] = members
        __props__['project'] = project
        __props__['region'] = region
        if role is None:
            raise TypeError("Missing required property 'role'")
        __props__['role'] = role
        if subnetwork is None:
            raise TypeError("Missing required property 'subnetwork'")
        __props__['subnetwork'] = subnetwork
        __props__['etag'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(SubnetworkIAMBinding, __self__).__init__(
            'gcp:compute/subnetworkIAMBinding:SubnetworkIAMBinding',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

