// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetInstancesHealthStatusInput struct {
	_ struct{} `type:"structure"`

	// An array that contains the IDs of all the instances that you want to get
	// the health status for.
	//
	// If you omit Instances, AWS Cloud Map returns the health status for all the
	// instances that are associated with the specified service.
	//
	// To get the IDs for the instances that you've registered by using a specified
	// service, submit a ListInstances request.
	Instances []string `min:"1" type:"list"`

	// The maximum number of instances that you want AWS Cloud Map to return in
	// the response to a GetInstancesHealthStatus request. If you don't specify
	// a value for MaxResults, AWS Cloud Map returns up to 100 instances.
	MaxResults *int64 `min:"1" type:"integer"`

	// For the first GetInstancesHealthStatus request, omit this value.
	//
	// If more than MaxResults instances match the specified criteria, you can submit
	// another GetInstancesHealthStatus request to get the next group of results.
	// Specify the value of NextToken from the previous response in the next request.
	NextToken *string `type:"string"`

	// The ID of the service that the instance is associated with.
	//
	// ServiceId is a required field
	ServiceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetInstancesHealthStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetInstancesHealthStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetInstancesHealthStatusInput"}
	if s.Instances != nil && len(s.Instances) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Instances", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ServiceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetInstancesHealthStatusOutput struct {
	_ struct{} `type:"structure"`

	// If more than MaxResults instances match the specified criteria, you can submit
	// another GetInstancesHealthStatus request to get the next group of results.
	// Specify the value of NextToken from the previous response in the next request.
	NextToken *string `type:"string"`

	// A complex type that contains the IDs and the health status of the instances
	// that you specified in the GetInstancesHealthStatus request.
	Status map[string]HealthStatus `type:"map"`
}

// String returns the string representation
func (s GetInstancesHealthStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetInstancesHealthStatus = "GetInstancesHealthStatus"

// GetInstancesHealthStatusRequest returns a request value for making API operation for
// AWS Cloud Map.
//
// Gets the current health status (Healthy, Unhealthy, or Unknown) of one or
// more instances that are associated with a specified service.
//
// There is a brief delay between when you register an instance and when the
// health status for the instance is available.
//
//    // Example sending a request using GetInstancesHealthStatusRequest.
//    req := client.GetInstancesHealthStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/GetInstancesHealthStatus
func (c *Client) GetInstancesHealthStatusRequest(input *GetInstancesHealthStatusInput) GetInstancesHealthStatusRequest {
	op := &aws.Operation{
		Name:       opGetInstancesHealthStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetInstancesHealthStatusInput{}
	}

	req := c.newRequest(op, input, &GetInstancesHealthStatusOutput{})
	return GetInstancesHealthStatusRequest{Request: req, Input: input, Copy: c.GetInstancesHealthStatusRequest}
}

// GetInstancesHealthStatusRequest is the request type for the
// GetInstancesHealthStatus API operation.
type GetInstancesHealthStatusRequest struct {
	*aws.Request
	Input *GetInstancesHealthStatusInput
	Copy  func(*GetInstancesHealthStatusInput) GetInstancesHealthStatusRequest
}

// Send marshals and sends the GetInstancesHealthStatus API request.
func (r GetInstancesHealthStatusRequest) Send(ctx context.Context) (*GetInstancesHealthStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInstancesHealthStatusResponse{
		GetInstancesHealthStatusOutput: r.Request.Data.(*GetInstancesHealthStatusOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetInstancesHealthStatusRequestPaginator returns a paginator for GetInstancesHealthStatus.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetInstancesHealthStatusRequest(input)
//   p := servicediscovery.NewGetInstancesHealthStatusRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetInstancesHealthStatusPaginator(req GetInstancesHealthStatusRequest) GetInstancesHealthStatusPaginator {
	return GetInstancesHealthStatusPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetInstancesHealthStatusInput
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

// GetInstancesHealthStatusPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetInstancesHealthStatusPaginator struct {
	aws.Pager
}

func (p *GetInstancesHealthStatusPaginator) CurrentPage() *GetInstancesHealthStatusOutput {
	return p.Pager.CurrentPage().(*GetInstancesHealthStatusOutput)
}

// GetInstancesHealthStatusResponse is the response type for the
// GetInstancesHealthStatus API operation.
type GetInstancesHealthStatusResponse struct {
	*GetInstancesHealthStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInstancesHealthStatus request.
func (r *GetInstancesHealthStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
