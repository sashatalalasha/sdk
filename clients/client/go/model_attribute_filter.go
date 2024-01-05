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
)

// checks if the AttributeFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributeFilter{}

// AttributeFilter struct for AttributeFilter
type AttributeFilter struct {
	Attribute *string `json:"attribute,omitempty"`
	Condition *string `json:"condition,omitempty"`
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AttributeFilter AttributeFilter

// NewAttributeFilter instantiates a new AttributeFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeFilter() *AttributeFilter {
	this := AttributeFilter{}
	return &this
}

// NewAttributeFilterWithDefaults instantiates a new AttributeFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeFilterWithDefaults() *AttributeFilter {
	this := AttributeFilter{}
	return &this
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *AttributeFilter) GetAttribute() string {
	if o == nil || IsNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeFilter) GetAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *AttributeFilter) HasAttribute() bool {
	if o != nil && !IsNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *AttributeFilter) SetAttribute(v string) {
	o.Attribute = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *AttributeFilter) GetCondition() string {
	if o == nil || IsNil(o.Condition) {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeFilter) GetConditionOk() (*string, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *AttributeFilter) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *AttributeFilter) SetCondition(v string) {
	o.Condition = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AttributeFilter) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeFilter) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AttributeFilter) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AttributeFilter) SetValue(v string) {
	o.Value = &v
}

func (o AttributeFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributeFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AttributeFilter) UnmarshalJSON(bytes []byte) (err error) {
	varAttributeFilter := _AttributeFilter{}

	err = json.Unmarshal(bytes, &varAttributeFilter)

	if err != nil {
		return err
	}

	*o = AttributeFilter(varAttributeFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "condition")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAttributeFilter struct {
	value *AttributeFilter
	isSet bool
}

func (v NullableAttributeFilter) Get() *AttributeFilter {
	return v.value
}

func (v *NullableAttributeFilter) Set(val *AttributeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeFilter(val *AttributeFilter) *NullableAttributeFilter {
	return &NullableAttributeFilter{value: val, isSet: true}
}

func (v NullableAttributeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


