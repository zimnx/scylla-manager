// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/scylla/models"
)

// LsaCompactPostReader is a Reader for the LsaCompactPost structure.
type LsaCompactPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LsaCompactPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLsaCompactPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLsaCompactPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLsaCompactPostOK creates a LsaCompactPostOK with default headers values
func NewLsaCompactPostOK() *LsaCompactPostOK {
	return &LsaCompactPostOK{}
}

/*LsaCompactPostOK handles this case with default header values.

LsaCompactPostOK lsa compact post o k
*/
type LsaCompactPostOK struct {
}

func (o *LsaCompactPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLsaCompactPostDefault creates a LsaCompactPostDefault with default headers values
func NewLsaCompactPostDefault(code int) *LsaCompactPostDefault {
	return &LsaCompactPostDefault{
		_statusCode: code,
	}
}

/*LsaCompactPostDefault handles this case with default header values.

internal server error
*/
type LsaCompactPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the lsa compact post default response
func (o *LsaCompactPostDefault) Code() int {
	return o._statusCode
}

func (o *LsaCompactPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *LsaCompactPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *LsaCompactPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
