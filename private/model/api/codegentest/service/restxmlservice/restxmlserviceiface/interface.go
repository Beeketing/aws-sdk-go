// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package restxmlserviceiface provides an interface to enable mocking the REST XML Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package restxmlserviceiface

import (
	github.com/Beeketing/aws-sdk-go/aws"
	github.com/Beeketing/aws-sdk-go/aws/request"
	github.com/Beeketing/aws-sdk-go/private/model/api/codegentest/service/restxmlservice"
)

// RESTXMLServiceAPI provides an interface to enable mocking the
// restxmlservice.RESTXMLService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // REST XML Service.
//    func myFunc(svc restxmlserviceiface.RESTXMLServiceAPI) bool {
//        // Make svc.EmptyStream request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := restxmlservice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRESTXMLServiceClient struct {
//        restxmlserviceiface.RESTXMLServiceAPI
//    }
//    func (m *mockRESTXMLServiceClient) EmptyStream(input *restxmlservice.EmptyStreamInput) (*restxmlservice.EmptyStreamOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRESTXMLServiceClient{}
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
type RESTXMLServiceAPI interface {
	EmptyStream(*restxmlservice.EmptyStreamInput) (*restxmlservice.EmptyStreamOutput, error)
	EmptyStreamWithContext(aws.Context, *restxmlservice.EmptyStreamInput, ...request.Option) (*restxmlservice.EmptyStreamOutput, error)
	EmptyStreamRequest(*restxmlservice.EmptyStreamInput) (*request.Request, *restxmlservice.EmptyStreamOutput)

	GetEventStream(*restxmlservice.GetEventStreamInput) (*restxmlservice.GetEventStreamOutput, error)
	GetEventStreamWithContext(aws.Context, *restxmlservice.GetEventStreamInput, ...request.Option) (*restxmlservice.GetEventStreamOutput, error)
	GetEventStreamRequest(*restxmlservice.GetEventStreamInput) (*request.Request, *restxmlservice.GetEventStreamOutput)
}

var _ RESTXMLServiceAPI = (*restxmlservice.RESTXMLService)(nil)
