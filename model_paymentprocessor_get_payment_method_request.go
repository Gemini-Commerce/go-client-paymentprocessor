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

// checks if the PaymentprocessorGetPaymentMethodRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentprocessorGetPaymentMethodRequest{}

// PaymentprocessorGetPaymentMethodRequest struct for PaymentprocessorGetPaymentMethodRequest
type PaymentprocessorGetPaymentMethodRequest struct {
	TenantId string `json:"tenantId"`
	Code string `json:"code"`
	AuthenticateMethod *bool `json:"authenticateMethod,omitempty"`
}

type _PaymentprocessorGetPaymentMethodRequest PaymentprocessorGetPaymentMethodRequest

// NewPaymentprocessorGetPaymentMethodRequest instantiates a new PaymentprocessorGetPaymentMethodRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentprocessorGetPaymentMethodRequest(tenantId string, code string) *PaymentprocessorGetPaymentMethodRequest {
	this := PaymentprocessorGetPaymentMethodRequest{}
	this.TenantId = tenantId
	this.Code = code
	return &this
}

// NewPaymentprocessorGetPaymentMethodRequestWithDefaults instantiates a new PaymentprocessorGetPaymentMethodRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentprocessorGetPaymentMethodRequestWithDefaults() *PaymentprocessorGetPaymentMethodRequest {
	this := PaymentprocessorGetPaymentMethodRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *PaymentprocessorGetPaymentMethodRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *PaymentprocessorGetPaymentMethodRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *PaymentprocessorGetPaymentMethodRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetCode returns the Code field value
func (o *PaymentprocessorGetPaymentMethodRequest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *PaymentprocessorGetPaymentMethodRequest) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *PaymentprocessorGetPaymentMethodRequest) SetCode(v string) {
	o.Code = v
}

// GetAuthenticateMethod returns the AuthenticateMethod field value if set, zero value otherwise.
func (o *PaymentprocessorGetPaymentMethodRequest) GetAuthenticateMethod() bool {
	if o == nil || IsNil(o.AuthenticateMethod) {
		var ret bool
		return ret
	}
	return *o.AuthenticateMethod
}

// GetAuthenticateMethodOk returns a tuple with the AuthenticateMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentprocessorGetPaymentMethodRequest) GetAuthenticateMethodOk() (*bool, bool) {
	if o == nil || IsNil(o.AuthenticateMethod) {
		return nil, false
	}
	return o.AuthenticateMethod, true
}

// HasAuthenticateMethod returns a boolean if a field has been set.
func (o *PaymentprocessorGetPaymentMethodRequest) HasAuthenticateMethod() bool {
	if o != nil && !IsNil(o.AuthenticateMethod) {
		return true
	}

	return false
}

// SetAuthenticateMethod gets a reference to the given bool and assigns it to the AuthenticateMethod field.
func (o *PaymentprocessorGetPaymentMethodRequest) SetAuthenticateMethod(v bool) {
	o.AuthenticateMethod = &v
}

func (o PaymentprocessorGetPaymentMethodRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentprocessorGetPaymentMethodRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["code"] = o.Code
	if !IsNil(o.AuthenticateMethod) {
		toSerialize["authenticateMethod"] = o.AuthenticateMethod
	}
	return toSerialize, nil
}

func (o *PaymentprocessorGetPaymentMethodRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPaymentprocessorGetPaymentMethodRequest := _PaymentprocessorGetPaymentMethodRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymentprocessorGetPaymentMethodRequest)

	if err != nil {
		return err
	}

	*o = PaymentprocessorGetPaymentMethodRequest(varPaymentprocessorGetPaymentMethodRequest)

	return err
}

type NullablePaymentprocessorGetPaymentMethodRequest struct {
	value *PaymentprocessorGetPaymentMethodRequest
	isSet bool
}

func (v NullablePaymentprocessorGetPaymentMethodRequest) Get() *PaymentprocessorGetPaymentMethodRequest {
	return v.value
}

func (v *NullablePaymentprocessorGetPaymentMethodRequest) Set(val *PaymentprocessorGetPaymentMethodRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentprocessorGetPaymentMethodRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentprocessorGetPaymentMethodRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentprocessorGetPaymentMethodRequest(val *PaymentprocessorGetPaymentMethodRequest) *NullablePaymentprocessorGetPaymentMethodRequest {
	return &NullablePaymentprocessorGetPaymentMethodRequest{value: val, isSet: true}
}

func (v NullablePaymentprocessorGetPaymentMethodRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentprocessorGetPaymentMethodRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

