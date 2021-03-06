// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListTagsLogGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the log group.
	//
	// LogGroupName is a required field
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsLogGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsLogGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsLogGroupInput"}

	if s.LogGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogGroupName"))
	}
	if s.LogGroupName != nil && len(*s.LogGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListTagsLogGroupOutput struct {
	_ struct{} `type:"structure"`

	// The tags for the log group.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`
}

// String returns the string representation
func (s ListTagsLogGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTagsLogGroup = "ListTagsLogGroup"

// ListTagsLogGroupRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Lists the tags for the specified log group.
//
//    // Example sending a request using ListTagsLogGroupRequest.
//    req := client.ListTagsLogGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/ListTagsLogGroup
func (c *Client) ListTagsLogGroupRequest(input *ListTagsLogGroupInput) ListTagsLogGroupRequest {
	op := &aws.Operation{
		Name:       opListTagsLogGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListTagsLogGroupInput{}
	}

	req := c.newRequest(op, input, &ListTagsLogGroupOutput{})
	return ListTagsLogGroupRequest{Request: req, Input: input, Copy: c.ListTagsLogGroupRequest}
}

// ListTagsLogGroupRequest is the request type for the
// ListTagsLogGroup API operation.
type ListTagsLogGroupRequest struct {
	*aws.Request
	Input *ListTagsLogGroupInput
	Copy  func(*ListTagsLogGroupInput) ListTagsLogGroupRequest
}

// Send marshals and sends the ListTagsLogGroup API request.
func (r ListTagsLogGroupRequest) Send(ctx context.Context) (*ListTagsLogGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsLogGroupResponse{
		ListTagsLogGroupOutput: r.Request.Data.(*ListTagsLogGroupOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTagsLogGroupResponse is the response type for the
// ListTagsLogGroup API operation.
type ListTagsLogGroupResponse struct {
	*ListTagsLogGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTagsLogGroup request.
func (r *ListTagsLogGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
