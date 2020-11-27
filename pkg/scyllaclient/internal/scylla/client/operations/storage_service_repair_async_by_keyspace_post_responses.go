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

// StorageServiceRepairAsyncByKeyspacePostReader is a Reader for the StorageServiceRepairAsyncByKeyspacePost structure.
type StorageServiceRepairAsyncByKeyspacePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceRepairAsyncByKeyspacePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceRepairAsyncByKeyspacePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceRepairAsyncByKeyspacePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceRepairAsyncByKeyspacePostOK creates a StorageServiceRepairAsyncByKeyspacePostOK with default headers values
func NewStorageServiceRepairAsyncByKeyspacePostOK() *StorageServiceRepairAsyncByKeyspacePostOK {
	return &StorageServiceRepairAsyncByKeyspacePostOK{}
}

/*StorageServiceRepairAsyncByKeyspacePostOK handles this case with default header values.

StorageServiceRepairAsyncByKeyspacePostOK storage service repair async by keyspace post o k
*/
type StorageServiceRepairAsyncByKeyspacePostOK struct {
	Payload int32
}

func (o *StorageServiceRepairAsyncByKeyspacePostOK) GetPayload() int32 {
	return o.Payload
}

func (o *StorageServiceRepairAsyncByKeyspacePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceRepairAsyncByKeyspacePostDefault creates a StorageServiceRepairAsyncByKeyspacePostDefault with default headers values
func NewStorageServiceRepairAsyncByKeyspacePostDefault(code int) *StorageServiceRepairAsyncByKeyspacePostDefault {
	return &StorageServiceRepairAsyncByKeyspacePostDefault{
		_statusCode: code,
	}
}

/*StorageServiceRepairAsyncByKeyspacePostDefault handles this case with default header values.

internal server error
*/
type StorageServiceRepairAsyncByKeyspacePostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service repair async by keyspace post default response
func (o *StorageServiceRepairAsyncByKeyspacePostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceRepairAsyncByKeyspacePostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceRepairAsyncByKeyspacePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceRepairAsyncByKeyspacePostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
