// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package globalacceleratoriface provides an interface to enable mocking the AWS Global Accelerator service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package globalacceleratoriface

import (
	"github.com/Beeketing/aws-sdk-go/aws"
	"github.com/Beeketing/aws-sdk-go/aws/request"
	"github.com/Beeketing/aws-sdk-go/service/globalaccelerator"
)

// GlobalAcceleratorAPI provides an interface to enable mocking the
// globalaccelerator.GlobalAccelerator service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Global Accelerator.
//    func myFunc(svc globalacceleratoriface.GlobalAcceleratorAPI) bool {
//        // Make svc.CreateAccelerator request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := globalaccelerator.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockGlobalAcceleratorClient struct {
//        globalacceleratoriface.GlobalAcceleratorAPI
//    }
//    func (m *mockGlobalAcceleratorClient) CreateAccelerator(input *globalaccelerator.CreateAcceleratorInput) (*globalaccelerator.CreateAcceleratorOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockGlobalAcceleratorClient{}
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
type GlobalAcceleratorAPI interface {
	CreateAccelerator(*globalaccelerator.CreateAcceleratorInput) (*globalaccelerator.CreateAcceleratorOutput, error)
	CreateAcceleratorWithContext(aws.Context, *globalaccelerator.CreateAcceleratorInput, ...request.Option) (*globalaccelerator.CreateAcceleratorOutput, error)
	CreateAcceleratorRequest(*globalaccelerator.CreateAcceleratorInput) (*request.Request, *globalaccelerator.CreateAcceleratorOutput)

	CreateEndpointGroup(*globalaccelerator.CreateEndpointGroupInput) (*globalaccelerator.CreateEndpointGroupOutput, error)
	CreateEndpointGroupWithContext(aws.Context, *globalaccelerator.CreateEndpointGroupInput, ...request.Option) (*globalaccelerator.CreateEndpointGroupOutput, error)
	CreateEndpointGroupRequest(*globalaccelerator.CreateEndpointGroupInput) (*request.Request, *globalaccelerator.CreateEndpointGroupOutput)

	CreateListener(*globalaccelerator.CreateListenerInput) (*globalaccelerator.CreateListenerOutput, error)
	CreateListenerWithContext(aws.Context, *globalaccelerator.CreateListenerInput, ...request.Option) (*globalaccelerator.CreateListenerOutput, error)
	CreateListenerRequest(*globalaccelerator.CreateListenerInput) (*request.Request, *globalaccelerator.CreateListenerOutput)

	DeleteAccelerator(*globalaccelerator.DeleteAcceleratorInput) (*globalaccelerator.DeleteAcceleratorOutput, error)
	DeleteAcceleratorWithContext(aws.Context, *globalaccelerator.DeleteAcceleratorInput, ...request.Option) (*globalaccelerator.DeleteAcceleratorOutput, error)
	DeleteAcceleratorRequest(*globalaccelerator.DeleteAcceleratorInput) (*request.Request, *globalaccelerator.DeleteAcceleratorOutput)

	DeleteEndpointGroup(*globalaccelerator.DeleteEndpointGroupInput) (*globalaccelerator.DeleteEndpointGroupOutput, error)
	DeleteEndpointGroupWithContext(aws.Context, *globalaccelerator.DeleteEndpointGroupInput, ...request.Option) (*globalaccelerator.DeleteEndpointGroupOutput, error)
	DeleteEndpointGroupRequest(*globalaccelerator.DeleteEndpointGroupInput) (*request.Request, *globalaccelerator.DeleteEndpointGroupOutput)

	DeleteListener(*globalaccelerator.DeleteListenerInput) (*globalaccelerator.DeleteListenerOutput, error)
	DeleteListenerWithContext(aws.Context, *globalaccelerator.DeleteListenerInput, ...request.Option) (*globalaccelerator.DeleteListenerOutput, error)
	DeleteListenerRequest(*globalaccelerator.DeleteListenerInput) (*request.Request, *globalaccelerator.DeleteListenerOutput)

	DescribeAccelerator(*globalaccelerator.DescribeAcceleratorInput) (*globalaccelerator.DescribeAcceleratorOutput, error)
	DescribeAcceleratorWithContext(aws.Context, *globalaccelerator.DescribeAcceleratorInput, ...request.Option) (*globalaccelerator.DescribeAcceleratorOutput, error)
	DescribeAcceleratorRequest(*globalaccelerator.DescribeAcceleratorInput) (*request.Request, *globalaccelerator.DescribeAcceleratorOutput)

	DescribeAcceleratorAttributes(*globalaccelerator.DescribeAcceleratorAttributesInput) (*globalaccelerator.DescribeAcceleratorAttributesOutput, error)
	DescribeAcceleratorAttributesWithContext(aws.Context, *globalaccelerator.DescribeAcceleratorAttributesInput, ...request.Option) (*globalaccelerator.DescribeAcceleratorAttributesOutput, error)
	DescribeAcceleratorAttributesRequest(*globalaccelerator.DescribeAcceleratorAttributesInput) (*request.Request, *globalaccelerator.DescribeAcceleratorAttributesOutput)

	DescribeEndpointGroup(*globalaccelerator.DescribeEndpointGroupInput) (*globalaccelerator.DescribeEndpointGroupOutput, error)
	DescribeEndpointGroupWithContext(aws.Context, *globalaccelerator.DescribeEndpointGroupInput, ...request.Option) (*globalaccelerator.DescribeEndpointGroupOutput, error)
	DescribeEndpointGroupRequest(*globalaccelerator.DescribeEndpointGroupInput) (*request.Request, *globalaccelerator.DescribeEndpointGroupOutput)

	DescribeListener(*globalaccelerator.DescribeListenerInput) (*globalaccelerator.DescribeListenerOutput, error)
	DescribeListenerWithContext(aws.Context, *globalaccelerator.DescribeListenerInput, ...request.Option) (*globalaccelerator.DescribeListenerOutput, error)
	DescribeListenerRequest(*globalaccelerator.DescribeListenerInput) (*request.Request, *globalaccelerator.DescribeListenerOutput)

	ListAccelerators(*globalaccelerator.ListAcceleratorsInput) (*globalaccelerator.ListAcceleratorsOutput, error)
	ListAcceleratorsWithContext(aws.Context, *globalaccelerator.ListAcceleratorsInput, ...request.Option) (*globalaccelerator.ListAcceleratorsOutput, error)
	ListAcceleratorsRequest(*globalaccelerator.ListAcceleratorsInput) (*request.Request, *globalaccelerator.ListAcceleratorsOutput)

	ListEndpointGroups(*globalaccelerator.ListEndpointGroupsInput) (*globalaccelerator.ListEndpointGroupsOutput, error)
	ListEndpointGroupsWithContext(aws.Context, *globalaccelerator.ListEndpointGroupsInput, ...request.Option) (*globalaccelerator.ListEndpointGroupsOutput, error)
	ListEndpointGroupsRequest(*globalaccelerator.ListEndpointGroupsInput) (*request.Request, *globalaccelerator.ListEndpointGroupsOutput)

	ListListeners(*globalaccelerator.ListListenersInput) (*globalaccelerator.ListListenersOutput, error)
	ListListenersWithContext(aws.Context, *globalaccelerator.ListListenersInput, ...request.Option) (*globalaccelerator.ListListenersOutput, error)
	ListListenersRequest(*globalaccelerator.ListListenersInput) (*request.Request, *globalaccelerator.ListListenersOutput)

	UpdateAccelerator(*globalaccelerator.UpdateAcceleratorInput) (*globalaccelerator.UpdateAcceleratorOutput, error)
	UpdateAcceleratorWithContext(aws.Context, *globalaccelerator.UpdateAcceleratorInput, ...request.Option) (*globalaccelerator.UpdateAcceleratorOutput, error)
	UpdateAcceleratorRequest(*globalaccelerator.UpdateAcceleratorInput) (*request.Request, *globalaccelerator.UpdateAcceleratorOutput)

	UpdateAcceleratorAttributes(*globalaccelerator.UpdateAcceleratorAttributesInput) (*globalaccelerator.UpdateAcceleratorAttributesOutput, error)
	UpdateAcceleratorAttributesWithContext(aws.Context, *globalaccelerator.UpdateAcceleratorAttributesInput, ...request.Option) (*globalaccelerator.UpdateAcceleratorAttributesOutput, error)
	UpdateAcceleratorAttributesRequest(*globalaccelerator.UpdateAcceleratorAttributesInput) (*request.Request, *globalaccelerator.UpdateAcceleratorAttributesOutput)

	UpdateEndpointGroup(*globalaccelerator.UpdateEndpointGroupInput) (*globalaccelerator.UpdateEndpointGroupOutput, error)
	UpdateEndpointGroupWithContext(aws.Context, *globalaccelerator.UpdateEndpointGroupInput, ...request.Option) (*globalaccelerator.UpdateEndpointGroupOutput, error)
	UpdateEndpointGroupRequest(*globalaccelerator.UpdateEndpointGroupInput) (*request.Request, *globalaccelerator.UpdateEndpointGroupOutput)

	UpdateListener(*globalaccelerator.UpdateListenerInput) (*globalaccelerator.UpdateListenerOutput, error)
	UpdateListenerWithContext(aws.Context, *globalaccelerator.UpdateListenerInput, ...request.Option) (*globalaccelerator.UpdateListenerOutput, error)
	UpdateListenerRequest(*globalaccelerator.UpdateListenerInput) (*request.Request, *globalaccelerator.UpdateListenerOutput)
}

var _ GlobalAcceleratorAPI = (*globalaccelerator.GlobalAccelerator)(nil)
