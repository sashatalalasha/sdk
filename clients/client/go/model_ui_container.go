/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.7
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UiContainer Container represents a HTML Form. The container can work with both HTTP Form and JSON requests
type UiContainer struct {
	// Action should be used as the form action URL `<form action=\"{{ .Action }}\" method=\"post\">`.
	Action string `json:"action"`
	Messages []UiText `json:"messages,omitempty"`
	// Method is the form method (e.g. POST)
	Method string `json:"method"`
	Nodes []UiNode `json:"nodes"`
	AdditionalProperties map[string]interface{}
}

type _UiContainer UiContainer

// NewUiContainer instantiates a new UiContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiContainer(action string, method string, nodes []UiNode) *UiContainer {
	this := UiContainer{}
	this.Action = action
	this.Method = method
	this.Nodes = nodes
	return &this
}

// NewUiContainerWithDefaults instantiates a new UiContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiContainerWithDefaults() *UiContainer {
	this := UiContainer{}
	return &this
}

// GetAction returns the Action field value
func (o *UiContainer) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *UiContainer) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *UiContainer) SetAction(v string) {
	o.Action = v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *UiContainer) GetMessages() []UiText {
	if o == nil || o.Messages == nil {
		var ret []UiText
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiContainer) GetMessagesOk() ([]UiText, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *UiContainer) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []UiText and assigns it to the Messages field.
func (o *UiContainer) SetMessages(v []UiText) {
	o.Messages = v
}

// GetMethod returns the Method field value
func (o *UiContainer) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *UiContainer) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *UiContainer) SetMethod(v string) {
	o.Method = v
}

// GetNodes returns the Nodes field value
func (o *UiContainer) GetNodes() []UiNode {
	if o == nil {
		var ret []UiNode
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *UiContainer) GetNodesOk() ([]UiNode, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nodes, true
}

// SetNodes sets field value
func (o *UiContainer) SetNodes(v []UiNode) {
	o.Nodes = v
}

func (o UiContainer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["action"] = o.Action
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["nodes"] = o.Nodes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UiContainer) UnmarshalJSON(bytes []byte) (err error) {
	varUiContainer := _UiContainer{}

	if err = json.Unmarshal(bytes, &varUiContainer); err == nil {
		*o = UiContainer(varUiContainer)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "messages")
		delete(additionalProperties, "method")
		delete(additionalProperties, "nodes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUiContainer struct {
	value *UiContainer
	isSet bool
}

func (v NullableUiContainer) Get() *UiContainer {
	return v.value
}

func (v *NullableUiContainer) Set(val *UiContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableUiContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableUiContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiContainer(val *UiContainer) *NullableUiContainer {
	return &NullableUiContainer{value: val, isSet: true}
}

func (v NullableUiContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


