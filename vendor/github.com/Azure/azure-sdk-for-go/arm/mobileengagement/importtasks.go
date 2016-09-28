package mobileengagement

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ImportTasksClient is the microsoft Azure Mobile Engagement REST APIs.
type ImportTasksClient struct {
	ManagementClient
}

// NewImportTasksClient creates an instance of the ImportTasksClient client.
func NewImportTasksClient(subscriptionID string, resourceGroupName string, appCollection string, appName string) ImportTasksClient {
	return NewImportTasksClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, appCollection, appName)
}

// NewImportTasksClientWithBaseURI creates an instance of the
// ImportTasksClient client.
func NewImportTasksClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, appCollection string, appName string) ImportTasksClient {
	return ImportTasksClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, appCollection, appName)}
}

// Create creates a job to import the specified data to a storageUrl.
//
func (client ImportTasksClient) Create(parameters ImportTask) (result ImportTaskResult, err error) {
	req, err := client.CreatePreparer(parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "Create", nil, "Failure preparing request")
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "Create", resp, "Failure sending request")
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ImportTasksClient) CreatePreparer(parameters ImportTask) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appCollection":     autorest.Encode("path", client.AppCollection),
		"appName":           autorest.Encode("path", client.AppName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileEngagement/appcollections/{appCollection}/apps/{appName}/devices/importTasks", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ImportTasksClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ImportTasksClient) CreateResponder(resp *http.Response) (result ImportTaskResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get the Get import job operation retrieves information about a previously
// created import job.
//
// id is import job identifier.
func (client ImportTasksClient) Get(id string) (result ImportTaskResult, err error) {
	req, err := client.GetPreparer(id)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ImportTasksClient) GetPreparer(id string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appCollection":     autorest.Encode("path", client.AppCollection),
		"appName":           autorest.Encode("path", client.AppName),
		"id":                autorest.Encode("path", id),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileEngagement/appcollections/{appCollection}/apps/{appName}/devices/importTasks/{id}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ImportTasksClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ImportTasksClient) GetResponder(resp *http.Response) (result ImportTaskResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get the list of import jobs.
//
// skip is control paging of import jobs, start results at the given offset,
// defaults to 0 (1st page of data). top is control paging of import jobs,
// number of import jobs to return with each call. By default, it returns all
// import jobs with a default paging of 20.
// The response contains a `nextLink` property describing the path to get the
// next page if there are more results.
// The maximum paging limit for $top is 40. orderby is sort results by an
// expression which looks like `$orderby=jobId asc` (default when not
// specified).
// The syntax is orderby={property} {direction} or just orderby={property}.
// Properties that can be specified for sorting: jobId, errorDetails,
// dateCreated, jobStatus, and dateCreated.
// The available directions are asc (for ascending order) and desc (for
// descending order).
// When not specified the asc direction is used.
// Only one orderby property can be specified.
func (client ImportTasksClient) List(skip *int32, top *int32, orderby string) (result ImportTaskListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{skip,
			[]validation.Constraint{{"skip", validation.Null, false,
				[]validation.Constraint{{"skip", validation.InclusiveMinimum, 0, nil}}}}},
		{top,
			[]validation.Constraint{{"top", validation.Null, false,
				[]validation.Constraint{{"top", validation.InclusiveMaximum, 40, nil},
					{"top", validation.InclusiveMinimum, 1, nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "mobileengagement.ImportTasksClient", "List")
	}

	req, err := client.ListPreparer(skip, top, orderby)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ImportTasksClient) ListPreparer(skip *int32, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appCollection":     autorest.Encode("path", client.AppCollection),
		"appName":           autorest.Encode("path", client.AppName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileEngagement/appcollections/{appCollection}/apps/{appName}/devices/importTasks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ImportTasksClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ImportTasksClient) ListResponder(resp *http.Response) (result ImportTaskListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client ImportTasksClient) ListNextResults(lastResults ImportTaskListResult) (result ImportTaskListResult, err error) {
	req, err := lastResults.ImportTaskListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "List", resp, "Failure responding to next results request")
	}

	return
}
