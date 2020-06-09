# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Lien(pulumi.CustomResource):
    create_time: pulumi.Output[str]
    """
    Time of creation
    """
    name: pulumi.Output[str]
    """
    A system-generated unique identifier for this Lien.
    """
    origin: pulumi.Output[str]
    """
    A stable, user-visible/meaningful string identifying the origin
    of the Lien, intended to be inspected programmatically. Maximum length of
    200 characters.
    """
    parent: pulumi.Output[str]
    """
    A reference to the resource this Lien is attached to.
    The server will validate the parent against those for which Liens are supported.
    Since a variety of objects can have Liens against them, you must provide the type
    prefix (e.g. "projects/my-project-name").
    """
    reason: pulumi.Output[str]
    """
    Concise user-visible strings indicating why an action cannot be performed
    on a resource. Maximum length of 200 characters.
    """
    restrictions: pulumi.Output[list]
    """
    The types of operations which should be blocked as a result of this Lien.
    Each value should correspond to an IAM permission. The server will validate
    the permissions against those for which Liens are supported.  An empty
    list is meaningless and will be rejected.
    e.g. ['resourcemanager.projects.delete']
    """
    def __init__(__self__, resource_name, opts=None, origin=None, parent=None, reason=None, restrictions=None, __props__=None, __name__=None, __opts__=None):
        """
        A Lien represents an encumbrance on the actions that can be performed on a resource.



        ## Example Usage

        ### Resource Manager Lien

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.organizations.Project("project", project_id="staging-project")
        lien = gcp.resourcemanager.Lien("lien",
            origin="machine-readable-explanation",
            parent=project.number.apply(lambda number: f"projects/{number}"),
            reason="This project is an important environment",
            restrictions=["resourcemanager.projects.delete"])
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] origin: A stable, user-visible/meaningful string identifying the origin
               of the Lien, intended to be inspected programmatically. Maximum length of
               200 characters.
        :param pulumi.Input[str] parent: A reference to the resource this Lien is attached to.
               The server will validate the parent against those for which Liens are supported.
               Since a variety of objects can have Liens against them, you must provide the type
               prefix (e.g. "projects/my-project-name").
        :param pulumi.Input[str] reason: Concise user-visible strings indicating why an action cannot be performed
               on a resource. Maximum length of 200 characters.
        :param pulumi.Input[list] restrictions: The types of operations which should be blocked as a result of this Lien.
               Each value should correspond to an IAM permission. The server will validate
               the permissions against those for which Liens are supported.  An empty
               list is meaningless and will be rejected.
               e.g. ['resourcemanager.projects.delete']
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

            if origin is None:
                raise TypeError("Missing required property 'origin'")
            __props__['origin'] = origin
            if parent is None:
                raise TypeError("Missing required property 'parent'")
            __props__['parent'] = parent
            if reason is None:
                raise TypeError("Missing required property 'reason'")
            __props__['reason'] = reason
            if restrictions is None:
                raise TypeError("Missing required property 'restrictions'")
            __props__['restrictions'] = restrictions
            __props__['create_time'] = None
            __props__['name'] = None
        super(Lien, __self__).__init__(
            'gcp:resourcemanager/lien:Lien',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, create_time=None, name=None, origin=None, parent=None, reason=None, restrictions=None):
        """
        Get an existing Lien resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Time of creation
        :param pulumi.Input[str] name: A system-generated unique identifier for this Lien.
        :param pulumi.Input[str] origin: A stable, user-visible/meaningful string identifying the origin
               of the Lien, intended to be inspected programmatically. Maximum length of
               200 characters.
        :param pulumi.Input[str] parent: A reference to the resource this Lien is attached to.
               The server will validate the parent against those for which Liens are supported.
               Since a variety of objects can have Liens against them, you must provide the type
               prefix (e.g. "projects/my-project-name").
        :param pulumi.Input[str] reason: Concise user-visible strings indicating why an action cannot be performed
               on a resource. Maximum length of 200 characters.
        :param pulumi.Input[list] restrictions: The types of operations which should be blocked as a result of this Lien.
               Each value should correspond to an IAM permission. The server will validate
               the permissions against those for which Liens are supported.  An empty
               list is meaningless and will be rejected.
               e.g. ['resourcemanager.projects.delete']
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["create_time"] = create_time
        __props__["name"] = name
        __props__["origin"] = origin
        __props__["parent"] = parent
        __props__["reason"] = reason
        __props__["restrictions"] = restrictions
        return Lien(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

