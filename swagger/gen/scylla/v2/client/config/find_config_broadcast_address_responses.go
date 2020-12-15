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

	"github.com/scylladb/scylla-manager/swagger/gen/scylla/v2/models"
)

// FindConfigBroadcastAddressReader is a Reader for the FindConfigBroadcastAddress structure.
type FindConfigBroadcastAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigBroadcastAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigBroadcastAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigBroadcastAddressDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigBroadcastAddressOK creates a FindConfigBroadcastAddressOK with default headers values
func NewFindConfigBroadcastAddressOK() *FindConfigBroadcastAddressOK {
	return &FindConfigBroadcastAddressOK{}
}

/*FindConfigBroadcastAddressOK handles this case with default header values.

Config value
*/
type FindConfigBroadcastAddressOK struct {
	Payload string
}

func (o *FindConfigBroadcastAddressOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigBroadcastAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigBroadcastAddressDefault creates a FindConfigBroadcastAddressDefault with default headers values
func NewFindConfigBroadcastAddressDefault(code int) *FindConfigBroadcastAddressDefault {
	return &FindConfigBroadcastAddressDefault{
		_statusCode: code,
	}
}

/*FindConfigBroadcastAddressDefault handles this case with default header values.

unexpected error
*/
type FindConfigBroadcastAddressDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config broadcast address default response
func (o *FindConfigBroadcastAddressDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigBroadcastAddressDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigBroadcastAddressDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigBroadcastAddressDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}