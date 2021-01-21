# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'access_approval_custom_endpoint',
    'access_context_manager_custom_endpoint',
    'access_token',
    'active_directory_custom_endpoint',
    'api_gateway_custom_endpoint',
    'apigee_custom_endpoint',
    'app_engine_custom_endpoint',
    'artifact_registry_custom_endpoint',
    'batching',
    'big_query_custom_endpoint',
    'bigquery_connection_custom_endpoint',
    'bigquery_data_transfer_custom_endpoint',
    'bigquery_reservation_custom_endpoint',
    'bigtable_custom_endpoint',
    'billing_custom_endpoint',
    'billing_project',
    'binary_authorization_custom_endpoint',
    'cloud_asset_custom_endpoint',
    'cloud_billing_custom_endpoint',
    'cloud_build_custom_endpoint',
    'cloud_functions_custom_endpoint',
    'cloud_identity_custom_endpoint',
    'cloud_iot_custom_endpoint',
    'cloud_run_custom_endpoint',
    'cloud_scheduler_custom_endpoint',
    'cloud_tasks_custom_endpoint',
    'composer_custom_endpoint',
    'compute_beta_custom_endpoint',
    'compute_custom_endpoint',
    'container_analysis_custom_endpoint',
    'container_beta_custom_endpoint',
    'container_custom_endpoint',
    'credentials',
    'data_catalog_custom_endpoint',
    'data_fusion_custom_endpoint',
    'data_loss_prevention_custom_endpoint',
    'dataflow_custom_endpoint',
    'dataproc_beta_custom_endpoint',
    'dataproc_custom_endpoint',
    'datastore_custom_endpoint',
    'deployment_manager_custom_endpoint',
    'dialogflow_custom_endpoint',
    'dns_beta_custom_endpoint',
    'dns_custom_endpoint',
    'filestore_custom_endpoint',
    'firebase_custom_endpoint',
    'firestore_custom_endpoint',
    'game_services_custom_endpoint',
    'healthcare_custom_endpoint',
    'iam_beta_custom_endpoint',
    'iam_credentials_custom_endpoint',
    'iam_custom_endpoint',
    'iap_custom_endpoint',
    'identity_platform_custom_endpoint',
    'impersonate_service_account',
    'impersonate_service_account_delegates',
    'kms_custom_endpoint',
    'logging_custom_endpoint',
    'memcache_custom_endpoint',
    'ml_engine_custom_endpoint',
    'monitoring_custom_endpoint',
    'network_management_custom_endpoint',
    'notebooks_custom_endpoint',
    'os_config_custom_endpoint',
    'os_login_custom_endpoint',
    'project',
    'pubsub_custom_endpoint',
    'pubsub_lite_custom_endpoint',
    'redis_custom_endpoint',
    'region',
    'request_timeout',
    'resource_manager_custom_endpoint',
    'resource_manager_v2beta1_custom_endpoint',
    'runtime_config_custom_endpoint',
    'runtimeconfig_custom_endpoint',
    'scopes',
    'secret_manager_custom_endpoint',
    'security_center_custom_endpoint',
    'security_scanner_custom_endpoint',
    'service_directory_custom_endpoint',
    'service_management_custom_endpoint',
    'service_networking_custom_endpoint',
    'service_usage_custom_endpoint',
    'source_repo_custom_endpoint',
    'spanner_custom_endpoint',
    'sql_custom_endpoint',
    'storage_custom_endpoint',
    'storage_transfer_custom_endpoint',
    'tpu_custom_endpoint',
    'user_project_override',
    'vpc_access_custom_endpoint',
    'zone',
]

__config__ = pulumi.Config('gcp')

access_approval_custom_endpoint = __config__.get('accessApprovalCustomEndpoint')

access_context_manager_custom_endpoint = __config__.get('accessContextManagerCustomEndpoint')

access_token = __config__.get('accessToken')

active_directory_custom_endpoint = __config__.get('activeDirectoryCustomEndpoint')

api_gateway_custom_endpoint = __config__.get('apiGatewayCustomEndpoint')

apigee_custom_endpoint = __config__.get('apigeeCustomEndpoint')

app_engine_custom_endpoint = __config__.get('appEngineCustomEndpoint')

artifact_registry_custom_endpoint = __config__.get('artifactRegistryCustomEndpoint')

batching = __config__.get('batching')

big_query_custom_endpoint = __config__.get('bigQueryCustomEndpoint')

bigquery_connection_custom_endpoint = __config__.get('bigqueryConnectionCustomEndpoint')

bigquery_data_transfer_custom_endpoint = __config__.get('bigqueryDataTransferCustomEndpoint')

bigquery_reservation_custom_endpoint = __config__.get('bigqueryReservationCustomEndpoint')

bigtable_custom_endpoint = __config__.get('bigtableCustomEndpoint')

billing_custom_endpoint = __config__.get('billingCustomEndpoint')

billing_project = __config__.get('billingProject')

binary_authorization_custom_endpoint = __config__.get('binaryAuthorizationCustomEndpoint')

