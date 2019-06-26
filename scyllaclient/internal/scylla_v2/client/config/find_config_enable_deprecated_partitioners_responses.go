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

// FindConfigEnableDeprecatedPartitionersReader is a Reader for the FindConfigEnableDeprecatedPartitioners structure.
type FindConfigEnableDeprecatedPartitionersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigEnableDeprecatedPartitionersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindConfigEnableDeprecatedPartitionersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFindConfigEnableDeprecatedPartitionersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigEnableDeprecatedPartitionersOK creates a FindConfigEnableDeprecatedPartitionersOK with default headers values
func NewFindConfigEnableDeprecatedPartitionersOK() *FindConfigEnableDeprecatedPartitionersOK {
	return &FindConfigEnableDeprecatedPartitionersOK{}
}

/*FindConfigEnableDeprecatedPartitionersOK handles this case with default header values.

Config value
*/
type FindConfigEnableDeprecatedPartitionersOK struct {
	Payload bool
}

func (o *FindConfigEnableDeprecatedPartitionersOK) Error() string {
	return fmt.Sprintf("[GET /config/enable_deprecated_partitioners][%d] findConfigEnableDeprecatedPartitionersOK  %+v", 200, o.Payload)
}

func (o *FindConfigEnableDeprecatedPartitionersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigEnableDeprecatedPartitionersDefault creates a FindConfigEnableDeprecatedPartitionersDefault with default headers values
func NewFindConfigEnableDeprecatedPartitionersDefault(code int) *FindConfigEnableDeprecatedPartitionersDefault {
	return &FindConfigEnableDeprecatedPartitionersDefault{
		_statusCode: code,
	}
}

/*FindConfigEnableDeprecatedPartitionersDefault handles this case with default header values.

unexpected error
*/
type FindConfigEnableDeprecatedPartitionersDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config enable deprecated partitioners default response
func (o *FindConfigEnableDeprecatedPartitionersDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigEnableDeprecatedPartitionersDefault) Error() string {
	return fmt.Sprintf("[GET /config/enable_deprecated_partitioners][%d] find_config_enable_deprecated_partitioners default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigEnableDeprecatedPartitionersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
