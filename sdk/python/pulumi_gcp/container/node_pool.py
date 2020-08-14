# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class NodePool(pulumi.CustomResource):
    autoscaling: pulumi.Output[dict]
    """
    Configuration required by cluster autoscaler to adjust
    the size of the node pool to the current cluster usage. Structure is documented below.

      * `maxNodeCount` (`float`) - Maximum number of nodes in the NodePool. Must be >= min_node_count.
      * `minNodeCount` (`float`) - Minimum number of nodes in the NodePool. Must be >=0 and
        <= `max_node_count`.
    """
    cluster: pulumi.Output[str]
    """
    The cluster to create the node pool for. Cluster must be present in `location` provided for zonal clusters.
    """
    initial_node_count: pulumi.Output[float]
    """
    The initial number of nodes for the pool. In
    regional or multi-zonal clusters, this is the number of nodes per zone. Changing
    this will force recreation of the resource.
    """
    instance_group_urls: pulumi.Output[list]
    """
    The resource URLs of the managed instance groups associated with this node pool.
    """
    location: pulumi.Output[str]
    """
    The location (region or zone) of the cluster.
    """
    management: pulumi.Output[dict]
    """
    Node management configuration, wherein auto-repair and
    auto-upgrade is configured. Structure is documented below.

      * `autoRepair` (`bool`) - Whether the nodes will be automatically repaired.
      * `autoUpgrade` (`bool`) - Whether the nodes will be automatically upgraded.
    """
    max_pods_per_node: pulumi.Output[float]
    """
    The maximum number of pods per node in this node pool.
    Note that this does not work on node pools which are "route-based" - that is, node
    pools belonging to clusters that do not have IP Aliasing enabled.
    See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
    for more information.
    """
    name: pulumi.Output[str]
    """
    The name of the node pool. If left blank, the provider will
    auto-generate a unique name.
    """
    name_prefix: pulumi.Output[str]
    """
    Creates a unique name for the node pool beginning
    with the specified prefix. Conflicts with `name`.
    """
    node_config: pulumi.Output[dict]
    """
    The node configuration of the pool. See
    container.Cluster for schema.

      * `bootDiskKmsKey` (`str`)
      * `disk_size_gb` (`float`)
      * `diskType` (`str`)
      * `guest_accelerators` (`list`)
        * `count` (`float`)
        * `type` (`str`)

      * `imageType` (`str`)
      * `labels` (`dict`)
      * `localSsdCount` (`float`)
      * `machine_type` (`str`)
      * `metadata` (`dict`)
      * `min_cpu_platform` (`str`)
      * `oauthScopes` (`list`)
      * `preemptible` (`bool`)
      * `sandboxConfig` (`dict`)
        * `sandboxType` (`str`)

      * `service_account` (`str`)
      * `shielded_instance_config` (`dict`)
        * `enableIntegrityMonitoring` (`bool`)
        * `enableSecureBoot` (`bool`)

      * `tags` (`list`)
      * `taints` (`list`)
        * `effect` (`str`)
        * `key` (`str`)
        * `value` (`str`)

      * `workloadMetadataConfig` (`dict`)
        * `nodeMetadata` (`str`)
    """
    node_count: pulumi.Output[float]
    """
    The number of nodes per instance group. This field can be used to
    update the number of nodes per instance group but should not be used alongside `autoscaling`.
    """
    node_locations: pulumi.Output[list]
    """
    The list of zones in which the node pool's nodes should be located. Nodes must
    be in the region of their regional cluster or in the same region as their
    cluster's zone for zonal clusters. If unspecified, the cluster-level
    `node_locations` will be used.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which to create the node pool. If blank,
    the provider-configured project will be used.
    """
    upgrade_settings: pulumi.Output[dict]
    """
    Specify node upgrade settings to change how many nodes GKE attempts to
    upgrade at once. The number of nodes upgraded simultaneously is the sum of `max_surge` and `max_unavailable`.
    The maximum number of nodes upgraded simultaneously is limited to 20.

      * `maxSurge` (`float`) - The number of additional nodes that can be added to the node pool during
        an upgrade. Increasing `max_surge` raises the number of nodes that can be upgraded simultaneously.
        Can be set to 0 or greater.
      * `maxUnavailable` (`float`) - The number of nodes that can be simultaneously unavailable during
        an upgrade. Increasing `max_unavailable` raises the number of nodes that can be upgraded in
        parallel. Can be set to 0 or greater.
    """
    version: pulumi.Output[str]
    """
    The Kubernetes version for the nodes in this pool. Note that if this field
    and `auto_upgrade` are both specified, they will fight each other for what the node version should
    be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
    recommended that you specify explicit versions as the provider will see spurious diffs
    when fuzzy versions are used. See the `container.getEngineVersions` data source's
    `version_prefix` field to approximate fuzzy versions in a provider-compatible way.
    """
    def __init__(__self__, resource_name, opts=None, autoscaling=None, cluster=None, initial_node_count=None, location=None, management=None, max_pods_per_node=None, name=None, name_prefix=None, node_config=None, node_count=None, node_locations=None, project=None, upgrade_settings=None, version=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a node pool in a Google Kubernetes Engine (GKE) cluster separately from
        the cluster control plane. For more information see [the official documentation](https://cloud.google.com/container-engine/docs/node-pools)
        and [the API reference](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1beta1/projects.locations.clusters.nodePools).

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] autoscaling: Configuration required by cluster autoscaler to adjust
               the size of the node pool to the current cluster usage. Structure is documented below.
        :param pulumi.Input[str] cluster: The cluster to create the node pool for. Cluster must be present in `location` provided for zonal clusters.
        :param pulumi.Input[float] initial_node_count: The initial number of nodes for the pool. In
               regional or multi-zonal clusters, this is the number of nodes per zone. Changing
               this will force recreation of the resource.
        :param pulumi.Input[str] location: The location (region or zone) of the cluster.
        :param pulumi.Input[dict] management: Node management configuration, wherein auto-repair and
               auto-upgrade is configured. Structure is documented below.
        :param pulumi.Input[float] max_pods_per_node: The maximum number of pods per node in this node pool.
               Note that this does not work on node pools which are "route-based" - that is, node
               pools belonging to clusters that do not have IP Aliasing enabled.
               See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
               for more information.
        :param pulumi.Input[str] name: The name of the node pool. If left blank, the provider will
               auto-generate a unique name.
        :param pulumi.Input[str] name_prefix: Creates a unique name for the node pool beginning
               with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[dict] node_config: The node configuration of the pool. See
               container.Cluster for schema.
        :param pulumi.Input[float] node_count: The number of nodes per instance group. This field can be used to
               update the number of nodes per instance group but should not be used alongside `autoscaling`.
        :param pulumi.Input[list] node_locations: The list of zones in which the node pool's nodes should be located. Nodes must
               be in the region of their regional cluster or in the same region as their
               cluster's zone for zonal clusters. If unspecified, the cluster-level
               `node_locations` will be used.
        :param pulumi.Input[str] project: The ID of the project in which to create the node pool. If blank,
               the provider-configured project will be used.
        :param pulumi.Input[dict] upgrade_settings: Specify node upgrade settings to change how many nodes GKE attempts to
               upgrade at once. The number of nodes upgraded simultaneously is the sum of `max_surge` and `max_unavailable`.
               The maximum number of nodes upgraded simultaneously is limited to 20.
        :param pulumi.Input[str] version: The Kubernetes version for the nodes in this pool. Note that if this field
               and `auto_upgrade` are both specified, they will fight each other for what the node version should
               be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
               recommended that you specify explicit versions as the provider will see spurious diffs
               when fuzzy versions are used. See the `container.getEngineVersions` data source's
               `version_prefix` field to approximate fuzzy versions in a provider-compatible way.

        The **autoscaling** object supports the following:

          * `maxNodeCount` (`pulumi.Input[float]`) - Maximum number of nodes in the NodePool. Must be >= min_node_count.
          * `minNodeCount` (`pulumi.Input[float]`) - Minimum number of nodes in the NodePool. Must be >=0 and
            <= `max_node_count`.

        The **management** object supports the following:

          * `autoRepair` (`pulumi.Input[bool]`) - Whether the nodes will be automatically repaired.
          * `autoUpgrade` (`pulumi.Input[bool]`) - Whether the nodes will be automatically upgraded.

        The **node_config** object supports the following:

          * `bootDiskKmsKey` (`pulumi.Input[str]`)
          * `disk_size_gb` (`pulumi.Input[float]`)
          * `diskType` (`pulumi.Input[str]`)
          * `guest_accelerators` (`pulumi.Input[list]`)
            * `count` (`pulumi.Input[float]`)
            * `type` (`pulumi.Input[str]`)

          * `imageType` (`pulumi.Input[str]`)
          * `labels` (`pulumi.Input[dict]`)
          * `localSsdCount` (`pulumi.Input[float]`)
          * `machine_type` (`pulumi.Input[str]`)
          * `metadata` (`pulumi.Input[dict]`)
          * `min_cpu_platform` (`pulumi.Input[str]`)
          * `oauthScopes` (`pulumi.Input[list]`)
          * `preemptible` (`pulumi.Input[bool]`)
          * `sandboxConfig` (`pulumi.Input[dict]`)
            * `sandboxType` (`pulumi.Input[str]`)

          * `service_account` (`pulumi.Input[str]`)
          * `shielded_instance_config` (`pulumi.Input[dict]`)
            * `enableIntegrityMonitoring` (`pulumi.Input[bool]`)
            * `enableSecureBoot` (`pulumi.Input[bool]`)

          * `tags` (`pulumi.Input[list]`)
          * `taints` (`pulumi.Input[list]`)
            * `effect` (`pulumi.Input[str]`)
            * `key` (`pulumi.Input[str]`)
            * `value` (`pulumi.Input[str]`)

          * `workloadMetadataConfig` (`pulumi.Input[dict]`)
            * `nodeMetadata` (`pulumi.Input[str]`)

        The **upgrade_settings** object supports the following:

          * `maxSurge` (`pulumi.Input[float]`) - The number of additional nodes that can be added to the node pool during
            an upgrade. Increasing `max_surge` raises the number of nodes that can be upgraded simultaneously.
            Can be set to 0 or greater.
          * `maxUnavailable` (`pulumi.Input[float]`) - The number of nodes that can be simultaneously unavailable during
            an upgrade. Increasing `max_unavailable` raises the number of nodes that can be upgraded in
            parallel. Can be set to 0 or greater.
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
            __props__['node_locations'] = node_locations
            __props__['project'] = project
            __props__['upgrade_settings'] = upgrade_settings
            __props__['version'] = version
            __props__['instance_group_urls'] = None
        super(NodePool, __self__).__init__(
            'gcp:container/nodePool:NodePool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, autoscaling=None, cluster=None, initial_node_count=None, instance_group_urls=None, location=None, management=None, max_pods_per_node=None, name=None, name_prefix=None, node_config=None, node_count=None, node_locations=None, project=None, upgrade_settings=None, version=None):
        """
        Get an existing NodePool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] autoscaling: Configuration required by cluster autoscaler to adjust
               the size of the node pool to the current cluster usage. Structure is documented below.
        :param pulumi.Input[str] cluster: The cluster to create the node pool for. Cluster must be present in `location` provided for zonal clusters.
        :param pulumi.Input[float] initial_node_count: The initial number of nodes for the pool. In
               regional or multi-zonal clusters, this is the number of nodes per zone. Changing
               this will force recreation of the resource.
        :param pulumi.Input[list] instance_group_urls: The resource URLs of the managed instance groups associated with this node pool.
        :param pulumi.Input[str] location: The location (region or zone) of the cluster.
        :param pulumi.Input[dict] management: Node management configuration, wherein auto-repair and
               auto-upgrade is configured. Structure is documented below.
        :param pulumi.Input[float] max_pods_per_node: The maximum number of pods per node in this node pool.
               Note that this does not work on node pools which are "route-based" - that is, node
               pools belonging to clusters that do not have IP Aliasing enabled.
               See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
               for more information.
        :param pulumi.Input[str] name: The name of the node pool. If left blank, the provider will
               auto-generate a unique name.
        :param pulumi.Input[str] name_prefix: Creates a unique name for the node pool beginning
               with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[dict] node_config: The node configuration of the pool. See
               container.Cluster for schema.
        :param pulumi.Input[float] node_count: The number of nodes per instance group. This field can be used to
               update the number of nodes per instance group but should not be used alongside `autoscaling`.
        :param pulumi.Input[list] node_locations: The list of zones in which the node pool's nodes should be located. Nodes must
               be in the region of their regional cluster or in the same region as their
               cluster's zone for zonal clusters. If unspecified, the cluster-level
               `node_locations` will be used.
        :param pulumi.Input[str] project: The ID of the project in which to create the node pool. If blank,
               the provider-configured project will be used.
        :param pulumi.Input[dict] upgrade_settings: Specify node upgrade settings to change how many nodes GKE attempts to
               upgrade at once. The number of nodes upgraded simultaneously is the sum of `max_surge` and `max_unavailable`.
               The maximum number of nodes upgraded simultaneously is limited to 20.
        :param pulumi.Input[str] version: The Kubernetes version for the nodes in this pool. Note that if this field
               and `auto_upgrade` are both specified, they will fight each other for what the node version should
               be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
               recommended that you specify explicit versions as the provider will see spurious diffs
               when fuzzy versions are used. See the `container.getEngineVersions` data source's
               `version_prefix` field to approximate fuzzy versions in a provider-compatible way.

        The **autoscaling** object supports the following:

          * `maxNodeCount` (`pulumi.Input[float]`) - Maximum number of nodes in the NodePool. Must be >= min_node_count.
          * `minNodeCount` (`pulumi.Input[float]`) - Minimum number of nodes in the NodePool. Must be >=0 and
            <= `max_node_count`.

        The **management** object supports the following:

          * `autoRepair` (`pulumi.Input[bool]`) - Whether the nodes will be automatically repaired.
          * `autoUpgrade` (`pulumi.Input[bool]`) - Whether the nodes will be automatically upgraded.

        The **node_config** object supports the following:

          * `bootDiskKmsKey` (`pulumi.Input[str]`)
          * `disk_size_gb` (`pulumi.Input[float]`)
          * `diskType` (`pulumi.Input[str]`)
          * `guest_accelerators` (`pulumi.Input[list]`)
            * `count` (`pulumi.Input[float]`)
            * `type` (`pulumi.Input[str]`)

          * `imageType` (`pulumi.Input[str]`)
          * `labels` (`pulumi.Input[dict]`)
          * `localSsdCount` (`pulumi.Input[float]`)
          * `machine_type` (`pulumi.Input[str]`)
          * `metadata` (`pulumi.Input[dict]`)
          * `min_cpu_platform` (`pulumi.Input[str]`)
          * `oauthScopes` (`pulumi.Input[list]`)
          * `preemptible` (`pulumi.Input[bool]`)
          * `sandboxConfig` (`pulumi.Input[dict]`)
            * `sandboxType` (`pulumi.Input[str]`)

          * `service_account` (`pulumi.Input[str]`)
          * `shielded_instance_config` (`pulumi.Input[dict]`)
            * `enableIntegrityMonitoring` (`pulumi.Input[bool]`)
            * `enableSecureBoot` (`pulumi.Input[bool]`)

          * `tags` (`pulumi.Input[list]`)
          * `taints` (`pulumi.Input[list]`)
            * `effect` (`pulumi.Input[str]`)
            * `key` (`pulumi.Input[str]`)
            * `value` (`pulumi.Input[str]`)

          * `workloadMetadataConfig` (`pulumi.Input[dict]`)
            * `nodeMetadata` (`pulumi.Input[str]`)

        The **upgrade_settings** object supports the following:

          * `maxSurge` (`pulumi.Input[float]`) - The number of additional nodes that can be added to the node pool during
            an upgrade. Increasing `max_surge` raises the number of nodes that can be upgraded simultaneously.
            Can be set to 0 or greater.
          * `maxUnavailable` (`pulumi.Input[float]`) - The number of nodes that can be simultaneously unavailable during
            an upgrade. Increasing `max_unavailable` raises the number of nodes that can be upgraded in
            parallel. Can be set to 0 or greater.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["autoscaling"] = autoscaling
        __props__["cluster"] = cluster
        __props__["initial_node_count"] = initial_node_count
        __props__["instance_group_urls"] = instance_group_urls
        __props__["location"] = location
        __props__["management"] = management
        __props__["max_pods_per_node"] = max_pods_per_node
        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["node_config"] = node_config
        __props__["node_count"] = node_count
        __props__["node_locations"] = node_locations
        __props__["project"] = project
        __props__["upgrade_settings"] = upgrade_settings
        __props__["version"] = version
        return NodePool(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
