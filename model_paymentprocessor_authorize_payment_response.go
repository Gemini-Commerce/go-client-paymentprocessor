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

// checks if the PaymentprocessorAuthorizePaymentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentprocessorAuthorizePaymentResponse{}

// PaymentprocessorAuthorizePaymentResponse struct for PaymentprocessorAuthorizePaymentResponse
type PaymentprocessorAuthorizePaymentResponse struct {
	Transaction *PaymentprocessorTransaction `json:"transaction,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentprocessorAuthorizePaymentResponse PaymentprocessorAuthorizePaymentResponse

// NewPaymentprocessorAuthorizePaymentResponse instantiates a new PaymentprocessorAuthorizePaymentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentprocessorAuthorizePaymentResponse() *PaymentprocessorAuthorizePaymentResponse {
	this := PaymentprocessorAuthorizePaymentResponse{}
	return &this
}

// NewPaymentprocessorAuthorizePaymentResponseWithDefaults instantiates a new PaymentprocessorAuthorizePaymentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentprocessorAuthorizePaymentResponseWithDefaults() *PaymentprocessorAuthorizePaymentResponse {
	this := PaymentprocessorAuthorizePaymentResponse{}
	return &this
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *PaymentprocessorAuthorizePaymentResponse) GetTransaction() PaymentprocessorTransaction {
	if o == nil || IsNil(o.Transaction) {
		var ret PaymentprocessorTransaction
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentprocessorAuthorizePaymentResponse) GetTransactionOk() (*PaymentprocessorTransaction, bool) {
	if o == nil || IsNil(o.Transaction) {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *PaymentprocessorAuthorizePaymentResponse) HasTransaction() bool {
	if o != nil && !IsNil(o.Transaction) {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given PaymentprocessorTransaction and assigns it to the Transaction field.
func (o *PaymentprocessorAuthorizePaymentResponse) SetTransaction(v PaymentprocessorTransaction) {
	o.Transaction = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *PaymentprocessorAuthorizePaymentResponse) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentprocessorAuthorizePaymentResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *PaymentprocessorAuthorizePaymentResponse) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *PaymentprocessorAuthorizePaymentResponse) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o PaymentprocessorAuthorizePaymentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentprocessorAuthorizePaymentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Transaction) {
		toSerialize["transaction"] = o.Transaction
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentprocessorAuthorizePaymentResponse) UnmarshalJSON(data []byte) (err error) {
	varPaymentprocessorAuthorizePaymentResponse := _PaymentprocessorAuthorizePaymentResponse{}

	err = json.Unmarshal(data, &varPaymentprocessorAuthorizePaymentResponse)

	if err != nil {
		return err
	}

	*o = PaymentprocessorAuthorizePaymentResponse(varPaymentprocessorAuthorizePaymentResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "transaction")
		delete(additionalProperties, "errorMessage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PaymentprocessorAuthorizePaymentResponse) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *PaymentprocessorAuthorizePaymentResponse) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullablePaymentprocessorAuthorizePaymentResponse struct {
	value *PaymentprocessorAuthorizePaymentResponse
	isSet bool
}

func (v NullablePaymentprocessorAuthorizePaymentResponse) Get() *PaymentprocessorAuthorizePaymentResponse {
	return v.value
}

func (v *NullablePaymentprocessorAuthorizePaymentResponse) Set(val *PaymentprocessorAuthorizePaymentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentprocessorAuthorizePaymentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentprocessorAuthorizePaymentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentprocessorAuthorizePaymentResponse(val *PaymentprocessorAuthorizePaymentResponse) *NullablePaymentprocessorAuthorizePaymentResponse {
	return &NullablePaymentprocessorAuthorizePaymentResponse{value: val, isSet: true}
}

func (v NullablePaymentprocessorAuthorizePaymentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentprocessorAuthorizePaymentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


