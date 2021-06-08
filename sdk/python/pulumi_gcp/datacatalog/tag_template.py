# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['TagTemplateArgs', 'TagTemplate']

@pulumi.input_type
class TagTemplateArgs:
    def __init__(__self__, *,
                 fields: pulumi.Input[Sequence[pulumi.Input['TagTemplateFieldArgs']]],
                 tag_template_id: pulumi.Input[str],
                 display_name: Optional[pulumi.Input[str]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TagTemplate resource.
        :param pulumi.Input[Sequence[pulumi.Input['TagTemplateFieldArgs']]] fields: Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.
               Structure is documented below.
        :param pulumi.Input[str] tag_template_id: The id of the tag template to create.
        :param pulumi.Input[str] display_name: The display name for this template.
        :param pulumi.Input[bool] force_delete: This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Template location region.
        """
        pulumi.set(__self__, "fields", fields)
        pulumi.set(__self__, "tag_template_id", tag_template_id)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if force_delete is not None:
            pulumi.set(__self__, "force_delete", force_delete)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Input[Sequence[pulumi.Input['TagTemplateFieldArgs']]]:
        """
        Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.
        Structure is documented below.
        """
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: pulumi.Input[Sequence[pulumi.Input['TagTemplateFieldArgs']]]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter(name="tagTemplateId")
    def tag_template_id(self) -> pulumi.Input[str]:
        """
        The id of the tag template to create.
        """
        return pulumi.get(self, "tag_template_id")

    @tag_template_id.setter
    def tag_template_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tag_template_id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name for this template.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="forceDelete")
    def force_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
        """
        return pulumi.get(self, "force_delete")

    @force_delete.setter
    def force_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_delete", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Template location region.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _TagTemplateState:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input['TagTemplateFieldArgs']]]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tag_template_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TagTemplate resources.
        :param pulumi.Input[str] display_name: The display name for this template.
        :param pulumi.Input[Sequence[pulumi.Input['TagTemplateFieldArgs']]] fields: Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.
               Structure is documented below.
        :param pulumi.Input[bool] force_delete: This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
        :param pulumi.Input[str] name: -
               The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Template location region.
        :param pulumi.Input[str] tag_template_id: The id of the tag template to create.
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if fields is not None:
            pulumi.set(__self__, "fields", fields)
        if force_delete is not None:
            pulumi.set(__self__, "force_delete", force_delete)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tag_template_id is not None:
            pulumi.set(__self__, "tag_template_id", tag_template_id)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name for this template.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def fields(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TagTemplateFieldArgs']]]]:
        """
        Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.
        Structure is documented below.
        """
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TagTemplateFieldArgs']]]]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter(name="forceDelete")
    def force_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
        """
        return pulumi.get(self, "force_delete")

    @force_delete.setter
    def force_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_delete", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        -
        The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Template location region.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="tagTemplateId")
    def tag_template_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the tag template to create.
        """
        return pulumi.get(self, "tag_template_id")

    @tag_template_id.setter
    def tag_template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag_template_id", value)


class TagTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TagTemplateFieldArgs']]]]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tag_template_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A tag template defines a tag, which can have one or more typed fields.
        The template is used to create and attach the tag to GCP resources.

        To get more information about TagTemplate, see:

        * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.tagTemplates)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/data-catalog/docs)

        ## Example Usage
        ### Data Catalog Tag Template Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic_tag_template = gcp.datacatalog.TagTemplate("basicTagTemplate",
            display_name="Demo Tag Template",
            fields=[
                gcp.datacatalog.TagTemplateFieldArgs(
                    display_name="Source of data asset",
                    field_id="source",
                    is_required=True,
                    type=gcp.datacatalog.TagTemplateFieldTypeArgs(
                        primitive_type="STRING",
                    ),
                ),
                gcp.datacatalog.TagTemplateFieldArgs(
                    display_name="Number of rows in the data asset",
                    field_id="num_rows",
                    type=gcp.datacatalog.TagTemplateFieldTypeArgs(
                        primitive_type="DOUBLE",
                    ),
                ),
                gcp.datacatalog.TagTemplateFieldArgs(
                    display_name="PII type",
                    field_id="pii_type",
                    type=gcp.datacatalog.TagTemplateFieldTypeArgs(
                        enum_type=gcp.datacatalog.TagTemplateFieldTypeEnumTypeArgs(
                            allowed_values=[
                                gcp.datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArgs(
                                    display_name="EMAIL",
                                ),
                                gcp.datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArgs(
                                    display_name="SOCIAL SECURITY NUMBER",
                                ),
                                gcp.datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArgs(
                                    display_name="NONE",
                                ),
                            ],
                        ),
                    ),
                ),
            ],
            force_delete=False,
            region="us-central1",
            tag_template_id="my_template")
        ```

        ## Import

        TagTemplate can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:datacatalog/tagTemplate:TagTemplate default {{name}}
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: The display name for this template.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TagTemplateFieldArgs']]]] fields: Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.
               Structure is documented below.
        :param pulumi.Input[bool] force_delete: This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Template location region.
        :param pulumi.Input[str] tag_template_id: The id of the tag template to create.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: TagTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A tag template defines a tag, which can have one or more typed fields.
        The template is used to create and attach the tag to GCP resources.

        To get more information about TagTemplate, see:

        * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.tagTemplates)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/data-catalog/docs)

        ## Example Usage
        ### Data Catalog Tag Template Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic_tag_template = gcp.datacatalog.TagTemplate("basicTagTemplate",
            display_name="Demo Tag Template",
            fields=[
                gcp.datacatalog.TagTemplateFieldArgs(
                    display_name="Source of data asset",
                    field_id="source",
                    is_required=True,
                    type=gcp.datacatalog.TagTemplateFieldTypeArgs(
                        primitive_type="STRING",
                    ),
                ),
                gcp.datacatalog.TagTemplateFieldArgs(
                    display_name="Number of rows in the data asset",
                    field_id="num_rows",
                    type=gcp.datacatalog.TagTemplateFieldTypeArgs(
                        primitive_type="DOUBLE",
                    ),
                ),
                gcp.datacatalog.TagTemplateFieldArgs(
                    display_name="PII type",
                    field_id="pii_type",
                    type=gcp.datacatalog.TagTemplateFieldTypeArgs(
                        enum_type=gcp.datacatalog.TagTemplateFieldTypeEnumTypeArgs(
                            allowed_values=[
                                gcp.datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArgs(
                                    display_name="EMAIL",
                                ),
                                gcp.datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArgs(
                                    display_name="SOCIAL SECURITY NUMBER",
                                ),
                                gcp.datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArgs(
                                    display_name="NONE",
                                ),
                            ],
                        ),
                    ),
                ),
            ],
            force_delete=False,
            region="us-central1",
            tag_template_id="my_template")
        ```

        ## Import

        TagTemplate can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:datacatalog/tagTemplate:TagTemplate default {{name}}
        ```

        :param str resource_name_: The name of the resource.
        :param TagTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TagTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TagTemplateFieldArgs']]]]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tag_template_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TagTemplateArgs.__new__(TagTemplateArgs)

            __props__.__dict__["display_name"] = display_name
            if fields is None and not opts.urn:
                raise TypeError("Missing required property 'fields'")
            __props__.__dict__["fields"] = fields
            __props__.__dict__["force_delete"] = force_delete
            __props__.__dict__["project"] = project
            __props__.__dict__["region"] = region
            if tag_template_id is None and not opts.urn:
                raise TypeError("Missing required property 'tag_template_id'")
            __props__.__dict__["tag_template_id"] = tag_template_id
            __props__.__dict__["name"] = None
        super(TagTemplate, __self__).__init__(
            'gcp:datacatalog/tagTemplate:TagTemplate',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            fields: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TagTemplateFieldArgs']]]]] = None,
            force_delete: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            tag_template_id: Optional[pulumi.Input[str]] = None) -> 'TagTemplate':
        """
        Get an existing TagTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: The display name for this template.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TagTemplateFieldArgs']]]] fields: Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.
               Structure is documented below.
        :param pulumi.Input[bool] force_delete: This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
        :param pulumi.Input[str] name: -
               The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Template location region.
        :param pulumi.Input[str] tag_template_id: The id of the tag template to create.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TagTemplateState.__new__(_TagTemplateState)

        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["fields"] = fields
        __props__.__dict__["force_delete"] = force_delete
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["region"] = region
        __props__.__dict__["tag_template_id"] = tag_template_id
        return TagTemplate(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        The display name for this template.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Output[Sequence['outputs.TagTemplateField']]:
        """
        Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.
        Structure is documented below.
        """
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter(name="forceDelete")
    def force_delete(self) -> pulumi.Output[Optional[bool]]:
        """
        This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
        """
        return pulumi.get(self, "force_delete")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        -
        The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Template location region.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="tagTemplateId")
    def tag_template_id(self) -> pulumi.Output[str]:
        """
        The id of the tag template to create.
        """
        return pulumi.get(self, "tag_template_id")

