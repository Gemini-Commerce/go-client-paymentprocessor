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

// checks if the PaymentprocessorPerformRefundRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentprocessorPerformRefundRequest{}

// PaymentprocessorPerformRefundRequest struct for PaymentprocessorPerformRefundRequest
type PaymentprocessorPerformRefundRequest struct {
	TenantId string `json:"tenantId"`
	RefundId string `json:"refundId"`
	Payment PaymentprocessorPayment `json:"payment"`
	Amount *PaymentprocessorMoney `json:"amount,omitempty"`
	Currency *PaymentprocessorCurrency `json:"currency,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentprocessorPerformRefundRequest PaymentprocessorPerformRefundRequest

// NewPaymentprocessorPerformRefundRequest instantiates a new PaymentprocessorPerformRefundRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentprocessorPerformRefundRequest(tenantId string, refundId string, payment PaymentprocessorPayment) *PaymentprocessorPerformRefundRequest {
	this := PaymentprocessorPerformRefundRequest{}
	this.TenantId = tenantId
	this.RefundId = refundId
	this.Payment = payment
	var currency PaymentprocessorCurrency = PAYMENTPROCESSORCURRENCY_XXX
	this.Currency = &currency
	return &this
}

// NewPaymentprocessorPerformRefundRequestWithDefaults instantiates a new PaymentprocessorPerformRefundRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentprocessorPerformRefundRequestWithDefaults() *PaymentprocessorPerformRefundRequest {
	this := PaymentprocessorPerformRefundRequest{}
	var currency PaymentprocessorCurrency = PAYMENTPROCESSORCURRENCY_XXX
	this.Currency = &currency
	return &this
}

// GetTenantId returns the TenantId field value
func (o *PaymentprocessorPerformRefundRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *PaymentprocessorPerformRefundRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *PaymentprocessorPerformRefundRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetRefundId returns the RefundId field value
func (o *PaymentprocessorPerformRefundRequest) GetRefundId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefundId
}

// GetRefundIdOk returns a tuple with the RefundId field value
// and a boolean to check if the value has been set.
func (o *PaymentprocessorPerformRefundRequest) GetRefundIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefundId, true
}

// SetRefundId sets field value
func (o *PaymentprocessorPerformRefundRequest) SetRefundId(v string) {
	o.RefundId = v
}

// GetPayment returns the Payment field value
func (o *PaymentprocessorPerformRefundRequest) GetPayment() PaymentprocessorPayment {
	if o == nil {
		var ret PaymentprocessorPayment
		return ret
	}

	return o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value
// and a boolean to check if the value has been set.
func (o *PaymentprocessorPerformRefundRequest) GetPaymentOk() (*PaymentprocessorPayment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payment, true
}

// SetPayment sets field value
func (o *PaymentprocessorPerformRefundRequest) SetPayment(v PaymentprocessorPayment) {
	o.Payment = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PaymentprocessorPerformRefundRequest) GetAmount() PaymentprocessorMoney {
	if o == nil || IsNil(o.Amount) {
		var ret PaymentprocessorMoney
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentprocessorPerformRefundRequest) GetAmountOk() (*PaymentprocessorMoney, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PaymentprocessorPerformRefundRequest) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given PaymentprocessorMoney and assigns it to the Amount field.
func (o *PaymentprocessorPerformRefundRequest) SetAmount(v PaymentprocessorMoney) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PaymentprocessorPerformRefundRequest) GetCurrency() PaymentprocessorCurrency {
	if o == nil || IsNil(o.Currency) {
		var ret PaymentprocessorCurrency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentprocessorPerformRefundRequest) GetCurrencyOk() (*PaymentprocessorCurrency, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PaymentprocessorPerformRefundRequest) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given PaymentprocessorCurrency and assigns it to the Currency field.
func (o *PaymentprocessorPerformRefundRequest) SetCurrency(v PaymentprocessorCurrency) {
	o.Currency = &v
}

func (o PaymentprocessorPerformRefundRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentprocessorPerformRefundRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["refundId"] = o.RefundId
	toSerialize["payment"] = o.Payment
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentprocessorPerformRefundRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"refundId",
		"payment",
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

	varPaymentprocessorPerformRefundRequest := _PaymentprocessorPerformRefundRequest{}

	err = json.Unmarshal(data, &varPaymentprocessorPerformRefundRequest)

	if err != nil {
		return err
	}

	*o = PaymentprocessorPerformRefundRequest(varPaymentprocessorPerformRefundRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "refundId")
		delete(additionalProperties, "payment")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "currency")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PaymentprocessorPerformRefundRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *PaymentprocessorPerformRefundRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullablePaymentprocessorPerformRefundRequest struct {
	value *PaymentprocessorPerformRefundRequest
	isSet bool
}

func (v NullablePaymentprocessorPerformRefundRequest) Get() *PaymentprocessorPerformRefundRequest {
	return v.value
}

func (v *NullablePaymentprocessorPerformRefundRequest) Set(val *PaymentprocessorPerformRefundRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentprocessorPerformRefundRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentprocessorPerformRefundRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentprocessorPerformRefundRequest(val *PaymentprocessorPerformRefundRequest) *NullablePaymentprocessorPerformRefundRequest {
	return &NullablePaymentprocessorPerformRefundRequest{value: val, isSet: true}
}

func (v NullablePaymentprocessorPerformRefundRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentprocessorPerformRefundRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


