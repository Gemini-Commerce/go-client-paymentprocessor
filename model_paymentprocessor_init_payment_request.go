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
	"bytes"
	"fmt"
)

// checks if the PaymentprocessorInitPaymentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentprocessorInitPaymentRequest{}

// PaymentprocessorInitPaymentRequest struct for PaymentprocessorInitPaymentRequest
type PaymentprocessorInitPaymentRequest struct {
	TenantId string `json:"tenantId"`
	Context *PaymentprocessorPaymentContext `json:"context,omitempty"`
	Code string `json:"code"`
	Amount *PaymentprocessorMoney `json:"amount,omitempty"`
	Currency *PaymentprocessorCurrency `json:"currency,omitempty"`
	AdditionalInfo *string `json:"additionalInfo,omitempty"`
}

type _PaymentprocessorInitPaymentRequest PaymentprocessorInitPaymentRequest

// NewPaymentprocessorInitPaymentRequest instantiates a new PaymentprocessorInitPaymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentprocessorInitPaymentRequest(tenantId string, code string) *PaymentprocessorInitPaymentRequest {
	this := PaymentprocessorInitPaymentRequest{}
	this.TenantId = tenantId
	this.Code = code
	var currency PaymentprocessorCurrency = XXX
	this.Currency = &currency
	return &this
}

// NewPaymentprocessorInitPaymentRequestWithDefaults instantiates a new PaymentprocessorInitPaymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentprocessorInitPaymentRequestWithDefaults() *PaymentprocessorInitPaymentRequest {
	this := PaymentprocessorInitPaymentRequest{}
	var currency PaymentprocessorCurrency = XXX
	this.Currency = &currency
	return &this
}

// GetTenantId returns the TenantId field value
func (o *PaymentprocessorInitPaymentRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *PaymentprocessorInitPaymentRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *PaymentprocessorInitPaymentRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *PaymentprocessorInitPaymentRequest) GetContext() PaymentprocessorPaymentContext {
	if o == nil || IsNil(o.Context) {
		var ret PaymentprocessorPaymentContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentprocessorInitPaymentRequest) GetContextOk() (*PaymentprocessorPaymentContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *PaymentprocessorInitPaymentRequest) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given PaymentprocessorPaymentContext and assigns it to the Context field.
func (o *PaymentprocessorInitPaymentRequest) SetContext(v PaymentprocessorPaymentContext) {
	o.Context = &v
}

// GetCode returns the Code field value
func (o *PaymentprocessorInitPaymentRequest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *PaymentprocessorInitPaymentRequest) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *PaymentprocessorInitPaymentRequest) SetCode(v string) {
	o.Code = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PaymentprocessorInitPaymentRequest) GetAmount() PaymentprocessorMoney {
	if o == nil || IsNil(o.Amount) {
		var ret PaymentprocessorMoney
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentprocessorInitPaymentRequest) GetAmountOk() (*PaymentprocessorMoney, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PaymentprocessorInitPaymentRequest) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given PaymentprocessorMoney and assigns it to the Amount field.
func (o *PaymentprocessorInitPaymentRequest) SetAmount(v PaymentprocessorMoney) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PaymentprocessorInitPaymentRequest) GetCurrency() PaymentprocessorCurrency {
	if o == nil || IsNil(o.Currency) {
		var ret PaymentprocessorCurrency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentprocessorInitPaymentRequest) GetCurrencyOk() (*PaymentprocessorCurrency, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PaymentprocessorInitPaymentRequest) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given PaymentprocessorCurrency and assigns it to the Currency field.
func (o *PaymentprocessorInitPaymentRequest) SetCurrency(v PaymentprocessorCurrency) {
	o.Currency = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *PaymentprocessorInitPaymentRequest) GetAdditionalInfo() string {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentprocessorInitPaymentRequest) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *PaymentprocessorInitPaymentRequest) HasAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *PaymentprocessorInitPaymentRequest) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

func (o PaymentprocessorInitPaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentprocessorInitPaymentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	toSerialize["code"] = o.Code
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	return toSerialize, nil
}

func (o *PaymentprocessorInitPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
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

	varPaymentprocessorInitPaymentRequest := _PaymentprocessorInitPaymentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymentprocessorInitPaymentRequest)

	if err != nil {
		return err
	}

	*o = PaymentprocessorInitPaymentRequest(varPaymentprocessorInitPaymentRequest)

	return err
}

type NullablePaymentprocessorInitPaymentRequest struct {
	value *PaymentprocessorInitPaymentRequest
	isSet bool
}

func (v NullablePaymentprocessorInitPaymentRequest) Get() *PaymentprocessorInitPaymentRequest {
	return v.value
}

func (v *NullablePaymentprocessorInitPaymentRequest) Set(val *PaymentprocessorInitPaymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentprocessorInitPaymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentprocessorInitPaymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentprocessorInitPaymentRequest(val *PaymentprocessorInitPaymentRequest) *NullablePaymentprocessorInitPaymentRequest {
	return &NullablePaymentprocessorInitPaymentRequest{value: val, isSet: true}
}

func (v NullablePaymentprocessorInitPaymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentprocessorInitPaymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


