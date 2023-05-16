/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.29
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateRegistrationFlowWithWebAuthnMethod Update Registration Flow with WebAuthn Method
type UpdateRegistrationFlowWithWebAuthnMethod struct {
	// CSRFToken is the anti-CSRF token
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method  Should be set to \"webauthn\" when trying to add, update, or remove a webAuthn pairing.
	Method string `json:"method"`
	// The identity's traits
	Traits map[string]interface{} `json:"traits"`
	// Transient data to pass along to any webhooks
	TransientPayload map[string]interface{} `json:"transient_payload,omitempty"`
	// Register a WebAuthn Security Key  It is expected that the JSON returned by the WebAuthn registration process is included here.
	WebauthnRegister *string `json:"webauthn_register,omitempty"`
	// Name of the WebAuthn Security Key to be Added  A human-readable name for the security key which will be added.
	WebauthnRegisterDisplayname *string `json:"webauthn_register_displayname,omitempty"`
}

// NewUpdateRegistrationFlowWithWebAuthnMethod instantiates a new UpdateRegistrationFlowWithWebAuthnMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRegistrationFlowWithWebAuthnMethod(method string, traits map[string]interface{}) *UpdateRegistrationFlowWithWebAuthnMethod {
	this := UpdateRegistrationFlowWithWebAuthnMethod{}
	this.Method = method
	this.Traits = traits
	return &this
}

// NewUpdateRegistrationFlowWithWebAuthnMethodWithDefaults instantiates a new UpdateRegistrationFlowWithWebAuthnMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRegistrationFlowWithWebAuthnMethodWithDefaults() *UpdateRegistrationFlowWithWebAuthnMethod {
	this := UpdateRegistrationFlowWithWebAuthnMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *UpdateRegistrationFlowWithWebAuthnMethod) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithWebAuthnMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *UpdateRegistrationFlowWithWebAuthnMethod) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *UpdateRegistrationFlowWithWebAuthnMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *UpdateRegistrationFlowWithWebAuthnMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithWebAuthnMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *UpdateRegistrationFlowWithWebAuthnMethod) SetMethod(v string) {
	o.Method = v
}

// GetTraits returns the Traits field value
func (o *UpdateRegistrationFlowWithWebAuthnMethod) GetTraits() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithWebAuthnMethod) GetTraitsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Traits, true
}

// SetTraits sets field value
func (o *UpdateRegistrationFlowWithWebAuthnMethod) SetTraits(v map[string]interface{}) {
	o.Traits = v
}

// GetTransientPayload returns the TransientPayload field value if set, zero value otherwise.
func (o *UpdateRegistrationFlowWithWebAuthnMethod) GetTransientPayload() map[string]interface{} {
	if o == nil || o.TransientPayload == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.TransientPayload
}

// GetTransientPayloadOk returns a tuple with the TransientPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithWebAuthnMethod) GetTransientPayloadOk() (map[string]interface{}, bool) {
	if o == nil || o.TransientPayload == nil {
		return nil, false
	}
	return o.TransientPayload, true
}

// HasTransientPayload returns a boolean if a field has been set.
func (o *UpdateRegistrationFlowWithWebAuthnMethod) HasTransientPayload() bool {
	if o != nil && o.TransientPayload != nil {
		return true
	}

	return false
}

// SetTransientPayload gets a reference to the given map[string]interface{} and assigns it to the TransientPayload field.
func (o *UpdateRegistrationFlowWithWebAuthnMethod) SetTransientPayload(v map[string]interface{}) {
	o.TransientPayload = v
}

// GetWebauthnRegister returns the WebauthnRegister field value if set, zero value otherwise.
func (o *UpdateRegistrationFlowWithWebAuthnMethod) GetWebauthnRegister() string {
	if o == nil || o.WebauthnRegister == nil {
		var ret string
		return ret
	}
	return *o.WebauthnRegister
}

// GetWebauthnRegisterOk returns a tuple with the WebauthnRegister field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithWebAuthnMethod) GetWebauthnRegisterOk() (*string, bool) {
	if o == nil || o.WebauthnRegister == nil {
		return nil, false
	}
	return o.WebauthnRegister, true
}

// HasWebauthnRegister returns a boolean if a field has been set.
func (o *UpdateRegistrationFlowWithWebAuthnMethod) HasWebauthnRegister() bool {
	if o != nil && o.WebauthnRegister != nil {
		return true
	}

	return false
}

// SetWebauthnRegister gets a reference to the given string and assigns it to the WebauthnRegister field.
func (o *UpdateRegistrationFlowWithWebAuthnMethod) SetWebauthnRegister(v string) {
	o.WebauthnRegister = &v
}

// GetWebauthnRegisterDisplayname returns the WebauthnRegisterDisplayname field value if set, zero value otherwise.
func (o *UpdateRegistrationFlowWithWebAuthnMethod) GetWebauthnRegisterDisplayname() string {
	if o == nil || o.WebauthnRegisterDisplayname == nil {
		var ret string
		return ret
	}
	return *o.WebauthnRegisterDisplayname
}

// GetWebauthnRegisterDisplaynameOk returns a tuple with the WebauthnRegisterDisplayname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithWebAuthnMethod) GetWebauthnRegisterDisplaynameOk() (*string, bool) {
	if o == nil || o.WebauthnRegisterDisplayname == nil {
		return nil, false
	}
	return o.WebauthnRegisterDisplayname, true
}

// HasWebauthnRegisterDisplayname returns a boolean if a field has been set.
func (o *UpdateRegistrationFlowWithWebAuthnMethod) HasWebauthnRegisterDisplayname() bool {
	if o != nil && o.WebauthnRegisterDisplayname != nil {
		return true
	}

	return false
}

// SetWebauthnRegisterDisplayname gets a reference to the given string and assigns it to the WebauthnRegisterDisplayname field.
func (o *UpdateRegistrationFlowWithWebAuthnMethod) SetWebauthnRegisterDisplayname(v string) {
	o.WebauthnRegisterDisplayname = &v
}

func (o UpdateRegistrationFlowWithWebAuthnMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["traits"] = o.Traits
	}
	if o.TransientPayload != nil {
		toSerialize["transient_payload"] = o.TransientPayload
	}
	if o.WebauthnRegister != nil {
		toSerialize["webauthn_register"] = o.WebauthnRegister
	}
	if o.WebauthnRegisterDisplayname != nil {
		toSerialize["webauthn_register_displayname"] = o.WebauthnRegisterDisplayname
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateRegistrationFlowWithWebAuthnMethod struct {
	value *UpdateRegistrationFlowWithWebAuthnMethod
	isSet bool
}

func (v NullableUpdateRegistrationFlowWithWebAuthnMethod) Get() *UpdateRegistrationFlowWithWebAuthnMethod {
	return v.value
}

func (v *NullableUpdateRegistrationFlowWithWebAuthnMethod) Set(val *UpdateRegistrationFlowWithWebAuthnMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRegistrationFlowWithWebAuthnMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRegistrationFlowWithWebAuthnMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRegistrationFlowWithWebAuthnMethod(val *UpdateRegistrationFlowWithWebAuthnMethod) *NullableUpdateRegistrationFlowWithWebAuthnMethod {
	return &NullableUpdateRegistrationFlowWithWebAuthnMethod{value: val, isSet: true}
}

func (v NullableUpdateRegistrationFlowWithWebAuthnMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRegistrationFlowWithWebAuthnMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


