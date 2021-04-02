# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ConsentStoreIamPolicyArgs', 'ConsentStoreIamPolicy']

@pulumi.input_type
class ConsentStoreIamPolicyArgs:
    def __init__(__self__, *,
                 consent_store_id: pulumi.Input[str],
                 dataset: pulumi.Input[str],
                 policy_data: pulumi.Input[str]):
        """
        The set of arguments for constructing a ConsentStoreIamPolicy resource.
        :param pulumi.Input[str] consent_store_id: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
               Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        """
        pulumi.set(__self__, "consent_store_id", consent_store_id)
        pulumi.set(__self__, "dataset", dataset)
        pulumi.set(__self__, "policy_data", policy_data)

    @property
    @pulumi.getter(name="consentStoreId")
    def consent_store_id(self) -> pulumi.Input[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "consent_store_id")

    @consent_store_id.setter
    def consent_store_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "consent_store_id", value)

    @property
    @pulumi.getter
    def dataset(self) -> pulumi.Input[str]:
        """
        Identifies the dataset addressed by this request. Must be in the format
        'projects/{project}/locations/{location}/datasets/{dataset}'
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "dataset")

    @dataset.setter
    def dataset(self, value: pulumi.Input[str]):
        pulumi.set(self, "dataset", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Input[str]:
        """
        The policy data generated by
        a `organizations.getIAMPolicy` data source.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_data", value)


