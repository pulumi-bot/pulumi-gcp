// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring
{
    /// <summary>
    /// A Service-Level Objective (SLO) describes the level of desired good
    /// service. It consists of a service-level indicator (SLI), a performance
    /// goal, and a period over which the objective is to be evaluated against
    /// that goal. The SLO can use SLIs defined in a number of different manners.
    /// Typical SLOs might include "99% of requests in each rolling week have
    /// latency below 200 milliseconds" or "99.5% of requests in each calendar
    /// month return successfully."
    /// 
    /// To get more information about Slo, see:
    /// 
    /// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services.serviceLevelObjectives)
    /// * How-to Guides
    ///     * [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
    ///     * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// Slo can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:monitoring/slo:Slo default {{name}}
    /// ```
    /// </summary>
    public partial class Slo : Pulumi.CustomResource
    {
        /// <summary>
        /// Basic Service-Level Indicator (SLI) on a well-known service type.
        /// Performance will be computed on the basis of pre-defined metrics.
        /// SLIs are used to measure and calculate the quality of the Service's
        /// performance with respect to a single aspect of service quality.
        /// Exactly one of the following must be set:
        /// `basic_sli`, `request_based_sli`, `windows_based_sli`
        /// Structure is documented below.
        /// </summary>
        [Output("basicSli")]
        public Output<Outputs.SloBasicSli?> BasicSli { get; private set; } = null!;

        /// <summary>
        /// A calendar period, semantically "since the start of the current
        /// &lt;calendarPeriod&gt;".
        /// Possible values are `DAY`, `WEEK`, `FORTNIGHT`, and `MONTH`.
        /// </summary>
        [Output("calendarPeriod")]
        public Output<string?> CalendarPeriod { get; private set; } = null!;

        /// <summary>
        /// Name used for UI elements listing this SLO.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The fraction of service that must be good in order for this objective
        /// to be met. 0 &lt; goal &lt;= 0.999
        /// </summary>
        [Output("goal")]
        public Output<double> Goal { get; private set; } = null!;

        /// <summary>
        /// The full resource name for this service. The syntax is:
        /// projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// A request-based SLI defines a SLI for which atomic units of
        /// service are counted directly.
        /// A SLI describes a good service.
        /// It is used to measure and calculate the quality of the Service's
        /// performance with respect to a single aspect of service quality.
        /// Exactly one of the following must be set:
        /// `basic_sli`, `request_based_sli`, `windows_based_sli`
        /// Structure is documented below.
        /// </summary>
        [Output("requestBasedSli")]
        public Output<Outputs.SloRequestBasedSli?> RequestBasedSli { get; private set; } = null!;

        /// <summary>
        /// A rolling time period, semantically "in the past X days".
        /// Must be between 1 to 30 days, inclusive.
        /// </summary>
        [Output("rollingPeriodDays")]
        public Output<int?> RollingPeriodDays { get; private set; } = null!;

        /// <summary>
        /// ID of the service to which this SLO belongs.
        /// </summary>
        [Output("service")]
        public Output<string> Service { get; private set; } = null!;

        /// <summary>
        /// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
        /// </summary>
        [Output("sloId")]
        public Output<string> SloId { get; private set; } = null!;

        /// <summary>
        /// A windows-based SLI defines the criteria for time windows.
        /// good_service is defined based off the count of these time windows
        /// for which the provided service was of good quality.
        /// A SLI describes a good service. It is used to measure and calculate
        /// the quality of the Service's performance with respect to a single
        /// aspect of service quality.
        /// Exactly one of the following must be set:
        /// `basic_sli`, `request_based_sli`, `windows_based_sli`
        /// Structure is documented below.
        /// </summary>
        [Output("windowsBasedSli")]
        public Output<Outputs.SloWindowsBasedSli?> WindowsBasedSli { get; private set; } = null!;


        /// <summary>
        /// Create a Slo resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Slo(string name, SloArgs args, CustomResourceOptions? options = null)
            : base("gcp:monitoring/slo:Slo", name, args ?? new SloArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Slo(string name, Input<string> id, SloState? state = null, CustomResourceOptions? options = null)
            : base("gcp:monitoring/slo:Slo", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Slo resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Slo Get(string name, Input<string> id, SloState? state = null, CustomResourceOptions? options = null)
        {
            return new Slo(name, id, state, options);
        }
    }

    public sealed class SloArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Basic Service-Level Indicator (SLI) on a well-known service type.
        /// Performance will be computed on the basis of pre-defined metrics.
        /// SLIs are used to measure and calculate the quality of the Service's
        /// performance with respect to a single aspect of service quality.
        /// Exactly one of the following must be set:
        /// `basic_sli`, `request_based_sli`, `windows_based_sli`
        /// Structure is documented below.
        /// </summary>
        [Input("basicSli")]
        public Input<Inputs.SloBasicSliArgs>? BasicSli { get; set; }

        /// <summary>
        /// A calendar period, semantically "since the start of the current
        /// &lt;calendarPeriod&gt;".
        /// Possible values are `DAY`, `WEEK`, `FORTNIGHT`, and `MONTH`.
        /// </summary>
        [Input("calendarPeriod")]
        public Input<string>? CalendarPeriod { get; set; }

        /// <summary>
        /// Name used for UI elements listing this SLO.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The fraction of service that must be good in order for this objective
        /// to be met. 0 &lt; goal &lt;= 0.999
        /// </summary>
        [Input("goal", required: true)]
        public Input<double> Goal { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A request-based SLI defines a SLI for which atomic units of
        /// service are counted directly.
        /// A SLI describes a good service.
        /// It is used to measure and calculate the quality of the Service's
        /// performance with respect to a single aspect of service quality.
        /// Exactly one of the following must be set:
        /// `basic_sli`, `request_based_sli`, `windows_based_sli`
        /// Structure is documented below.
        /// </summary>
        [Input("requestBasedSli")]
        public Input<Inputs.SloRequestBasedSliArgs>? RequestBasedSli { get; set; }

        /// <summary>
        /// A rolling time period, semantically "in the past X days".
        /// Must be between 1 to 30 days, inclusive.
        /// </summary>
        [Input("rollingPeriodDays")]
        public Input<int>? RollingPeriodDays { get; set; }

        /// <summary>
        /// ID of the service to which this SLO belongs.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        /// <summary>
        /// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
        /// </summary>
        [Input("sloId")]
        public Input<string>? SloId { get; set; }

        /// <summary>
        /// A windows-based SLI defines the criteria for time windows.
        /// good_service is defined based off the count of these time windows
        /// for which the provided service was of good quality.
        /// A SLI describes a good service. It is used to measure and calculate
        /// the quality of the Service's performance with respect to a single
        /// aspect of service quality.
        /// Exactly one of the following must be set:
        /// `basic_sli`, `request_based_sli`, `windows_based_sli`
        /// Structure is documented below.
        /// </summary>
        [Input("windowsBasedSli")]
        public Input<Inputs.SloWindowsBasedSliArgs>? WindowsBasedSli { get; set; }

        public SloArgs()
        {
        }
    }

    public sealed class SloState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Basic Service-Level Indicator (SLI) on a well-known service type.
        /// Performance will be computed on the basis of pre-defined metrics.
        /// SLIs are used to measure and calculate the quality of the Service's
        /// performance with respect to a single aspect of service quality.
        /// Exactly one of the following must be set:
        /// `basic_sli`, `request_based_sli`, `windows_based_sli`
        /// Structure is documented below.
        /// </summary>
        [Input("basicSli")]
        public Input<Inputs.SloBasicSliGetArgs>? BasicSli { get; set; }

        /// <summary>
        /// A calendar period, semantically "since the start of the current
        /// &lt;calendarPeriod&gt;".
        /// Possible values are `DAY`, `WEEK`, `FORTNIGHT`, and `MONTH`.
        /// </summary>
        [Input("calendarPeriod")]
        public Input<string>? CalendarPeriod { get; set; }

        /// <summary>
        /// Name used for UI elements listing this SLO.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The fraction of service that must be good in order for this objective
        /// to be met. 0 &lt; goal &lt;= 0.999
        /// </summary>
        [Input("goal")]
        public Input<double>? Goal { get; set; }

        /// <summary>
        /// The full resource name for this service. The syntax is:
        /// projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A request-based SLI defines a SLI for which atomic units of
        /// service are counted directly.
        /// A SLI describes a good service.
        /// It is used to measure and calculate the quality of the Service's
        /// performance with respect to a single aspect of service quality.
        /// Exactly one of the following must be set:
        /// `basic_sli`, `request_based_sli`, `windows_based_sli`
        /// Structure is documented below.
        /// </summary>
        [Input("requestBasedSli")]
        public Input<Inputs.SloRequestBasedSliGetArgs>? RequestBasedSli { get; set; }

        /// <summary>
        /// A rolling time period, semantically "in the past X days".
        /// Must be between 1 to 30 days, inclusive.
        /// </summary>
        [Input("rollingPeriodDays")]
        public Input<int>? RollingPeriodDays { get; set; }

        /// <summary>
        /// ID of the service to which this SLO belongs.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
        /// </summary>
        [Input("sloId")]
        public Input<string>? SloId { get; set; }

        /// <summary>
        /// A windows-based SLI defines the criteria for time windows.
        /// good_service is defined based off the count of these time windows
        /// for which the provided service was of good quality.
        /// A SLI describes a good service. It is used to measure and calculate
        /// the quality of the Service's performance with respect to a single
        /// aspect of service quality.
        /// Exactly one of the following must be set:
        /// `basic_sli`, `request_based_sli`, `windows_based_sli`
        /// Structure is documented below.
        /// </summary>
        [Input("windowsBasedSli")]
        public Input<Inputs.SloWindowsBasedSliGetArgs>? WindowsBasedSli { get; set; }

        public SloState()
        {
        }
    }
}
