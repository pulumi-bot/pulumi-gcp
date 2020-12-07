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

__all__ = ['WorkloadIdentityPoolProvider']


class WorkloadIdentityPoolProvider(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attribute_condition: Optional[pulumi.Input[str]] = None,
                 attribute_mapping: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 aws: Optional[pulumi.Input[pulumi.InputType['WorkloadIdentityPoolProviderAwsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 oidc: Optional[pulumi.Input[pulumi.InputType['WorkloadIdentityPoolProviderOidcArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 workload_identity_pool_id: Optional[pulumi.Input[str]] = None,
                 workload_identity_pool_provider_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Import

        WorkloadIdentityPoolProvider can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:iam/workloadIdentityPoolProvider:WorkloadIdentityPoolProvider default projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}/providers/{{workload_identity_pool_provider_id}}
        ```

        ```sh
         $ pulumi import gcp:iam/workloadIdentityPoolProvider:WorkloadIdentityPoolProvider default {{project}}/{{workload_identity_pool_id}}/{{workload_identity_pool_provider_id}}
        ```

        ```sh
         $ pulumi import gcp:iam/workloadIdentityPoolProvider:WorkloadIdentityPoolProvider default {{workload_identity_pool_id}}/{{workload_identity_pool_provider_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attribute_condition: [A Common Expression Language](https://opensource.google/projects/cel) expression, in
               plain text, to restrict what otherwise valid authentication credentials issued by the
               provider should not be accepted.
               The expression must output a boolean representing whether to allow the federation.
               The following keywords may be referenced in the expressions:
               * `assertion`: JSON representing the authentication credential issued by the provider.
               * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`.
               * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`.
               The maximum length of the attribute condition expression is 4096 characters. If
               unspecified, all valid authentication credential are accepted.
               The following example shows how to only allow credentials with a mapped `google.groups`
               value of `admins`:
               ```python
               import pulumi
               ```
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] attribute_mapping: Maps attributes from authentication credentials issued by an external identity provider
               to Google Cloud attributes, such as `subject` and `segment`.
               Each key must be a string specifying the Google Cloud IAM attribute to map to.
               The following keys are supported:
               * `google.subject`: The principal IAM is authenticating. You can reference this value
               in IAM bindings. This is also the subject that appears in Cloud Logging logs.
               Cannot exceed 127 characters.
               * `google.groups`: Groups the external identity belongs to. You can grant groups
               access to resources using an IAM `principalSet` binding; access applies to all
               members of the group.
               You can also provide custom attributes by specifying `attribute.{custom_attribute}`,
               where `{custom_attribute}` is the name of the custom attribute to be mapped. You can
               define a maximum of 50 custom attributes. The maximum length of a mapped attribute key
               is 100 characters, and the key may only contain the characters [a-z0-9_].
               You can reference these attributes in IAM policies to define fine-grained access for a
               workload to Google Cloud resources. For example:
               * `google.subject`:
               `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}`
               * `google.groups`:
               `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}`
               * `attribute.{custom_attribute}`:
               `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}`
               Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
               function that maps an identity provider credential to the normalized attribute specified
               by the corresponding map key.
               You can use the `assertion` keyword in the expression to access a JSON representation of
               the authentication credential issued by the provider.
               The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
               the total size of all mapped attributes must not exceed 8KB.
               For AWS providers, the following rules apply:
               - If no attribute mapping is defined, the following default mapping applies:
               ```python
               import pulumi
               ```
               - If any custom attribute mappings are defined, they must include a mapping to the
               `google.subject` attribute.
               For OIDC providers, the following rules apply:
               - Custom attribute mappings must be defined, and must include a mapping to the
               `google.subject` attribute. For example, the following maps the `sub` claim of the
               incoming credential to the `subject` attribute on a Google token.
               ```python
               import pulumi
               ```
        :param pulumi.Input[pulumi.InputType['WorkloadIdentityPoolProviderAwsArgs']] aws: An Amazon Web Services identity provider. Not compatible with the property oidc.
               Structure is documented below.
        :param pulumi.Input[str] description: A description for the provider. Cannot exceed 256 characters.
        :param pulumi.Input[bool] disabled: Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
               However, existing tokens still grant access.
        :param pulumi.Input[str] display_name: A display name for the provider. Cannot exceed 32 characters.
        :param pulumi.Input[pulumi.InputType['WorkloadIdentityPoolProviderOidcArgs']] oidc: An OpenId Connect 1.0 identity provider. Not compatible with the property aws.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] workload_identity_pool_id: The ID used for the pool, which is the final component of the pool resource name. This
               value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
               `gcp-` is reserved for use by Google, and may not be specified.
        :param pulumi.Input[str] workload_identity_pool_provider_id: The ID for the provider, which becomes the final component of the resource name. This
               value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
               `gcp-` is reserved for use by Google, and may not be specified.
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

            __props__['attribute_condition'] = attribute_condition
            __props__['attribute_mapping'] = attribute_mapping
            __props__['aws'] = aws
            __props__['description'] = description
            __props__['disabled'] = disabled
            __props__['display_name'] = display_name
            __props__['oidc'] = oidc
            __props__['project'] = project
            if workload_identity_pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'workload_identity_pool_id'")
            __props__['workload_identity_pool_id'] = workload_identity_pool_id
            if workload_identity_pool_provider_id is None and not opts.urn:
                raise TypeError("Missing required property 'workload_identity_pool_provider_id'")
            __props__['workload_identity_pool_provider_id'] = workload_identity_pool_provider_id
            __props__['name'] = None
            __props__['state'] = None
        super(WorkloadIdentityPoolProvider, __self__).__init__(
            'gcp:iam/workloadIdentityPoolProvider:WorkloadIdentityPoolProvider',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attribute_condition: Optional[pulumi.Input[str]] = None,
            attribute_mapping: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            aws: Optional[pulumi.Input[pulumi.InputType['WorkloadIdentityPoolProviderAwsArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            oidc: Optional[pulumi.Input[pulumi.InputType['WorkloadIdentityPoolProviderOidcArgs']]] = None,
            project: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            workload_identity_pool_id: Optional[pulumi.Input[str]] = None,
            workload_identity_pool_provider_id: Optional[pulumi.Input[str]] = None) -> 'WorkloadIdentityPoolProvider':
        """
        Get an existing WorkloadIdentityPoolProvider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attribute_condition: [A Common Expression Language](https://opensource.google/projects/cel) expression, in
               plain text, to restrict what otherwise valid authentication credentials issued by the
               provider should not be accepted.
               The expression must output a boolean representing whether to allow the federation.
               The following keywords may be referenced in the expressions:
               * `assertion`: JSON representing the authentication credential issued by the provider.
               * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`.
               * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`.
               The maximum length of the attribute condition expression is 4096 characters. If
               unspecified, all valid authentication credential are accepted.
               The following example shows how to only allow credentials with a mapped `google.groups`
               value of `admins`:
               ```python
               import pulumi
               ```
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] attribute_mapping: Maps attributes from authentication credentials issued by an external identity provider
               to Google Cloud attributes, such as `subject` and `segment`.
               Each key must be a string specifying the Google Cloud IAM attribute to map to.
               The following keys are supported:
               * `google.subject`: The principal IAM is authenticating. You can reference this value
               in IAM bindings. This is also the subject that appears in Cloud Logging logs.
               Cannot exceed 127 characters.
               * `google.groups`: Groups the external identity belongs to. You can grant groups
               access to resources using an IAM `principalSet` binding; access applies to all
               members of the group.
               You can also provide custom attributes by specifying `attribute.{custom_attribute}`,
               where `{custom_attribute}` is the name of the custom attribute to be mapped. You can
               define a maximum of 50 custom attributes. The maximum length of a mapped attribute key
               is 100 characters, and the key may only contain the characters [a-z0-9_].
               You can reference these attributes in IAM policies to define fine-grained access for a
               workload to Google Cloud resources. For example:
               * `google.subject`:
               `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}`
               * `google.groups`:
               `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}`
               * `attribute.{custom_attribute}`:
               `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}`
               Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
               function that maps an identity provider credential to the normalized attribute specified
               by the corresponding map key.
               You can use the `assertion` keyword in the expression to access a JSON representation of
               the authentication credential issued by the provider.
               The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
               the total size of all mapped attributes must not exceed 8KB.
               For AWS providers, the following rules apply:
               - If no attribute mapping is defined, the following default mapping applies:
               ```python
               import pulumi
               ```
               - If any custom attribute mappings are defined, they must include a mapping to the
               `google.subject` attribute.
               For OIDC providers, the following rules apply:
               - Custom attribute mappings must be defined, and must include a mapping to the
               `google.subject` attribute. For example, the following maps the `sub` claim of the
               incoming credential to the `subject` attribute on a Google token.
               ```python
               import pulumi
               ```
        :param pulumi.Input[pulumi.InputType['WorkloadIdentityPoolProviderAwsArgs']] aws: An Amazon Web Services identity provider. Not compatible with the property oidc.
               Structure is documented below.
        :param pulumi.Input[str] description: A description for the provider. Cannot exceed 256 characters.
        :param pulumi.Input[bool] disabled: Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
               However, existing tokens still grant access.
        :param pulumi.Input[str] display_name: A display name for the provider. Cannot exceed 32 characters.
        :param pulumi.Input[str] name: The resource name of the provider as
               'projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}/providers/{workload_identity_pool_provider_id}'.
        :param pulumi.Input[pulumi.InputType['WorkloadIdentityPoolProviderOidcArgs']] oidc: An OpenId Connect 1.0 identity provider. Not compatible with the property aws.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] state: The state of the provider. * STATE_UNSPECIFIED: State unspecified. * ACTIVE: The provider is active, and may be used to
               validate authentication credentials. * DELETED: The provider is soft-deleted. Soft-deleted providers are permanently
               deleted after approximately 30 days. You can restore a soft-deleted provider using UndeleteWorkloadIdentityPoolProvider.
               You cannot reuse the ID of a soft-deleted provider until it is permanently deleted.
        :param pulumi.Input[str] workload_identity_pool_id: The ID used for the pool, which is the final component of the pool resource name. This
               value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
               `gcp-` is reserved for use by Google, and may not be specified.
        :param pulumi.Input[str] workload_identity_pool_provider_id: The ID for the provider, which becomes the final component of the resource name. This
               value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
               `gcp-` is reserved for use by Google, and may not be specified.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["attribute_condition"] = attribute_condition
        __props__["attribute_mapping"] = attribute_mapping
        __props__["aws"] = aws
        __props__["description"] = description
        __props__["disabled"] = disabled
        __props__["display_name"] = display_name
        __props__["name"] = name
        __props__["oidc"] = oidc
        __props__["project"] = project
        __props__["state"] = state
        __props__["workload_identity_pool_id"] = workload_identity_pool_id
        __props__["workload_identity_pool_provider_id"] = workload_identity_pool_provider_id
        return WorkloadIdentityPoolProvider(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attributeCondition")
    def attribute_condition(self) -> pulumi.Output[Optional[str]]:
        """
        [A Common Expression Language](https://opensource.google/projects/cel) expression, in
        plain text, to restrict what otherwise valid authentication credentials issued by the
        provider should not be accepted.
        The expression must output a boolean representing whether to allow the federation.
        The following keywords may be referenced in the expressions:
        * `assertion`: JSON representing the authentication credential issued by the provider.
        * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`.
        * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`.
        The maximum length of the attribute condition expression is 4096 characters. If
        unspecified, all valid authentication credential are accepted.
        The following example shows how to only allow credentials with a mapped `google.groups`
        value of `admins`:
        ```python
        import pulumi
        ```
        """
        return pulumi.get(self, "attribute_condition")

    @property
    @pulumi.getter(name="attributeMapping")
    def attribute_mapping(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Maps attributes from authentication credentials issued by an external identity provider
        to Google Cloud attributes, such as `subject` and `segment`.
        Each key must be a string specifying the Google Cloud IAM attribute to map to.
        The following keys are supported:
        * `google.subject`: The principal IAM is authenticating. You can reference this value
        in IAM bindings. This is also the subject that appears in Cloud Logging logs.
        Cannot exceed 127 characters.
        * `google.groups`: Groups the external identity belongs to. You can grant groups
        access to resources using an IAM `principalSet` binding; access applies to all
        members of the group.
        You can also provide custom attributes by specifying `attribute.{custom_attribute}`,
        where `{custom_attribute}` is the name of the custom attribute to be mapped. You can
        define a maximum of 50 custom attributes. The maximum length of a mapped attribute key
        is 100 characters, and the key may only contain the characters [a-z0-9_].
        You can reference these attributes in IAM policies to define fine-grained access for a
        workload to Google Cloud resources. For example:
        * `google.subject`:
        `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}`
        * `google.groups`:
        `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}`
        * `attribute.{custom_attribute}`:
        `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}`
        Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
        function that maps an identity provider credential to the normalized attribute specified
        by the corresponding map key.
        You can use the `assertion` keyword in the expression to access a JSON representation of
        the authentication credential issued by the provider.
        The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
        the total size of all mapped attributes must not exceed 8KB.
        For AWS providers, the following rules apply:
        - If no attribute mapping is defined, the following default mapping applies:
        ```python
        import pulumi
        ```
        - If any custom attribute mappings are defined, they must include a mapping to the
        `google.subject` attribute.
        For OIDC providers, the following rules apply:
        - Custom attribute mappings must be defined, and must include a mapping to the
        `google.subject` attribute. For example, the following maps the `sub` claim of the
        incoming credential to the `subject` attribute on a Google token.
        ```python
        import pulumi
        ```
        """
        return pulumi.get(self, "attribute_mapping")

    @property
    @pulumi.getter
    def aws(self) -> pulumi.Output[Optional['outputs.WorkloadIdentityPoolProviderAws']]:
        """
        An Amazon Web Services identity provider. Not compatible with the property oidc.
        Structure is documented below.
        """
        return pulumi.get(self, "aws")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description for the provider. Cannot exceed 256 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
        However, existing tokens still grant access.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        A display name for the provider. Cannot exceed 32 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the provider as
        'projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}/providers/{workload_identity_pool_provider_id}'.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def oidc(self) -> pulumi.Output[Optional['outputs.WorkloadIdentityPoolProviderOidc']]:
        """
        An OpenId Connect 1.0 identity provider. Not compatible with the property aws.
        Structure is documented below.
        """
        return pulumi.get(self, "oidc")

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
    def state(self) -> pulumi.Output[str]:
        """
        The state of the provider. * STATE_UNSPECIFIED: State unspecified. * ACTIVE: The provider is active, and may be used to
        validate authentication credentials. * DELETED: The provider is soft-deleted. Soft-deleted providers are permanently
        deleted after approximately 30 days. You can restore a soft-deleted provider using UndeleteWorkloadIdentityPoolProvider.
        You cannot reuse the ID of a soft-deleted provider until it is permanently deleted.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="workloadIdentityPoolId")
    def workload_identity_pool_id(self) -> pulumi.Output[str]:
        """
        The ID used for the pool, which is the final component of the pool resource name. This
        value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
        `gcp-` is reserved for use by Google, and may not be specified.
        """
        return pulumi.get(self, "workload_identity_pool_id")

    @property
    @pulumi.getter(name="workloadIdentityPoolProviderId")
    def workload_identity_pool_provider_id(self) -> pulumi.Output[str]:
        """
        The ID for the provider, which becomes the final component of the resource name. This
        value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
        `gcp-` is reserved for use by Google, and may not be specified.
        """
        return pulumi.get(self, "workload_identity_pool_provider_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

