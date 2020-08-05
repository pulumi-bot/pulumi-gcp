# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class InstanceGroupManager(pulumi.CustomResource):
    auto_healing_policies: pulumi.Output[dict]
    """
    The autohealing policies for this managed instance
    group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).

      * `healthCheck` (`str`) - The health check resource that signals autohealing.
      * `initialDelaySec` (`float`) - The number of seconds that the managed instance group waits before
        it applies autohealing policies to new instances or recently recreated instances. Between 0 and 3600.
    """
    base_instance_name: pulumi.Output[str]
    """
    The base instance name to use for
    instances in this group. The value must be a valid
    [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
    are lowercase letters, numbers, and hyphens (-). Instances are named by
    appending a hyphen and a random four-character string to the base instance
    name.
    """
    description: pulumi.Output[str]
    """
    An optional textual description of the instance
    group manager.
    """
    fingerprint: pulumi.Output[str]
    """
    The fingerprint of the instance group manager.
    """
    instance_group: pulumi.Output[str]
    """
    The full URL of the instance group created by the manager.
    """
    name: pulumi.Output[str]
    """
    - Version name.
    """
    named_ports: pulumi.Output[list]
    """
    The named port configuration. See the section below
    for details on configuration.

      * `name` (`str`) - - Version name.
      * `port` (`float`) - The port number.
        - - -
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    self_link: pulumi.Output[str]
    """
    The URL of the created resource.
    """
    stateful_disks: pulumi.Output[list]
    """
    Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).

      * `deleteRule` (`str`) - , A value that prescribes what should happen to the stateful disk when the VM instance is deleted. The available options are `NEVER` and `ON_PERMANENT_INSTANCE_DELETION`. `NEVER` detatch the disk when the VM is deleted, but not delete the disk. `ON_PERMANENT_INSTANCE_DELETION` will delete the stateful disk when the VM is permanently deleted from the instance group. The default is `NEVER`.
      * `device_name` (`str`) - , The device name of the disk to be attached.
    """
    target_pools: pulumi.Output[list]
    """
    The full URL of all target pools to which new
    instances in the group are added. Updating the target pools attribute does
    not affect existing instances.
    """
    target_size: pulumi.Output[float]
    """
    - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
    """
    update_policy: pulumi.Output[dict]
    """
    The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/instanceGroupManagers/patch)

      * `maxSurgeFixed` (`float`) - , The maximum number of instances that can be created above the specified targetSize during the update process. Conflicts with `max_surge_percent`. If neither is set, defaults to 1
      * `maxSurgePercent` (`float`) - , The maximum number of instances(calculated as percentage) that can be created above the specified targetSize during the update process. Conflicts with `max_surge_fixed`.
      * `maxUnavailableFixed` (`float`) - , The maximum number of instances that can be unavailable during the update process. Conflicts with `max_unavailable_percent`. If neither is set, defaults to 1
      * `maxUnavailablePercent` (`float`) - , The maximum number of instances(calculated as percentage) that can be unavailable during the update process. Conflicts with `max_unavailable_fixed`.
      * `minReadySec` (`float`) - , Minimum number of seconds to wait for after a newly created instance becomes available. This value must be from range [0, 3600]
        - - -
      * `minimal_action` (`str`) - - Minimal action to be taken on an instance. You can specify either `RESTART` to restart existing instances or `REPLACE` to delete and create new instances from the target template. If you specify a `RESTART`, the Updater will attempt to perform that action only. However, if the Updater determines that the minimal action you specify is not enough to perform the update, it might perform a more disruptive action.
      * `type` (`str`) - - The type of update process. You can specify either `PROACTIVE` so that the instance group manager proactively executes actions in order to bring instances to their target versions or `OPPORTUNISTIC` so that no action is proactively executed but the update will be performed as part of other actions (for example, resizes or recreateInstances calls).
    """
    versions: pulumi.Output[list]
    """
    Application versions managed by this instance group. Each
    version deals with a specific instance template, allowing canary release scenarios.
    Structure is documented below.

      * `instanceTemplate` (`str`) - - The full URL to an instance template from which all new instances of this version will be created.
      * `name` (`str`) - - Version name.
      * `target_size` (`dict`) - - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
        * `fixed` (`float`) - , The number of instances which are managed for this version. Conflicts with `percent`.
        * `percent` (`float`) - , The number of instances (calculated as percentage) which are managed for this version. Conflicts with `fixed`.
          Note that when using `percent`, rounding will be in favor of explicitly set `target_size` values; a managed instance group with 2 instances and 2 `version`s,
          one of which has a `target_size.percent` of `60` will create 2 instances of that `version`.
    """
    wait_for_instances: pulumi.Output[bool]
    """
    Whether to wait for all instances to be created/updated before
    returning. Note that if this is set to true and the operation does not succeed, this provider will
    continue trying until it times out.
    """
    zone: pulumi.Output[str]
    """
    The zone that instances in this group should be created
    in.
    """
    def __init__(__self__, resource_name, opts=None, auto_healing_policies=None, base_instance_name=None, description=None, name=None, named_ports=None, project=None, stateful_disks=None, target_pools=None, target_size=None, update_policy=None, versions=None, wait_for_instances=None, zone=None, __props__=None, __name__=None, __opts__=None):
        """
        The Google Compute Engine Instance Group Manager API creates and manages pools
        of homogeneous Compute Engine virtual machine instances from a common instance
        template. For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/manager)
        and [API](https://cloud.google.com/compute/docs/reference/latest/instanceGroupManagers)

        > **Note:** Use [compute.RegionInstanceGroupManager](https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager.html) to create a regional (multi-zone) instance group manager.

        ## Example Usage
        ### With Top Level Instance Template (`Google` Provider)

        ```python
        import pulumi
        import pulumi_gcp as gcp

        autohealing = gcp.compute.HealthCheck("autohealing",
            check_interval_sec=5,
            timeout_sec=5,
            healthy_threshold=2,
            unhealthy_threshold=10,
            http_health_check={
                "request_path": "/healthz",
                "port": 8080,
            })
        appserver = gcp.compute.InstanceGroupManager("appserver",
            base_instance_name="app",
            zone="us-central1-a",
            versions=[{
                "instanceTemplate": google_compute_instance_template["appserver"]["id"],
            }],
            target_pools=[google_compute_target_pool["appserver"]["id"]],
            target_size=2,
            named_ports=[{
                "name": "customHTTP",
                "port": 8888,
            }],
            auto_healing_policies={
                "healthCheck": autohealing.id,
                "initialDelaySec": 300,
            })
        ```
        ### With Multiple Versions (`Google-Beta` Provider)
        ```python
        import pulumi
        import pulumi_gcp as gcp

        appserver = gcp.compute.InstanceGroupManager("appserver",
            base_instance_name="app",
            zone="us-central1-a",
            target_size=5,
            versions=[
                {
                    "name": "appserver",
                    "instanceTemplate": google_compute_instance_template["appserver"]["id"],
                },
                {
                    "name": "appserver-canary",
                    "instanceTemplate": google_compute_instance_template["appserver-canary"]["id"],
                    "target_size": {
                        "fixed": 1,
                    },
                },
            ],
            opts=ResourceOptions(provider=google_beta))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] auto_healing_policies: The autohealing policies for this managed instance
               group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
        :param pulumi.Input[str] base_instance_name: The base instance name to use for
               instances in this group. The value must be a valid
               [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
               are lowercase letters, numbers, and hyphens (-). Instances are named by
               appending a hyphen and a random four-character string to the base instance
               name.
        :param pulumi.Input[str] description: An optional textual description of the instance
               group manager.
        :param pulumi.Input[str] name: - Version name.
        :param pulumi.Input[list] named_ports: The named port configuration. See the section below
               for details on configuration.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[list] stateful_disks: Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
        :param pulumi.Input[list] target_pools: The full URL of all target pools to which new
               instances in the group are added. Updating the target pools attribute does
               not affect existing instances.
        :param pulumi.Input[float] target_size: - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
        :param pulumi.Input[dict] update_policy: The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/instanceGroupManagers/patch)
        :param pulumi.Input[list] versions: Application versions managed by this instance group. Each
               version deals with a specific instance template, allowing canary release scenarios.
               Structure is documented below.
        :param pulumi.Input[bool] wait_for_instances: Whether to wait for all instances to be created/updated before
               returning. Note that if this is set to true and the operation does not succeed, this provider will
               continue trying until it times out.
        :param pulumi.Input[str] zone: The zone that instances in this group should be created
               in.

        The **auto_healing_policies** object supports the following:

          * `healthCheck` (`pulumi.Input[str]`) - The health check resource that signals autohealing.
          * `initialDelaySec` (`pulumi.Input[float]`) - The number of seconds that the managed instance group waits before
            it applies autohealing policies to new instances or recently recreated instances. Between 0 and 3600.

        The **named_ports** object supports the following:

          * `name` (`pulumi.Input[str]`) - - Version name.
          * `port` (`pulumi.Input[float]`) - The port number.
            - - -

        The **stateful_disks** object supports the following:

          * `deleteRule` (`pulumi.Input[str]`) - , A value that prescribes what should happen to the stateful disk when the VM instance is deleted. The available options are `NEVER` and `ON_PERMANENT_INSTANCE_DELETION`. `NEVER` detatch the disk when the VM is deleted, but not delete the disk. `ON_PERMANENT_INSTANCE_DELETION` will delete the stateful disk when the VM is permanently deleted from the instance group. The default is `NEVER`.
          * `device_name` (`pulumi.Input[str]`) - , The device name of the disk to be attached.

        The **update_policy** object supports the following:

          * `maxSurgeFixed` (`pulumi.Input[float]`) - , The maximum number of instances that can be created above the specified targetSize during the update process. Conflicts with `max_surge_percent`. If neither is set, defaults to 1
          * `maxSurgePercent` (`pulumi.Input[float]`) - , The maximum number of instances(calculated as percentage) that can be created above the specified targetSize during the update process. Conflicts with `max_surge_fixed`.
          * `maxUnavailableFixed` (`pulumi.Input[float]`) - , The maximum number of instances that can be unavailable during the update process. Conflicts with `max_unavailable_percent`. If neither is set, defaults to 1
          * `maxUnavailablePercent` (`pulumi.Input[float]`) - , The maximum number of instances(calculated as percentage) that can be unavailable during the update process. Conflicts with `max_unavailable_fixed`.
          * `minReadySec` (`pulumi.Input[float]`) - , Minimum number of seconds to wait for after a newly created instance becomes available. This value must be from range [0, 3600]
            - - -
          * `minimal_action` (`pulumi.Input[str]`) - - Minimal action to be taken on an instance. You can specify either `RESTART` to restart existing instances or `REPLACE` to delete and create new instances from the target template. If you specify a `RESTART`, the Updater will attempt to perform that action only. However, if the Updater determines that the minimal action you specify is not enough to perform the update, it might perform a more disruptive action.
          * `type` (`pulumi.Input[str]`) - - The type of update process. You can specify either `PROACTIVE` so that the instance group manager proactively executes actions in order to bring instances to their target versions or `OPPORTUNISTIC` so that no action is proactively executed but the update will be performed as part of other actions (for example, resizes or recreateInstances calls).

        The **versions** object supports the following:

          * `instanceTemplate` (`pulumi.Input[str]`) - - The full URL to an instance template from which all new instances of this version will be created.
          * `name` (`pulumi.Input[str]`) - - Version name.
          * `target_size` (`pulumi.Input[dict]`) - - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
            * `fixed` (`pulumi.Input[float]`) - , The number of instances which are managed for this version. Conflicts with `percent`.
            * `percent` (`pulumi.Input[float]`) - , The number of instances (calculated as percentage) which are managed for this version. Conflicts with `fixed`.
              Note that when using `percent`, rounding will be in favor of explicitly set `target_size` values; a managed instance group with 2 instances and 2 `version`s,
              one of which has a `target_size.percent` of `60` will create 2 instances of that `version`.
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

            __props__['auto_healing_policies'] = auto_healing_policies
            if base_instance_name is None:
                raise TypeError("Missing required property 'base_instance_name'")
            __props__['base_instance_name'] = base_instance_name
            __props__['description'] = description
            __props__['name'] = name
            __props__['named_ports'] = named_ports
            __props__['project'] = project
            __props__['stateful_disks'] = stateful_disks
            __props__['target_pools'] = target_pools
            __props__['target_size'] = target_size
            __props__['update_policy'] = update_policy
            if versions is None:
                raise TypeError("Missing required property 'versions'")
            __props__['versions'] = versions
            __props__['wait_for_instances'] = wait_for_instances
            __props__['zone'] = zone
            __props__['fingerprint'] = None
            __props__['instance_group'] = None
            __props__['self_link'] = None
        super(InstanceGroupManager, __self__).__init__(
            'gcp:compute/instanceGroupManager:InstanceGroupManager',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, auto_healing_policies=None, base_instance_name=None, description=None, fingerprint=None, instance_group=None, name=None, named_ports=None, project=None, self_link=None, stateful_disks=None, target_pools=None, target_size=None, update_policy=None, versions=None, wait_for_instances=None, zone=None):
        """
        Get an existing InstanceGroupManager resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] auto_healing_policies: The autohealing policies for this managed instance
               group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
        :param pulumi.Input[str] base_instance_name: The base instance name to use for
               instances in this group. The value must be a valid
               [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
               are lowercase letters, numbers, and hyphens (-). Instances are named by
               appending a hyphen and a random four-character string to the base instance
               name.
        :param pulumi.Input[str] description: An optional textual description of the instance
               group manager.
        :param pulumi.Input[str] fingerprint: The fingerprint of the instance group manager.
        :param pulumi.Input[str] instance_group: The full URL of the instance group created by the manager.
        :param pulumi.Input[str] name: - Version name.
        :param pulumi.Input[list] named_ports: The named port configuration. See the section below
               for details on configuration.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URL of the created resource.
        :param pulumi.Input[list] stateful_disks: Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
        :param pulumi.Input[list] target_pools: The full URL of all target pools to which new
               instances in the group are added. Updating the target pools attribute does
               not affect existing instances.
        :param pulumi.Input[float] target_size: - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
        :param pulumi.Input[dict] update_policy: The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/instanceGroupManagers/patch)
        :param pulumi.Input[list] versions: Application versions managed by this instance group. Each
               version deals with a specific instance template, allowing canary release scenarios.
               Structure is documented below.
        :param pulumi.Input[bool] wait_for_instances: Whether to wait for all instances to be created/updated before
               returning. Note that if this is set to true and the operation does not succeed, this provider will
               continue trying until it times out.
        :param pulumi.Input[str] zone: The zone that instances in this group should be created
               in.

        The **auto_healing_policies** object supports the following:

          * `healthCheck` (`pulumi.Input[str]`) - The health check resource that signals autohealing.
          * `initialDelaySec` (`pulumi.Input[float]`) - The number of seconds that the managed instance group waits before
            it applies autohealing policies to new instances or recently recreated instances. Between 0 and 3600.

        The **named_ports** object supports the following:

          * `name` (`pulumi.Input[str]`) - - Version name.
          * `port` (`pulumi.Input[float]`) - The port number.
            - - -

        The **stateful_disks** object supports the following:

          * `deleteRule` (`pulumi.Input[str]`) - , A value that prescribes what should happen to the stateful disk when the VM instance is deleted. The available options are `NEVER` and `ON_PERMANENT_INSTANCE_DELETION`. `NEVER` detatch the disk when the VM is deleted, but not delete the disk. `ON_PERMANENT_INSTANCE_DELETION` will delete the stateful disk when the VM is permanently deleted from the instance group. The default is `NEVER`.
          * `device_name` (`pulumi.Input[str]`) - , The device name of the disk to be attached.

        The **update_policy** object supports the following:

          * `maxSurgeFixed` (`pulumi.Input[float]`) - , The maximum number of instances that can be created above the specified targetSize during the update process. Conflicts with `max_surge_percent`. If neither is set, defaults to 1
          * `maxSurgePercent` (`pulumi.Input[float]`) - , The maximum number of instances(calculated as percentage) that can be created above the specified targetSize during the update process. Conflicts with `max_surge_fixed`.
          * `maxUnavailableFixed` (`pulumi.Input[float]`) - , The maximum number of instances that can be unavailable during the update process. Conflicts with `max_unavailable_percent`. If neither is set, defaults to 1
          * `maxUnavailablePercent` (`pulumi.Input[float]`) - , The maximum number of instances(calculated as percentage) that can be unavailable during the update process. Conflicts with `max_unavailable_fixed`.
          * `minReadySec` (`pulumi.Input[float]`) - , Minimum number of seconds to wait for after a newly created instance becomes available. This value must be from range [0, 3600]
            - - -
          * `minimal_action` (`pulumi.Input[str]`) - - Minimal action to be taken on an instance. You can specify either `RESTART` to restart existing instances or `REPLACE` to delete and create new instances from the target template. If you specify a `RESTART`, the Updater will attempt to perform that action only. However, if the Updater determines that the minimal action you specify is not enough to perform the update, it might perform a more disruptive action.
          * `type` (`pulumi.Input[str]`) - - The type of update process. You can specify either `PROACTIVE` so that the instance group manager proactively executes actions in order to bring instances to their target versions or `OPPORTUNISTIC` so that no action is proactively executed but the update will be performed as part of other actions (for example, resizes or recreateInstances calls).

        The **versions** object supports the following:

          * `instanceTemplate` (`pulumi.Input[str]`) - - The full URL to an instance template from which all new instances of this version will be created.
          * `name` (`pulumi.Input[str]`) - - Version name.
          * `target_size` (`pulumi.Input[dict]`) - - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
            * `fixed` (`pulumi.Input[float]`) - , The number of instances which are managed for this version. Conflicts with `percent`.
            * `percent` (`pulumi.Input[float]`) - , The number of instances (calculated as percentage) which are managed for this version. Conflicts with `fixed`.
              Note that when using `percent`, rounding will be in favor of explicitly set `target_size` values; a managed instance group with 2 instances and 2 `version`s,
              one of which has a `target_size.percent` of `60` will create 2 instances of that `version`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_healing_policies"] = auto_healing_policies
        __props__["base_instance_name"] = base_instance_name
        __props__["description"] = description
        __props__["fingerprint"] = fingerprint
        __props__["instance_group"] = instance_group
        __props__["name"] = name
        __props__["named_ports"] = named_ports
        __props__["project"] = project
        __props__["self_link"] = self_link
        __props__["stateful_disks"] = stateful_disks
        __props__["target_pools"] = target_pools
        __props__["target_size"] = target_size
        __props__["update_policy"] = update_policy
        __props__["versions"] = versions
        __props__["wait_for_instances"] = wait_for_instances
        __props__["zone"] = zone
        return InstanceGroupManager(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
