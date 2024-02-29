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

// PaymentprocessorCurrency Stands for Albania, not all ;-)
type PaymentprocessorCurrency string

// List of paymentprocessorCurrency
const (
	XXX PaymentprocessorCurrency = "XXX"
	ALL PaymentprocessorCurrency = "ALL"
	DZD PaymentprocessorCurrency = "DZD"
	ARS PaymentprocessorCurrency = "ARS"
	AUD PaymentprocessorCurrency = "AUD"
	BSD PaymentprocessorCurrency = "BSD"
	BHD PaymentprocessorCurrency = "BHD"
	BDT PaymentprocessorCurrency = "BDT"
	AMD PaymentprocessorCurrency = "AMD"
	BBD PaymentprocessorCurrency = "BBD"
	BMD PaymentprocessorCurrency = "BMD"
	BTN PaymentprocessorCurrency = "BTN"
	BOB PaymentprocessorCurrency = "BOB"
	BWP PaymentprocessorCurrency = "BWP"
	BZD PaymentprocessorCurrency = "BZD"
	SBD PaymentprocessorCurrency = "SBD"
	BND PaymentprocessorCurrency = "BND"
	MMK PaymentprocessorCurrency = "MMK"
	BIF PaymentprocessorCurrency = "BIF"
	KHR PaymentprocessorCurrency = "KHR"
	CAD PaymentprocessorCurrency = "CAD"
	CVE PaymentprocessorCurrency = "CVE"
	KYD PaymentprocessorCurrency = "KYD"
	LKR PaymentprocessorCurrency = "LKR"
	CLP PaymentprocessorCurrency = "CLP"
	CNY PaymentprocessorCurrency = "CNY"
	COP PaymentprocessorCurrency = "COP"
	KMF PaymentprocessorCurrency = "KMF"
	CRC PaymentprocessorCurrency = "CRC"
	HRK PaymentprocessorCurrency = "HRK"
	CUP PaymentprocessorCurrency = "CUP"
	CZK PaymentprocessorCurrency = "CZK"
	DKK PaymentprocessorCurrency = "DKK"
	DOP PaymentprocessorCurrency = "DOP"
	SVC PaymentprocessorCurrency = "SVC"
	ETB PaymentprocessorCurrency = "ETB"
	ERN PaymentprocessorCurrency = "ERN"
	FKP PaymentprocessorCurrency = "FKP"
	FJD PaymentprocessorCurrency = "FJD"
	DJF PaymentprocessorCurrency = "DJF"
	GMD PaymentprocessorCurrency = "GMD"
	GIP PaymentprocessorCurrency = "GIP"
	GTQ PaymentprocessorCurrency = "GTQ"
	GNF PaymentprocessorCurrency = "GNF"
	GYD PaymentprocessorCurrency = "GYD"
	HTG PaymentprocessorCurrency = "HTG"
	HNL PaymentprocessorCurrency = "HNL"
	HKD PaymentprocessorCurrency = "HKD"
	HUF PaymentprocessorCurrency = "HUF"
	ISK PaymentprocessorCurrency = "ISK"
	INR PaymentprocessorCurrency = "INR"
	IDR PaymentprocessorCurrency = "IDR"
	IRR PaymentprocessorCurrency = "IRR"
	IQD PaymentprocessorCurrency = "IQD"
	ILS PaymentprocessorCurrency = "ILS"
	JMD PaymentprocessorCurrency = "JMD"
	JPY PaymentprocessorCurrency = "JPY"
	KZT PaymentprocessorCurrency = "KZT"
	JOD PaymentprocessorCurrency = "JOD"
	KES PaymentprocessorCurrency = "KES"
	KPW PaymentprocessorCurrency = "KPW"
	KRW PaymentprocessorCurrency = "KRW"
	KWD PaymentprocessorCurrency = "KWD"
	KGS PaymentprocessorCurrency = "KGS"
	LAK PaymentprocessorCurrency = "LAK"
	LBP PaymentprocessorCurrency = "LBP"
	LSL PaymentprocessorCurrency = "LSL"
	LRD PaymentprocessorCurrency = "LRD"
	LYD PaymentprocessorCurrency = "LYD"
	LTL PaymentprocessorCurrency = "LTL"
	MOP PaymentprocessorCurrency = "MOP"
	MWK PaymentprocessorCurrency = "MWK"
	MYR PaymentprocessorCurrency = "MYR"
	MVR PaymentprocessorCurrency = "MVR"
	MRO PaymentprocessorCurrency = "MRO"
	MUR PaymentprocessorCurrency = "MUR"
	MXN PaymentprocessorCurrency = "MXN"
	MNT PaymentprocessorCurrency = "MNT"
	MDL PaymentprocessorCurrency = "MDL"
	MAD PaymentprocessorCurrency = "MAD"
	OMR PaymentprocessorCurrency = "OMR"
	NAD PaymentprocessorCurrency = "NAD"
	NPR PaymentprocessorCurrency = "NPR"
	ANG PaymentprocessorCurrency = "ANG"
	AWG PaymentprocessorCurrency = "AWG"
	VUV PaymentprocessorCurrency = "VUV"
	NZD PaymentprocessorCurrency = "NZD"
	NIO PaymentprocessorCurrency = "NIO"
	NGN PaymentprocessorCurrency = "NGN"
	NOK PaymentprocessorCurrency = "NOK"
	PKR PaymentprocessorCurrency = "PKR"
	PAB PaymentprocessorCurrency = "PAB"
	PGK PaymentprocessorCurrency = "PGK"
	PYG PaymentprocessorCurrency = "PYG"
	PEN PaymentprocessorCurrency = "PEN"
	PHP PaymentprocessorCurrency = "PHP"
	QAR PaymentprocessorCurrency = "QAR"
	RUB PaymentprocessorCurrency = "RUB"
	RWF PaymentprocessorCurrency = "RWF"
	SHP PaymentprocessorCurrency = "SHP"
	STD PaymentprocessorCurrency = "STD"
	SAR PaymentprocessorCurrency = "SAR"
	SCR PaymentprocessorCurrency = "SCR"
	SLL PaymentprocessorCurrency = "SLL"
	SGD PaymentprocessorCurrency = "SGD"
	VND PaymentprocessorCurrency = "VND"
	SOS PaymentprocessorCurrency = "SOS"
	ZAR PaymentprocessorCurrency = "ZAR"
	SSP PaymentprocessorCurrency = "SSP"
	SZL PaymentprocessorCurrency = "SZL"
	SEK PaymentprocessorCurrency = "SEK"
	CHF PaymentprocessorCurrency = "CHF"
	SYP PaymentprocessorCurrency = "SYP"
	THB PaymentprocessorCurrency = "THB"
	TOP PaymentprocessorCurrency = "TOP"
	TTD PaymentprocessorCurrency = "TTD"
	AED PaymentprocessorCurrency = "AED"
	TND PaymentprocessorCurrency = "TND"
	UGX PaymentprocessorCurrency = "UGX"
	MKD PaymentprocessorCurrency = "MKD"
	EGP PaymentprocessorCurrency = "EGP"
	GBP PaymentprocessorCurrency = "GBP"
	TZS PaymentprocessorCurrency = "TZS"
	USD PaymentprocessorCurrency = "USD"
	UYU PaymentprocessorCurrency = "UYU"
	UZS PaymentprocessorCurrency = "UZS"
	WST PaymentprocessorCurrency = "WST"
	YER PaymentprocessorCurrency = "YER"
	TWD PaymentprocessorCurrency = "TWD"
	CUC PaymentprocessorCurrency = "CUC"
	ZWL PaymentprocessorCurrency = "ZWL"
	TMT PaymentprocessorCurrency = "TMT"
	GHS PaymentprocessorCurrency = "GHS"
	VEF PaymentprocessorCurrency = "VEF"
	SDG PaymentprocessorCurrency = "SDG"
	UYI PaymentprocessorCurrency = "UYI"
	RSD PaymentprocessorCurrency = "RSD"
	MZN PaymentprocessorCurrency = "MZN"
	AZN PaymentprocessorCurrency = "AZN"
	RON PaymentprocessorCurrency = "RON"
	CHE PaymentprocessorCurrency = "CHE"
	CHW PaymentprocessorCurrency = "CHW"
	TRY PaymentprocessorCurrency = "TRY"
	XAF PaymentprocessorCurrency = "XAF"
	XCD PaymentprocessorCurrency = "XCD"
	XOF PaymentprocessorCurrency = "XOF"
	XPF PaymentprocessorCurrency = "XPF"
	XBA PaymentprocessorCurrency = "XBA"
	XBB PaymentprocessorCurrency = "XBB"
	XBC PaymentprocessorCurrency = "XBC"
	XBD PaymentprocessorCurrency = "XBD"
	XAU PaymentprocessorCurrency = "XAU"
	XDR PaymentprocessorCurrency = "XDR"
	XAG PaymentprocessorCurrency = "XAG"
	XPT PaymentprocessorCurrency = "XPT"
	XPD PaymentprocessorCurrency = "XPD"
	XUA PaymentprocessorCurrency = "XUA"
	ZMW PaymentprocessorCurrency = "ZMW"
	SRD PaymentprocessorCurrency = "SRD"
	MGA PaymentprocessorCurrency = "MGA"
	COU PaymentprocessorCurrency = "COU"
	AFN PaymentprocessorCurrency = "AFN"
	TJS PaymentprocessorCurrency = "TJS"
	AOA PaymentprocessorCurrency = "AOA"
	BYR PaymentprocessorCurrency = "BYR"
	BGN PaymentprocessorCurrency = "BGN"
	CDF PaymentprocessorCurrency = "CDF"
	BAM PaymentprocessorCurrency = "BAM"
	EUR PaymentprocessorCurrency = "EUR"
	MXV PaymentprocessorCurrency = "MXV"
	UAH PaymentprocessorCurrency = "UAH"
	GEL PaymentprocessorCurrency = "GEL"
	BOV PaymentprocessorCurrency = "BOV"
	PLN PaymentprocessorCurrency = "PLN"
	BRL PaymentprocessorCurrency = "BRL"
	CLF PaymentprocessorCurrency = "CLF"
	XSU PaymentprocessorCurrency = "XSU"
	USN PaymentprocessorCurrency = "USN"
)

