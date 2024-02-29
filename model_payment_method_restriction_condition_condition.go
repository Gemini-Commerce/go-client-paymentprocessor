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

// PaymentMethodRestrictionConditionCondition the model 'PaymentMethodRestrictionConditionCondition'
type PaymentMethodRestrictionConditionCondition string

// List of PaymentMethodRestrictionConditionCondition
const (
	IN PaymentMethodRestrictionConditionCondition = "IN"
	NOT_IN PaymentMethodRestrictionConditionCondition = "NOT_IN"
)

// All allowed values of PaymentMethodRestrictionConditionCondition enum
var AllowedPaymentMethodRestrictionConditionConditionEnumValues = []PaymentMethodRestrictionConditionCondition{
	"IN",
	"NOT_IN",
}

func (v *PaymentMethodRestrictionConditionCondition) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentMethodRestrictionConditionCondition(value)
	for _, existing := range AllowedPaymentMethodRestrictionConditionConditionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentMethodRestrictionConditionCondition", value)
}

// NewPaymentMethodRestrictionConditionConditionFromValue returns a pointer to a valid PaymentMethodRestrictionConditionCondition
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentMethodRestrictionConditionConditionFromValue(v string) (*PaymentMethodRestrictionConditionCondition, error) {
	ev := PaymentMethodRestrictionConditionCondition(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentMethodRestrictionConditionCondition: valid values are %v", v, AllowedPaymentMethodRestrictionConditionConditionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentMethodRestrictionConditionCondition) IsValid() bool {
	for _, existing := range AllowedPaymentMethodRestrictionConditionConditionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentMethodRestrictionConditionCondition value
func (v PaymentMethodRestrictionConditionCondition) Ptr() *PaymentMethodRestrictionConditionCondition {
	return &v
}

type NullablePaymentMethodRestrictionConditionCondition struct {
	value *PaymentMethodRestrictionConditionCondition
	isSet bool
}

func (v NullablePaymentMethodRestrictionConditionCondition) Get() *PaymentMethodRestrictionConditionCondition {
	return v.value
}

func (v *NullablePaymentMethodRestrictionConditionCondition) Set(val *PaymentMethodRestrictionConditionCondition) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodRestrictionConditionCondition) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodRestrictionConditionCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodRestrictionConditionCondition(val *PaymentMethodRestrictionConditionCondition) *NullablePaymentMethodRestrictionConditionCondition {
	return &NullablePaymentMethodRestrictionConditionCondition{value: val, isSet: true}
}

func (v NullablePaymentMethodRestrictionConditionCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodRestrictionConditionCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
