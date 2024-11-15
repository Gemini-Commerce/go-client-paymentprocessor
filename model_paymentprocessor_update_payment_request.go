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
	"fmt"
)

// checks if the PaymentprocessorUpdatePaymentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentprocessorUpdatePaymentRequest{}

// PaymentprocessorUpdatePaymentRequest struct for PaymentprocessorUpdatePaymentRequest
type PaymentprocessorUpdatePaymentRequest struct {
	TenantId string `json:"tenantId"`
	PaymentId string `json:"paymentId"`
	// payment type stripe, paypal..
	Code string `json:"code"`
	AdditionalInfo *string `json:"additionalInfo,omitempty"`
	Transactions []PaymentprocessorTransaction `json:"transactions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentprocessorUpdatePaymentRequest PaymentprocessorUpdatePaymentRequest

// NewPaymentprocessorUpdatePaymentRequest instantiates a new PaymentprocessorUpdatePaymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentprocessorUpdatePaymentRequest(tenantId string, paymentId string, code string) *PaymentprocessorUpdatePaymentRequest {
	this := PaymentprocessorUpdatePaymentRequest{}
	this.TenantId = tenantId
	this.PaymentId = paymentId
	this.Code = code
	return &this
}

// NewPaymentprocessorUpdatePaymentRequestWithDefaults instantiates a new PaymentprocessorUpdatePaymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentprocessorUpdatePaymentRequestWithDefaults() *PaymentprocessorUpdatePaymentRequest {
	this := PaymentprocessorUpdatePaymentRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *PaymentprocessorUpdatePaymentRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *PaymentprocessorUpdatePaymentRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *PaymentprocessorUpdatePaymentRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetPaymentId returns the PaymentId field value
func (o *PaymentprocessorUpdatePaymentRequest) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *PaymentprocessorUpdatePaymentRequest) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *PaymentprocessorUpdatePaymentRequest) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetCode returns the Code field value
func (o *PaymentprocessorUpdatePaymentRequest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *PaymentprocessorUpdatePaymentRequest) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *PaymentprocessorUpdatePaymentRequest) SetCode(v string) {
	o.Code = v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *PaymentprocessorUpdatePaymentRequest) GetAdditionalInfo() string {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentprocessorUpdatePaymentRequest) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *PaymentprocessorUpdatePaymentRequest) HasAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *PaymentprocessorUpdatePaymentRequest) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *PaymentprocessorUpdatePaymentRequest) GetTransactions() []PaymentprocessorTransaction {
	if o == nil || IsNil(o.Transactions) {
		var ret []PaymentprocessorTransaction
		return ret
	}
	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentprocessorUpdatePaymentRequest) GetTransactionsOk() ([]PaymentprocessorTransaction, bool) {
	if o == nil || IsNil(o.Transactions) {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *PaymentprocessorUpdatePaymentRequest) HasTransactions() bool {
	if o != nil && !IsNil(o.Transactions) {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given []PaymentprocessorTransaction and assigns it to the Transactions field.
func (o *PaymentprocessorUpdatePaymentRequest) SetTransactions(v []PaymentprocessorTransaction) {
	o.Transactions = v
}

func (o PaymentprocessorUpdatePaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentprocessorUpdatePaymentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["paymentId"] = o.PaymentId
	toSerialize["code"] = o.Code
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	if !IsNil(o.Transactions) {
		toSerialize["transactions"] = o.Transactions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentprocessorUpdatePaymentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"paymentId",
		"code",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPaymentprocessorUpdatePaymentRequest := _PaymentprocessorUpdatePaymentRequest{}

	err = json.Unmarshal(data, &varPaymentprocessorUpdatePaymentRequest)

	if err != nil {
		return err
	}

	*o = PaymentprocessorUpdatePaymentRequest(varPaymentprocessorUpdatePaymentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "paymentId")
		delete(additionalProperties, "code")
		delete(additionalProperties, "additionalInfo")
		delete(additionalProperties, "transactions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PaymentprocessorUpdatePaymentRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *PaymentprocessorUpdatePaymentRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullablePaymentprocessorUpdatePaymentRequest struct {
	value *PaymentprocessorUpdatePaymentRequest
	isSet bool
}

func (v NullablePaymentprocessorUpdatePaymentRequest) Get() *PaymentprocessorUpdatePaymentRequest {
	return v.value
}

func (v *NullablePaymentprocessorUpdatePaymentRequest) Set(val *PaymentprocessorUpdatePaymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentprocessorUpdatePaymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentprocessorUpdatePaymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentprocessorUpdatePaymentRequest(val *PaymentprocessorUpdatePaymentRequest) *NullablePaymentprocessorUpdatePaymentRequest {
	return &NullablePaymentprocessorUpdatePaymentRequest{value: val, isSet: true}
}

func (v NullablePaymentprocessorUpdatePaymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentprocessorUpdatePaymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


