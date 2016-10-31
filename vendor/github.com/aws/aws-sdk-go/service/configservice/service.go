// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package configservice

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

// AWS Config provides a way to keep track of the configurations of all the
// AWS resources associated with your AWS account. You can use AWS Config to
// get the current and historical configurations of each AWS resource and also
// to get information about the relationship between the resources. An AWS resource
// can be an Amazon Compute Cloud (Amazon EC2) instance, an Elastic Block Store
// (EBS) volume, an Elastic network Interface (ENI), or a security group. For
// a complete list of resources currently supported by AWS Config, see Supported
// AWS Resources (http://docs.aws.amazon.com/config/latest/developerguide/resource-config-reference.html#supported-resources).
//
// You can access and manage AWS Config through the AWS Management Console,
// the AWS Command Line Interface (AWS CLI), the AWS Config API, or the AWS
// SDKs for AWS Config
//
// This reference guide contains documentation for the AWS Config API and the
// AWS CLI commands that you can use to manage AWS Config.
//
// The AWS Config API uses the Signature Version 4 protocol for signing requests.
// For more information about how to sign a request with this protocol, see
// Signature Version 4 Signing Process (http://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
//
// For detailed information about AWS Config features and their associated actions
// or commands, as well as how to work with AWS Management Console, see What
// Is AWS Config? (http://docs.aws.amazon.com/config/latest/developerguide/WhatIsConfig.html)
// in the AWS Config Developer Guide.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type ConfigService struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "config"

// New creates a new instance of the ConfigService client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ConfigService client from just a session.
//     svc := configservice.New(mySession)
//
//     // Create a ConfigService client with additional configuration
//     svc := configservice.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ConfigService {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *ConfigService {
	svc := &ConfigService{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-11-12",
				JSONVersion:   "1.1",
				TargetPrefix:  "StarlingDoveService",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a ConfigService operation and runs any
// custom request initialization.
func (c *ConfigService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
