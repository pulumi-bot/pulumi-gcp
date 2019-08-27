# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class OrganizationPolicy(pulumi.CustomResource):
    boolean_policy: pulumi.Output[dict]
    """
    A boolean policy is a constraint that is either enforced or not. Structure is documented below.
    
      * `enforced` (`bool`)
    """
    constraint: pulumi.Output[str]
    """
    The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
    """
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
    """
    folder: pulumi.Output[str]
    """
    The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
    """
    list_policy: pulumi.Output[dict]
    """
    A policy that can define specific values that are allowed or denied for the given constraint. It
    can also be used to allow or deny all values. Structure is documented below.
    
      * `allow` (`dict`)
    
        * `all` (`bool`)
        * `values` (`list`)
    
      * `deny` (`dict`)
    
        * `all` (`bool`)
        * `values` (`list`)
    
      * `inherit_from_parent` (`bool`)
      * `suggested_value` (`str`)
    """
    restore_policy: pulumi.Output[dict]
    """
    A restore policy is a constraint to restore the default policy. Structure is documented below.
    
      * `default` (`bool`)
    """
    update_time: pulumi.Output[str]
    """
    (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
    """
    version: pulumi.Output[float]
    """
    Version of the Policy. Default version is 0.
    """
    def __init__(__self__, resource_name, opts=None, boolean_policy=None, constraint=None, folder=None, list_policy=None, restore_policy=None, version=None, __props__=None, __name__=None, __opts__=None):
        """
        Allows management of Organization policies for a Google Folder. For more information see
        [the official
        documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview) and
        [API](https://cloud.google.com/resource-manager/reference/rest/v1/folders/setOrgPolicy).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] boolean_policy: A boolean policy is a constraint that is either enforced or not. Structure is documented below.
        :param pulumi.Input[str] constraint: The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
        :param pulumi.Input[str] folder: The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
        :param pulumi.Input[dict] list_policy: A policy that can define specific values that are allowed or denied for the given constraint. It
               can also be used to allow or deny all values. Structure is documented below.
        :param pulumi.Input[dict] restore_policy: A restore policy is a constraint to restore the default policy. Structure is documented below.
        :param pulumi.Input[float] version: Version of the Policy. Default version is 0.
        
        The **boolean_policy** object supports the following:
        
          * `enforced` (`pulumi.Input[bool]`)
        
        The **list_policy** object supports the following:
        
          * `allow` (`pulumi.Input[dict]`)
        
            * `all` (`pulumi.Input[bool]`)
            * `values` (`pulumi.Input[list]`)
        
          * `deny` (`pulumi.Input[dict]`)
        
            * `all` (`pulumi.Input[bool]`)
            * `values` (`pulumi.Input[list]`)
        
          * `inherit_from_parent` (`pulumi.Input[bool]`)
          * `suggested_value` (`pulumi.Input[str]`)
        
        The **restore_policy** object supports the following:
        
          * `default` (`pulumi.Input[bool]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/folder_organization_policy.html.markdown.
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

            __props__['boolean_policy'] = boolean_policy
            if constraint is None:
                raise TypeError("Missing required property 'constraint'")
            __props__['constraint'] = constraint
            if folder is None:
                raise TypeError("Missing required property 'folder'")
            __props__['folder'] = folder
            __props__['list_policy'] = list_policy
            __props__['restore_policy'] = restore_policy
            __props__['version'] = version
            __props__['etag'] = None
            __props__['update_time'] = None
        super(OrganizationPolicy, __self__).__init__(
            'gcp:folder/organizationPolicy:OrganizationPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, boolean_policy=None, constraint=None, etag=None, folder=None, list_policy=None, restore_policy=None, update_time=None, version=None):
        """
        Get an existing OrganizationPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] boolean_policy: A boolean policy is a constraint that is either enforced or not. Structure is documented below.
        :param pulumi.Input[str] constraint: The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
        :param pulumi.Input[str] etag: (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
        :param pulumi.Input[str] folder: The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
        :param pulumi.Input[dict] list_policy: A policy that can define specific values that are allowed or denied for the given constraint. It
               can also be used to allow or deny all values. Structure is documented below.
        :param pulumi.Input[dict] restore_policy: A restore policy is a constraint to restore the default policy. Structure is documented below.
        :param pulumi.Input[str] update_time: (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
        :param pulumi.Input[float] version: Version of the Policy. Default version is 0.
        
        The **boolean_policy** object supports the following:
        
          * `enforced` (`pulumi.Input[bool]`)
        
        The **list_policy** object supports the following:
        
          * `allow` (`pulumi.Input[dict]`)
        
            * `all` (`pulumi.Input[bool]`)
            * `values` (`pulumi.Input[list]`)
        
          * `deny` (`pulumi.Input[dict]`)
        
            * `all` (`pulumi.Input[bool]`)
            * `values` (`pulumi.Input[list]`)
        
          * `inherit_from_parent` (`pulumi.Input[bool]`)
          * `suggested_value` (`pulumi.Input[str]`)
        
        The **restore_policy** object supports the following:
        
          * `default` (`pulumi.Input[bool]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/folder_organization_policy.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["boolean_policy"] = boolean_policy
        __props__["constraint"] = constraint
        __props__["etag"] = etag
        __props__["folder"] = folder
        __props__["list_policy"] = list_policy
        __props__["restore_policy"] = restore_policy
        __props__["update_time"] = update_time
        __props__["version"] = version
        return OrganizationPolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

