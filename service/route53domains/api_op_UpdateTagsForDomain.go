// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The UpdateTagsForDomainRequest includes the following elements.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/UpdateTagsForDomainRequest
type UpdateTagsForDomainInput struct {
	_ struct{} `type:"structure"`

	// The domain for which you want to add or update tags.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`

	// A list of the tag keys and values that you want to add or update. If you
	// specify a key that already exists, the corresponding value will be replaced.
	TagsToUpdate []Tag `type:"list"`
}

// String returns the string representation
func (s UpdateTagsForDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTagsForDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTagsForDomainInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/UpdateTagsForDomainResponse
type UpdateTagsForDomainOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateTagsForDomainOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateTagsForDomain = "UpdateTagsForDomain"

// UpdateTagsForDomainRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// This operation adds or updates tags for a specified domain.
//
// All tag operations are eventually consistent; subsequent operations might
// not immediately represent all issued operations.
//
//    // Example sending a request using UpdateTagsForDomainRequest.
//    req := client.UpdateTagsForDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/UpdateTagsForDomain
func (c *Client) UpdateTagsForDomainRequest(input *UpdateTagsForDomainInput) UpdateTagsForDomainRequest {
	op := &aws.Operation{
		Name:       opUpdateTagsForDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateTagsForDomainInput{}
	}

	req := c.newRequest(op, input, &UpdateTagsForDomainOutput{})
	return UpdateTagsForDomainRequest{Request: req, Input: input, Copy: c.UpdateTagsForDomainRequest}
}

// UpdateTagsForDomainRequest is the request type for the
// UpdateTagsForDomain API operation.
type UpdateTagsForDomainRequest struct {
	*aws.Request
	Input *UpdateTagsForDomainInput
	Copy  func(*UpdateTagsForDomainInput) UpdateTagsForDomainRequest
}

// Send marshals and sends the UpdateTagsForDomain API request.
func (r UpdateTagsForDomainRequest) Send(ctx context.Context) (*UpdateTagsForDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTagsForDomainResponse{
		UpdateTagsForDomainOutput: r.Request.Data.(*UpdateTagsForDomainOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTagsForDomainResponse is the response type for the
// UpdateTagsForDomain API operation.
type UpdateTagsForDomainResponse struct {
	*UpdateTagsForDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTagsForDomain request.
func (r *UpdateTagsForDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}