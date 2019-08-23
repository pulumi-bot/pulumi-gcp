# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ClusterIAMPolicy(pulumi.CustomResource):
    cluster: pulumi.Output[str]
    """
    The name or relative resource id of the cluster to manage IAM policies for.
    """
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the clusters's IAM policy.
    """
    policy_data: pulumi.Output[str]
    """
    The policy data generated by a `organizations.getIAMPolicy` data source.
    """
    project: pulumi.Output[str]
    """
    The project in which the cluster belongs. If it
    is not provided, this provider will use the provider default.
    """
    region: pulumi.Output[str]
    """
    The region in which the cluster belongs. If it
    is not provided, this provider will use the provider default.
    """
    def __init__(__self__, resource_name, opts=None, cluster=None, policy_data=None, project=None, region=None, __props__=None, __name__=None, __opts__=None):
        """
        Three different resources help you manage IAM policies on dataproc clusters. Each of these resources serves a different use case:
        
        * `dataproc.ClusterIAMPolicy`: Authoritative. Sets the IAM policy for the cluster and replaces any existing policy already attached.
        * `dataproc.ClusterIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the cluster are preserved.
        * `dataproc.ClusterIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the cluster are preserved.
        
        > **Note:** `dataproc.ClusterIAMPolicy` **cannot** be used in conjunction with `dataproc.ClusterIAMBinding` and `dataproc.ClusterIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentaly unset ownership of the cluster as `dataproc.ClusterIAMPolicy` replaces the entire policy.
        
        > **Note:** `dataproc.ClusterIAMBinding` resources **can be** used in conjunction with `dataproc.ClusterIAMMember` resources **only if** they do not grant privilege to the same role.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster: The name or relative resource id of the cluster to manage IAM policies for.
        :param pulumi.Input[str] policy_data: The policy data generated by a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] project: The project in which the cluster belongs. If it
               is not provided, this provider will use the provider default.
        :param pulumi.Input[str] region: The region in which the cluster belongs. If it
               is not provided, this provider will use the provider default.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataproc_cluster_iam_policy.html.markdown.
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

            if cluster is None:
                raise TypeError("Missing required property 'cluster'")
            __props__['cluster'] = cluster
            if policy_data is None:
                raise TypeError("Missing required property 'policy_data'")
            __props__['policy_data'] = policy_data
            __props__['project'] = project
            __props__['region'] = region
            __props__['etag'] = None
        super(ClusterIAMPolicy, __self__).__init__(
            'gcp:dataproc/clusterIAMPolicy:ClusterIAMPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, cluster=None, etag=None, policy_data=None, project=None, region=None):
        """
        Get an existing ClusterIAMPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster: The name or relative resource id of the cluster to manage IAM policies for.
        :param pulumi.Input[str] etag: (Computed) The etag of the clusters's IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] project: The project in which the cluster belongs. If it
               is not provided, this provider will use the provider default.
        :param pulumi.Input[str] region: The region in which the cluster belongs. If it
               is not provided, this provider will use the provider default.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataproc_cluster_iam_policy.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["cluster"] = cluster
        __props__["etag"] = etag
        __props__["policy_data"] = policy_data
        __props__["project"] = project
        __props__["region"] = region
        return ClusterIAMPolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

