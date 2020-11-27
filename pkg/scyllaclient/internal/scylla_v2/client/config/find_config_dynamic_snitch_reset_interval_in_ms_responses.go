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

// FindConfigDynamicSnitchResetIntervalInMsReader is a Reader for the FindConfigDynamicSnitchResetIntervalInMs structure.
type FindConfigDynamicSnitchResetIntervalInMsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigDynamicSnitchResetIntervalInMsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigDynamicSnitchResetIntervalInMsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigDynamicSnitchResetIntervalInMsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigDynamicSnitchResetIntervalInMsOK creates a FindConfigDynamicSnitchResetIntervalInMsOK with default headers values
func NewFindConfigDynamicSnitchResetIntervalInMsOK() *FindConfigDynamicSnitchResetIntervalInMsOK {
	return &FindConfigDynamicSnitchResetIntervalInMsOK{}
}

/*FindConfigDynamicSnitchResetIntervalInMsOK handles this case with default header values.

Config value
*/
type FindConfigDynamicSnitchResetIntervalInMsOK struct {
	Payload int64
}

func (o *FindConfigDynamicSnitchResetIntervalInMsOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigDynamicSnitchResetIntervalInMsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigDynamicSnitchResetIntervalInMsDefault creates a FindConfigDynamicSnitchResetIntervalInMsDefault with default headers values
func NewFindConfigDynamicSnitchResetIntervalInMsDefault(code int) *FindConfigDynamicSnitchResetIntervalInMsDefault {
	return &FindConfigDynamicSnitchResetIntervalInMsDefault{
		_statusCode: code,
	}
}

/*FindConfigDynamicSnitchResetIntervalInMsDefault handles this case with default header values.

unexpected error
*/
type FindConfigDynamicSnitchResetIntervalInMsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config dynamic snitch reset interval in ms default response
func (o *FindConfigDynamicSnitchResetIntervalInMsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigDynamicSnitchResetIntervalInMsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigDynamicSnitchResetIntervalInMsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigDynamicSnitchResetIntervalInMsDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
