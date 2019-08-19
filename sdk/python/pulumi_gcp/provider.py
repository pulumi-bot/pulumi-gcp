# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class Provider(pulumi.ProviderResource):
    def __init__(__self__, resource_name, opts=None, access_context_manager_custom_endpoint=None, access_token=None, app_engine_custom_endpoint=None, batching=None, bigquery_custom_endpoint=None, bigtable_custom_endpoint=None, binary_authorization_custom_endpoint=None, cloud_billing_custom_endpoint=None, cloud_build_custom_endpoint=None, cloud_functions_custom_endpoint=None, cloud_iot_custom_endpoint=None, cloud_run_custom_endpoint=None, cloud_scheduler_custom_endpoint=None, composer_custom_endpoint=None, compute_beta_custom_endpoint=None, compute_custom_endpoint=None, container_analysis_custom_endpoint=None, container_beta_custom_endpoint=None, container_custom_endpoint=None, credentials=None, dataflow_custom_endpoint=None, dataproc_beta_custom_endpoint=None, dataproc_custom_endpoint=None, dns_beta_custom_endpoint=None, dns_custom_endpoint=None, filestore_custom_endpoint=None, firestore_custom_endpoint=None, healthcare_custom_endpoint=None, iam_credentials_custom_endpoint=None, iam_custom_endpoint=None, iap_custom_endpoint=None, kms_custom_endpoint=None, logging_custom_endpoint=None, monitoring_custom_endpoint=None, project=None, pubsub_custom_endpoint=None, redis_custom_endpoint=None, region=None, resource_manager_custom_endpoint=None, resource_manager_v2beta1_custom_endpoint=None, runtimeconfig_custom_endpoint=None, scopes=None, security_scanner_custom_endpoint=None, service_management_custom_endpoint=None, service_networking_custom_endpoint=None, service_usage_custom_endpoint=None, source_repo_custom_endpoint=None, spanner_custom_endpoint=None, sql_custom_endpoint=None, storage_custom_endpoint=None, storage_transfer_custom_endpoint=None, tpu_custom_endpoint=None, zone=None, __props__=None, __name__=None, __opts__=None):
        """
        The provider type for the google-beta package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google-beta/blob/master/website/docs/index.html.markdown.
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

            __props__['access_context_manager_custom_endpoint'] = access_context_manager_custom_endpoint
            __props__['access_token'] = access_token
            __props__['app_engine_custom_endpoint'] = app_engine_custom_endpoint
            __props__['batching'] = pulumi.Output.from_input(batching).apply(json.dumps) if batching is not None else None
            __props__['bigquery_custom_endpoint'] = bigquery_custom_endpoint
            __props__['bigtable_custom_endpoint'] = bigtable_custom_endpoint
            __props__['binary_authorization_custom_endpoint'] = binary_authorization_custom_endpoint
            __props__['cloud_billing_custom_endpoint'] = cloud_billing_custom_endpoint
            __props__['cloud_build_custom_endpoint'] = cloud_build_custom_endpoint
            __props__['cloud_functions_custom_endpoint'] = cloud_functions_custom_endpoint
            __props__['cloud_iot_custom_endpoint'] = cloud_iot_custom_endpoint
            __props__['cloud_run_custom_endpoint'] = cloud_run_custom_endpoint
            __props__['cloud_scheduler_custom_endpoint'] = cloud_scheduler_custom_endpoint
            __props__['composer_custom_endpoint'] = composer_custom_endpoint
            __props__['compute_beta_custom_endpoint'] = compute_beta_custom_endpoint
            __props__['compute_custom_endpoint'] = compute_custom_endpoint
            __props__['container_analysis_custom_endpoint'] = container_analysis_custom_endpoint
            __props__['container_beta_custom_endpoint'] = container_beta_custom_endpoint
            __props__['container_custom_endpoint'] = container_custom_endpoint
            if credentials is None:
                credentials = utilities.get_env('GOOGLE_CREDENTIALS', 'GOOGLE_CLOUD_KEYFILE_JSON', 'GCLOUD_KEYFILE_JSON')
            __props__['credentials'] = credentials
            __props__['dataflow_custom_endpoint'] = dataflow_custom_endpoint
            __props__['dataproc_beta_custom_endpoint'] = dataproc_beta_custom_endpoint
            __props__['dataproc_custom_endpoint'] = dataproc_custom_endpoint
            __props__['dns_beta_custom_endpoint'] = dns_beta_custom_endpoint
            __props__['dns_custom_endpoint'] = dns_custom_endpoint
            __props__['filestore_custom_endpoint'] = filestore_custom_endpoint
            __props__['firestore_custom_endpoint'] = firestore_custom_endpoint
            __props__['healthcare_custom_endpoint'] = healthcare_custom_endpoint
            __props__['iam_credentials_custom_endpoint'] = iam_credentials_custom_endpoint
            __props__['iam_custom_endpoint'] = iam_custom_endpoint
            __props__['iap_custom_endpoint'] = iap_custom_endpoint
            __props__['kms_custom_endpoint'] = kms_custom_endpoint
            __props__['logging_custom_endpoint'] = logging_custom_endpoint
            __props__['monitoring_custom_endpoint'] = monitoring_custom_endpoint
            if project is None:
                project = utilities.get_env('GOOGLE_PROJECT', 'GOOGLE_CLOUD_PROJECT', 'GCLOUD_PROJECT', 'CLOUDSDK_CORE_PROJECT')
            __props__['project'] = project
            __props__['pubsub_custom_endpoint'] = pubsub_custom_endpoint
            __props__['redis_custom_endpoint'] = redis_custom_endpoint
            if region is None:
                region = utilities.get_env('GOOGLE_REGION', 'GCLOUD_REGION', 'CLOUDSDK_COMPUTE_REGION')
            __props__['region'] = region
            __props__['resource_manager_custom_endpoint'] = resource_manager_custom_endpoint
            __props__['resource_manager_v2beta1_custom_endpoint'] = resource_manager_v2beta1_custom_endpoint
            __props__['runtimeconfig_custom_endpoint'] = runtimeconfig_custom_endpoint
            __props__['scopes'] = pulumi.Output.from_input(scopes).apply(json.dumps) if scopes is not None else None
            __props__['security_scanner_custom_endpoint'] = security_scanner_custom_endpoint
            __props__['service_management_custom_endpoint'] = service_management_custom_endpoint
            __props__['service_networking_custom_endpoint'] = service_networking_custom_endpoint
            __props__['service_usage_custom_endpoint'] = service_usage_custom_endpoint
            __props__['source_repo_custom_endpoint'] = source_repo_custom_endpoint
            __props__['spanner_custom_endpoint'] = spanner_custom_endpoint
            __props__['sql_custom_endpoint'] = sql_custom_endpoint
            __props__['storage_custom_endpoint'] = storage_custom_endpoint
            __props__['storage_transfer_custom_endpoint'] = storage_transfer_custom_endpoint
            __props__['tpu_custom_endpoint'] = tpu_custom_endpoint
            if zone is None:
                zone = utilities.get_env('GOOGLE_ZONE', 'GCLOUD_ZONE', 'CLOUDSDK_COMPUTE_ZONE')
            __props__['zone'] = zone
        super(Provider, __self__).__init__(
            'gcp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Provider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google-beta/blob/master/website/docs/index.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        return Provider(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

