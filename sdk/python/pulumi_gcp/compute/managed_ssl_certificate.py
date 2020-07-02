# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ManagedSslCertificate(pulumi.CustomResource):
    certificate_id: pulumi.Output[float]
    """
    The unique identifier for the resource.
    """
    creation_timestamp: pulumi.Output[str]
    """
    Creation timestamp in RFC3339 text format.
    """
    description: pulumi.Output[str]
    """
    An optional description of this resource.
    """
    expire_time: pulumi.Output[str]
    """
    Expire time of the certificate.
    """
    managed: pulumi.Output[dict]
    """
    Properties relevant to a managed certificate.  These will be used if the
    certificate is managed (as indicated by a value of `MANAGED` in `type`).  Structure is documented below.

      * `domains` (`list`) - Domains for which a managed SSL certificate will be valid.  Currently,
        there can be up to 100 domains in this list.
    """
    name: pulumi.Output[str]
    """
    Name of the resource. Provided by the client when the resource is
    created. The name must be 1-63 characters long, and comply with
    RFC1035. Specifically, the name must be 1-63 characters long and match
    the regular expression `a-z?` which means the
    first character must be a lowercase letter, and all following
    characters must be a dash, lowercase letter, or digit, except the last
    character, which cannot be a dash.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    subject_alternative_names: pulumi.Output[list]
    """
    Domains associated with the certificate via Subject Alternative Name.
    """
    type: pulumi.Output[str]
    """
    Enum field whose value is always `MANAGED` - used to signal to the API
    which type this is.
    """
    def __init__(__self__, resource_name, opts=None, certificate_id=None, description=None, managed=None, name=None, project=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        An SslCertificate resource, used for HTTPS load balancing.  This resource
        represents a certificate for which the certificate secrets are created and
        managed by Google.

        For a resource where you provide the key, see the
        SSL Certificate resource.

        To get more information about ManagedSslCertificate, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/sslCertificates)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/load-balancing/docs/ssl-certificates)

        > **Warning:** This resource should be used with extreme caution!  Provisioning an SSL
        certificate is complex.  Ensure that you understand the lifecycle of a
        certificate before attempting complex tasks like cert rotation automatically.
        This resource will "return" as soon as the certificate object is created,
        but post-creation the certificate object will go through a "provisioning"
        process.  The provisioning process can complete only when the domain name
        for which the certificate is created points to a target pool which, itself,
        points at the certificate.  Depending on your DNS provider, this may take
        some time, and migrating from self-managed certificates to Google-managed
        certificates may entail some downtime while the certificate provisions.

        In conclusion: Be extremely cautious.

        ## Example Usage
        ### Managed Ssl Certificate Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_managed_ssl_certificate = gcp.compute.ManagedSslCertificate("defaultManagedSslCertificate", managed={
            "domains": ["sslcert.tf-test.club."],
        },
        opts=ResourceOptions(provider=google_beta))
        default_http_health_check = gcp.compute.HttpHealthCheck("defaultHttpHealthCheck",
            request_path="/",
            check_interval_sec=1,
            timeout_sec=1,
            opts=ResourceOptions(provider=google_beta))
        default_backend_service = gcp.compute.BackendService("defaultBackendService",
            port_name="http",
            protocol="HTTP",
            timeout_sec=10,
            health_checks=[default_http_health_check.id],
            opts=ResourceOptions(provider=google_beta))
        default_url_map = gcp.compute.URLMap("defaultURLMap",
            description="a description",
            default_service=default_backend_service.id,
            host_rules=[{
                "hosts": ["sslcert.tf-test.club"],
                "pathMatcher": "allpaths",
            }],
            path_matchers=[{
                "name": "allpaths",
                "default_service": default_backend_service.id,
                "pathRules": [{
                    "paths": ["/*"],
                    "service": default_backend_service.id,
                }],
            }],
            opts=ResourceOptions(provider=google_beta))
        default_target_https_proxy = gcp.compute.TargetHttpsProxy("defaultTargetHttpsProxy",
            url_map=default_url_map.id,
            ssl_certificates=[default_managed_ssl_certificate.id],
            opts=ResourceOptions(provider=google_beta))
        zone = gcp.dns.ManagedZone("zone", dns_name="sslcert.tf-test.club.",
        opts=ResourceOptions(provider=google_beta))
        default_global_forwarding_rule = gcp.compute.GlobalForwardingRule("defaultGlobalForwardingRule",
            target=default_target_https_proxy.id,
            port_range=443,
            opts=ResourceOptions(provider=google_beta))
        set = gcp.dns.RecordSet("set",
            type="A",
            ttl=3600,
            managed_zone=zone.name,
            rrdatas=[default_global_forwarding_rule.ip_address],
            opts=ResourceOptions(provider=google_beta))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] certificate_id: The unique identifier for the resource.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[dict] managed: Properties relevant to a managed certificate.  These will be used if the
               certificate is managed (as indicated by a value of `MANAGED` in `type`).  Structure is documented below.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] type: Enum field whose value is always `MANAGED` - used to signal to the API
               which type this is.

        The **managed** object supports the following:

          * `domains` (`pulumi.Input[list]`) - Domains for which a managed SSL certificate will be valid.  Currently,
            there can be up to 100 domains in this list.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['certificate_id'] = certificate_id
            __props__['description'] = description
            __props__['managed'] = managed
            __props__['name'] = name
            __props__['project'] = project
            __props__['type'] = type
            __props__['creation_timestamp'] = None
            __props__['expire_time'] = None
            __props__['self_link'] = None
            __props__['subject_alternative_names'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="gcp:compute/mangedSslCertificate:MangedSslCertificate")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ManagedSslCertificate, __self__).__init__(
            'gcp:compute/managedSslCertificate:ManagedSslCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, certificate_id=None, creation_timestamp=None, description=None, expire_time=None, managed=None, name=None, project=None, self_link=None, subject_alternative_names=None, type=None):
        """
        Get an existing ManagedSslCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] certificate_id: The unique identifier for the resource.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] expire_time: Expire time of the certificate.
        :param pulumi.Input[dict] managed: Properties relevant to a managed certificate.  These will be used if the
               certificate is managed (as indicated by a value of `MANAGED` in `type`).  Structure is documented below.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[list] subject_alternative_names: Domains associated with the certificate via Subject Alternative Name.
        :param pulumi.Input[str] type: Enum field whose value is always `MANAGED` - used to signal to the API
               which type this is.

        The **managed** object supports the following:

          * `domains` (`pulumi.Input[list]`) - Domains for which a managed SSL certificate will be valid.  Currently,
            there can be up to 100 domains in this list.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["certificate_id"] = certificate_id
        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["expire_time"] = expire_time
        __props__["managed"] = managed
        __props__["name"] = name
        __props__["project"] = project
        __props__["self_link"] = self_link
        __props__["subject_alternative_names"] = subject_alternative_names
        __props__["type"] = type
        return ManagedSslCertificate(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
