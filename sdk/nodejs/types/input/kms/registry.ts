// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface RegistryCredential {
    publicKeyCertificate?: pulumi.Input<inputApi.kms.RegistryCredentialPublicKeyCertificate>;
}

export interface RegistryCredentialPublicKeyCertificate {
    certificate: pulumi.Input<string>;
    format: pulumi.Input<string>;
}

export interface RegistryEventNotificationConfig {
    pubsubTopicName: pulumi.Input<string>;
}

export interface RegistryHttpConfig {
    httpEnabledState: pulumi.Input<string>;
}

export interface RegistryMqttConfig {
    mqttEnabledState: pulumi.Input<string>;
}

export interface RegistryStateNotificationConfig {
    pubsubTopicName: pulumi.Input<string>;
}
