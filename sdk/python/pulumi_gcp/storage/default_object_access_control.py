# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class DefaultObjectAccessControl(pulumi.CustomResource):
    bucket: pulumi.Output[str]
    """
    -
    (Required)
    The name of the bucket.
    """
    domain: pulumi.Output[str]
    """
    The domain associated with the entity.
    """
    email: pulumi.Output[str]
    """
    The email address associated with the entity.
    """
    entity: pulumi.Output[str]
    """
    -
    (Required)
    The entity holding the permission, in one of the following forms:
    * user-{{userId}}
    * user-{{email}} (such as "user-liz@example.com")
    * group-{{groupId}}
    * group-{{email}} (such as "group-example@googlegroups.com")
    * domain-{{domain}} (such as "domain-example.com")
    * project-team-{{projectId}}
    * allUsers
    * allAuthenticatedUsers
    """
    entity_id: pulumi.Output[str]
    """
    The ID for the entity
    """
    generation: pulumi.Output[float]
    """
    The content generation of the object, if applied to an object.
    """
    object: pulumi.Output[str]
    """
    -
    (Optional)
    The name of the object, if applied to an object.
    """
    project_team: pulumi.Output[dict]
    """
    The project team associated with the entity

      * `project_number` (`str`)
      * `team` (`str`)
    """
    role: pulumi.Output[str]
    """
    -
    (Required)
    The access permission for the entity.
    """
    def __init__(__self__, resource_name, opts=None, bucket=None, entity=None, object=None, role=None, __props__=None, __name__=None, __opts__=None):
        """
        The DefaultObjectAccessControls resources represent the Access Control
        Lists (ACLs) applied to a new object within a Google Cloud Storage bucket
        when no ACL was provided for that object. ACLs let you specify who has
        access to your bucket contents and to what extent.

        There are two roles that can be assigned to an entity:

        READERs can get an object, though the acl property will not be revealed.
        OWNERs are READERs, and they can get the acl property, update an object,
        and call all objectAccessControls methods on the object. The owner of an
        object is always an OWNER.
        For more information, see Access Control, with the caveat that this API
        uses READER and OWNER instead of READ and FULL_CONTROL.


        To get more information about DefaultObjectAccessControl, see:

        * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/defaultObjectAccessControls)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/storage/docs/access-control/create-manage-lists)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: -
               (Required)
               The name of the bucket.
        :param pulumi.Input[str] entity: -
               (Required)
               The entity holding the permission, in one of the following forms:
               * user-{{userId}}
               * user-{{email}} (such as "user-liz@example.com")
               * group-{{groupId}}
               * group-{{email}} (such as "group-example@googlegroups.com")
               * domain-{{domain}} (such as "domain-example.com")
               * project-team-{{projectId}}
               * allUsers
               * allAuthenticatedUsers
        :param pulumi.Input[str] object: -
               (Optional)
               The name of the object, if applied to an object.
        :param pulumi.Input[str] role: -
               (Required)
               The access permission for the entity.
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

            if bucket is None:
                raise TypeError("Missing required property 'bucket'")
            __props__['bucket'] = bucket
            if entity is None:
                raise TypeError("Missing required property 'entity'")
            __props__['entity'] = entity
            __props__['object'] = object
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['domain'] = None
            __props__['email'] = None
            __props__['entity_id'] = None
            __props__['generation'] = None
            __props__['project_team'] = None
        super(DefaultObjectAccessControl, __self__).__init__(
            'gcp:storage/defaultObjectAccessControl:DefaultObjectAccessControl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, bucket=None, domain=None, email=None, entity=None, entity_id=None, generation=None, object=None, project_team=None, role=None):
        """
        Get an existing DefaultObjectAccessControl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: -
               (Required)
               The name of the bucket.
        :param pulumi.Input[str] domain: The domain associated with the entity.
        :param pulumi.Input[str] email: The email address associated with the entity.
        :param pulumi.Input[str] entity: -
               (Required)
               The entity holding the permission, in one of the following forms:
               * user-{{userId}}
               * user-{{email}} (such as "user-liz@example.com")
               * group-{{groupId}}
               * group-{{email}} (such as "group-example@googlegroups.com")
               * domain-{{domain}} (such as "domain-example.com")
               * project-team-{{projectId}}
               * allUsers
               * allAuthenticatedUsers
        :param pulumi.Input[str] entity_id: The ID for the entity
        :param pulumi.Input[float] generation: The content generation of the object, if applied to an object.
        :param pulumi.Input[str] object: -
               (Optional)
               The name of the object, if applied to an object.
        :param pulumi.Input[dict] project_team: The project team associated with the entity
        :param pulumi.Input[str] role: -
               (Required)
               The access permission for the entity.

        The **project_team** object supports the following:

          * `project_number` (`pulumi.Input[str]`)
          * `team` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bucket"] = bucket
        __props__["domain"] = domain
        __props__["email"] = email
        __props__["entity"] = entity
        __props__["entity_id"] = entity_id
        __props__["generation"] = generation
        __props__["object"] = object
        __props__["project_team"] = project_team
        __props__["role"] = role
        return DefaultObjectAccessControl(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

