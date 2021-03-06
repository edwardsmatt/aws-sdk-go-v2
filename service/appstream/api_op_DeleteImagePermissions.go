// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteImagePermissionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the private image.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The 12-digit identifier of the AWS account for which to delete image permissions.
	//
	// SharedAccountId is a required field
	SharedAccountId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteImagePermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteImagePermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteImagePermissionsInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.SharedAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SharedAccountId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteImagePermissionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteImagePermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteImagePermissions = "DeleteImagePermissions"

// DeleteImagePermissionsRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Deletes permissions for the specified private image. After you delete permissions
// for an image, AWS accounts to which you previously granted these permissions
// can no longer use the image.
//
//    // Example sending a request using DeleteImagePermissionsRequest.
//    req := client.DeleteImagePermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DeleteImagePermissions
func (c *Client) DeleteImagePermissionsRequest(input *DeleteImagePermissionsInput) DeleteImagePermissionsRequest {
	op := &aws.Operation{
		Name:       opDeleteImagePermissions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteImagePermissionsInput{}
	}

	req := c.newRequest(op, input, &DeleteImagePermissionsOutput{})
	return DeleteImagePermissionsRequest{Request: req, Input: input, Copy: c.DeleteImagePermissionsRequest}
}

// DeleteImagePermissionsRequest is the request type for the
// DeleteImagePermissions API operation.
type DeleteImagePermissionsRequest struct {
	*aws.Request
	Input *DeleteImagePermissionsInput
	Copy  func(*DeleteImagePermissionsInput) DeleteImagePermissionsRequest
}

// Send marshals and sends the DeleteImagePermissions API request.
func (r DeleteImagePermissionsRequest) Send(ctx context.Context) (*DeleteImagePermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteImagePermissionsResponse{
		DeleteImagePermissionsOutput: r.Request.Data.(*DeleteImagePermissionsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteImagePermissionsResponse is the response type for the
// DeleteImagePermissions API operation.
type DeleteImagePermissionsResponse struct {
	*DeleteImagePermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteImagePermissions request.
func (r *DeleteImagePermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
