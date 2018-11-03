# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class RegionInstanceGroupManager(pulumi.CustomResource):
    """
    The Google Compute Engine Regional Instance Group Manager API creates and manages pools
    of homogeneous Compute Engine virtual machine instances from a common instance
    template. For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups)
    and [API](https://cloud.google.com/compute/docs/reference/latest/regionInstanceGroupManagers)
    
    ~> **Note:** Use [google_compute_instance_group_manager](https://www.terraform.io/docs/providers/google/r/compute_instance_group_manager.html) to create a single-zone instance group manager.
    """
    def __init__(__self__, __name__, __opts__=None, auto_healing_policies=None, base_instance_name=None, description=None, distribution_policy_zones=None, instance_template=None, name=None, named_ports=None, project=None, region=None, rolling_update_policy=None, target_pools=None, target_size=None, update_strategy=None, versions=None, wait_for_instances=None):
        """Create a RegionInstanceGroupManager resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['autoHealingPolicies'] = auto_healing_policies

        if not base_instance_name:
            raise TypeError('Missing required property base_instance_name')
        __props__['baseInstanceName'] = base_instance_name

        __props__['description'] = description

        __props__['distributionPolicyZones'] = distribution_policy_zones

        __props__['instanceTemplate'] = instance_template

        __props__['name'] = name

        __props__['namedPorts'] = named_ports

        __props__['project'] = project

        if not region:
            raise TypeError('Missing required property region')
        __props__['region'] = region

        __props__['rollingUpdatePolicy'] = rolling_update_policy

        __props__['targetPools'] = target_pools

        __props__['targetSize'] = target_size

        __props__['updateStrategy'] = update_strategy

        __props__['versions'] = versions

        __props__['waitForInstances'] = wait_for_instances

        __props__['fingerprint'] = None
        __props__['instance_group'] = None
        __props__['self_link'] = None

        super(RegionInstanceGroupManager, __self__).__init__(
            'gcp:compute/regionInstanceGroupManager:RegionInstanceGroupManager',
            __name__,
            __props__,
            __opts__)

