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

// ColumnFamilyMetricsMemtableSwitchCountGetReader is a Reader for the ColumnFamilyMetricsMemtableSwitchCountGet structure.
type ColumnFamilyMetricsMemtableSwitchCountGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsMemtableSwitchCountGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsMemtableSwitchCountGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsMemtableSwitchCountGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsMemtableSwitchCountGetOK creates a ColumnFamilyMetricsMemtableSwitchCountGetOK with default headers values
func NewColumnFamilyMetricsMemtableSwitchCountGetOK() *ColumnFamilyMetricsMemtableSwitchCountGetOK {
	return &ColumnFamilyMetricsMemtableSwitchCountGetOK{}
}

/*ColumnFamilyMetricsMemtableSwitchCountGetOK handles this case with default header values.

ColumnFamilyMetricsMemtableSwitchCountGetOK column family metrics memtable switch count get o k
*/
type ColumnFamilyMetricsMemtableSwitchCountGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsMemtableSwitchCountGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsMemtableSwitchCountGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsMemtableSwitchCountGetDefault creates a ColumnFamilyMetricsMemtableSwitchCountGetDefault with default headers values
func NewColumnFamilyMetricsMemtableSwitchCountGetDefault(code int) *ColumnFamilyMetricsMemtableSwitchCountGetDefault {
	return &ColumnFamilyMetricsMemtableSwitchCountGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsMemtableSwitchCountGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsMemtableSwitchCountGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics memtable switch count get default response
func (o *ColumnFamilyMetricsMemtableSwitchCountGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsMemtableSwitchCountGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsMemtableSwitchCountGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsMemtableSwitchCountGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
