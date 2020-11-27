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

// FindConfigHintedHandoffEnabledReader is a Reader for the FindConfigHintedHandoffEnabled structure.
type FindConfigHintedHandoffEnabledReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigHintedHandoffEnabledReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigHintedHandoffEnabledOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigHintedHandoffEnabledDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigHintedHandoffEnabledOK creates a FindConfigHintedHandoffEnabledOK with default headers values
func NewFindConfigHintedHandoffEnabledOK() *FindConfigHintedHandoffEnabledOK {
	return &FindConfigHintedHandoffEnabledOK{}
}

/*FindConfigHintedHandoffEnabledOK handles this case with default header values.

Config value
*/
type FindConfigHintedHandoffEnabledOK struct {
	Payload string
}

func (o *FindConfigHintedHandoffEnabledOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigHintedHandoffEnabledOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigHintedHandoffEnabledDefault creates a FindConfigHintedHandoffEnabledDefault with default headers values
func NewFindConfigHintedHandoffEnabledDefault(code int) *FindConfigHintedHandoffEnabledDefault {
	return &FindConfigHintedHandoffEnabledDefault{
		_statusCode: code,
	}
}

/*FindConfigHintedHandoffEnabledDefault handles this case with default header values.

unexpected error
*/
type FindConfigHintedHandoffEnabledDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config hinted handoff enabled default response
func (o *FindConfigHintedHandoffEnabledDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigHintedHandoffEnabledDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigHintedHandoffEnabledDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigHintedHandoffEnabledDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
