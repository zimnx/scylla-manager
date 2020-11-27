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

// StorageServiceIsInitializedGetReader is a Reader for the StorageServiceIsInitializedGet structure.
type StorageServiceIsInitializedGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceIsInitializedGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceIsInitializedGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceIsInitializedGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceIsInitializedGetOK creates a StorageServiceIsInitializedGetOK with default headers values
func NewStorageServiceIsInitializedGetOK() *StorageServiceIsInitializedGetOK {
	return &StorageServiceIsInitializedGetOK{}
}

/*StorageServiceIsInitializedGetOK handles this case with default header values.

StorageServiceIsInitializedGetOK storage service is initialized get o k
*/
type StorageServiceIsInitializedGetOK struct {
	Payload bool
}

func (o *StorageServiceIsInitializedGetOK) GetPayload() bool {
	return o.Payload
}

func (o *StorageServiceIsInitializedGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceIsInitializedGetDefault creates a StorageServiceIsInitializedGetDefault with default headers values
func NewStorageServiceIsInitializedGetDefault(code int) *StorageServiceIsInitializedGetDefault {
	return &StorageServiceIsInitializedGetDefault{
		_statusCode: code,
	}
}

/*StorageServiceIsInitializedGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceIsInitializedGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service is initialized get default response
func (o *StorageServiceIsInitializedGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceIsInitializedGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceIsInitializedGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceIsInitializedGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
