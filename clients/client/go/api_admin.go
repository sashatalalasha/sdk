/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.60
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type AdminApi interface {

	/*
	UpdateOAuth2ClientLifespans Method for UpdateOAuth2ClientLifespans

	UpdateLifespans an existing OAuth 2.0 client's token lifespan configuration. This
client configuration takes precedence over the instance-wide token lifespan
configuration.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The id of the OAuth 2.0 Client.
	@return AdminApiUpdateOAuth2ClientLifespansRequest
	*/
	UpdateOAuth2ClientLifespans(ctx context.Context, id string) AdminApiUpdateOAuth2ClientLifespansRequest

	// UpdateOAuth2ClientLifespansExecute executes the request
	//  @return OAuth2Client
	UpdateOAuth2ClientLifespansExecute(r AdminApiUpdateOAuth2ClientLifespansRequest) (*OAuth2Client, *http.Response, error)
}

// AdminApiService AdminApi service
type AdminApiService service

type AdminApiUpdateOAuth2ClientLifespansRequest struct {
	ctx context.Context
	ApiService AdminApi
	id string
	updateOAuth2ClientLifespans *UpdateOAuth2ClientLifespans
}

func (r AdminApiUpdateOAuth2ClientLifespansRequest) UpdateOAuth2ClientLifespans(updateOAuth2ClientLifespans UpdateOAuth2ClientLifespans) AdminApiUpdateOAuth2ClientLifespansRequest {
	r.updateOAuth2ClientLifespans = &updateOAuth2ClientLifespans
	return r
}

func (r AdminApiUpdateOAuth2ClientLifespansRequest) Execute() (*OAuth2Client, *http.Response, error) {
	return r.ApiService.UpdateOAuth2ClientLifespansExecute(r)
}

/*
UpdateOAuth2ClientLifespans Method for UpdateOAuth2ClientLifespans

UpdateLifespans an existing OAuth 2.0 client's token lifespan configuration. This
client configuration takes precedence over the instance-wide token lifespan
configuration.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The id of the OAuth 2.0 Client.
 @return AdminApiUpdateOAuth2ClientLifespansRequest
*/
func (a *AdminApiService) UpdateOAuth2ClientLifespans(ctx context.Context, id string) AdminApiUpdateOAuth2ClientLifespansRequest {
	return AdminApiUpdateOAuth2ClientLifespansRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return OAuth2Client
func (a *AdminApiService) UpdateOAuth2ClientLifespansExecute(r AdminApiUpdateOAuth2ClientLifespansRequest) (*OAuth2Client, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OAuth2Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdminApiService.UpdateOAuth2ClientLifespans")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/clients/{id}/lifespans"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateOAuth2ClientLifespans
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
