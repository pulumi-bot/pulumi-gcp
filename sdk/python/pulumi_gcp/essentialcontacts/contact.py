# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ContactArgs', 'Contact']

@pulumi.input_type
class ContactArgs:
    def __init__(__self__, *,
                 email: pulumi.Input[str],
                 language_tag: pulumi.Input[str],
                 notification_category_subscriptions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 parent: pulumi.Input[str]):
        """
        The set of arguments for constructing a Contact resource.
        :param pulumi.Input[str] email: The email address to send notifications to. This does not need to be a Google account.
        :param pulumi.Input[str] language_tag: The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported languages.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_category_subscriptions: The categories of notifications that the contact will receive communications for.
        :param pulumi.Input[str] parent: The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or projects/{project_id}
        """
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "language_tag", language_tag)
        pulumi.set(__self__, "notification_category_subscriptions", notification_category_subscriptions)
        pulumi.set(__self__, "parent", parent)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[str]:
        """
        The email address to send notifications to. This does not need to be a Google account.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[str]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="languageTag")
    def language_tag(self) -> pulumi.Input[str]:
        """
        The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported languages.
        """
        return pulumi.get(self, "language_tag")

    @language_tag.setter
    def language_tag(self, value: pulumi.Input[str]):
        pulumi.set(self, "language_tag", value)

    @property
    @pulumi.getter(name="notificationCategorySubscriptions")
    def notification_category_subscriptions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The categories of notifications that the contact will receive communications for.
        """
        return pulumi.get(self, "notification_category_subscriptions")

    @notification_category_subscriptions.setter
    def notification_category_subscriptions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "notification_category_subscriptions", value)

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Input[str]:
        """
        The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or projects/{project_id}
        """
        return pulumi.get(self, "parent")

    @parent.setter
    def parent(self, value: pulumi.Input[str]):
        pulumi.set(self, "parent", value)


