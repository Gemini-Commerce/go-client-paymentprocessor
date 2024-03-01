# # PaymentprocessorUpdatePaymentMethodRequest


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId**| **string** |   |
**Code**| **string** |   |
**Title**| [**PaymentprocessorLocalizedText**](PaymentprocessorLocalizedText.md) |   | [optional]
**Label**| [**PaymentprocessorLocalizedText**](PaymentprocessorLocalizedText.md) |   | [optional]
**Enabled**| **bool** |   | [optional]
**Amount**| [**PaymentprocessorMoney**](PaymentprocessorMoney.md) |   | [optional]
**Currency**| [**PaymentprocessorCurrency**](PaymentprocessorCurrency.md) |  for more information please, see Model/PaymentprocessorCurrency.php  | [optional] [default to PAYMENTPROCESSORCURRENCY_XXX]
**Configuration**| **map[string]interface{}** |   | [optional]
**IsUpfront**| **bool** |   | [optional]
**Description**| [**PaymentprocessorLocalizedText**](PaymentprocessorLocalizedText.md) |   | [optional]
**Restrictions**| [**[]PaymentprocessorPaymentMethodRestriction**](PaymentprocessorPaymentMethodRestriction.md) |   | [optional]
**FieldMask**| **string** |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

