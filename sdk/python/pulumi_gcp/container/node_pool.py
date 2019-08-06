# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class NodePool(pulumi.CustomResource):
    autoscaling: pulumi.Output[dict]
    """
    Configuration required by cluster autoscaler to adjust
    the size of the node pool to the current cluster usage. Structure is documented below.
    """
    cluster: pulumi.Output[str]
    """
    The cluster to create the node pool for.  Cluster must be present in `zone` provided for zonal clusters.
    """
    initial_node_count: pulumi.Output[float]
    """
    The initial node count for the pool. Changing this will force
    recreation of the resource.
    """
    instance_group_urls: pulumi.Output[list]
    location: pulumi.Output[str]
    """
    The location (region or zone) in which the cluster
    resides.
    """
    management: pulumi.Output[dict]
    """
    Node management configuration, wherein auto-repair and
    auto-upgrade is configured. Structure is documented below.
    """
    max_pods_per_node: pulumi.Output[float]
    """
    ) The maximum number of pods per node in this node pool.
    Note that this does not work on node pools which are "route-based" - that is, node
    pools belonging to clusters that do not have IP Aliasing enabled.
    See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
    for more information.
    """
    name: pulumi.Output[str]
    """
    The name of the node pool. If left blank, this provider will
    auto-generate a unique name.
    """
    name_prefix: pulumi.Output[str]
    node_config: pulumi.Output[dict]
    """
    The node configuration of the pool. See
    container.Cluster for schema.
    """
    node_count: pulumi.Output[float]
    """
    The number of nodes per instance group. This field can be used to
    update the number of nodes per instance group but should not be used alongside `autoscaling`.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which to create the node pool. If blank,
    the provider-configured project will be used.
    """
    region: pulumi.Output[str]
    """
    The region in which the cluster resides (for
    regional clusters). `zone` has been deprecated in favor of `location`.
    """
    version: pulumi.Output[str]
    """
    The Kubernetes version for the nodes in this pool. Note that if this field
    and `auto_upgrade` are both specified, they will fight each other for what the node version should
    be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
    recommended that you specify explicit versions as this provider will see spurious diffs
    when fuzzy versions are used. See the `google_container_engine_versions` data source's
    `version_prefix` field to approximate fuzzy versions.
    """
    zone: pulumi.Output[str]
    """
    The zone in which the cluster resides. `zone`
    has been deprecated in favor of `location`.
    """
    def __init__(__self__, resource_name, opts=None, autoscaling=None, cluster=None, initial_node_count=None, location=None, management=None, max_pods_per_node=None, name=None, name_prefix=None, node_config=None, node_count=None, project=None, region=None, version=None, zone=None, __name__=None, __opts__=None):
        """
        Manages a node pool in a Google Kubernetes Engine (GKE) cluster separately from
        the cluster control plane. For more information see [the official documentation](https://cloud.google.com/container-engine/docs/node-pools)
        and [the API reference](https://cloud.google.com/container-engine/reference/rest/v1/projects.zones.clusters.nodePools).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] autoscaling: Configuration required by cluster autoscaler to adjust
               the size of the node pool to the current cluster usage. Structure is documented below.
        :param pulumi.Input[str] cluster: The cluster to create the node pool for.  Cluster must be present in `zone` provided for zonal clusters.
        :param pulumi.Input[float] initial_node_count: The initial node count for the pool. Changing this will force
               recreation of the resource.
        :param pulumi.Input[str] location: The location (region or zone) in which the cluster
               resides.
        :param pulumi.Input[dict] management: Node management configuration, wherein auto-repair and
               auto-upgrade is configured. Structure is documented below.
        :param pulumi.Input[float] max_pods_per_node: ) The maximum number of pods per node in this node pool.
               Note that this does not work on node pools which are "route-based" - that is, node
               pools belonging to clusters that do not have IP Aliasing enabled.
               See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
               for more information.
        :param pulumi.Input[str] name: The name of the node pool. If left blank, this provider will
               auto-generate a unique name.
        :param pulumi.Input[dict] node_config: The node configuration of the pool. See
               container.Cluster for schema.
        :param pulumi.Input[float] node_count: The number of nodes per instance group. This field can be used to
               update the number of nodes per instance group but should not be used alongside `autoscaling`.
        :param pulumi.Input[str] project: The ID of the project in which to create the node pool. If blank,
               the provider-configured project will be used.
        :param pulumi.Input[str] region: The region in which the cluster resides (for
               regional clusters). `zone` has been deprecated in favor of `location`.
        :param pulumi.Input[str] version: The Kubernetes version for the nodes in this pool. Note that if this field
               and `auto_upgrade` are both specified, they will fight each other for what the node version should
               be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
               recommended that you specify explicit versions as this provider will see spurious diffs
               when fuzzy versions are used. See the `google_container_engine_versions` data source's
               `version_prefix` field to approximate fuzzy versions.
        :param pulumi.Input[str] zone: The zone in which the cluster resides. `zone`
               has been deprecated in favor of `location`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/container_node_pool.html.markdown.
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

        __props__['autoscaling'] = autoscaling

        if cluster is None:
            raise TypeError("Missing required property 'cluster'")
        __props__['cluster'] = cluster

        __props__['initial_node_count'] = initial_node_count

        __props__['location'] = location

        __props__['management'] = management

        __props__['max_pods_per_node'] = max_pods_per_node

        __props__['name'] = name

        __props__['name_prefix'] = name_prefix

        __props__['node_config'] = node_config

        __props__['node_count'] = node_count

        __props__['project'] = project

        __props__['region'] = region

        __props__['version'] = version

        __props__['zone'] = zone

        __props__['instance_group_urls'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(NodePool, __self__).__init__(
            'gcp:container/nodePool:NodePool',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

