// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface DatabaseInstanceIpAddress {
    ipAddress?: pulumi.Input<string>;
    timeToRetire?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
}

export interface DatabaseInstanceReplicaConfiguration {
    caCertificate?: pulumi.Input<string>;
    clientCertificate?: pulumi.Input<string>;
    clientKey?: pulumi.Input<string>;
    connectRetryInterval?: pulumi.Input<number>;
    dumpFilePath?: pulumi.Input<string>;
    failoverTarget?: pulumi.Input<boolean>;
    masterHeartbeatPeriod?: pulumi.Input<number>;
    password?: pulumi.Input<string>;
    sslCipher?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    verifyServerCertificate?: pulumi.Input<boolean>;
}

export interface DatabaseInstanceServerCaCert {
    cert?: pulumi.Input<string>;
    commonName?: pulumi.Input<string>;
    createTime?: pulumi.Input<string>;
    expirationTime?: pulumi.Input<string>;
    sha1Fingerprint?: pulumi.Input<string>;
}

export interface DatabaseInstanceSettings {
    activationPolicy?: pulumi.Input<string>;
    authorizedGaeApplications?: pulumi.Input<pulumi.Input<string>[]>;
    availabilityType?: pulumi.Input<string>;
    backupConfiguration?: pulumi.Input<inputApi.sql.DatabaseInstanceSettingsBackupConfiguration>;
    crashSafeReplication?: pulumi.Input<boolean>;
    databaseFlags?: pulumi.Input<pulumi.Input<inputApi.sql.DatabaseInstanceSettingsDatabaseFlag>[]>;
    diskAutoresize?: pulumi.Input<boolean>;
    diskSize?: pulumi.Input<number>;
    diskType?: pulumi.Input<string>;
    ipConfiguration?: pulumi.Input<inputApi.sql.DatabaseInstanceSettingsIpConfiguration>;
    locationPreference?: pulumi.Input<inputApi.sql.DatabaseInstanceSettingsLocationPreference>;
    maintenanceWindow?: pulumi.Input<inputApi.sql.DatabaseInstanceSettingsMaintenanceWindow>;
    pricingPlan?: pulumi.Input<string>;
    replicationType?: pulumi.Input<string>;
    tier: pulumi.Input<string>;
    userLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    version?: pulumi.Input<number>;
}

export interface DatabaseInstanceSettingsBackupConfiguration {
    binaryLogEnabled?: pulumi.Input<boolean>;
    enabled?: pulumi.Input<boolean>;
    startTime?: pulumi.Input<string>;
}

export interface DatabaseInstanceSettingsDatabaseFlag {
    /**
     * The name of the instance. If the name is left
     * blank, this provider will randomly generate one when the instance is first
     * created. This is done because after a name is used, it cannot be reused for
     * up to [one week](https://cloud.google.com/sql/docs/delete-instance).
     */
    name?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface DatabaseInstanceSettingsIpConfiguration {
    authorizedNetworks?: pulumi.Input<pulumi.Input<inputApi.sql.DatabaseInstanceSettingsIpConfigurationAuthorizedNetwork>[]>;
    ipv4Enabled?: pulumi.Input<boolean>;
    privateNetwork?: pulumi.Input<string>;
    requireSsl?: pulumi.Input<boolean>;
}

export interface DatabaseInstanceSettingsIpConfigurationAuthorizedNetwork {
    expirationTime?: pulumi.Input<string>;
    /**
     * The name of the instance. If the name is left
     * blank, this provider will randomly generate one when the instance is first
     * created. This is done because after a name is used, it cannot be reused for
     * up to [one week](https://cloud.google.com/sql/docs/delete-instance).
     */
    name?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface DatabaseInstanceSettingsLocationPreference {
    followGaeApplication?: pulumi.Input<string>;
    zone?: pulumi.Input<string>;
}

export interface DatabaseInstanceSettingsMaintenanceWindow {
    day?: pulumi.Input<number>;
    hour?: pulumi.Input<number>;
    updateTrack?: pulumi.Input<string>;
}
