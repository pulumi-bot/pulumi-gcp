# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Project(pulumi.CustomResource):
    auto_create_network: pulumi.Output[bool]
    """
    Create the 'default' network automatically.  Default `true`.
    If set to `false`, the default network will be deleted.  Note that, for quota purposes, you
    will still need to have 1 network slot available to create the project successfully, even if
    you set `auto_create_network` to `false`, since the network will exist momentarily.
    """
    billing_account: pulumi.Output[str]
    """
    The alphanumeric ID of the billing account this project
    belongs to. The user or service account performing this operation with the provider
    must have Billing Account Administrator privileges (`roles/billing.admin`) in
    the organization. See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/v1/how-tos/access-control)
    for more details.
    """
    folder_id: pulumi.Output[str]
    """
    The numeric ID of the folder this project should be
    created under. Only one of `org_id` or `folder_id` may be
    specified. If the `folder_id` is specified, then the project is
    created under the specified folder. Changing this forces the
    project to be migrated to the newly specified folder.
    """
    labels: pulumi.Output[dict]
    """
    A set of key/value label pairs to assign to the project.
    """
    name: pulumi.Output[str]
    """
    The display name of the project.
    """
    number: pulumi.Output[str]
    """
    The numeric identifier of the project.
    """
    org_id: pulumi.Output[str]
    """
    The numeric ID of the organization this project belongs to.
    Changing this forces a new project to be created.  Only one of
    `org_id` or `folder_id` may be specified. If the `org_id` is
    specified then the project is created at the top level. Changing
    this forces the project to be migrated to the newly specified
    organization.
    """
    project_id: pulumi.Output[str]
    """
    The project ID. Changing this forces a new project to be created.
    """
    skip_delete: pulumi.Output[bool]
    """
    If true, the resource can be deleted
    without deleting the Project via the Google API.
    """
    def __init__(__self__, resource_name, opts=None, auto_create_network=None, billing_account=None, folder_id=None, labels=None, name=None, org_id=None, project_id=None, skip_delete=None, __props__=None, __name__=None, __opts__=None):
        """
        Allows creation and management of a Google Cloud Platform project.

        Projects created with this resource must be associated with an Organization.
        See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.

        The service account used to run this provider when creating a `organizations.Project`
        resource must have `roles/resourcemanager.projectCreator`. See the
        [Access Control for Organizations Using IAM](https://cloud.google.com/resource-manager/docs/access-control-org)
        doc for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_create_network: Create the 'default' network automatically.  Default `true`.
               If set to `false`, the default network will be deleted.  Note that, for quota purposes, you
               will still need to have 1 network slot available to create the project successfully, even if
               you set `auto_create_network` to `false`, since the network will exist momentarily.
        :param pulumi.Input[str] billing_account: The alphanumeric ID of the billing account this project
               belongs to. The user or service account performing this operation with the provider
               must have Billing Account Administrator privileges (`roles/billing.admin`) in
               the organization. See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/v1/how-tos/access-control)
               for more details.
        :param pulumi.Input[str] folder_id: The numeric ID of the folder this project should be
               created under. Only one of `org_id` or `folder_id` may be
               specified. If the `folder_id` is specified, then the project is
               created under the specified folder. Changing this forces the
               project to be migrated to the newly specified folder.
        :param pulumi.Input[dict] labels: A set of key/value label pairs to assign to the project.
        :param pulumi.Input[str] name: The display name of the project.
        :param pulumi.Input[str] org_id: The numeric ID of the organization this project belongs to.
               Changing this forces a new project to be created.  Only one of
               `org_id` or `folder_id` may be specified. If the `org_id` is
               specified then the project is created at the top level. Changing
               this forces the project to be migrated to the newly specified
               organization.
        :param pulumi.Input[str] project_id: The project ID. Changing this forces a new project to be created.
        :param pulumi.Input[bool] skip_delete: If true, the resource can be deleted
               without deleting the Project via the Google API.
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

            __props__['auto_create_network'] = auto_create_network
            __props__['billing_account'] = billing_account
            __props__['folder_id'] = folder_id
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['org_id'] = org_id
            if project_id is None:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            __props__['skip_delete'] = skip_delete
            __props__['number'] = None
        super(Project, __self__).__init__(
            'gcp:organizations/project:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, auto_create_network=None, billing_account=None, folder_id=None, labels=None, name=None, number=None, org_id=None, project_id=None, skip_delete=None):
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_create_network: Create the 'default' network automatically.  Default `true`.
               If set to `false`, the default network will be deleted.  Note that, for quota purposes, you
               will still need to have 1 network slot available to create the project successfully, even if
               you set `auto_create_network` to `false`, since the network will exist momentarily.
        :param pulumi.Input[str] billing_account: The alphanumeric ID of the billing account this project
               belongs to. The user or service account performing this operation with the provider
               must have Billing Account Administrator privileges (`roles/billing.admin`) in
               the organization. See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/v1/how-tos/access-control)
               for more details.
        :param pulumi.Input[str] folder_id: The numeric ID of the folder this project should be
               created under. Only one of `org_id` or `folder_id` may be
               specified. If the `folder_id` is specified, then the project is
               created under the specified folder. Changing this forces the
               project to be migrated to the newly specified folder.
        :param pulumi.Input[dict] labels: A set of key/value label pairs to assign to the project.
        :param pulumi.Input[str] name: The display name of the project.
        :param pulumi.Input[str] number: The numeric identifier of the project.
        :param pulumi.Input[str] org_id: The numeric ID of the organization this project belongs to.
               Changing this forces a new project to be created.  Only one of
               `org_id` or `folder_id` may be specified. If the `org_id` is
               specified then the project is created at the top level. Changing
               this forces the project to be migrated to the newly specified
               organization.
        :param pulumi.Input[str] project_id: The project ID. Changing this forces a new project to be created.
        :param pulumi.Input[bool] skip_delete: If true, the resource can be deleted
               without deleting the Project via the Google API.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_create_network"] = auto_create_network
        __props__["billing_account"] = billing_account
        __props__["folder_id"] = folder_id
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["number"] = number
        __props__["org_id"] = org_id
        __props__["project_id"] = project_id
        __props__["skip_delete"] = skip_delete
        return Project(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
