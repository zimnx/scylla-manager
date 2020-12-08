// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/scylla/models"
)

// StorageServiceRepairStatusReader is a Reader for the StorageServiceRepairStatus structure.
type StorageServiceRepairStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceRepairStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceRepairStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceRepairStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceRepairStatusOK creates a StorageServiceRepairStatusOK with default headers values
func NewStorageServiceRepairStatusOK() *StorageServiceRepairStatusOK {
	return &StorageServiceRepairStatusOK{}
}

/*StorageServiceRepairStatusOK handles this case with default header values.

Repair status value
*/
type StorageServiceRepairStatusOK struct {
	Payload models.RepairAsyncStatusResponse
}

func (o *StorageServiceRepairStatusOK) GetPayload() models.RepairAsyncStatusResponse {
	return o.Payload
}

func (o *StorageServiceRepairStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceRepairStatusDefault creates a StorageServiceRepairStatusDefault with default headers values
func NewStorageServiceRepairStatusDefault(code int) *StorageServiceRepairStatusDefault {
	return &StorageServiceRepairStatusDefault{
		_statusCode: code,
	}
}

/*StorageServiceRepairStatusDefault handles this case with default header values.

internal server error
*/
type StorageServiceRepairStatusDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service repair status default response
func (o *StorageServiceRepairStatusDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceRepairStatusDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceRepairStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceRepairStatusDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}