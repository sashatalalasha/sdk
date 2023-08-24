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
	"time"
)

// SelfServiceFlowExpiredError Is sent when a flow is expired
type SelfServiceFlowExpiredError struct {
	Error *GenericError `json:"error,omitempty"`
	// When the flow has expired
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
	// A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years.
	Since *int64 `json:"since,omitempty"`
	// The flow ID that should be used for the new flow as it contains the correct messages.
	UseFlowId *string `json:"use_flow_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SelfServiceFlowExpiredError SelfServiceFlowExpiredError

// NewSelfServiceFlowExpiredError instantiates a new SelfServiceFlowExpiredError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceFlowExpiredError() *SelfServiceFlowExpiredError {
	this := SelfServiceFlowExpiredError{}
	return &this
}

// NewSelfServiceFlowExpiredErrorWithDefaults instantiates a new SelfServiceFlowExpiredError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceFlowExpiredErrorWithDefaults() *SelfServiceFlowExpiredError {
	this := SelfServiceFlowExpiredError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SelfServiceFlowExpiredError) GetError() GenericError {
	if o == nil || o.Error == nil {
		var ret GenericError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceFlowExpiredError) GetErrorOk() (*GenericError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SelfServiceFlowExpiredError) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given GenericError and assigns it to the Error field.
func (o *SelfServiceFlowExpiredError) SetError(v GenericError) {
	o.Error = &v
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise.
func (o *SelfServiceFlowExpiredError) GetExpiredAt() time.Time {
	if o == nil || o.ExpiredAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceFlowExpiredError) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiredAt == nil {
		return nil, false
	}
	return o.ExpiredAt, true
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *SelfServiceFlowExpiredError) HasExpiredAt() bool {
	if o != nil && o.ExpiredAt != nil {
		return true
	}

	return false
}

// SetExpiredAt gets a reference to the given time.Time and assigns it to the ExpiredAt field.
func (o *SelfServiceFlowExpiredError) SetExpiredAt(v time.Time) {
	o.ExpiredAt = &v
}

// GetSince returns the Since field value if set, zero value otherwise.
func (o *SelfServiceFlowExpiredError) GetSince() int64 {
	if o == nil || o.Since == nil {
		var ret int64
		return ret
	}
	return *o.Since
}

// GetSinceOk returns a tuple with the Since field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceFlowExpiredError) GetSinceOk() (*int64, bool) {
	if o == nil || o.Since == nil {
		return nil, false
	}
	return o.Since, true
}

// HasSince returns a boolean if a field has been set.
func (o *SelfServiceFlowExpiredError) HasSince() bool {
	if o != nil && o.Since != nil {
		return true
	}

	return false
}

// SetSince gets a reference to the given int64 and assigns it to the Since field.
func (o *SelfServiceFlowExpiredError) SetSince(v int64) {
	o.Since = &v
}

// GetUseFlowId returns the UseFlowId field value if set, zero value otherwise.
func (o *SelfServiceFlowExpiredError) GetUseFlowId() string {
	if o == nil || o.UseFlowId == nil {
		var ret string
		return ret
	}
	return *o.UseFlowId
}

// GetUseFlowIdOk returns a tuple with the UseFlowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceFlowExpiredError) GetUseFlowIdOk() (*string, bool) {
	if o == nil || o.UseFlowId == nil {
		return nil, false
	}
	return o.UseFlowId, true
}

// HasUseFlowId returns a boolean if a field has been set.
func (o *SelfServiceFlowExpiredError) HasUseFlowId() bool {
	if o != nil && o.UseFlowId != nil {
		return true
	}

	return false
}

// SetUseFlowId gets a reference to the given string and assigns it to the UseFlowId field.
func (o *SelfServiceFlowExpiredError) SetUseFlowId(v string) {
	o.UseFlowId = &v
}

func (o SelfServiceFlowExpiredError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.ExpiredAt != nil {
		toSerialize["expired_at"] = o.ExpiredAt
	}
	if o.Since != nil {
		toSerialize["since"] = o.Since
	}
	if o.UseFlowId != nil {
		toSerialize["use_flow_id"] = o.UseFlowId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SelfServiceFlowExpiredError) UnmarshalJSON(bytes []byte) (err error) {
	varSelfServiceFlowExpiredError := _SelfServiceFlowExpiredError{}

	if err = json.Unmarshal(bytes, &varSelfServiceFlowExpiredError); err == nil {
		*o = SelfServiceFlowExpiredError(varSelfServiceFlowExpiredError)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "expired_at")
		delete(additionalProperties, "since")
		delete(additionalProperties, "use_flow_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSelfServiceFlowExpiredError struct {
	value *SelfServiceFlowExpiredError
	isSet bool
}

func (v NullableSelfServiceFlowExpiredError) Get() *SelfServiceFlowExpiredError {
	return v.value
}

func (v *NullableSelfServiceFlowExpiredError) Set(val *SelfServiceFlowExpiredError) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceFlowExpiredError) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceFlowExpiredError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceFlowExpiredError(val *SelfServiceFlowExpiredError) *NullableSelfServiceFlowExpiredError {
	return &NullableSelfServiceFlowExpiredError{value: val, isSet: true}
}

func (v NullableSelfServiceFlowExpiredError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceFlowExpiredError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


