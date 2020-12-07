# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['InstanceGroupManager']


class InstanceGroupManager(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_healing_policies: Optional[pulumi.Input[pulumi.InputType['InstanceGroupManagerAutoHealingPoliciesArgs']]] = None,
                 base_instance_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 named_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceGroupManagerNamedPortArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 stateful_disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceGroupManagerStatefulDiskArgs']]]]] = None,
                 target_pools: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_size: Optional[pulumi.Input[int]] = None,
                 update_policy: Optional[pulumi.Input[pulumi.InputType['InstanceGroupManagerUpdatePolicyArgs']]] = None,
                 versions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceGroupManagerVersionArgs']]]]] = None,
                 wait_for_instances: Optional[pulumi.Input[bool]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            http_health_check=gcp.compute.HealthCheckHttpHealthCheckArgs(
                request_path="/healthz",
                port=8080,
            ))
        appserver = gcp.compute.InstanceGroupManager("appserver",
            base_instance_name="app",
            zone="us-central1-a",
            versions=[gcp.compute.InstanceGroupManagerVersionArgs(
                instance_template=google_compute_instance_template["appserver"]["id"],
            )],
            target_pools=[google_compute_target_pool["appserver"]["id"]],
            target_size=2,
            named_ports=[gcp.compute.InstanceGroupManagerNamedPortArgs(
                name="customHTTP",
                port=8888,
            )],
            auto_healing_policies=gcp.compute.InstanceGroupManagerAutoHealingPoliciesArgs(
                health_check=autohealing.id,
                initial_delay_sec=300,
            ))
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
                gcp.compute.InstanceGroupManagerVersionArgs(
                    name="appserver",
                    instance_template=google_compute_instance_template["appserver"]["id"],
                ),
                gcp.compute.InstanceGroupManagerVersionArgs(
                    name="appserver-canary",
                    instance_template=google_compute_instance_template["appserver-canary"]["id"],
                    target_size={
                        "fixed": 1,
                    },
                ),
            ],
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        Instance group managers can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager appserver projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager appserver {{project}}/{{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager appserver {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager appserver {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['InstanceGroupManagerAutoHealingPoliciesArgs']] auto_healing_policies: The autohealing policies for this managed instance
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
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceGroupManagerNamedPortArgs']]]] named_ports: The named port configuration. See the section below
               for details on configuration.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceGroupManagerStatefulDiskArgs']]]] stateful_disks: Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_pools: The full URL of all target pools to which new
               instances in the group are added. Updating the target pools attribute does
               not affect existing instances.
        :param pulumi.Input[int] target_size: - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
        :param pulumi.Input[pulumi.InputType['InstanceGroupManagerUpdatePolicyArgs']] update_policy: The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceGroupManagerVersionArgs']]]] versions: Application versions managed by this instance group. Each
               version deals with a specific instance template, allowing canary release scenarios.
               Structure is documented below.
        :param pulumi.Input[bool] wait_for_instances: Whether to wait for all instances to be created/updated before
               returning. Note that if this is set to true and the operation does not succeed, this provider will
               continue trying until it times out.
        :param pulumi.Input[str] zone: The zone that instances in this group should be created
               in.
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
            if base_instance_name is None and not opts.urn:
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
            if versions is None and not opts.urn:
                raise TypeError("Missing required property 'versions'")
            __props__['versions'] = versions
            __props__['wait_for_instances'] = wait_for_instances
            __props__['zone'] = zone
            __props__['fingerprint'] = None
            __props__['instance_group'] = None
            __props__['operation'] = None
            __props__['self_link'] = None
        super(InstanceGroupManager, __self__).__init__(
            'gcp:compute/instanceGroupManager:InstanceGroupManager',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_healing_policies: Optional[pulumi.Input[pulumi.InputType['InstanceGroupManagerAutoHealingPoliciesArgs']]] = None,
            base_instance_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            fingerprint: Optional[pulumi.Input[str]] = None,
            instance_group: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            named_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceGroupManagerNamedPortArgs']]]]] = None,
            operation: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            stateful_disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceGroupManagerStatefulDiskArgs']]]]] = None,
            target_pools: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            target_size: Optional[pulumi.Input[int]] = None,
            update_policy: Optional[pulumi.Input[pulumi.InputType['InstanceGroupManagerUpdatePolicyArgs']]] = None,
            versions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceGroupManagerVersionArgs']]]]] = None,
            wait_for_instances: Optional[pulumi.Input[bool]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'InstanceGroupManager':
        """
        Get an existing InstanceGroupManager resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['InstanceGroupManagerAutoHealingPoliciesArgs']] auto_healing_policies: The autohealing policies for this managed instance
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
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceGroupManagerNamedPortArgs']]]] named_ports: The named port configuration. See the section below
               for details on configuration.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URL of the created resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceGroupManagerStatefulDiskArgs']]]] stateful_disks: Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_pools: The full URL of all target pools to which new
               instances in the group are added. Updating the target pools attribute does
               not affect existing instances.
        :param pulumi.Input[int] target_size: - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
        :param pulumi.Input[pulumi.InputType['InstanceGroupManagerUpdatePolicyArgs']] update_policy: The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceGroupManagerVersionArgs']]]] versions: Application versions managed by this instance group. Each
               version deals with a specific instance template, allowing canary release scenarios.
               Structure is documented below.
        :param pulumi.Input[bool] wait_for_instances: Whether to wait for all instances to be created/updated before
               returning. Note that if this is set to true and the operation does not succeed, this provider will
               continue trying until it times out.
        :param pulumi.Input[str] zone: The zone that instances in this group should be created
               in.
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
        __props__["operation"] = operation
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

    @property
    @pulumi.getter(name="autoHealingPolicies")
    def auto_healing_policies(self) -> pulumi.Output[Optional['outputs.InstanceGroupManagerAutoHealingPolicies']]:
        """
        The autohealing policies for this managed instance
        group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
        """
        return pulumi.get(self, "auto_healing_policies")

    @property
    @pulumi.getter(name="baseInstanceName")
    def base_instance_name(self) -> pulumi.Output[str]:
        """
        The base instance name to use for
        instances in this group. The value must be a valid
        [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
        are lowercase letters, numbers, and hyphens (-). Instances are named by
        appending a hyphen and a random four-character string to the base instance
        name.
        """
        return pulumi.get(self, "base_instance_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional textual description of the instance
        group manager.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        """
        The fingerprint of the instance group manager.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter(name="instanceGroup")
    def instance_group(self) -> pulumi.Output[str]:
        """
        The full URL of the instance group created by the manager.
        """
        return pulumi.get(self, "instance_group")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        - Version name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namedPorts")
    def named_ports(self) -> pulumi.Output[Optional[Sequence['outputs.InstanceGroupManagerNamedPort']]]:
        """
        The named port configuration. See the section below
        for details on configuration.
        """
        return pulumi.get(self, "named_ports")

    @property
    @pulumi.getter
    def operation(self) -> pulumi.Output[str]:
        return pulumi.get(self, "operation")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URL of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="statefulDisks")
    def stateful_disks(self) -> pulumi.Output[Optional[Sequence['outputs.InstanceGroupManagerStatefulDisk']]]:
        """
        Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
        """
        return pulumi.get(self, "stateful_disks")

    @property
    @pulumi.getter(name="targetPools")
    def target_pools(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The full URL of all target pools to which new
        instances in the group are added. Updating the target pools attribute does
        not affect existing instances.
        """
        return pulumi.get(self, "target_pools")

    @property
    @pulumi.getter(name="targetSize")
    def target_size(self) -> pulumi.Output[int]:
        """
        - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
        """
        return pulumi.get(self, "target_size")

    @property
    @pulumi.getter(name="updatePolicy")
    def update_policy(self) -> pulumi.Output['outputs.InstanceGroupManagerUpdatePolicy']:
        """
        The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
        """
        return pulumi.get(self, "update_policy")

    @property
    @pulumi.getter
    def versions(self) -> pulumi.Output[Sequence['outputs.InstanceGroupManagerVersion']]:
        """
        Application versions managed by this instance group. Each
        version deals with a specific instance template, allowing canary release scenarios.
        Structure is documented below.
        """
        return pulumi.get(self, "versions")

    @property
    @pulumi.getter(name="waitForInstances")
    def wait_for_instances(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to wait for all instances to be created/updated before
        returning. Note that if this is set to true and the operation does not succeed, this provider will
        continue trying until it times out.
        """
        return pulumi.get(self, "wait_for_instances")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        The zone that instances in this group should be created
        in.
        """
        return pulumi.get(self, "zone")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