// All allowed values of PaymentprocessorCurrency enum
var AllowedPaymentprocessorCurrencyEnumValues = []PaymentprocessorCurrency{
	"XXX",
	"ALL",
	"DZD",
	"ARS",
	"AUD",
	"BSD",
	"BHD",
	"BDT",
	"AMD",
	"BBD",
	"BMD",
	"BTN",
	"BOB",
	"BWP",
	"BZD",
	"SBD",
	"BND",
	"MMK",
	"BIF",
	"KHR",
	"CAD",
	"CVE",
	"KYD",
	"LKR",
	"CLP",
	"CNY",
	"COP",
	"KMF",
	"CRC",
	"HRK",
	"CUP",
	"CZK",
	"DKK",
	"DOP",
	"SVC",
	"ETB",
	"ERN",
	"FKP",
	"FJD",
	"DJF",
	"GMD",
	"GIP",
	"GTQ",
	"GNF",
	"GYD",
	"HTG",
	"HNL",
	"HKD",
	"HUF",
	"ISK",
	"INR",
	"IDR",
	"IRR",
	"IQD",
	"ILS",
	"JMD",
	"JPY",
	"KZT",
	"JOD",
	"KES",
	"KPW",
	"KRW",
	"KWD",
	"KGS",
	"LAK",
	"LBP",
	"LSL",
	"LRD",
	"LYD",
	"LTL",
	"MOP",
	"MWK",
	"MYR",
	"MVR",
	"MRO",
	"MUR",
	"MXN",
	"MNT",
	"MDL",
	"MAD",
	"OMR",
	"NAD",
	"NPR",
	"ANG",
	"AWG",
	"VUV",
	"NZD",
	"NIO",
	"NGN",
	"NOK",
	"PKR",
	"PAB",
	"PGK",
	"PYG",
	"PEN",
	"PHP",
	"QAR",
	"RUB",
	"RWF",
	"SHP",
	"STD",
	"SAR",
	"SCR",
	"SLL",
	"SGD",
	"VND",
	"SOS",
	"ZAR",
	"SSP",
	"SZL",
	"SEK",
	"CHF",
	"SYP",
	"THB",
	"TOP",
	"TTD",
	"AED",
	"TND",
	"UGX",
	"MKD",
	"EGP",
	"GBP",
	"TZS",
	"USD",
	"UYU",
	"UZS",
	"WST",
	"YER",
	"TWD",
	"CUC",
	"ZWL",
	"TMT",
	"GHS",
	"VEF",
	"SDG",
	"UYI",
	"RSD",
	"MZN",
	"AZN",
	"RON",
	"CHE",
	"CHW",
	"TRY",
	"XAF",
	"XCD",
	"XOF",
	"XPF",
	"XBA",
	"XBB",
	"XBC",
	"XBD",
	"XAU",
	"XDR",
	"XAG",
	"XPT",
	"XPD",
	"XUA",
	"ZMW",
	"SRD",
	"MGA",
	"COU",
	"AFN",
	"TJS",
	"AOA",
	"BYR",
	"BGN",
	"CDF",
	"BAM",
	"EUR",
	"MXV",
	"UAH",
	"GEL",
	"BOV",
	"PLN",
	"BRL",
	"CLF",
	"XSU",
	"USN",
}

