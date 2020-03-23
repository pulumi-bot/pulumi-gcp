# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Folder(pulumi.CustomResource):
    create_time: pulumi.Output[str]
    """
    Timestamp when the Folder was created. Assigned by the server.
    A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
    """
    display_name: pulumi.Output[str]
    """
    The folder’s display name.
    A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
    """
    lifecycle_state: pulumi.Output[str]
    """
    The lifecycle state of the folder such as `ACTIVE` or `DELETE_REQUESTED`.
    """
    name: pulumi.Output[str]
    """
    The resource name of the Folder. Its format is folders/{folder_id}.
    """
    parent: pulumi.Output[str]
    """
    The resource name of the parent Folder or Organization.
    Must be of the form `folders/{folder_id}` or `organizations/{org_id}`.
    """
    def __init__(__self__, resource_name, opts=None, display_name=None, parent=None, __props__=None, __name__=None, __opts__=None):
        """
        Allows management of a Google Cloud Platform folder. For more information see 
        [the official documentation](https://cloud.google.com/resource-manager/docs/creating-managing-folders)
        and 
        [API](https://cloud.google.com/resource-manager/reference/rest/v2/folders).

        A folder can contain projects, other folders, or a combination of both. You can use folders to group projects under an organization in a hierarchy. For example, your organization might contain multiple departments, each with its own set of Cloud Platform resources. Folders allows you to group these resources on a per-department basis. Folders are used to group resources that share common IAM policies.

        Folders created live inside an Organization. See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.

        The service account used to run this provider when creating a `organizations.Folder`
        resource must have `roles/resourcemanager.folderCreator`. See the
        [Access Control for Folders Using IAM](https://cloud.google.com/resource-manager/docs/access-control-folders)
        doc for more information.

        {{% examples %}}
        {{% /examples %}}

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/google_folder.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: The folder’s display name.
               A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
        :param pulumi.Input[str] parent: The resource name of the parent Folder or Organization.
               Must be of the form `folders/{folder_id}` or `organizations/{org_id}`.
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

            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            if parent is None:
                raise TypeError("Missing required property 'parent'")
            __props__['parent'] = parent
            __props__['create_time'] = None
            __props__['lifecycle_state'] = None
            __props__['name'] = None
        super(Folder, __self__).__init__(
            'gcp:organizations/folder:Folder',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, create_time=None, display_name=None, lifecycle_state=None, name=None, parent=None):
        """
        Get an existing Folder resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Timestamp when the Folder was created. Assigned by the server.
               A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        :param pulumi.Input[str] display_name: The folder’s display name.
               A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
        :param pulumi.Input[str] lifecycle_state: The lifecycle state of the folder such as `ACTIVE` or `DELETE_REQUESTED`.
        :param pulumi.Input[str] name: The resource name of the Folder. Its format is folders/{folder_id}.
        :param pulumi.Input[str] parent: The resource name of the parent Folder or Organization.
               Must be of the form `folders/{folder_id}` or `organizations/{org_id}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["create_time"] = create_time
        __props__["display_name"] = display_name
        __props__["lifecycle_state"] = lifecycle_state
        __props__["name"] = name
        __props__["parent"] = parent
        return Folder(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

