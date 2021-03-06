// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeScheduledActionsInput struct {
	_ struct{} `type:"structure"`

	// If true, retrieve only active scheduled actions. If false, retrieve only
	// disabled scheduled actions.
	Active *bool `type:"boolean"`

	// The end time in UTC of the scheduled action to retrieve. Only active scheduled
	// actions that have invocations before this time are retrieved.
	EndTime *time.Time `type:"timestamp"`

	// List of scheduled action filters.
	Filters []ScheduledActionFilter `locationNameList:"ScheduledActionFilter" type:"list"`

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeScheduledActions request
	// exceed the value specified in MaxRecords, AWS returns a value in the Marker
	// field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the Marker parameter and retrying
	// the request.
	Marker *string `type:"string"`

	// The maximum number of response records to return in each call. If the number
	// of remaining response records exceeds the specified MaxRecords value, a value
	// is returned in a marker field of the response. You can retrieve the next
	// set of records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The name of the scheduled action to retrieve.
	ScheduledActionName *string `type:"string"`

	// The start time in UTC of the scheduled actions to retrieve. Only active scheduled
	// actions that have invocations after this time are retrieved.
	StartTime *time.Time `type:"timestamp"`

	// The type of the scheduled actions to retrieve.
	TargetActionType ScheduledActionTypeValues `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeScheduledActionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeScheduledActionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeScheduledActionsInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeScheduledActionsOutput struct {
	_ struct{} `type:"structure"`

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeScheduledActions request
	// exceed the value specified in MaxRecords, AWS returns a value in the Marker
	// field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the Marker parameter and retrying
	// the request.
	Marker *string `type:"string"`

	// List of retrieved scheduled actions.
	ScheduledActions []ScheduledAction `locationNameList:"ScheduledAction" type:"list"`
}

// String returns the string representation
func (s DescribeScheduledActionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeScheduledActions = "DescribeScheduledActions"

// DescribeScheduledActionsRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Describes properties of scheduled actions.
//
//    // Example sending a request using DescribeScheduledActionsRequest.
//    req := client.DescribeScheduledActionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeScheduledActions
func (c *Client) DescribeScheduledActionsRequest(input *DescribeScheduledActionsInput) DescribeScheduledActionsRequest {
	op := &aws.Operation{
		Name:       opDescribeScheduledActions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeScheduledActionsInput{}
	}

	req := c.newRequest(op, input, &DescribeScheduledActionsOutput{})
	return DescribeScheduledActionsRequest{Request: req, Input: input, Copy: c.DescribeScheduledActionsRequest}
}

// DescribeScheduledActionsRequest is the request type for the
// DescribeScheduledActions API operation.
type DescribeScheduledActionsRequest struct {
	*aws.Request
	Input *DescribeScheduledActionsInput
	Copy  func(*DescribeScheduledActionsInput) DescribeScheduledActionsRequest
}

// Send marshals and sends the DescribeScheduledActions API request.
func (r DescribeScheduledActionsRequest) Send(ctx context.Context) (*DescribeScheduledActionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeScheduledActionsResponse{
		DescribeScheduledActionsOutput: r.Request.Data.(*DescribeScheduledActionsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeScheduledActionsRequestPaginator returns a paginator for DescribeScheduledActions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeScheduledActionsRequest(input)
//   p := redshift.NewDescribeScheduledActionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeScheduledActionsPaginator(req DescribeScheduledActionsRequest) DescribeScheduledActionsPaginator {
	return DescribeScheduledActionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeScheduledActionsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeScheduledActionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeScheduledActionsPaginator struct {
	aws.Pager
}

func (p *DescribeScheduledActionsPaginator) CurrentPage() *DescribeScheduledActionsOutput {
	return p.Pager.CurrentPage().(*DescribeScheduledActionsOutput)
}

// DescribeScheduledActionsResponse is the response type for the
// DescribeScheduledActions API operation.
type DescribeScheduledActionsResponse struct {
	*DescribeScheduledActionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeScheduledActions request.
func (r *DescribeScheduledActionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