cloud_asset_custom_endpoint = __config__.get('cloudAssetCustomEndpoint')

cloud_billing_custom_endpoint = __config__.get('cloudBillingCustomEndpoint')

cloud_build_custom_endpoint = __config__.get('cloudBuildCustomEndpoint')

cloud_functions_custom_endpoint = __config__.get('cloudFunctionsCustomEndpoint')

cloud_identity_custom_endpoint = __config__.get('cloudIdentityCustomEndpoint')

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

credentials = __config__.get('credentials') or _utilities.get_env('GOOGLE_CREDENTIALS', 'GOOGLE_CLOUD_KEYFILE_JSON', 'GCLOUD_KEYFILE_JSON')

data_catalog_custom_endpoint = __config__.get('dataCatalogCustomEndpoint')

data_fusion_custom_endpoint = __config__.get('dataFusionCustomEndpoint')

data_loss_prevention_custom_endpoint = __config__.get('dataLossPreventionCustomEndpoint')

dataflow_custom_endpoint = __config__.get('dataflowCustomEndpoint')

dataproc_beta_custom_endpoint = __config__.get('dataprocBetaCustomEndpoint')

dataproc_custom_endpoint = __config__.get('dataprocCustomEndpoint')

datastore_custom_endpoint = __config__.get('datastoreCustomEndpoint')

deployment_manager_custom_endpoint = __config__.get('deploymentManagerCustomEndpoint')

dialogflow_custom_endpoint = __config__.get('dialogflowCustomEndpoint')

dns_beta_custom_endpoint = __config__.get('dnsBetaCustomEndpoint')

dns_custom_endpoint = __config__.get('dnsCustomEndpoint')

filestore_custom_endpoint = __config__.get('filestoreCustomEndpoint')

firebase_custom_endpoint = __config__.get('firebaseCustomEndpoint')

firestore_custom_endpoint = __config__.get('firestoreCustomEndpoint')

game_services_custom_endpoint = __config__.get('gameServicesCustomEndpoint')

healthcare_custom_endpoint = __config__.get('healthcareCustomEndpoint')

iam_beta_custom_endpoint = __config__.get('iamBetaCustomEndpoint')

iam_credentials_custom_endpoint = __config__.get('iamCredentialsCustomEndpoint')

iam_custom_endpoint = __config__.get('iamCustomEndpoint')

iap_custom_endpoint = __config__.get('iapCustomEndpoint')

identity_platform_custom_endpoint = __config__.get('identityPlatformCustomEndpoint')

impersonate_service_account = __config__.get('impersonateServiceAccount')

impersonate_service_account_delegates = __config__.get('impersonateServiceAccountDelegates')

kms_custom_endpoint = __config__.get('kmsCustomEndpoint')

logging_custom_endpoint = __config__.get('loggingCustomEndpoint')

memcache_custom_endpoint = __config__.get('memcacheCustomEndpoint')

ml_engine_custom_endpoint = __config__.get('mlEngineCustomEndpoint')

monitoring_custom_endpoint = __config__.get('monitoringCustomEndpoint')

network_management_custom_endpoint = __config__.get('networkManagementCustomEndpoint')

notebooks_custom_endpoint = __config__.get('notebooksCustomEndpoint')

os_config_custom_endpoint = __config__.get('osConfigCustomEndpoint')

os_login_custom_endpoint = __config__.get('osLoginCustomEndpoint')

project = __config__.get('project') or _utilities.get_env('GOOGLE_PROJECT', 'GOOGLE_CLOUD_PROJECT', 'GCLOUD_PROJECT', 'CLOUDSDK_CORE_PROJECT')

pubsub_custom_endpoint = __config__.get('pubsubCustomEndpoint')

pubsub_lite_custom_endpoint = __config__.get('pubsubLiteCustomEndpoint')

redis_custom_endpoint = __config__.get('redisCustomEndpoint')

region = __config__.get('region') or _utilities.get_env('GOOGLE_REGION', 'GCLOUD_REGION', 'CLOUDSDK_COMPUTE_REGION')

request_timeout = __config__.get('requestTimeout')

resource_manager_custom_endpoint = __config__.get('resourceManagerCustomEndpoint')

resource_manager_v2beta1_custom_endpoint = __config__.get('resourceManagerV2beta1CustomEndpoint')

runtime_config_custom_endpoint = __config__.get('runtimeConfigCustomEndpoint')

runtimeconfig_custom_endpoint = __config__.get('runtimeconfigCustomEndpoint')

scopes = __config__.get('scopes')

secret_manager_custom_endpoint = __config__.get('secretManagerCustomEndpoint')

security_center_custom_endpoint = __config__.get('securityCenterCustomEndpoint')

security_scanner_custom_endpoint = __config__.get('securityScannerCustomEndpoint')

service_directory_custom_endpoint = __config__.get('serviceDirectoryCustomEndpoint')

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

zone = __config__.get('zone') or _utilities.get_env('GOOGLE_ZONE', 'GCLOUD_ZONE', 'CLOUDSDK_COMPUTE_ZONE')

