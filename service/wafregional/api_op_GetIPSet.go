// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/waf"
)

type GetIPSetInput struct {
	_ struct{} `type:"structure"`

	// The IPSetId of the IPSet that you want to get. IPSetId is returned by CreateIPSet
	// and by ListIPSets.
	//
	// IPSetId is a required field
	IPSetId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetIPSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetIPSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetIPSetInput"}

	if s.IPSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IPSetId"))
	}
	if s.IPSetId != nil && len(*s.IPSetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IPSetId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetIPSetOutput struct {
	_ struct{} `type:"structure"`

	// Information about the IPSet that you specified in the GetIPSet request. For
	// more information, see the following topics:
	//
	//    * IPSet: Contains IPSetDescriptors, IPSetId, and Name
	//
	//    * IPSetDescriptors: Contains an array of IPSetDescriptor objects. Each
	//    IPSetDescriptor object contains Type and Value
	IPSet *waf.IPSet `type:"structure"`
}

// String returns the string representation
func (s GetIPSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetIPSet = "GetIPSet"

// GetIPSetRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Returns the IPSet that is specified by IPSetId.
//
//    // Example sending a request using GetIPSetRequest.
//    req := client.GetIPSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetIPSet
func (c *Client) GetIPSetRequest(input *GetIPSetInput) GetIPSetRequest {
	op := &aws.Operation{
		Name:       opGetIPSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetIPSetInput{}
	}

	req := c.newRequest(op, input, &GetIPSetOutput{})
	return GetIPSetRequest{Request: req, Input: input, Copy: c.GetIPSetRequest}
}

// GetIPSetRequest is the request type for the
// GetIPSet API operation.
type GetIPSetRequest struct {
	*aws.Request
	Input *GetIPSetInput
	Copy  func(*GetIPSetInput) GetIPSetRequest
}

// Send marshals and sends the GetIPSet API request.
func (r GetIPSetRequest) Send(ctx context.Context) (*GetIPSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetIPSetResponse{
		GetIPSetOutput: r.Request.Data.(*GetIPSetOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetIPSetResponse is the response type for the
// GetIPSet API operation.
type GetIPSetResponse struct {
	*GetIPSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetIPSet request.
func (r *GetIPSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
