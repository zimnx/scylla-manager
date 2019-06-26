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

// FindConfigCompactionStaticSharesReader is a Reader for the FindConfigCompactionStaticShares structure.
type FindConfigCompactionStaticSharesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigCompactionStaticSharesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindConfigCompactionStaticSharesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFindConfigCompactionStaticSharesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigCompactionStaticSharesOK creates a FindConfigCompactionStaticSharesOK with default headers values
func NewFindConfigCompactionStaticSharesOK() *FindConfigCompactionStaticSharesOK {
	return &FindConfigCompactionStaticSharesOK{}
}

/*FindConfigCompactionStaticSharesOK handles this case with default header values.

Config value
*/
type FindConfigCompactionStaticSharesOK struct {
	Payload float64
}

func (o *FindConfigCompactionStaticSharesOK) Error() string {
	return fmt.Sprintf("[GET /config/compaction_static_shares][%d] findConfigCompactionStaticSharesOK  %+v", 200, o.Payload)
}

func (o *FindConfigCompactionStaticSharesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigCompactionStaticSharesDefault creates a FindConfigCompactionStaticSharesDefault with default headers values
func NewFindConfigCompactionStaticSharesDefault(code int) *FindConfigCompactionStaticSharesDefault {
	return &FindConfigCompactionStaticSharesDefault{
		_statusCode: code,
	}
}

/*FindConfigCompactionStaticSharesDefault handles this case with default header values.

unexpected error
*/
type FindConfigCompactionStaticSharesDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config compaction static shares default response
func (o *FindConfigCompactionStaticSharesDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigCompactionStaticSharesDefault) Error() string {
	return fmt.Sprintf("[GET /config/compaction_static_shares][%d] find_config_compaction_static_shares default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigCompactionStaticSharesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
