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

// FindConfigMemtableFlushQueueSizeReader is a Reader for the FindConfigMemtableFlushQueueSize structure.
type FindConfigMemtableFlushQueueSizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigMemtableFlushQueueSizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigMemtableFlushQueueSizeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigMemtableFlushQueueSizeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigMemtableFlushQueueSizeOK creates a FindConfigMemtableFlushQueueSizeOK with default headers values
func NewFindConfigMemtableFlushQueueSizeOK() *FindConfigMemtableFlushQueueSizeOK {
	return &FindConfigMemtableFlushQueueSizeOK{}
}

/*FindConfigMemtableFlushQueueSizeOK handles this case with default header values.

Config value
*/
type FindConfigMemtableFlushQueueSizeOK struct {
	Payload int64
}

func (o *FindConfigMemtableFlushQueueSizeOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigMemtableFlushQueueSizeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigMemtableFlushQueueSizeDefault creates a FindConfigMemtableFlushQueueSizeDefault with default headers values
func NewFindConfigMemtableFlushQueueSizeDefault(code int) *FindConfigMemtableFlushQueueSizeDefault {
	return &FindConfigMemtableFlushQueueSizeDefault{
		_statusCode: code,
	}
}

/*FindConfigMemtableFlushQueueSizeDefault handles this case with default header values.

unexpected error
*/
type FindConfigMemtableFlushQueueSizeDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config memtable flush queue size default response
func (o *FindConfigMemtableFlushQueueSizeDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigMemtableFlushQueueSizeDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigMemtableFlushQueueSizeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigMemtableFlushQueueSizeDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
