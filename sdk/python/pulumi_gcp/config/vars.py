# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

__config__ = pulumi.Config('gcp')

access_context_manager_custom_endpoint = __config__.get('accessContextManagerCustomEndpoint')

access_token = __config__.get('accessToken')

app_engine_custom_endpoint = __config__.get('appEngineCustomEndpoint')

batching = __config__.get('batching')

big_query_custom_endpoint = __config__.get('bigQueryCustomEndpoint')

bigquery_data_transfer_custom_endpoint = __config__.get('bigqueryDataTransferCustomEndpoint')

bigtable_custom_endpoint = __config__.get('bigtableCustomEndpoint')

billing_custom_endpoint = __config__.get('billingCustomEndpoint')

binary_authorization_custom_endpoint = __config__.get('binaryAuthorizationCustomEndpoint')

cloud_billing_custom_endpoint = __config__.get('cloudBillingCustomEndpoint')

cloud_build_custom_endpoint = __config__.get('cloudBuildCustomEndpoint')

cloud_functions_custom_endpoint = __config__.get('cloudFunctionsCustomEndpoint')

cloud_iot_custom_endpoint = __config__.get('cloudIotCustomEndpoint')

cloud_run_custom_endpoint = __config__.get('cloudRunCustomEndpoint')

cloud_scheduler_custom_endpoint = __config__.get('cloudSchedulerCustomEndpoint')

cloud_tasks_custom_endpoint = __config__.get('cloudTasksCustomEndpoint')

composer_custom_endpoint = __config__.get('composerCustomEndpoint')

compute_beta_custom_endpoint = __config__.get('computeBetaCustomEndpoint')

compute_custom_endpoint = __config__.get('computeCustomEndpoint')

container_analysis_custom_endpoint = __config__.get('containerAnalysisCustomEndpoint')

container_beta_custom_endpoint = __config__.get('containerBetaCustomEndpoint')

container_custom_endpoint = __config__.get('containerCustomEndpoint')

credentials = __config__.get('credentials') or utilities.get_env('GOOGLE_CREDENTIALS', 'GOOGLE_CLOUD_KEYFILE_JSON', 'GCLOUD_KEYFILE_JSON')

data_fusion_custom_endpoint = __config__.get('dataFusionCustomEndpoint')

dataflow_custom_endpoint = __config__.get('dataflowCustomEndpoint')

dataproc_beta_custom_endpoint = __config__.get('dataprocBetaCustomEndpoint')

dataproc_custom_endpoint = __config__.get('dataprocCustomEndpoint')

deployment_manager_custom_endpoint = __config__.get('deploymentManagerCustomEndpoint')

dns_beta_custom_endpoint = __config__.get('dnsBetaCustomEndpoint')

dns_custom_endpoint = __config__.get('dnsCustomEndpoint')

filestore_custom_endpoint = __config__.get('filestoreCustomEndpoint')

firestore_custom_endpoint = __config__.get('firestoreCustomEndpoint')

healthcare_custom_endpoint = __config__.get('healthcareCustomEndpoint')

iam_credentials_custom_endpoint = __config__.get('iamCredentialsCustomEndpoint')

iam_custom_endpoint = __config__.get('iamCustomEndpoint')

iap_custom_endpoint = __config__.get('iapCustomEndpoint')

identity_platform_custom_endpoint = __config__.get('identityPlatformCustomEndpoint')

kms_custom_endpoint = __config__.get('kmsCustomEndpoint')

logging_custom_endpoint = __config__.get('loggingCustomEndpoint')

ml_engine_custom_endpoint = __config__.get('mlEngineCustomEndpoint')

monitoring_custom_endpoint = __config__.get('monitoringCustomEndpoint')

project = __config__.get('project') or utilities.get_env('GOOGLE_PROJECT', 'GOOGLE_CLOUD_PROJECT', 'GCLOUD_PROJECT', 'CLOUDSDK_CORE_PROJECT')

pubsub_custom_endpoint = __config__.get('pubsubCustomEndpoint')

redis_custom_endpoint = __config__.get('redisCustomEndpoint')

region = __config__.get('region') or utilities.get_env('GOOGLE_REGION', 'GCLOUD_REGION', 'CLOUDSDK_COMPUTE_REGION')

request_timeout = __config__.get('requestTimeout')

resource_manager_custom_endpoint = __config__.get('resourceManagerCustomEndpoint')

resource_manager_v2beta1_custom_endpoint = __config__.get('resourceManagerV2beta1CustomEndpoint')

runtime_config_custom_endpoint = __config__.get('runtimeConfigCustomEndpoint')

runtimeconfig_custom_endpoint = __config__.get('runtimeconfigCustomEndpoint')

scopes = __config__.get('scopes')

security_center_custom_endpoint = __config__.get('securityCenterCustomEndpoint')

security_scanner_custom_endpoint = __config__.get('securityScannerCustomEndpoint')

service_management_custom_endpoint = __config__.get('serviceManagementCustomEndpoint')

service_networking_custom_endpoint = __config__.get('serviceNetworkingCustomEndpoint')

service_usage_custom_endpoint = __config__.get('serviceUsageCustomEndpoint')

source_repo_custom_endpoint = __config__.get('sourceRepoCustomEndpoint')

spanner_custom_endpoint = __config__.get('spannerCustomEndpoint')

sql_custom_endpoint = __config__.get('sqlCustomEndpoint')

storage_custom_endpoint = __config__.get('storageCustomEndpoint')

storage_transfer_custom_endpoint = __config__.get('storageTransferCustomEndpoint')

tpu_custom_endpoint = __config__.get('tpuCustomEndpoint')

user_project_override = __config__.get('userProjectOverride')

vpc_access_custom_endpoint = __config__.get('vpcAccessCustomEndpoint')

zone = __config__.get('zone') or utilities.get_env('GOOGLE_ZONE', 'GCLOUD_ZONE', 'CLOUDSDK_COMPUTE_ZONE')

