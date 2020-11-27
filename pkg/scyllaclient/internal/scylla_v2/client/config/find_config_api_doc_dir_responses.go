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

// FindConfigAPIDocDirReader is a Reader for the FindConfigAPIDocDir structure.
type FindConfigAPIDocDirReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigAPIDocDirReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigAPIDocDirOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigAPIDocDirDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigAPIDocDirOK creates a FindConfigAPIDocDirOK with default headers values
func NewFindConfigAPIDocDirOK() *FindConfigAPIDocDirOK {
	return &FindConfigAPIDocDirOK{}
}

/*FindConfigAPIDocDirOK handles this case with default header values.

Config value
*/
type FindConfigAPIDocDirOK struct {
	Payload string
}

func (o *FindConfigAPIDocDirOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigAPIDocDirOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigAPIDocDirDefault creates a FindConfigAPIDocDirDefault with default headers values
func NewFindConfigAPIDocDirDefault(code int) *FindConfigAPIDocDirDefault {
	return &FindConfigAPIDocDirDefault{
		_statusCode: code,
	}
}

/*FindConfigAPIDocDirDefault handles this case with default header values.

unexpected error
*/
type FindConfigAPIDocDirDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config api doc dir default response
func (o *FindConfigAPIDocDirDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigAPIDocDirDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigAPIDocDirDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigAPIDocDirDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
