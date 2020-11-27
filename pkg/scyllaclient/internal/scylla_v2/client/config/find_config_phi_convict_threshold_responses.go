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

// FindConfigPhiConvictThresholdReader is a Reader for the FindConfigPhiConvictThreshold structure.
type FindConfigPhiConvictThresholdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigPhiConvictThresholdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigPhiConvictThresholdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigPhiConvictThresholdDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigPhiConvictThresholdOK creates a FindConfigPhiConvictThresholdOK with default headers values
func NewFindConfigPhiConvictThresholdOK() *FindConfigPhiConvictThresholdOK {
	return &FindConfigPhiConvictThresholdOK{}
}

/*FindConfigPhiConvictThresholdOK handles this case with default header values.

Config value
*/
type FindConfigPhiConvictThresholdOK struct {
	Payload int64
}

func (o *FindConfigPhiConvictThresholdOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigPhiConvictThresholdOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigPhiConvictThresholdDefault creates a FindConfigPhiConvictThresholdDefault with default headers values
func NewFindConfigPhiConvictThresholdDefault(code int) *FindConfigPhiConvictThresholdDefault {
	return &FindConfigPhiConvictThresholdDefault{
		_statusCode: code,
	}
}

/*FindConfigPhiConvictThresholdDefault handles this case with default header values.

unexpected error
*/
type FindConfigPhiConvictThresholdDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config phi convict threshold default response
func (o *FindConfigPhiConvictThresholdDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigPhiConvictThresholdDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigPhiConvictThresholdDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigPhiConvictThresholdDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
