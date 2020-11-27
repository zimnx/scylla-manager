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

// StorageServicePartitionerNameGetReader is a Reader for the StorageServicePartitionerNameGet structure.
type StorageServicePartitionerNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServicePartitionerNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServicePartitionerNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServicePartitionerNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServicePartitionerNameGetOK creates a StorageServicePartitionerNameGetOK with default headers values
func NewStorageServicePartitionerNameGetOK() *StorageServicePartitionerNameGetOK {
	return &StorageServicePartitionerNameGetOK{}
}

/*StorageServicePartitionerNameGetOK handles this case with default header values.

StorageServicePartitionerNameGetOK storage service partitioner name get o k
*/
type StorageServicePartitionerNameGetOK struct {
	Payload string
}

func (o *StorageServicePartitionerNameGetOK) GetPayload() string {
	return o.Payload
}

func (o *StorageServicePartitionerNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServicePartitionerNameGetDefault creates a StorageServicePartitionerNameGetDefault with default headers values
func NewStorageServicePartitionerNameGetDefault(code int) *StorageServicePartitionerNameGetDefault {
	return &StorageServicePartitionerNameGetDefault{
		_statusCode: code,
	}
}

/*StorageServicePartitionerNameGetDefault handles this case with default header values.

internal server error
*/
type StorageServicePartitionerNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service partitioner name get default response
func (o *StorageServicePartitionerNameGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServicePartitionerNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServicePartitionerNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServicePartitionerNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
