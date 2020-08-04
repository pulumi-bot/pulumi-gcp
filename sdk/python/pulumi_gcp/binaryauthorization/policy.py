# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Policy(pulumi.CustomResource):
    admission_whitelist_patterns: pulumi.Output[list]
    """
    A whitelist of image patterns to exclude from admission rules. If an
    image's name matches a whitelist pattern, the image's admission
    requests will always be permitted regardless of your admission rules.  Structure is documented below.

      * `namePattern` (`str`) - An image name pattern to whitelist, in the form
        `registry/path/to/image`. This supports a trailing * as a
        wildcard, but this is allowed only in text after the registry/
        part.
    """
    cluster_admission_rules: pulumi.Output[list]
    """
    Per-cluster admission rules. An admission rule specifies either that
    all container images used in a pod creation request must be attested
    to by one or more attestors, that all pod creations will be allowed,
    or that all pod creations will be denied. There can be at most one
    admission rule per cluster spec.

      * `cluster` (`str`) - The identifier for this object. Format specified above.
      * `enforcementMode` (`str`) - The action when a pod creation is denied by the admission rule.
      * `evaluationMode` (`str`) - How this admission rule will be evaluated.
      * `requireAttestationsBies` (`list`) - The resource names of the attestors that must attest to a
        container image. If the attestor is in a different project from the
        policy, it should be specified in the format `projects/*/attestors/*`.
        Each attestor must exist before a policy can reference it. To add an
        attestor to a policy the principal issuing the policy change
        request must be able to read the attestor resource.
        Note: this field must be non-empty when the evaluation_mode field
        specifies REQUIRE_ATTESTATION, otherwise it must be empty.
    """
    default_admission_rule: pulumi.Output[dict]
    """
    Default admission rule for a cluster without a per-cluster admission
    rule.  Structure is documented below.

      * `enforcementMode` (`str`) - The action when a pod creation is denied by the admission rule.
      * `evaluationMode` (`str`) - How this admission rule will be evaluated.
      * `requireAttestationsBies` (`list`) - The resource names of the attestors that must attest to a
        container image. If the attestor is in a different project from the
        policy, it should be specified in the format `projects/*/attestors/*`.
        Each attestor must exist before a policy can reference it. To add an
        attestor to a policy the principal issuing the policy change
        request must be able to read the attestor resource.
        Note: this field must be non-empty when the evaluation_mode field
        specifies REQUIRE_ATTESTATION, otherwise it must be empty.
    """
    description: pulumi.Output[str]
    """
    A descriptive comment.
    """
    global_policy_evaluation_mode: pulumi.Output[str]
    """
    Controls the evaluation of a Google-maintained global admission policy
    for common system-level images. Images not covered by the global
    policy will be subject to the project admission policy.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    def __init__(__self__, resource_name, opts=None, admission_whitelist_patterns=None, cluster_admission_rules=None, default_admission_rule=None, description=None, global_policy_evaluation_mode=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        A policy for container image binary authorization.

        To get more information about Policy, see:

        * [API documentation](https://cloud.google.com/binary-authorization/docs/reference/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/binary-authorization/)

        ## Example Usage
        ### Binary Authorization Policy Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        note = gcp.containeranalysis.Note("note", attestation_authority=gcp.containeranalysis.NoteAttestationAuthorityArgs(
            hint=gcp.containeranalysis.NoteAttestationAuthorityHintArgs(
                human_readable_name="My attestor",
            ),
        ))
        attestor = gcp.binaryauthorization.Attestor("attestor", attestation_authority_note=gcp.binaryauthorization.AttestorAttestationAuthorityNoteArgs(
            note_reference=note.name,
        ))
        policy = gcp.binaryauthorization.Policy("policy",
            admission_whitelist_patterns=[gcp.binaryauthorization.PolicyAdmissionWhitelistPatternArgs(
                name_pattern="gcr.io/google_containers/*",
            )],
            default_admission_rule=gcp.binaryauthorization.PolicyDefaultAdmissionRuleArgs(
                evaluation_mode="ALWAYS_ALLOW",
                enforcement_mode="ENFORCED_BLOCK_AND_AUDIT_LOG",
            ),
            cluster_admission_rules=[gcp.binaryauthorization.PolicyClusterAdmissionRuleArgs(
                cluster="us-central1-a.prod-cluster",
                evaluation_mode="REQUIRE_ATTESTATION",
                enforcement_mode="ENFORCED_BLOCK_AND_AUDIT_LOG",
                require_attestations_bies=[attestor.name],
            )])
        ```
        ### Binary Authorization Policy Global Evaluation

        ```python
        import pulumi
        import pulumi_gcp as gcp

        note = gcp.containeranalysis.Note("note", attestation_authority=gcp.containeranalysis.NoteAttestationAuthorityArgs(
            hint=gcp.containeranalysis.NoteAttestationAuthorityHintArgs(
                human_readable_name="My attestor",
            ),
        ))
        attestor = gcp.binaryauthorization.Attestor("attestor", attestation_authority_note=gcp.binaryauthorization.AttestorAttestationAuthorityNoteArgs(
            note_reference=note.name,
        ))
        policy = gcp.binaryauthorization.Policy("policy",
            default_admission_rule=gcp.binaryauthorization.PolicyDefaultAdmissionRuleArgs(
                evaluation_mode="REQUIRE_ATTESTATION",
                enforcement_mode="ENFORCED_BLOCK_AND_AUDIT_LOG",
                require_attestations_bies=[attestor.name],
            ),
            global_policy_evaluation_mode="ENABLE")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] admission_whitelist_patterns: A whitelist of image patterns to exclude from admission rules. If an
               image's name matches a whitelist pattern, the image's admission
               requests will always be permitted regardless of your admission rules.  Structure is documented below.
        :param pulumi.Input[list] cluster_admission_rules: Per-cluster admission rules. An admission rule specifies either that
               all container images used in a pod creation request must be attested
               to by one or more attestors, that all pod creations will be allowed,
               or that all pod creations will be denied. There can be at most one
               admission rule per cluster spec.
        :param pulumi.Input[dict] default_admission_rule: Default admission rule for a cluster without a per-cluster admission
               rule.  Structure is documented below.
        :param pulumi.Input[str] description: A descriptive comment.
        :param pulumi.Input[str] global_policy_evaluation_mode: Controls the evaluation of a Google-maintained global admission policy
               for common system-level images. Images not covered by the global
               policy will be subject to the project admission policy.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        The **admission_whitelist_patterns** object supports the following:

          * `namePattern` (`pulumi.Input[str]`) - An image name pattern to whitelist, in the form
            `registry/path/to/image`. This supports a trailing * as a
            wildcard, but this is allowed only in text after the registry/
            part.

        The **cluster_admission_rules** object supports the following:

          * `cluster` (`pulumi.Input[str]`) - The identifier for this object. Format specified above.
          * `enforcementMode` (`pulumi.Input[str]`) - The action when a pod creation is denied by the admission rule.
          * `evaluationMode` (`pulumi.Input[str]`) - How this admission rule will be evaluated.
          * `requireAttestationsBies` (`pulumi.Input[list]`) - The resource names of the attestors that must attest to a
            container image. If the attestor is in a different project from the
            policy, it should be specified in the format `projects/*/attestors/*`.
            Each attestor must exist before a policy can reference it. To add an
            attestor to a policy the principal issuing the policy change
            request must be able to read the attestor resource.
            Note: this field must be non-empty when the evaluation_mode field
            specifies REQUIRE_ATTESTATION, otherwise it must be empty.

        The **default_admission_rule** object supports the following:

          * `enforcementMode` (`pulumi.Input[str]`) - The action when a pod creation is denied by the admission rule.
          * `evaluationMode` (`pulumi.Input[str]`) - How this admission rule will be evaluated.
          * `requireAttestationsBies` (`pulumi.Input[list]`) - The resource names of the attestors that must attest to a
            container image. If the attestor is in a different project from the
            policy, it should be specified in the format `projects/*/attestors/*`.
            Each attestor must exist before a policy can reference it. To add an
            attestor to a policy the principal issuing the policy change
            request must be able to read the attestor resource.
            Note: this field must be non-empty when the evaluation_mode field
            specifies REQUIRE_ATTESTATION, otherwise it must be empty.
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

            __props__['admission_whitelist_patterns'] = admission_whitelist_patterns
            __props__['cluster_admission_rules'] = cluster_admission_rules
            if default_admission_rule is None:
                raise TypeError("Missing required property 'default_admission_rule'")
            __props__['default_admission_rule'] = default_admission_rule
            __props__['description'] = description
            __props__['global_policy_evaluation_mode'] = global_policy_evaluation_mode
            __props__['project'] = project
        super(Policy, __self__).__init__(
            'gcp:binaryauthorization/policy:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, admission_whitelist_patterns=None, cluster_admission_rules=None, default_admission_rule=None, description=None, global_policy_evaluation_mode=None, project=None):
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] admission_whitelist_patterns: A whitelist of image patterns to exclude from admission rules. If an
               image's name matches a whitelist pattern, the image's admission
               requests will always be permitted regardless of your admission rules.  Structure is documented below.
        :param pulumi.Input[list] cluster_admission_rules: Per-cluster admission rules. An admission rule specifies either that
               all container images used in a pod creation request must be attested
               to by one or more attestors, that all pod creations will be allowed,
               or that all pod creations will be denied. There can be at most one
               admission rule per cluster spec.
        :param pulumi.Input[dict] default_admission_rule: Default admission rule for a cluster without a per-cluster admission
               rule.  Structure is documented below.
        :param pulumi.Input[str] description: A descriptive comment.
        :param pulumi.Input[str] global_policy_evaluation_mode: Controls the evaluation of a Google-maintained global admission policy
               for common system-level images. Images not covered by the global
               policy will be subject to the project admission policy.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        The **admission_whitelist_patterns** object supports the following:

          * `namePattern` (`pulumi.Input[str]`) - An image name pattern to whitelist, in the form
            `registry/path/to/image`. This supports a trailing * as a
            wildcard, but this is allowed only in text after the registry/
            part.

        The **cluster_admission_rules** object supports the following:

          * `cluster` (`pulumi.Input[str]`) - The identifier for this object. Format specified above.
          * `enforcementMode` (`pulumi.Input[str]`) - The action when a pod creation is denied by the admission rule.
          * `evaluationMode` (`pulumi.Input[str]`) - How this admission rule will be evaluated.
          * `requireAttestationsBies` (`pulumi.Input[list]`) - The resource names of the attestors that must attest to a
            container image. If the attestor is in a different project from the
            policy, it should be specified in the format `projects/*/attestors/*`.
            Each attestor must exist before a policy can reference it. To add an
            attestor to a policy the principal issuing the policy change
            request must be able to read the attestor resource.
            Note: this field must be non-empty when the evaluation_mode field
            specifies REQUIRE_ATTESTATION, otherwise it must be empty.

        The **default_admission_rule** object supports the following:

          * `enforcementMode` (`pulumi.Input[str]`) - The action when a pod creation is denied by the admission rule.
          * `evaluationMode` (`pulumi.Input[str]`) - How this admission rule will be evaluated.
          * `requireAttestationsBies` (`pulumi.Input[list]`) - The resource names of the attestors that must attest to a
            container image. If the attestor is in a different project from the
            policy, it should be specified in the format `projects/*/attestors/*`.
            Each attestor must exist before a policy can reference it. To add an
            attestor to a policy the principal issuing the policy change
            request must be able to read the attestor resource.
            Note: this field must be non-empty when the evaluation_mode field
            specifies REQUIRE_ATTESTATION, otherwise it must be empty.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["admission_whitelist_patterns"] = admission_whitelist_patterns
        __props__["cluster_admission_rules"] = cluster_admission_rules
        __props__["default_admission_rule"] = default_admission_rule
        __props__["description"] = description
        __props__["global_policy_evaluation_mode"] = global_policy_evaluation_mode
        __props__["project"] = project
        return Policy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
