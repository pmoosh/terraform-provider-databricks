package eventgrid

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DomainTopicsClient is the azure EventGrid Management Client
type DomainTopicsClient struct {
	BaseClient
}

// NewDomainTopicsClient creates an instance of the DomainTopicsClient client.
func NewDomainTopicsClient(subscriptionID string) DomainTopicsClient {
	return NewDomainTopicsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDomainTopicsClientWithBaseURI creates an instance of the DomainTopicsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDomainTopicsClientWithBaseURI(baseURI string, subscriptionID string) DomainTopicsClient {
	return DomainTopicsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get properties of a domain topic
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// domainName - name of the domain
// topicName - name of the topic
func (client DomainTopicsClient) Get(ctx context.Context, resourceGroupName string, domainName string, topicName string) (result DomainTopic, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainTopicsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, domainName, topicName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DomainTopicsClient) GetPreparer(ctx context.Context, resourceGroupName string, domainName string, topicName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainName":        autorest.Encode("path", domainName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"topicName":         autorest.Encode("path", topicName),
	}

	const APIVersion = "2018-09-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{topicName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DomainTopicsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DomainTopicsClient) GetResponder(resp *http.Response) (result DomainTopic, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDomain list all the topics in a domain.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// domainName - domain name.
func (client DomainTopicsClient) ListByDomain(ctx context.Context, resourceGroupName string, domainName string) (result DomainTopicsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainTopicsClient.ListByDomain")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByDomainPreparer(ctx, resourceGroupName, domainName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "ListByDomain", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDomainSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "ListByDomain", resp, "Failure sending request")
		return
	}

	result, err = client.ListByDomainResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "ListByDomain", resp, "Failure responding to request")
	}

	return
}

// ListByDomainPreparer prepares the ListByDomain request.
func (client DomainTopicsClient) ListByDomainPreparer(ctx context.Context, resourceGroupName string, domainName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainName":        autorest.Encode("path", domainName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDomainSender sends the ListByDomain request. The method will close the
// http.Response Body if it receives an error.
func (client DomainTopicsClient) ListByDomainSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDomainResponder handles the response to the ListByDomain request. The method always
// closes the http.Response Body.
func (client DomainTopicsClient) ListByDomainResponder(resp *http.Response) (result DomainTopicsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}