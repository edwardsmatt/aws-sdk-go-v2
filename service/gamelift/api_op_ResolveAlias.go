// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type ResolveAliasInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for the alias you want to resolve.
	//
	// AliasId is a required field
	AliasId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ResolveAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResolveAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResolveAliasInput"}

	if s.AliasId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AliasId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type ResolveAliasOutput struct {
	_ struct{} `type:"structure"`

	// Fleet identifier that is associated with the requested alias.
	FleetId *string `type:"string"`
}

// String returns the string representation
func (s ResolveAliasOutput) String() string {
	return awsutil.Prettify(s)
}

const opResolveAlias = "ResolveAlias"

// ResolveAliasRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves the fleet ID that a specified alias is currently pointing to.
//
//    * CreateAlias
//
//    * ListAliases
//
//    * DescribeAlias
//
//    * UpdateAlias
//
//    * DeleteAlias
//
//    * ResolveAlias
//
//    // Example sending a request using ResolveAliasRequest.
//    req := client.ResolveAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/ResolveAlias
func (c *Client) ResolveAliasRequest(input *ResolveAliasInput) ResolveAliasRequest {
	op := &aws.Operation{
		Name:       opResolveAlias,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResolveAliasInput{}
	}

	req := c.newRequest(op, input, &ResolveAliasOutput{})
	return ResolveAliasRequest{Request: req, Input: input, Copy: c.ResolveAliasRequest}
}

// ResolveAliasRequest is the request type for the
// ResolveAlias API operation.
type ResolveAliasRequest struct {
	*aws.Request
	Input *ResolveAliasInput
	Copy  func(*ResolveAliasInput) ResolveAliasRequest
}

// Send marshals and sends the ResolveAlias API request.
func (r ResolveAliasRequest) Send(ctx context.Context) (*ResolveAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResolveAliasResponse{
		ResolveAliasOutput: r.Request.Data.(*ResolveAliasOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResolveAliasResponse is the response type for the
// ResolveAlias API operation.
type ResolveAliasResponse struct {
	*ResolveAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResolveAlias request.
func (r *ResolveAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
