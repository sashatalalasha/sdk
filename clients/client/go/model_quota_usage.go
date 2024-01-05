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

// checks if the QuotaUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuotaUsage{}

// QuotaUsage struct for QuotaUsage
type QuotaUsage struct {
	// The additional price per unit in cents.
	AdditionalPrice int64 `json:"additional_price"`
	CanUseMore bool `json:"can_use_more"`
	//  region_eu RegionEU region_us RegionUS region_apac RegionAPAC region_global RegionGlobal production_projects ProductionProjects daily_active_users DailyActiveUsers custom_domains CustomDomains sla SLA collaborator_seats CollaboratorSeats edge_cache EdgeCache branding_themes BrandingThemes zendesk_support ZendeskSupport project_metrics ProjectMetrics project_metrics_time_window ProjectMetricsTimeWindow project_metrics_events_history ProjectMetricsEventsHistory organizations Organizations rop_grant ResourceOwnerPasswordGrant rate_limit_tier RateLimitTier session_rate_limit_tier RateLimitTierSessions identities_list_rate_limit_tier RateLimitTierIdentitiesList
	Feature string `json:"feature"`
	FeatureAvailable bool `json:"feature_available"`
	Included int64 `json:"included"`
	Used int64 `json:"used"`
	AdditionalProperties map[string]interface{}
}

type _QuotaUsage QuotaUsage

// NewQuotaUsage instantiates a new QuotaUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaUsage(additionalPrice int64, canUseMore bool, feature string, featureAvailable bool, included int64, used int64) *QuotaUsage {
	this := QuotaUsage{}
	this.AdditionalPrice = additionalPrice
	this.CanUseMore = canUseMore
	this.Feature = feature
	this.FeatureAvailable = featureAvailable
	this.Included = included
	this.Used = used
	return &this
}

// NewQuotaUsageWithDefaults instantiates a new QuotaUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaUsageWithDefaults() *QuotaUsage {
	this := QuotaUsage{}
	return &this
}

// GetAdditionalPrice returns the AdditionalPrice field value
func (o *QuotaUsage) GetAdditionalPrice() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AdditionalPrice
}

// GetAdditionalPriceOk returns a tuple with the AdditionalPrice field value
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetAdditionalPriceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdditionalPrice, true
}

// SetAdditionalPrice sets field value
func (o *QuotaUsage) SetAdditionalPrice(v int64) {
	o.AdditionalPrice = v
}

// GetCanUseMore returns the CanUseMore field value
func (o *QuotaUsage) GetCanUseMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanUseMore
}

// GetCanUseMoreOk returns a tuple with the CanUseMore field value
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetCanUseMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanUseMore, true
}

// SetCanUseMore sets field value
func (o *QuotaUsage) SetCanUseMore(v bool) {
	o.CanUseMore = v
}

// GetFeature returns the Feature field value
func (o *QuotaUsage) GetFeature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetFeatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Feature, true
}

// SetFeature sets field value
func (o *QuotaUsage) SetFeature(v string) {
	o.Feature = v
}

// GetFeatureAvailable returns the FeatureAvailable field value
func (o *QuotaUsage) GetFeatureAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FeatureAvailable
}

// GetFeatureAvailableOk returns a tuple with the FeatureAvailable field value
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetFeatureAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureAvailable, true
}

// SetFeatureAvailable sets field value
func (o *QuotaUsage) SetFeatureAvailable(v bool) {
	o.FeatureAvailable = v
}

// GetIncluded returns the Included field value
func (o *QuotaUsage) GetIncluded() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetIncludedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Included, true
}

// SetIncluded sets field value
func (o *QuotaUsage) SetIncluded(v int64) {
	o.Included = v
}

// GetUsed returns the Used field value
func (o *QuotaUsage) GetUsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Used
}

// GetUsedOk returns a tuple with the Used field value
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Used, true
}

// SetUsed sets field value
func (o *QuotaUsage) SetUsed(v int64) {
	o.Used = v
}

func (o QuotaUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuotaUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["additional_price"] = o.AdditionalPrice
	toSerialize["can_use_more"] = o.CanUseMore
	toSerialize["feature"] = o.Feature
	toSerialize["feature_available"] = o.FeatureAvailable
	toSerialize["included"] = o.Included
	toSerialize["used"] = o.Used

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuotaUsage) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"additional_price",
		"can_use_more",
		"feature",
		"feature_available",
		"included",
		"used",
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

	varQuotaUsage := _QuotaUsage{}

	err = json.Unmarshal(bytes, &varQuotaUsage)

	if err != nil {
		return err
	}

	*o = QuotaUsage(varQuotaUsage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "additional_price")
		delete(additionalProperties, "can_use_more")
		delete(additionalProperties, "feature")
		delete(additionalProperties, "feature_available")
		delete(additionalProperties, "included")
		delete(additionalProperties, "used")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuotaUsage struct {
	value *QuotaUsage
	isSet bool
}

func (v NullableQuotaUsage) Get() *QuotaUsage {
	return v.value
}

func (v *NullableQuotaUsage) Set(val *QuotaUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaUsage(val *QuotaUsage) *NullableQuotaUsage {
	return &NullableQuotaUsage{value: val, isSet: true}
}

func (v NullableQuotaUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


