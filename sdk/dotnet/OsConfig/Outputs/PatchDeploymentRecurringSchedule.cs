// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.OsConfig.Outputs
{

    [OutputType]
    public sealed class PatchDeploymentRecurringSchedule
    {
        /// <summary>
        /// The end time at which a recurring patch deployment schedule is no longer active.
        /// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        public readonly string? EndTime;
        /// <summary>
        /// -
        /// The time the last patch job ran successfully.
        /// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        public readonly string? LastExecuteTime;
        /// <summary>
        /// Schedule with monthly executions.  Structure is documented below.
        /// </summary>
        public readonly Outputs.PatchDeploymentRecurringScheduleMonthly? Monthly;
        /// <summary>
        /// -
        /// The time the next patch job is scheduled to run.
        /// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        public readonly string? NextExecuteTime;
        /// <summary>
        /// The time that the recurring schedule becomes effective. Defaults to createTime of the patch deployment.
        /// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        public readonly string? StartTime;
        /// <summary>
        /// Time of the day to run a recurring deployment.  Structure is documented below.
        /// </summary>
        public readonly Outputs.PatchDeploymentRecurringScheduleTimeOfDay TimeOfDay;
        /// <summary>
        /// Defines the time zone that timeOfDay is relative to. The rules for daylight saving time are
        /// determined by the chosen time zone.  Structure is documented below.
        /// </summary>
        public readonly Outputs.PatchDeploymentRecurringScheduleTimeZone TimeZone;
        /// <summary>
        /// Schedule with weekly executions.  Structure is documented below.
        /// </summary>
        public readonly Outputs.PatchDeploymentRecurringScheduleWeekly? Weekly;

        [OutputConstructor]
        private PatchDeploymentRecurringSchedule(
            string? endTime,

            string? lastExecuteTime,

            Outputs.PatchDeploymentRecurringScheduleMonthly? monthly,

            string? nextExecuteTime,

            string? startTime,

            Outputs.PatchDeploymentRecurringScheduleTimeOfDay timeOfDay,

            Outputs.PatchDeploymentRecurringScheduleTimeZone timeZone,

            Outputs.PatchDeploymentRecurringScheduleWeekly? weekly)
        {
            EndTime = endTime;
            LastExecuteTime = lastExecuteTime;
            Monthly = monthly;
            NextExecuteTime = nextExecuteTime;
            StartTime = startTime;
            TimeOfDay = timeOfDay;
            TimeZone = timeZone;
            Weekly = weekly;
        }
    }
}