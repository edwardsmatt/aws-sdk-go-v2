// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/PostCommentForPullRequestInput
type PostCommentForPullRequestInput struct {
	_ struct{} `type:"structure"`

	// The full commit ID of the commit in the source branch that is the current
	// tip of the branch for the pull request when you post the comment.
	//
	// AfterCommitId is a required field
	AfterCommitId *string `locationName:"afterCommitId" type:"string" required:"true"`

	// The full commit ID of the commit in the destination branch that was the tip
	// of the branch at the time the pull request was created.
	//
	// BeforeCommitId is a required field
	BeforeCommitId *string `locationName:"beforeCommitId" type:"string" required:"true"`

	// A unique, client-generated idempotency token that when provided in a request,
	// ensures the request cannot be repeated with a changed parameter. If a request
	// is received with the same parameters and a token is included, the request
	// will return information about the initial request that used that token.
	ClientRequestToken *string `locationName:"clientRequestToken" type:"string" idempotencyToken:"true"`

	// The content of your comment on the change.
	//
	// Content is a required field
	Content *string `locationName:"content" type:"string" required:"true"`

	// The location of the change where you want to post your comment. If no location
	// is provided, the comment will be posted as a general comment on the pull
	// request difference between the before commit ID and the after commit ID.
	Location *Location `locationName:"location" type:"structure"`

	// The system-generated ID of the pull request. To get this ID, use ListPullRequests.
	//
	// PullRequestId is a required field
	PullRequestId *string `locationName:"pullRequestId" type:"string" required:"true"`

	// The name of the repository where you want to post a comment on a pull request.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PostCommentForPullRequestInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PostCommentForPullRequestInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PostCommentForPullRequestInput"}

	if s.AfterCommitId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AfterCommitId"))
	}

	if s.BeforeCommitId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BeforeCommitId"))
	}

	if s.Content == nil {
		invalidParams.Add(aws.NewErrParamRequired("Content"))
	}

	if s.PullRequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PullRequestId"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/PostCommentForPullRequestOutput
type PostCommentForPullRequestOutput struct {
	_ struct{} `type:"structure"`

	// In the directionality of the pull request, the blob ID of the 'after' blob.
	AfterBlobId *string `locationName:"afterBlobId" type:"string"`

	// The full commit ID of the commit in the destination branch where the pull
	// request will be merged.
	AfterCommitId *string `locationName:"afterCommitId" type:"string"`

	// In the directionality of the pull request, the blob ID of the 'before' blob.
	BeforeBlobId *string `locationName:"beforeBlobId" type:"string"`

	// The full commit ID of the commit in the source branch used to create the
	// pull request, or in the case of an updated pull request, the full commit
	// ID of the commit used to update the pull request.
	BeforeCommitId *string `locationName:"beforeCommitId" type:"string"`

	// The content of the comment you posted.
	Comment *Comment `locationName:"comment" type:"structure"`

	// The location of the change where you posted your comment.
	Location *Location `locationName:"location" type:"structure"`

	// The system-generated ID of the pull request.
	PullRequestId *string `locationName:"pullRequestId" type:"string"`

	// The name of the repository where you posted a comment on a pull request.
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string"`
}

// String returns the string representation
func (s PostCommentForPullRequestOutput) String() string {
	return awsutil.Prettify(s)
}

const opPostCommentForPullRequest = "PostCommentForPullRequest"

// PostCommentForPullRequestRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Posts a comment on a pull request.
//
//    // Example sending a request using PostCommentForPullRequestRequest.
//    req := client.PostCommentForPullRequestRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/PostCommentForPullRequest
func (c *Client) PostCommentForPullRequestRequest(input *PostCommentForPullRequestInput) PostCommentForPullRequestRequest {
	op := &aws.Operation{
		Name:       opPostCommentForPullRequest,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PostCommentForPullRequestInput{}
	}

	req := c.newRequest(op, input, &PostCommentForPullRequestOutput{})
	return PostCommentForPullRequestRequest{Request: req, Input: input, Copy: c.PostCommentForPullRequestRequest}
}

// PostCommentForPullRequestRequest is the request type for the
// PostCommentForPullRequest API operation.
type PostCommentForPullRequestRequest struct {
	*aws.Request
	Input *PostCommentForPullRequestInput
	Copy  func(*PostCommentForPullRequestInput) PostCommentForPullRequestRequest
}

// Send marshals and sends the PostCommentForPullRequest API request.
func (r PostCommentForPullRequestRequest) Send(ctx context.Context) (*PostCommentForPullRequestResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PostCommentForPullRequestResponse{
		PostCommentForPullRequestOutput: r.Request.Data.(*PostCommentForPullRequestOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PostCommentForPullRequestResponse is the response type for the
// PostCommentForPullRequest API operation.
type PostCommentForPullRequestResponse struct {
	*PostCommentForPullRequestOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PostCommentForPullRequest request.
func (r *PostCommentForPullRequestResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}