@pulumi.input_type
class _ConsentStoreIamPolicyState:
    def __init__(__self__, *,
                 consent_store_id: Optional[pulumi.Input[str]] = None,
                 dataset: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ConsentStoreIamPolicy resources.
        :param pulumi.Input[str] consent_store_id: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
               Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        """
        if consent_store_id is not None:
            pulumi.set(__self__, "consent_store_id", consent_store_id)
        if dataset is not None:
            pulumi.set(__self__, "dataset", dataset)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if policy_data is not None:
            pulumi.set(__self__, "policy_data", policy_data)

    @property
    @pulumi.getter(name="consentStoreId")
    def consent_store_id(self) -> Optional[pulumi.Input[str]]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "consent_store_id")

    @consent_store_id.setter
    def consent_store_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "consent_store_id", value)

    @property
    @pulumi.getter
    def dataset(self) -> Optional[pulumi.Input[str]]:
        """
        Identifies the dataset addressed by this request. Must be in the format
        'projects/{project}/locations/{location}/datasets/{dataset}'
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "dataset")

    @dataset.setter
    def dataset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dataset", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the IAM policy.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> Optional[pulumi.Input[str]]:
        """
        The policy data generated by
        a `organizations.getIAMPolicy` data source.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_data", value)


class ConsentStoreIamPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consent_store_id: Optional[pulumi.Input[str]] = None,
                 dataset: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Three different resources help you manage your IAM policy for Cloud Healthcare ConsentStore. Each of these resources serves a different use case:

        * `healthcare.ConsentStoreIamPolicy`: Authoritative. Sets the IAM policy for the consentstore and replaces any existing policy already attached.
        * `healthcare.ConsentStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the consentstore are preserved.
        * `healthcare.ConsentStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the consentstore are preserved.

        > **Note:** `healthcare.ConsentStoreIamPolicy` **cannot** be used in conjunction with `healthcare.ConsentStoreIamBinding` and `healthcare.ConsentStoreIamMember` or they will fight over what your policy should be.

        > **Note:** `healthcare.ConsentStoreIamBinding` resources **can be** used in conjunction with `healthcare.ConsentStoreIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\_healthcare\_consent\_store\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/viewer",
            members=["user:jane@example.com"],
        )])
        policy = gcp.healthcare.ConsentStoreIamPolicy("policy",
            dataset=google_healthcare_consent_store["my-consent"]["dataset"],
            consent_store_id=google_healthcare_consent_store["my-consent"]["name"],
            policy_data=admin.policy_data)
        ```

        ## google\_healthcare\_consent\_store\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.healthcare.ConsentStoreIamBinding("binding",
            dataset=google_healthcare_consent_store["my-consent"]["dataset"],
            consent_store_id=google_healthcare_consent_store["my-consent"]["name"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## google\_healthcare\_consent\_store\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.healthcare.ConsentStoreIamMember("member",
            dataset=google_healthcare_consent_store["my-consent"]["dataset"],
            consent_store_id=google_healthcare_consent_store["my-consent"]["name"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* {{dataset}}/consentStores/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Healthcare consentstore IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:healthcare/consentStoreIamPolicy:ConsentStoreIamPolicy editor "{{dataset}}/consentStores/{{consent_store}} roles/viewer user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:healthcare/consentStoreIamPolicy:ConsentStoreIamPolicy editor "{{dataset}}/consentStores/{{consent_store}} roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:healthcare/consentStoreIamPolicy:ConsentStoreIamPolicy editor {{dataset}}/consentStores/{{consent_store}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] consent_store_id: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
               Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConsentStoreIamPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for Cloud Healthcare ConsentStore. Each of these resources serves a different use case:

        * `healthcare.ConsentStoreIamPolicy`: Authoritative. Sets the IAM policy for the consentstore and replaces any existing policy already attached.
        * `healthcare.ConsentStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the consentstore are preserved.
        * `healthcare.ConsentStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the consentstore are preserved.

        > **Note:** `healthcare.ConsentStoreIamPolicy` **cannot** be used in conjunction with `healthcare.ConsentStoreIamBinding` and `healthcare.ConsentStoreIamMember` or they will fight over what your policy should be.

        > **Note:** `healthcare.ConsentStoreIamBinding` resources **can be** used in conjunction with `healthcare.ConsentStoreIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\_healthcare\_consent\_store\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/viewer",
            members=["user:jane@example.com"],
        )])
        policy = gcp.healthcare.ConsentStoreIamPolicy("policy",
            dataset=google_healthcare_consent_store["my-consent"]["dataset"],
            consent_store_id=google_healthcare_consent_store["my-consent"]["name"],
            policy_data=admin.policy_data)
        ```

        ## google\_healthcare\_consent\_store\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.healthcare.ConsentStoreIamBinding("binding",
            dataset=google_healthcare_consent_store["my-consent"]["dataset"],
            consent_store_id=google_healthcare_consent_store["my-consent"]["name"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## google\_healthcare\_consent\_store\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.healthcare.ConsentStoreIamMember("member",
            dataset=google_healthcare_consent_store["my-consent"]["dataset"],
            consent_store_id=google_healthcare_consent_store["my-consent"]["name"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* {{dataset}}/consentStores/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Healthcare consentstore IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:healthcare/consentStoreIamPolicy:ConsentStoreIamPolicy editor "{{dataset}}/consentStores/{{consent_store}} roles/viewer user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:healthcare/consentStoreIamPolicy:ConsentStoreIamPolicy editor "{{dataset}}/consentStores/{{consent_store}} roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:healthcare/consentStoreIamPolicy:ConsentStoreIamPolicy editor {{dataset}}/consentStores/{{consent_store}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param ConsentStoreIamPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConsentStoreIamPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consent_store_id: Optional[pulumi.Input[str]] = None,
                 dataset: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
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
            __props__ = ConsentStoreIamPolicyArgs.__new__(ConsentStoreIamPolicyArgs)

            if consent_store_id is None and not opts.urn:
                raise TypeError("Missing required property 'consent_store_id'")
            __props__.__dict__['consent_store_id'] = consent_store_id
            if dataset is None and not opts.urn:
                raise TypeError("Missing required property 'dataset'")
            __props__.__dict__['dataset'] = dataset
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__['policy_data'] = policy_data
            __props__.__dict__['etag'] = None
        super(ConsentStoreIamPolicy, __self__).__init__(
            'gcp:healthcare/consentStoreIamPolicy:ConsentStoreIamPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            consent_store_id: Optional[pulumi.Input[str]] = None,
            dataset: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None) -> 'ConsentStoreIamPolicy':
        """
        Get an existing ConsentStoreIamPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] consent_store_id: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
               Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConsentStoreIamPolicyState.__new__(_ConsentStoreIamPolicyState)

        __props__.__dict__['consent_store_id'] = consent_store_id
        __props__.__dict__['dataset'] = dataset
        __props__.__dict__['etag'] = etag
        __props__.__dict__['policy_data'] = policy_data
        return ConsentStoreIamPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="consentStoreId")
    def consent_store_id(self) -> pulumi.Output[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "consent_store_id")

    @property
    @pulumi.getter
    def dataset(self) -> pulumi.Output[str]:
        """
        Identifies the dataset addressed by this request. Must be in the format
        'projects/{project}/locations/{location}/datasets/{dataset}'
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "dataset")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Output[str]:
        """
        The policy data generated by
        a `organizations.getIAMPolicy` data source.
        """
        return pulumi.get(self, "policy_data")

