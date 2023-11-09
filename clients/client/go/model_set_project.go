/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.17
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SetProject struct for SetProject
type SetProject struct {
	CorsAdmin ProjectCors `json:"cors_admin"`
	CorsPublic ProjectCors `json:"cors_public"`
	// The name of the project.
	Name string `json:"name"`
	Services ProjectServices `json:"services"`
	AdditionalProperties map[string]interface{}
}

type _SetProject SetProject

// NewSetProject instantiates a new SetProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetProject(corsAdmin ProjectCors, corsPublic ProjectCors, name string, services ProjectServices) *SetProject {
	this := SetProject{}
	this.CorsAdmin = corsAdmin
	this.CorsPublic = corsPublic
	this.Name = name
	this.Services = services
	return &this
}

// NewSetProjectWithDefaults instantiates a new SetProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetProjectWithDefaults() *SetProject {
	this := SetProject{}
	return &this
}

// GetCorsAdmin returns the CorsAdmin field value
func (o *SetProject) GetCorsAdmin() ProjectCors {
	if o == nil {
		var ret ProjectCors
		return ret
	}

	return o.CorsAdmin
}

// GetCorsAdminOk returns a tuple with the CorsAdmin field value
// and a boolean to check if the value has been set.
func (o *SetProject) GetCorsAdminOk() (*ProjectCors, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CorsAdmin, true
}

// SetCorsAdmin sets field value
func (o *SetProject) SetCorsAdmin(v ProjectCors) {
	o.CorsAdmin = v
}

// GetCorsPublic returns the CorsPublic field value
func (o *SetProject) GetCorsPublic() ProjectCors {
	if o == nil {
		var ret ProjectCors
		return ret
	}

	return o.CorsPublic
}

// GetCorsPublicOk returns a tuple with the CorsPublic field value
// and a boolean to check if the value has been set.
func (o *SetProject) GetCorsPublicOk() (*ProjectCors, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CorsPublic, true
}

// SetCorsPublic sets field value
func (o *SetProject) SetCorsPublic(v ProjectCors) {
	o.CorsPublic = v
}

// GetName returns the Name field value
func (o *SetProject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SetProject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SetProject) SetName(v string) {
	o.Name = v
}

// GetServices returns the Services field value
func (o *SetProject) GetServices() ProjectServices {
	if o == nil {
		var ret ProjectServices
		return ret
	}

	return o.Services
}

// GetServicesOk returns a tuple with the Services field value
// and a boolean to check if the value has been set.
func (o *SetProject) GetServicesOk() (*ProjectServices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Services, true
}

// SetServices sets field value
func (o *SetProject) SetServices(v ProjectServices) {
	o.Services = v
}

func (o SetProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cors_admin"] = o.CorsAdmin
	}
	if true {
		toSerialize["cors_public"] = o.CorsPublic
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["services"] = o.Services
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SetProject) UnmarshalJSON(bytes []byte) (err error) {
	varSetProject := _SetProject{}

	if err = json.Unmarshal(bytes, &varSetProject); err == nil {
		*o = SetProject(varSetProject)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cors_admin")
		delete(additionalProperties, "cors_public")
		delete(additionalProperties, "name")
		delete(additionalProperties, "services")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetProject struct {
	value *SetProject
	isSet bool
}

func (v NullableSetProject) Get() *SetProject {
	return v.value
}

func (v *NullableSetProject) Set(val *SetProject) {
	v.value = val
	v.isSet = true
}

func (v NullableSetProject) IsSet() bool {
	return v.isSet
}

func (v *NullableSetProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetProject(val *SetProject) *NullableSetProject {
	return &NullableSetProject{value: val, isSet: true}
}

func (v NullableSetProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


