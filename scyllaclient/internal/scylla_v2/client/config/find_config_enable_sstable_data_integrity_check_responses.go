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

// FindConfigEnableSstableDataIntegrityCheckReader is a Reader for the FindConfigEnableSstableDataIntegrityCheck structure.
type FindConfigEnableSstableDataIntegrityCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigEnableSstableDataIntegrityCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindConfigEnableSstableDataIntegrityCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFindConfigEnableSstableDataIntegrityCheckDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigEnableSstableDataIntegrityCheckOK creates a FindConfigEnableSstableDataIntegrityCheckOK with default headers values
func NewFindConfigEnableSstableDataIntegrityCheckOK() *FindConfigEnableSstableDataIntegrityCheckOK {
	return &FindConfigEnableSstableDataIntegrityCheckOK{}
}

/*FindConfigEnableSstableDataIntegrityCheckOK handles this case with default header values.

Config value
*/
type FindConfigEnableSstableDataIntegrityCheckOK struct {
	Payload bool
}

func (o *FindConfigEnableSstableDataIntegrityCheckOK) Error() string {
	return fmt.Sprintf("[GET /config/enable_sstable_data_integrity_check][%d] findConfigEnableSstableDataIntegrityCheckOK  %+v", 200, o.Payload)
}

func (o *FindConfigEnableSstableDataIntegrityCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigEnableSstableDataIntegrityCheckDefault creates a FindConfigEnableSstableDataIntegrityCheckDefault with default headers values
func NewFindConfigEnableSstableDataIntegrityCheckDefault(code int) *FindConfigEnableSstableDataIntegrityCheckDefault {
	return &FindConfigEnableSstableDataIntegrityCheckDefault{
		_statusCode: code,
	}
}

/*FindConfigEnableSstableDataIntegrityCheckDefault handles this case with default header values.

unexpected error
*/
type FindConfigEnableSstableDataIntegrityCheckDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config enable sstable data integrity check default response
func (o *FindConfigEnableSstableDataIntegrityCheckDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigEnableSstableDataIntegrityCheckDefault) Error() string {
	return fmt.Sprintf("[GET /config/enable_sstable_data_integrity_check][%d] find_config_enable_sstable_data_integrity_check default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigEnableSstableDataIntegrityCheckDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
