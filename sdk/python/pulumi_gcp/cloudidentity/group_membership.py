# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GroupMembership(pulumi.CustomResource):
    create_time: pulumi.Output[str]
    """
    The time when the Membership was created.
    """
    group: pulumi.Output[str]
    """
    The name of the Group to create this membership in.
    """
    member_key: pulumi.Output[dict]
    """
    EntityKey of the member.
    Structure is documented below.

      * `id` (`str`) - The ID of the entity.
        For Google-managed entities, the id must be the email address of an existing
        group or user.
        For external-identity-mapped entities, the id must be a string conforming
        to the Identity Source's requirements.
        Must be unique within a namespace.
      * `namespace` (`str`) - The namespace in which the entity exists.
        If not specified, the EntityKey represents a Google-managed entity
        such as a Google user or a Google Group.
        If specified, the EntityKey represents an external-identity-mapped group.
        The namespace must correspond to an identity source created in Admin Console
        and must be in the form of `identitysources/{identity_source_id}`.
    """
    name: pulumi.Output[str]
    """
    The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER.
    Possible values are `OWNER`, `MANAGER`, and `MEMBER`.
    """
    preferred_member_key: pulumi.Output[dict]
    """
    EntityKey of the member.
    Structure is documented below.

      * `id` (`str`) - The ID of the entity.
        For Google-managed entities, the id must be the email address of an existing
        group or user.
        For external-identity-mapped entities, the id must be a string conforming
        to the Identity Source's requirements.
        Must be unique within a namespace.
      * `namespace` (`str`) - The namespace in which the entity exists.
        If not specified, the EntityKey represents a Google-managed entity
        such as a Google user or a Google Group.
        If specified, the EntityKey represents an external-identity-mapped group.
        The namespace must correspond to an identity source created in Admin Console
        and must be in the form of `identitysources/{identity_source_id}`.
    """
    roles: pulumi.Output[list]
    """
    The MembershipRoles that apply to the Membership.
    Must not contain duplicate MembershipRoles with the same name.
    Structure is documented below.

      * `name` (`str`) - The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER.
        Possible values are `OWNER`, `MANAGER`, and `MEMBER`.
    """
    type: pulumi.Output[str]
    """
    The type of the membership.
    """
    update_time: pulumi.Output[str]
    """
    The time when the Membership was last updated.
    """
    def __init__(__self__, resource_name, opts=None, group=None, member_key=None, preferred_member_key=None, roles=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a GroupMembership resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group: The name of the Group to create this membership in.
        :param pulumi.Input[dict] member_key: EntityKey of the member.
               Structure is documented below.
        :param pulumi.Input[dict] preferred_member_key: EntityKey of the member.
               Structure is documented below.
        :param pulumi.Input[list] roles: The MembershipRoles that apply to the Membership.
               Must not contain duplicate MembershipRoles with the same name.
               Structure is documented below.

        The **member_key** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ID of the entity.
            For Google-managed entities, the id must be the email address of an existing
            group or user.
            For external-identity-mapped entities, the id must be a string conforming
            to the Identity Source's requirements.
            Must be unique within a namespace.
          * `namespace` (`pulumi.Input[str]`) - The namespace in which the entity exists.
            If not specified, the EntityKey represents a Google-managed entity
            such as a Google user or a Google Group.
            If specified, the EntityKey represents an external-identity-mapped group.
            The namespace must correspond to an identity source created in Admin Console
            and must be in the form of `identitysources/{identity_source_id}`.

        The **preferred_member_key** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ID of the entity.
            For Google-managed entities, the id must be the email address of an existing
            group or user.
            For external-identity-mapped entities, the id must be a string conforming
            to the Identity Source's requirements.
            Must be unique within a namespace.
          * `namespace` (`pulumi.Input[str]`) - The namespace in which the entity exists.
            If not specified, the EntityKey represents a Google-managed entity
            such as a Google user or a Google Group.
            If specified, the EntityKey represents an external-identity-mapped group.
            The namespace must correspond to an identity source created in Admin Console
            and must be in the form of `identitysources/{identity_source_id}`.

        The **roles** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER.
            Possible values are `OWNER`, `MANAGER`, and `MEMBER`.
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

            if group is None:
                raise TypeError("Missing required property 'group'")
            __props__['group'] = group
            __props__['member_key'] = member_key
            __props__['preferred_member_key'] = preferred_member_key
            if roles is None:
                raise TypeError("Missing required property 'roles'")
            __props__['roles'] = roles
            __props__['create_time'] = None
            __props__['name'] = None
            __props__['type'] = None
            __props__['update_time'] = None
        super(GroupMembership, __self__).__init__(
            'gcp:cloudidentity/groupMembership:GroupMembership',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, create_time=None, group=None, member_key=None, name=None, preferred_member_key=None, roles=None, type=None, update_time=None):
        """
        Get an existing GroupMembership resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: The time when the Membership was created.
        :param pulumi.Input[str] group: The name of the Group to create this membership in.
        :param pulumi.Input[dict] member_key: EntityKey of the member.
               Structure is documented below.
        :param pulumi.Input[str] name: The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER.
               Possible values are `OWNER`, `MANAGER`, and `MEMBER`.
        :param pulumi.Input[dict] preferred_member_key: EntityKey of the member.
               Structure is documented below.
        :param pulumi.Input[list] roles: The MembershipRoles that apply to the Membership.
               Must not contain duplicate MembershipRoles with the same name.
               Structure is documented below.
        :param pulumi.Input[str] type: The type of the membership.
        :param pulumi.Input[str] update_time: The time when the Membership was last updated.

        The **member_key** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ID of the entity.
            For Google-managed entities, the id must be the email address of an existing
            group or user.
            For external-identity-mapped entities, the id must be a string conforming
            to the Identity Source's requirements.
            Must be unique within a namespace.
          * `namespace` (`pulumi.Input[str]`) - The namespace in which the entity exists.
            If not specified, the EntityKey represents a Google-managed entity
            such as a Google user or a Google Group.
            If specified, the EntityKey represents an external-identity-mapped group.
            The namespace must correspond to an identity source created in Admin Console
            and must be in the form of `identitysources/{identity_source_id}`.

        The **preferred_member_key** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ID of the entity.
            For Google-managed entities, the id must be the email address of an existing
            group or user.
            For external-identity-mapped entities, the id must be a string conforming
            to the Identity Source's requirements.
            Must be unique within a namespace.
          * `namespace` (`pulumi.Input[str]`) - The namespace in which the entity exists.
            If not specified, the EntityKey represents a Google-managed entity
            such as a Google user or a Google Group.
            If specified, the EntityKey represents an external-identity-mapped group.
            The namespace must correspond to an identity source created in Admin Console
            and must be in the form of `identitysources/{identity_source_id}`.

        The **roles** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER.
            Possible values are `OWNER`, `MANAGER`, and `MEMBER`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["create_time"] = create_time
        __props__["group"] = group
        __props__["member_key"] = member_key
        __props__["name"] = name
        __props__["preferred_member_key"] = preferred_member_key
        __props__["roles"] = roles
        __props__["type"] = type
        __props__["update_time"] = update_time
        return GroupMembership(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
