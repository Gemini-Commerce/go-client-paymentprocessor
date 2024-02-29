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

// checks if the PaymentprocessorGetAvailablePaymentMethodRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentprocessorGetAvailablePaymentMethodRequest{}

// PaymentprocessorGetAvailablePaymentMethodRequest struct for PaymentprocessorGetAvailablePaymentMethodRequest
type PaymentprocessorGetAvailablePaymentMethodRequest struct {
	TenantId string `json:"tenantId"`
	Code string `json:"code"`
	Context PaymentprocessorAvailabilityContext `json:"context"`
	AuthenticateMethod *bool `json:"authenticateMethod,omitempty"`
}

type _PaymentprocessorGetAvailablePaymentMethodRequest PaymentprocessorGetAvailablePaymentMethodRequest

// NewPaymentprocessorGetAvailablePaymentMethodRequest instantiates a new PaymentprocessorGetAvailablePaymentMethodRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentprocessorGetAvailablePaymentMethodRequest(tenantId string, code string, context PaymentprocessorAvailabilityContext) *PaymentprocessorGetAvailablePaymentMethodRequest {
	this := PaymentprocessorGetAvailablePaymentMethodRequest{}
	this.TenantId = tenantId
	this.Code = code
	this.Context = context
	return &this
}

// NewPaymentprocessorGetAvailablePaymentMethodRequestWithDefaults instantiates a new PaymentprocessorGetAvailablePaymentMethodRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentprocessorGetAvailablePaymentMethodRequestWithDefaults() *PaymentprocessorGetAvailablePaymentMethodRequest {
	this := PaymentprocessorGetAvailablePaymentMethodRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *PaymentprocessorGetAvailablePaymentMethodRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *PaymentprocessorGetAvailablePaymentMethodRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *PaymentprocessorGetAvailablePaymentMethodRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetCode returns the Code field value
func (o *PaymentprocessorGetAvailablePaymentMethodRequest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *PaymentprocessorGetAvailablePaymentMethodRequest) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *PaymentprocessorGetAvailablePaymentMethodRequest) SetCode(v string) {
	o.Code = v
}

// GetContext returns the Context field value
func (o *PaymentprocessorGetAvailablePaymentMethodRequest) GetContext() PaymentprocessorAvailabilityContext {
	if o == nil {
		var ret PaymentprocessorAvailabilityContext
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *PaymentprocessorGetAvailablePaymentMethodRequest) GetContextOk() (*PaymentprocessorAvailabilityContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *PaymentprocessorGetAvailablePaymentMethodRequest) SetContext(v PaymentprocessorAvailabilityContext) {
	o.Context = v
}

// GetAuthenticateMethod returns the AuthenticateMethod field value if set, zero value otherwise.
func (o *PaymentprocessorGetAvailablePaymentMethodRequest) GetAuthenticateMethod() bool {
	if o == nil || IsNil(o.AuthenticateMethod) {
		var ret bool
		return ret
	}
	return *o.AuthenticateMethod
}

// GetAuthenticateMethodOk returns a tuple with the AuthenticateMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentprocessorGetAvailablePaymentMethodRequest) GetAuthenticateMethodOk() (*bool, bool) {
	if o == nil || IsNil(o.AuthenticateMethod) {
		return nil, false
	}
	return o.AuthenticateMethod, true
}

// HasAuthenticateMethod returns a boolean if a field has been set.
func (o *PaymentprocessorGetAvailablePaymentMethodRequest) HasAuthenticateMethod() bool {
	if o != nil && !IsNil(o.AuthenticateMethod) {
		return true
	}

	return false
}

// SetAuthenticateMethod gets a reference to the given bool and assigns it to the AuthenticateMethod field.
func (o *PaymentprocessorGetAvailablePaymentMethodRequest) SetAuthenticateMethod(v bool) {
	o.AuthenticateMethod = &v
}

func (o PaymentprocessorGetAvailablePaymentMethodRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentprocessorGetAvailablePaymentMethodRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["code"] = o.Code
	toSerialize["context"] = o.Context
	if !IsNil(o.AuthenticateMethod) {
		toSerialize["authenticateMethod"] = o.AuthenticateMethod
	}
	return toSerialize, nil
}

func (o *PaymentprocessorGetAvailablePaymentMethodRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"code",
		"context",
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

	varPaymentprocessorGetAvailablePaymentMethodRequest := _PaymentprocessorGetAvailablePaymentMethodRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymentprocessorGetAvailablePaymentMethodRequest)

	if err != nil {
		return err
	}

	*o = PaymentprocessorGetAvailablePaymentMethodRequest(varPaymentprocessorGetAvailablePaymentMethodRequest)

	return err
}

type NullablePaymentprocessorGetAvailablePaymentMethodRequest struct {
	value *PaymentprocessorGetAvailablePaymentMethodRequest
	isSet bool
}

func (v NullablePaymentprocessorGetAvailablePaymentMethodRequest) Get() *PaymentprocessorGetAvailablePaymentMethodRequest {
	return v.value
}

func (v *NullablePaymentprocessorGetAvailablePaymentMethodRequest) Set(val *PaymentprocessorGetAvailablePaymentMethodRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentprocessorGetAvailablePaymentMethodRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentprocessorGetAvailablePaymentMethodRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentprocessorGetAvailablePaymentMethodRequest(val *PaymentprocessorGetAvailablePaymentMethodRequest) *NullablePaymentprocessorGetAvailablePaymentMethodRequest {
	return &NullablePaymentprocessorGetAvailablePaymentMethodRequest{value: val, isSet: true}
}

func (v NullablePaymentprocessorGetAvailablePaymentMethodRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentprocessorGetAvailablePaymentMethodRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


