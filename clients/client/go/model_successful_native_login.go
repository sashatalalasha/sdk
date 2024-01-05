/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.4.9
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the SuccessfulNativeLogin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuccessfulNativeLogin{}

// SuccessfulNativeLogin The Response for Login Flows via API
type SuccessfulNativeLogin struct {
	Session Session `json:"session"`
	// The Session Token  A session token is equivalent to a session cookie, but it can be sent in the HTTP Authorization Header:  Authorization: bearer ${session-token}  The session token is only issued for API flows, not for Browser flows!
	SessionToken *string `json:"session_token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SuccessfulNativeLogin SuccessfulNativeLogin

// NewSuccessfulNativeLogin instantiates a new SuccessfulNativeLogin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessfulNativeLogin(session Session) *SuccessfulNativeLogin {
	this := SuccessfulNativeLogin{}
	this.Session = session
	return &this
}

// NewSuccessfulNativeLoginWithDefaults instantiates a new SuccessfulNativeLogin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessfulNativeLoginWithDefaults() *SuccessfulNativeLogin {
	this := SuccessfulNativeLogin{}
	return &this
}

// GetSession returns the Session field value
func (o *SuccessfulNativeLogin) GetSession() Session {
	if o == nil {
		var ret Session
		return ret
	}

	return o.Session
}

// GetSessionOk returns a tuple with the Session field value
// and a boolean to check if the value has been set.
func (o *SuccessfulNativeLogin) GetSessionOk() (*Session, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Session, true
}

// SetSession sets field value
func (o *SuccessfulNativeLogin) SetSession(v Session) {
	o.Session = v
}

// GetSessionToken returns the SessionToken field value if set, zero value otherwise.
func (o *SuccessfulNativeLogin) GetSessionToken() string {
	if o == nil || IsNil(o.SessionToken) {
		var ret string
		return ret
	}
	return *o.SessionToken
}

// GetSessionTokenOk returns a tuple with the SessionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulNativeLogin) GetSessionTokenOk() (*string, bool) {
	if o == nil || IsNil(o.SessionToken) {
		return nil, false
	}
	return o.SessionToken, true
}

// HasSessionToken returns a boolean if a field has been set.
func (o *SuccessfulNativeLogin) HasSessionToken() bool {
	if o != nil && !IsNil(o.SessionToken) {
		return true
	}

	return false
}

// SetSessionToken gets a reference to the given string and assigns it to the SessionToken field.
func (o *SuccessfulNativeLogin) SetSessionToken(v string) {
	o.SessionToken = &v
}

func (o SuccessfulNativeLogin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuccessfulNativeLogin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["session"] = o.Session
	if !IsNil(o.SessionToken) {
		toSerialize["session_token"] = o.SessionToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SuccessfulNativeLogin) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"session",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSuccessfulNativeLogin := _SuccessfulNativeLogin{}

	err = json.Unmarshal(bytes, &varSuccessfulNativeLogin)

	if err != nil {
		return err
	}

	*o = SuccessfulNativeLogin(varSuccessfulNativeLogin)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "session")
		delete(additionalProperties, "session_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSuccessfulNativeLogin struct {
	value *SuccessfulNativeLogin
	isSet bool
}

func (v NullableSuccessfulNativeLogin) Get() *SuccessfulNativeLogin {
	return v.value
}

func (v *NullableSuccessfulNativeLogin) Set(val *SuccessfulNativeLogin) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessfulNativeLogin) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessfulNativeLogin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessfulNativeLogin(val *SuccessfulNativeLogin) *NullableSuccessfulNativeLogin {
	return &NullableSuccessfulNativeLogin{value: val, isSet: true}
}

func (v NullableSuccessfulNativeLogin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessfulNativeLogin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


