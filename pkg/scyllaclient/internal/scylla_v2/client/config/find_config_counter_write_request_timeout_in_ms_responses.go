// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigCounterWriteRequestTimeoutInMsReader is a Reader for the FindConfigCounterWriteRequestTimeoutInMs structure.
type FindConfigCounterWriteRequestTimeoutInMsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigCounterWriteRequestTimeoutInMsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigCounterWriteRequestTimeoutInMsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigCounterWriteRequestTimeoutInMsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigCounterWriteRequestTimeoutInMsOK creates a FindConfigCounterWriteRequestTimeoutInMsOK with default headers values
func NewFindConfigCounterWriteRequestTimeoutInMsOK() *FindConfigCounterWriteRequestTimeoutInMsOK {
	return &FindConfigCounterWriteRequestTimeoutInMsOK{}
}

/*FindConfigCounterWriteRequestTimeoutInMsOK handles this case with default header values.

Config value
*/
type FindConfigCounterWriteRequestTimeoutInMsOK struct {
	Payload int64
}

func (o *FindConfigCounterWriteRequestTimeoutInMsOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigCounterWriteRequestTimeoutInMsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigCounterWriteRequestTimeoutInMsDefault creates a FindConfigCounterWriteRequestTimeoutInMsDefault with default headers values
func NewFindConfigCounterWriteRequestTimeoutInMsDefault(code int) *FindConfigCounterWriteRequestTimeoutInMsDefault {
	return &FindConfigCounterWriteRequestTimeoutInMsDefault{
		_statusCode: code,
	}
}

/*FindConfigCounterWriteRequestTimeoutInMsDefault handles this case with default header values.

unexpected error
*/
type FindConfigCounterWriteRequestTimeoutInMsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config counter write request timeout in ms default response
func (o *FindConfigCounterWriteRequestTimeoutInMsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigCounterWriteRequestTimeoutInMsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigCounterWriteRequestTimeoutInMsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigCounterWriteRequestTimeoutInMsDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
