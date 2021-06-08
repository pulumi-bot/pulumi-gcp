# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SSLPolicyArgs', 'SSLPolicy']

@pulumi.input_type
class SSLPolicyArgs:
    def __init__(__self__, *,
                 custom_features: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 min_tls_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SSLPolicy resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] custom_features: Profile specifies the set of SSL features that can be used by the
               load balancer when negotiating SSL with clients. This can be one of
               `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
               the set of SSL features to enable must be specified in the
               `customFeatures` field.
               See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
               for which ciphers are available to use. **Note**: this argument
               *must* be present when using the `CUSTOM` profile. This argument
               *must not* be present when using any other profile.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] min_tls_version: The minimum version of SSL protocol that can be used by the clients
               to establish a connection with the load balancer.
               Default value is `TLS_1_0`.
               Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] profile: Profile specifies the set of SSL features that can be used by the
               load balancer when negotiating SSL with clients. If using `CUSTOM`,
               the set of SSL features to enable must be specified in the
               `customFeatures` field.
               See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
               for information on what cipher suites each profile provides. If
               `CUSTOM` is used, the `custom_features` attribute **must be set**.
               Default value is `COMPATIBLE`.
               Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        if custom_features is not None:
            pulumi.set(__self__, "custom_features", custom_features)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if min_tls_version is not None:
            pulumi.set(__self__, "min_tls_version", min_tls_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if profile is not None:
            pulumi.set(__self__, "profile", profile)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="customFeatures")
    def custom_features(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Profile specifies the set of SSL features that can be used by the
        load balancer when negotiating SSL with clients. This can be one of
        `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
        the set of SSL features to enable must be specified in the
        `customFeatures` field.
        See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
        for which ciphers are available to use. **Note**: this argument
        *must* be present when using the `CUSTOM` profile. This argument
        *must not* be present when using any other profile.
        """
        return pulumi.get(self, "custom_features")

    @custom_features.setter
    def custom_features(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "custom_features", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="minTlsVersion")
    def min_tls_version(self) -> Optional[pulumi.Input[str]]:
        """
        The minimum version of SSL protocol that can be used by the clients
        to establish a connection with the load balancer.
        Default value is `TLS_1_0`.
        Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
        """
        return pulumi.get(self, "min_tls_version")

    @min_tls_version.setter
    def min_tls_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "min_tls_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035. Specifically, the name must be 1-63 characters long and match
        the regular expression `a-z?` which means the
        first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def profile(self) -> Optional[pulumi.Input[str]]:
        """
        Profile specifies the set of SSL features that can be used by the
        load balancer when negotiating SSL with clients. If using `CUSTOM`,
        the set of SSL features to enable must be specified in the
        `customFeatures` field.
        See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
        for information on what cipher suites each profile provides. If
        `CUSTOM` is used, the `custom_features` attribute **must be set**.
        Default value is `COMPATIBLE`.
        Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
        """
        return pulumi.get(self, "profile")

    @profile.setter
    def profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "profile", value)

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


@pulumi.input_type
class _SSLPolicyState:
    def __init__(__self__, *,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 custom_features: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled_features: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 min_tls_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SSLPolicy resources.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] custom_features: Profile specifies the set of SSL features that can be used by the
               load balancer when negotiating SSL with clients. This can be one of
               `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
               the set of SSL features to enable must be specified in the
               `customFeatures` field.
               See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
               for which ciphers are available to use. **Note**: this argument
               *must* be present when using the `CUSTOM` profile. This argument
               *must not* be present when using any other profile.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] enabled_features: The list of features enabled in the SSL policy.
        :param pulumi.Input[str] fingerprint: Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
        :param pulumi.Input[str] min_tls_version: The minimum version of SSL protocol that can be used by the clients
               to establish a connection with the load balancer.
               Default value is `TLS_1_0`.
               Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] profile: Profile specifies the set of SSL features that can be used by the
               load balancer when negotiating SSL with clients. If using `CUSTOM`,
               the set of SSL features to enable must be specified in the
               `customFeatures` field.
               See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
               for information on what cipher suites each profile provides. If
               `CUSTOM` is used, the `custom_features` attribute **must be set**.
               Default value is `COMPATIBLE`.
               Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        """
        if creation_timestamp is not None:
            pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if custom_features is not None:
            pulumi.set(__self__, "custom_features", custom_features)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled_features is not None:
            pulumi.set(__self__, "enabled_features", enabled_features)
        if fingerprint is not None:
            pulumi.set(__self__, "fingerprint", fingerprint)
        if min_tls_version is not None:
            pulumi.set(__self__, "min_tls_version", min_tls_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if profile is not None:
            pulumi.set(__self__, "profile", profile)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> Optional[pulumi.Input[str]]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @creation_timestamp.setter
    def creation_timestamp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_timestamp", value)

    @property
    @pulumi.getter(name="customFeatures")
    def custom_features(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Profile specifies the set of SSL features that can be used by the
        load balancer when negotiating SSL with clients. This can be one of
        `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
        the set of SSL features to enable must be specified in the
        `customFeatures` field.
        See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
        for which ciphers are available to use. **Note**: this argument
        *must* be present when using the `CUSTOM` profile. This argument
        *must not* be present when using any other profile.
        """
        return pulumi.get(self, "custom_features")

    @custom_features.setter
    def custom_features(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "custom_features", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enabledFeatures")
    def enabled_features(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of features enabled in the SSL policy.
        """
        return pulumi.get(self, "enabled_features")

    @enabled_features.setter
    def enabled_features(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "enabled_features", value)

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[pulumi.Input[str]]:
        """
        Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
        """
        return pulumi.get(self, "fingerprint")

    @fingerprint.setter
    def fingerprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fingerprint", value)

    @property
    @pulumi.getter(name="minTlsVersion")
    def min_tls_version(self) -> Optional[pulumi.Input[str]]:
        """
        The minimum version of SSL protocol that can be used by the clients
        to establish a connection with the load balancer.
        Default value is `TLS_1_0`.
        Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
        """
        return pulumi.get(self, "min_tls_version")

    @min_tls_version.setter
    def min_tls_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "min_tls_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035. Specifically, the name must be 1-63 characters long and match
        the regular expression `a-z?` which means the
        first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def profile(self) -> Optional[pulumi.Input[str]]:
        """
        Profile specifies the set of SSL features that can be used by the
        load balancer when negotiating SSL with clients. If using `CUSTOM`,
        the set of SSL features to enable must be specified in the
        `customFeatures` field.
        See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
        for information on what cipher suites each profile provides. If
        `CUSTOM` is used, the `custom_features` attribute **must be set**.
        Default value is `COMPATIBLE`.
        Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
        """
        return pulumi.get(self, "profile")

    @profile.setter
    def profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "profile", value)

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
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)


class SSLPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_features: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 min_tls_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Represents a SSL policy. SSL policies give you the ability to control the
        features of SSL that your SSL proxy or HTTPS load balancer negotiates.

        To get more information about SslPolicy, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/sslPolicies)
        * How-to Guides
            * [Using SSL Policies](https://cloud.google.com/compute/docs/load-balancing/ssl-policies)

        ## Example Usage
        ### Ssl Policy Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        prod_ssl_policy = gcp.compute.SSLPolicy("prod-ssl-policy", profile="MODERN")
        nonprod_ssl_policy = gcp.compute.SSLPolicy("nonprod-ssl-policy",
            min_tls_version="TLS_1_2",
            profile="MODERN")
        custom_ssl_policy = gcp.compute.SSLPolicy("custom-ssl-policy",
            custom_features=[
                "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
                "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
            ],
            min_tls_version="TLS_1_2",
            profile="CUSTOM")
        ```

        ## Import

        SslPolicy can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/sSLPolicy:SSLPolicy default projects/{{project}}/global/sslPolicies/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/sSLPolicy:SSLPolicy default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/sSLPolicy:SSLPolicy default {{name}}
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] custom_features: Profile specifies the set of SSL features that can be used by the
               load balancer when negotiating SSL with clients. This can be one of
               `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
               the set of SSL features to enable must be specified in the
               `customFeatures` field.
               See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
               for which ciphers are available to use. **Note**: this argument
               *must* be present when using the `CUSTOM` profile. This argument
               *must not* be present when using any other profile.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] min_tls_version: The minimum version of SSL protocol that can be used by the clients
               to establish a connection with the load balancer.
               Default value is `TLS_1_0`.
               Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] profile: Profile specifies the set of SSL features that can be used by the
               load balancer when negotiating SSL with clients. If using `CUSTOM`,
               the set of SSL features to enable must be specified in the
               `customFeatures` field.
               See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
               for information on what cipher suites each profile provides. If
               `CUSTOM` is used, the `custom_features` attribute **must be set**.
               Default value is `COMPATIBLE`.
               Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: Optional[SSLPolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a SSL policy. SSL policies give you the ability to control the
        features of SSL that your SSL proxy or HTTPS load balancer negotiates.

        To get more information about SslPolicy, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/sslPolicies)
        * How-to Guides
            * [Using SSL Policies](https://cloud.google.com/compute/docs/load-balancing/ssl-policies)

        ## Example Usage
        ### Ssl Policy Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        prod_ssl_policy = gcp.compute.SSLPolicy("prod-ssl-policy", profile="MODERN")
        nonprod_ssl_policy = gcp.compute.SSLPolicy("nonprod-ssl-policy",
            min_tls_version="TLS_1_2",
            profile="MODERN")
        custom_ssl_policy = gcp.compute.SSLPolicy("custom-ssl-policy",
            custom_features=[
                "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
                "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
            ],
            min_tls_version="TLS_1_2",
            profile="CUSTOM")
        ```

        ## Import

        SslPolicy can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/sSLPolicy:SSLPolicy default projects/{{project}}/global/sslPolicies/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/sSLPolicy:SSLPolicy default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/sSLPolicy:SSLPolicy default {{name}}
        ```

        :param str resource_name_: The name of the resource.
        :param SSLPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SSLPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_features: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 min_tls_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
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
            __props__ = SSLPolicyArgs.__new__(SSLPolicyArgs)

            __props__.__dict__["custom_features"] = custom_features
            __props__.__dict__["description"] = description
            __props__.__dict__["min_tls_version"] = min_tls_version
            __props__.__dict__["name"] = name
            __props__.__dict__["profile"] = profile
            __props__.__dict__["project"] = project
            __props__.__dict__["creation_timestamp"] = None
            __props__.__dict__["enabled_features"] = None
            __props__.__dict__["fingerprint"] = None
            __props__.__dict__["self_link"] = None
        super(SSLPolicy, __self__).__init__(
            'gcp:compute/sSLPolicy:SSLPolicy',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            creation_timestamp: Optional[pulumi.Input[str]] = None,
            custom_features: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enabled_features: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            fingerprint: Optional[pulumi.Input[str]] = None,
            min_tls_version: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            profile: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None) -> 'SSLPolicy':
        """
        Get an existing SSLPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] custom_features: Profile specifies the set of SSL features that can be used by the
               load balancer when negotiating SSL with clients. This can be one of
               `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
               the set of SSL features to enable must be specified in the
               `customFeatures` field.
               See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
               for which ciphers are available to use. **Note**: this argument
               *must* be present when using the `CUSTOM` profile. This argument
               *must not* be present when using any other profile.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] enabled_features: The list of features enabled in the SSL policy.
        :param pulumi.Input[str] fingerprint: Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
        :param pulumi.Input[str] min_tls_version: The minimum version of SSL protocol that can be used by the clients
               to establish a connection with the load balancer.
               Default value is `TLS_1_0`.
               Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] profile: Profile specifies the set of SSL features that can be used by the
               load balancer when negotiating SSL with clients. If using `CUSTOM`,
               the set of SSL features to enable must be specified in the
               `customFeatures` field.
               See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
               for information on what cipher suites each profile provides. If
               `CUSTOM` is used, the `custom_features` attribute **must be set**.
               Default value is `COMPATIBLE`.
               Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SSLPolicyState.__new__(_SSLPolicyState)

        __props__.__dict__["creation_timestamp"] = creation_timestamp
        __props__.__dict__["custom_features"] = custom_features
        __props__.__dict__["description"] = description
        __props__.__dict__["enabled_features"] = enabled_features
        __props__.__dict__["fingerprint"] = fingerprint
        __props__.__dict__["min_tls_version"] = min_tls_version
        __props__.__dict__["name"] = name
        __props__.__dict__["profile"] = profile
        __props__.__dict__["project"] = project
        __props__.__dict__["self_link"] = self_link
        return SSLPolicy(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter(name="customFeatures")
    def custom_features(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Profile specifies the set of SSL features that can be used by the
        load balancer when negotiating SSL with clients. This can be one of
        `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
        the set of SSL features to enable must be specified in the
        `customFeatures` field.
        See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
        for which ciphers are available to use. **Note**: this argument
        *must* be present when using the `CUSTOM` profile. This argument
        *must not* be present when using any other profile.
        """
        return pulumi.get(self, "custom_features")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description of this resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enabledFeatures")
    def enabled_features(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of features enabled in the SSL policy.
        """
        return pulumi.get(self, "enabled_features")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        """
        Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter(name="minTlsVersion")
    def min_tls_version(self) -> pulumi.Output[Optional[str]]:
        """
        The minimum version of SSL protocol that can be used by the clients
        to establish a connection with the load balancer.
        Default value is `TLS_1_0`.
        Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
        """
        return pulumi.get(self, "min_tls_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. Provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035. Specifically, the name must be 1-63 characters long and match
        the regular expression `a-z?` which means the
        first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def profile(self) -> pulumi.Output[Optional[str]]:
        """
        Profile specifies the set of SSL features that can be used by the
        load balancer when negotiating SSL with clients. If using `CUSTOM`,
        the set of SSL features to enable must be specified in the
        `customFeatures` field.
        See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
        for information on what cipher suites each profile provides. If
        `CUSTOM` is used, the `custom_features` attribute **must be set**.
        Default value is `COMPATIBLE`.
        Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
        """
        return pulumi.get(self, "profile")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

