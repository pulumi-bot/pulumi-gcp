// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Database instances can be imported using one of any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:sql/databaseInstance:DatabaseInstance master projects/{{project}}/instances/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:sql/databaseInstance:DatabaseInstance master {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:sql/databaseInstance:DatabaseInstance master {{name}}
 * ```
 *
 *  config and set on the server. When importing, double-check that your config has all the fields set that you expect- just seeing no diff isn't sufficient to know that your config could reproduce the imported resource.
 */
export class DatabaseInstance extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseInstanceState, opts?: pulumi.CustomResourceOptions): DatabaseInstance {
        return new DatabaseInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:sql/databaseInstance:DatabaseInstance';

    /**
     * Returns true if the given object is an instance of DatabaseInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseInstance.__pulumiType;
    }

    /**
     * Configuration for creating a new instance as a clone of another instance.
     */
    public readonly clone!: pulumi.Output<outputs.sql.DatabaseInstanceClone | undefined>;
    /**
     * The connection name of the instance to be used in
     * connection strings. For example, when connecting with [Cloud SQL Proxy](https://cloud.google.com/sql/docs/mysql/connect-admin-proxy).
     */
    public /*out*/ readonly connectionName!: pulumi.Output<string>;
    /**
     * The MySQL, PostgreSQL or
     * SQL Server (beta) version to use. Supported values include `MYSQL_5_6`,
     * `MYSQL_5_7`, `MYSQL_8_0`, `POSTGRES_9_6`,`POSTGRES_10`, `POSTGRES_11`,
     * `POSTGRES_12`, `POSTGRES_13`, `SQLSERVER_2017_STANDARD`,
     * `SQLSERVER_2017_ENTERPRISE`, `SQLSERVER_2017_EXPRESS`, `SQLSERVER_2017_WEB`.
     * [Database Version Policies](https://cloud.google.com/sql/docs/db-versions)
     * includes an up-to-date reference of supported versions.
     */
    public readonly databaseVersion!: pulumi.Output<string | undefined>;
    /**
     * Used to block Terraform from deleting a SQL Instance.
     */
    public readonly deletionProtection!: pulumi.Output<boolean | undefined>;
    /**
     * The full path to the encryption key used for the CMEK disk encryption.  Setting
     * up disk encryption currently requires manual steps outside of this provider.
     * The provided key must be in the same region as the SQL instance.  In order
     * to use this feature, a special kind of service account must be created and
     * granted permission on this key.  This step can currently only be done
     * manually, please see [this step](https://cloud.google.com/sql/docs/mysql/configure-cmek#service-account).
     * That service account needs the `Cloud KMS > Cloud KMS CryptoKey Encrypter/Decrypter` role on your
     * key - please see [this step](https://cloud.google.com/sql/docs/mysql/configure-cmek#grantkey).
     */
    public readonly encryptionKeyName!: pulumi.Output<string>;
    /**
     * The first IPv4 address of any type assigned.
     */
    public /*out*/ readonly firstIpAddress!: pulumi.Output<string>;
    public /*out*/ readonly ipAddresses!: pulumi.Output<outputs.sql.DatabaseInstanceIpAddress[]>;
    /**
     * The name of the instance that will act as
     * the master in the replication setup. Note, this requires the master to have
     * `binaryLogEnabled` set, as well as existing backups.
     */
    public readonly masterInstanceName!: pulumi.Output<string>;
    /**
     * A name for this whitelist entry.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The first private (`PRIVATE`) IPv4 address assigned.
     */
    public /*out*/ readonly privateIpAddress!: pulumi.Output<string>;
    /**
     * The full project ID of the source instance.`
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The first public (`PRIMARY`) IPv4 address assigned.
     */
    public /*out*/ readonly publicIpAddress!: pulumi.Output<string>;
    /**
     * The region the instance will sit in. Note, Cloud SQL is not
     * available in all regions - choose from one of the options listed [here](https://cloud.google.com/sql/docs/mysql/instance-locations).
     * A valid region must be provided to use this resource. If a region is not provided in the resource definition,
     * the provider region will be used instead, but this will be an apply-time error for instances if the provider
     * region is not supported with Cloud SQL. If you choose not to provide the `region` argument for this resource,
     * make sure you understand this.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The configuration for replication. The
     * configuration is detailed below.
     */
    public readonly replicaConfiguration!: pulumi.Output<outputs.sql.DatabaseInstanceReplicaConfiguration>;
    public readonly restoreBackupContext!: pulumi.Output<outputs.sql.DatabaseInstanceRestoreBackupContext | undefined>;
    /**
     * Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
     */
    public readonly rootPassword!: pulumi.Output<string | undefined>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    public /*out*/ readonly serverCaCerts!: pulumi.Output<outputs.sql.DatabaseInstanceServerCaCert[]>;
    /**
     * The service account email address assigned to the
     * instance.
     */
    public /*out*/ readonly serviceAccountEmailAddress!: pulumi.Output<string>;
    /**
     * The settings to use for the database. The
     * configuration is detailed below. Required if `clone` is not set.
     */
    public readonly settings!: pulumi.Output<outputs.sql.DatabaseInstanceSettings>;

    /**
     * Create a DatabaseInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DatabaseInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseInstanceArgs | DatabaseInstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseInstanceState | undefined;
            inputs["clone"] = state ? state.clone : undefined;
            inputs["connectionName"] = state ? state.connectionName : undefined;
            inputs["databaseVersion"] = state ? state.databaseVersion : undefined;
            inputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            inputs["encryptionKeyName"] = state ? state.encryptionKeyName : undefined;
            inputs["firstIpAddress"] = state ? state.firstIpAddress : undefined;
            inputs["ipAddresses"] = state ? state.ipAddresses : undefined;
            inputs["masterInstanceName"] = state ? state.masterInstanceName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["privateIpAddress"] = state ? state.privateIpAddress : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["publicIpAddress"] = state ? state.publicIpAddress : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["replicaConfiguration"] = state ? state.replicaConfiguration : undefined;
            inputs["restoreBackupContext"] = state ? state.restoreBackupContext : undefined;
            inputs["rootPassword"] = state ? state.rootPassword : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["serverCaCerts"] = state ? state.serverCaCerts : undefined;
            inputs["serviceAccountEmailAddress"] = state ? state.serviceAccountEmailAddress : undefined;
            inputs["settings"] = state ? state.settings : undefined;
        } else {
            const args = argsOrState as DatabaseInstanceArgs | undefined;
            inputs["clone"] = args ? args.clone : undefined;
            inputs["databaseVersion"] = args ? args.databaseVersion : undefined;
            inputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            inputs["encryptionKeyName"] = args ? args.encryptionKeyName : undefined;
            inputs["masterInstanceName"] = args ? args.masterInstanceName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["replicaConfiguration"] = args ? args.replicaConfiguration : undefined;
            inputs["restoreBackupContext"] = args ? args.restoreBackupContext : undefined;
            inputs["rootPassword"] = args ? args.rootPassword : undefined;
            inputs["settings"] = args ? args.settings : undefined;
            inputs["connectionName"] = undefined /*out*/;
            inputs["firstIpAddress"] = undefined /*out*/;
            inputs["ipAddresses"] = undefined /*out*/;
            inputs["privateIpAddress"] = undefined /*out*/;
            inputs["publicIpAddress"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["serverCaCerts"] = undefined /*out*/;
            inputs["serviceAccountEmailAddress"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DatabaseInstance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabaseInstance resources.
 */
export interface DatabaseInstanceState {
    /**
     * Configuration for creating a new instance as a clone of another instance.
     */
    readonly clone?: pulumi.Input<inputs.sql.DatabaseInstanceClone>;
    /**
     * The connection name of the instance to be used in
     * connection strings. For example, when connecting with [Cloud SQL Proxy](https://cloud.google.com/sql/docs/mysql/connect-admin-proxy).
     */
    readonly connectionName?: pulumi.Input<string>;
    /**
     * The MySQL, PostgreSQL or
     * SQL Server (beta) version to use. Supported values include `MYSQL_5_6`,
     * `MYSQL_5_7`, `MYSQL_8_0`, `POSTGRES_9_6`,`POSTGRES_10`, `POSTGRES_11`,
     * `POSTGRES_12`, `POSTGRES_13`, `SQLSERVER_2017_STANDARD`,
     * `SQLSERVER_2017_ENTERPRISE`, `SQLSERVER_2017_EXPRESS`, `SQLSERVER_2017_WEB`.
     * [Database Version Policies](https://cloud.google.com/sql/docs/db-versions)
     * includes an up-to-date reference of supported versions.
     */
    readonly databaseVersion?: pulumi.Input<string>;
    /**
     * Used to block Terraform from deleting a SQL Instance.
     */
    readonly deletionProtection?: pulumi.Input<boolean>;
    /**
     * The full path to the encryption key used for the CMEK disk encryption.  Setting
     * up disk encryption currently requires manual steps outside of this provider.
     * The provided key must be in the same region as the SQL instance.  In order
     * to use this feature, a special kind of service account must be created and
     * granted permission on this key.  This step can currently only be done
     * manually, please see [this step](https://cloud.google.com/sql/docs/mysql/configure-cmek#service-account).
     * That service account needs the `Cloud KMS > Cloud KMS CryptoKey Encrypter/Decrypter` role on your
     * key - please see [this step](https://cloud.google.com/sql/docs/mysql/configure-cmek#grantkey).
     */
    readonly encryptionKeyName?: pulumi.Input<string>;
    /**
     * The first IPv4 address of any type assigned.
     */
    readonly firstIpAddress?: pulumi.Input<string>;
    readonly ipAddresses?: pulumi.Input<pulumi.Input<inputs.sql.DatabaseInstanceIpAddress>[]>;
    /**
     * The name of the instance that will act as
     * the master in the replication setup. Note, this requires the master to have
     * `binaryLogEnabled` set, as well as existing backups.
     */
    readonly masterInstanceName?: pulumi.Input<string>;
    /**
     * A name for this whitelist entry.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The first private (`PRIVATE`) IPv4 address assigned.
     */
    readonly privateIpAddress?: pulumi.Input<string>;
    /**
     * The full project ID of the source instance.`
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The first public (`PRIMARY`) IPv4 address assigned.
     */
    readonly publicIpAddress?: pulumi.Input<string>;
    /**
     * The region the instance will sit in. Note, Cloud SQL is not
     * available in all regions - choose from one of the options listed [here](https://cloud.google.com/sql/docs/mysql/instance-locations).
     * A valid region must be provided to use this resource. If a region is not provided in the resource definition,
     * the provider region will be used instead, but this will be an apply-time error for instances if the provider
     * region is not supported with Cloud SQL. If you choose not to provide the `region` argument for this resource,
     * make sure you understand this.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The configuration for replication. The
     * configuration is detailed below.
     */
    readonly replicaConfiguration?: pulumi.Input<inputs.sql.DatabaseInstanceReplicaConfiguration>;
    readonly restoreBackupContext?: pulumi.Input<inputs.sql.DatabaseInstanceRestoreBackupContext>;
    /**
     * Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
     */
    readonly rootPassword?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly serverCaCerts?: pulumi.Input<pulumi.Input<inputs.sql.DatabaseInstanceServerCaCert>[]>;
    /**
     * The service account email address assigned to the
     * instance.
     */
    readonly serviceAccountEmailAddress?: pulumi.Input<string>;
    /**
     * The settings to use for the database. The
     * configuration is detailed below. Required if `clone` is not set.
     */
    readonly settings?: pulumi.Input<inputs.sql.DatabaseInstanceSettings>;
}

/**
 * The set of arguments for constructing a DatabaseInstance resource.
 */
export interface DatabaseInstanceArgs {
    /**
     * Configuration for creating a new instance as a clone of another instance.
     */
    readonly clone?: pulumi.Input<inputs.sql.DatabaseInstanceClone>;
    /**
     * The MySQL, PostgreSQL or
     * SQL Server (beta) version to use. Supported values include `MYSQL_5_6`,
     * `MYSQL_5_7`, `MYSQL_8_0`, `POSTGRES_9_6`,`POSTGRES_10`, `POSTGRES_11`,
     * `POSTGRES_12`, `POSTGRES_13`, `SQLSERVER_2017_STANDARD`,
     * `SQLSERVER_2017_ENTERPRISE`, `SQLSERVER_2017_EXPRESS`, `SQLSERVER_2017_WEB`.
     * [Database Version Policies](https://cloud.google.com/sql/docs/db-versions)
     * includes an up-to-date reference of supported versions.
     */
    readonly databaseVersion?: pulumi.Input<string>;
    /**
     * Used to block Terraform from deleting a SQL Instance.
     */
    readonly deletionProtection?: pulumi.Input<boolean>;
    /**
     * The full path to the encryption key used for the CMEK disk encryption.  Setting
     * up disk encryption currently requires manual steps outside of this provider.
     * The provided key must be in the same region as the SQL instance.  In order
     * to use this feature, a special kind of service account must be created and
     * granted permission on this key.  This step can currently only be done
     * manually, please see [this step](https://cloud.google.com/sql/docs/mysql/configure-cmek#service-account).
     * That service account needs the `Cloud KMS > Cloud KMS CryptoKey Encrypter/Decrypter` role on your
     * key - please see [this step](https://cloud.google.com/sql/docs/mysql/configure-cmek#grantkey).
     */
    readonly encryptionKeyName?: pulumi.Input<string>;
    /**
     * The name of the instance that will act as
     * the master in the replication setup. Note, this requires the master to have
     * `binaryLogEnabled` set, as well as existing backups.
     */
    readonly masterInstanceName?: pulumi.Input<string>;
    /**
     * A name for this whitelist entry.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The full project ID of the source instance.`
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region the instance will sit in. Note, Cloud SQL is not
     * available in all regions - choose from one of the options listed [here](https://cloud.google.com/sql/docs/mysql/instance-locations).
     * A valid region must be provided to use this resource. If a region is not provided in the resource definition,
     * the provider region will be used instead, but this will be an apply-time error for instances if the provider
     * region is not supported with Cloud SQL. If you choose not to provide the `region` argument for this resource,
     * make sure you understand this.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The configuration for replication. The
     * configuration is detailed below.
     */
    readonly replicaConfiguration?: pulumi.Input<inputs.sql.DatabaseInstanceReplicaConfiguration>;
    readonly restoreBackupContext?: pulumi.Input<inputs.sql.DatabaseInstanceRestoreBackupContext>;
    /**
     * Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
     */
    readonly rootPassword?: pulumi.Input<string>;
    /**
     * The settings to use for the database. The
     * configuration is detailed below. Required if `clone` is not set.
     */
    readonly settings?: pulumi.Input<inputs.sql.DatabaseInstanceSettings>;
}
