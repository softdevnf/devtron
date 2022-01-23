/*
Devtron Labs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AppEnvironmentDetailAllOf struct for AppEnvironmentDetailAllOf
type AppEnvironmentDetailAllOf struct {
	// cluster corresponding to the environemt where application is deployed
	ClusterName *string `json:"clusterName,omitempty"`
	// clusterId corresponding to the environemt where application is deployed
	ClusterId *int32 `json:"clusterId,omitempty"`
}

// NewAppEnvironmentDetailAllOf instantiates a new AppEnvironmentDetailAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEnvironmentDetailAllOf() *AppEnvironmentDetailAllOf {
	this := AppEnvironmentDetailAllOf{}
	return &this
}

// NewAppEnvironmentDetailAllOfWithDefaults instantiates a new AppEnvironmentDetailAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEnvironmentDetailAllOfWithDefaults() *AppEnvironmentDetailAllOf {
	this := AppEnvironmentDetailAllOf{}
	return &this
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *AppEnvironmentDetailAllOf) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEnvironmentDetailAllOf) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *AppEnvironmentDetailAllOf) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *AppEnvironmentDetailAllOf) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *AppEnvironmentDetailAllOf) GetClusterId() int32 {
	if o == nil || o.ClusterId == nil {
		var ret int32
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEnvironmentDetailAllOf) GetClusterIdOk() (*int32, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *AppEnvironmentDetailAllOf) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given int32 and assigns it to the ClusterId field.
func (o *AppEnvironmentDetailAllOf) SetClusterId(v int32) {
	o.ClusterId = &v
}

func (o AppEnvironmentDetailAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.ClusterId != nil {
		toSerialize["clusterId"] = o.ClusterId
	}
	return json.Marshal(toSerialize)
}

type NullableAppEnvironmentDetailAllOf struct {
	value *AppEnvironmentDetailAllOf
	isSet bool
}

func (v NullableAppEnvironmentDetailAllOf) Get() *AppEnvironmentDetailAllOf {
	return v.value
}

func (v *NullableAppEnvironmentDetailAllOf) Set(val *AppEnvironmentDetailAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEnvironmentDetailAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEnvironmentDetailAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEnvironmentDetailAllOf(val *AppEnvironmentDetailAllOf) *NullableAppEnvironmentDetailAllOf {
	return &NullableAppEnvironmentDetailAllOf{value: val, isSet: true}
}

func (v NullableAppEnvironmentDetailAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEnvironmentDetailAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
