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

// checks if the ContinueWithSettingsUiFlow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContinueWithSettingsUiFlow{}

// ContinueWithSettingsUiFlow struct for ContinueWithSettingsUiFlow
type ContinueWithSettingsUiFlow struct {
	// The ID of the settings flow
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _ContinueWithSettingsUiFlow ContinueWithSettingsUiFlow

// NewContinueWithSettingsUiFlow instantiates a new ContinueWithSettingsUiFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinueWithSettingsUiFlow(id string) *ContinueWithSettingsUiFlow {
	this := ContinueWithSettingsUiFlow{}
	this.Id = id
	return &this
}

// NewContinueWithSettingsUiFlowWithDefaults instantiates a new ContinueWithSettingsUiFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinueWithSettingsUiFlowWithDefaults() *ContinueWithSettingsUiFlow {
	this := ContinueWithSettingsUiFlow{}
	return &this
}

// GetId returns the Id field value
func (o *ContinueWithSettingsUiFlow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContinueWithSettingsUiFlow) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContinueWithSettingsUiFlow) SetId(v string) {
	o.Id = v
}

func (o ContinueWithSettingsUiFlow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContinueWithSettingsUiFlow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContinueWithSettingsUiFlow) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varContinueWithSettingsUiFlow := _ContinueWithSettingsUiFlow{}

	err = json.Unmarshal(bytes, &varContinueWithSettingsUiFlow)

	if err != nil {
		return err
	}

	*o = ContinueWithSettingsUiFlow(varContinueWithSettingsUiFlow)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContinueWithSettingsUiFlow struct {
	value *ContinueWithSettingsUiFlow
	isSet bool
}

func (v NullableContinueWithSettingsUiFlow) Get() *ContinueWithSettingsUiFlow {
	return v.value
}

func (v *NullableContinueWithSettingsUiFlow) Set(val *ContinueWithSettingsUiFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableContinueWithSettingsUiFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableContinueWithSettingsUiFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinueWithSettingsUiFlow(val *ContinueWithSettingsUiFlow) *NullableContinueWithSettingsUiFlow {
	return &NullableContinueWithSettingsUiFlow{value: val, isSet: true}
}

func (v NullableContinueWithSettingsUiFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinueWithSettingsUiFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


