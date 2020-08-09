# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['GuestPolicies']


class GuestPolicies(pulumi.CustomResource):
    assignment: pulumi.Output['outputs.GuestPoliciesAssignment'] = pulumi.property("assignment")
    """
    Specifies the VM instances that are assigned to this policy. This allows you to target sets
    or groups of VM instances by different parameters such as labels, names, OS, or zones.
    If left empty, all VM instances underneath this policy are targeted.
    At the same level in the resource hierarchy (that is within a project), the service prevents
    the creation of multiple policies that conflict with each other.
    For more information, see how the service
    [handles assignment conflicts](https://cloud.google.com/compute/docs/os-config-management/create-guest-policy#handle-conflicts).  Structure is documented below.
    """

    create_time: pulumi.Output[str] = pulumi.property("createTime")
    """
    Time this guest policy was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example:
    "2014-10-02T15:01:23.045123456Z".
    """

    description: pulumi.Output[Optional[str]] = pulumi.property("description")
    """
    Description of the guest policy. Length of the description is limited to 1024 characters.
    """

    etag: pulumi.Output[str] = pulumi.property("etag")
    """
    The etag for this guest policy. If this is provided on update, it must match the server's etag.
    """

    guest_policy_id: pulumi.Output[str] = pulumi.property("guestPolicyId")
    """
    The logical name of the guest policy in the project with the following restrictions:
    * Must contain only lowercase letters, numbers, and hyphens.
    * Must start with a letter.
    * Must be between 1-63 characters.
    * Must end with a number or a letter.
    * Must be unique within the project.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    Unique identifier for the recipe. Only one recipe with a given name is installed on an instance.
    Names are also used to identify resources which helps to determine whether guest policies have conflicts.
    This means that requests to create multiple recipes with the same name and version are rejected since they
    could potentially have conflicting assignments.
    """

    package_repositories: pulumi.Output[Optional[List['outputs.GuestPoliciesPackageRepository']]] = pulumi.property("packageRepositories")
    """
    A list of package repositories to configure on the VM instance.
    This is done before any other configs are applied so they can use these repos.
    Package repositories are only configured if the corresponding package manager(s) are available.  Structure is documented below.
    """

    packages: pulumi.Output[Optional[List['outputs.GuestPoliciesPackage']]] = pulumi.property("packages")
    """
    The software packages to be managed by this policy.  Structure is documented below.
    """

    project: pulumi.Output[str] = pulumi.property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """

    recipes: pulumi.Output[Optional[List['outputs.GuestPoliciesRecipe']]] = pulumi.property("recipes")
    """
    A list of Recipes to install on the VM instance.  Structure is documented below.
    """

    update_time: pulumi.Output[str] = pulumi.property("updateTime")
    """
    Last time this guest policy was updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example:
    "2014-10-02T15:01:23.045123456Z".
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assignment: Optional[pulumi.Input[pulumi.InputType['GuestPoliciesAssignmentArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 guest_policy_id: Optional[pulumi.Input[str]] = None,
                 package_repositories: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['GuestPoliciesPackageRepositoryArgs']]]]] = None,
                 packages: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['GuestPoliciesPackageArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 recipes: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['GuestPoliciesRecipeArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a GuestPolicies resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GuestPoliciesAssignmentArgs']] assignment: Specifies the VM instances that are assigned to this policy. This allows you to target sets
               or groups of VM instances by different parameters such as labels, names, OS, or zones.
               If left empty, all VM instances underneath this policy are targeted.
               At the same level in the resource hierarchy (that is within a project), the service prevents
               the creation of multiple policies that conflict with each other.
               For more information, see how the service
               [handles assignment conflicts](https://cloud.google.com/compute/docs/os-config-management/create-guest-policy#handle-conflicts).  Structure is documented below.
        :param pulumi.Input[str] description: Description of the guest policy. Length of the description is limited to 1024 characters.
        :param pulumi.Input[str] etag: The etag for this guest policy. If this is provided on update, it must match the server's etag.
        :param pulumi.Input[str] guest_policy_id: The logical name of the guest policy in the project with the following restrictions:
               * Must contain only lowercase letters, numbers, and hyphens.
               * Must start with a letter.
               * Must be between 1-63 characters.
               * Must end with a number or a letter.
               * Must be unique within the project.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['GuestPoliciesPackageRepositoryArgs']]]] package_repositories: A list of package repositories to configure on the VM instance.
               This is done before any other configs are applied so they can use these repos.
               Package repositories are only configured if the corresponding package manager(s) are available.  Structure is documented below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['GuestPoliciesPackageArgs']]]] packages: The software packages to be managed by this policy.  Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['GuestPoliciesRecipeArgs']]]] recipes: A list of Recipes to install on the VM instance.  Structure is documented below.
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

            if assignment is None:
                raise TypeError("Missing required property 'assignment'")
            __props__['assignment'] = assignment
            __props__['description'] = description
            __props__['etag'] = etag
            if guest_policy_id is None:
                raise TypeError("Missing required property 'guest_policy_id'")
            __props__['guest_policy_id'] = guest_policy_id
            __props__['package_repositories'] = package_repositories
            __props__['packages'] = packages
            __props__['project'] = project
            __props__['recipes'] = recipes
            __props__['create_time'] = None
            __props__['name'] = None
            __props__['update_time'] = None
        super(GuestPolicies, __self__).__init__(
            'gcp:osconfig/guestPolicies:GuestPolicies',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            assignment: Optional[pulumi.Input[pulumi.InputType['GuestPoliciesAssignmentArgs']]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            guest_policy_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            package_repositories: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['GuestPoliciesPackageRepositoryArgs']]]]] = None,
            packages: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['GuestPoliciesPackageArgs']]]]] = None,
            project: Optional[pulumi.Input[str]] = None,
            recipes: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['GuestPoliciesRecipeArgs']]]]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'GuestPolicies':
        """
        Get an existing GuestPolicies resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GuestPoliciesAssignmentArgs']] assignment: Specifies the VM instances that are assigned to this policy. This allows you to target sets
               or groups of VM instances by different parameters such as labels, names, OS, or zones.
               If left empty, all VM instances underneath this policy are targeted.
               At the same level in the resource hierarchy (that is within a project), the service prevents
               the creation of multiple policies that conflict with each other.
               For more information, see how the service
               [handles assignment conflicts](https://cloud.google.com/compute/docs/os-config-management/create-guest-policy#handle-conflicts).  Structure is documented below.
        :param pulumi.Input[str] create_time: Time this guest policy was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example:
               "2014-10-02T15:01:23.045123456Z".
        :param pulumi.Input[str] description: Description of the guest policy. Length of the description is limited to 1024 characters.
        :param pulumi.Input[str] etag: The etag for this guest policy. If this is provided on update, it must match the server's etag.
        :param pulumi.Input[str] guest_policy_id: The logical name of the guest policy in the project with the following restrictions:
               * Must contain only lowercase letters, numbers, and hyphens.
               * Must start with a letter.
               * Must be between 1-63 characters.
               * Must end with a number or a letter.
               * Must be unique within the project.
        :param pulumi.Input[str] name: Unique identifier for the recipe. Only one recipe with a given name is installed on an instance.
               Names are also used to identify resources which helps to determine whether guest policies have conflicts.
               This means that requests to create multiple recipes with the same name and version are rejected since they
               could potentially have conflicting assignments.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['GuestPoliciesPackageRepositoryArgs']]]] package_repositories: A list of package repositories to configure on the VM instance.
               This is done before any other configs are applied so they can use these repos.
               Package repositories are only configured if the corresponding package manager(s) are available.  Structure is documented below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['GuestPoliciesPackageArgs']]]] packages: The software packages to be managed by this policy.  Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['GuestPoliciesRecipeArgs']]]] recipes: A list of Recipes to install on the VM instance.  Structure is documented below.
        :param pulumi.Input[str] update_time: Last time this guest policy was updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example:
               "2014-10-02T15:01:23.045123456Z".
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["assignment"] = assignment
        __props__["create_time"] = create_time
        __props__["description"] = description
        __props__["etag"] = etag
        __props__["guest_policy_id"] = guest_policy_id
        __props__["name"] = name
        __props__["package_repositories"] = package_repositories
        __props__["packages"] = packages
        __props__["project"] = project
        __props__["recipes"] = recipes
        __props__["update_time"] = update_time
        return GuestPolicies(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

