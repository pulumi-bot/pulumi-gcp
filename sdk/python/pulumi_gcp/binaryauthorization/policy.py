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

__all__ = ['Policy']


class Policy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admission_whitelist_patterns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAdmissionWhitelistPatternArgs']]]]] = None,
                 cluster_admission_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyClusterAdmissionRuleArgs']]]]] = None,
                 default_admission_rule: Optional[pulumi.Input[pulumi.InputType['PolicyDefaultAdmissionRuleArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 global_policy_evaluation_mode: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

        note = gcp.containeranalysis.Note("note", attestation_authority={
            "hint": {
                "humanReadableName": "My attestor",
            },
        })
        attestor = gcp.binaryauthorization.Attestor("attestor", attestation_authority_note={
            "noteReference": note.name,
        })
        policy = gcp.binaryauthorization.Policy("policy",
            admission_whitelist_patterns=[{
                "namePattern": "gcr.io/google_containers/*",
            }],
            default_admission_rule={
                "evaluationMode": "ALWAYS_ALLOW",
                "enforcementMode": "ENFORCED_BLOCK_AND_AUDIT_LOG",
            },
            cluster_admission_rules=[{
                "cluster": "us-central1-a.prod-cluster",
                "evaluationMode": "REQUIRE_ATTESTATION",
                "enforcementMode": "ENFORCED_BLOCK_AND_AUDIT_LOG",
                "requireAttestationsBies": [attestor.name],
            }])
        ```
        ### Binary Authorization Policy Global Evaluation

        ```python
        import pulumi
        import pulumi_gcp as gcp

        note = gcp.containeranalysis.Note("note", attestation_authority={
            "hint": {
                "humanReadableName": "My attestor",
            },
        })
        attestor = gcp.binaryauthorization.Attestor("attestor", attestation_authority_note={
            "noteReference": note.name,
        })
        policy = gcp.binaryauthorization.Policy("policy",
            default_admission_rule={
                "evaluationMode": "REQUIRE_ATTESTATION",
                "enforcementMode": "ENFORCED_BLOCK_AND_AUDIT_LOG",
                "requireAttestationsBies": [attestor.name],
            },
            global_policy_evaluation_mode="ENABLE")
        ```

        ## Import

        Policy can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:binaryauthorization/policy:Policy default projects/{{project}}
        ```

        ```sh
         $ pulumi import gcp:binaryauthorization/policy:Policy default {{project}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAdmissionWhitelistPatternArgs']]]] admission_whitelist_patterns: A whitelist of image patterns to exclude from admission rules. If an
               image's name matches a whitelist pattern, the image's admission
               requests will always be permitted regardless of your admission rules.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyClusterAdmissionRuleArgs']]]] cluster_admission_rules: Per-cluster admission rules. An admission rule specifies either that
               all container images used in a pod creation request must be attested
               to by one or more attestors, that all pod creations will be allowed,
               or that all pod creations will be denied. There can be at most one
               admission rule per cluster spec.
        :param pulumi.Input[pulumi.InputType['PolicyDefaultAdmissionRuleArgs']] default_admission_rule: Default admission rule for a cluster without a per-cluster admission
               rule.
               Structure is documented below.
        :param pulumi.Input[str] description: A descriptive comment.
        :param pulumi.Input[str] global_policy_evaluation_mode: Controls the evaluation of a Google-maintained global admission policy
               for common system-level images. Images not covered by the global
               policy will be subject to the project admission policy.
               Possible values are `ENABLE` and `DISABLE`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
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
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admission_whitelist_patterns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAdmissionWhitelistPatternArgs']]]]] = None,
            cluster_admission_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyClusterAdmissionRuleArgs']]]]] = None,
            default_admission_rule: Optional[pulumi.Input[pulumi.InputType['PolicyDefaultAdmissionRuleArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            global_policy_evaluation_mode: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'Policy':
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAdmissionWhitelistPatternArgs']]]] admission_whitelist_patterns: A whitelist of image patterns to exclude from admission rules. If an
               image's name matches a whitelist pattern, the image's admission
               requests will always be permitted regardless of your admission rules.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyClusterAdmissionRuleArgs']]]] cluster_admission_rules: Per-cluster admission rules. An admission rule specifies either that
               all container images used in a pod creation request must be attested
               to by one or more attestors, that all pod creations will be allowed,
               or that all pod creations will be denied. There can be at most one
               admission rule per cluster spec.
        :param pulumi.Input[pulumi.InputType['PolicyDefaultAdmissionRuleArgs']] default_admission_rule: Default admission rule for a cluster without a per-cluster admission
               rule.
               Structure is documented below.
        :param pulumi.Input[str] description: A descriptive comment.
        :param pulumi.Input[str] global_policy_evaluation_mode: Controls the evaluation of a Google-maintained global admission policy
               for common system-level images. Images not covered by the global
               policy will be subject to the project admission policy.
               Possible values are `ENABLE` and `DISABLE`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
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

    @property
    @pulumi.getter(name="admissionWhitelistPatterns")
    def admission_whitelist_patterns(self) -> pulumi.Output[Optional[Sequence['outputs.PolicyAdmissionWhitelistPattern']]]:
        """
        A whitelist of image patterns to exclude from admission rules. If an
        image's name matches a whitelist pattern, the image's admission
        requests will always be permitted regardless of your admission rules.
        Structure is documented below.
        """
        return pulumi.get(self, "admission_whitelist_patterns")

    @property
    @pulumi.getter(name="clusterAdmissionRules")
    def cluster_admission_rules(self) -> pulumi.Output[Optional[Sequence['outputs.PolicyClusterAdmissionRule']]]:
        """
        Per-cluster admission rules. An admission rule specifies either that
        all container images used in a pod creation request must be attested
        to by one or more attestors, that all pod creations will be allowed,
        or that all pod creations will be denied. There can be at most one
        admission rule per cluster spec.
        """
        return pulumi.get(self, "cluster_admission_rules")

    @property
    @pulumi.getter(name="defaultAdmissionRule")
    def default_admission_rule(self) -> pulumi.Output['outputs.PolicyDefaultAdmissionRule']:
        """
        Default admission rule for a cluster without a per-cluster admission
        rule.
        Structure is documented below.
        """
        return pulumi.get(self, "default_admission_rule")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A descriptive comment.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="globalPolicyEvaluationMode")
    def global_policy_evaluation_mode(self) -> pulumi.Output[str]:
        """
        Controls the evaluation of a Google-maintained global admission policy
        for common system-level images. Images not covered by the global
        policy will be subject to the project admission policy.
        Possible values are `ENABLE` and `DISABLE`.
        """
        return pulumi.get(self, "global_policy_evaluation_mode")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

