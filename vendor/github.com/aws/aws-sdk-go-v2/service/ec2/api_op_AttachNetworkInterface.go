// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for AttachNetworkInterface.
type AttachNetworkInterfaceInput struct {
	_ struct{} `type:"structure"`

	// The index of the device for the network interface attachment.
	//
	// DeviceIndex is a required field
	DeviceIndex *int64 `locationName:"deviceIndex" type:"integer" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the instance.
	//
	// InstanceId is a required field
	InstanceId *string `locationName:"instanceId" type:"string" required:"true"`

	// The ID of the network interface.
	//
	// NetworkInterfaceId is a required field
	NetworkInterfaceId *string `locationName:"networkInterfaceId" type:"string" required:"true"`
}

// String returns the string representation
func (s AttachNetworkInterfaceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachNetworkInterfaceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachNetworkInterfaceInput"}

	if s.DeviceIndex == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceIndex"))
	}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if s.NetworkInterfaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkInterfaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of AttachNetworkInterface.
type AttachNetworkInterfaceOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the network interface attachment.
	AttachmentId *string `locationName:"attachmentId" type:"string"`
}

// String returns the string representation
func (s AttachNetworkInterfaceOutput) String() string {
	return awsutil.Prettify(s)
}

const opAttachNetworkInterface = "AttachNetworkInterface"

// AttachNetworkInterfaceRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Attaches a network interface to an instance.
//
//    // Example sending a request using AttachNetworkInterfaceRequest.
//    req := client.AttachNetworkInterfaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AttachNetworkInterface
func (c *Client) AttachNetworkInterfaceRequest(input *AttachNetworkInterfaceInput) AttachNetworkInterfaceRequest {
	op := &aws.Operation{
		Name:       opAttachNetworkInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachNetworkInterfaceInput{}
	}

	req := c.newRequest(op, input, &AttachNetworkInterfaceOutput{})

	return AttachNetworkInterfaceRequest{Request: req, Input: input, Copy: c.AttachNetworkInterfaceRequest}
}

// AttachNetworkInterfaceRequest is the request type for the
// AttachNetworkInterface API operation.
type AttachNetworkInterfaceRequest struct {
	*aws.Request
	Input *AttachNetworkInterfaceInput
	Copy  func(*AttachNetworkInterfaceInput) AttachNetworkInterfaceRequest
}

// Send marshals and sends the AttachNetworkInterface API request.
func (r AttachNetworkInterfaceRequest) Send(ctx context.Context) (*AttachNetworkInterfaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachNetworkInterfaceResponse{
		AttachNetworkInterfaceOutput: r.Request.Data.(*AttachNetworkInterfaceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachNetworkInterfaceResponse is the response type for the
// AttachNetworkInterface API operation.
type AttachNetworkInterfaceResponse struct {
	*AttachNetworkInterfaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachNetworkInterface request.
func (r *AttachNetworkInterfaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
