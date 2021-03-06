// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeguruprofiler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request for ConfigureAgent operation.
type ConfigureAgentInput struct {
	_ struct{} `type:"structure"`

	// Identifier of the instance of compute fleet being profiled by the agent.
	// For instance, host name in EC2, task id for ECS, function name for AWS Lambda
	FleetInstanceId *string `locationName:"fleetInstanceId" min:"1" type:"string"`

	// The name of the profiling group.
	//
	// ProfilingGroupName is a required field
	ProfilingGroupName *string `location:"uri" locationName:"profilingGroupName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ConfigureAgentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConfigureAgentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ConfigureAgentInput"}
	if s.FleetInstanceId != nil && len(*s.FleetInstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetInstanceId", 1))
	}

	if s.ProfilingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProfilingGroupName"))
	}
	if s.ProfilingGroupName != nil && len(*s.ProfilingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProfilingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ConfigureAgentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FleetInstanceId != nil {
		v := *s.FleetInstanceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "fleetInstanceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProfilingGroupName != nil {
		v := *s.ProfilingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "profilingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Response for ConfigureAgent operation.
type ConfigureAgentOutput struct {
	_ struct{} `type:"structure" payload:"Configuration"`

	// The configuration for the agent to use.
	//
	// Configuration is a required field
	Configuration *AgentConfiguration `locationName:"configuration" type:"structure" required:"true"`
}

// String returns the string representation
func (s ConfigureAgentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ConfigureAgentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Configuration != nil {
		v := s.Configuration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "configuration", v, metadata)
	}
	return nil
}

const opConfigureAgent = "ConfigureAgent"

// ConfigureAgentRequest returns a request value for making API operation for
// Amazon CodeGuru Profiler.
//
// Provides the configuration to use for an agent of the profiling group.
//
//    // Example sending a request using ConfigureAgentRequest.
//    req := client.ConfigureAgentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeguruprofiler-2019-07-18/ConfigureAgent
func (c *Client) ConfigureAgentRequest(input *ConfigureAgentInput) ConfigureAgentRequest {
	op := &aws.Operation{
		Name:       opConfigureAgent,
		HTTPMethod: "POST",
		HTTPPath:   "/profilingGroups/{profilingGroupName}/configureAgent",
	}

	if input == nil {
		input = &ConfigureAgentInput{}
	}

	req := c.newRequest(op, input, &ConfigureAgentOutput{})
	return ConfigureAgentRequest{Request: req, Input: input, Copy: c.ConfigureAgentRequest}
}

// ConfigureAgentRequest is the request type for the
// ConfigureAgent API operation.
type ConfigureAgentRequest struct {
	*aws.Request
	Input *ConfigureAgentInput
	Copy  func(*ConfigureAgentInput) ConfigureAgentRequest
}

// Send marshals and sends the ConfigureAgent API request.
func (r ConfigureAgentRequest) Send(ctx context.Context) (*ConfigureAgentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ConfigureAgentResponse{
		ConfigureAgentOutput: r.Request.Data.(*ConfigureAgentOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ConfigureAgentResponse is the response type for the
// ConfigureAgent API operation.
type ConfigureAgentResponse struct {
	*ConfigureAgentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ConfigureAgent request.
func (r *ConfigureAgentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
