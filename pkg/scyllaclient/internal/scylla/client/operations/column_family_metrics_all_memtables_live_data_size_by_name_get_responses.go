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

// ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetReader is a Reader for the ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGet structure.
type ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK creates a ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK with default headers values
func NewColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK() *ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK {
	return &ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK{}
}

/*ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK handles this case with default header values.

ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK column family metrics all memtables live data size by name get o k
*/
type ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetDefault creates a ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetDefault with default headers values
func NewColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetDefault(code int) *ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetDefault {
	return &ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics all memtables live data size by name get default response
func (o *ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
