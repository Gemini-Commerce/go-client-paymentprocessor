/*
Payment Processor Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paymentprocessor

import (
	"encoding/json"
)

// checks if the PaymentprocessorGetPaymentMethodConfigurationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentprocessorGetPaymentMethodConfigurationResponse{}

// PaymentprocessorGetPaymentMethodConfigurationResponse struct for PaymentprocessorGetPaymentMethodConfigurationResponse
type PaymentprocessorGetPaymentMethodConfigurationResponse struct {
	Configuration map[string]interface{} `json:"configuration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentprocessorGetPaymentMethodConfigurationResponse PaymentprocessorGetPaymentMethodConfigurationResponse

// NewPaymentprocessorGetPaymentMethodConfigurationResponse instantiates a new PaymentprocessorGetPaymentMethodConfigurationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentprocessorGetPaymentMethodConfigurationResponse() *PaymentprocessorGetPaymentMethodConfigurationResponse {
	this := PaymentprocessorGetPaymentMethodConfigurationResponse{}
	return &this
}

// NewPaymentprocessorGetPaymentMethodConfigurationResponseWithDefaults instantiates a new PaymentprocessorGetPaymentMethodConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentprocessorGetPaymentMethodConfigurationResponseWithDefaults() *PaymentprocessorGetPaymentMethodConfigurationResponse {
	this := PaymentprocessorGetPaymentMethodConfigurationResponse{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *PaymentprocessorGetPaymentMethodConfigurationResponse) GetConfiguration() map[string]interface{} {
	if o == nil || IsNil(o.Configuration) {
		var ret map[string]interface{}
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentprocessorGetPaymentMethodConfigurationResponse) GetConfigurationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Configuration) {
		return map[string]interface{}{}, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *PaymentprocessorGetPaymentMethodConfigurationResponse) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]interface{} and assigns it to the Configuration field.
func (o *PaymentprocessorGetPaymentMethodConfigurationResponse) SetConfiguration(v map[string]interface{}) {
	o.Configuration = v
}

func (o PaymentprocessorGetPaymentMethodConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentprocessorGetPaymentMethodConfigurationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentprocessorGetPaymentMethodConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	varPaymentprocessorGetPaymentMethodConfigurationResponse := _PaymentprocessorGetPaymentMethodConfigurationResponse{}

	err = json.Unmarshal(data, &varPaymentprocessorGetPaymentMethodConfigurationResponse)

	if err != nil {
		return err
	}

	*o = PaymentprocessorGetPaymentMethodConfigurationResponse(varPaymentprocessorGetPaymentMethodConfigurationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "configuration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PaymentprocessorGetPaymentMethodConfigurationResponse) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *PaymentprocessorGetPaymentMethodConfigurationResponse) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullablePaymentprocessorGetPaymentMethodConfigurationResponse struct {
	value *PaymentprocessorGetPaymentMethodConfigurationResponse
	isSet bool
}

func (v NullablePaymentprocessorGetPaymentMethodConfigurationResponse) Get() *PaymentprocessorGetPaymentMethodConfigurationResponse {
	return v.value
}

func (v *NullablePaymentprocessorGetPaymentMethodConfigurationResponse) Set(val *PaymentprocessorGetPaymentMethodConfigurationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentprocessorGetPaymentMethodConfigurationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentprocessorGetPaymentMethodConfigurationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentprocessorGetPaymentMethodConfigurationResponse(val *PaymentprocessorGetPaymentMethodConfigurationResponse) *NullablePaymentprocessorGetPaymentMethodConfigurationResponse {
	return &NullablePaymentprocessorGetPaymentMethodConfigurationResponse{value: val, isSet: true}
}

func (v NullablePaymentprocessorGetPaymentMethodConfigurationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentprocessorGetPaymentMethodConfigurationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


