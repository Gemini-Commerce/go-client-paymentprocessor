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

// checks if the PaymentprocessorInitPaymentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentprocessorInitPaymentResponse{}

// PaymentprocessorInitPaymentResponse struct for PaymentprocessorInitPaymentResponse
type PaymentprocessorInitPaymentResponse struct {
	Payload *string `json:"payload,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentprocessorInitPaymentResponse PaymentprocessorInitPaymentResponse

// NewPaymentprocessorInitPaymentResponse instantiates a new PaymentprocessorInitPaymentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentprocessorInitPaymentResponse() *PaymentprocessorInitPaymentResponse {
	this := PaymentprocessorInitPaymentResponse{}
	return &this
}

// NewPaymentprocessorInitPaymentResponseWithDefaults instantiates a new PaymentprocessorInitPaymentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentprocessorInitPaymentResponseWithDefaults() *PaymentprocessorInitPaymentResponse {
	this := PaymentprocessorInitPaymentResponse{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *PaymentprocessorInitPaymentResponse) GetPayload() string {
	if o == nil || IsNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentprocessorInitPaymentResponse) GetPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// &#39;Has&#39;Payload returns a boolean if a field has been set.
func (o *PaymentprocessorInitPaymentResponse) &#39;Has&#39;Payload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *PaymentprocessorInitPaymentResponse) SetPayload(v string) {
	o.Payload = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *PaymentprocessorInitPaymentResponse) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentprocessorInitPaymentResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// &#39;Has&#39;ErrorMessage returns a boolean if a field has been set.
func (o *PaymentprocessorInitPaymentResponse) &#39;Has&#39;ErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *PaymentprocessorInitPaymentResponse) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o PaymentprocessorInitPaymentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentprocessorInitPaymentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentprocessorInitPaymentResponse) UnmarshalJSON(data []byte) (err error) {
	varPaymentprocessorInitPaymentResponse := _PaymentprocessorInitPaymentResponse{}

	err = json.Unmarshal(data, &varPaymentprocessorInitPaymentResponse)

	if err != nil {
		return err
	}

	*o = PaymentprocessorInitPaymentResponse(varPaymentprocessorInitPaymentResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "payload")
		delete(additionalProperties, "errorMessage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PaymentprocessorInitPaymentResponse) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *PaymentprocessorInitPaymentResponse) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullablePaymentprocessorInitPaymentResponse struct {
	value *PaymentprocessorInitPaymentResponse
	isSet bool
}

func (v NullablePaymentprocessorInitPaymentResponse) Get() *PaymentprocessorInitPaymentResponse {
	return v.value
}

func (v *NullablePaymentprocessorInitPaymentResponse) Set(val *PaymentprocessorInitPaymentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentprocessorInitPaymentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentprocessorInitPaymentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentprocessorInitPaymentResponse(val *PaymentprocessorInitPaymentResponse) *NullablePaymentprocessorInitPaymentResponse {
	return &NullablePaymentprocessorInitPaymentResponse{value: val, isSet: true}
}

func (v NullablePaymentprocessorInitPaymentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentprocessorInitPaymentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


