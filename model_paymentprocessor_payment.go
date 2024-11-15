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

// checks if the PaymentprocessorPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentprocessorPayment{}

// PaymentprocessorPayment struct for PaymentprocessorPayment
type PaymentprocessorPayment struct {
	Id string `json:"id"`
	// payment type stripe, paypal..
	Code string `json:"code"`
	AdditionalInfo *string `json:"additionalInfo,omitempty"`
	Transactions []PaymentprocessorTransaction `json:"transactions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentprocessorPayment PaymentprocessorPayment

// NewPaymentprocessorPayment instantiates a new PaymentprocessorPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentprocessorPayment(id string, code string) *PaymentprocessorPayment {
	this := PaymentprocessorPayment{}
	this.Id = id
	this.Code = code
	return &this
}

// NewPaymentprocessorPaymentWithDefaults instantiates a new PaymentprocessorPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentprocessorPaymentWithDefaults() *PaymentprocessorPayment {
	this := PaymentprocessorPayment{}
	return &this
}

// GetId returns the Id field value
func (o *PaymentprocessorPayment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaymentprocessorPayment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaymentprocessorPayment) SetId(v string) {
	o.Id = v
}

// GetCode returns the Code field value
func (o *PaymentprocessorPayment) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *PaymentprocessorPayment) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *PaymentprocessorPayment) SetCode(v string) {
	o.Code = v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *PaymentprocessorPayment) GetAdditionalInfo() string {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentprocessorPayment) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *PaymentprocessorPayment) HasAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *PaymentprocessorPayment) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *PaymentprocessorPayment) GetTransactions() []PaymentprocessorTransaction {
	if o == nil || IsNil(o.Transactions) {
		var ret []PaymentprocessorTransaction
		return ret
	}
	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentprocessorPayment) GetTransactionsOk() ([]PaymentprocessorTransaction, bool) {
	if o == nil || IsNil(o.Transactions) {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *PaymentprocessorPayment) HasTransactions() bool {
	if o != nil && !IsNil(o.Transactions) {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given []PaymentprocessorTransaction and assigns it to the Transactions field.
func (o *PaymentprocessorPayment) SetTransactions(v []PaymentprocessorTransaction) {
	o.Transactions = v
}

func (o PaymentprocessorPayment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentprocessorPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
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

func (o *PaymentprocessorPayment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varPaymentprocessorPayment := _PaymentprocessorPayment{}

	err = json.Unmarshal(data, &varPaymentprocessorPayment)

	if err != nil {
		return err
	}

	*o = PaymentprocessorPayment(varPaymentprocessorPayment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "code")
		delete(additionalProperties, "additionalInfo")
		delete(additionalProperties, "transactions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PaymentprocessorPayment) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *PaymentprocessorPayment) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullablePaymentprocessorPayment struct {
	value *PaymentprocessorPayment
	isSet bool
}

func (v NullablePaymentprocessorPayment) Get() *PaymentprocessorPayment {
	return v.value
}

func (v *NullablePaymentprocessorPayment) Set(val *PaymentprocessorPayment) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentprocessorPayment) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentprocessorPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentprocessorPayment(val *PaymentprocessorPayment) *NullablePaymentprocessorPayment {
	return &NullablePaymentprocessorPayment{value: val, isSet: true}
}

func (v NullablePaymentprocessorPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentprocessorPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


