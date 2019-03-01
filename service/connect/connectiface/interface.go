// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package connectiface provides an interface to enable mocking the Amazon Connect Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package connectiface

import (
	"github.com/Beeketing/aws-sdk-go/aws"
	"github.com/Beeketing/aws-sdk-go/aws/request"
	"github.com/Beeketing/aws-sdk-go/service/connect"
)

// ConnectAPI provides an interface to enable mocking the
// connect.Connect service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Connect Service.
//    func myFunc(svc connectiface.ConnectAPI) bool {
//        // Make svc.CreateUser request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := connect.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockConnectClient struct {
//        connectiface.ConnectAPI
//    }
//    func (m *mockConnectClient) CreateUser(input *connect.CreateUserInput) (*connect.CreateUserOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockConnectClient{}
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
type ConnectAPI interface {
	CreateUser(*connect.CreateUserInput) (*connect.CreateUserOutput, error)
	CreateUserWithContext(aws.Context, *connect.CreateUserInput, ...request.Option) (*connect.CreateUserOutput, error)
	CreateUserRequest(*connect.CreateUserInput) (*request.Request, *connect.CreateUserOutput)

	DeleteUser(*connect.DeleteUserInput) (*connect.DeleteUserOutput, error)
	DeleteUserWithContext(aws.Context, *connect.DeleteUserInput, ...request.Option) (*connect.DeleteUserOutput, error)
	DeleteUserRequest(*connect.DeleteUserInput) (*request.Request, *connect.DeleteUserOutput)

	DescribeUser(*connect.DescribeUserInput) (*connect.DescribeUserOutput, error)
	DescribeUserWithContext(aws.Context, *connect.DescribeUserInput, ...request.Option) (*connect.DescribeUserOutput, error)
	DescribeUserRequest(*connect.DescribeUserInput) (*request.Request, *connect.DescribeUserOutput)

	DescribeUserHierarchyGroup(*connect.DescribeUserHierarchyGroupInput) (*connect.DescribeUserHierarchyGroupOutput, error)
	DescribeUserHierarchyGroupWithContext(aws.Context, *connect.DescribeUserHierarchyGroupInput, ...request.Option) (*connect.DescribeUserHierarchyGroupOutput, error)
	DescribeUserHierarchyGroupRequest(*connect.DescribeUserHierarchyGroupInput) (*request.Request, *connect.DescribeUserHierarchyGroupOutput)

	DescribeUserHierarchyStructure(*connect.DescribeUserHierarchyStructureInput) (*connect.DescribeUserHierarchyStructureOutput, error)
	DescribeUserHierarchyStructureWithContext(aws.Context, *connect.DescribeUserHierarchyStructureInput, ...request.Option) (*connect.DescribeUserHierarchyStructureOutput, error)
	DescribeUserHierarchyStructureRequest(*connect.DescribeUserHierarchyStructureInput) (*request.Request, *connect.DescribeUserHierarchyStructureOutput)

	GetContactAttributes(*connect.GetContactAttributesInput) (*connect.GetContactAttributesOutput, error)
	GetContactAttributesWithContext(aws.Context, *connect.GetContactAttributesInput, ...request.Option) (*connect.GetContactAttributesOutput, error)
	GetContactAttributesRequest(*connect.GetContactAttributesInput) (*request.Request, *connect.GetContactAttributesOutput)

	GetCurrentMetricData(*connect.GetCurrentMetricDataInput) (*connect.GetCurrentMetricDataOutput, error)
	GetCurrentMetricDataWithContext(aws.Context, *connect.GetCurrentMetricDataInput, ...request.Option) (*connect.GetCurrentMetricDataOutput, error)
	GetCurrentMetricDataRequest(*connect.GetCurrentMetricDataInput) (*request.Request, *connect.GetCurrentMetricDataOutput)

	GetCurrentMetricDataPages(*connect.GetCurrentMetricDataInput, func(*connect.GetCurrentMetricDataOutput, bool) bool) error
	GetCurrentMetricDataPagesWithContext(aws.Context, *connect.GetCurrentMetricDataInput, func(*connect.GetCurrentMetricDataOutput, bool) bool, ...request.Option) error

	GetFederationToken(*connect.GetFederationTokenInput) (*connect.GetFederationTokenOutput, error)
	GetFederationTokenWithContext(aws.Context, *connect.GetFederationTokenInput, ...request.Option) (*connect.GetFederationTokenOutput, error)
	GetFederationTokenRequest(*connect.GetFederationTokenInput) (*request.Request, *connect.GetFederationTokenOutput)

	GetMetricData(*connect.GetMetricDataInput) (*connect.GetMetricDataOutput, error)
	GetMetricDataWithContext(aws.Context, *connect.GetMetricDataInput, ...request.Option) (*connect.GetMetricDataOutput, error)
	GetMetricDataRequest(*connect.GetMetricDataInput) (*request.Request, *connect.GetMetricDataOutput)

	GetMetricDataPages(*connect.GetMetricDataInput, func(*connect.GetMetricDataOutput, bool) bool) error
	GetMetricDataPagesWithContext(aws.Context, *connect.GetMetricDataInput, func(*connect.GetMetricDataOutput, bool) bool, ...request.Option) error

	ListRoutingProfiles(*connect.ListRoutingProfilesInput) (*connect.ListRoutingProfilesOutput, error)
	ListRoutingProfilesWithContext(aws.Context, *connect.ListRoutingProfilesInput, ...request.Option) (*connect.ListRoutingProfilesOutput, error)
	ListRoutingProfilesRequest(*connect.ListRoutingProfilesInput) (*request.Request, *connect.ListRoutingProfilesOutput)

	ListSecurityProfiles(*connect.ListSecurityProfilesInput) (*connect.ListSecurityProfilesOutput, error)
	ListSecurityProfilesWithContext(aws.Context, *connect.ListSecurityProfilesInput, ...request.Option) (*connect.ListSecurityProfilesOutput, error)
	ListSecurityProfilesRequest(*connect.ListSecurityProfilesInput) (*request.Request, *connect.ListSecurityProfilesOutput)

	ListUserHierarchyGroups(*connect.ListUserHierarchyGroupsInput) (*connect.ListUserHierarchyGroupsOutput, error)
	ListUserHierarchyGroupsWithContext(aws.Context, *connect.ListUserHierarchyGroupsInput, ...request.Option) (*connect.ListUserHierarchyGroupsOutput, error)
	ListUserHierarchyGroupsRequest(*connect.ListUserHierarchyGroupsInput) (*request.Request, *connect.ListUserHierarchyGroupsOutput)

	ListUsers(*connect.ListUsersInput) (*connect.ListUsersOutput, error)
	ListUsersWithContext(aws.Context, *connect.ListUsersInput, ...request.Option) (*connect.ListUsersOutput, error)
	ListUsersRequest(*connect.ListUsersInput) (*request.Request, *connect.ListUsersOutput)

	StartOutboundVoiceContact(*connect.StartOutboundVoiceContactInput) (*connect.StartOutboundVoiceContactOutput, error)
	StartOutboundVoiceContactWithContext(aws.Context, *connect.StartOutboundVoiceContactInput, ...request.Option) (*connect.StartOutboundVoiceContactOutput, error)
	StartOutboundVoiceContactRequest(*connect.StartOutboundVoiceContactInput) (*request.Request, *connect.StartOutboundVoiceContactOutput)

	StopContact(*connect.StopContactInput) (*connect.StopContactOutput, error)
	StopContactWithContext(aws.Context, *connect.StopContactInput, ...request.Option) (*connect.StopContactOutput, error)
	StopContactRequest(*connect.StopContactInput) (*request.Request, *connect.StopContactOutput)

	UpdateContactAttributes(*connect.UpdateContactAttributesInput) (*connect.UpdateContactAttributesOutput, error)
	UpdateContactAttributesWithContext(aws.Context, *connect.UpdateContactAttributesInput, ...request.Option) (*connect.UpdateContactAttributesOutput, error)
	UpdateContactAttributesRequest(*connect.UpdateContactAttributesInput) (*request.Request, *connect.UpdateContactAttributesOutput)

	UpdateUserHierarchy(*connect.UpdateUserHierarchyInput) (*connect.UpdateUserHierarchyOutput, error)
	UpdateUserHierarchyWithContext(aws.Context, *connect.UpdateUserHierarchyInput, ...request.Option) (*connect.UpdateUserHierarchyOutput, error)
	UpdateUserHierarchyRequest(*connect.UpdateUserHierarchyInput) (*request.Request, *connect.UpdateUserHierarchyOutput)

	UpdateUserIdentityInfo(*connect.UpdateUserIdentityInfoInput) (*connect.UpdateUserIdentityInfoOutput, error)
	UpdateUserIdentityInfoWithContext(aws.Context, *connect.UpdateUserIdentityInfoInput, ...request.Option) (*connect.UpdateUserIdentityInfoOutput, error)
	UpdateUserIdentityInfoRequest(*connect.UpdateUserIdentityInfoInput) (*request.Request, *connect.UpdateUserIdentityInfoOutput)

	UpdateUserPhoneConfig(*connect.UpdateUserPhoneConfigInput) (*connect.UpdateUserPhoneConfigOutput, error)
	UpdateUserPhoneConfigWithContext(aws.Context, *connect.UpdateUserPhoneConfigInput, ...request.Option) (*connect.UpdateUserPhoneConfigOutput, error)
	UpdateUserPhoneConfigRequest(*connect.UpdateUserPhoneConfigInput) (*request.Request, *connect.UpdateUserPhoneConfigOutput)

	UpdateUserRoutingProfile(*connect.UpdateUserRoutingProfileInput) (*connect.UpdateUserRoutingProfileOutput, error)
	UpdateUserRoutingProfileWithContext(aws.Context, *connect.UpdateUserRoutingProfileInput, ...request.Option) (*connect.UpdateUserRoutingProfileOutput, error)
	UpdateUserRoutingProfileRequest(*connect.UpdateUserRoutingProfileInput) (*request.Request, *connect.UpdateUserRoutingProfileOutput)

	UpdateUserSecurityProfiles(*connect.UpdateUserSecurityProfilesInput) (*connect.UpdateUserSecurityProfilesOutput, error)
	UpdateUserSecurityProfilesWithContext(aws.Context, *connect.UpdateUserSecurityProfilesInput, ...request.Option) (*connect.UpdateUserSecurityProfilesOutput, error)
	UpdateUserSecurityProfilesRequest(*connect.UpdateUserSecurityProfilesInput) (*request.Request, *connect.UpdateUserSecurityProfilesOutput)
}

var _ ConnectAPI = (*connect.Connect)(nil)
