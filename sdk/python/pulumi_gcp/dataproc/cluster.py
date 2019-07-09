# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Cluster(pulumi.CustomResource):
    cluster_config: pulumi.Output[dict]
    """
    Allows you to configure various aspects of the cluster.
    Structure defined below.
    """
    labels: pulumi.Output[dict]
    """
    The list of labels (key/value pairs) to be applied to
    instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
    which is the name of the cluster.
    """
    name: pulumi.Output[str]
    """
    The name of the cluster, unique within the project and
    zone.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the `cluster` will exist. If it
    is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    The region in which the cluster and associated nodes will be created in.
    Defaults to `global`.
    """
    def __init__(__self__, resource_name, opts=None, cluster_config=None, labels=None, name=None, project=None, region=None, __name__=None, __opts__=None):
        """
        Manages a Cloud Dataproc cluster resource within GCP. For more information see
        [the official dataproc documentation](https://cloud.google.com/dataproc/).
        
        
        !> **Warning:** Due to limitations of the API, all arguments except
        `labels`,`cluster_config.worker_config.num_instances` and `cluster_config.preemptible_worker_config.num_instances` are non-updateable. Changing others will cause recreation of the
        whole cluster!
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] cluster_config: Allows you to configure various aspects of the cluster.
               Structure defined below.
        :param pulumi.Input[dict] labels: The list of labels (key/value pairs) to be applied to
               instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
               which is the name of the cluster.
        :param pulumi.Input[str] name: The name of the cluster, unique within the project and
               zone.
        :param pulumi.Input[str] project: The ID of the project in which the `cluster` will exist. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region in which the cluster and associated nodes will be created in.
               Defaults to `global`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataproc_cluster.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['cluster_config'] = cluster_config

        __props__['labels'] = labels

        __props__['name'] = name

        __props__['project'] = project

        __props__['region'] = region

        super(Cluster, __self__).__init__(
            'gcp:dataproc/cluster:Cluster',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

