# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'DomainMappingMetadata',
    'DomainMappingSpec',
    'DomainMappingStatus',
    'DomainMappingStatusCondition',
    'DomainMappingStatusResourceRecord',
    'IamBindingCondition',
    'IamMemberCondition',
    'ServiceMetadata',
    'ServiceStatus',
    'ServiceStatusCondition',
    'ServiceTemplate',
    'ServiceTemplateMetadata',
    'ServiceTemplateSpec',
    'ServiceTemplateSpecContainer',
    'ServiceTemplateSpecContainerEnv',
    'ServiceTemplateSpecContainerEnvFrom',
    'ServiceTemplateSpecContainerEnvFromConfigMapRef',
    'ServiceTemplateSpecContainerEnvFromConfigMapRefLocalObjectReference',
    'ServiceTemplateSpecContainerEnvFromSecretRef',
    'ServiceTemplateSpecContainerEnvFromSecretRefLocalObjectReference',
    'ServiceTemplateSpecContainerResources',
    'ServiceTraffic',
]

@pulumi.output_type
class DomainMappingMetadata(dict):
    annotations: Optional[Dict[str, str]] = pulumi.output_property("annotations")
    """
    Annotations is a key value map stored with a resource that
    may be set by external tools to store and retrieve arbitrary metadata. More
    info: http://kubernetes.io/docs/user-guide/annotations
    """
    generation: Optional[float] = pulumi.output_property("generation")
    """
    -
    A sequence number representing a specific generation of the desired state.
    """
    labels: Optional[Dict[str, str]] = pulumi.output_property("labels")
    """
    Map of string keys and values that can be used to organize and categorize
    (scope and select) objects. May match selectors of replication controllers
    and routes.
    More info: http://kubernetes.io/docs/user-guide/labels
    """
    namespace: str = pulumi.output_property("namespace")
    """
    In Cloud Run the namespace must be equal to either the
    project ID or project number.
    """
    resource_version: Optional[str] = pulumi.output_property("resourceVersion")
    """
    -
    An opaque value that represents the internal version of this object that
    can be used by clients to determine when objects have changed. May be used
    for optimistic concurrency, change detection, and the watch operation on a
    resource or set of resources. They may only be valid for a
    particular resource or set of resources.
    More info:
    https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
    """
    self_link: Optional[str] = pulumi.output_property("selfLink")
    """
    -
    SelfLink is a URL representing this object.
    """
    uid: Optional[str] = pulumi.output_property("uid")
    """
    -
    UID is a unique id generated by the server on successful creation of a resource and is not
    allowed to change on PUT operations.
    More info: http://kubernetes.io/docs/user-guide/identifiers#uids
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DomainMappingSpec(dict):
    certificate_mode: Optional[str] = pulumi.output_property("certificateMode")
    """
    The mode of the certificate.
    """
    force_override: Optional[bool] = pulumi.output_property("forceOverride")
    """
    If set, the mapping will override any mapping set before this spec was set.
    It is recommended that the user leaves this empty to receive an error
    warning about a potential conflict and only set it once the respective UI
    has given such a warning.
    """
    route_name: str = pulumi.output_property("routeName")
    """
    The name of the Cloud Run Service that this DomainMapping applies to.
    The route must exist.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DomainMappingStatus(dict):
    conditions: Optional[List['outputs.DomainMappingStatusCondition']] = pulumi.output_property("conditions")
    mapped_route_name: Optional[str] = pulumi.output_property("mappedRouteName")
    observed_generation: Optional[float] = pulumi.output_property("observedGeneration")
    resource_records: Optional[List['outputs.DomainMappingStatusResourceRecord']] = pulumi.output_property("resourceRecords")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DomainMappingStatusCondition(dict):
    message: Optional[str] = pulumi.output_property("message")
    reason: Optional[str] = pulumi.output_property("reason")
    status: Optional[str] = pulumi.output_property("status")
    type: Optional[str] = pulumi.output_property("type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DomainMappingStatusResourceRecord(dict):
    name: Optional[str] = pulumi.output_property("name")
    """
    Name should be a verified domain
    """
    rrdata: Optional[str] = pulumi.output_property("rrdata")
    type: Optional[str] = pulumi.output_property("type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IamBindingCondition(dict):
    description: Optional[str] = pulumi.output_property("description")
    expression: str = pulumi.output_property("expression")
    title: str = pulumi.output_property("title")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IamMemberCondition(dict):
    description: Optional[str] = pulumi.output_property("description")
    expression: str = pulumi.output_property("expression")
    title: str = pulumi.output_property("title")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceMetadata(dict):
    annotations: Optional[Dict[str, str]] = pulumi.output_property("annotations")
    """
    Annotations is a key value map stored with a resource that
    may be set by external tools to store and retrieve arbitrary metadata. More
    info: http://kubernetes.io/docs/user-guide/annotations
    """
    generation: Optional[float] = pulumi.output_property("generation")
    """
    -
    A sequence number representing a specific generation of the desired state.
    """
    labels: Optional[Dict[str, str]] = pulumi.output_property("labels")
    """
    Map of string keys and values that can be used to organize and categorize
    (scope and select) objects. May match selectors of replication controllers
    and routes.
    More info: http://kubernetes.io/docs/user-guide/labels
    """
    namespace: Optional[str] = pulumi.output_property("namespace")
    """
    In Cloud Run the namespace must be equal to either the
    project ID or project number.
    """
    resource_version: Optional[str] = pulumi.output_property("resourceVersion")
    """
    -
    An opaque value that represents the internal version of this object that
    can be used by clients to determine when objects have changed. May be used
    for optimistic concurrency, change detection, and the watch operation on a
    resource or set of resources. They may only be valid for a
    particular resource or set of resources.
    More info:
    https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
    """
    self_link: Optional[str] = pulumi.output_property("selfLink")
    """
    -
    SelfLink is a URL representing this object.
    """
    uid: Optional[str] = pulumi.output_property("uid")
    """
    -
    UID is a unique id generated by the server on successful creation of a resource and is not
    allowed to change on PUT operations.
    More info: http://kubernetes.io/docs/user-guide/identifiers#uids
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceStatus(dict):
    conditions: Optional[List['outputs.ServiceStatusCondition']] = pulumi.output_property("conditions")
    latest_created_revision_name: Optional[str] = pulumi.output_property("latestCreatedRevisionName")
    latest_ready_revision_name: Optional[str] = pulumi.output_property("latestReadyRevisionName")
    observed_generation: Optional[float] = pulumi.output_property("observedGeneration")
    url: Optional[str] = pulumi.output_property("url")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceStatusCondition(dict):
    message: Optional[str] = pulumi.output_property("message")
    reason: Optional[str] = pulumi.output_property("reason")
    status: Optional[str] = pulumi.output_property("status")
    type: Optional[str] = pulumi.output_property("type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplate(dict):
    metadata: Optional['outputs.ServiceTemplateMetadata'] = pulumi.output_property("metadata")
    """
    Metadata associated with this Service, including name, namespace, labels,
    and annotations.  Structure is documented below.
    """
    spec: Optional['outputs.ServiceTemplateSpec'] = pulumi.output_property("spec")
    """
    RevisionSpec holds the desired state of the Revision (from the client).  Structure is documented below.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateMetadata(dict):
    annotations: Optional[Dict[str, str]] = pulumi.output_property("annotations")
    """
    Annotations is a key value map stored with a resource that
    may be set by external tools to store and retrieve arbitrary metadata. More
    info: http://kubernetes.io/docs/user-guide/annotations
    """
    generation: Optional[float] = pulumi.output_property("generation")
    """
    -
    A sequence number representing a specific generation of the desired state.
    """
    labels: Optional[Dict[str, str]] = pulumi.output_property("labels")
    """
    Map of string keys and values that can be used to organize and categorize
    (scope and select) objects. May match selectors of replication controllers
    and routes.
    More info: http://kubernetes.io/docs/user-guide/labels
    """
    name: Optional[str] = pulumi.output_property("name")
    """
    Name of the environment variable.
    """
    namespace: Optional[str] = pulumi.output_property("namespace")
    """
    In Cloud Run the namespace must be equal to either the
    project ID or project number.
    """
    resource_version: Optional[str] = pulumi.output_property("resourceVersion")
    """
    -
    An opaque value that represents the internal version of this object that
    can be used by clients to determine when objects have changed. May be used
    for optimistic concurrency, change detection, and the watch operation on a
    resource or set of resources. They may only be valid for a
    particular resource or set of resources.
    More info:
    https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
    """
    self_link: Optional[str] = pulumi.output_property("selfLink")
    """
    -
    SelfLink is a URL representing this object.
    """
    uid: Optional[str] = pulumi.output_property("uid")
    """
    -
    UID is a unique id generated by the server on successful creation of a resource and is not
    allowed to change on PUT operations.
    More info: http://kubernetes.io/docs/user-guide/identifiers#uids
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpec(dict):
    container_concurrency: Optional[float] = pulumi.output_property("containerConcurrency")
    """
    ContainerConcurrency specifies the maximum allowed in-flight (concurrent)
    requests per container of the Revision. Values are:
    - `0` thread-safe, the system should manage the max concurrency. This is
    the default value.
    - `1` not-thread-safe. Single concurrency
    - `2-N` thread-safe, max concurrency of N
    """
    containers: Optional[List['outputs.ServiceTemplateSpecContainer']] = pulumi.output_property("containers")
    """
    Container defines the unit of execution for this Revision.
    In the context of a Revision, we disallow a number of the fields of
    this Container, including: name, ports, and volumeMounts.
    The runtime contract is documented here:
    https://github.com/knative/serving/blob/master/docs/runtime-contract.md  Structure is documented below.
    """
    service_account_name: Optional[str] = pulumi.output_property("serviceAccountName")
    """
    Email address of the IAM service account associated with the revision of the
    service. The service account represents the identity of the running revision,
    and determines what permissions the revision has. If not provided, the revision
    will use the project's default service account.
    """
    serving_state: Optional[str] = pulumi.output_property("servingState")
    """
    -
    ServingState holds a value describing the state the resources
    are in for this Revision.
    It is expected
    that the system will manipulate this based on routability and load.
    """
    timeout_seconds: Optional[float] = pulumi.output_property("timeoutSeconds")
    """
    TimeoutSeconds holds the max duration the instance is allowed for responding to a request.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpecContainer(dict):
    args: Optional[List[str]] = pulumi.output_property("args")
    """
    Arguments to the entrypoint.
    The docker image's CMD is used if this is not provided.
    Variable references $(VAR_NAME) are expanded using the container's
    environment. If a variable cannot be resolved, the reference in the input
    string will be unchanged. The $(VAR_NAME) syntax can be escaped with a
    double $$, ie: $$(VAR_NAME). Escaped references will never be expanded,
    regardless of whether the variable exists or not.
    More info:
    https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
    """
    commands: Optional[List[str]] = pulumi.output_property("commands")
    """
    Entrypoint array. Not executed within a shell.
    The docker image's ENTRYPOINT is used if this is not provided.
    Variable references $(VAR_NAME) are expanded using the container's
    environment. If a variable cannot be resolved, the reference in the input
    string will be unchanged. The $(VAR_NAME) syntax can be escaped with a
    double $$, ie: $$(VAR_NAME). Escaped references will never be expanded,
    regardless of whether the variable exists or not.
    More info:
    https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
    """
    env_froms: Optional[List['outputs.ServiceTemplateSpecContainerEnvFrom']] = pulumi.output_property("envFroms")
    """
    -
    (Optional, Deprecated)
    List of sources to populate environment variables in the container.
    All invalid keys will be reported as an event when the container is starting.
    When a key exists in multiple sources, the value associated with the last source will
    take precedence. Values defined by an Env with a duplicate key will take
    precedence.  Structure is documented below.
    """
    envs: Optional[List['outputs.ServiceTemplateSpecContainerEnv']] = pulumi.output_property("envs")
    """
    List of environment variables to set in the container.  Structure is documented below.
    """
    image: str = pulumi.output_property("image")
    """
    Docker image name. This is most often a reference to a container located
    in the container registry, such as gcr.io/cloudrun/hello
    More info: https://kubernetes.io/docs/concepts/containers/images
    """
    resources: Optional['outputs.ServiceTemplateSpecContainerResources'] = pulumi.output_property("resources")
    """
    Compute Resources required by this container. Used to set values such as max memory
    More info:
    https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources  Structure is documented below.
    """
    working_dir: Optional[str] = pulumi.output_property("workingDir")
    """
    -
    (Optional, Deprecated)
    Container's working directory.
    If not specified, the container runtime's default will be used, which
    might be configured in the container image.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpecContainerEnv(dict):
    name: Optional[str] = pulumi.output_property("name")
    """
    Name of the environment variable.
    """
    value: Optional[str] = pulumi.output_property("value")
    """
    Variable references $(VAR_NAME) are expanded
    using the previous defined environment variables in the container and
    any route environment variables. If a variable cannot be resolved,
    the reference in the input string will be unchanged. The $(VAR_NAME)
    syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped
    references will never be expanded, regardless of whether the variable
    exists or not.
    Defaults to "".
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpecContainerEnvFrom(dict):
    config_map_ref: Optional['outputs.ServiceTemplateSpecContainerEnvFromConfigMapRef'] = pulumi.output_property("configMapRef")
    """
    The ConfigMap to select from.  Structure is documented below.
    """
    prefix: Optional[str] = pulumi.output_property("prefix")
    """
    An optional identifier to prepend to each key in the ConfigMap.
    """
    secret_ref: Optional['outputs.ServiceTemplateSpecContainerEnvFromSecretRef'] = pulumi.output_property("secretRef")
    """
    The Secret to select from.  Structure is documented below.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpecContainerEnvFromConfigMapRef(dict):
    local_object_reference: Optional['outputs.ServiceTemplateSpecContainerEnvFromConfigMapRefLocalObjectReference'] = pulumi.output_property("localObjectReference")
    """
    The Secret to select from.  Structure is documented below.
    """
    optional: Optional[bool] = pulumi.output_property("optional")
    """
    Specify whether the Secret must be defined
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpecContainerEnvFromConfigMapRefLocalObjectReference(dict):
    name: str = pulumi.output_property("name")
    """
    Name of the environment variable.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpecContainerEnvFromSecretRef(dict):
    local_object_reference: Optional['outputs.ServiceTemplateSpecContainerEnvFromSecretRefLocalObjectReference'] = pulumi.output_property("localObjectReference")
    """
    The Secret to select from.  Structure is documented below.
    """
    optional: Optional[bool] = pulumi.output_property("optional")
    """
    Specify whether the Secret must be defined
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpecContainerEnvFromSecretRefLocalObjectReference(dict):
    name: str = pulumi.output_property("name")
    """
    Name of the environment variable.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpecContainerResources(dict):
    limits: Optional[Dict[str, str]] = pulumi.output_property("limits")
    """
    Limits describes the maximum amount of compute resources allowed.
    The values of the map is string form of the 'quantity' k8s type:
    https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
    """
    requests: Optional[Dict[str, str]] = pulumi.output_property("requests")
    """
    Requests describes the minimum amount of compute resources required.
    If Requests is omitted for a container, it defaults to Limits if that is
    explicitly specified, otherwise to an implementation-defined value.
    The values of the map is string form of the 'quantity' k8s type:
    https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTraffic(dict):
    latest_revision: Optional[bool] = pulumi.output_property("latestRevision")
    """
    LatestRevision may be optionally provided to indicate that the latest ready
    Revision of the Configuration should be used for this traffic target. When
    provided LatestRevision must be true if RevisionName is empty; it must be
    false when RevisionName is non-empty.
    """
    percent: float = pulumi.output_property("percent")
    """
    Percent specifies percent of the traffic to this Revision or Configuration.
    """
    revision_name: Optional[str] = pulumi.output_property("revisionName")
    """
    RevisionName of a specific revision to which to send this portion of traffic.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


