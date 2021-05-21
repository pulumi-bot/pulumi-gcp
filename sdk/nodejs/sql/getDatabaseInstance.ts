// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about a Cloud SQL instance
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const qa = pulumi.output(gcp.sql.getDatabaseInstance({
 *     name: "test-sql-instance",
 * }, { async: true }));
 * ```
 */
export function getDatabaseInstance(args: GetDatabaseInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseInstanceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:sql/getDatabaseInstance:getDatabaseInstance", {
        "name": args.name,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabaseInstance.
 */
export interface GetDatabaseInstanceArgs {
    /**
     * The name of the instance.
     */
    name: string;
    /**
     * The ID of the project in which the resource belongs.
     */
    project?: string;
}

/**
 * A collection of values returned by getDatabaseInstance.
 */
export interface GetDatabaseInstanceResult {
    readonly clones: outputs.sql.GetDatabaseInstanceClone[];
    /**
     * The connection name of the instance to be used in connection strings.
     */
    readonly connectionName: string;
    /**
     * The MySQL, PostgreSQL or SQL Server (beta) version to use.
     */
    readonly databaseVersion: string;
    readonly deletionProtection: boolean;
    /**
     * The full path to the encryption key used for the CMEK disk encryption.
     */
    readonly encryptionKeyName: string;
    /**
     * The first IPv4 address of any type assigned.
     */
    readonly firstIpAddress: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipAddresses: outputs.sql.GetDatabaseInstanceIpAddress[];
    /**
     * The name of the instance that will act as
     * the master in the replication setup.
     */
    readonly masterInstanceName: string;
    /**
     * A name for this whitelist entry.
     */
    readonly name: string;
    /**
     * The first private (`PRIVATE`) IPv4 address assigned.
     */
    readonly privateIpAddress: string;
    readonly project?: string;
    /**
     * The first public (`PRIMARY`) IPv4 address assigned.
     */
    readonly publicIpAddress: string;
    readonly region: string;
    /**
     * The configuration for replication. The
     * configuration is detailed below.
     */
    readonly replicaConfigurations: outputs.sql.GetDatabaseInstanceReplicaConfiguration[];
    readonly restoreBackupContexts: outputs.sql.GetDatabaseInstanceRestoreBackupContext[];
    /**
     * Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
     */
    readonly rootPassword: string;
    /**
     * The URI of the created resource.
     */
    readonly selfLink: string;
    readonly serverCaCerts: outputs.sql.GetDatabaseInstanceServerCaCert[];
    /**
     * The service account email address assigned to the instance.
     */
    readonly serviceAccountEmailAddress: string;
    /**
     * The settings to use for the database. The
     * configuration is detailed below.
     */
    readonly settings: outputs.sql.GetDatabaseInstanceSetting[];
}

export function getDatabaseInstanceOutput(args: GetDatabaseInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabaseInstanceResult> {
    return pulumi.output(args).apply(a => getDatabaseInstance(a, opts))
}

/**
 * A collection of arguments for invoking getDatabaseInstance.
 */
export interface GetDatabaseInstanceOutputArgs {
    /**
     * The name of the instance.
     */
    name: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     */
    project?: pulumi.Input<string>;
}
