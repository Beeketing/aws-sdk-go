// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sagemakerruntimeiface provides an interface to enable mocking the Amazon SageMaker Runtime service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package sagemakerruntimeiface

import (
	github.com/Beeketing/aws-sdk-go/aws"
	github.com/Beeketing/aws-sdk-go/aws/request"
	github.com/Beeketing/aws-sdk-go/service/sagemakerruntime"
)

// SageMakerRuntimeAPI provides an interface to enable mocking the
// sagemakerruntime.SageMakerRuntime service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon SageMaker Runtime.
//    func myFunc(svc sagemakerruntimeiface.SageMakerRuntimeAPI) bool {
//        // Make svc.InvokeEndpoint request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := sagemakerruntime.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSageMakerRuntimeClient struct {
//        sagemakerruntimeiface.SageMakerRuntimeAPI
//    }
//    func (m *mockSageMakerRuntimeClient) InvokeEndpoint(input *sagemakerruntime.InvokeEndpointInput) (*sagemakerruntime.InvokeEndpointOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSageMakerRuntimeClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type SageMakerRuntimeAPI interface {
	InvokeEndpoint(*sagemakerruntime.InvokeEndpointInput) (*sagemakerruntime.InvokeEndpointOutput, error)
	InvokeEndpointWithContext(aws.Context, *sagemakerruntime.InvokeEndpointInput, ...request.Option) (*sagemakerruntime.InvokeEndpointOutput, error)
	InvokeEndpointRequest(*sagemakerruntime.InvokeEndpointInput) (*request.Request, *sagemakerruntime.InvokeEndpointOutput)
}

var _ SageMakerRuntimeAPI = (*sagemakerruntime.SageMakerRuntime)(nil)
