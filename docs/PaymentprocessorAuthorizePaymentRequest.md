# # PaymentprocessorAuthorizePaymentRequest


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId**| **string** |   |
**Context**| [**PaymentprocessorPaymentContext**](PaymentprocessorPaymentContext.md) |   | [optional]
**Code**| **string** | payment type stripe, paypal..  |
**Amount**| [**PaymentprocessorMoney**](PaymentprocessorMoney.md) |   | [optional]
**Currency**| [**PaymentprocessorCurrency**](PaymentprocessorCurrency.md) |  for more information please, see Model/PaymentprocessorCurrency.php  | [optional] [default to PAYMENTPROCESSORCURRENCY_XXX]
**AdditionalInfo**| **string** |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

