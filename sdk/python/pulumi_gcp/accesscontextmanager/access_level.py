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

__all__ = ['AccessLevel']


class AccessLevel(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 basic: Optional[pulumi.Input[pulumi.InputType['AccessLevelBasicArgs']]] = None,
                 custom: Optional[pulumi.Input[pulumi.InputType['AccessLevelCustomArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        An AccessLevel is a label that can be applied to requests to GCP services,
        along with a list of requirements necessary for the label to be applied.

        To get more information about AccessLevel, see:

        * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.accessLevels)
        * How-to Guides
            * [Access Policy Quickstart](https://cloud.google.com/access-context-manager/docs/quickstart)

        > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
        you must specify a `billing_project` and set `user_project_override` to true
        in the provider configuration. Otherwise the ACM API will return a 403 error.
        Your account must have the `serviceusage.services.use` permission on the
        `billing_project` you defined.

        ## Example Usage
        ### Access Context Manager Access Level Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        access_policy = gcp.accesscontextmanager.AccessPolicy("access-policy",
            parent="organizations/123456789",
            title="my policy")
        access_level = gcp.accesscontextmanager.AccessLevel("access-level",
            basic=gcp.accesscontextmanager.AccessLevelBasicArgs(
                conditions=[gcp.accesscontextmanager.AccessLevelBasicConditionArgs(
                    device_policy={
                        "osConstraints": [{
                            "osType": "DESKTOP_CHROME_OS",
                        }],
                        "requireScreenLock": True,
                    },
                    regions=[
                        "CH",
                        "IT",
                        "US",
                    ],
                )],
            ),
            parent=access_policy.name.apply(lambda name: f"accessPolicies/{name}"),
            title="chromeos_no_lock")
        ```

        ## Import

        AccessLevel can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:accesscontextmanager/accessLevel:AccessLevel default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AccessLevelBasicArgs']] basic: A set of predefined conditions for the access level and a combining function.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['AccessLevelCustomArgs']] custom: Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
               See CEL spec at: https://github.com/google/cel-spec.
               Structure is documented below.
        :param pulumi.Input[str] description: Description of the expression
        :param pulumi.Input[str] name: Resource name for the Access Level. The short_name component must begin
               with a letter and only include alphanumeric and '_'.
               Format: accessPolicies/{policy_id}/accessLevels/{short_name}
        :param pulumi.Input[str] parent: The AccessPolicy this AccessLevel lives in.
               Format: accessPolicies/{policy_id}
        :param pulumi.Input[str] title: Title for the expression, i.e. a short string describing its purpose.
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

            __props__['basic'] = basic
            __props__['custom'] = custom
            __props__['description'] = description
            __props__['name'] = name
            if parent is None and not opts.urn:
                raise TypeError("Missing required property 'parent'")
            __props__['parent'] = parent
            if title is None and not opts.urn:
                raise TypeError("Missing required property 'title'")
            __props__['title'] = title
        super(AccessLevel, __self__).__init__(
            'gcp:accesscontextmanager/accessLevel:AccessLevel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            basic: Optional[pulumi.Input[pulumi.InputType['AccessLevelBasicArgs']]] = None,
            custom: Optional[pulumi.Input[pulumi.InputType['AccessLevelCustomArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent: Optional[pulumi.Input[str]] = None,
            title: Optional[pulumi.Input[str]] = None) -> 'AccessLevel':
        """
        Get an existing AccessLevel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AccessLevelBasicArgs']] basic: A set of predefined conditions for the access level and a combining function.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['AccessLevelCustomArgs']] custom: Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
               See CEL spec at: https://github.com/google/cel-spec.
               Structure is documented below.
        :param pulumi.Input[str] description: Description of the expression
        :param pulumi.Input[str] name: Resource name for the Access Level. The short_name component must begin
               with a letter and only include alphanumeric and '_'.
               Format: accessPolicies/{policy_id}/accessLevels/{short_name}
        :param pulumi.Input[str] parent: The AccessPolicy this AccessLevel lives in.
               Format: accessPolicies/{policy_id}
        :param pulumi.Input[str] title: Title for the expression, i.e. a short string describing its purpose.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["basic"] = basic
        __props__["custom"] = custom
        __props__["description"] = description
        __props__["name"] = name
        __props__["parent"] = parent
        __props__["title"] = title
        return AccessLevel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def basic(self) -> pulumi.Output[Optional['outputs.AccessLevelBasic']]:
        """
        A set of predefined conditions for the access level and a combining function.
        Structure is documented below.
        """
        return pulumi.get(self, "basic")

    @property
    @pulumi.getter
    def custom(self) -> pulumi.Output[Optional['outputs.AccessLevelCustom']]:
        """
        Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
        See CEL spec at: https://github.com/google/cel-spec.
        Structure is documented below.
        """
        return pulumi.get(self, "custom")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the expression
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name for the Access Level. The short_name component must begin
        with a letter and only include alphanumeric and '_'.
        Format: accessPolicies/{policy_id}/accessLevels/{short_name}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Output[str]:
        """
        The AccessPolicy this AccessLevel lives in.
        Format: accessPolicies/{policy_id}
        """
        return pulumi.get(self, "parent")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[str]:
        """
        Title for the expression, i.e. a short string describing its purpose.
        """
        return pulumi.get(self, "title")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

