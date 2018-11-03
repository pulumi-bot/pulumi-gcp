# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Cluster(pulumi.CustomResource):
    """
    Manages a Cloud Dataproc cluster resource within GCP. For more information see
    [the official dataproc documentation](https://cloud.google.com/dataproc/).
    
    
    !> **Warning:** Due to limitations of the API, all arguments except
    `labels`,`cluster_config.worker_config.num_instances` and `cluster_config.preemptible_worker_config.num_instances` are non-updateable. Changing others will cause recreation of the
    whole cluster!
    """
    def __init__(__self__, __name__, __opts__=None, cluster_config=None, labels=None, name=None, project=None, region=None):
        """Create a Cluster resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['clusterConfig'] = cluster_config

        __props__['labels'] = labels

        __props__['name'] = name

        __props__['project'] = project

        __props__['region'] = region

        super(Cluster, __self__).__init__(
            'gcp:dataproc/cluster:Cluster',
            __name__,
            __props__,
            __opts__)

