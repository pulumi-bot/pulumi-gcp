# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DomainTrustArgs', 'DomainTrust']

@pulumi.input_type
class DomainTrustArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 target_dns_ip_addresses: pulumi.Input[Sequence[pulumi.Input[str]]],
                 target_domain_name: pulumi.Input[str],
                 trust_direction: pulumi.Input[str],
                 trust_handshake_secret: pulumi.Input[str],
                 trust_type: pulumi.Input[str],
                 project: Optional[pulumi.Input[str]] = None,
                 selective_authentication: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a DomainTrust resource.
        :param pulumi.Input[str] domain: The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
               https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_dns_ip_addresses: The target DNS server IP addresses which can resolve the remote domain involved in the trust.
        :param pulumi.Input[str] target_domain_name: The fully qualified target domain name which will be in trust with the current domain.
        :param pulumi.Input[str] trust_direction: The trust direction, which decides if the current domain is trusted, trusting, or both.
               Possible values are `INBOUND`, `OUTBOUND`, and `BIDIRECTIONAL`.
        :param pulumi.Input[str] trust_handshake_secret: The trust secret used for the handshake with the target domain. This will not be stored.
               **Note**: This property is sensitive and will not be displayed in the plan.
        :param pulumi.Input[str] trust_type: The type of trust represented by the trust resource.
               Possible values are `FOREST` and `EXTERNAL`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[bool] selective_authentication: Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "target_dns_ip_addresses", target_dns_ip_addresses)
        pulumi.set(__self__, "target_domain_name", target_domain_name)
        pulumi.set(__self__, "trust_direction", trust_direction)
        pulumi.set(__self__, "trust_handshake_secret", trust_handshake_secret)
        pulumi.set(__self__, "trust_type", trust_type)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if selective_authentication is not None:
            pulumi.set(__self__, "selective_authentication", selective_authentication)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
        https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="targetDnsIpAddresses")
    def target_dns_ip_addresses(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The target DNS server IP addresses which can resolve the remote domain involved in the trust.
        """
        return pulumi.get(self, "target_dns_ip_addresses")

    @target_dns_ip_addresses.setter
    def target_dns_ip_addresses(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "target_dns_ip_addresses", value)

    @property
    @pulumi.getter(name="targetDomainName")
    def target_domain_name(self) -> pulumi.Input[str]:
        """
        The fully qualified target domain name which will be in trust with the current domain.
        """
        return pulumi.get(self, "target_domain_name")

    @target_domain_name.setter
    def target_domain_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_domain_name", value)

    @property
    @pulumi.getter(name="trustDirection")
    def trust_direction(self) -> pulumi.Input[str]:
        """
        The trust direction, which decides if the current domain is trusted, trusting, or both.
        Possible values are `INBOUND`, `OUTBOUND`, and `BIDIRECTIONAL`.
        """
        return pulumi.get(self, "trust_direction")

    @trust_direction.setter
    def trust_direction(self, value: pulumi.Input[str]):
        pulumi.set(self, "trust_direction", value)

    @property
    @pulumi.getter(name="trustHandshakeSecret")
    def trust_handshake_secret(self) -> pulumi.Input[str]:
        """
        The trust secret used for the handshake with the target domain. This will not be stored.
        **Note**: This property is sensitive and will not be displayed in the plan.
        """
        return pulumi.get(self, "trust_handshake_secret")

    @trust_handshake_secret.setter
    def trust_handshake_secret(self, value: pulumi.Input[str]):
        pulumi.set(self, "trust_handshake_secret", value)

    @property
    @pulumi.getter(name="trustType")
    def trust_type(self) -> pulumi.Input[str]:
        """
        The type of trust represented by the trust resource.
        Possible values are `FOREST` and `EXTERNAL`.
        """
        return pulumi.get(self, "trust_type")

    @trust_type.setter
    def trust_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "trust_type", value)

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
    @pulumi.getter(name="selectiveAuthentication")
    def selective_authentication(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
        """
        return pulumi.get(self, "selective_authentication")

    @selective_authentication.setter
    def selective_authentication(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "selective_authentication", value)


@pulumi.input_type
class _DomainTrustState:
    def __init__(__self__, *,
                 domain: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 selective_authentication: Optional[pulumi.Input[bool]] = None,
                 target_dns_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_domain_name: Optional[pulumi.Input[str]] = None,
                 trust_direction: Optional[pulumi.Input[str]] = None,
                 trust_handshake_secret: Optional[pulumi.Input[str]] = None,
                 trust_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DomainTrust resources.
        :param pulumi.Input[str] domain: The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
               https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[bool] selective_authentication: Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_dns_ip_addresses: The target DNS server IP addresses which can resolve the remote domain involved in the trust.
        :param pulumi.Input[str] target_domain_name: The fully qualified target domain name which will be in trust with the current domain.
        :param pulumi.Input[str] trust_direction: The trust direction, which decides if the current domain is trusted, trusting, or both.
               Possible values are `INBOUND`, `OUTBOUND`, and `BIDIRECTIONAL`.
        :param pulumi.Input[str] trust_handshake_secret: The trust secret used for the handshake with the target domain. This will not be stored.
               **Note**: This property is sensitive and will not be displayed in the plan.
        :param pulumi.Input[str] trust_type: The type of trust represented by the trust resource.
               Possible values are `FOREST` and `EXTERNAL`.
        """
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if selective_authentication is not None:
            pulumi.set(__self__, "selective_authentication", selective_authentication)
        if target_dns_ip_addresses is not None:
            pulumi.set(__self__, "target_dns_ip_addresses", target_dns_ip_addresses)
        if target_domain_name is not None:
            pulumi.set(__self__, "target_domain_name", target_domain_name)
        if trust_direction is not None:
            pulumi.set(__self__, "trust_direction", trust_direction)
        if trust_handshake_secret is not None:
            pulumi.set(__self__, "trust_handshake_secret", trust_handshake_secret)
        if trust_type is not None:
            pulumi.set(__self__, "trust_type", trust_type)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
        https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

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
    @pulumi.getter(name="selectiveAuthentication")
    def selective_authentication(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
        """
        return pulumi.get(self, "selective_authentication")

    @selective_authentication.setter
    def selective_authentication(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "selective_authentication", value)

    @property
    @pulumi.getter(name="targetDnsIpAddresses")
    def target_dns_ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The target DNS server IP addresses which can resolve the remote domain involved in the trust.
        """
        return pulumi.get(self, "target_dns_ip_addresses")

    @target_dns_ip_addresses.setter
    def target_dns_ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "target_dns_ip_addresses", value)

    @property
    @pulumi.getter(name="targetDomainName")
    def target_domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        The fully qualified target domain name which will be in trust with the current domain.
        """
        return pulumi.get(self, "target_domain_name")

    @target_domain_name.setter
    def target_domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_domain_name", value)

    @property
    @pulumi.getter(name="trustDirection")
    def trust_direction(self) -> Optional[pulumi.Input[str]]:
        """
        The trust direction, which decides if the current domain is trusted, trusting, or both.
        Possible values are `INBOUND`, `OUTBOUND`, and `BIDIRECTIONAL`.
        """
        return pulumi.get(self, "trust_direction")

    @trust_direction.setter
    def trust_direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trust_direction", value)

    @property
    @pulumi.getter(name="trustHandshakeSecret")
    def trust_handshake_secret(self) -> Optional[pulumi.Input[str]]:
        """
        The trust secret used for the handshake with the target domain. This will not be stored.
        **Note**: This property is sensitive and will not be displayed in the plan.
        """
        return pulumi.get(self, "trust_handshake_secret")

    @trust_handshake_secret.setter
    def trust_handshake_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trust_handshake_secret", value)

    @property
    @pulumi.getter(name="trustType")
    def trust_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of trust represented by the trust resource.
        Possible values are `FOREST` and `EXTERNAL`.
        """
        return pulumi.get(self, "trust_type")

    @trust_type.setter
    def trust_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trust_type", value)


class DomainTrust(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 selective_authentication: Optional[pulumi.Input[bool]] = None,
                 target_dns_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_domain_name: Optional[pulumi.Input[str]] = None,
                 trust_direction: Optional[pulumi.Input[str]] = None,
                 trust_handshake_secret: Optional[pulumi.Input[str]] = None,
                 trust_type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Adds a trust between Active Directory domains

        To get more information about DomainTrust, see:

        * [API documentation](https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains/attachTrust)
        * How-to Guides
            * [Active Directory Trust](https://cloud.google.com/managed-microsoft-ad/docs/create-one-way-trust)

        > **Warning:** All arguments including `trust_handshake_secret` will be stored in the raw
        state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).

        ## Example Usage
        ### Active Directory Domain Trust Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        ad_domain_trust = gcp.activedirectory.DomainTrust("ad-domain-trust",
            domain="test-managed-ad.com",
            target_dns_ip_addresses=["10.1.0.100"],
            target_domain_name="example-gcp.com",
            trust_direction="OUTBOUND",
            trust_handshake_secret="Testing1!",
            trust_type="FOREST")
        ```

        ## Import

        DomainTrust can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:activedirectory/domainTrust:DomainTrust default projects/{{project}}/locations/global/domains/{{domain}}/{{target_domain_name}}
        ```

        ```sh
         $ pulumi import gcp:activedirectory/domainTrust:DomainTrust default {{project}}/{{domain}}/{{target_domain_name}}
        ```

        ```sh
         $ pulumi import gcp:activedirectory/domainTrust:DomainTrust default {{domain}}/{{target_domain_name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
               https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[bool] selective_authentication: Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_dns_ip_addresses: The target DNS server IP addresses which can resolve the remote domain involved in the trust.
        :param pulumi.Input[str] target_domain_name: The fully qualified target domain name which will be in trust with the current domain.
        :param pulumi.Input[str] trust_direction: The trust direction, which decides if the current domain is trusted, trusting, or both.
               Possible values are `INBOUND`, `OUTBOUND`, and `BIDIRECTIONAL`.
        :param pulumi.Input[str] trust_handshake_secret: The trust secret used for the handshake with the target domain. This will not be stored.
               **Note**: This property is sensitive and will not be displayed in the plan.
        :param pulumi.Input[str] trust_type: The type of trust represented by the trust resource.
               Possible values are `FOREST` and `EXTERNAL`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DomainTrustArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Adds a trust between Active Directory domains

        To get more information about DomainTrust, see:

        * [API documentation](https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains/attachTrust)
        * How-to Guides
            * [Active Directory Trust](https://cloud.google.com/managed-microsoft-ad/docs/create-one-way-trust)

        > **Warning:** All arguments including `trust_handshake_secret` will be stored in the raw
        state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).

        ## Example Usage
        ### Active Directory Domain Trust Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        ad_domain_trust = gcp.activedirectory.DomainTrust("ad-domain-trust",
            domain="test-managed-ad.com",
            target_dns_ip_addresses=["10.1.0.100"],
            target_domain_name="example-gcp.com",
            trust_direction="OUTBOUND",
            trust_handshake_secret="Testing1!",
            trust_type="FOREST")
        ```

        ## Import

        DomainTrust can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:activedirectory/domainTrust:DomainTrust default projects/{{project}}/locations/global/domains/{{domain}}/{{target_domain_name}}
        ```

        ```sh
         $ pulumi import gcp:activedirectory/domainTrust:DomainTrust default {{project}}/{{domain}}/{{target_domain_name}}
        ```

        ```sh
         $ pulumi import gcp:activedirectory/domainTrust:DomainTrust default {{domain}}/{{target_domain_name}}
        ```

        :param str resource_name: The name of the resource.
        :param DomainTrustArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainTrustArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 selective_authentication: Optional[pulumi.Input[bool]] = None,
                 target_dns_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_domain_name: Optional[pulumi.Input[str]] = None,
                 trust_direction: Optional[pulumi.Input[str]] = None,
                 trust_handshake_secret: Optional[pulumi.Input[str]] = None,
                 trust_type: Optional[pulumi.Input[str]] = None,
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
            __props__ = DomainTrustArgs.__new__(DomainTrustArgs)

            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__['domain'] = domain
            __props__.__dict__['project'] = project
            __props__.__dict__['selective_authentication'] = selective_authentication
            if target_dns_ip_addresses is None and not opts.urn:
                raise TypeError("Missing required property 'target_dns_ip_addresses'")
            __props__.__dict__['target_dns_ip_addresses'] = target_dns_ip_addresses
            if target_domain_name is None and not opts.urn:
                raise TypeError("Missing required property 'target_domain_name'")
            __props__.__dict__['target_domain_name'] = target_domain_name
            if trust_direction is None and not opts.urn:
                raise TypeError("Missing required property 'trust_direction'")
            __props__.__dict__['trust_direction'] = trust_direction
            if trust_handshake_secret is None and not opts.urn:
                raise TypeError("Missing required property 'trust_handshake_secret'")
            __props__.__dict__['trust_handshake_secret'] = trust_handshake_secret
            if trust_type is None and not opts.urn:
                raise TypeError("Missing required property 'trust_type'")
            __props__.__dict__['trust_type'] = trust_type
        super(DomainTrust, __self__).__init__(
            'gcp:activedirectory/domainTrust:DomainTrust',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            selective_authentication: Optional[pulumi.Input[bool]] = None,
            target_dns_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            target_domain_name: Optional[pulumi.Input[str]] = None,
            trust_direction: Optional[pulumi.Input[str]] = None,
            trust_handshake_secret: Optional[pulumi.Input[str]] = None,
            trust_type: Optional[pulumi.Input[str]] = None) -> 'DomainTrust':
        """
        Get an existing DomainTrust resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
               https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[bool] selective_authentication: Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_dns_ip_addresses: The target DNS server IP addresses which can resolve the remote domain involved in the trust.
        :param pulumi.Input[str] target_domain_name: The fully qualified target domain name which will be in trust with the current domain.
        :param pulumi.Input[str] trust_direction: The trust direction, which decides if the current domain is trusted, trusting, or both.
               Possible values are `INBOUND`, `OUTBOUND`, and `BIDIRECTIONAL`.
        :param pulumi.Input[str] trust_handshake_secret: The trust secret used for the handshake with the target domain. This will not be stored.
               **Note**: This property is sensitive and will not be displayed in the plan.
        :param pulumi.Input[str] trust_type: The type of trust represented by the trust resource.
               Possible values are `FOREST` and `EXTERNAL`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DomainTrustState.__new__(_DomainTrustState)

        __props__.__dict__['domain'] = domain
        __props__.__dict__['project'] = project
        __props__.__dict__['selective_authentication'] = selective_authentication
        __props__.__dict__['target_dns_ip_addresses'] = target_dns_ip_addresses
        __props__.__dict__['target_domain_name'] = target_domain_name
        __props__.__dict__['trust_direction'] = trust_direction
        __props__.__dict__['trust_handshake_secret'] = trust_handshake_secret
        __props__.__dict__['trust_type'] = trust_type
        return DomainTrust(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
        https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="selectiveAuthentication")
    def selective_authentication(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
        """
        return pulumi.get(self, "selective_authentication")

    @property
    @pulumi.getter(name="targetDnsIpAddresses")
    def target_dns_ip_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        The target DNS server IP addresses which can resolve the remote domain involved in the trust.
        """
        return pulumi.get(self, "target_dns_ip_addresses")

    @property
    @pulumi.getter(name="targetDomainName")
    def target_domain_name(self) -> pulumi.Output[str]:
        """
        The fully qualified target domain name which will be in trust with the current domain.
        """
        return pulumi.get(self, "target_domain_name")

    @property
    @pulumi.getter(name="trustDirection")
    def trust_direction(self) -> pulumi.Output[str]:
        """
        The trust direction, which decides if the current domain is trusted, trusting, or both.
        Possible values are `INBOUND`, `OUTBOUND`, and `BIDIRECTIONAL`.
        """
        return pulumi.get(self, "trust_direction")

    @property
    @pulumi.getter(name="trustHandshakeSecret")
    def trust_handshake_secret(self) -> pulumi.Output[str]:
        """
        The trust secret used for the handshake with the target domain. This will not be stored.
        **Note**: This property is sensitive and will not be displayed in the plan.
        """
        return pulumi.get(self, "trust_handshake_secret")

    @property
    @pulumi.getter(name="trustType")
    def trust_type(self) -> pulumi.Output[str]:
        """
        The type of trust represented by the trust resource.
        Possible values are `FOREST` and `EXTERNAL`.
        """
        return pulumi.get(self, "trust_type")

