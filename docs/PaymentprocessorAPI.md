# GeminiCommerce\Paymentprocessor\PaymentprocessorAPI

All URIs are relative to *https://payment-processor.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizePayment**](PaymentprocessorAPI.md#AuthorizePayment) | **Post** /paymentprocessor.Paymentprocessor/AuthorizePayment | Authorize Payment
[**CreatePaymentMethod**](PaymentprocessorAPI.md#CreatePaymentMethod) | **Post** /paymentprocessor.Paymentprocessor/CreatePaymentMethod | Create Payment Method
[**FinalizePayment**](PaymentprocessorAPI.md#FinalizePayment) | **Post** /paymentprocessor.Paymentprocessor/FinalizePayment | Finalize Payment
[**GetAvailablePaymentMethod**](PaymentprocessorAPI.md#GetAvailablePaymentMethod) | **Post** /paymentprocessor.Paymentprocessor/GetAvailablePaymentMethod | Get Available Payment Method
[**GetPaymentMethod**](PaymentprocessorAPI.md#GetPaymentMethod) | **Post** /paymentprocessor.Paymentprocessor/GetPaymentMethod | Get Payment Method
[**GetPaymentMethodConfiguration**](PaymentprocessorAPI.md#GetPaymentMethodConfiguration) | **Post** /paymentprocessor.Paymentprocessor/GetPaymentMethodConfiguration | Get Payment Method Configuration
[**InitPayment**](PaymentprocessorAPI.md#InitPayment) | **Post** /paymentprocessor.Paymentprocessor/InitPayment | Init Payment
[**ListAvailablePaymentMethods**](PaymentprocessorAPI.md#ListAvailablePaymentMethods) | **Post** /paymentprocessor.Paymentprocessor/ListAvailablePaymentMethods | List Available Payment Methods
[**ListPaymentMethods**](PaymentprocessorAPI.md#ListPaymentMethods) | **Post** /paymentprocessor.Paymentprocessor/ListPaymentMethods | List Payment Methods
[**PerformPayment**](PaymentprocessorAPI.md#PerformPayment) | **Post** /paymentprocessor.Paymentprocessor/PerformPayment | Perform Payment
[**PerformRefund**](PaymentprocessorAPI.md#PerformRefund) | **Post** /paymentprocessor.Paymentprocessor/PerformRefund | Perform Refund
[**UpdatePayment**](PaymentprocessorAPI.md#UpdatePayment) | **Post** /paymentprocessor.Paymentprocessor/UpdatePayment | Update Payment
[**UpdatePaymentMethod**](PaymentprocessorAPI.md#UpdatePaymentMethod) | **Post** /paymentprocessor.Paymentprocessor/UpdatePaymentMethod | Update Payment Method
[**VoidPayment**](PaymentprocessorAPI.md#VoidPayment) | **Post** /paymentprocessor.Paymentprocessor/VoidPayment | Void Payment



## AuthorizePayment

> PaymentprocessorAuthorizePaymentResponse AuthorizePayment(ctx).Body(body).Execute()

Authorize Payment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-paymentprocessor"
)

func main() {
	body := *openapiclient.NewPaymentprocessorAuthorizePaymentRequest("TenantId_example", "Code_example") // PaymentprocessorAuthorizePaymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentprocessorAPI.AuthorizePayment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentprocessorAPI.AuthorizePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorizePayment`: PaymentprocessorAuthorizePaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentprocessorAPI.AuthorizePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaymentprocessorAuthorizePaymentRequest**](PaymentprocessorAuthorizePaymentRequest.md) |  | 

### Return type

[**PaymentprocessorAuthorizePaymentResponse**](PaymentprocessorAuthorizePaymentResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePaymentMethod

> PaymentprocessorPaymentMethod CreatePaymentMethod(ctx).Body(body).Execute()

Create Payment Method

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-paymentprocessor"
)

func main() {
	body := *openapiclient.NewPaymentprocessorCreatePaymentMethodRequest("TenantId_example", "Code_example", *openapiclient.NewPaymentprocessorLocalizedText()) // PaymentprocessorCreatePaymentMethodRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentprocessorAPI.CreatePaymentMethod(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentprocessorAPI.CreatePaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaymentMethod`: PaymentprocessorPaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `PaymentprocessorAPI.CreatePaymentMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaymentprocessorCreatePaymentMethodRequest**](PaymentprocessorCreatePaymentMethodRequest.md) |  | 

### Return type

[**PaymentprocessorPaymentMethod**](PaymentprocessorPaymentMethod.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinalizePayment

> PaymentprocessorFinalizePaymentResponse FinalizePayment(ctx).Body(body).Execute()

Finalize Payment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-paymentprocessor"
)

func main() {
	body := *openapiclient.NewPaymentprocessorFinalizePaymentRequest("TenantId_example", "PaymentId_example", "Code_example") // PaymentprocessorFinalizePaymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentprocessorAPI.FinalizePayment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentprocessorAPI.FinalizePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinalizePayment`: PaymentprocessorFinalizePaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentprocessorAPI.FinalizePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFinalizePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaymentprocessorFinalizePaymentRequest**](PaymentprocessorFinalizePaymentRequest.md) |  | 

### Return type

[**PaymentprocessorFinalizePaymentResponse**](PaymentprocessorFinalizePaymentResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailablePaymentMethod

> PaymentprocessorPaymentMethod GetAvailablePaymentMethod(ctx).Body(body).Execute()

Get Available Payment Method

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-paymentprocessor"
)

func main() {
	body := *openapiclient.NewPaymentprocessorGetAvailablePaymentMethodRequest("TenantId_example", "Code_example", *openapiclient.NewPaymentprocessorAvailabilityContext()) // PaymentprocessorGetAvailablePaymentMethodRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentprocessorAPI.GetAvailablePaymentMethod(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentprocessorAPI.GetAvailablePaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailablePaymentMethod`: PaymentprocessorPaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `PaymentprocessorAPI.GetAvailablePaymentMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailablePaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaymentprocessorGetAvailablePaymentMethodRequest**](PaymentprocessorGetAvailablePaymentMethodRequest.md) |  | 

### Return type

[**PaymentprocessorPaymentMethod**](PaymentprocessorPaymentMethod.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentMethod

> PaymentprocessorPaymentMethod GetPaymentMethod(ctx).Body(body).Execute()

Get Payment Method

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-paymentprocessor"
)

func main() {
	body := *openapiclient.NewPaymentprocessorGetPaymentMethodRequest("TenantId_example", "Code_example") // PaymentprocessorGetPaymentMethodRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentprocessorAPI.GetPaymentMethod(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentprocessorAPI.GetPaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentMethod`: PaymentprocessorPaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `PaymentprocessorAPI.GetPaymentMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaymentprocessorGetPaymentMethodRequest**](PaymentprocessorGetPaymentMethodRequest.md) |  | 

### Return type

[**PaymentprocessorPaymentMethod**](PaymentprocessorPaymentMethod.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentMethodConfiguration

> PaymentprocessorGetPaymentMethodConfigurationResponse GetPaymentMethodConfiguration(ctx).Body(body).Execute()

Get Payment Method Configuration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-paymentprocessor"
)

func main() {
	body := *openapiclient.NewPaymentprocessorGetPaymentMethodConfigurationRequest("TenantId_example", "Code_example") // PaymentprocessorGetPaymentMethodConfigurationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentprocessorAPI.GetPaymentMethodConfiguration(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentprocessorAPI.GetPaymentMethodConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentMethodConfiguration`: PaymentprocessorGetPaymentMethodConfigurationResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentprocessorAPI.GetPaymentMethodConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentMethodConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaymentprocessorGetPaymentMethodConfigurationRequest**](PaymentprocessorGetPaymentMethodConfigurationRequest.md) |  | 

### Return type

[**PaymentprocessorGetPaymentMethodConfigurationResponse**](PaymentprocessorGetPaymentMethodConfigurationResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitPayment

> PaymentprocessorInitPaymentResponse InitPayment(ctx).Body(body).Execute()

Init Payment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-paymentprocessor"
)

func main() {
	body := *openapiclient.NewPaymentprocessorInitPaymentRequest("TenantId_example", "Code_example") // PaymentprocessorInitPaymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentprocessorAPI.InitPayment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentprocessorAPI.InitPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitPayment`: PaymentprocessorInitPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentprocessorAPI.InitPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaymentprocessorInitPaymentRequest**](PaymentprocessorInitPaymentRequest.md) |  | 

### Return type

[**PaymentprocessorInitPaymentResponse**](PaymentprocessorInitPaymentResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailablePaymentMethods

> PaymentprocessorListAvailablePaymentMethodsResponse ListAvailablePaymentMethods(ctx).Body(body).Execute()

List Available Payment Methods

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-paymentprocessor"
)

func main() {
	body := *openapiclient.NewPaymentprocessorListAvailablePaymentMethodsRequest("TenantId_example", *openapiclient.NewPaymentprocessorAvailabilityContext()) // PaymentprocessorListAvailablePaymentMethodsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentprocessorAPI.ListAvailablePaymentMethods(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentprocessorAPI.ListAvailablePaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAvailablePaymentMethods`: PaymentprocessorListAvailablePaymentMethodsResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentprocessorAPI.ListAvailablePaymentMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAvailablePaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaymentprocessorListAvailablePaymentMethodsRequest**](PaymentprocessorListAvailablePaymentMethodsRequest.md) |  | 

### Return type

[**PaymentprocessorListAvailablePaymentMethodsResponse**](PaymentprocessorListAvailablePaymentMethodsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentMethods

> PaymentprocessorListPaymentMethodsResponse ListPaymentMethods(ctx).Body(body).Execute()

List Payment Methods

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-paymentprocessor"
)

func main() {
	body := *openapiclient.NewPaymentprocessorListPaymentMethodsRequest("TenantId_example") // PaymentprocessorListPaymentMethodsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentprocessorAPI.ListPaymentMethods(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentprocessorAPI.ListPaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPaymentMethods`: PaymentprocessorListPaymentMethodsResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentprocessorAPI.ListPaymentMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaymentprocessorListPaymentMethodsRequest**](PaymentprocessorListPaymentMethodsRequest.md) |  | 

### Return type

[**PaymentprocessorListPaymentMethodsResponse**](PaymentprocessorListPaymentMethodsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformPayment

> PaymentprocessorPerformPaymentResponse PerformPayment(ctx).Body(body).Execute()

Perform Payment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-paymentprocessor"
)

func main() {
	body := *openapiclient.NewPaymentprocessorPerformPaymentRequest("TenantId_example", "PaymentId_example", "Code_example") // PaymentprocessorPerformPaymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentprocessorAPI.PerformPayment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentprocessorAPI.PerformPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PerformPayment`: PaymentprocessorPerformPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentprocessorAPI.PerformPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPerformPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaymentprocessorPerformPaymentRequest**](PaymentprocessorPerformPaymentRequest.md) |  | 

### Return type

[**PaymentprocessorPerformPaymentResponse**](PaymentprocessorPerformPaymentResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformRefund

> map[string]interface{} PerformRefund(ctx).Body(body).Execute()

Perform Refund

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-paymentprocessor"
)

func main() {
	body := *openapiclient.NewPaymentprocessorPerformRefundRequest("TenantId_example", "RefundId_example", *openapiclient.NewPaymentprocessorPayment("Id_example", "Code_example")) // PaymentprocessorPerformRefundRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentprocessorAPI.PerformRefund(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentprocessorAPI.PerformRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PerformRefund`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaymentprocessorAPI.PerformRefund`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPerformRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaymentprocessorPerformRefundRequest**](PaymentprocessorPerformRefundRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePayment

> map[string]interface{} UpdatePayment(ctx).Body(body).Execute()

Update Payment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-paymentprocessor"
)

func main() {
	body := *openapiclient.NewPaymentprocessorUpdatePaymentRequest("TenantId_example", "PaymentId_example", "Code_example") // PaymentprocessorUpdatePaymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentprocessorAPI.UpdatePayment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentprocessorAPI.UpdatePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePayment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaymentprocessorAPI.UpdatePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaymentprocessorUpdatePaymentRequest**](PaymentprocessorUpdatePaymentRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePaymentMethod

> PaymentprocessorPaymentMethod UpdatePaymentMethod(ctx).Body(body).Execute()

Update Payment Method

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-paymentprocessor"
)

func main() {
	body := *openapiclient.NewPaymentprocessorUpdatePaymentMethodRequest("TenantId_example", "Code_example") // PaymentprocessorUpdatePaymentMethodRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentprocessorAPI.UpdatePaymentMethod(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentprocessorAPI.UpdatePaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePaymentMethod`: PaymentprocessorPaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `PaymentprocessorAPI.UpdatePaymentMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaymentprocessorUpdatePaymentMethodRequest**](PaymentprocessorUpdatePaymentMethodRequest.md) |  | 

### Return type

[**PaymentprocessorPaymentMethod**](PaymentprocessorPaymentMethod.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidPayment

> PaymentprocessorVoidPaymentResponse VoidPayment(ctx).Body(body).Execute()

Void Payment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-paymentprocessor"
)

func main() {
	body := *openapiclient.NewPaymentprocessorVoidPaymentRequest("TenantId_example", "Code_example") // PaymentprocessorVoidPaymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentprocessorAPI.VoidPayment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentprocessorAPI.VoidPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoidPayment`: PaymentprocessorVoidPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentprocessorAPI.VoidPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVoidPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaymentprocessorVoidPaymentRequest**](PaymentprocessorVoidPaymentRequest.md) |  | 

### Return type

[**PaymentprocessorVoidPaymentResponse**](PaymentprocessorVoidPaymentResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

