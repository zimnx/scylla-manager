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

	"github.com/scylladb/scylla-manager/swagger/gen/scylla/v1/models"
)

// StorageServiceLoggingLevelPostReader is a Reader for the StorageServiceLoggingLevelPost structure.
type StorageServiceLoggingLevelPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceLoggingLevelPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceLoggingLevelPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceLoggingLevelPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceLoggingLevelPostOK creates a StorageServiceLoggingLevelPostOK with default headers values
func NewStorageServiceLoggingLevelPostOK() *StorageServiceLoggingLevelPostOK {
	return &StorageServiceLoggingLevelPostOK{}
}

/*StorageServiceLoggingLevelPostOK handles this case with default header values.

StorageServiceLoggingLevelPostOK storage service logging level post o k
*/
type StorageServiceLoggingLevelPostOK struct {
}

func (o *StorageServiceLoggingLevelPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceLoggingLevelPostDefault creates a StorageServiceLoggingLevelPostDefault with default headers values
func NewStorageServiceLoggingLevelPostDefault(code int) *StorageServiceLoggingLevelPostDefault {
	return &StorageServiceLoggingLevelPostDefault{
		_statusCode: code,
	}
}

/*StorageServiceLoggingLevelPostDefault handles this case with default header values.

internal server error
*/
type StorageServiceLoggingLevelPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service logging level post default response
func (o *StorageServiceLoggingLevelPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceLoggingLevelPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceLoggingLevelPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceLoggingLevelPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
