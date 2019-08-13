// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface ServiceApi {
    methods: outputApi.endpoints.ServiceApiMethod[];
    name: string;
    syntax: string;
    version: string;
}

export interface ServiceApiMethod {
    name: string;
    requestType: string;
    responseType: string;
    syntax: string;
}

export interface ServiceEndpoint {
    address: string;
    name: string;
}
