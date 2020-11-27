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

// ColumnFamilyMetricsLiveDiskSpaceUsedGetReader is a Reader for the ColumnFamilyMetricsLiveDiskSpaceUsedGet structure.
type ColumnFamilyMetricsLiveDiskSpaceUsedGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsLiveDiskSpaceUsedGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsLiveDiskSpaceUsedGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsLiveDiskSpaceUsedGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsLiveDiskSpaceUsedGetOK creates a ColumnFamilyMetricsLiveDiskSpaceUsedGetOK with default headers values
func NewColumnFamilyMetricsLiveDiskSpaceUsedGetOK() *ColumnFamilyMetricsLiveDiskSpaceUsedGetOK {
	return &ColumnFamilyMetricsLiveDiskSpaceUsedGetOK{}
}

/*ColumnFamilyMetricsLiveDiskSpaceUsedGetOK handles this case with default header values.

ColumnFamilyMetricsLiveDiskSpaceUsedGetOK column family metrics live disk space used get o k
*/
type ColumnFamilyMetricsLiveDiskSpaceUsedGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsLiveDiskSpaceUsedGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsLiveDiskSpaceUsedGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsLiveDiskSpaceUsedGetDefault creates a ColumnFamilyMetricsLiveDiskSpaceUsedGetDefault with default headers values
func NewColumnFamilyMetricsLiveDiskSpaceUsedGetDefault(code int) *ColumnFamilyMetricsLiveDiskSpaceUsedGetDefault {
	return &ColumnFamilyMetricsLiveDiskSpaceUsedGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsLiveDiskSpaceUsedGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsLiveDiskSpaceUsedGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics live disk space used get default response
func (o *ColumnFamilyMetricsLiveDiskSpaceUsedGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsLiveDiskSpaceUsedGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsLiveDiskSpaceUsedGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsLiveDiskSpaceUsedGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
