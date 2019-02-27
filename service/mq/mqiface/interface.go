// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mqiface provides an interface to enable mocking the AmazonMQ service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mqiface

import (
	github.com/Beeketing/aws-sdk-go/aws"
	github.com/Beeketing/aws-sdk-go/aws/request"
	github.com/Beeketing/aws-sdk-go/service/mq"
)

// MQAPI provides an interface to enable mocking the
// mq.MQ service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AmazonMQ.
//    func myFunc(svc mqiface.MQAPI) bool {
//        // Make svc.CreateBroker request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := mq.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMQClient struct {
//        mqiface.MQAPI
//    }
//    func (m *mockMQClient) CreateBroker(input *mq.CreateBrokerRequest) (*mq.CreateBrokerResponse, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMQClient{}
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
type MQAPI interface {
	CreateBroker(*mq.CreateBrokerRequest) (*mq.CreateBrokerResponse, error)
	CreateBrokerWithContext(aws.Context, *mq.CreateBrokerRequest, ...request.Option) (*mq.CreateBrokerResponse, error)
	CreateBrokerRequest(*mq.CreateBrokerRequest) (*request.Request, *mq.CreateBrokerResponse)

	CreateConfiguration(*mq.CreateConfigurationRequest) (*mq.CreateConfigurationResponse, error)
	CreateConfigurationWithContext(aws.Context, *mq.CreateConfigurationRequest, ...request.Option) (*mq.CreateConfigurationResponse, error)
	CreateConfigurationRequest(*mq.CreateConfigurationRequest) (*request.Request, *mq.CreateConfigurationResponse)

	CreateTags(*mq.CreateTagsInput) (*mq.CreateTagsOutput, error)
	CreateTagsWithContext(aws.Context, *mq.CreateTagsInput, ...request.Option) (*mq.CreateTagsOutput, error)
	CreateTagsRequest(*mq.CreateTagsInput) (*request.Request, *mq.CreateTagsOutput)

	CreateUser(*mq.CreateUserRequest) (*mq.CreateUserOutput, error)
	CreateUserWithContext(aws.Context, *mq.CreateUserRequest, ...request.Option) (*mq.CreateUserOutput, error)
	CreateUserRequest(*mq.CreateUserRequest) (*request.Request, *mq.CreateUserOutput)

	DeleteBroker(*mq.DeleteBrokerInput) (*mq.DeleteBrokerResponse, error)
	DeleteBrokerWithContext(aws.Context, *mq.DeleteBrokerInput, ...request.Option) (*mq.DeleteBrokerResponse, error)
	DeleteBrokerRequest(*mq.DeleteBrokerInput) (*request.Request, *mq.DeleteBrokerResponse)

	DeleteTags(*mq.DeleteTagsInput) (*mq.DeleteTagsOutput, error)
	DeleteTagsWithContext(aws.Context, *mq.DeleteTagsInput, ...request.Option) (*mq.DeleteTagsOutput, error)
	DeleteTagsRequest(*mq.DeleteTagsInput) (*request.Request, *mq.DeleteTagsOutput)

	DeleteUser(*mq.DeleteUserInput) (*mq.DeleteUserOutput, error)
	DeleteUserWithContext(aws.Context, *mq.DeleteUserInput, ...request.Option) (*mq.DeleteUserOutput, error)
	DeleteUserRequest(*mq.DeleteUserInput) (*request.Request, *mq.DeleteUserOutput)

	DescribeBroker(*mq.DescribeBrokerInput) (*mq.DescribeBrokerResponse, error)
	DescribeBrokerWithContext(aws.Context, *mq.DescribeBrokerInput, ...request.Option) (*mq.DescribeBrokerResponse, error)
	DescribeBrokerRequest(*mq.DescribeBrokerInput) (*request.Request, *mq.DescribeBrokerResponse)

	DescribeConfiguration(*mq.DescribeConfigurationInput) (*mq.DescribeConfigurationOutput, error)
	DescribeConfigurationWithContext(aws.Context, *mq.DescribeConfigurationInput, ...request.Option) (*mq.DescribeConfigurationOutput, error)
	DescribeConfigurationRequest(*mq.DescribeConfigurationInput) (*request.Request, *mq.DescribeConfigurationOutput)

	DescribeConfigurationRevision(*mq.DescribeConfigurationRevisionInput) (*mq.DescribeConfigurationRevisionResponse, error)
	DescribeConfigurationRevisionWithContext(aws.Context, *mq.DescribeConfigurationRevisionInput, ...request.Option) (*mq.DescribeConfigurationRevisionResponse, error)
	DescribeConfigurationRevisionRequest(*mq.DescribeConfigurationRevisionInput) (*request.Request, *mq.DescribeConfigurationRevisionResponse)

	DescribeUser(*mq.DescribeUserInput) (*mq.DescribeUserResponse, error)
	DescribeUserWithContext(aws.Context, *mq.DescribeUserInput, ...request.Option) (*mq.DescribeUserResponse, error)
	DescribeUserRequest(*mq.DescribeUserInput) (*request.Request, *mq.DescribeUserResponse)

	ListBrokers(*mq.ListBrokersInput) (*mq.ListBrokersResponse, error)
	ListBrokersWithContext(aws.Context, *mq.ListBrokersInput, ...request.Option) (*mq.ListBrokersResponse, error)
	ListBrokersRequest(*mq.ListBrokersInput) (*request.Request, *mq.ListBrokersResponse)

	ListConfigurationRevisions(*mq.ListConfigurationRevisionsInput) (*mq.ListConfigurationRevisionsResponse, error)
	ListConfigurationRevisionsWithContext(aws.Context, *mq.ListConfigurationRevisionsInput, ...request.Option) (*mq.ListConfigurationRevisionsResponse, error)
	ListConfigurationRevisionsRequest(*mq.ListConfigurationRevisionsInput) (*request.Request, *mq.ListConfigurationRevisionsResponse)

	ListConfigurations(*mq.ListConfigurationsInput) (*mq.ListConfigurationsResponse, error)
	ListConfigurationsWithContext(aws.Context, *mq.ListConfigurationsInput, ...request.Option) (*mq.ListConfigurationsResponse, error)
	ListConfigurationsRequest(*mq.ListConfigurationsInput) (*request.Request, *mq.ListConfigurationsResponse)

	ListTags(*mq.ListTagsInput) (*mq.ListTagsOutput, error)
	ListTagsWithContext(aws.Context, *mq.ListTagsInput, ...request.Option) (*mq.ListTagsOutput, error)
	ListTagsRequest(*mq.ListTagsInput) (*request.Request, *mq.ListTagsOutput)

	ListUsers(*mq.ListUsersInput) (*mq.ListUsersResponse, error)
	ListUsersWithContext(aws.Context, *mq.ListUsersInput, ...request.Option) (*mq.ListUsersResponse, error)
	ListUsersRequest(*mq.ListUsersInput) (*request.Request, *mq.ListUsersResponse)

	RebootBroker(*mq.RebootBrokerInput) (*mq.RebootBrokerOutput, error)
	RebootBrokerWithContext(aws.Context, *mq.RebootBrokerInput, ...request.Option) (*mq.RebootBrokerOutput, error)
	RebootBrokerRequest(*mq.RebootBrokerInput) (*request.Request, *mq.RebootBrokerOutput)

	UpdateBroker(*mq.UpdateBrokerRequest) (*mq.UpdateBrokerResponse, error)
	UpdateBrokerWithContext(aws.Context, *mq.UpdateBrokerRequest, ...request.Option) (*mq.UpdateBrokerResponse, error)
	UpdateBrokerRequest(*mq.UpdateBrokerRequest) (*request.Request, *mq.UpdateBrokerResponse)

	UpdateConfiguration(*mq.UpdateConfigurationRequest) (*mq.UpdateConfigurationResponse, error)
	UpdateConfigurationWithContext(aws.Context, *mq.UpdateConfigurationRequest, ...request.Option) (*mq.UpdateConfigurationResponse, error)
	UpdateConfigurationRequest(*mq.UpdateConfigurationRequest) (*request.Request, *mq.UpdateConfigurationResponse)

	UpdateUser(*mq.UpdateUserRequest) (*mq.UpdateUserOutput, error)
	UpdateUserWithContext(aws.Context, *mq.UpdateUserRequest, ...request.Option) (*mq.UpdateUserOutput, error)
	UpdateUserRequest(*mq.UpdateUserRequest) (*request.Request, *mq.UpdateUserOutput)
}

var _ MQAPI = (*mq.MQ)(nil)
