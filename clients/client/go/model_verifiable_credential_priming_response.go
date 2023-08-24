/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.50
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// VerifiableCredentialPrimingResponse struct for VerifiableCredentialPrimingResponse
type VerifiableCredentialPrimingResponse struct {
	CNonce *string `json:"c_nonce,omitempty"`
	CNonceExpiresIn *int64 `json:"c_nonce_expires_in,omitempty"`
	Error *string `json:"error,omitempty"`
	ErrorDebug *string `json:"error_debug,omitempty"`
	ErrorDescription *string `json:"error_description,omitempty"`
	ErrorHint *string `json:"error_hint,omitempty"`
	Format *string `json:"format,omitempty"`
	StatusCode *int64 `json:"status_code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VerifiableCredentialPrimingResponse VerifiableCredentialPrimingResponse

// NewVerifiableCredentialPrimingResponse instantiates a new VerifiableCredentialPrimingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifiableCredentialPrimingResponse() *VerifiableCredentialPrimingResponse {
	this := VerifiableCredentialPrimingResponse{}
	return &this
}

// NewVerifiableCredentialPrimingResponseWithDefaults instantiates a new VerifiableCredentialPrimingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifiableCredentialPrimingResponseWithDefaults() *VerifiableCredentialPrimingResponse {
	this := VerifiableCredentialPrimingResponse{}
	return &this
}

// GetCNonce returns the CNonce field value if set, zero value otherwise.
func (o *VerifiableCredentialPrimingResponse) GetCNonce() string {
	if o == nil || o.CNonce == nil {
		var ret string
		return ret
	}
	return *o.CNonce
}

// GetCNonceOk returns a tuple with the CNonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiableCredentialPrimingResponse) GetCNonceOk() (*string, bool) {
	if o == nil || o.CNonce == nil {
		return nil, false
	}
	return o.CNonce, true
}

// HasCNonce returns a boolean if a field has been set.
func (o *VerifiableCredentialPrimingResponse) HasCNonce() bool {
	if o != nil && o.CNonce != nil {
		return true
	}

	return false
}

// SetCNonce gets a reference to the given string and assigns it to the CNonce field.
func (o *VerifiableCredentialPrimingResponse) SetCNonce(v string) {
	o.CNonce = &v
}

// GetCNonceExpiresIn returns the CNonceExpiresIn field value if set, zero value otherwise.
func (o *VerifiableCredentialPrimingResponse) GetCNonceExpiresIn() int64 {
	if o == nil || o.CNonceExpiresIn == nil {
		var ret int64
		return ret
	}
	return *o.CNonceExpiresIn
}

// GetCNonceExpiresInOk returns a tuple with the CNonceExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiableCredentialPrimingResponse) GetCNonceExpiresInOk() (*int64, bool) {
	if o == nil || o.CNonceExpiresIn == nil {
		return nil, false
	}
	return o.CNonceExpiresIn, true
}

// HasCNonceExpiresIn returns a boolean if a field has been set.
func (o *VerifiableCredentialPrimingResponse) HasCNonceExpiresIn() bool {
	if o != nil && o.CNonceExpiresIn != nil {
		return true
	}

	return false
}

// SetCNonceExpiresIn gets a reference to the given int64 and assigns it to the CNonceExpiresIn field.
func (o *VerifiableCredentialPrimingResponse) SetCNonceExpiresIn(v int64) {
	o.CNonceExpiresIn = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *VerifiableCredentialPrimingResponse) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiableCredentialPrimingResponse) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *VerifiableCredentialPrimingResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *VerifiableCredentialPrimingResponse) SetError(v string) {
	o.Error = &v
}

// GetErrorDebug returns the ErrorDebug field value if set, zero value otherwise.
func (o *VerifiableCredentialPrimingResponse) GetErrorDebug() string {
	if o == nil || o.ErrorDebug == nil {
		var ret string
		return ret
	}
	return *o.ErrorDebug
}

// GetErrorDebugOk returns a tuple with the ErrorDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiableCredentialPrimingResponse) GetErrorDebugOk() (*string, bool) {
	if o == nil || o.ErrorDebug == nil {
		return nil, false
	}
	return o.ErrorDebug, true
}

// HasErrorDebug returns a boolean if a field has been set.
func (o *VerifiableCredentialPrimingResponse) HasErrorDebug() bool {
	if o != nil && o.ErrorDebug != nil {
		return true
	}

	return false
}

// SetErrorDebug gets a reference to the given string and assigns it to the ErrorDebug field.
func (o *VerifiableCredentialPrimingResponse) SetErrorDebug(v string) {
	o.ErrorDebug = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *VerifiableCredentialPrimingResponse) GetErrorDescription() string {
	if o == nil || o.ErrorDescription == nil {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiableCredentialPrimingResponse) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || o.ErrorDescription == nil {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *VerifiableCredentialPrimingResponse) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription != nil {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *VerifiableCredentialPrimingResponse) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetErrorHint returns the ErrorHint field value if set, zero value otherwise.
func (o *VerifiableCredentialPrimingResponse) GetErrorHint() string {
	if o == nil || o.ErrorHint == nil {
		var ret string
		return ret
	}
	return *o.ErrorHint
}

// GetErrorHintOk returns a tuple with the ErrorHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiableCredentialPrimingResponse) GetErrorHintOk() (*string, bool) {
	if o == nil || o.ErrorHint == nil {
		return nil, false
	}
	return o.ErrorHint, true
}

// HasErrorHint returns a boolean if a field has been set.
func (o *VerifiableCredentialPrimingResponse) HasErrorHint() bool {
	if o != nil && o.ErrorHint != nil {
		return true
	}

	return false
}

// SetErrorHint gets a reference to the given string and assigns it to the ErrorHint field.
func (o *VerifiableCredentialPrimingResponse) SetErrorHint(v string) {
	o.ErrorHint = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *VerifiableCredentialPrimingResponse) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiableCredentialPrimingResponse) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *VerifiableCredentialPrimingResponse) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *VerifiableCredentialPrimingResponse) SetFormat(v string) {
	o.Format = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *VerifiableCredentialPrimingResponse) GetStatusCode() int64 {
	if o == nil || o.StatusCode == nil {
		var ret int64
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiableCredentialPrimingResponse) GetStatusCodeOk() (*int64, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *VerifiableCredentialPrimingResponse) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int64 and assigns it to the StatusCode field.
func (o *VerifiableCredentialPrimingResponse) SetStatusCode(v int64) {
	o.StatusCode = &v
}

func (o VerifiableCredentialPrimingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CNonce != nil {
		toSerialize["c_nonce"] = o.CNonce
	}
	if o.CNonceExpiresIn != nil {
		toSerialize["c_nonce_expires_in"] = o.CNonceExpiresIn
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.ErrorDebug != nil {
		toSerialize["error_debug"] = o.ErrorDebug
	}
	if o.ErrorDescription != nil {
		toSerialize["error_description"] = o.ErrorDescription
	}
	if o.ErrorHint != nil {
		toSerialize["error_hint"] = o.ErrorHint
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.StatusCode != nil {
		toSerialize["status_code"] = o.StatusCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VerifiableCredentialPrimingResponse) UnmarshalJSON(bytes []byte) (err error) {
	varVerifiableCredentialPrimingResponse := _VerifiableCredentialPrimingResponse{}

	if err = json.Unmarshal(bytes, &varVerifiableCredentialPrimingResponse); err == nil {
		*o = VerifiableCredentialPrimingResponse(varVerifiableCredentialPrimingResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "c_nonce")
		delete(additionalProperties, "c_nonce_expires_in")
		delete(additionalProperties, "error")
		delete(additionalProperties, "error_debug")
		delete(additionalProperties, "error_description")
		delete(additionalProperties, "error_hint")
		delete(additionalProperties, "format")
		delete(additionalProperties, "status_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVerifiableCredentialPrimingResponse struct {
	value *VerifiableCredentialPrimingResponse
	isSet bool
}

func (v NullableVerifiableCredentialPrimingResponse) Get() *VerifiableCredentialPrimingResponse {
	return v.value
}

func (v *NullableVerifiableCredentialPrimingResponse) Set(val *VerifiableCredentialPrimingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifiableCredentialPrimingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifiableCredentialPrimingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifiableCredentialPrimingResponse(val *VerifiableCredentialPrimingResponse) *NullableVerifiableCredentialPrimingResponse {
	return &NullableVerifiableCredentialPrimingResponse{value: val, isSet: true}
}

func (v NullableVerifiableCredentialPrimingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifiableCredentialPrimingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


