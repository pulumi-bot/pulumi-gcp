// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Billing.Inputs
{

    public sealed class BudgetBudgetFilterGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies how credits should be treated when determining spend
        /// for threshold calculations.
        /// </summary>
        [Input("creditTypesTreatment")]
        public Input<string>? CreditTypesTreatment { get; set; }

        [Input("projects")]
        private InputList<string>? _projects;

        /// <summary>
        /// A set of projects of the form projects/{project_id},
        /// specifying that usage from only this set of projects should be
        /// included in the budget. If omitted, the report will include
        /// all usage for the billing account, regardless of which project
        /// the usage occurred on. Only zero or one project can be
        /// specified currently.
        /// </summary>
        public InputList<string> Projects
        {
            get => _projects ?? (_projects = new InputList<string>());
            set => _projects = value;
        }

        [Input("services")]
        private InputList<string>? _services;

        /// <summary>
        /// A set of services of the form services/{service_id},
        /// specifying that usage from only this set of services should be
        /// included in the budget. If omitted, the report will include
        /// usage for all the services. The service names are available
        /// through the Catalog API:
        /// https://cloud.google.com/billing/v1/how-tos/catalog-api.
        /// </summary>
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        public BudgetBudgetFilterGetArgs()
        {
        }
    }
}
