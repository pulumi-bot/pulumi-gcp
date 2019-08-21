# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class IAMPolicy(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the folder's IAM policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
    """
    folder: pulumi.Output[str]
    """
    The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
    """
    policy_data: pulumi.Output[str]
    """
    The `organizations.getIAMPolicy` data source that represents
    the IAM policy that will be applied to the folder. This policy overrides any existing
    policy applied to the folder.
    """
    def __init__(__self__, resource_name, opts=None, folder=None, policy_data=None, __props__=None, __name__=None, __opts__=None):
        """
        Allows creation and management of the IAM policy for an existing Google Cloud
        Platform folder.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] folder: The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        :param pulumi.Input[str] policy_data: The `organizations.getIAMPolicy` data source that represents
               the IAM policy that will be applied to the folder. This policy overrides any existing
               policy applied to the folder.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/folder_iam_policy.html.markdown.
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

            if folder is None:
                raise TypeError("Missing required property 'folder'")
            __props__['folder'] = folder
            if policy_data is None:
                raise TypeError("Missing required property 'policy_data'")
            __props__['policy_data'] = policy_data
            __props__['etag'] = None
        super(IAMPolicy, __self__).__init__(
            'gcp:folder/iAMPolicy:IAMPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, etag=None, folder=None, policy_data=None):
        """
        Get an existing IAMPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the folder's IAM policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
        :param pulumi.Input[str] folder: The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        :param pulumi.Input[str] policy_data: The `organizations.getIAMPolicy` data source that represents
               the IAM policy that will be applied to the folder. This policy overrides any existing
               policy applied to the folder.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/folder_iam_policy.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["etag"] = etag
        __props__["folder"] = folder
        __props__["policy_data"] = policy_data
        return IAMPolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

