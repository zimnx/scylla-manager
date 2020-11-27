// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/agent/models"
)

// OperationsAboutReader is a Reader for the OperationsAbout structure.
type OperationsAboutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperationsAboutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOperationsAboutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOperationsAboutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOperationsAboutOK creates a OperationsAboutOK with default headers values
func NewOperationsAboutOK() *OperationsAboutOK {
	return &OperationsAboutOK{}
}

/*OperationsAboutOK handles this case with default header values.

File system details
*/
type OperationsAboutOK struct {
	Payload *models.FileSystemDetails
	JobID   int64
}

func (o *OperationsAboutOK) GetPayload() *models.FileSystemDetails {
	return o.Payload
}

func (o *OperationsAboutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FileSystemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

// NewOperationsAboutDefault creates a OperationsAboutDefault with default headers values
func NewOperationsAboutDefault(code int) *OperationsAboutDefault {
	return &OperationsAboutDefault{
		_statusCode: code,
	}
}

/*OperationsAboutDefault handles this case with default header values.

Server error
*/
type OperationsAboutDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
	JobID   int64
}

// Code gets the status code for the operations about default response
func (o *OperationsAboutDefault) Code() int {
	return o._statusCode
}

func (o *OperationsAboutDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *OperationsAboutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

func (o *OperationsAboutDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