@pulumi.input_type
class _ContactState:
    def __init__(__self__, *,
                 email: Optional[pulumi.Input[str]] = None,
                 language_tag: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_category_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 parent: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Contact resources.
        :param pulumi.Input[str] email: The email address to send notifications to. This does not need to be a Google account.
        :param pulumi.Input[str] language_tag: The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported languages.
        :param pulumi.Input[str] name: The identifier for the contact. Format: {resourceType}/{resource_id}/contacts/{contact_id}
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_category_subscriptions: The categories of notifications that the contact will receive communications for.
        :param pulumi.Input[str] parent: The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or projects/{project_id}
        """
        if email is not None:
            pulumi.set(__self__, "email", email)
        if language_tag is not None:
            pulumi.set(__self__, "language_tag", language_tag)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notification_category_subscriptions is not None:
            pulumi.set(__self__, "notification_category_subscriptions", notification_category_subscriptions)
        if parent is not None:
            pulumi.set(__self__, "parent", parent)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        The email address to send notifications to. This does not need to be a Google account.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="languageTag")
    def language_tag(self) -> Optional[pulumi.Input[str]]:
        """
        The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported languages.
        """
        return pulumi.get(self, "language_tag")

    @language_tag.setter
    def language_tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "language_tag", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier for the contact. Format: {resourceType}/{resource_id}/contacts/{contact_id}
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationCategorySubscriptions")
    def notification_category_subscriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The categories of notifications that the contact will receive communications for.
        """
        return pulumi.get(self, "notification_category_subscriptions")

    @notification_category_subscriptions.setter
    def notification_category_subscriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notification_category_subscriptions", value)

    @property
    @pulumi.getter
    def parent(self) -> Optional[pulumi.Input[str]]:
        """
        The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or projects/{project_id}
        """
        return pulumi.get(self, "parent")

    @parent.setter
    def parent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent", value)


class Contact(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 language_tag: Optional[pulumi.Input[str]] = None,
                 notification_category_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A contact that will receive notifications from Google Cloud.

        To get more information about Contact, see:

        * [API documentation](https://cloud.google.com/resource-manager/docs/reference/essentialcontacts/rest/v1beta1/projects.contacts)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/resource-manager/docs/managing-notification-contacts)

        > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
        you must specify a `billing_project` and set `user_project_override` to true
        in the provider configuration. Otherwise the Essential Contacts API will return a 403 error.
        Your account must have the `serviceusage.services.use` permission on the
        `billing_project` you defined.

        ## Example Usage
        ### Essential Contact

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.organizations.get_project()
        contact = gcp.essentialcontacts.Contact("contact",
            parent=project.id,
            email="foo@bar.com",
            language_tag="en-GB",
            notification_category_subscriptions=["ALL"],
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        Contact can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:essentialcontacts/contact:Contact default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] email: The email address to send notifications to. This does not need to be a Google account.
        :param pulumi.Input[str] language_tag: The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported languages.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_category_subscriptions: The categories of notifications that the contact will receive communications for.
        :param pulumi.Input[str] parent: The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or projects/{project_id}
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ContactArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A contact that will receive notifications from Google Cloud.

        To get more information about Contact, see:

        * [API documentation](https://cloud.google.com/resource-manager/docs/reference/essentialcontacts/rest/v1beta1/projects.contacts)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/resource-manager/docs/managing-notification-contacts)

        > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
        you must specify a `billing_project` and set `user_project_override` to true
        in the provider configuration. Otherwise the Essential Contacts API will return a 403 error.
        Your account must have the `serviceusage.services.use` permission on the
        `billing_project` you defined.

        ## Example Usage
        ### Essential Contact

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.organizations.get_project()
        contact = gcp.essentialcontacts.Contact("contact",
            parent=project.id,
            email="foo@bar.com",
            language_tag="en-GB",
            notification_category_subscriptions=["ALL"],
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        Contact can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:essentialcontacts/contact:Contact default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param ContactArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContactArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 language_tag: Optional[pulumi.Input[str]] = None,
                 notification_category_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            __props__ = ContactArgs.__new__(ContactArgs)

            if email is None and not opts.urn:
                raise TypeError("Missing required property 'email'")
            __props__.__dict__['email'] = email
            if language_tag is None and not opts.urn:
                raise TypeError("Missing required property 'language_tag'")
            __props__.__dict__['language_tag'] = language_tag
            if notification_category_subscriptions is None and not opts.urn:
                raise TypeError("Missing required property 'notification_category_subscriptions'")
            __props__.__dict__['notification_category_subscriptions'] = notification_category_subscriptions
            if parent is None and not opts.urn:
                raise TypeError("Missing required property 'parent'")
            __props__.__dict__['parent'] = parent
            __props__.__dict__['name'] = None
        super(Contact, __self__).__init__(
            'gcp:essentialcontacts/contact:Contact',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            email: Optional[pulumi.Input[str]] = None,
            language_tag: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notification_category_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            parent: Optional[pulumi.Input[str]] = None) -> 'Contact':
        """
        Get an existing Contact resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] email: The email address to send notifications to. This does not need to be a Google account.
        :param pulumi.Input[str] language_tag: The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported languages.
        :param pulumi.Input[str] name: The identifier for the contact. Format: {resourceType}/{resource_id}/contacts/{contact_id}
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_category_subscriptions: The categories of notifications that the contact will receive communications for.
        :param pulumi.Input[str] parent: The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or projects/{project_id}
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ContactState.__new__(_ContactState)

        __props__.__dict__['email'] = email
        __props__.__dict__['language_tag'] = language_tag
        __props__.__dict__['name'] = name
        __props__.__dict__['notification_category_subscriptions'] = notification_category_subscriptions
        __props__.__dict__['parent'] = parent
        return Contact(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[str]:
        """
        The email address to send notifications to. This does not need to be a Google account.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="languageTag")
    def language_tag(self) -> pulumi.Output[str]:
        """
        The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported languages.
        """
        return pulumi.get(self, "language_tag")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The identifier for the contact. Format: {resourceType}/{resource_id}/contacts/{contact_id}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationCategorySubscriptions")
    def notification_category_subscriptions(self) -> pulumi.Output[Sequence[str]]:
        """
        The categories of notifications that the contact will receive communications for.
        """
        return pulumi.get(self, "notification_category_subscriptions")

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Output[str]:
        """
        The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or projects/{project_id}
        """
        return pulumi.get(self, "parent")