func (v *PaymentprocessorCurrency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentprocessorCurrency(value)
	for _, existing := range AllowedPaymentprocessorCurrencyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentprocessorCurrency", value)
}

// NewPaymentprocessorCurrencyFromValue returns a pointer to a valid PaymentprocessorCurrency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentprocessorCurrencyFromValue(v string) (*PaymentprocessorCurrency, error) {
	ev := PaymentprocessorCurrency(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentprocessorCurrency: valid values are %v", v, AllowedPaymentprocessorCurrencyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentprocessorCurrency) IsValid() bool {
	for _, existing := range AllowedPaymentprocessorCurrencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to paymentprocessorCurrency value
func (v PaymentprocessorCurrency) Ptr() *PaymentprocessorCurrency {
	return &v
}

type NullablePaymentprocessorCurrency struct {
	value *PaymentprocessorCurrency
	isSet bool
}

func (v NullablePaymentprocessorCurrency) Get() *PaymentprocessorCurrency {
	return v.value
}

func (v *NullablePaymentprocessorCurrency) Set(val *PaymentprocessorCurrency) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentprocessorCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentprocessorCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentprocessorCurrency(val *PaymentprocessorCurrency) *NullablePaymentprocessorCurrency {
	return &NullablePaymentprocessorCurrency{value: val, isSet: true}
}

func (v NullablePaymentprocessorCurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentprocessorCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

