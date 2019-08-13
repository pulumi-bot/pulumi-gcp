// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface ManagedZoneDnssecConfig {
    defaultKeySpecs: outputApi.dns.ManagedZoneDnssecConfigDefaultKeySpec[];
    kind?: string;
    nonExistence: string;
    state?: string;
}

export interface ManagedZoneDnssecConfigDefaultKeySpec {
    algorithm?: string;
    keyLength?: number;
    keyType?: string;
    kind?: string;
}

export interface ManagedZoneForwardingConfig {
    targetNameServers?: outputApi.dns.ManagedZoneForwardingConfigTargetNameServer[];
}

export interface ManagedZoneForwardingConfigTargetNameServer {
    ipv4Address?: string;
}

export interface ManagedZonePeeringConfig {
    targetNetwork?: outputApi.dns.ManagedZonePeeringConfigTargetNetwork;
}

export interface ManagedZonePeeringConfigTargetNetwork {
    networkUrl?: string;
}

export interface ManagedZonePrivateVisibilityConfig {
    networks?: outputApi.dns.ManagedZonePrivateVisibilityConfigNetwork[];
}

export interface ManagedZonePrivateVisibilityConfigNetwork {
    networkUrl?: string;
}
