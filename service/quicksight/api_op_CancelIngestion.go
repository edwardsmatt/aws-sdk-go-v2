// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CancelIngestionInput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID of the dataset used in the ingestion.
	//
	// DataSetId is a required field
	DataSetId *string `location:"uri" locationName:"DataSetId" type:"string" required:"true"`

	// An ID for the ingestion.
	//
	// IngestionId is a required field
	IngestionId *string `location:"uri" locationName:"IngestionId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelIngestionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelIngestionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelIngestionInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.DataSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSetId"))
	}

	if s.IngestionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IngestionId"))
	}
	if s.IngestionId != nil && len(*s.IngestionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IngestionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelIngestionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DataSetId != nil {
		v := *s.DataSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DataSetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IngestionId != nil {
		v := *s.IngestionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "IngestionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CancelIngestionOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) for the data ingestion.
	Arn *string `type:"string"`

	// An ID for the ingestion.
	IngestionId *string `min:"1" type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s CancelIngestionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelIngestionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IngestionId != nil {
		v := *s.IngestionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IngestionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opCancelIngestion = "CancelIngestion"

// CancelIngestionRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Cancels an ongoing ingestion of data into SPICE.
//
//    // Example sending a request using CancelIngestionRequest.
//    req := client.CancelIngestionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/CancelIngestion
func (c *Client) CancelIngestionRequest(input *CancelIngestionInput) CancelIngestionRequest {
	op := &aws.Operation{
		Name:       opCancelIngestion,
		HTTPMethod: "DELETE",
		HTTPPath:   "/accounts/{AwsAccountId}/data-sets/{DataSetId}/ingestions/{IngestionId}",
	}

	if input == nil {
		input = &CancelIngestionInput{}
	}

	req := c.newRequest(op, input, &CancelIngestionOutput{})
	return CancelIngestionRequest{Request: req, Input: input, Copy: c.CancelIngestionRequest}
}

// CancelIngestionRequest is the request type for the
// CancelIngestion API operation.
type CancelIngestionRequest struct {
	*aws.Request
	Input *CancelIngestionInput
	Copy  func(*CancelIngestionInput) CancelIngestionRequest
}

// Send marshals and sends the CancelIngestion API request.
func (r CancelIngestionRequest) Send(ctx context.Context) (*CancelIngestionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelIngestionResponse{
		CancelIngestionOutput: r.Request.Data.(*CancelIngestionOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelIngestionResponse is the response type for the
// CancelIngestion API operation.
type CancelIngestionResponse struct {
	*CancelIngestionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelIngestion request.
func (r *CancelIngestionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
