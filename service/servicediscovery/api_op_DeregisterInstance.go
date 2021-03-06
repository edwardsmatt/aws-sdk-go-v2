// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeregisterInstanceInput struct {
	_ struct{} `type:"structure"`

	// The value that you specified for Id in the RegisterInstance request.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// The ID of the service that the instance is associated with.
	//
	// ServiceId is a required field
	ServiceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeregisterInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterInstanceInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if s.ServiceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeregisterInstanceOutput struct {
	_ struct{} `type:"structure"`

	// A value that you can use to determine whether the request completed successfully.
	// For more information, see GetOperation.
	OperationId *string `type:"string"`
}

// String returns the string representation
func (s DeregisterInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeregisterInstance = "DeregisterInstance"

// DeregisterInstanceRequest returns a request value for making API operation for
// AWS Cloud Map.
//
// Deletes the Amazon Route 53 DNS records and health check, if any, that AWS
// Cloud Map created for the specified instance.
//
//    // Example sending a request using DeregisterInstanceRequest.
//    req := client.DeregisterInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/DeregisterInstance
func (c *Client) DeregisterInstanceRequest(input *DeregisterInstanceInput) DeregisterInstanceRequest {
	op := &aws.Operation{
		Name:       opDeregisterInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterInstanceInput{}
	}

	req := c.newRequest(op, input, &DeregisterInstanceOutput{})
	return DeregisterInstanceRequest{Request: req, Input: input, Copy: c.DeregisterInstanceRequest}
}

// DeregisterInstanceRequest is the request type for the
// DeregisterInstance API operation.
type DeregisterInstanceRequest struct {
	*aws.Request
	Input *DeregisterInstanceInput
	Copy  func(*DeregisterInstanceInput) DeregisterInstanceRequest
}

// Send marshals and sends the DeregisterInstance API request.
func (r DeregisterInstanceRequest) Send(ctx context.Context) (*DeregisterInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterInstanceResponse{
		DeregisterInstanceOutput: r.Request.Data.(*DeregisterInstanceOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterInstanceResponse is the response type for the
// DeregisterInstance API operation.
type DeregisterInstanceResponse struct {
	*DeregisterInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterInstance request.
func (r *DeregisterInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
