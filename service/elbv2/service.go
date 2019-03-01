// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elbv2

import (
	"github.com/Beeketing/aws-sdk-go/aws"
	"github.com/Beeketing/aws-sdk-go/aws/client"
	"github.com/Beeketing/aws-sdk-go/aws/client/metadata"
	"github.com/Beeketing/aws-sdk-go/aws/request"
	"github.com/Beeketing/aws-sdk-go/aws/signer/v4"
	"github.com/Beeketing/aws-sdk-go/private/protocol/query"
)

// ELBV2 provides the API operation methods for making requests to
// Elastic Load Balancing. See this package's package overview docs
// for details on the service.
//
// ELBV2 methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type ELBV2 struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "elasticloadbalancing"      // Name of service.
	EndpointsID = ServiceName                 // ID to lookup a service endpoint with.
	ServiceID   = "Elastic Load Balancing v2" // ServiceID is a unique identifer of a specific service.
)

// New creates a new instance of the ELBV2 client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ELBV2 client from just a session.
//     svc := elbv2.New(mySession)
//
//     // Create a ELBV2 client with additional configuration
//     svc := elbv2.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ELBV2 {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *ELBV2 {
	svc := &ELBV2{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-12-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a ELBV2 operation and runs any
// custom request initialization.
func (c *ELBV2) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
