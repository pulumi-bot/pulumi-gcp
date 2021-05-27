// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Sql
{
    public static class GetBackupRun
    {
        /// <summary>
        /// Use this data source to get information about a Cloud SQL instance backup run.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var backup = Output.Create(Gcp.Sql.GetBackupRun.InvokeAsync(new Gcp.Sql.GetBackupRunArgs
        ///         {
        ///             Instance = google_sql_database_instance.Master.Name,
        ///             MostRecent = true,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBackupRunResult> InvokeAsync(GetBackupRunArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBackupRunResult>("gcp:sql/getBackupRun:getBackupRun", args ?? new GetBackupRunArgs(), options.WithVersion());

        public static Output<GetBackupRunResult> Apply(GetBackupRunApplyArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.BackupId.Box(),
                args.Instance.Box(),
                args.MostRecent.Box()
            ).Apply(a => {
                    var args = new GetBackupRunArgs();
                    a[0].Set(args, nameof(args.BackupId));
                    a[1].Set(args, nameof(args.Instance));
                    a[2].Set(args, nameof(args.MostRecent));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetBackupRunArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier for this backup run. Unique only for a specific Cloud SQL instance.
        /// If left empty and multiple backups exist for the instance, `most_recent` must be set to `true`.
        /// </summary>
        [Input("backupId")]
        public int? BackupId { get; set; }

        /// <summary>
        /// The name of the instance the backup is taken from.
        /// </summary>
        [Input("instance", required: true)]
        public string Instance { get; set; } = null!;

        /// <summary>
        /// Toggles use of the most recent backup run if multiple backups exist for a 
        /// Cloud SQL instance.
        /// </summary>
        [Input("mostRecent")]
        public bool? MostRecent { get; set; }

        public GetBackupRunArgs()
        {
        }
    }

    public sealed class GetBackupRunApplyArgs
    {
        /// <summary>
        /// The identifier for this backup run. Unique only for a specific Cloud SQL instance.
        /// If left empty and multiple backups exist for the instance, `most_recent` must be set to `true`.
        /// </summary>
        [Input("backupId")]
        public Input<int>? BackupId { get; set; }

        /// <summary>
        /// The name of the instance the backup is taken from.
        /// </summary>
        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        /// <summary>
        /// Toggles use of the most recent backup run if multiple backups exist for a 
        /// Cloud SQL instance.
        /// </summary>
        [Input("mostRecent")]
        public Input<bool>? MostRecent { get; set; }

        public GetBackupRunApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBackupRunResult
    {
        public readonly int BackupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Instance;
        /// <summary>
        /// Location of the backups.
        /// </summary>
        public readonly string Location;
        public readonly bool? MostRecent;
        /// <summary>
        /// The time the backup operation actually started in UTC timezone in RFC 3339 format, for 
        /// example 2012-11-15T16:19:00.094Z.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// The status of this run. Refer to [API reference](https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1beta4/backupRuns#SqlBackupRunStatus) for possible status values.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetBackupRunResult(
            int backupId,

            string id,

            string instance,

            string location,

            bool? mostRecent,

            string startTime,

            string status)
        {
            BackupId = backupId;
            Id = id;
            Instance = instance;
            Location = location;
            MostRecent = mostRecent;
            StartTime = startTime;
            Status = status;
        }
    }
}
