// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEndpointInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Number (ARN) of the endpoint being described.
	//
	// EndpointArn is a required field
	EndpointArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEndpointInput"}

	if s.EndpointArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeEndpointOutput struct {
	_ struct{} `type:"structure"`

	// Describes information associated with the specific endpoint.
	EndpointProperties *EndpointProperties `type:"structure"`
}

// String returns the string representation
func (s DescribeEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEndpoint = "DescribeEndpoint"

// DescribeEndpointRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Gets the properties associated with a specific endpoint. Use this operation
// to get the status of an endpoint.
//
//    // Example sending a request using DescribeEndpointRequest.
//    req := client.DescribeEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/DescribeEndpoint
func (c *Client) DescribeEndpointRequest(input *DescribeEndpointInput) DescribeEndpointRequest {
	op := &aws.Operation{
		Name:       opDescribeEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEndpointInput{}
	}

	req := c.newRequest(op, input, &DescribeEndpointOutput{})
	return DescribeEndpointRequest{Request: req, Input: input, Copy: c.DescribeEndpointRequest}
}

// DescribeEndpointRequest is the request type for the
// DescribeEndpoint API operation.
type DescribeEndpointRequest struct {
	*aws.Request
	Input *DescribeEndpointInput
	Copy  func(*DescribeEndpointInput) DescribeEndpointRequest
}

// Send marshals and sends the DescribeEndpoint API request.
func (r DescribeEndpointRequest) Send(ctx context.Context) (*DescribeEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEndpointResponse{
		DescribeEndpointOutput: r.Request.Data.(*DescribeEndpointOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEndpointResponse is the response type for the
// DescribeEndpoint API operation.
type DescribeEndpointResponse struct {
	*DescribeEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEndpoint request.
func (r *DescribeEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}