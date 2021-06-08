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

__all__ = ['PreventionDeidentifyTemplateArgs', 'PreventionDeidentifyTemplate']

@pulumi.input_type
class PreventionDeidentifyTemplateArgs:
    def __init__(__self__, *,
                 deidentify_config: pulumi.Input['PreventionDeidentifyTemplateDeidentifyConfigArgs'],
                 parent: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PreventionDeidentifyTemplate resource.
        :param pulumi.Input['PreventionDeidentifyTemplateDeidentifyConfigArgs'] deidentify_config: Configuration of the deidentify template
               Structure is documented below.
        :param pulumi.Input[str] parent: The parent of the template in any of the following formats:
               * `projects/{{project}}`
               * `projects/{{project}}/locations/{{location}}`
               * `organizations/{{organization_id}}`
               * `organizations/{{organization_id}}/locations/{{location}}`
        :param pulumi.Input[str] description: A description of the template.
        :param pulumi.Input[str] display_name: User set display name of the template.
        """
        pulumi.set(__self__, "deidentify_config", deidentify_config)
        pulumi.set(__self__, "parent", parent)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)

    @property
    @pulumi.getter(name="deidentifyConfig")
    def deidentify_config(self) -> pulumi.Input['PreventionDeidentifyTemplateDeidentifyConfigArgs']:
        """
        Configuration of the deidentify template
        Structure is documented below.
        """
        return pulumi.get(self, "deidentify_config")

    @deidentify_config.setter
    def deidentify_config(self, value: pulumi.Input['PreventionDeidentifyTemplateDeidentifyConfigArgs']):
        pulumi.set(self, "deidentify_config", value)

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Input[str]:
        """
        The parent of the template in any of the following formats:
        * `projects/{{project}}`
        * `projects/{{project}}/locations/{{location}}`
        * `organizations/{{organization_id}}`
        * `organizations/{{organization_id}}/locations/{{location}}`
        """
        return pulumi.get(self, "parent")

    @parent.setter
    def parent(self, value: pulumi.Input[str]):
        pulumi.set(self, "parent", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the template.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        User set display name of the template.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)


@pulumi.input_type
class _PreventionDeidentifyTemplateState:
    def __init__(__self__, *,
                 deidentify_config: Optional[pulumi.Input['PreventionDeidentifyTemplateDeidentifyConfigArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PreventionDeidentifyTemplate resources.
        :param pulumi.Input['PreventionDeidentifyTemplateDeidentifyConfigArgs'] deidentify_config: Configuration of the deidentify template
               Structure is documented below.
        :param pulumi.Input[str] description: A description of the template.
        :param pulumi.Input[str] display_name: User set display name of the template.
        :param pulumi.Input[str] name: Name of the information type.
        :param pulumi.Input[str] parent: The parent of the template in any of the following formats:
               * `projects/{{project}}`
               * `projects/{{project}}/locations/{{location}}`
               * `organizations/{{organization_id}}`
               * `organizations/{{organization_id}}/locations/{{location}}`
        """
        if deidentify_config is not None:
            pulumi.set(__self__, "deidentify_config", deidentify_config)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent is not None:
            pulumi.set(__self__, "parent", parent)

    @property
    @pulumi.getter(name="deidentifyConfig")
    def deidentify_config(self) -> Optional[pulumi.Input['PreventionDeidentifyTemplateDeidentifyConfigArgs']]:
        """
        Configuration of the deidentify template
        Structure is documented below.
        """
        return pulumi.get(self, "deidentify_config")

    @deidentify_config.setter
    def deidentify_config(self, value: Optional[pulumi.Input['PreventionDeidentifyTemplateDeidentifyConfigArgs']]):
        pulumi.set(self, "deidentify_config", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the template.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        User set display name of the template.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the information type.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parent(self) -> Optional[pulumi.Input[str]]:
        """
        The parent of the template in any of the following formats:
        * `projects/{{project}}`
        * `projects/{{project}}/locations/{{location}}`
        * `organizations/{{organization_id}}`
        * `organizations/{{organization_id}}/locations/{{location}}`
        """
        return pulumi.get(self, "parent")

    @parent.setter
    def parent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent", value)


class PreventionDeidentifyTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deidentify_config: Optional[pulumi.Input[pulumi.InputType['PreventionDeidentifyTemplateDeidentifyConfigArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows creation of templates to de-identify content.

        To get more information about DeidentifyTemplate, see:

        * [API documentation](https://cloud.google.com/dlp/docs/reference/rest/v2/projects.deidentifyTemplates)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/dlp/docs/concepts-templates)

        ## Example Usage
        ### Dlp Deidentify Template Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic = gcp.dataloss.PreventionDeidentifyTemplate("basic",
            deidentify_config=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigArgs(
                info_type_transformations=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsArgs(
                    transformations=[
                        gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationArgs(
                            info_types=[
                                gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs(
                                    name="PHONE_NUMBER",
                                ),
                                gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs(
                                    name="CREDIT_CARD_NUMBER",
                                ),
                            ],
                            primitive_transformation=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationArgs(
                                replace_config=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigArgs(
                                    new_value=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigNewValueArgs(
                                        integer_value=9,
                                    ),
                                ),
                            ),
                        ),
                        gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationArgs(
                            info_types=[
                                gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs(
                                    name="EMAIL_ADDRESS",
                                ),
                                gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs(
                                    name="LAST_NAME",
                                ),
                            ],
                            primitive_transformation=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationArgs(
                                character_mask_config=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCharacterMaskConfigArgs(
                                    characters_to_ignore=[{
                                        "commonCharactersToIgnore": "PUNCTUATION",
                                    }],
                                    masking_character="X",
                                    number_to_mask=4,
                                    reverse_order=True,
                                ),
                            ),
                        ),
                        gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationArgs(
                            info_types=[gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs(
                                name="DATE_OF_BIRTH",
                            )],
                            primitive_transformation=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationArgs(
                                replace_config=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigArgs(
                                    new_value=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigNewValueArgs(
                                        date_value=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigNewValueDateValueArgs(
                                            day=1,
                                            month=1,
                                            year=2020,
                                        ),
                                    ),
                                ),
                            ),
                        ),
                    ],
                ),
            ),
            description="Description",
            display_name="Displayname",
            parent="projects/my-project-name")
        ```

        ## Import

        DeidentifyTemplate can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate default {{parent}}/deidentifyTemplates/{{name}}
        ```

        ```sh
         $ pulumi import gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate default {{parent}}/{{name}}
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PreventionDeidentifyTemplateDeidentifyConfigArgs']] deidentify_config: Configuration of the deidentify template
               Structure is documented below.
        :param pulumi.Input[str] description: A description of the template.
        :param pulumi.Input[str] display_name: User set display name of the template.
        :param pulumi.Input[str] parent: The parent of the template in any of the following formats:
               * `projects/{{project}}`
               * `projects/{{project}}/locations/{{location}}`
               * `organizations/{{organization_id}}`
               * `organizations/{{organization_id}}/locations/{{location}}`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: PreventionDeidentifyTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows creation of templates to de-identify content.

        To get more information about DeidentifyTemplate, see:

        * [API documentation](https://cloud.google.com/dlp/docs/reference/rest/v2/projects.deidentifyTemplates)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/dlp/docs/concepts-templates)

        ## Example Usage
        ### Dlp Deidentify Template Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic = gcp.dataloss.PreventionDeidentifyTemplate("basic",
            deidentify_config=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigArgs(
                info_type_transformations=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsArgs(
                    transformations=[
                        gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationArgs(
                            info_types=[
                                gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs(
                                    name="PHONE_NUMBER",
                                ),
                                gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs(
                                    name="CREDIT_CARD_NUMBER",
                                ),
                            ],
                            primitive_transformation=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationArgs(
                                replace_config=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigArgs(
                                    new_value=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigNewValueArgs(
                                        integer_value=9,
                                    ),
                                ),
                            ),
                        ),
                        gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationArgs(
                            info_types=[
                                gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs(
                                    name="EMAIL_ADDRESS",
                                ),
                                gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs(
                                    name="LAST_NAME",
                                ),
                            ],
                            primitive_transformation=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationArgs(
                                character_mask_config=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCharacterMaskConfigArgs(
                                    characters_to_ignore=[{
                                        "commonCharactersToIgnore": "PUNCTUATION",
                                    }],
                                    masking_character="X",
                                    number_to_mask=4,
                                    reverse_order=True,
                                ),
                            ),
                        ),
                        gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationArgs(
                            info_types=[gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs(
                                name="DATE_OF_BIRTH",
                            )],
                            primitive_transformation=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationArgs(
                                replace_config=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigArgs(
                                    new_value=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigNewValueArgs(
                                        date_value=gcp.dataloss.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigNewValueDateValueArgs(
                                            day=1,
                                            month=1,
                                            year=2020,
                                        ),
                                    ),
                                ),
                            ),
                        ),
                    ],
                ),
            ),
            description="Description",
            display_name="Displayname",
            parent="projects/my-project-name")
        ```

        ## Import

        DeidentifyTemplate can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate default {{parent}}/deidentifyTemplates/{{name}}
        ```

        ```sh
         $ pulumi import gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate default {{parent}}/{{name}}
        ```

        :param str resource_name_: The name of the resource.
        :param PreventionDeidentifyTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PreventionDeidentifyTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deidentify_config: Optional[pulumi.Input[pulumi.InputType['PreventionDeidentifyTemplateDeidentifyConfigArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
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
            __props__ = PreventionDeidentifyTemplateArgs.__new__(PreventionDeidentifyTemplateArgs)

            if deidentify_config is None and not opts.urn:
                raise TypeError("Missing required property 'deidentify_config'")
            __props__.__dict__["deidentify_config"] = deidentify_config
            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            if parent is None and not opts.urn:
                raise TypeError("Missing required property 'parent'")
            __props__.__dict__["parent"] = parent
            __props__.__dict__["name"] = None
        super(PreventionDeidentifyTemplate, __self__).__init__(
            'gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            deidentify_config: Optional[pulumi.Input[pulumi.InputType['PreventionDeidentifyTemplateDeidentifyConfigArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent: Optional[pulumi.Input[str]] = None) -> 'PreventionDeidentifyTemplate':
        """
        Get an existing PreventionDeidentifyTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PreventionDeidentifyTemplateDeidentifyConfigArgs']] deidentify_config: Configuration of the deidentify template
               Structure is documented below.
        :param pulumi.Input[str] description: A description of the template.
        :param pulumi.Input[str] display_name: User set display name of the template.
        :param pulumi.Input[str] name: Name of the information type.
        :param pulumi.Input[str] parent: The parent of the template in any of the following formats:
               * `projects/{{project}}`
               * `projects/{{project}}/locations/{{location}}`
               * `organizations/{{organization_id}}`
               * `organizations/{{organization_id}}/locations/{{location}}`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PreventionDeidentifyTemplateState.__new__(_PreventionDeidentifyTemplateState)

        __props__.__dict__["deidentify_config"] = deidentify_config
        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["name"] = name
        __props__.__dict__["parent"] = parent
        return PreventionDeidentifyTemplate(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deidentifyConfig")
    def deidentify_config(self) -> pulumi.Output['outputs.PreventionDeidentifyTemplateDeidentifyConfig']:
        """
        Configuration of the deidentify template
        Structure is documented below.
        """
        return pulumi.get(self, "deidentify_config")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the template.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        User set display name of the template.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the information type.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Output[str]:
        """
        The parent of the template in any of the following formats:
        * `projects/{{project}}`
        * `projects/{{project}}/locations/{{location}}`
        * `organizations/{{organization_id}}`
        * `organizations/{{organization_id}}/locations/{{location}}`
        """
        return pulumi.get(self, "parent")

