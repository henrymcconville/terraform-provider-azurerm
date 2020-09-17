package automation

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// NodeReportsClient is the automation Client
type NodeReportsClient struct {
	BaseClient
}

// NewNodeReportsClient creates an instance of the NodeReportsClient client.
func NewNodeReportsClient(subscriptionID string) NodeReportsClient {
	return NewNodeReportsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewNodeReportsClientWithBaseURI creates an instance of the NodeReportsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewNodeReportsClientWithBaseURI(baseURI string, subscriptionID string) NodeReportsClient {
	return NodeReportsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get retrieve the Dsc node report data by node id and report id.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// nodeID - the Dsc node id.
// reportID - the report id.
func (client NodeReportsClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, reportID string) (result DscNodeReport, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NodeReportsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.NodeReportsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, automationAccountName, nodeID, reportID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client NodeReportsClient) GetPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, reportID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"nodeId":                autorest.Encode("path", nodeID),
		"reportId":              autorest.Encode("path", reportID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}/reports/{reportId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client NodeReportsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client NodeReportsClient) GetResponder(resp *http.Response) (result DscNodeReport, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetContent retrieve the Dsc node reports by node id and report id.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// nodeID - the Dsc node id.
// reportID - the report id.
func (client NodeReportsClient) GetContent(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, reportID string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NodeReportsClient.GetContent")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.NodeReportsClient", "GetContent", err.Error())
	}

	req, err := client.GetContentPreparer(ctx, resourceGroupName, automationAccountName, nodeID, reportID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "GetContent", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetContentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "GetContent", resp, "Failure sending request")
		return
	}

	result, err = client.GetContentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "GetContent", resp, "Failure responding to request")
	}

	return
}

// GetContentPreparer prepares the GetContent request.
func (client NodeReportsClient) GetContentPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, reportID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"nodeId":                autorest.Encode("path", nodeID),
		"reportId":              autorest.Encode("path", reportID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}/reports/{reportId}/content", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetContentSender sends the GetContent request. The method will close the
// http.Response Body if it receives an error.
func (client NodeReportsClient) GetContentSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetContentResponder handles the response to the GetContent request. The method always
// closes the http.Response Body.
func (client NodeReportsClient) GetContentResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByNode retrieve the Dsc node report list by node id.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// nodeID - the parameters supplied to the list operation.
// filter - the filter to apply on the operation.
func (client NodeReportsClient) ListByNode(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, filter string) (result DscNodeReportListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NodeReportsClient.ListByNode")
		defer func() {
			sc := -1
			if result.dnrlr.Response.Response != nil {
				sc = result.dnrlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.NodeReportsClient", "ListByNode", err.Error())
	}

	result.fn = client.listByNodeNextResults
	req, err := client.ListByNodePreparer(ctx, resourceGroupName, automationAccountName, nodeID, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "ListByNode", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByNodeSender(req)
	if err != nil {
		result.dnrlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "ListByNode", resp, "Failure sending request")
		return
	}

	result.dnrlr, err = client.ListByNodeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "ListByNode", resp, "Failure responding to request")
	}

	return
}

// ListByNodePreparer prepares the ListByNode request.
func (client NodeReportsClient) ListByNodePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"nodeId":                autorest.Encode("path", nodeID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}/reports", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByNodeSender sends the ListByNode request. The method will close the
// http.Response Body if it receives an error.
func (client NodeReportsClient) ListByNodeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByNodeResponder handles the response to the ListByNode request. The method always
// closes the http.Response Body.
func (client NodeReportsClient) ListByNodeResponder(resp *http.Response) (result DscNodeReportListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByNodeNextResults retrieves the next set of results, if any.
func (client NodeReportsClient) listByNodeNextResults(ctx context.Context, lastResults DscNodeReportListResult) (result DscNodeReportListResult, err error) {
	req, err := lastResults.dscNodeReportListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.NodeReportsClient", "listByNodeNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByNodeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.NodeReportsClient", "listByNodeNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByNodeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "listByNodeNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByNodeComplete enumerates all values, automatically crossing page boundaries as required.
func (client NodeReportsClient) ListByNodeComplete(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, filter string) (result DscNodeReportListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NodeReportsClient.ListByNode")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByNode(ctx, resourceGroupName, automationAccountName, nodeID, filter)
	return
}