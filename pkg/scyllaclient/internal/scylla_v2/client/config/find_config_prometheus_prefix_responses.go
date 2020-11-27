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

// FindConfigPrometheusPrefixReader is a Reader for the FindConfigPrometheusPrefix structure.
type FindConfigPrometheusPrefixReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigPrometheusPrefixReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigPrometheusPrefixOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigPrometheusPrefixDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigPrometheusPrefixOK creates a FindConfigPrometheusPrefixOK with default headers values
func NewFindConfigPrometheusPrefixOK() *FindConfigPrometheusPrefixOK {
	return &FindConfigPrometheusPrefixOK{}
}

/*FindConfigPrometheusPrefixOK handles this case with default header values.

Config value
*/
type FindConfigPrometheusPrefixOK struct {
	Payload string
}

func (o *FindConfigPrometheusPrefixOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigPrometheusPrefixOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigPrometheusPrefixDefault creates a FindConfigPrometheusPrefixDefault with default headers values
func NewFindConfigPrometheusPrefixDefault(code int) *FindConfigPrometheusPrefixDefault {
	return &FindConfigPrometheusPrefixDefault{
		_statusCode: code,
	}
}

/*FindConfigPrometheusPrefixDefault handles this case with default header values.

unexpected error
*/
type FindConfigPrometheusPrefixDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config prometheus prefix default response
func (o *FindConfigPrometheusPrefixDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigPrometheusPrefixDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigPrometheusPrefixDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigPrometheusPrefixDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
