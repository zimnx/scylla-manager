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

// FindConfigSavedCachesDirectoryReader is a Reader for the FindConfigSavedCachesDirectory structure.
type FindConfigSavedCachesDirectoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigSavedCachesDirectoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindConfigSavedCachesDirectoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFindConfigSavedCachesDirectoryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigSavedCachesDirectoryOK creates a FindConfigSavedCachesDirectoryOK with default headers values
func NewFindConfigSavedCachesDirectoryOK() *FindConfigSavedCachesDirectoryOK {
	return &FindConfigSavedCachesDirectoryOK{}
}

/*FindConfigSavedCachesDirectoryOK handles this case with default header values.

Config value
*/
type FindConfigSavedCachesDirectoryOK struct {
	Payload string
}

func (o *FindConfigSavedCachesDirectoryOK) Error() string {
	return fmt.Sprintf("[GET /config/saved_caches_directory][%d] findConfigSavedCachesDirectoryOK  %+v", 200, o.Payload)
}

func (o *FindConfigSavedCachesDirectoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigSavedCachesDirectoryDefault creates a FindConfigSavedCachesDirectoryDefault with default headers values
func NewFindConfigSavedCachesDirectoryDefault(code int) *FindConfigSavedCachesDirectoryDefault {
	return &FindConfigSavedCachesDirectoryDefault{
		_statusCode: code,
	}
}

/*FindConfigSavedCachesDirectoryDefault handles this case with default header values.

unexpected error
*/
type FindConfigSavedCachesDirectoryDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config saved caches directory default response
func (o *FindConfigSavedCachesDirectoryDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigSavedCachesDirectoryDefault) Error() string {
	return fmt.Sprintf("[GET /config/saved_caches_directory][%d] find_config_saved_caches_directory default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigSavedCachesDirectoryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
