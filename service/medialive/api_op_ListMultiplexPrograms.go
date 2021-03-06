// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListMultiplexProgramsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// MultiplexId is a required field
	MultiplexId *string `location:"uri" locationName:"multiplexId" type:"string" required:"true"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListMultiplexProgramsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMultiplexProgramsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMultiplexProgramsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.MultiplexId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MultiplexId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMultiplexProgramsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MultiplexId != nil {
		v := *s.MultiplexId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "multiplexId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListMultiplexProgramsOutput struct {
	_ struct{} `type:"structure"`

	MultiplexPrograms []MultiplexProgramSummary `locationName:"multiplexPrograms" type:"list"`

	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListMultiplexProgramsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMultiplexProgramsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.MultiplexPrograms != nil {
		v := s.MultiplexPrograms

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "multiplexPrograms", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListMultiplexPrograms = "ListMultiplexPrograms"

// ListMultiplexProgramsRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// List the programs that currently exist for a specific multiplex.
//
//    // Example sending a request using ListMultiplexProgramsRequest.
//    req := client.ListMultiplexProgramsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/ListMultiplexPrograms
func (c *Client) ListMultiplexProgramsRequest(input *ListMultiplexProgramsInput) ListMultiplexProgramsRequest {
	op := &aws.Operation{
		Name:       opListMultiplexPrograms,
		HTTPMethod: "GET",
		HTTPPath:   "/prod/multiplexes/{multiplexId}/programs",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListMultiplexProgramsInput{}
	}

	req := c.newRequest(op, input, &ListMultiplexProgramsOutput{})
	return ListMultiplexProgramsRequest{Request: req, Input: input, Copy: c.ListMultiplexProgramsRequest}
}

// ListMultiplexProgramsRequest is the request type for the
// ListMultiplexPrograms API operation.
type ListMultiplexProgramsRequest struct {
	*aws.Request
	Input *ListMultiplexProgramsInput
	Copy  func(*ListMultiplexProgramsInput) ListMultiplexProgramsRequest
}

// Send marshals and sends the ListMultiplexPrograms API request.
func (r ListMultiplexProgramsRequest) Send(ctx context.Context) (*ListMultiplexProgramsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMultiplexProgramsResponse{
		ListMultiplexProgramsOutput: r.Request.Data.(*ListMultiplexProgramsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListMultiplexProgramsRequestPaginator returns a paginator for ListMultiplexPrograms.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListMultiplexProgramsRequest(input)
//   p := medialive.NewListMultiplexProgramsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListMultiplexProgramsPaginator(req ListMultiplexProgramsRequest) ListMultiplexProgramsPaginator {
	return ListMultiplexProgramsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListMultiplexProgramsInput
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

// ListMultiplexProgramsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListMultiplexProgramsPaginator struct {
	aws.Pager
}

func (p *ListMultiplexProgramsPaginator) CurrentPage() *ListMultiplexProgramsOutput {
	return p.Pager.CurrentPage().(*ListMultiplexProgramsOutput)
}

// ListMultiplexProgramsResponse is the response type for the
// ListMultiplexPrograms API operation.
type ListMultiplexProgramsResponse struct {
	*ListMultiplexProgramsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMultiplexPrograms request.
func (r *ListMultiplexProgramsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
