// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/scylla_v2/models"
)

// FindConfigRequestTimeoutInMsReader is a Reader for the FindConfigRequestTimeoutInMs structure.
type FindConfigRequestTimeoutInMsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigRequestTimeoutInMsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindConfigRequestTimeoutInMsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFindConfigRequestTimeoutInMsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigRequestTimeoutInMsOK creates a FindConfigRequestTimeoutInMsOK with default headers values
func NewFindConfigRequestTimeoutInMsOK() *FindConfigRequestTimeoutInMsOK {
	return &FindConfigRequestTimeoutInMsOK{}
}

/*FindConfigRequestTimeoutInMsOK handles this case with default header values.

Config value
*/
type FindConfigRequestTimeoutInMsOK struct {
	Payload int64
}

func (o *FindConfigRequestTimeoutInMsOK) Error() string {
	return fmt.Sprintf("[GET /config/request_timeout_in_ms][%d] findConfigRequestTimeoutInMsOK  %+v", 200, o.Payload)
}

func (o *FindConfigRequestTimeoutInMsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigRequestTimeoutInMsDefault creates a FindConfigRequestTimeoutInMsDefault with default headers values
func NewFindConfigRequestTimeoutInMsDefault(code int) *FindConfigRequestTimeoutInMsDefault {
	return &FindConfigRequestTimeoutInMsDefault{
		_statusCode: code,
	}
}

/*FindConfigRequestTimeoutInMsDefault handles this case with default header values.

unexpected error
*/
type FindConfigRequestTimeoutInMsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config request timeout in ms default response
func (o *FindConfigRequestTimeoutInMsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigRequestTimeoutInMsDefault) Error() string {
	return fmt.Sprintf("[GET /config/request_timeout_in_ms][%d] find_config_request_timeout_in_ms default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigRequestTimeoutInMsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
